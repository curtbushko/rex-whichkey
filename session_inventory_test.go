package whichkey

import (
	"context"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"temp/rex"
)

const (
	inventoryLocalSession      = "session:00000000-0000-0000-0000-000000000001"
	inventorySwitchDescription = "Switch sessions"
)

func sessionInventoryClient(t *testing.T) (*rex.Client, string) {
	t.Helper()
	script, err := os.ReadFile("tests/fixtures/session-inventory.sh")
	require.NoError(t, err)
	directory := t.TempDir()
	executable := filepath.Join(directory, "rex")
	require.NoError(t, os.WriteFile(executable, script, 0700))
	logPath := filepath.Join(directory, "calls")
	t.Setenv("SESSION_INVENTORY_LOG", logPath)
	client, err := rex.NewContext(context.Background(), rex.Options{
		Selection: &rex.CLISelection{Path: executable}, Server: "unix:///tmp/session-inventory.sock", Timeout: 3 * time.Second,
	})
	require.NoError(t, err)
	return client, logPath
}

func inventoryMapping(t *testing.T, rows []Mapping, keys string) Mapping {
	t.Helper()
	for _, row := range rows {
		if row.Keys == keys {
			return row
		}
	}
	require.FailNow(t, "expected inventory key missing", keys)
	return Mapping{}
}

func TestRemoteSessionInventory(t *testing.T) {
	client, _ := sessionInventoryClient(t)
	entries, err := readSessionRows(context.Background(), client, inventoryLocalSession)
	require.NoError(t, err)
	require.Len(t, entries, 2)
	rows, err := expandResourceMarker(Mapping{Keys: "ws", Action: menuSessionSwitch, Desc: inventorySwitchDescription}, entries, nil)
	require.NoError(t, err)
	local := inventoryMapping(t, rows, "wsl1")
	remote := inventoryMapping(t, rows, "wss1")
	require.Equal(t, "[current] Local", local.Desc)
	require.Equal(t, "studio (1)", *inventoryMapping(t, rows, "wss").Group)
	require.Equal(t, "Editor", remote.Desc)
	require.NotContains(t, remote.Desc, "session:")
	require.NotContains(t, remote.Desc, "remote:")
	require.JSONEq(t, `{"session_id":"session:00000000-0000-0000-0000-000000000002"}`, string(remote.Args))
}

func TestRemoteSessionSelectionUsesRetainedServer(t *testing.T) {
	client, logPath := sessionInventoryClient(t)
	entries, err := readSessionRows(context.Background(), client, inventoryLocalSession)
	require.NoError(t, err)
	require.Len(t, entries, 2)
	marker := Mapping{Keys: "ws", Action: menuSessionSwitch, Desc: inventorySwitchDescription}
	rows, err := expandResourceMarker(marker, entries, nil)
	require.NoError(t, err)
	targets := make(map[string]providerTarget)
	retainProviderTargets(targets, rows, entries)
	choice := inventoryMapping(t, rows, "wss1")
	target := targets[choice.Keys]
	require.NotNil(t, target.session)
	require.Equal(t, "studio.example", target.session.Server())
	require.NoError(t, validateSessionChoice(context.Background(), client, choice, target.session), "validate on the captured remote even if its display name changes")
	changed := cloneMenuMapping(choice)
	changed.Args = []byte(`{"session_id":"` + inventoryLocalSession + `"}`)
	require.Error(t, validateSessionChoice(context.Background(), client, changed, target.session))
	t.Setenv("SESSION_REMOTE_GONE", "1")
	require.Error(t, validateSessionChoice(context.Background(), client, choice, target.session), "a missing remote must not resolve on another server")
	data, err := os.ReadFile(logPath)
	require.NoError(t, err)
	require.NotContains(t, string(data), "--server unix:///tmp/session-inventory.sock --autostart=false --timeout 3s session inspect")
}

func TestSessionInventoryWarningsRemainVisible(t *testing.T) {
	client, _ := sessionInventoryClient(t)
	t.Setenv("SESSION_INVENTORY_WARNING", "1")
	entries, err := readSessionRows(context.Background(), client, inventoryLocalSession)
	require.NoError(t, err)
	require.Len(t, entries, 3)
	require.Contains(t, *entries[1].mapping.Group, "cached home inventory")
	require.Contains(t, *entries[1].mapping.Group, "saved-home")
	marker := Mapping{Keys: "ws", Action: menuSessionSwitch, Desc: inventorySwitchDescription}
	rows, err := expandResourceMarker(marker, entries, nil)
	require.NoError(t, err)
	targets := make(map[string]providerTarget)
	retainProviderTargets(targets, rows, entries)
	require.NotNil(t, targets["wss1"].session)
	require.Equal(t, "studio.example", targets["wss1"].session.Server())
}

func TestSessionInventoryBoundsAndAmbiguousIDs(t *testing.T) {
	full := make([]providerRow, maxSessionMenuEntries)
	_, err := appendSessionRows(full, []rex.Session{{ID: inventoryLocalSession}}, "", "")
	require.ErrorIs(t, err, errProviderOverflow)
	_, err = appendSessionWarning(full, "offline")
	require.ErrorIs(t, err, errProviderOverflow)
	_, err = appendSessionRows([]providerRow{{id: inventoryLocalSession}}, []rex.Session{{ID: inventoryLocalSession}}, "copy", "")
	require.ErrorContains(t, err, "multiple hosts")
	event := MenuEvent{provider: &renameBinding{Kind: menuProviderSessions}}
	err = dispatchNativeSelection(context.Background(), nil, launchOrigin{}, rex.Block{}, "",
		Mapping{Keys: "ws1", Action: menuSessionSelect}, event, nil, nil)
	require.ErrorContains(t, err, "missing retained session destination")
}

func TestGroupedTargetsDoNotFollowDisplayOrder(t *testing.T) {
	client, _ := sessionInventoryClient(t)
	entries, err := readSessionRows(context.Background(), client, inventoryLocalSession)
	require.NoError(t, err)
	slices.Reverse(entries)
	marker := Mapping{Keys: "ws", Action: menuSessionSwitch, Desc: inventorySwitchDescription}
	rows, err := expandResourceMarker(marker, entries, nil)
	require.NoError(t, err)
	targets := make(map[string]providerTarget)
	retainProviderTargets(targets, rows, entries)
	require.Equal(t, inventoryLocalSession, targets["wsl1"].session.ID)
	require.Equal(t, "studio.example", targets["wss1"].session.Server())
	require.NoError(t, validateSessionChoice(context.Background(), client, inventoryMapping(t, rows, "wss1"), targets["wss1"].session))
}

func TestEmptyRemoteRemainsAHostGroup(t *testing.T) {
	client, _ := sessionInventoryClient(t)
	t.Setenv("SESSION_REMOTE_EMPTY", "1")
	entries, err := readSessionRows(context.Background(), client, inventoryLocalSession)
	require.NoError(t, err)
	rows, err := expandResourceMarker(Mapping{Keys: "ws", Action: menuSessionSwitch, Desc: inventorySwitchDescription}, entries, nil)
	require.NoError(t, err)
	require.Equal(t, "studio (0)", *inventoryMapping(t, rows, "wss").Group)
	notice := inventoryMapping(t, rows, "wss1")
	require.Empty(t, notice.Action)
	require.Contains(t, *notice.Group, "No sessions")
}

func TestUnavailableRemotePreservesLocalSessions(t *testing.T) {
	client, logPath := sessionInventoryClient(t)
	t.Setenv("SESSION_REMOTE_UNAVAILABLE", "1")
	entries, err := readSessionRows(context.Background(), client, inventoryLocalSession)
	require.NoError(t, err)
	require.Len(t, entries, 2)
	rows, err := expandResourceMarker(Mapping{Keys: "ws", Action: menuSessionSwitch, Desc: inventorySwitchDescription}, entries, nil)
	require.NoError(t, err)
	require.Equal(t, "[current] Local", inventoryMapping(t, rows, "wsl1").Desc)
	require.Equal(t, "studio (0)", *inventoryMapping(t, rows, "wss").Group)
	warning := inventoryMapping(t, rows, "wss1")
	require.Contains(t, *warning.Group, "studio")
	require.Contains(t, *warning.Group, "unavailable")
	require.Empty(t, warning.Action)
	log, err := os.ReadFile(logPath)
	require.NoError(t, err)
	require.Equal(t, 1, strings.Count(string(log), "--server studio.example"), "remote inventory reads must not retry")
}
