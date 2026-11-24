---
identity:
  kind: task
  taskID: rex-whichkey-21-01-lint-toolchain
  uuid: c2f6a044-4062-45eb-a2e2-db9468bd9172
structure:
  order: 1
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
implementation:
  language: nix
  ownership:
    paths:
      - flake.nix
      - flake.lock
      - nix/lint-toolchain.nix
      - nix/lint-toolchain-check.nix
    sharedPaths: []
    resources:
      - lint-toolchain-heavy-gate
    interfaces:
      - default-devshell-lint-toolchain
  reviewPlan:
    languages:
      - language: nix
        paths:
          - flake.nix
          - flake.lock
          - nix/lint-toolchain.nix
          - nix/lint-toolchain-check.nix
      - language: bash
        paths:
          - flake.nix
          - nix/lint-toolchain.nix
          - nix/lint-toolchain-check.nix
    specialists:
      architecture:
        required: false
        reason: >-
          Narrow developer-tool package compatibility repair; no
          application/runtime/interface architecture change. Maintenance reviews
          build-environment consistency.
      security:
        required: false
        reason: >-
          Pinned trusted development tooling; no coding
          authorization/secrets/trust-boundary change. Package provenance
          belongs to maintenance.
      usability:
        required: false
        reason: >-
          Restores existing task lint workflow without changing CLI/help/default
          interactions or shellHook semantics.
      maintenance:
        required: true
        reason: >-
          Nix dev shell, tool/compiler/source pins, regression checks and
          default quality workflow change.
  gates:
    - ID: nix-evaluation
      command: nix flake check --no-build --no-write-lock-file
      required: true
    - ID: lint-toolchain-regression
      command: >-
        nix build --no-link --max-jobs 1 --cores 2
        .#checks.aarch64-darwin.lint-toolchain
      required: true
    - ID: build
      command: task build
      required: true
    - ID: tests
      command: task test
      required: true
    - ID: full-lint
      command: task lint
      required: true
    - ID: whitespace
      command: git diff --check
      required: true
lifecycle:
  state: complete
  sourceEventID: rex-whichkey-p21-t01/release/dispatch/2
  changedAt: '2026-11-24T15:12:16.011Z'
title: Align dev-shell linters with pinned Go and add regression check
mtnCreateRetry:
  key: rex-whichkey-p21-task01-v1
  intent: aa2ea058eff420bf6550ddde4f0e45c11d05e4961411135c4c94361bf57af063
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-24T13:50:16.375Z'
dateModified: '2026-11-24T15:13:45.794Z'
tags:
  - task
type: task
id: 14964fae-d2af-4766-b50c-9c4eb6c09330
projects:
  - '[[phases/rex-whichkey-21-lint-toolchain.md]]'
completedDate: '2026-11-24'
---
# Align declared lint tools and add a compatibility regression

## Scope and authority
Implement user-approved flake.nix maintenance only, covering REQ-ENV-001..003. Read the governing lint-toolchain design and current quality-gate correction through the parent collection. Existing phase20 script behavior is not reopened. The user approved the repair; routine package/builder/hash choices and tests are delegated within the design constraints, not separately user-approved details.

## Deliverable
Research actual locked package/compiler definitions and both linter minimum supported Go versions. Make the default devShell select reproducibly pinned compatible golangci-lint and go-ai-lint with application Go1.26.8 unchanged. Prefer known working golangci-lint2.15.0 and minimal focused overrides; a different necessary compatible version must satisfy unchanged config/full gates. No hardcoded existing store output or .mtn-tmp tool in shipped code. Preserve unrelated pins and app/package semantics; no broad flake update. flake.lock may change only for narrowly necessary lint-only inputs, with exact impact justified.

Own a small checks.<system>.lint-toolchain regression with the same selected Go/lint packages as the devShell, exercising a Go1.26.8-targeted module and both tools. Minimal embedded Go fixture is incidental; do not add production Go changes. Add meaningful RED/GREEN evidence detecting original incompatibility, not a version-print-only assertion or an unconsumed alternate package.

## Ownership/boundaries
Only flake.nix, narrowly scoped flake.lock, optional nix/lint-toolchain.nix and nix/lint-toolchain-check.nix. Keep application source, go.mod/sum, Taskfile, hooks, lint rule files, user profiles/config and existing shellHook commands unchanged. Do not run shellHook or enter a nested/profile Nix environment; it auto-pulls Git and copies configs. Use supported environment introspection without executing hooks, then gates from that declared environment. No temporary compatible-tool PATH may count as proof of the repair.

## Verification
Evaluate preserved four supported systems and host devShell/check outputs without broad lock writes. Build host checks.aarch64-darwin.lint-toolchain with max-jobs1/cores2, then actual task build, task test (unit+race), full task lint from the generated default environment. Inspect command resolution/build provenance and ensure selected tools are the newly declared derivations. Confirm hook/config/input/application invariants and Git whitespace. Required source commits use existing hooks under the new declared environment; no bypass. Reuse exact existing protected lint config files in worktree without running copy hooks. No install, activation, profile changes, push, remote/prod trials, GC or unrelated cleanup.

## Review and delivery
Separate Nix code/test, Bash code/test (embedded check logic) and maintenance reviewers, fresh contexts/exact frozen source tuple. No architecture/security/usability scope is triggered by this narrow restoration of existing commands/package compatibility without runtime/trust/UI changes; coordinator reassesses actual diff. Builder never writes mtn or integrates; coordinator owns publication, exact-tree speculative checks, main integration and owned stopped-resource cleanup. Return actual revision-bound build/evidence with provenance and limitations; do not invent package compatibility or cross-platform execution.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDE1OjEzOjQ1Ljc4MFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMS10MDEvc3VtbWFyeS8yIiwiaW50ZW50IjoiYjg5MGVmZWY4OWExOWNjNmE2ODc1YmVhNjI0NWI5YWVhMWY5ZThiNzcyOThjOTAxMGNiYmY3ZmM2OGZjYmYwNyJ9 -->
### 2026-11-24T15:13:45.780Z · summary · unspecified

{"acceptedScope":"Fix default Nix lint-toolchain compatibility reproducibly: source-pinned golangci-lint2.15.0 built with the still-pinned Go1.26.8, shared exact Go/lint derivations and a deterministic real-analyzer regression check. GoAI1.0/compiler, app Go/deps, lock, rules/configs/hooks and shellHook behavior preserved.","approvalSummary":["User explicitly approved yes, fix flake.nix after reported old-linter failure and recommendation to align reproducible compatible tools","Coordinator bounded phase/task plan under that authority; no claim of separate user approval of routine details","Five distinct current-build Nix/Bash code/test and maintenance keys APPROVED on exact frozen head/tree"],"attempt":2,"cleanup":{"branchesAndTransportPreserved":true,"facadeCompatibility":"Reused previously validated private cleanup release-guard syntax correction; installed tools/repository implementation untouched","integrationWorktreeRemoved":true,"retainedRecoveryWindows":[{"reason":"Stopped negative-review window retained by policy","rexWindowID":"window:01a20dea-1087-724d-876d-319a5bf56d12","runID":"rex-whichkey-p21-t01-nix-test-reviewer-a01"}],"successfulWorkerWindowsClosed":true,"taskWorktreesRemoved":true,"unlaunchedPreparedWorktreeRemoved":true},"eventID":"rex-whichkey-p21-t01/summary/2","files":["flake.nix","nix/lint-toolchain-check.nix",".mtn-tmp/p21/integration-accepted-a02.json",".mtn-tmp/p21/main-gates/",".mtn-tmp/p21/integration-a02/",".mtn-tmp/p21/builder-a02-evidence/",".mtn-tmp/p21/review-nix-code-a02-evidence/",".mtn-tmp/p21/review-nix-test-a02-evidence/",".mtn-tmp/p21/review-bash-code-a02-evidence/",".mtn-tmp/p21/review-bash-test-a02-evidence/",".mtn-tmp/p21/review-maintenance-a02-evidence/",".mtn-tmp/p21/cleanup/"],"kind":"summary","limitations":["Actual runtime/build verification on aarch64-darwin; other three supported systems evaluated only","Independent fresh same-family model contexts, not cross-family review","User shell/profile/config not activated or reloaded; caller must refresh its development environment to select declared tools","One stopped negative-review window retained by recovery policy; no delivery blocker"],"recordRef":"tasks/rex-whichkey-21-01-lint-toolchain.md","requirements":["REQ-ENV-001","REQ-ENV-002","REQ-ENV-003"],"reviewEventIDs":["rex-whichkey-p21-t01/review/nix-code/2","rex-whichkey-p21-t01/review/nix-test/2","rex-whichkey-p21-t01/review/bash-code/2","rex-whichkey-p21-t01/review/bash-test/2","rex-whichkey-p21-t01/review/maintenance/2"],"schemaVersion":1,"specHash":"sha256:75444801add9039e7f937bbe0f5387991287688a2594f670b017e65bec4ce80b","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"699b53e05d966c3f9e1326bbcf6a9e3e201598ba","targetTree":"95f434d0b08e07fa4146d0df2ad4b10e33f840eb","unresolvedBlockers":[],"verification":[{"command":"nix flake check --all-systems --no-build --no-write-lock-file","result":"pass"},{"command":"host lint-toolchain check, repeated actual --rebuild, old Go1.25-linter rejection and both real analyzer diagnostics","result":"pass"},{"command":"fresh default declared environment command/compiler provenance and unchanged protected input/Go/checksum/shellHook assertions","result":"pass"},{"command":"task build on combined candidate and final main","result":"pass"},{"command":"task test unit and race on combined candidate and final main","result":"pass"},{"command":"FULL unchanged task lint on combined candidate and final main","result":"pass"},{"command":"exact reviewed source tree, target-base recheck, ff-only integration and target ancestry","result":"pass"},{"command":"stopped-runtime releases, successful-window closure verification and non-forced owned-worktree cleanup","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
