---
identity:
  kind: phase
  phaseID: rex-whichkey-17-unnecessary-parts-audit
  uuid: 4ede1be7-ff7d-4b13-ac87-ae37dae00b01
structure:
  order: 17
scope:
  requirements:
    - REQ-001
    - REQ-002
    - REQ-003
    - REQ-004
    - REQ-014
    - REQ-015
    - REQ-016
  documents:
    - ref: designs/rex-whichkey-cli-only-menu.md
      hash: 'sha256:23add062dac4326fe35e7e178a6a39f4c4d7631b4c8e164feb988b9aba07e21b'
lifecycle:
  state: complete
  sourceEventID: rex-whichkey-p17/summary/1
  changedAt: '2026-11-07T01:31:20.513Z'
title: 'Phase 17: Audit unnecessary parts (speed excluded)'
mtnCreateRetry:
  key: rex-whichkey-phase17-unnecessary-parts-audit-create-v1
  intent: f0cdc2f306d516065028c4585c20eee835719d5c54264e4a0a52992e63cfb9b6
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-06T23:49:04.081Z'
dateModified: '2026-11-07T01:31:21.150Z'
tags:
  - task
type: task
id: d7d1a0df-bf21-4f20-a5a0-f44859a6e155
projects:
  - '[[projects/rex-whichkey.md]]'
completedDate: '2026-11-06'
---
# Phase 17 — Audit unnecessary parts

## Authority and boundary

User request: "create a phase for a audit of unnessesary parts. we will discuss the speed afterwards".

This defines one audit-only phase. The deliverable is an evidence-backed inventory and cleanup recommendations, not authorization to remove dependencies/code, change behavior, or start implementation now. Startup speed, deferred Rex calls, goroutines, asynchronous menus, caching and performance/security-policy redesign are explicitly outside this phase and reserved for later discussion. No speed phase or performance task is created here.

Use the accepted CLI-only design in designs/rex-whichkey-cli-only-menu.md as the current product boundary: flat short-lived Go application, Go-readable data configuration, concrete Rex resources through the argument-array CLI boundary, minimal launch-only Lua, and deferred hooks/callbacks. Older callback-oriented project/PRD prose and test fixtures do not establish current production requirements. Preserve their provenance; do not migrate legacy workflow records or reopen unrelated phases.

## Baseline and known starting points

Phase16 is delivered and native-complete. Its summary distinguishes source review, actual public binary/menu/dynamic navigation/owned rename/dismissal/cleanup evidence, untested behavior and undeployed user configuration. Bind the future audit to an explicit Git revision and separately describe working-tree changes. At definition time HEAD is 5d15997; flake.nix has a separate uncommitted Darwin pre-commit/NuGet packaging fix from the dev-shell investigation. Preserve that change; it is not an approved cleanup or part of this planning delivery. The individual fixed pre-commit package built, but full dev-shell entry was not verified.

Discussion identified candidates, not an approved deletion list:

- Dev shell: Docker; the make-to-Task wrapper despite no current Makefile or documented make workflow; Python pre-commit framework/configuration alongside the active direct Task-based Git hook; gotools with no in-repository command consumer but unknown editor/personal usage.
- Shell entry side effects: automatic git pull, hook-path changes and unconditional copying of shared lint configurations. Distinguish unnecessary entry-time mutation from the required lint rules/configuration and quality gates themselves.
- Legacy Go wire/compatibility surfaces: Completion/capability types and constants, legacy DecodeConfig, launchOutput, readPopupResult/nativeTextSelection and nativeRenameMapping. Establish actual callers; do not confuse these with live strict JSON/Unicode helpers, retained resource validation or generic Go popup transport.
- Documentation: README.md and docs/acceptance.md retain pending-verification statements predating the completed Phase16 runtime evidence; historical sections and installed-document relative-link limitations also need clear classification.
- Historical/test-only content: archived Phase15 Lua fixtures and docs/experiments are intentionally non-production material. Establish test/packaging/provenance dependencies before proposing any disposition; exclusion from installation is not the same as safe deletion.

## Scope and deliverable

Create a reviewed audit report, expected at docs/unnecessary-parts-audit.md when execution is authorized, covering runtime code and dependencies, development tools and hooks, packaging/shipped assets, examples, documentation and retained test/history artifacts. This phase definition does not create that report or claim a comprehensive audit has already occurred.

For each item record: exact path/package/symbol, production/build/test/tool/editor/docs/packaging consumers, evidence and limitations, classification, reason to retain or consider removal, effects on existing assertions/quality gates/platforms/public contracts, risks and any validation or user choice still needed.

Use distinct classifications: required production; required build/quality tooling; optional developer convenience; historical/test-only; confirmed unreferenced/redundant; unresolved. No repository reference is a useful lead, not proof of non-use: account for build tags, reflection/serialization, embed directives, Nix exports, examples, external tooling/editor use and platform-specific paths. A filename containing runtime/legacy or a transitive dependency is not sufficient evidence by itself.

## Acceptance criteria

1. **Bounded, reproducible inventory.** Report the inspected revision, local differences and scope. Inventory actual dependency declarations, shell packages/hooks, active code, installed assets and historical fixtures, with evidence references rather than generic template-cleanup claims.
2. **Meaningful consumer tracing.** Classify items using actual caller/command/gate/test/module/install relationships. Separate confirmed findings from unknown editor/external usage and assumptions. Explain exclusions and limitations; do not claim a complete audit from a single textual search.
3. **Development tooling audit.** Address Docker, make-wrapper, pre-commit duplication and gotools explicitly. Identify tools actually used by Taskfile.yml and existing quality gates, including Go, Task, golangci-lint, go-ai-lint and shared configuration. Assess shellHook side effects without executing them or weakening/removing hook/test/lint standards.
4. **Legacy code and artifact audit.** Trace the discussed Go compatibility symbols and their historical tests. Identify live shared validators, CLI/resource/lifecycle helpers and minimal production Lua separately. Document fixture/experiment/package boundaries and preservation needs; historical or test-only does not automatically mean removable.
5. **Documentation audit.** Identify stale status, obsolete active instructions, duplication, historical material and installed-link limitations. Compare claims against recorded Phase16 source/runtime evidence; do not infer physical keybinding, package activation, platform coverage or broad test success. Recommend corrections without silently rewriting historic delivery results.
6. **Actionable disposition matrix.** For each removal/simplification candidate provide evidence, risk, affected consumers, required assertion/documentation adaptations and the smallest remaining validation/decision. Produce bounded follow-up cleanup groups for later approval; no approved removals, feature changes or performance benefit are implied by the audit.
7. **Explicit exclusions and preservation.** No startup-speed profiling/benchmarks, async/goroutine/deferred-call/cache design, executable-copy/hash/server-check optimization or safety-policy changes. No production/tooling/config changes, dependency/lock updates, test/fixture/assertion deletion or weakening, user config/profile edits, Rex reload/restart, deployment/push, cache/worktree/history deletion or global Nix cleanup.
8. **Honest reviewed completion.** Completion means the audit report and independently reviewed inventory/recommendations are delivered with exact revision/evidence/uncertainty, not that cleanup or performance work was implemented. Preserve accepted CLI-only behavior, tests and standards, and report unresolved items instead of treating them as accepted decisions.

## Expected inspection paths

flake.nix, flake.lock, .envrc, Taskfile.yml, .githooks/pre-commit, .pre-commit-config.yaml, go.mod/go.sum, nix/package.nix, nix/home-manager.nix, root Go files, rex/, cmd/, minimal production lua/, examples/, active docs/, tests/fixtures/phase15-lua and its tests/lua loaders. Read native .tasks records for requirements/evidence only; the coordinator remains their sole writer. Existing .direnv, .mtn-tmp, profiles, caches, failed-review resources and history are not deletion targets.

## Requirements and prerequisites

Complementary maintenance/audit traceability: REQ-001/002/003/004/014/015/016 under the current CLI-only design. The report evaluates necessity and consistency; it does not claim new functional implementation of those requirements. REQ-003's former Lua-loading mechanism remains superseded, executable REQ-004 portions and REQ-005 remain deferred.

Parent: projects/rex-whichkey.md; order17. Phase16 is the documented delivered baseline, not a new whole-phase barrier: this read-only audit needs no completion of older incomplete phases and has no native prerequisites. No speed milestone dependency is introduced. Future tasks should own audit/report scopes, not an unrestricted cleanup rewrite.

## Verification and execution status

Planning/phase creation only. No task graph, workers, report execution, removals or product changes are authorized by this request. A later execution instruction may authorize the audit and its bounded task decomposition, not automatic cleanup.

Use source/Git/native metadata inspection and a fresh independent source/report review. Do not run prohibited unit/race/TDD/Lua/parser/FIFO harnesses, benchmarks, experiments, live/remote trials, shellHook, package/profile activation or indirect test suites. Existing prior build/run evidence remains historical verification, not a new test pass. Source documentation checks and native record/graph/coverage validation must be accurately reported. Performance discussion remains postponed.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDAxOjMxOjE5Ljc0NloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxNy9zdW1tYXJ5LzEiLCJpbnRlbnQiOiI1NTc5NjA3ZDM4MjZlZjAzNmI2MWNjMDU0MWNmNTExYjZjYWQ5ZmMyMGU3Zjk2YWFiZWU2MzVmYmMxNjhjYTYxIn0 -->
### 2026-11-07T01:31:19.746Z · summary · unspecified

{"acceptedScope":"Baseline-bound unnecessary-parts audit only: dependency/tool/hook/package/live-vs-historical consumers, six classifications, stale-doc evidence, disposition/risk/remaining-decision matrix. Correct Darwin select/Linux poll fact independently re-reviewed; no removals or performance changes.","approvalSummary":["User audit phase and start-next-go-team/finish-all/commit-each-phase execution authority","Current exact approved one-task Phase17 plan and CLI-only design","Fresh independent same-family combined review2 APPROVED exact bcb0a934/tree8650a859 after REX-P17-T01-001; reviewed source integrated with no adaptation"],"attempt":1,"cleanup":{"eligibleWorktreesRemoved":["p17-planner-a01","p17-t01-builder-a01","p17-t01-builder-a02","p17-t01-tasks-a02","p17-t01-integration-a02"],"facade":"Official complete native/protocol succeeds, cleanup jq syntax defect persists; manually removed only clean delivered ancestor worktrees after stopped-consumer verification. Ordinary team resource cleanup, not audit-recommended removal.","retainedNegativeReviewer":{"reason":"Negative review recovery resource; deletion not authorized","window":"window:01a1b382-f6d0-72da-9342-a148e9e6103b","worktree":"p17-t01-tasks-a01"},"successfulWorkerWindowsConfirmedAbsent":true},"eventID":"rex-whichkey-p17/summary/1","files":["docs/unnecessary-parts-audit.md",".mtn-tmp/p17/build-result.json",".mtn-tmp/p17/review-result.json",".mtn-tmp/p17/fix-build-result.json",".mtn-tmp/p17/fix-review-result.json",".mtn-tmp/p17/integration.json"],"kind":"summary","limitations":["Audit-only, no cleanup/removals/runtime/performance/test/build/environment changes authorized by Phase17","Fresh same-family report/source review, not cross-family verification","Report binds9634246 source; later Phase18 changes require fresh tracing before any cleanup recommendation is executed","External/editor use, deployed state, full transitive/compiler reachability remain uncertain as report states","Existing flaky/full lint and actual Phase16 evidence are historical, not rerun by audit","User-config/fixture/assertion/tooling/production removals and recommendation acceptance not inferred; separate dirty flake.nix fix preserved"],"nextScope":"Phase18 independently underway under existing coordinator; audit recommendations are NOT authorized removals. Historical phases11/14/15 not reopened or falsely completed.","recordRef":"phases/rex-whichkey-17-unnecessary-parts-audit.md","records":["tasks/rex-whichkey-17-01-unnecessary-parts-audit.md"],"requirements":["REQ-001","REQ-002","REQ-003","REQ-004","REQ-014","REQ-015","REQ-016"],"schemaVersion":1,"specHash":"sha256:69c39657459baf2f3373a3a9b28b15a147033e2a680a5e34596777310165eb2f","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"62284533cd777d1501f292ca6b62f981bdd57474","targetTree":"ad81571de00d15ac0ab0bc2069e74cc64cff393d","verification":[{"artifact":".mtn-tmp/p17/integration-whitespace.log","command":"Combined git diff --check TARGET_BASE HEAD","result":"pass"},{"artifacts":[".mtn-tmp/p17/candidate-report-identity.log",".mtn-tmp/p17/candidate-preservation.log"],"command":"Frozen reviewed report equals candidate report; sole source change is report; existing startup source/metadata and separate dirty flake.nix preserved","result":"pass"},{"artifact":".mtn-tmp/p17/main-fast-forward.log","command":"Target-base recheck, verified candidate fast-forward and reviewed report head ancestor of main","result":"pass"},{"artifacts":[".mtn-tmp/p17/final-hierarchy.json",".mtn-tmp/p17/final-collection-validation.json"],"command":"Native nonempty1/1 complete hierarchy; seven requirements/eight criteria covered; current plan/design approvals and acyclic valid native collection","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
