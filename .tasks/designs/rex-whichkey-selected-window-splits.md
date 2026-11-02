---
workflowKind: design
title: Native splits target currently selected app window
mtnCreateRetry:
  key: rex-whichkey-selected-window-design-v1
  intent: 7edef83c36591b93c3653034c48fb439f0b9188d2736ffcd755285dba4d44927
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T20:38:07.021Z'
dateModified: '2026-11-02T20:38:12.364Z'
tags:
  - task
type: task
id: 989f8303-7126-4c78-bdc3-da234a89fe45
completedDate: '2026-11-02'
---
# Selected-window native splits (REQ-032)

Latest explicit user correction: "split down or vertical should be of the current window selected". For native v/s, split the focused pane in currently selected tab/window of chosen Rex app, NOT popup child, launch-origin block/window, or shared server active_window_id. Prior origin-forced native pane.split behavior is superseded only for split action target context; original immutable selection/args/confirmation and strict GUI target generation remain. Down argument is confirmed down, right is right; do not use CLI below with pane.split.

Installed client.perform schema explicitly says omitted block_id uses client's focused block. Omitted session_id lets client current context apply. Current native.go unconditionally supplies --session origin.SessionID --block origin.BlockID; that overrides desired target. The generic rex.Client.Call uses scoped Session to inject session_id, so create a separate UNscoped exact-server short-timeout client if using documented client.perform directly. Pass only client_id/action/retained args/origin/timeout as appropriate, NO session_id/block_id. Do not rely on omitting CLI flags while inheriting REX_SESSION/BLOCK/WINDOW; ambient CLI context must not leak back in. Preserve original scoped client for exact invocation/session/app-generation validation immediately before one dispatch. The source origin authorizes invocation/liveness, not native split destination. Generic Call accepts nil result to discard validated JSON safely. Uncertain API/decode/timeout errors remain not retried and sanitized. No duplicate native dispatch or speculative fallback.

Scope native pane.split only (v/right and s/down); don't silently change destructive pane/session/tab close targets or configured authenticated originalVM route. Other actions retain their existing target semantics pending audit/user requirements. All UI/target choice/files/leases gone before dispatch; target app exact fullID/connection generation must remain validated and user chooses when ambiguous. Basic display and Escape remain app-independent; submenu geometry fixed.

Coordinator actual gate: owned launcher tab A and separate owned selected tab B; B has two fixture panes with known focused first pane. Present B in exact work app immediately before final native split selection/target choice; drive owned hidden A menu via terminal keys. For both v/s assert A tiled layout unchanged, B gets correct directional split of its focused original pane and its other anchor survives. Test ordinary same-tab routing too via preceding bare target-picker gate. No use of server active_window_id as app-selected state, no mutation of any user/coordinator/worker pane. Cleanup exact fixture IDs after actual process exits/resource evidence. Source leaves no builds/tests/RPC/GUI. No general tests/settings/config/fonts/platform trials/push.

Additional verified audit correction: native client.tab.goto returns null and does not update server active_window_id; 11 tab navigation bindings are GUI-only observations, not certified resource effects. Actual first resource matrix produced six effects v/s/x/tx/wn/wx; tn exit0 but no new server window or session observed. tn remains unverified (may be client-local UI), not automatically pass or conclusively broken. Thus six resource outcomes verified,19 native GUI outcomes/presentation pending. Keep all25 verification requirement open; distinguish source/registry/receipt/effect/visibility. REQ026 native chrome remains independently blocked.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDIwOjM4OjExLjQ1OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9hY3Rpb24tYXVkaXQvc2VsZWN0ZWQtd2luZG93LWFwcHJvdmFsLzEiLCJpbnRlbnQiOiIzNDZiZDM4ZDg1MDMxYjhiNzYzZTNjMzg0YWUyZTZiZTRhNGVjYmU0YTM1NzQ4YTQwZTgwYTlhMGQ3ZjlmYTNmIn0 -->
### 2026-11-02T20:38:11.459Z · approval · unspecified

{"approvedBy":"Explicit latest user split down or vertical should be of current window selected","attempt":1,"decision":"approved","decisionEvidence":"Native s/v act on chosen app current selected tab/focused pane, not forced launcher origin/shared server focus. API client.perform explicitly supports omitted focused-block override. Original invocation/generation/action/cleanup safety and all other configured/destructive target contracts retained. Corrected audit classification is six verified effects versus19 unobserved UI outcomes, no waiver.","decisionRegisterEntry":{"decision":"Native pane.split uses chosen app current selected tab/focused pane without session/block overrides; origin still invocation/liveness proof, not split destination.","id":"REXWHICHKEY-20261102-153808","rationale":"Latest explicit user correction; existing code always forces launch-origin resource.","refines":"REXWHICHKEY-20261102-151552","state":"Accepted"},"eventId":"rex-whichkey/action-audit/selected-window-approval/1","kind":"approval","recordRef":"designs/rex-whichkey-selected-window-splits.md","relatedDecisions":[{"authority":"Bounded implementation choice under prior explicit delegation, not a claim the user personally chose this UI design","decision":"Late explicit app chooser for ambiguous native action target; retain immutable original action/args/confirmation, strict explicit selectors, exact generation/origin checks and cleanup before one dispatch. Verify resource effects and keep GUI-only tests blocked until observable.","id":"REXWHICHKEY-20261102-151552","previousState":"Accepted","rationale":"Shared refusal blocks all native actions with two apps. Never select an arbitrary device or hide unverified GUI behavior.","refinedBy":"REXWHICHKEY-20261102-153808","refinement":"Late target chooser stays; native split destination is current selected app context. Audit18/7 assumption corrected to actual six effects/19 pending native GUI outcomes.","state":"Refined"}],"schemaVersion":1,"specHash":"sha256:1d5e6297352b97b84e3ba7730e586130b4103ddb4ad32ed03d35b6ac2ad6e4f2"}
<!-- mtn:end-entry:v1 -->
