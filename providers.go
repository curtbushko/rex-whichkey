package whichkey

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"sync"
	"time"

	"temp/rex"
)

const (
	menuPaneFocus = "go.pane.focus"
	menuSessionID = "session_id"
	menuWindowID  = "window_id"
)

// providerTarget is retained in-process. A key selects this exact resource;
// labels, terminal input and current focus supply no identity.
type providerTarget struct{ block rex.Block }

type providerRow struct {
	id      string
	label   string
	mapping Mapping
	block   rex.Block
}

// canonicalMenuSnapshot validates in original order before changing spelling.
// Retained matching, generated prefixes and the menu trie then share one path.
func canonicalMenuSnapshot(config Config) (Config, error) {
	if err := config.Validate(); err != nil {
		return Config{}, err
	}
	spec := make([]Mapping, len(config.Spec))
	for i, mapping := range config.Spec {
		mapping = cloneMenuMapping(mapping)
		canonical, keyErr := NormalizeKeys(mapping.Keys)
		if keyErr != nil {
			return Config{}, keyErr
		}
		mapping.Keys = canonical
		spec[i] = mapping
	}
	config.Spec = spec
	return config, nil
}

func menuSnapshotBound(config Config) error {
	data, err := json.Marshal(config)
	if err != nil || len(data) > menuConfigLimit {
		return errors.New("whichkey: expanded menu exceeds 4 MiB; not truncated")
	}
	return nil
}

// menuKeyBudget admits complete canonical keys before generated concatenation.
// Subtraction keeps both the byte sum and pagination arithmetic overflow-safe.
// Serialized batch/combined-menu bounds remain separate checks.
type menuKeyBudget struct {
	rowsLeft, bytesLeft int
}

func (budget *menuKeyBudget) reserve(prefixBytes, suffixBytes int) error {
	if budget.rowsLeft <= 0 || prefixBytes > budget.bytesLeft || suffixBytes > budget.bytesLeft-prefixBytes {
		return errors.New("whichkey: provider rows exceed transport bounds; not truncated")
	}
	budget.rowsLeft--
	budget.bytesLeft -= prefixBytes + suffixBytes
	return nil
}

// reserveList mirrors hostMenuListKey, including its intermediate More row.
func (budget *menuKeyBudget) reserveList(prefixBytes, index int) (int, error) {
	suffixBytes := index / 9
	if index > 0 && index%9 == 0 {
		if err := budget.reserve(prefixBytes, suffixBytes); err != nil {
			return 0, err
		}
	}
	suffixBytes++
	if err := budget.reserve(prefixBytes, suffixBytes); err != nil {
		return 0, err
	}
	return prefixBytes + suffixBytes, nil
}

// providerRows formats one bounded inventory. Full IDs lead labels; metadata is
// escaped rather than treated as a resource selector. No row is truncated.
func providerRows(marker Mapping, entries []providerRow) ([]Mapping, error) {
	budget := menuKeyBudget{rowsLeft: 1024, bytesLeft: menuConfigLimit}
	if len(entries) == 0 {
		if err := budget.reserve(len(marker.Keys), 1); err != nil {
			return nil, err
		}
	}
	for index := range entries {
		if _, err := budget.reserveList(len(marker.Keys), index); err != nil {
			return nil, err
		}
	}
	rows := make([]Mapping, 0, 1024-budget.rowsLeft)
	add := func(mapping Mapping) {
		mapping.Hidden = marker.Hidden
		rows = append(rows, mapping)
	}
	addGroup := func(keys, label string) { add(Mapping{Keys: keys, Group: &label}) }
	if len(entries) == 0 {
		addGroup(marker.Keys+"0", "No resources in this snapshot")
	}
	for index, entry := range entries {
		keys := hostMenuListKey(marker.Keys, index, len(entries), menuProvider(marker), addGroup)
		mapping := entry.mapping
		mapping.Keys = keys
		mapping.Desc = "[" + entry.id + "] " + hostMenuText(entry.label)
		add(mapping)
	}
	return rows, nil
}

func readSessionRows(ctx context.Context, client *rex.Client, currentID string) ([]providerRow, error) {
	sessions, err := client.Sessions(ctx)
	if err != nil {
		return nil, err
	}
	if len(sessions) > maxSessionMenuEntries {
		return nil, errProviderOverflow
	}
	rows := make([]providerRow, 0, len(sessions))
	for _, session := range sessions {
		mapping, encodeErr := sessionRowMapping(session.ID, "")
		if encodeErr != nil {
			return nil, encodeErr
		}
		label := session.Label
		if session.ID == currentID {
			label = "[current] " + label
		}
		rows = append(rows, providerRow{id: session.ID, label: label, mapping: mapping})
	}
	return rows, nil
}

func readWindowRows(view rex.SessionSnapshot) ([]providerRow, error) {
	sessionID := view.Session().ID
	windows := view.Windows()
	if len(windows) > maxSessionMenuEntries {
		return nil, errProviderOverflow
	}
	rows := make([]providerRow, 0, len(windows))
	for _, window := range windows {
		mapping, encodeErr := sessionRowMapping(sessionID, window.ID)
		if encodeErr != nil {
			return nil, encodeErr
		}
		label := window.Label
		if window.Active {
			label = "[current] " + label
		}
		rows = append(rows, providerRow{id: window.ID, label: label, mapping: mapping})
	}
	return rows, nil
}

func readPaneRows(view rex.SessionSnapshot) ([]providerRow, error) {
	blocks := view.Blocks()
	if len(blocks) > maxSessionMenuEntries {
		return nil, errProviderOverflow
	}
	currentID := ""
	for _, window := range view.Windows() {
		if window.Active {
			currentID = window.FocusedBlockID
		}
	}
	rows := make([]providerRow, 0, len(blocks))
	for _, block := range blocks {
		label := "Window " + block.WindowID + " / layer " + block.LayerID
		if block.ID == currentID {
			label = "[current] " + label
		}
		rows = append(rows, providerRow{id: block.ID, label: label,
			mapping: Mapping{Action: menuPaneFocus}, block: block})
	}
	return rows, nil
}

func sessionRowMapping(sessionID, windowID string) (Mapping, error) {
	destination := map[string]string{menuSessionID: sessionID}
	if windowID != "" {
		destination[menuWindowID] = windowID
	}
	args, err := json.Marshal(destination)
	if err != nil {
		return Mapping{}, errors.New("whichkey: cannot encode retained session choice")
	}
	return Mapping{Action: menuSessionSelect, Args: args}, nil
}

var errProviderOverflow = errors.New("whichkey: provider exceeds 512 resources; not truncated")

func expandResourceMarker(marker Mapping, entries []providerRow, readErr error) ([]Mapping, error) {
	if errors.Is(readErr, errProviderOverflow) {
		return nil, readErr
	}
	root := marker
	root.Provider, root.Action, root.Args = "", "", nil
	if root.Group == nil {
		label := marker.Desc
		root.Group = &label
		root.Desc = ""
	}
	var children []Mapping
	if readErr != nil {
		budget := menuKeyBudget{rowsLeft: 1024, bytesLeft: menuConfigLimit}
		if err := budget.reserve(len(marker.Keys), 1); err != nil {
			return nil, err
		}
		label := "Warning: inventory unavailable; not retried this invocation"
		children = []Mapping{{Keys: marker.Keys + "0", Group: &label, Hidden: marker.Hidden}}
	} else {
		var formatErr error
		children, formatErr = providerRows(marker, entries)
		if formatErr != nil {
			return nil, formatErr
		}
	}
	if err := ValidateMenuProviderChildren(marker, children); err != nil {
		return nil, err
	}
	return append([]Mapping{root}, children...), nil
}

func retainPaneTargets(targets map[string]providerTarget, rows []Mapping, entries []providerRow) {
	index := 0
	for _, row := range rows {
		if row.Action == menuPaneFocus {
			targets[row.Keys] = providerTarget{block: entries[index].block}
			index++
		}
	}
}

// Four finite inventory kinds share the same single-publication contract as
// rename labels. The last job kind loads both rename labels from the origin view.
const metadataRenameKind = 4

var providerKinds = [4]string{menuProviderSessions, menuProviderWindows, menuProviderPanes, menuProviderHosts}

func providerKind(mapping Mapping) int {
	for kind, name := range providerKinds {
		if menuProvider(mapping) == name {
			return kind
		}
	}
	return -1
}

type providerUpdate struct {
	Binding renameBinding `json:"binding"`
	Rows    []Mapping     `json:"rows"`
}

type metadataResult struct {
	kind    int
	entries []providerRow
	hosts   rex.HostInventory
	labels  renameLabels
	err     error
}

// Only the owning UI loop changes scheduling state. Workers receive
// captured read-only context and send one detached value. There are at most two
// executing goroutines, two result slots and five once-only jobs; no work queue
// can grow with key presses or repeated markers.
type metadataJobs struct {
	ctx             context.Context
	cancel          context.CancelFunc
	workers         sync.WaitGroup
	wanted, started [5]bool
	active          int
	results         chan metadataResult
	request         popupRequest
	client          *rex.Client
}

func newMetadataJobs(ctx context.Context, request popupRequest, client *rex.Client) *metadataJobs {
	ctx, cancel := context.WithCancel(ctx)
	return &metadataJobs{ctx: ctx, cancel: cancel, results: make(chan metadataResult, 2), request: request, client: client}
}

func (jobs *metadataJobs) stop() {
	jobs.cancel()
	jobs.workers.Wait()
}

func (jobs *metadataJobs) start() {
	for kind, wanted := range jobs.wanted {
		if jobs.ctx.Err() != nil || jobs.active >= 2 {
			return
		}
		if !wanted || jobs.started[kind] {
			continue
		}
		jobs.started[kind] = true
		jobs.active++
		jobs.workers.Add(1)
		go func() {
			defer jobs.workers.Done()
			result := loadMetadata(jobs.ctx, jobs.request, jobs.client, kind)
			select {
			case jobs.results <- result:
			case <-jobs.ctx.Done():
			}
		}()
	}
}

func loadMetadata(parent context.Context, request popupRequest, client *rex.Client, kind int) metadataResult {
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()
	result := metadataResult{kind: kind}
	if kind == metadataRenameKind {
		result.labels = loadRenameLabels(ctx, client, request)
		return result
	}
	if kind == 3 {
		result.hosts, result.err = rex.Hosts(ctx, rex.InventoryOptions{Selection: request.selection, Timeout: 3 * time.Second})
		return result
	}
	if client == nil || request.snapshot == nil || request.snapshot.Session().ID != request.SessionID {
		result.err = errors.New("whichkey: native resource origin required")
		return result
	}
	switch kind {
	case 0:
		result.entries, result.err = readSessionRows(ctx, client, request.SessionID)
	case 1:
		result.entries, result.err = readWindowRows(*request.snapshot)
	case 2:
		result.entries, result.err = readPaneRows(*request.snapshot)
	}
	return result
}

// A provider root is always a group, including the legacy Switch action marker.
// The placeholder cannot execute or reserve a future resource key for replay.
func pendingProviderRows(marker Mapping) ([]Mapping, error) {
	root := marker
	root.Provider, root.Action, root.Args = "", "", nil
	if root.Group == nil {
		label := root.Desc
		root.Group, root.Desc = &label, ""
	}
	budget := menuKeyBudget{rowsLeft: 1024, bytesLeft: menuConfigLimit}
	if err := budget.reserve(len(root.Keys), 1); err != nil {
		return nil, err
	}
	label := "Loading inventory; choices disabled"
	return []Mapping{root, {Keys: root.Keys + "0", Group: &label, Hidden: root.Hidden}}, nil
}

func providerMarkerRows(marker Mapping, ready [4]*providerUpdate) ([]Mapping, error) {
	kind := providerKind(marker)
	if !marker.Enabled() || kind < 0 {
		return []Mapping{marker}, nil
	}
	if ready[kind] == nil {
		return pendingProviderRows(marker)
	}
	return providerSubtree(marker.Keys, ready[kind].Rows), nil
}

func providerSubtree(prefix string, batch []Mapping) []Mapping {
	var rows []Mapping
	for _, row := range batch {
		if row.Keys == prefix || strings.HasPrefix(row.Keys, prefix) {
			rows = append(rows, row)
		}
	}
	return rows
}

// Combine complete kind batches in original marker order. Validation precedes
// publication, so overflow/conflict never becomes a plausible truncated menu.
func providerConfig(base Config, ready [4]*providerUpdate) (Config, error) {
	config := base
	config.Spec = nil
	bytesUsed := 0
	for _, marker := range base.Spec {
		rows, err := providerMarkerRows(marker, ready)
		if err != nil {
			return Config{}, err
		}
		data, err := json.Marshal(rows)
		if err != nil {
			return Config{}, err
		}
		bytesUsed += len(data)
		if bytesUsed > menuConfigLimit {
			return Config{}, errors.New("whichkey: combined menu exceeds 4 MiB; not truncated")
		}
		config.Spec = append(config.Spec, rows...)
	}
	if err := validateMenuConflicts(config.Spec); err != nil {
		return Config{}, err
	}
	if err := menuSnapshotBound(config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func buildProviderUpdate(base Config, demand renameDemand, result metadataResult) (providerUpdate, map[string]providerTarget, error) {
	update := providerUpdate{Binding: demand.binding()}
	targets := make(map[string]providerTarget)
	bytesUsed := 0
	for _, marker := range base.Spec {
		if !marker.Enabled() || providerKind(marker) != result.kind {
			continue
		}
		var rows []Mapping
		var err error
		if result.kind == 3 {
			rows, err = expandedHostRows(marker, result.hosts, result.err)
		} else {
			rows, err = expandResourceMarker(marker, result.entries, result.err)
		}
		if err != nil {
			return providerUpdate{}, nil, err
		}
		data, err := json.Marshal(rows)
		if err != nil {
			return providerUpdate{}, nil, err
		}
		bytesUsed += len(data)
		if bytesUsed > menuConfigLimit {
			return providerUpdate{}, nil, errors.New("whichkey: provider batch exceeds 4 MiB; not truncated")
		}
		retainPaneTargets(targets, rows, result.entries)
		update.Rows = append(update.Rows, rows...)
	}
	return update, targets, nil
}

func validateProviderDemand(demand renameDemand, invocation string, kind int, config Config) error {
	if demand.Invocation != invocation || demand.Kind != providerKinds[kind] || demand.Load != 1 {
		return errors.New("whichkey: stale or unknown provider demand")
	}
	for _, marker := range config.Spec {
		if marker.Enabled() && providerKind(marker) == kind && marker.Keys == demand.Keys {
			return nil
		}
	}
	return errors.New("whichkey: provider demand is not an enabled retained marker")
}

// Validate every complete marker expansion, its original root and all children.
// No child can introduce another provider, escape the subtree or nominate IDs
// for parent authority. The parent's separately retained batch is authoritative.
func validateProviderUpdate(base Config, demand renameDemand, update providerUpdate) error {
	if update.Binding != demand.binding() {
		return errors.New("whichkey: stale provider publication")
	}
	seen := make(map[string]bool)
	for _, marker := range base.Spec {
		if !marker.Enabled() || menuProvider(marker) != demand.Kind {
			continue
		}
		rows := providerSubtree(marker.Keys, update.Rows)
		for _, row := range rows {
			if seen[row.Keys] {
				return errors.New("whichkey: duplicate provider publication key")
			}
			seen[row.Keys] = true
		}
		if err := validateProviderSubtree(marker, rows); err != nil {
			return err
		}
	}
	if len(seen) != len(update.Rows) {
		return errors.New("whichkey: provider publication escaped retained markers")
	}
	return nil
}

func validateProviderSubtree(marker Mapping, rows []Mapping) error {
	if len(rows) < 2 {
		return errors.New("whichkey: incomplete provider publication")
	}
	pending, err := pendingProviderRows(marker)
	if err != nil {
		return err
	}
	root, err := json.Marshal(pending[0])
	if err != nil {
		return err
	}
	actual, err := json.Marshal(rows[0])
	if err != nil || string(root) != string(actual) {
		return errors.New("whichkey: changed provider root")
	}
	return ValidateMenuProviderChildren(marker, rows[1:])
}

func validatePaneChoice(ctx context.Context, target providerTarget) (rex.Block, error) {
	block, err := target.block.Inspect(ctx)
	if err != nil || block.ID != target.block.ID || block.WindowID != target.block.WindowID || block.LayerID != target.block.LayerID {
		return rex.Block{}, errors.New("whichkey: retained pane is unavailable or moved; not retargeted")
	}
	return block, nil
}
