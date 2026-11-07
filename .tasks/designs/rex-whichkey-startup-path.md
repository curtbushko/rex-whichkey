---
identity:
  kind: design
  designID: rex-whichkey-startup-path
  uuid: 56880036-5329-411a-b9ef-2bf4fe7ae988
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
  documents:
    - ref: designs/rex-whichkey-cli-only-menu.md
      hash: 'sha256:23add062dac4326fe35e7e178a6a39f4c4d7631b4c8e164feb988b9aba07e21b'
lifecycle:
  state: approved
  sourceEventID: rex-whichkey/startup-design/approval/1
  changedAt: '2026-11-07T00:16:17.182Z'
title: 'Startup path: PATH-only lookup, draw-first menu and asynchronous data'
mtnCreateRetry:
  key: rex-whichkey-startup-path-design-create-v1
  intent: bf95b7b7b11353df1eae08d9a8c4917e92263543502be65f72ed7d2250389aef
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T00:16:10.823Z'
dateModified: '2026-11-07T00:16:18.210Z'
tags:
  - task
type: task
id: d3eb6582-6a04-4514-be07-9f118719acb2
---
# Startup path — PATH lookup, draw-first menu and asynchronous data

## Authority and status

Decision REXWHICHKEY-20261106-191113 records the current user-approved startup direction. The ID is assigned while recording this decision, not claimed as the exact time of earlier discussion. User explicitly approved: "remove the validation. Only check that the binaries are in the path"; separate background/on-demand data; load asynchronously; remove duplicate queries; and verify. After explanation of temporary-executable disk sync and asynchronous UI updates, the user requested: "add these changes as a new phase".

Accepted planning constraints are recorded here; implementation and runtime timing have not started. The precise async update mechanism remains an implementation detail to resolve in bounded task planning, not an already selected protocol or permission to add a framework. Creation of the phase does not authorize dispatch, source changes, user configuration/deployment/reload or broad test/benchmark execution.

## Problem and verified source observations

The delivered Phase16 program works, but performs unnecessary blocking startup work before displaying the menu. Source tracing shows an ordinary successful resource operation currently has a server.status subprocess before and after the requested Rex command; each of those three processes rechecks executable content before/after, resulting in six full executable hashes. Startup also eagerly loads inventories and rename labels, repeats metadata reads across layers, and copies/syncs/verifies the popup image. These are source observations, not a measured allocation of elapsed startup time. Do not describe them as unavoidable Rex CLI requirements or claim a measured speedup.

Baseline: current committed Phase16 implementation and its actual public menu/owned-rename/dismissal/cleanup evidence. Phase17 remains a separate audit-only phase. A local flake.nix Darwin pre-commit packaging fix exists outside this planning work and must remain separate.

## Accepted changes and supersession

1. **STARTUP-PATH-LOOKUP.** Discover required binaries through normal execution-host PATH using ordinary executable lookup and clear missing-binary errors. Remove startup executable SHA256 validation, binary receipt revalidation, repeated server.status/config-generation wrapper checks, version/API-capability/schema-negotiation gates and manual executable-location heuristics. Do not retain REX_BIN_DIR, Lua anchors or private libexec lookup as alternative executable search authorities. Actual CLI execution errors remain errors; no uncertain operation retry or alternate-server fallback. This intentionally replaces the previous heavier executable/compatibility policy, rather than silently weakening or claiming to preserve it.
2. **STARTUP-FIRST-FRAME.** Draw the static validated menu before nonessential metadata is ready. Establish only the endpoint/origin/resource and presentation information actually required to create/display the correct popup. Preserve local JSON configuration validation, explicit endpoint/client intent, target identity, confirmation, action-specific data checks, safe result decoding, terminal restoration and exact owned cleanup. Removing transport-level preflight is not authorization to dispatch against the wrong origin or drop safety/behavior checks.
3. **STARTUP-ASYNC-DATA.** Separate background and on-demand work. Sessions/windows/panes/known-host inventories and rename presentation labels are candidates for asynchronous loading; choose which start after first frame versus submenu/form demand during task planning. Only real metadata reads are moved, not action effects. Loading/unavailable states must be explicit; entries requiring pending metadata cannot dispatch. Known hosts remain inventory-only, without remote probes/attach.
4. **STARTUP-QUERY-REUSE.** Reuse per-invocation resource/inventory snapshots and in-flight loads across consumers instead of duplicate startup queries. No daemon, cross-run persistent cache or service framework. Fresh origin/client-generation/destination checks immediately before a selected effect remain semantically necessary and are not duplicate startup reads to delete. Built-in app-action metadata is already deferred until selection; do not move it into startup or claim it as a newly deferred cost.
5. **STARTUP-NO-DURABILITY-SYNC.** Remove startup fsync/durability waits for the invocation-private copied executable: power-loss persistence is unnecessary for this short-lived popup. Complete ordinary writes/close and process startup correctly. This does not approve blanket removal of every Sync, discarding reliable selection handoff, or a predetermined removal of the controller/renderer split or executable copy itself. Remove hashing wherever it merely reinstates the rejected startup binary-validation mechanism.
6. **STARTUP-UI-UPDATES.** Fetch in bounded goroutines, return immutable results to the UI owner, and redraw from the UI event loop. Current controller/renderer exchange a single immutable initial snapshot, so safe updates and matching parent-retained target publication need a small coordinated mechanism. Preserve single-writer menu state and consistent selected keys/targets; workers cannot mutate live menu slices concurrently, overwrite a user's in-progress rename input, execute actions or publish stale results after dismissal. Exact channel/IPC layout is not yet decided. Cancel/join workers, bound loads/results and preserve errors during teardown. No new Lua control path, custom Rex registration or hook facility.
7. **STARTUP-VERIFICATION.** Verify actual ordinary public startup and interaction, distinguishing launch-to-first-visible-frame, required-data-ready and relevant actions-ready. Record measurements and uncertainty on comparable source/configuration/host context before and after the change; no numeric latency SLA has been specified. Use targeted actual app observations/timing, not a new benchmark framework, private substitute harness, broad experiments or a fictitious performance claim. Check responsiveness during loading, correct loaded choices and benign owned effect, dismissal/cancellation and exact cleanup.

## Stable behavior and security boundary

The earlier designs/rex-whichkey-cli-only-menu.md remains the product baseline except where STARTUP-PATH-LOOKUP and STARTUP-NO-DURABILITY-SYNC explicitly replace its implementation policy and later Phase15/16 receipts/compatibility assumptions. The executable trust model now relies on normal PATH/environment and the operating system, not application rehashing; this is an explicit user choice. Preserve endpoint selection and execution-host/viewing-host separation, concrete Go Rex resources, argument-array execution, current defaults/data schema/navigation/theme/inline rename, retained origin/GUI identity and once-only effects after owned cleanup. Resolve ready-state action availability without retargeting or renewing the wrong app generation. Keep fixed invocation bounds and pagination rather than introducing a resize-on-navigation behavior change merely because data arrives later.

Hooks/callbacks, executable Lua configuration, direct socket clients, daemons, new transport frameworks, automatic endpoint/executable fallback, user config/profile changes and deployment remain out of scope. Preserve tests/assertions/Taskfile/hooks/lint settings; do not lower standards or hide inherited lint failures. Phase17's general unnecessary-parts audit is not folded into this phase.

## Interfaces, expected files and task boundaries

Likely surfaces: rex/runtime.go, rex/exec.go, rex/rex.go, rex/endpoint.go and concrete resource helpers; launch.go/native.go; providers.go/hosts.go; popup.go/runtime_bootstrap.go; terminal.go/menu.go/text-entry handling; discovery/package and active documentation as necessary to expose normal PATH requirements. No data-config or Lua-launcher public rewrite is assumed. Assign literal ownership and a minimal ready/update contract before consumers; produce the CLI/async-update interface before integrating async provider and first-frame consumers. Each implementation task owns relevant source/static/build checks and combined real-app verification; do not create a disconnected benchmark/research phase.

## Verification, rollout and unresolved details

Allowed future verification is source review, unchanged formatting/static/normal build gates and targeted actual public-app timing/behavior/owned cleanup under the user's explicit "yes, verify". Existing unit/race/TDD/Lua/parser/FIFO harness and broad benchmark/remote-experiment restrictions are not lifted. No user configuration/reload/profile activation or push without separate authority. Record a real failure as failure, not a compile/help/private ACK pass.

Before implementation resolve the minimal controller/renderer update interface, per-provider background-versus-demand placement, loading/error display integration and timing collection boundaries. These are routine bounded task-planning choices within accepted constraints, not invented approvals or new prerequisite research. Material public-contract/security departures still need the user. Keep prior source, decisions and delivered evidence intact; use reviewed isolated source and coordinator-only integration if execution is later authorized.

## Requirement interpretation and related records

Complementary scope: REQ-001/002/003/004/006/008/009/010/011/012/013/014/015/016/017, preserving IDs. Req003/004 follow the delivered data-only interpretation; REQ005/hooks and executable portions remain deferred. This record narrows earlier runtime compatibility/discovery implementation choices under REXWHICHKEY-20261104-193124 and refines REXWHICHKEY-20261106-114935 without reversing Go/CLI ownership, cleanup or deferred hooks. Phase16 delivery is historical baseline; Phase17 audit stays independent. New phase: phases/rex-whichkey-18-startup-speed.md.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDAwOjE2OjE2LjAzOFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9zdGFydHVwLWRlc2lnbi9hcHByb3ZhbC8xIiwiaW50ZW50IjoiZTVmNTRmYWQwNDVhOTY2NzE2NThjOGY4YjNkZmQxNzJhNzdlNjgxM2NiZmUyY2RjNWNhZjliOTAxODQxNWIwMiJ9 -->
### 2026-11-07T00:16:16.038Z · approval · unspecified

{"approvalScope":"Accepted startup constraints and design/phase planning only. No exact async IPC mechanism, executable-copy removal, task dispatch, implementation, latency SLA, broad benchmark/test expansion, user config/deployment/reload or unrelated audit cleanup inferred.","approvedBy":"user","attempt":1,"decision":"approved","decisionEvidence":"User explicitly approved PATH-only binaries/no validation, separate background/on-demand work, async loading, removing duplicate queries and verification; after fsync/UI-update explanation requested add these changes as a new phase.","decisionID":"REXWHICHKEY-20261106-191113","eventID":"rex-whichkey/startup-design/approval/1","kind":"approval","recordRef":"designs/rex-whichkey-startup-path.md","schemaVersion":1,"specHash":"sha256:c7a6d0d736fbbb3925173391431a05eb2c8c65273c2a81e2e5e7d0ac23e1c85c"}
<!-- mtn:end-entry:v1 -->
