---
workflowKind: task
workflowLanguage: go
workflowOrder: 4
workflowRequirements:
  - REQ-012
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - hosts.go
    - hosts_test.go
    - testdata/host-menu
  sharedPaths: []
  resources: []
  interfaces:
    - known-host-menu
title: Present Rex-known hosts and inventory warnings
mtnCreateRetry:
  key: rex-whichkey/draft-task/05-04-host-menu/1
  intent: eecf6aef9cb5c5b70d0734b67aecc7029d71e26772ff4447e5d59e6f5e891876
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:41.477Z'
dateModified: '2026-10-31T19:58:35.750Z'
tags:
  - task
type: task
id: e193860b-8de7-43d9-b666-4eaa86359806
projects:
  - '[[phases/rex-whichkey-05-parity-remote.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-02-04-host-inventory.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-04-05-popup-integration.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-03-05-mapping-format.md]]'
    reltype: FINISHTOSTART
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Present the effective Rex inventory in the menu according to documented inventory ownership, preserving labels/IDs and warnings without implementing remote attach/bootstrap.


## Functional criteria

REQ-012, REQ-013
Empty, cached/held-home, warning and duplicate-label values are intelligible and stable. No independent registry, proactive endpoint probes, host mutation, remote attach, WOL or SSH is introduced. Host selection behavior stays within the approved listing-only contract.


## Ownership

Allowed source/test edits: hosts.go, hosts_test.go, testdata/host-menu. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Authoritative format input

Consume mapping.go and docs/mapping-format.md from tasks/rex-whichkey-03-05-mapping-format.md after its native producer prerequisite is delivered. These Go-owned files define the shared mapping/configuration standard; do not define or copy an independent Lua/Nix/consumer schema. Executable configuration and callbacks remain in the originating Rex Lua runtime. Follow the existing functional requirements; raise a consequential new public interface or dependency choice for user approval instead of silently selecting it.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDExOjI1OjQ0LjE0M1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3doaWNoa2V5LWZpcnN0L3ByZXNlcnZlL3Rhc2tzX19yZXgtd2hpY2hrZXktMDUtMDQtaG9zdC1tZW51Lm1kIiwiaW50ZW50IjoiYmEyNjAwM2QxODJkOGRhYTk4ZTI5MDI3ZmM4OGYwZDU4ZmQ4NTgwNTY4NDc4ZGIzZGRlMjg2MGIxNDQ4ZmRlYiJ9 -->
### 2026-10-31T11:25:44.143Z · historical · unspecified

{"schemaVersion":1,"eventId":"rex-whichkey/whichkey-first/preserve/tasks__rex-whichkey-05-04-host-menu.md","kind":"historical","recordRef":"tasks/rex-whichkey-05-04-host-menu.md","specHash":"sha256:6d6b283df5c98f24626b29317123b75be4baaa5d573d1a035bc7d17497510891","attempt":1,"content":"# Draft task\n\nDRAFT ONLY: user approved creation of this decomposition, not final interfaces or implementation dispatch. Governing decision REXWHICHKEY-20261030-094737 is Accepted for draft creation. PRD evidence repair and fingerprint-bound design approval remain gates. Paths and contracts below are provisional planning reservations, not accepted public APIs; reassess them after the RFC and republish fingerprints before dispatch. Follow the accepted flat concrete Go architecture, not generic hexagonal defaults. Go workflow language denotes this Go project; Lua/Nix work also requires its specialist checks. Workers never edit .tasks; coordinator alone publishes design/approval evidence.\n\n## Deliverable and scope\nPresent the effective Rex inventory in the menu according to approved IF-HOSTS ownership, preserving labels/IDs and warnings without implementing remote attach/bootstrap.\n\n## Requirement criteria\nREQ-012, REQ-013\nEmpty, cached/held-home, warning and duplicate-label fixtures are intelligible and stable. No independent registry, proactive endpoint probes, host mutation, remote attach, WOL or SSH is introduced. Host selection behavior stays within the approved listing-only contract.\n\n## Ownership\nAllowed source/test edits: hosts.go, hosts_test.go, testdata/host-menu. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.\n\n## Inputs, outputs and prerequisites\nConsume the approved RFC and relevant frozen producer contracts. Output the behavior and tests above. Explicit task prerequisites: tasks/rex-whichkey-02-04-host-inventory.md, tasks/rex-whichkey-04-05-popup-integration.md. Check native parent-phase dependencies too; membership alone does not confer readiness.\n\n## TDD and verification\nRED/GREEN menu fixtures and actual read-only Rex-known host display; root quality gates. Each code change includes its own failing-first tests, GREEN evidence and refactor checks; do not split tests into a separately completed task. Root Go gates are task build, task test (unit and race), task lint. Lua requires its isolated real-Rex/config fixtures; Nix requires package/module evaluation.\n\n## Integration and completion\nRequire independent Go-project combined specification/quality review plus relevant Lua/Nix expertise. Coordinator integrates reviewed results serially, runs affected checks on the combined target and publishes evidence. Final-phase gates belong to the parent milestone; unavailable GUI/remote infrastructure remains unmet. No implementation, dispatch, completion or interface adoption is authorized by this draft.","provenance":"Exact description before user-directed implementation-first replanning REXWHICHKEY-20261031-072107; previous evidence remains untouched.","gateEligible":false}
<!-- mtn:end-entry:v1 -->
