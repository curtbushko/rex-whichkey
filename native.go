package whichkey

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"temp/rex"
)

// App eligibility never gates resource-only display. A successful binding is
// frozen now, not resolved from whoever owns/focuses the pane after interaction.
func captureNativeApp(ctx context.Context, client *rex.Client, origin launchOrigin, placement rex.Block, requested string) (launchOrigin, error) {
	connection, err := nativePaneAppTarget(ctx, client, origin, placement, requested)
	if err != nil {
		return origin, err
	}
	origin.ClientID, origin.ConnectedAt = connection.ID, connection.ConnectedAt
	bound := popupRequest{Server: origin.Server, SessionID: origin.SessionID, BlockID: origin.BlockID,
		ClientID: connection.ID, ConnectedAt: connection.ConnectedAt}
	if originErr := validatePaneAppOrigin(ctx, client, bound, placement, requested); originErr != nil {
		return origin, originErr
	}
	return origin, nil
}

func retainedNativeMapping(message MenuEvent, config Config) (Mapping, error) {
	if message.Selection == nil || message.Dismissed {
		return Mapping{}, errors.New("whichkey: invalid retained selection")
	}
	selected := message.Selection
	if _, err := retainedNativeEvent(config, selected.mapping.Keys, selected.confirmed, selected.textInput); err != nil {
		return Mapping{}, err
	}
	for _, mapping := range config.Spec {
		if mapping.Keys != selected.mapping.Keys {
			continue
		}
		if mapping.CallbackRef != "" || mapping.Action == "" || !mapping.Enabled() {
			return Mapping{}, errors.New("whichkey: selection is not a retained Go action")
		}
		return mapping, nil
	}
	return Mapping{}, errors.New("whichkey: selection missing from retained snapshot")
}

func dispatchNativeSelection(ctx context.Context, client *rex.Client, origin launchOrigin, placement rex.Block, requested string, mapping Mapping, message MenuEvent, targets map[string]providerTarget, appErr error) error {
	if inlineRenameMapping(mapping) {
		if message.Selection.textInput == nil {
			return errors.New("whichkey: rename requires private validated text")
		}
		return nativeResourceRename(ctx, client, origin, placement, requested, mapping, *message.Selection.textInput)
	}
	switch mapping.Action {
	case menuPaneClose, menuTabClose, menuSessionClose:
		return nativeResourceClose(ctx, client, origin, placement, requested, mapping)
	case menuPaneFocus:
		target, exists := targets[mapping.Keys]
		if !exists {
			return errors.New("whichkey: missing retained pane target")
		}
		dispatch, done := context.WithTimeout(ctx, 3*time.Second)
		defer done()
		if _, _, err := nativeResourceOrigin(dispatch, client, origin, placement, requested); err != nil {
			return err
		}
		block, err := validatePaneChoice(dispatch, target)
		if err != nil {
			return err
		}
		return block.Focus(dispatch)
	default:
		if appErr != nil {
			return fmt.Errorf("whichkey: app action refused after cleanup: %w", appErr)
		}
		return dispatchNativeApp(ctx, client, origin, placement, requested, mapping)
	}
}

func dispatchNativeApp(parent context.Context, client *rex.Client, origin launchOrigin, placement rex.Block, requested string, mapping Mapping) error {
	ctx, done := context.WithTimeout(parent, 3*time.Second)
	defer done()
	bound := popupRequest{Server: origin.Server, SessionID: origin.SessionID, BlockID: origin.BlockID, ClientID: origin.ClientID, ConnectedAt: origin.ConnectedAt}
	actions, err := client.Actions(ctx, origin.ClientID)
	if err != nil {
		return errors.New("whichkey: cannot read retained app built-in actions after cleanup")
	}
	action := mapping.Action
	if action == menuTabNew {
		action = menuSessionSelect
	}
	if actions[action] != "client" {
		return fmt.Errorf("whichkey: retained app does not provide built-in client action %s", action)
	}
	if mapping.Action == menuTabNew {
		return nativeNewTab(ctx, client, bound, placement, requested, origin.selection, mapping)
	}
	perform, err := rex.NewContext(ctx, rex.Options{Selection: origin.selection, Server: origin.Server, Timeout: 3 * time.Second})
	if err != nil {
		return err
	}
	if mapping.Action == menuSessionSelect {
		if choiceErr := validateSessionChoice(ctx, client, mapping); choiceErr != nil {
			return choiceErr
		}
	}
	if originErr := validatePaneAppOrigin(ctx, client, bound, placement, requested); originErr != nil {
		return originErr
	}
	args := mapping.Args
	if len(args) == 0 {
		args = json.RawMessage(`{}`)
	}
	sessionID, blockID := origin.SessionID, origin.BlockID
	if mapping.Action == menuPaneSplit || mapping.Action == menuSessionSelect {
		sessionID, blockID = "", ""
	}
	if performErr := perform.Perform(ctx, origin.ClientID, mapping.Action, args, sessionID, blockID); performErr != nil {
		return errors.New("whichkey: native action failed or is uncertain; check app Remote Control/action compatibility; not retried")
	}
	return nil
}

func validateNativeNewTab(mapping Mapping) error {
	if mapping.Action != menuTabNew {
		return errors.New("whichkey: native new tab refused: invalid retained new-tab intent")
	}
	if len(mapping.Args) != 0 {
		if _, err := completionObject(mapping.Args); err != nil {
			return errors.New("whichkey: native new tab refused: retained new-tab arguments must be empty")
		}
	}
	return nil
}

// nativeNewTab compiles only Go-owned new-tab intent after popup cleanup and the
// caller's session.select registration check. Creation and app activation are
// distinct effects, never a retryable transaction.
func nativeNewTab(ctx context.Context, client *rex.Client, bound popupRequest, placement rex.Block, requested string, selection *rex.CLISelection, mapping Mapping) error {
	if err := validateNativeNewTab(mapping); err != nil {
		return err
	}
	unscoped, err := rex.NewContext(ctx, rex.Options{Selection: selection, Server: bound.Server,
		Timeout: 3 * time.Second})
	if err != nil {
		return errors.New("whichkey: native new tab refused: cannot construct the captured endpoint client")
	}
	session, err := client.Session(ctx, bound.SessionID)
	if err != nil {
		return errors.New("whichkey: native new tab refused: cannot inspect the captured origin session")
	}
	if originErr := validatePaneAppOrigin(ctx, client, bound, placement, requested); originErr != nil {
		return originErr
	}
	// Exactly one default-terminal creation in the captured source session.
	// NewWindow can retain a validated ID even when later response checks fail.
	window, err := session.NewWindow(ctx, "")
	if err != nil {
		if window.ID != "" {
			return fmt.Errorf("whichkey: native new tab creation response is incomplete; inspect preserved tab %s in session %s manually; activation not attempted, no retry or rollback", window.ID, bound.SessionID)
		}
		return errors.New("whichkey: native new tab creation failed or is uncertain; inspect the origin session for a possible new tab before any manual action; no retry or rollback")
	}
	// Freeze this exact returned ID. Never create again, close it on failure,
	// resolve a position, or fall back to the original client.tab.new action.
	actions, err := client.Actions(ctx, bound.ClientID)
	if err != nil || actions["session.select"] != "client" {
		return fmt.Errorf("whichkey: native new tab %s preserved in session %s; activation refused because session.select registration could not be verified; select that tab manually, no retry or rollback", window.ID, bound.SessionID)
	}
	if _, inspectErr := window.Inspect(ctx); inspectErr != nil {
		return fmt.Errorf("whichkey: native new tab %s created in session %s; membership could not be verified before activation; inspect that exact tab manually, no retry or rollback", window.ID, bound.SessionID)
	}
	// Creation can legitimately change GUI ownership. Keep the already-bound
	// app; revalidate its generation and origin placement, never pick a new owner.
	current, err := bound.originBlock(ctx, client)
	if err != nil || current.WindowID != placement.WindowID || current.LayerID != placement.LayerID {
		return fmt.Errorf("whichkey: native new tab %s preserved in session %s; origin or chosen app connection changed before activation; select that tab manually, no retry or rollback", window.ID, bound.SessionID)
	}
	// Only parent-derived destination args select the new tab. No outer
	// session/block override can redirect the chosen app away from that ID.
	if performErr := unscoped.Perform(ctx, bound.ClientID, menuSessionSelect,
		map[string]string{menuSessionID: bound.SessionID, menuWindowID: window.ID}, "", ""); performErr != nil {
		return fmt.Errorf("whichkey: native new tab %s preserved in session %s; activation failed or is uncertain; check app Remote Control/action compatibility and select that tab manually, no retry or rollback", window.ID, bound.SessionID)
	}
	return nil
}
