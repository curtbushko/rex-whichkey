package whichkey

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

// DecodeConfig reads one resolved, data-only configuration from the originating
// Lua runtime. It does not load modules, evaluate conditions or execute callbacks.
// The wire fields and validation belong to mapping.go, not a second Lua schema.
func DecodeConfig(input io.Reader) (Config, error) {
	const limit = 4 << 20
	data, err := io.ReadAll(io.LimitReader(input, limit+1))
	if err != nil {
		return Config{}, fmt.Errorf("whichkey: read configuration: %w", err)
	}
	if len(data) > limit {
		return Config{}, errors.New("whichkey: configuration exceeds 4 MiB")
	}
	if !utf8.Valid(data) || !json.Valid(data) {
		return Config{}, errors.New("whichkey: configuration must be one valid UTF-8 JSON value")
	}
	if err := checkConfigJSON(json.NewDecoder(bytes.NewReader(data)), false); err != nil {
		return Config{}, err
	}
	// A resolved configuration supplies every display field, including explicit
	// empty values. Do not accidentally turn missing data into display defaults.
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return Config{}, fmt.Errorf("whichkey: configuration must be an object: %w", err)
	}
	for _, name := range []string{"position", "title", "icon", "icons", "spec"} {
		if _, present := fields[name]; !present {
			return Config{}, fmt.Errorf("whichkey: resolved configuration is missing %s", name)
		}
	}
	var markers map[string]json.RawMessage
	if err := json.Unmarshal(fields["icons"], &markers); err != nil {
		return Config{}, fmt.Errorf("whichkey: icons must be an object: %w", err)
	}
	for _, name := range []string{"breadcrumb", "group"} {
		if _, present := markers[name]; !present {
			return Config{}, fmt.Errorf("whichkey: resolved icons are missing %s", name)
		}
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	var config Config
	if err := decoder.Decode(&config); err != nil {
		return Config{}, fmt.Errorf("whichkey: decode configuration: %w", err)
	}
	if err := config.Validate(); err != nil {
		return Config{}, err
	}
	return config, nil
}

// encoding/json otherwise accepts duplicate fields, null scalar coercions and
// case-insensitive struct names. Reject those before decoding into the shared
// types. Args remain arbitrary data objects with case-sensitive user keys.
func checkConfigJSON(decoder *json.Decoder, args bool) error {
	token, err := decoder.Token()
	if err != nil {
		return fmt.Errorf("whichkey: read configuration value: %w", err)
	}
	if token == nil {
		return errors.New("whichkey: configuration must not contain null")
	}
	delimiter, container := token.(json.Delim)
	if !container {
		return nil
	}
	seen := make(map[string]bool)
	for decoder.More() {
		childArgs := args
		if delimiter == '{' {
			token, err := decoder.Token()
			if err != nil {
				return fmt.Errorf("whichkey: read configuration field: %w", err)
			}
			name, ok := token.(string)
			if !ok || seen[name] || !args && name != strings.ToLower(name) {
				return errors.New("whichkey: duplicate or noncanonical configuration field")
			}
			seen[name] = true
			childArgs = args || name == "args"
		}
		if err := checkConfigJSON(decoder, childArgs); err != nil {
			return err
		}
	}
	if _, err := decoder.Token(); err != nil {
		return fmt.Errorf("whichkey: read configuration delimiter: %w", err)
	}
	return nil
}

// WriteMenuResult emits only the selection fields accepted by the retained Lua
// invocation. It never sends replacement action names or arguments. A dismissal
// is JSON null, not an empty selection or a request to execute a default action.
// The caller must dispatch only after successful terminal restoration and output.
func WriteMenuResult(output io.Writer, event MenuEvent) error {
	if event.Dismissed && event.Selection == nil {
		if err := json.NewEncoder(output).Encode(nil); err != nil {
			return fmt.Errorf("whichkey: write dismissal: %w", err)
		}
		return nil
	}
	if event.Dismissed || event.Selection == nil {
		return errors.New("whichkey: menu returned no unambiguous outcome")
	}
	mapping := event.Selection.Mapping()
	if err := ValidateMappings([]Mapping{mapping}); err != nil {
		return err
	}
	if !mapping.Enabled() || mapping.Action == "" && mapping.CallbackRef == "" || mapping.Confirm && !event.Selection.Confirmed() {
		return errors.New("whichkey: menu returned an ineligible selection")
	}
	result := struct {
		Keys        string `json:"keys"`
		CallbackRef string `json:"callback_ref,omitempty"`
		Confirmed   bool   `json:"confirmed"`
	}{mapping.Keys, mapping.CallbackRef, event.Selection.Confirmed()}
	if err := json.NewEncoder(output).Encode(result); err != nil {
		return fmt.Errorf("whichkey: write selection: %w", err)
	}
	return nil
}
