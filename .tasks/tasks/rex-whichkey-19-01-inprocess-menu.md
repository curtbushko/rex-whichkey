---
identity:
  kind: task
  taskID: rex-whichkey-19-01-inprocess-menu
  uuid: a92251c4-eb13-4f3e-bff7-2ebda111e7a3
structure:
  order: 1
scope:
  requirements:
    - REQ-001
    - REQ-002
    - REQ-003
    - REQ-004
    - REQ-006
    - REQ-008
    - REQ-009
    - REQ-010
    - REQ-011
    - REQ-012
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
      - lua/rex_whichkey_launch.lua
      - menu.go
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
      - render.go
      - rex/block.go
      - rex/direct_popup.go
      - rex/layer.go
      - rex/session.go
      - rex/snapshot.go
      - runtime_bootstrap.go
      - runtime_legacy_support_test.go
      - terminal.go
      - text_entry.go
      - theme.go
    sharedPaths: []
    resources: []
    interfaces:
      - P19-DIRECT-LIFETIME
      - P19-FILTERED-ORIGIN-SNAPSHOT
      - P19-INPROCESS-MENU
      - P19-LUA-ORIGIN-ARGV
      - P19-OWN-PTY
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
  sourceEventID: rex-whichkey-p19-t01/summary/4
  changedAt: '2026-11-07T08:29:30.791Z'
title: 'Phase19.1: In-process menu and own-PTY producer'
mtnCreateRetry:
  key: rex-whichkey-p19/task-create/inprocess-menu/1
  intent: 5bc882008f001efa9d437bbc293dc95ddaebfeb84d5f7e670e0c74bc3173787f
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T04:48:21.351Z'
dateModified: '2026-11-07T08:29:31.411Z'
tags:
  - task
type: task
id: 148d116e-23cd-4431-9fa0-6568fe058ff6
projects:
  - '[[phases/rex-whichkey-19-lua-float-launch.md]]'
completedDate: '2026-11-07'
---
Implement a CURRENT USED vertical slice: direct Go in-popup entry plus shared Lua float creation/launcher, consuming original origin as data independently from popup REX_* IDs. Accepted decisions1a/1b now remove earlier blocked-consumer conflict. Use defaults .8/.3 with validated configurable fractions/centered placement; exact originating window, no focus fallback; default native chrome only. Shipped Lua fixed program rex-whichkey, shell=none, documented exit on_completion=true quick_exit_threshold_ms=0. A scoped development PATH prepend must select checkout binary, not arbitrary public executable nomination.

Go owns real in-process metadata/rename/provider menu state and exact original action targets, max2 bounded jobs/once-per-kind/errors, one UI owner, current palette/own PTY/SIGWINCH/fixed rectangle/pagination/Unicode/confirm/limits/input handling. Filter exact own popup from post-create pane snapshots/counts/targets. New functions/fields have real current CLI/Lua callers; no dormant framework. Default/native existing path can remain transitional until task2 removes it, but new direct entry MUST function now and not create a nested popup. Ambiguous no-argument/origin data cannot detach a caller's ordinary pane.

Implement accepted direct lifetime in this slice: validate exact own sole-block floating placement distinct from original IDs; cancel/join workers and restore terminal; detach exact own popup, verify visible layer retired/detached host, restore/revalidate original origin/app generation/focus then invoke selected CLI effect once. Uncertain detach/focus/restore/ownership suppresses effects; no retarget/retry. Rex retires detached host after Go exit, not before dispatch; do not claim Go can verify its own post-exit deletion. Dismissal/cancel/error also ends with no orphan block. Session-closing effects can destroy host; never invent successful exit or retry an uncertain mutation. No Go launcher parent/renderer/copy/private-IPC controller for this new path.

Only listed literal source/Lua paths. Preserve test assertions/current ordinary path until consumer switches; no production stubs or framework/Cobra/Viper/hexagonal/direct socket/callback control plane. Existing JSON/defaults/Go actions/palette and PATH-only CLI policy intact, no binary hashes/status/schema gate. Static/build required: changed-source gofmt, git diff --check, task build, task lint:vet with pinned14CPU/Go5GiB; full unchanged lint optional inherited62 diagnostic. No tests/harnesses/benchmark/fake delays/environment hooks/Nix activation/config/reload/live leaf run.

Coordinator integration check on exact reviewed combined source: use documented rex do FILE --action shared Lua entry (no user init load/reload) with explicit owned original origin/current built Go; observe direct popup/menu, colors/own-PTY sizing, dynamic rows/current targets/rename editing, benign rename exact owned tab, cancellation/dismissal and list_blocks before/after for detached leaks, visible retirement before effect. This is real application use, not a custom harness; literal make run final path is delivered/actually checked by task3. Do not call intermediate source full Phase19 completion or native chrome customization. Leave current controller removal to task2, no latent direct APIs awaiting blocked consumer.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA4OjI5OjI5Ljc0NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDEvc3VtbWFyeS80IiwiaW50ZW50IjoiMDU4MzBjMjE4MTYyM2Y2NjlmOGUwNTM2NDVmMjhkODJiODhkMjE1OGVjY2JhZWNlMjI4NDRjYzg4YWNiMjY3OSJ9 -->
### 2026-11-07T08:29:29.745Z · summary · unspecified

{"acceptedScope":"Current-used shared Lua action/module with centered .8/.3/default chrome/direct Go entry; in-process bounded menu/metadata/current targets/palette/own-PTY display, exact own-popup filtering; verified visible detach/terminal/worker/origin focus before once-only effect with approved host-through-dispatch/post-exit Rex cleanup.","approvalSummary":["Explicit user Phase19 execute/task/commit/actual-make-run authority plus1a native chrome default deferral and1b host-through-effect exception","Current design/phase/plan exact; fresh independent combined review4 approved exact source3262beaf, same-family not cross-family","Combined normal build/vet plus actual default root/dynamic/full-ID owned inline rename/visible retirement ordering and separate private --config pane-filter/edit-cancel/root-dismiss runs passed before target advance"],"attempt":4,"cleanup":{"eligibleStoppedCleanDeliveredWorktreesRemoved":true,"facade":"Official completion attempted/native delivery verified; immutable helper cleanup jq defect preserved, exact eligible clean stopped delivered worktrees removed manually.","ownedValidationWindowClosed":true,"retained":["p19-t01-builder-a01 blocked/no-change recovery window/worktree","p19-t01-integration-a02 failed real source run","p19-t01-integration-a03 failed real source run"],"runtimeReservationsReleased":true,"successfulWindowsAbsent":true},"eventID":"rex-whichkey-p19-t01/summary/4","files":["cmd/main.go","direct_menu.go","lua/rex_whichkey.lua","examples/rex-whichkey-keybinding.lua","providers.go","rex/direct_popup.go","terminal.go",".mtn-tmp/p19/resume/task1-fix4-build-result.json",".mtn-tmp/p19/resume/task1-fix4-review-result.json",".mtn-tmp/p19/resume/task1-public-evidence.json",".mtn-tmp/p19/resume/task1-observed-rex-whichkey",".mtn-tmp/p19/resume/task1-integration-2-failed.json",".mtn-tmp/p19/resume/task1-integration-3-failed.json"],"kind":"summary","limitations":["Current controller retirement and actual make/package route are authorized downstream tasks; native chrome customization deferred by user1a","No prohibited tests/harnesses/benchmark/config/deploy; fresh same-family review, not cross-family","Owned native resize captured but no page transition claim; final own-PTY resize evidence pending Task03"],"recordRef":"tasks/rex-whichkey-19-01-inprocess-menu.md","requirements":["REQ-001","REQ-002","REQ-003","REQ-004","REQ-006","REQ-008","REQ-009","REQ-010","REQ-011","REQ-012","REQ-013","REQ-016","REQ-017","REQ-025","REQ-028","REQ-029"],"schemaVersion":1,"specHash":"sha256:17541d5d15d1a5e1c5e85cb4321734b835099389f1139cd6d9739e2d6b9aa6be","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"3b9293334a498d65f459afc2173839d67562e174","targetTree":"6e5a5b5afcdfbd0874fdc0c75bcc09e1efc04552","verification":[{"artifacts":[".mtn-tmp/p19/resume/task1-fix4-integration-build.log",".mtn-tmp/p19/resume/task1-fix4-integration-vet.log"],"command":"Combined normal task build and task lint:vet with pinned budget","result":"pass"},{"command":"Frozen reviewed source equality/owned paths/whitespace; exact target-base recheck and fast-forward; reviewed head retained in main ancestry","result":"pass"},{"artifact":".mtn-tmp/p19/resume/task1-public-evidence.json","command":"ACTUAL combined-source shared Lua/direct Go root/dynamic/rename/cancel/dismiss/once-only exact owned effect and detached host cleanup","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
