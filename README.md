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
source changes alone do not verify stable runtime bounds. Phase14 historically
added an app chooser and numbered-tab adapter; both were subsequently removed at
the user's direction. Public app routing uses the launching pane's validated GUI
owner unless explicitly overridden; key/palette launch retains its actual app.
New Tab uses the typed source-session resource API and exact-ID activation; app dispatch uses explicit `client.perform`.
These source repairs do not certify every default action or native UI behavior.
**Hiding native Rex pane/tab/window controls (REQ-026) remains blocked**,
not waived or fulfilled;
Phase11/project delivery is partial while that capability is unresolved.

Phase15 adds code-owned discovery and startup with an original-VM compatibility
bridge. Prior runtime source review/compilation is **not observed runtime verification**;
this packaging/docs cutover has parse/static evidence only;
no new installation, callback, remote, reload or rename result is claimed here.

## Install on the Rex server host

The helper and `rex` CLI must be available on the **execution/server host**,
even when the graphical client is on another machine. Optional customization
also needs Lua modules available to the user/process running the server. The
implementation targets Linux and macOS; declared Nix outputs are not platform
certification. Basic startup creates a native floating terminal; selected app
actions use explicitly targeted `client.perform` API requests through the Rex CLI.
Native New Tab first creates one source-session tab through the typed resource API.
Configured Lua launch additionally uses
configuration actions and detached process execution.

The flake exports `packages.<system>.rex-whichkey` (also `default`) for
`x86_64-linux`, `aarch64-linux`, `x86_64-darwin` and `aarch64-darwin`.
With a licensed CLI dependency, normal installation includes:

```text
bin/rex-whichkey -> ../share/rex-whichkey/lua/rex_whichkey/runtime/rex-whichkey
share/rex-whichkey/lua/rex_whichkey.lua
share/rex-whichkey/lua/rex_whichkey/...
share/rex-whichkey/lua/rex_whichkey/runtime/anchor.lua
share/rex-whichkey/lua/rex_whichkey/runtime/rex-whichkey (real executable)
share/rex-whichkey/lua/rex_whichkey/runtime/deps/rex -> licensed Rex package's bin/rex
share/doc/rex-whichkey/{mapping-format,runtime-bootstrap}.md
```

This is **layout revision 1**, declared by the normal anchor module, not a
generated runtime configuration. In a complete installation the real executable
lives beside the anchor; the public `bin` entry links to it, not the reverse.
Go's implicit layout uses that executable location to reach `deps/rex` when
`REX_BIN_DIR` is absent. Lua uses its explicit file-loader anchor to reach the
same image and dependency.
No wrapper, generated runtime path or constructor wiring implements discovery.
This topology is source evidence, not an installed public/key launch result.
On the two Linux architectures and Apple-silicon
macOS, the installer packages the official pinned `0.1.0-preview.13` Rex CLI
dependency described in [runtime bootstrap](docs/runtime-bootstrap.md).
It preserves the released macOS signature, Linux loader/glibc/PAM paths and
upstream terminfo/SSH-guard assets. These assets do not enable a service or edit
SSH/user configuration. Rex is proprietary: applicable upstream terms and Nix's
unfree-package policy apply. WhichKey's MIT license does **not** relicense Rex or
authorize redistribution of its binary/cache closure. No Rex bytes are vendored
in this repository. The pin is an installation source, **not** a minimum tested
release or a version users must pair with the server.

**Intel macOS:** that pinned upstream source has no artifact, so the default
package omits `runtime/deps/rex`. It keeps the real executable at
`bin/rex-whichkey`, outside the incomplete integration layout, and links
`runtime/rex-whichkey -> ../../../../../bin/rex-whichkey`. Both configured
public launch and key/palette launch reject that installation before popup,
even with a valid `REX_BIN_DIR` or PATH CLI: the anchored layout check always
requires the dependency. PATH repair cannot complete it, and configured failure
never becomes native fallback. The existing package override
`rex-whichkey.override { rex-cli = licensedHostRexPackage; }` can install a
legitimately available execution-host dependency exporting `bin/rex`; the name
here is a user-supplied derivation, not a claimed upstream Intel package. Home
Manager can select that result with `programs.rex-whichkey.package`. This is
ordinary dependency installation, not runtime selection or a required constructor
setting; no new artifact, platform or licensing entitlement is asserted.

Default packaged native operation on Intel macOS is retained, as is direct/local
standalone use. With no `REX_BIN_DIR` preference and no applicable implicit layout,
Go uses an already installed execution-host CLI from a valid unambiguous absolute
PATH. No environment repair or dependency override is needed solely for native
use. Normal origin/endpoint/CLI prerequisites still apply; the implicit default
must be truly absent and no explicit action required. A valid terminal
`REX_BIN_DIR` remains an optional preference. This is not fallback from a broken
applicable layout or a configured failure: both configured routes still reject
until their licensed dependency is installed.

These are **UNEXECUTED source scenarios**, not installed runtime results:

| Installation / intent | Source outcome |
| --- | --- |
| Default Intel package, no environment preference, valid execution-host PATH CLI, absent implicit default/no explicit action | Standalone public image has no implicit layout; normal native route remains available. |
| Complete dependency installation, no environment preference | Public image's implicit layout selects its normal dependency; key/palette uses the explicit anchor. |
| Default Intel configured public or key/palette | Explicit anchored dependency validation fails before popup; never native fallback. |
| Legitimate `rex-cli` package override | Installs the dependency and complete topology; Go still validates assets, compatibility and endpoint intent. |
| Invalid environment preference or broken applicable layout | Reject without trying a lower-precedence CLI. |

The basic menu needs no Lua configuration. Optional customization needs the
normal Lua namespace exposed to Rex; merely installing a package does not
register a menu. Non-Nix installers must preserve the same relative assets and
licensed dependency, not embed helper paths in configuration.

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
own declarative layout. Enabled/disabled behavior, custom locations, symlink
updates, conflicts and uninstall preservation are unexecuted source scenarios,
not Home Manager activation results.

For Lua customization without Home Manager, expose the complete installed Lua
entry point and `rex_whichkey/` directory together through Rex's normal module
lookup beside your user-owned configuration. Use the standard file loader; do not
copy just the entry point, synthesize an anchor, or mix module trees. The
file-backed anchor starts its fixed sibling helper even with restricted server
PATH when the complete host dependency layout above is installed; Go validates
and selects the normal dependency. Home Manager links only
this namespace, with no runtime-directory injection or per-host discovery.
Neither installation method edits or reloads `init.lua`.

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

No helper or CLI-directory constructor wiring is needed. Present
`execution.helper` or `execution.rex_bin_dir` fields fail with migration guidance;
users remove those obsolete fields from their own configuration.

### Development and normal upgrades

A direct local build uses the running Go binary for public preparation, controller
and popup; a registered menu no longer substitutes an installed popup helper.
The installed helper is only a finite relay for public launch, preserving the
original VM's custom closures. Key/palette launch instead uses its anchored helper.
An arbitrary standalone binary outside that layout needs terminal `REX_BIN_DIR`
or an unambiguous execution-host PATH for initial CLI discovery. This guarantee
for terminal children is **not** a configuration-action environment guarantee.

Install software/assets through your ordinary installer when upgrading. Do not
copy store paths or repeatedly refresh a consuming flake to repair normal
startup. Compatible public/helper/loaded-Lua build IDs may differ: Go negotiates
real protocols and the **loaded** consumer's features before any popup. Missing
or unsupported old bridges fail before invoking an old launch action. Installation
cannot replace already-loaded closures; the user may load the updated integration
through their normal Rex configuration lifecycle. This tool never edits, reloads
or restarts configuration and does not ask users to pair component versions.

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
suffixes fail early. Native resource closes (`x`, `tx`, `wx`) use their captured
pane/tab/session IDs directly after confirmation and complete popup cleanup;
they do not need an app chooser or native action registration. Placement is
captured before the menu and rechecked before closing, so a moved origin refuses
rather than closing a different tab. An explicit app selector is still checked
strictly; it is not silently ignored.

For actions that invoke app UI, native mode resolves app intent **only after
executable selection and exact owned-popup/file cleanup**; Escape returns without
app lookup. **All app actions** use the same pane-owner route below; there is no
app-selection menu or per-action fallback. This includes session/tab previous
and next, new session/tab, splits, sidebar, palette and Switch Session. Default
rename instead uses its exact captured resource API.
Configured CLI launch uses that same binding before activating the original Lua
VM; Lua-invoked menus keep their actual initiating client context.

Rex's native `info.kind == "app"` excludes CLI/automation and unknown kinds.
Names, focus and last-resize reports never select a target. The documented block
owner supplies app routing as described below. `REX_CLIENT` is normally absent.
No eligible app or a stale/non-app/unattached explicit selector refuses **native
action targeting**, not popup display. An invalid explicit target never falls back
to another app. The retained action, arguments and confirmation stay unchanged;
native New Tab intent is compiled as described below. The popup, files and
lifetime pipe are cleaned before dispatch. The actual
dispatched action must be available; its exact client ID and captured
`connected_at` generation plus origin placement are rechecked immediately
before dispatch, with no reconnect retargeting. Interaction remains deadline-free
and signal-cancellable; the short dispatch budget starts after selection and
cleanup.

Native `v` (Split right) and `s` (Split down) split the focused pane in the
**chosen app's currently selected tab/window**, not necessarily the launch terminal's
window. These splits omit destination session/block overrides when asking the app
to act; the captured origin still authorizes invocation and liveness, not the
destination.

Numbered “Go to tab” entries (`t1`–`t9`) are not included in the default menu.
Tabs retain New, Previous, Next, Rename and Close.

### Tab rename input screen

In both the native menu and the registered Lua defaults, Tabs → Rename (`tr`) opens a text-entry screen
inside the same popup and fixed bounds, prefilled with the current tab label.
Typing replaces the selected name; arrows/Home/End edit it. Enter confirms,
Escape returns to Tabs, and Ctrl-U clears the field. Backspace/Delete handle
whole graphemes; bounded paste cannot submit or trigger mappings.

After popup cleanup, native dispatch renames the captured tab through the Rex CLI;
configured dispatch authenticates the name against the original Lua invocation
and calls `session.set_window_label` with the retained tab ID. Neither route opens
Rex's rename dialog. Names must be nonblank text without control characters, at
most 4096 bytes and 128 Unicode characters (Rex's label limit). Custom callbacks
or replacement actions at `tr` keep their own behavior; only the default empty-argument rename intent opens this input page.
See [prototype design and scope](docs/tab-rename-input.md).

### Session switch and rename screens

In the **built-in native menu**, Sessions → Switch (`ws`) is our submenu,
populated from Rex's session list each time the menu opens. Rows identify the
session, show its name and mark the current session. Use numbered keys, More
(`n`) groups and normal Tab paging; Escape/Backspace returns to Sessions. The
parent retains exact destination IDs and revalidates the choice before sending
`session.select` to the pane's GUI app. It does not open Rex's session picker.

Sessions → Rename (`wr`) uses the same text-entry screen as tab rename, prefilled
with the current session name. Enter confirms; Escape returns to Sessions. After
cleanup, the parent renames the captured session directly—no native dialog or
app chooser. The registered Lua defaults use the same input page and authenticated
`session.set_label` dispatch. Custom callbacks/replacement actions remain unchanged.
The loaded consumer must advertise text/resource-rename support before popup creation;
unsupported consumers reject rather than opening a native dialog or dropping text.
See [session UI details](docs/session-owned-ui.md).

### Native app actions

Native `tn` retains Go-owned `client.tab.new` intent with empty arguments, but Go
creates exactly one default-terminal tab in the captured origin session through
`Session.NewWindow`, then activates its exact returned ID with `session.select`.
The chosen app must register that actual action before creation and again before
activation; origin/app generation is checked before each effect, and window
membership before activation. All UI/files/leases are gone before creation.
If creation is uncertain, inspect the origin session before any manual action.
The pane owner is checked before creation. Creation may legitimately change
focus/ownership, so activation retains the same bound app and rechecks its
connection and origin placement instead of choosing a new owner.
If activation fails, the tab is preserved and the error includes its exact ID
for manual selection. There is no retry, duplicate creation, automatic close or
fallback to `client.tab.new`. This translates intent rather than claiming Rex's
native New Tab action executed; configured Lua routing is unchanged.

Built-in tab/session rename and session selection use our screens above. The
command palette still uses Rex's native UI. Other app-routed actions retain their
action semantics; selecting a session uses the API directly, not Rex's picker. An explicit `--client` or
`REX_CLIENT` binding takes precedence. Otherwise Go reads the
launching pane's documented `owner` from `session.describe_block`, then validates
that exact ID as a live graphical app attached to the captured session. It does
not open the session-wide app chooser. Immediately before dispatch it rechecks
app generation, original tab/layer placement and unchanged owner; changes refuse
rather than retarget. An absent/non-GUI owner requires an exact app binding,
never a first-app, label, last-resize or remembered-choice fallback. Ownership is
an explicit routing policy, not proof of which device sent a keystroke and not
permission to bypass Rex's Remote Control checks.

Other app-routed native actions use one documented `client.perform`
request with the original action/JSON arguments and explicit captured origin
`session_id`/`block_id`. A separate unscoped client pins the exact server and
three-second timeout; splits and exact session selection omit both context fields. Errors, invalid replies
and uncertain effects are not retried, and there is no `rex do` fallback.
Configured menus retain their original-VM contract. Server-wide active-window
state is not used to choose a split destination.

Configured success acknowledges activation, not display or callback completion.
Native success follows child exit, exact cleanup and optional one-shot action
acceptance (for `tn`, one creation followed by exact-ID activation). Failed or
uncertain activation is never retried automatically.

### Linux/macOS endpoints and overrides

Go first freezes one execution-host CLI: a supplied trusted internal selection,
otherwise nonempty `REX_BIN_DIR`, a validated applicable layout dependency, or
an unambiguous absolute PATH candidate. Invalid preferences/broken layouts fail
without fallback; empty, relative/cwd PATH entries and distinct candidates reject.
Go retains CLI identity/content and read-only API-schema evidence through endpoint
discovery, controller, PTY, resource calls and completion.

Endpoint selection is independent: present public `--server` (or key constructor
intent) > nonempty `REX_SERVER` > bounded no-autostart Linux/macOS native Rex
discovery. No socket/home/store/repository guesses, shell wrappers or server
autostart are used. Accepted endpoints are concrete absolute Unix socket paths/
URIs or numeric-port `host:port`; opaque Lua `ctx.server` is never an address.

`popup.new(defaults, endpoint)` remains optional immutable deployment intent,
outside mapping options. Key launch uses that precedence. Public Go has already
selected its endpoint; a present constructor value must canonically agree with
it and is retained byte-for-byte through all relay calls. Empty, malformed or
conflicting intent fails before evaluation/authority/popup, never ignored,
overridden or used to retarget. For a deployment actually using the shared socket,
`popup.new(defaults, "/run/rex/server.sock")` may agree with public
`--server /run/rex/server.sock`; this is not a platform default or mandatory setup.
Explicit client intent remains strict as described above.

Public launch and the original VM must share execution-host private files.
A remote viewing Mac supplies neither Linux paths nor execution authority.
A remote endpoint alone does not deploy code or make cross-host public launch
supported. Use the server-host terminal or key action. See the
[launch contract](docs/launch-contract.md), [runtime bootstrap](docs/runtime-bootstrap.md)
and [CLI/GUI delegation](docs/cli-gui-delegation.md) for exact boundaries.

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
| `wn`, `ws`, `wh`, `wl`, `wr`, `wx` | Session new, switch, previous, next, rename, close. Built-in Switch uses our dynamic session submenu; Rename uses our text-entry screen. |
| `tn`, `th`, `tl`, `tr`, `tx` | Tab new, previous, next, rename, close. |
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
- **Helper/layout or CLI failure:** install the normal execution-host asset tree
  and licensed dependency; use the standard unshadowed Lua namespace. A standalone
  binary may use terminal `REX_BIN_DIR` or valid unambiguous PATH. Do not repair
  constructor paths or copy viewing-host binaries. Invalid preferences do not
  fall back. Check compatible read-only CLI/server metadata, not matching release
  strings; raw child diagnostics/private handoffs are suppressed.
- **Runtime compatibility failure:** update the integration normally and load it
  through your own configuration lifecycle if the loaded bridge is unsupported.
  Compatible unequal build IDs are allowed. Check setup/unique completion and
  launch registration; do not auto-reload or switch a failed custom menu to native.
- **Missing/stale origin:** launch inside the original Rex terminal with its real
  full `REX_SESSION`/`REX_BLOCK` and execution endpoint. No focus fallback exists.
- **Native action-targeting refusal / configured client-selection failure:**
  app actions require the launching pane's owner to be a live graphical app in
  that session, or a valid explicit override. Several attached apps do not cause
  an app-selection prompt. CLI/automation/unknown owners do not count. Fix malformed `--client`/empty
  `REX_CLIENT` before launch; stale but valid IDs are checked only after native
  selection, or before configured
  activation. Native display/Escape need neither a singleton app nor live selector.
  An absent `REX_CLIENT` is normal. Never substitute current focus or guess an app;
  changed generation prevents dispatch. Historical one-app captures do not certify
  other sessions or this deferred-binding revision.
- **Discovery/endpoint failure:** check the already-running execution-host server,
  Rex runtime/home settings and CLI compatibility. Fix malformed explicit or
  `REX_SERVER` values; custom/shared endpoints require canonical agreement with
  any optional Lua constructor intent. A viewing client's socket path
  or a different server realm is not a substitute. No server is autostarted and
  no alternate endpoint is tried.
- **Activation failed or uncertain:** verify app Remote Control permission and
  loaded feature compatibility; configured launch also needs its registered action/setup.
  Metadata, creation, wait, cleanup and dispatch failures have distinct safe
  diagnostics; arbitrary configured-action output is never printed. Configured
  success requires finite controller startup acknowledgment, not observed display
  or later callback completion. Do not replay an uncertain launch/action or
  switch execution routes to compensate.
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
  label to compensate. Once controller start is possible, lost acknowledgment does
  not authorize retry, killing an uncertain child or deleting its workspace.
  Malformed/unauthenticated handoffs may also leave artifacts; expiry revokes
  authority but does not prove artifact cleanup. See the source-linked
  [fixed owned cleanup inventory](docs/launch-contract.md#fixed-owned-cleanup-inventory);
  retired prompt resources are not part of cleanup.

Do not share private handoff/result files, capabilities or captured preparation
output. Terminal restoration and cleanup are bounded best effort, not guarantees
after SIGKILL or server/terminal loss. See [remote compatibility](docs/remote-compatibility.md)
for server-local paths, connection identity, moved-block handling and race limits,
and [the release inventory](docs/acceptance.md) for what has and has not been verified.
