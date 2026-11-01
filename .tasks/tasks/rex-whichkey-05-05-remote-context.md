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
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:43.485Z'
dateModified: '2026-11-01T07:51:14.461Z'
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
completedDate: '2026-11-01'
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

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDA3OjUxOjE0LjQ0OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNS10MDUtYnVpbGRlci1hMDIvc3VtbWFyeS8yIiwiaW50ZW50IjoiMjhkZmVmOTIwYmY5NTljYTA2YmI0MzFiMmM2OWFiOWNkMzY5YjM4OTE4MDE0YmU0YmY2ZTE0ZWE3MDM2NmFkZCJ9 -->
### 2026-11-01T07:51:14.449Z · summary · unspecified

{"acceptedScope":"Document remote-context implementation and limitations","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p05-t05-tasks-a02/review/2 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p05-t05-builder-a02/integration/integration.json"],"attempt":2,"eventId":"rex-whichkey-p05-t05-builder-a02/summary/2","files":["docs/remote-compatibility.md"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-05-05-remote-context.md","requirements":["REQ-003","REQ-005","REQ-011","REQ-013"],"schemaVersion":1,"specHash":"sha256:290e7d31b1345a3311d9148cb6c745bb8f193d2d32793080efd973a8a23ac167","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"cddf43711feeb2625be8f3ba5e719dde3bafab28","targetTree":"3f93623363683776bf507a2192558ba69aae4f54","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
