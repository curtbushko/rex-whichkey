---
identity:
  kind: task
  taskID: rex-whichkey-17-01-unnecessary-parts-audit
  uuid: fe036ddb-e877-458c-ac17-082828b70a3a
structure:
  order: 1
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
implementation:
  language: docs
  ownership:
    paths:
      - docs/unnecessary-parts-audit.md
    sharedPaths: []
    resources: []
    interfaces: []
  gates:
    - ID: documentation-whitespace
      command: git diff --check
      required: true
lifecycle:
  state: complete
  sourceEventID: rex-whichkey-p17-t01/summary/2
  changedAt: '2026-11-07T01:31:08.017Z'
title: Document the evidence-backed unnecessary-parts audit
mtnCreateRetry:
  key: rex-whichkey/phase17/task01-unnecessary-parts-audit/create/1
  intent: 222dd430a01a228668aed88cf8f720334f58d8d73980c47d0fcbffdc0eda40d3
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T00:43:47.063Z'
dateModified: '2026-11-07T01:31:08.656Z'
tags:
  - task
type: task
id: 7c33688d-8aa6-46bc-ab17-6840bec62296
projects:
  - '[[phases/rex-whichkey-17-unnecessary-parts-audit.md]]'
completedDate: '2026-11-06'
---
Deliver only docs/unnecessary-parts-audit.md, implementing Phase17 criteria 1-8 as one coherent report. This is maintenance/audit traceability for the seven scoped REQ IDs, not new functional implementation. The accepted CLI-only design governs: flat short-lived Go/data configuration, concrete argument-array Rex CLI operations and launch-only Lua; hooks/callbacks and executable REQ-004 portions remain deferred, REQ-003 Lua loading is superseded, and REQ-005 is not claimed.

Acceptance 1: Bind the inventory to baseline commit 9634246b7e4edbc8bed59210466cac492ddf2f0f and its tree; record actual inspected checkout/revision and separately describe local/parent differences. Inspect dependency declarations/locks, shell tools/hooks, active root Go/rex/cmd, production Lua, package/HM exports, installed assets/examples/docs, retained tests/fixtures and nested historical experiment modules. Explain bounded exclusions. Report the parent-only Darwin pre-commit/NuGet flake.nix diff separately without incorporating or treating it as approved cleanup.

Acceptance 2: For each item give exact path/package/symbol and source-line or revision-stable references; trace production, build, test, tool, editor/external, docs and packaging consumers or explicitly state unknown/not applicable. Connect declarations to imports/callers, Task commands/hooks, fixture loaders and source/install filters. Consider build tags/platform files, reflection/serialization, embed directives, Nix exports, transitive dependencies and examples. A name or zero textual hits is not proof of non-use. Use exactly distinct classifications: required production; required build/quality tooling; optional developer convenience; historical/test-only; confirmed unreferenced/redundant; unresolved.

Acceptance 3: Explicitly audit Docker, make-to-Task wrapper, Python pre-commit versus direct Task hook, and gotools. Trace Go, Task, golangci-lint, go-ai-lint and shared configuration through existing gates. Describe shellHook git pull, hook-path mutation and unconditional config copies from source only. Distinguish questionable entry-time side effects from required lint rules/quality standards. Personal/editor use remains unknown unless evidenced; do not assert non-use.

Acceptance 4: Trace Completion/capability types/constants, legacy DecodeConfig, launchOutput, readPopupResult/nativeTextSelection and nativeRenameMapping to real callers and historical assertions. Separately identify live strict JSON/Unicode helpers, resource validation, generic Go popup transport, CLI/lifecycle helpers and minimal production Lua. Follow gopher-lua/test dependencies through lua_context_test.go and tests/lua loaders into tests/fixtures/phase15-lua. Inspect docs/experiments provenance and nested module/package boundaries. Test-only or excluded-from-installation is not safe-to-delete evidence.

Acceptance 5: Inventory stale status, obsolete active instructions, duplication, historical sections and installed-relative-link limitations, especially README.md and docs/acceptance.md. Compare claims against Phase16 delivered summary and available referenced source/runtime evidence, distinguishing observed public menu/dynamic navigation/owned rename/dismissal/cleanup from untested actions, physical keybinding, undeployed profiles/config and Linux/remote coverage. Preserve the failed older phases and historical prose provenance. Recommend status/doc corrections only inside this report; do not edit other docs or claim tests/lint globally passed.

Acceptance 6: Supply a disposition matrix with each candidate, evidence, classification, retention/removal rationale, affected consumers/platforms/public contracts, risks, required assertion/documentation adaptations and smallest remaining validation or user choice. Propose bounded follow-up cleanup groups for later approval. No removal, adaptation or performance benefit is approved or implemented by this report.

Acceptance 7: Preserve all production/tooling/config/dependency/lock files, tests/assertions/fixtures, hooks, lint standards, user state and history. No code/tool/dependency/config/doc-status fixes or removals; no startup profiling/benchmarks, asynchronous/goroutine/deferred-call/cache or executable-copy/hash/server-check optimization, performance or safety-policy design. No environment entry, shellHook, build, normal run, harness, experiment, activation, config/reload/restart/deployment/push or cache/worktree/history/global cleanup.

Acceptance 8 and verification: Give explicit 1-8 criterion and seven-REQ traceability, source/evidence limitations, confirmed findings versus uncertainty, and a concise verification ledger. Use read-only source/Git/local-document checks plus required git diff --check only; do not invoke Task, hooks, Nix or tests indirectly. Before completion, a separate fresh reviewer must inspect the frozen report and referenced source independently for scope, classification, consumer tracing, evidence accuracy and preserved CLI-only behavior. Report reviewer findings and exact reviewed revision through the coordinator; the author must not claim independent review. Coordinator alone publishes native records, integrates and handles the requested phase commit. Only the report is writable; .tasks and parent evidence are read-only. Unavailable external/editor/platform facts remain explicit unresolved items rather than invented approval or a reason to expand scope.
<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDAxOjMxOjA3LjIxNVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxNy10MDEvc3VtbWFyeS8yIiwiaW50ZW50IjoiZThiYTg4ZTEwNTMyZDEyNWJkNGVkZjBlNTU0MWRiYTYxNTNiOTM5N2I1NjBkZjI0MDVkZDA0YTMxMzk5ZGQyNSJ9 -->
### 2026-11-07T01:31:07.215Z · summary · unspecified

{"acceptedScope":"Baseline-bound unnecessary-parts audit only: dependency/tool/hook/package/live-vs-historical consumers, six classifications, stale-doc evidence, disposition/risk/remaining-decision matrix. Correct Darwin select/Linux poll fact independently re-reviewed; no removals or performance changes.","approvalSummary":["User audit phase and start-next-go-team/finish-all/commit-each-phase execution authority","Current exact approved one-task Phase17 plan and CLI-only design","Fresh independent same-family combined review2 APPROVED exact bcb0a934/tree8650a859 after REX-P17-T01-001; reviewed source integrated with no adaptation"],"attempt":2,"cleanup":{"eligibleWorktreesRemoved":["p17-planner-a01","p17-t01-builder-a01","p17-t01-builder-a02","p17-t01-tasks-a02","p17-t01-integration-a02"],"facade":"Official complete native/protocol succeeds, cleanup jq syntax defect persists; manually removed only clean delivered ancestor worktrees after stopped-consumer verification. Ordinary team resource cleanup, not audit-recommended removal.","retainedNegativeReviewer":{"reason":"Negative review recovery resource; deletion not authorized","window":"window:01a1b382-f6d0-72da-9342-a148e9e6103b","worktree":"p17-t01-tasks-a01"},"successfulWorkerWindowsConfirmedAbsent":true},"eventID":"rex-whichkey-p17-t01/summary/2","files":["docs/unnecessary-parts-audit.md",".mtn-tmp/p17/build-result.json",".mtn-tmp/p17/review-result.json",".mtn-tmp/p17/fix-build-result.json",".mtn-tmp/p17/fix-review-result.json",".mtn-tmp/p17/integration.json"],"kind":"summary","limitations":["Audit-only, no cleanup/removals/runtime/performance/test/build/environment changes authorized by Phase17","Fresh same-family report/source review, not cross-family verification","Report binds9634246 source; later Phase18 changes require fresh tracing before any cleanup recommendation is executed","External/editor use, deployed state, full transitive/compiler reachability remain uncertain as report states","Existing flaky/full lint and actual Phase16 evidence are historical, not rerun by audit","User-config/fixture/assertion/tooling/production removals and recommendation acceptance not inferred; separate dirty flake.nix fix preserved"],"recordRef":"tasks/rex-whichkey-17-01-unnecessary-parts-audit.md","requirements":["REQ-001","REQ-002","REQ-003","REQ-004","REQ-014","REQ-015","REQ-016"],"schemaVersion":1,"specHash":"sha256:efc85700f7bc5613ecb7c423eb19628f095c5669dec4a5e30b203c90829d74fb","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"62284533cd777d1501f292ca6b62f981bdd57474","targetTree":"ad81571de00d15ac0ab0bc2069e74cc64cff393d","verification":[{"artifact":".mtn-tmp/p17/integration-whitespace.log","command":"Combined git diff --check TARGET_BASE HEAD","result":"pass"},{"artifacts":[".mtn-tmp/p17/candidate-report-identity.log",".mtn-tmp/p17/candidate-preservation.log"],"command":"Frozen reviewed report equals candidate report; sole source change is report; existing startup source/metadata and separate dirty flake.nix preserved","result":"pass"},{"artifact":".mtn-tmp/p17/main-fast-forward.log","command":"Target-base recheck, verified candidate fast-forward and reviewed report head ancestor of main","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
