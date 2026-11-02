package whichkey

import (
	"fmt"
	"strings"

	"github.com/rivo/uniseg"
)

// MenuFrame is a constrained terminal-cell rectangle at zero-based X/Y. Lines
// contain only text and Theme.Style's optional SGR sequences, without newlines
// or cursor movement. Each line fits Width; len(Lines) is Height.
// The terminal owner places these lines without wrapping, clears the previous
// frame after resize, and passes Pages to Menu.HandleKey for Tab navigation.
// A zero-sized viewport or inactive menu returns no lines and Pages == 1.
type MenuFrame struct {
	X, Y, Width, Height int
	Page, Pages         int
	Lines               []string

	// PreferredHeight is the natural row count at the supplied viewport width:
	// all current entries (or the confirmation/empty message), both status
	// borders, and no blank padding. It is independent of viewport height,
	// page clipping and the selected page; it can exceed Height, even when
	// Height is zero. Counts saturate at the largest int. Inactive menus and
	// nonpositive widths report zero.
	PreferredHeight int
}

type menuSpan struct {
	text string
	role ThemeRole
}

type menuLine []menuSpan

// RenderMenu renders an already-resolved Config and Menu snapshot. It performs
// no I/O and does not mutate navigation, including during zero-size transitions.
// Page is normalized modulo the resized page count, matching HandleKey. Only
// whole grapheme clusters are emitted; clipping never normalizes user text.
// Theme controls all styling and optional icons/ASCII decorations. ASCII mode
// preserves user titles, keys and labels, rather than transliterating them.
func RenderMenu(menu *Menu, config Config, theme Theme, width, height int) MenuFrame {
	frame := MenuFrame{Pages: 1}
	if menu == nil || menu.Done() || width <= 0 {
		return frame
	}
	header := menuHeader(menu, config, theme)

	entries := menu.Entries()
	items := make([]menuLine, 0, max(1, len(entries)))
	for _, entry := range entries {
		items = append(items, menuEntryLine(entry, config, theme))
	}
	pending, confirming := menu.Confirmation()
	if confirming {
		// Keep confirmation keys on the content row as well as help: content
		// is the only guaranteed row in an extremely short terminal.
		items = []menuLine{{menuText(ThemeCritical, "y=yes / other=cancel: "+pending.Label)}}
	} else if len(items) == 0 {
		items = []menuLine{{menuText(ThemeMuted, "No visible mappings")}}
	}
	itemWidths := make([]int, len(items))
	for i, item := range items {
		itemWidths[i] = uniseg.StringWidth(item.plain())
	}
	layout := layoutMenu(width, height, config.Position, itemWidths, menu.Page())
	frame.PreferredHeight = layout.preferredHeight
	if height <= 0 {
		return frame
	}
	frame.X, frame.Y = layout.x, layout.y
	frame.Width, frame.Height = layout.width, layout.height
	frame.Page, frame.Pages = layout.page, layout.pages

	if layout.header {
		frame.Lines = append(frame.Lines, menuStatus(header, theme, layout.width))
	}
	cellWidth := (layout.width - (layout.columns-1)*menuColumnGap) / layout.columns
	start := layout.page * layout.capacity
	end := start + min(layout.capacity, len(items)-start)
	for row := 0; row < layout.rows; row++ {
		var line strings.Builder
		for column := 0; column < layout.columns; column++ {
			index := start + row*layout.columns
			if column >= end-index {
				break
			}
			index += column
			text, used := items[index].render(theme, cellWidth)
			line.WriteString(text)
			if column+1 < layout.columns && index+1 < end {
				line.WriteString(theme.Style(ThemePrimary, strings.Repeat(" ", cellWidth-used+menuColumnGap)))
			}
		}
		frame.Lines = append(frame.Lines, line.String())
	}
	if layout.help {
		role := ThemeMuted
		if confirming {
			role = ThemeCritical
		}
		help := menuHelp(menu, confirming, layout.page, layout.pages, layout.width)
		frame.Lines = append(frame.Lines, menuStatus(menuLine{menuText(role, help)}, theme, layout.width))
	}
	return frame
}

func menuHeader(menu *Menu, config Config, theme Theme) menuLine {
	var header menuLine
	if icon := menuDecoration(theme, config.Icon, ""); icon != "" {
		header = append(header, menuText(ThemeIcon, icon))
	}
	if config.Title != "" {
		if len(header) > 0 {
			header = append(header, menuText(ThemePrimary, " "))
		}
		header = append(header, menuText(ThemePrimary, config.Title))
	}
	separator := menuDecoration(theme, config.Icons.Breadcrumb, ">")
	for _, crumb := range menu.Breadcrumbs() {
		if len(header) > 0 {
			header = append(header, menuText(ThemeIcon, " "+separator+" "))
		}
		header = append(header, menuText(ThemeGroup, crumb.Label))
	}
	return header
}

func menuEntryLine(entry MenuEntry, config Config, theme Theme) menuLine {
	key := entry.Key
	if key == " " {
		key = "Space"
	}
	line := menuLine{menuText(ThemeKey, key), menuText(ThemeMuted, "  ")}
	if icon := menuDecoration(theme, entry.Icon, ""); icon != "" {
		line = append(line, menuText(ThemeIcon, icon+" "))
	}
	role := ThemePrimary
	if entry.Group {
		role = ThemeGroup
		if marker := menuDecoration(theme, config.Icons.Group, "+"); marker != "" {
			line = append(line, menuText(ThemeIcon, marker+" "))
		}
	} else if entry.Confirm {
		role = ThemeCritical
		// Place the textual warning before the label so a long description
		// cannot push it out of an otherwise usable cell.
		line = append(line, menuText(role, "[confirm] "))
	}
	return append(line, menuText(role, entry.Label))
}

// Choose complete hints before grapheme clipping, using the whole row so
// decoration yields before essential keys. At tiny widths prefer Escape over
// a partial page counter; confirmation always keeps a cancellation hint.
func menuHelp(menu *Menu, confirming bool, page, pages, width int) string {
	var hints []string
	if confirming {
		hints = []string{
			"y: confirm  Any other key: cancel",
			"y=yes other=cancel",
			"y=yes Esc=cancel",
			"Esc: cancel",
		}
	} else {
		navigation := "back"
		if menu.Prefix() == "" {
			navigation = "quit"
		}
		navigationHint := "Esc/Backspace: " + navigation
		full := navigationHint + "  key: select/open"
		if menu.QuitWithQ() {
			full += "  q: quit"
		}
		if pages > 1 {
			counter := fmt.Sprintf("%d/%d", page+1, pages)
			paging := counter + " Tab: next  "
			hints = []string{
				paging + full,
				paging + navigationHint,
				counter + " Tab Esc: " + navigation,
				counter + " Tab Esc",
				"Tab Esc",
			}
		} else {
			hints = []string{full, navigationHint, "Esc: " + navigation}
		}
	}
	for _, hint := range hints {
		if uniseg.StringWidth(hint) <= width {
			return hint
		}
	}
	return "Esc"
}

// Status text shares its row with a horizontal rule, not an extra border row.
// Narrow terminals spend their cells on hints rather than decorative padding.
func menuStatus(line menuLine, theme Theme, width int) string {
	if width <= 0 {
		return ""
	}
	rule := theme.Decoration("─", "-")
	if line.plain() == "" {
		return theme.Style(ThemeMuted, strings.Repeat(rule, width))
	}
	if uniseg.StringWidth(line.plain()) > width-4 {
		text, used := line.render(theme, width)
		return text + theme.Style(ThemeMuted, strings.Repeat(rule, width-used))
	}
	text, used := line.render(theme, width-4)
	return theme.Style(ThemeMuted, rule+" ") + text +
		theme.Style(ThemeMuted, " "+strings.Repeat(rule, width-used-3))
}

func menuDecoration(theme Theme, glyph, fallback string) string {
	if glyph == "" {
		return ""
	}
	return theme.Decoration(glyph, fallback)
}

func menuText(role ThemeRole, text string) menuSpan {
	// Use the theme's own plain-output sanitization before measurement, not
	// after, so an unchecked caller cannot smuggle control widths into layout.
	text = (Theme{}).Style(role, text)
	if text != "" {
		_, _, width, _ := uniseg.FirstGraphemeClusterInString(text, -1)
		if width == 0 {
			// Standalone combining input tokens remain intact, but must not
			// modify a preceding column or terminal cell outside this frame.
			text = " " + text
		}
	}
	return menuSpan{text: text, role: role}
}

func (line menuLine) plain() string {
	var text strings.Builder
	for _, span := range line {
		text.WriteString(span.text)
	}
	return text.String()
}

// Clip the complete plain line before introducing SGR. Grapheme segmentation
// therefore sees adjacent spans together, including unusual boundary characters.
func (line menuLine) render(theme Theme, width int) (string, int) {
	plain := line.plain()
	marker := ""
	limit := max(0, width)
	if uniseg.StringWidth(plain) > limit {
		marker = theme.Decoration("…", "~")
		if uniseg.StringWidth(marker) > limit {
			marker = ""
		}
		limit -= uniseg.StringWidth(marker)
	}
	state, used, end := -1, 0, 0
	for rest := plain; rest != ""; {
		cluster, next, cells, nextState := uniseg.FirstGraphemeClusterInString(rest, state)
		if cells > limit-used {
			break
		}
		used += cells
		end += len(cluster)
		rest, state = next, nextState
	}

	var output, run strings.Builder
	spanIndex, spanEnd, offset := 0, 0, 0
	role := ThemePrimary
	if len(line) > 0 {
		spanEnd = len(line[0].text)
		role = line[0].role
	}
	state = -1
	for rest := plain[:end]; rest != ""; {
		cluster, next, _, nextState := uniseg.FirstGraphemeClusterInString(rest, state)
		for spanIndex+1 < len(line) && offset >= spanEnd {
			spanIndex++
			spanEnd += len(line[spanIndex].text)
		}
		if line[spanIndex].role != role {
			output.WriteString(theme.Style(role, run.String()))
			run.Reset()
			role = line[spanIndex].role
		}
		run.WriteString(cluster)
		offset += len(cluster)
		rest, state = next, nextState
	}
	output.WriteString(theme.Style(role, run.String()))
	output.WriteString(theme.Style(ThemeMuted, marker))
	return output.String(), used + uniseg.StringWidth(marker)
}
