package rex

import (
	"context"
	"errors"
	"fmt"
)

// DirectPopup binds the caller's exact terminal to a sole-block floating layer,
// separate from the invocation origin. The returned view excludes that layer:
// popup creation must not consume a pane-provider slot or become a target.
func (v SessionSnapshot) DirectPopup(ownID string, origin Block) (Layer, SessionSnapshot, error) {
	if ownID == origin.ID || origin.Session().ID != v.session.ID {
		return Layer{}, SessionSnapshot{}, errors.New("rex: popup and origin must be distinct blocks in the captured session")
	}
	current, err := v.Block(origin.ID)
	if err != nil || current.WindowID != origin.WindowID || current.LayerID != origin.LayerID {
		return Layer{}, SessionSnapshot{}, errors.New("rex: original placement changed before popup binding")
	}
	own, err := v.Block(ownID)
	if err != nil || own.WindowID != origin.WindowID || own.LayerID == origin.LayerID {
		return Layer{}, SessionSnapshot{}, errors.New("rex: own terminal is not a separate float in the original window")
	}
	filtered, err := v.withoutDirectPopup(own)
	if err != nil {
		return Layer{}, SessionSnapshot{}, err
	}
	// Creation took focus. Provider current markers describe invocation intent,
	// not the popup's new focus, without altering the raw ownership snapshot.
	filtered.restoreOriginMarkers(origin)
	return Layer{ID: own.LayerID, BlockID: ownID, WindowID: own.WindowID,
		client: own.client, layerID: own.LayerID, blockID: ownID}, filtered, nil
}

func (v SessionSnapshot) withoutDirectPopup(own Block) (SessionSnapshot, error) {
	filtered := v
	filtered.layers = make([]placedLayer, 0, len(v.layers)-1)
	found := false
	for _, layer := range v.layers {
		if layer.id != own.LayerID {
			filtered.layers = append(filtered.layers, layer)
			continue
		}
		if layer.kind != "floating" || len(layer.blocks) != 1 || layer.blocks[0].ID != own.ID {
			return SessionSnapshot{}, errors.New("rex: own popup must be the sole block in its floating layer")
		}
		found = true
	}
	if !found {
		return SessionSnapshot{}, errors.New("rex: own popup layer is missing")
	}
	return filtered, nil
}

func (v *SessionSnapshot) restoreOriginMarkers(origin Block) {
	v.windows = v.Windows()
	for i := range v.windows {
		window := &v.windows[i]
		window.Active = window.ID == origin.WindowID
		if window.ID == origin.WindowID {
			window.FocusedBlockID = origin.ID
		}
	}
}

// Detach retires only the bound sole-block float, keeping its program alive.
// It is a single mutation, never retried, followed by independent checks of
// visible-layer retirement and the exact terminal's detached placement. Rex's
// terminal exit policy, not this method, deletes the host after Go exits.
func (l Layer) Detach(ctx context.Context) error {
	if err := l.checkOwned(ctx); err != nil {
		return err
	}
	if err := l.terminalPlacement(ctx, "placed"); err != nil {
		return err
	}
	if err := l.client.Call(ctx, "session.detach_block", map[string]string{
		fieldSession: l.client.session, "block_id": l.blockID,
	}, nil); err != nil {
		return fmt.Errorf("rex: popup detach failed or is uncertain; not retried: %w", err)
	}
	view, err := l.client.SessionSnapshot(ctx, l.client.session)
	if err != nil {
		return err
	}
	for _, layer := range view.layers {
		if layer.id == l.layerID {
			return errors.New("rex: detached popup layer is still visible")
		}
		for _, block := range layer.blocks {
			if block.ID == l.blockID {
				return errors.New("rex: popup block is still placed after detach")
			}
		}
	}
	return l.terminalPlacement(ctx, "detached")
}

func (l Layer) terminalPlacement(ctx context.Context, placement string) error {
	var result struct {
		Blocks []struct {
			ID        string `json:"block_id"`
			Placement string `json:"placement"`
			WindowID  string `json:"window_id"`
			Creator   string `json:"creator_name"`
			Flavor    string `json:"flavor"`
		} `json:"blocks"`
	}
	if err := l.client.Call(ctx, "session.list_blocks", map[string]string{fieldSession: l.client.session}, &result); err != nil {
		return err
	}
	if result.Blocks == nil {
		return errors.New("rex: missing block inventory for popup verification")
	}
	seen, found := make(map[string]bool, len(result.Blocks)), false
	for _, block := range result.Blocks {
		if !resourceID(block.ID, "block") || seen[block.ID] {
			return errors.New("rex: invalid or duplicate block inventory identity")
		}
		seen[block.ID] = true
		if block.ID != l.blockID {
			continue
		}
		if block.Placement != placement || block.Creator != "com.superlogical.terminal" ||
			block.Flavor != "com.superlogical.terminal.shell" ||
			(placement == "placed" && block.WindowID != l.WindowID) ||
			(placement == "detached" && block.WindowID != "") {
			return errors.New("rex: own terminal placement or flavor could not be verified")
		}
		found = true
	}
	if !found {
		return errors.New("rex: own popup terminal is missing")
	}
	return nil
}
