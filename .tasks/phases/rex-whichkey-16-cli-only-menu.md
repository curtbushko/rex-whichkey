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
  dependencies: []
lifecycle:
  state: complete
  sourceEventID: rex-whichkey-p16/summary/1
  changedAt: '2026-11-06T23:03:58.149Z'
title: 'Phase 16: CLI-only Go menu and data configuration'
mtnCreateRetry:
  key: rex-whichkey/phase16-cli-only/create/1
  intent: eb9446133b288b4bd4e65cead419ff6743151bd9673fa1e6924c1e20574b68c0
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-06T17:00:37.511Z'
dateModified: '2026-11-06T23:03:58.766Z'
tags:
  - task
type: task
id: 2e653386-0a3e-4f75-8d1d-fcd517899d7b
projects:
  - '[[projects/rex-whichkey.md]]'
completedDate: '2026-11-06'
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

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA2VDIzOjAzOjU3LjQwNFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxNi9zdW1tYXJ5LzEiLCJpbnRlbnQiOiIwMzhkZTRiYmFhYzRhMmUwNzk4ZWQ2MDA1ZDhkMzhhYWU3MWZiYmQxYjJhYWJmMmQwZjEwYzI2MjYwMTI5YWYyIn0 -->
### 2026-11-06T23:03:57.404Z · summary · unspecified

{"acceptedScope":"Phase16 CLI-only Go menu delivered: version1 JSON data config/defaults, bounded Go dynamic providers and retained-resource real effects/inline rename, one normal public CLI path, no mandatory Lua registry/advertisement/completion/relay, launch-only Lua binding example, ordinary binary/optional CLI dependency package and safe Home Manager/migration documentation. Hooks and executable Lua remain deferred; no deployed user configuration change.","actualRuntime":{"additionalOwnedEffectEvidence":".mtn-tmp/p16/task3-public-runtime-evidence.json: actual inline rename of exact disposable tab to p16-owned-cli-rename","artifacts":[".mtn-tmp/p16/final-main-build.log",".mtn-tmp/p16/final-main-menu-root.txt",".mtn-tmp/p16/final-main-menu-dynamic-sessions.txt",".mtn-tmp/p16/final-main-exit.json",".mtn-tmp/p16/final-main-after-session.json",".mtn-tmp/p16/final-main-parent-process.json",".mtn-tmp/p16/final-main-parent-modes.json",".mtn-tmp/p16/final-main-after-owned-window-cleanup.json"],"binary":"bin/darwin-arm64/rex-whichkey","binarySHA256":"76be2e1ee915a91409ade98d09f41f2a9298289fdce385b5dbdde6a36bc40e3d","command":"Normal task build then task run from main in an explicitly owned Rex terminal in existing session","completedAtUTC":"2026-11-06T23:01:07Z","created":{"block_ids":["block:01a1b316-4610-75cc-900a-a19527941e91"],"revision":230,"window_id":"window:01a1b316-4611-7a21-bb21-8ec9b3121c0a"},"exitCode":0,"kind":"actual-public-runtime-evidence","limitations":["Physical keybinding and installed profile/config not migrated/reloaded; source-only Lua/Nix integration.","No prohibited tests/harnesses/benchmarks, custom JSON probes, remote/Linux trials or unrelated mutation.","Full unchanged lint remains failing62 inherited diagnostics; lint:ai not reached."],"normalPublicArgs":[],"observed":["Actual root Rex menu/current defaults","Actual Go dynamic Switch session inventory/full IDs/current marker","Backspace navigation and ordinary Escape dismissal","Public process exit0 and primary terminal restored","Exact owned popup layer/block/private directory absent before coordinator closes only its disposable validation window"],"ownedCleanupPassed":true,"resources":{"label":"Window b3121c0a","layers":[{"blocks":["block:01a1b316-4610-75cc-900a-a19527941e91"],"kind":"tiled","layer_id":"layer:01a1b316-4611-7a25-8c61-b7c4e7ea2fd0"},{"blocks":["block:01a1b316-5006-7616-b380-b46bf1c8deb3"],"kind":"floating","layer_id":"layer:01a1b316-5007-7772-8573-8f5a7e7a738a"}],"window_id":"window:01a1b316-4611-7a21-bb21-8ec9b3121c0a"},"schemaVersion":1,"targetBranch":"main","targetCommit":"9fa9310b1398672fdc0f78823d96cff0f6ec0292","targetTree":"51327ec9528c6060ebb34fb7cc1710ac57cbe91b"},"approvalSummary":["User explicitly accepted CLI-only proposal REXWHICHKEY-20261106-114935 and then authorized execution with start it.","Coordinator current phase authority2 and bounded approved task-plan1 cover routine interfaces/decomposition/review/integration, not invented user approval of exact JSON details.","Completed native Task01 summary2, Task02 summary2 and Task03 summary1 retain exact current source review/integration proofs; source reviews are fresh same-family, not cross-family.","Final main normal task build/run observed actual menu/dynamic navigation/Escape/exit0/terminal/exact cleanup; combined final reviewed source also observed real inline rename of exact disposable owned window."],"attempt":1,"cleanup":{"actualOwnedValidationWindowsRemoved":4,"cleanupPlan":".mtn-tmp/p16/final-cleanup-plan.json","globalOrUserCleanup":false,"inventory":".mtn-tmp/p16/final-resource-inventory.json","retained":"Stopped negative Task01/reviewer1 and Task02/reviewer1 exact windows/worktrees retained; original coordinator/session and older Phase15/unowned resources preserved","successfulDeliveredWorktreesRemoved":12,"successfulWorkerWindowsConfirmedClosed":9,"tooling":"Immutable complete facade cleanup jq defect recorded; native delivery verified, eligible successful owned resources manually cleaned with exact scope/clean/ancestry checks; compaction through official helper succeeded"},"coverage":{"REQ-001":["tasks/rex-whichkey-16-02-go-cli-menu.md","tasks/rex-whichkey-16-03-launch-package-docs.md"],"REQ-002":["tasks/rex-whichkey-16-02-go-cli-menu.md"],"REQ-003":["tasks/rex-whichkey-16-01-data-config.md","tasks/rex-whichkey-16-02-go-cli-menu.md","tasks/rex-whichkey-16-03-launch-package-docs.md"],"REQ-004":["tasks/rex-whichkey-16-01-data-config.md","tasks/rex-whichkey-16-02-go-cli-menu.md"],"REQ-006":["tasks/rex-whichkey-16-01-data-config.md","tasks/rex-whichkey-16-02-go-cli-menu.md"],"REQ-007":["tasks/rex-whichkey-16-01-data-config.md","tasks/rex-whichkey-16-02-go-cli-menu.md"],"REQ-008":["tasks/rex-whichkey-16-02-go-cli-menu.md"],"REQ-009":["tasks/rex-whichkey-16-02-go-cli-menu.md"],"REQ-010":["tasks/rex-whichkey-16-01-data-config.md","tasks/rex-whichkey-16-02-go-cli-menu.md"],"REQ-011":["tasks/rex-whichkey-16-02-go-cli-menu.md"],"REQ-012":["tasks/rex-whichkey-16-01-data-config.md","tasks/rex-whichkey-16-02-go-cli-menu.md"],"REQ-013":["tasks/rex-whichkey-16-02-go-cli-menu.md","tasks/rex-whichkey-16-03-launch-package-docs.md"],"REQ-014":["tasks/rex-whichkey-16-03-launch-package-docs.md"],"REQ-015":["tasks/rex-whichkey-16-03-launch-package-docs.md"],"REQ-016":["tasks/rex-whichkey-16-02-go-cli-menu.md","tasks/rex-whichkey-16-03-launch-package-docs.md"],"REQ-017":["tasks/rex-whichkey-16-02-go-cli-menu.md","tasks/rex-whichkey-16-03-launch-package-docs.md"]},"deferredRequirements":[{"reason":"Explicit CLI-only decision defers hooks/callback execution and transport. No task claims it, adds schema placeholders or repairs the old bridge.","requirement":"REQ-005"},{"reason":"Function-valued labels/conditions, callback refs and arbitrary executable configuration are deferred, not supported via data aliases.","requirement":"REQ-004 executable portions"}],"eventID":"rex-whichkey-p16/summary/1","files":["bin/darwin-arm64/rex-whichkey","data_config.go","providers.go","cmd/main.go","launch.go","native.go","rex/actions.go","rex/wait.go","lua/rex_whichkey.lua","examples/rex-whichkey-keybinding.lua","examples/my_whichkey.json","nix/package.nix","nix/home-manager.nix","README.md","docs/mapping-format.md","docs/launch-contract.md","tests/fixtures/phase15-lua",".mtn-tmp/p16/final-main-public-runtime-evidence.json",".mtn-tmp/p16/task3-public-runtime-evidence.json",".mtn-tmp/p16/final-plan-validation-result.json",".mtn-tmp/p16/final-collection-validation.json",".mtn-tmp/p16/final-resource-inventory.json",".mtn-tmp/p16/final-cleanup-plan.json"],"kind":"summary","limitations":["No user config/profile/package install or activation/reload/physical keybinding test.","No tests/unit/race/TDD/Lua/parser/FIFO harnesses/benchmarks or remote/Linux action trials; existing assertions/standards preserved.","REQ005/hooks/callbacks and executable REQ004 portions remain explicitly deferred; REQ003 Lua loading superseded by Go-readable JSON data, original IDs preserved.","Two stopped negative-review windows/worktrees retained per failed-resource policy; no global or user cleanup.","Nonblocking installed Markdown-link portability REX-P16-T03-001 remains.","Phase15 original failed bridge runtime/history and unrelated phases remain uncompleted; this is the approved replacement, not retroactive bridge success."],"milestoneEventID":"rex-whichkey-p16/milestone/1","nativeEvidenceCommit":"72309319a6909535c492023bba02160e0f4e718b","optionalValidation":[{"artifact":".mtn-tmp/p16/final-main-optional-lint.log","command":"Final main task lint","diagnostics":62,"exitCode":201,"limitation":"Inherited baseline, no global lint pass or AI-lint completion claimed","result":"fail"}],"recordRef":"phases/rex-whichkey-16-cli-only-menu.md","records":["tasks/rex-whichkey-16-01-data-config.md","tasks/rex-whichkey-16-02-go-cli-menu.md","tasks/rex-whichkey-16-03-launch-package-docs.md"],"requirements":["REQ-001","REQ-002","REQ-003","REQ-004","REQ-006","REQ-007","REQ-008","REQ-009","REQ-010","REQ-011","REQ-012","REQ-013","REQ-014","REQ-015","REQ-016","REQ-017"],"schemaVersion":1,"specHash":"sha256:794f84db12b1642de93d2e72d37902acf5ef3e0076de8b27a7aa3ca1df435d16","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"9fa9310b1398672fdc0f78823d96cff0f6ec0292","targetTree":"51327ec9528c6060ebb34fb7cc1710ac57cbe91b","verification":[{"artifact":".mtn-tmp/p16/final-main-build.log","command":"Final main normal task build, owned Go formatting/whitespace and task lint:vet","result":"pass"},{"artifact":".mtn-tmp/p16/final-main-nix-format.log","command":"Final Nix syntax parse and alejandra format checks only","result":"pass"},{"artifact":".mtn-tmp/p16/final-main-public-runtime-evidence.json","command":"Actual main public task run observed root/dynamic menu/navigation/Escape, exit0, primary terminal and exact owned popup/block/private-directory/window cleanup","result":"pass"},{"artifact":".mtn-tmp/p16/task3-public-runtime-evidence.json","command":"Combined reviewed source actual inline tab rename on exact owned disposable target; observed changed label, exit0 and exact cleanup","result":"pass"},{"artifact":".mtn-tmp/p16/final-plan-validation-result.json","command":"Current approved design/phase plan/spec hashes/full9-criterion16-requirement coverage and completed child delivery ancestry; nonempty3/3 acyclic native hierarchy; native collection valid","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
