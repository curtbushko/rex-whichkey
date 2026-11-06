package whichkey

import "encoding/json"

// nativeDefaults supplies the current reduced menu without a custom file or Lua
// registration. session.switch remains a Go sessions-menu intent for expansion
// by the consumer; loading defaults performs no inventory query.
func nativeDefaults() (Config, error) {
	group := func(keys, title, icon string) Mapping {
		return Mapping{Keys: keys, Group: &title, Icon: icon}
	}
	config := DefaultConfig()
	config.Spec = []Mapping{
		group("w", "Sessions", "󰙅"),
		{Keys: "wn", Action: menuSessionNew, Desc: "New session", Icon: ""},
		{Keys: "ws", Action: menuSessionSwitch, Desc: "Switch session", Icon: "󰒺"},
		{Keys: "wh", Action: menuSessionPrevious, Desc: "Previous", Icon: ""},
		{Keys: "wl", Action: menuSessionNext, Desc: "Next", Icon: ""},
		{Keys: "wr", Action: menuSessionRename, Desc: "Rename", Icon: "󰏫"},
		{Keys: "wx", Action: menuSessionClose, Desc: "Close session", Icon: "󰅖", Confirm: true},
		group("t", "Tabs", ""),
		{Keys: "tn", Action: menuTabNew, Desc: "New tab", Icon: ""},
		{Keys: "th", Action: menuTabPrevious, Desc: "Previous", Icon: ""},
		{Keys: "tl", Action: menuTabNext, Desc: "Next", Icon: ""},
		{Keys: "tr", Action: menuTabRename, Desc: "Rename", Icon: "󰏫"},
		{Keys: "tx", Action: menuTabClose, Desc: "Close tab", Icon: "󰅖", Confirm: true},
		{Keys: "v", Action: menuPaneSplit, Desc: "Split right", Icon: "", Args: json.RawMessage(`{"direction":"right"}`)},
		{Keys: "s", Action: menuPaneSplit, Desc: "Split down", Icon: "", Args: json.RawMessage(`{"direction":"down"}`)},
		{Keys: "x", Action: menuPaneClose, Desc: "Close pane", Icon: "󰅖", Confirm: true},
		{Keys: "e", Action: menuSidebarToggle, Desc: "Sidebar", Icon: "󰙀"},
		{Keys: "c", Action: menuPaletteToggle, Desc: "Command palette", Icon: ""},
	}
	return ResolveConfig(config, nil, nil)
}
