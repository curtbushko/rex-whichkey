package whichkey

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func groupedSessionEntry(t *testing.T, host string, index int) providerRow {
	t.Helper()
	id := fmt.Sprintf("session:00000000-0000-0000-0000-%012x", index+1)
	mapping, err := sessionRowMapping(id, "")
	require.NoError(t, err)
	return providerRow{id: id, host: host, label: fmt.Sprintf("Project %d", index+1), mapping: mapping}
}

func TestSessionsGroupedByHost(t *testing.T) {
	entries := []providerRow{groupedSessionEntry(t, "", 0)}
	for index := 1; index <= 30; index++ {
		entries = append(entries, groupedSessionEntry(t, "gamingrig", index))
	}
	marker := Mapping{Keys: "s", Provider: menuProviderSessions}
	rows, err := providerRows(marker, entries)
	require.NoError(t, err)
	menu, err := NewMenu(rows)
	require.NoError(t, err)
	menu.HandleKey("s", 1)
	root := menu.Entries()
	require.Len(t, root, 2)
	require.Equal(t, "l", root[0].Key)
	require.Equal(t, "This Mac (1)", root[0].Label)
	require.Equal(t, "g", root[1].Key)
	require.Equal(t, "gamingrig (30)", root[1].Label)
}

func TestGroupedSessionNavigationRetainsEveryChoice(t *testing.T) {
	var entries []providerRow
	for index := 0; index < 30; index++ {
		entries = append(entries, groupedSessionEntry(t, "gamingrig", index))
	}
	rows, err := providerRows(Mapping{Keys: "s", Provider: menuProviderSessions}, entries)
	require.NoError(t, err)
	for index, entry := range entries {
		menu, menuErr := NewMenu(rows)
		require.NoError(t, menuErr)
		menu.HandleKey("s", 1)
		menu.HandleKey("g", 1)
		trail := menu.Breadcrumbs()
		require.NotEmpty(t, trail)
		require.Equal(t, "gamingrig", trail[len(trail)-1].Label)
		for range index / 9 {
			menu.HandleKey("n", 1)
		}
		require.LessOrEqual(t, len(menu.Entries()), 10)
		frame := RenderMenu(menu, DefaultConfig(), NewTheme(TerminalPalette{}, true, false), 45, 5)
		require.LessOrEqual(t, frame.Height, 5)
		require.NotContains(t, strings.Join(frame.Lines, "\n"), "remote:")
		key := strconv.Itoa(index%9 + 1)
		event := menu.HandleKey(key, frame.Pages)
		require.NotNil(t, event.Selection)
		require.JSONEq(t, string(entry.mapping.Args), string(event.Selection.Mapping().Args))
	}
}

func TestSessionHostShortcutCollisionsAndOverflow(t *testing.T) {
	entries := make([]providerRow, 0, 40)
	for index := range 40 {
		entries = append(entries, groupedSessionEntry(t, fmt.Sprintf("gamingrig-%d", index), index))
	}
	rows, err := providerRows(Mapping{Keys: "s", Provider: menuProviderSessions}, entries)
	require.NoError(t, err)
	require.NoError(t, ValidateMappings(rows))
	choices := 0
	for _, row := range rows {
		if row.Action == menuSessionSelect {
			choices++
		}
	}
	require.Equal(t, 40, choices)
	menu, err := NewMenu(rows)
	require.NoError(t, err)
	menu.HandleKey("s", 1)
	require.True(t, menu.QuitWithQ())
	var more bool
	for _, entry := range menu.Entries() {
		if entry.Key == "n" {
			more = true
		}
	}
	require.True(t, more, "many hosts must stay reachable through More hosts")
}

func TestGroupedSessionKeyBudget(t *testing.T) {
	entries := []providerRow{groupedSessionEntry(t, "gamingrig", 0)}
	_, err := providerRows(Mapping{Keys: strings.Repeat("s", menuConfigLimit/2), Provider: menuProviderSessions}, entries)
	require.Error(t, err, "admit the complete grouped key budget before generating long keys")
}
