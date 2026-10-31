---
workflowKind: task
workflowLanguage: go
workflowOrder: 3
workflowRequirements:
  - REQ-002
  - REQ-011
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - rex/block.go
    - rex/block_test.go
    - rex/layer.go
    - rex/layer_test.go
    - rex/testdata/layers
  sharedPaths: []
  resources: []
  interfaces:
    - rex-block-layer
title: Implement moved-block focus and owned layer lifecycle
mtnCreateRetry:
  key: rex-whichkey/draft-task/02-03-block-layers/1
  intent: a023ce25bab658bbdd983ebffe687ad76f7f95db08e95ed852e1b8493e14fee9
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:13.382Z'
dateModified: '2026-10-31T20:51:18.797Z'
tags:
  - task
type: task
id: 3915799c-1e1d-403d-9be9-5462f646138b
projects:
  - '[[phases/rex-whichkey-02-rex-client.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-02-02-resources.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-10-31'
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Implement block lookup/focus and layer create/resize/close methods required by the popup contract. Resolve current block location rather than trusting cached window membership.


## Functional criteria

REQ-002, REQ-011, REQ-013
Moved blocks focus in their current window while session identity remains correct. Missing origins, invalid bounds, process exits and already-vanished wait targets fail safely without blind layer recreation. Cleanup touches only returned owned IDs.


## Ownership

Allowed source/test edits: rex/block.go, rex/block_test.go, rex/layer.go, rex/layer_test.go, rex/testdata/layers. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Remote implementation responsibility

rex/block.go and rex/layer.go resolve origin/block location in the selected server/session and propagate exact owned layer/block identities for popup lifecycle and cleanup. Do not fall back to a different client/server when the origin is unavailable. 05-05 only documents this implementation.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDIwOjUxOjE4Ljc4NloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMi10MDMtYnVpbGRlci1hMDIvc3VtbWFyeS8yIiwiaW50ZW50IjoiYmFjMTRmMWFjNTE2ODg5OTg0YzhmOTE2OTgxNWUwOWRmZGNlNDkwZTYzMTg0YzM4NjZjMzU4NTdiY2U2OWE3NiJ9 -->
### 2026-10-31T20:51:18.786Z · summary · unspecified

{"acceptedScope":"Implement moved-block focus and owned layer lifecycle","approvalSummary":["User-approved corrected Phase2 plan and explicit implementation continuation.","Fresh independent combined review rex-whichkey-p02-t03-reviewer-a02/review/2 APPROVED current attempt after cleanup race P02-T03-R1 was fixed."],"artifacts":[".mtn-tmp/phase2-delivery/rex-whichkey-p02-t03-reviewer-a02/review.json"],"attempt":2,"eventId":"rex-whichkey-p02-t03-builder-a02/summary/2","files":["rex/block.go","rex/layer.go"],"kind":"summary","limitations":["No project tests, TDD, experiments or live Rex/remote trials for this implementation.","Source review and standard compilation are not runtime behavior verification; cross-family review unavailable.","Cleanup refuses unsafe remaining layer removal rather than issue a cascading close that could destroy unrelated blocks."],"recordRef":"tasks/rex-whichkey-02-03-block-layers.md","requirements":["REQ-002","REQ-011","REQ-013"],"schemaVersion":1,"specHash":"sha256:74177525d85656ad88ef2f18025e0b803585fb5589ccf3cb75425851fe13ada2","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"fb7ed1228620a335d0e325e88a66ee3c024dfef6","targetTree":"c9985ba3a14e69000dd9b9a4b341b7bb07c67080","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./rex (combined integration candidate)","result":"pass"},{"command":"git rev-parse candidate^{tree} equals current approved build tree; git diff --check target-base candidate; git merge-base --is-ancestor integrated-commit main","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
