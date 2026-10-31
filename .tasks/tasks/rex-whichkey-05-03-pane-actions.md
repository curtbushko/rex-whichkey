---
workflowKind: task
workflowLanguage: go
workflowOrder: 3
workflowRequirements:
  - REQ-010
  - REQ-011
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - lua/rex_whichkey/defaults/panes.lua
    - lua/rex_whichkey/defaults/app.lua
    - tests/integration/panes
  sharedPaths: []
  resources: []
  interfaces:
    - default-pane-app-mappings
title: Implement pane and application action mappings
mtnCreateRetry:
  key: rex-whichkey/draft-task/05-03-pane-actions/1
  intent: f839c8e11280684c6dfa28c19820fe7bd9551a1ff6f19705edc92453e1741934
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:39.458Z'
dateModified: '2026-10-31T11:26:28.395Z'
tags:
  - task
type: task
id: 8f9db26f-4413-4a4b-860b-99b399b6d2fd
projects:
  - '[[phases/rex-whichkey-05-parity-remote.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-04-05-popup-integration.md]]'
    reltype: FINISHTOSTART
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. This planning update does not execute implementation or approve provisional IF-* signatures.

## Deliverable and scope

Implement split, directional focus, zoom, sidebar, command palette and pane close with origin-aware action dispatch.


## Functional criteria

REQ-010, REQ-011
Do not confuse pane.split up/down action arguments with CLI above/below. Implement split/focus/window behavior and destructive close cancellation/confirmation using origin-aware API targeting.


## Ownership

Allowed source/test edits: lua/rex_whichkey/defaults/panes.lua, lua/rex_whichkey/defaults/app.lua, tests/integration/panes. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve existing files and evidence. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. This request updates plans only; no agents or implementation are dispatched.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDExOjI1OjQxLjY1NFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3doaWNoa2V5LWZpcnN0L3ByZXNlcnZlL3Rhc2tzX19yZXgtd2hpY2hrZXktMDUtMDMtcGFuZS1hY3Rpb25zLm1kIiwiaW50ZW50IjoiNWIzZjRkMDhhYTJlYTQxNzM2Nzk4NTI1ZTZlOWQ5NjQ5NDZlYmQ0NmEyZmEwMDM1NjdlMTJkYTljODZmMWYyNiJ9 -->
### 2026-10-31T11:25:41.654Z · historical · unspecified

{"schemaVersion":1,"eventId":"rex-whichkey/whichkey-first/preserve/tasks__rex-whichkey-05-03-pane-actions.md","kind":"historical","recordRef":"tasks/rex-whichkey-05-03-pane-actions.md","specHash":"sha256:0c48b1afe0c74452c1b08e1e4b5f44d3e2353ce364a34e889e0ddf5f091039a5","attempt":1,"content":"# Draft task\n\nDRAFT ONLY: user approved creation of this decomposition, not final interfaces or implementation dispatch. Governing decision REXWHICHKEY-20261030-094737 is Accepted for draft creation. PRD evidence repair and fingerprint-bound design approval remain gates. Paths and contracts below are provisional planning reservations, not accepted public APIs; reassess them after the RFC and republish fingerprints before dispatch. Follow the accepted flat concrete Go architecture, not generic hexagonal defaults. Go workflow language denotes this Go project; Lua/Nix work also requires its specialist checks. Workers never edit .tasks; coordinator alone publishes design/approval evidence.\n\n## Deliverable and scope\nImplement split, directional focus, zoom, sidebar, command palette and pane close with origin-aware action dispatch.\n\n## Requirement criteria\nREQ-010, REQ-011\nDo not confuse pane.split up/down action arguments with CLI above/below. Verify split/focus/window behavior and destructive close cancellation/confirmation in isolated app resources.\n\n## Ownership\nAllowed source/test edits: lua/rex_whichkey/defaults/panes.lua, lua/rex_whichkey/defaults/app.lua, tests/integration/panes. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.\n\n## Inputs, outputs and prerequisites\nConsume the approved RFC and relevant frozen producer contracts. Output the behavior and tests above. Explicit task prerequisites: tasks/rex-whichkey-04-05-popup-integration.md. Check native parent-phase dependencies too; membership alone does not confer readiness.\n\n## TDD and verification\nRED/GREEN argument/default mapping tests plus real-app split/focus/zoom/sidebar/palette/close evidence; root quality gates. Each code change includes its own failing-first tests, GREEN evidence and refactor checks; do not split tests into a separately completed task. Root Go gates are task build, task test (unit and race), task lint. Lua requires its isolated real-Rex/config fixtures; Nix requires package/module evaluation.\n\n## Integration and completion\nRequire independent Go-project combined specification/quality review plus relevant Lua/Nix expertise. Coordinator integrates reviewed results serially, runs affected checks on the combined target and publishes evidence. Final-phase gates belong to the parent milestone; unavailable GUI/remote infrastructure remains unmet. No implementation, dispatch, completion or interface adoption is authorized by this draft.","provenance":"Exact description before user-directed implementation-first replanning REXWHICHKEY-20261031-072107; previous evidence remains untouched.","gateEligible":false}
<!-- mtn:end-entry:v1 -->
