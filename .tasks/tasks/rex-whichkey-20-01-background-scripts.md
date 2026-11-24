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
    sharedPaths: []
    resources: []
    interfaces:
      - SCRIPT-ACTION
      - SCRIPT-ORIGIN-CONTEXT
      - SCRIPT-STARTUP-LIFETIME
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
lifecycle:
  state: planned
  sourceEventID: null
  changedAt: '2026-11-24T02:57:58.553Z'
title: Background scripts with exact origin context and reusable-menu startup errors
mtnCreateRetry:
  key: rex-whichkey-p20-task01-create-v1
  intent: 502f156fb519df5a55c0aa240bca15e6f1ba8098842f749d0fd4c7c5e1d2550b
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-24T02:57:59.652Z'
dateModified: '2026-11-24T02:58:00.722Z'
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