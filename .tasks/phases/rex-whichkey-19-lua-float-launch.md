---
identity:
  kind: phase
  phaseID: rex-whichkey-19-lua-float-launch
  uuid: 0882c39f-bcb7-409b-b7c6-b0237ef9b788
structure:
  order: 19
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
    - REQ-026
    - REQ-028
    - REQ-029
  documents:
    - ref: designs/rex-whichkey-lua-float-launch.md
      hash: 'sha256:a3657116206eb455b158b9727099b7c3859882d18f759171205eb830dda163c8'
lifecycle:
  state: incomplete
  sourceEventID: rex-whichkey-p19/execution-authority/1
  changedAt: '2026-11-07T04:34:23.494Z'
title: 'Phase 19: Lua-owned float and direct Go menu'
mtnCreateRetry:
  key: rex-whichkey-phase19-lua-float-v1
  intent: b58d8a34b85ced0cf7714d8ed6df8f98abfe3daa01a6b3a77400140e6c58deda
status: in-progress
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T04:34:18.916Z'
dateModified: '2026-11-07T04:34:24.736Z'
tags:
  - task
type: task
id: 1ce0bc56-e223-4a35-8ea4-477e9a28fb3f
projects:
  - '[[projects/rex-whichkey.md]]'
---
# Phase 19 — Lua-owned float and direct Go menu

## Authority and goal

User explicitly requests creation, task decomposition, implementation, commit on completion and actual make run. Accepted direction REXWHICHKEY-20261106-232754: Lua owns float style/size/placement/launch; Go owns menus, terminal palette matching and CLI effects; remove the current Go launching controller. Current main19b76e5, Phase17/18 delivered, separate uncommitted flake.nix Darwin fix preserved.

Governing supplement designs/rex-whichkey-lua-float-launch.md captures exact intent, the supported native detach-block lifetime candidate and the unresolved chrome dependency. This is execution authority inside accepted constraints, not permission to invent unsupported Rex appearance controls, silently defer native chrome or weaken cleanup. Use the existing sole go-team coordinator and mixed-language task skills; no new session/coordinator or legacy migration. Order19 under projects/rex-whichkey.md, no whole-phase prerequisites on older unfinished history.

## Acceptance criteria

1. **Installed Lua float owner.** Shared installed Lua/action implementation computes validated configured fractional width/height and centered x/y, creates one native floating layer/terminal in the exact initiating window and starts installed/built Go directly. Original session/block/client intent survives separately from the popup's own IDs. No focused-window fallback, copied executable, Go popup-launch controller, nested popup or Lua menu/callback framework.
2. **Native chrome ownership, genuinely implemented.** Lua config controls native Rex title bar, border, corners and shadow through verified supported per-float APIs. This is not terminal colors or Go-drawn header/footer. Currently BLOCKED by installed Rex capabilities. No ignored knobs, fake ANSI substitutes, global settings changes or unsupported fields; no completion/deferral without real support or explicit later user change.
3. **Go menu/operations/palette.** One direct Go menu process preserves JSON/defaults/navigation/providers/rename/confirmation/limits/full targets and current terminal-palette matching. UI changes remain single-writer, workers bounded/immutable/canceled/joined; fits own PTY and resizes, no Go origin-geometry round trips for float sizing. Resource operations still use concrete Rex CLI argument arrays, PATH-only binaries, exact endpoint/app generation and no uncertain retries.
4. **Controller-free safe lifecycle.** Remove ordinary production Go launcher/renderer controller, self image capture/copy, private lease/result/update IPC. Verify documented detach_block can retire the exact visible popup while the same restored Go process dispatches once, followed by Rex terminal exit cleanup of its detached host; no leaked block/layer/private files or background supervisor. Keep original resource/freshness/cleanup safeguards. If not achievable without a material change, record the exact blocker rather than weaken ordering or disguise a new controller.
5. **One real make run path.** Actual make run builds/uses checkout Go and calls the shared Lua entry without user configuration action registration/reload. Installed keybinding and development command share the launch logic. Inspect the real PATH make-to-Task wrapper and support conventional Makefile use if needed. Only run/build routing explicitly changes; preserve every test/lint/hook gate and assertion.
6. **Actual runtime verification.** On reviewed combined source actually invoke make run in an explicitly owned real Rex origin, observe the resulting direct menu/one popup/current binary, keyboard/palette/dynamic readiness/inline benign owned rename, normal dismissal/cancel, terminal/focus and visible-layer plus detached-host cleanup. Compile/help/dry-run/private ACK is insufficient. Preserve user state; no broad experiments, prohibited test/harness suites or synthetic delays. Source+actual evidence and visual native-chrome verification are distinct.
7. **Honest packaging/docs/history.** Install normal Go/Lua assets, update examples/active launch/migration/acceptance docs with reviewed revision and actual results/limits, no init takeover or runtime policy in Nix. Native task/phase delivery reflects partial versus complete truth. Preserve prior decisions/failed runs, old tests/fixtures and unrelated flake.nix. Commit reviewed delivered changes; no push/activation/reload/global cleanup.

## Planning and independent work

Task Manager proposes a small serial Go/launcher/runtime graph with literal ownership and all seven criteria/requirements covered, plus a separate genuinely blocked native-chrome task. Chrome must remain in accepted required scope; do not fabricate an implementation or mark full phase done merely because other work succeeds. Supported work may proceed only where safe with the direct-lifetime requirements; inspect documented safe alternatives before escalating.

The installed session.new_layer/terminal-creator schemas expose no native chrome knobs. A supported Rex capability or explicit chrome deferral/change is needed for full phase completion. Source evidence: .mtn-tmp/p19/rex-api-list.json, rex-block-types.json, terminal-options-summary.json. API session.detach_block explicitly keeps a detached block running and removes a sole-block float; terminal exit.on_completion true plus quick_exit_threshold_ms0 are documented. These are inspected capabilities, not a runtime pass.

## Scope and constraints

Trace REQ001/002/003/004/006/008/009/010/011/012/013/014/015/016/017/025/026/028/029 using existing hyphenated stable IDs. Older callback REQ005 stays deferred. No old Phase11/14/15 reopening or retroactive closure. One builder/reviewer/heavy gate/integration;14CPU/6GiB total, Go memory5GiB, pinned existing Go1.26.8/lint2.15. No unit/race/TDD/Lua/parser/FIFO harnesses/benchmarks, nix develop mutating shellHook, dependency/profile/input changes, user config edits/reload/restart, remote trials/deploy/push/failed-resource deletion. Current request explicitly authorizes actual make run and owned runtime verification, not those broader actions.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA0OjM0OjIyLjc2OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMTkvZXhlY3V0aW9uLWF1dGhvcml0eS8xIiwiaW50ZW50IjoiODFjNzk2YjlkYjcyZjBlZmQ4MDYzZmMwMTc2NGFjMDg5OWIwNjJlZWY0YzIzNmIyNjdiMzVlNGIyYzJiM2M5MiJ9 -->
### 2026-11-07T04:34:22.769Z · approval · unspecified

{"approvalScope":"Phase definition, bounded task planning, source execution/fixes/reviews/integration and actual normal make run with owned resources. No implicit native-chrome deferral, material safety change, hidden unsupported controls, broader tests/benchmarks/user config/deploy/reload/migration authority.","approvedBy":"user","attempt":1,"decision":"approved","decisionEvidence":"Create a phase for this, break it down into tasks, implement it and commit the changes when complete. Make sure to actually make run rex-whichkey to test it.","decisionID":"REXWHICHKEY-20261106-232754","eventID":"rex-whichkey-p19/execution-authority/1","kind":"approval","recordRef":"phases/rex-whichkey-19-lua-float-launch.md","schemaVersion":1,"specHash":"sha256:653fe55ccda83b83583ec44188e37e3dba27dee457e72f4b002df84d592b921e"}
<!-- mtn:end-entry:v1 -->
