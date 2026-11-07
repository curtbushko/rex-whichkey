---
identity:
  kind: phase
  phaseID: rex-whichkey-18-startup-speed
  uuid: ba8818f4-01b3-4ed2-ac65-c89d0c2b024c
structure:
  order: 18
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
    - ref: designs/rex-whichkey-startup-path.md
      hash: 'sha256:c7a6d0d736fbbb3925173391431a05eb2c8c65273c2a81e2e5e7d0ac23e1c85c'
    - ref: designs/rex-whichkey-cli-only-menu.md
      hash: 'sha256:23add062dac4326fe35e7e178a6a39f4c4d7631b4c8e164feb988b9aba07e21b'
lifecycle:
  state: complete
  sourceEventID: rex-whichkey-p18/summary/1
  changedAt: '2026-11-07T02:44:55.265Z'
title: 'Phase 18: Faster startup and asynchronous menu data'
mtnCreateRetry:
  key: rex-whichkey-phase18-startup-speed-create-v1
  intent: 2d016db0f24e932bb1167398c657a9ae4a3a6d15350ff1f4b299d1ccae598d0c
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T00:16:20.329Z'
dateModified: '2026-11-07T02:44:55.887Z'
tags:
  - task
type: task
id: bf64e407-caf0-4537-bff8-638b6c6afae8
projects:
  - '[[projects/rex-whichkey.md]]'
completedDate: '2026-11-06'
---
# Phase 18 — Faster startup and asynchronous menu data

## Context and authority

User explicitly approved the startup changes: "remove the validation. Only check that the binaries are in the path"; separate background/on-demand work; "load asynchronously"; "remove duplicate queries"; and "yes, verify". After explanations of executable fsync and UI updates, the user requested "add these changes as a new phase". Accepted planning decision: REXWHICHKEY-20261106-191113; assigned at capture, not an exact historical conversation timestamp.

This is one planned implementation phase, not an instruction to dispatch workers or start code/timing work now. Its authoritative narrow startup design is designs/rex-whichkey-startup-path.md, supplementing the accepted CLI-only product baseline. Phase17 remains the independent unnecessary-parts audit; it is not a prerequisite or authorization for this phase to perform general cleanup. Preserve the current separate uncommitted flake.nix Darwin dev-shell fix.

The existing menu has real Phase16 public-startup/dynamic-navigation/inline-owned-rename/dismissal/cleanup evidence, but startup timing contributions were not profiled. Source analysis found three processes/six executable hashes around each successful ordinary resource operation, eager provider/label reads, duplicate startup queries and copied-executable durability waits. Do not turn this source analysis into invented timing measurements.

## Deliverable

A visibly earlier, responsive static menu using normal execution-host PATH binaries, with nonessential data loaded asynchronously/on demand and shared once per invocation. Remove the agreed expensive startup validation and unnecessary executable durability sync. Preserve current public menu/configuration/actions, safe target/cleanup behavior and the flat CLI-only Go architecture, with minimal Lua launch integration and no hooks.

## Acceptance criteria

1. **PATH-only binaries; no rejected preflight.** Required executables are located on execution-host PATH with clear missing-binary handling. Remove startup/full per-command executable hashing, frozen binary receipt revalidation, repeated server.status/config-generation sandwich calls and version/API/schema-negotiation startup gates. No REX_BIN_DIR/anchor/private-libexec search fallback or new arbitrary executable path policy. Functional CLI requests still fail honestly; no uncertain effect retry or endpoint retargeting. Update affected discovery/package/user docs consistently.
2. **Draw first with minimal necessary work.** Validated static root content is visible before inventories and other nonessential labels complete. Keep required endpoint/origin/resource context and local data validation; startup simplification does not remove configuration errors, confirmations, precise action targets, result integrity, terminal restoration or exact owned cleanup. Loading does not freeze normal input or block dismissal.
3. **Async/background versus demand.** Load session/window/pane/known-host data asynchronously, with explicit per-kind background/on-demand choices and rename-label loading when relevant. Inactive/disabled/unused providers do not trigger unnecessary work. Show loading/unavailable status and prevent dispatch of incomplete targets. Preserve full IDs, bounded inventory, warning/empty/pagination/hidden/condition behavior and listing-only known hosts, without probes/attach.
4. **No duplicate startup queries.** Consumers share bounded per-invocation snapshots/in-flight loads for the same information. Independently needed reads may proceed concurrently. Do not delete fresh origin/GUI-generation/destination checks needed immediately before a selected effect; these protect a changed interaction context rather than duplicate initial reads. No daemon or persistent cross-run cache. Already deferred app-action metadata stays off startup.
5. **No copied-executable durability wait.** Remove startup fsync that exists only to make an invocation-private executable copy survive power loss; normal completed write/close/start still works. Do not retain duplicate hashing elsewhere under another name. Whether the copy itself is still needed is not predetermined, nor is blanket removal of every result/file Sync or lifecycle boundary authorized.
6. **Safe incremental UI and worker lifetime.** Specify the minimal controller/renderer update contract before implementing consumers. Background workers return immutable results; the UI event loop alone changes live menu state/redraws and the controller retains corresponding authoritative resource targets. No data races, stale-result resurrection, background action dispatch or overwritten in-progress input. Pending actions cannot execute until their context is ready. Cancellation/dismissal stops and joins workers; cleanup remains bounded and precedes effects. Avoid a new daemon/framework/Lua registry or unnecessary protocol machinery.
7. **Existing behavior remains.** Preserve current keys/defaults, data-only JSON, explicit endpoint/client intent, execution-host separation, correct origin/app/destination identity, confirmation, inline WhichKey rename, once-only effects, fixed invocation bounds/pagination and terminal/popup/private cleanup. No no-op success, restored removed plugin entries, callback support, native rename dialog, extra background resize policy or silent target renewal.
8. **Actual verification and evidence.** Build and run the ordinary public binary from a real owned Rex origin. Record launch-to-first-visible-frame separately from data-ready/actions-ready, comparable baseline/candidate context and limits; no invented latency threshold or unmeasured speed claim. Observe first-frame responsiveness, navigation during/after loading, loaded identities, a benign effect on an explicitly owned disposable resource, normal dismissal/cancellation and exact cleanup. Source review must trace eliminated validation/duplicate work and new concurrency/ownership boundaries. A compile/help/private startup ACK is not completion.
9. **Accurate docs, scope and delivery.** Update active PATH/startup/async/error/lifecycle guidance and current evidence with the reviewed integrated source. Clearly record what remains unverified. Preserve tests, assertions, hooks, lint standards, previous decisions/commits and Phase16/earlier evidence. Phase completion is this delivered implementation with actual verification, not completion of Phase17 or unrelated phases.

## Expected surfaces and workstreams

Inspect rex/runtime.go, rex/exec.go, rex/rex.go, rex/endpoint.go and concrete resource metadata helpers; launch.go/native.go; providers.go/hosts.go; popup.go/runtime_bootstrap.go; terminal.go/menu.go/text-entry handling; normal package/discovery wiring and active launch/data/runtime/acceptance docs. Existing Taskfile/lint/hooks/test assertions are protected. Literal ownership is assigned later; no unrestricted repository rewrite.

Bounded future workstreams: simplify PATH/CLI startup; define a small UI-ready/update contract; implement async provider/label loads and snapshot reuse; integrate actual draw-first startup and lifecycle; independently review and verify normal runtime timing/behavior on combined source. Producers precede consumers. Checks stay with each implementation deliverable, not a separate test/benchmark project. Exact IPC/channel format, per-provider scheduling and timing instrumentation are open implementation details, not already approved APIs.

## Requirements and relationships

Complementary scope REQ-001/002/003/004/006/008/009/010/011/012/013/014/015/016/017, interpreted by the startup supplement and current CLI-only design. Required origin/target/confirmation/cleanup semantics remain; only the rejected executable/compatibility preflight policy is replaced. REQ003 remains data-config loading; executable REQ004 portions and REQ005/hooks remain deferred.

Parent projects/rex-whichkey.md, unique order18. Delivered Phase16 is the source baseline, not a waiting condition on older incomplete work. No native whole-phase prerequisites: do not block on Phase17 audit, Phase15's replaced bridge, or a future performance/research milestone. Audit and startup work must bind their own revision and coordinate if they later touch shared files. Do not reopen Phases11/14 or migrate legacy records.

## Verification restrictions and readiness

Current authority creates this phase/design and records the accepted changes; no task graph, workers, implementation, timing runs or completion claim is created now. A later execution instruction governs dispatch.

The user's "yes, verify" authorizes targeted normal public-app timing/behavior verification for this startup work. It does not reinstate broad benchmarks, unit/race/TDD/Lua/parser/FIFO harnesses, remote experiments, user init/profile edits, config reloads, deployment/push, input updates or global cleanup. Build/format/static checks and independent source/concurrency review remain; inherited global lint failures must be reported honestly and standards not weakened. Only exact explicitly owned resources may be used for benign interaction/effect/cleanup verification. Retain the existing sole coordinator/resource limits if execution starts.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDAyOjQ0OjU0LjU0OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOC9zdW1tYXJ5LzEiLCJpbnRlbnQiOiJmNzAwY2JiMzIwYjA4MWMyOTAwNDYxMDM4NTI0ZDk0M2E0MGFlMDFlNzI5NmNlOWNjODlkOThkNWFjYTk0YTlkIn0 -->
### 2026-11-07T02:44:54.549Z · summary · unspecified

{"acceptedScope":"PATH-only normal binary lookup with rejected executable/hash/status/version/API-schema preflight removed; unnecessary executable-copy durability Sync removed, functional config/result/target/confirmation/cleanup preserved. Actual static first draw, bounded background/on-demand inventories and labels, invocation snapshot/query reuse, immutable parent-retained ready updates/single UI writer and joined lifetime. Active evidence docs current, reviewed integrated source and actual normal public MAIN plus before/after/default/opt-in/owned-rename/cancel/cleanup verified.","approvalSummary":["User accepted startup direction REXWHICHKEY-20261106-191113 then explicitly start next phase with go-team, commit each phase, finish all phases","Current exact approved serial three-task plan and startup/CLI-only governing designs; same existing sole coordinator with bounded autonomy","Each current task independently reviewed on exact latest build, integrated and actual combined public verification before native completion; final main rebuilt/observed; no compile-only closure"],"attempt":1,"cleanup":"All new successful planner/builder/reviewer/integration/baseline worktrees and validation windows removed after stopped consumers/delivery/ancestry/evidence checks. Retained old negative reviewers and Phase17 negative reviewer per policy, no GC/history/profile deletion.","eventID":"rex-whichkey-p18/summary/1","files":[".tasks/phases/rex-whichkey-18-startup-speed.md",".tasks/tasks/rex-whichkey-18-01-path-startup.md",".tasks/tasks/rex-whichkey-18-02-incremental-ui.md",".tasks/tasks/rex-whichkey-18-03-async-inventories.md","README.md","docs/acceptance.md","docs/launch-contract.md","docs/mapping-format.md",".mtn-tmp/p18/phase18-milestone.json",".mtn-tmp/p18/task3-public-evidence.json",".mtn-tmp/p18/final-main-observed-rex-whichkey",".mtn-tmp/p18/final-cleanup-plan.json"],"kind":"summary","limitations":["First-visible single paired upper bounds~2.69s baseline/~0.54s final include CLI/poll latency; no exact internal timestamp, universal speedup, SLA or benchmark statistics","Full unchanged lint62 inherited diagnostics and AI lint not reached, not global green; all prohibited test/harness/race/benchmark suites not run","Darwin/local normal app and owned input/effect verification only, no user configuration/profile/keybinding/deployment/reload/remote/platform claims","Some failure/replay/disconnect/overflow paths source-reviewed rather than artificially forced; hidden prefix visited with intentionally invisible rows but no hidden target dispatched","Fresh independent same-family reviews, not cross-family validation","Known immutable cleanup jq defect handled by exact stopped/clean/ancestry-verified manual task worktree cleanup after native completion; failed reviewer resources/history retained","Historical phases11/14/15 unchanged; Phase17 audit is complete recommendations-only, no automatic removals; separate flake.nix dev-shell fix remains uncommitted"],"recordRef":"phases/rex-whichkey-18-startup-speed.md","records":["tasks/rex-whichkey-18-01-path-startup.md","tasks/rex-whichkey-18-02-incremental-ui.md","tasks/rex-whichkey-18-03-async-inventories.md"],"requirements":["REQ-001","REQ-002","REQ-003","REQ-004","REQ-006","REQ-008","REQ-009","REQ-010","REQ-011","REQ-012","REQ-013","REQ-014","REQ-015","REQ-016","REQ-017"],"schemaVersion":1,"specHash":"sha256:c5dc3f0ed59c1982958b6ada9b24757589089555ad9b4db01407d08bd9dc2e08","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"765265884a1a94c4084d6be4075c89cf5af029a4","targetTree":"45131e6eb08b82ddfacc7689666a61db6e0ed756","verification":[{"artifacts":[".mtn-tmp/p18/final-main-build.log",".mtn-tmp/p18/final-main-vet.log",".mtn-tmp/p18/final-main-root.txt",".mtn-tmp/p18/final-main-inventory.txt",".mtn-tmp/p18/final-main-after.json",".mtn-tmp/p18/final-main-modes.json"],"command":"Final main normal build/vet and actual ordinary public root/dynamic navigation/Escape exit0/terminal/focus/exact popup-private cleanup","result":"pass"},{"artifact":".mtn-tmp/p18/task3-public-evidence.json","command":"Normal public before/after separate first-visible/data/action-ready observations, opt-in providers/loading/IDs/current pagination/fixed bounds/exact owned rename and cleanup","result":"pass"},{"artifacts":[".mtn-tmp/p18/final-hierarchy.json",".mtn-tmp/p18/final-plan-guard-result.json",".mtn-tmp/p18/final-collection-validation.json",".mtn-tmp/p18/final-tasks.json"],"command":"Nonempty native3/3 complete scope, exact approved current task plan/designs,15 requirements/9 criteria, valid acyclic graph and all delivered summary target ancestry","result":"pass"},{"artifacts":[".mtn-tmp/p18/final-protected-source-check.log",".mtn-tmp/p18/final-worktree-inventory.txt",".mtn-tmp/p18/final-private-dir-inventory.txt"],"command":"Protected tests/assertions/dependencies/Taskfile/hooks/lint config unchanged; separate flake.nix fix byte-preserved; exact eligible cleanup and no popup-private dirs","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
