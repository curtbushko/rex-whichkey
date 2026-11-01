---
workflowKind: task
workflowLanguage: lua
workflowOrder: 3
workflowRequirements:
  - REQ-004
  - REQ-005
  - REQ-007
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - lua/rex_whichkey/evaluate.lua
    - tests/lua/evaluate
  sharedPaths: []
  resources: []
  interfaces:
    - whichkey-mapping-data
title: Evaluate conditional and dynamic mapping metadata safely
mtnCreateRetry:
  key: rex-whichkey/draft-task/03-03-dynamic-mappings/1
  intent: 5031bd3f7f72bc035a76bb5e95482332acf752755738f6405538ee8cc88f06ef
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:21.389Z'
dateModified: '2026-11-01T01:21:47.154Z'
tags:
  - task
type: task
id: c5a227ec-8ed1-454a-b60d-a29f9fa6b34f
projects:
  - '[[phases/rex-whichkey-03-lua-mappings.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-03-02-mapping-keys.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-03-05-mapping-format.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-10-31'
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Evaluate boolean/function conditions and string-valued description/group/icon functions when producing invocation metadata; retain hidden-but-active entries and opaque executable callback identities.


## Functional criteria

REQ-004, REQ-005, REQ-007
Hidden entries dispatch through keys without appearing in views; dynamic/conditional values inherit attributes correctly. Callback functions never enter JSON. Stale conditions, callback references and evaluation errors fail safely within the invocation, including confirmation flags.


## Ownership

Allowed source/test edits: lua/rex_whichkey/evaluate.lua, tests/lua/evaluate. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Authoritative format input

Consume mapping.go and docs/mapping-format.md from tasks/rex-whichkey-03-05-mapping-format.md after its native producer prerequisite is delivered. These Go-owned files define the shared mapping/configuration standard; do not define or copy an independent Lua/Nix/consumer schema. Executable configuration and callbacks remain in the originating Rex Lua runtime. Follow the existing functional requirements; raise a consequential new public interface or dependency choice for user approval instead of silently selecting it.

## Worker routing

Intended workflowLanguage: lua. Use the supported matching language/documentation worker and independent source-review route once the coordinator confirms the narrow tooling capability update. Do not dispatch through Go TDD gates or substitute Go metadata while support is unavailable. The project no-tests/no-experiments policy overrides worker defaults.

## Boundary

lua/rex_whichkey/evaluate.lua evaluates configuration functions in the originating Rex runtime and produces data-only invocation values according to the Go-owned standard. Callback identities are opaque data references, not serialized functions. Do not introduce a second Lua schema or an independent Go Lua runtime for user callbacks.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDAxOjIxOjQ3LjE0MloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy10MDMtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiZWY2NWQ4YTVhZDhjMjVkZWUzNWUwMmY5N2YxMTYwOTE3NGVkYzIwOGI0NTY1MThkNWUxMTdkNTVmYjc3ZWZjNyJ9 -->
### 2026-11-01T01:21:47.142Z · summary · unspecified

{"acceptedScope":"Evaluate conditional and dynamic mapping metadata safely","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p03-t03-tasks-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p03-t03-builder-a01/integration/integration.json"],"attempt":1,"eventId":"rex-whichkey-p03-t03-builder-a01/summary/1","files":["lua/rex_whichkey/evaluate.lua"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-03-03-dynamic-mappings.md","requirements":["REQ-004","REQ-005","REQ-007"],"schemaVersion":1,"specHash":"sha256:cc1e5010665b1432299c5f8709d039cf697cd8a30f278abb06c09594c41b9fdc","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"44f22c10e22abc2c5d354ff7c74530a6985cffe1","targetTree":"bfd0b846865e7dca6617c2dcda53b6317aee5d76","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
