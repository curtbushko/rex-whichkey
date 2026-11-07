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
  state: planned
  sourceEventID: rex-whichkey-p19-t05/build/1/resume/2
  changedAt: '2026-11-07T16:12:44.801Z'
title: Repair faithful archived fixture test setup
mtnCreateRetry:
  key: rex-whichkey-19-05-test-repair-v1
  intent: aba3cde7720dd26e51fe585bb2f5d75139d1ebbba8ef10be7d96727ff344a877
status: in-progress
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T16:00:04.596Z'
dateModified: '2026-11-07T16:12:45.424Z'
tags:
  - task
type: task
id: 8faa7318-b09c-437e-894a-dbc449443a9c
projects:
  - '[[phases/rex-whichkey-19-lua-float-launch.md]]'
---
Fix the four actual failing automated test groups shown in .mtn-tmp/p19/post-delivery-gates/unit.log and race.log: configured rename dispatch, CLI delegation, launch-origin diagnostics and action contexts. The user's explicit always-tests-and-lint/fix-the-problems directive supersedes earlier test prohibition. Read current supplementary quality design and original phase architecture. Existing archived Phase15 Lua helpers are historical test targets, NOT production; correct stale fixture/harness setup and module/source loading faithfully, preserve every assertion/scenario/security guarantee and negative path. No dropping/skipping cases, weakening guards, fabricating success, dummy compatibility objects or restoring old callbacks/controller to live source. Trace root causes and actual documented fixture interfaces before editing. Minimal genuine test-support changes, add current production regression coverage if genuinely required by a fix. No public behavior change.

First reproduce actual failing tests as RED, then fix; required normal task build and task test (unit+race) GREEN. Run full task lint after changes and record current diagnostics honestly; Task06 owns the existing lint debt, not a final exemption. No new lint failures introduced by this producer. Fresh independent spec/quality review with actual tests/build/lint runs, coordinator combined tests before target advance. User now authorizes existing automated unit/race/Lua/parser/FIFO sandbox fixtures/mocks/TDD. No live leaf Rex mutations, profile/config/install/reload/remote/deploy/push/benchmarks/global cleanup, new deps/pins, disabling lint rules or assertion weakening. Preserve separate dirty flake.nix.

Literal ownership listed includes Go Lua-test drivers and tests/lua harnesses. Archived tests/fixtures bytes stay unchanged unless concrete evidence proves a genuinely required non-weakening correction: ask coordinator for exact file scope first, not user for routine details. Go production and launcher behavior unchanged.14CPU/6GiB shared/Go5GiB, pinned existing toolchain. Commit own source hooksPath=/dev/null until full gates are repaired. Native build payload actual pass/fail, source/base/head/tree/spec/run IDs, commands/artifacts and remaining lint counts; tests and lint now MUST execute. Coordinator only .tasks/integration/real app. No phase completion from compile, optional lint or fixtures alone.

## Authorized exact historical fixture target

Coordinator traced git history: all four existing Lua test harnesses originated in unreviewed preservation checkpoint30a4439302ade4e08487d176ba6823206ddc326d; its exact Lua source supplies execution.helper/rex_bin_dir, field-specific origin diagnostics and matching registry issue/bind APIs. Current phase15-lua archive came from a later incompatible API. Add separate tests/fixtures/legacy-context-checkpoint containing the complete byte-identical Lua implementation tree from30a, with manifest source commit/Git blob identities and explicit unreviewed-historical-only label. Redirect ONLY these historical harness loaders to their matching authentic implementation; retain all42 scenarios/assertions/security/side-effect counts. Existing phase15-lua archive stays untouched. No copied code becomes production, no invented adapter or retroactive delivery claim. If exact original snapshot does not match, inspect further authentic history, do not weaken tests. Staging/loader correction and accurate fixture context remain allowed. This resolves the stopped no-change builder blocker under existing user fix authority, not another approval request.


<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDE2OjAzOjU4LjY3OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiZGlzcGF0Y2giLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMTktdDA1L2Rpc3BhdGNoLzEiLCJpbnRlbnQiOiI2NzZlMjdkNWE5MWEzOWYwOWM1OTEyOTkxNzI2ZjQ3ZWFlZThkYjU5ODhjYjdiY2NmMjEyNzRjYmIxYjdiY2E1In0 -->
### 2026-11-07T16:03:58.679Z · dispatch · unspecified

{"allowedCommands":["Owned faithful harness/source corrections; no assertion weakening","task build","task test","task lint mandatory run; remaining known lint debt Task06","format/whitespace/hooks-disabled own commit"],"allowedScope":{"interfaces":["P19-QUALITY-GATES"],"paths":["lua_context_test.go","tests/lua","direct_menu_test.go"],"resources":[],"sharedPaths":[]},"attempt":1,"baseCommit":"852fdbd7d210667dc2707e6bf8c5a0c7f9ca39f0","branch":"work/p19-t05-builder-a01","constraints":["User explicit tests/lint every change/fix problems supersedes ALL earlier test prohibitions","No old archive guard weakening/skips/deleted assertions/production callback restoration","Only coordinator mutations/integration/live app; no leaf agents","No profile/config/reload/install/remote/deploy/push/benchmarks/deps/pins/global cleanup; separate flake preserved"],"eventID":"rex-whichkey-p19-t05/dispatch/1","kind":"dispatch","limits":{"builders":1,"cpu":14,"goMemoryLimit":"5GiB","heavyGates":1,"integrations":1,"reviewers":1,"totalMemoryGiB":6},"recordRef":"tasks/rex-whichkey-19-05-test-repair.md","runID":"rex-whichkey-p19-t05-builder-a01","schemaVersion":1,"specHash":"sha256:3660f97dcb65216ebb856b0c5161e34882c11d9bf121895cef4b4acaea256520","targetBranch":"main","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t05-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDE2OjA0OjA3LjM3OFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDUtYnVpbGRlci1hMDEvcnVudGltZS9yZXNlcnZlZCIsImludGVudCI6Ijk0ZWJlYTAzMTZiYjUxZWU2OGY2ZjQyOWUzYWIxYmU2MTBhZDI0MDNmZDBmZmM2MWIxMzUxZTVjYWZhYTA3NjIifQ -->
### 2026-11-07T16:04:07.378Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-15-code-owned-runtime.md","runID":"rex-whichkey-p15-coordinator-a01"},"dispatchEventID":"rex-whichkey-p19-t05/dispatch/1","eventID":"rex-whichkey-p19-t05-builder-a01/runtime/reserved","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"75ff9ad4-a558-4295-8b20-855dafe77a8f"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-19-lua-float-launch.md","piSessionID":"rex-whichkey-p19-t05-builder-a01","piSessionName":"p19-t05-builder-a01","recordRef":"tasks/rex-whichkey-19-05-test-repair.md","rexSessionName":"rex-whichkey-main","rexWindowName":"p19-t05-builder-a01","role":"builder","runID":"rex-whichkey-p19-t05-builder-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-builder","specHash":"sha256:3660f97dcb65216ebb856b0c5161e34882c11d9bf121895cef4b4acaea256520","stage":"reserved","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p19-t05-builder-a01","windowName":"p19-t05-builder-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t05-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDE2OjA0OjEwLjcyN1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDUtYnVpbGRlci1hMDEvcnVudGltZS9sYXVuY2hlZCIsImludGVudCI6IjEyNjk5YWZmZWI2N2MyODE2ZWQxYmRlZmQ2OWUxOWNlYjBjMzAzN2M2NmJlZWUzNmUwNzYxYTg0NzAzY2Q2YTEifQ -->
### 2026-11-07T16:04:10.727Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-15-code-owned-runtime.md","runID":"rex-whichkey-p15-coordinator-a01"},"dispatchEventID":"rex-whichkey-p19-t05/dispatch/1","eventID":"rex-whichkey-p19-t05-builder-a01/runtime/launched","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"75ff9ad4-a558-4295-8b20-855dafe77a8f"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-19-lua-float-launch.md","piSessionID":"rex-whichkey-p19-t05-builder-a01","piSessionName":"p19-t05-builder-a01","provider":{"displayName":"p19-t05-builder-a01","name":"rex","resources":{"block":"block:01a1b6c0-c3b6-77d3-8bb4-e620832c92d5","endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","session":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","window":"window:01a1b6c0-c3b7-7c7f-abd2-0b665588a854"}},"recordRef":"tasks/rex-whichkey-19-05-test-repair.md","rexBlockID":"block:01a1b6c0-c3b6-77d3-8bb4-e620832c92d5","rexSessionID":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","rexSessionName":"rex-whichkey-main","rexWindowID":"window:01a1b6c0-c3b7-7c7f-abd2-0b665588a854","rexWindowName":"p19-t05-builder-a01","role":"builder","runID":"rex-whichkey-p19-t05-builder-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-builder","specHash":"sha256:3660f97dcb65216ebb856b0c5161e34882c11d9bf121895cef4b4acaea256520","stage":"launched","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p19-t05-builder-a01","windowName":"p19-t05-builder-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t05-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDE2OjEyOjI3LjkyOFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYnVpbGQiLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMTktdDA1L2J1aWxkLzEiLCJpbnRlbnQiOiJiYmJhYjJmMmVmODBiNzMyNDA3ODY0Njk2Mjc1MTBkZThhODk3YzI5ZTQ0ZjZiZmUwNTIxYWRmMmNkMTEyOTQ3In0 -->
### 2026-11-07T16:12:27.928Z · build · unspecified

{"artifactsRoot":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t05-builder-a01","attempt":1,"baseCommit":"852fdbd7d210667dc2707e6bf8c5a0c7f9ca39f0","blocker":{"evidence":["tests/lua/action_context.lua:43-46,55-76","tests/fixtures/phase15-lua/rex_whichkey.lua:258-264,287-299,307-340,414-422"],"path":"tests/lua/action_context.lua","reason":"Existing assertions require field-specific launch-origin diagnostics absent from the frozen archived implementation. The harness also supplies execution.helper/rex_bin_dir and origin, which the archive explicitly rejects. Authentic staging/loader correction cannot resolve this API/assertion conflict.","smallestCoordinatorAction":"Authorize a coherent non-weakening migration of the obsolete diagnostic assertions to the frozen archive contract, or provide an explicitly authorized matching historical fixture target. No frozen fixture or out-of-scope source was edited."},"blockerArtifact":".mtn-tmp/p19-t05/blocker.json","checks":[{"ID":"red-unit","artifact":".mtn-tmp/p19-t05/red-unit.log","command":"task test:unit","exitCode":201,"failingGroups":4,"failingScenarios":42,"result":"fail"},{"ID":"red-race","artifact":".mtn-tmp/p19-t05/red-race.log","command":"task test:race","exitCode":201,"failingGroups":4,"failingScenarios":42,"result":"fail"},{"ID":"build","artifact":".mtn-tmp/p19-t05/build.log","command":"task build","exitCode":0,"result":"pass"},{"ID":"tests","artifact":".mtn-tmp/p19-t05/test.log","command":"task test","exitCode":201,"note":"Aggregate stops at unit failure; race independently executed and failed above.","result":"fail"},{"ID":"full-lint","artifact":".mtn-tmp/p19-t05/lint.log","command":"task lint","diagnostics":53,"exitCode":201,"note":"Vet passes; golangci fails; aggregate never reaches AI lint. Remaining lint repair belongs to Task06, not a passing waiver.","result":"fail"},{"ID":"whitespace","artifact":".mtn-tmp/p19-t05/whitespace.log","command":"git diff --check","exitCode":0,"result":"pass"}],"commits":[],"dispatchEventID":"rex-whichkey-p19-t05/dispatch/1","environment":{"GOMAXPROCS":14,"GOMEMLIMIT":"5GiB","GOTOOLCHAIN":"local","go":"1.26.8","golangciLint":"2.15.0","heavyGatesConcurrent":1},"eventID":"rex-whichkey-p19-t05/build/1","files":[],"head":"852fdbd7d210667dc2707e6bf8c5a0c7f9ca39f0","headCommit":"852fdbd7d210667dc2707e6bf8c5a0c7f9ca39f0","kind":"build","limitations":["No GREEN unit/race suites; task not delivered.","No architecture configuration exists.","No independent review or coordinator combined gates performed.","No mtn writes, integration, agent launches, Rex mutations or parent source changes."],"nextSteps":["Coordinator resolves exact assertion/archive conflict before implementation continues.","Repair faithful harness, rerun build/unit/race/full lint, commit owned source, then obtain fresh independent full-test review.","Task06 and coordinator must establish final combined lint/test/build GREEN; no whole-phase GREEN claimed."],"outcome":"blocked","preservation":{"allAssertionsUnchanged":true,"archivedLuaFilesByteIdenticalToOriginal":12,"assertionGroups":["TestConfiguredRenameDispatch","TestCLIDelegation","TestLaunchOriginPreflightDiagnostics","TestRexActionContexts"],"lintConfigurationByteIdenticalToParent":true,"originalArchiveCommit":"cd5cbce818e07c264738ddcbc591ad49fdf0e9c8","productionUnchanged":true,"scenariosUnchanged":42,"setupCorrectionsApplied":[],"trackedFixtureFilesByteIdentical":13},"protocolNormalization":{"change":"Native outcome/head/validation aliases for blocked no-change result; original status/headCommit/checks preserved, no outcome altered","rawArtifact":".mtn-tmp/p19/quality-fix/task5-build-raw.json"},"recordRef":"tasks/rex-whichkey-19-05-test-repair.md","rootCauses":["Configured rename uses incomplete GUI placement/metadata/expiry, token-only binding and v1 completion; archived dispatch requires exact GUI binding, preparation receipt and v2 receipt hash.","Gopher-Lua LoadFile preserves its filename as source without the @ prefix required by archived discovery. Faithful file-backed loader adaptation is possible, but does not resolve obsolete launch APIs.","CLI delegation mocks the older origin/json.decode/start interface; the archive uses strict raw JSON describe/relay/prepare replies and controller wait acceptance.","Action-context diagnostic assertions target an older interface and cannot execute successfully against the frozen archive without changing their contract."],"runID":"rex-whichkey-p19-t05-builder-a01","schemaVersion":1,"specHash":"sha256:3660f97dcb65216ebb856b0c5161e34882c11d9bf121895cef4b4acaea256520","status":"blocked","summary":"Reproduced failures and traced incompatible historical harness/archive APIs. No source changes or fabricated GREEN evidence.","tree":"c18afc6a5865b552c8e6226141f27eb70b0b7934","validation":[{"ID":"red-unit","artifact":".mtn-tmp/p19-t05/red-unit.log","command":"task test:unit","exitCode":201,"failingGroups":4,"failingScenarios":42,"result":"fail"},{"ID":"red-race","artifact":".mtn-tmp/p19-t05/red-race.log","command":"task test:race","exitCode":201,"failingGroups":4,"failingScenarios":42,"result":"fail"},{"ID":"build","artifact":".mtn-tmp/p19-t05/build.log","command":"task build","exitCode":0,"result":"pass"},{"ID":"tests","artifact":".mtn-tmp/p19-t05/test.log","command":"task test","exitCode":201,"note":"Aggregate stops at unit failure; race independently executed and failed above.","result":"fail"},{"ID":"full-lint","artifact":".mtn-tmp/p19-t05/lint.log","command":"task lint","diagnostics":53,"exitCode":201,"note":"Vet passes; golangci fails; aggregate never reaches AI lint. Remaining lint repair belongs to Task06, not a passing waiver.","result":"fail"},{"ID":"whitespace","artifact":".mtn-tmp/p19-t05/whitespace.log","command":"git diff --check","exitCode":0,"result":"pass"}],"worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t05-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDE2OjEyOjMxLjI3OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVsZWFzZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDUtYnVpbGRlci1hMDEvcmVsZWFzZS9yZXNlcnZlZCIsImludGVudCI6ImU3ZGIxNGI0ODlmMWZmYTFhYjM3MWY2YjFjOTMxNDEzMmM3ZWQxNWY5YjE0ZjQ0MGRiYTc3ZTA1YTQyYWE2YjgifQ -->
### 2026-11-07T16:12:31.279Z · release · unspecified

{"attempt":1,"dispatchEventID":"rex-whichkey-p19-t05-builder-a01/runtime/reserved","eventID":"rex-whichkey-p19-t05-builder-a01/release/reserved","evidence":"Trusted blocked final result/zero worker exit retrieved. No source edits; exact original harness/archive contract mismatch traced. Blocked recovery retained; coordinator authorizes authentic matching historical fixture under fix scope.","kind":"release","outcome":"finished","recordRef":"tasks/rex-whichkey-19-05-test-repair.md","runID":"rex-whichkey-p19-t05-builder-a01","schemaVersion":1,"specHash":"sha256:3660f97dcb65216ebb856b0c5161e34882c11d9bf121895cef4b4acaea256520"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDE2OjEyOjM0LjY3MFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVsZWFzZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDUtYnVpbGRlci1hMDEvcmVsZWFzZS9sYXVuY2hlZCIsImludGVudCI6ImJkN2RhZWMzMmQ0MjZiMTIyMGMwYTI1NzRkYTNiOWI5NjQ1NWMxZTQwNWIyMzE2NDQxMzdiMDA4NTM4ODdiMzEifQ -->
### 2026-11-07T16:12:34.670Z · release · unspecified

{"attempt":1,"dispatchEventID":"rex-whichkey-p19-t05-builder-a01/runtime/launched","eventID":"rex-whichkey-p19-t05-builder-a01/release/launched","evidence":"Trusted blocked final result/zero worker exit retrieved. No source edits; exact original harness/archive contract mismatch traced. Blocked recovery retained; coordinator authorizes authentic matching historical fixture under fix scope.","kind":"release","outcome":"finished","recordRef":"tasks/rex-whichkey-19-05-test-repair.md","runID":"rex-whichkey-p19-t05-builder-a01","schemaVersion":1,"specHash":"sha256:3660f97dcb65216ebb856b0c5161e34882c11d9bf121895cef4b4acaea256520"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDE2OjEyOjM4LjAyNFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVsZWFzZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDUvcmVsZWFzZS8xIiwiaW50ZW50IjoiNzhiZjE1MTg1M2VkMTYxNDVkYjIzNzFkZjgyYzk3ZWVlY2M5MjAyOGYzZGVmMDQxNmQ4ZmJiMDRiMjQ1NmZhYyJ9 -->
### 2026-11-07T16:12:38.024Z · release · unspecified

{"attempt":1,"dispatchEventID":"rex-whichkey-p19-t05/dispatch/1","eventID":"rex-whichkey-p19-t05/release/1","evidence":"Blocked no-change builder stopped, safe matching original checkpoint fixture approach identified by coordinator; no test weakening/new production behavior required.","kind":"release","outcome":"stopped","recordRef":"tasks/rex-whichkey-19-05-test-repair.md","schemaVersion":1,"specHash":"sha256:3660f97dcb65216ebb856b0c5161e34882c11d9bf121895cef4b4acaea256520"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDE2OjEyOjQ0LjAyMFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVzdW1lIiwicmV0cnlLZXkiOiJyZXgtd2hpY2hrZXktcDE5LXQwNS9idWlsZC8xL3Jlc3VtZS8yIiwiaW50ZW50IjoiMjIwNGNhZDExNmQ2YTUxNGZjZDYxODhhZTE4ZGNjOGM2ZjE3YWY0MzA1ZWRjZjRhZDY0NDA2NjA0ODUwNTc5NCJ9 -->
### 2026-11-07T16:12:44.020Z · resume · unspecified

{"attempt":2,"authority":"Existing user fix-the-problems instruction and bounded fixture repair scope","blockerEventID":"rex-whichkey-p19-t05/build/1","eventID":"rex-whichkey-p19-t05/build/1/resume/2","kind":"resume","recordRef":"tasks/rex-whichkey-19-05-test-repair.md","resolution":"Matching original30a checkpoint Lua source and harness origin confirmed by git log/show. Precise separate test-only authentic fixture ownership authorized; no assertion weakening/production change/archive rewrite required. Original checkpoint remains unreviewed historical source, not delivered code.","schemaVersion":1,"specHash":"sha256:01e0cf590b932f47a52360ef3bee7dd33632531beefb3f85ee637c793998c97f"}
<!-- mtn:end-entry:v1 -->
