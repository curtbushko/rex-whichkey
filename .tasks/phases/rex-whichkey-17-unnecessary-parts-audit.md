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
  state: incomplete
  sourceEventID: rex-whichkey-p17/planning-validation/1
  changedAt: '2026-11-06T23:51:12.489Z'
title: 'Phase 17: Audit unnecessary parts (speed excluded)'
mtnCreateRetry:
  key: rex-whichkey-phase17-unnecessary-parts-audit-create-v1
  intent: f0cdc2f306d516065028c4585c20eee835719d5c54264e4a0a52992e63cfb9b6
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-06T23:49:04.081Z'
dateModified: '2026-11-06T23:51:13.517Z'
tags:
  - task
type: task
id: d7d1a0df-bf21-4f20-a5a0-f44859a6e155
projects:
  - '[[projects/rex-whichkey.md]]'
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

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA2VDIzOjUwOjAzLjU5N1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMTcvZGVmaW5pdGlvbi1hcHByb3ZhbC8xIiwiaW50ZW50IjoiZjFkOWRhNWIyNmEyZTc5ZTYxZjEyZDExMjA4MGRjMjc2YmQyZGRiOTllYjM5OGM3ZGY0Nzc1NjQ3ZjY2ZjgxNyJ9 -->
### 2026-11-06T23:50:03.597Z · approval · unspecified

{"approvalScope":"One audit-only phase definition/planning: evidence-backed inventory/classification and cleanup recommendations. No audit execution/task dispatch, automatic removals, behavior/standards changes or startup-speed/async/deferred-call work authorized.","approvedBy":"user","attempt":1,"decision":"approved","decisionEvidence":"User explicitly requests: create a phase for a audit of unnessesary parts. we will discuss the speed afterwards","eventID":"rex-whichkey-p17/definition-approval/1","kind":"approval","limitations":["Prior discussion is starting evidence, not an approved deletion list or completed comprehensive audit.","Existing uncommitted flake.nix Darwin pre-commit fix is preserved separately, not adopted/staged by this planning action."],"recordRef":"phases/rex-whichkey-17-unnecessary-parts-audit.md","requirements":["REQ-001","REQ-002","REQ-003","REQ-004","REQ-014","REQ-015","REQ-016"],"schemaVersion":1,"specHash":"sha256:69c39657459baf2f3373a3a9b28b15a147033e2a680a5e34596777310165eb2f"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA2VDIzOjUxOjExLjI3N1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxNy9wbGFubmluZy12YWxpZGF0aW9uLzEiLCJpbnRlbnQiOiJhOGRlNjI3ZWM1NDVkZjI0ZDM3MjhmNzkxYjUzZGE2ZjUxNTlmZTg0NzQ1YWRiODlkMzVlMTRmYzY0ZGQwMzVhIn0 -->
### 2026-11-06T23:51:11.277Z · historical · unspecified

{"artifacts":[".mtn-tmp/audit-phase/final-phase-snapshot.json",".mtn-tmp/audit-phase/phase-hierarchy.json",".mtn-tmp/audit-phase/phase-dependencies.json",".mtn-tmp/audit-phase/order-validation.json",".mtn-tmp/audit-phase/collection-validation.json",".mtn-tmp/audit-phase/project-plan-native-readback.json"],"attempt":1,"content":"Created one open audit-only Phase17 under existing project, unique order17, zero prerequisites/tasks/workers. Current approved CLI-only governing design and phase-definition authority read back; seven maintenance-trace requirement IDs and eight audit criteria covered by additive planning evidence. Native collection valid and graph acyclic. Legacy root project facade reports unknown workflow kind after plan persistence; exact native plan equality verified without duplicate append, migration or invented global readiness pass. No audit report/execution/removals/performance work started, no new sessions/coordinators, no tests/build/runtime/config actions. Existing uncommitted flake.nix Darwin fix preserved outside metadata commit.","eventID":"rex-whichkey-p17/planning-validation/1","gateEligible":false,"kind":"historical","provenance":"Coordinator native mtn creation/membership/hierarchy/dependency/collection checks, current accepted CLI-only design fingerprint/approval, exact additive project plan readback and sibling order validation.","recordRef":"phases/rex-whichkey-17-unnecessary-parts-audit.md","schemaVersion":1,"specHash":"sha256:69c39657459baf2f3373a3a9b28b15a147033e2a680a5e34596777310165eb2f"}
<!-- mtn:end-entry:v1 -->
