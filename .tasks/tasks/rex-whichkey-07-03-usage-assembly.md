---
workflowKind: task
workflowLanguage: docs
workflowOrder: 3
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
    - README.md
    - examples/README.md
    - docs/remote-compatibility.md
    - docs/mapping-format.md
    - docs/acceptance.md
    - docs/launch-contract.md
  sharedPaths: []
  resources: []
  interfaces:
    - public-startup-documentation
    - usability-acceptance-inventory
projects:
  - '[[phases/rex-whichkey-07-usability.md]]'
status: done
priority: normal
title: Document ordinary startup and assemble usability delivery
blockedBy:
  - uid: '[[tasks/rex-whichkey-07-01-endpoint-policy.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-07-02-direct-launch.md]]'
    reltype: FINISHTOSTART
mtnCreateRetry:
  key: rex-whichkey/usability/task07-03/v1
  intent: 2d219edc8b664e668ef365ef5ee714f63466f2656e4a0dce78fdf75e3d6954e4
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-01T15:15:20.700Z'
dateModified: '2026-11-01T16:22:52.367Z'
tags:
  - task
type: task
id: f7724bb4-4f07-4f79-b87d-98ce766e0ac3
completedDate: '2026-11-01'
---
# Document ordinary startup and assemble usability delivery

## Authority and inputs

REQ-016 and REQ-017; Phase07 criteria 1-5. Depends on reviewed/integrated task07-01 and task07-02, not just their descriptions or builder exits. Read the exact phase/usability PRD, `.mtn-tmp/usability/planner-source-contracts.md`, actual integrated launch/endpoint/Lua APIs, `docs/launch-contract.md`, retained mapping/remote/safety and packaging sources. Preserve old completed phase identities, decision history and historical verification limitations.

## Literal ownership and outputs

Markdown-only writes: `README.md`, `examples/README.md`, `docs/remote-compatibility.md`, `docs/mapping-format.md`, `docs/acceptance.md`, `docs/launch-contract.md`. Primary language docs. Read Go/Lua and Nix packaging source for exact statements, but do not write them or tests/Taskfile/lint/hook/user/managed files. `docs/launch-contract.md` overlap is serialized after both producers. Mapping semantics remain Go-owned and unchanged; edits to mapping-format.md only repair composition/startup/endpoint guidance.

## Acceptance criteria

1. Give an executable, internally consistent one-time user-owned setup example using installed Lua modules/defaults, completion registration, `popup.new(defaults)` with no mandatory socket, successful setup and `menu.register_launch("whichkey")`. Explain exact constructor/public endpoint and action-name overrides actually implemented. Preserve direct original-runtime callback semantics and existing custom mapping/default precedence.
2. Show `rex-whichkey` as the ordinary terminal command that activates a Rex floating popup using the terminal's own captured session/block and an explicit or uniquely resolved attached client. Distinguish that popup path from standalone low-level JSON menu mode. Do not recommend hidden subcommands, user-written resolved JSON, a script VM, a greeting/demo, auto-configuration or reload by this tool.
3. Document native Linux/macOS default discovery delegated to the installed Rex CLI and environment/explicit precedence, execution/server-host resolution, custom/shared-server overrides, and the difference between concrete endpoint routing and opaque ctx.server authority. Do not invent exact filesystem formulas: Rex owns those algorithms, and Go captures/validates its returned concrete endpoint. Explain normal default-server availability and actionable no-autostart discovery failures; no viewing-Mac/home guesses or unverified platform certification. Revise obsolete mandatory-endpoint statements in README, examples, mapping and remote docs without silently rewriting historical decision evidence. Explain refinement of old explicit-only choice, retaining explicit overrides/fail-closed rationale.
4. Explain module/PATH/helper/CLI/setup/launch-registration prerequisites, Remote Control, missing or stale REX_SESSION/REX_BLOCK, absent/ambiguous client generation, exact-client override, malformed endpoint/missing runtime/home, wrong server realm and failed/uncertain activation. No ambient focus/current-client workaround or action retry. Keep user init.lua ownership, no daemon/reload/takeover, owned-only cleanup and security warnings explicit.
5. Update acceptance inventory for REQ-016/017 with actual reviewed source and coordinator-supplied compilation/integration evidence. Separate planned contracts, implemented source traces, actual static/compile results, and runtime-unverified behavior. Preserve historical completed phases/requirements/decisions and limitations; do not turn old evidence or skipped runtime gates into passing usability acceptance.
6. Fresh independent documentation/source-consistency review must trace the documented no-argument route through integrated Go and Lua and catch dead/fictional API examples. Coordinator performs final assembled-target compilation under allocation using existing `task build` without tests, verifies exact Git candidate/target ancestry and integrated reviewed diffs, then commits actual Phase07 delivery on main with test-running hooks bypassed only per invocation under retained policy. Builders/reviewers do not integrate, commit phase completion or write mtn. Coordinator publishes actual delivery evidence and compaction only after criteria are met; a docs task cannot itself declare the phase complete.

## Verification constraints and outputs

No project tests/TDD/experiments/product execution/GUI/live/remote/API probes/config reloads/dependency installs or package activation trials. No lint/Taskfile/hook weakening. Preserve tests. Allowed documentation evidence is static review/whitespace/link-to-local-source consistency, not executing example snippets or probing documented endpoints. Return changed Markdown paths, exact revision, reviewable source-to-doc traces, actual static checks and explicit remaining limitations. Combined compilation and Git integration are coordinator gates; never report them passed without supplied actual evidence. If a producer defect requires source changes, return the exact finding to that owner and serial coordinator instead of exceeding literal ownership.



## Final delivery-status accuracy
The coordinator final audit found stale ephemeral verification-status claims in README.md after reviewed/compiled phase delivery. Read .mtn-tmp/usability/final-doc-correction.md and correct only existing owned Markdown: source review and actual Go/Nix compilation must be accurately distinguished from prohibited/unexecuted tests and runtime trials. Remove self-staling claims that this revision still awaits review/assembly; no future passes or fake verification. Existing implementation remains unchanged. A fresh frozen review/integration is required for this scoped correction.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDE2OjIyOjUyLjM1NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNy10MDMtYnVpbGRlci1hMDIvc3VtbWFyeS8yIiwiaW50ZW50IjoiNDFlM2U2ODIwNzMwOTJjOWVmNzY3YjQyNDUzOWM5YmZkOWUwNmZhYjVhYzg5MTgzMTg4ODk3MDI3NGU4MGU4NCJ9 -->
### 2026-11-01T16:22:52.355Z · summary · unspecified

{"acceptedScope":"Document ordinary startup and assemble usability delivery","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p07-t03-tasks-a02/review/2 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/usability/rex-whichkey-p07-t03-builder-a02/integration/integration.json"],"attempt":2,"eventId":"rex-whichkey-p07-t03-builder-a02/summary/2","files":["README.md","docs/acceptance.md","docs/launch-contract.md"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-07-03-usage-assembly.md","requirements":["REQ-016","REQ-017"],"schemaVersion":1,"specHash":"sha256:11d9db51ec6fc503620a996210a9cc7d024aaa83b7189b230e2102b01f174423","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"1a43a590540dfd81c53ab9c097f72b8c8ec1b520","targetTree":"74b33db063d47b29025fefcabe8712c9bbcedc5f","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
