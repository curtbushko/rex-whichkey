---
identity:
  kind: task
  taskID: rex-whichkey-20-01-background-scripts
  uuid: afbec1d2-ee52-400e-8ecd-9b8127f2995c
structure:
  order: 1
scope:
  requirements:
    - REQ-SCRIPT-001
    - REQ-SCRIPT-002
    - REQ-SCRIPT-003
    - REQ-SCRIPT-004
  documents:
    - ref: designs/rex-whichkey-background-scripts.md
      hash: 'sha256:1a7cb4ce61a8d2723f031e956b35b9a3e758285a41152d04c6d63debf2e99eb9'
    - ref: designs/rex-whichkey-quality-gates.md
      hash: 'sha256:f3120762f7075c7c767e5f39552ba9c2b38b2e213d33d2472d07f0bc0c1cbdb2'
  dependencies: []
implementation:
  language: go
  ownership:
    paths:
      - data_config.go
      - mapping.go
      - action.go
      - direct_menu.go
      - native.go
      - launch.go
      - script_action.go
      - script_process_unix.go
      - script_process_darwin.go
      - script_process_linux.go
      - script_menu.go
      - script_action_test.go
      - script_process_test.go
      - script_process_darwin_test.go
      - script_process_linux_test.go
      - script_menu_test.go
      - direct_lifecycle_darwin_test.go
      - rex/script_context.go
      - rex/script_context_test.go
      - rex/block.go
      - rex/layer.go
      - rex/direct_popup.go
      - rex/direct_popup_test.go
      - rex/snapshot.go
      - rex/snapshot_test.go
      - README.md
      - docs/actions.md
      - docs/mapping-format.md
      - docs/launch-contract.md
      - docs/acceptance.md
      - examples/background-scripts.json
      - rex/window.go
    sharedPaths: []
    resources: []
    interfaces:
      - SCRIPT-ACTION
      - SCRIPT-ORIGIN-CONTEXT
      - SCRIPT-STARTUP-LIFETIME
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
      includes:
        - 'task test:unit'
        - 'task test:race'
    - ID: full-lint
      command: task lint
      required: true
      includes:
        - 'lint:vet'
        - 'lint:golangci'
        - 'lint:ai'
  reviewPlan:
    languages:
      - language: go
        paths:
          - data_config.go
          - mapping.go
          - action.go
          - direct_menu.go
          - native.go
          - launch.go
          - script_action.go
          - script_process_unix.go
          - script_process_darwin.go
          - script_process_linux.go
          - script_menu.go
          - script_action_test.go
          - script_process_test.go
          - script_process_darwin_test.go
          - script_process_linux_test.go
          - script_menu_test.go
          - direct_lifecycle_darwin_test.go
          - rex/script_context.go
          - rex/script_context_test.go
          - rex/block.go
          - rex/layer.go
          - rex/direct_popup.go
          - rex/direct_popup_test.go
          - rex/snapshot.go
          - rex/snapshot_test.go
          - rex/window.go
      - language: bash
        paths:
          - script_action_test.go
          - script_process_test.go
          - script_process_darwin_test.go
          - script_process_linux_test.go
          - script_menu_test.go
          - direct_lifecycle_darwin_test.go
          - rex/script_context_test.go
    specialists:
      architecture:
        required: true
        reason: >-
          New executable-config contract, origin-context/process boundary and
          script-only minimized lifetime must preserve public cleanup
          invariants.
      security:
        required: true
        reason: >-
          Trusted configuration gains executable authority; literal argv/path,
          host/account proof, filesystem suitability, detached process and
          no-effect refusal paths.
      usability:
        required: true
        reason: >-
          New configuration workflow, visible startup errors and same-menu
          acknowledgement, host-local path/environment docs and no-success-focus
          behavior.
      maintenance:
        required: false
        reason: >-
          No dependencies, CI, packaging, build commands, Nix or quality
          configuration changes planned; existing environment/config reused.
lifecycle:
  state: planned
  sourceEventID: rex-whichkey-p20-t01/resume/2
  changedAt: '2026-11-24T03:08:04.284Z'
title: Background scripts with exact origin context and reusable-menu startup errors
mtnCreateRetry:
  key: rex-whichkey-p20-task01-create-v1
  intent: 502f156fb519df5a55c0aa240bca15e6f1ba8098842f749d0fd4c7c5e1d2550b
status: in-progress
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-24T02:57:59.652Z'
dateModified: '2026-11-24T03:08:05.314Z'
tags:
  - task
type: task
id: db947c61-8db3-43ac-8b1c-556c4b15be24
projects:
  - '[[phases/rex-whichkey-20-background-scripts.md]]'
---
Implement one coherent end-to-end script feature: strict script.run path/literal-args config; real retained-selection dispatch consumer; host/account/cwd proof and host-local target preflight; Unix detached native process startup; script-only verified minimized-menu/error-only acknowledgement lifecycle; tests and user documentation. Keep existing built-ins on their current exact detach path, standalone menu result-only and launch-only Lua unchanged. Use requirements-design.draft.md and discovery.md as proposal inputs, not evidence of explicit design approval. No parser-only producer or generic process API. Follow focused TDD and preserve existing lifecycle/security/confirmation assertions; final acceptance requires all gates and independent reviews.

## Acceptance
AC1: Strict valid/invalid JSON vectors incl null/unknown/duplicate/case aliases/NUL/Unicode/wrong args, no executable work on load, current defaults/providers/standalone menu unchanged; literal argv round-trip tests.

AC2: Production metadata/dispatch caller uses captured endpoint/session/block; fresh exact host/account/native execution authority and foreground cwd. Table failures for missing/mismatched identity, unreadable cwd, moved original and GUI reconnection produce zero filesystem launch/start effects. Same remote host IDs in fixtures are routing evidence only, not live remote proof.

AC3: Relative/absolute/spaces/metacharacter/Unicode targets; no PATH/tilde/env expansion. Host-local missing/directory/FIFO/device/nonexecutable/dangling/invalid-interpreter targets refuse. OS exec errors remain precise/sanitized; no shell fallback, no retry.

AC4: Isolated real child process survives launcher exit and PTY teardown, has exact cwd/argv, disconnected stdio and separate session; caller returns before fixture completion. No completion wait, menu-context kill, duplicate start or resource leak; builder owns test fixture child cleanup.

AC5: Real direct-menu orchestration plus concrete CLI fixtures proves worker join/TTY restoration/exact minimized same-layer ownership and original focus before one start. ACK without minimized verification, failed minimize, visible state, changed origin, cancellation and cleanup failure suppress start. Success never reveals/focuses menu/script UI, detaches exact owned host and exits. Preserve all prior built-in detach assertions.

AC6: Missing/preflight failure readable in existing menu; late certainly-failed OS startup restores only same layer/block error-only UI for acknowledgement, never executes/retries/reloads config/creates pane. Verify final retirement/focus/TTY and no unintended actions on acknowledgement. Unknown post-start cleanup failure cannot claim no execution/rollback or retry. Prove actual error visibility/invisibility in safe isolated UI evidence, not detached stderr alone.

AC7: Usable JSON example and docs explain host-local placement, shebang/executable requirement, literal vector, trusted-config boundary, host-process not pane-shell environment, discarded stdio/no completion monitor, same-menu error recovery and safe manual recovery. Exact-revision independent review and all task/combined/main gates required.

## Non-goals
No Lua execution/callback registry or generic Rex escape hatch
No daemon/service framework or new dependency
No script copying/install/bootstrap/reachability/SSH fallback
No new script-visible block/window/input injection
No user configuration/reload/deploy/push/production remote trial
No gate/rule/assertion weakening

## Governing authority and planning
User-authorized bounded execution, not separate user assent to every technical detail. Read governing design and the independent assessment .mtn-tmp/p20/architecture/assessment.md from the parent checkout. Preserve existing built-in detachment. Successful script post-start cleanup is detach-only: never re-focus after execution. If native Start fails certainly with no child, restore only the exact owned menu for error acknowledgement and retire it safely; uncertain start is never retried. Docs are covered by Go public-contract and independent usability review. Nontrivial embedded shell fixtures require separate Bash code/test review. No source-only/mock remote routing claim is remote-host certification.
<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAzOjAzOjQ4LjUzOFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiZGlzcGF0Y2giLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMjAtdDAxL2Rpc3BhdGNoLzEiLCJpbnRlbnQiOiJiMjU2Y2E5MWE3ZTY3YzEyMzg4MGQzOWViZDZiMjlmODM3ODgxNmFmNzBlYzMwNDRmZTZmOTEyZjk0YWM0YjE2In0 -->
### 2026-11-24T03:03:48.538Z · dispatch · unspecified

{"attempt":1,"baseCommit":"e5a9ba524b74c84ba3c0d672cec505a985bb91cc","branch":"work/p20-t01-builder-a01","collectionRoot":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.tasks","commands":["task build","task test","task lint","git diff --check"],"commitPolicy":"Source-only commit with unchanged hooks; no .tasks/push/integration","coordinator":"rex-whichkey-p20-coordinator-a01","eventID":"rex-whichkey-p20-t01/dispatch/1","kind":"dispatch","ownership":{"interfaces":["SCRIPT-ACTION","SCRIPT-ORIGIN-CONTEXT","SCRIPT-STARTUP-LIFETIME"],"paths":["data_config.go","mapping.go","action.go","direct_menu.go","native.go","launch.go","script_action.go","script_process_unix.go","script_process_darwin.go","script_process_linux.go","script_menu.go","script_action_test.go","script_process_test.go","script_process_darwin_test.go","script_process_linux_test.go","script_menu_test.go","direct_lifecycle_darwin_test.go","rex/script_context.go","rex/script_context_test.go","rex/block.go","rex/layer.go","rex/direct_popup.go","rex/direct_popup_test.go","rex/snapshot.go","rex/snapshot_test.go","README.md","docs/actions.md","docs/mapping-format.md","docs/launch-contract.md","docs/acceptance.md","examples/background-scripts.json"],"resources":[],"sharedPaths":[]},"readiness":{"facadeLimitation":"ready-phase rejects unchanged flat historical project parent. Read-only projection verifies native workflowKind/project UUID/ref; no migration or scope change.","metadata":".mtn-tmp/p20/scoped-metadata-validation.json","native":".mtn-tmp/p20/native-ready.json","plan":".mtn-tmp/p20/scoped-plan-validation.json"},"recordRef":"tasks/rex-whichkey-20-01-background-scripts.md","resourceAllocation":{"GOMAXPROCS":14,"GOMEMLIMIT":"2GiB","cpu":14,"heavyGates":1,"memoryGiB":3},"runID":"rex-whichkey-p20-t01-builder-a01","schemaVersion":1,"specHash":"sha256:9f2ef6e168a87f200c462f922af39383e525968aca2f0222c2c22e4d6f94d783","targetBranch":"main","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p20-t01-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAzOjA0OjAzLjg3NFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC10MDEtYnVpbGRlci1hMDEvcnVudGltZS9yZXNlcnZlZCIsImludGVudCI6IjI5NmQ0OGVmMjI2NThmOTVkNTE3NjYzMmNkMWIyNDM1NGJlNTQ4NTRhYTU2ZGI0ZGZjOWNjMjM0OWU2NTdiMGYifQ -->
### 2026-11-24T03:04:03.874Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-20-background-scripts.md","runID":"rex-whichkey-p20-coordinator-a01"},"dispatchEventID":"rex-whichkey-p20-t01/dispatch/1","eventID":"rex-whichkey-p20-t01-builder-a01/runtime/reserved","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"03e1c49e-f1a3-4d8e-9f55-7d5f0f4ef38e"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-20-background-scripts.md","piSessionID":"rex-whichkey-p20-t01-builder-a01","piSessionName":"p20-t01-builder-a01","recordRef":"tasks/rex-whichkey-20-01-background-scripts.md","rexSessionName":"rex-whichkey-main","rexWindowName":"p20-t01-builder-a01","role":"builder","runID":"rex-whichkey-p20-t01-builder-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-developer","specHash":"sha256:9f2ef6e168a87f200c462f922af39383e525968aca2f0222c2c22e4d6f94d783","stage":"reserved","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-t01-builder-a01","windowName":"p20-t01-builder-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p20-t01-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAzOjA0OjEwLjA5MloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC10MDEtYnVpbGRlci1hMDEvcnVudGltZS9sYXVuY2hlZCIsImludGVudCI6ImQxOWY4ZWU3ZWM1ODkxY2EwMDgwMThmNDM1NWMyYTc5YWVjM2E3Mjc5YWNmNDMwYTczOWIxMzAyZmZmMmFiYjEifQ -->
### 2026-11-24T03:04:10.092Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-20-background-scripts.md","runID":"rex-whichkey-p20-coordinator-a01"},"dispatchEventID":"rex-whichkey-p20-t01/dispatch/1","eventID":"rex-whichkey-p20-t01-builder-a01/runtime/launched","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"03e1c49e-f1a3-4d8e-9f55-7d5f0f4ef38e"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-20-background-scripts.md","piSessionID":"rex-whichkey-p20-t01-builder-a01","piSessionName":"p20-t01-builder-a01","provider":{"displayName":"p20-t01-builder-a01","name":"rex","resources":{"block":"block:01a20b82-ba3b-7b3e-9929-771e2b703ef7","endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","session":"session:01a20b68-50b0-7c48-81eb-9d5947ecb20d","window":"window:01a20b82-ba4c-76c2-9241-aec958b86aeb"}},"recordRef":"tasks/rex-whichkey-20-01-background-scripts.md","rexBlockID":"block:01a20b82-ba3b-7b3e-9929-771e2b703ef7","rexSessionID":"session:01a20b68-50b0-7c48-81eb-9d5947ecb20d","rexSessionName":"rex-whichkey-main","rexWindowID":"window:01a20b82-ba4c-76c2-9241-aec958b86aeb","rexWindowName":"p20-t01-builder-a01","role":"builder","runID":"rex-whichkey-p20-t01-builder-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-developer","specHash":"sha256:9f2ef6e168a87f200c462f922af39383e525968aca2f0222c2c22e4d6f94d783","stage":"launched","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-t01-builder-a01","windowName":"p20-t01-builder-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p20-t01-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAzOjA3OjMzLjc5M1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYnVpbGQiLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMjAtdDAxL2J1aWxkLzEiLCJpbnRlbnQiOiI1YjIwMmViZGU4MzA5M2RiNjYxMWY2MWM0NWEwODZkODllNDM3NzA4N2I1YzQ5ZWU3N2I0YWYwMmMyMWNiMWJkIn0 -->
### 2026-11-24T03:07:33.793Z · build · unspecified

{"attempt":1,"base":"e5a9ba524b74c84ba3c0d672cec505a985bb91cc","blocker":{"action":"Coordinator: add rex/window.go to implementation ownership and Go review scope; redispatch with updated specHash.","evidence":".mtn-tmp/p20/build-a01/blocker.md","path":"rex/window.go","reason":"Shared layerView discards state; minimal shared minimized-state verification requires editing this file, which exact ownership omits. Architecture assessment section 4 names this seam."},"coverage":[],"dispatchEventID":"rex-whichkey-p20-t01/dispatch/1","eventID":"rex-whichkey-p20-t01/build/1","files":[],"head":"e5a9ba524b74c84ba3c0d672cec505a985bb91cc","integrationCommands":[],"kind":"build","limitations":["Feature unimplemented; no feature acceptance verified.","No configured architecture command found in inspected Taskfile/hooks/lint configurations.","No live feature mutations or remote trials; no source commit or integration."],"outcome":"blocked","recordRef":"tasks/rex-whichkey-20-01-background-scripts.md","redGreen":{"green":[],"reason":"Stopped before source changes at ownership boundary.","red":[]},"runID":"rex-whichkey-p20-t01-builder-a01","schemaVersion":1,"specHash":"sha256:9f2ef6e168a87f200c462f922af39383e525968aca2f0222c2c22e4d6f94d783","summary":"Ownership expansion needed for shared layer-state decoding. Source unchanged; no partial parser-only implementation.","tree":"985f508cb7ebbcc9fbf8ef019f7b5a910dd7d7be","validation":[{"command":"git diff --check","evidence":".mtn-tmp/p20/build-a01/diff-check.log","result":"pass"},{"command":"git status --porcelain=v1","evidence":".mtn-tmp/p20/build-a01/source-identity.txt; source/index clean","result":"pass"},{"command":"task build","evidence":"Not run: ownership blocker before implementation; .mtn-tmp/p20/build-a01/blocker.md","result":"unavailable"},{"command":"task test","evidence":"Unit/race not run: ownership blocker before implementation","result":"unavailable"},{"command":"task lint","evidence":"Not run: ownership blocker before implementation","result":"unavailable"}]}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAzOjA3OjQwLjA5NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVsZWFzZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC10MDEtYnVpbGRlci1hMDEvcmVsZWFzZS9kaXNwYXRjaCIsImludGVudCI6ImMxNzQ5NDUwM2E0M2NmZTZmNGJkNTU3Y2U0MzM0Y2YwMTAxNzI4YzFmYWMxYjM5ZDA4NmZlODNjNDI4NDY0OTkifQ -->
### 2026-11-24T03:07:40.095Z · release · unspecified

{"attempt":1,"dispatchEventID":"rex-whichkey-p20-t01/dispatch/1","eventID":"rex-whichkey-p20-t01-builder-a01/release/dispatch","kind":"release","outcome":"stopped","processEvidence":{"exitCode":0,"exitFile":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-t01-builder-a01/exit-code","resultFile":".mtn-tmp/p20/builder-result-a01.json","resultValidatedBy":"mtn-rex result","source":"clean, no source edits; missing ownership seam","window":"blocked worker window retained under recovery policy"},"recordRef":"tasks/rex-whichkey-20-01-background-scripts.md","runID":"rex-whichkey-p20-t01-builder-a01","schemaVersion":1,"specHash":"sha256:9f2ef6e168a87f200c462f922af39383e525968aca2f0222c2c22e4d6f94d783"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAzOjA3OjQ2LjMzN1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVsZWFzZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC10MDEtYnVpbGRlci1hMDEvcmVsZWFzZS9yZXNlcnZlZCIsImludGVudCI6IjVkMTNmNTZiOTUzNjY0OGJhODBjOWUzZWU0YWRlZTA1ZjI0MWQxMzg4MDE0NDRiZmEzYmI4MWNiOTg5YTMxMjYifQ -->
### 2026-11-24T03:07:46.337Z · release · unspecified

{"attempt":1,"dispatchEventID":"rex-whichkey-p20-t01-builder-a01/runtime/reserved","eventID":"rex-whichkey-p20-t01-builder-a01/release/reserved","kind":"release","outcome":"stopped","processEvidence":{"exitCode":0,"exitFile":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-t01-builder-a01/exit-code","resultFile":".mtn-tmp/p20/builder-result-a01.json","resultValidatedBy":"mtn-rex result","source":"clean, no source edits; missing ownership seam","window":"blocked worker window retained under recovery policy"},"recordRef":"tasks/rex-whichkey-20-01-background-scripts.md","runID":"rex-whichkey-p20-t01-builder-a01","schemaVersion":1,"specHash":"sha256:9f2ef6e168a87f200c462f922af39383e525968aca2f0222c2c22e4d6f94d783"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAzOjA3OjUyLjc5NFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVsZWFzZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC10MDEtYnVpbGRlci1hMDEvcmVsZWFzZS9sYXVuY2hlZCIsImludGVudCI6IjM5Njk3ZmQwMzE0NTJiY2E4NGU5ODhkMTc1MDUwZTQ0OTcyNjMzYzFlYzM0M2JlZWVhZmQxNGUxNDkyNDExMTgifQ -->
### 2026-11-24T03:07:52.794Z · release · unspecified

{"attempt":1,"dispatchEventID":"rex-whichkey-p20-t01-builder-a01/runtime/launched","eventID":"rex-whichkey-p20-t01-builder-a01/release/launched","kind":"release","outcome":"stopped","processEvidence":{"exitCode":0,"exitFile":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-t01-builder-a01/exit-code","resultFile":".mtn-tmp/p20/builder-result-a01.json","resultValidatedBy":"mtn-rex result","source":"clean, no source edits; missing ownership seam","window":"blocked worker window retained under recovery policy"},"recordRef":"tasks/rex-whichkey-20-01-background-scripts.md","runID":"rex-whichkey-p20-t01-builder-a01","schemaVersion":1,"specHash":"sha256:9f2ef6e168a87f200c462f922af39383e525968aca2f0222c2c22e4d6f94d783"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAzOjA4OjAyLjY1N1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVzdW1lIiwicmV0cnlLZXkiOiJyZXgtd2hpY2hrZXktcDIwLXQwMS9yZXN1bWUvMiIsImludGVudCI6IjgzYmZlMGNlOTM2OWU1MjI5MDM0MDQwMTZiNTkwOTg0YzlmY2E2OGUyNTI0ZTY0Mjc0NWI3YzlmMGI0ZTBkMjUifQ -->
### 2026-11-24T03:08:02.657Z · resume · unspecified

{"attempt":2,"authorizedBy":"User bounded execution authority for routine task decomposition and implementation fixes","blockerEventID":"rex-whichkey-p20-t01/build/1","eventID":"rex-whichkey-p20-t01/resume/2","evidence":[".mtn-tmp/p20/builder-result-a01.json",".mtn-tmp/p20/task-patched-a02.json"],"kind":"resume","recordRef":"tasks/rex-whichkey-20-01-background-scripts.md","resolution":"Added rex/window.go to exact ownership and Go review paths; native shared layerView state is required by already-authorized minimization design. Old worker stopped with clean source.","schemaVersion":1,"specHash":"sha256:b124c17007a38f92ed6da08f6b03ade418bef32ef20b8f0354233598bab51abd"}
<!-- mtn:end-entry:v1 -->
