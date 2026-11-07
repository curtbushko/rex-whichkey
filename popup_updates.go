package whichkey

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"

	"golang.org/x/sys/unix"

	"temp/rex"
)

// Exactly two form kinds, each published once per invocation. Slot names are
// literals, never derived from child-supplied prefixes or resource identifiers.
const (
	renameStateReady       = "ready"
	renameStateUnavailable = "unavailable"
)

var renameDemandFiles = [2]string{"rename-tab-demand.json", "rename-session-demand.json"}
var renameUpdateFiles = [2]string{"rename-tab-update.json", "rename-session-update.json"}

// Literal transport inventory shared by creation and exact cleanup.
func popupUpdateFiles() []string {
	names := make([]string, 0, 12)
	names = append(names, renameDemandFiles[:]...)
	names = append(names, renameUpdateFiles[:]...)
	names = append(names, providerDemandFiles[:]...)
	return append(names, providerUpdateFiles[:]...)
}

type renameDemand struct {
	Invocation string `json:"invocation"`
	Kind       string `json:"kind"`
	Load       int    `json:"load"`
	Keys       string `json:"keys"`
}

// The final binding omits demand keys: selected full keys already travel in
// the result. This keeps even near-limit configured rename keys within its bound.
type renameBinding struct {
	Invocation string `json:"invocation"`
	Kind       string `json:"kind"`
	Load       int    `json:"load"`
}

func (d renameDemand) binding() renameBinding {
	return renameBinding{Invocation: d.Invocation, Kind: d.Kind, Load: d.Load}
}

type renameUpdate struct {
	Demand  renameDemand `json:"demand"`
	State   string       `json:"state"`
	Initial string       `json:"initial"`
	Problem string       `json:"problem"`
}

func renameKind(mapping Mapping) int {
	if inlineRenameMapping(mapping) {
		switch mapping.Action {
		case menuTabRename:
			return 0
		case menuSessionRename:
			return 1
		}
	}
	return -1
}

func renameKindName(kind int) string {
	if kind == 0 {
		return "tab"
	}
	return "session"
}

func validateRenameDemand(demand renameDemand, invocation string, kind int, config Config) error {
	if demand.Invocation != invocation || demand.Kind != renameKindName(kind) || demand.Load != 1 {
		return errors.New("whichkey: stale or unknown rename demand")
	}
	canonical, err := NormalizeKeys(demand.Keys)
	if err != nil || canonical != demand.Keys {
		return errors.New("whichkey: noncanonical rename demand")
	}
	for _, mapping := range config.Spec {
		if mapping.Keys == demand.Keys && mapping.Enabled() && renameKind(mapping) == kind {
			return nil
		}
	}
	return errors.New("whichkey: rename demand is not a retained form")
}

// transferRenameSlot never waits for a peer/lock. Empty/contended slots are
// retried by the owner's normal poll. A writer publishes once while holding the
// exclusive lock; a reader sees either empty or the whole bounded JSON value.
func transferRenameSlot(root *os.Root, name string, write []byte) (data []byte, complete bool, outcome error) {
	flags, lock := os.O_RDONLY, unix.LOCK_SH|unix.LOCK_NB
	if write != nil {
		flags, lock = os.O_RDWR, unix.LOCK_EX|unix.LOCK_NB
		if len(write) == 0 || len(write) > popupMessageLimit {
			return nil, false, errRuntimeRelay
		}
	}
	file, err := root.OpenFile(name, flags|unix.O_NONBLOCK|unix.O_NOFOLLOW, 0)
	if err != nil {
		return nil, false, errors.Join(errRuntimeRelay, err)
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	if slotErr := checkPopupSlot(file, false); slotErr != nil {
		return nil, false, slotErr
	}
	if lockErr := unix.Flock(int(file.Fd()), lock); lockErr != nil {
		if errors.Is(lockErr, unix.EWOULDBLOCK) || errors.Is(lockErr, unix.EAGAIN) {
			return nil, false, nil
		}
		return nil, false, errors.Join(errRuntimeRelay, lockErr)
	}
	info, err := file.Stat()
	if err != nil || info.Size() > popupMessageLimit {
		return nil, false, errors.Join(errRuntimeRelay, err)
	}
	if write != nil {
		return publishRenameSlot(file, info.Size(), write)
	}
	if info.Size() == 0 {
		return nil, false, nil
	}
	data, err = io.ReadAll(io.LimitReader(file, popupMessageLimit+1))
	if err != nil || len(data) > popupMessageLimit {
		return nil, false, errors.Join(errRuntimeRelay, err)
	}
	return data, true, nil
}

func publishRenameSlot(file *os.File, size int64, data []byte) ([]byte, bool, error) {
	if size != 0 {
		return nil, false, errors.New("whichkey: duplicate metadata publication")
	}
	n, err := file.Write(data)
	if err != nil || n != len(data) {
		return nil, false, errors.Join(errRuntimeRelay, err, io.ErrShortWrite)
	}
	return nil, true, nil
}

type renameLabels struct {
	initial [2]string
	problem [2]string
	err     error
}

// This worker receives no mutable config/menu/targets. It returns value-only
// presentation from an immutable view, never executes an effect or publishes UI.
func loadRenameLabels(ctx context.Context, client *rex.Client, request popupRequest) renameLabels {
	view := request.snapshot
	if view == nil {
		read, done := context.WithTimeout(ctx, 3*time.Second)
		snapshot, err := client.SessionSnapshot(read, request.SessionID)
		done()
		if err != nil {
			return renameLabels{err: errors.New("whichkey: captured rename context unavailable")}
		}
		view = &snapshot
	}
	if err := ctx.Err(); err != nil {
		return renameLabels{err: err}
	}
	block, err := view.Block(request.BlockID)
	if err != nil || view.Session().ID != request.SessionID || request.placement == nil ||
		block.WindowID != request.placement.WindowID || block.LayerID != request.placement.LayerID {
		return renameLabels{err: errors.New("whichkey: captured rename placement unavailable")}
	}
	window, err := view.Window(block.WindowID)
	if err != nil {
		return renameLabels{err: err}
	}
	labels := renameLabels{initial: [2]string{window.Label, view.Session().Label}}
	for kind, value := range labels.initial {
		if len(value) > renameTextLimit || validateText(value) != nil {
			labels.initial[kind], labels.problem[kind] = "", "Enter a new name"
		}
	}
	return labels
}

// Only run's goroutine changes these fields. The lifecycle reads retained state
// only after stop joins it; the worker's sole publication is its value channel.
type popupRenameUpdates struct {
	cancel            context.CancelFunc
	stopped           chan struct{}
	outcome           error
	invocation        string
	retained          [2]*renameUpdate
	published         [2]bool
	config            Config
	targets           map[string]providerTarget
	providerDemands   [4]*renameDemand
	providerRetained  [4]*providerUpdate
	providerPublished [4]bool
}

func startPopupRenameUpdates(parent context.Context, cancelLifecycle context.CancelCauseFunc, directory string,
	request popupRequest, client *rex.Client, config Config,
) (*popupRenameUpdates, error) {
	root, err := openPopupDirectory(directory)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(parent)
	initial, err := providerConfig(config, [4]*providerUpdate{})
	if err != nil {
		cancel()
		return nil, errors.Join(err, root.Close())
	}
	updates := &popupRenameUpdates{cancel: cancel, stopped: make(chan struct{}), invocation: request.invocationID,
		config: initial, targets: make(map[string]providerTarget)}
	go func() {
		defer close(updates.stopped)
		updates.outcome = errors.Join(updates.run(ctx, root, request, client, config), root.Close())
		if updates.outcome != nil {
			cancelLifecycle(updates.outcome)
		}
	}()
	return updates, nil
}

func (u *popupRenameUpdates) stop() error {
	u.cancel()
	<-u.stopped
	return u.outcome
}

func (u *popupRenameUpdates) run(ctx context.Context, root *os.Root, request popupRequest, client *rex.Client, config Config) error {
	jobs := newMetadataJobs(ctx, request, client)
	defer jobs.stop()
	var demands [2]*renameDemand
	var labels *renameLabels
	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		case result := <-jobs.results:
			jobs.active--
			if ctx.Err() != nil {
				return nil
			}
			if result.kind == metadataRenameKind {
				labels = &result.labels
			} else if err := u.retainProviderResult(config, result); err != nil {
				return err
			}
		case <-ticker.C:
		}
		newDemand, err := u.readDemands(ctx, root, config, &demands)
		if err != nil {
			return err
		}
		if newDemand {
			jobs.wanted[metadataRenameKind] = true
		}
		if err := u.readProviderDemands(ctx, root, config); err != nil {
			return err
		}
		for kind, demand := range u.providerDemands {
			jobs.wanted[kind] = demand != nil
		}
		jobs.start()
		if labels != nil {
			if err := u.publish(ctx, root, demands, *labels); err != nil {
				return err
			}
		}
		if err := u.publishProviders(ctx, root); err != nil {
			return err
		}
	}
}

func (u *popupRenameUpdates) retainProviderResult(config Config, result metadataResult) error {
	demand := u.providerDemands[result.kind]
	if demand == nil {
		return errRuntimeStartup
	}
	update, targets, err := buildProviderUpdate(config, *demand, result)
	if err != nil {
		return err
	}
	ready := u.providerRetained
	ready[result.kind] = &update
	next, err := providerConfig(config, ready)
	if err != nil {
		return err
	}
	// The controller is the sole owner until stop joins it. Exact mappings,
	// args and placement are retained BEFORE the first ready publication byte.
	u.config, u.providerRetained = next, ready
	for key, target := range targets {
		u.targets[key] = target
	}
	return nil
}

func (u *popupRenameUpdates) publishProviders(ctx context.Context, root *os.Root) error {
	for kind, update := range u.providerRetained {
		if ctx.Err() != nil {
			return nil
		}
		if update == nil || u.providerPublished[kind] {
			continue
		}
		data, err := json.Marshal(update)
		if err != nil {
			return err
		}
		_, u.providerPublished[kind], err = transferRenameSlot(root, providerUpdateFiles[kind], data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *popupRenameUpdates) readProviderDemands(ctx context.Context, root *os.Root, config Config) error {
	for kind, demand := range u.providerDemands {
		if ctx.Err() != nil || demand != nil {
			continue
		}
		data, ready, err := transferRenameSlot(root, providerDemandFiles[kind], nil)
		if err != nil {
			return err
		}
		if !ready {
			continue
		}
		var next renameDemand
		if err := decodeRuntime(data, &next, popupMessageLimit); err != nil {
			return err
		}
		if err := validateProviderDemand(next, u.invocation, kind, config); err != nil {
			return err
		}
		u.providerDemands[kind] = &next
	}
	return nil
}

func (u *popupRenameUpdates) readDemands(ctx context.Context, root *os.Root, config Config, demands *[2]*renameDemand) (bool, error) {
	found := false
	for kind, demand := range demands {
		if ctx.Err() != nil || demand != nil {
			continue
		}
		data, ready, err := transferRenameSlot(root, renameDemandFiles[kind], nil)
		if err != nil {
			return false, err
		}
		if !ready {
			continue
		}
		var next renameDemand
		if err := decodeRuntime(data, &next, popupMessageLimit); err != nil {
			return false, err
		}
		if err := validateRenameDemand(next, u.invocation, kind, config); err != nil {
			return false, err
		}
		demands[kind], found = &next, true
	}
	return found, nil
}

func (u *popupRenameUpdates) publish(ctx context.Context, root *os.Root, demands [2]*renameDemand, labels renameLabels) error {
	for kind, demand := range demands {
		if ctx.Err() != nil || demand == nil || u.published[kind] {
			continue
		}
		update := renameUpdate{Demand: *demand, State: renameStateReady, Initial: labels.initial[kind], Problem: labels.problem[kind]}
		if labels.err != nil {
			update.State, update.Initial, update.Problem = renameStateUnavailable, "", "Rename context unavailable; Esc: back"
		}
		u.retained[kind] = &update // Retain authority BEFORE renderer availability.
		data, err := json.Marshal(update)
		if err != nil {
			return err
		}
		_, u.published[kind], err = transferRenameSlot(root, renameUpdateFiles[kind], data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *popupRenameUpdates) validateProviderSelection(event MenuEvent, mapping Mapping) error {
	for kind, update := range u.providerRetained {
		if update == nil {
			continue
		}
		for _, row := range update.Rows {
			if row.Keys != mapping.Keys || row.Action == "" {
				continue
			}
			if !u.providerPublished[kind] || event.provider == nil || *event.provider != update.Binding {
				return errors.New("whichkey: provider selection has no matching retained ready publication")
			}
			return nil
		}
	}
	if event.provider != nil {
		return errors.New("whichkey: unexpected provider selection binding")
	}
	return nil
}

// Called only after stop. The binding adds no resource/action authority: keys
// still select the original retained mapping/args and exact origin target.
func (u *popupRenameUpdates) validateSelection(event MenuEvent, config Config) error {
	if event.Dismissed {
		if event.update != nil || event.provider != nil {
			return errRuntimeStartup
		}
		return nil
	}
	mapping, err := retainedNativeMapping(event, config)
	if err != nil {
		return err
	}
	if err := u.validateProviderSelection(event, mapping); err != nil {
		return err
	}
	kind := renameKind(mapping)
	if kind < 0 {
		if event.update != nil {
			return errors.New("whichkey: unexpected selection update binding")
		}
		return nil
	}
	if event.update == nil || !u.published[kind] || u.retained[kind] == nil ||
		u.retained[kind].State != renameStateReady || *event.update != u.retained[kind].Demand.binding() {
		return errors.New("whichkey: rename selection has no matching retained ready publication")
	}
	return nil
}
