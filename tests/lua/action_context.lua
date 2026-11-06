package.path = "tests/fixtures/phase15-lua/?.lua;tests/fixtures/phase15-lua/?/init.lua;" .. package.path

local session = "session:11111111-1111-1111-1111-111111111111"
local block = "block:22222222-2222-2222-2222-222222222222"
local client = "client:33333333-3333-3333-3333-333333333333"
local window = "window:44444444-4444-4444-4444-444444444444"
local capability = string.rep("a", 64)
local actions, launches = {}, 0
rex = {
    action = function(action) actions[action.name] = action.run end,
    session = {
        view = function()
            return { session_id = session, windows = {
                { window_id = window, layers = { { blocks = { { block_id = block } } } } },
            } }
        end,
        list_clients = function()
            return { clients = { { client_id = client, connected_at = "generation-1", info = { kind = "app" } } } }
        end,
    },
    json = {
        encode = function() return "{}" end,
        decode = function(text)
            if text == "expected-origin" then return launch_expected end
            return { capability = capability, executable = "/bin/rex-whichkey", server = "unix:///tmp/rex.sock" }
        end,
    },
}
rex.exec = setmetatable({
    start = function(argv, opts)
        assert(argv[1] == "/bin/rex-whichkey")
        assert(opts.env.REX_BIN_DIR == "/opt/rex/bin")
        launches = launches + 1
        return {}
    end,
}, { __call = function(_, argv, opts)
    assert(argv[1] == "/opt/whichkey/bin/rex-whichkey")
    assert(opts.env.REX_BIN_DIR == "/opt/rex/bin")
    return { ok = true, stdout = "prepared" }
end })

local popup = require("rex_whichkey")
popup.register_completion("whichkey_complete")
local menu = popup.new({ position = "bottom", title = "Keys", icon = "", icons = {}, spec = {} }, nil, {
    helper = "/opt/whichkey/bin/rex-whichkey",
    rex_bin_dir = "/opt/rex/bin",
})
menu.setup()
menu.register_launch("whichkey")

local function context(server)
    return { session_id = session, block_id = block, client_id = client, server = server }
end
if test_origin_mismatch then
    launch_expected = {
        server = "unix:///tmp/rex.sock", session_id = session, block_id = block,
        client_id = client, connected_at = "generation-1",
    }
    local field = test_origin_mismatch
    if field == "connected_at" then
        launch_expected[field] = "older-generation"
    else
        local kind = field:gsub("_id$", "")
        launch_expected[field] = kind .. ":55555555-5555-5555-5555-555555555555"
    end
    local ok, err = pcall(actions.whichkey, context(nil), { origin = "expected-origin" })
    assert(not ok, "mismatched origin must fail")
    if field == "connected_at" then
        assert(err:find("originating client connection changed", 1, true), err)
    elseif field == "client_id" then
        assert(err:find("originating client connection is missing or ambiguous", 1, true), err)
    else
        assert(err:find("launch origin " .. field .. " does not match", 1, true), err)
    end
    assert(launches == 0, "preflight failure must not launch")
    return
end
local cancel = '{"version":1,"capability":"' .. capability .. '","kind":"cancel"}'

-- Installed Rex omits server even with all three exact resource IDs. Newer
-- contexts may include an opaque identity; both must survive fresh activation.
do
    local server
    if test_server_kind == "string" then server = "opaque-server" end
    if test_server_kind == "table" then server = {} end
    local before = launches
    actions.whichkey(context(server), {})
    assert(launches == before + 1, "context must reach detached launch: " .. before .. " -> " .. launches)
    local wrong = context(server)
    wrong.block_id = "block:55555555-5555-5555-5555-555555555555"
    assert(not pcall(actions.whichkey_complete, wrong, { message = cancel }))
    local changed_server = context({})
    assert(not pcall(actions.whichkey_complete, changed_server, { message = cancel }))
    actions.whichkey_complete(context(server), { message = cancel })
    assert(not pcall(actions.whichkey_complete, context(server), { message = cancel }), "no replay")
end

-- Optional server does not make resource identity optional or accept malformed
-- supplied identities. Rejected launches must never start a controller.
for _, field in ipairs({ "session_id", "block_id", "client_id" }) do
    local ctx = context(nil)
    ctx[field] = nil
    local before = launches
    assert(not pcall(actions.whichkey, ctx, {}))
    assert(launches == before)
end
for _, server in ipairs({ false, 12, "" }) do
    assert(not pcall(actions.whichkey, context(server), {}))
end
