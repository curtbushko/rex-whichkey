---
workflowKind: task
workflowLanguage: docs
workflowOrder: 5
workflowRequirements:
  - REQ-003
  - REQ-005
  - REQ-011
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - docs/remote-compatibility.md
  sharedPaths: []
  resources: []
  interfaces: []
title: Document remote-context implementation and limitations
mtnCreateRetry:
  key: rex-whichkey/draft-task/05-05-remote-context/1
  intent: 8f0038b537789d88bcbdadbce533286b75c4345d06c3535ecd212a9bd9a483a3
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:43.485Z'
dateModified: '2026-10-31T19:58:40.268Z'
tags:
  - task
type: task
id: 22ee694a-5c4a-45c7-afbe-96084a9641ef
projects:
  - '[[phases/rex-whichkey-05-parity-remote.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-04-05-popup-integration.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-02-04-host-inventory.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-02-01-cli-execution.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-02-02-resources.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-02-03-block-layers.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-03-01-lua-config.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-03-04-invocation-dispatch.md]]'
    reltype: FINISHTOSTART
---
# Product documentation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Document the actually delivered remote-safe executable/module resolution and server/session/block/client identity propagation in docs/remote-compatibility.md. This task owns documentation only, not runtime implementation or an acceptance harness. Identify residual source-review risks and unexecuted remote behavior honestly.

## Functional criteria

Retain REQ-003, REQ-005, REQ-011, REQ-013 functional intent. Record implementation limitations honestly; runtime verification is not scheduled.

## Ownership

Allowed documentation edits: docs/remote-compatibility.md. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read delivered source and documentation as needed; document the integration the menu needs without taking runtime ownership. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered documentation within this task's ownership counts as its delivery; it does not establish delivery or runtime verification of the implementation it describes. Cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Worker routing

Intended workflowLanguage: docs. Use the supported matching language/documentation worker and independent source-review route once the coordinator confirms the narrow tooling capability update. Do not dispatch through Go TDD gates or substitute Go metadata while support is unavailable. The project no-tests/no-experiments policy overrides worker defaults.

## Runtime owners and documentation inputs

Trace delivered code from 02-01 rex/rex.go and rex/exec.go (CLI executable and targeting), 02-02 rex/session.go and rex/window.go (stable identities), 02-03 rex/block.go and rex/layer.go (current origin/owned resources), 03-01 lua/rex_whichkey/config.lua (server module lookup), 03-04 lua/rex_whichkey/dispatch.lua (origin-bound Rex-runtime callback/action dispatch), and 04-05 config.go, cmd/main.go and lua/rex_whichkey.lua (popup executable/config resolution and identity propagation). These existing runtime owners implement REQ-003/005/011/013; raise missing implementation to them rather than editing outside docs ownership. Keep the existing 02-04 inventory input to explain its distinct targeting behavior.

Reference mapping.go and docs/mapping-format.md when describing shared data. Do not duplicate a schema, propose remote attach/bootstrap/WOL, reserve a remote server, or claim compatibility demonstrated by a trial. Completion means independently reviewed accurate documentation of delivered source and its limitations, not a remote feature pass.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDExOjI1OjQ2LjYyNVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3doaWNoa2V5LWZpcnN0L3ByZXNlcnZlL3Rhc2tzX19yZXgtd2hpY2hrZXktMDUtMDUtcmVtb3RlLWNvbnRleHQubWQiLCJpbnRlbnQiOiI0OTlkNmVjMTgxNjU0YjJjMzJjMTBkNTBjNzllZTA1MzFiZjNhYTM1YjNkMDA0ZTEzMGEyODk3NDI5MDJmY2QzIn0 -->
### 2026-10-31T11:25:46.625Z · historical · unspecified

{"schemaVersion":1,"eventId":"rex-whichkey/whichkey-first/preserve/tasks__rex-whichkey-05-05-remote-context.md","kind":"historical","recordRef":"tasks/rex-whichkey-05-05-remote-context.md","specHash":"sha256:b66ff7ea8f5cc13d5651fab2541fed3fa1f4b292da57cd27cb34ceb2a19e0af8","attempt":1,"content":"# Draft task\n\nDRAFT ONLY: user approved creation of this decomposition, not final interfaces or implementation dispatch. Governing decision REXWHICHKEY-20261030-094737 is Accepted for draft creation. PRD evidence repair and fingerprint-bound design approval remain gates. Paths and contracts below are provisional planning reservations, not accepted public APIs; reassess them after the RFC and republish fingerprints before dispatch. Follow the accepted flat concrete Go architecture, not generic hexagonal defaults. Go workflow language denotes this Go project; Lua/Nix work also requires its specialist checks. Workers never edit .tasks; coordinator alone publishes design/approval evidence.\n\n## Deliverable and scope\nAdd a reproducible real remote-server acceptance harness for installed binary/module resolution, inline/imported configuration and popup-to-origin callback/action targeting. Own remote-specific fixes only in this scope; changes to other components require coordinator reassignment.\n\n## Requirement criteria\nREQ-003, REQ-005, REQ-011, REQ-013\nNo client-local absolute paths are needed; menu runtime/inventory ownership and original server/session/block/client IDs remain correct. Verify moved-block, dismissal and error cleanup remotely. Unavailable remote infrastructure is explicitly unmet, not a mock-only pass. Do not add attach/bootstrap workflows.\n\n## Ownership\nAllowed source/test edits: tests/integration/remote, docs/remote-compatibility.md. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.\n\n## Inputs, outputs and prerequisites\nConsume the approved RFC and relevant frozen producer contracts. Output the behavior and tests above. Explicit task prerequisites: tasks/rex-whichkey-04-05-popup-integration.md, tasks/rex-whichkey-02-04-host-inventory.md. Check native parent-phase dependencies too; membership alone does not confer readiness.\n\n## TDD and verification\nFailing-first remote-context fixtures plus real remote popup/callback/app-routing checks with version/ID evidence and owned-resource cleanup; root quality gates. Own phase-final parity/remote acceptance matrix. Each code change includes its own failing-first tests, GREEN evidence and refactor checks; do not split tests into a separately completed task. Root Go gates are task build, task test (unit and race), task lint. Lua requires its isolated real-Rex/config fixtures; Nix requires package/module evaluation.\n\n## Integration and completion\nRequire independent Go-project combined specification/quality review plus relevant Lua/Nix expertise. Coordinator integrates reviewed results serially, runs affected checks on the combined target and publishes evidence. Final-phase gates belong to the parent milestone; unavailable GUI/remote infrastructure remains unmet. No implementation, dispatch, completion or interface adoption is authorized by this draft.","provenance":"Exact description before user-directed implementation-first replanning REXWHICHKEY-20261031-072107; previous evidence remains untouched.","gateEligible":false}
<!-- mtn:end-entry:v1 -->
