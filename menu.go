package whichkey

import (
	"fmt"
	"slices"
	"strings"
)

// MenuEntry is renderer-independent display data. Key is one literal input token;
// Keys is its canonical full sequence. Label falls back to the key's display name.
// Group and Confirm let renderers retain textual submenu and critical-action hints
// without depending on color. Display data never contains executable targets.
type MenuEntry struct {
	Key     string
	Keys    string
	Label   string
	Icon    string
	Group   bool
	Confirm bool
}

// MenuEvent reports a terminal outcome. An empty event means read the updated
// menu state and continue. Selection and Dismissed are mutually exclusive.
type MenuEvent struct {
	Selection *Selection
	Dismissed bool
	update    *renameBinding // Private publication binding, never target authority.
}

// Menu owns one invocation's evaluated mappings and navigation state. It performs
// no I/O or dispatch and must be used by a single input/event loop. Construct it
// with NewMenu; a zero Menu is inactive. Once selected or dismissed it is closed
// and emits no further selections, even if more input is queued.
type Menu struct {
	current     *menuNode
	pending     *menuNode
	pendingText *string
	page        int
	done        bool

	textSpec *textEntrySpec
	entry    *menuTextEntry
	updates  *menuRenameUpdates // Written only by the terminal UI loop.
}

type menuNode struct {
	parent   *menuNode
	children []*menuNode
	lookup   map[string]*menuNode
	key      string
	keys     string
	mapping  Mapping
}

// NewMenu validates and snapshots the Go-owned mapping format. Disabled entries
// neither claim keys nor create prefixes. First-seen sibling order is retained
// when a later explicit record decorates an already-created implicit group.
func NewMenu(mappings []Mapping) (*Menu, error) {
	if err := ValidateMappings(mappings); err != nil {
		return nil, fmt.Errorf("whichkey: create menu: %w", err)
	}
	root := &menuNode{lookup: make(map[string]*menuNode)}
	for _, mapping := range mappings {
		if !mapping.Enabled() {
			continue
		}
		keys, err := ParseKeys(mapping.Keys)
		if err != nil {
			return nil, fmt.Errorf("whichkey: create menu: %w", err)
		}
		// One immutable canonical path backs every newly-created prefix. Slicing
		// only at token boundaries avoids cumulative full-prefix allocation.
		var path strings.Builder
		for _, key := range keys {
			path.WriteString(keyLiteral(key))
		}
		canonical := path.String()
		node := root
		boundary := 0
		for _, key := range keys {
			boundary += len(keyLiteral(key))
			child := node.lookup[key]
			if child == nil {
				child = &menuNode{
					parent: node,
					lookup: make(map[string]*menuNode),
					key:    key,
					keys:   canonical[:boundary],
				}
				node.lookup[key] = child
				node.children = append(node.children, child)
			}
			node = child
		}
		node.mapping = cloneMenuMapping(mapping)
		node.mapping.Keys = node.keys
	}
	return &Menu{current: root}, nil
}

// Entries returns visible immediate children in spec order, across all pages.
// Each child's resolved Hidden value controls only that child's display;
// inheritance is already evaluated by the mapping compiler, not key prefixes.
// Renderers choose page capacity; paging never restricts which keys are active.
func (m *Menu) Entries() []MenuEntry {
	if m.current == nil {
		return nil
	}
	var entries []MenuEntry
	for _, child := range m.current.children {
		if !child.mapping.Hidden {
			entries = append(entries, child.entry())
		}
	}
	return entries
}

// Breadcrumbs returns outermost-to-innermost submenu entries, excluding the root.
// Renderers prepend Config.Title and apply Config.Icons.Breadcrumb themselves.
// No recursive tree walk or fixed depth limit is used.
func (m *Menu) Breadcrumbs() []MenuEntry {
	var trail []MenuEntry
	for node := m.current; node != nil && node.parent != nil; node = node.parent {
		trail = append(trail, node.entry())
	}
	slices.Reverse(trail)
	return trail
}

// Prefix returns the canonical current submenu sequence; the root is empty.
func (m *Menu) Prefix() string {
	if m.current == nil {
		return ""
	}
	return m.current.keys
}

// Page is the zero-based requested page. Renderers clamp it to their current page
// count after resize; pass that same count to HandleKey for Tab wraparound.
func (m *Menu) Page() int { return m.page }

// Confirmation identifies the pending action for a y/cancel prompt. While a
// prompt is active, no ordinary mapping or navigation key can dispatch an action.
func (m *Menu) Confirmation() (MenuEntry, bool) {
	if m.pending == nil {
		return MenuEntry{}, false
	}
	return m.pending.entry(), true
}

// QuitWithQ tells the footer whether an unmapped q currently dismisses the menu.
// A hidden action or implicit/explicit group on q suppresses this fallback too.
// In a confirmation prompt q only cancels that prompt.
func (m *Menu) QuitWithQ() bool {
	return !m.Done() && m.pending == nil && m.entry == nil && m.current.lookup["q"] == nil
}

// Done reports selection or dismissal, or an uninitialized menu.
func (m *Menu) Done() bool { return m.done || m.current == nil }

// HandleKey consumes one decoded literal Unicode token, or Esc (\x1b),
// Backspace (\b or \x7f), or Tab (\t). It does not decode terminal byte streams:
// the input owner must separate terminal replies, paste and escape sequences
// before calling it, never replaying their bytes as mapping/confirmation keys.
// Esc/Backspace go up one group, or dismiss at root. Unmapped q dismisses at any
// depth; mapped q follows ordinary lookup. Other unknown keys leave state alone.
// Tab advances modulo pages (at least one); entering/leaving a group or prompt
// resets the page. While prompting, exactly y selects, and every other key
// cancels without selection or fallthrough to navigation or another mapping.
// Inline text entry instead owns all keys: Esc returns to its submenu, Enter
// validates/submits, and ordinary characters edit data rather than mappings.
func (m *Menu) HandleKey(key string, pages int) MenuEvent {
	if m.Done() {
		return MenuEvent{}
	}
	if m.entry != nil {
		return m.handleTextKey(key)
	}
	if m.pending != nil {
		target := m.pending
		text := m.pendingText
		m.pending, m.pendingText = nil, nil
		m.page = 0
		if key == "y" {
			event := m.selectNode(target, true)
			event.Selection.textInput = text
			return event
		}
		return MenuEvent{}
	}
	switch key {
	case "\x1b", "\b", "\x7f":
		m.page = 0
		if m.current.parent == nil {
			m.done = true
			return MenuEvent{Dismissed: true}
		}
		m.current = m.current.parent
		return MenuEvent{}
	case "\t":
		pages = max(1, pages)
		// Reduce before incrementing to avoid overflow for any page count.
		m.page = (m.page%pages + 1) % pages
		return MenuEvent{}
	}

	target := m.current.lookup[key]
	if target == nil {
		if key == "q" {
			m.done = true
			return MenuEvent{Dismissed: true}
		}
		return MenuEvent{}
	}
	if target.mapping.Action == "" && target.mapping.CallbackRef == "" {
		m.current = target
		m.page = 0
		return MenuEvent{}
	}
	if m.textSpec != nil && inlineRenameMapping(target.mapping) {
		m.openTextEntry(target)
		return MenuEvent{}
	}
	if target.mapping.Confirm {
		m.pending = target
		m.page = 0
		return MenuEvent{}
	}
	return m.selectNode(target, false)
}

func (m *Menu) selectNode(target *menuNode, confirmed bool) MenuEvent {
	m.done = true
	event := MenuEvent{Selection: &Selection{
		mapping:   cloneMenuMapping(target.mapping),
		confirmed: confirmed,
	}}
	if m.updates != nil && m.renameReady(target) && renameKind(target.mapping) >= 0 {
		binding := m.updates.ready[renameKind(target.mapping)].Demand.binding()
		event.update = &binding
	}
	return event
}

func (n *menuNode) entry() MenuEntry {
	label := n.mapping.Desc
	if n.mapping.Group != nil && *n.mapping.Group != "" {
		label = *n.mapping.Group
	}
	if label == "" {
		label = n.key
		if label == " " {
			label = "Space"
		}
	}
	return MenuEntry{
		Key:     n.key,
		Keys:    n.keys,
		Label:   label,
		Icon:    n.mapping.Icon,
		Group:   n.mapping.Action == "" && n.mapping.CallbackRef == "",
		Confirm: n.mapping.Confirm,
	}
}
