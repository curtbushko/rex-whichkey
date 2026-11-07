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
  state: planned
  sourceEventID: null
  changedAt: '2026-11-07T04:48:24.572Z'
title: 'Phase19.3: Normal source run, assets and honest delivery documentation'
mtnCreateRetry:
  key: rex-whichkey-p19/task-create/source-run-delivery/1
  intent: 8488c12b9be3ff93a27569cd55efd29974af00b5db87bdafc049a705e3597fad
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T04:48:25.218Z'
dateModified: '2026-11-07T07:21:23.451Z'
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
