package whichkey

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMenuHeader(t *testing.T) {
	for _, test := range []struct {
		name   string
		prefix string
		title  string
		icon   string
		want   string
	}{
		{name: "root separator only"},
		{name: "submenu breadcrumb", prefix: "t", want: "Tabs"},
		{name: "nested breadcrumbs", prefix: "ts", want: "Tabs » Split"},
		{name: "explicit custom heading", title: "Custom", icon: "*", want: "* Custom"},
	} {
		t.Run(test.name, func(t *testing.T) {
			config := DefaultConfig()
			if test.title != "" {
				config.Title, config.Icon = test.title, test.icon
			}
			tabs, split := "Tabs", "Split"
			config.Spec = []Mapping{
				{Keys: "t", Group: &tabs},
				{Keys: "ts", Group: &split},
				{Keys: "tsr", Action: menuPaletteToggle, Desc: "Run"},
			}
			menu, err := NewMenu(config.Spec)
			require.NoError(t, err)
			for _, key := range test.prefix {
				require.Nil(t, menu.HandleKey(string(key), 1).Selection)
			}
			theme := NewTheme(TerminalPalette{}, true, false)
			require.Equal(t, test.want, menuHeader(menu, config, theme).plain())
			frame := RenderMenu(menu, config, theme, 80, 10)
			require.NotEmpty(t, frame.Lines)
			if test.want == "" {
				require.Equal(t, strings.Repeat("─", frame.Width), frame.Lines[0])
			} else {
				require.Contains(t, frame.Lines[0], test.want)
			}
		})
	}
}
