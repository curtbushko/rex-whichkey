package whichkey

import (
	"fmt"
	"strings"
	"testing"

	"github.com/rivo/uniseg"
	"github.com/stretchr/testify/require"
)

func TestMenuHasNoDrawnRulesAndFooterHasIdentity(t *testing.T) {
	config := DefaultConfig()
	config.Spec = []Mapping{{Keys: "x", Action: menuPaletteToggle, Desc: "Palette"}}
	menu, err := NewMenu(config.Spec)
	require.NoError(t, err)
	for _, ascii := range []bool{false, true} {
		frame := RenderMenu(menu, config, NewTheme(TerminalPalette{}, true, ascii), 100, 20)
		text := strings.Join(frame.Lines, "\n")
		require.NotContains(t, text, "─")
		require.NotContains(t, text, strings.Repeat("-", 8))
		require.NotEmpty(t, frame.Lines[1], "an empty title must not reserve blank header rows beyond outer padding")
		footer := frame.Lines[len(frame.Lines)-2]
		require.Contains(t, footer, "esc/backspace: quit")
		require.True(t, strings.HasSuffix(footer, "rex-whichkey"))
	}
}

func TestFooterKeepsNavigationAndPagingAtNarrowWidths(t *testing.T) {
	config := DefaultConfig()
	for i, key := range "abcdefghijklmnopqrst" {
		config.Spec = append(config.Spec, Mapping{Keys: string(key), Action: menuPaletteToggle, Desc: fmt.Sprintf("Entry %d", i)})
	}
	menu, err := NewMenu(config.Spec)
	require.NoError(t, err)
	for _, width := range []int{1, 3, 8, 12, 20, 40, 80} {
		frame := RenderMenu(menu, config, NewTheme(TerminalPalette{}, true, false), width, 3)
		for _, line := range frame.Lines {
			require.LessOrEqual(t, uniseg.StringWidth(line), width)
		}
		footer := frame.Lines[len(frame.Lines)-1]
		if width >= 3 {
			require.Contains(t, strings.ToLower(footer), "esc")
		}
		if frame.Pages > 1 && width >= 8 {
			require.Contains(t, footer, "Tab")
		}
	}
}

func TestFooterKeepsWarningsCriticalAndIdentityMuted(t *testing.T) {
	t.Setenv("NO_COLOR", "")
	config := DefaultConfig()
	config.Spec = []Mapping{{Keys: "x", Action: menuPaneClose, Desc: "Close pane", Confirm: true}}
	menu, err := NewMenu(config.Spec)
	require.NoError(t, err)
	menu.HandleKey("x", 1)
	theme := NewTheme(TerminalPalette{}, false, false)
	frame := RenderMenu(menu, config, theme, 100, 10)
	footer := frame.Lines[len(frame.Lines)-2]
	require.Contains(t, footer, "\x1b[31my: confirm")
	require.Contains(t, footer, "\x1b[90m")
	require.Contains(t, footer, "rex-whichkey")
}

func TestMutedHintsUseTerminalNeutralInsteadOfBodyColor(t *testing.T) {
	const black, white = "#000000", "#ffffff"
	t.Setenv("NO_COLOR", "")
	for _, palette := range []TerminalPalette{
		{Foreground: white, Background: black},
		{Foreground: black, Background: white},
		{Foreground: white, Background: black, Colors: [16]string{7: white, 8: "#222222"}},
	} {
		theme := NewTheme(palette, false, false)
		require.NotEqual(t, theme.Style(ThemePrimary, "esc/backspace"), theme.Style(ThemeMuted, "esc/backspace"))
		muted := theme.accents[ThemeMuted-ThemeKey]
		require.GreaterOrEqual(t, themeContrast(*muted, *theme.background), 4.5)
		require.Less(t, themeContrast(*muted, *theme.background), themeContrast(*theme.foreground, *theme.background))
	}
	palette := TerminalPalette{Foreground: white, Background: black, Colors: [16]string{8: "#888888"}}
	theme := NewTheme(palette, false, false)
	require.Contains(t, theme.Style(ThemeMuted, "esc/backspace"), "\x1b[38;2;136;136;136m")
	theme = NewTheme(TerminalPalette{}, false, false)
	require.Contains(t, theme.Style(ThemeMuted, "esc/backspace"), "\x1b[90m")
	require.Equal(t, "esc/backspace", NewTheme(TerminalPalette{}, true, false).Style(ThemeMuted, "esc/backspace"))
}
