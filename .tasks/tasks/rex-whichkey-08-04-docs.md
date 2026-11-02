---
workflowKind: task
workflowLanguage: docs
workflowOrder: 4
workflowRequirements:
  - REQ-018
  - REQ-019
  - REQ-020
workflowDocuments:
  - ref: designs/rex-whichkey-popup-polish.md
    hash: 'sha256:ce42e479546a1181ce61138f0f5d37c88e5ed18ae3b6c39987fca90d89aee8f0'
workflowOwnership:
  paths:
    - README.md
    - docs/launch-contract.md
    - docs/acceptance.md
    - docs/appearance.md
  sharedPaths: []
  resources: []
  interfaces:
    - popup-polish-runtime-assembly
title: Document and verify assembled popup polish
mtnCreateRetry:
  key: rex-whichkey-phase08-task-4-v1
  intent: c46f2507c578373946f4ec7774ff6ea585b5e92cb1d6cadb3cba2ba2baadac68
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T13:17:08.493Z'
dateModified: '2026-11-02T14:11:57.272Z'
tags:
  - task
type: task
id: 6aba6d79-bd10-45de-becf-4a6a0a2ea892
projects:
  - '[[phases/rex-whichkey-08-popup-polish.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-08-01-lifetime.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-08-02-defaults.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-08-03-theme.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-02'
---
# Document and assemble popup polish

REQ-018/019/020. Depends on reviewed/integrated lifecycle/defaults/theme tasks. Own README.md, docs/launch-contract.md, docs/acceptance.md and docs/appearance.md (new) only. Read actual delivered APIs and user-linked Neovim snapshot; no source/gate/test/config/dependency or managed edits.

Describe native idle-until-dismissal, retained configured-Lua authenticated expiry and clean expiry behavior; distinguish per-operation/cleanup timeouts from user idle. Remove Focus and automatic Rex-known-hosts from default key tables/setup wording while preserving custom focus mappings, origin-focus restoration and optional explicit host provider. Explain distinct key/icon/label semantic colors, yellow/cyan inspiration, terminal-native fallback without palette replies and NO_COLOR/ASCII behavior. Do not invent exact hex settings or per-item configuration absent from source. Docs must reflect actual implementation, not proposed helper names or temporal 'this revision still awaits review' claims.

Acceptance inventory separates fresh source review/actual compilation from runtime observations. The final coordinator gate must run normal popup in existing Linux Rex session, leave it idle35seconds, verify still alive with reduced menu and VT/HTML semantic color spans, navigate Tabs and Escape-dismiss with exit0/exact cleanup. Builders/reviewers do not execute it and must not preclaim its success; use durable source/evidence pointers rather than self-staling pending statuses. No destructive/native action test, configured Lua callback trial, Darwin/remote certification, config reload or unit/race/TDD/general suite. Coordinator compiles the frozen combined candidate and runs this user-authorized targeted UI gate before completing final assembly. Fresh documentation reviewer compares every statement/example with actual integrated source, including optional host/lifetime/no-color constraints. No worker agents/integration/source commits outside owned Markdown, managed writes or push. Return full protocol object and concise visible file pointer.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE0OjExOjU3LjI2MFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwOC10MDQtYnVpbGRlci1hMDIvc3VtbWFyeS1lcXVpdmFsZW5jZS1yZWNvdmVyeS8yIiwiaW50ZW50IjoiZDNmOTk0NDJiNWI2NjViMWEzYTVkOWQwMGFkZjRkYjI2Y2E2NTkyNzJmZmNjMTg3Yjg3MWMyMDM1MTE4OTYyYSJ9 -->
### 2026-11-02T14:11:57.260Z · summary · unspecified

{"acceptedScope":"Document and verify assembled popup polish","approvalSummary":["User requests popup lifetime/default/color polish under bounded delegation; exact verification lineage approved after external app disconnection.","Independent fresh docs review rex-whichkey-p08-t04-tasks-a02/review/2 APPROVED exact frozen build after R1 fix."],"artifactRefs":[".mtn-tmp/menu-polish/rex-whichkey-p08-t04-builder-a02/integration-equivalence-recovery/integration.json",".mtn-tmp/menu-polish/runtime/combined-source-check-2/result.json"],"attempt":2,"eventId":"rex-whichkey-p08-t04-builder-a02/summary-equivalence-recovery/2","files":["README.md","docs/acceptance.md","docs/appearance.md","docs/launch-contract.md"],"kind":"summary","limitations":["Final documentation-only candidate fresh display refused because graphical app disconnected; no new-frame/full-tree runtime pass claimed. Prior real UI evidence reused only after mechanically checking all tracked runtime source identical.","No project tests/TDD, destructive/native action execution, configured Lua callbacks, remote/Darwin runtime or configuration reloads/changes; independent same-family review."],"recordRef":"tasks/rex-whichkey-08-04-docs.md","requirements":["REQ-018","REQ-019","REQ-020"],"schemaVersion":1,"specHash":"sha256:d489e0bc7508abf668a3b063539b0d828f38d957dba53e9e87283392324974d2","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"6090379edb0d43f7d56b01b81cc413ac9cc5fa5b","targetTree":"c98f2454d8e4d9288353b1dfd26e6944f71f82da","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./...; task build in frozen final candidate","result":"pass"},{"command":"Reviewed documentation contents unchanged; git diff --check; target commit/tree/ancestry verified","result":"pass"},{"command":"All tracked runtime/config/build source identical to actual runtime-passed source; differences limited to four owned Markdown files/managed records","result":"pass"},{"command":"Actual make run idle35seconds/reduced defaults/VT and HTML foreground colors/Tabs/Escape/exit0/exact owned cleanup on identical runtime source","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
