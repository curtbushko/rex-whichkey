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
      hash: 'sha256:a3657116206eb455b158b9727099b7c3859882d18f759171205eb830dda163c8'
implementation:
  language: go
  ownership:
    paths:
      - terminal.go
      - menu_updates.go
      - providers.go
      - direct_menu.go
      - cmd/main.go
      - launch.go
      - native.go
      - native_app_owner.go
      - native_resources.go
      - native_sessions.go
      - native_targets.go
      - native_text_entry.go
      - popup.go
      - popup_initial.go
      - popup_updates.go
      - runtime_bootstrap.go
      - rex/block.go
      - rex/layer.go
      - rex/session.go
      - lua/rex_whichkey.lua
      - examples/rex-whichkey-keybinding.lua
      - launch_legacy_support_test.go
      - popup_legacy_support_test.go
      - runtime_legacy_support_test.go
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
  state: blocked
  sourceEventID: rex-whichkey-p19-t02/lifetime-blocker/1
  changedAt: '2026-11-07T04:48:41.527Z'
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
dateModified: '2026-11-07T04:48:42.155Z'
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
BLOCKED for activation/dispatch by P19-DIRECT-LIFETIME; consume reviewed task19.1 only after the exact cleanup-order decision or a documented within-scope capability resolves it. Lua computes finite 0<w,h<=1 fractions and centered x/y, captures actual initiating window from exact original session/block, creates one sole-block floating terminal with explicit window, shell=none, command argv starting fixed rex-whichkey, exit.on_completion=true and quick_exit_threshold_ms=0. Forward original origin/endpoint/client intent as validated arguments distinct from Rex-provided own execution IDs. No focused-window fallback. Module launch and script-local documented rex do FILE --action entry share implementation; not user init registration/reload or a selected-effect callback registry. Go captures original placement/app generation before interaction, independently validates exact own sole-popup membership, consumes task19.1, joins workers/restores tty before retirement, revalidates original origin/destination/generation, and dispatches once through concrete argument-array Rex CLI calls. Refuse effects on cancellation, tty/join/ownership/detach/focus uncertainty; no uncertain retry. If explicitly accepted, detach only the exact own block, confirm detached placement and visible layer absence before focus/effects, then exit under native terminal policy; do not equate this with pre-effect host deletion. Plain binary execution uses the existing terminal or fails explicitly, never launches another float. Remove production parent/renderer launch, private popup command, self-image copy, lease/result/update IPC and geometry creation; preserve genuine shared strict validators and concrete Rex APIs still used. Preserve ALL existing assertions and fixture bytes. Move launchOutput and genuinely obsolete support into precise new _test.go support files if needed for compilation; no production stubs or deletion of assertions. Coordinator actual make run on assembled source must verify one Go process, exact origin effect/cancel, visible and detached cleanup for this task; static review cannot certify it.
<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA0OjQ4OjQwLjYxNloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYmxvY2tlciIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDIvbGlmZXRpbWUtYmxvY2tlci8xIiwiaW50ZW50IjoiMmY0ZTkzNmE4MWVhMGU3YWE1ZDA4ZDVkODRiMzM5YTkwM2M2NmQ4OTY3NDM3MjZmZWE0Y2YyZjI0YjBiNTQxNSJ9 -->
### 2026-11-07T04:48:40.616Z · blocker · unspecified

{"affectedRefs":["tasks/rex-whichkey-19-02-direct-lua-lifetime.md","tasks/rex-whichkey-19-03-source-run-delivery.md"],"attempt":1,"decisionID":"REXWHICHKEY-20261106-234820","eventID":"rex-whichkey-p19-t02/lifetime-blocker/1","evidence":".mtn-tmp/p19/lifetime-assessment.json; installed close/detach/exit schemas; current docs/launch-contract.md cleanup inventory.","kind":"blocker","reason":"Documented detach removes visible float but retains executing block through dispatch. Existing accepted all-owned-block cleanup-before-effect and suppress-on-cleanup-failure cannot be preserved by the same process; close kills Go. No permitted same-process host retirement API found.","recordRef":"tasks/rex-whichkey-19-02-direct-lua-lifetime.md","requiredDecision":"Explicitly accept visible-layer/TTY/worker/focus cleanup before effect with detached execution host retired after Go exits (post-exit cleanup cannot suppress prior effect), OR retain strict all-resource pre-effect rule and supply a supported same-process capability/approve a different architecture.","schemaVersion":1,"specHash":"sha256:94903530c10ca9ae9af67231b30b1927c1c2363eea6b3d2e483e811883c2415d"}
<!-- mtn:end-entry:v1 -->
