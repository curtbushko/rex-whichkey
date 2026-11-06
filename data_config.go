package whichkey

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"unicode/utf8"
)

const menuConfigLimit = 4 << 20

// Shared vocabulary for public data and the Go-owned consumer. These are names,
// not a generic Rex action registry or permission to dispatch arbitrary actions.
const (
	menuSessionNew       = "session.new"
	menuSessionPrevious  = "session.previous"
	menuSessionNext      = "session.next"
	menuSessionSwitch    = "session.switch"
	menuSessionRename    = "session.rename"
	menuSessionClose     = "session.close"
	menuSessionSelect    = "session.select"
	menuTabNew           = "client.tab.new"
	menuTabPrevious      = "client.tab.previous"
	menuTabNext          = "client.tab.next"
	menuTabRename        = "client.tab.rename"
	menuTabClose         = "client.tab.close"
	menuPaneSplit        = "pane.split"
	menuPaneClose        = "pane.close"
	menuSidebarToggle    = "client.sidebar.toggle"
	menuPaletteToggle    = "client.palette.toggle"
	menuProviderSessions = "sessions"
	menuProviderWindows  = "windows"
	menuProviderPanes    = "panes"
	menuProviderHosts    = "known-hosts"
	menuSplitRight       = "right"
	menuSplitDown        = "down"
)

// LoadMenuConfig reads the highest-priority selected data file over the current
// built-ins. A nil explicit path means --config was absent, not explicitly empty.
// Loading never queries Rex, expands providers, executes code or writes a file.
func LoadMenuConfig(explicit *string) (Config, error) {
	defaults, err := nativeDefaults()
	if err != nil {
		return Config{}, err
	}
	path, implicit, err := menuConfigPath(explicit)
	if err != nil {
		return Config{}, err
	}
	data, err := readMenuConfigFile(path)
	if err != nil {
		if implicit && errors.Is(err, os.ErrNotExist) {
			// A dangling leaf symlink is a broken selected file, not absence.
			if _, linkErr := os.Lstat(path); errors.Is(linkErr, os.ErrNotExist) {
				return resolveMenuConfig(defaults, nil)
			}
		}
		return Config{}, fmt.Errorf("whichkey: config %q: %w", path, err)
	}
	options, err := decodeMenuOptions(data)
	if err != nil {
		return Config{}, fmt.Errorf("whichkey: config %q: %w", path, err)
	}
	return resolveMenuConfig(defaults, &options)
}

func menuConfigPath(explicit *string) (string, bool, error) {
	if explicit != nil {
		if *explicit == "" {
			return "", false, errors.New("whichkey: --config path must not be empty")
		}
		return *explicit, false, nil
	}
	if path := os.Getenv("REX_WHICHKEY_CONFIG"); path != "" {
		return path, false, nil
	}
	if root, present := os.LookupEnv("XDG_CONFIG_HOME"); present {
		if !filepath.IsAbs(root) {
			return "", false, errors.New("whichkey: XDG_CONFIG_HOME must be an absolute nonempty path")
		}
		return filepath.Join(root, "rex-whichkey", "config.json"), true, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", false, fmt.Errorf("whichkey: locate home configuration: %w", err)
	}
	return filepath.Join(home, ".config", "rex-whichkey", "config.json"), true, nil
}

func readMenuConfigFile(path string) (data []byte, outcome error) {
	// Normal user/Nix symlinks are allowed. O_NONBLOCK prevents a substituted
	// FIFO from blocking open before the opened descriptor's type is checked.
	file, err := os.OpenFile(path, os.O_RDONLY|syscall.O_NONBLOCK, 0)
	if err != nil {
		return nil, fmt.Errorf("open data file: %w", err)
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("stat data file: %w", err)
	}
	if !info.Mode().IsRegular() {
		return nil, errors.New("configuration must be a regular data file")
	}
	if info.Size() > menuConfigLimit {
		return nil, errors.New("configuration exceeds 4 MiB")
	}
	data, err = io.ReadAll(io.LimitReader(file, menuConfigLimit+1))
	if err != nil {
		return nil, fmt.Errorf("read data file: %w", err)
	}
	if len(data) > menuConfigLimit {
		return nil, errors.New("configuration exceeds 4 MiB")
	}
	return data, nil
}

func decodeMenuOptions(data []byte) (ConfigOptions, error) {
	if err := checkMenuJSON(data); err != nil {
		return ConfigOptions{}, err
	}
	fields, err := menuObject(data, "version", "position", "title", "icon", "icons", "spec")
	if err != nil {
		return ConfigOptions{}, err
	}
	// Require the canonical integer scalar, not 1.0, 1e0, a string or bool.
	if !bytes.Equal(bytes.TrimSpace(fields["version"]), []byte("1")) {
		return ConfigOptions{}, errors.New("menu configuration requires version: 1")
	}
	if raw, present := fields["icons"]; present {
		if _, checkErr := menuObject(raw, "breadcrumb", "group"); checkErr != nil {
			return ConfigOptions{}, fmt.Errorf("icons: %w", checkErr)
		}
	}
	if raw, present := fields["spec"]; present {
		var rows []json.RawMessage
		if decodeErr := json.Unmarshal(raw, &rows); decodeErr != nil {
			return ConfigOptions{}, errors.New("spec must be an array of mappings")
		}
		for index, row := range rows {
			if _, checkErr := menuObject(row, "keys", "desc", "group", "icon", "action", "args", "hidden", "confirm", "provider", "cond"); checkErr != nil {
				return ConfigOptions{}, fmt.Errorf("spec[%d]: %w", index+1, checkErr)
			}
		}
	}
	// Exact member checks above precede struct decoding; DisallowUnknownFields
	// alone would still admit Unicode case-fold aliases such as long-s in spec.
	delete(fields, "version")
	optionsJSON, err := json.Marshal(fields)
	if err != nil {
		return ConfigOptions{}, fmt.Errorf("encode menu options: %w", err)
	}
	var options ConfigOptions
	if err := json.Unmarshal(optionsJSON, &options); err != nil {
		return ConfigOptions{}, fmt.Errorf("invalid menu option type: %w", err)
	}
	return options, nil
}

func menuObject(data []byte, allowed ...string) (map[string]json.RawMessage, error) {
	fields, err := completionObject(data, allowed...)
	if err != nil {
		return nil, errors.New("expected an object with exact known, unique, non-null members")
	}
	return fields, nil
}

// checkMenuJSON rejects what encoding/json normally repairs or silently accepts.
// Object keys remain case-sensitive data here; structural allowlists are applied
// separately so arbitrary data is never lowercased into a different meaning.
func checkMenuJSON(data []byte) error {
	if len(data) > menuConfigLimit || !utf8.Valid(data) || !json.Valid(data) {
		return errors.New("menu configuration must be valid UTF-8 JSON, at most 4 MiB")
	}
	for index := 0; index < len(data); index++ {
		if data[index] != '"' {
			continue
		}
		start := index
		index++
		for index < len(data) && data[index] != '"' {
			if data[index] == '\\' {
				index++
			}
			index++
		}
		if _, err := completionString(data[start : index+1]); err != nil {
			return errors.New("menu configuration contains malformed Unicode string escapes")
		}
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	return checkMenuValue(decoder, 0)
}

func checkMenuValue(decoder *json.Decoder, depth int) error {
	if depth > 64 {
		return errors.New("menu configuration exceeds 64 levels of JSON nesting")
	}
	token, err := decoder.Token()
	if err != nil {
		return fmt.Errorf("read menu JSON: %w", err)
	}
	if token == nil {
		return errors.New("menu configuration must not contain null")
	}
	delimiter, container := token.(json.Delim)
	if !container {
		return nil
	}
	seen := make(map[string]bool)
	for decoder.More() {
		if delimiter == '{' {
			token, err := decoder.Token()
			if err != nil {
				return fmt.Errorf("read menu JSON member: %w", err)
			}
			name, ok := token.(string)
			if !ok || seen[name] {
				return errors.New("menu configuration contains a duplicate object member")
			}
			seen[name] = true
		}
		if err := checkMenuValue(decoder, depth+1); err != nil {
			return err
		}
	}
	if _, err := decoder.Token(); err != nil {
		return fmt.Errorf("read menu JSON delimiter: %w", err)
	}
	return nil
}

func resolveMenuConfig(defaults Config, options *ConfigOptions) (Config, error) {
	config, err := ResolveConfig(defaults, options, nil)
	if err != nil {
		return Config{}, err
	}
	for index, mapping := range config.Spec {
		if err := validateMenuAction(mapping); err != nil {
			return Config{}, fmt.Errorf("whichkey: spec[%d] %q: %w", index+1, mapping.Keys, err)
		}
	}
	if err := validateMenuConflicts(config.Spec); err != nil {
		return Config{}, err
	}
	return config, nil
}

func validateMenuAction(mapping Mapping) error {
	if mapping.CallbackRef != "" {
		return errors.New("callbacks are not supported in menu data")
	}
	if mapping.Action == "" {
		return nil // Shared Mapping validation owns group/provider-only shapes.
	}
	args := mapping.Args
	if len(args) == 0 {
		args = json.RawMessage(`{}`)
	}
	switch mapping.Action {
	case menuPaneSplit:
		fields, err := menuObject(args, "direction")
		if err != nil {
			return err
		}
		direction, err := completionString(fields["direction"])
		if err != nil || direction != menuSplitRight && direction != menuSplitDown {
			return errors.New("pane.split requires direction right or down")
		}
	case menuSessionSelect:
		return validateMenuSessionArgs(args)
	case menuPaneClose, menuTabClose, menuSessionClose:
		if !mapping.Confirm {
			return errors.New("destructive resource actions require confirm: true")
		}
		_, err := menuObject(args)
		return err
	case menuSessionSwitch:
		if mapping.Confirm {
			return errors.New("session.switch is a Go menu intent, not a confirmable action")
		}
		_, err := menuObject(args)
		return err
	case menuSessionNew, menuSessionPrevious, menuSessionNext, menuSessionRename,
		menuTabNew, menuTabPrevious, menuTabNext, menuTabRename,
		menuSidebarToggle, menuPaletteToggle:
		_, err := menuObject(args)
		return err
	default:
		return errors.New("unsupported menu action (custom Lua actions are not supported)")
	}
	return nil
}

func validateMenuSessionArgs(args json.RawMessage) error {
	fields, err := menuObject(args, "session_id", "window_id")
	if err != nil {
		return err
	}
	id, err := completionString(fields["session_id"])
	if err != nil || !popupResourceID(id, "session") {
		return errors.New("session.select requires a full session:UUID session_id")
	}
	if raw, present := fields["window_id"]; present {
		id, err := completionString(raw)
		if err != nil || !popupResourceID(id, "window") {
			return errors.New("session.select window_id must be a full window:UUID")
		}
	}
	return nil // Live destination membership is checked by the dispatch owner.
}

func knownMenuProvider(provider string) bool {
	switch provider {
	case menuProviderSessions, menuProviderWindows, menuProviderPanes, menuProviderHosts:
		return true
	default:
		return false
	}
}

func menuProvider(mapping Mapping) string {
	if mapping.Action == menuSessionSwitch {
		return menuProviderSessions
	}
	return mapping.Provider
}

func validateMenuConflicts(mappings []Mapping) error {
	// Legacy evaluated configurations omit disabled keys from the active trie.
	// Public data is stricter: a later condition edit must not reveal conflicts.
	all := make([]Mapping, len(mappings))
	for index, mapping := range mappings {
		mapping.Cond = nil
		if mapping.Action == menuSessionSwitch {
			label := mapping.Desc
			mapping.Action, mapping.Args = "", nil
			mapping.Provider, mapping.Group = menuProviderSessions, &label
		}
		all[index] = mapping
	}
	return ValidateMappings(all)
}

// ValidateMenuProviderChildren checks a generated batch before it can replace a
// reserved marker. It validates structure, not authority to dispatch: the parent
// must retain exact resource targets separately and validate the combined menu.
// Generated internal intents are never accepted by the public file loader.
func ValidateMenuProviderChildren(marker Mapping, children []Mapping) error {
	const maxRows = 1024 // Includes pagination/metadata rows, not only resources.
	provider := menuProvider(marker)
	if provider == "" || !marker.Enabled() || len(children) > maxRows {
		return errors.New("whichkey: expansion needs an enabled provider and at most 1024 rows")
	}
	if err := ValidateMappings([]Mapping{marker}); err != nil {
		return err
	}
	if err := validateMenuAction(marker); err != nil {
		return err
	}
	prefix, err := NormalizeKeys(marker.Keys)
	if err != nil {
		return err
	}
	keyBytes := 0
	for _, child := range children {
		keyBytes += len(child.Keys)
		if keyBytes > menuConfigLimit {
			return errors.New("whichkey: generated rows exceed 4 MiB of key data")
		}
		if checkErr := validateMenuProviderChild(provider, prefix, child); checkErr != nil {
			return checkErr
		}
	}
	return validateMenuConflicts(children)
}

func validateMenuProviderChild(provider, prefix string, child Mapping) error {
	path, err := NormalizeKeys(child.Keys)
	if err != nil {
		return err
	}
	if path == prefix || !strings.HasPrefix(path, prefix) {
		return errors.New("whichkey: generated rows must stay strictly inside their reserved subtree")
	}
	if child.Provider != "" || child.Action == menuSessionSwitch || child.CallbackRef != "" {
		return errors.New("whichkey: generated rows cannot create providers or callbacks")
	}
	if provider == menuProviderHosts && child.Action != "" {
		return errors.New("whichkey: known-hosts rows are listing-only")
	}
	return nil
}
