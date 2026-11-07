-- Raw spec fragment consumed by rex_whichkey.mapping under the Go-owned
-- mapping.go/docs/mapping-format.md rules. The defaults composer appends it to
-- its spec; importing this module neither registers nor executes Rex actions.
--
-- Keep these as client actions: the authenticated dispatcher restores the origin
-- and targets that app explicitly. The default wr rename intent instead uses
-- WhichKey's input page and the retained session's set_label resource method.
-- session.switch delegates the picker (and its cancellation) to Rex; it is not
-- an enumerated submenu or a request to attach to a remote host.
return {
    { "w", group = "Sessions", icon = "󰙅" },
    { "wn", "session.new", desc = "New session", icon = "" },
    { "ws", "session.switch", desc = "Switch session", icon = "󰒺" },
    { "wh", "session.previous", desc = "Previous", icon = "" },
    { "wl", "session.next", desc = "Next", icon = "" },
    { "wr", "session.rename", desc = "Rename", icon = "󰏫" },
    { "wx", "session.close", desc = "Close session", icon = "󰅖", confirm = true },
}
