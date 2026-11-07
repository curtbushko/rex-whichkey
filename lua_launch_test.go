package whichkey_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	lua "github.com/yuin/gopher-lua"
)

func TestLuaFloatLabel(t *testing.T) {
	state := lua.NewState()
	defer state.Close()
	require.NoError(t, state.DoString(`
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
            assert(args.session_id == session and args.window_id == window)
            assert(args.layout.block == captured_block)
            return { layer_id = "layer:00000000-0000-0000-0000-000000000002",
                block_ids = { "block:00000000-0000-0000-0000-000000000002" } }
        end,
    },
}
`))
	require.NoError(t, state.DoFile("lua/rex_whichkey.lua"))
	require.NoError(t, state.DoString(`launch_action.run(launch_context, {})`))
	block, ok := state.GetGlobal("captured_block").(*lua.LTable)
	require.True(t, ok, "launch must create a terminal block")
	require.Equal(t, lua.LString("rex-whichkey"), block.RawGetString("label"))
	require.Equal(t, lua.LString("com.superlogical.terminal.shell"), block.RawGetString("flavor"))
}
