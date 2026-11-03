// Package whichkey defines the menu's shared configuration and mapping data.
// Executable configuration and callbacks remain in the originating Rex Lua runtime.
package whichkey

import (
	"bytes"
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"
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

	// Native-only presentation, sent through a separate private file. It is
	// deliberately absent from the canonical Config/Lua wire schema.
	textEntry *textEntrySpec
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
	// Provider is absent or "known-hosts": an inventory-only group expanded by
	// the terminal owner before NewMenu, never an executable target.
	Provider string `json:"provider,omitempty"`
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
	providers := make(map[string]bool)
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
		if mapping.Provider != "" {
			providers[canonical] = true
		}
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
		if providers[canonical] && prefixes[canonical] {
			return fmt.Errorf("whichkey: provider %q reserves its subtree; explicit children are not allowed", canonical)
		}
	}
	return nil
}

func (m Mapping) validate() error {
	if m.Action != "" && m.CallbackRef != "" {
		return errors.New("action and callback_ref are mutually exclusive")
	}
	actionable := m.Action != "" || m.CallbackRef != ""
	if m.Provider != "" {
		if m.Provider != "known-hosts" {
			return errors.New("unknown mapping provider")
		}
		if actionable || m.Group == nil || len(m.Args) != 0 || m.Confirm {
			return errors.New("known-hosts provider requires a group without action, callback_ref, args or confirm")
		}
	}
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

// CompletionVersion identifies the data-only post-return bridge format.
const CompletionVersion = 1

// Completion kinds select one retained action or release the invocation.
const (
	CompletionComplete = "complete"
	CompletionCancel   = "cancel"
)

// InvocationLifetimeSeconds and MaxPendingInvocations bound the originating Lua
// registry. Expiry starts at reservation, not at completion or the last redraw.
const (
	InvocationLifetimeSeconds = 30
	MaxPendingInvocations     = 64
)

// Completion carries only an invocation capability and a selection or cancellation.
// Origin, actions, arguments, callbacks and configuration identity stay retained
// in Rex Lua. Validation is not authentication or proof of a live origin.
type Completion struct {
	Version    int                  `json:"version"`
	Capability string               `json:"capability"`
	Kind       string               `json:"kind"`
	Selection  *CompletionSelection `json:"selection,omitempty"`

	// Private native form data; never serialized to the original Lua bridge.
	textInput *string
}

// CompletionSelection identifies a retained mapping. Confirmed records y at that
// mapping's prompt; it never overrides the retained mapping's confirmation flag.
type CompletionSelection struct {
	Keys        string `json:"keys"`
	CallbackRef string `json:"callback_ref,omitempty"`
	Confirmed   bool   `json:"confirmed"`
}

// NewInvocationCapability returns 256 unpredictable bits as 64 lowercase hex
// characters. RNG failure must abort invocation creation; there is no fallback.
func NewInvocationCapability() (string, error) {
	var data [32]byte
	if _, err := rand.Read(data[:]); err != nil {
		return "", fmt.Errorf("whichkey: invocation capability randomness: %w", err)
	}
	return hex.EncodeToString(data[:]), nil
}

// ValidateInvocationCapability checks the canonical representation, not authority.
func ValidateInvocationCapability(capability string) error {
	if len(capability) != 64 {
		return errors.New("whichkey: capability must be 64 lowercase hex characters")
	}
	for _, c := range capability {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
			return errors.New("whichkey: capability must be 64 lowercase hex characters")
		}
	}
	return nil
}

// InvocationCapabilityMatches compares canonical capabilities in constant time.
// The expected value must come from trusted retained invocation state.
func InvocationCapabilityMatches(presented, expected string) bool {
	if ValidateInvocationCapability(presented) != nil || ValidateInvocationCapability(expected) != nil {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(presented), []byte(expected)) == 1
}

// Validate checks message structure. It does not check the Lua registry, expiry,
// origin, replay state, connection or configuration generation.
func (c Completion) Validate() error {
	if c.Version != CompletionVersion {
		return errors.New("whichkey: unsupported completion version")
	}
	if err := ValidateInvocationCapability(c.Capability); err != nil {
		return err
	}
	switch c.Kind {
	case CompletionCancel:
		if c.Selection != nil {
			return errors.New("whichkey: cancellation must not carry a selection")
		}
	case CompletionComplete:
		if c.Selection == nil {
			return errors.New("whichkey: completion requires a selection")
		}
		canonical, err := NormalizeKeys(c.Selection.Keys)
		if err != nil || canonical != c.Selection.Keys {
			return errors.New("whichkey: completion keys must be a canonical full sequence")
		}
		if err := validateText(c.Selection.CallbackRef); err != nil {
			return errors.New("whichkey: invalid completion callback reference")
		}
		if c.Selection.CallbackRef != "" && strings.TrimSpace(c.Selection.CallbackRef) == "" {
			return errors.New("whichkey: completion callback reference must not be blank")
		}
	default:
		return errors.New("whichkey: invalid completion kind")
	}
	return nil
}

// ValidateSelection checks identity and confirmation against a retained, evaluated
// snapshot. Hidden actions remain selectable. Lua must still resolve its OWN
// retained snapshot and recheck conditions, authority and origin before dispatch.
// A valid cancellation needs no mapping and dispatches nothing.
func (c Completion) ValidateSelection(mappings []Mapping) error {
	if err := c.Validate(); err != nil {
		return err
	}
	if c.Kind == CompletionCancel {
		return nil
	}
	if err := ValidateMappings(mappings); err != nil {
		return err
	}
	for _, mapping := range mappings {
		if !mapping.Enabled() || (mapping.Action == "" && mapping.CallbackRef == "") {
			continue
		}
		canonical, err := NormalizeKeys(mapping.Keys)
		if err != nil {
			return err
		}
		if canonical != c.Selection.Keys {
			continue
		}
		if mapping.CallbackRef != c.Selection.CallbackRef {
			return errors.New("whichkey: completion callback reference does not match selection")
		}
		if mapping.Confirm != c.Selection.Confirmed {
			return errors.New("whichkey: completion confirmation does not match retained requirement")
		}
		return nil
	}
	return errors.New("whichkey: completion does not select an enabled retained action")
}

// MarshalJSON emits a validated data-only message. Capability-bearing output is
// secret transport data, never a log or a user-visible diagnostic.
func (c Completion) MarshalJSON() ([]byte, error) {
	if c.textInput != nil {
		return nil, errors.New("whichkey: native form data cannot use the configured completion wire format")
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	type message Completion
	return json.Marshal(message(c))
}

// UnmarshalJSON rejects duplicate/unknown fields, nulls, missing required fields,
// invalid UTF-8/surrogates, coercions and trailing values, including in selection.
// The receiver is unchanged on failure. Registry authentication remains mandatory.
func (c *Completion) UnmarshalJSON(data []byte) error {
	if c == nil {
		return errors.New("whichkey: completion receiver is nil")
	}
	fields, err := completionObject(data, "version", "capability", "kind", "selection")
	if err != nil {
		return err
	}
	var candidate Completion
	if err := json.Unmarshal(fields["version"], &candidate.Version); err != nil {
		return errors.New("whichkey: completion version must be an integer")
	}
	if candidate.Capability, err = completionString(fields["capability"]); err != nil {
		return err
	}
	if candidate.Kind, err = completionString(fields["kind"]); err != nil {
		return err
	}
	if raw, present := fields["selection"]; present {
		selection, err := completionObject(raw, "keys", "callback_ref", "confirmed")
		if err != nil {
			return err
		}
		value := new(CompletionSelection)
		if value.Keys, err = completionString(selection["keys"]); err != nil {
			return err
		}
		if raw, present := selection["callback_ref"]; present {
			if value.CallbackRef, err = completionString(raw); err != nil {
				return err
			}
			if value.CallbackRef == "" {
				return errors.New("whichkey: omit an absent completion callback reference")
			}
		}
		if err := json.Unmarshal(selection["confirmed"], &value.Confirmed); err != nil {
			return errors.New("whichkey: completion confirmed must be an explicit boolean")
		}
		candidate.Selection = value
	}
	if err := candidate.Validate(); err != nil {
		return err
	}
	*c = candidate
	return nil
}

// Decode fields individually: decoding straight into a struct would accept case
// aliases, duplicate fields and nulls as zero values at this trust boundary.
func completionObject(data []byte, allowed ...string) (map[string]json.RawMessage, error) {
	invalid := errors.New("whichkey: completion requires a strict JSON object with known non-null fields")
	if !utf8.Valid(data) {
		return nil, invalid
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	token, err := decoder.Token()
	if err != nil || token != json.Delim('{') {
		return nil, invalid
	}
	fields := make(map[string]json.RawMessage)
	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			return nil, invalid
		}
		name, ok := token.(string)
		if !ok || !slices.Contains(allowed, name) {
			return nil, invalid
		}
		if _, duplicate := fields[name]; duplicate {
			return nil, invalid
		}
		var raw json.RawMessage
		if err := decoder.Decode(&raw); err != nil || bytes.Equal(bytes.TrimSpace(raw), []byte("null")) {
			return nil, invalid
		}
		fields[name] = raw
	}
	if token, err := decoder.Token(); err != nil || token != json.Delim('}') {
		return nil, invalid
	}
	if _, err := decoder.Token(); err != io.EOF {
		return nil, invalid
	}
	return fields, nil
}

func completionString(data []byte) (string, error) {
	data = bytes.TrimSpace(data)
	invalid := errors.New("whichkey: completion string must be present and valid Unicode JSON")
	var value string
	if len(data) == 0 || data[0] != '"' || json.Unmarshal(data, &value) != nil {
		return "", invalid
	}
	// encoding/json replaces unpaired UTF-16 surrogates with U+FFFD. Reject
	// those instead of accidentally selecting a different, valid U+FFFD key.
	for i := 1; i < len(data)-1; i++ {
		if data[i] != '\\' {
			continue
		}
		i++
		if data[i] != 'u' {
			continue
		}
		r, err := strconv.ParseUint(string(data[i+1:i+5]), 16, 16)
		if err != nil {
			return "", invalid
		}
		i += 4
		if r >= 0xDC00 && r <= 0xDFFF {
			return "", invalid
		}
		if r >= 0xD800 && r <= 0xDBFF {
			if i+6 >= len(data) || data[i+1] != '\\' || data[i+2] != 'u' {
				return "", invalid
			}
			low, err := strconv.ParseUint(string(data[i+3:i+7]), 16, 16)
			if err != nil || low < 0xDC00 || low > 0xDFFF {
				return "", invalid
			}
			i += 6
		}
	}
	return value, nil
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
