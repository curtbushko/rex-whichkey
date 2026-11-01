---
workflowKind: task
workflowLanguage: lua
workflowOrder: 2
workflowRequirements:
  - REQ-004
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - lua/rex_whichkey/mapping.lua
    - tests/lua/mapping
  sharedPaths: []
  resources: []
  interfaces:
    - whichkey-mapping-compiler
title: 'Compile key sequences, groups and inherited scopes'
mtnCreateRetry:
  key: rex-whichkey/draft-task/03-02-mapping-keys/1
  intent: 95e0eae9f536611a4c99352a126ef0e3555c0d7f2124caead5e05c09ce7a47ff
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:19.386Z'
dateModified: '2026-11-01T01:12:11.685Z'
tags:
  - task
type: task
id: 098f027d-48e8-4799-81b5-ec65e92d1c89
projects:
  - '[[phases/rex-whichkey-03-lua-mappings.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-03-01-lua-config.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-03-05-mapping-format.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-10-31'
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Compile full sequences, implicit groups and nested inherited attributes into the data-only mapping contract. Preserve behavior, not the old compiler architecture.


## Functional criteria

REQ-004
Support arbitrary depth, named actions/args, labels/icons and <Space>/<lt>; reject duplicate action mappings, group/action conflicts and unsupported key notation with actionable errors.


## Ownership

Allowed source/test edits: lua/rex_whichkey/mapping.lua, tests/lua/mapping. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


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

lua/rex_whichkey/mapping.lua compiles the user-facing Lua notation into the Go-owned representation. Keep only the necessary Lua adaptation/compilation; it neither defines a parallel schema nor moves Rex-runtime executable configuration into Go. Conflicts with the authoritative files must be resolved with the Go owner, not copied as divergent field definitions.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDAxOjEyOjExLjY3NFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy10MDItYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiODVhZGE0NzhlMjZlZDY2Nzc2NzViZWYxN2M1OTljOGY2ZTE2ZGIxY2E1YjY4MGMzNDE4NWE1ZDMwNGMxNDM3MiJ9 -->
### 2026-11-01T01:12:11.674Z · summary · unspecified

{"acceptedScope":"Compile key sequences, groups and inherited scopes","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p03-t02-tasks-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p03-t02-builder-a01/integration/integration.json"],"attempt":1,"eventId":"rex-whichkey-p03-t02-builder-a01/summary/1","files":["lua/rex_whichkey/mapping.lua"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-03-02-mapping-keys.md","requirements":["REQ-004"],"schemaVersion":1,"specHash":"sha256:ce55b445d2d930d8eca0d2b32cd57fca97a8034e6087cfedb70b45dec0fb3f6d","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"39ca52f3eace66a0a3a4df41b1fb6ac6d0efff3c","targetTree":"dd387e23de1f397a7dfe8a145085874fceef9659","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
