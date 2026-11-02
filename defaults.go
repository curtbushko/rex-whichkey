package whichkey

import "encoding/json"

// nativeDefaults supplies the basic menu without requiring a server Lua action.
// It uses the canonical Mapping/Config types and the same keys and native action
// contracts as lua/rex_whichkey/defaults. Callbacks never enter this route.
func nativeDefaults() (Config, error) {
	group := func(keys, title, icon string) Mapping {
		return Mapping{Keys: keys, Group: &title, Icon: icon}
	}
	config := DefaultConfig()
	config.Spec = []Mapping{
		group("w", "Sessions", "󰙅"),
		{Keys: "wn", Action: "session.new", Desc: "New session", Icon: ""},
		{Keys: "ws", Action: "session.switch", Desc: "Switch session", Icon: "󰒺"},
		{Keys: "wh", Action: "session.previous", Desc: "Previous", Icon: ""},
		{Keys: "wl", Action: "session.next", Desc: "Next", Icon: ""},
		{Keys: "wr", Action: "session.rename", Desc: "Rename", Icon: "󰏫"},
		{Keys: "wx", Action: "session.close", Desc: "Close session", Icon: "󰅖", Confirm: true},
		group("t", "Tabs", ""),
		{Keys: "tn", Action: "client.tab.new", Desc: "New tab", Icon: ""},
		{Keys: "th", Action: "client.tab.previous", Desc: "Previous", Icon: ""},
		{Keys: "tl", Action: "client.tab.next", Desc: "Next", Icon: ""},
		{Keys: "tr", Action: "client.tab.rename", Desc: "Rename", Icon: "󰏫"},
		{Keys: "tx", Action: "client.tab.close", Desc: "Close tab", Icon: "󰅖", Confirm: true},
		{Keys: "v", Action: "pane.split", Desc: "Split right", Icon: "", Args: json.RawMessage(`{"direction":"right"}`)},
		{Keys: "s", Action: "pane.split", Desc: "Split down", Icon: "", Args: json.RawMessage(`{"direction":"down"}`)},
		{Keys: "x", Action: "pane.close", Desc: "Close pane", Icon: "󰅖", Confirm: true},
		{Keys: "e", Action: "client.sidebar.toggle", Desc: "Sidebar", Icon: "󰙀"},
		{Keys: "c", Action: "client.palette.toggle", Desc: "Command palette", Icon: ""},
	}
	return ResolveConfig(config, nil, nil)
}
