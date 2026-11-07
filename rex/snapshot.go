package rex

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

// SessionSnapshot retains one exact, validated session view for an invocation.
// Its collections are private; lookups return detached typed metadata. It is
// never used by Inspect, whose freshness contract remains a new server read.
type SessionSnapshot struct {
	session Session
	windows []Window
	layers  []placedLayer
}

// SessionSnapshot reads labels, tabs and block placement in one bounded CLI read.
// A full target ID must match exactly; a human target is resolved only once.
func (c *Client) SessionSnapshot(ctx context.Context, target string) (SessionSnapshot, error) {
	if c == nil {
		return SessionSnapshot{}, errors.New("rex: client is required")
	}
	if target == "" {
		target = c.session
	}
	if strings.TrimSpace(target) == "" || strings.ContainsRune(target, '\x00') {
		return SessionSnapshot{}, errors.New("rex: session target is required and must not contain NUL")
	}
	var view sessionView
	if err := c.resourceCommand(ctx, &view, "session", "inspect", "--json", "--", target); err != nil {
		return SessionSnapshot{}, err
	}
	if err := view.validate(); err != nil {
		return SessionSnapshot{}, err
	}
	if resourceID(target, "session") && view.ID != target {
		return SessionSnapshot{}, fmt.Errorf("rex session inspect: requested %q, received %q", target, view.ID)
	}
	session := c.sessionResource(view.ID, view.Label)
	layers, err := session.snapshotLayers(view)
	if err != nil {
		return SessionSnapshot{}, err
	}
	windows := make([]Window, 0, len(view.Windows))
	for _, item := range view.Windows {
		window := session.windowResource(item.ID, item.Label)
		window.FocusedBlockID, window.Active = item.FocusedBlockID, item.ID == view.ActiveWindowID
		windows = append(windows, window)
	}
	return SessionSnapshot{session: session, windows: windows, layers: layers}, nil
}

// Session returns the exact captured session and its label at read time.
func (v SessionSnapshot) Session() Session { return v.session }

// Windows returns detached metadata in captured server tab order.
func (v SessionSnapshot) Windows() []Window {
	return append([]Window{}, v.windows...)
}

// Blocks returns detached placement metadata in captured window/layer order.
func (v SessionSnapshot) Blocks() []Block {
	blocks := make([]Block, 0)
	for _, layer := range v.layers {
		blocks = append(blocks, layer.blocks...)
	}
	return blocks
}

// Window returns an exact member, never a label, position or active-tab match.
func (v SessionSnapshot) Window(id string) (Window, error) {
	if !resourceID(id, "window") {
		return Window{}, errors.New("rex: a full window ID is required")
	}
	for _, window := range v.windows {
		if window.windowID == id {
			return window, nil
		}
	}
	return Window{}, fmt.Errorf("rex: window %q not found in snapshot", id)
}

// Block returns a uniquely placed exact member from this view.
func (v SessionSnapshot) Block(id string) (Block, error) {
	if !resourceID(id, "block") {
		return Block{}, errors.New("rex: a full block ID is required")
	}
	for _, layer := range v.layers {
		for _, block := range layer.blocks {
			if block.blockID == id {
				return block, nil
			}
		}
	}
	return Block{}, fmt.Errorf("rex: block %q is not placed in snapshot", id)
}
