---
identity:
  kind: task
  taskID: rex-whichkey-19-02-direct-lua-lifetime
  uuid: 0179e681-6c80-4986-afed-82327b72a636
structure:
  order: 2
scope:
  requirements:
    - REQ-001
    - REQ-002
    - REQ-003
    - REQ-004
    - REQ-009
    - REQ-010
    - REQ-011
    - REQ-013
    - REQ-016
    - REQ-017
    - REQ-025
    - REQ-028
    - REQ-029
  documents:
    - ref: designs/rex-whichkey-lua-float-launch.md
      hash: 'sha256:f7e38d49287de319753732b09fe472a40dfadcef81e5245e2deac208e968ab70'
implementation:
  language: go
  ownership:
    paths:
      - cmd/controller_legacy_support_test.go
      - cmd/main.go
      - config.go
      - controller_legacy_support_test.go
      - direct_menu.go
      - examples/rex-whichkey-keybinding.lua
      - launch.go
      - launch_legacy_support_test.go
      - lua/rex_whichkey.lua
      - mapping.go
      - menu_updates.go
      - native.go
      - native_app_owner.go
      - native_resources.go
      - native_sessions.go
      - native_targets.go
      - native_text_entry.go
      - popup.go
      - popup_initial.go
      - popup_legacy_support_test.go
      - popup_updates.go
      - providers.go
      - rex/block.go
      - rex/controller_legacy_support_test.go
      - rex/direct_popup.go
      - rex/geometry.go
      - rex/layer.go
      - rex/session.go
      - rex/terminal.go
      - rex/wait.go
      - runtime_bootstrap.go
      - runtime_legacy_support_test.go
      - terminal.go
      - text_entry.go
    sharedPaths: []
    resources: []
    interfaces:
      - P19-INPROCESS-MENU
      - P19-FILTERED-ORIGIN-SNAPSHOT
      - P19-LUA-ORIGIN-ARGV
      - P19-DIRECT-LIFETIME
      - P19-SHARED-LUA-ENTRY
  gates:
    - ID: whitespace
      command: git diff --check
      required: true
    - ID: source-format
      command: test -z "$(gofmt -l ./*.go ./rex/*.go ./cmd/*.go)"
      required: true
    - ID: normal-build
      command: task build
      required: true
    - ID: go-vet
      command: 'task lint:vet'
      required: true
    - ID: full-lint-diagnostic
      command: task lint
      required: false
lifecycle:
  state: complete
  sourceEventID: rex-whichkey-p19-t02/summary/4
  changedAt: '2026-11-07T09:44:47.284Z'
title: 'Phase19.2: Lua float and controller-free direct Go lifecycle'
mtnCreateRetry:
  key: rex-whichkey-p19/task-create/direct-lua-lifetime/1
  intent: 8de180373eb195a56756314a9bca5e7f6d6c65f00a2718af593dc568256b0248
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T04:48:23.279Z'
dateModified: '2026-11-07T09:44:47.912Z'
tags:
  - task
type: task
id: 1b2c1e8e-409b-4e80-8f7a-132bb7326bd0
projects:
  - '[[phases/rex-whichkey-19-lua-float-launch.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-19-01-inprocess-menu.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-07'
---
After delivered task1 direct Go/Lua vertical slice, retire the OLD ordinary Go launch controller/renderer/image-copy/lease/result/update production plane. Wire the public direct entry intentionally; no no-argument/caller ambiguity may treat ordinary original pane as owned popup. Preserve task1 direct lifecycle/current behavior. Lua owns native float creation/sizing/default-chrome policy, Go menu/operations/palette only. No launching Go parent, hidden helper, renamed supervisor, registry/relay, private IPC/copy path still used by normal shared Lua flow. No Go origin-geometry startup sizing.

Trace real callers before removal. Keep live strict JSON/Unicode, concrete Rex resources/current app targeting/data validation/provider/rename/terminal code. Preserve ALL historical assertions/test files; where old controller symbols are necessary to compile unchanged tests, isolate genuine original support into approved *_test.go files, not production stubs. Do not remove unrelated legacy assets/dependencies or claim tests ran. Owned files only; add a precise literal support-file ownership request to coordinator if genuinely needed.

All original target/generation/freshness/confirm/no-retry and accepted lifecycle semantics retained: visible float/tty/workers/focus before selected effects, detached host retires after Go exit under Rex. No all-block-before-effects claim. Current callbacks/executable Lua/hooks remain deferred. Go normal build/vet/format/whitespace and fresh independent source/dead-production-path review required; optional inherited lint diagnostic. No tests/TDD/race/Lua/parser/FIFO harnesses/benchmark/environment/user config/profile/reload/remote/deploy/push.

Coordinator actual shared-Lua public run on combined source: direct program/no Go parent/copy/private transport, root/dynamic/inline exact owned rename, cancel/dismiss, terminal/focus/visible and detached cleanup. Actual final make run is task3 integration criterion, not a private ACK or compile substitute. Preserve separate dirty flake.nix, quality gates/history.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA5OjQ0OjQ2LjIyMloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDIvc3VtbWFyeS80IiwiaW50ZW50IjoiNWViOWZlYWQzYzVkYzg0ZWI1YmY0MGQ0ZDlkZTZjYmVkMTM4MDk2MGZkZGI2YzllODJkOWZhMTk3MWVhZjE3NCJ9 -->
### 2026-11-07T09:44:46.222Z · summary · unspecified

{"acceptedScope":"Retired production Go launching controller/renderer entry/selfimage/lifetime FIFO/result/update file IPC and Go startup geometry/child wait. Removed dead invocation/provider wire, test-only Completion/nativeRenameMapping/capability/resolved-config closure; kept live strict JSON/Unicode/resource/config/selection validators. Genuine original historical assertion support isolated *_test.go without assertion edits. Direct Lua/Go/accepted detached-host lifetime retained.","approvalSummary":["Original explicit Phase19 execute/commit/make-run authority; user1a defaults/deferred native chrome and1b detached host through action","Exact current design/plan and delivered prerequisite Task01","Independent findings001/002/003 corrected under precise scope; fresh same-family combined review4 approved exact6b2f674 source; combined build/vet and real menu/owned effect/cancel/resize/cleanup passed before target advance"],"attempt":4,"cleanup":{"eligibleStoppedCleanDeliveredWorktreesRemoved":true,"facade":"Official protocol/native completion verified; immutable helper cleanup jq defect preserved, eligible exact worktrees removed manually after stops/clean/ancestry proof.","ownedValidationWindowClosed":true,"retained":["p19-t02-reviewer-a01 negative001 recovery","p19-t02-reviewer-a02 negative002 recovery","p19-t02-reviewer-a03 negative003 recovery"],"runtimeReservationsReleased":true,"successfulWorkerWindowsAbsent":true},"eventID":"rex-whichkey-p19-t02/summary/4","files":["cmd/main.go","launch.go","native.go","popup.go","popup_initial.go","popup_updates.go","runtime_bootstrap.go","mapping.go","config.go","text_entry.go","providers.go","menu_updates.go","terminal.go","rex/layer.go","rex/block.go","rex/geometry.go","rex/terminal.go","rex/wait.go","controller_legacy_support_test.go","launch_legacy_support_test.go","popup_legacy_support_test.go",".mtn-tmp/p19/resume/task2-fix4-build-result.json",".mtn-tmp/p19/resume/task2-fix4-review-result.json",".mtn-tmp/p19/resume/task2-public-evidence.json",".mtn-tmp/p19/resume/task2-observed-rex-whichkey"],"kind":"summary","limitations":["Normal make route/packaging next Task03; default-native-chrome policy final docs next Task04, no customization implementation","No prohibited tests/harnesses/benchmark/user config/deploy; source and observed cases not universal coverage or timing SLA","Real terminal resize without GUI owner observed, not physical GUI resize/keybinding activation","Fresh same-family review, inherited full lint not green"],"recordRef":"tasks/rex-whichkey-19-02-direct-lua-lifetime.md","requirements":["REQ-001","REQ-002","REQ-003","REQ-004","REQ-009","REQ-010","REQ-011","REQ-013","REQ-016","REQ-017","REQ-025","REQ-028","REQ-029"],"schemaVersion":1,"specHash":"sha256:ae53292288a392fe5343bca5502c322409646efe9d7ff571a56ce23bd8d46f74","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"75c75d26923573a4b4b3a67f18cafdd3336fa82f","targetTree":"8a91f833344f3af0eb33c84170631111891959ab","verification":[{"artifacts":[".mtn-tmp/p19/resume/task2-integration-build.log",".mtn-tmp/p19/resume/task2-integration-vet.log"],"command":"Combined normal task build/task lint:vet/whitespace and frozen source/ownership equality","result":"pass"},{"artifact":".mtn-tmp/p19/resume/protected-source-manifest.txt","command":"All original protected assertion/test/hook/dependency blobs identical; target-base recheck and ancestry-preserving fast-forward","result":"pass"},{"artifact":".mtn-tmp/p19/resume/task2-public-evidence.json","command":"ACTUAL combined controller-free shared Lua/direct Go providers/edit cancellation/exact owned rename/live own-PTY resize/default root dismissal and pre-effect visible/post-exit host cleanup","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
