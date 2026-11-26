local uuid = "00000000-0000-0000-0000-000000000001"
local session, block = "session:" .. uuid, "block:" .. uuid
local window, layer = "window:" .. uuid, "layer:" .. uuid
launch_context = { session_id = session, block_id = block }
rex = {
    action = function(action) launch_action = action end,
    layout = { block = function(spec)
        captured_block = spec
        return { block = spec }
    end },
    session = {
        view = function(args)
            assert(args.session_id == session)
            return { session_id = session, windows = {
                { window_id = window, layers = {
                    { layer_id = layer, blocks = { { block_id = block } } }
                } }
            } }
        end,
        new_layer = function(args)
            captured_layer = args
            assert(args.session_id == session and args.window_id == window)
            return { layer_id = "layer:00000000-0000-0000-0000-000000000002",
                block_ids = { "block:00000000-0000-0000-0000-000000000002" } }
        end,
    },
}
