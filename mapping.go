// Package whichkey defines the menu's shared configuration and mapping data.
// Executable configuration and callbacks remain in the originating Rex Lua runtime.
package whichkey

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Icons supplies the breadcrumb separator and submenu marker.
// An explicitly empty string suppresses that marker.
type Icons struct {
	Breadcrumb string `json:"breadcrumb"`
	Group      string `json:"group"`
}

// Config is a resolved, data-only invocation configuration. Spec order determines
// first-seen sibling order, including implicit prefixes. See docs/mapping-format.md.
type Config struct {
	Position string    `json:"position"`
	Title    string    `json:"title"`
	Icon     string    `json:"icon"`
	Icons    Icons     `json:"icons"`
	Spec     []Mapping `json:"spec"`
}

// ConfigOptions preserves absence separately from explicit empty/false values.
// Each present field replaces the corresponding default; Spec and Icons replace
// whole values, rather than merging entries or individual markers.
type ConfigOptions struct {
	Position *string    `json:"position,omitempty"`
	Title    *string    `json:"title,omitempty"`
	Icon     *string    `json:"icon,omitempty"`
	Icons    *Icons     `json:"icons,omitempty"`
	Spec     *[]Mapping `json:"spec,omitempty"`
}

// Mapping is one evaluated, inherited, full-sequence record, not executable Lua.
// The Lua compiler coalesces repeated metadata declarations into one record per
// normalized sequence. Missing prefixes are implicit groups. A nil Group means
// no explicit group; a pointer to "" is still an explicit group declaration.
type Mapping struct {
	Keys        string          `json:"keys"`
	Desc        string          `json:"desc,omitempty"`
	Group       *string         `json:"group,omitempty"`
	Icon        string          `json:"icon,omitempty"`
	Action      string          `json:"action,omitempty"`
	CallbackRef string          `json:"callback_ref,omitempty"`
	Args        json.RawMessage `json:"args,omitempty"`
	Hidden      bool            `json:"hidden,omitempty"`
	Confirm     bool            `json:"confirm,omitempty"`
	// Cond is an evaluated boolean, never a function. Nil means absent/enabled.
	// A Lua condition function returning nil is disabled: encode it as a pointer
	// to false (or omit the disabled record), never as Cond == nil.
	Cond *bool `json:"cond,omitempty"`
}

// Enabled reports eligibility for both display and key dispatch. Hidden only
// suppresses display; consumers must not remove hidden mappings from key lookup.
func (m Mapping) Enabled() bool { return m.Cond == nil || *m.Cond }

// DefaultConfig returns the baseline display defaults. The built-in-action owner
// supplies Spec before resolving user options; this function invents no actions.
func DefaultConfig() Config {
	return Config{
		Position: "center",
		Title:    "Rex",
		Icon:     "",
		Icons:    Icons{Breadcrumb: "»", Group: "+"},
	}
}

// ResolveConfig applies imported options, then inline options, over defaults.
// It validates the complete replacement before returning an independent copy.
// An error must leave the caller's last working configuration installed.
func ResolveConfig(defaults Config, imported, inline *ConfigOptions) (Config, error) {
	config := defaults
	for _, options := range []*ConfigOptions{imported, inline} {
		if options == nil {
			continue
		}
		if options.Position != nil {
			config.Position = *options.Position
		}
		if options.Title != nil {
			config.Title = *options.Title
		}
		if options.Icon != nil {
			config.Icon = *options.Icon
		}
		if options.Icons != nil {
			config.Icons = *options.Icons
		}
		if options.Spec != nil {
			config.Spec = *options.Spec
		}
	}
	if err := config.Validate(); err != nil {
		return Config{}, err
	}
	config.Spec = slices.Clone(config.Spec)
	for i := range config.Spec {
		mapping := &config.Spec[i]
		mapping.Args = slices.Clone(mapping.Args)
		if mapping.Group != nil {
			group := *mapping.Group
			mapping.Group = &group
		}
		if mapping.Cond != nil {
			cond := *mapping.Cond
			mapping.Cond = &cond
		}
	}
	return config, nil
}

// Validate rejects invalid resolved values. Loading Lua, executing functions,
// allocating callback references and checking action-specific args are not Go
// format validation and must be performed by their owning Rex-runtime consumers.
func (c Config) Validate() error {
	switch c.Position {
	case "top", "bottom", "left", "right", "center":
	default:
		return fmt.Errorf("whichkey: invalid position %q", c.Position)
	}
	for _, text := range []struct{ name, value string }{
		{"title", c.Title}, {"icon", c.Icon},
		{"icons.breadcrumb", c.Icons.Breadcrumb}, {"icons.group", c.Icons.Group},
	} {
		if err := validateText(text.value); err != nil {
			return fmt.Errorf("whichkey: %s: %w", text.name, err)
		}
	}
	return ValidateMappings(c.Spec)
}

// ParseKeys splits a full sequence into literal UTF-8 input tokens. Space and
// less-than notation is case-insensitive; every other angle-bracket notation is
// rejected. Esc, Backspace, Tab and other control characters are not mappable.
func ParseKeys(sequence string) ([]string, error) {
	if sequence == "" || !utf8.ValidString(sequence) {
		return nil, errors.New("whichkey: keys must be nonempty valid UTF-8")
	}
	var keys []string
	for len(sequence) > 0 {
		if sequence[0] == '<' {
			end := strings.IndexByte(sequence, '>')
			if end < 0 {
				return nil, errors.New("whichkey: unclosed key notation (use <lt> for literal <)")
			}
			notation := sequence[:end+1]
			switch strings.ToLower(notation) {
			case "<space>":
				keys = append(keys, " ")
			case "<lt>":
				keys = append(keys, "<")
			default:
				return nil, fmt.Errorf("whichkey: unsupported key notation %q", notation)
			}
			sequence = sequence[end+1:]
			continue
		}
		r, size := utf8.DecodeRuneInString(sequence)
		if unicode.IsControl(r) {
			return nil, fmt.Errorf("whichkey: control key U+%04X is reserved", r)
		}
		keys = append(keys, sequence[:size])
		sequence = sequence[size:]
	}
	return keys, nil
}

// NormalizeKeys returns a stable full sequence using <Space> and <lt> literals.
func NormalizeKeys(sequence string) (string, error) {
	keys, err := ParseKeys(sequence)
	if err != nil {
		return "", err
	}
	var normalized strings.Builder
	for _, key := range keys {
		normalized.WriteString(keyLiteral(key))
	}
	return normalized.String(), nil
}

// ValidateMappings checks evaluated records, canonical duplicates and all
// action/submenu conflicts, regardless of declaration order. Disabled records
// still require valid data, but do not claim active keys or create prefixes.
func ValidateMappings(mappings []Mapping) error {
	seen := make(map[string]bool)
	actions := make(map[string]bool)
	prefixes := make(map[string]bool)
	for i, mapping := range mappings {
		keys, err := ParseKeys(mapping.Keys)
		if err != nil {
			return fmt.Errorf("whichkey: spec[%d]: %w", i+1, err)
		}
		if err := mapping.validate(); err != nil {
			return fmt.Errorf("whichkey: spec[%d] %q: %w", i+1, mapping.Keys, err)
		}
		if !mapping.Enabled() {
			continue
		}
		var path strings.Builder
		for j, key := range keys {
			path.WriteString(keyLiteral(key))
			if j < len(keys)-1 {
				prefixes[path.String()] = true
			}
		}
		canonical := path.String()
		if seen[canonical] {
			return fmt.Errorf("whichkey: duplicate compiled mapping %q", canonical)
		}
		seen[canonical] = true
		if mapping.Action != "" || mapping.CallbackRef != "" {
			actions[canonical] = true
		}
	}
	// Visit source order so diagnostics do not depend on map iteration order.
	for _, mapping := range mappings {
		if !mapping.Enabled() {
			continue
		}
		canonical, err := NormalizeKeys(mapping.Keys)
		if err != nil {
			return err
		}
		if actions[canonical] && prefixes[canonical] {
			return fmt.Errorf("whichkey: mapping %q cannot be both an action and submenu", canonical)
		}
	}
	return nil
}

func (m Mapping) validate() error {
	if m.Action != "" && m.CallbackRef != "" {
		return errors.New("action and callback_ref are mutually exclusive")
	}
	actionable := m.Action != "" || m.CallbackRef != ""
	if (m.Action != "" && strings.TrimSpace(m.Action) == "") ||
		(m.CallbackRef != "" && strings.TrimSpace(m.CallbackRef) == "") {
		return errors.New("action and callback_ref must not be blank")
	}
	if actionable && (m.Group != nil || strings.TrimSpace(m.Desc) == "") {
		return errors.New("action needs a nonempty desc and cannot declare group")
	}
	if !actionable && (len(m.Args) != 0 || m.Confirm) {
		return errors.New("only actions may carry args or confirm")
	}
	for _, text := range []string{m.Desc, m.Icon, m.Action, m.CallbackRef} {
		if err := validateText(text); err != nil {
			return err
		}
	}
	if m.Group != nil {
		if err := validateText(*m.Group); err != nil {
			return fmt.Errorf("group: %w", err)
		}
	}
	if len(m.Args) != 0 {
		data := bytes.TrimSpace(m.Args)
		if !utf8.Valid(data) || !json.Valid(data) || len(data) == 0 || data[0] != '{' {
			return errors.New("args must be a JSON object, not executable Lua or null")
		}
	}
	return nil
}

func keyLiteral(key string) string {
	switch key {
	case " ":
		return "<Space>"
	case "<":
		return "<lt>"
	default:
		return key
	}
}

func validateText(text string) error {
	if !utf8.ValidString(text) {
		return errors.New("text must be valid UTF-8")
	}
	for _, r := range text {
		if unicode.IsControl(r) {
			return errors.New("text must not contain terminal control characters")
		}
	}
	return nil
}
