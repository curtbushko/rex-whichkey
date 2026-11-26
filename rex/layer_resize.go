package rex

import (
	"context"
	"errors"
	"math"
)

// InitialBounds returns a copy of the popup's validated launch rectangle. Missing
// geometry leaves legacy/unsupported views usable, but disables automatic fitting.
func (l Layer) InitialBounds() (Bounds, bool) {
	if l.bounds == nil {
		return Bounds{}, false
	}
	return *l.bounds, true
}

// Resize changes only the verified sole-block float, never focus or layer order.
// Expected bounds prevent a stale sizing worker from undoing a manual move/resize.
// An uncertain mutation/readback is fatal to fitting, never a reason to retry.
func (l Layer) Resize(ctx context.Context, expected, bounds Bounds) error {
	if err := bounds.valid(); err != nil {
		return err
	}
	layer, err := l.owned(ctx)
	if err != nil {
		return err
	}
	if layer.state != "" && layer.state != layerStateNormal || layer.bounds == nil || !sameBounds(*layer.bounds, expected) {
		return errors.New("rex: owned popup geometry changed or is unavailable; resize refused")
	}
	if resizeErr := l.client.Call(ctx, "session.resize_layer", map[string]any{
		fieldSession: l.client.session, fieldLayer: l.layerID, "bounds": bounds,
	}, nil); resizeErr != nil {
		return errors.Join(errors.New("rex: popup resize failed or uncertain; not retried"), resizeErr)
	}
	layer, err = l.owned(ctx)
	if err != nil {
		return err
	}
	if layer.bounds == nil || !sameBounds(*layer.bounds, bounds) {
		return errors.New("rex: popup resize could not be verified; not retried")
	}
	return nil
}

func sameBounds(left, right Bounds) bool {
	return math.Abs(left.X-right.X) < 1e-9 && math.Abs(left.Y-right.Y) < 1e-9 &&
		math.Abs(left.W-right.W) < 1e-9 && math.Abs(left.H-right.H) < 1e-9
}
