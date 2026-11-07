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
      hash: 'sha256:a3657116206eb455b158b9727099b7c3859882d18f759171205eb830dda163c8'
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
dateModified: '2026-11-07T04:48:29.055Z'
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
Consume task19.2's direct binary/shared Lua entry; no native-chrome prerequisite. Update only Taskfile build/run routing, adding a standard Makefile delegation if useful. Actual PATH make is the fixed Task wrapper: run must depend on checkout build and execute documented rex do FILE --action with explicit full origin IDs, never stale installed binary, old Go launcher, user action registration or init reload. Use only a scoped terminal path_prepend to checkout bin with fixed rex-whichkey program name; no REX_BIN_DIR or runtime Nix path policy. Preserve every existing test/lint/hook command, target, assertion and standard. Package ordinary Go/Lua assets and examples using the same launch implementation; no shell wrapper, activation, pin/lock change or init takeover. Update active docs to distinguish reviewed source, actual runtime evidence, pending native chrome and failure limits, preserving prior history and deferred hooks. Coordinator must perform actual normal make run from an explicitly owned real origin on the reviewed combined source and record source/binary identity, command, exact Rex endpoint/session/window/block/layer IDs, one Go process, own-PTY/palette/keyboard/resize/fixed navigation, dynamic readiness/rename editing, one benign owned rename, cancel/rapid cancel/normal dismissal, original focus and both visible-layer and detached-block disappearance. Attribute checks to affected source tasks rather than a separate test project/private ACK. No actual run until lifecycle blocker resolves. No chrome pass or full phase completion while task19.4 remains blocked. Commit reviewed source and honest docs only via coordinator; exclude parent flake.nix fix.