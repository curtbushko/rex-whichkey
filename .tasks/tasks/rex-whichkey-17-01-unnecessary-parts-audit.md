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
  state: building
  sourceEventID: rex-whichkey-p17-t01-builder-a01/runtime/launched
  changedAt: '2026-11-07T00:46:04.313Z'
title: Document the evidence-backed unnecessary-parts audit
mtnCreateRetry:
  key: rex-whichkey/phase17/task01-unnecessary-parts-audit/create/1
  intent: 222dd430a01a228668aed88cf8f720334f58d8d73980c47d0fcbffdc0eda40d3
status: in-progress
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T00:43:47.063Z'
dateModified: '2026-11-07T00:46:04.968Z'
tags:
  - task
type: task
id: 7c33688d-8aa6-46bc-ab17-6840bec62296
projects:
  - '[[phases/rex-whichkey-17-unnecessary-parts-audit.md]]'
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

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDAwOjQ1OjUxLjQzOVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiZGlzcGF0Y2giLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMTctdDAxL2Rpc3BhdGNoLzEiLCJpbnRlbnQiOiI4OGY1ZGM2Nzk2NjQ1OWJhNWU0Y2RlYjg0YjkwZmUzYjAzZDViMjhmODkzMmI4OWViN2UwYzQ5MjE5Y2FlMzVjIn0 -->
### 2026-11-07T00:45:51.439Z · dispatch · unspecified

{"allowedCommands":["read-only source/Git/native-record inspection","git diff --check","git -c core.hooksPath=/dev/null commit -- docs/unnecessary-parts-audit.md"],"allowedScope":{"interfaces":[],"paths":["docs/unnecessary-parts-audit.md"],"resources":[],"sharedPaths":[]},"attempt":1,"baseCommit":"cef25463b4dd58c99228d8f9b857ed3554a3164b","branch":"work/p17-t01-builder-a01","constraints":["No tests/build/run/benchmarks/shellHook or source/config/removal changes","Only existing coordinator publishes/integrates; one docs builder, no heavy gate","Existing unrelated parent flake.nix preserved outside worktree/commits"],"eventID":"rex-whichkey-p17-t01/dispatch/1","kind":"dispatch","readiness":"Native task/phase readiness, exact design approval and phase task plan_approved guard passed; legacy root ready-phase metadata facade unsupported and not migrated.","recordRef":"tasks/rex-whichkey-17-01-unnecessary-parts-audit.md","runID":"rex-whichkey-p17-t01-builder-a01","schemaVersion":1,"specHash":"sha256:efc85700f7bc5613ecb7c423eb19628f095c5669dec4a5e30b203c90829d74fb","targetBranch":"main","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p17-t01-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDAwOjQ2OjAwLjE5NloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxNy10MDEtYnVpbGRlci1hMDEvcnVudGltZS9yZXNlcnZlZCIsImludGVudCI6IjI5N2U4MjA0ZGYzMDU3NDAwNTVlOTFhMmNlNTg0MzZlODNjNzRmM2I0ZjE5ODcyNmY4NjY1OTNhYTVmZTYwY2IifQ -->
### 2026-11-07T00:46:00.196Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-15-code-owned-runtime.md","runID":"rex-whichkey-p15-coordinator-a01"},"dispatchEventID":"rex-whichkey-p17-t01/dispatch/1","eventID":"rex-whichkey-p17-t01-builder-a01/runtime/reserved","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"75ff9ad4-a558-4295-8b20-855dafe77a8f"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-17-unnecessary-parts-audit.md","piSessionID":"rex-whichkey-p17-t01-builder-a01","piSessionName":"p17-t01-builder-a01","recordRef":"tasks/rex-whichkey-17-01-unnecessary-parts-audit.md","rexSessionName":"rex-whichkey-main","rexWindowName":"p17-t01-builder-a01","role":"builder","runID":"rex-whichkey-p17-t01-builder-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"mtn-tasks","specHash":"sha256:efc85700f7bc5613ecb7c423eb19628f095c5669dec4a5e30b203c90829d74fb","stage":"reserved","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p17-t01-builder-a01","windowName":"p17-t01-builder-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p17-t01-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDAwOjQ2OjAzLjU0NFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxNy10MDEtYnVpbGRlci1hMDEvcnVudGltZS9sYXVuY2hlZCIsImludGVudCI6IjIzNDA0Y2Y5NThiYjE5MDk0YjcwNTk0Yzk1YjU1OTZhZDAxYmIxNWRlMGFlMzYwNjNiODMwYWQ3YTU5M2QxODAifQ -->
### 2026-11-07T00:46:03.544Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-15-code-owned-runtime.md","runID":"rex-whichkey-p15-coordinator-a01"},"dispatchEventID":"rex-whichkey-p17-t01/dispatch/1","eventID":"rex-whichkey-p17-t01-builder-a01/runtime/launched","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"75ff9ad4-a558-4295-8b20-855dafe77a8f"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-17-unnecessary-parts-audit.md","piSessionID":"rex-whichkey-p17-t01-builder-a01","piSessionName":"p17-t01-builder-a01","provider":{"displayName":"p17-t01-builder-a01","name":"rex","resources":{"block":"block:01a1b378-3375-75b8-8687-5ab015ecee6f","endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","session":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","window":"window:01a1b378-3376-7829-93ba-61099a3cd6a9"}},"recordRef":"tasks/rex-whichkey-17-01-unnecessary-parts-audit.md","rexBlockID":"block:01a1b378-3375-75b8-8687-5ab015ecee6f","rexSessionID":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","rexSessionName":"rex-whichkey-main","rexWindowID":"window:01a1b378-3376-7829-93ba-61099a3cd6a9","rexWindowName":"p17-t01-builder-a01","role":"builder","runID":"rex-whichkey-p17-t01-builder-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"mtn-tasks","specHash":"sha256:efc85700f7bc5613ecb7c423eb19628f095c5669dec4a5e30b203c90829d74fb","stage":"launched","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p17-t01-builder-a01","windowName":"p17-t01-builder-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p17-t01-builder-a01"}
<!-- mtn:end-entry:v1 -->
