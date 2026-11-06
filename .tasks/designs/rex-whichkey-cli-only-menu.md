---
identity:
  kind: design
  designID: rex-whichkey-cli-only-menu
  uuid: d1e5dec4-c459-4a61-8c96-1e88daba1b39
scope:
  requirements:
    - REQ-001
    - REQ-002
    - REQ-003
    - REQ-004
    - REQ-006
    - REQ-007
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
  documents:
    - ref: prds/rex-whichkey.md
      hash: 'sha256:97614db611a6f35a44d0db32ccf2bf4d4854d6ffea3dfd5375959e59259f01ff'
    - ref: prds/rex-whichkey-usability.md
      hash: 'sha256:966ef70c264360b4820e5faf510496166e91f0effeae479ab7400a3792de88e9'
lifecycle:
  state: approved
  sourceEventID: rex-whichkey/cli-only-design/approval/1
  changedAt: '2026-11-06T17:00:35.631Z'
title: CLI-only Go menu and data configuration
mtnCreateRetry:
  key: rex-whichkey/cli-only-design/create/1
  intent: 64bf94e608294e9d5d665bb0a8291467e0e05e8764b689d3670bafc057be4f3e
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-06T17:00:31.795Z'
dateModified: '2026-11-06T17:00:36.242Z'
tags:
  - task
type: task
id: 7025a17c-e5a3-48d9-b93a-9a154d364e46
---
# CLI-only Go menu — accepted direction and implementation boundary

## Authority and status

Accepted architectural direction: REXWHICHKEY-20261106-114935. The user explicitly replied, "yes, plan out a new phase for this proposal. Use CLI only. hooks are for the future" after the proposal specified Go-owned configuration/menu/Rex operations, a launch-only Lua keybinding, no runtime/completion registrations, and migration to Go-readable data configuration.

This record captures that accepted direction for planning. It is NOT implemented, runtime verified, a task-dispatch authorization, or permission to modify/deploy/reload the user's configuration. Concrete data-schema, file-location/precedence and CLI spelling details remain implementation planning work; they are not retrospectively described as user decisions.

## Context and problem

The application has a concrete Go Rex CLI resource client and Go menu, but its current configured startup also depends on Lua configuration evaluation, original-VM registries, runtime advertisement, private receipts/relays and completion actions. Actual builds/runs exposed API metadata admission bugs and then failed at the required whichkey_runtime Lua registry check. The source corrections are integrated in 6f71ccf528c9cd6b6cbb869e70f7d86c4b9dccb4; workflow checkpoint 445c5fb records the unresolved runtime failure. Neither commit proves the menu launches.

History inspection found callbacks/functions first written into docs/rex-whichkey.prd in d0ac581 and implemented during Phase 3. The earlier document's Agreed label is not independent evidence of user approval. Current explicit direction governs the new phase; do not use an inherited callback requirement to recreate the rejected bridge.

## Accepted responsibility boundary

1. Go reads and validates data-only menu configuration, supplies defaults, builds dynamic entries from Rex CLI results, renders/navigates the menu, and dispatches real operations through the Rex CLI.
2. All application-owned Rex interaction uses the existing concrete Go rex package and shared argument-array CLI boundary: metadata, session/window/pane inventory, origin/client selection, focus, popup creation/cleanup, rename and ordinary actions. Do not split these responsibilities into Lua calls.
3. Lua's only integration responsibility is the existing Rex keybinding starting the Go binary. No custom action registration, configuration compiler, callback/completion registry, metadata advertisement, API client, helper process controller or second Lua VM is needed.
4. Ordinary CLI invocation and keybinding invocation reach the same Go-owned production menu. Existing registered legacy whichkey actions must neither be required nor redirect/block this path. Do not call rex do with a custom server Lua action name as an indirect substitute for CLI-only operation. Built-in Rex CLI/API operations remain legitimate.
5. Hooks/callbacks, executable Lua configuration, dynamic Lua functions, arbitrary custom-code execution and any hook transport/schema are explicitly deferred. Do not implement stubs or preserve a startup dependency for a future hook facility.
6. Keep the application flat and short-lived. The Go menu remains active for normal interaction until a selection/dismissal/cancellation; do not inherit the old Lua action or callback-ticket lifetime as an idle menu timeout. Operation timeouts and bounded owned cleanup remain.

## Stable contract boundaries

- CLI-ONLY-LAUNCH: a minimal Rex Lua keybinding invokes the installed binary through documented Rex execution facilities. The normal Go entrypoint needs no internal prepare/controller modes, resolved JSON, named Lua action, manually supplied runtime executable directories or compatibility handler. The initiating Lua action's lifetime must not accidentally kill an otherwise active menu.
- DATA-MENU-CONFIG: Go consumes data, not Lua source/functions or callback_ref capabilities. Built-in defaults work without a mandatory custom file. Configuration covers ordinary static keys/groups/labels/flags/arguments and supported Go-owned dynamic providers; invalid/conflicting entries fail before interaction/effects. Reuse sound existing Go Mapping/Config behavior where it fits, but remove obsolete Lua transport fields rather than pretending callbacks are supported. Specify exact schema/version/path precedence and migration examples in the first bounded implementation contract task. JSON is a starting candidate because Config already has a JSON form, not an additional accepted user-format decision.
- REX-CLI-EXECUTION: one frozen execution-host CLI/endpoint policy; concrete Client/Session/Window/Block resources; bounded read-only metadata and accurate supported-schema admission; exact arguments and safe diagnostics. No second direct socket/RPC client, shell interpolation, arbitrary config-driven executable nomination or alternate-server retry.
- GO-MENU-LIFECYCLE: Go owns popup/input/terminal state, retained resource/GUI intent, current origin checks, confirmation, selection and cleanup. User/CLI client identity remains distinct from an intended app target. Preserve moved-resource safeguards, connection/placement freshness where relevant, once-only effects and no uncertain retries. Cleanup addresses exact owned resources, not labels or whole windows/sessions.

## Requirement interpretation and compatibility impact

Existing IDs are not renamed or reused. This accepted direction supplies complementary Phase 16 coverage for REQ-001/002/003/004/006/007/008/009/010/011/012/013/014/015/016/017.

- REQ-003's Lua loading mechanism is superseded for Phase 16 by Go-readable data configuration; preserve configurable menus without an embedded Lua evaluator or inline/imported executable config contract.
- REQ-004's non-executable mapping/navigation intent remains. Function-valued conditions/labels and Lua callback target forms are deferred, not claimed implemented through a renamed data feature.
- REQ-005 original-VM callback execution and callback-specific completion obligations are explicitly deferred to the future. No Phase 16 completion credit is claimed for them.
- Callback/completion-specific clauses in the usability design/REQ-016 retained constraints are superseded for this phase. No-argument direct menu startup, platform endpoint intent, origin/cleanup/user-config ownership remain.
- Retain the current menu/defaults and later accepted menu corrections; the initial old-plugin inventory is not authority to resurrect removed entries or reopen Phases 11/14.

The prior PRD/design text and delivery history remain preserved. This record's accepted CLI-only direction takes precedence for this phase only; it does not silently establish approval of historic callback assertions or migrate legacy task metadata.

## Expected implementation surface

Review cmd/main.go, launch.go, native.go, popup.go, terminal.go, mapping.go, defaults/builtin provider/dispatch files, rex/*.go and current runtime_bootstrap.go. Reuse working Go behavior; remove obsolete bidirectional Lua startup/completion/relay dependencies from production paths and retire dead code/assets only after checking real consumers. Review lua/rex_whichkey* as removal/reduction candidates, not a new Lua subsystem. Nix/flake/Home Manager/package/docs/examples must expose the ordinary binary/data assets and minimal launcher, not runtime policy, helper paths or generated callback wiring. Exact task ownership is to be assigned before execution.

## Rollout, security and alternatives

Document migration from old Lua M.new/setup/register_launch/callback configuration to the data contract and minimal keybinding. Do not overwrite the user's init.lua, mutate immutable installed assets, activate profiles, reload/restart Rex or deploy without separate authorization. Preserve prior commits and records; no destructive cleanup or blanket legacy-record migration. Native/standalone operation must remain supported on existing platforms, with truthful dependency limitations rather than fabricated Rex artifacts or PATH repair.

Rejected for this phase: mandatory whichkey_runtime/completion registrations; Go -> CLI -> custom Lua handler -> Rex control paths; a new callback framework; automatic old-VM hot loading; silently discarding configured menu data; embedded Lua configuration execution; relocating runtime paths into Nix/wrappers/ldflags.

## Verification and residual uncertainty

Implementation completion requires reviewed, integrated source AND the actual built public binary running in Rex, an observed menu, ordinary dismissal and owned cleanup. A zero exit or finite-launch acknowledgement alone does not prove menu display. Exercise a benign owned dynamic/menu operation only within the user's actual-run authorization; preserve unrelated state. Record exact source/build/run results and failures. A preserved older Lua registration cannot be used as a completion prerequisite for the new normal Go path.

Build/format/static checks are allowed. The user's explicit build-and-run request authorizes the actual application validation; it does not reinstate unit/race/TDD/Lua harnesses, broad experiments/remote trials, user-config edits/reloads or Nix activation. Existing test files/quality standards remain intact. If a needed interaction or deployment is outside authority, report that precise unmet criterion; never mark a failing run complete.

Open implementation details: the smallest data format and versioned schema; config path/precedence; exact documented dynamic-provider/action IDs; inventory of removable vs still-consumed transport code. Resolve these in bounded contract/task planning inside the accepted boundary. Do not reopen CLI-only ownership or deferred hooks merely because old source implements the opposite.

## Related records

prds/rex-whichkey.md; prds/rex-whichkey-usability.md; designs/rex-whichkey-runtime-ownership.md (historical bridge baseline); designs/rex-whichkey-decision-register.md; phases/rex-whichkey-15-code-owned-runtime.md.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA2VDE3OjAwOjM0LjkyOFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9jbGktb25seS1kZXNpZ24vYXBwcm92YWwvMSIsImludGVudCI6ImMwMDI1M2JmMmZlNjIxNzI1MTAxY2VjN2U4ZGI0M2IwMzY1ODkzNDRlNDE2ZWRiMzhmZmMwYmJhNjJhZDUzN2MifQ -->
### 2026-11-06T17:00:34.928Z · approval · unspecified

{"approvalScope":"CLI-only responsibility boundary, Go-readable data config, no custom Lua runtime/completion registrations, deferred hooks; phase planning/documentation only. Exact config format/path/CLI details remain explicit open implementation choices. No implementation dispatch/config mutation/deploy/reload approval.","approvedBy":"user","attempt":1,"decision":"approved","decisionEvidence":"User explicitly approved proposal then requested: yes, plan out a new phase for this proposal. Use CLI only. hooks are for the future.","decisionID":"REXWHICHKEY-20261106-114935","deferredRequirements":["REQ-005"],"eventID":"rex-whichkey/cli-only-design/approval/1","kind":"approval","recordRef":"designs/rex-whichkey-cli-only-menu.md","requirementInterpretation":"REQ003 Lua mechanism and REQ004 function-valued features superseded/deferred only as explicitly described in this accepted CLI-only target; IDs and prior history preserved.","schemaVersion":1,"specHash":"sha256:23add062dac4326fe35e7e178a6a39f4c4d7631b4c8e164feb988b9aba07e21b"}
<!-- mtn:end-entry:v1 -->
