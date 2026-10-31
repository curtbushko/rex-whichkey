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
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:31.417Z'
dateModified: '2026-10-31T19:58:11.332Z'
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

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDExOjI1OjMxLjcwN1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3doaWNoa2V5LWZpcnN0L3ByZXNlcnZlL3Rhc2tzX19yZXgtd2hpY2hrZXktMDQtMDQtdGVybWluYWwtbGlmZWN5Y2xlLm1kIiwiaW50ZW50IjoiMTdiMjlmZmJlOTFiYzNkYmEzZWM2ZTYyNmE1NGIxZTQzNjcwMTQyMTRmZGQxZDA1NThjMGNhYjg3NDliNGM2YSJ9 -->
### 2026-10-31T11:25:31.707Z · historical · unspecified

{"schemaVersion":1,"eventId":"rex-whichkey/whichkey-first/preserve/tasks__rex-whichkey-04-04-terminal-lifecycle.md","kind":"historical","recordRef":"tasks/rex-whichkey-04-04-terminal-lifecycle.md","specHash":"sha256:0998b5520c95d0a30d2dd1436a853432c4c39eb68357034a4fede9bf0be1e63a","attempt":1,"content":"# Draft task\n\nDRAFT ONLY: user approved creation of this decomposition, not final interfaces or implementation dispatch. Governing decision REXWHICHKEY-20261030-094737 is Accepted for draft creation. PRD evidence repair and fingerprint-bound design approval remain gates. Paths and contracts below are provisional planning reservations, not accepted public APIs; reassess them after the RFC and republish fingerprints before dispatch. Follow the accepted flat concrete Go architecture, not generic hexagonal defaults. Go workflow language denotes this Go project; Lua/Nix work also requires its specialist checks. Workers never edit .tasks; coordinator alone publishes design/approval evidence.\n\n## Deliverable and scope\nIntegrate the approved terminal backend with navigation/rendering, input decoding and resize/redraw; own cancellation and terminal-state restoration without a daemon.\n\n## Requirement criteria\nREQ-001, REQ-008, REQ-009\nSelection, q/Esc dismissal, Tab/Backspace, terminal replies, resize races, errors and catchable signals restore configured terminal modes. Document unavoidable SIGKILL limitations; Darwin comparison may exclude only the verified transient PENDIN bit. Compatible dependency versions remain pinned.\n\n## Ownership\nAllowed source/test edits: terminal.go, terminal_test.go, testdata/terminal, go.mod, go.sum. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.\n\n## Inputs, outputs and prerequisites\nConsume the approved RFC and relevant frozen producer contracts. Output the behavior and tests above. Explicit task prerequisites: tasks/rex-whichkey-04-02-layout.md, tasks/rex-whichkey-04-03-theme.md. Check native parent-phase dependencies too; membership alone does not confer readiness.\n\n## TDD and verification\nFailing-first PTY tests with compiled application code, strict Unicode output, mode snapshots and injected replies/errors/signals; root quality gates. Each code change includes its own failing-first tests, GREEN evidence and refactor checks; do not split tests into a separately completed task. Root Go gates are task build, task test (unit and race), task lint. Lua requires its isolated real-Rex/config fixtures; Nix requires package/module evaluation.\n\n## Integration and completion\nRequire independent Go-project combined specification/quality review plus relevant Lua/Nix expertise. Coordinator integrates reviewed results serially, runs affected checks on the combined target and publishes evidence. Final-phase gates belong to the parent milestone; unavailable GUI/remote infrastructure remains unmet. No implementation, dispatch, completion or interface adoption is authorized by this draft.","provenance":"Exact description before user-directed implementation-first replanning REXWHICHKEY-20261031-072107; previous evidence remains untouched.","gateEligible":false}
<!-- mtn:end-entry:v1 -->
