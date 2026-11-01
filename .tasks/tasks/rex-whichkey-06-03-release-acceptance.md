---
workflowKind: task
workflowLanguage: docs
workflowOrder: 3
workflowRequirements:
  - REQ-011
  - REQ-015
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - docs/acceptance.md
    - README.md
  sharedPaths: []
  resources: []
  interfaces: []
title: Document release delivery and limitations
mtnCreateRetry:
  key: rex-whichkey/draft-task/06-03-release-acceptance/1
  intent: a034182ce339b4443b409d7768daa21f0929c288af47136f3bffb690decc73de
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:49.498Z'
dateModified: '2026-11-01T08:16:08.272Z'
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
completedDate: '2026-11-01'
---
# Product documentation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Prepare usage/troubleshooting documentation and an integrated delivery inventory mapping REQ-001–REQ-015 to actual code, outstanding limitations and available historical evidence. Do not implement or run an acceptance runner, GUI matrix, remote trial or disconnected-origin test.

## Functional criteria

Retain REQ-011, REQ-015 functional intent. Record implementation limitations honestly; runtime verification is not scheduled.

## Ownership

Allowed documentation edits: docs/acceptance.md, README.md. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read delivered source and documentation as needed; document the integration the menu needs without taking runtime ownership. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered documentation within this task's ownership counts as its delivery; it does not establish delivery or runtime verification of the implementation it describes. Cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Worker routing

Intended workflowLanguage: docs. Use the supported matching language/documentation worker and independent source-review route once the coordinator confirms the narrow tooling capability update. Do not dispatch through Go TDD gates or substitute Go metadata while support is unavailable. The project no-tests/no-experiments policy overrides worker defaults.

## Honest release inventory

Document actual delivered Go/Lua/Nix files and independent source-review outcomes, distinguish historical evidence from current execution, and state that project tests, package builds, graphical checks, remote trials and disconnected-origin behavior were not executed under current authority. No graphical-client or remote-test-server resource is reserved. Completion records reviewed documentation and a source-based delivery inventory, not an acceptance-test pass or inferred functional delivery.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDA4OjE2OjA4LjI2MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNi10MDMtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiNDVmNDU5NDYwYWY4NzliNWZmMjEwNDNiOTViNjUxMzZkZjg3MjUzMzNiZTg2YjZlYzFkZGFlYzFhOGZhNGM2YiJ9 -->
### 2026-11-01T08:16:08.261Z · summary · unspecified

{"acceptedScope":"Document release delivery and limitations","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p06-t03-tasks-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p06-t03-builder-a01/integration/integration.json"],"attempt":1,"eventId":"rex-whichkey-p06-t03-builder-a01/summary/1","files":["README.md","docs/acceptance.md"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-06-03-release-acceptance.md","requirements":["REQ-011","REQ-015"],"schemaVersion":1,"specHash":"sha256:5d5e7f07b261f1e45f7ec1626993c7e667c5bd921dacc9980259b7c0b9a36c1e","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"40c4f53a8b12f71245057cecaa2c18286d7d9c25","targetTree":"5b8ac6d2d75121426c26188c53d73e5ee261d313","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
