---
identity:
  kind: task
  taskID: rex-whichkey-19-05-test-repair
  uuid: f5c572c0-934c-4ff4-a163-34cd2cf5dcac
structure:
  order: 5
scope:
  requirements:
    - REQ-001
    - REQ-002
    - REQ-003
    - REQ-004
    - REQ-006
    - REQ-008
    - REQ-009
    - REQ-010
    - REQ-011
    - REQ-012
    - REQ-013
    - REQ-014
    - REQ-015
    - REQ-016
    - REQ-017
    - REQ-025
    - REQ-028
    - REQ-029
  documents:
    - ref: designs/rex-whichkey-lua-float-launch.md
      hash: 'sha256:f7e38d49287de319753732b09fe472a40dfadcef81e5245e2deac208e968ab70'
    - ref: designs/rex-whichkey-quality-gates.md
      hash: 'sha256:f3120762f7075c7c767e5f39552ba9c2b38b2e213d33d2472d07f0bc0c1cbdb2'
implementation:
  language: go
  ownership:
    paths:
      - direct_menu_test.go
      - lua_context_test.go
      - tests/fixtures/legacy-context-checkpoint
      - tests/lua
    sharedPaths: []
    resources: []
    interfaces:
      - P19-QUALITY-GATES
  gates:
    - ID: whitespace
      command: git diff --check
      required: true
    - ID: build
      command: task build
      required: true
    - ID: tests
      command: task test
      required: true
    - ID: full-lint
      command: task lint
      required: false
lifecycle:
  state: complete
  sourceEventID: rex-whichkey-p19-t05/summary/2
  changedAt: '2026-11-07T16:30:40.898Z'
title: Repair faithful archived fixture test setup
mtnCreateRetry:
  key: rex-whichkey-19-05-test-repair-v1
  intent: aba3cde7720dd26e51fe585bb2f5d75139d1ebbba8ef10be7d96727ff344a877
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T16:00:04.596Z'
dateModified: '2026-11-07T16:30:41.530Z'
tags:
  - task
type: task
id: 8faa7318-b09c-437e-894a-dbc449443a9c
projects:
  - '[[phases/rex-whichkey-19-lua-float-launch.md]]'
completedDate: '2026-11-07'
---
Fix the four actual failing automated test groups shown in .mtn-tmp/p19/post-delivery-gates/unit.log and race.log: configured rename dispatch, CLI delegation, launch-origin diagnostics and action contexts. The user's explicit always-tests-and-lint/fix-the-problems directive supersedes earlier test prohibition. Read current supplementary quality design and original phase architecture. Existing archived Phase15 Lua helpers are historical test targets, NOT production; correct stale fixture/harness setup and module/source loading faithfully, preserve every assertion/scenario/security guarantee and negative path. No dropping/skipping cases, weakening guards, fabricating success, dummy compatibility objects or restoring old callbacks/controller to live source. Trace root causes and actual documented fixture interfaces before editing. Minimal genuine test-support changes, add current production regression coverage if genuinely required by a fix. No public behavior change.

First reproduce actual failing tests as RED, then fix; required normal task build and task test (unit+race) GREEN. Run full task lint after changes and record current diagnostics honestly; Task06 owns the existing lint debt, not a final exemption. No new lint failures introduced by this producer. Fresh independent spec/quality review with actual tests/build/lint runs, coordinator combined tests before target advance. User now authorizes existing automated unit/race/Lua/parser/FIFO sandbox fixtures/mocks/TDD. No live leaf Rex mutations, profile/config/install/reload/remote/deploy/push/benchmarks/global cleanup, new deps/pins, disabling lint rules or assertion weakening. Preserve separate dirty flake.nix.

Literal ownership listed includes Go Lua-test drivers and tests/lua harnesses. Archived tests/fixtures bytes stay unchanged unless concrete evidence proves a genuinely required non-weakening correction: ask coordinator for exact file scope first, not user for routine details. Go production and launcher behavior unchanged.14CPU/6GiB shared/Go5GiB, pinned existing toolchain. Commit own source hooksPath=/dev/null until full gates are repaired. Native build payload actual pass/fail, source/base/head/tree/spec/run IDs, commands/artifacts and remaining lint counts; tests and lint now MUST execute. Coordinator only .tasks/integration/real app. No phase completion from compile, optional lint or fixtures alone.

## Authorized exact historical fixture target

Coordinator traced git history: all four existing Lua test harnesses originated in unreviewed preservation checkpoint30a4439302ade4e08487d176ba6823206ddc326d; its exact Lua source supplies execution.helper/rex_bin_dir, field-specific origin diagnostics and matching registry issue/bind APIs. Current phase15-lua archive came from a later incompatible API. Add separate tests/fixtures/legacy-context-checkpoint containing the complete byte-identical Lua implementation tree from30a, with manifest source commit/Git blob identities and explicit unreviewed-historical-only label. Redirect ONLY these historical harness loaders to their matching authentic implementation; retain all42 scenarios/assertions/security/side-effect counts. Existing phase15-lua archive stays untouched. No copied code becomes production, no invented adapter or retroactive delivery claim. If exact original snapshot does not match, inspect further authentic history, do not weaken tests. Staging/loader correction and accurate fixture context remain allowed. This resolves the stopped no-change builder blocker under existing user fix authority, not another approval request.


<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDE2OjMwOjQwLjA4M1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDUvc3VtbWFyeS8yIiwiaW50ZW50IjoiNmI1MDU5MWQ5Njk3MWFhYTRmZThmZGRlZTQ5ZDlhNTk0YmExYWEzNzYyOTA1MDBlODQ0OWFmNmI1OWNjYzMyMiJ9 -->
### 2026-11-07T16:30:40.083Z · summary · unspecified

{"acceptedScope":"Faithful historical Lua test repair: complete exact original unreviewed30a Lua snapshot in separate test-only fixture+manifest, three matching historical harness loader changes only, all42 original assertions/scenarios/guards unchanged. Later archive/production untouched. Full unit and race suites now green, remaining full-lint debt assigned06.","approvalSummary":["Explicit user always-tests/full-lint/fix-problems supersedes prior test prohibition","Precise original authentic fixture target authorized after genuine API/archive mismatch, no assertion weakening","Fresh independent exact-head combined review2 plus coordinator combined build/tasktest(unit+race)green; full tasklint executed and reported53, not final acceptance"],"attempt":2,"cleanup":{"blockedAttempt1RecoveryRetained":true,"officialFacade":"Native delivery verified despite existing helper cleanup jq defect, exact eligible resources removed manually","successfulStoppedWorktreesRemoved":true},"diagnosticValidation":[{"artifact":".mtn-tmp/p19/quality-fix/task5-combined-lint.log","command":"FULL task lint executed","detail":"Known lint debt explicitly owned by dependent Task06, not final phase exception or green claim","exitCode":201,"issues":53,"required":false,"result":"fail"}],"eventID":"rex-whichkey-p19-t05/summary/2","files":["tests/fixtures/legacy-context-checkpoint","tests/lua/action_context.lua","tests/lua/cli_delegation.lua","tests/lua/configured_rename.lua",".mtn-tmp/p19/quality-fix/task5-retry-build-result.json",".mtn-tmp/p19/quality-fix/task5-review-result.json",".mtn-tmp/p19/quality-fix/task5-combined-tests.log",".mtn-tmp/p19/quality-fix/task5-builder-evidence",".mtn-tmp/p19/quality-fix/task5-reviewer-evidence"],"kind":"summary","limitations":["Historical matching original30a source is test-only, unreviewed original checkpoint not retroactively delivery; all42 scenarios/assertions preserved","Current source full-lint53/AI2 pending Task06; no current final phase quality acceptance","Fresh same-family independent test/source review, not cross-family; no production/profile/config/activation/deployment change"],"recordRef":"tasks/rex-whichkey-19-05-test-repair.md","requirements":["REQ-001","REQ-002","REQ-003","REQ-004","REQ-006","REQ-008","REQ-009","REQ-010","REQ-011","REQ-012","REQ-013","REQ-014","REQ-015","REQ-016","REQ-017","REQ-025","REQ-028","REQ-029"],"schemaVersion":1,"specHash":"sha256:01e0cf590b932f47a52360ef3bee7dd33632531beefb3f85ee637c793998c97f","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"88e8f9ae98adf7677a14fac1bf63bf257c5e18dd","targetTree":"9e0c4b440e47571285498cc108417e94a409efbf","verification":[{"artifacts":[".mtn-tmp/p19/quality-fix/task5-combined-build.log",".mtn-tmp/p19/quality-fix/task5-combined-tests.log"],"command":"Combined normal task build and task test including unit+race","result":"pass"},{"artifacts":[".mtn-tmp/p19/quality-fix/task5-production-archive-identity.log",".mtn-tmp/p19/quality-fix/task5-frozen-source-identity.log"],"command":"Frozen reviewed source equality/whitespace and unchanged production/later archive; target-base recheck/ancestry fast-forward","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
