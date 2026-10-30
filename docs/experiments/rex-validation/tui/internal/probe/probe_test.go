package probe_test

import (
	"encoding/json"
	"testing"

	"example.com/whichkey-research/internal/probe"
	"github.com/stretchr/testify/require"
)

func TestMenu(t *testing.T) {
	for _, tc := range []struct {
		key, selection string
		page           int
	}{
		{"tab", "", 1}, {"backspace", "", 0}, {"esc", "cancel", 0}, {"q", "cancel", 0}, {"y", "choose", 0},
	} {
		t.Run(tc.key, func(t *testing.T) {
			m := probe.Menu{}
			m.Key(tc.key)
			require.Equal(t, tc.selection, m.Selection)
			require.Equal(t, tc.page, m.Page)
		})
	}
}

func TestDispatchArgumentsAreData(t *testing.T) {
	args, err := probe.DispatchArgs("unix:///tmp/probe.sock", "session:origin", "block:origin", "$(touch /tmp/never); 'quoted'", "choose")
	require.NoError(t, err)
	require.Equal(t, []string{"-S", "unix:///tmp/probe.sock", "--autostart=false", "do", "-s", "session:origin", "-b", "block:origin", "wk_probe_dispatch", "--args"}, args[:10])
	var payload struct {
		ID        string `json:"id"`
		Selection string `json:"selection"`
		Confirmed bool   `json:"confirmed"`
	}
	require.NoError(t, json.Unmarshal([]byte(args[10]), &payload))
	require.Equal(t, "$(touch /tmp/never); 'quoted'", payload.ID)
	require.Equal(t, "choose", payload.Selection)
	require.True(t, payload.Confirmed)
}
