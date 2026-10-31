---
workflowKind: phase
workflowOrder: 3
workflowRequirements:
  - REQ-002
  - REQ-011
  - REQ-012
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
title: 'Phase 2: Concrete Rex CLI client'
mtnCreateRetry:
  key: rex-whichkey/draft-phase/02-rex-client/1
  intent: 101950907959ee46e985b3d1be8b22feb042fcdef98be290fea65a9277e575ec
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:55:53.493Z'
dateModified: '2026-10-31T20:55:29.672Z'
tags:
  - task
type: task
id: 27b3cef3-37f5-4e51-9430-6a72d0dc1a6f
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy: []
completedDate: '2026-10-31'
---
# Product implementation milestone

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable

Implement the small concrete Rex client through documented CLI/API calls: safe process execution, resource operations, current block lookup, popup lifecycle and known-host listing. Build only the operations consumed by which-key; no independent Rex validation program.

## Requirement coverage

REQ-002, REQ-011, REQ-012, REQ-013. Child records retain their complementary functional responsibilities.

## Inputs and dependencies

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

No Phase1 or final-RFC prerequisite remains. Existing task-level product dependencies still apply; phase numbering does not create a dependency. Prior research is optional historical context, never an approval-required input.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Phase delivery policy

Follow projects/rex-whichkey.md: no project tests/builds/experiments/probes/remote trials; require independent source review. Commit each actually completed phase on main, separately from any metadata-only checkpoint. Completed tasks/phases retain concise actual-delivery summaries with reviewed files, requirements, limitations and commit pointers; active records retain necessary current state only. No proposed plan, source review alone without delivered implementation, successful worker exit or cancelled research establishes phase delivery. Coordinator uses supported native compaction at real completion, not a hand-edited history rewrite or duplicate archive.

## Remote implementation ownership

Child 02-01 owns remote executable selection and server/session CLI targeting in rex/rex.go and rex/exec.go; 02-02 owns stable targeting in rex/session.go and rex/window.go; 02-03 owns current remote origin resolution and owned-resource cleanup in rex/block.go and rex/layer.go. 02-04 retains the documented effective-inventory behavior in rex/hosts.go; remote server targeting must not be misapplied to rex hosts. Remote documentation does not implement these paths.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDIwOjU1OjI5LjY2MFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcGhhc2Utc3VtbWFyeS8zIiwiaW50ZW50IjoiMWUzYzg5ZjUyY2ExY2Y0NDZjMDc0YTZlZGRmOTRhMmMzZGQxZmNiMWVlOTU3NzI1ZmFjNTdlYTQ0NTkyYjYxZSJ9 -->
### 2026-10-31T20:55:29.660Z · summary · unspecified

{"acceptedScope":"Phase2 concrete Rex client: safe targeted CLI/generic calls, session/window resources, moved-block focus, owned-only popup layer lifecycle and effective known-host inventory.","approvalSummary":["User explicitly approved corrected records and resumed Phase2 implementation; no tests/experiments and per-phase commit policy retained.","rex-whichkey-p02-t01-reviewer-a01/review/1 APPROVED current revision in a fresh independent source-review context.","rex-whichkey-p02-t02-reviewer-a01/review/1 APPROVED current revision in a fresh independent source-review context.","rex-whichkey-p02-t04-reviewer-a01/review/1 APPROVED current revision in a fresh independent source-review context.","rex-whichkey-p02-t03-reviewer-a02/review/2 APPROVED current revision in a fresh independent source-review context."],"artifactRefs":[".mtn-tmp/phase2-delivery/final/native-validation.json",".mtn-tmp/phase2-delivery/final/combined-build.log",".mtn-tmp/phase2-delivery/final/worktree-inventory.txt",".mtn-tmp/phase2-delivery/final/cleanup-plan.json"],"attempt":1,"coordinator":{"processStatus":"Running finalization; phase scheduling and source integration finished. No process stoppage claimed.","runId":"rex-whichkey-p03-coordinator-a01"},"coverage":{"REQ-002":["tasks/rex-whichkey-02-01-cli-execution.md","tasks/rex-whichkey-02-02-resources.md","tasks/rex-whichkey-02-03-block-layers.md","tasks/rex-whichkey-02-04-host-inventory.md"],"REQ-011":["tasks/rex-whichkey-02-03-block-layers.md"],"REQ-012":["tasks/rex-whichkey-02-04-host-inventory.md"],"REQ-013":["tasks/rex-whichkey-02-01-cli-execution.md","tasks/rex-whichkey-02-02-resources.md","tasks/rex-whichkey-02-03-block-layers.md","tasks/rex-whichkey-02-04-host-inventory.md"]},"eventId":"rex-whichkey-p03-coordinator-a01/phase-summary/3","files":["rex/block.go","rex/exec.go","rex/hosts.go","rex/layer.go","rex/rex.go","rex/session.go","rex/window.go"],"finalizationNote":"Native summary republication recovers compaction readback after an idempotent prepublished summary was no longer the trailing note; no source, authority or delivery proof changed.","kind":"summary","limitations":["No project tests, TDD, experiments, GUI probes or live/remote Rex trials were run for this implementation. Default team test/lint execution gates were not claimed passed.","Independent source review and standard compilation do not establish runtime or cross-platform behavior. Reviews used fresh contexts in the same authenticated model family, not cross-family verification.","Cleanup refuses unsafe remaining layer removal instead of cascading into unrelated blocks; an unprovably safe remaining layer can require later cleanup.","Earlier metadata checkpoint 7d4cff8 inadvertently ran repository build/lint/unit/race hooks. Earlier no-tests reporting was corrected in project evidence; those runs are not acceptance evidence for this delivery. Later commits suppressed hooks per invocation without changing their configuration.","No Go-Lua dependency introduced or selected. Other phases were not executed. The project remains incomplete.","Inherited remote process stoppage remains unverified. Retained worktrees and transport artifacts were inventoried, not deleted; no push or global cleanup."],"recordRef":"phases/rex-whichkey-02-rex-client.md","records":["tasks/rex-whichkey-02-01-cli-execution.md","tasks/rex-whichkey-02-02-resources.md","tasks/rex-whichkey-02-03-block-layers.md","tasks/rex-whichkey-02-04-host-inventory.md"],"requirements":["REQ-002","REQ-011","REQ-012","REQ-013"],"schemaVersion":1,"specHash":"sha256:adec86cfe4d5dd970e4611a50f5e14bbebf3a46c8769697c741a3757a5d14d04","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"fb7ed1228620a335d0e325e88a66ee3c024dfef6","targetTree":"c9985ba3a14e69000dd9b9a4b341b7bb07c67080","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... on combined delivered Phase2 revision","result":"pass"},{"command":"gofmt -l rex/*.go (empty output); git diff --check","result":"pass"},{"command":"mtn agent validate --json; phase hierarchy has four native-completed children and no cycles","result":"pass"},{"command":"Verify all four child delivery commit/tree identities and git merge-base --is-ancestor delivered-commit main","result":"pass"},{"command":"git diff --exit-code e89bfed HEAD -- go.mod go.sum","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
