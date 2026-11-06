# Launch and JSON menu examples

These are source examples, not installed user configuration or runtime evidence.
Go owns configuration, providers, menu lifetime and all Rex operations.

## Optional function keybinding

Install the launch module on the execution host where Rex loads sibling modules,
then, only when you choose to edit your own init.lua, use
[rex-whichkey-keybinding.lua](rex-whichkey-keybinding.lua):

```lua
local whichkey = require("rex_whichkey")
rex.bind("ctrl+space", function(ctx)
    return whichkey.launch(ctx)
end)
```

Choose an unused key: this spelling is only an example and can replace a shortcut.
There is no named action/palette registration. Rex's process PATH must already
contain the installed `rex-whichkey`; a terminal's newly changed PATH is not proof
of the running server's environment. Home Manager links only the module and does
not install this binding, reload Rex or own init.lua.

The module uses documented argument-array `rex.exec.start`, `detach=true`, no
timeout and no lifetime wait. `env` overlays the execution environment; `false`
removes missing `ctx.session_id/block_id/client_id` values from
`REX_SESSION/REX_BLOCK/REX_CLIENT`. It never derives endpoint authority from
`ctx.server`. A start handle is not menu-display evidence; later failures belong
to the detached Go process. Immediate start errors are raised, not called success.

## Data configuration

[my_whichkey.json](my_whichkey.json) is a **whole-spec replacement**, including
explicit optional resource-list providers. It is not the unchanged built-in menu.
Copy/adapt it yourself to your execution-host config location or pass
`rex-whichkey --config /absolute/path/menu.json`. The minimal binding starts the
same public binary without flags; it uses Go's environment/default file selection.

Selection: present `--config` > nonempty `REX_WHICHKEY_CONFIG` >
`$XDG_CONFIG_HOME/rex-whichkey/config.json`, or
`$HOME/.config/rex-whichkey/config.json` if XDG_CONFIG_HOME is absent.
Only missing implicit data means built-ins. Relative explicit/environment paths
use the execution process cwd; prefer absolute paths for server launch.
XDG_CONFIG_HOME, if present, must be absolute and nonempty.

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

Omit `spec` to retain all built-ins; `{"version":1}` is sufficient.
Present `spec`/`icons` replace whole values, not per-key merges. Empty strings
and false booleans are literal data. `hidden:true` hides but does not disable;
`cond:false` disables. No source file is generated automatically.
The [schema/action catalogue](../docs/mapping-format.md) defines strict names,
Unicode/key rules, bounds, destructive confirmation and provider subtrees.

Providers `sessions/windows/panes/known-hosts` require groups, no args/actions.
Windows/panes refer to the origin session, not a viewer's current session.
Generated choices retain full IDs; hosts are listing-only, not SSH/connect.
Only built-in `ws` is dynamic by default.

## Migration from executable Lua options

Preserve old user files before changing wiring. Replace old
`M.new/setup/register_launch/register_completion` calls and runtime registrations
with the function binding. Manually translate literal data into flat full-key
JSON mappings. The old `my_whichkey.lua` example is retired, not auto-converted.
Review every old action: the JSON catalogue deliberately rejects arbitrary Lua
actions, Zoom/directional-focus actions and executable callbacks. Do not silently
discard unsupported user data or present a partial translation as equivalent.

Nested Lua inheritance, functions/computed text/conditions, callbacks and hooks
are deferred. There is no new hook schema or stub. Old originals remain in Git;
legacy Lua test modules are archived byte-for-byte under
`tests/fixtures/phase15-lua`. They are not installation or runtime assets.
Changing source declarations is not activation/deployment permission.

## Evidence and references

The coordinator observed Task2's direct public default menu and dismissal on
Darwin arm64, not this example binding or JSON replacement. Final combined runtime,
selected effects and physical keyboard launch are not established by these files.
No user binding, config, profile or Rex runtime was changed by writing them.

- [Launch contract](../docs/launch-contract.md)
- [Ordinary package and CLI discovery](../docs/runtime-bootstrap.md)
- [Release evidence and limits](../docs/acceptance.md)
- Rex [function binding](https://www.superlogical.com/rex/docs/reference/lua#rexbindkey-fn),
  [execution options](https://www.superlogical.com/rex/docs/reference/lua#rexexecargv-opts),
  [detached start](https://www.superlogical.com/rex/docs/reference/lua#rexexecstartargv-opts)
