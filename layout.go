package whichkey

// menuLayout uses terminal cells, never bytes or Unicode code-point counts.
// Heading and footer yield to at least one content row in tiny viewports.
type menuLayout struct {
	x, y, width, height    int
	preferredHeight        int
	columns, rows          int
	capacity               int
	page, pages            int
	header, help           bool
	padding, bottomPadding bool
	topSpace, bottomSpace  bool
}

const menuColumnGap = 2

// Popups keep one top blank row and put the footer on the last row.
// Standalone frames retain their outer padding at both edges.
// Tiny viewports prioritize usable content and navigation over padding.
func layoutMenu(width, height int, position string, itemWidths []int, requestedPage int, hasHeader, anchorFooter bool) menuLayout {
	paddingRows := 2
	if anchorFooter {
		paddingRows = 1
	}
	innerHeight := height
	padding := width > 0 && height >= paddingRows+2
	if padding {
		innerHeight -= paddingRows
	}
	layout := layoutMenuContent(width, innerHeight, position, itemWidths, requestedPage, hasHeader, anchorFooter)
	if width > 0 {
		layout.preferredHeight = min(layout.preferredHeight, int(^uint(0)>>1)-paddingRows) + paddingRows
	}
	if padding {
		layout.padding, layout.bottomPadding = true, !anchorFooter
		layout.height += paddingRows
	}
	if anchorFooter && width > 0 && height > 0 {
		layout.y, layout.height = 0, height
	}
	return layout
}

func layoutMenuContent(width, height int, position string, itemWidths []int, requestedPage int, hasHeader, anchorFooter bool) menuLayout {
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
	// An absent heading reserves neither a row nor a spacer. The footer and
	// its spacer remain; saturate rather than overflow near MaxInt.
	chromeRows := 2
	if hasHeader {
		chromeRows += 2
	}
	layout.preferredHeight = min(neededRows, int(^uint(0)>>1)-chromeRows) + chromeRows
	if height <= 0 {
		return layout
	}

	var chrome int
	layout.header, layout.help, layout.bottomSpace, chrome = menuStatusRows(height, hasHeader, anchorFooter)
	// Establish page capacity without spacers so separation cannot displace
	// content or change page count/Tab normalization.
	layout.rows = min(height-chrome, neededRows)
	layout.capacity = count
	if layout.rows <= count/layout.columns {
		layout.capacity = layout.rows * layout.columns
	}
	// The page indicator shares the footer: it never reserves another
	// row, so discovering multiple pages cannot itself cause pagination.
	layout.pages = 1 + (count-1)/layout.capacity
	// Match Menu.HandleKey's modulo arithmetic, so Tab advances the displayed
	// page even when a resize temporarily reduces the number of pages.
	layout.page = max(0, requestedPage) % layout.pages
	remaining := count - layout.page*layout.capacity
	layout.rows = min(layout.rows, 1+(remaining-1)/layout.columns)

	chrome = layout.addSpacers(height, chrome)

	// Rows span the viewport; horizontal anchors therefore coincide.
	layout.width = width
	layout.height = layout.rows + chrome
	layout.y = (height - layout.height) / 2
	switch position {
	case positionTop:
		layout.y = 0
	case "bottom":
		layout.y = height - layout.height
	}
	return layout
}

func (layout *menuLayout) addSpacers(height, chrome int) int {
	spare := height - chrome - layout.rows
	if !layout.bottomSpace && layout.help && spare >= 1 {
		layout.bottomSpace = true
		chrome++
		spare--
	}
	if layout.header && spare >= 1 {
		layout.topSpace = true
		chrome++
	}
	return chrome
}

func menuStatusRows(height int, hasHeader, anchorFooter bool) (header, help, gap bool, rows int) {
	header, help = hasHeader && height >= 3, height >= 2
	if header {
		rows++
	}
	if help {
		rows++
	}
	// A popup reserves separation before page capacity is calculated. Tiny
	// viewports still prioritize a content row and readable navigation.
	gap = anchorFooter && help && height-rows >= 2
	if gap {
		rows++
	}
	return header, help, gap, rows
}
