-- User options module: place beside the Rex SERVER's init.lua, not on the client.
-- Do not name this file rex_whichkey.lua: that namespace belongs to integration.
-- In your own init.lua (not in this returned options module):
-- local popup = require("rex_whichkey")
-- popup.register_completion("whichkey_complete")
-- local menu = popup.new(require("rex_whichkey.defaults"))
-- menu.setup({}, "my_whichkey")
-- menu.register_launch("whichkey")
-- No helper/rex_bin_dir constructor fields are needed or accepted.
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
