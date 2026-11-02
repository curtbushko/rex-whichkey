---
workflowKind: design
title: Native popup tab-return reliability
mtnCreateRetry:
  key: rex-whichkey-tab-return-design-v1
  intent: bae2bd3c4878152efe7c8be552715cf269cb30534eb68aec2ee25c4051b2b0ce
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T15:22:51.482Z'
dateModified: '2026-11-02T15:22:56.840Z'
tags:
  - task
type: task
id: 436ab3c5-924c-4fd3-b1a0-e93c55aef4bc
completedDate: '2026-11-02'
---
# Native popup tab-return reliability and remove default Zoom

## Authority and requirements
User requests removing Zoom and reports that returning to a tab with rex-whichkey open no longer shows its popup. User explicitly corrected coordinator: tab-return visibility is the bug; the separately reported several-app launch error is not its explanation. Existing delegated bounded technical authority applies. Preserve all completed phases02–08; this is Phase09, not a retrospective rewrite.

REQ-021: Remove z/Zoom pane from native Go and composed Lua defaults and current default documentation. Keep custom pane.zoom capability and unrelated session/tab/split/close/app behavior.
REQ-022: A native foreground popup open in its originating tab should remain usable and visible when switching away and back. Preserve root/submenu/navigation state, one owned popup and exact cleanup, no idle timer, no focus stealing while away, no selection/action from focus/terminal reports, no duplicate popup recreation or unsolicited client retargeting. Tab visibility alone is not an invocation cancellation. Preserve configured original-Lua-VM authenticated30second expiry and explicit selectors/generation checks before any executable selection.

## Verified evidence versus uncertainty
At fedbeb3, native RunTerminal does not enable focus reports; terminalKeys consumes CSI focus reports without a repaint. Redraw occurs on keys/palette/resize/CONT only. Current native waitPopup additionally polls complete app/session membership/generation every second and cancels upon its loss; this must not be confused with merely changing active windows. Actual coordinator reproduction with one live attached app, both server window.focus round-trip and native app client.tab.goto round-trip, retained same process/layer/menu contents. This is NOT proof the remote app displays its floating layer correctly. Artifacts .mtn-tmp/tab-recovery/reproduction and published native API metadata are read-only inputs. The app's GUI frame cannot be inferred solely from terminal capture. A previous several-app error was consistent with two real app identities; current list has work connected and m4-pro lost. Never guess labels, first/last app or assume tab switching creates clients.

## Bounded implementation direction
Investigate actual foreground return paths and implement the smallest defensible in-repository fix. A standard terminal focus-in repaint path is within scope if native/published contracts support it: save/enable/restore focus-reporting mode, consume complete focus-in notifications as redraw events rather than mapping keys, preserve parser framing and NO_COLOR/theme/state. Do not invent a Rex layer persistence flag: current session.new_layer schema has none. Do not run a periodic full-screen repaint, steal global focus, recreate layers on a timer, add daemon/config/Lua VM, or modify Rex itself. If transient app visibility/attachment is found to cancel the native UI, separate idle resource-liveness checks from executable-action authorization; do not weaken exact endpoint/origin/action/generation validation or silently retarget actions. Preserve cancellation/parent FIFO/process-loss/owned-only cleanup and bounded operations.

## Verification and exclusions
Leaf builders/reviewers: formatting/whitespace/source review only, no heavy/compilation slot, tests/TDD/general suites/product/GUI/API/config/remote trials. Coordinator compiles exact reviewed combined candidate using Taskfile and Go, then performs bounded existing Linux/Rex popup, root/submenu actual tab-away/back via native app actions on only owned diagnostic tabs, captured process/layer/focus/frame observations, Escape exit0 and exact cleanup. Verify any new focus-in repaint path through actual terminal behavior without calling source review runtime proof; record whether genuine disappearance was reproduced and limits of server snapshots versus app visuals. No tests/race suites/destructive actions/config reloads/settings exports/remote feature or Darwin trials. Standard metadata/source/help research is allowed. No push or global cleanup.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE1OjIyOjU1LjkzMFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS90YWItcmV0dXJuL2Rlc2lnbi1hcHByb3ZhbC8xIiwiaW50ZW50IjoiNzU1M2I3NjViYWZjZTg1MWY4NzZkNWRkY2M3NmRhNzJiMTgxMGY3ODVlMGI2ZmVhZjBkM2M4Y2NkMWFkZDlkZSJ9 -->
### 2026-11-02T15:22:55.930Z · approval · unspecified

{"approvedBy":"explicit user bug correction and Zoom removal; bounded technical implementation delegation","attempt":1,"decision":"approved","decisionEvidence":"User says remove Zoom and explicitly clarifies that returning to a tab no longer shows the popup. Coordinator verified ordinary server/app round-trip retains menu/process/layer, not remote GUI visibility; narrow foreground redraw investigation/fix and honest runtime limits authorized. No routine approval pause, no guessed crash/client-ambiguity cause, no new selection policy or test/config/destructive/remote authority.","decisionRegisterEntry":{"decision":"Remove default Zoom and improve native popup tab-return reliability; trace real focus/input/render lifecycle, retain configured expiry/security and record actual server-versus-GUI evidence limits.","id":"REXWHICHKEY-20261102-102253","state":"Accepted"},"eventId":"rex-whichkey/tab-return/design-approval/1","kind":"approval","recordRef":"designs/rex-whichkey-tab-return.md","requirements":["REQ-021","REQ-022"],"schemaVersion":1,"specHash":"sha256:a616c835e37588e3c673b289849b40573db242126626bcdc1597a9c2a76607c9"}
<!-- mtn:end-entry:v1 -->
