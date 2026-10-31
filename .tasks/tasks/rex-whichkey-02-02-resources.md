---
workflowKind: task
workflowLanguage: go
workflowOrder: 2
workflowRequirements:
  - REQ-002
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - rex/session.go
    - rex/session_test.go
    - rex/window.go
    - rex/window_test.go
    - rex/testdata/resources
  sharedPaths: []
  resources: []
  interfaces:
    - rex-session-window
title: Implement concrete session and window resource operations
mtnCreateRetry:
  key: rex-whichkey/draft-task/02-02-resources/1
  intent: 816a3ab804adb4b4ff38adde3c1862ffa9f65691d3915f917b692031d7e23806
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:11.389Z'
dateModified: '2026-10-31T20:35:42.954Z'
tags:
  - task
type: task
id: 0f2c44dd-c914-4d76-941b-9089bf03282e
projects:
  - '[[phases/rex-whichkey-02-rex-client.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-02-01-cli-execution.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-10-31'
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Implement only the Session/Window operations needed by the menu, preserving stable server/session identities and exposing ordinary concrete values.


## Functional criteria

REQ-002, REQ-013
Published schemas govern list/create/inspect and required window operations, empty collections and malformed responses. Returned resources retain targeting; method names must not invent API request fields. Client app actions remain distinct from server resource operations.


## Ownership

Allowed source/test edits: rex/session.go, rex/session_test.go, rex/window.go, rex/window_test.go, rex/testdata/resources. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Remote implementation responsibility

rex/session.go and rex/window.go retain the selected server and stable session identity on returned concrete resources. Never infer targeting from a client-local process or transient window membership. This existing runtime owner implements the behavior; 05-05 only documents it.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDIwOjM1OjQyLjk0MloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMi10MDItYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiZWU4ZTcyYWI3ZjhiMzI5YWU2OWViZTc0NGVjZWJhY2E1OGNmODIxNjUyYmFlYWI0MjM1NmQzYTM5MTRkYjNiZiJ9 -->
### 2026-10-31T20:35:42.942Z · summary · unspecified

{"acceptedScope":"Implement concrete session and window resource operations","approvalSummary":["User-approved corrected Phase2 plan and explicit implementation continuation.","Independent combined source review rex-whichkey-p02-t02-reviewer-a01/review/1 APPROVED; fresh same-family context."],"attempt":1,"eventId":"rex-whichkey-p02-t02-builder-a01/summary/1","files":["rex/session.go","rex/window.go"],"kind":"summary","limitations":["No project tests, TDD, experiments or live Rex/remote trials for this implementation.","Source review and standard compilation are not runtime behavior verification; cross-family review unavailable."],"recordRef":"tasks/rex-whichkey-02-02-resources.md","requirements":["REQ-002","REQ-013"],"schemaVersion":1,"specHash":"sha256:73c8b30a2cfe7ede6403c3f98d2579d447c0f5a1f46e1cd9505d3c532a3199aa","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"3ec92e9a5b8a983679324932cb3100a7ccd5f26a","targetTree":"8272102c650f8170f10bc6d10ca573d56ad85571","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./rex (combined integration candidate)","result":"pass"},{"command":"git diff --exit-code reviewed-head candidate -- owned-files; git diff --check target-base candidate; git merge-base --is-ancestor integrated-commit main","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
