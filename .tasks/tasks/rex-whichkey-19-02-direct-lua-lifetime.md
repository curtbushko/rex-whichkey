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
      - cmd/main.go
      - controller_legacy_support_test.go
      - direct_menu.go
      - examples/rex-whichkey-keybinding.lua
      - launch.go
      - launch_legacy_support_test.go
      - lua/rex_whichkey.lua
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
      - rex/direct_popup.go
      - rex/layer.go
      - rex/session.go
      - runtime_bootstrap.go
      - runtime_legacy_support_test.go
      - terminal.go
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
  state: planned
  sourceEventID: rex-whichkey-p19-t02/lifetime-blocker/1/resume/2
  changedAt: '2026-11-07T07:21:41.625Z'
title: 'Phase19.2: Lua float and controller-free direct Go lifecycle'
mtnCreateRetry:
  key: rex-whichkey-p19/task-create/direct-lua-lifetime/1
  intent: 8de180373eb195a56756314a9bca5e7f6d6c65f00a2718af593dc568256b0248
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T04:48:23.279Z'
dateModified: '2026-11-07T07:21:42.259Z'
tags:
  - task
type: task
id: 1b2c1e8e-409b-4e80-8f7a-132bb7326bd0
projects:
  - '[[phases/rex-whichkey-19-lua-float-launch.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-19-01-inprocess-menu.md]]'
    reltype: FINISHTOSTART
---
After delivered task1 direct Go/Lua vertical slice, retire the OLD ordinary Go launch controller/renderer/image-copy/lease/result/update production plane. Wire the public direct entry intentionally; no no-argument/caller ambiguity may treat ordinary original pane as owned popup. Preserve task1 direct lifecycle/current behavior. Lua owns native float creation/sizing/default-chrome policy, Go menu/operations/palette only. No launching Go parent, hidden helper, renamed supervisor, registry/relay, private IPC/copy path still used by normal shared Lua flow. No Go origin-geometry startup sizing.

Trace real callers before removal. Keep live strict JSON/Unicode, concrete Rex resources/current app targeting/data validation/provider/rename/terminal code. Preserve ALL historical assertions/test files; where old controller symbols are necessary to compile unchanged tests, isolate genuine original support into approved *_test.go files, not production stubs. Do not remove unrelated legacy assets/dependencies or claim tests ran. Owned files only; add a precise literal support-file ownership request to coordinator if genuinely needed.

All original target/generation/freshness/confirm/no-retry and accepted lifecycle semantics retained: visible float/tty/workers/focus before selected effects, detached host retires after Go exit under Rex. No all-block-before-effects claim. Current callbacks/executable Lua/hooks remain deferred. Go normal build/vet/format/whitespace and fresh independent source/dead-production-path review required; optional inherited lint diagnostic. No tests/TDD/race/Lua/parser/FIFO harnesses/benchmark/environment/user config/profile/reload/remote/deploy/push.

Coordinator actual shared-Lua public run on combined source: direct program/no Go parent/copy/private transport, root/dynamic/inline exact owned rename, cancel/dismiss, terminal/focus/visible and detached cleanup. Actual final make run is task3 integration criterion, not a private ACK or compile substitute. Preserve separate dirty flake.nix, quality gates/history.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA0OjQ4OjQwLjYxNloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYmxvY2tlciIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDIvbGlmZXRpbWUtYmxvY2tlci8xIiwiaW50ZW50IjoiMmY0ZTkzNmE4MWVhMGU3YWE1ZDA4ZDVkODRiMzM5YTkwM2M2NmQ4OTY3NDM3MjZmZWE0Y2YyZjI0YjBiNTQxNSJ9 -->
### 2026-11-07T04:48:40.616Z · blocker · unspecified

{"affectedRefs":["tasks/rex-whichkey-19-02-direct-lua-lifetime.md","tasks/rex-whichkey-19-03-source-run-delivery.md"],"attempt":1,"decisionID":"REXWHICHKEY-20261106-234820","eventID":"rex-whichkey-p19-t02/lifetime-blocker/1","evidence":".mtn-tmp/p19/lifetime-assessment.json; installed close/detach/exit schemas; current docs/launch-contract.md cleanup inventory.","kind":"blocker","reason":"Documented detach removes visible float but retains executing block through dispatch. Existing accepted all-owned-block cleanup-before-effect and suppress-on-cleanup-failure cannot be preserved by the same process; close kills Go. No permitted same-process host retirement API found.","recordRef":"tasks/rex-whichkey-19-02-direct-lua-lifetime.md","requiredDecision":"Explicitly accept visible-layer/TTY/worker/focus cleanup before effect with detached execution host retired after Go exits (post-exit cleanup cannot suppress prior effect), OR retain strict all-resource pre-effect rule and supply a supported same-process capability/approve a different architecture.","schemaVersion":1,"specHash":"sha256:94903530c10ca9ae9af67231b30b1927c1c2363eea6b3d2e483e811883c2415d"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA3OjIxOjQwLjY3NFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVzdW1lIiwicmV0cnlLZXkiOiJyZXgtd2hpY2hrZXktcDE5LXQwMi9saWZldGltZS1ibG9ja2VyLzEvcmVzdW1lLzIiLCJpbnRlbnQiOiJkYTk1MWIyMWJkNTgzZTcxNGE4NjE3NmYxZDgwMjA4NDE2YzY4MmViY2E5NTdkODUwMjA5NGE4ZDY4NTZiMGZkIn0 -->
### 2026-11-07T07:21:40.674Z · resume · unspecified

{"attempt":2,"blockerEventID":"rex-whichkey-p19-t02/lifetime-blocker/1","decisionEvidence":"1a - yes, accept it for now; 1b - yes","eventID":"rex-whichkey-p19-t02/lifetime-blocker/1/resume/2","kind":"resume","recordRef":"tasks/rex-whichkey-19-02-direct-lua-lifetime.md","resolution":"User explicitly approves1a default chrome deferral and1b host-through-dispatch exception. Replanned current-used Go/Lua vertical slice provides real consumer with widened literal ownership; no dormant API or unapproved ordering change. Prior source attempt stopped/no edits; original blocker history preserved.","schemaVersion":1,"specHash":"sha256:31a2ab209e370f04195026c90db164569768106bcb65d0dc92d07b18f4661ab5"}
<!-- mtn:end-entry:v1 -->
