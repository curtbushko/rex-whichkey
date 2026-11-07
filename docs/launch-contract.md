# CLI-only popup launch contract

## Public entrypoints

`rex-whichkey [--config PATH] [--server ENDPOINT] [--client FULL_ID]
[--ascii] [--no-color]` opens the Go-owned popup.
`--help`/`-h` prints help without effects.
The documented [function keybinding](../examples/rex-whichkey-keybinding.lua)
starts this same installed binary with no arguments.

Production Lua exports only `launch(ctx)`. It calls documented
`rex.exec.start({ "rex-whichkey" }, { detach=true, env=... })`, raises an
immediate start error and returns true when started. It does not wait, set a
command timeout, query metadata/session state, call Rex operations or register
actions/completion. A handle acknowledges process start, **not** menu display,
valid origin, successful later dispatch or cleanup.

Rex's documented environment semantics are overlay-by-default:
`REX_SESSION=ctx.session_id or false`,
`REX_BLOCK=ctx.block_id or false`,
`REX_CLIENT=ctx.client_id or false`.
`false` removes an inherited variable. Thus an unknown context field cannot
nominate stale session/block/CLI/GUI identity. Other execution-host environment
settings remain available to Go. Lua never stringifies `ctx.server` or resolves
an executable/store/helper path; Rex finds the ordinary binary through its PATH.

References: Rex [bind](https://www.superlogical.com/rex/docs/reference/lua#rexbindkey-fn),
[exec options](https://www.superlogical.com/rex/docs/reference/lua#rexexecargv-opts),
[start/lifetime](https://www.superlogical.com/rex/docs/reference/lua#rexexecstartargv-opts).
These are published API contracts, not a physical-key observation.

## Go startup and retained authority

[cmd/main.go](../cmd/main.go) → [LaunchPopup](../launch.go) →
[launchNativePopup](../native.go) is the one public resource path.

1. Load/validate version-1 [data configuration](mapping-format.md), canonicalize
   keys, and reject invalid data before process/resource startup.
2. Require full origin session/block IDs; capture the actual running public image.
3. Select `rex` through ordinary execution-host PATH lookup and independently
   retain a canonical endpoint; inspect exact origin Session/Block/Window/Layer
   placement. No binary receipts, startup version/API/schema negotiation or
   server.status/config-generation wrappers run.
4. Capture explicit app intent or the pane owner if eligible. An explicit invalid
   or unavailable app rejects; absent implicit GUI eligibility does not gate
   resource-only display.
5. Expand enabled Go providers once and retain targets; retain the required origin
   session-view snapshot without extra rename-label reads. Calculate one initial
   rectangle and create the actual popup terminal.
6. Wait for ordinary interaction/dismissal/cancellation without idle expiry.
   Restore terminal state, validate the result against the retained snapshot,
   finish exact cleanup, then dispatch a selected effect once.

Neither old Lua registrations nor an installed old bridge select, redirect or
gate startup. There is no custom `rex do whichkey` route, runtime advertisement,
original-VM registry, callback/completion ticket, preparation relay or second Lua VM.
The retained `menu` command is data-only **UI/result**, not popup/resource
dispatch. `popup --directory ... --native` is a private Go PTY child, not user
wiring. Removed prepare/controller/runtime-relay modes are not prerequisites.

## Exact origin and GUI intent

Origin is full `REX_SESSION/REX_BLOCK`, never a focused-window search or
`REX_WINDOW` shortcut. Go captures unique original Window/Layer placement and
rechecks it; moved or detached resources fail rather than retarget.
The endpoint is execution-host intent; [discovery](runtime-bootstrap.md) documents
the independent CLI/endpoint policies.

Present `--client` wins over present `REX_CLIENT`; either must be a full
`client:UUID`, including rejecting an empty value. Without explicit intent,
app routing uses the original pane's validated owner: attached live `app` kind,
exact ID and connection generation. No app chooser, singleton/list-order/current
focus/label fallback exists. A CLI caller is not automatically a graphical app.
See [CLI/GUI delegation](cli-gui-delegation.md).

Display, origin-focus restoration, resource rename/close and generated pane focus
do not require an implicit GUI app. Explicit app intent remains strict even for
resource actions. App-routed selections require the retained eligible app, fresh
connection/placement/owner and built-in `kind=client` action metadata. Calls use
explicit `client.perform`, never custom Lua invocation; Remote Control and
authentication remain Rex policy.

## Actions and providers

The closed [catalogue](mapping-format.md#complete-public-action-catalogue) is
authoritative; arbitrary configured executable/action names are not supported.
Go expands sessions/windows/panes/known-hosts before sizing. Generated full IDs
and action intents remain in the parent; the PTY returns canonical keys,
confirmation and authorized rename text, never replacement target authority.

Session/window selection revalidates destination membership and sends exact
`session.select` args without outer origin session/block overrides. Split also
omits those outer fields. Other app actions retain captured origin context.
New Tab performs exactly one source `Session.NewWindow`, then activates its
returned ID on the retained app. Activation failure preserves the created tab
and reports its ID; uncertain creation/activation is not retried or rolled back.
Rename and destructive close use exact typed resource methods after cleanup.
[Actions](actions.md), [rename](tab-rename-input.md) and
[session UI](session-owned-ui.md) describe the retained semantics.

## Private incremental rename contract (STARTUP-UI-UPDATES)

This bounded contract is specified before its consumers. It is private Go
invocation plumbing, not a public config/protocol/service extension. Inventories
still expand eagerly in this intermediate revision; this is not final async
first-frame startup.

- The parent creates a random invocation ID in `runtime.json`. Four fixed regular
  slots (`rename-tab-demand.json`, `rename-session-demand.json`,
  `rename-tab-update.json`, `rename-session-update.json`) start empty. Each slot
  carries at most one strict JSON object, bounded by `popupMessageLimit`. Owned
  0600/no-follow/link-count checks and nonblocking flock serialize full reads and
  writes; contention is retried by the existing polling loops, never by waiting
  on the terminal thread. No append log, arbitrary filenames or unbounded queue.
- Opening an inline form demands canonical kind `tab` or `session`, full canonical
  mapping keys as its prefix, invocation ID and load number 1. At most two demands
  and one metadata worker exist. The worker reads a detached origin session-view
  snapshot already required for placement, or obtains that exact view on demand
  when none is supplied. Its immutable result contains both labels or a bounded
  unavailable reason; it cannot mutate Config, targets, live Menu or terminal.
- The controller alone validates demands against enabled retained empty-args
  rename mappings. It retains each kind/load/prefix result **before** writing a
  ready update. Existing mappings/args/full target IDs and original GUI intent
  were retained before child creation and never come from presentation. A ready
  kind covers only the parent's already retained same-kind rename mappings;
  alternate configured keys share that publication rather than start new loads.
- UI state is loading, unavailable or ready per kind. Loading/unavailable status
  is inline and non-executable. Forms accept edits and Escape while loading, but
  Enter is discarded until ready: it is never buffered, confirmed or replayed.
  Only the UI loop reads updates, changes live state and redraws. Unrelated
  prefix/page/confirmation and partial UTF-8/escape/paste decoder state survive.
  A label initializes only a wholly untouched form, not edits, cursor movement,
  validation/error state, pending input bytes/paste or confirmed text.
- A rename result adds only `update` (invocation/kind/load) to bounded
  keys/confirmation/text. The parent joins the controller, then matches this
  binding to its retained ready publication/demand prefix and selected mapping
  kind. Demand keys are not repeated in the result, preserving the existing
  near-limit full-key result budget. Unknown, malformed, stale, loading or
  unavailable bindings refuse; non-rename results
  cannot carry one. Labels, ordering, child actions/args/IDs grant no authority.
  No refresh/retry generations exist in this revision; each kind publishes once.
- Child terminal loss/cancel/parent-pipe EOF stops interaction; its lifetime
  watcher is canceled/joined before result commit. Parent selection/dismissal,
  error, origin loss or cancellation cancels and joins its metadata worker and
  controller before result acceptance and exact popup/focus/transport cleanup.
  Late worker results are discarded after cancellation. Transport/join/close/
  cleanup errors suppress effects, not hidden as successful selection. Existing
  fresh origin/app-generation/destination checks and once-only dispatch remain.
  No idle menu deadline is introduced.

## Fixed navigation bounds and liveness

Go computes content-sized initial bounds and seeds PTY cells in a single creation
call. Submenu/back/paging/focus/idle never resize or replace the layer. Large menus
paginate inside current cells. External font/window changes may change terminal
cells; SIGWINCH/focus-in redraw do not resize application-owned bounds.

There is no Lua action timeout, callback-ticket expiry or menu idle deadline.
The Go parent watches captured resource placement/liveness; the PTY child watches
its parent lifetime pipe. Long terminal wait does not inherit a short RPC timeout.
Tab switching or absent GUI eligibility alone does not cancel display.
Operation checks/RPCs and cancellation-independent cleanup remain bounded.
See [appearance](appearance.md) for measurement/focus/native-controls limits.

## Fixed owned cleanup inventory

[popup.go](../popup.go) owns `runPopupLifecycle/removePopupFiles`;
[runtime_bootstrap.go](../runtime_bootstrap.go) owns source-handle capture/copy/retirement.
Cleanup targets only:

| Owned item | Boundary |
| --- | --- |
| Returned popup terminal Block and Layer | Inspect exact Block; close through the retained Layer; no cascading window/session or label cleanup |
| Original focus | Revalidate exact origin; never select a replacement focus target |
| Parent lifetime writer, child reader/watcher, wait goroutine | Close/join; no surviving helper daemon |
| `config.json`, `result.json` | Private resolved snapshot/result, not user files |
| `native-input.json` | Private inline-form enablement; legacy label presentation remains private, not identity |
| `rename-tab-demand.json`, `rename-session-demand.json` | Bounded single-publication form demands |
| `rename-tab-update.json`, `rename-session-update.json` | Bounded retained presentation publications |
| `lifetime` | Invocation-local FIFO |
| `runtime-image`, `runtime.json` | Invocation-local image copy and selected CLI path/independent endpoint context, not an executable receipt or Lua bridge |
| Private `rex-whichkey-*` directory and open source-image handle | Remove fixed names through the owned root, then empty directory; retire source handle before effect |

Directories are owned mode 0700; transport slots are owned mode 0600, the copied
image 0700. Strict JSON, bounded copy/input, fixed-slot owner/mode/link-count and
no-follow checks preserve private transport integrity. The child reuses the
selected CLI path without rediscovery; no binary hashing or installation/link
receipt is required. Copy/write/close completes before start without image Sync.
The child result cannot supply replacement action/resource authority. Result
Sync, locking, lifetime checks and exact source-handle retirement remain.
Unexpected directory contents prevent removal; there is no recursive scavenger.

The terminal restores modes/termios before result commit. Exact popup removal,
origin focus, pipe/files and source-handle retirement must succeed before selected
effects, including actions that close their caller. Uncertain cleanup suppresses
dispatch. Deferred cleanup survives caller cancellation with its own five-second
budget. Unknown creation IDs, server loss, SIGKILL or I/O failures can leave
resources/artifacts. No broad cleanup, automatic uncertain retry or guaranteed
post-crash cleanup is claimed. Checks and effects are separate, not atomic
against external changes.

## Evidence and historical boundary

Task2's coordinator actually observed default public root/Sessions/Switch menus,
Backspace/Escape, exit0, primary terminal, exact popup/block/private-directory
cleanup and origin focus on Darwin arm64 at `7bdbda2`.
See [acceptance](acceptance.md) for the exact commit and artifacts.
The final launcher/package combined candidate still needs independent review and
coordinator normal public runtime, including a benign effect only on an explicitly
owned disposable target. No leaf product invocation, physical-key/reload,
Nix build/activation, remote trial or selected effect is claimed here.

The earlier bridge contract is historical at
`cd5cbce818e07c264738ddcbc591ad49fdf0e9c8:docs/launch-contract.md`.
Its original-VM/completion/controller/anchor requirements do not apply to this
CLI-only path. Prior evidence/commits and [historical acceptance](acceptance.md#historical-delivery-records)
remain preserved; Phase15 is not retroactively declared complete.
