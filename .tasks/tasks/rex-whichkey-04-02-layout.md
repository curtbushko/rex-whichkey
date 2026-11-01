---
workflowKind: task
workflowLanguage: go
workflowOrder: 2
workflowRequirements:
  - REQ-006
  - REQ-008
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - layout.go
    - layout_test.go
    - render.go
    - render_test.go
    - testdata/render
  sharedPaths:
    - go.mod
    - go.sum
  resources: []
  interfaces:
    - menu-layout
title: Implement constrained Unicode-aware layout and paging
mtnCreateRetry:
  key: rex-whichkey/draft-task/04-02-layout/1
  intent: 7268f8b6ef3d4ebd79893dee3b777c4a552bd616c12b2539e8133a51d35acd1d
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:27.393Z'
dateModified: '2026-11-01T01:29:51.297Z'
tags:
  - task
type: task
id: 003b96b8-24d9-4902-9e1f-5e3614f3a947
projects:
  - '[[phases/rex-whichkey-04-menu-ui.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-04-01-navigation.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-03-05-mapping-format.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-04-03-theme.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-10-31'
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Render menu state using the product terminal approach, adaptive columns and Unicode display/grapheme widths; compute stable paging with icons, breadcrumbs, help and footer.


## Functional criteria

REQ-006, REQ-008
Combining marks, wide characters and icons are preserved; no out-of-bounds cells for narrow/tiny/zero-resize transitions. All five positions have constrained layouts and stable Tab page navigation. Preserve original Unicode codepoints and correct display widths.


## Ownership

Allowed source/test edits: layout.go, layout_test.go, render.go, render_test.go, testdata/render. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Authoritative format input

Consume mapping.go and docs/mapping-format.md from tasks/rex-whichkey-03-05-mapping-format.md after its native producer prerequisite is delivered. These Go-owned files define the shared mapping/configuration standard; do not define or copy an independent Lua/Nix/consumer schema. Executable configuration and callbacks remain in the originating Rex Lua runtime. Follow the existing functional requirements; raise a consequential new public interface or dependency choice for user approval instead of silently selecting it.


## Explicit producer/consumer delivery boundary
Consume theme.go from 04-03 after its actual native producer prerequisite. Actual production rendering must route styled text and optional icons/ASCII fallback decorations through the theme's semantic Style/Decoration APIs, rather than merely documenting future use. Keep rendering/layout ownership here, not in the theme producer. Preserve existing Unicode, width, paging, plain-output and accessibility criteria; source review must verify real callers before this task and Phase4 complete.


## Routine Unicode dependency authority
Current user delegates routine implementation fixes within REQ-008. Shared go.mod/go.sum are reserved exclusively here while this task runs; terminal ownership is naturally serialized by its existing prerequisite. A pinned github.com/rivo/uniseg dependency may implement real grapheme segmentation/display widths after checking its versioned documentation/source, license and Go compatibility. Record actual pin/checksums and evidence, not an unverified selection. This does not authorize a Go–Lua dependency, new public product contract, tests/experiments or handwritten Unicode approximations. Ordinary dependency fetching is build preparation, not product execution.


## Explicit frame producer and loop/composition consumers
This is the frame-producing layout/rendering component. Its complete RenderMenu API must perform the accepted geometry, Unicode segmentation/width, theme Style/Decoration and paging behavior against the actual menu/theme inputs; no placeholders or token callers. The source owning the live render/input loop is 04-04 terminal.go, and 04-05 cmd/main.go composes that loop into the executable. They are downstream native consumers, not prerequisites of this producer. This task may complete after its usable implementation and independent source review, without falsely claiming those future consumers or graphical/runtime verification. Phase4 cannot complete without their real production callers.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDAxOjI5OjUxLjI4NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNC10MDItYnVpbGRlci1hMDMvc3VtbWFyeS8zIiwiaW50ZW50IjoiZTc0MzU0ZWQwNzM2MjY2M2U3NDc4MzcxOTc0MTVhMTIxMWJiMjJjZThjODc1ZjQzZDY0OWRjYmIwMDVhZDgxNCJ9 -->
### 2026-11-01T01:29:51.285Z · summary · unspecified

{"acceptedScope":"Implement constrained Unicode-aware layout and paging","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p04-t02-reviewer-a03/review/3 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p04-t02-builder-a03/integration/integration.json"],"attempt":3,"eventId":"rex-whichkey-p04-t02-builder-a03/summary/3","files":["go.mod","go.sum","layout.go","render.go"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-04-02-layout.md","requirements":["REQ-006","REQ-008"],"schemaVersion":1,"specHash":"sha256:0daf24062f07dc13a93994ce177553cc48a27d79ccb6f10c64e7fc55ce73b0c2","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"43c62df85e946a05fed2458db7ad4eedd51399ed","targetTree":"7bdcac0637671cabc50878bc5d28a97b24689c4e","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
