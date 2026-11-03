# Popup launch, endpoint and preparation contract

## Ordinary startup and optional Lua setup

With the helper and matching Rex CLI installed on the execution/server host,
run inside the originating Rex terminal:

```sh
rex-whichkey
```

No arguments, custom action registration or hand-authored JSON are needed for
the basic Go-owned native-action menu. The production path creates a real Rex
floating popup, not a greeting, demo or standalone terminal menu. If read-only
server action metadata reports a registered Lua `whichkey` action, that original-VM
configured path is preferred instead. This routing is decided before activation;
a failed or uncertain activation never switches to native mode.
`rex-whichkey --help` (or `-h`) describes the public `rex-whichkey [options]`
entrypoint without contacting Rex. Public options are:

- `--server ENDPOINT`: a concrete execution-host endpoint override.
- `--client FULL_ID`: an exact `client:UUID` requested for app execution in this
  terminal's session. Syntax is checked early; liveness is checked after native
  executable selection or before configured launch.
- `--action NAME`: require an installed configuration launch action, including
  when NAME is `whichkey`. Missing registration is an actionable error, not a
  request for native fallback. Names contain only lowercase ASCII letters,
  digits and underscores.

The root parser rejects unexpected positional arguments, unknown flags, missing
flag values and an explicitly empty action before Rex calls. An explicitly empty
server/client is invalid, not an instruction to default. Existing private `menu`,
`prepare`, `controller` and `popup` modes remain compatible; they are not the
ordinary startup interface. Root launch does not read resolved configuration
from the terminal's stdin or run the standalone menu in place of a Rex popup.

The basic native menu needs no Lua modules, launch/completion registrations or
configuration reload. For **optional custom Lua mappings/callbacks**, install
matching server-host Lua modules and perform setup in the server's **user-owned**
`init.lua`:

```lua
local popup = require("rex_whichkey")
popup.register_completion("whichkey_complete")
local menu = popup.new(require("rex_whichkey.defaults"))
menu.setup() -- or menu.setup(inline_options, "my_whichkey")
menu.register_launch("whichkey") -- omitted name also defaults to whichkey
```

This registers a genuine `rex.action` whose `run(ctx, args)` receives the actual
fresh context in the original server configuration Lua VM. The named action also
works from keys/palette without an origin envelope; existing direct
`menu.launch(ctx)` calls remain supported. Completion and launch names must be
distinct; duplicate module launch registration fails. Use another registered name
with public `--action NAME`. The constructor's optional second endpoint override
remains available and must agree with the public caller's resolved endpoint.
For example, a deployment actually using Rex's documented shared socket can use
`popup.new(defaults, "/run/rex/server.sock")` and launch with
`rex-whichkey --server /run/rex/server.sock`. That is an explicit shared-server
example, **not** the personal Linux/macOS default. The public flag is not forwarded
as a Lua constructor override: preparation resolves separately in the server
process's environment. Matching canonical endpoints are required; a flag cannot
retarget a runtime hosted elsewhere. Changing a terminal's environment does not
change an already-running server's environment. No keybinding is installed;
bind the registered action through your existing Rex configuration if desired.
Modules never edit or reload `init.lua`. Remote Control must permit CLI activation
on the selected client; installing/registering these pieces is not automatic.

Missing helper/CLI, incompatible server metadata or unavailable session/block
can prevent basic startup. App availability, selected-action compatibility and
Remote Control gate native action execution, not native display/dismissal. The
configured path additionally needs a live selected app before launch, modules,
registration and successful setup. Diagnostics identify the failed stage without
forwarding raw child output/private handoffs. Nonzero/timeout/output-limit
activation failures are uncertain and never retried. A configured zero exit
means **activation was accepted**, not proof of display or callback execution.
Native success follows child exit/teardown and optional one-shot action acceptance;
RPC acceptance still is not proof of graphical display.

## Exact terminal origin and client selection

[`LaunchPopup`](../launch.go) captures execution-host `REX_SESSION` and `REX_BLOCK`
as full typed `session:UUID` and `block:UUID` IDs. Missing/malformed IDs fail with
instructions to run in the original Rex terminal. There is no focus/current-block
fallback, session label lookup, ID suffix or graphical-client filesystem lookup.
Resolve one canonical endpoint by the policy below, then use it explicitly for
all read-only preflight and activation calls.

Native display validates endpoint/session/block resources through
`popupRequest.resourceOriginBlock` before reading server registration metadata.
It does not call `session.list_clients` or query a selected-client action registry.
The native lifecycle uses resource-only liveness for creation, wait and server
focus restoration, and owned-resource checks for cleanup. Zero/several apps,
unknown attachment kinds and a valid-but-nonexistent explicit selector do not
prevent display/navigation/dismissal; they do not prove a GUI presented the layer.

Client precedence is present `--client`, then present `REX_CLIENT`, otherwise a
relationship lookup of **only the captured session's attached graphical apps**.
For native mode this lookup is deferred until an executable retained selection
has been confirmed and the owned popup/private transport cleaned. Cancellation
returns without looking up an app. Configured mode still performs it before
activation and retains that exact generation throughout its finite lifetime.
`session.list_clients` provides both exact session membership and Rex's native
`info.kind` classification in one response. Only the documented kind `app` is
eligible: `cli`, `agent`, `other`, missing and unknown kinds do not contribute to
automatic cardinality and cannot be explicitly selected. No label/name/platform
heuristics or unscoped global app list are used.

Present empty or malformed selectors and labels/suffixes fail early. A syntactically
valid native selector is Go-owned intent, not a fabricated live generation; its
liveness is deliberately not checked for display/Escape. At binding, malformed or
duplicate attached client IDs fail even for non-app/lost records; UUID case
variants cannot hide a duplicate. Each nonlost app attachment must have a nonblank
UTF-8 `connected_at` generation of at most 1024 bytes without control characters.
Rex's documented missing/unknown connection state counts as connected; only `lost`
is disconnected. The documented null/empty attachment list means no candidates;
a missing/malformed list is an inspection error, not a successful empty lookup.

Without a selector, exactly one live attached app binds directly, even when CLI
or automation connections are also attached. Rex terminals normally need not
export `REX_CLIENT`; its absence is not a prerequisite failure. No live attached
app produces guidance to connect the graphical app to the origin session.
Configured launch with several apps still requires an exact full app ID. Native
resource closes do not require app selection. Native rename binds the launching
pane's documented GUI owner, as described below; other native app actions use an
explicit target chooser when several apps qualify.
A supplied selector must match a live app in that exact session; stale, non-app or
unattached selectors refuse even if another app exists, never falling back to the
chooser. Native failures are **action-targeting refusals after popup cleanup**,
not launch prerequisites. No global, focused, last-resize or arbitrary-first
fallback is used. Rename's explicit block-owner routing policy is distinct from
using terminal-size reports as app identity.

`popupRequest.originBlock` combines `popupOriginConnection` generation revalidation
with `resourceOriginBlock`'s exact session and uniquely placed block check.
Configured launch/controllers use the combined check throughout. Native display
uses only resource checks; native app dispatch binds after selection and then uses
the combined check immediately before execution. Resource closes use resource
checks alone unless an explicit app selector was supplied. The Rex resource client rejects
malformed/duplicate window/layer/block placement. A wrong endpoint or stale origin
fails read-only before resource creation or activation. Preflight is not enduring
authority; each route independently validates at its mutation boundary.

## Native basic route and registration selection

`LaunchPopup` reads the captured server's `rex --server ENDPOINT
--autostart=false --timeout 3s actions --json` registry. It removes inherited
REX client/server/session/window/block selectors for that metadata process and
supplies captured session/block environment values; server metadata is deliberately
queried without a client selector. Client metadata uses the same explicit endpoint
plus `--client FULL_CLIENT`. These commands are read-only; they do not activate
or probe actions. Both have a three-second process deadline, 250 ms pipe shutdown
allowance, 1 MiB retained stdout limit, discarded 64 KiB stderr limit, and a
required single JSON object with an `actions` array. Invalid/duplicate identities
fail closed. Registry errors are never interpreted as absence of registration.

A matching `kind: "lua"` server action selects configured launch. A conflicting
non-Lua name is an error. Only absence of the **implicit** default selects
`launchNativePopup`; explicit `--action` absence fails. Once configured activation
begins, no outcome can select native mode or retry that activation.

[`defaults.go`](../defaults.go) constructs real canonical `Mapping`/`Config` data,
not an independent schema or a demo. Its ordered session/tab/pane/application keys,
icons, native picker/rename semantics and close confirmations match the shipped
Lua fragments. Neither default set includes Zoom (`z`), the Focus group or a
`known-hosts` provider. Explicit custom `pane.zoom` remains supported.
[`defaults/init.lua`](../lua/rex_whichkey/defaults/init.lua) composes only sessions,
tabs, panes and app fragments; no provider marker means no inventory
lookup in [`hosts.go`](../hosts.go). Custom `pane.focus` mappings and explicit
`provider = "known-hosts"` groups remain supported; the optional
[`defaults/hosts.lua`](../lua/rex_whichkey/defaults/hosts.lua) fragment is retained.
Numbered “Go to tab” bindings (`t1`–`t9`) are not included in either default
set; custom configured Lua actions remain supported. `pane.split` uses
`right`/`down`; custom `pane.focus` uses `left`/`down`/`up`/`right`.
No app registry is required to open the menu. Only the actual action to dispatch
must be reported as `kind: "client"` by the app selected after cleanup;
for native New Tab intent this is `session.select`, as described below.
Unrelated unavailable defaults do not block display or that selection.
Incompatibility refuses targeting with the actual Go-owned dispatch action name;
there is no speculative fallback. Rex remains responsible for final
action/schema/Remote Control checks.

[`native.go`](../native.go) retains an independent immutable parent snapshot and
passes canonical configuration to `runPopupWithLifetime` in
[`popup.go`](../popup.go) with the trusted native lifetime policy. The native CLI
stays in the foreground while its PTY child is usable, with **no fixed interaction
or idle deadline**. `RunNativePopupTerminal` requires a private 0600 lifetime pipe;
parent loss cancels the child even if parent defers cannot run. `waitPopup` waits
for the child's exit without capping interactive time at the RPC timeout, while
checking exact captured session/block placement every second. Native wait never
checks app attachments/generation; graphical disconnection/ambiguity cannot cancel
it. Cancellation, signals, terminal/process failure or lost origin still end the
invocation; there is no persistent daemon or detached indefinite controller.
The watcher does not inspect the active tab. Merely switching away/back is not
cancellation and causes no Rex focus mutation or layer recreation. Configured
wait alone also checks its pre-bound app generation.

The actual native path is `LaunchPopup` -> `launchNativePopup` ->
`runPopupWithLifetime(popupNative)` -> `RunNativePopupTerminal` ->
`runPopupTerminal` -> `RunTerminal`. Terminal ownership saves/enables/restores
xterm mode 1004. After ordered `PaletteDecoder` processing, `terminalKeys` consumes
complete unparameterized `CSI I` as a separate focus-in repaint signal, not a
mapping token. `RunTerminal` calls `drawTerminal` -> `RenderMenu` on the same
menu/theme, bypassing `HandleKey`, so submenu/page/confirmation survive. Fragmented
CSI and OSC-interleaved input retain framing; focus-out (`CSI O`) and other reports
cannot select or cancel. Existing Escape, mouse and paste protection is unchanged.
There is no unconditional repaint timer, tab-return replacement layer or
focus-stealing callback. A separate later action-target prompt is described below.
Focus reporting is a terminal protocol, not a claim of native-app layer persistence;
apps that do not emit it will not trigger this repaint.

Configured controllers use `runPopupLifecycle`, which always selects the finite
configured policy. Handoffs require the original future expiry within the canonical
30-second bound; absent/zero expiry cannot enable native mode. The launcher starts
its lifetime before evaluation/preparation, and the UI deadline is five seconds
before expiry to reserve teardown/completion time. `RunPopupTerminal` treats only
the exact expected deadline error, with successful terminal restoration, as null
dismissal. Expected controller expiry performs cleanup without stale completion;
no capability is renewed and no expired selection is dispatched. Invalid deadlines,
unexpected child status, I/O and restoration failures remain errors. Both child
routes share `runPopupTerminal` -> `RunTerminal` -> `RenderMenu`/`HandleKey`.

Preflight and individual RPCs remain bounded (20-second public preflight context,
three-second metadata/resource calls); selected dispatch is bounded to three
seconds and deferred cleanup to five. These operation limits are not native user
idle limits. Removing the Focus group does not remove exact origin-focus safety.

The shared lifecycle creates 0700 invocation-owned directories and exclusive
0600 config/result files and the native-only lifetime pipe. There are no
preferred-row hint files or geometry IPC. The child returns
only keys, optional opaque callback identity and explicit confirmation (or null
dismissal). The strict canonical
completion decoder rejects replacement action/args and checks confirmation and
identity against the retained snapshot. The native snapshot has **no callbacks**;
Go refuses callback references rather than evaluating Lua. Capabilities stay in
the parent; transport/config/result paths contain no action credentials.

After successful child exit (which follows terminal restoration), the controller
closes only its returned owned terminal, confirms its layer vanished, restores
exact origin focus and removes private config/result files, the native lifetime
pipe when present and their empty directory **before**
selected dispatch or any native app lookup. The parent's lifetime-pipe writer is closed in that cleanup
path too; close failure prevents dispatch. Native owned close inspects the exact
returned terminal in the captured session and can clean it even if the origin
block/app disappeared;
restoring focus still requires the exact origin block. No app is required for
native cleanup/focus. Configured mode retains its combined origin/app checks
before creation, close, focus and completion.

Native finish only captures the validated retained selection/confirmation; it
never opens another UI inside its short operation context. After the whole first
popup lifecycle returns successfully, cancellation returns without app lookup.
Only the selected immutable Go-owned mapping is retained for execution.

[`native_targets.go`](../native_targets.go) shares strict eligible-app enumeration
with `popupOriginConnection` without changing configured selector/cardinality
semantics. An exact requested selector remains strict. With no selector, zero
apps refuse, one binds its full ID/`connected_at`, and several create a private
Go-owned choice snapshot. A separate canonical `Config`/`Mapping` menu displays
full app IDs followed by bounded printable-ASCII label hints. Duplicate, empty or
unsafe labels cannot become identity; invalid label data gets a neutral fallback.
Keys use 1–9 and explicit More groups, not conflicting decimal prefixes. The
private selectable action tag is never dispatched to Rex. Child Completion is
validated against this retained config and its keys resolve only a retained app
ID/generation, never replacement original action/args/confirmation.

The chooser starts only after original popup teardown and uses the same native
FIFO/foreground ownership, resource-only liveness, initial sizing/spacers and
fixed-rectangle navigation. Its interaction has no idle deadline; Escape at the
root or unmapped `q` cancels with success and no action. Returning from it requires
successful owned UI/file/lease cleanup and exact origin focus restoration. Neither
stage caches an app preference. An outer native signal context spans both stages,
lookups, cleanup gaps and dispatch, while per-stage cleanup stays bounded and
parent loss closes each child's FIFO writer. The configured original-VM route,
capability/expiry and finite controller policy are unchanged.

After choice and cleanup, a fresh three-second dispatch context remains a child
of that outer signal context. Native dispatch revalidates the original selection,
checks the actual dispatch action's metadata on the bound app, then rechecks its
exact captured generation plus origin immediately before execution. Lost attachments,
reconnects or stale choices fail; no re-enumeration silently replaces the target.
The chooser cannot renew the captured generation.

Native actions other than numbered-tab and New Tab intent use documented `client.perform`
through a separate **unscoped** `rex.Client` pinned to the captured endpoint and a
three-second timeout. Each request preserves retained `client_id`, `action`, JSON
`args`, `origin: "cli"` and `timeout_ms: 3000`; empty args become `{}`. All non-split
requests also carry explicit captured origin `session_id` and `block_id`, preserving
origin-bound context including destructive close targets. Only `pane.split`
(`v`/`right`, `s`/`down`) omits both fields: the chosen app splits its currently
selected tab/window's focused pane.
`api call` does not default to ambient `REX_SESSION`/`REX_BLOCK`/`REX_WINDOW`, and
omitting `Options.Session` prevents the generic adapter injecting an origin session.
The original scoped client still checks invocation placement, session attachment
and exact app generation immediately before this one call. For splits, origin is
liveness and authorization evidence, not destination; server `active_window_id`
is not used as app-selected state. `Call(..., nil)` discards a validated JSON
result; timeout or invalid JSON may follow a successful mutation and never
permits retry or CLI fallback. The captured contract is
`.mtn-tmp/action-audit/client-perform-api.json`; missing `block_id` explicitly
means the client's focused block. Split arguments stay `right` or `down`, never
the CLI spelling `below`.

Accepted **REXWHICHKEY-20261102-171343** in the
[explicit-dispatch supplement](../.tasks/designs/rex-whichkey-client-perform.md)
extends this transport to the remaining native leaves. The coordinator recorded
the exact same installed executable returning success from `rex do client.tab.new`
without creating a window, while direct documented `client.perform` with the same
app/context created one. This is not version skew or proof that every CLI action
is broken; the underlying CLI cause is not established. That transport-only a01
candidate subsequently compiled and passed source review but **failed** the actual
`tn` effect gate: only six of seven required effects were observed. Tracing showed
`client.perform client.tab.new {}` with the retained exact app/origin; an interactive
zsh caller with valid OSC PWD also created no tab. The earlier comparison does not
establish reliable New Tab behavior or a shell/context diagnosis. Source changes
alone do not prove resource effects.

### Native New Tab intent compilation

Accepted **REXWHICHKEY-20261102-173105** in the
[New Tab supplement](../.tasks/designs/rex-whichkey-new-tab-intent.md) refines the
transport-only repair only for Go's built-in `tn`. [`native.go`](../native.go)
validates retained `client.tab.new`, key `tn` and absent or strictly empty-object
arguments. Neither the child nor the target chooser supplies replacement intent.
The chosen app must register native `session.select` before any mutation.

After both UI lifecycles and all private files/leases are gone, the parent resolves
the captured source session with `rex.Client.Session`, rechecks origin placement
and exact chosen app attachment/generation, then invokes `Session.NewWindow` once
with an empty label for a default terminal. The existing typed helper owns the
documented `window new` resource command and response validation; no guessed
creation schema or native action arguments are introduced.

The returned exact window ID is frozen. Before one activation the parent checks
`session.select` registration again, uses `Window.Inspect` to validate that ID's
source-session membership, and rechecks the original origin/app generation.
A separate unscoped exact-endpoint client sends `client.perform` with the chosen
`client_id`, `action: "session.select"`, explicit
`args: {session_id: ORIGIN_SESSION, window_id: CREATED_WINDOW}`, `origin: "cli"`
and `timeout_ms: 3000`, without outer session/block overrides. Creation and
activation share the fresh three-second signal-cancellable dispatch budget;
these checks are not an atomic resource lock.

An uncertain creation is reported as a possible effect without retry. If the
typed helper returns a valid ID alongside an incomplete creation response, the
error retains that ID and activation is not attempted. Every later registration,
membership, origin/generation or activation failure reports the exact created ID
and manual inspection/selection guidance. The tab is not automatically closed,
recreated or replaced; there is no fallback to `client.tab.new`. Success denotes
one creation plus accepted exact-ID activation, not a generic GUI certificate.
This compiles Go-owned intent rather than claiming native `client.tab.new` ran.
Configured Lua original-VM dispatch, split selected-app context and destructive
origin targets retain their existing contracts. Fresh review,
coordinator compilation and all seven actual resource effects remain required;
the failed a01 gate is not relabeled a pass.

### Historical native numbered-tab intent compilation

The user subsequently removed `t1`–`t9` from both default menus. This adapter and
its native routing branch were retired; the description below preserves the
previous design, not current source behavior.

Accepted **REXWHICHKEY-20261102-165714** in the
[numbered-tabs supplement](../.tasks/designs/rex-whichkey-numbered-tabs.md)
supersedes verbatim dispatch only for native `t1`–`t9`.
Historical `native_tabs.go` at commit `54a577bee6481ebe5110d7b280c03bcf307542c3` strictly decodes the original parent-retained
`client.tab.goto` Args as one integer `index` from 1 through 9, matching its retained
`tN` key. Missing/null/duplicate/unknown fields, coercions or out-of-range indices
refuse; no guessed zero-based/string args are sent. The canonical Go/Lua defaults,
original completion and confirmation remain immutable. Child data still carries
only selection identity/confirmation, never replacement actions or arguments.

Only after both menu/chooser lifecycles and all private files/leases are cleaned,
the parent uses existing typed `rex.Client.Session` / `Session.Windows` APIs to
read the captured origin session's ordered windows. It resolves the one-based
position once and freezes that exact window ID. A missing index refuses. Before
one dispatch it rechecks exact origin placement, chosen app attachment and
`connected_at`, then uses `Window.Inspect` to recheck the frozen ID's membership
in the same origin session. Removal or movement out of the session refuses;
reordering does not cause index re-resolution or substitute another window.
These are bounded pre-dispatch checks, not an atomic lock on Rex resources.

The actual app action must be registered as native `session.select`; checking only
`client.tab.goto` registration would not authorize this translated request. The
retained action contract in `.mtn-tmp/action-audit/work-actions.json` documents
`session.select` with required `session_id` and optional `window_id`. Go sends both
explicitly as action args, using one documented `client.perform` through a separate
unscoped exact-endpoint client: `client_id`, `action: "session.select"`,
`args: {session_id: ORIGIN_SESSION, window_id: RESOLVED_WINDOW}`, `origin: "cli"`,
`timeout_ms: 3000`. There is no outer session/block override or ambient current
client/server selection. All lookup/recheck/dispatch work shares the fresh
three-second signal-cancellable context. No missing-window/app fallback or retry
follows errors, including uncertain mutation/timeout/decode failures. A successful
receipt alone is not evidence that the app displayed the target tab.

This is intentional Go-owned intent compilation, not a repair of Rex's numeric
`client.tab.goto` implementation. Configured authenticated original-VM dispatch
and Lua defaults remain direct native actions. Native previous/next, splits and
all destructive actions retain their existing behavior.

Native `x`, `tx` and `wx` bypass app selection and native action registration.
After validated confirmation and full UI/file/lease cleanup, existing typed Rex
resource operations close the exact captured pane, tab or session. The origin's
window/layer placement is captured before the menu and revalidated before the
single close; a move refuses rather than selecting a new tab. Explicit app
selectors remain strict, with their connection checked before closing; without
one, these resource operations need no app lookup. Errors/uncertain closes never
retry. The route is compiled, not runtime-certified after the user's stop-testing
instruction.

Native tab/session rename (`tr`/`wr`) preserves Rex's native dialogs and original
session/block context. Explicit app selectors take precedence. Without one,
[`native_rename.go`](../native_rename.go) reads the launching pane's documented
`session.describe_block` owner and binds that exact live graphical app attached
to the captured session, including its connection generation. This is the
user-requested pane-owner routing policy, not proof of input provenance or extra
permissions. It bypasses the session-wide app chooser. Immediately before the
one `client.perform`, Go revalidates the app generation, original tab/layer and
unchanged owner. A missing/non-GUI owner, move or changed binding refuses; no
label, resize reporter, first/last app or remembered choice substitutes. Explicit
selectors keep their existing strict validation and need not match the owner.
No replacement rename prompt was introduced. The new route is compiled, not
runtime-tested under the user's stop-testing instruction.

Apart from resource closes and the retained New Tab intent adapter, every
remaining native action sends trusted action/JSON args once through the
explicit `client.perform` request above, with captured origin session/block
context; no native `rex do` dispatch wrapper remains. Metadata still legitimately
uses the read-only Rex `actions --json` command. Destructive close targets and
configured original-VM routing do not change. All native dispatch paths share
the fresh three-second signal-cancellable context, sanitized failure or
uncertain-effect errors and no retries or fallback. Dismissal of either stage
executes no native action.
Cleanup errors prevent selection dispatch; partial/uncertain creation never triggers broad
layer/window/session deletion. Deferred error/cancellation cleanup is bounded to
five seconds and never retries an uncertain close/focus/dispatch. Stale origins
can prevent safe cleanup; SIGKILL/server loss cannot be guaranteed recoverable.

`rex/layer.go` calls the documented `session.new_layer` API directly with the
exact inspected origin window and a `layout.block` shell creator, explicit argv,
`shell = "none"` and `exit.on_completion = false`. No script or callback VM is
created. Inspect/create/focus/dispatch are separate operations: concurrent moves,
disconnects or external focus changes can race the final check; there is no
claimed atomic transaction or current-client fallback.

## Initial popup geometry and terminal seed

[`popup_initial.go`](../popup_initial.go) computes bounds **before** the single
`session.new_layer` attempt. It constructs the retained canonical menu and calls
`RenderMenu` at the estimated popup columns with height zero to obtain natural
rows, including both ideal content/status spacers. The original `popupBounds`
rectangle is a maximum, not the initial rectangle.

[`rex/block.go`](../rex/block.go) reads layout metadata from the captured session's
view; [`rex/geometry.go`](../rex/geometry.go) validates positive, finite rectangles
and bounded, unambiguous leaf traversal. A leaf `layout.rect` is already in
**window-relative** fractions, including floating layers: do not multiply it by
layer bounds, assume the origin fills its window, or use active-window focus.
The documented zoomed-block ID overrides that block's normal layout with the
whole-window rectangle; hidden nonzoomed blocks do not supply visible geometry.
Missing/malformed layout disables geometry only, not resource-only native display.
`Block.Geometry` rechecks snapshot placement around terminal measurement.

A usable origin measurement requires the existing `TerminalSize` current-owner,
matching applied/desired rows/columns and timestamp provenance. Dividing the
origin rectangle by those cells estimates normalized units per cell. Estimated
popup columns reserve two cells for native horizontal insets; natural rows are
capped to estimated available rows, reserving four inset cells and half a row of
rounding vertically. The extra two vertical cells conservatively protect root
spacers when the new pane differs, without a later correction. Height is those
rows plus that cell-level headroom, clamped to the position maximum. Origins and
new panes can have different fonts/chrome; this is not a pixel-accurate claim.
Optional valid paired cell pixels are startup
metadata, not a font preference or a new viewer measurement.

For unavailable/unmeasured geometry, a hypothetical 100×40 window-cell grid and
30%-height startup cap (or the smaller position cap) supply a compact fallback.
No unowned default is treated as a viewer report. Cell counts are bounded to
1…65,535 before integer conversion; oversized menus paginate. Top/bottom edges
and center/side vertical center retain the original position anchors.

`Block.NewLayerWithSize` preserves the caller-compatible `NewLayer` wrapper and
adds documented shell `options.size:{columns,rows,cell_width_px?,cell_height_px?}`
in the same creation RPC as the estimated bounds. Rows/columns are seeded from
that calculation, not unrelated 80×24 defaults. This method rechecks the origin
snapshot before creation and refuses changed placement, without a creation retry.
A move after final inspection can still race; no atomic placement transaction is
claimed. There is no hidden/offscreen layer, delayed show/recreate, cross-launch
calibration, app lookup, font/settings edit or new mapping schema.

## Fixed navigation bounds and liveness

REQ-029 retains the one precomputed initial rectangle for the invocation. Neither
native nor configured popup navigation, paging, back, focus or idle changes layer
bounds. Larger menus use `RenderMenu`'s existing pagination in current cells, not
an all-submenu maximum. This supersedes the navigation-driven grow/shrink portion
of REQ-025, not initial geometry/seed or status/content spacers.

Both popup routes call public `RunTerminal` with a local top-positioned config:
the layer already carries the requested placement, so content starts at row zero
instead of being centered twice. Standalone callers retain their positioning and
API. SIGWINCH and focus-in still redraw retained menu/page/confirmation normally.
An external native window/font change can naturally alter terminal cells; content
adapts without calling `Layer.Resize` or publishing geometry. There is no private
terminal observer, preferred-row publisher, size-hint file or parent sizer.
Retired `popup_size.go` is preserved locally in ignored `.trash`, not compiled.
Generic public Rex bounds/resize/terminal geometry APIs remain available; popup
production uses origin geometry only before creation. A terminal startup seed
never acquires viewer-measurement provenance.

The existing one-second watcher now checks only captured resource liveness and,
for configured mode, the captured app generation. Each check is bounded by three
seconds and joined before cleanup; it does not sample popup geometry or mutate
bounds. Native display remains app-independent, parent/FIFO-bound and free of an
idle deadline. Configured original-VM binding and authenticated 30-second expiry/
five-second teardown reserve remain unchanged.

Cleanup unlinks exactly `config.json`, `result.json` and native `lifetime` when
present through the private root, then removes the empty directory. It does not
follow child symlinks, recursively delete unexpected contents or scavenge by label.
File/cleanup failure prevents selected dispatch. Strict retained selection,
confirmation, exact targeting and cleanup-before-action remain unchanged; no
mapping schema, public flag, daemon or renewed authority is added.

Historical measured growth/shrinkage is not fixed-navigation evidence. Fresh
frozen review, coordinator compilation and exact first/root/submenu/paging/back/
idle bounds, status spacers and Escape exit0/owned cleanup are required for
REQ-029. Leaves perform source/format/whitespace checks only.

Native Rex pane/tab/window controls are **not hidden**. REQ-026 remains blocked:
no documented per-popup control option was found in retained creation/resize/
terminal schemas, app actions or Lua layout docs. No global workaround or user
deferral is accepted. See [appearance](appearance.md) and the
[delivery inventory](acceptance.md) for evidence and remaining limits.

## Go-owned configured launch origin envelope

`launch.go` owns a small fixed `launchOrigin` DTO, separate from mapping options
and the authenticated completion protocol. Go validates each field before strict
`encoding/json` serialization; a single bounded object is emitted, never unknown,
missing, null or duplicate fields. Its five string fields are:

| Field | Meaning and validation |
| --- | --- |
| `server` | Already canonical concrete endpoint; no ambient defaulting in the envelope |
| `session_id` | Captured full typed session ID |
| `block_id` | Captured full typed block ID, uniquely placed in that session |
| `client_id` | Exact full typed attached client ID |
| `connected_at` | Captured nonlost connection generation, compared byte-for-byte, never renewed |

The envelope is at most 64 KiB, encoded as the string argument `origin` inside the
outer JSON argument object. Only stdin carries it. The caller makes one installed
configuration-action activation using an argument array equivalent to:

```text
rex --server CANONICAL_ENDPOINT --client FULL_CLIENT --autostart=false --timeout 10s do ACTION --session FULL_SESSION --block FULL_BLOCK --args -
```

There is no shell, `do -e`, Lua file, second VM, fake/token-only command, daemon or
uncertain retry. Preflight/activation share a 20-second context; individual
preflight CLI calls are bounded to three seconds and activation to ten seconds,
with earlier cancellation/deadlines taking precedence. Activation stdout/stderr
share a discarded 64 KiB total output budget and a 250 ms pipe shutdown allowance.
Errors are fixed safe text; top-level diagnostics quote errors against terminal
escape injection. No private completion/capability data is exposed.

`menu.register_launch` declares `origin` as an optional string, rejects other
arguments, and passes the **actual ctx unchanged** to `menu.launch`. Lua consumes
only the five documented fields, rejecting missing/unknown fields, wrong types,
malformed full IDs, blank generation and oversized envelope data. It compares
session/block/client with actual ctx and checks live `connected_at` **before**
issuing/evaluating a reservation. Opaque `ctx.server` is retained by identity,
never stringified, serialized or replaced by `server`.

Across yielding issue/evaluation, connection inspection, preparation and bind
calls, the launcher rechecks actual IDs/opaque server identity, connection and
configuration generation and original finite lifetime. Before controller start,
Go's canonical preparation `server` must exactly equal the expected envelope
endpoint. A detected mismatch cancels any newly issued reservation without
starting a controller or dispatching/focusing a selected action. The action also
rechecks after the yielding controller-start call and cancels on failure. As in
the retained bridge, no API provides an atomic check-and-start/focus/dispatch
transaction: external disconnects/moves can race a final check. The controller
therefore performs its own exact-origin checks before mutation and completion
still requires a live authenticated retained reservation. No failed reservation
is reopened; no selected action is retried.

This envelope is transport context, **not** executable Lua, a capability,
authority to fabricate ctx, or a second mapping schema. Other callers supplying
it cannot change the actual action activation or bypass retained registry checks.

## Configured production route (source trace)

1. [`cmd/main.go`](../cmd/main.go) parses no arguments/public flags through
   `runLaunch`, then calls `LaunchPopup`, not private-mode usage or a JSON demo.
2. [`launch.go`](../launch.go) resolves/fixes endpoint, validates exact session/block
   placement and reads server registration metadata without GUI lookup. An absent
   implicit default selects the native branch above; otherwise it binds/rechecks
   the exact app connection and invokes the installed named CONFIG action.
3. [`lua/rex_whichkey.lua`](../lua/rex_whichkey.lua) registers that action in the
   original VM; its real `run(ctx,args)` calls `menu.launch(ctx, optional_origin)`.
   [`config.lua`](../lua/rex_whichkey/config.lua) setup previously resolves
   defaults/import/inline options transactionally and compiles the raw mappings.
   [`defaults/init.lua`](../lua/rex_whichkey/defaults/init.lua) supplies the shipped
   menu; no new hard-coded demo or mapping schema is involved.
4. [`dispatch.lua`](../lua/rex_whichkey/dispatch.lua) `issue` reserves and evaluates
   the retained mapping snapshot in the original VM. The launcher encodes actual
   resolved configuration, captures private Go `prepare` output, validates endpoint
   and origin again, binds the secret capability, then starts the detached finite
   controller with a private stdin handoff.
5. [`popup.go`](../popup.go) `RunPopupController` verifies the frozen origin before
   mutation, creates the owned popup terminal layer and invokes `RunPopupTerminal`.
   [`terminal.go`](../terminal.go) `RunTerminal` uses real `RenderMenu` redraws and
   `HandleKey` navigation/selection; terminal cleanup succeeds before result commit.
6. The independent controller closes only its owned popup, revalidates/focuses the
   exact origin, removes private transport files, and sends authenticated fresh
   completion to the installed action.
   The original VM checks one-shot capability, connection/configuration generation,
   snapshot, expiry, confirmation and focus; original closures remain there and
   named/client actions use the actual completion activation. No popup-origin or
   fabricated-context dispatch is substituted.

## Execution-host routing

`rex.CanonicalEndpoint` in [`rex/endpoint.go`](../rex/endpoint.go) is the
popup's single Go-owned transport validator. `rex.ResolveEndpoint(ctx, explicit)`
selects one endpoint in this order:

1. A present explicit override, including an explicitly empty value (which fails
   validation). No invalid override falls through to another source.
2. A nonempty execution-host `REX_SERVER`, validated by the same rules. A malformed
   environment value fails rather than trying the platform default.
3. On Linux or macOS only, the installed matching Rex CLI's native default:
   `rex --autostart=false --timeout 2s whoami --json`. No `--server` is supplied.
   The command is an argv invocation, not a shell command. Its context is bounded
   to two seconds, with a 250 ms process-pipe shutdown allowance and a 64 KiB
   stdout bound. Child stdin is closed and stderr is discarded.

An empty `REX_SERVER` permits native discovery, unlike an explicitly empty
constructor/public override. A nonempty malformed environment value never falls
through. For custom/shared deployments, set matching concrete overrides or
execution-host environments on both public caller and server-side preparation.

The native branch removes inherited `REX_SERVER`, `REX_CLIENT`, `REX_SESSION`,
`REX_WINDOW` and `REX_BLOCK` selectors from the child environment. It leaves Rex's
native home, runtime, authentication and XDG environment intact. The platform
socket-directory/filename algorithms belong to Rex, not rex-whichkey: there is
no copied path table, viewing-client filesystem lookup, socket scan, explicit
directory creation or server autostart in the resolver. This branch is an ordinary
read-only identity query during product startup, not a second service. It requires
a reachable default server and valid Rex runtime/home context on that host.

The identity response must be one UTF-8 JSON object with a top-level string
`endpoint`. Other identity fields are ignored; duplicate top-level fields,
invalid/missing endpoints, oversized output and CLI failures are rejected.
Errors are sanitized and actionable, never raw identity JSON or child diagnostics.
There is no retry, alternate socket or alternate server after failure. Other
platforms require an explicit endpoint or `REX_SERVER` rather than automatic
native discovery; the product's existing platform build constraints still apply.

## Concrete endpoints, not identities

Accepted forms are an absolute Unix socket path (normalized by adding `unix://`),
a strict absolute `unix:///...` URI, or `host:port`, including bracketed IPv6.
Ports must be decimal, nonzero and fit 16 bits. Unix URIs must have an empty host,
a non-root absolute path, no user/query/fragment and no encoded-path ambiguity.
Invalid UTF-8, whitespace, controls, mutable Rex host labels and missing addresses
are rejected. Socket paths retain their spelling; no filesystem or symlink lookup
is used to rewrite them. This validation retains the previous popup URI/address
checks and adds absolute-path input normalization. Transport authentication is
still Rex's responsibility.

After resolution, consumers pass the concrete endpoint explicitly. The controller
calls `CanonicalEndpoint` directly, **not** `ResolveEndpoint`: an absent controller
server is always invalid and can never acquire a default. Generic `rex.New` keeps
its existing broader behavior; the popup supplies its validated server explicitly.
A routable endpoint is not the opaque Lua `ctx.server`. That value stays unchanged
in the original action context and remains subject to the dispatcher's original-VM
identity checks.

## Captured preparation API

The existing Go `PreparePopup(input, output)` remains a wrapper for callers without
an override. `PreparePopupWithServer(ctx, input, output, server *string)` adds
cancellation and a present/absent override. The private CLI forms are:

- `rex-whichkey prepare`
- `rex-whichkey prepare --server ENDPOINT`

Both read the actual resolved Go configuration on stdin, validate it through
`DecodeConfig`, require the server-host Rex executable, resolve the endpoint, and
only then generate an invocation capability. Captured stdout is one JSON object:

| Field | Meaning |
| --- | --- |
| `capability` | Go-generated 64-character lowercase hexadecimal single-use invocation secret |
| `executable` | Absolute path of this rex-whichkey executable on the execution host |
| `server` | Concrete canonical endpoint selected by the policy above |

This response is private plumbing, not ordinary user output. A terminal output
file is refused; the launcher must capture stdout. Capability/configuration and
controller/completion data stay off argv and logs. Discovery errors cannot include
private handoff data. `REX_BIN_DIR`, if supplied to popup preparation, must be an
absolute execution-host directory containing the matching Rex CLI; otherwise the
installed helper uses that host's PATH. No client-local executable is discovered.

## Actual Lua consumer and retained safety

`popup.new(defaults)` is valid. `popup.new(defaults, endpoint)` retains an optional
immutable constructor-only override outside mapping options. `menu.setup` cannot
replace it. Lua checks only the optional argument's type; Go owns transport syntax
and precedence. The real `menu.launch(ctx)` preparation argv includes `--server`
only when that constructor argument was supplied. An empty supplied string is not
silently treated as absent.

Lua validates the preparation response's capability/executable/server shape and
forwards **Go's returned `server`**, not the constructor spelling or `ctx.server`,
in the existing detached controller's private stdin handoff. It still reserves a
ticket before evaluating mappings, binds after yielding preparation through the
registry's generation/connection rechecks, and cancels the ticket on launch failure.
Preparation failures, including invalid response data, reach the caller as fixed
safe guidance: check the execution-host rex-whichkey helper, matching Rex CLI,
running server/runtime environment and endpoint override or `REX_SERVER`. Lua
still discards child stderr and never forwards child stdout or arbitrary error
objects to diagnostics. Later launch failures retain the generic cancellation
message; neither failure path retries or selects a fallback server.
The original `ctx`, opaque server identity, session/block/client identities and
connection generation are not synthesized from routing data.

The controller still checks exact placement and connection generation through the
explicit endpoint before arming mutations or cleanup. One-shot capability, expiry,
configuration generation, retained snapshot, confirmation, focus revalidation,
original-VM callbacks, owned-only bounded cleanup and no uncertain action retries
are unchanged. This does not install a daemon, copy callbacks into a new Lua VM,
select a focused/default client, or edit/reload the user's `init.lua`. The helper
and Rex CLI remain prerequisites for display. App eligibility and Remote Control gate
selected native execution or configured activation, not native display/Escape.
Original-server Lua modules and explicit user-owned action registration are
additional prerequisites only for the configured route, not the basic menu.

## Source evidence and verification limits

### Contract inputs and historical observations

Native startup uses coordinator-captured `.mtn-tmp/live-launch/actions.json`
(server actions), `client-actions.json` (native client names and `.args` JSON
schemas), `new-layer-api.json` (`session.new_layer` input/output/layout contract)
and `block-types.json` (the terminal shell creator's command, shell and exit
options) in the authoritative checkout. The installed Rex
`0.0.1792609031-gc465ec` read-only `actions --help` confirms server versus selected
client registry semantics; `do --help` confirms configured/native name routing,
explicit session/block/client selectors and JSON stdin arguments. No new live
API request or product run was made by the builder. Captured normal invocation
failed for absent custom actions; renderer-only diagnostic display/navigation/
dismissal is not evidence of successful normal floating-popup startup. The
[07-05 delivery record](../.tasks/tasks/rex-whichkey-07-05-native-default-popup.md)
subsequently records an independent source review, compilation and a real native
popup launch/Tabs navigation/Escape/exit0/owned-cleanup pass at `9f9b903`. Its runtime
artifact is `.mtn-tmp/live-launch/rex-whichkey-p07-t05-builder-a01/result.json`.
That observation predates Phase08 lifetime/default/color changes; it is not their
runtime result.

The native attachment classification is documented in the retained Rex
`session.list_clients` method schema at
`.mtn-tmp/remaining-delivery/rex-whichkey-p03-t04-builder-a02/session-list-clients-schema.json`
in the authoritative delivery checkout. Its output has `clients[].client_id`,
`connected_at`, `connection_state` and `info.kind`; kind is documented as `app`,
`cli`, `agent` or `other`, with unknown kinds treated as other. The schema's
example explicitly uses `info.kind = "app"`. Its validation notes limit results
to that session's `session.attach`/`session.bootstrap` attachments (not
`session.attach_all`). Reading this already-retained schema is source evidence,
not a new live API request. This same-response classification avoids a second
global `client.list` call and a cross-response membership intersection.

Installed Rex `0.0.1792609031-gc465ec` read-only `client ls --help` confirms that
ordinary text hides command-line clients while JSON includes the full server
reply. Counting all JSON attachments therefore does not implement graphical-app
selection. That original selection correction did not inspect user attachment
counts. The later captured native-startup evidence above does record one attached
app, without retroactively making the earlier source correction a runtime pass.
Historical review and compilation below do not verify later revisions.

The published [CLI contract](https://www.superlogical.com/rex/docs/automate/cli)
documents execution-host `REX_SERVER`, `REX_BIN_DIR`, explicit server selectors and
machine-readable JSON. The published [shared-server instructions](https://www.superlogical.com/rex/docs/linux/shared-server)
use absolute socket paths with `-S`. Installed Rex
`0.0.1792609031-gc465ec` read-only `whoami --help` documents identity/server/location
JSON and `--autostart`/`--timeout`. The coordinator-retained identity fixture has a
top-level concrete `endpoint`; this consumer requires that shape and rejects
unsupported shapes rather than guessing field alternatives. These are source/help
contracts, not a claim that every CLI version has been exercised.

No exact undocumented Linux/macOS path formula is asserted here. The earlier
explicit-only endpoint decision REXWHICHKEY-20261101-015342 was refined by accepted
REXWHICHKEY-20261101-101502 under delegated authority to meet REQ-017. Native CLI
discovery replaces the draft duplicated-path proposal REXWHICHKEY-20261101-100106
(Superseded); explicit overrides, opaque identity separation and fail-closed
origin safeguards remain. This document reflects that choice, not a new approval
or rewrite of historical evidence.

Endpoint task 07-01 and launch task 07-02 have recorded independent mixed Go/Lua
APPROVED source reviews and coordinator `go build ./...`/Git integration evidence
at `d1164e7` and `69078fc`, respectively; see the [delivery inventory](acceptance.md)
for exact identities and artifact refs. The six-file usage documentation received
independent source-consistency review and was integrated at `1af28738`; the
inventory also records coordinator Go compilation and Nix package builds of the
delivered source. These are supplied revision-bound results, not commands run by
this documentation builder or approval of later corrections. No unrecorded
`task build` result or runtime acceptance is inferred.

### Popup-polish evidence boundary

The authoritative [lifetime](../.tasks/tasks/rex-whichkey-08-01-lifetime.md),
[defaults](../.tasks/tasks/rex-whichkey-08-02-defaults.md) and
[theme](../.tasks/tasks/rex-whichkey-08-03-theme.md) delivery records bind fresh
source reviews and coordinator compilation to `ca03b46`, `3ad78d6` and `3e9a1fd`,
respectively. [Appearance](appearance.md) traces actual semantic render roles,
immediate ANSI fallback, palette updates and NO_COLOR/ASCII behavior.

The [assembly record](../.tasks/tasks/rex-whichkey-08-04-docs.md) and
[phase record](../.tasks/phases/rex-whichkey-08-popup-polish.md) are the durable
pointers for the reviewed frozen combined candidate, actual compilation and the
coordinator's user-authorized Linux runtime gate: normal popup, at least 35 seconds
without input, still-alive UI, absent default Focus/host groups, VT/HTML color spans,
Tabs navigation and Escape dismissal with exit0/exact owned cleanup. See the
[acceptance inventory](acceptance.md) for artifact locations and interpretation.
Only recorded observations count as runtime evidence, not this gate description.

Documentation builders/reviewers use static source/whitespace inspection only;
they do not run compilation, project tests/TDD, experiments, product/API/GUI/remote
probes or configuration reloads. The targeted coordinator gate executes no native
menu action or configured callback and certifies neither Darwin nor remote use.
Private preparation is not the normal entrypoint.

### Phase09 tab-return evidence boundary

Accepted **REXWHICHKEY-20261102-102253** adds REQ-021/022 in the
[tab-return design](../.tasks/designs/rex-whichkey-tab-return.md). Coordinator captures
under `.mtn-tmp/tab-recovery/reproduction/` at the pre-change source retained the
same native process, layer and terminal menu content through both server window
focus and app tab-goto round-trips. They do not establish what the app actually
painted, nor a crash/attachment-loss cause. The captured `session.new_layer` schema
has no persistence option; Phase09 did not invent one or change client selection
in response to a separate several-app diagnostic. Phase10's separately approved
binding correction below refines that native display prerequisite.

This source adds the focus-in repaint path and removes default Zoom in both real
generators. It does not weaken exact origin/app-generation authorization before
selection or alter configured original-VM expiry. Fresh mixed Go/Lua/docs review,
coordinator compilation and the bounded root/submenu tab-away/back, focus,
navigation/Escape/owned-cleanup gate must bind to the new frozen revision.
Earlier Phase07/08 gates are not Phase09 passes; terminal captures alone are not
native GUI visual certification. No builder tests, compilation or product/API
experiments were run for Phase09. See the [inventory](acceptance.md) for lineage
and remaining uncertainty.

### Phase11 compact-presentation evidence boundary

Accepted **REXWHICHKEY-20261102-123334** under the
[compact-presentation design](../.tasks/designs/rex-whichkey-compact-presentation.md)
adds REQ-024/025 and retains REQ-026 as blocked. The
[acceptance inventory](acceptance.md#compact-popup-presentation-req-024025-and-blocked-req-026)
indexes producer source reviews, combined Go/Taskfile compilation and supplied
coordinator measurements at `3c68b784`: actual root/Tabs/back normalized layer
bounds, current-owner terminal rows/cell pixels, settling/stability, VT/HTML status
frames, Escape exit0 and exact owned cleanup. This is not merely cropped text,
not a runtime result produced by documentation leaves, and not proof of native
controls hiding or general font/viewport/platform behavior. Task 11-04/Phase11
own revision-bound final assembly evidence; no later gate is preclaimed here.
Phase13 decision **REXWHICHKEY-20261102-143748** supersedes navigation-driven
resizing only: these historical grow/shrink measurements cannot accept REQ-029.
The current fixed-bounds contract above retains initial sizing and spacers.

Retained `.mtn-tmp/popup-presentation/{new-layer-api.json,resize-layer-api.json,
terminal-creator.json,client-actions.json,lua-reference.md}` in the authoritative
checkout establish the inspected capability boundary, not a new live API probe.
Neither server frames nor this source description certify remote native GUI
chrome. A supported per-popup Rex capability/update or explicit user deferral is
needed for REQ-026; the phase/project cannot be fully complete while unresolved.

### Phase10 native display/action boundary

Accepted **REXWHICHKEY-20261102-111430** and REQ-023 refine native launch-time
app requirements, not action safety or configured original-VM binding. The
[design](../.tasks/designs/rex-whichkey-app-independent-popup.md) and
[task](../.tasks/tasks/rex-whichkey-10-01-app-independent-popup.md) index the exact
source review, coordinator compilation and bounded runtime evidence. The native
request carries no live client/generation; its optional selector stays in a
trusted Go closure. `popupNative` is selected only by that Go route, never a
handoff field or missing configured expiry. Decoded configured handoffs still
require live origin identities, original authenticated expiry and completion.

Source inspection/formatting is not compilation or runtime proof. Earlier
single-app gates do not verify zero/several/unknown-app display or deferred target
refusal. The authorized coordinator gate covers normal startup/navigation/Escape,
valid-but-nonexistent selector startup/Escape, and one non-destructive selection
with that nonexistent target refusing after cleanup with no action dispatch.
Actual app counts and terminal/server frames must be distinguished from native
GUI presentation; no app is disconnected to fabricate cardinality. No builder
tests, compilation, product/API/GUI/config/remote experiments or actual native
selected-action execution are claimed.
