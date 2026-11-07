# Rex WhichKey

A short-lived keyboard menu for [Rex](https://www.superlogical.com/rex).
Go reads JSON menu data, builds dynamic resource lists, renders the popup and
performs operations through the Rex CLI. Lua only starts the installed binary.
There is no daemon, executable Lua configuration or hook/callback facility.

**Evidence:** Phase16 CLI-only delivery is complete, including actual public
menu/navigation/owned rename/dismissal/cleanup on Darwin arm64. Phase18 tasks1/2
also have reviewed, integrated public-run evidence, including rename edits while
loading and preservation after ready. Phase18 async source `d60c976` passed
independent source review; coordinator public runs on combined candidate
`397ce9a` captured coarse first-visible upper bounds of 2.69s baseline versus
0.54s candidate, provider readiness/navigation, an exact owned rename and
cancel/dismissal/cleanup. These are CLI/capture/poll observations, not exact
render/load timestamps, a benchmark or an SLA. Later documentation-only revisions
preserve the reviewed Go bytes; native task/phase summaries bind final delivery.
No installed-profile, physical-keybinding, all-provider or cross-platform success
is claimed. See [acceptance](docs/acceptance.md) for exact revisions and limits.

## Install on the Rex server host

The binary and compatible licensed Rex CLI must be available on the
**execution/server host**, not merely the machine displaying Rex.
The package installs ordinary assets:

```text
bin/rex-whichkey                         # real executable
bin/rex                                # optional licensed CLI symlink
share/rex-whichkey/lua/rex_whichkey.lua   # launch only
share/rex-whichkey/examples/             # JSON data and keybinding examples
share/doc/rex-whichkey/                  # documentation
```

The flake exports `packages.<system>.rex-whichkey` and `default` for
x86_64/aarch64 Linux and macOS. These are declared outputs, not platform
certification. Pinned official CLI artifacts exist for x86_64/aarch64 Linux and
aarch64 macOS, **not Intel macOS**. A legitimate `rex-cli` package override
can supply `bin/rex`; otherwise provide a legitimate `rex` on the
execution-host PATH. Licensing/unfree policy still applies.
Linux loader/PAM/terminfo/SSH assets and the released Darwin signature are
preserved. Nix installs inputs; Go alone owns runtime discovery and policy.
No wrapper, Lua-tree binary, anchor or generated helper layout is required.
A source-only Go build supplies neither the licensed CLI nor installed Lua/assets;
provide `rex` separately and expose the named `rex-whichkey` binary on PATH.
The optional package CLI is the same pinned dependency, now exposed as `bin/rex`;
no dependency, pin or profile change is implied.

Optional Home Manager integration:

```nix
{ inputs, ... }: {
  imports = [ inputs.rex-whichkey.homeManagerModules.default ];
  programs.rex-whichkey.enable = true;
}
```

It contributes the package and **only** `rex_whichkey.lua` beside the user's
Rex `init.lua` (default `${config.xdg.configHome}/rex`).
`programs.rex-whichkey.package` and `configDirectory` are overridable.
It does not manage the whole Rex directory, menu JSON, init.lua or a binding.
Conflicting unmanaged files are not forcibly replaced. Importing is opt-in;
nothing here activates a profile, reloads Rex or changes a running environment.
See the [Home Manager source guide](tests/packaging/home-manager/README.md).

## Launch

Inside the originating execution-host Rex terminal:

```sh
rex-whichkey
# Optional public flags:
rex-whichkey --config /absolute/path/menu.json --ascii --no-color
```

Public flags are `--config PATH`, `--server ENDPOINT`, `--client FULL_ID`,
`--ascii`, `--no-color`, and `--help`/`-h`.
Full `REX_SESSION` and `REX_BLOCK` are required; no current/focused-pane
fallback exists. Help has no effects.

For a user-chosen key, the [example binding](examples/rex-whichkey-keybinding.lua)
uses Rex's documented function binding:

```lua
local whichkey = require("rex_whichkey")
rex.bind("ctrl+space", function(ctx)
    return whichkey.launch(ctx)
end)
```

The key is **only an example**. It has not been installed or pressed as part of
this delivery. Choose and wire it yourself through your normal configuration
lifecycle; a binding may replace an existing shortcut. The server process must
already find both `rex-whichkey` and `rex` on its execution-host PATH.
Installing a package or updating a terminal profile does not change a running
server's inherited environment. No command palette registration is added.

The launch function calls `rex.exec.start({ "rex-whichkey" }, { detach = true,
env = ... })`. It passes exact context session/block/client IDs and explicitly
removes missing ones from the inherited environment. It never queries Rex,
stringifies `ctx.server`, waits for the menu, or supplies a command timeout.
Immediate start errors raise an error; a returned process handle proves only
process start, not a displayed menu or eventual success. The binding returns true.
Both entrypoints use the same ordinary Go path.

## Menu data and migration

Built-ins need no file. Go overlays exactly one version-1 JSON file, selected by:

1. Present `--config PATH`.
2. Nonempty `REX_WHICHKEY_CONFIG`.
3. `$XDG_CONFIG_HOME/rex-whichkey/config.json`, or
   `$HOME/.config/rex-whichkey/config.json` when XDG_CONFIG_HOME is absent.

Only a missing implicit default means built-ins. Empty explicit paths, invalid
XDG paths, missing selected files or invalid data fail, not silently fall back.
No configuration is written. `spec` and `icons` replace whole values;
`spec: []` deliberately clears mappings. See the complete
[data/action/provider contract](docs/mapping-format.md) and
[JSON example](examples/my_whichkey.json).

For migration, preserve your old Lua files first. Manually flatten literal full
key sequences/groups/labels/boolean conditions into JSON and check every action
against the closed catalogue. Replace old `M.new/setup/register_launch`,
runtime/completion registrations and helper-path wiring with the function binding.
`examples/my_whichkey.lua` is retired; JSON is data, not an executable import.
There is **no automatic conversion** or promise to retain unsupported Lua
semantics. Functions, callbacks, custom Lua actions, computed labels/conditions,
nested inheritance and hooks are deferred; do not silently drop them from a
user's configuration. Retain the old file until you have accounted for its data.
Historical source is preserved in Git and test-only
`tests/fixtures/phase15-lua`, never shipped or loaded by the launcher.

## Use the menu

| Keys | Current built-ins |
| --- | --- |
| `wn/ws/wh/wl/wr/wx` | Session new / Go switch list / previous / next / inline rename / close |
| `tn/th/tl/tr/tx` | Tab new / previous / next / inline rename / close |
| `v/s` | Split right / down |
| `x/e/c` | Close pane / sidebar / command palette |

Close actions require `y` confirmation; any other prompt key cancels.
Esc/Backspace go up, then dismiss at root. Unmapped `q` dismisses at any depth;
Tab pages and wraps. Hidden/off-page mappings stay active.
There is no Zoom, directional-focus menu, numbered-tab list, app chooser or
default hosts row. Optional providers are sessions/windows/panes/known-hosts;
hosts are informational only. Inline rename uses the existing popup, not a
native dialog. See [actions](docs/actions.md) and [rename](docs/tab-rename-input.md).

Go retains one initial popup rectangle through navigation; larger menus paginate.
There is no idle menu expiry. Per-operation deadlines, cancellation and bounded
cleanup remain. Terminal focus-in repaints; switching tabs alone is not dismissal.
Native Rex controls are **not hidden** (REQ-026 remains unresolved).
See [appearance](docs/appearance.md).

## Targeting and safety

Go selects the first `rex` found by ordinary execution-host PATH lookup, then
reuses that path for the invocation and private child. `REX_BIN_DIR`, private
libexec layouts and executable overrides are not search authorities. There is
no binary hashing/receipt, startup version/API/schema negotiation or per-command
server.status/config-generation sandwich. Actual CLI failures are reported
without retrying or selecting another binary/server; selected-action metadata
and strict functional result/identity checks remain.
Endpoint precedence is present `--server` > nonempty `REX_SERVER` >
bounded no-autostart native Linux/macOS `whoami` discovery. No viewing-host socket
or Lua server object is authority. See [discovery](docs/runtime-bootstrap.md).

A present `--client`/`REX_CLIENT` must be a full valid app ID and stays strict.
Without it, display/resource-only actions do not require a GUI; app actions use
the captured pane's validated app owner and connection generation, never whichever
app is focused. CLI identity is not GUI identity or additional permission.
Rex authentication and Remote Control still apply.

The terminal, exact returned popup, origin focus, lifetime pipe, private files and
captured image handle are cleaned before selected effects. Failure or uncertain
cleanup suppresses effects; uncertain mutations are never retried. New Tab creates
one exact source-session window and activates its returned ID; a failed activation
preserves that tab for manual inspection, not rollback or duplicate creation.

For diagnosis, distinguish invalid data, missing/stale origin, CLI/endpoint
discovery, app ownership/Remote Control, and uncertain cleanup/effects. Do not
repair failures by selecting a focused resource, reloading the user's config or
closing resources by label. SIGKILL/server loss can leave owned artifacts; cleanup
is bounded best effort. Read the [launch contract](docs/launch-contract.md),
[CLI/GUI boundary](docs/cli-gui-delegation.md) and
[remote limitations](docs/remote-compatibility.md) before acting on an error.

Historical Phase15 bridge guidance is not a startup prerequisite. This work does
not mark Phase15 or unrelated blocked requirements complete, and claims no tests,
Nix build/activation, deployment, remote trial or physical-key observation.
