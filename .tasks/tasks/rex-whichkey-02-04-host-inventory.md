---
workflowKind: task
workflowLanguage: go
workflowOrder: 4
workflowRequirements:
  - REQ-002
  - REQ-012
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - rex/hosts.go
    - rex/hosts_test.go
    - rex/testdata/hosts
  sharedPaths: []
  resources: []
  interfaces:
    - rex-host-inventory
title: Decode effective Rex-known host inventory
mtnCreateRetry:
  key: rex-whichkey/draft-task/02-04-host-inventory/1
  intent: f32ba3182e5c6d7e88394a0c32ec9f399618af8fce884641f8a2c5456977abbe
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:15.389Z'
dateModified: '2026-10-31T20:32:23.973Z'
tags:
  - task
type: task
id: ea04ff70-e7da-4e95-994e-6a10afc78702
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

Wrap rex hosts --json as the effective CLI inventory, not raw hosts.list; follow documented inventory ownership and keep it separate from retargetable server operations.


## Functional criteria

REQ-002, REQ-012, REQ-013
Decode stable IDs/labels, nested host records, warnings, held/cached-home state, empty inventory and malformed output. Do not pretend --server retargets rex hosts; do not add hosts, enumerate Tailscale or probe every endpoint.


## Ownership

Allowed source/test edits: rex/hosts.go, rex/hosts_test.go, rex/testdata/hosts. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDIwOjMyOjIzLjk2MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMi10MDQtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiODQyYzFkYjQ2ZTI1OWNjMDE3MDA2NTQ0NWRkNTJjNDRkMmU4NDkxYWIyOGFmNGUwYTgzNjk3ZWYzNzRkZjZiYiJ9 -->
### 2026-10-31T20:32:23.961Z · summary · unspecified

{"acceptedScope":"Decode effective Rex-known host inventory","approvalSummary":["User-approved corrected Phase2 plan and explicit implementation continuation.","Independent combined source review rex-whichkey-p02-t04-reviewer-a01/review/1 APPROVED; fresh same-family context."],"attempt":1,"eventId":"rex-whichkey-p02-t04-builder-a01/summary/1","files":["rex/hosts.go"],"kind":"summary","limitations":["No project tests, TDD, experiments or live Rex/remote trials for this implementation.","Source review and standard compilation are not runtime behavior verification; cross-family review unavailable."],"recordRef":"tasks/rex-whichkey-02-04-host-inventory.md","requirements":["REQ-002","REQ-012","REQ-013"],"schemaVersion":1,"specHash":"sha256:76ad67c9a797ebe2f570d1173f2e6f96dd2c768fdd953e18f87c9e5bbdcd5c54","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"43352560fb2b248c8ce360b9675a66ff277bac98","targetTree":"6f2ee75ed5a244856220fdf08fc57a025126b956","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./rex (combined integration candidate)","result":"pass"},{"command":"git diff --check target-base candidate; git merge-base --is-ancestor integrated-commit main","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
