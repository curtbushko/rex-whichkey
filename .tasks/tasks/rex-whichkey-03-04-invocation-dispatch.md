---
workflowKind: task
workflowLanguage: lua
workflowOrder: 4
workflowRequirements:
  - REQ-005
  - REQ-007
  - REQ-011
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - lua/rex_whichkey/dispatch.lua
    - tests/lua/dispatch
  sharedPaths: []
  resources: []
  interfaces:
    - whichkey-invocation-dispatch
title: Implement origin-aware invocation registry and dispatch
mtnCreateRetry:
  key: rex-whichkey/draft-task/03-04-invocation-dispatch/1
  intent: 87ba5fc64c398c84e93f03101eaad0ca1c9f6c7bc7b23af9839a0d8b86a93c64
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:23.380Z'
dateModified: '2026-11-01T06:25:49.891Z'
tags:
  - task
type: task
id: 3127581d-6d5d-4c7b-ac0e-dd3482abb0e2
projects:
  - '[[phases/rex-whichkey-03-lua-mappings.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-03-03-dynamic-mappings.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-03-05-mapping-format.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-01'
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Implement minimal invocation-bound callback/action dispatch and data-only selection validation; retain executable callbacks in the persistent Rex runtime and route named actions through the correct client/server mechanism.


## Functional criteria

REQ-005, REQ-007, REQ-011, REQ-013
Handle overlapping invocations in reverse order, replay, invalid/expired capabilities, wrong origin/client, reload invalidation, missing/moved origins, confirmation, cancellation, stale conditions, callback failures and reentrancy. State is bounded. Focus origin before execution; never automatically retry a timed-out destructive action.


## Ownership

Allowed source/test edits: lua/rex_whichkey/dispatch.lua, tests/lua/dispatch. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Authoritative format input

Consume mapping.go and docs/mapping-format.md from tasks/rex-whichkey-03-05-mapping-format.md after its native producer prerequisite is delivered. These Go-owned files define the shared mapping/configuration standard; do not define or copy an independent Lua/Nix/consumer schema. Executable configuration and callbacks remain in the originating Rex Lua runtime. Follow the existing functional requirements; raise a consequential new public interface or dependency choice for user approval instead of silently selecting it.

## Worker routing

Intended workflowLanguage: lua. Use the supported matching language/documentation worker and independent source-review route once the coordinator confirms the narrow tooling capability update. Do not dispatch through Go TDD gates or substitute Go metadata while support is unavailable. The project no-tests/no-experiments policy overrides worker defaults.

## Remote implementation responsibility

lua/rex_whichkey/dispatch.lua owns binding and checking the originating server/session/block/client and invocation identity when issuing/validating data-only selections and invoking Rex-runtime functions. Preserve the existing bounded lifetime, replay, focus and no-destructive-retry requirements. Use the Go-owned format files, not a Lua-defined transport schema. Actual transport/API signatures remain subject to user approval when consequential; package research does not approve them. 05-05 documents, but cannot implement, this code.


## Accepted bridge and current implementation authority
Accepted decision REXWHICHKEY-20261031-220348 authorizes an authenticated one-shot post-return completion/cancellation entry point in the same server Lua VM. Consume the revised canonical Go bridge standard in mapping.go/docs/mapping-format.md; do not create a Lua-owned wire schema or move callback execution into another VM. Retain opaque Lua tickets/functions/original context privately; bind unpredictable capabilities to captured server/session/block/client identity, connection generation and configuration generation. A documented fresh completion action must be explicitly origin-targeted, not popup/current-client fallback. Verify authoritative fresh-action binding before side effects, consume/revoke before yielding/reentry, preserve snapshot selection/ref/confirmation validation, expiry/cancellation/invalidation, moved-origin focus, named-action queue and callback safety. Do not merely delete rawequal checks or assume copied ctx data recreates an activation. Validate actual Rex action/CLI contracts from source/docs; implement a minimal user-owned registration handler/snippet contract for 04-05 composition, without rewriting init.lua. Normal interface/implementation choices are now delegated within accepted scope; unsupported API capabilities remain genuine blockers, not invented success. No tests/experiments/runtime trials; new current source review/integration required.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDA2OjI1OjQ5Ljg3OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy10MDQtYnVpbGRlci1hMDIvc3VtbWFyeS8yIiwiaW50ZW50IjoiMmZlMTZhYzg3MmYyYjcyN2FmNTdiZDVhM2Q2N2EwNGRlM2I0NDhmMmQ2YzYwN2YyNzdlNzMxZTVmNjE5ZDI4YyJ9 -->
### 2026-11-01T06:25:49.879Z · summary · unspecified

{"acceptedScope":"Implement origin-aware invocation registry and dispatch","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p03-t04-tasks-a02/review/2 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p03-t04-builder-a02/integration/integration.json"],"attempt":2,"eventId":"rex-whichkey-p03-t04-builder-a02/summary/2","files":["lua/rex_whichkey/dispatch.lua"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-03-04-invocation-dispatch.md","requirements":["REQ-005","REQ-007","REQ-011","REQ-013"],"schemaVersion":1,"specHash":"sha256:03c113e08712f2425e255163be084e1d71164dce7d1efd98867b4528d3420ead","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"0a0ac897c1807322e158e61ed637ca0f6833bea8","targetTree":"c0b751a036db26087d749ef319df626541955da5","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
