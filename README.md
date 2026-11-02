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
[delivery inventory](docs/acceptance.md) for requirements and revision-bound
results. Phase09 adds default Zoom removal and focus-in repaint; earlier runtime
observations do not verify these changes or native GUI tab-return visibility.
Phase10 separates native display from deferred app-action targeting; earlier
observations do not verify that boundary. No minimum tested Rex build or general
runtime certification is claimed. Phase11 adds compact status borders, conditional
paging hints and measured content-fitting layer height. Phase12 adds compact
first-frame sizing and status spacers. The later Phase13 correction retains one
initial rectangle through navigation, superseding live grow/shrink behavior;
source changes alone do not verify stable runtime bounds. Phase14 adds an explicit
app chooser after native action selection when several apps are attached, and
translates native numbered-tab intent to exact window IDs. These source repairs
do not certify every default action or GUI-only behavior.
**Hiding native Rex pane/tab/window controls (REQ-026) remains blocked**,
not waived or fulfilled;
Phase11/project delivery is partial while that capability is unresolved.

## Install on the Rex server host

The helper and `rex` CLI must be available on the **execution/server host**,
even when the graphical client is on another machine. Optional customization
also needs Lua modules available to the user/process running the server. The
implementation targets Linux and macOS; declared Nix outputs are not platform
certification. Basic startup creates a native floating terminal; selected actions
use explicitly targeted CLI activation. Configured Lua launch additionally uses
configuration actions and detached process execution.

The flake exports `packages.<system>.rex-whichkey` (also `default`) for
`x86_64-linux`, `aarch64-linux`, `x86_64-darwin` and `aarch64-darwin`. The package
layout is:

- `bin/rex-whichkey`
- `share/rex-whichkey/lua/rex_whichkey.lua` and `rex_whichkey/`
- `share/doc/rex-whichkey/mapping-format.md`

Rex itself is not bundled. The executable plus a compatible server-host Rex CLI
and valid captured session/block suffice to create the basic menu; no graphical
app selection or custom action registration is required for display/dismissal.
Executing a selected native action does require an eligible live graphical app.
Installing the executable does not put optional Lua modules in Rex's search path
or register a customized menu.

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
| `--client FULL_ID` | Exact full `client:UUID` requested for app actions. Native display validates syntax only; configured launch requires a live attached app immediately. |
| `--action NAME` | Require an installed Lua launch action; no native fallback. Omit to prefer registered `whichkey`, otherwise use basic defaults. |

The command captures this terminal's full `REX_SESSION` and `REX_BLOCK` and checks
those resources through the captured endpoint before choosing native/configured
mode from server action metadata. **Native display, navigation and dismissal do
not inspect graphical apps or their action registry.** Zero/several apps, unknown
attachment kinds or a syntactically valid but nonexistent requested client do not
prevent native display/dismissal. A graphical app still has to present the server's
terminal for a user to see it; resource creation alone is not graphical proof.

App selector precedence is `--client`, then a present `REX_CLIENT`, otherwise
eligible live apps in the captured session. Empty/malformed selectors, labels and
suffixes fail early. Native mode resolves that intent **only after executable
selection/confirmation and exact owned-popup/file cleanup**; Escape returns without
app lookup. With no selector, one app binds directly; several apps open a separate
**choose app** menu with numbered choices, safe display labels and full app IDs.
Escape at its root (or unmapped `q`) cancels with no action. Its initial rectangle
also stays fixed through navigation. This is an action-target prompt, not live
submenu resizing or a saved app preference. Configured mode still binds before
launch and requires an exact selector when several apps are attached.

Rex's native `info.kind == "app"` excludes CLI/automation and unknown kinds; names,
focus and terminal ownership never select a target. `REX_CLIENT` is normally absent.
No eligible app or a stale/non-app/unattached explicit selector refuses **native
action targeting**, not popup display. An invalid explicit target never falls back
to the chooser. The retained action, arguments and confirmation survive target
choice unchanged; native numbered tabs are compiled as described below. Both
popups, files and lifetime pipes are cleaned before one dispatch. The actual
dispatched action must be available; its exact client ID and captured
`connected_at` generation plus origin placement are rechecked immediately
before dispatch, with no reconnect retargeting. Interaction remains deadline-free
and signal-cancellable across both stages; the short dispatch budget starts after
choice and cleanup.

Native `v` (Split right) and `s` (Split down) split the focused pane in the
**chosen app's currently selected tab/window**, not necessarily the launch terminal's
window. These splits omit destination session/block overrides when asking the app
to act; the captured origin still authorizes invocation and liveness, not the
destination.

Native `t1`–`t9` retain their canonical `client.tab.goto` one-based integer intent,
but Go compiles it to documented `session.select` with the captured origin session
and an exact parent-resolved window ID. After all UI cleanup, Go reads that
session's ordered tabs, resolves the index once and freezes the ID. It rechecks
origin/app generation and window membership before dispatch; a missing index or
removed/moved tab refuses without fallback. Reordering never substitutes a new
ID. The chosen app must register `session.select`; the child cannot replace the
action or args. This is intent translation, not a claim that native numeric goto
was repaired. Configured Lua defaults keep their direct native action route.

Other native actions (including close) retain their existing origin-targeted
context, and configured menus retain their original-VM contract. Server-wide
active-window state is not used to choose a split destination.

Configured success acknowledges activation, not display or callback completion.
Native success follows child exit, exact cleanup and optional one-shot action
acceptance. Failed or uncertain activation is never retried automatically.

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
exact server/session/block checks remain mandatory. Configured launch and native
selected-action dispatch additionally validate app identity/generation. This
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
Zoom also remains available explicitly as
`{ "z", "pane.zoom", desc = "Zoom pane" }`, but is not a default binding.
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
| Tab | Next page when the menu spans multiple pages, wrapping to the first. |
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
| `x`, `e`, `c` | Close pane, toggle sidebar, toggle command palette. |

Session/tab/pane close bindings require confirmation. Named actions and callbacks
share the confirmation path. Native interaction has **no fixed idle deadline**;
RPC/preflight, dispatch and cleanup operations still have bounded timeouts. The
foreground parent watches captured server/session/block liveness, not app
attachments or generation, and the child is tied to that parent by a private
lifetime pipe, not a daemon. Switching tabs alone is not cancellation. The terminal saves/enables/restores xterm focus-reporting
mode 1004; a complete focus-in report repaints the existing menu with its current
submenu/page/confirmation and theme, without navigation or action dispatch.
Focus-out does not cancel or steal focus. No periodic repaint or replacement
popup is created. This depends on the terminal delivering focus reports; it is
not a guarantee that every native app restores a floating layer visually.

Configured Lua invocations retain their authenticated 30-second lifetime, starting
before mapping evaluation/preparation, with five seconds reserved for teardown
and completion. Expected UI expiry dismisses without a selection; an expired
capability is never renewed or used for completion. Real child, terminal-restoration
or cleanup failures remain errors. Both routes restore exact origin focus, close
only the owned popup and unlink private transport before selected dispatch,
including actions that may close the caller. Removing the Focus menu does not
remove that safety restoration. Uncertain actions are never retried.

Positions are `top`, `bottom`, `left`, `right` and `center`. The title/breadcrumbs
and concise key hints occupy separate full-width status-border rows above/below
the entries, with one blank content/status spacer beside each border when room
permits, without a `Keys:` prefix or a redundant `Page1/1` row. A real multi-page
menu retains a counter and Tab hints when space permits; tiny viewports drop
spacers before content or essential hints. These terminal borders are not Rex's
native controls.

The parent computes a **content-sized first layer** before its single creation
RPC and seeds terminal startup rows/columns consistently. It uses renderer natural
rows (including spacers), the exact origin's window-relative leaf rectangle and
validated measured terminal dimensions. Two horizontal/four vertical cells of
inset headroom and half-row rounding estimate new-pane font/chrome differences;
they are not fixed font pixels or a pixel-tight guarantee. Missing/unmeasured
origins use a compact bounded fallback (hypothetical 100×40 window-cell grid,
at most 30% initial height).

That **initial rectangle stays fixed** through submenu entry, paging, back,
focus and idle. Larger menus use the current terminal cells and existing
pagination, not a larger box. External window/font changes can naturally change
terminal cells; SIGWINCH and focus-in still redraw, without application-driven
layer resizing. The former preferred-row hint IPC and live resize loop are
retired. No app lookup is added for geometry, and no font/settings changes are
made. Native controls are **not hidden**: no
supported per-popup option was found in the retained Rex API evidence, and no
global workaround is used. See [appearance](docs/appearance.md) and the
[compact-presentation evidence](docs/acceptance.md#compact-popup-presentation-req-024025-and-blocked-req-026)
for the one-context observations and remaining limits.

Unicode cell-width calculation is not proof of actual client glyph geometry.
Keys use cyan/blue accents, icons yellow and labels separate primary/group roles. Colors
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
- **Missing/stale origin:** launch inside the original Rex terminal with its real
  full `REX_SESSION`/`REX_BLOCK` and execution endpoint. No focus fallback exists.
- **Native action-targeting refusal / configured client-selection failure:** no
  eligible app means connecting one to that session before trying a new action;
  several eligible apps offer a native action-time chooser when no selector was
  supplied. Configured launch still requires an exact full app ID in that case.
  CLI/automation/unknown kinds do not count. Fix malformed `--client`/empty
  `REX_CLIENT` before launch; stale but valid IDs are checked only after native
  selection, or before configured
  activation. Native display/Escape need neither a singleton app nor live selector.
  An absent `REX_CLIENT` is normal. Never substitute current focus or guess an app;
  changed generation prevents dispatch. Historical one-app captures do not certify
  other sessions or this deferred-binding revision.
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
