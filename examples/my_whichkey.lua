-- User options module: place beside the Rex SERVER's init.lua, not on the client.
-- Do not name this file rex_whichkey.lua: that namespace belongs to integration.
-- The data shape and mapping rules are owned by docs/mapping-format.md.
return {
    position = "bottom",
    title = "My Rex",
    spec = {
        { "w", group = "Sessions" },
        { "wn", "session.new", desc = "New session" },
        { "z", "pane.zoom", desc = "Zoom pane" },
        { "x", "pane.close", desc = "Close pane", confirm = true },
    },
}
