---
workflowKind: task
workflowLanguage: go
workflowOrder: 3
workflowRequirements:
  - REQ-011
  - REQ-015
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - tests/acceptance
    - docs/acceptance.md
    - README.md
  sharedPaths: []
  resources:
    - rex-graphical-client
    - rex-remote-test-server
  interfaces: []
title: Document integrated delivery and release limitations
mtnCreateRetry:
  key: rex-whichkey/draft-task/06-03-release-acceptance/1
  intent: a034182ce339b4443b409d7768daa21f0929c288af47136f3bffb690decc73de
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:49.498Z'
dateModified: '2026-10-31T11:26:42.157Z'
tags:
  - task
type: task
id: 73fa56ee-f723-472e-b636-ac5dd1996cb8
projects:
  - '[[phases/rex-whichkey-06-packaging-release.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-06-02-home-manager.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-05-01-session-actions.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-05-02-tab-actions.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-05-03-pane-actions.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-05-04-host-menu.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-05-05-remote-context.md]]'
    reltype: FINISHTOSTART
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. This planning update does not execute implementation or approve provisional IF-* signatures.

## Deliverable and scope

Prepare usage/troubleshooting documentation and an integrated delivery inventory mapping REQ-001–REQ-015 to actual code, outstanding limitations and available historical evidence. Do not implement or run an acceptance runner, GUI matrix, remote trial or disconnected-origin test.

## Functional criteria

Retain REQ-011, REQ-015 functional intent. Record implementation limitations honestly; runtime verification is not scheduled.

## Ownership

Allowed source/test edits: tests/acceptance, docs/acceptance.md, README.md. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve existing files and evidence. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. This request updates plans only; no agents or implementation are dispatched.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDAyOjE4OjQwLjI2MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L2Rpc2Nvbm5lY3QtZGVmZXJyYWwvcHJlc2VydmUvdGFza3MtLXJleC13aGljaGtleS0wNi0wMy1yZWxlYXNlLWFjY2VwdGFuY2UubWQiLCJpbnRlbnQiOiIwMGNlOTc3MTc0NTJiYmYxN2FiYTRlMWM5NWZhYmI0YmY3Y2M2MGFkNDJhY2MxZjVhYTRlZTRhOGMyNGFiY2E2In0 -->
### 2026-10-31T02:18:40.261Z · historical · unspecified

{"attempt":1,"content":"# Draft task\n\nDRAFT ONLY: user approved creation of this decomposition, not final interfaces or implementation dispatch. Governing decision REXWHICHKEY-20261030-094737 is Accepted for draft creation. PRD evidence repair and fingerprint-bound design approval remain gates. Paths and contracts below are provisional planning reservations, not accepted public APIs; reassess them after the RFC and republish fingerprints before dispatch. Follow the accepted flat concrete Go architecture, not generic hexagonal defaults. Go workflow language denotes this Go project; Lua/Nix work also requires its specialist checks. Workers never edit .tasks; coordinator alone publishes design/approval evidence.\n\n## Deliverable and scope\nImplement a reproducible initial-release acceptance runner/matrix and usage/troubleshooting documentation, then run it against the combined integrated target rather than individual task branches.\n\n## Requirement criteria\nREQ-015\nTrace REQ-001–REQ-015 to delivered tasks and actual evidence for mapping/config/callbacks, all navigation and defaults, five-position GUI geometry, Unicode/theme/restoration, origin focus/cleanup, remote operation and packaging coexistence. Missing GUI/remote gates remain unmet. No coverage-percentage substitute or silent waivers. Record exact target commit/tree, versions, residual limitations and rollback instructions.\n\n## Ownership\nAllowed source/test edits: tests/acceptance, docs/acceptance.md, README.md. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.\n\n## Inputs, outputs and prerequisites\nConsume the approved RFC and relevant frozen producer contracts. Output the behavior and tests above. Explicit task prerequisites: tasks/rex-whichkey-06-02-home-manager.md, tasks/rex-whichkey-05-01-session-actions.md, tasks/rex-whichkey-05-02-tab-actions.md, tasks/rex-whichkey-05-03-pane-actions.md, tasks/rex-whichkey-05-04-host-menu.md, tasks/rex-whichkey-05-05-remote-context.md. Check native parent-phase dependencies too; membership alone does not confer readiness.\n\n## TDD and verification\nFailing-first acceptance assertions, task build, task test, task lint, nix build .#, supported nix flake check and actual local GUI/remote Rex scenarios. Coordinator publishes milestone evidence only after all child integrations and accepted requirement coverage pass; preserve experiment artifacts and request approval for cleanup. Each code change includes its own failing-first tests, GREEN evidence and refactor checks; do not split tests into a separately completed task. Root Go gates are task build, task test (unit and race), task lint. Lua requires its isolated real-Rex/config fixtures; Nix requires package/module evaluation.\n\n## Integration and completion\nRequire independent Go-project combined specification/quality review plus relevant Lua/Nix expertise. Coordinator integrates reviewed results serially, runs affected checks on the combined target and publishes evidence. Final-phase gates belong to the parent milestone; unavailable GUI/remote infrastructure remains unmet. No implementation, dispatch, completion or interface adoption is authorized by this draft.","eventId":"rex-whichkey/disconnect-deferral/preserve/tasks--rex-whichkey-06-03-release-acceptance.md","gateEligible":false,"kind":"historical","provenance":"Exact definition preserved before explicit user-directed disconnected-origin test deferral and trust in Rex API state. This is not a pass, blanket test waiver or new product/design approval.","recordRef":"tasks/rex-whichkey-06-03-release-acceptance.md","schemaVersion":1,"specHash":"sha256:ce52c2555b4e1bfff9cf5231129edf577579f6e1c6c15764cfed108bef2e6cb7"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDAyOjE5OjA1LjkxMVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9kaXNjb25uZWN0LWZpbmFsLW9ibGlnYXRpb24vdGFza3MtLXJleC13aGljaGtleS0wNi0wMy1yZWxlYXNlLWFjY2VwdGFuY2UubWQiLCJpbnRlbnQiOiI4OTUwZDgyZmJmOTk3MWRmMGM2MGFiYWE4MTJmODIyMTM2MWFhNzZkYjlhOTc1MzBkNTQ4OGUwYWRkNmQ4MTY2In0 -->
### 2026-10-31T02:19:05.911Z · approval · unspecified

{"approvedBy":"user","attempt":1,"decision":"approved","decisionEvidence":"User: skip the test for now and defer it to the end of the phases. There is so much more work to be done that does not need rex state. Follow-up: also, you can trust the rex api to tell you state.","decisionId":"REXWHICHKEY-20261030-221324","deferredRequirement":{"dueRef":"tasks/rex-whichkey-06-03-release-acceptance.md","dueStage":"after implementation, before final release/Phase6/project completion","finding":"T02-R002","status":"deferred-untested"},"eventId":"rex-whichkey/disconnect-final-obligation/tasks--rex-whichkey-06-03-release-acceptance.md","kind":"approval","recordRef":"tasks/rex-whichkey-06-03-release-acceptance.md","requirements":["REQ-011"],"schemaVersion":1,"scope":"Defer genuine disconnected-origin/no-fallback test to mandatory final release acceptance after implementation phases; not a Phase1 prerequisite and no second app needed now. Trust authoritative API state without redundant screenshots. No behavioral pass, cancellation, blanket waiver, app disconnect or downstream implementation/design approval inferred.","specHash":"sha256:4377ee02aeb579512b9a355baac8288c1b5ef357e436b1c2eba202242add2690"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDExOjI1OjU0LjQxNVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3doaWNoa2V5LWZpcnN0L3ByZXNlcnZlL3Rhc2tzX19yZXgtd2hpY2hrZXktMDYtMDMtcmVsZWFzZS1hY2NlcHRhbmNlLm1kIiwiaW50ZW50IjoiNzVhMThmYTkwYmU3YWQwOGJkODg0NDM4NDBkYzI0ZTJiODFkOWZkMzg0YTZmZDYxNDQ3YmY1YmE2ZTlkOWFiNyJ9 -->
### 2026-10-31T11:25:54.415Z · historical · unspecified

{"schemaVersion":1,"eventId":"rex-whichkey/whichkey-first/preserve/tasks__rex-whichkey-06-03-release-acceptance.md","kind":"historical","recordRef":"tasks/rex-whichkey-06-03-release-acceptance.md","specHash":"sha256:4377ee02aeb579512b9a355baac8288c1b5ef357e436b1c2eba202242add2690","attempt":1,"content":"# Draft task\n\nDRAFT ONLY: user approved creation of this decomposition, not final interfaces or implementation dispatch. Governing decision REXWHICHKEY-20261030-094737 is Accepted for draft creation. PRD evidence repair and fingerprint-bound design approval remain gates. Paths and contracts below are provisional planning reservations, not accepted public APIs; reassess them after the RFC and republish fingerprints before dispatch. Follow the accepted flat concrete Go architecture, not generic hexagonal defaults. Go workflow language denotes this Go project; Lua/Nix work also requires its specialist checks. Workers never edit .tasks; coordinator alone publishes design/approval evidence.\n\n## Deliverable and scope\nImplement a reproducible initial-release acceptance runner/matrix and usage/troubleshooting documentation, then run it against the combined integrated target rather than individual task branches.\n\n## Requirement criteria\nREQ-015\nTrace REQ-001–REQ-015 to delivered tasks and actual evidence for mapping/config/callbacks, all navigation and defaults, five-position GUI geometry, Unicode/theme/restoration, origin focus/cleanup, remote operation and packaging coexistence. Missing GUI/remote gates remain unmet. No coverage-percentage substitute or silent waivers. Record exact target commit/tree, versions, residual limitations and rollback instructions.\n\n## Ownership\nAllowed source/test edits: tests/acceptance, docs/acceptance.md, README.md. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.\n\n## Inputs, outputs and prerequisites\nConsume the approved RFC and relevant frozen producer contracts. Output the behavior and tests above. Explicit task prerequisites: tasks/rex-whichkey-06-02-home-manager.md, tasks/rex-whichkey-05-01-session-actions.md, tasks/rex-whichkey-05-02-tab-actions.md, tasks/rex-whichkey-05-03-pane-actions.md, tasks/rex-whichkey-05-04-host-menu.md, tasks/rex-whichkey-05-05-remote-context.md. Check native parent-phase dependencies too; membership alone does not confer readiness.\n\n## TDD and verification\nFailing-first acceptance assertions, task build, task test, task lint, nix build .#, supported nix flake check and actual local GUI/remote Rex scenarios. Coordinator publishes milestone evidence only after all child integrations and accepted requirement coverage pass; preserve experiment artifacts and request approval for cleanup. Each code change includes its own failing-first tests, GREEN evidence and refactor checks; do not split tests into a separately completed task. Root Go gates are task build, task test (unit and race), task lint. Lua requires its isolated real-Rex/config fixtures; Nix requires package/module evaluation.\n\n## Integration and completion\nRequire independent Go-project combined specification/quality review plus relevant Lua/Nix expertise. Coordinator integrates reviewed results serially, runs affected checks on the combined target and publishes evidence. Final-phase gates belong to the parent milestone; unavailable GUI/remote infrastructure remains unmet. No implementation, dispatch, completion or interface adoption is authorized by this draft.\n\n## Mandatory final disconnected-origin acceptance — REXWHICHKEY-20261030-221324\nThe user explicitly moved the genuine disconnected-origin/no-fallback test from Phase1 to the end of all implementation phases. Before final release-task, Phase6 or project completion, capture an actual originating app identity, exercise an authorized real disconnect, dispatch only to that exact departed/lost identity once, and verify rejection/no fallback or unintended owned-resource effects using authoritative Rex API state. Preserve the narrower stale-CLI-ID result and all earlier failures; neither substitutes for this pending real behavior. Record graceful disconnect separately from transport-loss/resumption coverage; do not invent passes. The safe lifecycle procedure is to be authorized at that time; no second app or current-app Quit is required now. An absent environment remains a final completion blocker, not a reason to block unrelated implementation. Keep this obligation pending in later approved planning and combined-target acceptance matrices. This amendment is accepted scheduling/traceability only; all original draft-design/dispatch approval gates remain.\n","provenance":"Exact description before user-directed implementation-first replanning REXWHICHKEY-20261031-072107; previous evidence remains untouched.","gateEligible":false}
<!-- mtn:end-entry:v1 -->
