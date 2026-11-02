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
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T17:33:42.834Z'
dateModified: '2026-11-02T17:52:55.983Z'
tags:
  - task
type: task
id: d74caad8-2278-458f-87b1-cd339cfe5b54
projects:
  - '[[phases/rex-whichkey-11-compact-presentation.md]]'
completedDate: '2026-11-02'
---
# Compact menu status borders and conditional pagination

REQ-024; produce MenuFrame.PreferredHeight interface for dependent geometry task. Read approved designs/rex-whichkey-compact-presentation.md, helios-design relevant colors/layout/accessibility guidance and actual render/layout/terminal callers. Own render.go and layout.go only. No tests/gates/dependencies/config/mapping schema or managed edits.

Remove redundant Page1/1 and its footer row; genuine multi-page page indicator and Tab navigation remain usable, preferably merged into bottom status bar rather than another line. Replace adjacent plain title/help rows with full-width status-like horizontal border bars separating title/breadcrumbs above, menu content, concise key hints below. No literal Keys: prefix in menu body. Use existing semantic roles and theme decoration/ASCII handling; no guessed dark background/palette mutation. Maintain confirmations/cancel hints, readable tiny viewports, content priority, zero-size transitions, width/grapheme clipping, no wrapping/alignment regressions, no color-only state. Avoid unnecessary blank padding contrary to compact height request.

Add MenuFrame.PreferredHeight int (natural row count at current width including actual border/status/padding), independent of current height/page clipping, no I/O and no navigation mutation. Existing Width/Height/Lines retain strict actual constrained-frame meaning. Shared geometry consumer will use this field: do not invent separate child-side schema/API. Document exact producer meaning in Go comments. Pagination capacity must use actual visible chrome rows with no circular single-page-footer bug, overflow/divide-by-zero or jumping capacity across pages. Do not touch terminal.go/controller/geometry; dependent task wires real publication and resizing.

Leaf formats/static/whitespace only, no heavy/compiler-backed checks, tests/TDD/general/unit/race suites, product/GUI/API/config/remote trials, user settings/reloads or push. Coordinator compiles reviewed combined producer, final dependent docs gate measures actual floating layer. Source-only hook-disabled commits, freeze clean head/tree, no integration/agents/mtn writes. Full protocol-v1 build at payloadFile and MTN_REX_RESULT_FILE; validation array actual pass objects, skipped work in limitations; final visible pointer only.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE3OjUyOjU1Ljk3MloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxMS10MDEtYnVpbGRlci1hMDIvc3VtbWFyeS8yIiwiaW50ZW50IjoiMzYzNzRlYWU1MWZhOTMwODcxMWU4NjJmNmU2NzU0MTk3Y2U2NzFmYjViYWYxZDRhYTBiMGE4ZjQ3Nzk0N2UzMSJ9 -->
### 2026-11-02T17:52:55.972Z · summary · unspecified

{"acceptedScope":"Render compact status borders and conditional pagination","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p11-t01-reviewer-a02/review/2 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/popup-presentation/rex-whichkey-p11-t01-builder-a02/integration/integration.json"],"attempt":2,"eventId":"rex-whichkey-p11-t01-builder-a02/summary/2","files":["layout.go","render.go"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-11-01-render-frame.md","requirements":["REQ-024"],"schemaVersion":1,"specHash":"sha256:139ba9d08e8cf3dcb82eb28aff52d1aac0f80a4dc641f607de6a8fd9b63790e2","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"f57dfa00f84c4fed0e56dfc5ccaa5f2af60d3b87","targetTree":"0d1a7aa375bc8b099b2100a344b1edf6ff4ef9dd","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
