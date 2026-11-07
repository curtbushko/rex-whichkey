---
identity:
  kind: task
  taskID: rex-whichkey-19-04-native-chrome
  uuid: 934a4988-4ad8-47b5-ac4b-edf7fc7c6256
structure:
  order: 4
scope:
  requirements:
    - REQ-026
    - REQ-014
    - REQ-015
  documents:
    - ref: designs/rex-whichkey-lua-float-launch.md
      hash: 'sha256:f7e38d49287de319753732b09fe472a40dfadcef81e5245e2deac208e968ab70'
implementation:
  language: docs
  ownership:
    paths:
      - README.md
      - docs/acceptance.md
      - docs/appearance.md
      - lua/rex_whichkey.lua
      - nix/package.nix
    sharedPaths: []
    resources: []
    interfaces:
      - P19-SHARED-LUA-ENTRY
      - P19-NATIVE-CHROME
  gates:
    - ID: whitespace
      command: git diff --check
      required: true
lifecycle:
  state: complete
  sourceEventID: rex-whichkey-p19-t04/summary/1
  changedAt: '2026-11-07T14:52:49.962Z'
title: 'Phase19.4: Required native float chrome via verified Rex capability'
mtnCreateRetry:
  key: rex-whichkey-p19/task-create/native-chrome/1
  intent: c650c07508d11dc05b7c59b8bc32fdc5a3efa821338a0fb12be741cd713c64bc
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T04:48:27.141Z'
dateModified: '2026-11-07T14:52:50.982Z'
tags:
  - task
type: task
id: 6dadf60f-878a-4332-95c0-fe56f0cf3c79
projects:
  - '[[phases/rex-whichkey-19-lua-float-launch.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-19-03-source-run-delivery.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-07'
---
Record/deliver the EXPLICITLY APPROVED default-native-chrome policy and final current runtime evidence after task3. User1a accepts Rex default chrome FOR NOW; REQ-026 native titlebar/border/corner/shadow customization is DEFERRED, NOT implemented or credited. This task's new accepted scope is accurate policy/docs/optional pure Lua default-policy clarity, not unsupported chrome controls. Preserve original required-feature/Open blocker/history and link accepted refinement. No ignored knobs, invented schema/API flags, global Rex settings changes, terminal-color theme substitution or painted fake borders. No custom chrome implementation without a real supported future capability.

Update current appearance/launch/acceptance/package guidance using actual task1-3 reviewed/integrated make run evidence, code/source/command/Rex IDs and limits. Current native defaults may still show headers/corners/shadows; say so. Go still matches terminal colors/menus. Explain accepted visible/tty/worker/focus-before-effects with detached host through dispatch and eventual Rex exit cleanup; never claim all resources deleted before effects or verified post-exit by Go itself. Preserve historical acceptance sections verbatim, no whole-project/keybinding/profile/platform/test/lint-green claims. Default-package Lua/Go entry consistent, no new runtime policies.

Only listed paths, changes minimal/current evidence. Required git diff --check plus independent docs/source/evidence accuracy review; Nix formatting/parser if package file actually changes (no build/eval/activation). No tests/harness/benchmark/live leaf calls/environment/config/reload/deploy/push. Coordinator performs final ordinary make run against final reviewed combined/main source, verifies evidence source equality and exact cleanup. Phase completion covers all nondeferred functionality and approved explicit Req026 deferral, never delivers unsupported native appearance control. Preserve separate dirty flake.nix and failed resources/history.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDE0OjUyOjQ4LjU4M1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDQvc3VtbWFyeS8xIiwiaW50ZW50IjoiNTU5MWU4MjI0Y2E0MjllOWM0ZDc4ZjcxOTk0OTg3M2IxMjg4NzYwYWZkYzgzNTIyMTc2YWM1NTU5MDY2NDEwNiJ9 -->
### 2026-11-07T14:52:48.583Z · summary · unspecified

{"acceptedScope":"Docs-only README/appearance/current acceptance refresh: native Rex default chrome explicitly accepted, Req026 customization deferred/not delivered, Lua presentation/sizing/launch vs Go menu/palette separated; accepted post-effect detached-host limit and real Task03 combined/main makerun evidence/lint53/uncertainty documented. Historical acceptance suffix and ALL runtime code/config/recipes/package/test/hook/dependency blobs unchanged.","approvalSummary":["Explicit user1a native chrome defaults for now and1b host-through-action retirement; original Phase19 finish/commit/real make run authority","Exact approved design/current plan and native delivered Task03 with actual combined/main source-route evidence","Fresh same-family docs/source/evidence combined review1 approved exact8a7fe5d; source identity/combined normal build-vet and final actual candidate+MAIN makerun/exit0/cleanup passed"],"attempt":1,"cleanup":{"eligibleStoppedCleanDeliveredWorktreesRemoved":true,"facade":"Official protocol/native completion verified; immutable helper cleanup jq defect persists, exact eligible stopped clean delivered worktrees removed manually","finalOwnedValidationWindowClosed":true,"runtimeReservationsReleased":true,"successfulWorkerWindowsAbsent":true},"deferredRequirements":[{"authority":"User1a accepts native defaults for now","requirement":"REQ-026","scope":"Native chrome customization NOT implemented; task delivers explicit policy/evidence","state":"Deferred"}],"eventID":"rex-whichkey-p19-t04/summary/1","files":["README.md","docs/appearance.md","docs/acceptance.md",".mtn-tmp/p19/resume/task4-build-result.json",".mtn-tmp/p19/resume/task4-review-result.json",".mtn-tmp/p19/resume/task4-public-evidence.json",".mtn-tmp/p19/resume/final-main-observed-rex-whichkey"],"kind":"summary","limitations":["Current full lint53 diagnostics/exit201/AI not reached, not green; no suppressed standards","User1b host-through-effect/post-exit retirement limit retained; no all-owned-block-before-effects claim","Source docs only; no prohibited tests/harnesses/benchmark/profile/Nix build/reload/deploy, physical keybinding/GUI/other-theme/platform or all error branches claimed","Fresh same-family independent review, not cross-family"],"recordRef":"tasks/rex-whichkey-19-04-native-chrome.md","requirements":["REQ-026","REQ-014","REQ-015"],"schemaVersion":1,"specHash":"sha256:dc43d6c8b83e211ae455872b44ee7c0c280187d76d4a42f2ef5ea15d38a31606","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"fe9746d0b3856e67c66f942b59d29959cba98730","targetTree":"bcdca673e8ed2d76fc0296eaf427b8bffaba5ad5","verification":[{"artifact":".mtn-tmp/p19/resume/final-runtime-source-identity.log","command":"Fresh exact-source combined docs/evidence review, frozen owned docs equality, historical acceptance suffix/runtime-source/protected old blobs identity and whitespace","result":"pass"},{"artifacts":[".mtn-tmp/p19/resume/task4-integration-build.log",".mtn-tmp/p19/resume/task4-integration-vet.log"],"command":"Combined final normal task build/task lint:vet and exact target-base recheck/ancestry-preserving fast-forward","result":"pass"},{"artifact":".mtn-tmp/p19/resume/task4-public-evidence.json","command":"ACTUAL final reviewed combined and rebuilt MAIN make run/menu/current Go/dismissal child exit0/exact host-focus-primary cleanup","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
