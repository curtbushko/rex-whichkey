---
workflowKind: phase
workflowOrder: 2
workflowRequirements:
  - REQ-001
  - REQ-005
  - REQ-006
  - REQ-007
  - REQ-008
  - REQ-009
  - REQ-011
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
title: 'Phase 4: Menu and terminal UI'
mtnCreateRetry:
  key: rex-whichkey/draft-phase/04-menu-ui/1
  intent: 62362c704da8782a00b2f83563a15cc3017d22a8e0a604c1f000d967951e8f78
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:55:57.446Z'
dateModified: '2026-11-01T07:04:52.258Z'
tags:
  - task
type: task
id: d7fce24c-6130-4931-b28c-cdbf6efcf34c
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy: []
completedDate: '2026-11-01'
---
# Product implementation milestone

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable

Deliver the core which-key menu: prefix navigation, arbitrary-depth groups, breadcrumbs, Esc/Backspace, conditional q, Tab paging, confirmation, constrained Unicode-aware layout, theme, input and terminal restoration. The Go menu owns behavior and rendering. Finish popup integration after mapping/dispatch and the minimal Rex client exist; no Rex experiment is a prerequisite for core menu work.

## Requirement coverage

REQ-001, REQ-003, REQ-004, REQ-005, REQ-006, REQ-007, REQ-008, REQ-009, REQ-011, REQ-013. Child records retain their complementary functional responsibilities.

## Inputs and dependencies

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

No Phase1 or final-RFC prerequisite remains. Existing task-level product dependencies still apply; phase numbering does not create a dependency. Prior research is optional historical context, never an approval-required input.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Phase delivery policy

Follow projects/rex-whichkey.md: no project tests/builds/experiments/probes/remote trials; require independent source review. Commit each actually completed phase on main, separately from any metadata-only checkpoint. Completed tasks/phases retain concise actual-delivery summaries with reviewed files, requirements, limitations and commit pointers; active records retain necessary current state only. No proposed plan, source review alone without delivered implementation, successful worker exit or cancelled research establishes phase delivery. Coordinator uses supported native compaction at real completion, not a hand-edited history rewrite or duplicate archive.


## Go-owned standard input
The Go format-only producer is tasks/rex-whichkey-03-05-mapping-format.md, owning mapping.go and docs/mapping-format.md. Navigation and downstream menu/popup consumers use that standard after its actual producer prerequisite; the whole menu milestone does not gate Phase3 configuration.
<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDA3OjA0OjUyLjI0N1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcGhhc2UwNC9zdW1tYXJ5LzEiLCJpbnRlbnQiOiJlZjc0OWIzN2M3N2FmYTBiNTA2MDY2Njk2OTI5Y2E3ZWRiYjhiNjIwYzI3YjM1YWFlNDJkMGQ2OWU3ZmM2Mzk5In0 -->
### 2026-11-01T07:04:52.247Z · summary · unspecified

{"acceptedScope":"Phase 4: Menu and terminal UI","approvalSummary":["User explicitly authorizes all remaining accepted phases3–6 and per-phase commits without routine confirmation.","All exact native child summaries bind current independent source approval and delivered target commit/tree; ancestry verified."],"artifactRefs":[".mtn-tmp/remaining-delivery/phase04-final/worktrees.txt",".mtn-tmp/remaining-delivery/phase04-final/validate.json"],"attempt":1,"coverage":{"REQ-001":["tasks/rex-whichkey-04-04-terminal-lifecycle.md","tasks/rex-whichkey-04-05-popup-integration.md"],"REQ-005":["tasks/rex-whichkey-04-05-popup-integration.md"],"REQ-006":["tasks/rex-whichkey-04-01-navigation.md","tasks/rex-whichkey-04-02-layout.md"],"REQ-007":["tasks/rex-whichkey-04-01-navigation.md","tasks/rex-whichkey-04-05-popup-integration.md"],"REQ-008":["tasks/rex-whichkey-04-02-layout.md","tasks/rex-whichkey-04-04-terminal-lifecycle.md","tasks/rex-whichkey-04-05-popup-integration.md"],"REQ-009":["tasks/rex-whichkey-04-03-theme.md","tasks/rex-whichkey-04-04-terminal-lifecycle.md"],"REQ-011":["tasks/rex-whichkey-04-05-popup-integration.md"],"REQ-013":["tasks/rex-whichkey-04-05-popup-integration.md"]},"eventId":"rex-whichkey-p03-coordinator-a01/phase04/summary/1","files":["action.go","cmd/main.go","cmd/main_test.go","config.go","go.mod","go.sum","layout.go","lua/rex_whichkey.lua","menu.go","popup.go","render.go","terminal.go","theme.go"],"kind":"summary","limitations":["No project tests, TDD, experiments, GUI probes, live API/remote trials or config reloads; compilation/source review are not runtime verification.","Independent fresh contexts used the same authenticated model family, not cross-family review.","Prior transferred coordinator process status remains unverified; ownership transfer did not assert remote/local stoppage.","No global cleanup or push; retained worktrees/artifacts remain available."],"recordRef":"phases/rex-whichkey-04-menu-ui.md","records":["tasks/rex-whichkey-04-01-navigation.md","tasks/rex-whichkey-04-02-layout.md","tasks/rex-whichkey-04-03-theme.md","tasks/rex-whichkey-04-04-terminal-lifecycle.md","tasks/rex-whichkey-04-05-popup-integration.md"],"requirements":["REQ-001","REQ-005","REQ-006","REQ-007","REQ-008","REQ-009","REQ-011","REQ-013"],"schemaVersion":1,"specHash":"sha256:827f079e127d3f65ebf45be9d03dd8fd95d2906b3896695a8ea5cd0f7f6d80a7","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"8f2a196e50fd58848540d6ba0ca6b192385b7c30","targetTree":"38b51a2b64eb33e7030aac609e525e23b5e486f1","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined delivered revision)","result":"pass"},{"command":"git diff --check; native graph/schema validation; exact completed child scope and Git delivery identities/ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
