# Rex WhichKey

A short-lived, keyboard-driven menu for [Rex](https://www.superlogical.com/rex).
Go owns navigation, terminal rendering and the basic native-action menu. Optional
server-side Lua integration loads custom mappings and executes callbacks in the
originating Rex runtime. There is no daemon.

**Verification:** recorded evidence distinguishes independent source review,
Git integration, Go/Nix builds and targeted runtime observations. Phase07 records
a normal native popup launch, Tabs navigation, Escape dismissal, exit0 and exact
owned-popup cleanup in one Linux Rex session. That does not verify the later
idle-lifetime/color changes or certify configured callbacks, remote/Darwin runtime
compatibility or Home Manager activation. The coordinator's Phase08 gate adds
35 seconds of idle time, reduced defaults and VT/HTML color evidence. See the
[delivery inventory](docs/acceptance.md) for REQ-001–REQ-020 and revision-bound
results. No minimum tested Rex build or general runtime certification is claimed.

## Install on the Rex server host

The helper and `rex` CLI must be available on the **execution/server host**,
even when the graphical client is on another machine. Optional customization
also needs Lua modules available to the user/process running the server. The
implementation targets Linux and macOS; declared Nix outputs are not platform
certification. Basic startup uses native floating terminals and explicitly
targeted CLI activation; configured Lua launch additionally uses configuration
actions and detached process execution.

The flake exports `packages.<system>.rex-whichkey` (also `default`) for
`x86_64-linux`, `aarch64-linux`, `x86_64-darwin` and `aarch64-darwin`. The package
layout is:

- `bin/rex-whichkey`
- `share/rex-whichkey/lua/rex_whichkey.lua` and `rex_whichkey/`
- `share/doc/rex-whichkey/mapping-format.md`

Rex itself is not bundled. The executable plus a compatible server-host Rex CLI
and attached graphical app suffice for the basic menu; no custom action
registration is required. Installing the executable does not put optional Lua
modules in Rex's search path or register a customized menu.

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

For Lua customization without Home Manager, make the package's Lua entry point
and module directory available to the server runtime's documented module lookup,
alongside your user-owned configuration. Both `rex-whichkey` and `rex` must resolve in the server
process's PATH. A newly installed profile does not change an already-running
server's environment automatically.

## Optional Lua customization

The basic menu requires no `init.lua` changes. To use custom Lua mappings and
callbacks, this is an **opt-in example**, not an automatically installed or reloaded file.
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

With the binaries installed, run this in the originating Rex terminal:

```sh
rex-whichkey
```

This opens a real **Rex floating popup terminal** with the Go-owned built-ins,
without custom `whichkey`/completion registration. If the server's read-only
action registry instead contains a Lua `whichkey` action, launch uses that
original-VM configured menu. An explicit `--action NAME` always requires that
registered action, even `--action whichkey`; missing registration fails clearly.
Failed or uncertain configured activation never falls back to native mode.

The native caller stays in the foreground until selection/dismissal, cancellation
or a real terminal/origin failure, with no idle timeout. It performs teardown
before any selected native action. The configured route keeps its detached
controller, prompt-return behavior and authenticated 30-second expiry. Neither
route is a greeting, demo or standalone JSON menu; no private subcommand, new
script VM or user-written resolved JSON is
required. `rex-whichkey --help` (or `-h`) describes public options without contacting Rex:

| Option | Meaning |
| --- | --- |
| `--server ENDPOINT` | Concrete execution-host endpoint override. |
| `--client FULL_ID` | Exact full `client:UUID` of a live graphical app attached to this terminal's session. |
| `--action NAME` | Require an installed Lua launch action; no native fallback. Omit to prefer registered `whichkey`, otherwise use basic defaults. |

The command captures this terminal's full `REX_SESSION` and `REX_BLOCK`. Client
selection is `--client`, then a present `REX_CLIENT`, otherwise **exactly one live
graphical app** attached to that captured session, with a valid `connected_at`
generation. Rex's native attachment `info.kind == "app"` classification excludes
CLI/automation connections and unknown kinds; labels/names/platforms are not kind
heuristics. `REX_CLIENT` is normally absent and is not required for the single-app
case, even with other CLI connections. No attached live app requires connecting
the app to that session; several attached apps require an exact full app client
ID. Explicit selectors must also identify an attached live app; labels/suffixes,
stale IDs and non-app clients are rejected. Missing/stale origins fail rather than
selecting the focused block, current client or arbitrary first match. Configured
success acknowledges activation, not popup display or callback completion. Native
success follows child exit, exact owned-popup/file cleanup and optional one-shot
action acceptance; RPC acceptance alone is never proof of graphical display.
Failed or uncertain activation is never retried automatically.

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

For a custom/shared server, use the actual hosting endpoint in the terminal
caller. On the configured Lua route, server-side preparation must agree too. For example, **only if your
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
    },
}
```

Then use `menu.setup({}, "my_whichkey")` in the wiring above. For inline-only
configuration, pass the options table directly to `menu.setup(options)`.
Precedence is **defaults < imported module < inline options**. Present `spec` and
`icons` replace whole values, not merge them: `spec = {}` intentionally clears
all bindings. A failed setup retains the previous installed configuration;
successful replacement invalidates old pending
invocations. Ordinary Lua `require` caching applies; this is not hot reload.

Full key sequences, nested inherited attributes, hidden-but-active mappings,
conditions, dynamic labels and Lua callbacks are supported by the source.
`<Space>` and `<lt>` are the only bracket key forms. Callbacks stay in Rex Lua as
`function(ctx, args)`; Go receives data and opaque references, never executable
Lua. Imported modules are trusted code, not a sandbox. See the
[Go-owned mapping standard](docs/mapping-format.md) for exact fields, validation,
provider rules, callback lifetimes and completion safety; the
[configuration examples](examples/README.md) explain loader semantics.

Focus actions remain available in custom mappings, for example
`{ "fl", "pane.focus", desc = "Focus right", args = { direction = "right" } }`.
The Focus group and Rex-known-hosts listing are **not defaults**. To opt into
inventory explicitly, include a provider group in your own `spec`:

```lua
{ "h", group = "Rex-known hosts (listing only)", provider = "known-hosts" }
```

The optional `rex_whichkey.defaults.hosts` fragment is also retained for explicit
composition. With no enabled provider marker, no host inventory lookup occurs.
Adding a custom `spec` replaces the defaults rather than appending to them.

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

Built-in bindings are Go-owned data in `defaults.go` for basic startup; the
optional `rex_whichkey.defaults` Lua fragments supply the same keys/intent:

| Keys | Actions |
| --- | --- |
| `wn`, `ws`, `wh`, `wl`, `wr`, `wx` | Session new, switch, previous, next, rename, close. Switch delegates Rex's picker; no enumerated session submenu is added. |
| `tn`, `th`, `tl`, `tr`, `tx` | Tab new, previous, next, rename, close. |
| `t1`–`t9` | Go to a tab by its one-based index. |
| `v`, `s` | Split right / down. |
| `z`, `x`, `e`, `c` | Zoom, close pane, toggle sidebar, toggle command palette. |

Session/tab/pane close bindings require confirmation. Named actions and callbacks
share the confirmation path. Native interaction has **no fixed idle deadline**;
RPC/preflight, dispatch and cleanup operations still have bounded timeouts. The
foreground parent watches the captured origin/app generation, and the child is
tied to that parent by a private lifetime pipe, not a daemon.

Configured Lua invocations retain their authenticated 30-second lifetime, starting
before mapping evaluation/preparation, with five seconds reserved for teardown
and completion. Expected UI expiry dismisses without a selection; an expired
capability is never renewed or used for completion. Real child, terminal-restoration
or cleanup failures remain errors. Both routes restore exact origin focus, close
only the owned popup and unlink private transport before selected dispatch,
including actions that may close the caller. Removing the Focus menu does not
remove that safety restoration. Uncertain actions are never retried.

Positions are `top`, `bottom`, `left`, `right` and `center`. Columns and paging
adapt to terminal size; very small viewports omit chrome before content. Unicode
cell-width calculation is not proof of actual client glyph geometry. Keys use
cyan/blue accents, icons yellow and labels separate primary/group roles. Colors
appear immediately using terminal-native ANSI accents even without palette replies;
valid replies enable contrast-aware styling. See [appearance](docs/appearance.md)
for fallback and accessibility limits. A nonempty `NO_COLOR` in the **rendering
process's** environment disables colors; normal rendering runs in a Rex-created
server-host child, and the launcher does not forward the caller's environment
as a popup override. Low-level `menu` and `popup` modes accept `--ascii` and
`--no-color`; these are not Lua setup options, and the delivered launchers do not
forward those flags. ASCII decorations do
not transliterate user labels. `menu` consumes resolved JSON and returns a
selection/null; it is not the normal Lua launch or an action executor. `prepare`,
`controller` and `popup` are internal plumbing, not manual invocation examples.

## Troubleshooting and safety

- **Module not found:** check the server's module installation/config directory
  and that the user module has a name other than reserved `rex_whichkey`.
- **Missing helper/CLI or action metadata:** check the invoking/server host's
  PATH and compatible installed Rex CLI. Basic startup needs no custom Lua
  actions. For optional configured launch, also check the server's Lua modules,
  completion registration, successful `menu.setup()` and `menu.register_launch`
  (matching an explicit `--action`). `REX_BIN_DIR`, if set, must name an absolute execution-host directory
  containing the matching `rex`. Raw child diagnostics are suppressed to avoid
  leaking capability data.
- **Missing/stale origin or client-selection failure:** launch inside the original
  Rex terminal with its real full `REX_SESSION`/`REX_BLOCK`. No attached live app
  means the graphical app must connect to that session; several attached live apps
  require `--client` with the exact full app ID. CLI/automation connections do not
  cause ambiguity. An invalid supplied selector requires correcting `--client` or
  a stale/empty `REX_CLIENT`; an absent `REX_CLIENT` is normal, not a setup failure.
  Never substitute focus/current-client selection. A reconnect's changed
  connection generation requires a new invocation. Captured delivery evidence
  later showed one attached app; that observation does not certify other sessions
  or later revisions.
- **Discovery/endpoint failure:** check the already-running execution-host server,
  Rex runtime/home settings and CLI compatibility. Fix malformed explicit or
  `REX_SERVER` values; a custom/shared endpoint may need both a public override and
  a matching Lua constructor/server environment. A viewing client's socket path
  or a different server realm is not a substitute. No server is autostarted and
  no alternate endpoint is tried.
- **Activation failed or uncertain:** verify app Remote Control permission and
  action compatibility; configured launch also needs its registered action/setup.
  Metadata, creation, wait, cleanup and dispatch failures have distinct safe
  diagnostics; arbitrary configured-action output is never printed. Configured
  activation acceptance can precede a later preparation/controller failure. Do
  not replay an uncertain launch/action or switch execution routes to compensate.
- **Mappings disappear or conflict:** `spec` replaces defaults; duplicate actions,
  action/group prefix conflicts and unsupported key forms are errors. Import,
  evaluation or validation failure is not a successful empty configuration.
- **Actions rejected:** Remote Control policy, missing/disconnected origins,
  connection-generation changes, expiry and configuration replacement can reject
  a selection. Do not work around this by retargeting the popup/current client,
  replaying completion or retrying an uncertain destructive operation.
- **Explicit host list empty or warning:** inventory is opt-in, not a default menu
  entry. It belongs to the execution host and its configured Rex home, not the graphical client or `--server` target. Listing can
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
