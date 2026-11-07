package whichkey

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	"temp/rex"
)

// DirectOptions is invocation data captured by Lua before creating the float.
// Popup REX_SESSION/REX_BLOCK identify only the execution host, never the origin.
// Bare invocation is rejected; only the explicit direct command enters this mode.
type DirectOptions struct {
	OriginSession string
	OriginBlock   string
	OriginWindow  string
	OriginLayer   string
	// Nil selectors are absent; explicitly empty selectors are invalid.
	Server  *string
	Client  *string
	Config  *string
	ASCII   bool
	NoColor bool
}

// RunDirectMenu runs inside the Lua-created terminal. It creates no menu child,
// image copy, private transport or nested popup. Rex owns terminal exit cleanup.
func RunDirectMenu(parent context.Context, options DirectOptions) error {
	ownID, requested, err := directInvocation(parent, options)
	if err != nil {
		return err
	}
	// Never read the popup's REX_CLIENT as original app intent.
	ctx, stop := signal.NotifyContext(parent, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGPIPE, syscall.SIGTSTP)
	defer stop()
	selection, err := rex.ResolveCLI(ctx, rex.CLIOptions{})
	if err != nil {
		return err
	}
	endpoint, err := rex.ResolveEndpoint(ctx, options.Server, selection)
	if err != nil {
		return err
	}
	client, err := rex.NewContext(ctx, rex.Options{Selection: selection, Server: endpoint,
		Session: options.OriginSession, Timeout: 3 * time.Second})
	if err != nil {
		return err
	}
	origin := launchOrigin{Server: endpoint, SessionID: options.OriginSession, BlockID: options.OriginBlock, selection: client.Selection()}
	view, err := client.SessionSnapshot(ctx, origin.SessionID)
	if err != nil {
		return err
	}
	placement, err := view.Block(origin.BlockID)
	if err != nil || placement.WindowID != options.OriginWindow || placement.LayerID != options.OriginLayer {
		return errors.New("whichkey: original placement changed since Lua invocation; not retargeted")
	}
	popup, snapshot, err := view.DirectPopup(ownID, placement)
	if err != nil {
		return err // In particular, never detach an ordinary tiled caller pane.
	}
	state := directMenuState{targets: make(map[string]providerTarget)}
	message, bound, appErr, err := directSelection(ctx, client, origin, placement, requested, popup, snapshot, options, &state)
	if err != nil {
		return err
	}
	if message.Dismissed {
		return nil
	}
	if contextErr := ctx.Err(); contextErr != nil {
		return contextErr
	}
	mapping, err := state.validateSelection(message)
	if err != nil {
		return err
	}
	return dispatchNativeSelection(ctx, client, bound, placement, requested, mapping, message, state.targets, appErr)
}

// Cleanup is armed as soon as ownership is bound, even if config/app/TTY setup
// fails. RunTerminal joins its workers and restores/closes its TTY before this
// defer can retire the float. A cleanup error clears selection authority.
func directSelection(ctx context.Context, client *rex.Client, origin launchOrigin, placement rex.Block, requested string,
	popup rex.Layer, snapshot rex.SessionSnapshot, options DirectOptions, state *directMenuState,
) (event MenuEvent, bound launchOrigin, appErr error, outcome error) {
	bound = origin
	defer func() {
		cleanup, done := context.WithTimeout(context.WithoutCancel(ctx), 9*time.Second)
		defer done()
		outcome = errors.Join(outcome, retireDirectPopup(cleanup, client, bound, placement, requested, popup))
		if outcome != nil {
			event = MenuEvent{}
		}
	}()
	bound, appErr = captureNativeApp(ctx, client, origin, placement, requested)
	if appErr != nil && requested != "" {
		return MenuEvent{}, bound, appErr, appErr
	}
	config, err := LoadMenuConfig(options.Config)
	if err != nil {
		return MenuEvent{}, bound, appErr, err
	}
	config, err = canonicalMenuSnapshot(config)
	if err != nil {
		return MenuEvent{}, bound, appErr, err
	}
	config.textEntry = &textEntrySpec{}
	config.Position = positionTop // Lua already placed the fixed outer viewport.
	state.request = popupRequest{Server: origin.Server, SessionID: origin.SessionID, BlockID: origin.BlockID,
		selection: origin.selection, placement: &placement, snapshot: &snapshot}
	state.client = client
	ctx = context.WithValue(ctx, terminalDirectKey{}, state)
	event, outcome = RunTerminal(ctx, config, options.NoColor, options.ASCII)
	return event, bound, appErr, outcome
}

func retireDirectPopup(ctx context.Context, client *rex.Client, origin launchOrigin, placement rex.Block, requested string, popup rex.Layer) error {
	// No cascading close fallback and no retry even when detach may have run.
	// The terminal's on_completion/quick_exit_threshold_ms=0 policy also covers
	// cancellation and failures. Only the coordinator can observe post-exit deletion.
	if err := popup.Detach(ctx); err != nil {
		return err
	}
	current, _, err := nativeResourceOrigin(ctx, client, origin, placement, requested)
	if err != nil {
		return err
	}
	if origin.ClientID != "" {
		bound := popupRequest{Server: origin.Server, SessionID: origin.SessionID, BlockID: origin.BlockID,
			ClientID: origin.ClientID, ConnectedAt: origin.ConnectedAt}
		if originErr := validatePaneAppOrigin(ctx, client, bound, placement, requested); originErr != nil {
			return originErr
		}
	}
	if focusErr := current.Focus(ctx); focusErr != nil {
		return errors.New("whichkey: original focus restoration failed or is uncertain; no effect, no retry")
	}
	view, err := client.SessionSnapshot(ctx, origin.SessionID)
	if err != nil {
		return err
	}
	block, err := view.Block(origin.BlockID)
	if err != nil || block.WindowID != placement.WindowID || block.LayerID != placement.LayerID {
		return errors.New("whichkey: original placement changed during focus restoration")
	}
	window, err := view.Window(placement.WindowID)
	if err != nil || !window.Active || window.FocusedBlockID != origin.BlockID {
		return errors.New("whichkey: exact original focus could not be verified; no selected effect")
	}
	if origin.ClientID != "" {
		return validatePaneAppOrigin(ctx, client, popupRequest{Server: origin.Server, SessionID: origin.SessionID,
			BlockID: origin.BlockID, ClientID: origin.ClientID, ConnectedAt: origin.ConnectedAt}, placement, requested)
	}
	return nil
}

func directInvocation(parent context.Context, options DirectOptions) (string, string, error) {
	if parent == nil || !popupResourceID(options.OriginSession, "session") ||
		!popupResourceID(options.OriginBlock, "block") || !popupResourceID(options.OriginWindow, "window") ||
		!popupResourceID(options.OriginLayer, "layer") {
		return "", "", errors.New("whichkey: direct requires full original session/block/window/layer IDs captured before popup creation")
	}
	ownID := os.Getenv("REX_BLOCK")
	if os.Getenv("REX_SESSION") != options.OriginSession || !popupResourceID(ownID, "block") || ownID == options.OriginBlock {
		return "", "", errors.New("whichkey: direct must run in its own popup terminal, not the original pane")
	}
	requested := ""
	if options.Client != nil {
		if !popupResourceID(*options.Client, "client") {
			return "", "", errors.New("whichkey: direct --client must be an original full client ID")
		}
		requested = *options.Client
	}
	return ownID, requested, nil
}

func (state *directMenuState) applyInventory(menu *Menu, result metadataResult) error {
	demand := menu.providers.wanted[result.kind]
	if demand == nil {
		return errors.New("whichkey: unexpected in-process inventory result")
	}
	update, targets, err := buildProviderUpdate(menu.providers.base, *demand, result)
	if err != nil {
		return err
	}
	ready := state.providers
	ready[result.kind] = &update
	next, err := providerConfig(menu.providers.base, ready)
	if err != nil {
		return err
	}
	state.config, state.providers = next, ready
	for key, target := range targets {
		state.targets[key] = target
	}
	return menu.applyProviderUpdate(result.kind, update)
}

func (state *directMenuState) applyRenameLabels(menu *Menu) (bool, error) {
	changed := false
	if state.labels != nil {
		for kind, demand := range menu.updates.wanted {
			if demand == nil || state.rename[kind] != nil {
				continue
			}
			update := renameUpdate{Demand: *demand, State: renameStateReady,
				Initial: state.labels.initial[kind], Problem: state.labels.problem[kind]}
			if state.labels.err != nil {
				update.State, update.Initial, update.Problem = renameStateUnavailable, "", "Rename context unavailable; Esc: back"
			}
			if err := validateRenameUpdate(update, *demand); err != nil {
				return false, err
			}
			state.rename[kind] = &update
			menu.applyRenameUpdate(kind, update)
			changed = true
		}
	}
	return changed, nil
}

func (state *directMenuState) validateRenameSelection(mapping Mapping, event MenuEvent) error {
	if kind := renameKind(mapping); kind >= 0 {
		update := state.rename[kind]
		if update == nil || update.State != renameStateReady || event.update == nil || *event.update != update.Demand.binding() {
			return errors.New("whichkey: rename selection lacks retained ready metadata")
		}
	} else if event.update != nil {
		return errors.New("whichkey: unexpected rename metadata on selection")
	}
	return nil
}

type terminalDirectKey struct{}

// All fields are owned by RunTerminal's one UI goroutine. Workers receive only
// immutable request snapshots; authority is retained before any ready UI update.
type directMenuState struct {
	request   popupRequest
	client    *rex.Client
	config    Config
	targets   map[string]providerTarget
	labels    *renameLabels
	rename    [2]*renameUpdate
	providers [4]*providerUpdate
}

func (state *directMenuState) poll(menu *Menu, jobs *metadataJobs) (bool, error) {
	for kind, demand := range menu.providers.wanted {
		jobs.wanted[kind] = demand != nil
	}
	for _, demand := range menu.updates.wanted {
		if demand != nil {
			jobs.wanted[metadataRenameKind] = true
		}
	}
	changed := false
	select {
	case result := <-jobs.results:
		jobs.active--
		if err := jobs.ctx.Err(); err != nil {
			return false, err
		}
		if result.kind == metadataRenameKind {
			state.labels = &result.labels
		} else {
			if err := state.applyInventory(menu, result); err != nil {
				return false, err
			}
			changed = true
		}
	default:
	}
	renameChanged, err := state.applyRenameLabels(menu)
	if err != nil {
		return false, err
	}
	changed = changed || renameChanged
	jobs.start()
	return changed, nil
}

func (state *directMenuState) validateSelection(event MenuEvent) (Mapping, error) {
	mapping, err := retainedNativeMapping(event, state.config)
	if err != nil {
		return Mapping{}, err
	}
	if err := state.validateRenameSelection(mapping, event); err != nil {
		return Mapping{}, err
	}
	if err := state.validateProviderSelection(mapping, event); err != nil {
		return Mapping{}, err
	}
	return mapping, nil
}

func (state *directMenuState) validateProviderSelection(mapping Mapping, event MenuEvent) error {
	for _, update := range state.providers {
		if update == nil {
			continue
		}
		for _, row := range update.Rows {
			if row.Keys == mapping.Keys && row.Action != "" {
				if event.provider == nil || *event.provider != update.Binding {
					return errors.New("whichkey: inventory selection lacks retained ready metadata")
				}
				return nil
			}
		}
	}
	if event.provider != nil {
		return errors.New("whichkey: unexpected inventory metadata on selection")
	}
	return nil
}
