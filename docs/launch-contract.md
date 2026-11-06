# Popup launch, endpoint and preparation contract

## Ordinary startup and optional Lua setup

Run `rex-whichkey` in the originating Rex terminal on the execution host.
Public Go selects itself, then freezes one Rex CLI, endpoint and read-only API
evidence. Only a truly absent **implicit** default `whichkey` action selects
the basic native menu. Registered/default or explicit `--action NAME` requires
the original-VM configured route; compatibility/activation/startup failures never
become native fallback or a retry. Help (`--help`/`-h`) contacts no Rex server.

- `--server ENDPOINT`: concrete execution-host transport intent.
- `--client FULL_ID`: exact full `client:UUID`; syntax checked early, live
  binding after native selection or before configured activation.
- `--action NAME`: require a registered Lua action (lowercase letters/digits/
  underscores), even when explicitly naming `whichkey`.

Unknown flags/positionals, missing values and explicit empty action/server/client
reject, not default. Internal describe/relay/release/prepare/controller/popup
commands are finite private plumbing. Standalone `menu` still consumes resolved
JSON, not ordinary startup or an action executor.

Basic native startup needs no Lua configuration. Optional customization uses the
normal standard-file-loader namespace on the server:

```lua
local popup = require("rex_whichkey")
popup.register_completion("whichkey_complete")
local menu = popup.new(require("rex_whichkey.defaults"))
menu.setup() -- or menu.setup(inline_options, "my_whichkey")
menu.register_launch("whichkey")
```

Register completion once at load in the same VM. Launch/completion names are
unique, distinct and not reserved `whichkey_runtime`. Keys/palette and direct
`menu.launch(ctx)` use their real app context; no keybinding is installed. User
modules live outside reserved `rex_whichkey`. Setup, callbacks, args, conditions
and confirmation follow the [mapping standard](mapping-format.md).

A complete normal layout revision 1 ships `rex_whichkey/runtime/anchor.lua`, the
real application executable at its fixed sibling `rex-whichkey`, and `deps/rex`
linked to the licensed execution-host CLI. With that dependency installed, the
public `bin/rex-whichkey` symlink points into the layout, not the reverse: Go's
implicit executable-relative dependency selection and Lua's explicit anchor
reach the same installed image. Go validates canonical file/content/link identity,
ownership and modes. Home Manager exposes
only the ordinary namespace (symlinks supported), not runtime directories/per-host
selectors. The installer preserves upstream license, signature, loader/libraries
and assets. The pinned source supplies x86_64/aarch64 Linux and Apple-silicon
macOS, **not** Intel macOS. This is installation/source analysis, not a runtime
or Nix build result.

The default Intel macOS package omits `deps/rex` and keeps its real public image
at `bin/rex-whichkey`; `runtime/rex-whichkey` links back to it via
`../../../../../bin/rex-whichkey`. Both configured public and
key/palette routes therefore fail the anchored installation check before popup,
even if public Go already selected a valid CLI or `REX_BIN_DIR`/PATH is usable.
Neither PATH repair nor native fallback repairs that partial Lua layout. The
existing installation override `rex-cli` can supply a legitimately available,
licensed execution-host package exporting `bin/rex` to complete it; no upstream
Intel artifact is invented. See [installation](../README.md#install-on-the-rex-server-host).
The default packaged public image is outside the implicit layout, preserving
Intel standalone native use with an existing valid unambiguous absolute
execution-host PATH CLI when `REX_BIN_DIR` and the implicit default are absent
and no explicit action is required. No environment repair or dependency override
is needed solely for native use; origin/endpoint/CLI checks remain. A valid
terminal `REX_BIN_DIR` is still an optional preference. This is absence of an
applicable implicit layout, not fallback from a broken layout or configured
failure. Neither native case enables the incomplete configured integration.

Standard file-loader absolute anchors work; relative anchors require unchanged
server cwd. Synthetic/preloaded/custom-loader anchors, shadowed/mixed trees and
broken assets reject without path guessing. Restricted server PATH uses the
anchored helper/dependency, not terminal-only environment assumptions or required
constructor paths. Present `execution.helper`/`execution.rex_bin_dir` fails with
migration guidance; users remove obsolete wiring from their own configuration.

`popup.new(defaults, endpoint)` remains optional immutable endpoint intent outside
setup/mappings. Key launch uses it for Go selection; public launch requires
canonical agreement with its frozen endpoint. No required socket literal or
manual component pairing exists. See [runtime bootstrap](runtime-bootstrap.md).

A public local rebuild prepares/starts its own controller/PTY; the installed
helper is only a finite original-VM transport relay. Keys use their anchored
helper. An arbitrary standalone build needs terminal `REX_BIN_DIR` or valid
unambiguous PATH when no normal layout applies. Installation is not runtime
discovery or replacement of loaded closures. Install updates normally; unsupported
loaded bridges require the user's ordinary configuration lifecycle, never auto
reload/restart, repeated flake refresh or store-path copying. Compatible public/
helper/loaded-consumer build IDs may differ. User `init.lua` stays untouched.

Rex Remote Control gates CLI-origin app effects separately from authentication.
Native display/dismissal needs no app binding; configured activation does.
Configured success requires finite controller startup acknowledgment, not observed
display or callback completion. Native success follows child exit/cleanup and
optional one-shot effect acceptance, not proof of visible GUI effects.
This cutover documents source-reviewed/compiled intent, **not runtime verification**.

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

Client precedence is present `--client`, then present `REX_CLIENT`, otherwise the
launching pane's documented GUI owner. The exact ID is validated against only
the captured session's attached graphical apps; list cardinality never selects
an app.
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

Without a selector, all native app actions and configured CLI launches bind the
launching pane's exact GUI owner through the same helper. This includes session
and tab previous/next, new session/tab, splits, sidebar, palette and
Switch Session. Default rename is an exact-resource operation, not app dispatch. No app chooser, singleton/list-order fallback or per-action
exception remains. An absent/non-GUI/disconnected owner refuses rather than
choosing another app. Native resource closes/rename need no app binding unless
one was explicitly supplied. `REX_CLIENT` normally need not be exported.
A supplied selector must match a live app in that exact session; stale, non-app or
unattached selectors refuse even if another app exists, never falling back to
another client. Native failures are **action-targeting refusals after popup cleanup**,
not launch prerequisites. No global, focused, last-resize or arbitrary-first
fallback is used. The block-owner routing policy is distinct from using terminal
size reports as app identity.

`popupRequest.originBlock` combines `popupOriginConnection` generation revalidation
with `resourceOriginBlock`'s exact session and uniquely placed block check.
Configured launch/controllers use the combined check throughout. Native display
uses only resource checks; native app dispatch binds after selection and then uses
the combined check immediately before execution. Resource closes/rename use
resource checks alone unless an explicit app selector was supplied. The Rex resource client rejects
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
icons and close confirmations align with the shipped Lua default intents.
Native Switch uses Go's inventory submenu; registered Lua retains its named
switch action. Both routes use WhichKey text for default rename. Neither default set includes Zoom (`z`), the Focus group or a
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
passes canonical configuration and frozen CLI selection to `runPopupWithLifetime` in
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
focus-stealing callback. Later action targeting uses the retained explicit selector
or pane owner without another prompt or interactive stage.
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
only keys, optional opaque callback identity, explicit confirmation and permitted
rename text (or null dismissal). The strict canonical
completion decoder rejects replacement action/args and checks confirmation and
identity against the retained snapshot. The native snapshot has **no callbacks**;
Go refuses callback references rather than evaluating Lua. Capabilities stay in
the parent; transport/config/result paths contain no action credentials.

After successful child exit (which follows terminal restoration), the controller
closes only its returned owned terminal, confirms its layer vanished, restores
exact origin focus and removes fixed owned config/result, rename presentation,
private runtime receipt/image and native lifetime-pipe files and their empty directory **before**
selected dispatch or any native app lookup (see the complete
[fixed inventory](#fixed-owned-cleanup-inventory)). The parent's lifetime-pipe writer is closed in that cleanup
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

[`native_app_owner.go`](../native_app_owner.go) binds every native app action and
configured CLI launch through one pane-owner/explicit-selector path.
[`native_targets.go`](../native_targets.go) only validates eligible attachments;
it no longer constructs an app picker or reads labels. `popupOriginConnection`
requires an exact binding and cannot choose a singleton app or prompt for one.
The earlier app chooser and its private selectable action were retired.

Native app binding starts only after original popup teardown. There is no second
interactive stage and no cached app preference. The outer native signal context
spans the menu, metadata lookups, cleanup and dispatch; interaction has no idle
deadline, cleanup remains bounded, and parent loss closes the child's FIFO writer.
Configured CLI binding happens before activation; original-VM capability, expiry
and finite controller behavior remain unchanged.

After selection and cleanup, a fresh three-second dispatch context remains a
child of that outer signal context. Every app action revalidates the selection,
actual dispatch action registration, captured app generation, original tab/layer
placement and unchanged pane owner before its first effect. Explicit overrides
skip only the owner-equality condition, never app/placement validation. Lost
attachments, reconnects, ownership changes or moves refuse without retargeting.

Native app actions other than New Tab intent use documented `client.perform`
through a separate **unscoped** `rex.Client` pinned to the captured endpoint and a
three-second timeout. Each request preserves retained `client_id`, `action`, JSON
`args`, `origin: "cli"` and `timeout_ms: 3000`; empty args become `{}`. All non-split
requests also carry explicit captured origin `session_id` and `block_id`, preserving
origin-bound context. Resource closes instead use exact server IDs without app
dispatch. App-focused `pane.split`
(`v`/`right`, `s`/`down`) omits both fields when splitting the app's focused pane.
Exact session selection also omits them: its retained destination is in args,
not the origin context.
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

After the popup lifecycle and all private files/leases are gone, the parent resolves
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

Built-in native `tr` uses the accepted [text-entry prototype](tab-rename-input.md).
The parent captures the tab's ID and label before opening the menu. A private
`native-input.json` presents the initial name; the child never receives another
resource target or authority to change it. The same terminal/menu instance enters
text editing and returns to Tabs on Escape without completing the invocation.
Enter emits bounded name data for retained `tr`/`wr` rename entries only. The parent validates and strips that
private field before canonical Completion decoding, cleans the popup/files/lease,
revalidates exact resource placement, and invokes typed `Window.Rename` or
`Session.Rename` once, according to the parent-retained intent.
Mapping/Config JSON remains unchanged. Registered Lua defaults now share the
same input page: the controller reads exact session/tab labels into that private
presentation file before opening the popup. Negotiated configured completion 2
binds the exact receipt/hash and permits validated `text` only for the retained
empty-argument `tr` tab or `wr` session rename intent (4096 bytes and 128 Unicode
characters). Loaded text/resource support is required before popup; legacy
completion 1 alone does not establish that feature. The original Lua registry validates it, rechecks
conditions and origin, and calls `session.set_window_label` or `session.set_label`
with retained resource IDs after popup cleanup. Missing text cannot fall through
to a native dialog; text supplied for callbacks or other actions is rejected.
The tab is frozen at reservation; a pane moved to another tab cannot redirect
rename. Custom callbacks, custom retained arguments and replacement actions retain
their normal dispatch. Compatible public/helper/loaded-consumer build IDs may
differ; unsupported loaded bridges require the user's normal update/load
lifecycle, not manual version pairing or automatic reload.

Native `x`, `tx` and `wx` bypass app selection and native action registration.
After validated confirmation and full UI/file/lease cleanup, existing typed Rex
resource operations close the exact captured pane, tab or session. The origin's
window/layer placement is captured before the menu and revalidated before the
single close; a move refuses rather than selecting a new tab. Explicit app
selectors remain strict, with their connection checked before closing; without
one, these resource operations need no app lookup. Errors/uncertain closes never
retry. The route is compiled, not runtime-certified after the user's stop-testing
instruction.

Built-in session rename uses the same resource/form route as tab rename. Built-in
Switch Session expands a live-at-opening session-list snapshot into a canonical
`ws` submenu, with retained `session.select` leaves and full ID args. The parent
revalidates the selected destination after cleanup; it never trusts a label,
position or child-supplied destination. No Rex session picker is opened.
Unavailable inventory yields an informative empty submenu, not a guessed list.
See [session UI details](session-owned-ui.md). Other app actions retain their
semantics; Command Palette is the remaining Rex-owned dialog/picker in this menu. Explicit app selectors take
precedence. Without one,
[`native_app_owner.go`](../native_app_owner.go) reads the launching pane's documented
`session.describe_block` owner and binds that exact live graphical app attached
to the captured session, including its connection generation. This is the
user-requested pane-owner routing policy, not proof of input provenance or extra
permissions. It bypasses the session-wide app chooser. Immediately before the
one `client.perform`, Go revalidates the app generation, original tab/layer and
unchanged owner. A missing/non-GUI owner, move or changed binding refuses; no
label, resize reporter, first/last app or remembered choice substitutes. Explicit
selectors keep their existing strict validation and need not match the owner.
New Tab validates the owner before creation. Creation can change focus/ownership;
activation retains the already-bound app and rechecks its generation, origin
placement and new-window membership without resolving another owner.
The native screens above replace the old rename prompts/session picker; the app
binding itself introduces no additional UI. The route is compiled, not
runtime-tested under the user's stop-testing instruction.

Apart from resource closes/rename and the retained New Tab intent adapter, every
remaining native action sends trusted action/JSON args once through the
explicit `client.perform` request above, with captured origin session/block
context; no native `rex do` dispatch wrapper remains. Metadata still legitimately
uses the read-only Rex `actions --json` command. Destructive close targets and
configured original-VM routing do not change. All native dispatch paths share
the fresh three-second signal-cancellable context, sanitized failure or
uncertain-effect errors and no retries or fallback. Dismissal executes no
native action.
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

### Fixed owned cleanup inventory

The inventories belong to the actual cleanup functions, not a second transport
schema or permission to delete caller-nominated paths:

- [`removePopupFiles`](../popup.go) removes only `config.json`, `result.json`,
  `lifetime` (native only), `native-input.json` (rename presentation),
  `runtime-image` and `runtime.json` when present in its authenticated private
  popup root, then the empty directory. The two symbolic names come from
  [`nativeTextFile`](../native_text_entry.go) and
  [`runtimeImageName`](../runtime_bootstrap.go).
- [`removeRuntimeWorkspace`](../runtime_bootstrap.go) removes only `request.json`,
  `menu.json`, `prepared.json`, `bound.json`, `lease` and `runtime-image` when
  present in its authenticated private rendezvous/preparation root, then the
  empty directory. Key preparation need not contain the public relay slots.
- [Original Lua](../lua/rex_whichkey.lua) cancels unclaimed local registry tickets
  on failure. Only before an attempted key controller start may it ask Go's
  authenticated `runtime-release` to remove that preparation workspace. Lua
  does not recursively remove files or scavenge expired workspaces. Original-VM
  expiry revokes authority/references; it does not prove artifact removal.

Root-scoped removal does not follow child symlinks or recursively delete
unexpected contents; unexpected entries prevent empty-directory removal.
Only exact known owned popup resources are closed, never a broad layer/window/
session search or a retired prompt resource. Retire owned source/transport/startup
handles as required by their lifecycle. File/handle/close/focus uncertainty
suppresses selected dispatch; cleanup completes before effects and uncertain
closes are never repeated. A possible controller start leaves workspace ownership
with that child even after a failed/lost acknowledgment. Unauthenticated input
cannot authorize nominated cleanup and may retain artifacts. No retry, renewed
expiry or stronger cleanup guarantee is implied. Strict retained selection,
confirmation and exact targeting remain unchanged.

Historical measured growth/shrinkage is not fixed-navigation evidence. Fresh
frozen review, coordinator compilation and exact first/root/submenu/paging/back/
idle bounds, status spacers and Escape exit0/owned cleanup are required for
REQ-029. Leaves perform source/format/whitespace checks only.

Native Rex pane/tab/window controls are **not hidden**. REQ-026 remains blocked:
no documented per-popup control option was found in retained creation/resize/
terminal schemas, app actions or Lua layout docs. No global workaround or user
deferral is accepted. See [appearance](appearance.md) and the
[delivery inventory](acceptance.md) for evidence and remaining limits.

## Code-owned configured startup and finite transport

[Runtime bootstrap](runtime-bootstrap.md) and the
[Go-owned mapping/transport standard](mapping-format.md) define exact fields,
strict decoding and bounds; this is an ownership summary, not a parallel schema.
Protocol/layout are 1 and configured completion is 2. Unknown/duplicate/aliased/
null/wrong-type/bad-Unicode/trailing or oversized protocol messages reject.

### Public route

1. `LaunchPopup` captures its open self-image, `rex.CLISelection`, endpoint,
   terminal placement and API evidence. Read-only server actions choose the
   route; `whichkey_runtime` must advertise actual loaded menu/config/server
   generation and required features before activation. Missing/old bridges fail
   before calling an old launch action. Compatible build IDs need not match.
2. Go binds the explicit or pane-owner GUI app and exact `connected_at`, then
   creates 0700 invocation state, fixed 0600 no-follow one-link JSON slots and a
   FIFO lease. The bounded opaque locator binds request digest/receipt; action
   arguments contain no capability or caller-nominated executable.
3. Go activates the original registered action once with captured CLI stdin.
   Lua freezes actual opaque server/session/block/caller, separately validates
   nominated GUI and unique window/layer placement through direct APIs, then
   calls only its anchored helper's finite preflight. Go authenticates local
   owned files/manifest/digest/receipt/live lease and every origin/consumer field
   before any reply, ticket or evaluation. Unshared execution-host files reject,
   not remote path guessing.
4. After certain compatibility/preflight, Lua reserves once and evaluates in the
   original VM, retaining closures/args/guards/snapshot/table ticket. The finite
   helper relays menu data; **public Go** prepares its own selected component,
   random capability and final receipt/hash. Lua binds those exact bytes and
   publishes bound state. It never starts the public returned executable.
5. Public Go requires matching bound and certain activation success, then retires
   finite relay/FIFO/source handles before starting its selected private
   controller image. Uncertain retirement prevents startup. No activation/relay
   effect is repeated; no handle/secret/private result escapes the action.

### Key/palette route

The actual app is both caller and GUI, with no public nomination/locator.
Original Lua validates actual context and loaded generation, invokes only fixed
anchor-sibling `runtime-describe`, and checks producer compatibility/original-VM
metadata before reservation/evaluation. Private `prepare` receives resolved Config
and frozen receipt/hash, revalidating rather than rediscovering CLI/endpoint.
Lua binds final receipt/capability/resources and starts the verified returned copy
as a finite Go launcher, with captured stdin/stdout, suppressed stderr,
`detach=true` and supported `h:wait(35)`. Success requires successful exit **and**
exact accepted acknowledgment. Before attempted start only, preparation/bind
failure may release ticket and authenticated preparation. After possible start
Lua cannot release the workspace or retry. No private output/handle is returned.

### Controller/PTY barrier and irreversible ownership

Both paths use finite private input through EOF. Child closes stdin, authenticates
receipt/copy/self-image/frozen CLI/endpoint, checks loaded consumer/origin and
rename presentation, then rechecks menu/origin/expiry. It writes accepted/rejected
plus EOF and closes that writer before effects. Parent requires the complete
bounded acknowledgment/EOF, closes duplicates/reader and releases the process;
key launcher also writes/closes accepted stdout at this barrier. Only then is a
positive proceed plus EOF sent on a separate finite pipe. Child checks/closes
that reader before lifecycle entry; EOF alone cancels. No startup descriptor
stays live during popup effects. Three-second startup bounds never renew expiry.

`Start` transfers workspace ownership irreversibly. Failed/lost/rejected/late
acknowledgment, cancellation or close/release error after possible start does not
authorize retry, killing an uncertain child or deleting workspace. Complete
authentic expired pre-child input can authorize fixed owned cleanup only after
receipt/copy/self-image authentication; startup still fails with no stale
completion. Partial/malformed/foreign input cannot authorize nominated deletion
and may retain preparation artifacts. Expiry revokes authority, not proof of
artifact removal. Cleanup errors remain errors; uncertain closes are not repeated.

Controller/PTY use verified private copies of the selected open source image;
Go checks content/file identity before/after copying. Rex's proprietary CLI stays
at its frozen package path, never copied/modified at runtime. The PTY receives
private CLI/endpoint evidence, including for provider inventory, not environment
rediscovery. Trusted stable installer/user-owned components are assumed; no
hostile-same-account isolation or portable atomic execute-by-descriptor guarantee.

## CLI caller versus GUI target

Accepted **REXWHICHKEY-20261104-172117** separates each actual action caller from
its retained GUI. Public launch/completion CLI connections need not equal the app
or each other. Actual opaque context is frozen within each activation, never
fabricated from nomination. Key launch uses its actual app. Check full origin and
GUI generation before private preflight, after every yield, in controller watching
and before dispatch; never refresh a connection or adopt a replacement app.

Dotted client actions use explicit retained-target/context/args `client.perform`;
configuration actions use original-VM `rex.invoke`. Callbacks receive logical
`client_id` for GUI and `request_client_id` for actual completion caller, not
fabricated ambient authority. See [CLI/GUI delegation](cli-gui-delegation.md).

## Execution-host routing and concrete endpoints

The concrete `rex` package owns selection: reuse trusted frozen selection;
otherwise trusted internal explicit executable, nonempty execution-host
`REX_BIN_DIR`, applicable validated layout dependency, then unambiguous absolute
PATH. Invalid preferences/broken layouts never fall through. Empty/relative/cwd
PATH entries and distinct candidates reject. Terminal-child documentation does
not promise that configuration actions receive `REX_BIN_DIR`.

Go freezes canonical content/file/link identity and platform before endpoint
operations, with bounded read-only CLI/server/consumed-method-schema evidence.
Unknown/missing evidence is incompatible, not guessed success. CLI/server release
strings need not equal; no minimum tested release is claimed. Every later
resource/completion/PTY call retains that CLI/endpoint; change aborts, not reselects.

Endpoint precedence: present public `--server` or key constructor intent >
nonempty `REX_SERVER` > bounded no-autostart Linux/macOS native Rex discovery.
Invalid/empty explicit values reject; empty environment permits discovery. Rex
owns home/runtime/XDG socket rules. No shell/PATH wrapper, store/repository/home
guess, alternate endpoint, directory creation or server autostart is added.
Accepted forms are concrete absolute Unix socket paths/strict URIs or numeric
nonzero-port `host:port` (bracket IPv6). Percent encoding preserves Unix spaces;
opaque `ctx.server` is identity, never an address or execution-host proof.

Constructor intent stays unchanged through describe/menu/bound. Key describe uses
its precedence. Public preflight canonicalizes a present string only after private/
origin authentication and requires agreement with its retained public endpoint.
Absence is normal, not discovery. Empty/malformed/conflicting or changed presence/
spelling rejects before authority/evaluation/popup, even for a changed equivalent
spelling. It is never ignored or used to retarget. The optional shared deployment
`/run/rex/server.sock` is an example, not a default or required configuration.

## Original-VM authority, rename and completion

Go owns Config/Mapping/completion/receipt semantics. Lua retains closures, args,
guards, confirmation, immutable mapping/resource snapshots and opaque tickets.
Go receipt/relay identities are **byte-exact strings**: timestamps/device/inode
integers may exceed Lua 5.1 precision, so never decode/re-encode them. Bind exact
final receipt/hash; parsed inspection objects are not authority.

Configured expiry is never renewed: original 30 seconds, 64 pending entries,
three-second stages and five-second teardown reserve. Generation replacement,
backwards time, reconnect/detach or moved origin revokes/suppresses effects.
Native display separately has no idle deadline and retains its lifetime pipe.
Loaded compatibility/menu generation and exact origin are rechecked after rename/
geometry yields, immediately before layer creation, in the watcher and before
completion.

Default `tr`/`wr` requires negotiated text/resource-rename support before popup.
WhichKey's text page supplies a nonblank control-free UTF-8 label (4096 bytes and
128 Unicode characters), not an ID/action, to only the retained window/session
resource API after cleanup. No native dialog, dropped text or generic action
false success. Custom callbacks/args/conditions/confirmation remain ordinary
mappings. See [rename](tab-rename-input.md) and [session UI](session-owned-ui.md).

Terminal restoration and exact owned-popup/file/source/handle cleanup finish
before selected dispatch. Restore only a valid origin; unknown creation identity
or uncertain close/focus/removal suppresses effects, never broad cleanup.
Fresh completion strictly decodes/authenticates capability, final receipt/hash,
real origin, retained state and original expiry before single-use claim **before**
yielding guards/APIs/callbacks. Malformed/foreign input cannot consume unrelated
entries; claimed failure is permanently consumed. Recheck GUI/generation/
placement/guards/confirmation immediately before one retained effect. No uncertain
retry, reopened popup, refreshed authority or cancellation masking possible
completion. Inspect/focus/effect remains non-atomic against external races.

## Cutover source scenarios (all UNEXECUTED)

Complete public symlink/real-image layout and dependency selection without
`REX_BIN_DIR`; explicit Lua anchoring and optional HM exposure; default Intel
packaged native use with no environment preference, a valid execution-host PATH
CLI and absent implicit default/no explicit action; incomplete Intel configured
public/key rejection and licensed dependency override; standalone public rebuild
versus unequal compatible installed/loaded builds; restricted-PATH original-VM keys with a complete layout;
broken/mixed/synthetic modules; old/missing/text-incapable bridges; valid/equivalent/
invalid/conflicting endpoints and explicit client intent; custom callbacks/args/
conditions/confirmation; retained-resource rename; namespace symlink upgrades;
unshared remote state; expiry/reconnect/moves/replay; uncertain activation/start/
cleanup/dispatch and retained artifacts remain unexecuted source scenarios.
This packaging/docs work performs only permitted Nix parse/static/format and
whitespace diligence, not runtime/TDD acceptance, deployment or phase completion.
Historical evidence below is revision/context-bound, not combined-cutover proof.

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
