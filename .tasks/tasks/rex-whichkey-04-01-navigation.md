---
workflowKind: task
workflowLanguage: go
workflowOrder: 1
workflowRequirements:
  - REQ-006
  - REQ-007
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - menu.go
    - menu_test.go
    - action.go
    - action_test.go
  sharedPaths: []
  resources: []
  interfaces:
    - menu-state
title: Implement menu navigation and confirmation state
mtnCreateRetry:
  key: rex-whichkey/draft-task/04-01-navigation/1
  intent: fe044398b4e8373e872cfcdd0a5c6bb7b980284edefdc3bbc5a04c6e7e1ab717
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:25.375Z'
dateModified: '2026-10-31T19:58:30.285Z'
tags:
  - task
type: task
id: 13823b90-5ac4-4572-9641-2a2a87814b54
projects:
  - '[[phases/rex-whichkey-04-menu-ui.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-03-05-mapping-format.md]]'
    reltype: FINISHTOSTART
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Implement renderer-independent menu state over mapping data, including arbitrary-depth prefixes, breadcrumbs, Esc/Backspace, conditional q, pagination events and confirmation.


## Functional criteria

REQ-003, REQ-004, REQ-005, REQ-006, REQ-007, REQ-013
Hidden active keys resolve; q dispatches when mapped and dismisses only as specified. Esc/Backspace return or dismiss correctly. Confirmed entries dispatch only on y; cancellation never produces a selection. State handling covers both named actions and callbacks.


## Ownership

Allowed source/test edits: menu.go, menu_test.go, action.go, action_test.go. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.


## Authoritative format input
Consume mapping.go and docs/mapping-format.md from tasks/rex-whichkey-03-05-mapping-format.md after its native prerequisite is delivered. Go owns the standard; this consumer must not define a separate schema. Consequential unapproved public contracts or dependency selection remain user decisions.
<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDExOjI1OjI0LjI0NloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3doaWNoa2V5LWZpcnN0L3ByZXNlcnZlL3Rhc2tzX19yZXgtd2hpY2hrZXktMDQtMDEtbmF2aWdhdGlvbi5tZCIsImludGVudCI6ImY4OGY4OGZkMTRjZWFjM2U5NWRhOThhOGM4NjM5MGM2NzNjMGFlYjg2ODQ4NWE0MzQzNjFiNWI0YTYwN2IxMGYifQ -->
### 2026-10-31T11:25:24.246Z · historical · unspecified

{"schemaVersion":1,"eventId":"rex-whichkey/whichkey-first/preserve/tasks__rex-whichkey-04-01-navigation.md","kind":"historical","recordRef":"tasks/rex-whichkey-04-01-navigation.md","specHash":"sha256:d313e4fa35a8c270a7baf460a665336e93333e09ea5fb437acc57009974d576c","attempt":1,"content":"# Draft task\n\nDRAFT ONLY: user approved creation of this decomposition, not final interfaces or implementation dispatch. Governing decision REXWHICHKEY-20261030-094737 is Accepted for draft creation. PRD evidence repair and fingerprint-bound design approval remain gates. Paths and contracts below are provisional planning reservations, not accepted public APIs; reassess them after the RFC and republish fingerprints before dispatch. Follow the accepted flat concrete Go architecture, not generic hexagonal defaults. Go workflow language denotes this Go project; Lua/Nix work also requires its specialist checks. Workers never edit .tasks; coordinator alone publishes design/approval evidence.\n\n## Deliverable and scope\nImplement renderer-independent menu state over approved mapping data, including arbitrary-depth prefixes, breadcrumbs, Esc/Backspace, conditional q, pagination events and confirmation.\n\n## Requirement criteria\nREQ-006, REQ-007\nHidden active keys resolve; q dispatches when mapped and dismisses only as specified. Esc/Backspace return or dismiss correctly. Confirmed entries dispatch only on y; cancellation never produces a selection. State tests cover both named actions and callbacks.\n\n## Ownership\nAllowed source/test edits: menu.go, menu_test.go, action.go, action_test.go. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.\n\n## Inputs, outputs and prerequisites\nConsume the approved RFC and relevant frozen producer contracts. Output the behavior and tests above. Explicit task prerequisites: none beyond parent phase barriers. Check native parent-phase dependencies too; membership alone does not confer readiness.\n\n## TDD and verification\nRED/GREEN table-driven transitions, deep prefixes, empty menus, key conflicts and confirmation cancellation; root quality gates. Each code change includes its own failing-first tests, GREEN evidence and refactor checks; do not split tests into a separately completed task. Root Go gates are task build, task test (unit and race), task lint. Lua requires its isolated real-Rex/config fixtures; Nix requires package/module evaluation.\n\n## Integration and completion\nRequire independent Go-project combined specification/quality review plus relevant Lua/Nix expertise. Coordinator integrates reviewed results serially, runs affected checks on the combined target and publishes evidence. Final-phase gates belong to the parent milestone; unavailable GUI/remote infrastructure remains unmet. No implementation, dispatch, completion or interface adoption is authorized by this draft.","provenance":"Exact description before user-directed implementation-first replanning REXWHICHKEY-20261031-072107; previous evidence remains untouched.","gateEligible":false}
<!-- mtn:end-entry:v1 -->
