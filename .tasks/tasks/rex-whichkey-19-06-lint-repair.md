---
identity:
  kind: task
  taskID: rex-whichkey-19-06-lint-repair
  uuid: 9be113f9-7278-4e40-a907-e2ecf9703d6f
structure:
  order: 6
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
      - action.go
      - cmd
      - config.go
      - configured_rename_test.go
      - configured_rename_wire_test.go
      - controller_legacy_support_test.go
      - data_config.go
      - defaults.go
      - direct_menu.go
      - hosts.go
      - launch.go
      - launch_legacy_support_test.go
      - launch_output_test.go
      - layout.go
      - lua_context_test.go
      - mapping.go
      - menu.go
      - menu_updates.go
      - native.go
      - native_app_owner.go
      - native_resources.go
      - native_sessions.go
      - native_targets.go
      - popup.go
      - popup_legacy_support_test.go
      - popup_updates.go
      - providers.go
      - render.go
      - render_entry.go
      - rex
      - terminal.go
      - terminal_input.go
      - terminal_keys.go
      - terminal_keys_test.go
      - terminal_modes.go
      - terminal_poll_darwin.go
      - terminal_poll_darwin_test.go
      - terminal_poll_linux.go
      - terminal_poll_test.go
      - terminal_setup.go
      - text_entry.go
      - theme.go
      - theme_palette.go
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
      required: true
lifecycle:
  state: planned
  sourceEventID: null
  changedAt: '2026-11-07T16:00:06.730Z'
title: Repair full lint pipeline and preserve tested behavior
mtnCreateRetry:
  key: rex-whichkey-19-06-lint-repair-v1
  intent: f58c2b9e295931b61cb893b74ab9a8501469e4474ad5fc5e51d260e4495bb7c8
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T16:00:07.807Z'
dateModified: '2026-11-07T16:00:09.893Z'
tags:
  - task
type: task
id: 156d2b1f-0a29-4d31-bdc9-e983bf627fb9
projects:
  - '[[phases/rex-whichkey-19-lua-float-launch.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-19-05-test-repair.md]]'
    reltype: FINISHTOSTART
---
After delivered Task05 test repair, fix ALL current full task lint diagnostics by code changes, not rules. Baseline53 covers contextcheck/goconst/gocyclo/govet-shadow/nestif/prealloc/revive/staticcheck; exact current log .mtn-tmp/p19/post-delivery-gates/lint.log, don't assume fixed count or inherited-only. Once golangci is green run lint:ai too as part of full task lint, fix any further real findings. User explicitly always tests/lint after changes and fix the problems: no diagnostic-only exemption at final acceptance.

TDD/regression checks mandatory: preserve real terminal input parser behavior (UTF8, escape, paste, timeouts/key handling), raw-mode restoration/error context, menu/render/palette/readiness/rename edits, bounded metadata/retained exact original/own/destination/GUI-generation validation/confirm/no retries/lifecycle. Split complex code into meaningful private functions with actual production callers, not token helpers, interfaces, compatibility stubs, new framework/Cobra/Viper/hexagonal layers or dependencies. Constants/comments/shadow/error strings corrected consistently without altering user-facing contract silently; keep strong assertions and add regression tests for behavioral changes. Archived historical support remains test-only. Defaults/native chrome1a deferral and detached host1b rule settled, don't reopen.

Required task build, task test (unit+race), FULL task lint all GREEN after source changes and on final head. Also architecture check only if existing config. No editing/suppressing lint config/Taskfile/hook/test checks, no deleted/skipped/weakened assertions, no count-as-pass on tool failure. Existing automated sandbox fixtures/mocks/parser/FIFO tests and TDD authorized by user's correction; no live leaf Rex changes/profile/config/install/reload/remote/deploy/push/benchmark/global cleanup/new deps/pins. Coordinator can resolve concrete toolchain mismatch using permitted existing tools, not weaken checks. Separate dirty flake.nix preserved.14CPU/6GiB shared/Go5GiB, one gate; pinned existing Go1.26.8/lint2.15, no nix develop/shellHook.

Literal ownership root Go/Go tests and rex/cmd Go files as listed. No unrelated docs/legacy migration. Report exact extra path before edit if needed. Fresh independent exact revision combined Go review must run ALL tests/full lint, zero blockers/no test exceptions, then coordinator normal combined/main gates and ACTUAL make run/menu/cancel/owned effect/cleanup. Commit owned source hooksPath=/dev/null; return actual native build commands/results/source IDs and caller list. No green claim until all pipeline stages pass; no unsupported chrome/profile/platform/global coverage claim.
