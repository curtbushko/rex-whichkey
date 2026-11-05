package whichkey

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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
