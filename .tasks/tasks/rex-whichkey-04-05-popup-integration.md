---
workflowKind: task
workflowLanguage: go
workflowOrder: 5
workflowRequirements:
  - REQ-001
  - REQ-005
  - REQ-007
  - REQ-008
  - REQ-011
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - cmd/main.go
    - cmd/main_test.go
    - config.go
    - config_test.go
    - lua/rex_whichkey.lua
    - tests/integration/popup
  sharedPaths: []
  resources: []
  interfaces:
    - popup-launch-selection
title: 'Wire short-lived popup launch, selection and cleanup'
mtnCreateRetry:
  key: rex-whichkey/draft-task/04-05-popup-integration/1
  intent: a01287e2e4f673506cf505b15f463871777c77ed2da0b7a45f291635965e1b1c
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:33.418Z'
dateModified: '2026-10-31T19:58:32.104Z'
tags:
  - task
type: task
id: 5f1a721b-86f5-4524-8ed3-e46951b8d2f9
projects:
  - '[[phases/rex-whichkey-04-menu-ui.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-04-04-terminal-lifecycle.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-03-04-invocation-dispatch.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-02-03-block-layers.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-03-05-mapping-format.md]]'
    reltype: FINISHTOSTART
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Wire the short launcher and terminal-owned Go process to launch/config/selection contracts. Launcher returns promptly; Go owns menu navigation and exits on dismissal or dispatch.


## Functional criteria

REQ-001, REQ-005, REQ-007, REQ-008, REQ-011, REQ-013
Popup survives beyond the configuration action lifetime, carries exact origin identities, executes one validated selection and cleans only owned resources on cancellation/errors. Implement all five graphical positions and focus restoration; moved origins resolve safely. Independent source review must trace layout, focus and cleanup paths; graphical/runtime behavior remains unexecuted, not passed. No background daemon or user init.lua takeover.


## Ownership

Allowed source/test edits: cmd/main.go, cmd/main_test.go, config.go, config_test.go, lua/rex_whichkey.lua, tests/integration/popup. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Authoritative format input

Consume mapping.go and docs/mapping-format.md from tasks/rex-whichkey-03-05-mapping-format.md after its native producer prerequisite is delivered. These Go-owned files define the shared mapping/configuration standard; do not define or copy an independent Lua/Nix/consumer schema. Executable configuration and callbacks remain in the originating Rex Lua runtime. Follow the existing functional requirements; raise a consequential new public interface or dependency choice for user approval instead of silently selecting it.

## Remote implementation and mixed-language review

Own server-side popup executable/config resolution and propagation of exact originating server/session/block/client identities in config.go and cmd/main.go, with only required launcher glue in lua/rex_whichkey.lua. Use the execution host's installed executable/module paths, not a client-local absolute path; 03-01 owns user module loading, 03-04 owns callback/action validation and Phase 2 owns CLI/resource targeting. Preserve no-fallback origin behavior and cleanup only resources returned as owned. 05-05 documents these source paths; it cannot supply missing implementation.

Keep workflowLanguage go for this mixed task. Independent source review must explicitly cover lua/rex_whichkey.lua and its loading/argv/origin/lifetime behavior alongside all owned Go files and the Lua config/dispatch boundary. No Lua work may be omitted because the primary route is Go; no Go TDD/runtime gates are authorized.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDExOjI1OjM0LjIwOVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3doaWNoa2V5LWZpcnN0L3ByZXNlcnZlL3Rhc2tzX19yZXgtd2hpY2hrZXktMDQtMDUtcG9wdXAtaW50ZWdyYXRpb24ubWQiLCJpbnRlbnQiOiJmNjg0NGFiODc0Mzk1NmJhOWM2MmYwM2ZiMDExZTRjY2M5OWY3NzU1YjViMTY3YTVmMjc0N2ZmN2Q3OTRjNWQ4In0 -->
### 2026-10-31T11:25:34.209Z · historical · unspecified

{"schemaVersion":1,"eventId":"rex-whichkey/whichkey-first/preserve/tasks__rex-whichkey-04-05-popup-integration.md","kind":"historical","recordRef":"tasks/rex-whichkey-04-05-popup-integration.md","specHash":"sha256:cc6c477afc7dc584a79a01d8c3f8c69b6d40a2acfffdbac7aa9df57fbe68b2ec","attempt":1,"content":"# Draft task\n\nDRAFT ONLY: user approved creation of this decomposition, not final interfaces or implementation dispatch. Governing decision REXWHICHKEY-20261030-094737 is Accepted for draft creation. PRD evidence repair and fingerprint-bound design approval remain gates. Paths and contracts below are provisional planning reservations, not accepted public APIs; reassess them after the RFC and republish fingerprints before dispatch. Follow the accepted flat concrete Go architecture, not generic hexagonal defaults. Go workflow language denotes this Go project; Lua/Nix work also requires its specialist checks. Workers never edit .tasks; coordinator alone publishes design/approval evidence.\n\n## Deliverable and scope\nWire the short launcher and terminal-owned Go process to approved launch/config/selection contracts. Launcher returns promptly; Go owns menu navigation and exits on dismissal or dispatch.\n\n## Requirement criteria\nREQ-001, REQ-005, REQ-007, REQ-008, REQ-011, REQ-013\nPopup survives beyond the configuration action lifetime, carries exact origin identities, executes one validated selection and cleans only owned resources on cancellation/errors. All five graphical positions and focus restoration pass; moved origins resolve safely. No background daemon or user init.lua takeover.\n\n## Ownership\nAllowed source/test edits: cmd/main.go, cmd/main_test.go, config.go, config_test.go, lua/rex_whichkey.lua, tests/integration/popup. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.\n\n## Inputs, outputs and prerequisites\nConsume the approved RFC and relevant frozen producer contracts. Output the behavior and tests above. Explicit task prerequisites: tasks/rex-whichkey-04-04-terminal-lifecycle.md, tasks/rex-whichkey-03-04-invocation-dispatch.md, tasks/rex-whichkey-02-03-block-layers.md. Check native parent-phase dependencies too; membership alone does not confer readiness.\n\n## TDD and verification\nRED/GREEN launcher/CLI boundary tests, actual popup-to-Lua callback and app action, 31-second lifetime check, graphical geometry/resize and error cleanup; root quality gates. Own phase-final UI integration evidence. Each code change includes its own failing-first tests, GREEN evidence and refactor checks; do not split tests into a separately completed task. Root Go gates are task build, task test (unit and race), task lint. Lua requires its isolated real-Rex/config fixtures; Nix requires package/module evaluation.\n\n## Integration and completion\nRequire independent Go-project combined specification/quality review plus relevant Lua/Nix expertise. Coordinator integrates reviewed results serially, runs affected checks on the combined target and publishes evidence. Final-phase gates belong to the parent milestone; unavailable GUI/remote infrastructure remains unmet. No implementation, dispatch, completion or interface adoption is authorized by this draft.","provenance":"Exact description before user-directed implementation-first replanning REXWHICHKEY-20261031-072107; previous evidence remains untouched.","gateEligible":false}
<!-- mtn:end-entry:v1 -->
