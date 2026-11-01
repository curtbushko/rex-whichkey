---
workflowKind: phase
workflowOrder: 1
workflowRequirements:
  - REQ-003
  - REQ-004
  - REQ-005
  - REQ-007
  - REQ-011
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
title: 'Phase 3: Lua integration and mapping compilation'
mtnCreateRetry:
  key: rex-whichkey/draft-phase/03-lua-mappings/1
  intent: 6c2ccbf6952c758b7bacc923a94233271905c1cda0137a9d0d376463cf2e1a80
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:55:55.471Z'
dateModified: '2026-11-01T06:26:41.263Z'
tags:
  - task
type: task
id: 4a761217-f5ea-4401-b5dc-159996e5ede3
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy: []
completedDate: '2026-11-01'
---
# Product implementation milestone

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable

Implement minimal Rex Lua configuration, full key-sequence compilation, implicit/nested groups, inherited attributes, dynamic labels/conditions, hidden active entries and named actions/callback references against the Go-owned standard in mapping.go and docs/mapping-format.md, produced by tasks/rex-whichkey-03-05-mapping-format.md. Mapping and configuration behavior comes first; executable callbacks stay in Rex Lua. Dispatch wiring follows the core menu and uses the documented API rather than an experimental protocol milestone.

## Requirement coverage

REQ-003, REQ-004, REQ-005, REQ-007, REQ-011, REQ-013. Child records retain their complementary functional responsibilities.

## Inputs and dependencies

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

No Phase1 or final-RFC prerequisite remains. Existing task-level product dependencies still apply; phase numbering does not create a dependency. Prior research is optional historical context, never an approval-required input.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Phase delivery policy

Follow projects/rex-whichkey.md: no project tests/builds/experiments/probes/remote trials; require independent source review. Commit each actually completed phase on main, separately from any metadata-only checkpoint. Completed tasks/phases retain concise actual-delivery summaries with reviewed files, requirements, limitations and commit pointers; active records retain necessary current state only. No proposed plan, source review alone without delivered implementation, successful worker exit or cancelled research establishes phase delivery. Coordinator uses supported native compaction at real completion, not a hand-edited history rewrite or duplicate archive.

## Producer, consumers and remote responsibilities

04-01 is the existing Go source-owning producer; its format/model work does not depend on Lua consumers. All Phase 3 tasks explicitly depend on that producer and use its files rather than copying schema prose or inventing a Lua-owned schema. Keep setup functions, dynamic functions and callback execution in the originating Rex Lua runtime. 03-01 owns server-side module resolution in lua/rex_whichkey/config.lua; 03-04 owns origin/server/session/client identity binding and callback dispatch in lua/rex_whichkey/dispatch.lua. Remote documentation only describes their delivered behavior.


## Go-owned standard producer
tasks/rex-whichkey-03-05-mapping-format.md owns mapping.go and docs/mapping-format.md as a small flat Go producer, independently of Phase4 menu completion. Lua configuration/compiler/evaluation/dispatch remain Lua consumers with real producer edges; no separately owned Lua schema or callback VM is introduced.

## Accepted bridge and current implementation authority
Accepted decision REXWHICHKEY-20261031-220348 resolves the prior interface choice, not the missing code. Reopened 03-05 Go standard and 03-04 Lua dispatch must deliver canonical authenticated post-return completion/cancellation for the detached popup, preserving the originating server Lua VM and captured identities instead of requiring the old action activation. Go owns shared bridge data/capability conventions; Lua retains opaque tickets and callbacks. Other delivered children remain valid. Complete and commit this phase again only after the revised producers are independently reviewed, integrated and checked, with a NEW current-spec final summary. Retained execution policy is in projects/rex-whichkey.md and .mtn-tmp/remaining-delivery/bridge-resume/accepted-authority.md.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDA2OjI2OjQxLjI1MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcGhhc2UwMy9zdW1tYXJ5LzIiLCJpbnRlbnQiOiI2ZGU0Y2VjMDUzZTkzZTI4Y2IyNGU3MTRmZDE0ZWNkMGQ3NTlkYjY5NzFmZTZhM2ZhYTQ0NjIzODRhZTk5OGQ4In0 -->
### 2026-11-01T06:26:41.251Z · summary · unspecified

{"acceptedScope":"Phase 3: Lua integration and mapping compilation","approvalSummary":["User explicitly authorizes all remaining accepted phases3–6 and per-phase commits without routine confirmation.","All exact native child summaries bind current independent source approval and delivered target commit/tree; ancestry verified."],"artifactRefs":[".mtn-tmp/remaining-delivery/phase03-final-r2/worktrees.txt",".mtn-tmp/remaining-delivery/phase03-final-r2/validate.json"],"attempt":2,"coverage":{"REQ-003":["tasks/rex-whichkey-03-01-lua-config.md","tasks/rex-whichkey-03-05-mapping-format.md"],"REQ-004":["tasks/rex-whichkey-03-01-lua-config.md","tasks/rex-whichkey-03-02-mapping-keys.md","tasks/rex-whichkey-03-03-dynamic-mappings.md","tasks/rex-whichkey-03-05-mapping-format.md"],"REQ-005":["tasks/rex-whichkey-03-03-dynamic-mappings.md","tasks/rex-whichkey-03-04-invocation-dispatch.md","tasks/rex-whichkey-03-05-mapping-format.md"],"REQ-007":["tasks/rex-whichkey-03-03-dynamic-mappings.md","tasks/rex-whichkey-03-04-invocation-dispatch.md"],"REQ-011":["tasks/rex-whichkey-03-04-invocation-dispatch.md"],"REQ-013":["tasks/rex-whichkey-03-01-lua-config.md","tasks/rex-whichkey-03-04-invocation-dispatch.md","tasks/rex-whichkey-03-05-mapping-format.md"]},"eventId":"rex-whichkey-p03-coordinator-a01/phase03/summary/2","files":["docs/mapping-format.md","examples/README.md","examples/my_whichkey.lua","lua/rex_whichkey/config.lua","lua/rex_whichkey/dispatch.lua","lua/rex_whichkey/evaluate.lua","lua/rex_whichkey/mapping.lua","mapping.go"],"kind":"summary","limitations":["No project tests, TDD, experiments, GUI probes, live API/remote trials or config reloads; compilation/source review are not runtime verification.","Independent fresh contexts used the same authenticated model family, not cross-family review.","Prior transferred coordinator process status remains unverified; ownership transfer did not assert remote/local stoppage.","No global cleanup or push; retained worktrees/artifacts remain available."],"recordRef":"phases/rex-whichkey-03-lua-mappings.md","records":["tasks/rex-whichkey-03-01-lua-config.md","tasks/rex-whichkey-03-02-mapping-keys.md","tasks/rex-whichkey-03-03-dynamic-mappings.md","tasks/rex-whichkey-03-04-invocation-dispatch.md","tasks/rex-whichkey-03-05-mapping-format.md"],"requirements":["REQ-003","REQ-004","REQ-005","REQ-007","REQ-011","REQ-013"],"schemaVersion":1,"specHash":"sha256:6a6413cb4d002fa08fb69c0a7542ff0e3218c8f77d27d6a0e7c63bbda27ed6af","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"0a0ac897c1807322e158e61ed637ca0f6833bea8","targetTree":"c0b751a036db26087d749ef319df626541955da5","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined delivered revision)","result":"pass"},{"command":"git diff --check; native graph/schema validation; exact completed child scope and Git delivery identities/ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
