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
      - lua_context_test.go
      - tests/lua
      - direct_menu_test.go
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
  sourceEventID: null
  changedAt: '2026-11-07T16:00:03.531Z'
title: Repair faithful archived fixture test setup
mtnCreateRetry:
  key: rex-whichkey-19-05-test-repair-v1
  intent: aba3cde7720dd26e51fe585bb2f5d75139d1ebbba8ef10be7d96727ff344a877
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T16:00:04.596Z'
dateModified: '2026-11-07T16:00:05.662Z'
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
