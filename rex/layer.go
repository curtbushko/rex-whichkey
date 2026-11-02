package rex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"
)

// Bounds is a floating rectangle in normalized window coordinates. The entire
// rectangle must fit in the window; zero sizes and nonfinite values are invalid.
type Bounds struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W float64 `json:"w"`
	H float64 `json:"h"`
}

func (b Bounds) valid() error {
	for _, value := range []float64{b.X, b.Y, b.W, b.H} {
		if math.IsNaN(value) || math.IsInf(value, 0) || value < 0 || value > 1 {
			return errors.New("rex: layer bounds must be finite normalized coordinates")
		}
	}
	if b.W == 0 || b.H == 0 || b.X+b.W > 1 || b.Y+b.H > 1 {
		return errors.New("rex: layer bounds must have positive size and fit within the window")
	}
	return nil
}

// Layer owns exactly the floating layer and terminal block returned by
// NewLayer. Exported IDs are snapshots: mutating them cannot retarget cleanup.
// Copies refer to the same resources; they do not grant additional ownership.
type Layer struct {
	ID       string
	BlockID  string
	WindowID string

	client  *Client
	layerID string
	blockID string
}

// NewLayer opens one owned floating terminal in the origin's current window.
// command is a program and its arguments on the selected server, not paths to
// resolve on the graphical client. The terminal runs directly without a shell
// and remains open after child exit so Wait can recover its status; call Close
// for dismissal/error cleanup. No user configuration is read or changed here.
//
// Creation is attempted once. Errors may follow successful creation; any known
// returned identity is preserved. Never blindly retry after a timeout or decode
// error. The placement lookup and creation are separate server operations, so
// concurrent moves/closure can still race with creation.
func (b Block) NewLayer(ctx context.Context, bounds Bounds, command []string) (Layer, error) {
	if err := b.valid(); err != nil {
		return Layer{}, err
	}
	if err := bounds.valid(); err != nil {
		return Layer{}, err
	}
	if len(command) == 0 || strings.TrimSpace(command[0]) == "" {
		return Layer{}, errors.New("rex: layer command requires a program on the selected server")
	}
	for _, argument := range command {
		if strings.ContainsRune(argument, '\x00') {
			return Layer{}, errors.New("rex: layer command must not contain NUL")
		}
	}
	current, err := b.Inspect(ctx)
	if err != nil {
		return Layer{}, err
	}
	// The captured session.new_layer schema explicitly defines layout.block
	// and the shell creator's command/shell/exit options. Call that API directly;
	// popup creation needs no script VM or configuration action registration.
	arguments := map[string]any{
		"session_id": b.client.session,
		"window_id":  current.WindowID,
		"bounds":     bounds,
		"focus":      true,
		"layout": map[string]any{"block": map[string]any{
			"flavor": "com.superlogical.terminal.shell",
			"options": map[string]any{
				"command": command,
				"shell":   "none",
				"exit":    map[string]bool{"on_completion": false},
			},
		}},
	}
	var result struct {
		ID       string   `json:"layer_id"`
		BlockIDs []string `json:"block_ids"`
	}
	if err := b.client.Call(ctx, "session.new_layer", arguments, &result); err != nil {
		return Layer{}, err
	}
	if !resourceID(result.ID, "layer") {
		return Layer{}, errors.New("rex new_layer: missing or invalid layer_id (creation may have succeeded)")
	}
	layer := Layer{ID: result.ID, client: b.client, layerID: result.ID}
	if len(result.BlockIDs) != 1 || !resourceID(result.BlockIDs[0], "block") || result.BlockIDs[0] == b.blockID {
		return layer, errors.New("rex new_layer: invalid owned block_ids (creation may have succeeded)")
	}
	layer.BlockID, layer.blockID = result.BlockIDs[0], result.BlockIDs[0]
	layer.WindowID = current.WindowID
	return layer, nil
}

// Block returns the exact terminal identity captured from creation. Its Inspect
// and Focus methods resolve current placement instead of following WindowID.
func (l Layer) Block() Block {
	return Block{ID: l.blockID, client: l.client, blockID: l.blockID}
}

// Resize changes only the owned floating layer, after checking it still contains
// exactly the returned terminal. A vanished layer is an error, not a request to
// create a replacement. The server remains responsible for the final mutation.
func (l Layer) Resize(ctx context.Context, bounds Bounds) error {
	if err := bounds.valid(); err != nil {
		return err
	}
	if err := l.checkOwned(ctx); err != nil {
		return err
	}
	return l.client.Call(ctx, "session.resize_layer", struct {
		SessionID string `json:"session_id"`
		LayerID   string `json:"layer_id"`
		Bounds    Bounds `json:"bounds"`
	}{l.client.session, l.layerID, bounds}, nil)
}

// Close closes only the captured terminal, even if it has moved. Rex's
// session.close_layer also closes every block in the layer, so a prior membership
// snapshot cannot make that operation safe against concurrent layout changes.
// Without a documented atomic ownership condition, Close never calls it.
// After closing the terminal, Close checks whether the returned layer is gone.
// A remaining layer or failed inspection returns an error without further
// mutation; no window/session, label match or replacement popup is targeted.
func (l Layer) Close(ctx context.Context) error {
	if err := l.valid(); err != nil {
		return err
	}
	if err := l.Block().Close(ctx); err != nil {
		return err
	}
	layers, err := l.Block().Session().placedLayers(ctx)
	if err != nil {
		return fmt.Errorf("rex: closed owned block %q but could not confirm layer removal: %w", l.blockID, err)
	}
	for _, layer := range layers {
		if layer.id == l.layerID {
			return fmt.Errorf("rex: closed owned block %q but layer %q remains; refusing cascading layer cleanup", l.blockID, l.layerID)
		}
	}
	return nil
}

// Wait reports the owned terminal's completion without creating, closing or
// replacing anything. Use Close separately even when the child exits nonzero.
func (l Layer) Wait(ctx context.Context) error {
	if err := l.valid(); err != nil {
		return err
	}
	return l.Block().Wait(ctx)
}

func (l Layer) valid() error {
	if err := l.Block().valid(); err != nil {
		return err
	}
	if !resourceID(l.layerID, "layer") {
		return errors.New("rex: layer must be returned by NewLayer")
	}
	return nil
}

// Bounds reads the current normalized rectangle after the same exact-owned
// membership check used by Resize. A missing or malformed rectangle is never
// substituted with a cached creation size.
func (l Layer) Bounds(ctx context.Context) (Bounds, error) {
	layer, err := l.owned(ctx)
	if err != nil {
		return Bounds{}, err
	}
	var coordinates []*float64
	if err := json.Unmarshal(layer.bounds, &coordinates); err != nil {
		var object struct {
			X *float64 `json:"x"`
			Y *float64 `json:"y"`
			W *float64 `json:"w"`
			H *float64 `json:"h"`
		}
		if err := json.Unmarshal(layer.bounds, &object); err != nil {
			return Bounds{}, errors.New("rex: owned layer has malformed bounds")
		}
		coordinates = []*float64{object.X, object.Y, object.W, object.H}
	}
	if len(coordinates) != 4 {
		return Bounds{}, errors.New("rex: owned layer has missing bounds")
	}
	for _, value := range coordinates {
		if value == nil {
			return Bounds{}, errors.New("rex: owned layer has incomplete bounds")
		}
	}
	bounds := Bounds{X: *coordinates[0], Y: *coordinates[1], W: *coordinates[2], H: *coordinates[3]}
	if err := bounds.valid(); err != nil {
		return Bounds{}, err
	}
	return bounds, nil
}

func (l Layer) checkOwned(ctx context.Context) error {
	_, err := l.owned(ctx)
	return err
}

func (l Layer) owned(ctx context.Context) (placedLayer, error) {
	if err := l.valid(); err != nil {
		return placedLayer{}, err
	}
	layers, err := l.Block().Session().placedLayers(ctx)
	if err != nil {
		return placedLayer{}, err
	}
	for _, layer := range layers {
		if layer.id == l.layerID {
			if layer.kind != "floating" || len(layer.blocks) != 1 || layer.blocks[0].blockID != l.blockID {
				return placedLayer{}, fmt.Errorf("rex: layer %q no longer contains exactly its owned terminal; refusing mutation", l.layerID)
			}
			return layer, nil
		}
	}
	return placedLayer{}, fmt.Errorf("rex: owned layer %q is no longer in session %q", l.layerID, l.client.session)
}
