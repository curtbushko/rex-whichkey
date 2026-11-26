package whichkey

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPopupFooterAnchorsToBottomWithoutMovingMenu(t *testing.T) {
	config := DefaultConfig()
	config.Position, config.footerAtBottom = positionTop, true
	group := "Actions"
	config.Spec = []Mapping{{Keys: "g", Group: &group}, {Keys: "gx", Action: menuPaletteToggle, Desc: "Popup action"}}
	menu, err := NewMenu(config.Spec)
	require.NoError(t, err)
	theme := NewTheme(TerminalPalette{}, true, false)
	for _, preferred := range []int{4, 6} {
		for _, height := range []int{8, 14, 24} {
			frame := RenderMenu(menu, config, theme, 100, height)
			require.Zero(t, frame.Y)
			require.Equal(t, height, frame.Height)
			require.Len(t, frame.Lines, height)
			require.Empty(t, frame.Lines[0])
			require.NotEmpty(t, frame.Lines[1], "menu/heading stays at the top")
			require.Contains(t, frame.Lines[height-1], "esc/backspace", "instructions occupy the final row, with no bottom padding")
			require.Contains(t, frame.Lines[height-1], "rex-whichkey")
			require.Empty(t, frame.Lines[height-2], "blank space separates menu and instructions")
			require.Equal(t, preferred, frame.PreferredHeight, "footer anchoring must not inflate cached/fitted height")
		}
		menu.HandleKey("g", 1)
	}
}

func TestPopupFooterReservesGapWhilePaging(t *testing.T) {
	config := DefaultConfig()
	config.Position, config.footerAtBottom = positionTop, true
	for _, key := range "abcdefghijklmnopqrst" {
		config.Spec = append(config.Spec, Mapping{Keys: string(key), Action: menuPaletteToggle, Desc: "Entry with a long label"})
	}
	menu, err := NewMenu(config.Spec)
	require.NoError(t, err)
	theme := NewTheme(TerminalPalette{}, true, false)
	frame := RenderMenu(menu, config, theme, 40, 8)
	require.Greater(t, frame.Pages, 1)
	require.Empty(t, frame.Lines[6], "a full page still reserves menu/footer separation")
	require.Contains(t, frame.Lines[7], "Tab")
	menu.HandleKey("\t", frame.Pages)
	next := RenderMenu(menu, config, theme, 40, 8)
	require.Equal(t, frame.Pages, next.Pages)
	require.Equal(t, 1, next.Page)
	for height := 1; height <= 5; height++ {
		frame := RenderMenu(menu, config, theme, 40, height)
		require.Len(t, frame.Lines, height)
		require.Contains(t, strings.Join(frame.Lines, "\n"), "Entry with a long label")
	}
}
