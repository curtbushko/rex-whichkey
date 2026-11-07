package whichkey_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	lua "github.com/yuin/gopher-lua"
)

const testCallbackScenario = "callback"

func TestConfiguredRenameDispatch(t *testing.T) {
	for _, scenario := range []string{"tab", "session", testCallbackScenario, "missing", "blank", "control", "custom_args", "injected_callback", "moved", "uncertain"} {
		t.Run(scenario, func(t *testing.T) {
			state := lua.NewState()
			defer state.Close()
			state.SetGlobal("scenario", lua.LString(scenario))
			require.NoError(t, state.DoFile("tests/lua/configured_rename.lua"))
			invoke := func() error {
				return state.CallByParam(lua.P{Fn: state.GetGlobal("try_rename"), Protect: true, NRet: 0})
			}
			if scenario == "tab" || scenario == "session" || scenario == testCallbackScenario {
				require.NoError(t, invoke())
			} else {
				require.Error(t, invoke())
			}
			require.NoError(t, state.CallByParam(lua.P{Fn: state.GetGlobal("check_rename_result"), Protect: true, NRet: 0}))
			require.Error(t, invoke(), "no replay or uncertain-action retry")
			require.NoError(t, state.CallByParam(lua.P{Fn: state.GetGlobal("check_rename_result"), Protect: true, NRet: 0}))
		})
	}
}

func TestCLIDelegation(t *testing.T) {
	for _, scenario := range []string{
		"app_action", "key_launch", "confirmed", "unconfirmed", "guard_changed", testCallbackScenario, "named", "cancel",
		"wrong_session", "wrong_block", "wrong_server", "invalid_caller", "wrong_capability", "malformed",
		"reconnect", "lost", "expired", "reload", "caller_mutation", "rpc_error", "rpc_wrong_client",
		"non_app", "no_target", "kind_changed", "launch_caller_mutation",
	} {
		t.Run(scenario, func(t *testing.T) {
			state := lua.NewState()
			defer state.Close()
			state.SetGlobal("scenario", lua.LString(scenario))
			require.NoError(t, state.DoFile("tests/lua/cli_delegation.lua"))
			// Separate protected host calls model distinct Rex activations, not
			// nested actions inside a single Lua activation.
			if bad := state.GetGlobal("bad_completion"); bad != lua.LNil {
				require.Error(t, state.CallByParam(lua.P{Fn: bad, Protect: true, NRet: 0}))
			}
			if finish := state.GetGlobal("finish_test"); finish != lua.LNil {
				require.NoError(t, state.CallByParam(lua.P{Fn: finish, Protect: true, NRet: 0}))
			}
		})
	}
}

func TestLaunchOriginPreflightDiagnostics(t *testing.T) {
	for _, field := range []string{"session_id", "block_id", "client_id", "connected_at"} {
		t.Run(field, func(t *testing.T) {
			state := lua.NewState()
			defer state.Close()
			state.SetGlobal("test_origin_mismatch", lua.LString(field))
			require.NoError(t, state.DoFile("tests/lua/action_context.lua"))
		})
	}
}

func TestRexActionContexts(t *testing.T) {
	for _, identity := range []string{"absent", "string", "table"} {
		t.Run(identity, func(t *testing.T) {
			state := lua.NewState()
			defer state.Close()
			state.SetGlobal("test_server_kind", lua.LString(identity))
			require.NoError(t, state.DoFile("tests/lua/action_context.lua"))
		})
	}
}
