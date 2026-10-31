---
workflowKind: task
workflowLanguage: go
workflowOrder: 4
workflowRequirements:
  - REQ-002
  - REQ-012
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - rex/hosts.go
    - rex/hosts_test.go
    - rex/testdata/hosts
  sharedPaths: []
  resources: []
  interfaces:
    - rex-host-inventory
title: Decode effective Rex-known host inventory
mtnCreateRetry:
  key: rex-whichkey/draft-task/02-04-host-inventory/1
  intent: f32ba3182e5c6d7e88394a0c32ec9f399618af8fce884641f8a2c5456977abbe
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:15.389Z'
dateModified: '2026-10-31T11:26:20.569Z'
tags:
  - task
type: task
id: ea04ff70-e7da-4e95-994e-6a10afc78702
projects:
  - '[[phases/rex-whichkey-02-rex-client.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-02-01-cli-execution.md]]'
    reltype: FINISHTOSTART
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. This planning update does not execute implementation or approve provisional IF-* signatures.

## Deliverable and scope

Wrap rex hosts --json as the effective CLI inventory, not raw hosts.list; follow documented inventory ownership and keep it separate from retargetable server operations.


## Functional criteria

REQ-002, REQ-012, REQ-013
Decode stable IDs/labels, nested host records, warnings, held/cached-home state, empty inventory and malformed output. Do not pretend --server retargets rex hosts; do not add hosts, enumerate Tailscale or probe every endpoint.


## Ownership

Allowed source/test edits: rex/hosts.go, rex/hosts_test.go, rex/testdata/hosts. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve existing files and evidence. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. This request updates plans only; no agents or implementation are dispatched.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDExOjI1OjExLjg5MloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3doaWNoa2V5LWZpcnN0L3ByZXNlcnZlL3Rhc2tzX19yZXgtd2hpY2hrZXktMDItMDQtaG9zdC1pbnZlbnRvcnkubWQiLCJpbnRlbnQiOiJkODBmZjJjYmJlZTY3YTk2OGI2MTZhMDg1MjI3Y2ZkMGVhM2QwOTc4NDc5NDkwNGY5NDhhNGQxM2Y1OTU0ZGQ0In0 -->
### 2026-10-31T11:25:11.892Z · historical · unspecified

{"schemaVersion":1,"eventId":"rex-whichkey/whichkey-first/preserve/tasks__rex-whichkey-02-04-host-inventory.md","kind":"historical","recordRef":"tasks/rex-whichkey-02-04-host-inventory.md","specHash":"sha256:5c9bf786c22cc609f5b3674eac8b0c7d3babca33bfcc825cd6c651908246eb42","attempt":1,"content":"# Draft task\n\nDRAFT ONLY: user approved creation of this decomposition, not final interfaces or implementation dispatch. Governing decision REXWHICHKEY-20261030-094737 is Accepted for draft creation. PRD evidence repair and fingerprint-bound design approval remain gates. Paths and contracts below are provisional planning reservations, not accepted public APIs; reassess them after the RFC and republish fingerprints before dispatch. Follow the accepted flat concrete Go architecture, not generic hexagonal defaults. Go workflow language denotes this Go project; Lua/Nix work also requires its specialist checks. Workers never edit .tasks; coordinator alone publishes design/approval evidence.\n\n## Deliverable and scope\nWrap rex hosts --json as the effective CLI inventory, not raw hosts.list; follow approved IF-HOSTS ownership and keep it separate from retargetable server operations.\n\n## Requirement criteria\nREQ-002, REQ-012, REQ-013\nDecode stable IDs/labels, nested host records, warnings, held/cached-home state, empty inventory and malformed output. Do not pretend --server retargets rex hosts; do not add hosts, enumerate Tailscale or probe every endpoint.\n\n## Ownership\nAllowed source/test edits: rex/hosts.go, rex/hosts_test.go, rex/testdata/hosts. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.\n\n## Inputs, outputs and prerequisites\nConsume the approved RFC and relevant frozen producer contracts. Output the behavior and tests above. Explicit task prerequisites: tasks/rex-whichkey-02-01-cli-execution.md. Check native parent-phase dependencies too; membership alone does not confer readiness.\n\n## TDD and verification\nRED/GREEN effective-inventory fixtures and read-only real Rex discovery; root quality gates. Each code change includes its own failing-first tests, GREEN evidence and refactor checks; do not split tests into a separately completed task. Root Go gates are task build, task test (unit and race), task lint. Lua requires its isolated real-Rex/config fixtures; Nix requires package/module evaluation.\n\n## Integration and completion\nRequire independent Go-project combined specification/quality review plus relevant Lua/Nix expertise. Coordinator integrates reviewed results serially, runs affected checks on the combined target and publishes evidence. Final-phase gates belong to the parent milestone; unavailable GUI/remote infrastructure remains unmet. No implementation, dispatch, completion or interface adoption is authorized by this draft.","provenance":"Exact description before user-directed implementation-first replanning REXWHICHKEY-20261031-072107; previous evidence remains untouched.","gateEligible":false}
<!-- mtn:end-entry:v1 -->
