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
  state: building
  sourceEventID: rex-whichkey-p19-t03-builder-a01/runtime/launched
  changedAt: '2026-11-07T09:45:37.007Z'
title: 'Phase19.3: Normal source run, assets and honest delivery documentation'
mtnCreateRetry:
  key: rex-whichkey-p19/task-create/source-run-delivery/1
  intent: 8488c12b9be3ff93a27569cd55efd29974af00b5db87bdafc049a705e3597fad
status: in-progress
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T04:48:25.218Z'
dateModified: '2026-11-07T09:45:37.629Z'
tags:
  - task
type: task
id: 3610c9b7-cffc-4c96-932c-17331bcf52b7
projects:
  - '[[phases/rex-whichkey-19-lua-float-launch.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-19-02-direct-lua-lifetime.md]]'
    reltype: FINISHTOSTART
---
After delivered controller-free task2, deliver the actual source make run and normal package route. Current PATH make wrapper execs Task; modify Taskfile run/build routing only as explicitly authorized, preserve EVERY test/lint/hook quality definition. Add conventional Makefile delegating to Task for portable GNU make if needed. Normal make run MUST build/select this checkout Go binary, call shared Lua through documented rex do FILE --action path, create one float and run direct Go; not call old LaunchPopup, assume installed stale binary, require user action registration/config reload or use a fake/private probe.

Use scoped terminal PATH prepend for development fixed program name, same installed Lua implementation/keybinding, ordinary Go/Lua packaging. No Nix runtime executable-discovery policy, wrappers, new dependency/pin/flake.nix/profile activation, init takeover or user config edits. Update clear launch/migration/source-run instructions and direct lifecycle limitations/default native chrome; no pending/current-source success claims before actual evidence.

Source quality gates whitespace, existing Nix parser/formatter only without eval/build/activation; Go build/vet on combined source coordinator-owned, no unit/race/TDD/Lua/parser/FIFO harness/benchmark/synthetic delays. Tests/assertions/hook/lint config unchanged. Existing full lint62 failures remain honest.

MANDATORY coordinator actual make build/make run on reviewed combined checkout and rebuilt main, not dry-run/help/ACK. Capture exact resolved make wrapper/GNU path, Lua script/action/current Go program source, own origin IDs/popup IDs, root/palette/terminal size/keyboard/provider readiness, one benign exact owned rename, cancellation/dismissal. Verify no Go launching parent/copy/private transport, visible float retired before effect, list_blocks shows detached host eventually gone after Go exit and focus/terminal restored. Use only explicitly owned disposable resources, no broad/remote/destructive trials/user config reload. Report uncertainty/exit errors honestly; no physical-keybinding/installed-profile/platform claim from source. Actual artifacts feed task4 final evidence refresh so active docs not stale pending. This task does not implement deferred native chrome.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA5OjQ1OjIzLjg4MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiZGlzcGF0Y2giLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMTktdDAzL2Rpc3BhdGNoLzEiLCJpbnRlbnQiOiI4MDllMDZlZDc4MDI4MDQwOTllYzllYzNiNTVlMTFkYWM4Zjk1ZGUzMTM5ZWUzOTVlZDc5ZDQxZTIxYjFiODU5In0 -->
### 2026-11-07T09:45:23.881Z · dispatch · unspecified

{"allowedCommands":["owned source/whitespace","Nix static parse/format only, no eval/build/activation","shellcheck/shfmt on any actual new shell script","hooks-disabled owned commit"],"allowedScope":{"interfaces":["P19-SHARED-LUA-ENTRY","P19-NORMAL-SOURCE-RUN","P19-NORMAL-ASSETS","P19-REVISION-BOUND-RUNTIME"],"paths":["Taskfile.yml","Makefile","nix/package.nix","README.md","examples/README.md","docs/launch-contract.md","docs/runtime-bootstrap.md","docs/remote-compatibility.md","docs/session-owned-ui.md","docs/cli-gui-delegation.md","docs/tab-rename-input.md","docs/mapping-format.md","docs/appearance.md","docs/acceptance.md"],"resources":[],"sharedPaths":[]},"attempt":1,"baseCommit":"983f7df1565a9c6e72e640343d470b9d9a872c3f","branch":"work/p19-t03-builder-a01","constraints":["Task01/02 native delivery+review+actual public runtime/ancestry verified; actual make/package routing now","Language nix with mtn-tasks/nix support, not Go TDD defaults","No tests/harnesses/benchmark/live leaf/environment hook/Nix build/eval/config/activation/reload/remote/deploy/push","Preserve every Taskfile build/test/lint/hook/assertion/lint/dependency standard and dirty separate flake.nix","Coordinator alone mutations/integrates/actual make run with current checkout; not source/ACK runtime proof"],"eventID":"rex-whichkey-p19-t03/dispatch/1","kind":"dispatch","limits":{"builders":1,"cpu":14,"goMemoryLimit":"5GiB","heavyGates":1,"integrations":1,"reviewers":1,"totalMemoryGiB":6},"recordRef":"tasks/rex-whichkey-19-03-source-run-delivery.md","runID":"rex-whichkey-p19-t03-builder-a01","schemaVersion":1,"specHash":"sha256:7f59fb0e65b827f5486b3d20fefa9b23bf99d9ed20894bc520c2db72e0b113ea","targetBranch":"main","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t03-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA5OjQ1OjMyLjcyNFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDMtYnVpbGRlci1hMDEvcnVudGltZS9yZXNlcnZlZCIsImludGVudCI6ImU3N2U5Y2Y0ZjVlZTJjMGY5YTA3YzA4ODA5MGRmN2QwOGY0NzgyODIwYmQ3NWZkY2NjZjNkYWE5Y2FlNDJiYzgifQ -->
### 2026-11-07T09:45:32.724Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-15-code-owned-runtime.md","runID":"rex-whichkey-p15-coordinator-a01"},"dispatchEventID":"rex-whichkey-p19-t03/dispatch/1","eventID":"rex-whichkey-p19-t03-builder-a01/runtime/reserved","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"75ff9ad4-a558-4295-8b20-855dafe77a8f"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-19-lua-float-launch.md","piSessionID":"rex-whichkey-p19-t03-builder-a01","piSessionName":"p19-t03-builder-a01","recordRef":"tasks/rex-whichkey-19-03-source-run-delivery.md","rexSessionName":"rex-whichkey-main","rexWindowName":"p19-t03-builder-a01","role":"builder","runID":"rex-whichkey-p19-t03-builder-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"mtn-tasks","specHash":"sha256:7f59fb0e65b827f5486b3d20fefa9b23bf99d9ed20894bc520c2db72e0b113ea","stage":"reserved","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p19-t03-builder-a01","windowName":"p19-t03-builder-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t03-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA5OjQ1OjM2LjIwMVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDMtYnVpbGRlci1hMDEvcnVudGltZS9sYXVuY2hlZCIsImludGVudCI6ImI3MGU4MjJlZTkyNDk3OTIzMzdkNGM0MDllMWEwNjFiOGVjODYwOWI4NjIwN2U4YjljNDVhN2VjN2Q2MzRkNzEifQ -->
### 2026-11-07T09:45:36.201Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-15-code-owned-runtime.md","runID":"rex-whichkey-p15-coordinator-a01"},"dispatchEventID":"rex-whichkey-p19-t03/dispatch/1","eventID":"rex-whichkey-p19-t03-builder-a01/runtime/launched","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"75ff9ad4-a558-4295-8b20-855dafe77a8f"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-19-lua-float-launch.md","piSessionID":"rex-whichkey-p19-t03-builder-a01","piSessionName":"p19-t03-builder-a01","provider":{"displayName":"p19-t03-builder-a01","name":"rex","resources":{"block":"block:01a1b566-2add-7468-be43-ea1d1e8e948c","endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","session":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","window":"window:01a1b566-2ade-77e3-857d-fceddfab91df"}},"recordRef":"tasks/rex-whichkey-19-03-source-run-delivery.md","rexBlockID":"block:01a1b566-2add-7468-be43-ea1d1e8e948c","rexSessionID":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","rexSessionName":"rex-whichkey-main","rexWindowID":"window:01a1b566-2ade-77e3-857d-fceddfab91df","rexWindowName":"p19-t03-builder-a01","role":"builder","runID":"rex-whichkey-p19-t03-builder-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"mtn-tasks","specHash":"sha256:7f59fb0e65b827f5486b3d20fefa9b23bf99d9ed20894bc520c2db72e0b113ea","stage":"launched","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p19-t03-builder-a01","windowName":"p19-t03-builder-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t03-builder-a01"}
<!-- mtn:end-entry:v1 -->
