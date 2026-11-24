---
identity:
  kind: phase
  phaseID: rex-whichkey-21-lint-toolchain
  uuid: 049c931e-44f3-40f2-a730-73f20b23c5bf
structure:
  order: 21
scope:
  requirements:
    - REQ-ENV-001
    - REQ-ENV-002
    - REQ-ENV-003
  documents:
    - ref: designs/rex-whichkey-lint-toolchain.md
      hash: 'sha256:5d3ea3bebc1b203988b65e69ac1c0838025d7287509c668817fd4aa9e4a8e606'
    - ref: designs/rex-whichkey-quality-gates.md
      hash: 'sha256:f3120762f7075c7c767e5f39552ba9c2b38b2e213d33d2472d07f0bc0c1cbdb2'
  dependencies: []
lifecycle:
  state: complete
  sourceEventID: rex-whichkey-p21/summary/1
  changedAt: '2026-11-24T15:28:44.424Z'
title: 'Phase21: compatible reproducible lint environment'
mtnCreateRetry:
  key: rex-whichkey-p21-phase-v1
  intent: 24bd5ca7fbb06e0707028d7d0f26bdd4b7ed13486a4d2b88b47434e6ddb2370e
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-24T13:50:14.445Z'
dateModified: '2026-11-24T15:28:45.450Z'
tags:
  - task
type: task
id: 8a3850f5-28f3-4c2b-ab5e-cf8915fadd6a
projects:
  - '[[projects/rex-whichkey.md]]'
completedDate: '2026-11-24'
---
# Phase21: reproducible default lint environment

User explicitly approved fixing flake.nix after golangci-lint2.8.0 built with Go1.25 rejected Go1.26.8. Deliver only the bounded toolchain repair in designs/rex-whichkey-lint-toolchain.md, REQ-ENV-001..003, under current mandatory quality authority designs/rex-whichkey-quality-gates.md.

Preserve application Go1.26.8, current lint rules/assertions/hooks, runtime/application source, existing shellHook semantics and unrelated input pins. Compatible development tools and a meaningful same-tool regression check must replace temporary PATH workarounds. Verify actual declared environment without running the auto-pull/config-copy shellHook. No profile/config edits, installation/activation/reload, push, broad dependency refresh or GC.

One tightly coupled maintainer task owns Nix/Bash implementation plus tests; exact ownership/reviews and fingerprints recorded in its approved plan. Native parenting to projects/rex-whichkey.md; no numerical prerequisite on completed Phase20. The existing sole coordinator run rex-whichkey-p20-coordinator-a01 anchored on Phase20 continues this newly authorized scope in the existing Rex session; no new coordinator/session or ownership handoff.

Completion requires independent separate Nix and Bash code/test plus maintenance assessments on a frozen candidate, exact-tree speculative integration, actual declared-environment build/test/race/full lint and Nix check/evaluation evidence, owned task cleanup, and concise delivered scenario assessment. Commit reviewed delivery on main, with metadata separate from source.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDE1OjI4OjQzLjI2MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMS9zdW1tYXJ5LzEiLCJpbnRlbnQiOiI0YmUzOTM1ZGIyMjRiNDQzOWYxMTVjMzk4ZWRmNWJlMjQ2NzE3ZDc3MWE4Nzc2YWU3NDVkZWFlZjY4YTkzODE4In0 -->
### 2026-11-24T15:28:43.261Z · summary · unspecified

{"acceptedScope":"User-authorized bounded repair of reproducible default Nix lint tooling for application Go1.26.8; no application Go upgrade, broad input refresh, profile activation or quality-rule weakening.","approvalSummary":["User explicitly approved yes, fix flake.nix after reported old-linter failure and recommendation to align reproducible compatible tools","Coordinator bounded phase/task plan under that authority; no claim of separate user approval of routine details","Five distinct current-build Nix/Bash code/test and maintenance keys APPROVED on exact frozen head/tree","Read-only product-manager assessed actual delivered workflow evidence; all three requirements demonstrated, no unmet/unverified requirements, blockers or user decisions"],"attempt":1,"cleanup":{"branchesAndTransportPreserved":true,"facadeCompatibility":"Reused previously validated private cleanup release-guard syntax correction; installed tools/repository implementation untouched","integrationWorktreeRemoved":true,"plannerWorktreeRemoved":true,"retainedRecoveryWindows":[{"reason":"Stopped negative-review window retained by policy","rexWindowID":"window:01a20dea-1087-724d-876d-319a5bf56d12","runID":"rex-whichkey-p21-t01-nix-test-reviewer-a01"}],"successfulWorkerWindowsClosed":true,"taskWorktreesRemoved":true,"unlaunchedPreparedWorktreeRemoved":true},"coordinator":{"anchorPhaseRef":"phases/rex-whichkey-20-background-scripts.md","runID":"rex-whichkey-p20-coordinator-a01","state":"Live original coordinator intentionally retained; never falsely released for compaction"},"coverage":{"REQ-ENV-001":["tasks/rex-whichkey-21-01-lint-toolchain.md"],"REQ-ENV-002":["tasks/rex-whichkey-21-01-lint-toolchain.md"],"REQ-ENV-003":["tasks/rex-whichkey-21-01-lint-toolchain.md"]},"eventID":"rex-whichkey-p21/summary/1","files":["flake.nix","nix/lint-toolchain-check.nix",".mtn-tmp/p21/integration-accepted-a02.json",".mtn-tmp/p21/main-gates/",".mtn-tmp/p21/integration-a02/",".mtn-tmp/p21/builder-a02-evidence/",".mtn-tmp/p21/review-nix-code-a02-evidence/",".mtn-tmp/p21/review-nix-test-a02-evidence/",".mtn-tmp/p21/review-bash-code-a02-evidence/",".mtn-tmp/p21/review-bash-test-a02-evidence/",".mtn-tmp/p21/review-maintenance-a02-evidence/",".mtn-tmp/p21/cleanup/",".tasks/designs/rex-whichkey-lint-toolchain.md",".mtn-tmp/p21/task-summary.json",".mtn-tmp/p21/outcome-result.json"],"kind":"summary","limitations":["Actual runtime/build verification on aarch64-darwin; other three supported systems evaluated only","Independent fresh same-family model contexts, not cross-family review","User shell/profile/config not activated or reloaded; caller must refresh its development environment to select declared tools","One stopped negative-review window retained by recovery policy; no delivery blocker"],"outcomeAssessment":{"file":".mtn-tmp/p21/outcome-result.json","recommendation":"complete","unresolvedBlockers":[],"userDecisions":[]},"recordRef":"phases/rex-whichkey-21-lint-toolchain.md","records":["tasks/rex-whichkey-21-01-lint-toolchain.md"],"requirements":["REQ-ENV-001","REQ-ENV-002","REQ-ENV-003"],"reviewEventIDs":["rex-whichkey-p21-t01/review/nix-code/2","rex-whichkey-p21-t01/review/nix-test/2","rex-whichkey-p21-t01/review/bash-code/2","rex-whichkey-p21-t01/review/bash-test/2","rex-whichkey-p21-t01/review/maintenance/2"],"schemaVersion":1,"specHash":"sha256:7610ea29103ef180c5d9020abd27f8dd581bf1b6e5d72d8d83dfd40799836d5c","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"699b53e05d966c3f9e1326bbcf6a9e3e201598ba","targetTree":"95f434d0b08e07fa4146d0df2ad4b10e33f840eb","unresolvedBlockers":[],"verification":[{"command":"nix flake check --all-systems --no-build --no-write-lock-file","result":"pass"},{"command":"host lint-toolchain check, repeated actual --rebuild, old Go1.25-linter rejection and both real analyzer diagnostics","result":"pass"},{"command":"fresh default declared environment command/compiler provenance and unchanged protected input/Go/checksum/shellHook assertions","result":"pass"},{"command":"task build on combined candidate and final main","result":"pass"},{"command":"task test unit and race on combined candidate and final main","result":"pass"},{"command":"FULL unchanged task lint on combined candidate and final main","result":"pass"},{"command":"exact reviewed source tree, target-base recheck, ff-only integration and target ancestry","result":"pass"},{"command":"stopped-runtime releases, successful-window closure verification and non-forced owned-worktree cleanup","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
