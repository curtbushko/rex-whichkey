# Current release evidence: shared Lua float / direct Go

## Phase19 accepted policy and observed make route

Accepted direction `REXWHICHKEY-20261106-232754`: shared Lua owns native float
creation/fractional sizing; one Go process inside it owns menu/defaults/JSON,
providers/rename/confirmation, own-PTY palette/fitting and CLI effects. There is
no Go launching controller, renderer child, copied image, private FIFO/result/
update transport or wrapper supervisor. Bare/root-flag Go launch rejects safely.
Functions/callbacks/executable Lua menu config/hooks remain deferred.

User1a explicitly accepted default native chrome **for now**, deferring REQ-026
native titlebar/border/corners/shadow customization. It is not implemented or
credited complete. Rex's default native header, border, corners and shadow may
remain visible. Lua owns this presentation policy, sizing and launch; Go's menu
header/footer, semantic terminal colors and palette matching are separate, not
native chrome controls. No ignored knobs, invented schema/API fields, global Rex
settings/user config edits or painted borders implement customization.

The [accepted design](../.tasks/designs/rex-whichkey-lua-float-launch.md) records
explicit `1a - yes, accept it for now; 1b - yes` in
`rex-whichkey/lua-float-design/approval/2`; Task04's
`rex-whichkey-p19-t04/chrome-blocker/1/resume/2` retains the original unsupported-
API blocker `REXWHICHKEY-20261106-234819` and its history. Older blocked wording
below is historical, not current policy or feature credit.

User1b permits the detached execution-host Block through dispatch until Go exits
or Rex deletes it. Workers/tty/visible float/exact original focus must still be
cleaned and verified before an effect; uncertainty suppresses effects/no retry.
Go cannot verify its own post-exit deletion. A later cleanup failure cannot
suppress/undo an already-issued effect; exact `list_blocks` observation is needed.

Delivered prerequisites are recorded by native completed Task01/02/03 summaries in
the coordinator collection, not inferred from this documentation:

| Delivered source | Revision-bound observed scope |
| --- | --- |
| Task01 integrated `3b9293334a498d65f459afc2173839d67562e174`; `.mtn-tmp/p19/resume/task1-public-evidence.json` | Shared script/direct root, dynamic metadata, exact owned inline rename, separate configured pane-filter/edit-cancel/root-dismiss runs; visible retirement before effect and observed post-exit host deletion. No page-transition claim from its resize capture. |
| Task02 integrated `75c75d26923573a4b4b3a67f18cafdd3336fa82f`; `.mtn-tmp/p19/resume/task2-public-evidence.json` | Reviewed controller retirement; direct root/provider/own-popup exclusion, edited rename canceled then exact owned rename, real native terminal resize 80×24→60×9 and reflow with fixed `.8/.3` outer bounds, default dismissal, pre-effect detached host/no float while old label remained, eventual exit0/host absence/origin focus/tty restoration. Process executable/ancestry observed no launching Go parent. Resize had no GUI owner, not physical GUI/pixel evidence. |
| Task03 integrated/observed combined and main `6564d07ef690d591ea842fa9d986d0bcfa919dd4`, tree `53a91c6bd08b4ceae15f860492092eee6af6a038`; `.mtn-tmp/p19/resume/task3-summary.json` and `task3-public-evidence.json` | Actual normal source `make run` built current Go and invoked shared Lua; root/palette VT/HTML, nested Sessions full-ID/current-marker inventory, exact owned rename, separate edited cancellation/root dismissal, child exit0, pre-effect visible retirement and eventual exact host absence/focus/terminal restoration. Rebuilt-main default launch/dismissal separately observed. GNU custom-config/bounds observed, closure cause unobserved (below). |

Those coordinator artifacts live in the authoritative checkout; packaged docs do
not install workflow transport. Source equality, build/vet and independent reviews
are bound by native summaries. Same-family independent review is not cross-family
coverage. Earlier failed runs/blockers remain preserved, not silently counted as
success. These are selected cases on Darwin arm64, not a full failure/platform or
provider/action matrix.

Task03 changes the normal source route: `make run` builds this checkout and uses
`rex do ABSOLUTE_ROOT/lua/rex_whichkey.lua --action rex_whichkey_launch` with
explicit originating `-s/-b`, optional original `-C`, and trusted absolute
`development_directory=ABSOLUTE_ROOT/bin/GOOS-GOARCH`. The Lua terminal starts
fixed `rex-whichkey direct`; installed launch omits the directory. Home Manager
still links only the shared module, not init.lua/bindings/JSON. No activation,
reload, dependency/pin or runtime-discovery policy change is implied.

### Task03 actual commands, identities and observations

All artifact paths in this section are under `.mtn-tmp/p19/resume/` in the
authoritative coordinator checkout, not packaged workflow transport. Native
`rex-whichkey-p19-t03/build/1` and `rex-whichkey-p19-t03/review/1` bind the source
producer and fresh same-family review; `task3-summary.json` binds combined
build/vet/source equality and actual runtime. No cross-family review is claimed.

- Combined checkout: `/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t03-integration-a01`;
  actual `make run` resolved to
  `/nix/store/zfhkcz41zg7hks7h4gmz6x7cvcrc6ipn-make/bin/make` (Task wrapper).
  `task3-combined-make-path.txt`, `task3-combined-make.log` and
  `task3-combined-go-executable.txt` bind that route and the current
  `p19-t03-integration-a01/bin/darwin-arm64/rex-whichkey` executable/argv/ancestry,
  not an installed old binary. No launching whichkey parent/copy/IPC was observed.
- Actual GNU entry: `/usr/bin/make run ARGS='config=/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/p19/resume/optional-providers.json width=0.75 height=0.25'`.
  `task3-optional-make.log` records delegation to `task run --` with those data;
  `task3-optional-active.json` records centered bounds `[.125,.375,.75,.25]`.
  Listing-only known-host navigation did not attach/probe a host. The child exited 0,
  but both popup and original disposable window disappeared before planned cancel;
  **cause unobserved**, not a controlled cancellation/dismissal pass. Preserve
  `task3-optional-closure-note.json`; do not infer user or automatic closure.
- A fresh owned-origin PATH `make run -- config=/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/p19/resume/optional-providers.json`
  separately observed edited rename cancellation, unchanged `p19-make-cancel-owned`
  label, root dismissal, child exit0 and exact host/focus/visible cleanup
  (`task3-cancel-typed.txt`, `task3-cancel-process-exit.json`,
  `task3-cancel-after-blocks.json`, `task3-cancel-after-session.json`).
- Rebuilt main: `/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main`;
  actual `make run` displayed the default direct menu, then Escape dismissed it.
  `task3-main-go-executable.txt` identifies `main/bin/darwin-arm64/rex-whichkey`;
  `task3-main-process-exit.json`, `task3-main-after-blocks.json` and
  `task3-main-after-session.json` record exit0, exact host absence and original
  focus. `task3-main-origin-modes.json` and the supplied public evidence record
  primary-screen/terminal restoration.

Observed session: `session:01a1a935-0b64-7c7b-8e83-5d35011bdee2`. These are
historical observation locators, **not IDs to reuse in launch commands**:

| Run | Exact popup host Block | Exact floating Layer |
| --- | --- | --- |
| Combined default / owned rename | `block:01a1b664-a5bf-7340-b5d4-73b14f7ab60b` | `layer:01a1b664-a5c0-77bc-83fe-a98972fad916` |
| GNU custom-config/bounds | `block:01a1b666-a55c-7403-927b-5fd8ca1b4fd4` | `layer:01a1b666-a55d-791c-8e8a-f17dc58a2666` |
| Separate controlled config cancel | `block:01a1b669-b790-7743-9c8b-a25f1854152f` | `layer:01a1b669-b791-7d4a-8df5-b1f24dcfb5e7` |
| Rebuilt-main default dismiss | `block:01a1b66a-d635-7075-9ae4-83de763c5ac3` | `layer:01a1b66a-d636-7645-ad91-78620b7d138e` |

The combined origin was `block:01a1b65d-e397-7a83-a07f-769f9a2daaaf` in
`window:01a1b65d-e399-7ae8-abc3-908b7249e833`, original layer
`layer:01a1b65d-e399-7af0-adc5-4e84d256129c`. Its label changed exactly once
from `p19-source-make-owned` to `p19-owned-make-run` after Enter. In
`task3-combined-order.ndjson`, observations 2–4 show detached host, zero floating
layers and the unchanged old label; observation 5 shows the new label. This proves
observed visible retirement before the effect, **not host deletion before it**.
`task3-combined-process-exit.json` and `task3-combined-after-blocks.json` separately
record exit0 and host absence. Root VT/HTML and inventory/typed captures remain
`task3-combined-root.txt`, `task3-combined-root.html`,
`task3-combined-inventory.txt` and `task3-combined-rename-typed.txt`.

The rebuilt-main origin was `block:01a1b669-b690-735f-ab41-e48f132bd4d8` in
`window:01a1b669-b691-79d3-b24c-e1d2a5a62b0c`, original layer
`layer:01a1b669-b691-79db-9ac7-fb3c14e6a697`. Both default active snapshots show
one centered `.8/.3` float (normalized x≈.1, y=.35). Task03's unchanged Go/Lua
source carries forward Task02's own-PTY resize implementation/evidence; it did
not repeat every resize/palette/action/failure case. Terminal frames are not
native GUI pixel measurements or universal palette/contrast proof.

### Current verification boundary

This Task04 refresh edits documentation only. Go/Lua/Taskfile/Makefile/package
source remains byte-identical to its base `d4d917a4a9c6916d936526b54d887b994074cf71`
and the Task03 observed source above. This leaf ran no product or compilation.
Fresh independent docs/source/evidence review, combined equality/whitespace and
a final actual `make run` on rebuilt main after integration remain coordinator
gates before full phase completion. Task03's actual observations are delivered,
not pending and not a substitute for that final recheck. `rex do`/make launch ACK
is not eventual Go exit, display or cleanup proof; exact native watchers and
`session.list_blocks` establish only the observed cases.

Task04's required gate is `git diff --check`; Nix/Lua were not changed. Task03's
static Nix formatter/syntax parse and coordinator Go build/vet remain separately
recorded, not Nix evaluation/build/install/check/activation. Tests/unit/race/TDD/
Lua/parser/FIFO harnesses/benchmarks/fake servers/synthetic delays were not run.
Protected assertions, quality commands/hooks/dependencies/pins/lint configuration
remain unchanged, as does the parent's separate dirty `flake.nix`.

The current full-lint diagnostic is **53 diagnostics, exit 201; AI lint not reached**
(`current-main-full-lint.log`, `task3-public-evidence.json`). Full lint is not green;
no inherited-only classification is asserted. Phase18's historical 62 below stays
verbatim and is not the current count. There is no new latency statistic, benchmark
or performance SLA. No every-failure/action/platform, physical-keybinding/profile,
install/reload, remote/deploy, phase-completion or native-customization claim
follows from this documentation. Do not use older no-argument/controller commands
below as current usage: see [README](../README.md), [launch](launch-contract.md),
[package/discovery](runtime-bootstrap.md) and [appearance](appearance.md).

## Historical Phase16/18 delivery record (verbatim)

The following original inventory is preserved verbatim, including then-current
commands, pending statements and old lifecycle/chrome decisions. It describes
those revisions, not the active Phase19 route or accepted refinements above.

# Release evidence: CLI-only delivery and async startup

## Verified delivery and captured async evidence

Accepted CLI-only direction `REXWHICHKEY-20261106-114935`: Go owns data,
providers, menu/lifecycle and Rex CLI operations; Lua only starts the binary.
Functions/callbacks/hooks remain deferred, not credited as REQ-005 completion.
Accepted startup refinement `REXWHICHKEY-20261106-191113` permits PATH-only
executables, draw-first async/on-demand metadata and invocation query reuse.
Neither document intent nor a build is runtime proof.

**Phase16 is complete**, not waiting on a package/launcher cutover. Final main
`9fa9310b1398672fdc0f78823d96cff0f6ec0292` was built and run through the
ordinary no-argument public path on Darwin arm64. Evidence:
`.mtn-tmp/p16/final-main-public-runtime-evidence.json` in the coordinator
checkout, plus its linked `.mtn-tmp/p16/task3-public-runtime-evidence.json`.
Observed root, full-ID/current-marker Switch inventory, navigation, Escape,
exit0, terminal restoration, exact popup/private cleanup and origin focus.
An exact disposable owned tab was renamed to `p16-owned-cli-rename` in the
linked owned-effect run. Earlier task2 observation at `7bdbda2` remains valid
historical prerequisite evidence, not the final Phase16 boundary.

Phase18 task3 source `d60c976dcc21fea03a00dff0075491360e88ffcb` (tree
`e354bb8009f3d66e7e23222c9a613438d769211c`) received APPROVED independent
same-family source review `rex-whichkey-p18-t03/review/1`. The coordinator built
and observed combined candidate `397ce9a1591bab637c89978fa7670c023bea1111`;
`task3-public-evidence.json` and native historical event
`rex-whichkey-p18-t03/runtime-observed/1` index that actual evidence, not a new
approval or final-main completion. Later documentation-only revisions preserve
all reviewed Go bytes. The [task](../.tasks/tasks/rex-whichkey-18-03-async-inventories.md)
and [phase](../.tasks/phases/rex-whichkey-18-startup-speed.md) native delivery
summaries bind the exact final integrated head/tree, documentation review and
final normal-public recheck; this inventory does not assert those later results.

Phase17's independent audit at `b1902ca` is a baseline-bound recommendations
report, not authority for unrelated cleanup in this task.

## Phase18 delivered prerequisites

Artifacts below are coordinator-owned under `.mtn-tmp/p18/`:

| Revision/evidence | Actually observed scope |
| --- | --- |
| Baseline `9634246b7e4edbc8bed59210466cac492ddf2f0f`; `baseline-public-evidence.json` | Ordinary public root, later full-ID/current-marker Switch inventory and owned rename form (cancel only), dismissal/cleanup. Coarse observed first-visible-frame upper bound about **2.74s**, not exact rendering latency. |
| Task1 `cc3a320dbbf3ef05bbb959f71feed4f3dae1bced`; `task1-public-evidence.json` | Reviewed/integrated PATH-only path; actual public root/inventory/owned rename form cancel/dismissal/exact cleanup. Coarse upper bound about **0.80s**; REX_BIN_DIR removal was intentional PATH-policy verification. Inventories/labels were still eager. |
| Task2 `6c5c9ed67c8adde938a620cd86d55f777f665b18`; `task2-public-evidence.json` | Reviewed/integrated incremental rename path; actual loading form with Enter disabled and user edits, exact text preserved on ready, cancel without rename, revisit/one exact owned tab rename, separate dismissal, exit0/origin focus/exact popup/private cleanup. Inventories remained eager. |

Task2's native completion summary is committed at `17eb2f2`, the task3 source
base. These earlier timing observations were not a controlled baseline/final
pair; task1 intentionally changed REX_BIN_DIR. The comparable task3 pair below
is separate. All observations include CLI capture latency and poll jitter;
a previous negative observation is not a strict lower bound. There is **no
approved numeric SLA**. No synthetic delays or fake APIs manufactured loading.

## Current async source acceptance map

Task3 implements the remaining production consumers, not a new service. The
map below describes exact reviewed source `d60c976dcc21fea03a00dff0075491360e88ffcb`
under review/1; it does not claim a second documentation review has passed:

| Criterion | Source boundary |
| --- | --- |
| First draw before nonessential waits | `native.go` captures required endpoint/origin/app intent and a matching session view; no eager provider expansion remains. `terminal.go` writes the validated static root before demanding sessions or starting standalone jobs. `popup_initial.go` sizes static/loading groups once. |
| Scheduling and terminal reuse | `menu_updates.go` demands visible enabled sessions after successful first draw; hidden sessions/windows/panes/hosts on enabled prefix entry. `popup_updates.go` validates fixed demands; `providers.go` runs at most two metadata jobs, five once-only kinds, two result slots. All markers share a kind result/error, with no reopen retry. Disabled/absent markers do no work. |
| Query reuse versus freshness | `rex/snapshot.go` detached lists/lookups serve captured placement, window/pane rows and rename labels. Prior presentation-only Session + Windows/Blocks inspect pairs are removed. Sessions list and hosts each load once. Fresh Inspect/native-origin/app-generation/destination/pre-create geometry/liveness/cleanup checks remain uncached. Built-in actions are read only after selection. |
| Parent authority and atomic UI | `popup_updates.go` retains complete mappings/args/pane targets before writing ready bytes. Provider selections carry only keys/confirmation/invocation-kind-load binding. `menu_updates.go` validates complete batches, then rebinds the same prefix/page/confirmation/form state; only the UI loop writes the terminal. Partial decoding and edited value/cursor survive; pending/error rows cannot execute or replay Enter. |
| Bounds and cleanup | 512 resources, existing child/key-byte/4-MiB and known-host bounds remain; overflow/conflict fails, not truncates. No navigation/update resize. Jobs and transport cancel/join on exits; controller/root errors propagate before exact popup/focus/files/image retirement and effects. All twelve update/demand slots are in fixed cleanup. Standalone remains result-only with demand-only hosts, not fabricated native targets. |

The source map and review are distinct from the bounded observations below.
Required leaf gofmt/whitespace, normal `task build` and `task lint:vet` results
bind the exact committed head/tree in each build payload. Full lint on build1
reported **62 inherited failures; AI lint was not reached**. This docs-only
refresh does not repeat full lint or claim it passed; Go bytes remain identical.
No lint/test/Taskfile/hooks/dependency standard changed. No tests were run.

## Captured comparable public runs and optional providers

Coordinator artifacts are in `.mtn-tmp/p18/` in the authoritative checkout.
`task3-public-evidence.json`, `paired-baseline-timing.json` and
`paired-final-timing.json` bind baseline
`9634246b7e4edbc8bed59210466cac492ddf2f0f` versus observed combined candidate
`397ce9a1591bab637c89978fa7670c023bea1111`. Both used ordinary public
`./bin/darwin-arm64/rex-whichkey` with no arguments: same execution host,
PATH, REX_BIN_DIR, Rex server/session environment, built-in configuration and
normal GUI intent, with comparable owned origins and terminal context.
No heavy build ran concurrently with timing.

| Captured upper bound | Baseline | Candidate |
| --- | --- | --- |
| Launch to first visible capture | 2.69s | 0.54s |
| Inventory visible since launch | 2.79s | 0.63s |
| Inventory demand to visible | ~0.07s | ~0.07s |
| Rename demand to ready visible | ~0.09s | ~0.09s |

These are coarse wall-clock CLI/capture/poll upper bounds including jitter and
explicit demand/navigation, **not exact render or internal data-load timestamps,
a benchmark, an SLA or broad performance certification**. Baseline eager data
was ready before first draw by source; the inventory capture does not time that
internal readiness. Both runs showed root, dynamic full IDs/current marker,
ready inline rename, form cancellation/root dismissal, exit0, terminal/origin
restoration and exact popup/private cleanup. The ready captures do not establish
that every loading/error/cancellation branch occurred.

A separate public `--config` run used only coordinator-owned
`.mtn-tmp/p18/optional-providers.json`, without user configuration edits/reload.
The `optional-*` captures record:

- Windows explicitly loading with choices disabled, then ready; Tab reached
  page2 with the correct full owned current window ID. Panes full IDs, visible
  Sessions and listing-only known-host inventory/home/warnings groups appeared.
  No host connection, probe or remote operation was performed.
- The disabled key stayed at root. The hidden prefix opened with intentionally
  invisible inherited rows (`optional-hidden-ready.txt`: no visible mappings),
  not a visible hidden inventory or verified hidden-destination dispatch.
  Zero unused/disabled loads and same-kind reuse are source-review conclusions,
  not inferred from the screen alone.
- All provider loading/navigation retained identical floating bounds. Inline
  rename via the configured key retained entered `p18-owned-async-rename` and
  applied it to exact disposable window
  `window:01a1b3d2-b2c8-7e4c-b9c5-ee01aaf83a52`; `optional-after-session.json`
  records that exact ID/label. Exit0, primary terminal/origin restoration,
  exact floating/private cleanup and owned-window closure were captured.
  Earlier task2, separately, captured edits while loading and their exact
  preservation on ready; the optional run is not substituted for that evidence.

Source review establishes once-only dispatch and cleanup-before-effect ordering;
these bounded resource observations do not certify an exhaustive action/error
matrix. The native task/phase summaries link subsequent exact final delivery
rather than treating this observed candidate as final main.

## Residual limits and authoritative guides

No leaf live app/Rex/timing run, unit/race/TDD/Lua/parser/FIFO harness, benchmark,
fake server/delay, remote trial, user-config/reload, Nix eval/build/profile/input
update, deployment or push is part of this delivery. Physical keybinding,
installed-profile, Linux/remote and all-action/provider coverage are not claimed.
REQ-026 native controls remain unresolved. Phase15's original bridge failure is
not retroactively fixed. Fresh checks and effects are not atomic; uncertain
effects are never retried. Captured candidate evidence is not final-main,
installation or physical-keybinding evidence. Not every cancellation, error,
malformed/replay/disconnect/overflow or platform path was exercised; no synthetic
delay or harness was used to force them. Final integration/public recheck and
fresh documentation review are coordinator-owned, linked by the native summaries
above; neither phase completion nor review/2 approval is asserted here.

Current contracts: [data/actions/providers](mapping-format.md),
[launch/cleanup](launch-contract.md), [discovery/package](runtime-bootstrap.md),
[CLI/GUI](cli-gui-delegation.md), [remote limits](remote-compatibility.md),
[appearance](appearance.md), [rename](tab-rename-input.md),
[session UI](session-owned-ui.md), [migration examples](../examples/README.md).

## Historical delivery records

The original inventory below is preserved verbatim. All claims are scoped to
their recorded revisions. Old bridge setup, callback contracts, removed menus,
test runners and outstanding-cutover statements are historical, not active
Phase16 prerequisites or current verification. Current sections above take
precedence; the historical PRD and .tasks history are unchanged.

<details>
<summary>Historical pre-Phase16 delivery inventory and evidence</summary>

# Release delivery inventory and limitations

## Latest user correction

The user confirmed session rename (`wr`) and tab rename (`tr`) work, stopped
further testing, and requested removal of numbered “Go to tab” entries (`t1`–`t9`).
Both default menus now have 16 executable actions; the unused native numbered-tab
adapter was retired. Earlier all-25 counts and numbered-tab delivery sections
below are historical. See [current action status](actions.md#current-menu-and-user-confirmation)
for coordinator evidence versus user confirmation. No new runtime verification or
whole-project completion is claimed.

The latest app-routing audit removed the shared app chooser entirely. Every
native app action and configured CLI launch now uses the launching pane's exact
GUI owner unless explicitly overridden; resource closes remain app-independent.
This supersedes the historical chooser/unique-app policies below. App generation,
attachment, placement and owner are validated before the first effect; no labels,
resize reports or connection-list order select an app. New Tab retains its bound
app across its own creation/activation effects. Native dialogs and pickers remain
unchanged. These routing changes are compiled, not newly runtime-certified.

The user approved **REXWHICHKEY-20261103-085109**, a built-in Tabs → Rename
text-entry state within the same fixed popup. On confirmation and full cleanup,
Go renames the captured tab directly; no app picker or native tab rename dialog
is used for this native-menu prototype. Configured Lua behavior remains unchanged.
See [design and scope](tab-rename-input.md). Build/source checks are not a runtime
certificate; no new UI trials were run.

The user additionally requested our own session switch/rename UI. In the built-in
menu, Switch expands a fresh session-list snapshot into canonical retained-ID
choices within the same popup; Rename reuses the text-entry state and then typed
`Session.Rename`. Session selection revalidates its exact destination and uses
unscoped `session.select` without opening Rex's picker. Configured Lua menus stay
unchanged. See [session UI scope](session-owned-ui.md). No new runtime evidence is
claimed under the user's stop-testing instruction.

The user identified the remaining bar as Rex's **native floating-pane header**:
it shows `rex-whichkey`, `~` and a resize button. The user reports **Show pane
headers is already off**. This is not our rendered menu heading/breadcrumb, and
no screenshot is required to repeat that distinction; image drops are not
supported in their Rex environment. Do not remove our menu heading or suggest
turning the already-disabled preference off as a fix.

The documented [pane-header preference](https://www.superlogical.com/rex/docs/reference/mac-settings#setting-panes-show-pane-headers)
does not explain the floating-header behavior by itself. Our creation request
uses `session.new_layer` and does not request headers. Fresh installed metadata
exposes only bounds/focus/layout/session/window fields on creation and no
header/presentation control in terminal creator schemas or either app's actions.
The newer published `session.set_layer_state` only minimizes/restores layers;
it is not exposed by this installed server and is not a header-suppression API.
App metadata currently reports work build1758 and m4-pro build1670; no version
cause or client rendering bug is certified. Separate floating-layer chrome is
the likely explanation, pending Rex-side confirmation. Popup-only suppression
remains unresolved; no settings change or UI trial was performed.

## Fixed navigation bounds: REQ-029

Accepted **REXWHICHKEY-20261102-143748** under the
[Phase13 design](../.tasks/designs/rex-whichkey-fixed-navigation.md) supersedes
only navigation-driven resizing in REQ-025. Initial compact geometry/seed and
REQ-027 spacers remain. This section describes source, not a review or runtime
pass; whole-project completion remains blocked by native chrome REQ-026.

| Requirement | Source behavior | Verification boundary |
| --- | --- | --- |
| REQ-029 — one rectangle through navigation | [`popup.go`](../popup.go) creates the owned layer once with precomputed bounds/terminal seed, then watches only resource liveness and configured app generation. No popup production resize, preferred-row observer or hint files remain; `popup_size.go` is retired into ignored `.trash`. [`terminal.go`](../terminal.go) keeps public `RunTerminal`, focus-in/SIGWINCH redraw and existing renderer pagination. Popup content is top-aligned using a local config copy; standalone positioning is unchanged. [`popup_initial.go`](../popup_initial.go) reserves two horizontal/four vertical cells plus half-row rounding, not an all-submenu maximum or font-pixel guess. | Fresh frozen review and coordinator compilation/observations belong to [task 13-01](../.tasks/tasks/rex-whichkey-13-01-fixed-navigation.md). Leaves run only source/format/whitespace checks. No fixed-bounds runtime pass is asserted here. |

The coordinator must promptly retain first bounds and startup terminal seed,
then verify the **exact same rectangle** at root, Tabs submenu, paging when
available, back and idle. Retain root status/content spacers, Escape exit0 and
exact owned-resource/private-file cleanup before any selected dispatch. External
native window/font changes may naturally change terminal cells; redraw adapts
content but the application must not resize the layer. No font/settings change,
project tests/TDD/general/unit/race suite, destructive/native selection, remote
or Darwin trial is authorized for leaves.

The historical Phase12 result at
`.mtn-tmp/popup-initial-fit/runtime/rex-whichkey-p12-t01-builder-a02/result.json`
binds candidate `caff550e0e7079064f0d34d18abc32785cdfa4e3` / tree
`071100e2621621670de37b6d077ee3f29dd68ea6` to first H=0.1104651163 and
later root H=0.1311773256, with a 97×7 startup seed, observed status spacers,
Tabs/back and Escape cleanup. It records the former live corrections, **not**
REQ-029 fixed bounds. The added vertical cell headroom is a conservative estimate,
not a cross-invocation calibration cache or a measured new-pane guarantee.

## First-frame fit and status spacing: REQ-027/028

Accepted **REXWHICHKEY-20261102-141047** under the
[Phase12 design](../.tasks/designs/rex-whichkey-first-frame-fit.md) refines startup
and status spacing without waiving REQ-026. This section describes source and
required evidence, not an inferred review, compilation, runtime or project pass.

| Requirement | Source behavior | Verification boundary |
| --- | --- | --- |
| REQ-027 — status/content spacers | [`layout.go`](../layout.go)/[`render.go`](../render.go) include two ideal spacer rows in natural preferred height and render one beside each status border when room permits. Spacers yield before content/essential hints; no single-page counter or extra footer returns. | Source-only formatting/inspection is not a runtime matrix. Revision-bound review and coordinator frame observations belong to [task 12-01](../.tasks/tasks/rex-whichkey-12-01-first-frame-fit.md). |
| REQ-028 — compact first creation and terminal seed | [`popup_initial.go`](../popup_initial.go) computes natural root rows and content-sized anchored bounds before creation, using validated origin geometry or the documented compact fallback. [`rex/geometry.go`](../rex/geometry.go) consumes window-relative leaf rectangles directly, handles documented zoom, and checks placement around measurement. [`rex/layer.go`](../rex/layer.go) seeds shell startup size in that single creation RPC and rejects stale placement. Phase13 removes later fitting and retains the initial rectangle. | Initial cell-level inset/rounding is an estimate where new-pane font/chrome differs. The fallback is hypothetical, never viewer provenance. Exact frozen revision review, combined compilation and promptly captured first-layer/terminal geometry belong to task 12-01 and [Phase12](../.tasks/phases/rex-whichkey-12-first-frame-fit.md); no runtime pass is asserted here. |

The original Phase12 coordinator gate required sampling the **first** owned layer
promptly, before its then-active one-second resize watcher, and retaining bounds
plus initial startup terminal rows/columns. A final-small-only observation could
not accept REQ-028. The current Phase13 gate above requires identical bounds, not
settling through live corrections. Retain root/spacer/status frames, submenu/back,
Escape exit0 and exact owned-resource/private-file cleanup as well. No general tests/TDD/unit/race suite,
font/settings/config reload, native/destructive selection, remote trial or GUI
pixel experiment is authorized for leaves. Leaves perform source/format/whitespace
checks only; coordinator observations must identify the actual frozen revision.

See [appearance](appearance.md#actual-layer-height-and-measurement-limits) and
[initial geometry contract](launch-contract.md#initial-popup-geometry-and-terminal-seed)
for the 100×40 hypothetical fallback grid/30% startup cap, current two-horizontal/
four-vertical-cell inset estimate, half-row rounding and placement-race limits.
Later measured fitting is retired under REQ-029. Historical Phase11 settled-size
observations below are not first-frame or spacer evidence.
Server bounds/seed/frame observations do not certify native GUI chrome removal.
REQ-026 remains **BLOCKED**, and Phase11/project completion remains unavailable.

## Compact popup presentation: REQ-024/025 and blocked REQ-026

Accepted **REXWHICHKEY-20261102-123334** under the
[Phase11 design](../.tasks/designs/rex-whichkey-compact-presentation.md) adds
compact status borders and actual content-fitting layer height. The live fitting
portion below is historical and superseded by REQ-029 above. This is **partial
delivery**: hiding native Rex pane/tab/window controls is an accepted requirement,
not fulfilled or waived. Phase11/project completion remains blocked by REQ-026.

| Requirement | Source / capability state | Evidence boundary |
| --- | --- | --- |
| REQ-024 — compact status borders / conditional paging | [`render.go`](../render.go)/[`layout.go`](../layout.go) put title/breadcrumbs and key hints in separate full-width horizontal status borders, with no generated `Keys:` prefix, single-page counter or redundant page-footer row. Genuine multipage counters/Tab hints share the bottom border where width permits. Tiny heights yield status rows to content; critical confirmation text remains on the content row. `MenuFrame.PreferredHeight` reports natural current-width content plus both borders independently of clipping/page, now including the REQ-027 ideal spacers. | [Task 11-01](../.tasks/tasks/rex-whichkey-11-01-render-frame.md) records fresh APPROVED source review and coordinator Go compilation integrated at `f57dfa00` (tree `0d1a7aa375bc8b099b2100a344b1edf6ff4ef9dd`). Tiny/multipage/grapheme/ASCII/no-color paths are source-reviewed, not a runtime matrix. |
| REQ-025 — actual content-fitting owned layer | Historical Phase11 implementation published private natural-row hints and resized the exact owned floating layer using measured geometry. REQ-029 retires that popup feedback path and hint transport; current source keeps initial content-sized bounds and pagination instead. Public Rex geometry/resize APIs remain intact. | [Task 11-02](../.tasks/tasks/rex-whichkey-11-02-popup-height.md) records fresh APPROVED source review, combined Go compilation/`task build`, integration at `3c68b784bc0c90be94b0c4604040406f603557d3` (tree `7f6054689d0b505d5ed951b6c7f3e87b61eee4f0`) and the actual one-context coordinator gate below. Font-change generality, all positions and failure/transient paths are not runtime-certified. |
| REQ-026 — hide native controls for this popup only | **BLOCKED / not implemented.** Inspected Rex schemas/actions/docs provide no documented per-popup pane/tab/window control-hiding capability. Terminal text borders and resized bounds do not hide native chrome. | [Task 11-03](../.tasks/tasks/rex-whichkey-11-03-popup-chrome.md) owns blocker `rex-whichkey/presentation/native-chrome-blocker/1`. No hidden flag/global setting workaround or explicit user deferral is accepted. A supported Rex capability/update or explicit deferral is needed; source delivery cannot complete the phase/project. |

[Appearance](appearance.md) explains status-row priorities, semantic colors and
measurement limits. The current [launch contract](launch-contract.md#fixed-navigation-bounds-and-liveness)
documents fixed initial bounds and removal of the former private geometry IPC.
The following grow/shrink gate is historical evidence only. Initial geometry uses
validated origin rows/bounds and optional cell pixels, not guessed fonts; missing
measurements use the bounded fallback. Native display still needs no app lookup,
and native action binding remains deferred until cleanup. FIFO parent lifetime,
configured original-VM authenticated 30-second expiry, confirmation and no uncertain retry
are unchanged. No font/settings/init/config reload is part of this delivery.

### Historical producer gate, not fixed-navigation evidence

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
the historical revision-bound documentation review, combined Go/Taskfile
compilation and bounded root/nested-submenu/back assembly gate. That gate's
settling/grow-shrink criterion is superseded by Phase13's exact fixed-bounds
criterion. Earlier producer results do not preclaim a later assembly pass;
neither review nor compilation substitutes for observed layer geometry.

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

## Native action-time target choice: REQ-031

Accepted **REXWHICHKEY-20261102-151552** in the
[action-audit design](../.tasks/designs/rex-whichkey-action-audit.md) adds explicit
late app choice, not a new display prerequisite or a submenu resizing loop.
The later accepted **REXWHICHKEY-20261102-153808** in the
[selected-window supplement](../.tasks/designs/rex-whichkey-selected-window-splits.md)
refines split destination only: the chosen app's current selection replaces forced
launch-origin context for native `pane.split`.

| Requirement | Implemented source | Evidence boundary |
| --- | --- | --- |
| REQ-031 — explicit late native app choice | [`native.go`](../native.go) captures the first validated result and returns from popup cleanup before resolving apps. [`native_targets.go`](../native_targets.go) retains eligible full IDs/generations and routes multiple apps through a separate canonical menu; safe label hints never grant authority. An explicit selector stays strict, zero apps refuse, one binds directly, chooser cancellation dispatches nothing. Only the target identity changes; original mapping/args/confirmation are revalidated, app metadata and exact generation/origin checked before one dispatch. Outer signal ownership spans both UIs and gaps; the dispatch budget starts after choice/cleanup. Both stages retain native FIFO lifetime, initial sizing/spacers and fixed navigation bounds. Configured original-VM binding/expiry/auth remain unchanged. | Source-only implementation, not compilation or runtime proof. Fresh frozen review, coordinator compilation, actual bare `s` -> explicit work-app choice -> below-origin split, chooser cancellation with no pane, strict invalid selector refusal and exact UI/file/lease cleanup belong to [task 14-01](../.tasks/tasks/rex-whichkey-14-01-action-target.md). No such gate is claimed by this source entry. |
| REQ-032 — native splits use chosen app's selected window | [`native.go`](../native.go) sends native `pane.split` through documented `client.perform` using a separate exact-endpoint, unscoped client; only splits omit origin session/block context. No session/block override or server active-window lookup is used; retained `right`/`down` args and selection/confirmation survive unchanged. The original scoped client still rechecks origin, session attachment and exact app generation immediately before one dispatch. Other native/destructive actions, configured routing, target chooser and cleanup remain unchanged. | Source-only change, not compilation or behavioral proof. [Task 14-04](../.tasks/tasks/rex-whichkey-14-04-selected-window-splits.md) requires coordinator compilation and owned launcher A versus selected tab B with two known panes: both `s`/`v` must split B's focused pane in the expected direction, leave A and B's other anchor unchanged, and clean exact fixtures. Ordinary same-window splits, late choice/cancel and strict missing selector also require verification. No runtime pass is asserted here. |

The coordinator's recorded reproduction in the design identifies several attached
apps as the bare split-down failure; explicit app dispatch produced the expected
split. Native `pane.split` takes `direction: "down"`, not the CLI spelling `below`;
this repair intentionally does not alter that correct default mapping.

REQ-030's all-25-default audit is separate from this routing repair. Only the
coordinator may exercise authorized disposable fixtures and record actual
resource before/after effects. The selected-window supplement records six observed
resource effects (`v`, `s`, `x`, `tx`, `wn`, `wx`), not an all-actions pass or proof
of the revised split destination. The earlier assumed 18-resource/7-GUI partition
was disproved by actual probes: that snapshot left 19 native GUI outcomes
uncertified. The later numbered-tabs supplement records functional evidence for
`th` and `e`, but not an all-25 pass. Tab navigation did not update server
`active_window_id`; at that snapshot, `tn` returned success without an observed
new server window/session and remained unverified. The later explicit-dispatch
comparison below establishes a specific CLI/API effect difference, not an
all-action diagnosis. Source/registry acceptance or a successful receipt is not
GUI proof.
Native chrome REQ-026 remains independently blocked. Leaves run source/format/
whitespace checks only, with no tests/TDD, builds, native/GUI/API/config/font/
platform experiments or installs.

## Native numbered tabs by exact identity: REQ-030 bounded repair

Accepted **REXWHICHKEY-20261102-165714** in the
[numbered-tabs supplement](../.tasks/designs/rex-whichkey-numbered-tabs.md)
refines verbatim native dispatch only for `t1`–`t9`. The coordinator's recorded
fresh-hidden-target matrix acknowledged all nine numeric goto requests without
observing expected target geometry; explicit window-ID selection and previous/last
positive controls worked. This is not evidence that every native route is broken.

| Requirement | Implemented source | Evidence boundary |
| --- | --- | --- |
| REQ-030 — reliable native numbered tabs | [`native.go`](../native.go) checks the actual `session.select` native registration after original selection, optional target choice and all cleanup. Historical `native_tabs.go` at commit `54a577bee6481ebe5110d7b280c03bcf307542c3` strictly reads parent-retained integer `index` 1–9 matching its `tN` key; existing typed session/window APIs resolve ordered origin-session tabs once and freeze an exact ID. Immediately before one unscoped exact-endpoint `client.perform`, origin placement, chosen app attachment/generation and frozen window membership are rechecked. Action args explicitly name origin session and resolved window; missing/moved targets refuse, reorder never substitutes another ID, and errors/uncertain effects are not retried. Canonical defaults, original completion/confirmation, configured original-VM route, split current-app context, destructive targets and both UI lifecycles remain unchanged. | Source-only implementation; no leaf tests, compilation, RPC or behavioral pass. [Task 14-05](../.tasks/tasks/rex-whichkey-14-05-numbered-tabs.md) requires frozen review and coordinator combined compilation, then actual real-menu `t1`–`t9` against nine fresh hidden targets in exact verified order and a separate visible caller. Observe target-specific app Work geometry after each action, without explicit target selection between dispatch and measurement. Verify missing index, removed/moved target, stale app/origin and cleanup/refusal boundaries. Server focus or action receipts alone are insufficient; no screenshot/pixel certificate or all-25 completion is claimed. |

The translation is parent-owned intent compilation, not a claim to have repaired
Rex's native numeric `client.tab.goto` action. The all-25 action audit remains open;
REQ-026 native chrome remains independently blocked. Leaf validation is limited
to source, formatting and whitespace checks; runtime effects belong to the
coordinator's frozen candidate gate.

## Explicit native client.perform: REQ-030 bounded repair

Accepted **REXWHICHKEY-20261102-171343** in the
[explicit-dispatch supplement](../.tasks/designs/rex-whichkey-client-perform.md)
replaces the remaining native `rex do` transport, not retained action intent or
target authority. The coordinator recorded that the exact same installed Rex
executable returned success/empty output for `rex do client.tab.new` without a
new window across inspected sessions, while direct documented `client.perform`
with the same app/context created one. That earlier comparison was not a reliable
New Tab fix: the a01 candidate compiled and passed source review but **failed** the
actual `tn` resulting-effect gate, with only six of seven required effects.
Tracing showed the original `client.perform client.tab.new {}` and correct retained
app/origin; an interactive zsh caller with valid OSC PWD also created no tab.
The cause remains unknown; neither version skew nor a shell/context diagnosis is
established. The failed candidate and its artifacts remain failed, not certified.

| Requirement | Implemented source | Evidence boundary |
| --- | --- | --- |
| REQ-030 — explicit native dispatch | Except for numbered-tab and New Tab intent compilation, [`native.go`](../native.go) sends original action and `json.RawMessage` args through one documented `client.perform` using a separate unscoped exact-server client and the existing fresh three-second signal-cancellable budget. All non-split requests include explicit captured origin `session_id`/`block_id`; splits omit both and use the chosen app's selected window/focused pane. Original selection/confirmation, selected-action registration and exact app generation/origin recheck remain before dispatch, after both UI/file/lease cleanups. `Call(..., nil)` validates JSON before discarding; sanitized uncertainty errors never retry or fall back to CLI `do`. Read-only metadata still uses the Rex executable. The numbered exact-ID adapter, configured original-VM route/auth/expiry, late chooser, FIFO lifetime, no idle timeout and fixed initial dimensions/spacers are unchanged. | Source-only implementation, not compilation or behavior proof. [Task 14-06](../.tasks/tasks/rex-whichkey-14-06-client-perform.md) requires independent frozen source review and coordinator combined compilation, then actual native menu `v`/`s`/`tn`/`x`/`tx`/`wn`/`wx` with all seven resulting resource effects observed on exact owned fixtures. Selected-window split correctness and exact UI/files/fixture cleanup must hold. Destructive closes are limited to owned empty resources. A receipt or source review alone is insufficient; no leaf runtime pass is asserted. |

The earlier six observed resource effects and functional `th`/`e` evidence do
not certify the all-25 default matrix. Full action/API verification remains the
coordinator's responsibility after integration; GUI interactions remain
individually unresolved rather than a blanket screenshot requirement. REQ-026
native chrome remains independently blocked. No leaf builds, tests/TDD, RPC,
native/GUI/API/config/font/platform trials or installs were run for this repair.

## Native New Tab resource intent: REQ-030 refinement

Accepted **REXWHICHKEY-20261102-173105** in the
[New Tab supplement](../.tasks/designs/rex-whichkey-new-tab-intent.md) refines the
failed transport-only `tn` repair without changing the seven-effect acceptance gate.

| Requirement | Implemented source | Evidence boundary |
| --- | --- | --- |
| REQ-030 — one source-session New Tab | [`native.go`](../native.go) strictly validates Go-retained `tn` / `client.tab.new` intent with absent or empty-object args. The chosen app's actual `session.select` registration is checked before mutation. After both UI/file/lease cleanups, existing typed `Session.NewWindow` creates one default terminal in the captured source session following an origin/app-generation recheck. The returned exact ID is frozen; registration, exact window membership and original origin/app generation are checked before one unscoped exact-server `client.perform session.select` with explicit source session/new window args. All work shares the existing signal-cancellable dispatch budget. Uncertain creation refuses without retry; any known created ID survives response/recheck/activation errors with sanitized manual guidance. No duplicate create, automatic close, retry or original-action fallback exists. Configured Lua, numbered tabs, destructive targets and split context remain unchanged. This is Go intent translation, not execution of Rex's native `client.tab.new`. | Source-only change, not compiled or runtime-verified by this leaf. Fresh frozen review and coordinator combined compilation are required, followed by actual menu `v`/`s`/`tn`/`x`/`tx`/`wn`/`wx` and all seven resulting resource effects. Real `tn` must create exactly one tab in its captured own session and activate that exact ID; verify cleanup before creation and no duplicate/rollback on failure. The earlier a01 gate remains failed. All-25 modal/action verification and REQ-026 chrome remain incomplete. |

## Native display independent of app targeting: REQ-023

Accepted **REXWHICHKEY-20261102-111430** under the
[Phase10 design](../.tasks/designs/rex-whichkey-app-independent-popup.md) refines
native launch-time app requirements without weakening native action safety or
configured original-VM binding. This is separate from the Phase09 tab-return work.

| Requirement | Implemented source | Evidence boundary |
| --- | --- | --- |
| REQ-023 — app-independent native display | [`launch.go`](../launch.go) validates syntax and captured endpoint/session/block, then reads server registration before GUI binding. [`native.go`](../native.go) retains immutable canonical defaults and optional requested selector without app lookup. [`popup.go`](../popup.go) uses resource-only native creation/wait/focus, deadline-free interaction and FIFO parent lease; owned cleanup needs no app and precedes selected-action targeting. Cancellation needs no app. Executable retained selection/confirmation resolves an exact requested app, a unique eligible app, or the explicit action-time choice described under REQ-031; it checks only the selected action and rechecks full ID/generation/origin before one dispatch. Configured handoffs retain prelaunch binding, strict decoding, original authenticated expiry and same-VM completion. | Source trace and formatting do not prove compilation or runtime behavior. Earlier single-app gates do not verify the deferred boundary. Exact revision-bound review, coordinator compilation and runtime artifacts belong to [task 10-01](../.tasks/tasks/rex-whichkey-10-01-app-independent-popup.md) and its [phase](../.tasks/phases/rex-whichkey-10-app-independent-popup.md). No runtime pass is asserted here. |

Native display/navigation/Escape do not require a live requested client, a singleton
app, known attachment kinds or any selected-client action registry. Malformed
selector syntax still fails early. Missing/unavailable app, invalid explicit
selector or unsupported selected action refuses **action targeting after popup
cleanup**. Multiple apps without a selector now offer explicit choice (REQ-031),
never focus/first/last/terminal-owner inference or retry of uncertainty. Configured
launch retains its live app prerequisite. Resource creation is not proof that a
GUI showed the menu.

The earlier Phase10 coordinator-only runtime scope was normal startup/navigation/
Escape, valid-but-nonexistent selector startup/Escape, then one non-destructive selection
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
| REQ-008 — geometry/rendering | [`layout.go`](../layout.go)/[`render.go`](../render.go) constrain cell-based columns, paging, grapheme-safe clipping, breadcrumbs/icons/help; [`popup.go`](../popup.go) selects native top/bottom/left/right/center maximum bounds; [`popup_initial.go`](../popup_initial.go) computes compact initial bounds/terminal seed (REQ-028), retained unchanged through navigation (REQ-029; former live REQ-025 fitting is retired). [`terminal.go`](../terminal.go) consumes actual frame placements/Lines and passes Pages to key handling. Historical reviews 04-02/04-04/04-05; Phase11 refinement above. | One-context compact layer bounds are observed above; no five-position, clipping, resize/font-change or tiny-terminal geometry matrix is certified. No old full-window overlay workaround is adopted. L1/L5. |
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
  Configured mode binds app generation before activation and still refuses
  ambiguous app selection. Native ambiguity offers the explicit REQ-031 chooser;
  missing/lost/changed required identities fail closed at their respective boundary.
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
  final test matrix. Bounded coordinator gates include the historical
  Phase08/09/10/11 gates above and Phase14's explicitly authorized disposable-only
  action audit; none authorizes leaf runtime experiments or user-resource changes.
  REQ-024/025 partial delivery is not full Phase11/project completion while REQ-026
  remains blocked.
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
  Full captured session/block are mandatory for display. An exact eligible
  attached app with unchanged generation is mandatory for native selected-action
  execution or configured launch, not native display/dismissal (REQ-023). Native
  selection may bind an explicit chooser result when several apps exist (REQ-031);
  configured launch still needs a selector or unique eligible app. No
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

</details>
