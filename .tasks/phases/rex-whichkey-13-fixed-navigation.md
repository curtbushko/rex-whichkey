---
workflowKind: phase
workflowOrder: 13
workflowRequirements:
  - REQ-029
workflowDocuments:
  - ref: designs/rex-whichkey-fixed-navigation.md
    hash: 'sha256:7fbb1596666f9acf0ddb56ecf00ff29c09c8ae26781577268138e6a4ca05177e'
title: 'Phase13: Stable popup navigation'
mtnCreateRetry:
  key: rex-whichkey-phase13-v1
  intent: 3b1be4c5362d98bd3d22378a8033e292fa0aa1aaf08c07f8183aa2eb79ae66f7
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T19:37:53.010Z'
dateModified: '2026-11-02T19:48:52.838Z'
tags:
  - task
type: task
id: e8ef574b-9d5e-4a88-a73b-66ae42d1e4be
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy:
  - uid: '[[phases/rex-whichkey-12-first-frame-fit.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-02'
---
# Phase13: Stable popup navigation

Explicit latest user correction REQ029: keep initial popup rectangle stable across submenu entries, use existing pagination rather than live resize. Preserve completed phase12 first-frame/spacers and native/configured safety; retire unnecessary live geometry IPC. One source/docs task depends on phase12, fresh review/combined compilation/actual identical-bound root/submenu/paging/back/idle/Escape check. No tests/settings/fonts changes/native selections or remote/Darwin trials. REQ026 native title/dock capability remains blocked; no full project completion or push.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE5OjQ4OjUyLjgyN1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcGhhc2UxMy9zdW1tYXJ5LzEiLCJpbnRlbnQiOiIzZTk0MWMyNjhmNDQ3Y2RmYTQzNTlhZDg3N2I3MWY2MGNkNmExZjBjMzE3N2ZjODMxNTNjMjMxYWExMWJhYWQzIn0 -->
### 2026-11-02T19:48:52.827Z · summary · unspecified

{"acceptedScope":"Phase13: Stable popup navigation","approvalSummary":["User explicitly requests no popup resizing per submenu. Phase13 fixed initial rectangle/content pagination supersedes navigation resizing while preserving first-frame fit/spacers. Frozen review/combined compilation/exact-bound navigation/cleanup checks authorized; chrome still blocked.","All exact native child summaries bind current independent source approval and delivered target commit/tree; ancestry verified."],"artifactRefs":[".mtn-tmp/popup-fixed-navigation/phase13-final/validate.json",".mtn-tmp/popup-fixed-navigation/phase13-final/worktrees.txt",".mtn-tmp/popup-fixed-navigation/runtime/rex-whichkey-p13-t01-builder-a01/result.json"],"attempt":1,"coverage":{"REQ-029":["tasks/rex-whichkey-13-01-fixed-navigation.md"]},"eventId":"rex-whichkey-p03-coordinator-a01/phase13/summary/1","files":["README.md","docs/acceptance.md","docs/appearance.md","docs/launch-contract.md","popup.go","popup_initial.go","popup_size.go","terminal.go"],"kind":"summary","limitations":["No tests/TDD/general suites, native/destructive selections, configured Lua callbacks, settings/font/config changes/reloads or remote/Darwin trials. Exact fixed bounds/navigation observed in one measured context, not pixel-perfect remote GUI proof. REQ026 title/dock controls stays blocked; no project completion.","Independent fresh contexts used the same authenticated model family, not cross-family review.","Prior transferred coordinator process status remains unverified; ownership transfer did not assert remote/local stoppage.","No global cleanup or push; retained worktrees/artifacts remain available."],"recordRef":"phases/rex-whichkey-13-fixed-navigation.md","records":["tasks/rex-whichkey-13-01-fixed-navigation.md"],"requirements":["REQ-029"],"schemaVersion":1,"specHash":"sha256:2660682d6a3d7f52a00745012de6b45549c888675e0469841b8947991fc9acc2","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"eb9d2c719d8cd8f6fdf871c0a255f5a2e6f303d5","targetTree":"2a8f3b1ee09076a8abd232cb161715c37800f073","verification":[{"command":"Actual compact initial rectangle/terminal seed/root spacers; EXACT same layer bounds through root/Tabs/paging/back/idle, Escape exit0/owned cleanup; tracked source equivalence verified","result":"pass"},{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./...; make build (combined delivered revision and refreshed CLI artifact)","result":"pass"},{"command":"git diff --check; native graph/schema validation; exact completed child scope and Git delivery identities/ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
