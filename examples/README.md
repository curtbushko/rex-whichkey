# Launch and JSON menu examples

These are source examples, not installed user configuration or runtime proof.
Lua owns one native float/direct launch; Go owns JSON/defaults/providers/menu,
terminal palette and CLI effects. No executable Lua menu data or callbacks run.

## Normal source run

From the originating Rex terminal in this checkout:

```sh
make run
# Current PATH make execs Task; optional action data follows --:
make run -- width=0.8 height=0.3 config=/absolute/path/menu.json ascii=true no_color=true
# Equivalent Task invocation:
task run -- config=/absolute/path/menu.json
# GNU make reads Makefile instead:
/usr/bin/make run ARGS='config=/absolute/path/menu.json'
```

Every run rebuilds this checkout into its known `bin/GOOS-GOARCH/rex-whichkey`,
then calls the current shared script through `rex do FILE --action
rex_whichkey_launch` with an absolute `development_directory`. The new terminal
gets one scoped PATH prepend; the command stays fixed `rex-whichkey direct`.
No old installed module/binary, init.lua registration or reload is needed.

Originating nonempty `REX_SESSION/REX_BLOCK` are mandatory and passed as explicit
`-s/-b`. Missing origin fails, not focused-block selection outside Rex. A present
original `REX_CLIENT` is passed as `-C`; Go rejects invalid/non-app intent, never
uses popup `REX_CLIENT` as the original or substitutes a focused/sole app.
Use your own real origin IDs, not another terminal's environment.
See [exact source command](../README.md#run-this-checkout).

`make run`/`rex do` returns creation IDs/launch status, not eventual menu exit.
A returned ID does not prove display, action success or host cleanup.

## Optional installed function keybinding

Install the shared module where Rex can require it, then only if you choose to
edit user-owned init.lua use [rex-whichkey-keybinding.lua](rex-whichkey-keybinding.lua):

```lua
local whichkey = require("rex_whichkey")
rex.bind("ctrl+space", function(ctx)
    return whichkey.launch(ctx, {
        width = .8, height = .3,
        -- config = "/absolute/path/menu.json",
        -- ascii = true, no_color = true,
    })
end)
```

Choose your own unused key; this example may replace an existing shortcut.
Home Manager links only the module and package, not the binding/init.lua/menu
JSON, and does not activate/reload anything. Installed launch omits
`development_directory`; both `rex-whichkey` and `rex` must already be on the
execution/server process PATH. A terminal profile change does not update it.

The same file can be run as an installed script with explicit original `-s/-b`
and optional original `-C` and `--action rex_whichkey_launch`, without a global
registration or require lookup. Only canonical `require("rex_whichkey")` returns
its function table; ordinary script execution serializes no functions.
The locally defined action is launch-only, not a command-palette/config takeover
or menu callback registry. Defaults are centered `.8/.3`; fractions must be finite
and in `(0,1]`. Native chrome uses Rex defaults for now; customization/REQ-026
is explicitly user-deferred, not implemented by terminal colors/text borders.

## Data configuration

[my_whichkey.json](my_whichkey.json) is a **whole-spec replacement** with optional
resource providers, not the unchanged built-ins. Copy/adapt it yourself, or pass
launch `config=/absolute/path/menu.json` / Lua `config="/absolute/path/menu.json"`.
Bare `rex-whichkey --config ...` is no longer a launcher. `direct` belongs inside
the sole Lua-created terminal with distinct original IDs; do not paste it into
an ordinary pane. Standalone `rex-whichkey menu --config PATH` is UI/result-only.

Go selection: present explicit `--config` (from launch data) > nonempty rendering
host `REX_WHICHKEY_CONFIG` > `$XDG_CONFIG_HOME/rex-whichkey/config.json`, or
`$HOME/.config/rex-whichkey/config.json` if XDG_CONFIG_HOME is absent. Only a
missing implicit file means built-ins. XDG_CONFIG_HOME must be absolute/nonempty.
Prefer absolute paths for server launch. Caller-only exports are not guaranteed
to reach the server-created terminal; explicit launch data avoids that ambiguity.

```json
{
  "version": 1,
  "title": "My menu",
  "spec": [
    {"keys": "w", "group": "Sessions"},
    {"keys": "ws", "desc": "Switch session", "action": "session.switch"},
    {"keys": "wr", "desc": "Rename session", "action": "session.rename"},
    {"keys": "x", "desc": "Close pane", "action": "pane.close", "confirm": true}
  ]
}
```

Omit `spec` to retain built-ins; `{"version":1}` is sufficient. Present `spec`/
`icons` replace whole values; empty strings and false are literal.
`hidden:true` hides but does not disable; `cond:false` disables. No file is
written automatically. [Data contract](../docs/mapping-format.md) defines strict
names, Unicode/key rules, bounds, destructive confirmation and provider subtrees.

Providers `sessions/windows/panes/known-hosts` require groups, no args/actions.
Windows/panes refer to the original session, excluding the menu's own popup.
Choices retain full IDs; hosts are informational, not SSH/connect. Only built-in
`ws` is dynamic by default. All use bounded in-process Go metadata jobs, no IPC.

## Migration from executable Lua and old Go launch

Preserve old files before changing wiring. Replace old
`M.new/setup/register_launch/register_completion`, runtime/helper-path registration
and bare-Go launcher bindings with the shared require/script entry. Manually
translate literal data into flat full-key JSON. `my_whichkey.lua` is retired,
not auto-converted. Review each action: arbitrary Lua actions, callbacks,
Zoom/directional-focus and executable hooks are unsupported; do not silently drop
them or call a partial translation equivalent. Functions/computed text/conditions,
nested inheritance and hooks remain deferred, with no compatibility stub/schema.
Historical originals remain in Git and test-only `tests/fixtures/phase15-lua`,
not runtime assets. Source changes are not activation/deployment permission.

## Lifetime and evidence

Workers/tty/visible float/exact original focus are retired/restored and verified
before effects. Accepted user1b allows the detached execution-host Block through
dispatch until Go exits/Rex deletes it; a later failure cannot suppress an earlier
effect. Go cannot verify its own deletion. Observe exact `session.list_blocks`
after exit, not just visible-layer absence. No uncertain retry or broad cleanup.

Task01/02 direct-source observations are recorded in native summaries. Actual
Task03 make-route checks on reviewed combined source and rebuilt main remain
coordinator-owned/pending here; Task04 refreshes final evidence. These examples
prove no physical-key, installed-profile, remote/platform or activation result.

- [Launch contract](../docs/launch-contract.md)
- [Package and CLI discovery](../docs/runtime-bootstrap.md)
- [Evidence and limits](../docs/acceptance.md)
