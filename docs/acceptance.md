# Release delivery inventory and limitations

## First-frame fit and status spacing: REQ-027/028

Accepted **REXWHICHKEY-20261102-141047** under the
[Phase12 design](../.tasks/designs/rex-whichkey-first-frame-fit.md) refines startup
and status spacing without waiving REQ-026. This section describes source and
required evidence, not an inferred review, compilation, runtime or project pass.

| Requirement | Source behavior | Verification boundary |
| --- | --- | --- |
| REQ-027 — status/content spacers | [`layout.go`](../layout.go)/[`render.go`](../render.go) include two ideal spacer rows in natural preferred height and render one beside each status border when room permits. Spacers yield before content/essential hints; no single-page counter or extra footer returns. | Source-only formatting/inspection is not a runtime matrix. Revision-bound review and coordinator frame observations belong to [task 12-01](../.tasks/tasks/rex-whichkey-12-01-first-frame-fit.md). |
| REQ-028 — compact first creation and terminal seed | [`popup_initial.go`](../popup_initial.go) computes natural root rows and content-sized anchored bounds before creation, using validated origin geometry or the documented compact fallback. [`rex/geometry.go`](../rex/geometry.go) consumes window-relative leaf rectangles directly, handles documented zoom, and checks placement around measurement. [`rex/layer.go`](../rex/layer.go) seeds shell startup size in that single creation RPC and rejects stale placement. The existing watcher refines only the same owned layer. | Initial cell-level inset/rounding is an estimate where new-pane font/chrome differs. The fallback is hypothetical, never viewer provenance. Exact frozen revision review, combined compilation and promptly captured first-layer/terminal geometry belong to task 12-01 and [Phase12](../.tasks/phases/rex-whichkey-12-first-frame-fit.md); no runtime pass is asserted here. |

The coordinator gate must sample the **first** owned layer promptly, before the
one-second resize watcher, and retain its bounds plus initial startup terminal
rows/columns. A final-small-only observation cannot accept REQ-028. Retain
root/spacer/status frames, settled bounds, submenu/back, Escape exit0 and exact
owned-resource/private-file cleanup as well. No general tests/TDD/unit/race suite,
font/settings/config reload, native/destructive selection, remote trial or GUI
pixel experiment is authorized for leaves. Leaves perform source/format/whitespace
checks only; coordinator observations must identify the actual frozen revision.

See [appearance](appearance.md#actual-layer-height-and-measurement-limits) and
[initial geometry contract](launch-contract.md#initial-popup-geometry-and-terminal-seed)
for the 100×40 hypothetical fallback grid/30% startup cap, two-cell inset estimate,
half-row rounding, placement-race limits and later measured fitting. Historical
Phase11 settled-size observations below are not first-frame or spacer evidence.
Server bounds/seed/frame observations do not certify native GUI chrome removal.
REQ-026 remains **BLOCKED**, and Phase11/project completion remains unavailable.

## Compact popup presentation: REQ-024/025 and blocked REQ-026

Accepted **REXWHICHKEY-20261102-123334** under the
[Phase11 design](../.tasks/designs/rex-whichkey-compact-presentation.md) adds
compact status borders and actual content-fitting layer height. This is **partial
delivery**: hiding native Rex pane/tab/window controls is an accepted requirement,
not fulfilled or waived. Phase11/project completion remains blocked by REQ-026.

| Requirement | Source / capability state | Evidence boundary |
| --- | --- | --- |
| REQ-024 — compact status borders / conditional paging | [`render.go`](../render.go)/[`layout.go`](../layout.go) put title/breadcrumbs and key hints in separate full-width horizontal status borders, with no generated `Keys:` prefix, single-page counter or redundant page-footer row. Genuine multipage counters/Tab hints share the bottom border where width permits. Tiny heights yield status rows to content; critical confirmation text remains on the content row. `MenuFrame.PreferredHeight` reports natural current-width content plus both borders independently of clipping/page, now including the REQ-027 ideal spacers. | [Task 11-01](../.tasks/tasks/rex-whichkey-11-01-render-frame.md) records fresh APPROVED source review and coordinator Go compilation integrated at `f57dfa00` (tree `0d1a7aa375bc8b099b2100a344b1edf6ff4ef9dd`). Tiny/multipage/grapheme/ASCII/no-color paths are source-reviewed, not a runtime matrix. |
| REQ-025 — actual content-fitting owned layer | [`terminal.go`](../terminal.go)/[`popup_size.go`](../popup_size.go) publish bounded private natural-row hints on change. [`popup.go`](../popup.go) refreshes measured geometry through the existing watcher; [`rex/terminal.go`](../rex/terminal.go) requires current-owner matching applied/desired rows and timestamp provenance. [`rex/layer.go`](../rex/layer.go) reads/resizes only the exact owned floating layer. Position/max rectangle, bounded settling, cleanup and action/lifetime boundaries remain intact. | [Task 11-02](../.tasks/tasks/rex-whichkey-11-02-popup-height.md) records fresh APPROVED source review, combined Go compilation/`task build`, integration at `3c68b784bc0c90be94b0c4604040406f603557d3` (tree `7f6054689d0b505d5ed951b6c7f3e87b61eee4f0`) and the actual one-context coordinator gate below. Font-change generality, all positions and failure/transient paths are not runtime-certified. |
| REQ-026 — hide native controls for this popup only | **BLOCKED / not implemented.** Inspected Rex schemas/actions/docs provide no documented per-popup pane/tab/window control-hiding capability. Terminal text borders and resized bounds do not hide native chrome. | [Task 11-03](../.tasks/tasks/rex-whichkey-11-03-popup-chrome.md) owns blocker `rex-whichkey/presentation/native-chrome-blocker/1`. No hidden flag/global setting workaround or explicit user deferral is accepted. A supported Rex capability/update or explicit deferral is needed; source delivery cannot complete the phase/project. |

[Appearance](appearance.md) explains status-row priorities, semantic colors and
measurement limits. The [launch contract](launch-contract.md#private-preferred-row-hints-and-owned-layer-sizing)
traces 0700/0600 data-only hint publication/validation, parent-only owned resize,
anchor/max rules and exact hint/temp cleanup before dispatch. Geometry uses
measured rows/bounds and optional cell pixels, not guessed fonts; unmeasured
terminals pause adaptation. Native display still needs no app lookup, and native
action binding remains deferred until cleanup. FIFO parent lifetime, configured
original-VM authenticated 30-second expiry, confirmation and no uncertain retry
are unchanged. No font/settings/init/config reload is part of this delivery.

### Supplied producer gate, not general runtime certification

The coordinator-retained result is
`.mtn-tmp/popup-presentation/runtime/rex-whichkey-p11-t02-builder-a03/result.json`
in the authoritative checkout. It binds the `3c68b784` commit/tree above to an
actual Linux Rex normal native popup presented in a known viewing app, with
current-owner desired-size reports matching applied dimensions, not unclaimed
80×24 defaults. Its observed geometry is:

| State | Normalized height | Applied columns × rows | Nonblank terminal rows |
| --- | --- | --- | --- |
| Root | 0.103532609 | 98 × 6 | 5 |
| Tabs submenu | 0.119816160 | 98 × 7 | 7 |
| Back to root | 0.095390833 | 98 × 5 | 5 |

All three used measured 7×16-pixel cells and X=0.1/W=0.8. This records actual
layer growth/shrinkage, not just cropped `RenderMenu` text in the old H=0.6 box;
one extra root row is within the controller's rounding deadband. The artifact's
`root/`, `tabs/` and `back/` directories retain sequential `layout-*.json`,
`size-*.json`, `measurement-*.json` and screen samples over settling/stability,
plus `frame.vt` and `frame.html` with separated title/breadcrumb/hint borders and
no generated `Page1/1` or `Keys:`. `parent-final.json`, `after.json` and
`result.json` record Escape dismissal, exit0 and exact owned-popup cleanup.
These are supplied coordinator observations, not commands run by documentation
leaves, and do not certify native GUI controls removal.

The [11-04 assembly record](../.tasks/tasks/rex-whichkey-11-04-appearance-docs.md)
and [Phase11 record](../.tasks/phases/rex-whichkey-11-compact-presentation.md) own
revision-bound documentation review, exact combined Go/Taskfile compilation and
final bounded root/nested-submenu/back gate before integration. That gate must
record normalized layer bounds and terminal rows/cell pixels through settling
and stability, status-border VT/HTML frames, Escape exit0 and exact owned resource/
private-file cleanup. Earlier producer results do not preclaim a later assembly
pass; neither review nor compilation substitutes for measured layer geometry.

Only one existing font/viewport context was observed. No font/settings changes,
multipage/tiny/ASCII/no-color runtime trial, destructive/native action execution,
configured Lua callback, Darwin or remote trial is inferred. Font-change and
transient/missing-measurement behavior are source-described, not experimentally
certified. Builders/reviewers use static/whitespace checks only, with no project
tests/TDD/unit/race/general suite, compiler-backed checks or product/API/GUI/config/
remote experiments. L1–L9 remain applicable; L10 below records the chrome blocker.

Retained capability evidence is under `.mtn-tmp/popup-presentation/`:
`new-layer-api.json`, `resize-layer-api.json`, `terminal-creator.json`,
`client-actions.json` and `lua-reference.md`, indexed by `chrome-blocker.json`.
Creation/resize schemas expose bounds/layout/focus but no chrome option; terminal
creator options reject unknown properties; the captured connected-app registry
and published Lua layout reference offer no per-popup control suppression.
This is the inspected installed capability boundary, not a claim about every
Rex version. Server VT/HTML/screenshots cannot certify remote native GUI chrome.
Historical Phase07–10 evidence below is preserved, not promoted to REQ-024–026
acceptance.

## Native display independent of app targeting: REQ-023

Accepted **REXWHICHKEY-20261102-111430** under the
[Phase10 design](../.tasks/designs/rex-whichkey-app-independent-popup.md) refines
native launch-time app requirements without weakening native action safety or
configured original-VM binding. This is separate from the Phase09 tab-return work.

| Requirement | Implemented source | Evidence boundary |
| --- | --- | --- |
| REQ-023 — app-independent native display | [`launch.go`](../launch.go) validates syntax and captured endpoint/session/block, then reads server registration before GUI binding. [`native.go`](../native.go) retains immutable canonical defaults and optional requested selector without app lookup. [`popup.go`](../popup.go) uses resource-only native creation/wait/focus, deadline-free interaction and FIFO parent lease; owned cleanup needs no app and precedes selected-action targeting. Cancellation needs no app. Executable retained selection/confirmation resolves exact requested or unique eligible app, checks only the selected action and rechecks full ID/generation/origin before one dispatch. Configured handoffs retain prelaunch binding, strict decoding, original authenticated expiry and same-VM completion. | Source trace and formatting do not prove compilation or runtime behavior. Earlier single-app gates do not verify the deferred boundary. Exact revision-bound review, coordinator compilation and runtime artifacts belong to [task 10-01](../.tasks/tasks/rex-whichkey-10-01-app-independent-popup.md) and its [phase](../.tasks/phases/rex-whichkey-10-app-independent-popup.md). No runtime pass is asserted here. |

Native display/navigation/Escape do not require a live requested client, a singleton
app, known attachment kinds or any selected-client action registry. Malformed
selector syntax still fails early. Missing/ambiguous/unavailable app or unsupported
selected action refuses **action targeting after popup cleanup**, never silently
chooses focus/first/last/terminal owner, and never retries uncertainty. Configured
launch retains its live app prerequisite. Resource creation is not proof that a
GUI showed the menu.

The authorized coordinator-only runtime scope is normal startup/navigation/Escape,
valid-but-nonexistent selector startup/Escape, then one non-destructive selection
with that nonexistent target refusing after selection/cleanup with zero action
dispatch. Actual observed app counts remain distinct from source-reviewed zero/
multiple behavior; no user app disconnection is authorized. Builders use static
source/format/whitespace checks only: no tests/TDD, compilation, product/live API,
GUI/config/remote experiments or actual native action dispatch. Historical evidence
below is preserved and cannot certify this source revision. L1–L9 still apply.

## Tab-return and default Zoom removal: REQ-021/022

Accepted **REXWHICHKEY-20261102-102253** under the
[Phase09 design](../.tasks/designs/rex-whichkey-tab-return.md) removes default Zoom
and addresses native popup return to its originating tab. This is a new source
delivery, not a reinterpretation of Phase07/08 runtime passes.

| Requirement | Implemented source | Evidence boundary |
| --- | --- | --- |
| REQ-021 — remove default Zoom | [`defaults.go`](../defaults.go) and [`defaults/panes.lua`](../lua/rex_whichkey/defaults/panes.lua) no longer generate `z`/Zoom. The Lua composer still imports that real panes fragment. Split/close/session/tab/app bindings and confirmation remain; explicit custom `pane.zoom` remains available. | Static source trace only here. Fresh mixed Go/Lua/docs review and coordinator compilation/root-menu inspection must cover the frozen Phase09 revision. |
| REQ-022 — native tab return | [`terminal.go`](../terminal.go) saves/enables/restores xterm focus-reporting mode 1004. The real popup's ordered `PaletteDecoder` -> `terminalKeys` input path recognizes only complete unparameterized `CSI I` as focus-in. `RunTerminal` repaints via `drawTerminal` -> `RenderMenu` without `HandleKey`, preserving submenu/page/confirmation and current theme. Focus-out does nothing; no timer, focus stealing or replacement layer is added. | Decoder/render tracing is not runtime proof of native GUI visibility or terminal focus-report delivery. Fragmented/multiple/OSC-interleaved report framing, other terminal-input protection, resize/zero-size and restoration are statically inspected, not experimentally verified by the leaf. |

The native no-argument path remains `LaunchPopup` -> `launchNativePopup` ->
`runPopupWithLifetime(popupNative)` -> `RunNativePopupTerminal` -> `RunTerminal`
through `runPopupTerminal`. [`popup.go`](../popup.go) retains deadline-free native
interaction, the private parent-lifetime pipe, exact owned-only cleanup and no
uncertain retry. Phase09 retained endpoint/session/block/app-generation checks;
Phase10 above separates native resource liveness from deferred action binding.
Neither watcher checks active-tab visibility. No evidence here establishes
transient attachment loss as the tab-return cause. The configured original-VM route
still enforces authenticated 30-second expiry and graceful expected dismissal;
absence of native idle expiry does not renew configured authority.

Coordinator artifacts under `.mtn-tmp/tab-recovery/reproduction/` retained the same
native process/layer/menu through pre-change server focus and app tab-goto
round-trips. They did **not** establish actual GUI frame visibility, a crash or an
attachment-loss root cause. The captured new-layer API has no persistence option.
The separately reported several-app diagnostic is not an explanation of tab-return
visibility; Phase09 did not change client selection to work around it. Phase10's
separate native display correction is described above.

[Task 09-01](../.tasks/tasks/rex-whichkey-09-01-tab-return.md) and the
[phase record](../.tasks/phases/rex-whichkey-09-tab-return.md) own fresh review,
compilation and the coordinator-only bounded normal-popup gate: idle/color,
root/submenu tab-away/back via owned diagnostic tabs, genuine focus-in repaint,
retained navigation, Escape exit0 and exact popup cleanup. Actual observations
must identify frozen head/tree, process/layer/focus/frame and whether disappearance
was reproduced. Terminal/server captures alone do not certify the native app's
GUI or every platform. No Phase09 compilation, tests, TDD, product/API/GUI
experiments or configuration changes were run by this builder; no pending gate
is reported as passed. L1–L5/L8/L9 remain applicable.

## Popup polish: source delivery and targeted runtime evidence

Accepted **REXWHICHKEY-20261102-081654** adds REQ-018/019/020 under the
[popup-polish design](../.tasks/designs/rex-whichkey-popup-polish.md): native
idle-until-dismissal, reduced defaults and semantic key/icon/label colors. Native
interaction has no fixed deadline; configured original-VM Lua invocations retain
the authenticated 30-second expiry and graceful expected expiry without dispatch.
The Focus group and automatic host provider are absent from both default sets;
custom focus mappings, origin-focus restoration and explicit host providers remain.
[Appearance](appearance.md) documents actual rendering and fallback behavior.

### Fresh producer reviews and actual compilation

The authoritative task summaries record APPROVED independent source reviews,
Git integration and coordinator `GOTOOLCHAIN=local GOMAXPROCS=16 go build ./...`
on each combined candidate. These are supplied results, not commands run by the
documentation builder; same-family fresh review is not cross-family verification.

| Task / review event | Integrated commit / tree |
| --- | --- |
| [08-01 lifetime](../.tasks/tasks/rex-whichkey-08-01-lifetime.md) / `rex-whichkey-p08-t01-reviewer-a01/review/1` | `ca03b464ad8cbe1a4d31b367d73a2784fe4be56e` / `487b59d3e82b367730f65947835e6f4c7702f059` |
| [08-02 defaults](../.tasks/tasks/rex-whichkey-08-02-defaults.md) / `rex-whichkey-p08-t02-reviewer-a01/review/1` | `3ad78d627d6355c17731da53de6c710736f278e6` / `7a9972f5157393e53e20f42a340b16085702e7df` |
| [08-03 theme](../.tasks/tasks/rex-whichkey-08-03-theme.md) / `rex-whichkey-p08-t03-reviewer-a01/review/1` | `3e9a1fd19a6a930fac293966d0ee7b4114aab2dc` / `0550e40e9efa97ce41fadd81e13913e041c3be31` |

Their integration artifacts are in the authoritative checkout under
`.mtn-tmp/menu-polish/rex-whichkey-p08-t0N-builder-a01/integration/integration.json`
(N = 1, 2, 3); the matching reviewer directories retain `review.json`.
The [08-04 assembly record](../.tasks/tasks/rex-whichkey-08-04-docs.md) and
[Phase08 record](../.tasks/phases/rex-whichkey-08-popup-polish.md) bind documentation
review and final combined compilation to exact source identities. Earlier
producer reviews/builds do not approve later documentation or prove runtime use.

### Coordinator-only final gate and evidence interpretation

Before final assembly completion, the coordinator compiles the frozen reviewed
combined candidate and runs the user-authorized normal popup in the existing
Linux Rex session, without editing/reloading configuration:

1. Capture the real floating root menu with Sessions/Tabs and no default Focus or
   Rex-known-hosts group; a standalone renderer is not a substitute.
2. Leave it idle for at least **35 seconds**, then record the still-live parent
   and visible/usable popup beyond the former deadline.
3. Retain VT **and** HTML frames and compare key, icon and label/group spans with
   the semantic source roles. Plain text alone cannot prove color; a count of
   foreground styles alone does not identify which role received them.
4. Navigate into Tabs, then Escape back and Escape-dismiss without selecting an
   action. Record parent exit0 and disappearance of the exact owned popup
   block/layer, not broad resource cleanup.

The coordinator's capture layout is
`.mtn-tmp/menu-polish/runtime/rex-whichkey-p08-t04-builder-a01/`: `result.json`
binds candidate commit/tree and exact resource IDs; `root-screen.txt`,
`idle-parent.json`, `idle-screen.txt`, `color-frame.vt`, `color-frame.html`,
`foreground-styles.json`, `submenu-screen.txt`, `parent-final.json` and `after.json`
retain observations. The assembly/phase records are the durable index of actual
outcome and artifact references, including a replacement run if needed. Naming
these artifacts or describing the gate is **not a passing runtime result**.
Builders/reviewers do not run it. No project unit/race/TDD/general suite,
destructive/native action execution, configured callback trial, Darwin/remote
certification or configuration reload is authorized by this narrow gate.

### Historical native-startup observation

[Task 07-05](../.tasks/tasks/rex-whichkey-07-05-native-default-popup.md) implements
accepted **REXWHICHKEY-20261101-201834**, refining **REXWHICHKEY-20261101-101502**.
It records APPROVED review `rex-whichkey-p07-t05-reviewer-a01/review/1`, coordinator
Go compilation/`task build`, and an actual normal native popup launch, Tabs
navigation, Escape dismissal, exit0 and exact owned-popup cleanup in one Linux
Rex session at `9f9b903e7cb8a5ed526a54a506aee408e8f11fd0` (tree
`c124f58cbd1fa10c9f65795c786928d14b5c14fe`). Artifacts:
`.mtn-tmp/usability/rex-whichkey-p07-t05-builder-a01/integration/integration.json`
and `.mtn-tmp/live-launch/rex-whichkey-p07-t05-builder-a01/result.json`.

Earlier `.mtn-tmp/live-launch/` captures include a failed configured-only normal
invocation for absent custom actions and a successful standalone renderer
navigation/dismissal diagnostic. Neither is the later normal-popup pass. The
07-05 pass, in turn, predates Phase08; it does not establish 35-second idle,
reduced-default or semantic-color acceptance. No selected native action or Lua
callback success is inferred.

## Scope and evidence boundary

This inventory separates **source delivery, compilation and targeted runtime
observations**; it is not a general acceptance-test result.
The original Phase06 inventory described baseline commit
`4980e1f95a96731a1baee445562864c7e6e771bd` (tree
`68b72d7d8bdb2b4ccd99a7312c4bb9d3e772b2bb`), plus documentation task 06-03,
subsequently reviewed and integrated at `40c4f53` as recorded below. Those completed
phase/task identities and historical limitations are preserved.

Phase07 endpoint/launch baseline `69078fcee3dc0bbd9598e1b58e33920934e82497`
(tree `82dc432f715295a3b82829be8be509e83aa0197a`) and the six Markdown files
owned by task 07-03 were independently reviewed and integrated at
`1af28738a79fe20f2634c21a3405d8d1f5ff7bc2` (tree
`a256fa6e761d1f5cf160b448829d332f7cded754`). Subsequent metadata checkpoints
`a50ab6d8` and `673b2f28` record that delivery. Coordinator Go compilation and
Nix package-build results are recorded below. These are revision-bound results,
not approval of later edits or functional acceptance; this inventory cannot
approve its own changes or complete a managed task/phase.

Requirements come from `.tasks/prds/rex-whichkey.md`, with stable REQ-001–REQ-015.
`.tasks/prds/rex-whichkey-usability.md` adds REQ-016/017 without renumbering or
reopening completed phases02–06. The popup-polish design adds REQ-018–REQ-020;
its accepted lifetime/default/color refinements supersede earlier source descriptions
where noted below. The tab-return design adds REQ-021/022 without rewriting that
history. The app-independent-popup design adds REQ-023 and refines only native
display/action binding. The compact-presentation design adds REQ-024–REQ-026;
REQ-026 remains explicitly blocked, so earlier completion summaries do not mean
the current phase/project is fully delivered.

The later direction **REXWHICHKEY-20261031-072107** supersedes earlier experimental,
TDD, live-Rex and final acceptance-matrix execution requirements, not product
behavior or safety intent. REQ-015 now records integrated delivery, documentation
and honest limitations. The authenticated post-return bridge follows accepted
**REXWHICHKEY-20261031-220348**; the separate constructor-bound endpoint follows
**REXWHICHKEY-20261101-015342**, selected under delegated implementation authority.
The latter is now **Refined** by accepted **REXWHICHKEY-20261101-101502** for the
user-requested default endpoints: native Rex discovery replaces mandatory explicit
setup, retaining overrides, opaque-context separation and fail-closed origin
checks. The duplicated-path planning proposal **REXWHICHKEY-20261101-100106** is
**Superseded**, not retroactively approved or implemented. This reflects the
coordinator's recorded lineage under delegated authority, not a new decision by
this documentation task. Historical draft IF-* signatures are not adopted contracts.

**Not delivery gates under the retained policy:** project tests, TDD, experiments,
Nix package checks/install checks, Home Manager evaluation/activation, general graphical
checks, live API probes, remote trials, disconnected-origin exercises and
configuration reloads. Go compilation and Nix package builds did run, as recorded
below; they are not test or runtime passes. No new acceptance runner is
introduced. The targeted task `07-05` runtime gate above is the later explicit
exception, extended by Phase08 to the idle/default/color gate, Phase09 to the
bounded tab-return gate, Phase10 to startup/deferred nonexistent-target refusal,
and Phase11 to measured compact root/submenu/back layer geometry/status frames
as described above. Only revision-bound recorded observations count as passes. Existing test files and
historical evidence are preserved, not converted into current passes.

Historical exception: metadata commit `7d4cff8` inadvertently ran pre-commit
build/lint/unit/race hooks. The coordinator disclosed those runs; they are not
acceptance evidence for subsequent delivery. Later commits disabled hooks per
invocation without changing hook definitions.

Documentation leaves use source inspection and whitespace checks only;
compilation is not authorized for this leaf builder. Supplied coordinator
compilation is recorded separately below. Reading source/docs, compilation and
independent source review do not establish runtime compatibility, installed
configuration success, exact graphical geometry or end-to-end behavior.

## Requirement-to-source inventory

“Implemented source” means the named files contain the path described; it does
not mean that the requirement passed a runtime check. Review references in the
next section identify the independently reviewed deliveries. Limits L1–L8 apply
as indicated and across the assembled application; L9 additionally covers the
Phase07 startup/default-endpoint paths.

| Requirement | Implemented source and integration | Outstanding verification / limitations |
| --- | --- | --- |
| REQ-001 — rewrite/lifecycle | [`cmd/main.go`](../cmd/main.go) composes root package `whichkey` (module remains `temp`); [`popup.go`](../popup.go) shares lifecycle between native foreground parent/child (no idle deadline) and finite configured detached controller/child; [`lua/rex_whichkey.lua`](../lua/rex_whichkey.lua) launches configured mode and returns. No persistent daemon or service hierarchy. Reviews 04-04/04-05/08-01. | Targeted native observations are revision-bound above; configured prompt return, callbacks and failure/terminal cleanup are not generally runtime-certified. L1/L3. |
| REQ-002 — typed CLI client | [`rex/rex.go`](../rex/rex.go), [`exec.go`](../rex/exec.go), [`session.go`](../rex/session.go), [`window.go`](../rex/window.go), [`block.go`](../rex/block.go), [`layer.go`](../rex/layer.go): concrete Client/Session/Window/Block, generic Call, argv execution, JSON/error handling, deadlines and captured targets. Reviews 02-01–02-04. | No current fixture, cancellation, invalid-JSON or live API tests were run. Timeout/decode failure does not prove a mutation had no effect. L1/L2/L3. |
| REQ-003 — Lua loading | [`lua/rex_whichkey/config.lua`](../lua/rex_whichkey/config.lua), [`lua/rex_whichkey.lua`](../lua/rex_whichkey.lua), [`config.go`](../config.go): server-side module/inline loading and strict resolved-data decoding. Defaults < imported < inline; spec/icons replace whole values. Review 03-01 and composition 04-05. | Actual module lookup, PATH and inline/import deployment remain unexecuted. Ordinary require caching applies; imports are trusted executable Lua. L1/L4. |
| REQ-004 — mappings | [`mapping.go`](../mapping.go) and [mapping-format.md](mapping-format.md) own shared data/validation; [`mapping.lua`](../lua/rex_whichkey/mapping.lua) compiles full sequences, implicit groups, nested inheritance, conflicts and `<Space>`/`<lt>`; [`evaluate.lua`](../lua/rex_whichkey/evaluate.lua) resolves conditions/dynamic strings and retains hidden active entries. Reviews 03-02/03-03/03-05, provider revision 05-04. | No new positive/negative fixtures executed. Lua consumes the Go-owned standard, not a parallel schema or Go callback VM. L1/L4. |
| REQ-005 — runtime callbacks | [`evaluate.lua`](../lua/rex_whichkey/evaluate.lua), [`dispatch.lua`](../lua/rex_whichkey/dispatch.lua), [`mapping.go`](../mapping.go), [`popup.go`](../popup.go) and Lua launcher: callbacks/args/tickets remain in the original server Lua VM; a fresh origin-targeted configuration action handles authenticated one-shot completion. Reviews 03-03/03-04/03-05/04-05. | No end-to-end callback run. Opaque server identity, connection-generation semantics and actual queued-client routing remain runtime-unverified. L2/L3/L4. |
| REQ-006 — navigation | [`menu.go`](../menu.go), [`action.go`](../action.go), [`render.go`](../render.go), [`terminal.go`](../terminal.go): arbitrary-depth groups, breadcrumbs, Esc/Backspace back/root dismissal, unmapped-q dismissal and Tab paging; hidden/off-page mappings stay active. Reviews 04-01/04-02/04-04/04-05. | Only targeted native Tabs/Escape observations are recorded above, not a general state-transition, paging or interactive suite. L1/L5. |
| REQ-007 — confirmation | [`menu.go`](../menu.go) accepts exactly y at the prompt; any other key cancels without fallthrough. [`mapping.go`](../mapping.go), [`popup.go`](../popup.go) and [`dispatch.lua`](../lua/rex_whichkey/dispatch.lua) validate confirmation against the retained mapping for named actions and callbacks. Reviews 04-01/03-05/03-04/04-05. | No confirmation/cancellation/destructive-action runtime exercise. Confirmation is an authenticated popup assertion, not cryptographic proof of a physical keypress. L2/L4. |
| REQ-008 — geometry/rendering | [`layout.go`](../layout.go)/[`render.go`](../render.go) constrain cell-based columns, paging, grapheme-safe clipping, breadcrumbs/icons/help; [`popup.go`](../popup.go) selects native top/bottom/left/right/center maximum bounds; [`popup_size.go`](../popup_size.go) adds measured content-driven owned height (REQ-025). [`terminal.go`](../terminal.go) consumes actual frame placements/Lines and passes Pages to key handling. Historical reviews 04-02/04-04/04-05; Phase11 refinement above. | One-context compact layer bounds are observed above; no five-position, clipping, resize/font-change or tiny-terminal geometry matrix is certified. No old full-window overlay workaround is adopted. L1/L5. |
| REQ-009 — theme/terminal safety | [`theme.go`](../theme.go) supplies palette decoding, semantic contrast-aware styles, default/fallback colors, NO_COLOR and ASCII decorations; [`render.go`](../render.go) uses styles/decorations. [`terminal.go`](../terminal.go) orders palette metadata before keys and restores terminal state before returning a result. Reviews 04-03/04-02/04-04/04-05. | Real palette behavior, glyph widths and restoration were not exercised. Unknown background uses immediate ANSI semantic accents with default body text/background; measured palette contrast applies only when background metadata is valid. ASCII does not transliterate user labels. Phase08 review above. L1/L3/L5. |
| REQ-010 — built-ins/session parity | [`defaults/sessions.lua`](../lua/rex_whichkey/defaults/sessions.lua), [`tabs.lua`](../lua/rex_whichkey/defaults/tabs.lua), [`panes.lua`](../lua/rex_whichkey/defaults/panes.lua), [`app.lua`](../lua/rex_whichkey/defaults/app.lua), [`init.lua`](../lua/rex_whichkey/defaults/init.lua) and dispatcher provide session/tab lifecycle/navigation, split/close, sidebar and palette. Default Focus and host groups were removed in 08-02 and Zoom in Phase09; custom focus/zoom actions and explicit host providers remain. Reviews 05-01–05-03/05-04, 03-04 and 08-02. | Each actual Rex action, native rename/picker interaction and old-plugin runtime parity remain unexecuted. Session switch delegates Rex's picker, not an enumerated popup session list. L1/L2. |
| REQ-011 — origin/focus/cleanup | [`rex/block.go`](../rex/block.go)/[`layer.go`](../rex/layer.go), [`dispatch.lua`](../lua/rex_whichkey/dispatch.lua), [`popup.go`](../popup.go), Lua launcher: captured server/session/block, current placement lookup, origin focus and returned-owned-terminal cleanup. Configured binding retains app generation throughout; native app binding/generation begins only after executable selection and cleanup (REQ-023). Reviews 02-03/03-04/04-05 and documentation 05-05 describe their earlier revisions. | Recorded native dismissal cleanup is limited to the frozen targeted gate above. Moved-block, reconnect, error cleanup and disconnected-origin behavior remain unverified. No local/current-client fallback or broad cascading cleanup is intended. L2/L3/L4. |
| REQ-012 — known hosts | [`rex/hosts.go`](../rex/hosts.go) decodes `rex hosts --json`; [`hosts.go`](../hosts.go)/[`terminal.go`](../terminal.go), [`defaults/hosts.lua`](../lua/rex_whichkey/defaults/hosts.lua) expand explicit known-hosts provider groups once per invocation. IDs, labels, home metadata and warnings remain distinct; generated rows are non-executable. Reviews 02-04/05-04. | No current discovery or reachability trial. A failed listing produces a warning, not successful empty inventory. Execution-host inventory differs from origin-server selection. L6. |
| REQ-013 — remote compatibility | Captured targeting in [`rex/`](../rex), server-local loading/launch in Lua/config/popup paths, and [remote-compatibility.md](remote-compatibility.md). Reviews 02-01–02-04/03-01/03-04/04-05/05-05. | No remote deployment/trial. Helper/controller/child must share the Rex server host's executable and temporary filesystem; client-local installation does not install remotely. L1/L2/L3/L4/L6. |
| REQ-014 — packaging/config ownership | [`flake.nix`](../flake.nix), [`nix/package.nix`](../nix/package.nix), [`nix/home-manager.nix`](../nix/home-manager.nix) package the executable/Lua and opt-in sibling links. User init/mapping files are not managed; links are non-forcing. Reviews 06-01/06-02. | Recorded Nix package builds passed with checks/install checks disabled; no Home Manager evaluation/activation or existing-file byte-preservation trial. Declared four-system outputs are not certification. L7. |
| REQ-015 — integrated delivery/docs | [README](../README.md), this inventory, [mapping standard](mapping-format.md), [remote compatibility](remote-compatibility.md), [configuration examples](../examples/README.md) and [Home Manager cases](../tests/packaging/home-manager/README.md) describe the delivered source and retained limits. Recorded reviews below include 06-03 and the 07-03 revision integrated at `1af28738`; approvals apply only to their frozen revisions. | No acceptance runner, final GUI matrix, remote/disconnected-origin test or inferred functional pass. Historical evidence is not current execution. L8. |

### Usability extension: REQ-016/017

| Requirement | Implemented source and integration | Outstanding verification / limitations |
| --- | --- | --- |
| REQ-016 — ordinary popup launch | [`cmd/main.go`](../cmd/main.go) -> [`launch.go`](../launch.go) captures exact terminal origin, preflights placement and reads server registration before GUI binding. Native app binding is deferred to executable selection (REQ-023); configured launch still captures the app generation before activation. An absent implicit default uses [`native.go`](../native.go)/[`defaults.go`](../defaults.go) and the shared [`popup.go`](../popup.go) lifecycle; registered/explicit actions keep the original-VM configured route. Historical configured-only review 07-02/integration `69078fc` does not approve this correction. [Launch contract](launch-contract.md) traces both paths. | Historical configured-only launch failed; later 07-05 records a normal native popup/navigation/dismissal/exit0/cleanup pass. Neither that result nor historical producer review/compilation establishes Phase08 idle/default/color behavior; use the assembly evidence index above. L1–L5/L9. |
| REQ-017 — native platform endpoints | [`rex/endpoint.go`](../rex/endpoint.go) owns explicit override > nonempty execution-host REX_SERVER > bounded no-autostart native Linux/macOS Rex identity discovery, capturing/validating a concrete endpoint. [`popup.go`](../popup.go) preparation and the public caller consume it; Lua permits `popup.new(defaults)` and forwards preparation's canonical server. Private controller handoffs never ambient-default. Review 07-01, later composition review 07-02; integrated `d1164e7`/`69078fc`. | No default-discovery/platform matrix, custom-server or Darwin runtime certification. Exact native socket algorithms belong to Rex; no guessed path formula or platform certification. CLI identity shape requires compatible installed Rex; no autostart, scan or alternate-server retry. L1/L2/L4/L9. |

### Popup-polish extension: REQ-018/019/020

| Requirement | Implemented source | Evidence boundary |
| --- | --- | --- |
| REQ-018 — native idle / configured expiry | [`native.go`](../native.go) selects trusted native policy; [`popup.go`](../popup.go) `runPopupWithLifetime`, `waitPopup` and `RunNativePopupTerminal` keep interaction deadline-free but parent/origin-bound. `runPopupLifecycle`/`RunPopupTerminal` retain configured expiry and graceful expected cancellation, never stale dispatch. [`cmd/main.go`](../cmd/main.go) keeps native pipe and configured deadline modes distinct. | Fresh 08-01 review/compilation above; final assembly owns actual 35-second idle/navigation/dismissal evidence. No configured-expiry/callback runtime trial. L2–L4. |
| REQ-019 — reduced defaults | [`defaults.go`](../defaults.go), [`defaults/panes.lua`](../lua/rex_whichkey/defaults/panes.lua), [`defaults/init.lua`](../lua/rex_whichkey/defaults/init.lua): no Focus or automatic host group. [`hosts.go`](../hosts.go) performs no inventory lookup without an enabled explicit marker. Origin restoration and optional provider code remain. | Fresh mixed Go/Lua 08-02 review/compilation above; final assembly inspects the actual root menu and Tabs. No host inventory or focus action exercise. L2/L6. |
| REQ-020 — semantic color | [`theme.go`](../theme.go) defines `ThemeIcon`, immediate ANSI fallback and contrast-aware palette roles; [`render.go`](../render.go) styles sanitized/clipped key/icon/label spans; [`terminal.go`](../terminal.go) renders before querying and applies partial/late metadata. [Appearance](appearance.md) describes opt-outs and limits. | Fresh 08-03 source review/compilation above; only actual VT/HTML captures establish observed colors in that app. No palette matrix or geometry/platform certification. L3/L5. |

### Phase07 intent and implemented source traces

The accepted usability PRD/design define the intended product path; they are not
implementation evidence by themselves. The following traces were inspected in
the actual integrated baseline:

1. `cmd/main.go` `run`/`runLaunch` -> `LaunchPopup` -> one explicitly targeted
   configuration action -> `menu.register_launch` `run(ctx,args)` -> `menu.launch`.
   Public launch consumes no user JSON and never substitutes low-level `menu`.
2. Lua `config.new`/`mapping.compile` and shipped `defaults/init.lua` ->
   `dispatch.issue`/evaluation -> real `prepare` -> `PreparePopupWithServer` ->
   `rex.ResolveEndpoint` -> captured capability/executable/server -> `dispatch.bind`
   -> detached `RunPopupController`. The actual opaque ctx and original closures
   remain in the server VM; the launch envelope only constrains actual origin.
3. Controller exact-origin preflight -> owned floating terminal ->
   `RunPopupTerminal` -> `RunTerminal` -> `RenderMenu`/`HandleKey` -> terminal
   restoration/result commit -> owned popup cleanup/origin revalidation/focus ->
   authenticated fresh completion -> retained original-VM callback, `rex.invoke`
   or `rex.client.queue`, with unchanged replay/expiry/generation/snapshot/
   confirmation safeguards.

These traces are source-consistency observations, not runtime results. The
coordinator records below bind producer and documentation reviews to frozen
revisions and actual integration/compilation evidence. Later documentation
corrections require fresh independent review and coordinator integration; prior
approvals do not cover them. The coordinator alone records phase
delivery/completion. Those baseline traces predate the native correction and its
targeted coordinator-owned runtime gate above. No unsupplied `task build` result
is inferred from recorded Go/Nix commands.

## Recorded independent source-review outcomes

These are **reported coordinator delivery records**, not reviews performed by this
documentation builder. Each listed outcome is APPROVED for its frozen source
revision and subsequently recorded as integrated. The reviewer contexts were
fresh but same-family; cross-family verification was unavailable. Later changes
to shared files are represented by their later review row, especially 05-04.
No approval here means runtime acceptance.

Task shorthand `PP-TT` refers to the exact linked record in the authoritative
parent `.tasks` collection. Commit abbreviations identify integrated source, not
worker heads or runtime binaries. Full identities and artifact pointers remain
in each record's compact delivery summary.

| Task record | Recorded APPROVED review event | Integrated commit |
| --- | --- | --- |
| [02-01 CLI execution](../.tasks/tasks/rex-whichkey-02-01-cli-execution.md) | `rex-whichkey-p02-t01-reviewer-a01/review/1` | `020a31e` |
| [02-02 resources](../.tasks/tasks/rex-whichkey-02-02-resources.md) | `rex-whichkey-p02-t02-reviewer-a01/review/1` | `3ec92e9` |
| [02-03 blocks/layers](../.tasks/tasks/rex-whichkey-02-03-block-layers.md) | `rex-whichkey-p02-t03-reviewer-a02/review/2` | `fb7ed12` |
| [02-04 inventory](../.tasks/tasks/rex-whichkey-02-04-host-inventory.md) | `rex-whichkey-p02-t04-reviewer-a01/review/1` | `4335256` |
| [03-01 Lua loading](../.tasks/tasks/rex-whichkey-03-01-lua-config.md) | `rex-whichkey-p03-t01-tasks-a01/review/1` | `6109fff` |
| [03-02 mapping compiler](../.tasks/tasks/rex-whichkey-03-02-mapping-keys.md) | `rex-whichkey-p03-t02-tasks-a01/review/1` | `39ca52f` |
| [03-03 evaluation](../.tasks/tasks/rex-whichkey-03-03-dynamic-mappings.md) | `rex-whichkey-p03-t03-tasks-a01/review/1` | `44f22c1` |
| [03-04 dispatch](../.tasks/tasks/rex-whichkey-03-04-invocation-dispatch.md) | `rex-whichkey-p03-t04-tasks-a02/review/2` | `0a0ac89` |
| [03-05 Go-owned standard](../.tasks/tasks/rex-whichkey-03-05-mapping-format.md) | `rex-whichkey-p03-t05-reviewer-a03/review/3` | `55da14c` |
| [04-01 navigation](../.tasks/tasks/rex-whichkey-04-01-navigation.md) | `rex-whichkey-p04-t01-reviewer-a02/review/2` | `2d612d6` |
| [04-02 layout](../.tasks/tasks/rex-whichkey-04-02-layout.md) | `rex-whichkey-p04-t02-reviewer-a03/review/3` | `43c62df` |
| [04-03 theme](../.tasks/tasks/rex-whichkey-04-03-theme.md) | `rex-whichkey-p04-t03-reviewer-a02/review/2` | `f71fd47` |
| [04-04 terminal](../.tasks/tasks/rex-whichkey-04-04-terminal-lifecycle.md) | `rex-whichkey-p04-t04-reviewer-a03/review/3` | `0d8fc4c` |
| [04-05 popup composition](../.tasks/tasks/rex-whichkey-04-05-popup-integration.md) | `rex-whichkey-p04-t05-reviewer-a03/review/3` | `8f2a196` |
| [05-01 sessions](../.tasks/tasks/rex-whichkey-05-01-session-actions.md) | `rex-whichkey-p05-t01-tasks-a01/review/1` | `647b29b` |
| [05-02 tabs](../.tasks/tasks/rex-whichkey-05-02-tab-actions.md) | `rex-whichkey-p05-t02-tasks-a01/review/1` | `c8c810f` |
| [05-03 panes/application](../.tasks/tasks/rex-whichkey-05-03-pane-actions.md) | `rex-whichkey-p05-t03-tasks-a01/review/1` | `7525b2c` |
| [05-04 host menu/default composer](../.tasks/tasks/rex-whichkey-05-04-host-menu.md) | `rex-whichkey-p05-t04-reviewer-a02/review/2` | `cfba4cf` |
| [05-05 remote documentation](../.tasks/tasks/rex-whichkey-05-05-remote-context.md) | `rex-whichkey-p05-t05-tasks-a02/review/2` | `cddf437` |
| [06-01 Nix package](../.tasks/tasks/rex-whichkey-06-01-nix-package.md) | `rex-whichkey-p06-t01-tasks-a01/review/1` | `908a339` |
| [06-02 Home Manager](../.tasks/tasks/rex-whichkey-06-02-home-manager.md) | `rex-whichkey-p06-t02-tasks-a01/review/1` | `4980e1f` |
| [06-03 release documentation](../.tasks/tasks/rex-whichkey-06-03-release-acceptance.md) | `rex-whichkey-p06-t03-tasks-a01/review/1` | `40c4f53` |
| [07-01 endpoint policy](../.tasks/tasks/rex-whichkey-07-01-endpoint-policy.md) | `rex-whichkey-p07-t01-reviewer-a02/review/2` | `d1164e7` |
| [07-02 ordinary launch](../.tasks/tasks/rex-whichkey-07-02-direct-launch.md) | `rex-whichkey-p07-t02-reviewer-a01/review/1` | `69078fc` |
| [07-03 usage documentation](../.tasks/tasks/rex-whichkey-07-03-usage-assembly.md) | `rex-whichkey-p07-t03-tasks-a01/review/1` | `1af28738` |
| [07-05 native startup](../.tasks/tasks/rex-whichkey-07-05-native-default-popup.md) | `rex-whichkey-p07-t05-reviewer-a01/review/1` | `9f9b903` |

Phase08 fresh producer reviews/compilation and final assembly evidence pointers
are indexed separately at the top of this inventory.

The 02-03 summary explicitly records cleanup-race finding `P02-T03-R1` resolved
before its approved second review; it retains refusal of unsafe remaining-layer
removal. The later final approvals do not relabel earlier blocked attempts as
successful. Endpoint task 07-01's earlier actionable-error finding
`REXWK-P07-T01-R1` was resolved before its approved second review; the earlier
blocked attempt is not relabeled a pass. The 07-03 row covers only its earlier
frozen revision, not later corrections; inclusion here is not self-approval.

### Historical compilation versus current checks

The Phase 2 summaries record coordinator `go build ./rex` on their combined
integration candidates. Later task summaries record coordinator
`GOTOOLCHAIN=local GOMAXPROCS=16 go build ./...` plus source/identity checks on
combined candidates, including 06-03 and the two Phase07 producer records below.
Those are **supplied earlier Go compilation records**, not commands rerun by
this task. Go compilation does not compile Lua, evaluate Nix, exercise existing
tests or prove that the packaged program starts correctly. Separately recorded
Nix package builds below are actual build results, not inferences from Go
compilation or Nix source review.

For Phase07, the authoritative completed-task summaries and these coordinator
integration events report `GOTOOLCHAIN=local GOMAXPROCS=16 go build ./...` passed
on each combined candidate, unchanged reviewed file contents, `git diff --check`
and verified target identity/ancestry:

| Coordinator integration event | Integrated commit / tree | Retained artifact relative to parent checkout |
| --- | --- | --- |
| `rex-whichkey-p07-t01-builder-a02/integration/2` | `d1164e75028760421446a926421d50750ddade28` / `07f7faabe9248e41c94a6d9e05ede573301273cc` | `.mtn-tmp/usability/rex-whichkey-p07-t01-builder-a02/integration/integration.json` |
| `rex-whichkey-p07-t02-builder-a01/integration/1` | `69078fcee3dc0bbd9598e1b58e33920934e82497` / `82dc432f715295a3b82829be8be509e83aa0197a` | `.mtn-tmp/usability/rex-whichkey-p07-t02-builder-a01/integration/integration.json` |
| `rex-whichkey-p07-t03-builder-a01/integration/1` | `1af28738a79fe20f2634c21a3405d8d1f5ff7bc2` / `a256fa6e761d1f5cf160b448829d332f7cded754` | `.mtn-tmp/usability/rex-whichkey-p07-t03-builder-a01/integration/integration.json` |

The corresponding fresh mixed Go/Lua review artifacts are
`.mtn-tmp/usability/rex-whichkey-p07-t01-reviewer-a02/review.json` and
`.mtn-tmp/usability/rex-whichkey-p07-t02-reviewer-a01/review.json`. They approve
frozen worker trees matching the integrated trees above. These ignored transport
files are coordinator-retained evidence, not installed package files; durable
compact summaries remain in the linked authoritative task records. Documentation
review `rex-whichkey-p07-t03-tasks-a01/review/1` approved frozen head `3189ebbf`
with the same tree as integrated `1af28738`; its artifact is
`.mtn-tmp/usability/rex-whichkey-p07-t03-tasks-a01/review.json`.

The 06-01 package delivery at `908a3395` records a passing
`nix build --no-link --no-write-lock-file --cores 16 --max-jobs 1 .#default`
with project checks disabled. The later coordinator project summary
`rex-whichkey-p03-coordinator-a01/project/summary/2`, anchored at metadata commit
`a50ab6d86039d3d4a1fa674fc9289936a8ca1ccc` (tree
`436fde5c9c5ccd08fddd5e08a0e6312d3cb2d416`), records these final delivered-source
results:

| Recorded command | Result | Retained log relative to parent checkout |
| --- | --- | --- |
| `GOTOOLCHAIN=local GOMAXPROCS=16 go build ./...` | Passed | `.mtn-tmp/usability/project-final/go-build.log` |
| `nix build --no-link --no-write-lock-file --cores 16 --max-jobs 1 .#default` | Passed; checks/install checks disabled | `.mtn-tmp/usability/project-final/nix-build.log` |

The matching summary is retained in
`.mtn-tmp/usability/project-final/summary.json` and the authoritative project
record; metadata commit `673b2f28` checkpoints it. These supplied results do not
claim a later documentation correction was reviewed or built, a `task build`
command ran, or tests, Darwin builds, Home Manager activation or runtime trials
passed. The documentation leaf's actual static results and exact revision are
returned separately in its build payload.

### Preserved earlier experiments

The nested [terminal experiment](experiments/rex-validation/tui/VALIDATION.md) and
[attempt A02](experiments/rex-validation/tui/VALIDATION-A02.md),
[A03](experiments/rex-validation/tui/VALIDATION-A03.md),
[A04](experiments/rex-validation/tui/VALIDATION-A04.md) documents retain earlier
failures, corrective work and recorded unit/PTY/race/root-gate results. They
predate the stop instruction and describe an isolated diagnostic module, not the
assembled product. A04's green gates meant readiness for independent review only,
not GUI qualification, framework adoption or release acceptance. Its preserved
limits include an earlier unexplained transient poll-cancellation failure,
intersecting-wide-cell behavior and OSC-looking bracketed paste not claimed fixed
there. Those historical experiment limits/results must not be silently asserted
as either confirmed defects or passes in the different current product paths.
Root commands did not traverse that nested module. No historical command was
rerun or scheduled by this documentation task; cancelled research is not product
delivery.

## Outstanding limitations

- **L1 — deployment/API/platform:** no minimum tested Rex build or Linux/macOS
  runtime certification. The flake declares four systems; popup/terminal code
  uses Unix facilities. Compatible installed APIs, module lookup, server PATH,
  executable permissions and shared server-local temporary filesystem are
  prerequisites, not verified deployment outcomes. CLI-origin client/block
  actions require the graphical client's Remote Control setting. Configuration
  changes remain user-owned; no init.lua rewrite, automatic reload or restart.
- **L2 — origin/races:** source targets captured server/session/block, following
  block moves within that session, not searching other sessions. Native display
  uses resource liveness alone; app ID/`connected_at` bind after selection/cleanup.
  Configured mode binds app generation before activation. Missing/ambiguous/lost/
  changed required identities fail closed at their respective boundary.
  Inspect/focus/execute are separate Rex operations, not an atomic transaction;
  external moves, disconnects, focus or configuration changes can
  race side effects after the last check. Disconnected-origin rejection and real
  graphical-client action routing remain untested.
- **L3 — cleanup/termination/uncertainty:** only the returned owned terminal is
  closed; a remaining layer causes an error rather than cascading closure.
  Creation can succeed before timeout/decoding loses its resource identity.
  Network/server loss, SIGKILL/SIGSTOP and terminal loss can prevent cleanup.
  Terminal restoration is bounded best effort, with the documented transient
  Darwin PENDIN caveat, not a byte-for-byte platform comparison. The configured
  controller reserves five seconds of its original lifetime for UI teardown/
  completion; native interaction has no fixed expiry. Both have bounded deferred
  cleanup, and native parent loss cancels the child through its lifetime pipe.
  No retry of uncertain mutations/completion, label-based scavenging, recursive cleanup or compensating daemon is provided.
- **L4 — trust/lifetime:** configured source binds a Go crypto/rand capability,
  retained mapping/confirmation, original runtime context, config and connection
  generations to one invocation. Native app binding is deferred per REQ-023.
  Configured Lua lifetime is 30 seconds from
  reservation (the launcher starts its bound before evaluation/preparation), with
  at most 64 pending/in-progress entries per registry. Expected UI expiry dismisses
  without selection; completion is not activated after expiry. Native mode has no
  Lua callbacks or renewable authenticated handoff; its foreground lifetime is
  parent/origin-bound, not a 30-second interaction timer. Lua expiry is lazy on
  registry operations; backwards wall-clock changes invalidate the registry.
  Private pipes/files are not isolation from the same OS account, trusted user
  Lua or privileged processes. Capabilities and original tickets must not be
  logged/shared. Callback side effects cannot be rolled back. Rex's opaque
  server identity and connection-generation semantics remain unexercised.
- **L5 — terminal/geometry:** the targeted native frame/navigation observations
  above are not a direct-I/O, palette-reply, grapheme-clipping, resize, five-position
  or native-layer geometry matrix. Fonts, widths, Unicode presentation,
  palette/contrast across themes and tiny viewports are not generally certified.
  ASCII/no-color flags exist in low-level menu/popup modes but are not forwarded
  by the delivered launchers. Nonempty NO_COLOR is honored in the rendering
  child's environment; the launcher does not forward a caller-side environment
  override to Rex's new terminal. See [appearance](appearance.md).
- **L6 — inventory/deferred remote scope:** listing reads the execution host's
  effective Rex home inventory, not inventory selected by popup `--server`.
  Rex may contact home and forward held changes. Empty/failed/stale listings and
  endpoints are not reachability results. Rows cannot execute/attach or retarget
  the popup. No independent registry, remote attachment/connection, bootstrap,
  Wake-on-LAN, SSH, retries or reachability protocol is delivered.
- **L7 — packaging/user files:** package recipes have `doCheck = false` and
  `doInstallCheck = false` under current policy, not passing package tests.
  Home Manager defaults disabled, links only the reserved namespace and never
  forces unmanaged files. Existing-file preservation, conflicts and enabled/
  disabled activation have not been exercised. Whole-directory Rex symlinks need
  user-owned layout integration. Removing the module does not remove user-added
  require/setup wiring; users must account for that themselves.
- **L8 — delivery is not acceptance:** recorded reviewed/integrated tasks,
  coordinator compilation and this inventory do not assert a working installed
  release. Reviews and compilation apply to their recorded frozen revisions;
  later edits require fresh evidence. The narrow native runtime observations
  above are not general acceptance, a coverage claim or authority to execute a
  final test matrix. Only the bounded coordinator Phase08/09/10/11 gates described
  above are authorized. REQ-024/025 partial delivery is not full Phase11/project
  completion while REQ-026 remains blocked.
- **L9 — ordinary launch/native defaults:** the earlier producer revisions have
  recorded review/compilation/integration, but a later normal invocation failed
  for absent custom registrations. Task `07-05` supplies self-contained native
  defaults and records a normal native popup/navigation/dismissal pass at its
  frozen revision. The coordinator's Phase08 idle/default/color gate above must
  not be replaced by that earlier result, renderer-only evidence or compilation.
  A reachable already-running default server, valid execution-host Rex runtime/home settings and compatible CLI identity JSON
  are required for native discovery. Present invalid overrides fail rather than
  falling through; controller handoffs never default. Terminal and server-side
  preparation must resolve the same canonical endpoint for custom/shared setups.
  Full captured session/block are mandatory for display. An exact/unique eligible
  attached app with unchanged generation is mandatory for native selected-action
  execution or configured launch, not native display/dismissal (REQ-023). No
  focus/current-client/server substitution is authorized. No-argument configured
  activation may be accepted without popup display, and uncertainty is not retried. Recorded documentation review,
  integration and Go/Nix compilation do not convert skipped tests/runtime gates
  into passing usability criteria.
- **L10 — native controls capability:** REQ-026 remains blocked, not deferred or
  implemented. Inspected creation/resize schemas, terminal creator options,
  connected-app registry and published Lua layout docs expose no supported
  per-popup pane/tab/window controls-hiding option. No undocumented flags,
  global workaround, settings/font changes or Rex source changes are delivered.
  A supported Rex capability/update or explicit user deferral is required;
  terminal borders, compact bounds and server captures are not substitutes.

## Usage and published contracts

The [README](../README.md) covers basic startup without custom registration,
installation layout, optional user-owned Lua registration, imported/inline
configuration, default keys and troubleshooting.
[launch-contract.md](launch-contract.md) details ordinary no-argument popup launch,
exact client/origin selection and native/default/custom endpoint routing;
[appearance.md](appearance.md) describes semantic colors, palette fallback and
accessibility; [mapping-format.md](mapping-format.md) is the Go-owned shared standard;
[remote-compatibility.md](remote-compatibility.md) details endpoint/runtime identity,
server-local execution and cleanup boundaries. The
[Home Manager examples](../tests/packaging/home-manager/README.md) are source-review
cases, not an activation runner.

Published Rex references used by the implementation/documentation:
[CLI](https://www.superlogical.com/rex/docs/automate/cli),
[Server API](https://www.superlogical.com/rex/docs/automate/server-api),
[configuration](https://www.superlogical.com/rex/docs/customize/config),
[custom actions](https://www.superlogical.com/rex/docs/customize/actions),
[Lua API](https://www.superlogical.com/rex/docs/reference/lua) and
[Lua scripts](https://www.superlogical.com/rex/docs/automate/lua-scripts).
These contract references are not live-server observations or new verification
claims.
