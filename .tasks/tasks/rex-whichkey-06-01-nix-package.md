---
workflowKind: task
workflowLanguage: nix
workflowOrder: 1
workflowRequirements:
  - REQ-014
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - nix/package.nix
    - tests/packaging/package
  sharedPaths:
    - flake.nix
  resources: []
  interfaces:
    - nix-package-layout
title: Package the Go executable and Lua integration with Nix
mtnCreateRetry:
  key: rex-whichkey/draft-task/06-01-nix-package/1
  intent: 402a3b8e428c8918108b7f66f0cc70987fb4191654b8e99fd993cbd45619503a
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:45.490Z'
dateModified: '2026-11-01T07:59:26.794Z'
tags:
  - task
type: task
id: 84fb689e-9da7-4326-9c2c-65fd84a205d8
projects:
  - '[[phases/rex-whichkey-06-packaging-release.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-04-05-popup-integration.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-05-01-session-actions.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-05-02-tab-actions.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-05-03-pane-actions.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-05-04-host-menu.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-01'
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Expose a Nix package for the executable and Lua integration using the server-side installation/module paths. Retain pinned inputs; changes to flake.lock require separate authorization.


## Functional criteria

REQ-014
The package definition must install the executable and loadable Lua integration with complete integrated defaults. Package installation must not write init.lua or assume caller-local module paths. Use Nix source review, not Go project verification, package builds or Nix evaluation; build/load success remains unexecuted.


## Ownership

Allowed source/test edits: nix/package.nix, tests/packaging/package, flake.nix. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Worker routing

Intended workflowLanguage: nix. Use the supported matching language/documentation worker and independent source-review route once the coordinator confirms the narrow tooling capability update. Do not dispatch through Go TDD gates or substitute Go metadata while support is unavailable. The project no-tests/no-experiments policy overrides worker defaults.

## Actual producer inputs and final source assembly

Retain the popup prerequisite. Also require delivered 05-01 lua/rex_whichkey/defaults/sessions.lua, 05-02 lua/rex_whichkey/defaults/tabs.lua, 05-03 lua/rex_whichkey/defaults/panes.lua and lua/rex_whichkey/defaults/app.lua, and 05-04 hosts.go before completion. Review the final package source/install inventory against those files and the popup's Go/Lua entry points. Include the integrated Go-owned format from mapping.go and docs/mapping-format.md through these producer chains; do not redefine it in Nix. Do not claim an earlier popup-only package includes defaults produced later. Source review must identify missing inputs or stale install lists; no build result or runtime load pass is inferred.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDA3OjU5OjI2Ljc4M1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNi10MDEtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiNWI0NTA4NWY2NTVhYTU4MTIxN2Q3MWUwMmI5ODBjOGU2ODVhNmU0Y2MwZDkyZGYxNTE1MThjZjM4YmZjODdkZiJ9 -->
### 2026-11-01T07:59:26.783Z · summary · unspecified

{"acceptedScope":"Package the Go executable and Lua integration with Nix","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p06-t01-tasks-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p06-t01-builder-a01/integration/integration.json"],"attempt":1,"eventId":"rex-whichkey-p06-t01-builder-a01/summary/1","files":["flake.nix","nix/package.nix"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-06-01-nix-package.md","requirements":["REQ-014"],"schemaVersion":1,"specHash":"sha256:9adb80dc94df75def5a264b8032dd5a09ce2d1b98758aa751fb81fbf7a360e90","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"908a3395cf9c701d9aa1476e62119d9d1daebd21","targetTree":"ff0cf85ba78f3274e79e3096bcce50e2534bf94b","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"},{"command":"nix build --no-link --no-write-lock-file --cores 16 --max-jobs 1 .#default (project checks disabled)","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
