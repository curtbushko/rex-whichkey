---
workflowKind: task
workflowLanguage: docs
workflowOrder: 4
workflowRequirements:
  - REQ-024
  - REQ-025
  - REQ-026
workflowDocuments:
  - ref: designs/rex-whichkey-compact-presentation.md
    hash: 'sha256:10a72840e340a8bb53e744309fd518d564a9aac00a73364c3a554da06d05388f'
workflowOwnership:
  paths:
    - README.md
    - docs/appearance.md
    - docs/launch-contract.md
    - docs/acceptance.md
  sharedPaths: []
  resources: []
  interfaces:
    - compact-popup-appearance-assembly
title: Document and verify compact popup partial delivery
mtnCreateRetry:
  key: rex-whichkey-phase11-task4-v1
  intent: 0e988553cd9b84df4c43789ab727b74691bd7c79df4238af6e1c0ac3530c2c12
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T17:33:48.361Z'
dateModified: '2026-11-02T18:40:50.739Z'
tags:
  - task
type: task
id: c9952cbd-f828-4c45-85a3-0875a308033c
projects:
  - '[[phases/rex-whichkey-11-compact-presentation.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-11-01-render-frame.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-11-02-popup-height.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-02'
---
# Document and verify compact popup presentation (partial scope)

REQ-024/025 implementation and REQ-026 limitation documentation only. Depends on completed reviewed/integrated render and height tasks, not blocked chrome task. Own README.md, docs/appearance.md, docs/launch-contract.md and docs/acceptance.md. No source/test/gate/managed/user-config edits. Read exact producer/consumer code, approved designs/rex-whichkey-compact-presentation.md, Helios design and installed native API evidence.

Describe conditional page display, separate title/breadcrumb/key-hint status borders, actual content-fitting layer height using measured terminal/font geometry, preferred-row hint scope/cleanup, position/max/transient behavior, NO_COLOR/ASCII/tiny viewport constraints, unchanged app-independent startup and action/configured binding. Do not claim global font settings or per-popup chrome suppression unsupported by actual source. Explain native pane/tab/window controls remain a blocked Rex capability; no documented per-popup option found, no hidden flag/global changes, no user deferral accepted. Preserve historical proof versus current accepted requirements.

Coordinator gate before this task integration: exact combined Go/Taskfile compile; actual Linux Rex normal native root and nested submenu/back, recorded layer normalized bounds and terminal rows/cell pixels over settling/stability, status/border VT/HTML frames without Page1/1/Keys: and no giant fixed-height layer, Escape exit0 and exact owned cleanup. Must measure layer bounds, not just cropped RenderMenu text. No change of user font/settings or destructive/native selections to claim support; source-reviewed tiny/pagination/ASCII/no-color and font-change generality remain distinguished from actual one-context observations. Neither server screenshots nor this task certifies native GUI chrome removal. Builders/reviewers no heavy/product/API/GUI/config/remote experiments or tests/TDD/general suites. Do not preclaim future runtime passes or use self-staling pending-review prose.

Fresh independent docs review compares exact source/API/blocked scope. Source-only hook-disabled commits, freeze clean head/tree; no mtn writes/integration/agents/push. Full protocol-v1 JSON payload copied to MTN_REX_RESULT_FILE, validation array actual passing command/results, skipped work limitations; final visible pointer only. Phase/project must remain incomplete while chrome requirement remains unresolved.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE4OjQwOjUwLjcyN1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxMS10MDQtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiYzcxMTlhODI5MmY1NTZhNGFkODgxM2VkYmU5MzA4Yzg5YmNiMjQxZjVkYjAzNTc0MWNkMzY0ZTFlMmMyYzExOCJ9 -->
### 2026-11-02T18:40:50.727Z · summary · unspecified

{"acceptedScope":"Document and verify compact popup partial delivery","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p11-t04-tasks-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/popup-presentation/rex-whichkey-p11-t04-builder-a01/integration/integration.json",".mtn-tmp/popup-presentation/runtime/rex-whichkey-p11-t04-builder-a01/result.json"],"attempt":1,"eventId":"rex-whichkey-p11-t04-builder-a01/summary/1","files":["README.md","docs/acceptance.md","docs/appearance.md","docs/launch-contract.md"],"kind":"summary","limitations":["Compact root/submenu/back layer geometry/status bars/navigation/Escape/owned cleanup actually checked. Native pane controls remain an unresolved external capability, no user deferral or phase/project completion. No settings/font changes or general runtime/remote-GUI certificate.","No project tests/TDD, destructive actions, configured Lua callbacks, remote/Darwin or configuration changes/reloads; fresh same-family review."],"recordRef":"tasks/rex-whichkey-11-04-appearance-docs.md","requirements":["REQ-024","REQ-025","REQ-026"],"schemaVersion":1,"specHash":"sha256:b83d836886acc660a6d2750d40df204f2604c2dcead59c6c0e40d7164ef60141","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"178d3f80c1d7efeab898e17ae07c883439570f66","targetTree":"0b71f55d20052f408a1f9533dfe0ba4747b56969","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"},{"command":"task build; actual native compact root/nested submenu/back layer bounds and measured rows, separated status bars/no Page1/1, stable sizing, Escape/exit0/exact owned cleanup","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
