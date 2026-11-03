package whichkey

import (
	"context"
	"errors"
	"time"

	"temp/rex"
)

// nativePaneAppTarget binds all app actions and configured CLI launches to an
// explicit app or the launching pane's documented GUI owner. It never prompts
// for an app. Ownership is routing context, not extra permission; attachment,
// kind, generation and Rex Remote Control checks still apply.
func nativePaneAppTarget(parent context.Context, client *rex.Client, origin launchOrigin, placement rex.Block, requested string) (popupAppConnection, error) {
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()
	id := requested
	if id == "" {
		var err error
		id, err = nativePaneAppOwner(ctx, client, origin, placement.WindowID)
		if err != nil {
			return popupAppConnection{}, err
		}
	}
	id, generation, err := popupOriginConnection(ctx, client, origin.SessionID, id)
	if err != nil {
		return popupAppConnection{}, errors.New("whichkey: pane app routing refused: owner or explicit app is not a live graphical app attached to the origin session")
	}
	return popupAppConnection{ID: id, ConnectedAt: generation}, nil
}

// nativePaneAppOwner reads authoritative block metadata, not the last resize
// reporter or an app label. A missing/moved block or absent owner never selects
// another app. The owner is a routing policy, not proof of input provenance.
func nativePaneAppOwner(ctx context.Context, client *rex.Client, origin launchOrigin, windowID string) (string, error) {
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
		return "", errors.New("whichkey: pane app routing refused: cannot inspect the launching pane's owner")
	}
	block := result.Block
	if block.ID != origin.BlockID || block.Placement != "placed" || block.WindowID != windowID {
		return "", errors.New("whichkey: pane app routing refused: launching pane moved or its captured tab is unavailable")
	}
	if block.Owner == nil || !popupResourceID(*block.Owner, "client") {
		return "", errors.New("whichkey: pane app routing refused: launching pane has no GUI owner; an exact app binding is required")
	}
	return *block.Owner, nil
}

// validatePaneAppOrigin checks the frozen binding before the first effect.
// Own mutations may then change focus/ownership; later effects must retain this
// exact app and generation rather than resolving whichever owner appears next.
func validatePaneAppOrigin(ctx context.Context, client *rex.Client, bound popupRequest, placement rex.Block, requested string) error {
	current, err := bound.originBlock(ctx, client)
	if err != nil {
		return errors.New("whichkey: pane app routing refused: origin or app connection changed")
	}
	if current.ID != placement.ID || current.WindowID != placement.WindowID || current.LayerID != placement.LayerID {
		return errors.New("whichkey: pane app routing refused: launching pane moved after binding")
	}
	if requested == "" {
		origin := launchOrigin{SessionID: bound.SessionID, BlockID: bound.BlockID}
		owner, err := nativePaneAppOwner(ctx, client, origin, placement.WindowID)
		if err != nil {
			return err
		}
		if owner != bound.ClientID {
			return errors.New("whichkey: pane app routing refused: launching pane's GUI owner changed; not retargeted")
		}
	}
	return nil
}
