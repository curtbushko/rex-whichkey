package whichkey

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"temp/rex"
)

const (
	menuPaneFocus = "go.pane.focus"
	menuSessionID = "session_id"
	menuWindowID  = "window_id"
)

// providerTarget never crosses the PTY boundary. A key selects this exact
// retained resource; labels, child output and current focus supply no identity.
type providerTarget struct{ block rex.Block }

type providerRow struct {
	id      string
	label   string
	mapping Mapping
	block   rex.Block
}

// canonicalMenuSnapshot validates in original order before changing spelling.
// Parent matching, generated prefixes and the child trie then share one path.
func canonicalMenuSnapshot(config Config) (Config, error) {
	if err := config.Validate(); err != nil {
		return Config{}, err
	}
	spec := make([]Mapping, len(config.Spec))
	for i, mapping := range config.Spec {
		mapping = cloneMenuMapping(mapping)
		canonical, keyErr := NormalizeKeys(mapping.Keys)
		if keyErr != nil {
			return Config{}, keyErr
		}
		mapping.Keys = canonical
		spec[i] = mapping
	}
	config.Spec = spec
	return config, nil
}

// decodePopupConfig is an internal resolved snapshot, not the legacy public
// callback decoder. Only the parent can produce internal provider intents.
func decodePopupConfig(input io.Reader) (Config, error) {
	var config Config
	if err := readRuntime(input, &config, menuConfigLimit); err != nil {
		return Config{}, err
	}
	if err := config.Validate(); err != nil {
		return Config{}, err
	}
	for _, mapping := range config.Spec {
		if mapping.CallbackRef != "" || mapping.Enabled() && menuProvider(mapping) != "" {
			return Config{}, errors.New("whichkey: invalid executable popup snapshot")
		}
	}
	return canonicalMenuSnapshot(config)
}

func menuSnapshotBound(config Config) error {
	data, err := json.Marshal(config)
	if err != nil || len(data) > menuConfigLimit {
		return errors.New("whichkey: expanded menu exceeds 4 MiB; not truncated")
	}
	return nil
}

// providerRows formats one bounded inventory. Full IDs lead labels; metadata is
// escaped rather than treated as a resource selector. No row is truncated.
func providerRows(marker Mapping, entries []providerRow) ([]Mapping, error) {
	rows := make([]Mapping, 0, len(entries))
	keyBytes := 0
	overflow := false
	add := func(mapping Mapping) {
		keyBytes += len(mapping.Keys)
		if len(rows) >= 1024 || keyBytes > menuConfigLimit {
			overflow = true
			return
		}
		mapping.Hidden = marker.Hidden
		rows = append(rows, mapping)
	}
	addGroup := func(keys, label string) { add(Mapping{Keys: keys, Group: &label}) }
	if len(entries) == 0 {
		addGroup(marker.Keys+"0", "No resources in this snapshot")
	}
	for index, entry := range entries {
		if overflow {
			break
		}
		keys := hostMenuListKey(marker.Keys, index, len(entries), menuProvider(marker), addGroup)
		mapping := entry.mapping
		mapping.Keys = keys
		mapping.Desc = "[" + entry.id + "] " + hostMenuText(entry.label)
		add(mapping)
	}
	if overflow {
		return nil, errors.New("whichkey: provider rows exceed transport bounds; not truncated")
	}
	return rows, nil
}

// menuInventories caches exactly one attempt per resource kind. It is local to
// this launch, not a refresh mechanism, registry or shared daemon.
type menuInventories struct {
	sessions, windows, panes             []providerRow
	sessionsRead, windowsRead, panesRead bool
	sessionsErr, windowsErr, panesErr    error
}

func (inventory *menuInventories) read(ctx context.Context, client *rex.Client, sessionID, provider string) ([]providerRow, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	switch provider {
	case menuProviderSessions:
		if !inventory.sessionsRead {
			inventory.sessions, inventory.sessionsErr = readSessionRows(ctx, client, sessionID)
			inventory.sessionsRead = true
		}
		return inventory.sessions, inventory.sessionsErr
	case menuProviderWindows:
		if !inventory.windowsRead {
			inventory.windows, inventory.windowsErr = readWindowRows(ctx, client, sessionID)
			inventory.windowsRead = true
		}
		return inventory.windows, inventory.windowsErr
	case menuProviderPanes:
		if !inventory.panesRead {
			inventory.panes, inventory.panesErr = readPaneRows(ctx, client, sessionID)
			inventory.panesRead = true
		}
		return inventory.panes, inventory.panesErr
	default:
		return nil, errors.New("whichkey: unexpanded provider")
	}
}

func readSessionRows(ctx context.Context, client *rex.Client, currentID string) ([]providerRow, error) {
	sessions, err := client.Sessions(ctx)
	if err != nil {
		return nil, err
	}
	if len(sessions) > maxSessionMenuEntries {
		return nil, errProviderOverflow
	}
	rows := make([]providerRow, 0, len(sessions))
	for _, session := range sessions {
		mapping, encodeErr := sessionRowMapping(session.ID, "")
		if encodeErr != nil {
			return nil, encodeErr
		}
		label := session.Label
		if session.ID == currentID {
			label = "[current] " + label
		}
		rows = append(rows, providerRow{id: session.ID, label: label, mapping: mapping})
	}
	return rows, nil
}

func readWindowRows(ctx context.Context, client *rex.Client, sessionID string) ([]providerRow, error) {
	session, err := client.Session(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	windows, err := session.Windows(ctx)
	if err != nil {
		return nil, err
	}
	if len(windows) > maxSessionMenuEntries {
		return nil, errProviderOverflow
	}
	rows := make([]providerRow, 0, len(windows))
	for _, window := range windows {
		mapping, encodeErr := sessionRowMapping(sessionID, window.ID)
		if encodeErr != nil {
			return nil, encodeErr
		}
		label := window.Label
		if window.Active {
			label = "[current] " + label
		}
		rows = append(rows, providerRow{id: window.ID, label: label, mapping: mapping})
	}
	return rows, nil
}

func readPaneRows(ctx context.Context, client *rex.Client, sessionID string) ([]providerRow, error) {
	session, err := client.Session(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	blocks, err := session.Blocks(ctx)
	if err != nil {
		return nil, err
	}
	if len(blocks) > maxSessionMenuEntries {
		return nil, errProviderOverflow
	}
	rows := make([]providerRow, 0, len(blocks))
	for _, block := range blocks {
		rows = append(rows, providerRow{id: block.ID, label: "Window " + block.WindowID + " / layer " + block.LayerID,
			mapping: Mapping{Action: menuPaneFocus}, block: block})
	}
	return rows, nil
}

func sessionRowMapping(sessionID, windowID string) (Mapping, error) {
	destination := map[string]string{menuSessionID: sessionID}
	if windowID != "" {
		destination[menuWindowID] = windowID
	}
	args, err := json.Marshal(destination)
	if err != nil {
		return Mapping{}, errors.New("whichkey: cannot encode retained session choice")
	}
	return Mapping{Action: menuSessionSelect, Args: args}, nil
}

var errProviderOverflow = errors.New("whichkey: provider exceeds 512 resources; not truncated")

func expandResourceMarker(marker Mapping, entries []providerRow, readErr error) ([]Mapping, error) {
	if errors.Is(readErr, errProviderOverflow) {
		return nil, readErr
	}
	root := marker
	root.Provider, root.Action, root.Args = "", "", nil
	if root.Group == nil {
		label := marker.Desc
		root.Group = &label
		root.Desc = ""
	}
	var children []Mapping
	if readErr != nil {
		label := "Warning: inventory unavailable; reopen to retry"
		children = []Mapping{{Keys: marker.Keys + "0", Group: &label, Hidden: marker.Hidden}}
	} else {
		var formatErr error
		children, formatErr = providerRows(marker, entries)
		if formatErr != nil {
			return nil, formatErr
		}
	}
	if err := ValidateMenuProviderChildren(marker, children); err != nil {
		return nil, err
	}
	return append([]Mapping{root}, children...), nil
}

// nativeSessionMenu expands all providers once before sizing or handoff. Each
// kind shares one read-only inventory, even when configured at several prefixes.
func nativeSessionMenu(parent context.Context, client *rex.Client, origin launchOrigin, config Config) (Config, map[string]providerTarget, error) {
	expandedConfig, err := resolveHostMenu(parent, config, origin.selection)
	if err != nil {
		return Config{}, nil, err
	}
	config = expandedConfig
	targets := make(map[string]providerTarget)
	var inventory menuInventories
	expanded := make([]Mapping, 0, len(config.Spec))
	bytesUsed := 0
	for _, marker := range config.Spec {
		rows := []Mapping{marker}
		if marker.Enabled() && menuProvider(marker) != "" {
			entries, readErr := inventory.read(parent, client, origin.SessionID, menuProvider(marker))
			if cancelErr := parent.Err(); cancelErr != nil {
				return Config{}, nil, cancelErr
			}
			rows, err = expandResourceMarker(marker, entries, readErr)
			if err != nil {
				return Config{}, nil, err
			}
			retainPaneTargets(targets, rows, entries)
		}
		encoded, encodeErr := json.Marshal(rows)
		if encodeErr != nil {
			return Config{}, nil, encodeErr
		}
		bytesUsed += len(encoded)
		if bytesUsed > menuConfigLimit {
			return Config{}, nil, errors.New("whichkey: combined menu exceeds 4 MiB; not truncated")
		}
		expanded = append(expanded, rows...)
	}
	config.Spec = expanded
	if checkErr := validateMenuConflicts(config.Spec); checkErr != nil {
		return Config{}, nil, fmt.Errorf("whichkey: expanded menu: %w", checkErr)
	}
	if checkErr := menuSnapshotBound(config); checkErr != nil {
		return Config{}, nil, checkErr
	}
	return config, targets, nil
}

func retainPaneTargets(targets map[string]providerTarget, rows []Mapping, entries []providerRow) {
	index := 0
	for _, row := range rows {
		if row.Action == menuPaneFocus {
			targets[row.Keys] = providerTarget{block: entries[index].block}
			index++
		}
	}
}

func validatePaneChoice(ctx context.Context, target providerTarget) (rex.Block, error) {
	block, err := target.block.Inspect(ctx)
	if err != nil || block.ID != target.block.ID || block.WindowID != target.block.WindowID || block.LayerID != target.block.LayerID {
		return rex.Block{}, errors.New("whichkey: retained pane is unavailable or moved; not retargeted")
	}
	return block, nil
}
