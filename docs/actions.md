# Native default action audit

## Scope and current result

REQ-030 requires behavioral verification of **all 25 executable defaults**.
This is a partial audit, not an all-actions pass: **six bindings have observed
server resource effects; 19 still need native app observation**. Registry
compatibility, a successful invocation, a resource effect and visible GUI behavior
are separate claims. No remote GUI pixels or configured original-Lua-VM callback
behavior are certified here.

This document records coordinator evidence from 2026-11-02 UTC; the documentation
builder performed source/manual-spec and whitespace audits only, without running
builds, tests, native actions or new experiments. Source baseline is
`aa531748729b3ac7a7820d7d9499503647f17d39`, containing target-picker repair
`75a849cff9628d720ba74715491a32386afbfffa` and selected-window split correction
`7b5a40488353586dd90956bd792a1f7ddd4a3449`. The defaults and native routing files
are unchanged between `7b5a404` and that baseline. The current-source resource
matrix is **R** below; the stronger selected-window split gate **S** identifies
candidate `7b5a40488353586dd90956bd792a1f7ddd4a3449`, tree
`af20373b696b9c6df518aefc2228c75bc0dbefe7`.

[Task 14-03](../.tasks/tasks/rex-whichkey-14-03-gui-action-verification.md) remains
**REQUIRED / BLOCKED**, not deferred by the user. All-25 verification remains
incomplete. Native popup chrome REQ-026 is independently blocked; see
[acceptance status](acceptance.md). This document does not complete the phase or
project.

## Defaults, contracts and targeting

The manual 25-entry manifest **C** was independently compared with
[`defaults.go`](../defaults.go) and the Lua default fragments
([sessions](../lua/rex_whichkey/defaults/sessions.lua),
[tabs](../lua/rex_whichkey/defaults/tabs.lua),
[panes](../lua/rex_whichkey/defaults/panes.lua),
[app](../lua/rex_whichkey/defaults/app.lua)). Both sources specify six session
bindings, fourteen tab bindings and five pane/app bindings. The `t1`–`t9` loops
supply numeric, one-based indices. Only `wx`, `tx` and `x` require confirmation;
`w` and `t` are non-executable groups. The Lua
[composer](../lua/rex_whichkey/defaults/init.lua) includes those four fragments,
not the optional hosts provider. No callback is among these 25 defaults.

Every row's **W/H** means its action name and listed arguments match both captured
installed registries in **C**: **W** is the current `work` registry
(`client:01a19cd1-3d93-7a79-a4ac-f31396ebeb52`); **H** is the genuinely captured
historical `m4-pro` registry
(`client:01a19dcb-960b-7b66-b984-fc3259bdb7b6`). `m4-pro` was disconnected at the
latest capture; this is **not** two currently live apps or behavioral proof on
either app. No default action/argument contract difference was found between
those captures. `{}` means no binding arguments; native dispatch normalizes
omitted args to `{}`. No undocumented rename text or dialog arguments are added.

The accepted [late-targeting design](../.tasks/designs/rex-whichkey-action-audit.md)
is refined by the later
[selected-window supplement](../.tasks/designs/rex-whichkey-selected-window-splits.md).
Native display does not require an app choice. Executable selection resolves a
strict explicit selector, a unique attached app, or an explicit chooser when
several apps qualify; it never guesses a target. The original action, arguments
and confirmation remain retained. UI cleanup precedes dispatch, and exact origin,
attachment and app connection generation are revalidated before one dispatch.

Only native `pane.split` omits session/block overrides, using the chosen app's
currently selected window/tab and focused pane through an unscoped exact-server
`client.perform` call. Invocation origin remains a liveness/authorization check,
not the split destination. `v` remains `direction: "right"`; `s` remains
`direction: "down"`, **not** the CLI spelling `below`. Other native actions,
including destructive closes, retain their existing bound-origin context.
Configured authenticated original-VM routing is unchanged and unverified by these
native trials.

## All 25 bindings

**Effect verified** means only the stated server resource result was observed in
coordinator-owned disposable fixtures on the current integrated source. It does
not certify native GUI presentation. **Not exercised** means no menu-driven
behavioral result is supplied, even where a related direct CLI probe exists.
All GUI-local outcomes below remain unverified, not conclusively broken.

| Keys | Native action | Args | Confirm | Contracts | Expected behavior | Actual status / observation | Evidence |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `wn` | `session.new` | `{}` | No | W/H | Create a new session. | **Effect verified:** exactly one new session resource; selected app view not certified. | C; R/wn |
| `ws` | `session.switch` | `{}` | No | W/H | Open Rex's native session picker. | **Not exercised; GUI unverified.** | C; O |
| `wh` | `session.previous` | `{}` | No | W/H | Select previous session in the app. | **Not exercised; GUI unverified.** | C; O |
| `wl` | `session.next` | `{}` | No | W/H | Select next session in the app. | **Not exercised; GUI unverified.** | C; O |
| `wr` | `session.rename` | `{}` | No | W/H | Open native session rename interaction in bound context. | **Not exercised; GUI unverified.** No unclosable dialog opened for this audit. | C; O |
| `wx` | `session.close` | `{}` | Yes | W/H | After confirmation, close bound-origin session. | **Effect verified:** exact isolated audit session and fixture resources removed. | C; R/wx |
| `tn` | `client.tab.new` | `{}` | No | W/H | Open/create a new tab in the app. | **Invoked, unverified:** menu invocation exited 0, but no new server window or session observed. Receipt is not a pass or proof of breakage. | C; R/tn; O |
| `th` | `client.tab.previous` | `{}` | No | W/H | Select previous app tab. | **Not exercised; GUI unverified.** Shared server focus cannot certify selection. | C; O; N |
| `tl` | `client.tab.next` | `{}` | No | W/H | Select next app tab. | **Not exercised; GUI unverified.** Shared server focus cannot certify selection. | C; O; N |
| `tr` | `client.tab.rename` | `{}` | No | W/H | Open native tab rename interaction in bound context. | **Not exercised; GUI unverified.** No replacement-name args invented. | C; O |
| `tx` | `client.tab.close` | `{}` | Yes | W/H | After confirmation, close bound-origin tab. | **Effect verified:** origin tab and both fixture panes removed; audit session remained. | C; R/tx |
| `t1` | `client.tab.goto` | `{"index":1}` | No | W/H | Select app tab 1. | **Not exercised; GUI unverified.** Related direct CLI goto probe is not this menu binding's pass. | C; O; N |
| `t2` | `client.tab.goto` | `{"index":2}` | No | W/H | Select app tab 2. | **Not exercised; GUI unverified.** | C; O; N |
| `t3` | `client.tab.goto` | `{"index":3}` | No | W/H | Select app tab 3. | **Not exercised; GUI unverified.** | C; O; N |
| `t4` | `client.tab.goto` | `{"index":4}` | No | W/H | Select app tab 4. | **Not exercised; GUI unverified.** | C; O; N |
| `t5` | `client.tab.goto` | `{"index":5}` | No | W/H | Select app tab 5. | **Not exercised; GUI unverified.** | C; O; N |
| `t6` | `client.tab.goto` | `{"index":6}` | No | W/H | Select app tab 6. | **Not exercised; GUI unverified.** | C; O; N |
| `t7` | `client.tab.goto` | `{"index":7}` | No | W/H | Select app tab 7. | **Not exercised; GUI unverified.** | C; O; N |
| `t8` | `client.tab.goto` | `{"index":8}` | No | W/H | Select app tab 8. | **Not exercised; GUI unverified.** | C; O; N |
| `t9` | `client.tab.goto` | `{"index":9}` | No | W/H | Select app tab 9. | **Not exercised; GUI unverified.** | C; O; N |
| `v` | `pane.split` | `{"direction":"right"}` | No | W/H | Split chosen app's selected tab's focused pane right. | **Effect verified:** S/v split focused pane in selected B, not launcher A; A layout unchanged, existing B anchor retained. | C; R/v; S/v |
| `s` | `pane.split` | `{"direction":"down"}` | No | W/H | Split chosen app's selected tab's focused pane down. | **Effect verified:** S/s split focused pane in selected B, not launcher A; A layout unchanged, existing B anchor retained. | C; R/s; S/s |
| `x` | `pane.close` | `{}` | Yes | W/H | After confirmation, close bound-origin pane. | **Effect verified:** only origin pane closed; exact adjacent anchor remained in same tab. | C; R/x |
| `e` | `client.sidebar.toggle` | `{}` | No | W/H | Toggle native sidebar. | **Not exercised; GUI unverified.** | C; O |
| `c` | `client.palette.toggle` | `{}` | No | W/H | Toggle native command palette. | **Not exercised; GUI unverified.** | C; O |

## Evidence index and limits

Evidence paths below are relative to the canonical parent checkout's
`.mtn-tmp/action-audit/`, not a leaf worktree. These are local coordinator artifacts,
not bundled screenshots or a portable GUI certificate. For **R/key** and **S/key**,
use the corresponding binding subdirectory as well as the aggregate result.

- **C — contracts and manual manifest:**
  [`contracts/default-bindings.json`](../.mtn-tmp/action-audit/contracts/default-bindings.json),
  [`contracts/result.json`](../.mtn-tmp/action-audit/contracts/result.json),
  `contracts/work-actions.json` and `contracts/m4-pro-actions.json`. These certify
  names/arguments against the current work and historical m4-pro captures only.
- **R — current integrated resource matrix:**
  [`resource-runtime/integrated-selected-source/result.json`](../.mtn-tmp/action-audit/resource-runtime/integrated-selected-source/result.json).
  Each `v`, `s`, `tn`, `x`, `tx`, `wn`, `wx` directory retains `result.json`,
  `before.json`, `sessions-before.json`, `sessions-after.json`, action wait and menu
  evidence. Surviving sessions also have `after.json`; closing `wx` removes the
  session, so its session-list delta is the evidence instead. `x/anchor-pane.json`
  and `tx/anchor-pane.json` distinguish pane versus tab removal.
  `wn/new-session-ids.json` and `wn/new-session.json` identify the new resource.
  `tn/action-wait.json` records `exit_code: 0`; `tn/new-window-ids.json` is empty.
  Ordinary same-window split results here do not alone prove selected B routing.
- **S — strongest selected-window split proof on `7b5a404`:**
  [`selected-runtime/rex-whichkey-p14-t04-builder-a01-harness-recovery/result.json`](../.mtn-tmp/action-audit/selected-runtime/rex-whichkey-p14-t04-builder-a01-harness-recovery/result.json).
  Each `s`/`v` directory retains fixture creation, `before.json`, `after.json`,
  selected-target command evidence, action wait, result and exact cleanup outputs.
  Separate owned launcher A stays unchanged; selected B's focused pane receives
  the directional split, and the other anchor survives. Rex normally rebalances
  same-axis widths: anchor identity retention does **not** require identical width.
  The original private gate wrongly required unchanged anchor width for a right
  split; its failure is retained at
  `rex-whichkey-p14-t04-builder-a01/integration-harness-recovery/original-selected-gate/`.
  Product source was unchanged; a fresh corrected gate passed. The original failed
  gate is not relabeled as passing.
- **T — historical two-app target chooser on `75a849c`:**
  [`runtime/rex-whichkey-p14-t01-builder-a01/result.json`](../.mtn-tmp/action-audit/runtime/rex-whichkey-p14-t01-builder-a01/result.json).
  Real bare `s` offered explicit app choice; choosing work produced a below-origin
  split. Waiting beyond the short RPC budget still allowed dispatch; chooser
  cancellation made no pane; an invalid exact selector refused without fallback.
  Both UI layers were gone before action observation. This predates the selected-B
  correction and is **not** a fresh multi-app chooser run at `7b5a404`; newer actual
  cases had only one attached app, and shared target-routing files were unchanged.
  Historical `split-reproduction/` records the earlier bare `s` refusal with two
  apps and successful explicit-work split down. That was target ambiguity, not an
  invalid `down` argument.
- **N — tab-navigation observability correction:**
  `tab-navigation-before.json`, `tab-navigation-after.json` and
  [`tab-navigation-probe.json`](../.mtn-tmp/action-audit/tab-navigation-probe.json),
  explained in
  [`observability-correction.md`](../.mtn-tmp/action-audit/observability-correction.md).
  A direct CLI `client.tab.goto` to an owned tab returned `null`, while shared
  `session.active_window_id` remained the coordinator window. This is not a run of
  all eleven menu navigation bindings and does not expose selected-app tab state.
- **O — installed observation limits:** `client-inspect-api.json`,
  `client-key-api.json`, `client-perform-api.json` and `api-list.json`.
  `client.inspect` reports client metadata, attachments and action declarations,
  not selected tab/session, sidebar or dialog state. `client.key` targets a block
  through Lua bindings and then its PTY; it is not native app-modal key injection.
  The inspected API list supplies no native selected-view/sidebar/dialog observer.
  No screenshot capability, modal-dismiss mechanism, GUI pass or unrecorded action
  receipt is inferred from these schemas.

The initial design's 18-resource/7-GUI partition was a planning assumption, not
verified capability. N first corrected it to seven resource candidates and eighteen
GUI-local cases. R then left `tn` unverified despite exit 0: the current result is
**six observed resource effects and 19 pending GUI-local outcomes** (`tn`, `th`,
`tl`, `tr`, `t1`–`t9`, `ws`, `wh`, `wl`, `wr`, `e`, `c`). These revisions do not
waive the user's all-25 requirement or prove `tn` broken.

Remaining verification needs supported native app observation, including safe
interaction/dismissal for dialogs. Do not substitute server focus, registry
presence, RPC success or terminal output for it. Coordinator trials were restricted
to owned disposable IDs with exact cleanup; no user/coordinator/worker resources,
settings, fonts or configuration were changed. General/unit/race tests, remote
platform behavior and configured callbacks were outside this audit. Independent
frozen review and coordinator combined compilation remain delivery gates for this
documentation; neither would certify the missing GUI outcomes.
