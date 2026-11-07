package whichkey

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/rivo/uniseg"
	"github.com/stretchr/testify/require"

	"temp/rex"
)

func TestConfiguredRenameUsesMenuPage(t *testing.T) {
	for _, keys := range []string{"tr", "wr"} {
		t.Run(keys, func(t *testing.T) {
			action := "client.tab.rename"
			if keys == "wr" {
				action = "session.rename"
			}
			config := DefaultConfig()
			config.Spec = []Mapping{{Keys: keys, Action: action, Desc: "Rename"}}
			config.textEntry = &textEntrySpec{Initial: "Tab", SessionInitial: "Session"}
			menu, err := NewMenu(config.Spec)
			require.NoError(t, err)
			menu.textSpec = config.textEntry
			require.Nil(t, menu.HandleKey(keys[:1], 1).Selection)
			require.Nil(t, menu.HandleKey("r", 1).Selection)
			require.NotNil(t, menu.entry, "rename must open an inline page, not emit a Rex action")
			initial := "Tab"
			if keys == "wr" {
				initial = "Session"
			}
			require.Equal(t, initial, menu.entry.value)
			require.Nil(t, menu.HandleKey("\x1b", 1).Selection)
			require.Nil(t, menu.entry)
			require.False(t, menu.Done())
			menu.HandleKey("r", 1)
			for _, key := range "New name" {
				menu.HandleKey(string(key), 1)
			}
			event := menu.HandleKey("\r", 1)
			require.NotNil(t, event.Selection)
			require.Equal(t, "New name", *event.Selection.textInput)
			var result bytes.Buffer
			require.NoError(t, writePopupMenuResult(&result, event, config))
			directory := t.TempDir()
			require.NoError(t, os.Chmod(directory, 0700))
			require.NoError(t, os.WriteFile(filepath.Join(directory, "result.json"), result.Bytes(), 0600))
			completion, err := readPopupResult(directory, strings.Repeat("a", 64), config)
			require.NoError(t, err)
			require.NotNil(t, completion.Selection.Text)
			require.Equal(t, "New name", *completion.Selection.Text)
			completion.textInput = nil // Configured transport carries the validated typed field.
			encoded, err := json.Marshal(completion)
			require.NoError(t, err)
			var decoded Completion
			require.NoError(t, json.Unmarshal(encoded, &decoded))
			require.NoError(t, decoded.ValidateSelection(config.Spec))
			require.Equal(t, "New name", *decoded.Selection.Text)
		})
	}
}

func TestCustomRenameKeyKeepsItsCallback(t *testing.T) {
	config := DefaultConfig()
	config.Spec = []Mapping{{Keys: "tr", CallbackRef: "cb:1:1", Desc: "Custom rename"},
		{Keys: "wr", Action: "session.rename", Desc: "Rename session"}}
	config.textEntry = &textEntrySpec{Initial: "Callback tab", SessionInitial: "Other session"}
	menu, err := NewMenu(config.Spec)
	require.NoError(t, err)
	menu.textSpec = config.textEntry
	menu.HandleKey("t", 1)
	event := menu.HandleKey("r", 1)
	require.Nil(t, menu.entry)
	require.NotNil(t, event.Selection)
	var output bytes.Buffer
	require.NoError(t, writePopupMenuResult(&output, event, config))
	directory := t.TempDir()
	require.NoError(t, os.Chmod(directory, 0700))
	require.NoError(t, os.WriteFile(filepath.Join(directory, "result.json"), output.Bytes(), 0600))
	completion, err := readPopupResult(directory, strings.Repeat("a", 64), config)
	require.NoError(t, err)
	require.Equal(t, "cb:1:1", completion.Selection.CallbackRef)
	require.Nil(t, completion.Selection.Text)
}

func TestRenameCompletionRetainsMappingAuthority(t *testing.T) {
	name := "Name"
	completion := Completion{Version: 1, Capability: strings.Repeat("a", 64), Kind: CompletionComplete,
		Selection: &CompletionSelection{Keys: "tr", Text: &name}}
	rename := Mapping{Keys: "tr", Action: "client.tab.rename", Desc: "Rename"}
	require.NoError(t, completion.ValidateSelection([]Mapping{rename}))
	require.Error(t, completion.ValidateSelection([]Mapping{{Keys: "tr", Action: "pane.close", Desc: "Close"}}))
	require.Error(t, completion.ValidateSelection([]Mapping{{Keys: "tr", CallbackRef: "cb:1:1", Desc: "Callback"}}))
	custom := rename
	custom.Args = json.RawMessage(`{"label":"retained"}`)
	require.Error(t, completion.ValidateSelection([]Mapping{custom}))
	for _, invalid := range []string{"", "   ", "\u00a0", "a\n", "a\x00", strings.Repeat("x", renameTextLimit+1)} {
		completion.Selection.Text = &invalid
		require.Error(t, completion.Validate())
	}
	completion.Selection.Text = nil
	require.Error(t, completion.ValidateSelection([]Mapping{rename}), "rename may not fall through to Rex UI")
}

func TestDirectRenameReadinessRetainsEdits(t *testing.T) {
	config := DefaultConfig()
	config.Spec = []Mapping{{Keys: "r", Action: menuTabRename, Desc: "Ready rename", Confirm: true}}
	config.textEntry = &textEntrySpec{}
	menu, err := NewMenu(config.Spec)
	require.NoError(t, err)
	menu.textSpec, menu.updates = &textEntrySpec{}, &menuRenameUpdates{}
	menu.providers = &menuProviderUpdates{base: config}
	state := directMenuState{config: config, targets: make(map[string]providerTarget)}
	menu.HandleKey("r", 1)
	require.NotNil(t, menu.entry)
	menu.HandleKey("界", 1)
	require.Nil(t, menu.HandleKey("\r", 1).Selection, "Enter before metadata must never queue submission")
	jobs := newMetadataJobs(context.Background(), popupRequest{}, nil)
	defer jobs.stop()
	jobs.started[metadataRenameKind], jobs.active = true, 1
	jobs.results <- metadataResult{kind: metadataRenameKind, labels: renameLabels{initial: [2]string{"late server name", "session"}}}
	changed, err := state.poll(menu, jobs)
	require.NoError(t, err)
	require.True(t, changed)
	require.Equal(t, "界", menu.entry.value, "late metadata must preserve typed value")
	require.Equal(t, len("界"), menu.entry.cursor)
	require.True(t, menu.entry.touched)
	require.False(t, menu.Done())
	menu.HandleKey("\r", 1)
	_, confirming := menu.Confirmation()
	require.True(t, confirming)
	event := menu.HandleKey("y", 1)
	require.NotNil(t, event.Selection)
	require.Equal(t, "界", *event.Selection.textInput)
	retained, err := state.validateSelection(event)
	require.NoError(t, err)
	require.Equal(t, menuTabRename, retained.Action)
	stale := *event.update
	stale.Load++
	event.update = &stale
	retained, err = state.validateSelection(event)
	require.Error(t, err)
	require.Empty(t, retained.Keys)
	changed, err = state.poll(menu, jobs)
	require.NoError(t, err)
	require.False(t, changed, "cached metadata must publish once")
}

func TestMetadataJobsBoundedAndOncePerKind(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	jobs := newMetadataJobs(ctx, popupRequest{}, nil)
	defer jobs.stop()
	jobs.wanted[0], jobs.wanted[1], jobs.wanted[2] = true, true, true
	jobs.start()
	require.Equal(t, 2, jobs.active)
	require.Equal(t, 2, cap(jobs.results))
	require.False(t, jobs.started[2])
	seen := make(map[int]bool)
	for range 3 {
		select {
		case result := <-jobs.results:
			jobs.active--
			require.False(t, seen[result.kind], "each kind must return once")
			require.Error(t, result.err, "standalone has no native inventory authority")
			seen[result.kind] = true
			jobs.start()
			require.LessOrEqual(t, jobs.active, 2)
		case <-ctx.Done():
			t.Fatal("bounded metadata jobs did not finish")
		}
	}
	require.Len(t, seen, 3)
	jobs.start()
	require.Zero(t, jobs.active, "failed loads remain cached, not retried")
	select {
	case result := <-jobs.results:
		t.Fatalf("unexpected repeated result: %d", result.kind)
	default:
	}
}

func TestDirectInventoryBindingAndRetainedTarget(t *testing.T) {
	group := "Panes"
	config := DefaultConfig()
	config.Spec = []Mapping{{Keys: "p", Group: &group, Provider: menuProviderPanes}}
	base, err := canonicalMenuSnapshot(config)
	require.NoError(t, err)
	pending, err := providerConfig(base, [4]*providerUpdate{})
	require.NoError(t, err)
	menu, err := NewMenu(pending.Spec)
	require.NoError(t, err)
	menu.providers = &menuProviderUpdates{base: base}
	menu.updates = &menuRenameUpdates{}
	menu.HandleKey("p", 1)
	require.Equal(t, "p", menu.Prefix())
	state := directMenuState{config: pending, targets: make(map[string]providerTarget)}
	jobs := newMetadataJobs(context.Background(), popupRequest{}, nil)
	defer jobs.stop()
	jobs.started[2], jobs.active = true, 1
	block := rex.Block{ID: "block:00000000-0000-0000-0000-000000000001", WindowID: "captured-window", LayerID: "captured-layer"}
	jobs.results <- metadataResult{kind: 2, entries: []providerRow{{id: block.ID, label: "[current] pane", block: block, mapping: Mapping{Action: menuPaneFocus}}}}
	changed, err := state.poll(menu, jobs)
	require.NoError(t, err)
	require.True(t, changed)
	require.Equal(t, block, state.targets["p1"].block)
	require.Contains(t, menu.Entries()[0].Label, block.ID)
	require.Contains(t, menu.Entries()[0].Label, "[current]")
	event := menu.HandleKey("1", 1)
	require.NotNil(t, event.Selection)
	mapping, err := state.validateSelection(event)
	require.NoError(t, err)
	require.Equal(t, "p1", mapping.Keys)
	event.provider = nil
	mapping, err = state.validateSelection(event)
	require.Error(t, err)
	require.Empty(t, mapping.Keys)
}

func TestMappingValidationTrustBoundaries(t *testing.T) {
	const testAction = "one"
	group := "Group"
	unsafe := "bad\n"
	for _, mapping := range []Mapping{
		{Keys: "x", Action: testAction, CallbackRef: "two", Desc: "Conflict"},
		{Keys: "x", Provider: "unknown", Group: &group},
		{Keys: "x", Provider: menuProviderPanes, Group: &group, Confirm: true},
		{Keys: "x", Action: " ", Desc: "Blank"},
		{Keys: "x", Action: testAction, Group: &group, Desc: "Action group"},
		{Keys: "x", Args: json.RawMessage(`{}`)},
		{Keys: "x", Action: testAction, Desc: "Unsafe\x1b"},
		{Keys: "x", Group: &unsafe},
		{Keys: "x", Action: testAction, Desc: "Args", Args: json.RawMessage(`null`)},
		{Keys: "x", Action: testAction, Desc: "Args", Args: json.RawMessage(`[]`)},
	} {
		require.Error(t, ValidateMappings([]Mapping{mapping}))
	}
	for _, data := range []string{`{"x":"a","x":"b"}`, `{"X":"a"}`, `{"x":null}`, `{"x":"a"} {}`} {
		_, err := completionObject([]byte(data), "x")
		require.Error(t, err)
	}
	_, err := completionObject([]byte{'{', '"', 'x', '"', ':', '"', 0xff, '"', '}'}, "x")
	require.Error(t, err)
	require.NoError(t, ValidateMappings([]Mapping{{Keys: "x", Action: testAction, Desc: "Good", Args: json.RawMessage(`{}`)}}))
	_, err = decodeMenuOptions([]byte(`{"title":"one","title":"two"}`))
	require.Error(t, err)
	path := filepath.Join(t.TempDir(), "large.json")
	require.NoError(t, os.WriteFile(path, []byte(strings.Repeat(" ", menuConfigLimit+1)), 0600))
	_, err = readMenuConfigFile(path)
	require.ErrorContains(t, err, "4 MiB")
}

func TestRenderPaginationAndConfirmation(t *testing.T) {
	t.Setenv("NO_COLOR", "")
	config := DefaultConfig()
	config.Position = positionTop
	for index := range 20 {
		config.Spec = append(config.Spec, Mapping{Keys: string(rune('a' + index)), Action: "test", Desc: fmt.Sprintf("界 label %02d", index)})
	}
	menu, err := NewMenu(config.Spec)
	require.NoError(t, err)
	theme := NewTheme(TerminalPalette{}, true, true)
	frame := RenderMenu(menu, config, theme, 24, 5)
	require.Greater(t, frame.Pages, 1)
	require.Zero(t, frame.Y)
	require.Len(t, frame.Lines, frame.Height)
	require.Contains(t, strings.Join(frame.Lines, "\n"), "1/")
	menu.HandleKey("\t", frame.Pages)
	next := RenderMenu(menu, config, theme, 24, 5)
	require.Equal(t, 1, next.Page)
	require.NotEqual(t, frame.Lines, next.Lines)
	zero := RenderMenu(menu, config, theme, 24, 0)
	require.Empty(t, zero.Lines)
	require.Equal(t, 1, menu.Page(), "zero viewport must not mutate navigation")
	for _, line := range next.Lines {
		require.LessOrEqual(t, uniseg.StringWidth(line), next.Width)
	}
	config.Spec = []Mapping{{Keys: "x", Action: "close", Desc: "Close", Confirm: true}}
	menu, err = NewMenu(config.Spec)
	require.NoError(t, err)
	menu.HandleKey("x", 1)
	tiny := RenderMenu(menu, config, theme, 60, 1)
	require.Contains(t, strings.Join(tiny.Lines, ""), "y=yes / other=cancel")
	require.Nil(t, menu.HandleKey("x", 1).Selection)
	require.False(t, menu.Done(), "canceling confirmation must not dispatch or fall through")
}

func TestHostListingPreservesDetailsAndInput(t *testing.T) {
	inventory := rex.HostInventory{Home: rex.InventoryHome{"status": json.RawMessage(`"cached"`)},
		Warnings: []rex.InventoryWarning{{"message": json.RawMessage(`"not live"`)}}}
	for index := range 10 {
		inventory.Hosts = append(inventory.Hosts, rex.HostEntry{Held: true, From: "home", Host: rex.HostRecord{
			ID: fmt.Sprintf("host:%02d", 9-index), Label: "same label", HostID: "machine", Source: "source", AddedFrom: "added", UpdatedAt: "updated", Endpoints: []string{"ssh://one", "unsafe\nendpoint"}}})
	}
	originalFirst := inventory.Hosts[0].Host.ID
	rows, err := HostMenuMappings("h", inventory)
	require.NoError(t, err)
	require.Equal(t, originalFirst, inventory.Hosts[0].Host.ID, "formatting must not sort caller's inventory")
	labels := make(map[string]string)
	for _, row := range rows {
		require.Empty(t, row.Action)
		require.Empty(t, row.CallbackRef)
		require.NotNil(t, row.Group)
		labels[row.Keys] = *row.Group
	}
	require.Equal(t, "same label [host:00] [held: not yet sent to home]", labels["hl1"])
	require.Equal(t, "More hosts (10-10)", labels["hln"])
	require.Contains(t, labels["hln1"], "host:09")
	require.Equal(t, "Inventory owner: home", labels["hl1o"])
	require.Equal(t, "Machine/account ID (not inventory ID): machine", labels["hl1m"])
	require.Equal(t, `"unsafe\nendpoint"`, labels["hl1e2"])
	require.Contains(t, labels["hw1"], "not live")
	inventory.Warnings = []rex.InventoryWarning{nil}
	_, err = HostMenuMappings("h", inventory)
	require.ErrorContains(t, err, "warning 1 is null")
}

func TestDirectRenameFailureIsUnavailableAndCached(t *testing.T) {
	config := DefaultConfig()
	config.Spec = []Mapping{{Keys: "r", Action: menuTabRename, Desc: "Unavailable rename"}}
	menu, err := NewMenu(config.Spec)
	require.NoError(t, err)
	menu.textSpec, menu.updates = &textEntrySpec{}, &menuRenameUpdates{}
	menu.HandleKey("r", 1)
	state := directMenuState{labels: &renameLabels{err: errors.New("unavailable")}}
	changed, err := state.applyRenameLabels(menu)
	require.NoError(t, err)
	require.True(t, changed)
	require.Equal(t, renameStateUnavailable, state.rename[0].State)
	require.Nil(t, menu.HandleKey("\r", 1).Selection)
	changed, err = state.applyRenameLabels(menu)
	require.NoError(t, err)
	require.False(t, changed)
	require.Contains(t, menu.renameStatus(), "unavailable")
}

func TestDirectInvocationRejectsRetargeting(t *testing.T) {
	const sessionID = "session:00000000-0000-0000-0000-000000000001"
	const originID = "block:00000000-0000-0000-0000-000000000002"
	const ownID = "block:00000000-0000-0000-0000-000000000003"
	options := DirectOptions{OriginSession: sessionID, OriginBlock: originID,
		OriginWindow: "window:00000000-0000-0000-0000-000000000004",
		OriginLayer:  "layer:00000000-0000-0000-0000-000000000005"}
	t.Setenv("REX_SESSION", sessionID)
	t.Setenv("REX_BLOCK", ownID)
	t.Setenv("REX_CLIENT", "client:popup-client-is-not-origin-intent")
	own, requested, err := directInvocation(context.Background(), options)
	require.NoError(t, err)
	require.Equal(t, ownID, own)
	require.Empty(t, requested, "popup client environment must never become original app intent")
	blank := ""
	options.Client = &blank
	_, _, err = directInvocation(context.Background(), options)
	require.Error(t, err)
	options.Client = nil
	for _, test := range []struct {
		name   string
		parent context.Context
		valid  bool
	}{
		{name: "missing context"},
		{name: "present context", parent: context.Background(), valid: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			_, _, contextErr := directInvocation(test.parent, options)
			if test.valid {
				require.NoError(t, contextErr)
			} else {
				require.Error(t, contextErr)
			}
		})
	}
	t.Setenv("REX_BLOCK", originID)
	_, _, err = directInvocation(context.Background(), options)
	require.Error(t, err, "ordinary origin pane must never be detached")
	t.Setenv("REX_BLOCK", ownID)
	t.Setenv("REX_SESSION", "other")
	_, _, err = directInvocation(context.Background(), options)
	require.Error(t, err, "a different execution session must not retarget the origin")
}

func TestNativeNewTabIntentIsStrict(t *testing.T) {
	for _, args := range []json.RawMessage{nil, json.RawMessage(`{}`)} {
		require.NoError(t, validateNativeNewTab(Mapping{Action: menuTabNew, Args: args}))
	}
	for _, mapping := range []Mapping{
		{Action: menuTabClose},
		{Action: menuTabNew, Args: json.RawMessage(`{"index":1}`)},
		{Action: menuTabNew, Args: json.RawMessage(`null`)},
		{Action: menuTabNew, Args: json.RawMessage(`[]`)},
	} {
		require.Error(t, validateNativeNewTab(mapping))
	}
}
