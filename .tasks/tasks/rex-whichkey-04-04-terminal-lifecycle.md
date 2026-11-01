---
workflowKind: task
workflowLanguage: go
workflowOrder: 4
workflowRequirements:
  - REQ-001
  - REQ-008
  - REQ-009
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - terminal.go
    - terminal_test.go
    - testdata/terminal
  sharedPaths:
    - go.mod
    - go.sum
  resources: []
  interfaces:
    - terminal-runtime
title: 'Implement terminal input, resize and restoration lifecycle'
mtnCreateRetry:
  key: rex-whichkey/draft-task/04-04-terminal-lifecycle/1
  intent: d82c59aff1ee26718d5830592b0c1962c123b0224aa2e2ef6262b386648acc65
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:31.417Z'
dateModified: '2026-11-01T01:51:28.159Z'
tags:
  - task
type: task
id: e06bf853-b602-456e-a9dd-baa1418cdee2
projects:
  - '[[phases/rex-whichkey-04-menu-ui.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-04-02-layout.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-04-03-theme.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-10-31'
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Integrate the terminal backend with navigation/rendering, input decoding and resize/redraw; own cancellation and terminal-state restoration without a daemon.


## Functional criteria

REQ-001, REQ-008, REQ-009
Selection, q/Esc dismissal, Tab/Backspace, terminal replies, resize races, errors and catchable signals restore configured terminal modes. Document unavoidable SIGKILL limitations and preserve the documented transient Darwin PENDIN handling without asserting a new platform comparison or runtime pass. Compatible dependency versions remain pinned.


## Ownership

Allowed source/test edits: terminal.go, terminal_test.go, testdata/terminal, go.mod, go.sum. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.


## Explicit producer/consumer delivery boundary
Own actual palette acquisition/input integration using theme.go from prerequisite 04-03: send supported palette query bytes as normal product startup behavior, route the ordered input stream through PaletteDecoder before key/cancel handling, and apply returned metadata to the current theme while preserving key order and terminal restoration. This source wiring must exist, not just a future caller note. Do not execute palette probes, tests or the product during implementation; source review/compilation are not runtime verification.


## Explicit frame producer and loop/composition consumers
Own the actual production render/input loop using the delivered RenderMenu frame API: render initially and after navigation, resize and theme/palette changes; consume frame geometry/placement and Lines when writing the terminal, and pass frame.Pages to Menu.HandleKey. Integrate ordered palette decoding before key/cancel handling and restore terminal state on every exit. This is actual source wiring within terminal.go, not documentation or token calls. The executable composition owner 04-05 must invoke this loop; terminal component completion alone is not a claim that popup/main is already delivered.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDAxOjUxOjI4LjE0N1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNC10MDQtYnVpbGRlci1hMDMvc3VtbWFyeS8zIiwiaW50ZW50IjoiYmJmODZiOWY3N2RkYThmODVkZDhlYTJlNDMxOTlkNmI3Nzg1MWUxZDJlYmQyMTgzOWFjYjBhZDliMzQwOGFiNiJ9 -->
### 2026-11-01T01:51:28.147Z · summary · unspecified

{"acceptedScope":"Implement terminal input, resize and restoration lifecycle","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p04-t04-reviewer-a03/review/3 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p04-t04-builder-a03/integration/integration.json"],"attempt":3,"eventId":"rex-whichkey-p04-t04-builder-a03/summary/3","files":["go.mod","go.sum","terminal.go"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-04-04-terminal-lifecycle.md","requirements":["REQ-001","REQ-008","REQ-009"],"schemaVersion":1,"specHash":"sha256:4cc7af65a747f5df55e52b6758a34241238d77a2372fe0994f57d61f115bb431","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"0d8fc4c8c1a306742a30fa0cee3ff7800b0f65ca","targetTree":"48a293e32b5a3047e7983627cb05c0c906f3f583","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
