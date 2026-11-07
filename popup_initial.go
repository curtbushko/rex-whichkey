package whichkey

import (
	"context"
	"math"

	"temp/rex"
)

const (
	// Insets are an estimate in terminal cells, not fixed font pixels. Origin
	// bounds include native chrome; a new pane's chrome/font may differ. Allow
	// two horizontal and four vertical cells plus half a row for rounding.
	// The extra two vertical cells protect root spacers without later resizing;
	// they do not hide native controls or guarantee pixel-tight initial fit.
	popupHorizontalInsetCells = 2
	popupVerticalInsetCells   = 4
	popupRounding             = 0.5

	// An unmeasured origin supplies no scale. Use a hypothetical 100x40 window
	// grid, cap startup height at 30%, and paginate rather than opening at the
	// old position maximum. These constants are fallback policy, not a report
	// about any viewer, font or actual window pixels.
	popupFallbackColumns = 100
	popupFallbackRows    = 40
	popupFallbackHeight  = 0.3
)

// popupInitialSize runs before the single creation attempt. RenderMenu is the
// authoritative natural-row producer, including both content/status spacers.
// The terminal seed uses the same estimated width and capped content rows, so
// the child never starts in an unrelated server-default 80x24 viewport.
func popupInitialSize(ctx context.Context, origin rex.Block, config Config) (rex.Bounds, rex.TerminalStartSize, error) {
	maximum := popupBounds(config.Position)
	rowUnit, columnUnit := 1.0/popupFallbackRows, 1.0/popupFallbackColumns
	heightCap := min(maximum.H, popupFallbackHeight)
	var seed rex.TerminalStartSize
	rect, size, err := origin.Geometry(ctx)
	if ctx.Err() != nil {
		return rex.Bounds{}, seed, ctx.Err()
	}
	if err == nil && size.Owner != "" && !size.ReportedAt.IsZero() &&
		size.Rows > 0 && size.Rows <= 65535 && size.Columns > 0 && size.Columns <= 65535 &&
		rect.H/float64(size.Rows) > 0 && rect.W/float64(size.Columns) > 0 {
		rowUnit = rect.H / float64(size.Rows)
		columnUnit = rect.W / float64(size.Columns)
		heightCap = maximum.H
		// Optional cell metadata belongs only to this startup estimate. It
		// never creates measurement provenance for the new terminal.
		if size.CellWidthPx > 0 && size.CellWidthPx <= 65535 && size.CellHeightPx > 0 && size.CellHeightPx <= 65535 {
			seed.CellWidthPx, seed.CellHeightPx = size.CellWidthPx, size.CellHeightPx
		}
	}
	// Bound floating-point arithmetic before any integer conversion. Geometry
	// decoding has rejected nonfinite/zero rectangles; clamp very small splits
	// and very large terminals to the documented terminal dimension limit.
	columns := popupCellCount(maximum.W/columnUnit - popupHorizontalInsetCells)
	availableRows := popupCellCount(heightCap/rowUnit - popupVerticalInsetCells - popupRounding)
	config, err = providerConfig(config, [4]*providerUpdate{})
	if err != nil {
		return rex.Bounds{}, seed, err
	}
	menu, err := NewMenu(config.Spec)
	if err != nil {
		return rex.Bounds{}, seed, err
	}
	frame := RenderMenu(menu, config, NewTheme(TerminalPalette{}, true, false), columns, 0)
	rows := max(1, min(frame.PreferredHeight, availableRows))
	seed.Columns, seed.Rows = uint32(columns), uint32(rows)
	height := min(heightCap, (float64(rows)+popupVerticalInsetCells+popupRounding)*rowUnit)
	return popupHeightBounds(maximum, config.Position, height), seed, nil
}

func popupCellCount(cells float64) int {
	if math.IsNaN(cells) || cells < 1 {
		return 1
	}
	if math.IsInf(cells, 1) || cells >= 65535 {
		return 65535
	}
	return int(math.Floor(cells))
}

// Initial bounds retain the position maximum's anchors for the whole invocation.
func popupHeightBounds(maximum rex.Bounds, position string, height float64) rex.Bounds {
	next := maximum
	next.H = height
	switch position {
	case "top":
		// Keep the original top edge.
	case "bottom":
		next.Y += maximum.H - height
	default: // Center and side popups retain their original vertical center.
		next.Y += (maximum.H - height) / 2
	}
	next.Y = min(next.Y, 1-next.H) // Clamp floating-point edge rounding.
	return next
}
