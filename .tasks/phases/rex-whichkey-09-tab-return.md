---
workflowKind: phase
workflowOrder: 9
workflowRequirements:
  - REQ-021
  - REQ-022
workflowDocuments:
  - ref: designs/rex-whichkey-tab-return.md
    hash: 'sha256:a616c835e37588e3c673b289849b40573db242126626bcdc1597a9c2a76607c9'
title: 'Phase09: Native popup tab return'
mtnCreateRetry:
  key: rex-whichkey-phase09-v1
  intent: bd2997b0827c959b84c20447a09ed5c26902a37669a7f796410c0b3da0ccf3be
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T15:22:57.739Z'
dateModified: '2026-11-02T15:36:59.595Z'
tags:
  - task
type: task
id: a6c6e0e8-3f79-4336-bd8d-7ba9935c6546
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy:
  - uid: '[[phases/rex-whichkey-08-popup-polish.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-02'
---
# Phase09: Tab-return reliability

User correction: the popup stops showing when returning to its tab. Implement REQ-021 removing Zoom defaults and REQ-022 native tab-return/repaint reliability from approved designs/rex-whichkey-tab-return.md, preserving prior delivery and configured authenticated expiry. One bounded implementation task owns Go/Lua/default/docs interaction, with fresh mixed source review and coordinator-owned compilation plus actual owned-tab round-trip/navigation/Escape/cleanup check. Do not claim source/server capture proves remote graphical presentation. No tests/TDD/general suites, destructive/native selections, config reloads/settings edits, new client-selection policy, remote or Darwin trials. Same coordinator/session, serialized integration, no push or global cleanup.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE1OjM2OjU5LjU4NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcGhhc2UwOS9zdW1tYXJ5LzEiLCJpbnRlbnQiOiJmOWMzNjJmM2FmOGNmZDUzY2VmN2E2ZTIwZmM3YTUzMTI1M2E0YmU3OTIzNjc4NmFjNzZhZmFiZDFjZGRiYTkzIn0 -->
### 2026-11-02T15:36:59.585Z · summary · unspecified

{"acceptedScope":"Phase09: Native popup tab return","approvalSummary":["User requests removing Zoom and corrects bug scope to popup visibility on tab return. Phase09 implementation/source review and bounded actual owned-tab/focus/navigation/dismissal check are authorized under retained delegation; no new client-selection policy.","All exact native child summaries bind current independent source approval and delivered target commit/tree; ancestry verified."],"artifactRefs":[".mtn-tmp/tab-recovery/phase09-final/validate.json",".mtn-tmp/tab-recovery/phase09-final/worktrees.txt",".mtn-tmp/tab-recovery/runtime/rex-whichkey-p09-t01-builder-a01/result.json"],"attempt":1,"coverage":{"REQ-021":["tasks/rex-whichkey-09-01-tab-return.md"],"REQ-022":["tasks/rex-whichkey-09-01-tab-return.md"]},"eventId":"rex-whichkey-p03-coordinator-a01/phase09/summary/1","files":["README.md","defaults.go","docs/acceptance.md","docs/appearance.md","docs/launch-contract.md","lua/rex_whichkey/defaults/panes.lua","popup.go","terminal.go"],"kind":"summary","limitations":["No tests/TDD/general suites, destructive selections, configured Lua callbacks, config changes/reloads or remote/Darwin trials. Native app root/submenu tab-round-trip and focus reports were checked through server menu/process/layer state. Original remote graphical disappearance was not reproduced and the remote app presentation was not directly observed; no full GUI repair certificate. Several-app launch guard is unchanged.","Independent fresh contexts used the same authenticated model family, not cross-family review.","Prior transferred coordinator process status remains unverified; ownership transfer did not assert remote/local stoppage.","No global cleanup or push; retained worktrees/artifacts remain available."],"recordRef":"phases/rex-whichkey-09-tab-return.md","records":["tasks/rex-whichkey-09-01-tab-return.md"],"requirements":["REQ-021","REQ-022"],"schemaVersion":1,"specHash":"sha256:09e27ea5551f6f2f7417863af724d156049413d7452b7d608d111ee7ab61e7b8","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"4e03865091daff8a96144f878f959efddcc4d8b5","targetTree":"4cd9f3b8fbd551b872dd3069dddf78897621f5a2","verification":[{"command":"Actual native popup excludes Zoom; focus1004 enabled; native-app root/submenu tab-away/back and fragmented focus notifications retained same process/layer/menu; Escape exit0/exact owned cleanup","result":"pass"},{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./...; make build (combined delivered revision and refreshed CLI artifact)","result":"pass"},{"command":"git diff --check; native graph/schema validation; exact completed child scope and Git delivery identities/ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
