-- Example only: choose your own key in your user-owned Rex init.lua.
-- The execution host must have rex-whichkey on the Rex process's PATH.
-- The shared module creates one float in ctx's exact window and starts Go
-- directly. No user config reload is needed for the development entry:
-- rex do "$PWD/lua/rex_whichkey_launch.lua" --action rex_whichkey_launch development=true
-- The script loads the sibling shared module without returning its function table.
-- Width/height are finite fractions; native chrome uses Rex defaults.
local whichkey = require("rex_whichkey")

rex.bind("ctrl+space", function(ctx)
    return whichkey.launch(ctx, { width = .8, height = .3 })
end)
