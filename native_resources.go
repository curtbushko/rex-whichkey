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
	bound := popupRequest{Server: origin.Server, SessionID: origin.SessionID, BlockID: origin.BlockID}
	if requested != "" {
		// An explicit selector remains strict even though no GUI action is sent.
		// A missing/reconnected app never becomes permission to ignore that intent.
		id, generation, err := popupOriginConnection(ctx, client, origin.SessionID, requested)
		if err != nil {
			return errors.New("whichkey: native resource close refused: explicit app binding is unavailable")
		}
		bound.ClientID, bound.ConnectedAt = id, generation
	}
	window, err := placement.Window(ctx)
	if err != nil || window.ID != placement.WindowID {
		return errors.New("whichkey: native resource close refused: captured tab is unavailable or origin moved")
	}
	var current rex.Block
	if requested != "" {
		current, err = bound.originBlock(ctx, client)
	} else {
		current, err = bound.resourceOriginBlock(ctx, client)
	}
	if err != nil || current.ID != placement.ID || current.WindowID != placement.WindowID || current.LayerID != placement.LayerID {
		return errors.New("whichkey: native resource close refused: origin placement or explicit app connection changed")
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
