---
workflowKind: task
workflowLanguage: lua
workflowOrder: 1
workflowRequirements:
  - REQ-003
  - REQ-004
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - lua/rex_whichkey/config.lua
    - tests/lua/config
    - examples
  sharedPaths: []
  resources: []
  interfaces:
    - whichkey-config
title: Implement inline and imported Lua configuration
mtnCreateRetry:
  key: rex-whichkey/draft-task/03-01-lua-config/1
  intent: 83009c5611c704597de2303f7fd22e27e12260478540f1121d6f78fcf54cfe44
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:17.390Z'
dateModified: '2026-11-01T01:01:54.143Z'
tags:
  - task
type: task
id: 16a49694-0172-41b4-a7a6-aa6fb9ebe848
projects:
  - '[[phases/rex-whichkey-03-lua-mappings.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-03-05-mapping-format.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-10-31'
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Implement the setup/config loader with collision-free integration/user module names, absent-value defaults and explicit precedence/replacement behavior. Keep executable configuration functions in Rex Lua.


## Functional criteria

REQ-003, REQ-004, REQ-013
Inline and sibling-module paths work in real server init.lua; duplicate/invalid setup is handled against the Go-owned standard and any explicitly approved public loader contract without corrupting the last working configuration. Server-side paths work without client-local absolute assumptions.


## Ownership

Allowed source/test edits: lua/rex_whichkey/config.lua, tests/lua/config, examples. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Authoritative format input

Consume mapping.go and docs/mapping-format.md from tasks/rex-whichkey-03-05-mapping-format.md after its native producer prerequisite is delivered. These Go-owned files define the shared mapping/configuration standard; do not define or copy an independent Lua/Nix/consumer schema. Executable configuration and callbacks remain in the originating Rex Lua runtime. Follow the existing functional requirements; raise a consequential new public interface or dependency choice for user approval instead of silently selecting it.

## Worker routing

Intended workflowLanguage: lua. Use the supported matching language/documentation worker and independent source-review route once the coordinator confirms the narrow tooling capability update. Do not dispatch through Go TDD gates or substitute Go metadata while support is unavailable. The project no-tests/no-experiments policy overrides worker defaults.

## Server-side loader responsibility

Implement remote-safe integration/user-module lookup and explicit precedence in lua/rex_whichkey/config.lua using the server host's installed modules and user configuration context. Do not resolve user modules by a client-local absolute path or take over init.lua. The Go producer owns configuration/default/replacement semantics in mapping.go and docs/mapping-format.md; this file only adapts Lua setup/loading to that standard. Functions remain in Rex Lua.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDAxOjAxOjU0LjEzMVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy10MDEtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiY2U1OTUzYjcxYWY1MjFmZGZkZmYyN2NmZjgyNTgyNmRmNjRkZGNmN2I4ZjMyYzg0MWM0YjkwOGUxZDNiNTBlMiJ9 -->
### 2026-11-01T01:01:54.131Z · summary · unspecified

{"acceptedScope":"Implement inline and imported Lua configuration","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p03-t01-tasks-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p03-t01-builder-a01/integration/integration.json"],"attempt":1,"eventId":"rex-whichkey-p03-t01-builder-a01/summary/1","files":["examples/README.md","examples/my_whichkey.lua","lua/rex_whichkey/config.lua"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-03-01-lua-config.md","requirements":["REQ-003","REQ-004","REQ-013"],"schemaVersion":1,"specHash":"sha256:334b767966694c3f7940807254c3f04d071d8b10369f125331351aa93d4f1dfb","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"6109ffff2935711a2d53fe641d18ed185a7e0ef3","targetTree":"15dbd5c7a99edb9186b46cf747550a3fa23abd87","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
