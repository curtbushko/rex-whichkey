//go:build darwin || linux

package whichkey

import (
	"context"
	"math"
	"time"
)

type popupGeometrySample struct {
	height                   float64
	columns, rows, preferred int
}

type popupGeometryRecorder struct {
	save      func(float64)
	saved     popupGeometrySample
	nextCheck time.Time
}

// Cache the settled root, not whichever submenu happened to be open on exit.
// An ownerless bootstrap grid or a pending/failed resize cannot seed startup.
func (recorder *popupGeometryRecorder) remember(ctx context.Context, sizing popupSizing, demand popupSizingDemand,
	columns, rows int, presented func(context.Context, int, int) (bool, error),
) {
	if recorder.save == nil || !demand.root || demand.rows <= 0 || sizing.waiting || columns <= 0 || rows <= 0 {
		return
	}
	fits := rows >= demand.rows && rows-demand.rows <= 1
	atCap := math.Abs(sizing.bounds.H-sizing.maximum.H) < 1e-9
	if !fits && !atCap {
		return
	}
	sample := popupGeometrySample{height: sizing.bounds.H, columns: columns, rows: rows, preferred: demand.rows}
	if sample == recorder.saved || time.Now().Before(recorder.nextCheck) {
		return
	}
	recorder.nextCheck = time.Now().Add(250 * time.Millisecond)
	ready, err := presented(ctx, columns, rows)
	if err != nil || !ready || ctx.Err() != nil {
		return // Cache acquisition is optional, read-only and never selection authority.
	}
	recorder.save(sample.height)
	recorder.saved = sample
}
