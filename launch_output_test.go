package whichkey

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLaunchOutputRecognizesOnlyFixedLegacyDiagnostic(t *testing.T) {
	const diagnostic = "rex_whichkey: launch origin client_id does not match the actual action context"
	for split := 0; split <= len(diagnostic); split++ {
		output := &launchOutput{remaining: popupMessageLimit}
		_, err := output.Write([]byte("private arbitrary output\x1b[31m\nError: " + diagnostic[:split]))
		require.NoError(t, err)
		_, err = output.Write([]byte(diagnostic[split:] + "\nprivate suffix"))
		require.NoError(t, err)
		require.True(t, output.legacyClientBinding)
	}
	output := &launchOutput{remaining: popupMessageLimit}
	_, err := output.Write([]byte(strings.Repeat("a", 64) + "\x1b[31m arbitrary child error"))
	require.NoError(t, err)
	require.False(t, output.legacyClientBinding)
}

func TestLaunchOutputBudget(t *testing.T) {
	output := &launchOutput{remaining: 3}
	count, err := output.Write([]byte("abc"))
	require.NoError(t, err)
	require.Equal(t, 3, count)
	count, err = output.Write([]byte("d"))
	require.ErrorIs(t, err, io.ErrShortWrite)
	require.Zero(t, count)
}
