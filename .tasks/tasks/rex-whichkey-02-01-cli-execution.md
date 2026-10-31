---
workflowKind: task
workflowLanguage: go
workflowOrder: 1
workflowRequirements:
  - REQ-002
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - rex/rex.go
    - rex/rex_test.go
    - rex/exec.go
    - rex/exec_test.go
    - rex/testdata/cli
  sharedPaths:
    - go.mod
    - go.sum
  resources: []
  interfaces:
    - rex-cli-execution
title: Implement safe Rex CLI execution and generic calls
mtnCreateRetry:
  key: rex-whichkey/draft-task/02-01-cli-execution/1
  intent: bc8b5a8f23e937add2ddc261806611855b9f57f5b3a11943a73ba89c45167673
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:09.390Z'
dateModified: '2026-10-31T20:23:20.810Z'
tags:
  - task
type: task
id: befda34b-e61c-4162-9329-f7c635788f97
projects:
  - '[[phases/rex-whichkey-02-rex-client.md]]'
completedDate: '2026-10-31'
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Implement concrete client targeting, argument-array process execution, typed JSON/error handling and generic Call using documented API/CLI contracts. A small private process-runner seam is permitted; no service or transport hierarchy.


## Functional criteria

REQ-002, REQ-013
Implementation must preserve exact argv for server/session targeting, spaces/metacharacters and JSON values; invalid JSON, nonzero exits, missing executable, cancellation and timeouts return useful errors. Never shell-eval user configuration or automatically retry destructive requests.


## Ownership

Allowed source/test edits: rex/rex.go, rex/rex_test.go, rex/exec.go, rex/exec_test.go, rex/testdata/cli, go.mod, go.sum. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Remote implementation responsibility

Own executable resolution for the rex CLI and exact server/session targeting in rex/rex.go and rex/exec.go, including errors when the executable is unavailable in the execution host context. Preserve caller-supplied originating context; do not substitute client-local absolute paths, silently select another server, or use shell interpolation. Popup executable resolution belongs to 04-05. 05-05 documents this delivered behavior and cannot edit these runtime files.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDIwOjIzOjIwLjc5OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMi10MDEtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiOWU0OGVkNmM4ZTUwMGM0M2E0ZGY1NTliZTE4ODQ2YzBjNTJhMDIwMDBkMzI1NzdlZWE1YzUzNmVjMDczYTJmNCJ9 -->
### 2026-10-31T20:23:20.799Z · summary · unspecified

{"acceptedScope":"Implement safe Rex CLI execution and generic calls","approvalSummary":["User-approved corrected Phase2 plan; implementation explicitly resumed by user.","Independent combined source review rex-whichkey-p02-t01-reviewer-a01/review/1 APPROVED; fresh same-family reviewer context."],"artifacts":[".mtn-tmp/phase2-delivery/rex-whichkey-p02-t01-reviewer-a01/review.json",".mtn-tmp/phase2-delivery/integration-t01-a01/integration.json"],"attempt":1,"eventId":"rex-whichkey-p02-t01-builder-a01/summary/1","files":["rex/exec.go","rex/rex.go"],"kind":"summary","limitations":["No project tests, TDD, experiments or live Rex/remote trials performed for this implementation.","Source review and standard compilation are not runtime behavior verification; cross-family review unavailable."],"recordRef":"tasks/rex-whichkey-02-01-cli-execution.md","requirements":["REQ-002","REQ-013"],"schemaVersion":1,"specHash":"sha256:fdaea00d3856d8587fa74d80f5d05147510537f58f11eed452187cd68f7cbc7d","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"020a31e7271ef75c68ab3f432c256057f5d1d6b1","targetTree":"69ff536db6a4ff5ca6e21c9e5a4b28a48755389b","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./rex (combined integration candidate)","result":"pass"},{"command":"git diff --check target-base candidate; git merge-base --is-ancestor integrated-commit main","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
