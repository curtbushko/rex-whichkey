---
identity:
  kind: design
  designID: rex-whichkey-lua-float-launch
  uuid: d63c1985-918b-4dd6-bf9d-76d4b6a7fbfc
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
lifecycle:
  state: approved
  sourceEventID: rex-whichkey/lua-float-design/approval/2
  changedAt: '2026-11-07T07:21:15.951Z'
title: Lua-owned float launch and direct Go menu
mtnCreateRetry:
  key: rex-whichkey-lua-float-launch-design-v1
  intent: 55974fe3d4947a2b8557e467d9f34ec313e544ff470c89fadf1b7fddba9c9106
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T04:34:13.083Z'
dateModified: '2026-11-07T07:21:16.580Z'
tags:
  - task
type: task
id: 2eb82453-7283-4f35-9a66-bfd051b0424c
---
# Lua-owned floating launch and direct Go menu — accepted implementation

## Authority, refinements and delivery status

Accepted Phase19 direction REXWHICHKEY-20261106-232754: install Lua for native floating presentation/sizing/launch, Go for menus and palette matching, remove Go controller, make run invokes Lua, implement/review/commit and actually run it. User now explicitly replies "1a - yes, accept it for now; 1b - yes" to the two persisted unblock decisions. Their permanent IDs remain the values recorded in the decision register and blockers; they are not replaced by this document revision.

Current approved refinements take precedence over earlier blocked-candidate wording: (1a) use Rex default native chrome for now, explicitly DEFER native titlebar/border/corners/shadow customization/REQ-026 until supported; (1b) visible float, terminal state, workers and origin focus must be cleaned up/verified before an effect, but the detached execution-host Block may remain during dispatch and retire after Go exits. Such post-exit cleanup cannot suppress an already-issued effect. These are explicit user choices, not silent requirement weakening. Original Open states, design/phase revisions, blocked planner/build and earlier phases remain historical evidence.

This refines launch-only Lua/Go-owned popup creation in designs/rex-whichkey-cli-only-menu.md and designs/rex-whichkey-startup-path.md. All non-superseded target/config/confirmation/CLI/no-retry and palette behavior remain. No code or runtime success is established by approval or this prose.

## Required responsibilities and contracts

- LUA-FLOAT-LAUNCH: package a shared Lua module/action. Normal keybinding and actual make run invoke it without user init edits/reload or a preinstalled global action. Capture initiating session/block/client data before creating anything; resolve its exact window from authoritative placement, not active/focused fallback. Validate configured fractional w/h, center x=(1-w)/2,y=(1-h)/2. Defaults .8/.3 match the observed preview. Use native new_layer with one terminal, shell=none and the fixed rex-whichkey program. Development launch must select the newly built checkout program through a scoped normal PATH prepend, not a stale installed version, arbitrary code nomination or REX_BIN_DIR/Nix fallback.
- DEFAULT-NATIVE-CHROME: current Rex provides no per-float titlebar/border/corner/shadow API. Use its defaults. Lua owns future presentation policy, but add no fake/ignored chrome flags, unsupported API fields, global settings mutation or ANSI border substitute. Terminal theme/color options are not chrome. Req026 customization is approved deferred, not completed. A docs/policy task records that explicit disposition and honest visual limits.
- DIRECT-GO-MENU: one Go menu process is started INSIDE the Lua-created float. Public direct entry/data contract distinguishes original source IDs from its own Rex terminal IDs; no-argument/ambiguous use must never detach the caller's original pane. Go owns validated JSON/defaults/menu/navigation/providers/inline rename/confirm and concrete argument-array Rex CLI effects. No launching Go parent, renderer child, copied image, lease/result/update IPC, renamed supervisor, callback registry, executable Lua menu config, daemon or direct Rex socket.
- INPROCESS-MENU: bounded detached immutable metadata jobs with one UI owner; max two jobs and existing once-per-kind/error/snapshot reuse, limits and current markers. Own-popup Block excluded from pane inventories/counts/targets in post-create origin snapshots; raw view retained only for exact own-float validation. Join/cancel jobs and restore terminal on all exits. Current default keys, JSON contract, hidden/disabled behavior, pagination/fixed rectangle, Unicode and pending rename input protections remain. Go reads only its own PTY dimensions/palette and SIGWINCH for fitting; no Go origin geometry queries for float size. This is not pixel-perfect menu shrink wrapping or a first-size-final guarantee.
- DIRECT-LIFETIME: validate exact own placed Block is the only block in its floating Layer, distinct from retained original source. Go joins workers/restores tty, performs documented session.detach_block on that exact Block, verifies visible layer retirement and detached placement, restores/validates original focus/placement/app-generation, and dispatches a selection exactly once. Any uncertain pre-effect detach/focus/ownership/terminal/join error suppresses effects with no retries. No label/index/current-focus retargeting. Closing the original session may naturally destroy the detached host; never retry an uncertain operation or fabricate status. Lua creates terminal with exit.on_completion=true, quick_exit_threshold_ms=0 so Rex retires host after normal/error exit; actual source-run verification must inspect list_blocks, not just visible layers, for leaks. Go cannot prove its own post-exit deletion; this accepted limit must be documented rather than hidden.
- NORMAL-SOURCE-RUN: actual make run builds checkout source then uses shared Lua action/script; compile/help/dry-run/action-created acknowledgment alone is not display or completion. Current PATH make is a wrapper for Task. Task run routing may change and a conventional Makefile may be added, but build/lint/test/hook standards are protected. Installed normal Go/Lua assets expose the same entry without profile/config activation or executable-location policy in Nix.

## Implementation and review boundaries

Use a coherent current-used Go direct-mode plus Lua launcher vertical slice as the first producer; widen literal ownership to the actual CLI/native/Lua consumers rather than shipping an unused future-only producer. Subsequent task retires obsolete controller production paths and isolates truly historical support test-only while preserving assertions. Producer-before-consumer native dependencies and exact revision-bound reviews/integration remain. Current four task identities are retained with reapproved narrowed responsibilities; no task marked done from a blocked/no-change attempt. No blanket legacy migration or old-phase reopening.

Expected paths: cmd/main.go, direct_menu.go, terminal/menu/provider/text/native/resource files, concrete rex snapshot/block/layer helpers, lua module/example, obsolete popup/runtime transport source and precise test-only support files, Taskfile run/Makefile/package/active docs. No framework/Cobra/Viper/hexagonal rewrite or new dependency. Go CLI availability stays normal execution-host PATH only; do not restore hashing/status/version/schema preflight.

## Verification and preservation

Each code task requires formatting/whitespace, normal task build and task lint:vet, independent combined source review and appropriate real public owned runtime contribution. Final actual make run from the reviewed combined checkout AND rebuilt main must show Lua-created single float/current built Go, palette/own-PTY sizing/resize, static/dynamic menu readiness, bounded loading/rename editing, exact owned benign rename, dismissal/cancellation, pre-effect visible retirement and final detached Block cleanup. Retain exact sources/command/IDs/errors and distinguish functional results from timing/visual claims. No numerical SLA; no physical keybinding/deployment/platform claim without observation. Source-only checks cannot replace actual make run.

User's unit/race/TDD/Lua/parser/FIFO harness/benchmark/broad-experiment prohibition remains; actual normal source run is explicitly authorized. No fake servers/delays, shellHook/nix develop, dependencies/pins/profiles/config edits/reload/restart/remote/deploy/push/global cleanup or deletion of failed resources. Protected tests/assertions/hooks/Taskfile lint/test definitions and ignored lint config remain; inherited full lint62 failure is not green. Separate dirty flake.nix Darwin fix stays out of commits. Existing sole coordinator, isolated leaf worktrees, serial review/integration/heavy slot,14CPU/6GiB shared/Go5GiB and pinned reused tools.

Stable traceability REQ001/002/003/004/006/008/009/010/011/012/013/014/015/016/017/025/026/028/029 uses existing hyphenated IDs. Req026 native customization is explicitly deferred by1a; docs task records the decision, never implementation credit. Req025/028/029 now mean Lua fractional sizing/own-PTY fit/fixed bounds, not Go autosizing. Req003/004 data-only and deferred Req005/hooks unchanged. Full phase completion means all nondeferred scope and actual evidence delivered, with approved native-chrome deferral explicit.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA0OjM0OjE2LjMxOFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9sdWEtZmxvYXQtZGVzaWduL2FwcHJvdmFsLzEiLCJpbnRlbnQiOiIwZGM1NTI3ZGU1OTA1YzQxMDllMTQzNzRhNGJjZWMxM2M2MWNhMWMzMzQwZjA1NDlmOTgwMjUyYmM3MzcxZGVkIn0 -->
### 2026-11-07T04:34:16.318Z · approval · unspecified

{"approvalScope":"Accepted responsibility/public-launch requirements and bounded execution/task decomposition. Native chrome remains required but unsupported/blocked, NOT deferred. Documented detach-before-effect is an engineering candidate requiring review and real verification; no material safeguard weakening or external Rex/config changes approved.","approvedBy":"user","attempt":1,"decision":"approved","decisionEvidence":"User explicitly requests Lua native float style and size, Go menus/palette matching, removal of current controller, make run through Lua, create phase/tasks, implement/commit and actually make run.","decisionID":"REXWHICHKEY-20261106-232754","eventID":"rex-whichkey/lua-float-design/approval/1","kind":"approval","recordRef":"designs/rex-whichkey-lua-float-launch.md","schemaVersion":1,"specHash":"sha256:a3657116206eb455b158b9727099b7c3859882d18f759171205eb830dda163c8"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA3OjIxOjE1LjIwNVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9sdWEtZmxvYXQtZGVzaWduL2FwcHJvdmFsLzIiLCJpbnRlbnQiOiJjZTg0ZTNkZWU1NjMwYjIyZGVkNDIxNmNjMTQ3ZjNlMmU0ZGQ2MTE0ZGMyMTk3M2NjZDczNDg1ZGRkZDgwYzVlIn0 -->
### 2026-11-07T07:21:15.205Z · approval · unspecified

{"approvalScope":"Explicit native chrome deferral/use defaults and detached execution-host through dispatch/post-exit retirement exception. Original Phase19 create/tasks/implement/commit/actual make run authority retained; routine vertical task replan authorized. No other guards or test/config/deploy limits waived.","approvedBy":"user","attempt":2,"decision":"approved","decisionEvidence":"1a - yes, accept it for now; 1b - yes","eventID":"rex-whichkey/lua-float-design/approval/2","kind":"approval","recordRef":"designs/rex-whichkey-lua-float-launch.md","schemaVersion":1,"specHash":"sha256:f7e38d49287de319753732b09fe472a40dfadcef81e5245e2deac208e968ab70"}
<!-- mtn:end-entry:v1 -->
