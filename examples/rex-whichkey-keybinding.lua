-- Example only: choose your own key in your user-owned Rex init.lua.
-- The execution host must have rex-whichkey on the Rex process's PATH.
local whichkey = require("rex_whichkey")

rex.bind("ctrl+space", function(ctx)
    return whichkey.launch(ctx)
end)
