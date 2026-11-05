package.path = "lua/?.lua;lua/?/init.lua;" .. package.path

local session = "session:11111111-1111-1111-1111-111111111111"
local block = "block:22222222-2222-2222-2222-222222222222"
local app = "client:33333333-3333-3333-3333-333333333333"
local window = "window:44444444-4444-4444-4444-444444444444"
local launcher = "client:55555555-5555-5555-5555-555555555555"
local completer = "client:66666666-6666-6666-6666-666666666666"
local token = string.rep("a", 64)
local generation, state, app_kind = "generation-1", "connected", "app"
local actions, effects, launches, callbacks = {}, 0, 0, 0
local now = os.time()
os.time = function() return now end
local launch_ctx, completion_ctx
local selected_action = "pane.split"
local enabled = true
local function context(client)
    return { session_id = session, block_id = block, client_id = client }
end
local expected = { server = "unix:///tmp/rex.sock", session_id = session, block_id = block,
    client_id = app, connected_at = generation }
local mutate, rpc_error, rpc_wrong_client
rex = {
    action = function(action) actions[action.name] = action.run end,
    session = {
        view = function()
            if mutate == "caller" then completion_ctx.client_id = launcher end
            return { session_id = session, active_window_id = window, windows = {
                { window_id = window, focused_block_id = block,
                  layers = { { blocks = { { block_id = block } } } } },
            } }
        end,
        list_clients = function()
            return { clients = { { client_id = app, connected_at = generation,
                connection_state = state, info = { kind = app_kind } } } }
        end,
        focus_block = function() return {} end,
    },
    json = {
        encode = function(value)
            if value.completion_action then
                assert(value.client_id == app, "controller must retain GUI target, not CLI caller")
            end
            return "{}"
        end,
        decode = function(text)
            if text == "origin" then return expected end
            return { capability = token, executable = "/bin/rex-whichkey", server = expected.server }
        end,
    },
    call = function(method, args)
        assert(method == "client.perform", "app actions must use explicit routing")
        assert(args.client_id == app and args.session_id == session and args.block_id == block)
        assert(args.action == selected_action)
        effects = effects + 1
        if rpc_error then return nil, "refused" end
        return { performed_by = rpc_wrong_client and launcher or app }
    end,
    client = { queue = function() error("must not queue on the transient CLI") end },
    invoke = function(name) assert(name == "named_action"); callbacks = callbacks + 1 end,
}
rex.exec = setmetatable({ start = function() launches = launches + 1; return {} end }, {
    __call = function()
        if mutate == "launch" then launch_ctx.client_id = completer end
        return { ok = true, stdout = "prepared" }
    end,
})
local popup = require("rex_whichkey")
popup.register_completion("complete")
local callback = function(ctx)
    assert(ctx.client_id == app and ctx.request_client_id == completer)
    callbacks = callbacks + 1
end
local target = scenario == "callback" and callback or (scenario == "named" and "named_action" or selected_action)
local menu = popup.new({ position = "bottom", title = "Keys", icon = "", icons = {},
    spec = { { "x", target, desc = "Action", cond = function() return enabled end,
        confirm = scenario == "unconfirmed" or scenario == "confirmed" } } })
menu.setup()
menu.register_launch("launch")
launch_ctx, completion_ctx = context(launcher), context(completer)
if scenario == "non_app" then
    app_kind = "cli"
    assert(not pcall(actions.launch, launch_ctx, { origin = "origin" }))
    assert(launches == 0 and effects == 0)
    return
end
if scenario == "no_target" then
    assert(not pcall(actions.launch, launch_ctx, {}), "CLI without explicit GUI binding must fail")
    assert(launches == 0)
    return
end
if scenario == "launch_caller_mutation" then
    mutate = "launch"
    assert(not pcall(actions.launch, launch_ctx, { origin = "origin" }))
    assert(launches == 0 and effects == 0)
    return
end
-- This is the regression: actual CLI identity must remain intact, while the
-- controller and retained selection bind the independently validated app.
if scenario == "key_launch" then
    launch_ctx = context(app)
    actions.launch(launch_ctx, {})
else
    actions.launch(launch_ctx, { origin = "origin" })
end
assert(launch_ctx.client_id == (scenario == "key_launch" and app or launcher) and launches == 1)
local selection = '"keys":"x","confirmed":' .. (scenario == "confirmed" and "true" or "false")
if scenario == "callback" then selection = selection .. ',"callback_ref":"cb:1:1"' end
local message = '{"version":1,"capability":"' .. token .. '","kind":"complete","selection":{' .. selection .. '}}'
local function complete(ctx, raw) actions.complete(ctx, { message = raw or message }) end
local rejected = {
    wrong_session = "session_id", wrong_block = "block_id", wrong_server = "server",
    invalid_caller = "client_id", wrong_capability = "capability", malformed = "json",
}
if rejected[scenario] then
    local field = rejected[scenario]
    local wrong, raw = context(completer), message
    if field == "capability" then raw = message:gsub(token, string.rep("b", 64))
    elseif field == "json" then raw = message:sub(1, -2)
    elseif field == "server" then wrong.server = "another-server"
    elseif field == "client_id" then wrong.client_id = "label"
    else wrong[field] = field:gsub("_id$", "") .. ":77777777-7777-7777-7777-777777777777" end
    bad_completion = function() complete(wrong, raw) end
end
function finish_test()
    if scenario == "reconnect" then generation = "generation-2" end
    if scenario == "lost" then state = "lost" end
    if scenario == "kind_changed" then app_kind = "cli" end
    if scenario == "expired" then now = now + 31 end
    if scenario == "reload" then menu.setup() end
    if scenario == "caller_mutation" then mutate = "caller" end
    if scenario == "guard_changed" then enabled = false end
    if scenario == "rpc_error" then rpc_error = true end
    if scenario == "rpc_wrong_client" then rpc_wrong_client = true end
    local invalid = { reconnect=true, lost=true, kind_changed=true, expired=true, reload=true, caller_mutation=true,
        rpc_error=true, rpc_wrong_client=true, guard_changed=true, unconfirmed=true }
    if invalid[scenario] then
        assert(not pcall(complete, completion_ctx))
        assert(effects == ((rpc_error or rpc_wrong_client) and 1 or 0))
        assert(callbacks == 0)
    elseif scenario == "cancel" then
        complete(completion_ctx, '{"version":1,"capability":"' .. token .. '","kind":"cancel"}')
        assert(effects == 0 and callbacks == 0)
    else
        complete(completion_ctx)
        local lua_action = scenario == "callback" or scenario == "named"
        assert(effects == (lua_action and 0 or 1), "effect count for " .. tostring(scenario))
        assert(callbacks == (lua_action and 1 or 0))
    end
    local before = effects
    assert(not pcall(complete, context(completer)), "completion must never replay")
    assert(effects == before)
end
