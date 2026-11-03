package whichkey

import (
	"context"
	"errors"
	"time"

	"temp/rex"
)

// nativeResourceClose runs after confirmed selection and complete popup cleanup.
// The captured resources, not an app's current focus, determine what is closed.
func nativeResourceClose(parent context.Context, client *rex.Client, origin launchOrigin, placement rex.Block, requested string, mapping Mapping) error {
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()
	if !mapping.Confirm {
		return errors.New("whichkey: native resource close requires retained confirmation")
	}
	if len(mapping.Args) != 0 {
		if _, err := completionObject(mapping.Args); err != nil {
			return errors.New("whichkey: native resource close requires empty retained arguments")
		}
	}
	if mapping.Action != "pane.close" && mapping.Action != "client.tab.close" && mapping.Action != "session.close" {
		return errors.New("whichkey: unsupported retained resource close")
	}
	current, window, err := nativeResourceOrigin(ctx, client, origin, placement, requested)
	if err != nil {
		return err
	}
	// Each operation uses private full resource IDs on the captured server.
	// Concurrent moves can still race these checks; there is no atomic Rex
	// placement-and-close operation. Never retry an uncertain close.
	switch mapping.Action {
	case "pane.close":
		err = current.Close(ctx)
	case "client.tab.close":
		err = window.Close(ctx)
	case "session.close":
		err = current.Session().Close(ctx)
	}
	if err != nil {
		return errors.New("whichkey: native resource close failed or is uncertain; inspect the original resource; not retried")
	}
	return nil
}

// nativeResourceOrigin revalidates the parent-retained resource, not an app's
// selected view. An explicit app override stays strict but is never required.
func nativeResourceOrigin(ctx context.Context, client *rex.Client, origin launchOrigin, placement rex.Block, requested string) (rex.Block, rex.Window, error) {
	bound := popupRequest{Server: origin.Server, SessionID: origin.SessionID, BlockID: origin.BlockID}
	if requested != "" {
		// An explicit selector remains strict even though no GUI action is sent.
		// A missing/reconnected app never becomes permission to ignore that intent.
		id, generation, err := popupOriginConnection(ctx, client, origin.SessionID, requested)
		if err != nil {
			return rex.Block{}, rex.Window{}, errors.New("whichkey: resource action refused: explicit app binding is unavailable")
		}
		bound.ClientID, bound.ConnectedAt = id, generation
	}
	window, err := placement.Window(ctx)
	if err != nil || window.ID != placement.WindowID {
		return rex.Block{}, rex.Window{}, errors.New("whichkey: resource action refused: captured tab is unavailable or origin moved")
	}
	var current rex.Block
	if requested != "" {
		current, err = bound.originBlock(ctx, client)
	} else {
		current, err = bound.resourceOriginBlock(ctx, client)
	}
	if err != nil || current.ID != placement.ID || current.WindowID != placement.WindowID || current.LayerID != placement.LayerID {
		return rex.Block{}, rex.Window{}, errors.New("whichkey: resource action refused: origin placement or explicit app connection changed")
	}
	return current, window, nil
}

// nativeResourceRename receives name data after complete popup cleanup. The
// parent retains both tab/session identities and the action; none come from text.
func nativeResourceRename(parent context.Context, client *rex.Client, origin launchOrigin, placement rex.Block, requested string, mapping Mapping, name string) error {
	if !nativeRenameMapping(mapping) {
		return errors.New("whichkey: invalid retained rename intent")
	}
	if err := validateRenameText(name); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()
	current, window, err := nativeResourceOrigin(ctx, client, origin, placement, requested)
	if err != nil {
		return err
	}
	// Rename passes label as one argv value, never shell code. There is no
	// app action or retry, including when a failed reply may follow mutation.
	if mapping.Action == "session.rename" {
		err = current.Session().Rename(ctx, name)
	} else {
		err = window.Rename(ctx, name)
	}
	if err != nil {
		return errors.New("whichkey: rename failed or is uncertain; inspect the captured resource; not retried")
	}
	return nil
}
