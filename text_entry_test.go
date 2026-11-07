package whichkey

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func renameEditor(t *testing.T, action, initial string) *Menu {
	t.Helper()
	menu, err := NewMenu([]Mapping{{Keys: "r", Action: action, Desc: "Edit name"}})
	require.NoError(t, err)
	menu.textSpec = &textEntrySpec{Initial: initial, SessionInitial: initial}
	menu.HandleKey("r", 1)
	require.NotNil(t, menu.entry)
	return menu
}

func TestRenameBackspaceDeletesOneCharacter(t *testing.T) {
	for _, action := range []string{menuTabRename, menuSessionRename} {
		t.Run(action, func(t *testing.T) {
			for _, key := range []string{"\b", "\x7f"} {
				t.Run(key, func(t *testing.T) {
					for _, test := range []struct {
						name, initial, first, second string
					}{
						{"ASCII", "alpha", "alph", "alp"},
						{"Unicode", "café", "caf", "ca"},
						{"combining character", "Cafe\u0301", "Caf", "Ca"},
						{"emoji cluster", "name\U0001f468\u200d\U0001f469\u200d\U0001f467\u200d\U0001f466", "name", "nam"},
						{"single character", "界", "", ""},
						{"empty", "", "", ""},
					} {
						t.Run(test.name, func(t *testing.T) {
							menu := renameEditor(t, action, test.initial)
							require.Nil(t, menu.HandleKey(key, 1).Selection)
							require.Equal(t, test.first, menu.entry.value)
							require.Equal(t, len(test.first), menu.entry.cursor)
							require.True(t, menu.entry.touched)
							require.False(t, menu.entry.selected)
							require.Nil(t, menu.HandleKey(key, 1).Selection)
							require.Equal(t, test.second, menu.entry.value)
							require.False(t, menu.Done())
						})
					}
				})
			}
		})
	}
}

func TestRenameDeletionAtCursorAndExplicitClear(t *testing.T) {
	menu := renameEditor(t, menuTabRename, "abcd")
	menu.HandleKey(terminalDeleteKey, 1)
	require.Equal(t, "abcd", menu.entry.value, "Delete at the end must not clear the prefilled name")
	menu.HandleKey("\x1b[H", 1)
	menu.HandleKey("\x1b[C", 1)
	menu.HandleKey("\x1b[C", 1)
	menu.HandleKey("\x7f", 1)
	require.Equal(t, "acd", menu.entry.value)
	require.Equal(t, 1, menu.entry.cursor)
	menu.HandleKey(terminalDeleteKey, 1)
	require.Equal(t, "ad", menu.entry.value)
	menu.HandleKey("\x15", 1)
	require.Empty(t, menu.entry.value, "Ctrl-U remains the explicit clear-all key")
	require.Zero(t, menu.entry.cursor)
}
