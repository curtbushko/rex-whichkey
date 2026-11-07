---
identity:
  kind: task
  taskID: rex-whichkey-19-03-source-run-delivery
  uuid: f4eb32c7-30a6-4ce2-baae-67036d7cf299
structure:
  order: 3
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
    - REQ-014
    - REQ-015
    - REQ-016
    - REQ-017
    - REQ-025
    - REQ-028
    - REQ-029
  documents:
    - ref: designs/rex-whichkey-lua-float-launch.md
      hash: 'sha256:f7e38d49287de319753732b09fe472a40dfadcef81e5245e2deac208e968ab70'
implementation:
  language: nix
  ownership:
    paths:
      - Taskfile.yml
      - Makefile
      - nix/package.nix
      - README.md
      - examples/README.md
      - docs/launch-contract.md
      - docs/runtime-bootstrap.md
      - docs/remote-compatibility.md
      - docs/session-owned-ui.md
      - docs/cli-gui-delegation.md
      - docs/tab-rename-input.md
      - docs/mapping-format.md
      - docs/appearance.md
      - docs/acceptance.md
    sharedPaths: []
    resources: []
    interfaces:
      - P19-SHARED-LUA-ENTRY
      - P19-NORMAL-SOURCE-RUN
      - P19-NORMAL-ASSETS
      - P19-REVISION-BOUND-RUNTIME
  gates:
    - ID: whitespace
      command: git diff --check
      required: true
lifecycle:
  state: complete
  sourceEventID: rex-whichkey-p19-t03/summary/1
  changedAt: '2026-11-07T14:33:46.501Z'
title: 'Phase19.3: Normal source run, assets and honest delivery documentation'
mtnCreateRetry:
  key: rex-whichkey-p19/task-create/source-run-delivery/1
  intent: 8488c12b9be3ff93a27569cd55efd29974af00b5db87bdafc049a705e3597fad
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T04:48:25.218Z'
dateModified: '2026-11-07T14:33:47.558Z'
tags:
  - task
type: task
id: 3610c9b7-cffc-4c96-932c-17331bcf52b7
projects:
  - '[[phases/rex-whichkey-19-lua-float-launch.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-19-02-direct-lua-lifetime.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-07'
---
After delivered controller-free task2, deliver the actual source make run and normal package route. Current PATH make wrapper execs Task; modify Taskfile run/build routing only as explicitly authorized, preserve EVERY test/lint/hook quality definition. Add conventional Makefile delegating to Task for portable GNU make if needed. Normal make run MUST build/select this checkout Go binary, call shared Lua through documented rex do FILE --action path, create one float and run direct Go; not call old LaunchPopup, assume installed stale binary, require user action registration/config reload or use a fake/private probe.

Use scoped terminal PATH prepend for development fixed program name, same installed Lua implementation/keybinding, ordinary Go/Lua packaging. No Nix runtime executable-discovery policy, wrappers, new dependency/pin/flake.nix/profile activation, init takeover or user config edits. Update clear launch/migration/source-run instructions and direct lifecycle limitations/default native chrome; no pending/current-source success claims before actual evidence.

Source quality gates whitespace, existing Nix parser/formatter only without eval/build/activation; Go build/vet on combined source coordinator-owned, no unit/race/TDD/Lua/parser/FIFO harness/benchmark/synthetic delays. Tests/assertions/hook/lint config unchanged. Existing full lint62 failures remain honest.

MANDATORY coordinator actual make build/make run on reviewed combined checkout and rebuilt main, not dry-run/help/ACK. Capture exact resolved make wrapper/GNU path, Lua script/action/current Go program source, own origin IDs/popup IDs, root/palette/terminal size/keyboard/provider readiness, one benign exact owned rename, cancellation/dismissal. Verify no Go launching parent/copy/private transport, visible float retired before effect, list_blocks shows detached host eventually gone after Go exit and focus/terminal restored. Use only explicitly owned disposable resources, no broad/remote/destructive trials/user config reload. Report uncertainty/exit errors honestly; no physical-keybinding/installed-profile/platform claim from source. Actual artifacts feed task4 final evidence refresh so active docs not stale pending. This task does not implement deferred native chrome.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDE0OjMzOjQ1LjA2MloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDMvc3VtbWFyeS8xIiwiaW50ZW50IjoiY2NlNDM2NWRlNWU0YWFkMDc5ZGRlZmIwZTczZTUzZGE1Njc3ZmFlNDY2YzViZDUzZTUxOTMwMTk1YzQyOGU5MCJ9 -->
### 2026-11-07T14:33:45.062Z · summary · unspecified

{"acceptedScope":"Normal source make run requires exact Rex origin, builds this checkout fixed program, invokes shared Lua script/action and scoped bin directory; conventional GNU Make delegates Task. Installed Go/Lua asset layout unchanged and package/current launch/migration docs coherent; build/test/lint/hook standards preserved. ACTUAL normal combined/main source makerun and GNU custom-config/bounds routes observed, exact owned rename and controlled cancel verified, no Go controller/copy/IPC.","approvalSummary":["Original user Phase19 execute/commit/ACTUAL make run and explicit1a default native chrome deferral/1b host through effect","Current approved exact design/serial plan and native delivered Task01/02 with actual source+cleanup proof","Fresh same-family mtn-tasks+nix combined review1 approved exactbaa6237; combined build/vet/source equality and actual combined/main makerun/owned effect/cancel/exit/cleanup passed before completion"],"attempt":1,"cleanup":{"controlledValidationWindowClosed":true,"eligibleStoppedCleanDeliveredWorktreesRemoved":true,"facade":"Official protocol/native completion verified despite immutable cleanup jq defect; exact eligible stopped clean delivered worktrees removed manually","gnuOriginalWindow":"Already absent with child exit0 before planned cancel; cause unobserved, raw evidence preserved; no extra close by stale ID","runtimeReservationsReleased":true,"successfulWorkerWindowsAbsent":true},"diagnosticValidation":[{"artifact":".mtn-tmp/p19/resume/current-main-full-lint.log","command":"Unchanged current-main task lint","detail":"Ignored lint config byte-identical, AI lint not reached; no full-lint pass or standards weakening","exitCode":201,"issues":53,"required":false,"result":"fail"}],"eventID":"rex-whichkey-p19-t03/summary/1","files":["Taskfile.yml","Makefile","nix/package.nix","README.md","examples/README.md","docs/appearance.md","docs/acceptance.md","docs/launch-contract.md","docs/runtime-bootstrap.md","docs/session-owned-ui.md","docs/cli-gui-delegation.md","docs/remote-compatibility.md","docs/tab-rename-input.md","docs/mapping-format.md",".mtn-tmp/p19/resume/task3-build-result.json",".mtn-tmp/p19/resume/task3-review-result.json",".mtn-tmp/p19/resume/task3-public-evidence.json",".mtn-tmp/p19/resume/task3-main-observed-rex-whichkey"],"kind":"summary","limitations":["Native chrome remains default by explicit1a, customization deferred.1b host survives action until exit/retirement, not all resources before effect","Physical keybinding/profile/Nix build/activation/remote/all-platform and all failure branches not verified; no prohibited tests/harnesses/benchmarks","Optional GNU-run closure cause unobserved, not credited controlled cancellation; separate new PATH config-cancel and MAIN run passed","Fresh same-family mixed-language source review, not cross-family"],"recordRef":"tasks/rex-whichkey-19-03-source-run-delivery.md","requirements":["REQ-001","REQ-002","REQ-003","REQ-004","REQ-006","REQ-008","REQ-009","REQ-010","REQ-011","REQ-012","REQ-013","REQ-014","REQ-015","REQ-016","REQ-017","REQ-025","REQ-028","REQ-029"],"schemaVersion":1,"specHash":"sha256:7f59fb0e65b827f5486b3d20fefa9b23bf99d9ed20894bc520c2db72e0b113ea","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"6564d07ef690d591ea842fa9d986d0bcfa919dd4","targetTree":"53a91c6bd08b4ceae15f860492092eee6af6a038","verification":[{"command":"Fresh exact-spec/base/head/tree Nix/Make/docs combined source review and static Nix formatter/parser","result":"pass"},{"artifacts":[".mtn-tmp/p19/resume/task3-integration-build.log",".mtn-tmp/p19/resume/task3-integration-vet.log"],"command":"Combined normal task build/task lint:vet/whitespace/frozen source equality/protected old assertion+hook+dependency blobs; target-base recheck and ancestry fast-forward","result":"pass"},{"artifact":".mtn-tmp/p19/resume/task3-public-evidence.json","command":"ACTUAL normal make run combined and rebuilt MAIN; GNU Make optional data/bounds; owned rename/cancel/root dismiss/current Go path/exit0/visible-before-effect+host/focus/terminal cleanup","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
