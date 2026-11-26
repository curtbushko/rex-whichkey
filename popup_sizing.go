package whichkey

import (
	"math"

	"temp/rex"
)

// popupSizing learns cell-to-window scale from its own PTY acknowledgements.
// A second sample accounts for native chrome, which is outside the PTY rows.
// It never uses the origin pane's dimensions or enlarges the configured cap.
type popupSizing struct {
	bounds, maximum rex.Bounds
	previous        rex.Bounds
	rows, columns   int
	preferred       int
	rowFraction     float64
	waiting         bool
	adjustments     int
}

func newPopupSizing(bounds rex.Bounds) popupSizing {
	return popupSizing{bounds: bounds, maximum: bounds}
}

func (sizing *popupSizing) fit(preferred, columns, rows int) (rex.Bounds, bool) {
	if preferred <= 0 || columns <= 0 || rows <= 0 {
		return sizing.bounds, false
	}
	newContent := preferred != sizing.preferred || columns != sizing.columns
	if newContent {
		sizing.adjustments = 0
	}
	sizing.observe(columns, rows)
	if sizing.waiting && rows == sizing.rows && columns == sizing.columns && !newContent {
		return sizing.bounds, false
	}
	sizing.waiting = false
	sizing.preferred, sizing.columns, sizing.rows = preferred, columns, rows
	// Accept one spare row for pixel/cell rounding. Bound corrections even if a
	// client imposes a minimum native size or returns inconsistent dimensions.
	if rows >= preferred && rows-preferred <= 1 || sizing.adjustments >= 6 {
		return sizing.bounds, false
	}
	if sizing.rowFraction == 0 {
		sizing.rowFraction = sizing.bounds.H / float64(rows)
	}
	height := sizing.bounds.H + (float64(preferred)+.5-float64(rows))*sizing.rowFraction
	height = min(sizing.maximum.H, max(sizing.maximum.H/65535, height))
	if math.Abs(height-sizing.bounds.H) < 1e-9 {
		return sizing.bounds, false
	}
	sizing.previous = sizing.bounds
	sizing.bounds.H, sizing.bounds.Y = height, (1-height)/2
	sizing.waiting = true
	sizing.adjustments++
	return sizing.bounds, true
}

func (sizing *popupSizing) observe(columns, rows int) {
	if sizing.waiting && columns == sizing.columns && rows != sizing.rows {
		scale := (sizing.bounds.H - sizing.previous.H) / float64(rows-sizing.rows)
		if scale > 0 && !math.IsNaN(scale) && !math.IsInf(scale, 0) {
			sizing.rowFraction = scale
		}
	} else if columns != sizing.columns || !sizing.waiting && rows != sizing.rows {
		// A screen/font change is not an acknowledgement of our last resize.
		sizing.rowFraction = 0
	}
}
