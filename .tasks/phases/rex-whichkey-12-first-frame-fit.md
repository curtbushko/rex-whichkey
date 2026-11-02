---
workflowKind: phase
workflowOrder: 12
workflowRequirements:
  - REQ-027
  - REQ-028
workflowDocuments:
  - ref: designs/rex-whichkey-first-frame-fit.md
    hash: 'sha256:e40a82729585e2a35de60126c4a3771c4c651d92e287f55c2808ff32d244e777'
title: 'Phase12: First-frame compact fit'
mtnCreateRetry:
  key: rex-whichkey-phase12-v1
  intent: fcdb50de1a43cba096b727ec28988b023c76b4038910e58f518fc0c41e5b7769
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T19:10:51.539Z'
dateModified: '2026-11-02T19:33:54.177Z'
tags:
  - task
type: task
id: c258ac4f-0f45-4211-880f-a4577d07c5f5
projects:
  - '[[projects/rex-whichkey.md]]'
completedDate: '2026-11-02'
---
# Phase12: First-frame compact sizing

User requests status/menu spacer rows and no visible oversized popup before fitting. Implement REQ027/028 in one coherent source/docs task with exact precreation geometry/terminal seed, frozen independent review and actual early-first-layer/spacing/adaptation/cleanup check. Depend on completed phase11 renderer/height/docs tasks, not its whole milestone: REQ026 native title/dock controls remains blocked separately. No fake hide/global settings, tests, font changes, destructive/native selections or remote/Darwin trials. Same coordinator/session. Phase12 can complete on actual supported scope; phase11 and project cannot while chrome remains unresolved.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE5OjMzOjU0LjE2NloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcGhhc2UxMi9zdW1tYXJ5LzEiLCJpbnRlbnQiOiI2MTRlZmQ0MDRlZDdhMmU3YjFkZWM2ZWI2OTI5MTc3ZjU5MzcyN2ZhYzJjNTBlZGNlODUzMDdlMzA5NjJjNmZiIn0 -->
### 2026-11-02T19:33:54.166Z · summary · unspecified

{"acceptedScope":"Phase12: First-frame compact fit","approvalSummary":["User explicitly requests status spacers and initial sizing before the popup draws rather than large-then-small correction. Phase12 frozen review/combined compilation/early-first-layer and normal adaptation/cleanup checks authorized; native chrome remains separately blocked in phase11.","All exact native child summaries bind current independent source approval and delivered target commit/tree; ancestry verified."],"artifactRefs":[".mtn-tmp/popup-initial-fit/phase12-final/validate.json",".mtn-tmp/popup-initial-fit/phase12-final/worktrees.txt",".mtn-tmp/popup-initial-fit/runtime/rex-whichkey-p12-t01-builder-a02/result.json"],"attempt":1,"coverage":{"REQ-027":["tasks/rex-whichkey-12-01-first-frame-fit.md"],"REQ-028":["tasks/rex-whichkey-12-01-first-frame-fit.md"]},"eventId":"rex-whichkey-p03-coordinator-a01/phase12/summary/1","files":["README.md","docs/acceptance.md","docs/appearance.md","docs/launch-contract.md","layout.go","popup.go","popup_initial.go","popup_size.go","render.go","rex/block.go","rex/geometry.go","rex/layer.go","rex/terminal.go"],"kind":"summary","limitations":["No tests/TDD/general suites, native/destructive selections, configured Lua callbacks, font/settings/config changes/reloads or remote/Darwin trials. Initial and subsequent bounds/geometry/status spacing observed in one measured context; seed and first-snapshot evidence are not pixel-perfect remote GUI proof. REQ026 native title/dock controls remains blocked in phase11; no full project completion.","Independent fresh contexts used the same authenticated model family, not cross-family review.","Prior transferred coordinator process status remains unverified; ownership transfer did not assert remote/local stoppage.","No global cleanup or push; retained worktrees/artifacts remain available."],"recordRef":"phases/rex-whichkey-12-first-frame-fit.md","records":["tasks/rex-whichkey-12-01-first-frame-fit.md"],"requirements":["REQ-027","REQ-028"],"schemaVersion":1,"specHash":"sha256:88d07bd9eb72c870ccd9e95bf4ff1bc4053eb474c8cce2daa7fbbc2138339243","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"caff550e0e7079064f0d34d18abc32785cdfa4e3","targetTree":"071100e2621621670de37b6d077ee3f29dd68ea6","verification":[{"command":"Actual promptly observed first layer compact before watcher, startup terminal geometry seeded, status spacer rows, settled/submenu/back fit, Escape exit0/owned cleanup; tracked source equivalence verified","result":"pass"},{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./...; make build (combined delivered revision and refreshed CLI artifact)","result":"pass"},{"command":"git diff --check; native graph/schema validation; exact completed child scope and Git delivery identities/ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
