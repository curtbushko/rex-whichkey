package whichkey

// menuLayout uses terminal cells, never bytes or Unicode code-point counts.
// Chrome yields to at least one content row, even in a one-line viewport.
type menuLayout struct {
	x, y, width, height  int
	columns, rows        int
	capacity             int
	page, pages          int
	header, help, footer bool
}

const menuColumnGap = 2

func layoutMenu(width, height int, position string, itemWidths []int, headerWidth, helpWidth, footerWidth, requestedPage int) menuLayout {
	layout := menuLayout{pages: 1}
	if width <= 0 || height <= 0 {
		return layout
	}
	layout.footer = height >= 2
	layout.header = height >= 3 && headerWidth > 0
	layout.help = height >= 4 && helpWidth > 0
	chrome := 0
	for _, visible := range []bool{layout.header, layout.help, layout.footer} {
		if visible {
			chrome++
		}
	}

	cellWidth := 1
	for _, itemWidth := range itemWidths {
		cellWidth = max(cellWidth, itemWidth)
	}
	cellWidth = min(width, cellWidth)
	count := max(1, len(itemWidths))
	layout.columns = 1
	// Only form cellWidth+gap when at least two columns fit. Besides avoiding
	// a useless division, this keeps arithmetic safe for extreme resize values.
	if width > menuColumnGap && cellWidth <= (width-menuColumnGap)/2 {
		layout.columns = min(count, 1+(width-cellWidth)/(cellWidth+menuColumnGap))
	}
	neededRows := 1 + (count-1)/layout.columns
	layout.rows = min(height-chrome, neededRows)
	layout.capacity = count
	if layout.rows <= count/layout.columns {
		layout.capacity = layout.rows * layout.columns
	}
	layout.pages = 1 + (count-1)/layout.capacity
	// Match Menu.HandleKey's modulo arithmetic, so Tab advances the displayed
	// page even when a resize temporarily reduces the number of pages.
	layout.page = max(0, requestedPage) % layout.pages

	layout.width = layout.columns*cellWidth + (layout.columns-1)*menuColumnGap
	if layout.header {
		layout.width = max(layout.width, min(width, headerWidth))
	}
	if layout.help {
		layout.width = max(layout.width, min(width, helpWidth))
	}
	if layout.footer {
		layout.width = max(layout.width, min(width, footerWidth))
	}
	layout.height = layout.rows + chrome
	layout.x = (width - layout.width) / 2
	layout.y = (height - layout.height) / 2
	switch position {
	case "top":
		layout.y = 0
	case "bottom":
		layout.y = height - layout.height
	case "left":
		layout.x = 0
	case "right":
		layout.x = width - layout.width
	}
	return layout
}
