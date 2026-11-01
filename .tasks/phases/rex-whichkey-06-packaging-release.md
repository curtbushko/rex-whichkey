---
workflowKind: phase
workflowOrder: 5
workflowRequirements:
  - REQ-011
  - REQ-014
  - REQ-015
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
title: 'Phase 6: Packaging and release verification'
mtnCreateRetry:
  key: rex-whichkey/draft-phase/06-packaging-release/1
  intent: 88f5c36a6e027d5625f229f8ba3f07c5a1e37f2cec79e15c257c36af5f56e98e
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:01.413Z'
dateModified: '2026-11-01T08:16:31.776Z'
tags:
  - task
type: task
id: 2b321d15-5423-45af-95b5-7b09a3d5ec7a
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy: []
completedDate: '2026-11-01'
---
# Product implementation milestone

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable

Package the implemented application and Lua integration with Nix and opt-in Home Manager; document configuration, mappings, usage and known limitations. Preserve user init.lua. Prepare an honest release delivery inventory; do not create or run another acceptance harness or resurrect the deferred disconnected-origin experiment. No release pass is asserted.

## Requirement coverage

REQ-011, REQ-014, REQ-015. Child records retain their complementary functional responsibilities.

## Inputs and dependencies

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

No Phase1 or final-RFC prerequisite remains. Existing task-level product dependencies still apply; phase numbering does not create a dependency. Prior research is optional historical context, never an approval-required input.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Phase delivery policy

Follow projects/rex-whichkey.md: no project tests/builds/experiments/probes/remote trials; require independent source review. Commit each actually completed phase on main, separately from any metadata-only checkpoint. Completed tasks/phases retain concise actual-delivery summaries with reviewed files, requirements, limitations and commit pointers; active records retain necessary current state only. No proposed plan, source review alone without delivered implementation, successful worker exit or cancelled research establishes phase delivery. Coordinator uses supported native compaction at real completion, not a hand-edited history rewrite or duplicate archive.

## Packaging producers and review

06-01 depends on the actual popup, session/tab/pane defaults and host-menu producers before claiming an integrated package deliverable. Its source inventory includes the authoritative Go format and shipped Lua consumers/defaults, not an early partial assembly. Nix packaging and Home Manager route through Nix source review; release documentation routes through documentation review. No Go verification gates, package builds, Nix evaluation, graphical-client reservations or remote-test-server reservations are scheduled.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDA4OjE2OjMxLjc2NFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcGhhc2UwNi9zdW1tYXJ5LzEiLCJpbnRlbnQiOiJkMzdhNWNkNTdiMDI5OTY1OTJlNzQ0MDAxMDY3NzUyMGFmODYyM2ZiNDBiMmEzOGM2MDY3YjM1OTU2NWUxZWQ4In0 -->
### 2026-11-01T08:16:31.764Z · summary · unspecified

{"acceptedScope":"Phase 6: Packaging and release verification","approvalSummary":["User explicitly authorizes all remaining accepted phases3–6 and per-phase commits without routine confirmation.","All exact native child summaries bind current independent source approval and delivered target commit/tree; ancestry verified."],"artifactRefs":[".mtn-tmp/remaining-delivery/phase06-final/worktrees.txt",".mtn-tmp/remaining-delivery/phase06-final/validate.json"],"attempt":1,"coverage":{"REQ-011":["tasks/rex-whichkey-06-03-release-acceptance.md"],"REQ-014":["tasks/rex-whichkey-06-01-nix-package.md","tasks/rex-whichkey-06-02-home-manager.md"],"REQ-015":["tasks/rex-whichkey-06-03-release-acceptance.md"]},"eventId":"rex-whichkey-p03-coordinator-a01/phase06/summary/1","files":["README.md","docs/acceptance.md","flake.nix","nix/home-manager.nix","nix/package.nix","tests/packaging/home-manager/README.md"],"kind":"summary","limitations":["No project tests, TDD, experiments, GUI probes, live API/remote trials or config reloads; compilation/source review are not runtime verification.","Independent fresh contexts used the same authenticated model family, not cross-family review.","Prior transferred coordinator process status remains unverified; ownership transfer did not assert remote/local stoppage.","No global cleanup or push; retained worktrees/artifacts remain available."],"recordRef":"phases/rex-whichkey-06-packaging-release.md","records":["tasks/rex-whichkey-06-01-nix-package.md","tasks/rex-whichkey-06-02-home-manager.md","tasks/rex-whichkey-06-03-release-acceptance.md"],"requirements":["REQ-011","REQ-014","REQ-015"],"schemaVersion":1,"specHash":"sha256:cef9637bf97e5df947d17547f8b84669d41c7262ab99fc0ccac6ddacad8ef6ba","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"40c4f53a8b12f71245057cecaa2c18286d7d9c25","targetTree":"5b8ac6d2d75121426c26188c53d73e5ee261d313","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined delivered revision)","result":"pass"},{"command":"git diff --check; native graph/schema validation; exact completed child scope and Git delivery identities/ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
