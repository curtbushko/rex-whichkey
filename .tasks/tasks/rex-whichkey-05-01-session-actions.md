---
workflowKind: task
workflowLanguage: lua
workflowOrder: 1
workflowRequirements:
  - REQ-010
  - REQ-011
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - lua/rex_whichkey/defaults/sessions.lua
    - tests/integration/sessions
  sharedPaths: []
  resources: []
  interfaces:
    - default-session-mappings
title: Implement session actions and delegated picker parity
mtnCreateRetry:
  key: rex-whichkey/draft-task/05-01-session-actions/1
  intent: 95e42491c8cd1226409e06aa37f1fe7d32281d17c6429a44e475f203e91447a6
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:35.431Z'
dateModified: '2026-11-01T07:12:11.581Z'
tags:
  - task
type: task
id: 8e07941c-9c3c-4f1f-897a-5734639edb3e
projects:
  - '[[phases/rex-whichkey-05-parity-remote.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-04-05-popup-integration.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-03-05-mapping-format.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-01'
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Add session new/switch/previous/next/rename/close mappings with correct app routing and confirmation/context behavior. Preserve the observed delegated Rex session picker rather than adding an unrequested enumerated submenu.


## Functional criteria

REQ-010, REQ-011
Each action matches approved current API/action arguments and old functional behavior; picker opens in the originating client and cancellation is safe. Destructive close requires confirmation. No live resource mutation is scheduled.


## Ownership

Allowed source/test edits: lua/rex_whichkey/defaults/sessions.lua, tests/integration/sessions. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Authoritative format input

Consume mapping.go and docs/mapping-format.md from tasks/rex-whichkey-03-05-mapping-format.md after its native producer prerequisite is delivered. These Go-owned files define the shared mapping/configuration standard; do not define or copy an independent Lua/Nix/consumer schema. Executable configuration and callbacks remain in the originating Rex Lua runtime. Follow the existing functional requirements; raise a consequential new public interface or dependency choice for user approval instead of silently selecting it.

## Worker routing

Intended workflowLanguage: lua. Use the supported matching language/documentation worker and independent source-review route once the coordinator confirms the narrow tooling capability update. Do not dispatch through Go TDD gates or substitute Go metadata while support is unavailable. The project no-tests/no-experiments policy overrides worker defaults.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDA3OjEyOjExLjU3MFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNS10MDEtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiMGZjMWQxNmIzNWNhMWZjMzZkMTA3MjY5YmEwYmFlOGNlNWM4ZDQzYWUxZjExOTVjZjRkNWRlMTBjNmVjMDA2MCJ9 -->
### 2026-11-01T07:12:11.570Z · summary · unspecified

{"acceptedScope":"Implement session actions and delegated picker parity","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p05-t01-tasks-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p05-t01-builder-a01/integration/integration.json"],"attempt":1,"eventId":"rex-whichkey-p05-t01-builder-a01/summary/1","files":["lua/rex_whichkey/defaults/sessions.lua"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-05-01-session-actions.md","requirements":["REQ-010","REQ-011"],"schemaVersion":1,"specHash":"sha256:6b6cbd9e0bdd99a96126c60ffd0ba6704db2553273ad6924e7f37664f3bcc9f1","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"647b29b91a05aaa1ccb6cdf3aa520e1c90f961b0","targetTree":"e3d2848eb5145df3331ce5167204ec0955432226","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
