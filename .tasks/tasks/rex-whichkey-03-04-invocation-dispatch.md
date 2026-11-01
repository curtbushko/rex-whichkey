---
workflowKind: task
workflowLanguage: lua
workflowOrder: 4
workflowRequirements:
  - REQ-005
  - REQ-007
  - REQ-011
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - lua/rex_whichkey/dispatch.lua
    - tests/lua/dispatch
  sharedPaths: []
  resources: []
  interfaces:
    - whichkey-invocation-dispatch
title: Implement origin-aware invocation registry and dispatch
mtnCreateRetry:
  key: rex-whichkey/draft-task/03-04-invocation-dispatch/1
  intent: 87ba5fc64c398c84e93f03101eaad0ca1c9f6c7bc7b23af9839a0d8b86a93c64
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:23.380Z'
dateModified: '2026-11-01T02:07:25.341Z'
tags:
  - task
type: task
id: 3127581d-6d5d-4c7b-ac0e-dd3482abb0e2
projects:
  - '[[phases/rex-whichkey-03-lua-mappings.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-03-03-dynamic-mappings.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-03-05-mapping-format.md]]'
    reltype: FINISHTOSTART
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Implement minimal invocation-bound callback/action dispatch and data-only selection validation; retain executable callbacks in the persistent Rex runtime and route named actions through the correct client/server mechanism.


## Functional criteria

REQ-005, REQ-007, REQ-011, REQ-013
Handle overlapping invocations in reverse order, replay, invalid/expired capabilities, wrong origin/client, reload invalidation, missing/moved origins, confirmation, cancellation, stale conditions, callback failures and reentrancy. State is bounded. Focus origin before execution; never automatically retry a timed-out destructive action.


## Ownership

Allowed source/test edits: lua/rex_whichkey/dispatch.lua, tests/lua/dispatch. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Authoritative format input

Consume mapping.go and docs/mapping-format.md from tasks/rex-whichkey-03-05-mapping-format.md after its native producer prerequisite is delivered. These Go-owned files define the shared mapping/configuration standard; do not define or copy an independent Lua/Nix/consumer schema. Executable configuration and callbacks remain in the originating Rex Lua runtime. Follow the existing functional requirements; raise a consequential new public interface or dependency choice for user approval instead of silently selecting it.

## Worker routing

Intended workflowLanguage: lua. Use the supported matching language/documentation worker and independent source-review route once the coordinator confirms the narrow tooling capability update. Do not dispatch through Go TDD gates or substitute Go metadata while support is unavailable. The project no-tests/no-experiments policy overrides worker defaults.

## Remote implementation responsibility

lua/rex_whichkey/dispatch.lua owns binding and checking the originating server/session/block/client and invocation identity when issuing/validating data-only selections and invoking Rex-runtime functions. Preserve the existing bounded lifetime, replay, focus and no-destructive-retry requirements. Use the Go-owned format files, not a Lua-defined transport schema. Actual transport/API signatures remain subject to user approval when consequential; package research does not approve them. 05-05 documents, but cannot implement, this code.


## Reopened integration blocker — REXWHICHKEY-20261031-220348 (Open, not approved)
The accepted popup consumer must return promptly and survive its initiating configuration action. The current dispatch producer instead requires the same live action/rawequal context and a 30-second lifetime; integration cannot satisfy both. This is new source evidence, not a new research gate. The dispatcher task and Phase3 are reopened; other delivered child source and prior c3a0b284/961d8cfd references remain accurate historical deliveries. Current target scope is still original-runtime/captured-origin safe post-return dispatch, not same-action blocking. Exact authentication/re-entry/public completion contract is unresolved: do not merely relax context guards, select a new bridge, or change the accepted lifetime requirement without the pending user decision. Details: .mtn-tmp/remaining-delivery/popup-contract-proposal.md and the blocked popup source-notes.md. No new contract or implementation approval is established by this scope reconciliation.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDAxOjMzOjQ4LjI0OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy10MDQtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiZTAzYzcxZjAzYWJlOWQ5YWI0ZWU2YTM0ZjRjMDg2N2RhZGNhNjNhZjQ5ODZmZTE0MzEzOWU2Y2E5YTJlYTNjNSJ9 -->
### 2026-11-01T01:33:48.249Z · summary · unspecified

{"acceptedScope":"Implement origin-aware invocation registry and dispatch","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p03-t04-tasks-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p03-t04-builder-a01/integration/integration.json"],"attempt":1,"eventId":"rex-whichkey-p03-t04-builder-a01/summary/1","files":["lua/rex_whichkey/dispatch.lua"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-03-04-invocation-dispatch.md","requirements":["REQ-005","REQ-007","REQ-011","REQ-013"],"schemaVersion":1,"specHash":"sha256:55c89a8a5f8acfd7db30fca2fcf69d0454b67507613eb5f10a1c5cd5e6ddce73","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"c3a0b28435d00e675876f7f846492ce4b480f6f9","targetTree":"e752514be4bc43b9a46588be859e866cc4244874","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDAyOjA3OjI1LjMzMFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYmxvY2tlciIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvYmxvY2tlci9wb3N0LXJldHVybi1kaXNwYXRjaC90YXNrcy1yZXgtd2hpY2hrZXktMDMtMDQtaW52b2NhdGlvbi1kaXNwYXRjaC1tZCIsImludGVudCI6IjQ4MTU0MjA5MjNlYmU0ODcyOGNjMGIzMGY1MTQ1ZTkyODg2MTJkYmMwMWVmNzg1MjM5ZmRhNDE1ODJkYjNhY2IifQ -->
### 2026-11-01T02:07:25.330Z · blocker · unspecified

{"affectedRefs":["tasks/rex-whichkey-03-04-invocation-dispatch.md","phases/rex-whichkey-03-lua-mappings.md","tasks/rex-whichkey-04-05-popup-integration.md","phases/rex-whichkey-04-menu-ui.md","phases/rex-whichkey-05-parity-remote.md","phases/rex-whichkey-06-packaging-release.md"],"attempt":1,"decisionRegisterEntry":{"approved":false,"id":"REXWHICHKEY-20261031-220348","implemented":false,"label":"Proposal","proposalFile":".mtn-tmp/remaining-delivery/popup-contract-proposal.md","relatedDecisionIds":["REXWHICHKEY-20261031-072107"],"state":"Open"},"deliveredSource":{"phase4ThroughTerminalCommit":"0d8fc4c8","priorPhase3Commit":"c3a0b284","priorPhase3MetadataCheckpoint":"961d8cfd"},"eventId":"rex-whichkey-p03-coordinator-a01/blocker/post-return-dispatch/tasks-rex-whichkey-03-04-invocation-dispatch-md","kind":"blocker","limitations":["Phase3 and dispatcher reopened on new integration evidence; other four Phase3 child summaries remain delivered.","Phase4 incomplete; no Lua popup launcher or complete popup behavior claimed. Phases5–6 not executed.","No project tests/experiments, live or remote trials for this continuation; standard compilation and independent source reviews are not runtime verification.","No Go-Lua dependency, daemon, public bridge or relaxed origin guard adopted. No push or destructive cleanup."],"preservedUnintegratedSource":{"branch":"build/rex-whichkey-p04-t05-builder-a01","commit":"00fe2599ef12ccbc784ed10e992b450a453a6c4a","delivered":false,"reviewed":false},"reason":"Accepted popup lifetime and delivered dispatcher contract are incompatible: prompt-return detached UI must outlive the initiating configuration action, while dispatch.lua requires the same live action/rawequal context, a runtime-local ticket and a 30-second lease. rex.invoke/client.queue inherit ambient action context. A copied context or relaxed guard is not an authenticated origin-preserving post-return channel. No remaining accepted task is dependency-ready.","recordRef":"tasks/rex-whichkey-03-04-invocation-dispatch.md","requiredDecision":"Approve or reject the minimal authenticated, single-use post-return completion/cancellation bridge proposal. It executes a fresh origin-targeted action in the original server Lua VM, preserving captured identity, replay/expiry/configuration/confirmation/focus protections. This consequential security/interface contract has not been adopted. A synchronous same-action menu would instead require an explicit change to the accepted prompt-return requirement.","schemaVersion":1,"sourceEvidenceFile":".mtn-tmp/remaining-delivery/rex-whichkey-p04-t05-builder-a01/source-notes.md","sourceLocations":["lua/rex_whichkey/dispatch.lua:1-9","lua/rex_whichkey/dispatch.lua:132-140","lua/rex_whichkey/dispatch.lua:189-193","lua/rex_whichkey/dispatch.lua:244-294"],"specHash":"sha256:35ca71748243fe387179fbbfc6c80624edfce3fa2432eb3c12e95ffb6dd54dab"}
<!-- mtn:end-entry:v1 -->
