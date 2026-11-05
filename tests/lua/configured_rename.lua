package.path = "lua/?.lua;lua/?/init.lua;" .. package.path
local mapping = require("rex_whichkey.mapping")
local dispatch = require("rex_whichkey.dispatch")
local session = "session:11111111-1111-1111-1111-111111111111"
local block = "block:22222222-2222-2222-2222-222222222222"
local app = "client:33333333-3333-3333-3333-333333333333"
local window = "window:44444444-4444-4444-4444-444444444444"
local original_window = window
local caller = "client:55555555-5555-5555-5555-555555555555"
local token = string.rep("a", 64)
local actions, effects, callbacks = {}, 0, 0
local keys = scenario == "session" and "wr" or "tr"
local action = keys == "wr" and "session.rename" or "client.tab.rename"
rex = {
    action = function(def) actions[def.name] = def.run end,
    client = { queue = function() error("rename must not open native UI") end },
    call = function() error("rename must not use client.perform") end,
    session = {
        view = function()
            return { session_id = session, active_window_id = window, windows = {
                { window_id = window, focused_block_id = block,
                  layers = { { blocks = { { block_id = block } } } } },
            } }
        end,
        list_clients = function()
            return { clients = { { client_id = app, connected_at = "generation", info = { kind = "app" } } } }
        end,
        focus_block = function() return {} end,
        set_label = function(args)
            assert(keys == "wr" and args.session_id == session and args.label == 'New "name"')
            effects = effects + 1
            return { revision = 2 }
        end,
        set_window_label = function(args)
            assert(keys == "tr" and args.session_id == session and args.window_id == original_window)
            assert(args.label == 'New "name"')
            effects = effects + 1
            if scenario == "uncertain" then return nil, "lost reply" end
            return { revision = 2 }
        end,
    },
}
local target = action
if scenario == "callback" or scenario == "injected_callback" then
    target = function() callbacks = callbacks + 1 end
end
local record = { keys, target, desc = "Rename" }
if scenario == "custom_args" then record.args = { label = "retained" } end
local registry = dispatch.new(mapping.compile({ record }))
dispatch.register_completion("complete")
local ctx = { session_id = session, block_id = block, client_id = caller }
local records, ticket = registry.issue(ctx, { client_id = app, connected_at = "generation" })
registry.bind(ticket, ctx, token)
local selection = '"keys":"' .. keys .. '","confirmed":false'
if target ~= action then
    local reference
    for _, item in ipairs(records) do
        if item.keys == keys then reference = item.callback_ref end
    end
    selection = selection .. ',"callback_ref":"' .. reference .. '"'
end
if scenario ~= "missing" and scenario ~= "callback" then
    local text = 'New \\"name\\"'
    if scenario == "blank" then text = "   " end
    if scenario == "control" then text = "bad\\nname" end
    selection = selection .. ',"text":"' .. text .. '"'
end
local message = '{"version":1,"capability":"' .. token .. '","kind":"complete","selection":{' .. selection .. '}}'
function try_rename()
    if scenario == "moved" then window = "window:66666666-6666-6666-6666-666666666666" end
    actions.complete({ session_id = session, block_id = block,
        client_id = "client:77777777-7777-7777-7777-777777777777" }, { message = message })
end
function check_rename_result()
    local success = scenario == "tab" or scenario == "session" or scenario == "uncertain"
    assert(effects == (success and 1 or 0), "rename count")
    assert(callbacks == (scenario == "callback" and 1 or 0), "callback count")
end
