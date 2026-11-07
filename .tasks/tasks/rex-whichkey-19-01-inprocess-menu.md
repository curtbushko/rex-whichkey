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
    - REQ-003
    - REQ-004
    - REQ-006
    - REQ-008
    - REQ-009
    - REQ-010
    - REQ-011
    - REQ-012
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
      - menu.go
      - text_entry.go
      - render.go
      - theme.go
      - rex/snapshot.go
      - direct_menu.go
    sharedPaths: []
    resources: []
    interfaces:
      - P19-INPROCESS-MENU
      - P19-FILTERED-ORIGIN-SNAPSHOT
      - P19-OWN-PTY
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
  sourceEventID: null
  changedAt: '2026-11-07T04:48:20.704Z'
title: 'Phase19.1: In-process menu and own-PTY producer'
mtnCreateRetry:
  key: rex-whichkey-p19/task-create/inprocess-menu/1
  intent: 5bc882008f001efa9d437bbc293dc95ddaebfeb84d5f7e670e0c74bc3173787f
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T04:48:21.351Z'
dateModified: '2026-11-07T04:48:21.994Z'
tags:
  - task
type: task
id: 148d116e-23cd-4431-9fa0-6568fe058ff6
projects:
  - '[[phases/rex-whichkey-19-lua-float-launch.md]]'
---
Implement concrete in-process menu support without activating a new launcher/effect path. Reuse bounded metadataJobs: at most two workers, five once-only jobs and bounded immutable results; only the UI loop changes menu/retained targets. Add a concrete origin-aware menu entry that returns the retained config, MenuEvent and exact provider targets after cancel/join and terminal restoration. Keep existing public/standalone and current private-child behavior compatible until task19.2 removes the latter; no placeholder stubs or new transport framework. Read actual own PTY size/palette, retain SIGWINCH, fixed float rectangle/pagination, rename editing/readiness, confirmations, Unicode, defaults/JSON, limits and theme behavior. Derive original current markers from retained original IDs, not popup focus. Provide a filtered post-create origin snapshot and provider inventories excluding the exact own popup block, including counts/limits/targets. Shared raw view may be used solely for own-popup validation. No Go float geometry RPCs on the new path. Safe to prepare/review/build independently; do not claim direct-launch/runtime lifecycle delivery from this source-only producer. Coordinator's later combined normal make run must verify this task's menu contribution, with evidence attributed back to this task, before phase runtime acceptance. Existing RunTerminal/production callers must meaningfully exercise any extracted core. No exported origin-aware API with only a future blocked-task/test caller may be called delivered; if safe current integration cannot meet that invariant, return the exact blocker rather than creating dormant scaffolding. Current public result-only menu must never dispatch resources.