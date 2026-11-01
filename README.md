# Rex WhichKey

A short-lived, keyboard-driven menu for [Rex](https://www.superlogical.com/rex).
Go owns navigation and terminal rendering; a small server-side Lua integration
loads your mappings and executes actions/callbacks in the originating Rex runtime.
There is no daemon.

**Delivery status:** implementation has recorded independent source reviews,
Git integration, Go compilation and Nix package builds—not runtime acceptance.
Project tests and graphical/remote trials were not delivery gates; installed
startup, disconnected-origin behavior, Darwin runtime compatibility and Home
Manager activation remain unverified. See the
[delivery inventory and limitations](docs/acceptance.md) for REQ-001–REQ-017,
revision-bound evidence and historical caveats. No minimum tested Rex build is
claimed.

## Install on the Rex server host

The helper, `rex` CLI and Lua modules must be available to the **user/process
running the Rex server**, even when the graphical client is on another machine.
The implementation targets Linux and macOS; declared Nix outputs are not platform
certification. It expects compatible Rex configuration actions, detached process
execution, floating terminals and explicitly targeted CLI activation.

The flake exports `packages.<system>.rex-whichkey` (also `default`) for
`x86_64-linux`, `aarch64-linux`, `x86_64-darwin` and `aarch64-darwin`. The package
layout is:

- `bin/rex-whichkey`
- `share/rex-whichkey/lua/rex_whichkey.lua` and `rex_whichkey/`
- `share/doc/rex-whichkey/mapping-format.md`

Rex itself is not bundled. Installing the executable alone does not register the
menu or put its Lua modules in Rex's search path.

### Home Manager (opt in)

Add `github:curtbushko/rex-whichkey` as an input named `rex-whichkey` in your own
flake. In the server user's Home Manager configuration, with `inputs` supplied
by your configuration:

```nix
{ inputs, ... }: {
  imports = [ inputs.rex-whichkey.homeManagerModules.default ];
  programs.rex-whichkey.enable = true;
}
```

The module adds the package and links only `rex_whichkey.lua` and the
`rex_whichkey` module directory beside the server's `init.lua`, by default in
`${config.xdg.configHome}/rex`. For a custom `REX_CONFIG`, set
`programs.rex-whichkey.configDirectory` to its parent directory inside that user's
home. `programs.rex-whichkey.package` can override the package.

The module is disabled by default. It does **not** manage `init.lua`, user mapping
files, keybindings, server environment, services or reloads/restarts. Links use
`force = false`; do not overwrite unmanaged conflicts. If your configuration
already links the whole Rex directory, reconcile the sibling-module paths in your
own declarative layout. See the [Home Manager source-review cases](tests/packaging/home-manager/README.md)
for enabled/disabled behavior, custom locations and removal implications.

Without Home Manager, make the package's Lua entry point and module directory
available to the server runtime's documented module lookup, alongside your
user-owned configuration. Both `rex-whichkey` and `rex` must resolve in the server
process's PATH. A newly installed profile does not change an already-running
server's environment automatically.

## Wire your own Rex configuration

This is an **opt-in example**, not an automatically installed or reloaded file.
In your own server-side `init.lua`:

```lua
local popup = require("rex_whichkey")
popup.register_completion("whichkey_complete")

local menu = popup.new(require("rex_whichkey.defaults"))
menu.setup() -- built-ins; alternatively menu.setup(inline_options, "my_whichkey")
menu.register_launch("whichkey")
```

Register completion once, at configuration load time, in the **same server Lua
VM** as launch. Completion and launch action names must be distinct and unique
in your configuration. `menu.register_launch()` also defaults to `whichkey`;
`menu.register_launch("my_whichkey")` pairs with `rex-whichkey --action my_whichkey`.
Launch names contain only lowercase ASCII letters, digits and underscores. The
registered action also works from your normal Rex keys/palette without CLI origin
data; direct `menu.launch(ctx)` remains supported. No keybinding is installed.

CLI activation requires **Remote Control** in the selected graphical client's
settings, per Rex's custom-actions contract. The package changes neither that
setting nor transport authentication. Put this wiring in your own configuration
through ordinary Rex administration; rex-whichkey never edits or reloads it.

## Launch from a Rex terminal

After the one-time setup, run this in the originating Rex terminal:

```sh
rex-whichkey
```

This activates the registered action in the original Rex Lua VM, which starts a
real **Rex floating popup terminal** using your configured/default menu. It is not
a greeting, demo or standalone JSON menu. No private subcommand, script VM or
user-written resolved JSON is required. `rex-whichkey --help` (or `-h`) describes
the public options without contacting Rex:

| Option | Meaning |
| --- | --- |
| `--server ENDPOINT` | Concrete execution-host endpoint override. |
| `--client FULL_ID` | Exact full `client:UUID` attached to this terminal's session. |
| `--action NAME` | Installed configuration launch action; default `whichkey`. |

The command captures this terminal's full `REX_SESSION` and `REX_BLOCK`. Client
selection is `--client`, then a present `REX_CLIENT`, otherwise **exactly one**
eligible client attached to that captured session, with a live `connected_at`
generation. Rex terminals are not assumed to export `REX_CLIENT`. Zero or several
eligible clients require an exact attached client ID; labels/suffixes are not
accepted. Missing/stale origins fail rather than selecting the focused block,
current client or arbitrary first match. A successful command means activation
was accepted, not proof that a popup displayed or a callback completed. Failed or
uncertain activation is never retried automatically.

### Linux/macOS endpoints and overrides

Both public launch and Lua preparation use one Go-owned endpoint policy:

1. Present explicit override: public `--server`, or Lua's optional constructor
   argument `popup.new(defaults, endpoint)` for server-side preparation.
2. Nonempty `REX_SERVER` in that process's execution-host environment.
3. On Linux/macOS, the installed Rex CLI's native discovery:
   `rex --autostart=false --timeout 2s whoami --json`.

The helper captures and validates the returned concrete endpoint and targets it
explicitly afterward. Rex owns its platform socket paths and runtime/home/XDG
rules; rex-whichkey does not guess Linux filenames or macOS home directories,
scan sockets, create runtime directories or start a server. Discovery needs the
normal default server already running and reachable on the **execution/server
host**, not the viewing Mac. Missing/invalid runtime or home context, unavailable
CLI/server or unsupported identity JSON fails with guidance to fix that host's
environment or supply an endpoint. Invalid explicit/environment endpoints do not
fall through to native discovery. Explicit empty overrides are errors; an empty
`REX_SERVER` permits native discovery.

For a custom/shared server, configure the actual hosting endpoint consistently
in the terminal caller and server-side preparation. For example, **only if your
Rex server uses the documented shared socket** `/run/rex/server.sock`, use
`rex-whichkey --server /run/rex/server.sock` and construct the menu with
`popup.new(defaults, "/run/rex/server.sock")`, or supply matching `REX_SERVER`
environments. A public flag is not forwarded as a Lua constructor override;
Go's public endpoint and Lua preparation's canonical result must agree. Different
spellings that normalize to the same endpoint work; different server realms do
not. Setting an environment variable in a new terminal does not change an
already-running server's environment.

Accepted endpoints are absolute Unix socket paths (normalized to `unix:///...`),
strict absolute Unix URIs or `host:port` with a nonzero numeric port (bracket IPv6).
No whitespace, controls or mutable host labels are accepted. Routing is not
runtime authority: opaque `ctx.server` is never converted into an address, and
original server/session/block/client/connection checks remain mandatory. This
refines the earlier explicit-only setup, not its fail-closed safety rationale.
See the [launch contract](docs/launch-contract.md) for exact validation, discovery
bounds, source traces and runtime-unverified limits.

## Your mappings

For a separate configuration, put a differently named module such as
`my_whichkey.lua` beside the **server's** `init.lua`, returning an options table:

```lua
return {
    position = "bottom",
    spec = {
        { "w", group = "Sessions" },
        { "wn", "session.new", desc = "New session" },
        { "x", "pane.close", desc = "Close pane", confirm = true },
        { "h", group = "Rex-known hosts (listing only)", provider = "known-hosts" },
    },
}
```

Then use `menu.setup({}, "my_whichkey")` in the wiring above. For inline-only
configuration, pass the options table directly to `menu.setup(options)`.
Precedence is **defaults < imported module < inline options**. Present `spec` and
`icons` replace whole values, not merge them: `spec = {}` intentionally clears
all bindings, including built-ins and host listing. A failed setup retains the
previous installed configuration; successful replacement invalidates old pending
invocations. Ordinary Lua `require` caching applies; this is not hot reload.

Full key sequences, nested inherited attributes, hidden-but-active mappings,
conditions, dynamic labels and Lua callbacks are supported by the source.
`<Space>` and `<lt>` are the only bracket key forms. Callbacks stay in Rex Lua as
`function(ctx, args)`; Go receives data and opaque references, never executable
Lua. Imported modules are trusted code, not a sandbox. See the
[Go-owned mapping standard](docs/mapping-format.md) for exact fields, validation,
provider rules, callback lifetimes and completion safety; the
[configuration examples](examples/README.md) explain loader semantics.

## Use the menu

Type the displayed key to open a group or choose an action. Full sequences can
be typed without waiting for each redraw. Hidden bindings and keys on other pages
remain active.

| Key | Behavior |
| --- | --- |
| Esc / Backspace | Up one submenu; dismiss at the root. |
| `q` | Dismiss at any depth only when `q` is not mapped in that submenu. |
| Tab | Next page, wrapping to the first. |
| `y` in confirmation | Confirm the pending action. |
| Any other key in confirmation | Cancel that prompt without dispatch or navigation fallthrough. |

Built-in bindings from `rex_whichkey.defaults`:

| Keys | Actions |
| --- | --- |
| `wn`, `ws`, `wh`, `wl`, `wr`, `wx` | Session new, switch, previous, next, rename, close. Switch delegates Rex's picker; no enumerated session submenu is added. |
| `tn`, `th`, `tl`, `tr`, `tx` | Tab new, previous, next, rename, close. |
| `t1`–`t9` | Go to a tab by its one-based index. |
| `v`, `s` | Split right / down. |
| `fh`, `fj`, `fk`, `fl` | Focus left / down / up / right. |
| `z`, `x`, `e`, `c` | Zoom, close pane, toggle sidebar, toggle command palette. |
| `h` | Rex-known host inventory and warnings; details only, no connection action. |

Session/tab/pane close bindings require confirmation. Named actions and callbacks
share the confirmation path. The popup has a 30-second total invocation lifetime
starting before evaluation; preparation reduces available menu time, and the
controller reserves five seconds for teardown/completion. There is no automatic
retry of uncertain action execution.

Positions are `top`, `bottom`, `left`, `right` and `center`. Columns and paging
adapt to terminal size; very small viewports omit chrome before content. Unicode
cell-width calculation is not proof of actual client glyph geometry. A nonempty
`NO_COLOR` in the execution environment disables colors. Low-level `menu` and
`popup` modes accept `--ascii` and `--no-color`; these are not Lua setup options,
and the delivered launcher does not forward those flags. ASCII decorations do
not transliterate user labels. `menu` consumes resolved JSON and returns a
selection/null; it is not the normal Lua launch or an action executor. `prepare`,
`controller` and `popup` are internal plumbing, not manual invocation examples.

## Troubleshooting and safety

- **Module not found:** check the server's module installation/config directory
  and that the user module has a name other than reserved `rex_whichkey`.
- **Missing helper/CLI/setup:** check both the invoking terminal and server
  process's PATH, matching installed helper/CLI/modules, completion registration,
  successful `menu.setup()` and `menu.register_launch("whichkey")` (or matching
  `--action`). `REX_BIN_DIR`, if set, must name an absolute execution-host directory
  containing the matching `rex`. Raw child diagnostics are suppressed to avoid
  leaking capability data.
- **Missing/stale origin or ambiguous client:** launch inside the original Rex
  terminal with its real full `REX_SESSION`/`REX_BLOCK`. If several clients attach
  there, use `--client` with the exact full attached ID; correct a stale or empty
  `REX_CLIENT` rather than using focus/current-client selection. A reconnect's
  changed connection generation requires a new invocation.
- **Discovery/endpoint failure:** check the already-running execution-host server,
  Rex runtime/home settings and CLI compatibility. Fix malformed explicit or
  `REX_SERVER` values; a custom/shared endpoint may need both a public override and
  a matching Lua constructor/server environment. A viewing client's socket path
  or a different server realm is not a substitute. No server is autostarted and
  no alternate endpoint is tried.
- **Activation failed or uncertain:** verify the installed action name and client
  Remote Control permission. A zero exit status only acknowledges activation;
  later preparation/controller errors may still prevent display. Do not replay
  an uncertain launch/action to compensate.
- **Mappings disappear or conflict:** `spec` replaces defaults; duplicate actions,
  action/group prefix conflicts and unsupported key forms are errors. Import,
  evaluation or validation failure is not a successful empty configuration.
- **Actions rejected:** Remote Control policy, missing/disconnected origins,
  connection-generation changes, expiry and configuration replacement can reject
  a selection. Do not work around this by retargeting the popup/current client,
  replaying completion or retrying an uncertain destructive operation.
- **Host list empty or warning:** inventory belongs to the execution host and its
  configured Rex home, not the graphical client or `--server` target. Listing can
  include Rex-managed home synchronization; it is not a reachability check. It
  implements no attach, remote connection, Wake-on-LAN or SSH workflow.
- **Popup remains after failure:** cleanup closes only the returned owned terminal;
  it refuses broad layer/window/session cleanup. Server loss, process termination
  or uncertain creation can prevent cleanup. Never close unrelated resources by
  label to compensate.

Do not share private handoff/result files, capabilities or captured preparation
output. Terminal restoration and cleanup are bounded best effort, not guarantees
after SIGKILL or server/terminal loss. See [remote compatibility](docs/remote-compatibility.md)
for server-local paths, connection identity, moved-block handling and race limits,
and [the release inventory](docs/acceptance.md) for what has and has not been verified.
