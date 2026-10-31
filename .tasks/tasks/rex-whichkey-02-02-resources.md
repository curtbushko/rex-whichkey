---
workflowKind: task
workflowLanguage: go
workflowOrder: 2
workflowRequirements:
  - REQ-002
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - rex/session.go
    - rex/session_test.go
    - rex/window.go
    - rex/window_test.go
    - rex/testdata/resources
  sharedPaths: []
  resources: []
  interfaces:
    - rex-session-window
title: Implement concrete session and window resource operations
mtnCreateRetry:
  key: rex-whichkey/draft-task/02-02-resources/1
  intent: 816a3ab804adb4b4ff38adde3c1862ffa9f65691d3915f917b692031d7e23806
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:11.389Z'
dateModified: '2026-10-31T19:57:58.685Z'
tags:
  - task
type: task
id: 0f2c44dd-c914-4d76-941b-9089bf03282e
projects:
  - '[[phases/rex-whichkey-02-rex-client.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-02-01-cli-execution.md]]'
    reltype: FINISHTOSTART
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Implement only the Session/Window operations needed by the menu, preserving stable server/session identities and exposing ordinary concrete values.


## Functional criteria

REQ-002, REQ-013
Published schemas govern list/create/inspect and required window operations, empty collections and malformed responses. Returned resources retain targeting; method names must not invent API request fields. Client app actions remain distinct from server resource operations.


## Ownership

Allowed source/test edits: rex/session.go, rex/session_test.go, rex/window.go, rex/window_test.go, rex/testdata/resources. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Remote implementation responsibility

rex/session.go and rex/window.go retain the selected server and stable session identity on returned concrete resources. Never infer targeting from a client-local process or transient window membership. This existing runtime owner implements the behavior; 05-05 only documents it.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDExOjI1OjA2Ljk0NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3doaWNoa2V5LWZpcnN0L3ByZXNlcnZlL3Rhc2tzX19yZXgtd2hpY2hrZXktMDItMDItcmVzb3VyY2VzLm1kIiwiaW50ZW50IjoiNGNlZjg5NjVmMGIxOTI3NDlkOGZjYzlmNjM1YjdiNmU1ZDFmMjY3YzJjNTc3YTgzYWVhNjk0ZDM1MzZhMTQ4OCJ9 -->
### 2026-10-31T11:25:06.945Z · historical · unspecified

{"schemaVersion":1,"eventId":"rex-whichkey/whichkey-first/preserve/tasks__rex-whichkey-02-02-resources.md","kind":"historical","recordRef":"tasks/rex-whichkey-02-02-resources.md","specHash":"sha256:614cbea337f6fbc3a97d4c93eaad87a91ce11b140bbab1ee7875d8752bd94df5","attempt":1,"content":"# Draft task\n\nDRAFT ONLY: user approved creation of this decomposition, not final interfaces or implementation dispatch. Governing decision REXWHICHKEY-20261030-094737 is Accepted for draft creation. PRD evidence repair and fingerprint-bound design approval remain gates. Paths and contracts below are provisional planning reservations, not accepted public APIs; reassess them after the RFC and republish fingerprints before dispatch. Follow the accepted flat concrete Go architecture, not generic hexagonal defaults. Go workflow language denotes this Go project; Lua/Nix work also requires its specialist checks. Workers never edit .tasks; coordinator alone publishes design/approval evidence.\n\n## Deliverable and scope\nImplement only the approved Session/Window operations needed by the menu, preserving stable server/session identities and exposing ordinary concrete values.\n\n## Requirement criteria\nREQ-002, REQ-013\nPublished schema fixtures cover list/create/inspect and required window operations, empty collections and malformed responses. Returned resources retain targeting; method names must not invent API request fields. Client app actions remain distinct from server resource operations.\n\n## Ownership\nAllowed source/test edits: rex/session.go, rex/session_test.go, rex/window.go, rex/window_test.go, rex/testdata/resources. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.\n\n## Inputs, outputs and prerequisites\nConsume the approved RFC and relevant frozen producer contracts. Output the behavior and tests above. Explicit task prerequisites: tasks/rex-whichkey-02-01-cli-execution.md. Check native parent-phase dependencies too; membership alone does not confer readiness.\n\n## TDD and verification\nRED/GREEN schema/argv contract tests and controlled isolated-server resource checks; root quality gates. Each code change includes its own failing-first tests, GREEN evidence and refactor checks; do not split tests into a separately completed task. Root Go gates are task build, task test (unit and race), task lint. Lua requires its isolated real-Rex/config fixtures; Nix requires package/module evaluation.\n\n## Integration and completion\nRequire independent Go-project combined specification/quality review plus relevant Lua/Nix expertise. Coordinator integrates reviewed results serially, runs affected checks on the combined target and publishes evidence. Final-phase gates belong to the parent milestone; unavailable GUI/remote infrastructure remains unmet. No implementation, dispatch, completion or interface adoption is authorized by this draft.","provenance":"Exact description before user-directed implementation-first replanning REXWHICHKEY-20261031-072107; previous evidence remains untouched.","gateEligible":false}
<!-- mtn:end-entry:v1 -->
