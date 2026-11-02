---
workflowKind: design
title: Stable popup geometry during navigation
mtnCreateRetry:
  key: rex-whichkey-fixed-navigation-design-v1
  intent: 1fadafb961ef7e0a8bcaaa0b1627a5a69a7957494202787bd7341e2813e6673b
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T19:37:46.696Z'
dateModified: '2026-11-02T19:37:52.096Z'
tags:
  - task
type: task
id: 4e5b6b8f-131a-43e3-a477-3502529bedf7
completedDate: '2026-11-02'
---
# Stable popup geometry during navigation

REQ-029: user explicitly says do not resize the window on every submenu entry. Keep the one precomputed initial layer rectangle unchanged by menu navigation for the invocation. Larger menus use existing renderer/pagination, not Layer.Resize. User/native external window/font resize can still change terminal cells; redraw/adapt content inside the existing box without application-driven layer-bound changes.

This later correction refines/supersedes only the navigation-driven resize part of REQ-025/phase11. Preserve REQ-027 spacers/tiny-view degradation and REQ-028 calculated compact initial bounds plus terminal startup seed. Do not remove those startup fixes, choose a giant maximum rectangle to fit every submenu, reintroduce Page1/1, or require one GUI app. Keep sufficient conservative cell-sized initial headroom so root spacing fits when new-pane chrome/font differs; source12 live evidence showed first H0.1105 versus final H0.1312, not a pixel-perfect estimate. No fixed pixel font guesses or cross-invocation caches.

Remove the now-unnecessary private preferred-height hint publisher/file transport and parent live sizing callback/feedback state rather than retaining a dead or conditional resize loop. Existing liveness watcher remains for resource loss/configured generation, not geometry. Preserve direct RunTerminal API, focus-in/SIGWINCH redraw, native FIFO/no-idle/app-independent display and deferred strict executable targeting, configured original-VM authenticated expiry, and owned-only cleanup before action. Generic Rex Resize/geometry adapter APIs may remain for legitimate public resource use; popup production navigation must never call Resize. Retire deleted source by moving to ignored .trash, not destructive deletion.

REQ026 native title/dock controls remains blocked in phase11; do not fake hiding/global config or claim project complete. Prior phase12 first-frame/spacer delivery remains historical proof, not evidence of new stable-navigation requirement. One coherent source/docs task depends on completed phase12; fresh frozen review and coordinator combined compilation/actual initial bounds + stable EXACT bounds at root, Tabs, optional paging/back and idle, spacers, Escape exit0/owned cleanup. No general tests/TDD/unit/race, settings/font edits/reloads, native/destructive selections or remote/Darwin trial. No push. Phase13 may complete; whole project stays incomplete for REQ026.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE5OjM3OjUxLjE2OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9maXhlZC1uYXZpZ2F0aW9uL2Rlc2lnbi1hcHByb3ZhbC8xIiwiaW50ZW50IjoiZjRkNTg4ZDFmYzNkNDQ4ZDdlZWUxYjViMDkwYzQwYTVkMmM2MDJjYjJhNmY3NzBlOGE3MmQ2Y2VjMTBjMjgwZSJ9 -->
### 2026-11-02T19:37:51.169Z · approval · unspecified

{"approvedBy":"explicit latest user Do not resize the window on every submenu entry","attempt":1,"decision":"approved","decisionEvidence":"Latest explicit correction overrides navigation-driven resizing. Keep initial compact calculation/seed, stable rectangle through submenu/paging/back/idle and use content pagination. Retire unnecessary live geometry IPC; no settings/font/test/destructive/remote authority. Native title/dock remains blocked.","decisionRegisterEntry":{"decision":"Keep initial popup bounds unchanged through navigation; larger menus use existing cells/pagination, not live layer resize. Retain first-frame sizing/spacers and strict lifecycle/action safety.","id":"REXWHICHKEY-20261102-143748","rationale":"Explicit later user correction rejects changing box size at each submenu.","refines":"REXWHICHKEY-20261102-123334","state":"Accepted"},"eventId":"rex-whichkey/fixed-navigation/design-approval/1","kind":"approval","recordRef":"designs/rex-whichkey-fixed-navigation.md","relatedDecisions":[{"decision":"Compact status-border rendering and content-fitting owned layer using actual cells; retain native controls hiding as explicitly blocked requirement, no invented flags or full completion.","id":"REXWHICHKEY-20261102-123334","previousState":"Accepted","refinedBy":"REXWHICHKEY-20261102-143748","refinement":"Only navigation-driven height adaptation is superseded by fixed bounds; compact initial/content rendering and unresolved native chrome scope remain.","state":"Refined"}],"requirements":["REQ-029"],"schemaVersion":1,"specHash":"sha256:7fbb1596666f9acf0ddb56ecf00ff29c09c8ae26781577268138e6a4ca05177e"}
<!-- mtn:end-entry:v1 -->
