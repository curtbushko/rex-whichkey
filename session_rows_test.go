package whichkey

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSessionRowsHideIdentity(t *testing.T) {
	const sessionID = "session:00000000-0000-0000-0000-000000000001"
	mapping, err := sessionRowMapping(sessionID, "")
	require.NoError(t, err)
	for _, marker := range []Mapping{
		{Keys: "ws", Action: menuSessionSwitch},
		{Keys: "s", Provider: menuProviderSessions},
	} {
		rows, rowsErr := providerRows(marker, []providerRow{
			{id: sessionID, label: "[current] project", mapping: mapping},
		})
		require.NoError(t, rowsErr)
		require.Len(t, rows, 2)
		require.Equal(t, "This Mac (1)", *rows[0].Group)
		require.Equal(t, marker.Keys+"l1", rows[1].Keys)
		require.Equal(t, "[current] project", rows[1].Desc)
		require.JSONEq(t, `{"session_id":"`+sessionID+`"}`, string(rows[1].Args))
	}
}

func TestOtherResourceRowsKeepIdentity(t *testing.T) {
	const windowID = "window:00000000-0000-0000-0000-000000000001"
	rows, err := providerRows(Mapping{Keys: "t", Provider: menuProviderWindows}, []providerRow{
		{id: windowID, label: "Editor tab"},
	})
	require.NoError(t, err)
	require.Len(t, rows, 1)
	require.Equal(t, "["+windowID+"] Editor tab", rows[0].Desc)
}
