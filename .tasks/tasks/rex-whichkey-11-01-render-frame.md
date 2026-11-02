---
workflowKind: task
workflowLanguage: go
workflowOrder: 1
workflowRequirements:
  - REQ-024
workflowDocuments:
  - ref: designs/rex-whichkey-compact-presentation.md
    hash: 'sha256:10a72840e340a8bb53e744309fd518d564a9aac00a73364c3a554da06d05388f'
workflowOwnership:
  paths:
    - render.go
    - layout.go
  sharedPaths: []
  resources: []
  interfaces:
    - menu-frame-preferred-height
title: Render compact status borders and conditional pagination
mtnCreateRetry:
  key: rex-whichkey-phase11-task1-v1
  intent: 1312d67ab275a5bbfae5adc914bb4617b973419e2cce45c998d67d8051d90552
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T17:33:42.834Z'
dateModified: '2026-11-02T17:33:43.772Z'
tags:
  - task
type: task
id: d74caad8-2278-458f-87b1-cd339cfe5b54
projects:
  - '[[phases/rex-whichkey-11-compact-presentation.md]]'
---
# Compact menu status borders and conditional pagination

REQ-024; produce MenuFrame.PreferredHeight interface for dependent geometry task. Read approved designs/rex-whichkey-compact-presentation.md, helios-design relevant colors/layout/accessibility guidance and actual render/layout/terminal callers. Own render.go and layout.go only. No tests/gates/dependencies/config/mapping schema or managed edits.

Remove redundant Page1/1 and its footer row; genuine multi-page page indicator and Tab navigation remain usable, preferably merged into bottom status bar rather than another line. Replace adjacent plain title/help rows with full-width status-like horizontal border bars separating title/breadcrumbs above, menu content, concise key hints below. No literal Keys: prefix in menu body. Use existing semantic roles and theme decoration/ASCII handling; no guessed dark background/palette mutation. Maintain confirmations/cancel hints, readable tiny viewports, content priority, zero-size transitions, width/grapheme clipping, no wrapping/alignment regressions, no color-only state. Avoid unnecessary blank padding contrary to compact height request.

Add MenuFrame.PreferredHeight int (natural row count at current width including actual border/status/padding), independent of current height/page clipping, no I/O and no navigation mutation. Existing Width/Height/Lines retain strict actual constrained-frame meaning. Shared geometry consumer will use this field: do not invent separate child-side schema/API. Document exact producer meaning in Go comments. Pagination capacity must use actual visible chrome rows with no circular single-page-footer bug, overflow/divide-by-zero or jumping capacity across pages. Do not touch terminal.go/controller/geometry; dependent task wires real publication and resizing.

Leaf formats/static/whitespace only, no heavy/compiler-backed checks, tests/TDD/general/unit/race suites, product/GUI/API/config/remote trials, user settings/reloads or push. Coordinator compiles reviewed combined producer, final dependent docs gate measures actual floating layer. Source-only hook-disabled commits, freeze clean head/tree, no integration/agents/mtn writes. Full protocol-v1 build at payloadFile and MTN_REX_RESULT_FILE; validation array actual pass objects, skipped work in limitations; final visible pointer only.
