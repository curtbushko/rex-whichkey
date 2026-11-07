package whichkey

import (
	"errors"
	"strings"
	"unicode/utf8"

	"github.com/rivo/uniseg"
)

const (
	renameTextLimit   = 4096
	terminalDeleteKey = "\x1b[3~"
)

// textEntrySpec is Go-owned inline rename presentation, not a Mapping option
// or callback argument. Direct dispatch retains the executable target identity.
type textEntrySpec struct {
	Initial        string `json:"initial"`
	SessionInitial string `json:"session_initial"`
}

func inlineRenameMapping(mapping Mapping) bool {
	if len(mapping.Args) != 0 {
		if _, err := completionObject(mapping.Args); err != nil {
			return false // Custom retained arguments keep their ordinary action behavior.
		}
	}
	return mapping.CallbackRef == "" &&
		(mapping.Action == menuTabRename || mapping.Action == menuSessionRename)
}

type menuTextEntry struct {
	node     *menuNode
	value    string
	cursor   int  // UTF-8 byte offset at a grapheme boundary
	selected bool // Replace untouched prefill on typing, not on deletion.
	problem  string
	touched  bool // Edits, cursor keys, validation or partial input forbid late prefill.
}

func validateRenameText(value string) error {
	if len(value) > renameTextLimit {
		return errors.New("name exceeds 4096 bytes")
	}
	if err := validateText(value); err != nil {
		return errors.New("name must be valid text without control characters")
	}
	if utf8.RuneCountInString(value) > 128 {
		return errors.New("name exceeds Rex's 128-character limit")
	}
	if strings.TrimSpace(value) == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}

func (m *Menu) openTextEntry(node *menuNode) {
	value := m.textSpec.Initial
	if node.mapping.Action == menuSessionRename {
		value = m.textSpec.SessionInitial
	}
	problem := ""
	if m.updates != nil {
		m.updates.demand(node)
		value = ""
		if update := m.updates.ready[renameKind(node.mapping)]; update != nil && update.State == renameStateReady {
			value, problem = update.Initial, update.Problem
		}
	}
	if len(value) > renameTextLimit || validateText(value) != nil {
		// Never silently truncate or rewrite the existing name on submission.
		value, problem = "", "Enter a new name"
	}
	m.entry = &menuTextEntry{node: node, value: value, cursor: len(value), selected: true, problem: problem}
}

func (e *menuTextEntry) title() string {
	if e.node.mapping.Action == menuSessionRename {
		return "Rename session"
	}
	return "Rename tab"
}

func (m *Menu) handleTextKey(key string) MenuEvent {
	e := m.entry
	e.touched = true // Escape discards this form; all other keys protect its state.
	switch key {
	case "\x1b":
		m.entry = nil // Keep the Tabs node and its page; no result is committed.
		return MenuEvent{}
	case "\r", "\n":
		return m.submitTextEntry(e)
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
	case "\b", "\x7f", terminalDeleteKey:
		if key == terminalDeleteKey {
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

func (m *Menu) submitTextEntry(entry *menuTextEntry) MenuEvent {
	if !m.renameReady(entry.node) {
		return MenuEvent{} // Discard pending Enter; readiness never replays input.
	}
	if err := validateRenameText(entry.value); err != nil {
		entry.problem = err.Error()
		return MenuEvent{}
	}
	value := entry.value
	if entry.node.mapping.Confirm {
		m.pending, m.pendingText, m.entry = entry.node, &value, nil
		m.page = 0
		return MenuEvent{}
	}
	event := m.selectNode(entry.node, false)
	event.Selection.textInput = &value
	return event
}

// insertText is also the paste boundary. Pasted Enter/Escape/control bytes are
// rejected as data; they can never submit the form or fall through as mappings.
func (m *Menu) insertText(text string) {
	if m.entry == nil || text == "" {
		return
	}
	e := m.entry
	e.touched = true
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
	if m.renameStatus() != "" {
		label, role = "Name (not ready): ", ThemeMuted
	}
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
