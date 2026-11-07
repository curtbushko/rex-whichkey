-- Production rex do entry; use an absolute path so the sibling module and
-- development binary come from this checkout, not require's installed search path.
-- After task build:
--   rex do "$PWD/lua/rex_whichkey_launch.lua" --action rex_whichkey_launch development=true
-- Optional action data: width=.8 height=.3 config=/path/menu.json ascii=true no_color=true
-- Select an owned origin with -s SESSION -b BLOCK and -C CLIENT when needed.
local source = debug.getinfo(1, "S").source
local directory = source:match("^@(/.*/)")
if not directory then
    error("rex_whichkey: launch script requires an absolute file path", 0)
end

-- Loading registers the shared launch action. Discard M: rex do serializes
-- script returns, and M contains the launch function.
-- The module still returns M for installed require(...).launch callers.
dofile(directory .. "rex_whichkey.lua")
