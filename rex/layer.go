package rex

import (
	"context"
	"errors"
	"fmt"
	"math"
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

// Layer owns exactly the sole floating terminal bound by DirectPopup.
// Exported IDs are snapshots: mutating them cannot retarget cleanup.
// Copies refer to the same resources; they do not grant additional ownership.
type Layer struct {
	ID       string
	BlockID  string
	WindowID string

	client  *Client
	layerID string
	blockID string
}

// Block returns the exact terminal identity captured from popup binding. Its Inspect
// and Focus methods resolve current placement instead of following WindowID.
func (l Layer) Block() Block {
	return Block{ID: l.blockID, client: l.client, blockID: l.blockID}
}

func (l Layer) valid() error {
	if err := l.Block().valid(); err != nil {
		return err
	}
	if !resourceID(l.layerID, "layer") {
		return errors.New("rex: layer must be bound by DirectPopup")
	}
	return nil
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
