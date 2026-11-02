---
workflowKind: task
workflowLanguage: go
workflowOrder: 4
workflowRequirements:
  - REQ-032
workflowDocuments:
  - ref: designs/rex-whichkey-selected-window-splits.md
    hash: 'sha256:1d5e6297352b97b84e3ba7730e586130b4103ddb4ad32ed03d35b6ac2ad6e4f2'
workflowOwnership:
  paths:
    - native.go
    - README.md
    - docs/launch-contract.md
    - docs/acceptance.md
  sharedPaths: []
  resources: []
  interfaces:
    - native-selected-window-splits
title: Split chosen app current selected window
mtnCreateRetry:
  key: rex-whichkey-phase14-task04-v1
  intent: a04bccc3b58b818e52e50a33f522e4c4231658446a6865144650fb6cc6a1f654
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T20:38:16.882Z'
dateModified: '2026-11-02T20:50:45.749Z'
tags:
  - task
type: task
id: 691f0b64-4d0f-4e29-825c-5821070d795a
projects:
  - '[[phases/rex-whichkey-14-action-audit.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-14-01-action-target.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-02'
---
# Native splits operate on chosen app's currently selected window

REQ032 explicit latest user correction. Read designs/rex-whichkey-selected-window-splits.md and existing action audit design with later supplement taking precedence. Current native.go always passes origin session/block to rex do, forcing splits into launcher origin even when chosen app selects another tab. Fix ONLY native pane.split v/right and s/down to use chosen app current selected tab/focused pane. Never split popup terminal, launch-origin unconditionally, another arbitrary app or server shared active_window_id.

Use documented client.perform with context omitted (block_id/session_id absent), via exact-server UNscoped Rex client or equally proven CLI environment isolation; generic scoped Client.Call injects origin session. Retain scoped origin/session/app-generation validation immediately before one dispatch; selected action registration, retained original args down/right and capability/confirmation unchanged. Do not silently alter destructive native close targets or configured originalVM/auth lifetime. Full cleanup BEFORE dispatch, short RPC context after user choice, signal cancellation/no-idle/FIFO, late explicit app choice and strict --client behavior intact. No retries/fallback on uncertain effects. Error returns sanitized. No new API guessed/dependency/service/schema.

Source ownership native.go, README.md, docs/launch-contract.md, docs/acceptance.md. No defaults arg changes, generic Rex adapter changes, Lua/config changes or UI sizing changes. Leaves source/format/whitespace only; NO tests/TDD/general/unit/race/compiler/build/heavy lint/RPC/native GUI/action/settings/fonts/platform trials. Coordinator compiles frozen reviewed combined candidate and verifies launcher A versus chosen app selected B, B two known panes, s/v split focused B pane in expected direction, A unchanged/other B anchor unchanged and exact owned cleanup. Also normal same-window split/late chooser/cancel/missing-selector gate. Server active_window_id NOT app-selected evidence. Resource audit and GUI19 still incomplete; do not claim every action/project done.

Full protocol-v1 build exact record/spec/attempt/base/head/tree/dispatch; outcome ready-for-review, validation [{command,result:"pass"}] actual source checks only; limitations skipped gates. Copy FULL object to MTN_REX_RESULT_FILE. Source commit hooks disabled, no collection mutations/integration/subagents. Final concise pointer.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDIwOjUwOjQ1LjczN1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxNC10MDQtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiYjBjZTQ5YTRkNjJjNzZlOWU2MjU3MjUxMDdhNWM1NzFmODIzZGE3NTY5NjY5NmE0NDA2NDE3NTEyNGRlNmFkZCJ9 -->
### 2026-11-02T20:50:45.737Z · summary · unspecified

{"acceptedScope":"Native s/v split chosen app current selected window focused pane, not launch origin","approvalSummary":["Explicit latest selected-window correction","Fresh frozen independent Go review APPROVED unchanged product source"],"artifactRefs":[".mtn-tmp/action-audit/runtime/rex-whichkey-p14-t04-builder-a01-harness-recovery/result.json",".mtn-tmp/action-audit/selected-runtime/rex-whichkey-p14-t04-builder-a01-harness-recovery/result.json",".mtn-tmp/action-audit/rex-whichkey-p14-t04-builder-a01/integration-harness-recovery/integration.json",".mtn-tmp/action-audit/rex-whichkey-p14-t04-builder-a01/integration-harness-recovery/correction.json"],"attempt":1,"eventId":"rex-whichkey-p14-t04-builder-a01/summary/1","files":["README.md","docs/acceptance.md","docs/launch-contract.md","native.go"],"kind":"summary","limitations":["Original too-strict harness failure retained, no product change or review substitution","One current live app; earlier genuine two-app chooser gate separate, not fresh multi-app dispatch proof","Only split/resource checks;19 GUI outcomes/all25 verification and native chrome still unresolved; no general tests/settings/config/fonts/platform certificate"],"recordRef":"tasks/rex-whichkey-14-04-selected-window-splits.md","requirements":["REQ-032"],"schemaVersion":1,"specHash":"sha256:e45606dee536698ab72e65b3246de1ec7fdb375d1d3c80cdceb7eaf3e0fe50ad","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"7b5a40488353586dd90956bd792a1f7ddd4a3449","targetTree":"af20373b696b9c6df518aefc2228c75bc0dbefe7","verification":[{"command":"Combined candidate go build ./...; task build; reviewed source identity; git diff --check; FF ancestry","result":"pass"},{"command":"Actual available-app same-window split, strict invalid selector/no action and cleanup (optional multi-app cases per runtime limitations)","result":"pass"},{"command":"Actual s/v target app-selected B focused pane, launcher A layout unchanged, B anchor identity retained, exact owned cleanup","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
