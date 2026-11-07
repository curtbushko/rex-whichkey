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
  sourceEventID: rex-whichkey/lua-float-design/approval/1
  changedAt: '2026-11-07T04:34:17.016Z'
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
dateModified: '2026-11-07T04:34:17.634Z'
tags:
  - task
type: task
id: 2eb82453-7283-4f35-9a66-bfd051b0424c
---
# Lua-owned floating launch and direct Go menu

## Authority

REXWHICHKEY-20261106-232754 records the user's explicit request: installed Lua handles native float style (title bar, border, corners, shadow), calculates float size/placement; Go handles menus and reads/matches terminal colors; remove the existing Go controller behavior; make run calls Lua to launch the binary; create phase/tasks, implement and commit, and actually run make run. ID assigned at capture, not an exact earlier message timestamp. Requirements and bounded execution/planning are accepted; unavailable Rex features are not assumed approved deferrals or implemented.

This supplements designs/rex-whichkey-cli-only-menu.md and designs/rex-whichkey-startup-path.md. It replaces their launch-only Lua/Go-owned popup creation and parent-renderer/image-copy mechanism for this scope. It does NOT restore Go-to-Lua callbacks, configuration evaluation, custom action registries/relay receipts or a framework. Lua is the native floating launcher; Go menu data remains JSON and all selected application Rex effects remain argument-array CLI operations.

## Accepted responsibility and interfaces

1. LUA-FLOAT-LAUNCH: ship Lua with the package; the normal binding and make run use this same implementation without loading/reloading user init. Lua retains initiating full session/block/client intent before popup creation, resolves the origin's actual window (not current-focus fallback), validates finite fractional width/height and computes centered x=(1-w)/2,y=(1-h)/2. Existing preview used w=.8,h=.3; these are reasonable configurable defaults, not a claim of pixel-tight shrink wrapping. Explicit whole-window fractions replace Go origin-geometry/natural-height creation policy; Go adapts to its actual PTY and initial resize. No global settings mutation.
2. DIRECT-GO-MENU: native new_layer starts the installed/built rex-whichkey directly with shell=none and explicit original-origin data. The terminal's own IDs name only its owned popup execution block, never replace the original action target. One Go menu process at startup: no Go launching controller, renderer child, self executable copy, private result/lease/update files or IPC bridge. Async providers/rename are bounded in-process UI-owner updates, not renamed controller logic. Preserve current defaults/JSON/data/action behavior, limits, full IDs/current markers, confirmations, pagination/fixed rectangle, text editing, once-only dispatch and no uncertain retries.
3. TERMINAL-COLOR-MATCH: Go continues reading its own terminal palette and dimensions, applying current semantic theme/readability and redraw-on-resize. Native chrome is NOT ANSI colors, Go title/footer or painted fake borders. Do not confuse options.theme with header/border/corner/shadow control.
4. DIRECT-LIFETIME: examine the documented session.detach_block route before removing correctness safeguards. It removes the exact owned popup from its window without ending Go; a sole-block floating layer closes, so the same restored/joined Go process can dispatch outside the visible popup, then exit. Terminal options exit.on_completion=true and quick_exit_threshold_ms=0 are documented mechanisms for retiring its detached process block even on quick failure. Preserve original placement/app-generation checks, restore terminal/cancel+join workers BEFORE detach/effects, validate only its exact owned block/layer, avoid closing unrelated windows, and refuse dispatch after uncertain detach/focus. A detached execution host remains only until Go exits, unlike the removed two-process controller; validate that terminal colors/layout/lifetime are really safe on supported platforms. This is an engineering candidate requiring source/lifecycle review and real make run validation, NOT already verified behavior. If exact accepted cleanup guarantees cannot be met, record the concrete material conflict instead of silently weakening them or adding another supervisor under a new name.
5. NORMAL-MAKE-LAUNCH: actual make run must build/use the current checkout executable and invoke Lua through the documented Rex script/action entry. It must not invoke the old Go popup launcher, use an installed stale binary, depend on user init action registration/reload, or claim a dry-run/help/compile as menu verification. Existing PATH make is a three-line wrapper execing Task; no Makefile currently exists. Updating only Task's run route is possible, but a conventional Makefile may be added for GNU make portability. Keep build/lint/test targets and all assertions/standards intact. Any development PATH prepend is explicit scoped launch data using normal PATH, not new REX_BIN_DIR compatibility policy or Nix runtime wiring.

## Confirmed native-chrome blocker

Rechecked installed session.new_layer schema accepts only session/window/bounds/layout/focus. Terminal creator options cover command/shell/cwd/size/theme/exit/etc, but not native title-bar visibility, border, corners or shadow. options.theme is terminal colors only. Community has active native chrome feature requests. Thus full LUA-FLOAT-CHROME is a REQUIRED, currently BLOCKED deliverable. Do not implement unused/ignored chrome flags, pass guessed schema fields, change global Rex settings, draw fake chrome, edit Rex externally, or mark the criterion complete/deferred. Track a distinct task/blocker; continue genuinely independent supported work. Full completion needs an actual supported Rex capability/version or an explicit user decision to defer/change native-chrome scope. Older Phase11 chrome blocker/history remains untouched, not retrospectively completed.

## Requirements, verification and preservation

Preserve stable REQ001/002/003/004/006/008/009/010/011/012/013/014/015/016/017 plus complementary REQ025/026/028/029 presentation traceability. REQ003/004 remain data-only; executable functions/hooks REQ005 remain deferred. Req026 maps to the genuinely blocked native-chrome task. Req025/028/029 now use Lua fractional sizing and Go own-PTY fitting rather than Go origin geometry or a controller. Old phase delivery records are history.

Independent revision-bound source reviews, unchanged static/build/vet gates, actual normal make run from a real explicitly owned Rex origin and observed popup/menu/palette/keyboard/dynamic provider/owned benign effect/exit/popup plus detached-block cleanup are mandatory for runtime delivery. Verify the correct built binary and actual Lua path with exact source/command/Rex IDs; record detached/block/layer/focus transitions and failure limits. Test rapid cancel and normal dismissal without fake servers/delays or broad experiments. Existing unit/race/TDD/Lua/parser/FIFO harness/benchmark prohibition stands; real make run is expressly authorized. No user config/profile edits, installation/activation/reload/restart, remote trials, push, lock updates, historical/failed-resource deletion or global cleanup. Existing separate dirty flake.nix Darwin fix stays out of phase commits. Preserve historical assertions; obsolete implementation may be isolated as genuine test-only support where necessary, never production stubs. Do not weaken hooks/Taskfile quality/lint configs; inherited lint62 failures remain honestly reported.

Expected source surfaces: lua launcher/examples/package wiring; cmd/main.go/launch.go/native*.go; popup/runtime/terminal/provider/menu/text-entry files and concrete rex helpers; Taskfile run/Makefile, active docs. Tasks have literal ownership and serialized producer/consumer contracts under the existing sole coordinator. No second collection/session/coordinator. Resource caps one builder/reviewer/heavy gate/integration,14CPU/6GiB total with Go memory5GiB, pinned reused environment. Task Manager must assess the direct-lifetime candidate before code; routine safe implementation details do not require reapproval, material safety/public-contract changes do.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA0OjM0OjE2LjMxOFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9sdWEtZmxvYXQtZGVzaWduL2FwcHJvdmFsLzEiLCJpbnRlbnQiOiIwZGM1NTI3ZGU1OTA1YzQxMDllMTQzNzRhNGJjZWMxM2M2MWNhMWMzMzQwZjA1NDlmOTgwMjUyYmM3MzcxZGVkIn0 -->
### 2026-11-07T04:34:16.318Z · approval · unspecified

{"approvalScope":"Accepted responsibility/public-launch requirements and bounded execution/task decomposition. Native chrome remains required but unsupported/blocked, NOT deferred. Documented detach-before-effect is an engineering candidate requiring review and real verification; no material safeguard weakening or external Rex/config changes approved.","approvedBy":"user","attempt":1,"decision":"approved","decisionEvidence":"User explicitly requests Lua native float style and size, Go menus/palette matching, removal of current controller, make run through Lua, create phase/tasks, implement/commit and actually make run.","decisionID":"REXWHICHKEY-20261106-232754","eventID":"rex-whichkey/lua-float-design/approval/1","kind":"approval","recordRef":"designs/rex-whichkey-lua-float-launch.md","schemaVersion":1,"specHash":"sha256:a3657116206eb455b158b9727099b7c3859882d18f759171205eb830dda163c8"}
<!-- mtn:end-entry:v1 -->
