package whichkey

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMenuOuterSpacingAndPreferredHeight(t *testing.T) {
	config := DefaultConfig()
	group := "Actions"
	config.Spec = []Mapping{{Keys: "g", Group: &group}, {Keys: "gx", Action: menuPaletteToggle, Desc: "Palette"}}
	menu, err := NewMenu(config.Spec)
	require.NoError(t, err)
	theme := NewTheme(TerminalPalette{}, true, false)
	for _, preferred := range []int{5, 7} {
		frame := RenderMenu(menu, config, theme, 100, 20)
		require.Empty(t, frame.Lines[0])
		require.Empty(t, frame.Lines[len(frame.Lines)-1])
		require.NotEmpty(t, frame.Lines[1], "exactly one outer top blank row")
		require.NotEmpty(t, frame.Lines[len(frame.Lines)-2], "footer precedes the bottom blank row")
		require.Equal(t, preferred, frame.PreferredHeight)
		require.Len(t, frame.Lines, frame.Height)
		menu.HandleKey("g", frame.Pages)
	}
	for height := 1; height <= 4; height++ {
		frame := RenderMenu(menu, config, theme, 100, height)
		require.LessOrEqual(t, frame.Height, height)
		require.Len(t, frame.Lines, frame.Height)
		require.Contains(t, frame.Lines, "x  Palette", "padding must not hide the only content row")
	}
}
