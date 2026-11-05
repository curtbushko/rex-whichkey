package whichkey

import (
	"errors"
	"strings"

	"github.com/rivo/uniseg"
)

const renameTextLimit = 4096

// textEntrySpec is private popup presentation data, not a Mapping option
// or callback argument. The parent retains the only executable target identity.
type textEntrySpec struct {
	Initial        string `json:"initial"`
	SessionInitial string `json:"session_initial"`
}

func nativeRenameMapping(mapping Mapping) bool {
	if len(mapping.Args) != 0 {
		if _, err := completionObject(mapping.Args); err != nil {
			return false // Custom retained arguments keep their ordinary action behavior.
		}
	}
	return mapping.CallbackRef == "" && !mapping.Confirm &&
		(mapping.Keys == "tr" && mapping.Action == "client.tab.rename" ||
			mapping.Keys == "wr" && mapping.Action == "session.rename")
}

type menuTextEntry struct {
	node     *menuNode
	value    string
	cursor   int // UTF-8 byte offset at a grapheme boundary
	selected bool
	problem  string
}

func validateRenameText(value string) error {
	if len(value) > renameTextLimit {
		return errors.New("Name exceeds 4096 bytes")
	}
	if err := validateText(value); err != nil {
		return errors.New("Name must be valid text without control characters")
	}
	if strings.TrimSpace(value) == "" {
		return errors.New("Name cannot be empty")
	}
	return nil
}

func (m *Menu) openTextEntry(node *menuNode) {
	value := m.textSpec.Initial
	if node.mapping.Keys == "wr" {
		value = m.textSpec.SessionInitial
	}
	problem := ""
	if len(value) > renameTextLimit || validateText(value) != nil {
		// Never silently truncate or rewrite the existing name on submission.
		value, problem = "", "Enter a new name"
	}
	m.entry = &menuTextEntry{node: node, value: value, cursor: len(value), selected: true, problem: problem}
}

func (e *menuTextEntry) title() string {
	if e.node.mapping.Keys == "wr" {
		return "Rename session"
	}
	return "Rename tab"
}

func (m *Menu) handleTextKey(key string) MenuEvent {
	e := m.entry
	switch key {
	case "\x1b":
		m.entry = nil // Keep the Tabs node and its page; no result is committed.
		return MenuEvent{}
	case "\r", "\n":
		if err := validateRenameText(e.value); err != nil {
			e.problem = err.Error()
			return MenuEvent{}
		}
		event := m.selectNode(e.node, false)
		value := e.value
		event.Selection.textInput = &value
		return event
	case "\x15": // Ctrl-U
		e.value, e.cursor, e.selected, e.problem = "", 0, false, ""
	case "\x1b[H":
		e.cursor, e.selected = 0, false
	case "\x1b[F":
		e.cursor, e.selected = len(e.value), false
	case "\x1b[D":
		if e.selected {
			e.cursor = 0
		} else {
			e.cursor = previousTextBoundary(e.value, e.cursor)
		}
		e.selected = false
	case "\x1b[C":
		if e.selected {
			e.cursor = len(e.value)
		} else {
			e.cursor = nextTextBoundary(e.value, e.cursor)
		}
		e.selected = false
	case "\b", "\x7f", "\x1b[3~":
		if e.selected {
			e.value, e.cursor = "", 0
		} else if key == "\x1b[3~" {
			end := nextTextBoundary(e.value, e.cursor)
			e.value = e.value[:e.cursor] + e.value[end:]
		} else {
			start := previousTextBoundary(e.value, e.cursor)
			e.value = e.value[:start] + e.value[e.cursor:]
			e.cursor = start
		}
		if e.cursor > 0 {
			e.cursor = nextTextBoundary(e.value, e.cursor-1)
		}
		e.selected, e.problem = false, ""
	default:
		if key != "\t" {
			m.insertText(key)
		}
	}
	return MenuEvent{}
}

// insertText is also the paste boundary. Pasted Enter/Escape/control bytes are
// rejected as data; they can never submit the form or fall through as mappings.
func (m *Menu) insertText(text string) {
	if m.entry == nil || text == "" {
		return
	}
	e := m.entry
	value, cursor := e.value, e.cursor
	if e.selected {
		value, cursor = "", 0
	}
	if len(text) > renameTextLimit || len(value)+len(text) > renameTextLimit {
		e.problem = "Name exceeds 4096 bytes"
		return
	}
	if validateText(text) != nil {
		e.problem = "Name must be valid text without control characters"
		return
	}
	e.value = value[:cursor] + text + value[cursor:]
	e.cursor = cursor + len(text)
	// Combining/ZWJ input can merge with following text. Advance to a whole
	// grapheme boundary rather than leaving the next edit inside a cluster.
	if e.cursor > 0 {
		e.cursor = nextTextBoundary(e.value, e.cursor-1)
	}
	e.selected, e.problem = false, ""
}

func previousTextBoundary(value string, cursor int) int {
	previous := 0
	clusters := uniseg.NewGraphemes(value)
	for clusters.Next() {
		_, end := clusters.Positions()
		if end >= cursor {
			return previous
		}
		previous = end
	}
	return previous
}

func nextTextBoundary(value string, cursor int) int {
	clusters := uniseg.NewGraphemes(value)
	for clusters.Next() {
		_, end := clusters.Positions()
		if end > cursor {
			return end
		}
	}
	return len(value)
}

func (m *Menu) textEntryLine(theme Theme, width int) menuLine {
	e := m.entry
	label, role := "Name: ", ThemePrimary
	if e.problem != "" {
		label, role = "Name (!): ", ThemeCritical
	}
	if width <= uniseg.StringWidth(label) {
		label = "" // Extremely narrow views still prioritize the edit point.
	}
	left, right := e.value[:e.cursor], e.value[e.cursor:]
	budget := max(1, width-uniseg.StringWidth(label))
	// Reserve a visible software caret; the TUI already owns/hides the real
	// cursor. Scroll whole graphemes horizontally so the edit point stays visible.
	leftWidth, start := uniseg.StringWidth(left), 0
	clusters := uniseg.NewGraphemes(left)
	for leftWidth+1 > budget && clusters.Next() {
		_, start = clusters.Positions()
		leftWidth -= uniseg.StringWidth(clusters.Str())
	}
	left = left[start:]
	remaining, end := max(0, budget-leftWidth-1), 0
	clusters = uniseg.NewGraphemes(right)
	for clusters.Next() {
		cells := uniseg.StringWidth(clusters.Str())
		if cells > remaining {
			break
		}
		remaining -= cells
		_, end = clusters.Positions()
	}
	right = right[:end] // Avoid general clipping replacing a right-edge caret.
	valueRole := ThemePrimary
	if e.selected {
		valueRole = ThemeKey
	}
	return menuLine{menuText(role, label), menuText(valueRole, left),
		menuText(ThemeKey, theme.Decoration("│", "|")), menuText(valueRole, right)}
}
