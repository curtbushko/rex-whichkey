package whichkey

// menuLayout uses terminal cells, never bytes or Unicode code-point counts.
// Status borders yield to at least one content row, even in a one-line viewport.
type menuLayout struct {
	x, y, width, height int
	preferredHeight     int
	columns, rows       int
	capacity            int
	page, pages         int
	header, help        bool
}

const menuColumnGap = 2

func layoutMenu(width, height int, position string, itemWidths []int, requestedPage int) menuLayout {
	layout := menuLayout{pages: 1}
	if width <= 0 {
		return layout
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
	// Natural height always includes both status borders and no blank padding.
	// Saturate rather than overflow if the content row count approaches MaxInt.
	layout.preferredHeight = min(neededRows, int(^uint(0)>>1)-2) + 2
	if height <= 0 {
		return layout
	}

	layout.help = height >= 2
	layout.header = height >= 3
	chrome := 0
	if layout.help {
		chrome++
	}
	if layout.header {
		chrome++
	}
	layout.rows = min(height-chrome, neededRows)
	layout.capacity = count
	if layout.rows <= count/layout.columns {
		layout.capacity = layout.rows * layout.columns
	}
	// The page indicator shares the bottom border: it never reserves another
	// row, so discovering multiple pages cannot itself cause pagination.
	layout.pages = 1 + (count-1)/layout.capacity
	// Match Menu.HandleKey's modulo arithmetic, so Tab advances the displayed
	// page even when a resize temporarily reduces the number of pages.
	layout.page = max(0, requestedPage) % layout.pages
	remaining := count - layout.page*layout.capacity
	layout.rows = min(layout.rows, 1+(remaining-1)/layout.columns)

	// Status borders span the viewport; horizontal anchors therefore coincide.
	layout.width = width
	layout.height = layout.rows + chrome
	layout.y = (height - layout.height) / 2
	switch position {
	case "top":
		layout.y = 0
	case "bottom":
		layout.y = height - layout.height
	}
	return layout
}
