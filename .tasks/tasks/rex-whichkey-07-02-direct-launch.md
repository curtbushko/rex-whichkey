---
workflowKind: task
workflowLanguage: go
workflowOrder: 2
workflowRequirements:
  - REQ-016
  - REQ-017
workflowDocuments:
  - ref: prds/rex-whichkey-usability.md
    hash: 'sha256:c2df9b9404ce7f2d649eed40d5c219f7098e9389c6407159b6ae35eddb7746c5'
  - ref: designs/rex-whichkey-usability.md
    hash: 'sha256:c8553dd4ac48fcd34a43d7cb3cd842bb18c65310d0fc6133cf1b6e3f196f1b8a'
workflowOwnership:
  paths:
    - launch.go
    - cmd/main.go
    - popup.go
    - lua/rex_whichkey.lua
    - docs/launch-contract.md
  sharedPaths: []
  resources: []
  interfaces:
    - canonical-popup-endpoint
    - public-original-vm-launch-origin
projects:
  - '[[phases/rex-whichkey-07-usability.md]]'
status: done
priority: normal
title: Wire no-argument popup through original-runtime launch action
blockedBy:
  - uid: '[[tasks/rex-whichkey-07-01-endpoint-policy.md]]'
    reltype: FINISHTOSTART
mtnCreateRetry:
  key: rex-whichkey/usability/task07-02/v1
  intent: bdd04ddc81732b7bd3222560891194003927616a336b2d2bac13c837dc39f136
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-01T15:15:18.902Z'
dateModified: '2026-11-01T15:51:40.075Z'
tags:
  - task
type: task
id: e8a0fa34-695d-4967-82ff-9ef716541464
completedDate: '2026-11-01'
---
# Wire genuine no-argument original-runtime popup activation

## Authority and inputs

REQ-016 and REQ-017; Phase07 criteria 1, 2, 3, 4 and 5. Requires reviewed/integrated task07-01 endpoint/preparation producer. Read the exact phase/usability PRD, `.mtn-tmp/usability/planner-source-contracts.md`, task07-01 `docs/launch-contract.md`, `cmd/main.go`, `popup.go`, `lua/rex_whichkey.lua`, and retained `lua/rex_whichkey/dispatch.lua` and Rex resource APIs. Coordinator resolved the draft endpoint-information blocker by adopting native Rex CLI default discovery in the approved usability design. Consume the actual integrated Go resolver; do not reproduce unavailable platform path algorithms or re-escalate missing routine approval.

## Literal ownership and outputs

Owned paths only: new root `launch.go`, existing `cmd/main.go`, `popup.go`, `lua/rex_whichkey.lua`, and `docs/launch-contract.md`. Primary Go, substantive Lua and Markdown. `popup.go` ownership permits reuse/extraction of exact origin checks without duplicate inconsistent placement/generation validation. Do not write dispatcher, mapping format, dependency files, Taskfile/lint/hook definitions, tests, user configuration or managed records. Serialize after task07-01 and before docs/final assembly.

## Required implementation and acceptance

1. Root `rex-whichkey` with no arguments must call a real new Go `LaunchPopup` path, not return usage or run a JSON demo. Public root flags may include `--server`, `--client` and `--action` (default `whichkey`), plus `--help`/`-h`. Public help describes `rex-whichkey [options]`, real setup/context prerequisites and failures; private `menu`, `prepare`, `controller`, `popup` remain compatible but are not presented as the normal startup command. Help and argument errors must not contact Rex. Retain existing command implementation and private-pipe security behavior.
2. Capture server-host REX_SESSION and REX_BLOCK as full typed IDs; missing, malformed or stale values fail actionably without selecting focus/current targets. Resolve and freeze endpoint through task07-01's Go policy. Validate the captured session and uniquely placed block there before mutation. Resolve a supplied `--client` or REX_CLIENT only as an exact full ID attached to that session with nonempty connected_at and nonlost state; reject labels/suffixes/duplicates. If neither is supplied, inspect only THAT captured session's attached clients, reject malformed/duplicate entries and require exactly one eligible full ID/generation. Zero or several eligible clients fail with instructions to supply an exact client; never use Rex's implicit current/single-app-client targeting, global client selection, focused block or arbitrary first match. This relationship lookup is explicitly scoped to the captured terminal session, not ambient focus.
3. Implement the actual argument-array named-configuration-action caller in `launch.go`: invoke the installed Rex executable once with explicit concrete `--server`, `--client`, `--autostart=false`, bounded timeout, `do ACTION --session FULL_SESSION --block FULL_BLOCK --args -`. Use stdin for an origin envelope; do not run `rex do -e`, a Lua file, a second VM, a private fake/token-only command or a new daemon. No hand-authored user JSON is required. Bound child lifetime/output and quote/sanitize actionable errors without exposing private handoffs. Do not retry an uncertain activation or claim the menu was displayed merely because activation was accepted.
4. Define the small Go-owned launch origin envelope in `launch.go`, and document its fields/validation in `docs/launch-contract.md`: canonical endpoint, full session/block/client and connected_at. It is transport context, not executable Lua, a capability, a mapping option or authority to fabricate ctx. Own the strict Go serialization/validation; Lua consumes documented fields and compares against the actual action ctx. Do not add another mapping schema or generic framework.
5. Add a real `menu.register_launch(name)` method in the existing Lua composition. At configuration load it registers `rex.action` with optional origin-envelope data and a run function that uses the ACTUAL fresh ctx and calls delivered `menu.launch`. Default installed name `whichkey` matches Go; alternate registered names use public `--action`. Validate action/envelope inputs, compare full origin IDs and live connection generation before issuing/evaluating, compare expected endpoint with task07-01's canonical preparation result before starting the controller, and recheck expected identity/generation across yielding boundaries. On mismatch cancel any newly issued reservation and fail without controller creation/focus/action dispatch. Do not stringify/copy opaque ctx.server into the envelope or replace ctx with user data. Preserve `menu.launch(ctx)` and key/palette invocation without an envelope using their genuine actual context. Retain all registry/completion semantics and no-return-of-secrets rules.
6. Source-trace the complete real route: no-argument parse -> LaunchPopup -> explicitly targeted installed configuration action -> register_launch run -> menu.launch -> default loader/evaluation/prepare -> detached controller -> owned popup terminal -> existing RunTerminal/RenderMenu/HandleKey -> authenticated fresh completion -> original VM callbacks/actions. The user installs modules, registers completion, sets up defaults and registers launch once in user-owned init.lua through ordinary Rex administration; code must neither edit nor reload it. Missing registration/Remote Control/helper/context fail clearly, not as a generic usage-only path.

## Verification and handoff constraints

No tests/TDD/experiments, product execution, GUI/live/remote/API trials, config reloads or dependency installation. Preserve all existing tests/gates; never run hidden test hooks. Builder may format and statically inspect Go/Lua/Markdown and whitespace only. A fresh independent reviewer explicitly covers BOTH Go and Lua and the entire actual startup/callback/origin/error path, plus compatibility of internal modes. Coordinator compiles the actual combined revision with existing `task build` under the allocated gate and publishes exact review/combined Git identity, ancestry/diff and integration proof before completion. Compilation is not runtime verification. No worker mtn writes, integration or agent launch. Return revision-bound evidence/limitations, not an uncreated-record fingerprint or an inferred popup runtime pass.


<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDE1OjUxOjQwLjA2NFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNy10MDItYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiYjFlM2Y5YzE4YzY4NmY2ZjFlNTkzZmIxOWY4ODE4MTI5ZGUyYTU0NDhiMGM1YmJjNGMxN2U2YzYxOGRmMTFiMCJ9 -->
### 2026-11-01T15:51:40.064Z · summary · unspecified

{"acceptedScope":"Wire no-argument popup through original-runtime launch action","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p07-t02-reviewer-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/usability/rex-whichkey-p07-t02-builder-a01/integration/integration.json"],"attempt":1,"eventId":"rex-whichkey-p07-t02-builder-a01/summary/1","files":["cmd/main.go","docs/launch-contract.md","launch.go","lua/rex_whichkey.lua","popup.go"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-07-02-direct-launch.md","requirements":["REQ-016","REQ-017"],"schemaVersion":1,"specHash":"sha256:408dd4478e572880d5b6fe9d0c3354876e5d583f47569f398c3c14d40e60af6b","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"69078fcee3dc0bbd9598e1b58e33920934e82497","targetTree":"82dc432f715295a3b82829be8be509e83aa0197a","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
