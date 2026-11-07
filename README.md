# Rex WhichKey

A short-lived keyboard menu for [Rex](https://www.superlogical.com/rex).
Shared Lua creates one centered floating terminal; one Go process inside it reads
JSON menu data, renders the menu and performs operations through the Rex CLI.
There is no Go launcher/controller, renderer child, image copy, private IPC,
daemon or executable Lua menu configuration.

**Evidence:** Phase19 Task03 records actual `make run` on reviewed combined
source and rebuilt main at `6564d07ef690d591ea842fa9d986d0bcfa919dd4`: current
checkout Go executable, shared Lua float, root/providers, exact owned rename,
separate edited cancellation/dismissal, pre-effect visible retirement and observed
exit 0/post-exit host absence. Task02 separately observed real own-PTY resize.
This docs-only refresh runs no product; independent review and the coordinator's
final rebuilt-main run remain required before phase completion. Build/help/creation
acknowledgment is not menu-display or exit proof. Full lint remains unresolved
(53 diagnostics, exit 201; AI lint not reached). No installed-profile, physical-
keybinding, remote or all-platform success is claimed. See [acceptance](docs/acceptance.md).

## Run this checkout

From the **originating execution-host Rex terminal**, with Go, Task and a
compatible licensed `rex` already available:

```sh
make build
make run
```

`make run` always builds this checkout first, then invokes the current
`lua/rex_whichkey.lua` script with `--action rex_whichkey_launch`. It requires
nonempty originating `REX_SESSION` and `REX_BLOCK` and passes them explicitly as
`-s`/`-b`; outside Rex it fails rather than using whichever block is focused.
A present `REX_CLIENT` is forwarded as the original `-C` intent, never replaced
with the popup's client. It must be a full ID and an eligible graphical app;
a CLI caller ID is not permission to impersonate an app. Without explicit app
intent, Go can use the original pane's validated owner for app operations.
Do not copy this builder's or another terminal's IDs.

Task computes the known absolute `bin/GOOS-GOARCH` directory from this checkout
and supplies it as `development_directory` data. Lua prepends that single directory
to this terminal's PATH and starts the fixed program `rex-whichkey direct` with
separate original IDs. Lua does no source-path probing, `go env` discovery or
executable nomination. No installed old module/binary or init.lua reload is needed.

Two `make` implementations have different extra-argument syntax:

```sh
# Current development PATH: make is a tiny wrapper that execs Task.
make run -- width=0.8 height=0.3 config=/absolute/path/menu.json ascii=true no_color=true
# Task directly (same route):
task run -- width=0.8 height=0.3 config=/absolute/path/menu.json
# Conventional GNU make reads Makefile and delegates to Task:
/usr/bin/make run ARGS='width=0.8 height=0.3 config=/absolute/path/menu.json'
```

Use your actual GNU make path if it differs; `command -v make` identifies the PATH
entry. Extra arguments are Lua action `key=value` data, **not Go root flags**.
Shell-quote paths containing spaces normally (and retain quoting inside GNU
make's `ARGS`). Omit `development_directory` from user arguments: the source task
supplies it. Defaults are centered width `.8`, height `.3`; dimensions must be
finite fractions in `(0,1]`. `ascii`/`no_color` are optional booleans.

The equivalent explicit source entry, after `task build`, is:

```sh
: "${REX_SESSION:?originating Rex session required}"
: "${REX_BLOCK:?originating Rex block required}"
set --
if [ "${REX_CLIENT+x}" = x ]; then
  : "${REX_CLIENT:?original full client ID required when present}"
  set -- -C "$REX_CLIENT"
fi
rex "$@" do "$PWD/lua/rex_whichkey.lua" -s "$REX_SESSION" -b "$REX_BLOCK" \
  --action rex_whichkey_launch \
  development_directory="$PWD/bin/$(go env GOOS)-$(go env GOARCH)"
```

Run that from the checkout root; `make run` supplies its root independently of
shell cwd. `rex do` returns the created IDs, not the eventual Go exit status.
The menu has no idle deadline. Observe the exact popup, terminal exit and
`session.list_blocks` after completion to establish cleanup; see the
[launch contract](docs/launch-contract.md).

## Install on the Rex server host

Both executables belong on the **execution/server host**, not just the viewer.
The existing package installs ordinary assets:

```text
bin/rex-whichkey                         # real Go executable
bin/rex                                # optional licensed CLI symlink
share/rex-whichkey/lua/rex_whichkey.lua   # shared module/script
share/rex-whichkey/examples/             # JSON and keybinding examples
share/doc/rex-whichkey/                  # documentation
```

The flake declares x86_64/aarch64 Linux and macOS package outputs, not runtime
certification. Pinned official CLI artifacts cover x86_64/aarch64 Linux and
Apple Silicon macOS, **not Intel macOS**. A legitimate `rex-cli` override or
execution-host PATH CLI can supply `rex`; licensing/unfree policy applies.
Pins/vendor hash, Linux loader/PAM/terminfo/SSH assets and Darwin signature
preservation are unchanged. Nix installs assets, not runtime discovery policy.
A Go source build alone installs neither the CLI nor Lua/assets.

Optional Home Manager integration:

```nix
{ inputs, ... }: {
  imports = [ inputs.rex-whichkey.homeManagerModules.default ];
  programs.rex-whichkey.enable = true;
}
```

It contributes the package and **only** `rex_whichkey.lua` beside user-owned
`init.lua` (default `${config.xdg.configHome}/rex`). `package` and
`configDirectory` are overridable. It does not own the whole Rex directory,
JSON, init.lua or bindings, forcibly replace conflicts, activate a profile or
reload/restart Rex. See the [source guide](tests/packaging/home-manager/README.md).

## Installed public launch

With the package bin directory already on the Rex server's PATH, the shared
script is directly usable without user action registration or reload:

```sh
: "${REX_SESSION:?originating Rex session required}"
: "${REX_BLOCK:?originating Rex block required}"
# PACKAGE is the absolute package output.
set --
if [ "${REX_CLIENT+x}" = x ]; then
  : "${REX_CLIENT:?original full client ID required when present}"
  set -- -C "$REX_CLIENT"
fi
rex "$@" do "$PACKAGE/share/rex-whichkey/lua/rex_whichkey.lua" \
  -s "$REX_SESSION" -b "$REX_BLOCK" --action rex_whichkey_launch
```

Installed launches **omit `development_directory`** and use ordinary PATH.
Alternatively choose a key yourself in your normal user-owned configuration:

```lua
local whichkey = require("rex_whichkey")
rex.bind("ctrl+space", function(ctx)
    return whichkey.launch(ctx, { width = .8, height = .3 })
end)
```

The same module accepts optional `config`, `ascii` and `no_color` data.
Only canonical `require("rex_whichkey")` receives the function table; ordinary
script execution returns no function for Rex to serialize. The script defines
`rex_whichkey_launch` locally for `--action`; it is not a prerequisite global
registration, menu callback or command-palette takeover. This example key has
not been installed or pressed. Installing a package or changing a terminal
profile does not update an already-running server's environment.

Bare `rex-whichkey` and old root launcher flags are rejected safely. `direct`
is the explicit Lua-created-terminal entry, **not** a command to paste into an
ordinary pane: it requires original session/block/window/layer IDs distinct from
its own popup `REX_*` identity. `rex-whichkey menu [--config PATH] [--ascii]
[--no-color]` is in-terminal UI/result-only and performs no selected operation.

## Menu data and migration

Built-ins need no file. Go overlays one strict version-1 JSON file, selected by:

1. Present launch `config` (passed to Go `direct --config PATH`).
2. Nonempty rendering-host `REX_WHICHKEY_CONFIG`.
3. `$XDG_CONFIG_HOME/rex-whichkey/config.json`, or
   `$HOME/.config/rex-whichkey/config.json` when XDG_CONFIG_HOME is absent.

Only a missing implicit default means built-ins. Invalid/empty explicit paths,
missing selected files or invalid data fail; nothing writes configuration.
`spec` and `icons` replace whole values; `spec: []` clears mappings.
See the [data contract](docs/mapping-format.md) and [JSON example](examples/my_whichkey.json).
The server-created terminal's environment, not a caller's new exports, governs
implicit data and `NO_COLOR`; explicit launch data avoids that ambiguity.

Preserve old Lua files before migration. Manually flatten literal keys/groups/
labels/boolean conditions to JSON and check every action against the closed
catalogue. Replace old `M.new/setup/register_launch/register_completion`, helper
paths and bare-binary launch bindings with the shared module/script. No automatic
conversion is provided. Functions, callbacks, custom Lua actions, computed values,
nested inheritance and hooks remain deferred; never silently discard them.
Historical Lua stays in Git and test-only `tests/fixtures/phase15-lua`, not runtime
assets. See [examples and migration](examples/README.md).

## Use and safety

| Keys | Built-ins |
| --- | --- |
| `wn/ws/wh/wl/wr/wx` | Session new / Go switch list / previous / next / inline rename / close |
| `tn/th/tl/tr/tx` | Tab new / previous / next / inline rename / close |
| `v/s` | Split right / down |
| `x/e/c` | Close pane / sidebar / command palette |

Close requires `y`; another prompt key cancels. Esc/Backspace go up, then dismiss
at root; unmapped `q` dismisses. Tab pages/wraps; hidden/off-page mappings remain
active. Providers are sessions/windows/panes/known-hosts (hosts listing-only).
Inline rename uses this same menu. Lua keeps one fixed fractional rectangle;
Go fits its own PTY cells and handles SIGWINCH/pagination without resizing it.
Native chrome uses Rex defaults **for now**: user1a explicitly deferred REQ-026
native titlebar/border/corners/shadow customization; it is not implemented or
credited complete. Rex's native header, border, corners and shadow may remain
visible. Go's rendered header/footer and semantic terminal colors are separate;
no theme setting or painted border supplies native chrome control.

Go selects `rex` once by ordinary execution-host PATH and independently selects
its endpoint: explicit direct `--server` > nonempty `REX_SERVER` > bounded
no-autostart native Linux/macOS `whoami`. There is no `REX_BIN_DIR` lookup,
Nix runtime resolver, binary hash/status/version/schema preflight or retry route.
Original session/block/window/layer and explicit client intent are retained;
popup `REX_CLIENT` is never treated as original GUI intent. App actions require
fresh exact app connection/owner/generation and Rex permissions, not current focus.

Before an effect, Go joins workers, restores the tty, detaches/verifies the exact
sole-block float and restores/verifies exact original focus/placement. Uncertain
pre-effect cleanup suppresses effects. User1b permits the detached execution-host
Block to remain through dispatch until Go exits and Rex deletes it. Go cannot
verify its own post-exit deletion; a later cleanup failure cannot undo or suppress
an already-issued effect. No uncertain mutation is retried. SIGKILL/server loss
can leave resources; there is no broad cleanup guarantee.

See [launch/lifetime](docs/launch-contract.md), [discovery](docs/runtime-bootstrap.md),
[CLI/GUI](docs/cli-gui-delegation.md), [appearance](docs/appearance.md) and
[remote limits](docs/remote-compatibility.md). This source delivery claims no
prohibited tests, Nix build/activation, deployment, physical key or phase completion.
