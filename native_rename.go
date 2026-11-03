package whichkey

import (
	"context"
	"errors"
	"time"

	"temp/rex"
)

// nativeRenameTarget preserves Rex's native dialog, routing it to an explicit
// app or the launching pane's documented GUI owner, not a session-wide chooser.
// Ownership supplies routing context only; attachment, kind and generation must
// still pass the normal app checks, and Rex enforces Remote Control permissions.
func nativeRenameTarget(parent context.Context, client *rex.Client, origin launchOrigin, placement rex.Block, requested string) (popupAppConnection, error) {
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()
	id := requested
	if id == "" {
		var err error
		id, err = nativeRenameOwner(ctx, client, origin, placement.WindowID)
		if err != nil {
			return popupAppConnection{}, err
		}
	}
	id, generation, err := popupOriginConnection(ctx, client, origin.SessionID, id)
	if err != nil {
		return popupAppConnection{}, errors.New("whichkey: rename refused: pane owner or explicit app is not a live graphical app attached to the origin session")
	}
	return popupAppConnection{ID: id, ConnectedAt: generation}, nil
}

// nativeRenameOwner reads authoritative block metadata, not the last resize
// reporter or an app label. A missing/moved block or absent owner never selects
// another app. The owner is a routing policy, not proof of input provenance.
func nativeRenameOwner(ctx context.Context, client *rex.Client, origin launchOrigin, windowID string) (string, error) {
	var result struct {
		Block struct {
			ID        string  `json:"block_id"`
			WindowID  string  `json:"window_id"`
			Placement string  `json:"placement"`
			Owner     *string `json:"owner"`
		} `json:"block"`
	}
	if err := client.Call(ctx, "session.describe_block", map[string]string{
		"session_id": origin.SessionID, "block_id": origin.BlockID,
	}, &result); err != nil {
		return "", errors.New("whichkey: rename refused: cannot inspect the launching pane's owner")
	}
	block := result.Block
	if block.ID != origin.BlockID || block.Placement != "placed" || block.WindowID != windowID {
		return "", errors.New("whichkey: rename refused: launching pane moved or its captured tab is unavailable")
	}
	if block.Owner == nil || !popupResourceID(*block.Owner, "client") {
		return "", errors.New("whichkey: rename refused: launching pane has no GUI owner; an exact app binding is required")
	}
	return *block.Owner, nil
}
