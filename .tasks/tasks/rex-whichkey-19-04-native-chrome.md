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
  state: planned
  sourceEventID: rex-whichkey-p19-t04/chrome-blocker/1/resume/2
  changedAt: '2026-11-07T07:21:45.094Z'
title: 'Phase19.4: Required native float chrome via verified Rex capability'
mtnCreateRetry:
  key: rex-whichkey-p19/task-create/native-chrome/1
  intent: c650c07508d11dc05b7c59b8bc32fdc5a3efa821338a0fb12be741cd713c64bc
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T04:48:27.141Z'
dateModified: '2026-11-07T09:46:55.216Z'
tags:
  - task
type: task
id: 6dadf60f-878a-4332-95c0-fe56f0cf3c79
projects:
  - '[[phases/rex-whichkey-19-lua-float-launch.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-19-03-source-run-delivery.md]]'
    reltype: FINISHTOSTART
---
Record/deliver the EXPLICITLY APPROVED default-native-chrome policy and final current runtime evidence after task3. User1a accepts Rex default chrome FOR NOW; REQ-026 native titlebar/border/corner/shadow customization is DEFERRED, NOT implemented or credited. This task's new accepted scope is accurate policy/docs/optional pure Lua default-policy clarity, not unsupported chrome controls. Preserve original required-feature/Open blocker/history and link accepted refinement. No ignored knobs, invented schema/API flags, global Rex settings changes, terminal-color theme substitution or painted fake borders. No custom chrome implementation without a real supported future capability.

Update current appearance/launch/acceptance/package guidance using actual task1-3 reviewed/integrated make run evidence, code/source/command/Rex IDs and limits. Current native defaults may still show headers/corners/shadows; say so. Go still matches terminal colors/menus. Explain accepted visible/tty/worker/focus-before-effects with detached host through dispatch and eventual Rex exit cleanup; never claim all resources deleted before effects or verified post-exit by Go itself. Preserve historical acceptance sections verbatim, no whole-project/keybinding/profile/platform/test/lint-green claims. Default-package Lua/Go entry consistent, no new runtime policies.

Only listed paths, changes minimal/current evidence. Required git diff --check plus independent docs/source/evidence accuracy review; Nix formatting/parser if package file actually changes (no build/eval/activation). No tests/harness/benchmark/live leaf calls/environment/config/reload/deploy/push. Coordinator performs final ordinary make run against final reviewed combined/main source, verifies evidence source equality and exact cleanup. Phase completion covers all nondeferred functionality and approved explicit Req026 deferral, never delivers unsupported native appearance control. Preserve separate dirty flake.nix and failed resources/history.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA0OjQ4OjQ0LjE0OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYmxvY2tlciIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDQvY2hyb21lLWJsb2NrZXIvMSIsImludGVudCI6IjY1NmNjNWY5YzljY2I1MmJiNjdkY2NiYTM4Nzc3MGNlNTY5ZThkZWZmZjcxNWFhMGU1YTk3YzU2ZTI5ZjA5ZmUifQ -->
### 2026-11-07T04:48:44.149Z · blocker · unspecified

{"attempt":1,"decisionID":"REXWHICHKEY-20261106-234819","eventID":"rex-whichkey-p19-t04/chrome-blocker/1","evidence":[".mtn-tmp/p19/rex-api-list.json",".mtn-tmp/p19/rex-block-types.json","https://discord.com/channels/1546797889009881158/1566968591062929548"],"kind":"blocker","reason":"Installed Rex new_layer/terminal creation and published methods expose no per-float native titlebar/border/corners/shadow controls. Terminal theme is colors, not chrome; active Rex feature request confirms unresolved dependency.","recordRef":"tasks/rex-whichkey-19-04-native-chrome.md","requiredDecision":"Provide an actual supported Rex chrome capability/version within authorized access, or explicitly defer/change chrome requirement. Default chrome is not silently accepted and guessed/ignored flags are forbidden.","schemaVersion":1,"specHash":"sha256:b61c8b6635362677eaba352f4b01d34da26528a75031634721cebde8d9de6431"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA3OjIxOjQ0LjMwM1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVzdW1lIiwicmV0cnlLZXkiOiJyZXgtd2hpY2hrZXktcDE5LXQwNC9jaHJvbWUtYmxvY2tlci8xL3Jlc3VtZS8yIiwiaW50ZW50IjoiMjg1MTIzYzk0NGMyZTFhZTRkYzA3NmY4ZTFmNjRjMGQxMzdlNWYzYzM1ZTI0NTBiNmM5OTBhZjE2ODdiZWNjMyJ9 -->
### 2026-11-07T07:21:44.303Z · resume · unspecified

{"attempt":2,"blockerEventID":"rex-whichkey-p19-t04/chrome-blocker/1","decisionEvidence":"1a - yes, accept it for now; 1b - yes","eventID":"rex-whichkey-p19-t04/chrome-blocker/1/resume/2","kind":"resume","recordRef":"tasks/rex-whichkey-19-04-native-chrome.md","resolution":"User explicitly approves1a default chrome deferral and1b host-through-dispatch exception. Replanned current-used Go/Lua vertical slice provides real consumer with widened literal ownership; no dormant API or unapproved ordering change. Prior source attempt stopped/no edits; original blocker history preserved.","schemaVersion":1,"specHash":"sha256:f5706caed695e5e815779f2f8cdf6c0f581286a3bd98409342d8a0e9ecc7b0d4"}
<!-- mtn:end-entry:v1 -->
