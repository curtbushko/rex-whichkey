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
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:25.375Z'
dateModified: '2026-11-01T01:06:49.892Z'
tags:
  - task
type: task
id: 13823b90-5ac4-4572-9641-2a2a87814b54
projects:
  - '[[phases/rex-whichkey-04-menu-ui.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-03-05-mapping-format.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-10-31'
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

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDAxOjA2OjQ5Ljg4MFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNC10MDEtYnVpbGRlci1hMDIvc3VtbWFyeS8yIiwiaW50ZW50IjoiMjRhNzFkZGQwNjUzZjQwYjUwZWRkNThmYzgzNWY4YzZmZWUxZDM3YzNlZWVkMmNlNTY5MmU5YTgwYTY1ZTYxZiJ9 -->
### 2026-11-01T01:06:49.880Z · summary · unspecified

{"acceptedScope":"Implement menu navigation and confirmation state","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p04-t01-reviewer-a02/review/2 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p04-t01-builder-a02/integration/integration.json"],"attempt":2,"eventId":"rex-whichkey-p04-t01-builder-a02/summary/2","files":["action.go","menu.go"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-04-01-navigation.md","requirements":["REQ-006","REQ-007"],"schemaVersion":1,"specHash":"sha256:4983d2b17c9aaf2ffa5983c2d10c2e4a4f9856733b931e83b2d8bc58099e93b8","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"2d612d615255c3870bbfce20bcd6b92bfcd353ee","targetTree":"ffe4e3aefa5eeaac56b802c9f7edba05bd9c91f3","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
