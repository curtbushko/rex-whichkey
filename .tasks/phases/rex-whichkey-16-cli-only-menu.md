---
identity:
  kind: phase
  phaseID: rex-whichkey-16-cli-only-menu
  uuid: 6cb0e0ac-7038-4432-b9dd-36cb0585e1bd
structure:
  order: 16
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
    - ref: designs/rex-whichkey-cli-only-menu.md
      hash: 'sha256:23add062dac4326fe35e7e178a6a39f4c4d7631b4c8e164feb988b9aba07e21b'
    - ref: prds/rex-whichkey.md
      hash: 'sha256:97614db611a6f35a44d0db32ccf2bf4d4854d6ffea3dfd5375959e59259f01ff'
    - ref: prds/rex-whichkey-usability.md
      hash: 'sha256:966ef70c264360b4820e5faf510496166e91f0effeae479ab7400a3792de88e9'
lifecycle:
  state: incomplete
  sourceEventID: rex-whichkey-p16/planning-validation/1
  changedAt: '2026-11-06T17:05:48.511Z'
title: 'Phase 16: CLI-only Go menu and data configuration'
mtnCreateRetry:
  key: rex-whichkey/phase16-cli-only/create/1
  intent: eb9446133b288b4bd4e65cead419ff6743151bd9673fa1e6924c1e20574b68c0
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-06T17:00:37.511Z'
dateModified: '2026-11-06T17:05:49.124Z'
tags:
  - task
type: task
id: 2e653386-0a3e-4f75-8d1d-fcd517899d7b
projects:
  - '[[projects/rex-whichkey.md]]'
---
# Phase 16 — CLI-only Go menu and data configuration

## Context and accepted decision

The user approved proposal REXWHICHKEY-20261106-114935: "yes, plan out a new phase for this proposal. Use CLI only. hooks are for the future". This is one planned phase, not permission to start implementation, deploy assets, change the user's configuration or reload Rex.

Target flow: a Rex Lua keymapping starts rex-whichkey; Go reads configuration, queries the Rex CLI for session/window/pane information and dynamic items, and owns the displayed menu and dispatch. Lua is a thin launch integration only. The CLI-only design is authoritative where prior Lua-configuration/callback/registration documents conflict. Hooks and callbacks are explicitly future scope.

Phase 15's current public binary still fails before popup creation because its configured path requires a whichkey_runtime Lua action absent from the installed integration. Do not make that rejected bridge work as a prerequisite to replacing it. Source baseline includes 6f71ccf528c9cd6b6cbb869e70f7d86c4b9dccb4 (observed API admission corrections) and 445c5fb (honest blocked-runtime record); neither is proof the menu runs.

## Deliverable and scope

One flat, short-lived Go application with data-only configurable menus, Go-owned dynamic providers and Rex CLI resource operations, a minimal Lua launch keybinding, ordinary packaging and accurate migration instructions. Preserve delivered menu rendering/navigation, theme, confirmation, inline text rename and current defaults/corrections instead of rebuilding the UI or reverting to the old Lua plugin's feature list.

### Acceptance criteria

1. **Single Rex control path.** Application metadata queries, dynamic session/window/pane/known-host menus, focus, popup/resource creation, rename, dispatch and cleanup use the shared concrete Go rex client/argument-array Rex CLI boundary. Production Lua makes no direct Rex query/mutation calls beyond documented execution of the binary. No application-owned direct socket client or CLI-to-custom-Lua-action detour.
2. **Ordinary startup.** Both the no-argument public binary and the documented keybinding reach the same real Go menu. Startup does not query/require whichkey_runtime, custom completion actions, user action registration, the old M.new/setup/register_launch API, helper wiring, private handoff invocation or a Lua VM registry. Existing legacy registrations do not block or select the new path.
3. **Go-readable configuration.** Built-in defaults work without a required custom file. Go loads/validates data-only custom configuration with a documented schema/version, source selection and deterministic precedence. Preserve appropriate existing static mapping/group/label/args/confirmation/hidden semantics, Unicode/key validation and conflict errors. Invalid configuration fails clearly before effects. Do not execute/transpile Lua or accept callback_ref/executable function forms as supported new configuration. Specify the smallest format/loader contract before downstream implementation.
4. **Dynamic data and real effects.** Go builds supported dynamic items from bounded Rex CLI results and retains exact selected resource identities. Selection performs the intended Rex operation once; no no-op success, dropped rename text, callback delegation or wrong focused-resource substitute. Preserve current WhichKey text-entry rename pages and retained-resource CLI dispatch, not a native dialog.
5. **Lifecycle and safety.** The menu stays active for normal interaction and owns its process/input/popup lifetime; no accidental Lua-action timeout or legacy callback-ticket idle expiry. Operation deadlines, cancellation, exact-origin/GUI-client intent, confirmation, moved-resource freshness, terminal restoration, bounded owned cleanup and no uncertain-effect retries remain. Viewing-host paths/opaque ctx.server are not endpoint/executable authority.
6. **Retire obsolete integration.** Inventory and remove or explicitly isolate unneeded Lua evaluation/dispatch modules, callback registries/capabilities, read-only advertisement, completion-receipt negotiation, public file relay, finite bridge commands and assets from ordinary production paths. Do not leave them mandatory, silently drop configured data, create a second Lua VM, add a daemon/framework or implement future hook placeholders. Reuse useful generic Go lifecycle/CLI safety code rather than deleting blindly.
7. **Packaging and migration.** Nix builds/installs ordinary binary and necessary data/launcher assets, not runtime policy/wrappers/ldflags/generated helper directories. Update example keybinding and all active docs from Lua config/registration to the new data contract. Explain migration without editing/reloading the user's init.lua or activating/deploying anything. Preserve default standalone/native support and truthful existing platform/licensed dependency limitations.
8. **Actual build and run.** On the final combined reviewed source, run the normal project build and execute the actual public rex-whichkey binary in its Rex origin. Observe a real menu, dismiss it normally, and confirm exact owned popup/terminal cleanup. Do not substitute help output, compilation alone, a private harness, a finite-start ACK or a failing run. Record which entrypoint/actions were actually exercised; a keyboard-mapping/configuration reload needs separate permission if necessary. Keep the phase incomplete if the app still fails.
9. **Accurate records and history.** Link current source/review/integration/build/run evidence, carry unresolved issues explicitly, and distinguish accepted design from proposed format/CLI details and verified implementation. Preserve old commits/records and current quality standards; do not mark Phase 15's failed runtime as fixed or automatically complete unrelated phases.

## Implementation workstreams for later task decomposition

These are planning boundaries, not created/dispatchable tasks:

- Contract/data loader: define the minimal Go-readable configuration/actions/providers and migration mapping, using existing Go Config/Mapping code where appropriate; no hook API.
- Go execution/lifecycle: route all Rex operations through rex CLI resources and one Go entrypoint, remove registered-Lua discovery/completion from startup/dispatch, preserve exact targeting and cleanup.
- Dynamic menu behavior: move remaining Lua-owned data/dispatch responsibilities into Go, preserve navigation/current defaults/text rename and active-menu lifetime.
- Minimal integration/package/docs: reduce Lua to launch-only keybinding, retire obsolete shipped modules/assets, maintain normal optional dependencies and migration documentation.
- Combined verification/review: independently review integrated responsibility boundaries and run the actual built public application/menu/dismissal/cleanup. Each implementation task owns its relevant checks; verification is not a disconnected research phase.

Expected areas: cmd/main.go, mapping.go, launch.go, native.go, popup.go, terminal.go, defaults/builtin providers/dispatch, rex/*.go, runtime_bootstrap.go, Lua module tree (retirement/reduction), examples, nix/package.nix, nix/home-manager.nix, flake.nix and active launch/config/remote docs. Assign exact bounded ownership/interfaces/prerequisites before worker dispatch; do not hand an unrestricted repository rewrite to one worker.

## Requirements and supersession

Complementary scope: REQ-001/002/003/004/006/007/008/009/010/011/012/013/014/015/016/017, governed by the CLI-only design.

- REQ-003's old Lua loading mechanism is superseded here by data configuration; its ID is not renamed or repurposed silently.
- REQ-004 retains static compilation/navigation intent; function-valued strings/conditions are deferred.
- REQ-005 and all callback/hook execution/transport are explicitly deferred to a future separately agreed phase. Do not claim coverage/completion of them.
- Old callback-specific constraints in direct-launch/usability documents do not authorize recreating the bridge. Current Go CLI targeting, normal discovery, resource safety and user configuration ownership remain.

The accepted CLI-only direction supersedes the Phase 15 mandatory Lua runtime/completion control plane for this target. Deployment request REXWHICHKEY-20261106-101053 to install that bridge was never approved and is superseded by this direction; it is not new deployment permission.

## Prerequisites, readiness and parallelization

Native membership: projects/rex-whichkey.md, order 16. No whole-phase Phase 15 finish-to-start edge: its outstanding old-handler/deployment blocker is intentionally replaced, not a prerequisite. Reuse the already integrated current Go CLI/menu source and inspect actual compatibility/lifecycle changes at the recorded committed base. Do not reopen Phases 11/14 or migrate legacy records.

Planning is authorized now. Implementation is not started: create/approve a bounded task graph and resolve concrete loader/action interfaces before dispatch. Contract producer tasks precede their consumers. Shared main/launch/config/runtime/Lua removal paths serialize; documentation/packaging can proceed independently only after immutable contracts and disjoint ownership exist. Preserve the existing single-coordinator/resource limits.

## Verification policy and protected actions

Required future gates: source review, owned formatting/static checks, task build, exact combined Git revision/ownership, normal built application invocation, observed menu/dismissal/cleanup, accurate diagnostics and collection/graph/coverage validation. Actual app build/run is explicitly authorized by the user; broader unit/race/TDD/Lua harnesses, benchmarks, live/remote experiments and test-plan expansion remain prohibited unless explicitly changed. Existing tests/assertions/Taskfile/lint/hooks are preserved, not weakened or run indirectly by hooks/Nix checks.

No real user configuration mutation, package/profile activation, Rex restart/reload, push/deploy, remote trial or global/unowned cleanup is authorized by phase creation. Provide source/migration instructions and report any exact remaining access/deployment blocker rather than invent runtime success.

## Open implementation details

Choose/document the minimal data format (existing Go JSON form is a candidate, not an approved new mandatory format), config location/precedence, exact action/provider vocabulary and retained vs retired transport code in bounded task planning. These are not permission to reopen CLI-only ownership or deferred hooks. No additional product capabilities are assumed.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA2VDE3OjA0OjA4Ljg2MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMTYvZGVmaW5pdGlvbi1hcHByb3ZhbC8xIiwiaW50ZW50IjoiMGI2ODI3YWJkMGViNDc0MzhhMjVmYzA4ZTgzZGU5YjQyY2Y5ZjA2OWY4ODM2YzdkNWZkNGQ2Y2Y5ZjNjMmEzYiJ9 -->
### 2026-11-06T17:04:08.861Z · approval · unspecified

{"approvalScope":"Phase definition/accepted architecture only; no source implementation or real config/deploy/reload action.","approvedBy":"user","attempt":1,"decision":"approved","decisionEvidence":"User explicitly approved CLI-only proposal and requested one new phase, hooks deferred.","decisionID":"REXWHICHKEY-20261106-114935","documents":[{"hash":"sha256:23add062dac4326fe35e7e178a6a39f4c4d7631b4c8e164feb988b9aba07e21b","ref":"designs/rex-whichkey-cli-only-menu.md"}],"eventID":"rex-whichkey-p16/definition-approval/1","kind":"approval","recordRef":"phases/rex-whichkey-16-cli-only-menu.md","schemaVersion":1,"specHash":"sha256:e91df9a2e9764ff3ccc3434b4aa1818b9527b72c4b9853b2d01d10c484c7c4ad","taskPlanApproved":false}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA2VDE3OjA1OjQ3Ljc4OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxNi9wbGFubmluZy12YWxpZGF0aW9uLzEiLCJpbnRlbnQiOiIzNGU2MDM0MjNkYTA2NjhmYzIzZmVlNjU1MDJiOTg3NDMwZjQ1ZTA1ZjcxZGU3ZGM0YjViMDU3N2Q3ODk4NDk3In0 -->
### 2026-11-06T17:05:47.789Z · historical · unspecified

{"attempt":1,"content":"Created one Phase16 and its narrow CLI-only design, native project membership, no whole Phase15 prerequisite, unique order16 and acyclic graph. Current scope/document fingerprints and architecture/phase approval read back; REQ005 deferred, callback mechanisms not dispatch prerequisites. Additive project plan persistence verified through native API after legacy root facade post-write unknown-kind readback; no schema/history migration or blanket legacy approval. No implementation tasks/workers, source changes, tests/product runs, user configuration/deploy/reload actions in this planning round. Phase remains open until implementation and actual app validation.","eventID":"rex-whichkey-p16/planning-validation/1","gateEligible":false,"kind":"historical","provenance":"Coordinator native collection/hierarchy/dependency/order/coverage/fingerprint/approval readback after explicit phase-planning request.","recordRef":"phases/rex-whichkey-16-cli-only-menu.md","schemaVersion":1,"specHash":"sha256:e91df9a2e9764ff3ccc3434b4aa1818b9527b72c4b9853b2d01d10c484c7c4ad"}
<!-- mtn:end-entry:v1 -->
