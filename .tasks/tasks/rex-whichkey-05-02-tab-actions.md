---
workflowKind: task
workflowLanguage: lua
workflowOrder: 2
workflowRequirements:
  - REQ-010
  - REQ-011
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - lua/rex_whichkey/defaults/tabs.lua
    - tests/integration/tabs
  sharedPaths: []
  resources: []
  interfaces:
    - default-tab-mappings
title: Implement tab navigation and lifecycle mappings
mtnCreateRetry:
  key: rex-whichkey/draft-task/05-02-tab-actions/1
  intent: 21cc59b6ed2729ec237e3e458ce76a7c1a8a9da970a37c9f1bee87e3a4de6e79
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:37.435Z'
dateModified: '2026-10-31T19:58:33.933Z'
tags:
  - task
type: task
id: 933c3fbf-b29d-4396-9406-0248166b3f66
projects:
  - '[[phases/rex-whichkey-05-parity-remote.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-04-05-popup-integration.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-03-05-mapping-format.md]]'
    reltype: FINISHTOSTART
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Implement tab new/previous/next/rename/goto/close, retaining the old numbered-tab behavior within the mapping model.


## Functional criteria

REQ-010, REQ-011
Tab arguments are grounded in actual actions, numbering is correct, and dispatch targets the original app context. Close confirmation/cancel is handled safely without touching user tabs.


## Ownership

Allowed source/test edits: lua/rex_whichkey/defaults/tabs.lua, tests/integration/tabs. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Authoritative format input

Consume mapping.go and docs/mapping-format.md from tasks/rex-whichkey-03-05-mapping-format.md after its native producer prerequisite is delivered. These Go-owned files define the shared mapping/configuration standard; do not define or copy an independent Lua/Nix/consumer schema. Executable configuration and callbacks remain in the originating Rex Lua runtime. Follow the existing functional requirements; raise a consequential new public interface or dependency choice for user approval instead of silently selecting it.

## Worker routing

Intended workflowLanguage: lua. Use the supported matching language/documentation worker and independent source-review route once the coordinator confirms the narrow tooling capability update. Do not dispatch through Go TDD gates or substitute Go metadata while support is unavailable. The project no-tests/no-experiments policy overrides worker defaults.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDExOjI1OjM5LjE2NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3doaWNoa2V5LWZpcnN0L3ByZXNlcnZlL3Rhc2tzX19yZXgtd2hpY2hrZXktMDUtMDItdGFiLWFjdGlvbnMubWQiLCJpbnRlbnQiOiIwODU1N2YwZWQ2OTgzNGYyMTZiNzhlN2E4YmYzNzJhMWFiY2JjODY4Zjc5N2ZjZmI1ZTk1OWJmOWJlOTliZmY1In0 -->
### 2026-10-31T11:25:39.165Z · historical · unspecified

{"schemaVersion":1,"eventId":"rex-whichkey/whichkey-first/preserve/tasks__rex-whichkey-05-02-tab-actions.md","kind":"historical","recordRef":"tasks/rex-whichkey-05-02-tab-actions.md","specHash":"sha256:8b775e9482d06dc339f2b25b2066858c0ff75c1b9e4eec861a55d76f156dea5f","attempt":1,"content":"# Draft task\n\nDRAFT ONLY: user approved creation of this decomposition, not final interfaces or implementation dispatch. Governing decision REXWHICHKEY-20261030-094737 is Accepted for draft creation. PRD evidence repair and fingerprint-bound design approval remain gates. Paths and contracts below are provisional planning reservations, not accepted public APIs; reassess them after the RFC and republish fingerprints before dispatch. Follow the accepted flat concrete Go architecture, not generic hexagonal defaults. Go workflow language denotes this Go project; Lua/Nix work also requires its specialist checks. Workers never edit .tasks; coordinator alone publishes design/approval evidence.\n\n## Deliverable and scope\nImplement tab new/previous/next/rename/goto/close, retaining the old numbered-tab behavior within the approved mapping model.\n\n## Requirement criteria\nREQ-010, REQ-011\nTab arguments are grounded in actual actions, numbering is tested, and dispatch targets the original app context. Close confirmation/cancel is verified without touching user tabs.\n\n## Ownership\nAllowed source/test edits: lua/rex_whichkey/defaults/tabs.lua, tests/integration/tabs. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.\n\n## Inputs, outputs and prerequisites\nConsume the approved RFC and relevant frozen producer contracts. Output the behavior and tests above. Explicit task prerequisites: tasks/rex-whichkey-04-05-popup-integration.md. Check native parent-phase dependencies too; membership alone does not confer readiness.\n\n## TDD and verification\nRED/GREEN default mapping fixtures and real disposable-session tab lifecycle/navigation checks; root quality gates. Each code change includes its own failing-first tests, GREEN evidence and refactor checks; do not split tests into a separately completed task. Root Go gates are task build, task test (unit and race), task lint. Lua requires its isolated real-Rex/config fixtures; Nix requires package/module evaluation.\n\n## Integration and completion\nRequire independent Go-project combined specification/quality review plus relevant Lua/Nix expertise. Coordinator integrates reviewed results serially, runs affected checks on the combined target and publishes evidence. Final-phase gates belong to the parent milestone; unavailable GUI/remote infrastructure remains unmet. No implementation, dispatch, completion or interface adoption is authorized by this draft.","provenance":"Exact description before user-directed implementation-first replanning REXWHICHKEY-20261031-072107; previous evidence remains untouched.","gateEligible":false}
<!-- mtn:end-entry:v1 -->
