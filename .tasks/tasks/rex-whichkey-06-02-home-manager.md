---
workflowKind: task
workflowLanguage: nix
workflowOrder: 2
workflowRequirements:
  - REQ-014
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - nix/home-manager.nix
    - tests/packaging/home-manager
  sharedPaths:
    - flake.nix
  resources: []
  interfaces:
    - home-manager-options
title: Add opt-in Home Manager integration preserving user config
mtnCreateRetry:
  key: rex-whichkey/draft-task/06-02-home-manager/1
  intent: 2746c0f3d33d8f22c8690dbf0618c888dd73327961fc4f76c8b5eab367d53457
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:47.500Z'
dateModified: '2026-11-01T08:07:45.896Z'
tags:
  - task
type: task
id: ff24179c-3cd4-4976-a2da-b7644ed37c10
projects:
  - '[[phases/rex-whichkey-06-packaging-release.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-06-01-nix-package.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-01'
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Provide an opt-in Home Manager module that makes packaged integration available without taking ownership of the user's existing Rex init.lua.


## Functional criteria

REQ-014
Existing init.lua and separate user mapping files remain byte-for-byte unchanged. Source must support disabled/enabled configurations and explicit require/setup wiring on the server host; independent Nix source review checks both branches without claiming evaluation or runtime success. No activation reload/restart or unmanaged file overwrite.


## Ownership

Allowed source/test edits: nix/home-manager.nix, tests/packaging/home-manager, flake.nix. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Worker routing

Intended workflowLanguage: nix. Use the supported matching language/documentation worker and independent source-review route once the coordinator confirms the narrow tooling capability update. Do not dispatch through Go TDD gates or substitute Go metadata while support is unavailable. The project no-tests/no-experiments policy overrides worker defaults.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDA4OjA3OjQ1Ljg4NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNi10MDItYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiZTcyNzdiNWMwNDQ0ZDlkYjJiN2EzMmUzOTE3YTY5MTY2OGVkNWM1NjY2MzgyNWYyYzg1MTk3OGQ5MTY4ZDA0OSJ9 -->
### 2026-11-01T08:07:45.885Z · summary · unspecified

{"acceptedScope":"Add opt-in Home Manager integration preserving user config","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p06-t02-tasks-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p06-t02-builder-a01/integration/integration.json"],"attempt":1,"eventId":"rex-whichkey-p06-t02-builder-a01/summary/1","files":["flake.nix","nix/home-manager.nix","tests/packaging/home-manager/README.md"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-06-02-home-manager.md","requirements":["REQ-014"],"schemaVersion":1,"specHash":"sha256:8e1f696fc4d9db81290cdb8aeef392381a82fbc3e6b7aa2fd3ea6fdd5ffa4ff7","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"4980e1f95a96731a1baee445562864c7e6e771bd","targetTree":"68b72d7d8bdb2b4ccd99a7312c4bb9d3e772b2bb","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
