---
workflowKind: task
workflowLanguage: go
workflowOrder: 4
workflowRequirements:
  - REQ-012
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - docs/mapping-format.md
    - hosts.go
    - hosts_test.go
    - lua/rex_whichkey/defaults/hosts.lua
    - lua/rex_whichkey/defaults/init.lua
    - lua/rex_whichkey/evaluate.lua
    - lua/rex_whichkey/mapping.lua
    - mapping.go
    - popup.go
    - terminal.go
    - testdata/host-menu
  sharedPaths: []
  resources: []
  interfaces:
    - known-host-menu
    - known-host-provider
    - resolved-default-composition
title: Present Rex-known hosts and inventory warnings
mtnCreateRetry:
  key: rex-whichkey/draft-task/05-04-host-menu/1
  intent: eecf6aef9cb5c5b70d0734b67aecc7029d71e26772ff4447e5d59e6f5e891876
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:41.477Z'
dateModified: '2026-11-01T07:36:28.651Z'
tags:
  - task
type: task
id: e193860b-8de7-43d9-b666-4eaa86359806
projects:
  - '[[phases/rex-whichkey-05-parity-remote.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-02-04-host-inventory.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-04-05-popup-integration.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-03-05-mapping-format.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-05-03-pane-actions.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-01'
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Present the effective Rex inventory in the menu according to documented inventory ownership, preserving labels/IDs and warnings without implementing remote attach/bootstrap.


## Functional criteria

REQ-012, REQ-013
Empty, cached/held-home, warning and duplicate-label values are intelligible and stable. No independent registry, proactive endpoint probes, host mutation, remote attach, WOL or SSH is introduced. Host selection behavior stays within the approved listing-only contract.


## Ownership

Allowed source/test edits: hosts.go, hosts_test.go, testdata/host-menu. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Authoritative format input

Consume mapping.go and docs/mapping-format.md from tasks/rex-whichkey-03-05-mapping-format.md after its native producer prerequisite is delivered. These Go-owned files define the shared mapping/configuration standard; do not define or copy an independent Lua/Nix/consumer schema. Executable configuration and callbacks remain in the originating Rex Lua runtime. Follow the existing functional requirements; raise a consequential new public interface or dependency choice for user approval instead of silently selecting it.


## Accepted delegated production composition — REXWHICHKEY-20261101-022635
The previous pure formatter build is incomplete because it has no real UI caller. Literal ownership now includes the canonical Go field/docs, Lua forwarding adapters, actual terminal/popup caller, and defaults host/aggregation modules. Read .mtn-tmp/remaining-delivery/host-composition-guidance.md for the narrow chosen contract. Implement a single Go-owned provider=known-hosts marker and bounded real rex.Hosts acquisition/expansion before NewMenu; no marker (especially explicit empty/replaced spec) means no insertion/request. Validate collisions, keep every generated host row non-executable and all callbacks in the original VM. Add actual resolved-default aggregation of delivered session/tab/pane/app/host modules. Preserve endpoint binding/origin/expiry/privacy/owned-only cleanup. No independent Lua schema, generic framework, host attach/WOL/SSH or tests/experiments. This is approved under delegated technical authority, not another approval request. Mixed Go/Lua assembled source review is required.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDA3OjM2OjI4LjY0MFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNS10MDQtYnVpbGRlci1hMDIvc3VtbWFyeS8yIiwiaW50ZW50IjoiYzBhZTYyMjgyYmE1NDY5NjYyYjg3MWE2ZDA0YjY4Njg2NmM5MDU0Mzc2ZDZlOWZjN2Q2Y2Q5Y2NiYjY0MjcxOCJ9 -->
### 2026-11-01T07:36:28.640Z · summary · unspecified

{"acceptedScope":"Present Rex-known hosts and inventory warnings","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p05-t04-reviewer-a02/review/2 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p05-t04-builder-a02/integration/integration.json"],"attempt":2,"eventId":"rex-whichkey-p05-t04-builder-a02/summary/2","files":["docs/mapping-format.md","hosts.go","lua/rex_whichkey/defaults/hosts.lua","lua/rex_whichkey/defaults/init.lua","lua/rex_whichkey/evaluate.lua","lua/rex_whichkey/mapping.lua","mapping.go","terminal.go"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-05-04-host-menu.md","requirements":["REQ-012","REQ-013"],"schemaVersion":1,"specHash":"sha256:005489f744ce0c3785fab4cf9d901e690599ce6501e83348bd92a362e2abd626","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"cfba4cfd3e9a4192d52bdaa0cb9c4449c645964d","targetTree":"a61ab07fd504a43f6f1bb0bab9af78f591d29153","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
