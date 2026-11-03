# Native default action audit

## Current menu and user confirmation

The user subsequently requested removal of all numbered “Go to tab” entries
(`t1`–`t9`). Both Go and Lua defaults now contain **16 executable actions**:
six session actions, five tab actions and five pane/app actions. The unused native
numbered-tab adapter was retired along with those entries. Custom configured Lua
mappings can still invoke Rex actions; no user configuration was changed.

The user confirmed that **session rename (`wr`) and tab rename (`tr`) work**.
This is user-reported verification, separate from the coordinator's earlier
experiments. Of the retained defaults, twelve have recorded coordinator behavior
evidence and two rename actions have user confirmation; `ws` and `c` remain
unverified. The user stopped further testing; no new runtime checks were run for
this menu change.

A subsequent source correction makes `x`, `tx` and `wx` close the exact captured
pane/tab/session directly through existing typed Rex resource operations, without
an app chooser or native action registration. Confirmation and complete popup
cleanup remain mandatory. Origin placement is captured before the menu and
rechecked immediately before closing; explicit app selectors stay strict. This
new route is compiled but **not runtime-tested**, per the user's instruction;
the older native-dispatch resource evidence below predates this route.

Native rename dialogs were not replaced. The proposed whichkey rename prompt
(`REXWHICHKEY-20261102-193858`) was **rejected by the user**. Bare CLI launch
context contains the session and pane IDs but not an invoking GUI app ID. The
user directed use of the associated GUI app ID exposed by the CLI: native `tr`
and `wr` now use the launching pane's documented `session.describe_block` owner
unless an explicit app selector was supplied. The exact owner must be a live
attached graphical app with a valid connection generation. Rename bypasses the
session-wide app chooser, retaining the original session/block action context
and native dialog. App generation, original tab/layer placement and unchanged
owner are checked immediately before dispatch; absent/non-GUI owners, moves or
changes refuse rather than guessing another app. Ownership is routing context,
not input provenance or additional permission. This refines the earlier overly
broad claim that no associated app ID was available. Switch Session (`ws`) now
shares this pane-owner binding and revalidation path, so it opens Rex's session
picker without first opening whichkey's app chooser.

The subsequent full selector audit found the same old chooser in **ten more
native bindings**: `wh`, `wl`, `th`, `tl`, `wn`, `tn`, `v`, `s`, `e` and `c`.
All native app actions now use the same
[`pane-owner binding and validation`](../native_app_owner.go), not an action
allowlist. The app chooser, private selection tag, label formatter and
singleton/list-order selection fallback were removed. Configured CLI launch was
also corrected to derive its default exact binding from the pane owner before
original-VM activation. Existing internal `--client` flags still transport and
validate that known binding; they are not user-selection prompts.

Resource closes remain app-independent, explicit overrides remain strict, and
native dialogs/pickers are unchanged. New Tab freezes its app before creation;
its subsequent activation revalidates that connection and exact resources, not
a potentially changed owner caused by its own creation. These routes are built
but not runtime-tested, per the user's stop-testing instruction.

## Historical 25-action audit

The remainder records the earlier source and evidence, before that explicit menu
change and rename confirmation. Its counts, numbered bindings and outstanding
rename checks are historical, not the current default menu or current status.
Historical numbered-tab implementation is available at commit
`54a577bee6481ebe5110d7b280c03bcf307542c3`.

### Scope and result at that revision

REQ-030 requires behavioral verification of **all 25 executable defaults**.
Current coordinator evidence supports **21 functional outcomes: seven resource
effects and fourteen rendering/geometry effects**. Four interactions remain
unverified: `ws`, `wr`, `tr` and `c`. This is not an all-actions pass. Registry
compatibility, invocation receipts, resulting resources, app rendering and native
GUI pixels are distinct claims; no native pixel or configured original-Lua-VM
callback certificate is supplied here.

[Task 14-03](../.tasks/tasks/rex-whichkey-14-03-gui-action-verification.md) remains
**REQUIRED / IN PROGRESS** for those four individual outcomes, not blocked on
screenshots for 19 actions and not deferred by the user. Native popup chrome
REQ-026 remains independently blocked; see [acceptance status](acceptance.md).
Neither this partial audit nor the documentation delivery completes the phase or
project.

### Source identity

This refresh records coordinator evidence from 2026-11-02 UTC against integrated
source commit `f28531ec574ed3011d680e71d42979f84fc2d54b`, tree
`6d8b503c4bec5bb6dd79b4d437ee01cdb4e1173e`. It includes target-picker repair
`75a849cff9628d720ba74715491a32386afbfffa`, selected-window split correction
`7b5a40488353586dd90956bd792a1f7ddd4a3449`, numbered-tab intent adapter
`54a577bee6481ebe5110d7b280c03bcf307542c3`, and explicit native RPC/New Tab
intent adapter `f28531ec574ed3011d680e71d42979f84fc2d54b`.

The seven-effect producer gate **R** ran on candidate
`6f099b0bcafe413f16a28642ce5c5ef32b5b0ed5`, tree
`9456592454333021490aba0c415a675f40f8f301`; its Go/Lua product source is identical
to the integrated baseline. Final-main rendering gates **N/J/G/E**, selected-window
splits **S**, and normal-popup gate **P** use the final compiled product source.
These are supplied coordinator observations, not trials run by this documentation
builder. The builder performed source/manual-spec and whitespace audits only;
independent frozen review and coordinator combined compilation remain delivery
gates for this documentation revision.

Selected baseline file SHA-256 hashes (not binary or runtime hashes):

| File | SHA-256 |
| --- | --- |
| `defaults.go` | `ca73dda1b02a4690e88fb70d0ca3e9482a511fb025d8452156a68568e49ee4f0` |
| `native.go` | `a8d8576ac115f51e0a41b97fb698a51f7bf344e5d57dab23688b19293306667a` |
| `native_tabs.go` | `ec50a7a14aa2f5e7f435c93701f654d472eb90c516f6760fb75758d9feb5a147` |

## Defaults, contracts and dispatch

The manual 25-entry manifest **C** was independently compared with
[`defaults.go`](../defaults.go) and the Lua default fragments
([sessions](../lua/rex_whichkey/defaults/sessions.lua),
[tabs](../lua/rex_whichkey/defaults/tabs.lua),
[panes](../lua/rex_whichkey/defaults/panes.lua),
[app](../lua/rex_whichkey/defaults/app.lua)). Both specify six session bindings,
fourteen tab bindings and five pane/app bindings. The `t1`–`t9` loops supply
numeric, one-based indices. Only `wx`, `tx` and `x` require confirmation; `w` and
`t` are non-executable groups. The Lua
[composer](../lua/rex_whichkey/defaults/init.lua) includes those four fragments,
not the optional hosts provider. No callback is among these defaults. The default
schema and Lua fragments are unchanged by the native adapters.

Every row's **W/H** means its original action name and listed arguments match
both captured registries in **C**: **W** is the `work` capture
(`client:01a19cd1-3d93-7a79-a4ac-f31396ebeb52`); **H** is the genuinely captured
historical `m4-pro` registry
(`client:01a19dcb-960b-7b66-b984-fc3259bdb7b6`). At the latest contract capture,
work was live and that m4-pro connection was disconnected. This is not a fresh
cross-app behavioral certificate or a claim about present connectivity. No default
name/argument difference was found between those captures. `{}` denotes no binding
arguments; dispatch normalizes omitted args to `{}`. No rename text/dialog
arguments are invented. W/H alone does not certify the translated dispatch below.

The accepted [late-targeting design](../.tasks/designs/rex-whichkey-action-audit.md)
is refined by the later
[selected-window supplement](../.tasks/designs/rex-whichkey-selected-window-splits.md).
Display/Escape need no app choice. Executable selection resolves a strict explicit
selector, a unique attached app, or an explicit chooser when several qualify;
it never guesses a target. Original selection/args/confirmation remain retained.
Both UI/file/lease cleanups precede dispatch. Exact invocation origin, attachment
and app generation are revalidated, with no uncertain retry or silent fallback.

[`native.go`](../native.go) now uses documented `client.perform` through a separate
unscoped exact-server client. Among generic verbatim requests, only `pane.split`
omits outer session/block overrides: `v`/`s` split the chosen app's selected tab's
focused pane, not launcher A or shared
server `active_window_id`. `v` stays `direction: "right"`; `s` stays
`direction: "down"`, **not** CLI `below`. Other generic actions, especially
closes, retain explicit bound-origin context.

Two native intent families compile parent-owned selections rather than execute
their original Rex verbs verbatim:

- **`t1`–`t9`:** the historical `native_tabs.go` validates the retained integer
  and matching key, resolves the ordered **captured origin-session** windows once,
  freezes the exact window ID, rechecks membership/origin/app generation and sends
  `client.perform session.select` with exact `session_id`/`window_id` args and no
  outer context override. Missing targets refuse; reordering never substitutes an
  ID. This is not proof that Rex's opaque numeric `client.tab.goto` was repaired.
- **`tn`:** the parent validates retained empty-args New Tab intent and actual
  `session.select` registration, then uses typed `Session.NewWindow` once to create
  a default-terminal window in the **captured origin session**. It freezes the
  returned ID, rechecks registration/membership/origin/app generation, then sends
  exact-ID `session.select`. Creation and activation are separate effects. An
  activation failure preserves the known created ID for manual recovery: no retry,
  rollback/automatic close, second create or original-verb fallback.

Both adapters retain cleanup, targeting and the short signal-cancellable dispatch
budget. The configured authenticated original-VM route, Lua defaults and callback
semantics are unchanged and remain unverified by these native trials.

## All 25 bindings

**Resource verified** means the stated resulting resource effect was observed on
owned disposable fixtures. **Rendering verified** means a causal app-specific
measurement changed after the actual menu action, not that native pixels or
keyboard first-responder state were observed. For navigation, the fresh target
had **no Work desired-size report before** selection and acquired one **after**,
without explicit target control or input injection in between. Shared server
focus, an existing size report or RPC success alone is not this proof.

The action/args columns describe the unchanged defaults; `t1`–`t9` and `tn` use
the intent translations above. Each row has one status; **S** strengthens `v`/`s`
without adding bindings to the count.

| Keys | Native action | Args | Confirm | Contracts | Expected behavior | Actual status / observation | Evidence |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `wn` | `session.new` | `{}` | No | W/H | Create a new session. | **Resource verified:** exactly one new session; selected app view not certified. | C; R/wn |
| `ws` | `session.switch` | `{}` | No | W/H | Open Rex's native session picker. | **Unverified:** picker interaction not fully exercised. | C; U |
| `wh` | `session.previous` | `{}` | No | W/H | Select previous session in the app. | **Rendering verified:** exact Work app first measured the previously hidden adjacent owned session after menu action. | C; G/wh |
| `wl` | `session.next` | `{}` | No | W/H | Select next session in the app. | **Rendering verified:** exact Work app first measured the previously hidden adjacent owned session after menu action. | C; G/wl |
| `wr` | `session.rename` | `{}` | No | W/H | Open native session rename interaction in bound context. | **Unverified:** earlier attempted rename left labels unchanged; audit text/ESC reached the owned terminal, not a proven native modal. No completed rename verified. | C; U |
| `wx` | `session.close` | `{}` | Yes | W/H | After confirmation, close bound-origin session. | **Resource verified:** exact isolated session and its fixture resources removed. | C; R/wx |
| `tn` | `client.tab.new` | `{}` | No | W/H | Create a default-terminal tab in captured origin session and select its exact ID. | **Resource verified:** exactly one new window in intended owned session through parent-owned intent adapter; not a verbatim native New Tab pass. | C; R/tn; J/tn |
| `th` | `client.tab.previous` | `{}` | No | W/H | Select previous app tab. | **Rendering verified:** exact fresh previous target first acquired Work geometry after menu action. | C; N/th |
| `tl` | `client.tab.next` | `{}` | No | W/H | Select next app tab. | **Rendering verified:** interior next selected the exact previously unmeasured adjacent owned tab; no boundary-wrap claim. | C; J/tl |
| `tr` | `client.tab.rename` | `{}` | No | W/H | Open native tab rename interaction in bound context. | **Unverified:** rename interaction not fully exercised; no replacement-name args invented. | C; U |
| `tx` | `client.tab.close` | `{}` | Yes | W/H | After confirmation, close bound-origin tab. | **Resource verified:** origin tab and both fixture panes removed; audit session remained. | C; R/tx |
| `t1` | `client.tab.goto` | `{"index":1}` | No | W/H | Select exact tab 1 in captured origin-session order. | **Rendering verified:** fresh exact target first acquired Work geometry after menu selection via exact-ID adapter. | C; N/t1 |
| `t2` | `client.tab.goto` | `{"index":2}` | No | W/H | Select exact tab 2 in captured origin-session order. | **Rendering verified:** fresh exact target first acquired Work geometry after menu selection via exact-ID adapter. | C; N/t2 |
| `t3` | `client.tab.goto` | `{"index":3}` | No | W/H | Select exact tab 3 in captured origin-session order. | **Rendering verified:** fresh exact target first acquired Work geometry after menu selection via exact-ID adapter. | C; N/t3 |
| `t4` | `client.tab.goto` | `{"index":4}` | No | W/H | Select exact tab 4 in captured origin-session order. | **Rendering verified:** fresh exact target first acquired Work geometry after menu selection via exact-ID adapter. | C; N/t4 |
| `t5` | `client.tab.goto` | `{"index":5}` | No | W/H | Select exact tab 5 in captured origin-session order. | **Rendering verified:** fresh exact target first acquired Work geometry after menu selection via exact-ID adapter. | C; N/t5 |
| `t6` | `client.tab.goto` | `{"index":6}` | No | W/H | Select exact tab 6 in captured origin-session order. | **Rendering verified:** fresh exact target first acquired Work geometry after menu selection via exact-ID adapter. | C; N/t6 |
| `t7` | `client.tab.goto` | `{"index":7}` | No | W/H | Select exact tab 7 in captured origin-session order. | **Rendering verified:** fresh exact target first acquired Work geometry after menu selection via exact-ID adapter. | C; N/t7 |
| `t8` | `client.tab.goto` | `{"index":8}` | No | W/H | Select exact tab 8 in captured origin-session order. | **Rendering verified:** fresh exact target first acquired Work geometry after menu selection via exact-ID adapter. | C; N/t8 |
| `t9` | `client.tab.goto` | `{"index":9}` | No | W/H | Select exact tab 9 in captured origin-session order. | **Rendering verified:** fresh exact target first acquired Work geometry after menu selection via exact-ID adapter. | C; N/t9 |
| `v` | `pane.split` | `{"direction":"right"}` | No | W/H | Split chosen app's selected tab's focused pane right. | **Resource verified:** focused pane in selected B split right, not launcher A; A layout unchanged and B's other anchor retained. | C; R/v; S/v |
| `s` | `pane.split` | `{"direction":"down"}` | No | W/H | Split chosen app's selected tab's focused pane down. | **Resource verified:** focused pane in selected B split down, not launcher A; A layout unchanged and B's other anchor retained. | C; R/s; S/s |
| `x` | `pane.close` | `{}` | Yes | W/H | After confirmation, close bound-origin pane. | **Resource verified:** only origin pane closed; exact adjacent anchor remained in same tab. | C; R/x |
| `e` | `client.sidebar.toggle` | `{}` | No | W/H | Toggle native sidebar. | **Rendering verified:** two actual menu toggles changed terminal columns×rows 124×86 → 158×84 → 124×86 with the same reporting app. | C; E |
| `c` | `client.palette.toggle` | `{}` | No | W/H | Toggle native command palette. | **Unverified:** direct/menu acknowledgments and focus bytes do not establish exact palette open/closed state. | C; U |

## Evidence index

Paths below are relative to the canonical parent checkout's
`.mtn-tmp/action-audit/`, not a leaf worktree. They are local coordinator artifacts,
not bundled screenshots or a portable cross-platform GUI certificate. For
`label/key`, use that binding's retained subdirectory and the aggregate result;
**J** has `tl` and `tn` entries in a shared directory.

- **C — manual manifest and contract captures:**
  [`contracts/default-bindings.json`](../.mtn-tmp/action-audit/contracts/default-bindings.json),
  [`contracts/result.json`](../.mtn-tmp/action-audit/contracts/result.json),
  `contracts/work-actions.json` and `contracts/m4-pro-actions.json`. These certify
  original default names/args against the work and historical m4-pro captures,
  not behavior or universal installed-version compatibility.
- **R — seven actual resource effects:**
  [`resource-runtime/rex-whichkey-p14-t06-builder-a02/result.json`](../.mtn-tmp/action-audit/resource-runtime/rex-whichkey-p14-t06-builder-a02/result.json).
  Each `v`, `s`, `tn`, `x`, `tx`, `wn`, `wx` directory retains its own result,
  before/after session/resource evidence and menu/action wait outputs. Closed `wx`
  uses session-list removal, not a surviving-session snapshot. Anchor and exact
  new-ID artifacts distinguish pane/tab/session effects. R's inherited note about
  eighteen GUI-only bindings describes its resource-only scope, not the final
  unverified count after N/J/G/E.
- **N — final-main numbered tabs and previous:**
  [`api-capability-investigation/tab-matrix-final-api-main/result.json`](../.mtn-tmp/action-audit/api-capability-investigation/tab-matrix-final-api-main/result.json).
  All nine `t1`–`t9` and `th` entries passed with exact expected window IDs.
  Per-binding `before-size.json`/`after-size.json`, views, menu layouts and selection
  outputs retain the fresh-target causal sequence. The aggregate is **partial**:
  its `tl` boundary case is inconclusive, not passed. J supplies the separate
  successful interior-next gate. Initial 80×24/unclaimed geometry is not viewer
  proof; the new Work desired-size entry after dispatch is the rendering signal.
- **J — final-main interior Next and New Tab:**
  [`api-capability-investigation/next-new-tab-final-api-main/result.json`](../.mtn-tmp/action-audit/api-capability-investigation/next-new-tab-final-api-main/result.json).
  `next-before.json`/`next-after.json` retain the interior adjacency observation;
  `new-before.json`/`new-after.json` and `new-window-ids.json` retain exactly one
  newly created owned window. No end-of-list wrap behavior is assumed.
- **G — final-main adjacent owned sessions:**
  [`api-capability-investigation/session-navigation-final-api-main/result.json`](../.mtn-tmp/action-audit/api-capability-investigation/session-navigation-final-api-main/result.json).
  `wh`/`wl` retain source/target fixture creation, before/after target sizes, actual
  menu selection and exact cleanup outputs. The exact Work app first measures
  the fresh target after the action. The observed adjacent cases pass; a creation
  order hypothesis is not a universal session-order/wrap contract.
- **E — final-main sidebar geometry round-trip:**
  [`api-capability-investigation/sidebar-final-api-main/result.json`](../.mtn-tmp/action-audit/api-capability-investigation/sidebar-final-api-main/result.json),
  `before.json`, `after-14.json`, `after-15.json` and menu outputs. Both columns
  **and** rows change and restore; constant row count is not required.
- **S — final-source selected-window split proof:**
  [`selected-runtime/final-api-repairs/result.json`](../.mtn-tmp/action-audit/selected-runtime/final-api-repairs/result.json)
  explicitly binds the final commit/tree above. Separate owned launcher A remains
  unchanged while selected B's known focused pane receives `s`/`v`; its other anchor
  survives. Rex may rebalance same-axis widths normally. Anchor identity retention
  does not require identical width. Per-binding snapshots/action/cleanup outputs
  are retained with the aggregate.
- **P — final-source actual normal `make run` popup:**
  [`../popup-fixed-navigation/runtime/final-api-repairs-make-run/result.json`](../.mtn-tmp/popup-fixed-navigation/runtime/final-api-repairs-make-run/result.json)
  binds the same final commit/tree. Promptly sampled first bounds equal
  root/Tabs/back/idle bounds, with actual Work geometry, compact startup seed,
  status spacers, Escape
  exit0 and exact cleanup. This guards fixed navigation presentation, not additional
  action outcomes or native controls hiding.
- **U — unresolved interactions and API investigation:**
  [`api-capability-investigation/current-action-matrix.json`](../.mtn-tmp/action-audit/api-capability-investigation/current-action-matrix.json)
  aggregates **21 verified / four unverified** on the final source; the row-specific
  gates above supply the actual evidence. Under `api-capability-investigation/`,
  `findings.md`, `api-list.json`, published Lua/API references, `rename-api-probe/`
  and `palette-focus-final-api-main/` retain the investigation. The rename attempt left
  labels unchanged and delivered audit/ESC bytes to an owned terminal reader.
  `pane.send_key`/`client.key` are terminal/Lua-key routing, not a demonstrated
  native-modal controller. Palette focus traces can also reflect popup transitions
  and multiple client reports; neither receipts nor those bytes prove palette
  state. `ws`/`tr` interactions remain not fully exercised. No supported native
  modal text/input/state control or completed rename is inferred.

## Historical corrections and remaining limits

The initial 18-resource/7-GUI partition was a planning assumption. The direct tab
probe in [`observability-correction.md`](../.mtn-tmp/action-audit/observability-correction.md)
showed that shared `session.active_window_id` is not selected-app tab state,
correcting it to seven resource candidates/eighteen other outcomes. The original
[`resource-runtime/integrated-selected-source/result.json`](../.mtn-tmp/action-audit/resource-runtime/integrated-selected-source/result.json)
then observed six effects: `tn` exited 0 without a new window/session. That was
not a pass or conclusive proof of breakage. The subsequent blanket assertion that
19 actions required screenshots was too broad: full API investigation established
functional observation through controlled resource/rendering consequences. Those
historical counts do not describe the current 21/four matrix.

Failed assessments remain failed or inconclusive:

- Original opaque numeric `client.tab.goto` requests were acknowledged without the
  expected fresh targets; exact-ID positive controls worked. The parent adapter
  and actual nine-binding gate establish the repaired product behavior, not the
  underlying Rex cause or a repair of its numeric native verb.
- [`resource-runtime/rex-whichkey-p14-t06-builder-a01/result.json`](../.mtn-tmp/action-audit/resource-runtime/rex-whichkey-p14-t06-builder-a01/result.json)
  retains the **failed transport-only New Tab gate**, six of seven effects, despite
  direct `client.perform`. A02's explicit create/exact-ID activation intent is a
  material implementation refinement, not relabeling a01 as passing. Direct API
  and CLI creation effects varied by context with the same executable verified;
  neither transport alone nor a claim that all CLI actions are broken explains it.
  The underlying Rex cause remains unestablished.
- N's boundary `tl` assessment assumed wrap behavior; its inconclusive result is
  retained, not promoted to pass. J verifies a fresh **interior** next case.
- `api-capability-investigation/sidebar-probe/original-constant-row-assessment.json`
  retains the mistaken constant-row sidebar criterion. E's complete geometry
  round-trip is the current evidence, not that failed assessment.
- The original selected-window private gate wrongly required unchanged anchor
  width on a right split. Its failed outputs remain at
  `rex-whichkey-p14-t04-builder-a01/integration-harness-recovery/original-selected-gate/`.
  Product source was unchanged for the fresh corrected gate at
  `selected-runtime/rex-whichkey-p14-t04-builder-a01-harness-recovery/result.json`;
  S subsequently refreshed both splits on final source. Normal Rex rebalancing is
  not a split-target failure, and the original gate is not called passing.

Historical two-app target-choice evidence at
[`runtime/rex-whichkey-p14-t01-builder-a01/result.json`](../.mtn-tmp/action-audit/runtime/rex-whichkey-p14-t01-builder-a01/result.json)
was genuine on `75a849c`: bare `s` offered a choice, choosing work split down,
waiting beyond the short dispatch budget still allowed selection, cancellation
made no pane, invalid exact selection refused without fallback, and both UI
layers were gone before observation. The earlier bare refusal with two apps was
ambiguity, not invalid `down`. This is not a fresh multi-app chooser trial at the
final revision; later selected-window fixtures are separate evidence. No
unrecorded cross-app or remote-platform behavior is certified.

Remaining work is individual API/native-interaction verification of `ws`, `wr`,
`tr` and `c`, with safe observation/dismissal and resulting-state evidence. Missing
screenshots alone are not a blanket blocker; equally, registry entries, action
acknowledgments, focus bytes and terminal text are not substitutes for those four
outcomes. The all-25 requirement remains open without waiver, as does independent
chrome REQ-026. Coordinator trials used only owned disposable IDs and exact cleanup;
no user/coordinator/worker resources, settings, fonts or configuration were changed.
No general/unit/race/TDD suite, leaf build/native experiment, configured callback
trial or remote-OS certification is claimed. Source review and compilation cannot
certify the remaining interaction outcomes.
