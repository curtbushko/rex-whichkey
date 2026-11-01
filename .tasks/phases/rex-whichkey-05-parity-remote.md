---
workflowKind: phase
workflowOrder: 4
workflowRequirements:
  - REQ-003
  - REQ-005
  - REQ-010
  - REQ-011
  - REQ-012
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
title: 'Phase 5: Feature parity and remote context'
mtnCreateRetry:
  key: rex-whichkey/draft-phase/05-parity-remote/1
  intent: 4bb93002a38caa17fe98e882c53577c01a563d0a476601cf23810f7b56ccf301
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:55:59.429Z'
dateModified: '2026-11-01T07:51:45.228Z'
tags:
  - task
type: task
id: 8d4b8c61-eddb-4dc2-80c0-757f8f4ac038
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy: []
completedDate: '2026-11-01'
---
# Product implementation milestone

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable

Wire built-in session/tab/pane actions, delegated session picker behavior and Rex-known host listing into the implemented menu. Preserve origin targeting, confirmation and remote-safe paths without introducing connection/bootstrap/WOL workflows. Remote portability is implementation work, not a remote experiment harness.

## Requirement coverage

REQ-010, REQ-011, REQ-012, REQ-013. Child records retain their complementary functional responsibilities.

## Inputs and dependencies

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

No Phase1 or final-RFC prerequisite remains. Existing task-level product dependencies still apply; phase numbering does not create a dependency. Prior research is optional historical context, never an approval-required input.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Phase delivery policy

Follow projects/rex-whichkey.md: no project tests/builds/experiments/probes/remote trials; require independent source review. Commit each actually completed phase on main, separately from any metadata-only checkpoint. Completed tasks/phases retain concise actual-delivery summaries with reviewed files, requirements, limitations and commit pointers; active records retain necessary current state only. No proposed plan, source review alone without delivered implementation, successful worker exit or cancelled research establishes phase delivery. Coordinator uses supported native compaction at real completion, not a hand-edited history rewrite or duplicate archive.

## Implementation versus documentation

05-01/02/03 keep Lua default-mapping ownership and consume mapping.go and docs/mapping-format.md from 04-01, with explicit producer prerequisites. 05-04 keeps Go host-menu ownership. Remote execution/module resolution and context propagation are implemented by the existing Phase 2 client, 03-01 config, 03-04 dispatch and 04-05 popup owners. 05-05 owns docs/remote-compatibility.md only and records actual implementation, source-review findings and untested limitations. It is not a runtime implementation or remote trial task.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDA3OjUxOjQ1LjIxNloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcGhhc2UwNS9zdW1tYXJ5LzEiLCJpbnRlbnQiOiIyNjU0NDc1MTRjYjQ5ZmI0MjQyNjYyNGViMWMwMDAwOWVlZDZmMzY1OTM1YzA0MjNjYzJmYTEyZDM4MGFiNjYwIn0 -->
### 2026-11-01T07:51:45.216Z · summary · unspecified

{"acceptedScope":"Phase 5: Feature parity and remote context","approvalSummary":["User explicitly authorizes all remaining accepted phases3–6 and per-phase commits without routine confirmation.","All exact native child summaries bind current independent source approval and delivered target commit/tree; ancestry verified."],"artifactRefs":[".mtn-tmp/remaining-delivery/phase05-final/worktrees.txt",".mtn-tmp/remaining-delivery/phase05-final/validate.json"],"attempt":1,"coverage":{"REQ-003":["tasks/rex-whichkey-05-05-remote-context.md"],"REQ-005":["tasks/rex-whichkey-05-05-remote-context.md"],"REQ-010":["tasks/rex-whichkey-05-01-session-actions.md","tasks/rex-whichkey-05-02-tab-actions.md","tasks/rex-whichkey-05-03-pane-actions.md"],"REQ-011":["tasks/rex-whichkey-05-01-session-actions.md","tasks/rex-whichkey-05-02-tab-actions.md","tasks/rex-whichkey-05-03-pane-actions.md","tasks/rex-whichkey-05-05-remote-context.md"],"REQ-012":["tasks/rex-whichkey-05-04-host-menu.md"],"REQ-013":["tasks/rex-whichkey-05-04-host-menu.md","tasks/rex-whichkey-05-05-remote-context.md"]},"eventId":"rex-whichkey-p03-coordinator-a01/phase05/summary/1","files":["docs/mapping-format.md","docs/remote-compatibility.md","hosts.go","lua/rex_whichkey/defaults/app.lua","lua/rex_whichkey/defaults/hosts.lua","lua/rex_whichkey/defaults/init.lua","lua/rex_whichkey/defaults/panes.lua","lua/rex_whichkey/defaults/sessions.lua","lua/rex_whichkey/defaults/tabs.lua","lua/rex_whichkey/evaluate.lua","lua/rex_whichkey/mapping.lua","mapping.go","terminal.go"],"kind":"summary","limitations":["No project tests, TDD, experiments, GUI probes, live API/remote trials or config reloads; compilation/source review are not runtime verification.","Independent fresh contexts used the same authenticated model family, not cross-family review.","Prior transferred coordinator process status remains unverified; ownership transfer did not assert remote/local stoppage.","No global cleanup or push; retained worktrees/artifacts remain available."],"recordRef":"phases/rex-whichkey-05-parity-remote.md","records":["tasks/rex-whichkey-05-01-session-actions.md","tasks/rex-whichkey-05-02-tab-actions.md","tasks/rex-whichkey-05-03-pane-actions.md","tasks/rex-whichkey-05-04-host-menu.md","tasks/rex-whichkey-05-05-remote-context.md"],"requirements":["REQ-003","REQ-005","REQ-010","REQ-011","REQ-012","REQ-013"],"schemaVersion":1,"specHash":"sha256:17eea48be41caf6605d48d85dd7e67a1fbeb4c334585bad07e94513066c2b75a","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"cddf43711feeb2625be8f3ba5e719dde3bafab28","targetTree":"3f93623363683776bf507a2192558ba69aae4f54","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined delivered revision)","result":"pass"},{"command":"git diff --check; native graph/schema validation; exact completed child scope and Git delivery identities/ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
