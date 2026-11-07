package rex

import (
	"context"
	"errors"
	"fmt"
)

// Block is a placed block in one fixed server/session. Placement fields are
// snapshots; operations use the captured block ID, never a cached window ID.
type Block struct {
	ID       string
	WindowID string
	LayerID  string

	client  *Client
	blockID string
}

// Block finds a full block ID in this session's current layout. Detached or
// missing blocks are not usable origins and return an error, without fallback.
func (s Session) Block(ctx context.Context, id string) (Block, error) {
	if !resourceID(id, "block") {
		return Block{}, errors.New("rex: a full block ID is required")
	}
	layers, err := s.placedLayers(ctx)
	if err != nil {
		return Block{}, err
	}
	for _, layer := range layers {
		for _, block := range layer.blocks {
			if block.blockID == id {
				return block, nil
			}
		}
	}
	return Block{}, fmt.Errorf("rex: block %q is not placed in session %q", id, s.client.session)
}

// Blocks returns the exact placed blocks in session/window/layer source order.
// The bounded CLI output is validated before resource identities are exposed.
func (s Session) Blocks(ctx context.Context) ([]Block, error) {
	layers, err := s.placedLayers(ctx)
	if err != nil {
		return nil, err
	}
	blocks := make([]Block, 0)
	for _, layer := range layers {
		blocks = append(blocks, layer.blocks...)
	}
	return blocks, nil
}

// Inspect refreshes placement in the captured session, including after a move.
func (b Block) Inspect(ctx context.Context) (Block, error) {
	if err := b.valid(); err != nil {
		return Block{}, err
	}
	return b.Session().Block(ctx, b.blockID)
}

// Session returns the captured session, independent of window membership.
func (b Block) Session() Session {
	if b.client == nil {
		return Session{}
	}
	return b.client.sessionResource(b.client.session, "")
}

// Server returns the captured server selector, including for remote origins.
func (b Block) Server() string { return b.Session().Server() }

// Window resolves the current window rather than trusting the snapshot.
func (b Block) Window(ctx context.Context) (Window, error) {
	current, err := b.Inspect(ctx)
	if err != nil {
		return Window{}, err
	}
	return current.Session().windowResource(current.WindowID, ""), nil
}

// Focus restores server focus to this block in its current window. The final
// call names only the stable block/session IDs, so a move after Inspect cannot
// redirect focus to the old window. It does not switch a graphical client.
func (b Block) Focus(ctx context.Context) error {
	if _, err := b.Inspect(ctx); err != nil {
		return err
	}
	return b.client.Call(ctx, "session.focus_block", map[string]string{
		"session_id": b.client.session,
		"block_id":   b.blockID,
	}, nil)
}

// Close explicitly closes only this block, even if it has moved. No window,
// session, label match, or other block is used as a cleanup substitute.
func (b Block) Close(ctx context.Context) error {
	if err := b.valid(); err != nil {
		return err
	}
	return b.client.resourceCommand(ctx, nil, "block", "close", "--session", b.client.session, "--", b.blockID)
}

func (b Block) valid() error {
	if err := b.Session().valid(); err != nil {
		return err
	}
	if !resourceID(b.blockID, "block") {
		return errors.New("rex: block must be obtained from a Session or owned Layer")
	}
	return nil
}

type placedLayer struct {
	id     string
	kind   string
	blocks []Block
}

// Read the documented session view via session inspect, extending the existing
// tab-only view with placement. Layers/blocks follow Rex's resource model; no
// search by popup label or process-local REX_WINDOW is involved.
func (s Session) placedLayers(ctx context.Context) ([]placedLayer, error) {
	if err := s.valid(); err != nil {
		return nil, err
	}
	var view sessionView
	if err := s.client.resourceCommand(ctx, &view, "session", "inspect", "--json", "--", s.client.session); err != nil {
		return nil, err
	}
	if view.ID != s.client.session || view.Windows == nil {
		return nil, errors.New("rex session inspect: unexpected session_id or missing windows")
	}
	return s.snapshotLayers(view)
}

func (s Session) snapshotLayers(view sessionView) ([]placedLayer, error) {
	layers := make([]placedLayer, 0)
	seen := make(map[string]bool)
	for _, window := range view.Windows {
		if !resourceID(window.ID, "window") || seen[window.ID] || window.Layers == nil {
			return nil, errors.New("rex session inspect: invalid window placement")
		}
		seen[window.ID] = true
		for _, item := range window.Layers {
			if !resourceID(item.ID, "layer") || seen[item.ID] || item.Blocks == nil {
				return nil, errors.New("rex session inspect: invalid layer placement")
			}
			seen[item.ID] = true
			layer := placedLayer{id: item.ID, kind: item.Kind}
			// Keep parsing optional layout metadata, but never use it for float
			// sizing or placement authority. Missing/malformed geometry remains
			// nonfatal; the exact placed identities below are authoritative.
			_, _ = layoutRectangles(item.Layout)
			for _, block := range item.Blocks {
				if !resourceID(block.ID, "block") || seen[block.ID] {
					return nil, errors.New("rex session inspect: invalid or duplicate block placement")
				}
				seen[block.ID] = true
				layer.blocks = append(layer.blocks, Block{
					ID: block.ID, WindowID: window.ID, LayerID: item.ID,
					client: s.client, blockID: block.ID,
				})
			}
			layers = append(layers, layer)
		}
	}
	return layers, nil
}
