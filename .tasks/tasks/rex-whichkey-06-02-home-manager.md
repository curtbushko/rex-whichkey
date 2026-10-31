---
workflowKind: task
workflowLanguage: nix
workflowOrder: 2
workflowRequirements:
  - REQ-014
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - nix/home-manager.nix
    - tests/packaging/home-manager
  sharedPaths:
    - flake.nix
  resources: []
  interfaces:
    - home-manager-options
title: Add opt-in Home Manager integration preserving user config
mtnCreateRetry:
  key: rex-whichkey/draft-task/06-02-home-manager/1
  intent: 2746c0f3d33d8f22c8690dbf0618c888dd73327961fc4f76c8b5eab367d53457
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:47.500Z'
dateModified: '2026-10-31T19:58:23.953Z'
tags:
  - task
type: task
id: ff24179c-3cd4-4976-a2da-b7644ed37c10
projects:
  - '[[phases/rex-whichkey-06-packaging-release.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-06-01-nix-package.md]]'
    reltype: FINISHTOSTART
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Provide an opt-in Home Manager module that makes packaged integration available without taking ownership of the user's existing Rex init.lua.


## Functional criteria

REQ-014
Existing init.lua and separate user mapping files remain byte-for-byte unchanged. Source must support disabled/enabled configurations and explicit require/setup wiring on the server host; independent Nix source review checks both branches without claiming evaluation or runtime success. No activation reload/restart or unmanaged file overwrite.


## Ownership

Allowed source/test edits: nix/home-manager.nix, tests/packaging/home-manager, flake.nix. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Worker routing

Intended workflowLanguage: nix. Use the supported matching language/documentation worker and independent source-review route once the coordinator confirms the narrow tooling capability update. Do not dispatch through Go TDD gates or substitute Go metadata while support is unavailable. The project no-tests/no-experiments policy overrides worker defaults.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDExOjI1OjUxLjc2OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3doaWNoa2V5LWZpcnN0L3ByZXNlcnZlL3Rhc2tzX19yZXgtd2hpY2hrZXktMDYtMDItaG9tZS1tYW5hZ2VyLm1kIiwiaW50ZW50IjoiZmU5NzhjZjUwN2MyNjY4YjJiYWYwNzQ1YTBlOTUyMTk5MzgxMDE0OTg0YTE4MTFmYzRiZDY5ODNmNmI2ZGFlNyJ9 -->
### 2026-10-31T11:25:51.769Z · historical · unspecified

{"schemaVersion":1,"eventId":"rex-whichkey/whichkey-first/preserve/tasks__rex-whichkey-06-02-home-manager.md","kind":"historical","recordRef":"tasks/rex-whichkey-06-02-home-manager.md","specHash":"sha256:5f51bffebeca708d4987f4c716faa6735ab4288125449a59f90395c53d408c11","attempt":1,"content":"# Draft task\n\nDRAFT ONLY: user approved creation of this decomposition, not final interfaces or implementation dispatch. Governing decision REXWHICHKEY-20261030-094737 is Accepted for draft creation. PRD evidence repair and fingerprint-bound design approval remain gates. Paths and contracts below are provisional planning reservations, not accepted public APIs; reassess them after the RFC and republish fingerprints before dispatch. Follow the accepted flat concrete Go architecture, not generic hexagonal defaults. Go workflow language denotes this Go project; Lua/Nix work also requires its specialist checks. Workers never edit .tasks; coordinator alone publishes design/approval evidence.\n\n## Deliverable and scope\nProvide an opt-in Home Manager module that makes packaged integration available without taking ownership of the user's existing Rex init.lua.\n\n## Requirement criteria\nREQ-014\nExisting init.lua and separate user mapping files remain byte-for-byte unchanged. Disabled/enabled configurations evaluate; documented explicit require/setup wiring works on the server host. No activation reload/restart or unmanaged file overwrite.\n\n## Ownership\nAllowed source/test edits: nix/home-manager.nix, tests/packaging/home-manager, flake.nix. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.\n\n## Inputs, outputs and prerequisites\nConsume the approved RFC and relevant frozen producer contracts. Output the behavior and tests above. Explicit task prerequisites: tasks/rex-whichkey-06-01-nix-package.md. Check native parent-phase dependencies too; membership alone does not confer readiness.\n\n## TDD and verification\nFailing-first coexistence/evaluation checks, nix flake check with supported outputs and Home Manager evaluation against a fixture user setup; root quality gates. Use nix skill. Each code change includes its own failing-first tests, GREEN evidence and refactor checks; do not split tests into a separately completed task. Root Go gates are task build, task test (unit and race), task lint. Lua requires its isolated real-Rex/config fixtures; Nix requires package/module evaluation.\n\n## Integration and completion\nRequire independent Go-project combined specification/quality review plus relevant Lua/Nix expertise. Coordinator integrates reviewed results serially, runs affected checks on the combined target and publishes evidence. Final-phase gates belong to the parent milestone; unavailable GUI/remote infrastructure remains unmet. No implementation, dispatch, completion or interface adoption is authorized by this draft.","provenance":"Exact description before user-directed implementation-first replanning REXWHICHKEY-20261031-072107; previous evidence remains untouched.","gateEligible":false}
<!-- mtn:end-entry:v1 -->
