---
workflowKind: phase
workflowOrder: 10
workflowRequirements:
  - REQ-023
workflowDocuments:
  - ref: designs/rex-whichkey-app-independent-popup.md
    hash: 'sha256:7215d04e0b71c6cc6a8a593372ef03cec57ae6aeb16907a81f1f9407b5781b4c'
title: 'Phase10: App-independent popup display'
mtnCreateRetry:
  key: rex-whichkey-phase10-v1
  intent: 5b41ac2d723b7c4cd6210942fd8f9ae2e3acfddc83de08ff7bb97db0e70fe29a
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T16:14:34.788Z'
dateModified: '2026-11-02T16:30:27.709Z'
tags:
  - task
type: task
id: 708c5036-24cb-422e-b6f5-20952329a818
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy:
  - uid: '[[phases/rex-whichkey-09-tab-return.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-02'
---
# Phase10: App-independent native popup startup

Deliver REQ-023 from approved designs/rex-whichkey-app-independent-popup.md: remove app-target validation from basic menu display, retain it only where executable native selection requires a target, and preserve configured Lua binding/security. User explicitly requests this correction. One source/docs task, independent fresh review, coordinator exact combined compilation and bounded real normal/nonexistent-selector display/dismissal plus nonexistent-target refusal after cleanup. No broad tests, destructive action dispatch, configuration changes/reloads or arbitrary client selection. Existing coordinator/session/runtime identity and completed phases02–09 remain intact; no push/global cleanup.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE2OjMwOjI3LjY5N1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcGhhc2UxMC9zdW1tYXJ5LzEiLCJpbnRlbnQiOiI4ZTNhMDY3ZjAyNzk1NDZhMDM0ZGM2YzYyMmRlMzU2MDc5MDM5NTViNGM2OTE3ODU0MmRjMjdlMDI3ZTFiNDFmIn0 -->
### 2026-11-02T16:30:27.697Z · summary · unspecified

{"acceptedScope":"Phase10: App-independent popup display","approvalSummary":["User explicitly requests decoupling native popup display from GUI target selection. Phase10 implementation/fresh review and bounded normal/nonexistent-selector UI/refusal checks authorized; configured binding/action safety retained.","All exact native child summaries bind current independent source approval and delivered target commit/tree; ancestry verified."],"artifactRefs":[".mtn-tmp/startup-decoupling/phase10-final/validate.json",".mtn-tmp/startup-decoupling/phase10-final/worktrees.txt",".mtn-tmp/startup-decoupling/runtime/rex-whichkey-p10-t01-builder-a01/result.json"],"attempt":1,"coverage":{"REQ-023":["tasks/rex-whichkey-10-01-app-independent-popup.md"]},"eventId":"rex-whichkey-p03-coordinator-a01/phase10/summary/1","files":["README.md","docs/acceptance.md","docs/launch-contract.md","launch.go","native.go","popup.go"],"kind":"summary","limitations":["No tests/TDD/general suites, successful native/destructive dispatch, configured Lua callbacks, config changes/reloads or remote/Darwin trials. Normal and nonexistent-selector native display/dismissal and deliberately nonexistent-target selection refusal checked with recorded app counts, no app disconnections. Server capture is not remote GUI presentation proof. Native display no longer has a singleton-app guard; executable targeting and configured launch remain strict.","Independent fresh contexts used the same authenticated model family, not cross-family review.","Prior transferred coordinator process status remains unverified; ownership transfer did not assert remote/local stoppage.","No global cleanup or push; retained worktrees/artifacts remain available."],"recordRef":"phases/rex-whichkey-10-app-independent-popup.md","records":["tasks/rex-whichkey-10-01-app-independent-popup.md"],"requirements":["REQ-023"],"schemaVersion":1,"specHash":"sha256:6bdf93feca5f655089585868e7bdd56b5177a4689b6dcbc76f7ac5d7bfc2d16d","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"8a5f388444ae371a15ea89c4d54085d9f80180eb","targetTree":"bad2da66d266758ee52ac925dc9da3b8bd273565","verification":[{"command":"Actual normal and nonexistent-client-selector native popups opened/navigated/dismissed exit0; nonexistent-target selection refused only after owned popup cleanup; exact tracked runtime source verified","result":"pass"},{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./...; make build (combined delivered revision and refreshed CLI artifact)","result":"pass"},{"command":"git diff --check; native graph/schema validation; exact completed child scope and Git delivery identities/ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
