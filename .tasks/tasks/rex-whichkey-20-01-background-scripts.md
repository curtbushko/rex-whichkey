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
  state: complete
  sourceEventID: rex-whichkey-p20-t01/summary/6
  changedAt: '2026-11-24T06:36:08.369Z'
title: Background scripts with exact origin context and reusable-menu startup errors
mtnCreateRetry:
  key: rex-whichkey-p20-task01-create-v1
  intent: 502f156fb519df5a55c0aa240bca15e6f1ba8098842f749d0fd4c7c5e1d2550b
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-24T02:57:59.652Z'
dateModified: '2026-11-24T06:36:08.999Z'
tags:
  - task
type: task
id: db947c61-8db3-43ac-8b1c-556c4b15be24
projects:
  - '[[phases/rex-whichkey-20-background-scripts.md]]'
completedDate: '2026-11-24'
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

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDA2OjM2OjA3LjE2NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC10MDEvc3VtbWFyeS82IiwiaW50ZW50IjoiMzk1OGU4NDhmOGI5YTY5MTAzZGEyOGEyZWU1ZjliZDVkODJiZTNmNzVhYWU3NjE1MDg4NzEzZDRlZWJlNTcwYyJ9 -->
### 2026-11-24T06:36:07.165Z · summary · unspecified

{"acceptedScope":"REQ-SCRIPT-001..004: strict script.run path/literal argv, proven origin host/account/current cwd, detached background lifetime without a script pane/tab, missing/unsuitable/startup diagnostics in the same menu, compatibility regressions and user documentation/example.","approvalSummary":["User explicitly authorized background execution wherever the origin pane lives, current cwd, existence error, then phase/tasks/implementation; bounded routine design/reviews/integration/cleanup authorized. No separate user agreement to every mechanism is claimed.","Current task/phase manifest binds specHash and governing design/quality-gate fingerprints; scope correction added rex/window.go without changing accepted behavior.","Seven independent required current build6 scopes approved the exact reviewed head/tree: go:code, go:test, bash:code, bash:test, architecture, security, usability. Fresh language verdicts and evidenced current specialist attestations; same authenticated model family in independent contexts.","Source implemented and fixed by delegated Go developers. Build6 revalidated byte-identical build5 source after a confirmed stopped launcher-gate failure; no coordinator feature implementation or stale approval waiver."],"attempt":6,"buildEventID":"rex-whichkey-p20-t01/build/6","cleanup":{"allTaskWorktreesRemoved":true,"branchesAndTransportRetained":true,"eligibleOwnedCleanupComplete":true,"facadeCompatibility":"Private copy fixes invalid jq all-generator syntax while preserving every runtime-reservation release check; positive actual evidence and negative incomplete-release controls passed. Installed helper and repository implementation unchanged.","integrationWorktreeRemoved":true,"retainedOwnedWindows":[{"host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"03e1c49e-f1a3-4d8e-9f55-7d5f0f4ef38e"},"reason":"Stopped earlier blocked/negative/startup-failed attempt; retained by recovery policy","rexBlockID":"block:01a20b82-ba3b-7b3e-9929-771e2b703ef7","rexWindowID":"window:01a20b82-ba4c-76c2-9241-aec958b86aeb","runID":"rex-whichkey-p20-t01-builder-a01","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-t01-builder-a01"},{"host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"03e1c49e-f1a3-4d8e-9f55-7d5f0f4ef38e"},"reason":"Stopped earlier blocked/negative/startup-failed attempt; retained by recovery policy","rexBlockID":"block:01a20bab-5dca-7e34-8c07-13c66e41bc18","rexWindowID":"window:01a20bab-5ddb-7835-96e5-d3075e86c9f9","runID":"rex-whichkey-p20-t01-go-code-reviewer-a02","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-t01-go-code-reviewer-a02"},{"host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"03e1c49e-f1a3-4d8e-9f55-7d5f0f4ef38e"},"reason":"Stopped earlier blocked/negative/startup-failed attempt; retained by recovery policy","rexBlockID":"block:01a20bab-965d-774f-b1ba-ac3070bd6342","rexWindowID":"window:01a20bab-9668-7d0b-b09a-1f85d1b91f84","runID":"rex-whichkey-p20-t01-architecture-reviewer-a02","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-t01-architecture-reviewer-a02"},{"host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"03e1c49e-f1a3-4d8e-9f55-7d5f0f4ef38e"},"reason":"Stopped earlier blocked/negative/startup-failed attempt; retained by recovery policy","rexBlockID":"block:01a20bbd-246f-73ad-b8a4-136af0eb61c0","rexWindowID":"window:01a20bbd-247b-78b6-a7d8-19a95d10e835","runID":"rex-whichkey-p20-t01-go-code-reviewer-a03","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-t01-go-code-reviewer-a03"},{"host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"03e1c49e-f1a3-4d8e-9f55-7d5f0f4ef38e"},"reason":"Stopped earlier blocked/negative/startup-failed attempt; retained by recovery policy","rexBlockID":"block:01a20bd9-f43d-7399-9d82-02490bb3d5e3","rexWindowID":"window:01a20bd9-f448-76f1-9c51-1ad09a0b0821","runID":"rex-whichkey-p20-t01-usability-reviewer-a04","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-t01-usability-reviewer-a04"},{"host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"03e1c49e-f1a3-4d8e-9f55-7d5f0f4ef38e"},"reason":"Stopped earlier blocked/negative/startup-failed attempt; retained by recovery policy","rexBlockID":"block:01a20c01-ae1f-7693-b785-1df2f056d7cb","rexWindowID":"window:01a20c01-ae29-78ca-857a-bf6f29dc56b2","runID":"rex-whichkey-p20-t01-bash-test-reviewer-a05","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-t01-bash-test-reviewer-a05"}],"rexSessionID":"session:01a20b68-50b0-7c48-81eb-9d5947ecb20d","successfulClosedRunCount":26,"successfulWorkerWindowsClosed":true,"totalTaskRuns":32},"eventID":"rex-whichkey-p20-t01/summary/6","files":["README.md","data_config.go","direct_lifecycle_darwin_test.go","direct_menu.go","docs/acceptance.md","docs/actions.md","docs/launch-contract.md","docs/mapping-format.md","examples/background-scripts.json","rex/block.go","rex/direct_popup.go","rex/layer.go","rex/script_context.go","rex/script_context_test.go","rex/snapshot_test.go","rex/window.go","script_action.go","script_action_test.go","script_menu.go","script_menu_test.go","script_process_darwin.go","script_process_linux.go","script_process_test.go","script_process_unix.go",".mtn-tmp/p20/integration-accepted-a06.json",".mtn-tmp/p20/integration-a06/native-scenarios.json",".mtn-tmp/p20/cleanup/task-completed.json",".mtn-tmp/p20/cleanup/facade-compat.diff"],"integrationEventID":"rex-whichkey-p20-t01/integration/6-accepted","kind":"summary","limitations":["No live remote-host or Linux runtime certification; routing/foreign-host refusal is fixture-tested and Linux compilation is recorded separately. Scripts and interpreters must already exist on the pane host; no automatic copying or local fallback.","Native Rex wait reported exit0 after acknowledged error screens; external action-error status propagation is not certified. Actual UI cause/recovery/no-effect behavior and Go error-return regressions are verified; async rex do returns created IDs rather than eventual action status.","Actual local Rex API, terminal buffers, PTYs and OS children were exercised, not physical-key or human visual GUI research.","Six confirmed-stopped earlier blocked/negative/startup-failed worker windows are intentionally retained by recovery policy. All eligible task/integration worktrees and successful worker windows are removed; source branches, transport and Pi evidence are retained."],"recordRef":"tasks/rex-whichkey-20-01-background-scripts.md","requirements":["REQ-SCRIPT-001","REQ-SCRIPT-002","REQ-SCRIPT-003","REQ-SCRIPT-004"],"reviewEventIDs":["rex-whichkey-p20-t01/review/go-code/6","rex-whichkey-p20-t01/review/go-test/6","rex-whichkey-p20-t01/review/bash-code/6","rex-whichkey-p20-t01/review/bash-test/6","rex-whichkey-p20-t01/review/architecture/6","rex-whichkey-p20-t01/review/security/6","rex-whichkey-p20-t01/review/usability/6"],"reviewedHead":"544941280affe3ded577c399feab1a2911dc6396","reviewedTree":"9a2c09e067c5b66322c179ec4ef7b7a24d50a3fa","schemaVersion":1,"specHash":"sha256:b124c17007a38f92ed6da08f6b03ade418bef32ef20b8f0354233598bab51abd","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"ddd77484316bca0af63173b80bc912ec9a8c2648","targetTree":"9a2c09e067c5b66322c179ec4ef7b7a24d50a3fa","verification":[{"command":"main task build","evidence":".mtn-tmp/p20/main-gates/build.log","result":"pass"},{"command":"main task test (unit and race)","evidence":".mtn-tmp/p20/main-gates/test.log","result":"pass"},{"command":"main task lint (vet, golangci and AI stages)","evidence":".mtn-tmp/p20/main-gates/lint.log","result":"pass"},{"command":"reviewed/source integration exact Git tree comparison and main ancestry","evidence":".mtn-tmp/p20/integration-a06/tree-comparison.log and integration-accepted-a06.json","result":"pass"},{"command":"native local Rex background cwd/literal argv/child survival/release completion/no extra pane or tab","evidence":".mtn-tmp/p20/integration-a06/live/success-observed/steps.log","result":"pass"},{"command":"native missing-script error cause/recovery/same-menu/no-effect/retirement","evidence":".mtn-tmp/p20/integration-a06/live/missing-observed/accepted-checks.json","result":"pass"},{"command":"native startup failure: actual 64x5 terminal, visible cause/recovery, same layer, no script body effect and retirement","evidence":".mtn-tmp/p20/integration-a06/live/startup-observed/steps.log and terminal-size.json","result":"pass"},{"command":"owned task cleanup with all runtime releases, stopped-process evidence, closed-successful-window checks and non-forced worktree removal","evidence":".mtn-tmp/p20/cleanup/task-completed.json, runtime-guard-actual.json and retained-windows.json","result":"pass"},{"command":"private cleanup facade compatibility shellcheck/shfmt and complete/incomplete runtime-release guard controls","evidence":".mtn-tmp/p20/cleanup/facade-compat.diff and guard-*-result.json; installed tool unchanged","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
