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
  state: incomplete
  sourceEventID: rex-whichkey-p18-planner-a01/runtime/launched
  changedAt: '2026-11-07T00:42:59.505Z'
title: 'Phase 18: Faster startup and asynchronous menu data'
mtnCreateRetry:
  key: rex-whichkey-phase18-startup-speed-create-v1
  intent: 2d016db0f24e932bb1167398c657a9ae4a3a6d15350ff1f4b299d1ccae598d0c
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T00:16:20.329Z'
dateModified: '2026-11-07T00:43:00.119Z'
tags:
  - task
type: task
id: bf64e407-caf0-4537-bff8-638b6c6afae8
projects:
  - '[[projects/rex-whichkey.md]]'
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

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDAwOjE3OjE2Ljg0OFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMTgvZGVmaW5pdGlvbi1hcHByb3ZhbC8xIiwiaW50ZW50IjoiZDEwZDU5NmY5MzM3MDU3MjdlOTI4Yjc2OGNkMGY2YTRhMWIyNTExZDlmNDU4NDZlYmQ1MTk5ZTM0OTJiZmQ5YyJ9 -->
### 2026-11-07T00:17:16.848Z · approval · unspecified

{"approvalScope":"One startup implementation phase definition/planning, not execution or task dispatch. Exact async channel/IPC, executable-copy removal and numerical timing target not predetermined; preserved configuration/action/target/cleanup semantics and existing no-broad-test/deployment boundaries remain.","approvedBy":"user","attempt":1,"decision":"approved","decisionEvidence":"Explicit startup item approvals (PATH-only/no validation, background/on-demand, async, duplicate-query removal and verification); after executable fsync/UI update clarification, user requested add these changes as a new phase.","decisionID":"REXWHICHKEY-20261106-191113","eventID":"rex-whichkey-p18/definition-approval/1","kind":"approval","recordRef":"phases/rex-whichkey-18-startup-speed.md","schemaVersion":1,"specHash":"sha256:c5dc3f0ed59c1982958b6ada9b24757589089555ad9b4db01407d08bd9dc2e08"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDAwOjIzOjMxLjczMFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOC9wbGFubmluZy12YWxpZGF0aW9uLzEiLCJpbnRlbnQiOiJhMjQ1MTdkNWM4ZDQwNzFlYWYwZmNjYmFlYTc4OTE4Nzk0ZWM0ZWY2OTEwZTVlMTRlY2MxZGQwOGZmNDliZTgzIn0 -->
### 2026-11-07T00:23:31.730Z · historical · unspecified

{"attempt":1,"content":"{\"checks\":{\"phase\":\"phases/rex-whichkey-18-startup-speed.md\",\"order\":18,\"nativeStatus\":\"open\",\"childCount\":0,\"hierarchy\":{\"ref\":\"phases/rex-whichkey-18-startup-speed.md\",\"parents\":[{\"raw\":\"[[projects/rex-whichkey.md]]\",\"resolution\":{\"ref\":\"projects/rex-whichkey.md\"}}],\"children\":[],\"descendants\":[],\"parentStatus\":\"open\",\"progress\":{\"completed\":0,\"total\":1,\"percentage\":0},\"leafTasks\":[\"phases/rex-whichkey-18-startup-speed.md\"],\"archivedPolicy\":\"exclude-records-and-branches\",\"cycles\":[],\"snapshot\":\"2026-11-07T00:21:29.299Z\"},\"dependencies\":{\"ref\":\"phases/rex-whichkey-18-startup-speed.md\",\"status\":\"open\",\"blockedBy\":[],\"blocking\":[],\"nativeBlocked\":false,\"dependencyReady\":true,\"agentEligible\":null,\"activity\":null,\"policy\":null,\"dependencyPolicy\":{\"source\":\"tasknotes-spec defaults\",\"treat_missing_target_as_blocked\":true,\"require_resolved_uid_on_write\":false,\"enforce_unique_uid\":true,\"unresolved_target_severity\":\"warning\",\"default_reltype\":\"FINISHTOSTART\"},\"reasons\":[],\"blockerExplanation\":null,\"cycles\":[],\"snapshot\":\"2026-11-07T00:21:30.284Z\"},\"uniqueSiblingOrders\":true,\"highestOrder\":18,\"phaseApproval\":true,\"designApproval\":true,\"auditDefinitionUnchanged\":true,\"auditNativeStatus\":\"open\"},\"collectionValidation\":\"Native agent validate: valid core, no hierarchy or dependency cycles.\",\"projectPlanPersistence\":\"Legacy project facade returned unknown workflow record kind after persistence; exact native event equality readback verified. No retry or duplicate write.\",\"representationCorrection\":\"Initial private order check read only typed structure.order and omitted legacy workflowOrder; normalized old/new native representation and rechecked all sibling orders unique. No record rewrite required.\",\"scope\":\"Phase18 open/planned, order18 under existing project, no children/prerequisites/workers; separate Phase17 audit definition fingerprint unchanged. New startup/phase approvals bound to current exact fingerprints.\",\"notRun\":[\"implementation/source changes\",\"startup timing/public app run\",\"tests/benchmarks\",\"user config/reload/deploy/push\"],\"preserved\":\"Uncommitted flake.nix Darwin dev-shell dependency fix remains outside planning commit; prior delivery/failure history retained.\"}","eventID":"rex-whichkey-p18/planning-validation/1","gateEligible":false,"kind":"historical","provenance":"Sole coordinator native creation/readback/collection validation; planning only, not implementation or runtime evidence.","recordRef":"phases/rex-whichkey-18-startup-speed.md","schemaVersion":1,"specHash":"sha256:c5dc3f0ed59c1982958b6ada9b24757589089555ad9b4db01407d08bd9dc2e08"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDAwOjM3OjU5LjM4NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMTgvZXhlY3V0aW9uLWF1dGhvcml0eS8xIiwiaW50ZW50IjoiMzEwNzZjMjE2YjIwYWU4MWNhMDA0NDE1ZDdjY2IyMjM1YmNhZWI3ZjVjZjE2MmUwZTc2ZjFhNDk3NGVhMjllZiJ9 -->
### 2026-11-07T00:37:59.385Z · approval · unspecified

{"approvalScope":"Execute approved remaining Phase17 audit and Phase18 startup constraints end-to-end; routine bounded task decomposition, fresh independent reviews, serial isolated integration and phase commits authorized. Existing sole coordinator retained, no reopening historical superseded Phase15/11/14 work or general cleanup/removals, no test/benchmark/deploy/config restrictions lifted.","approvedBy":"user","attempt":1,"autonomy":"Bounded task/phase planning and routine fix/integration transitions inside approved definitions; ask only concrete material blockers.","decision":"approved","decisionEvidence":"start the next phase with go-team. commit after each phase. Finish all phases","eventID":"rex-whichkey-p18/execution-authority/1","kind":"approval","recordRef":"phases/rex-whichkey-18-startup-speed.md","schemaVersion":1,"specHash":"sha256:c5dc3f0ed59c1982958b6ada9b24757589089555ad9b4db01407d08bd9dc2e08"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDAwOjQyOjU1LjU1NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOC1wbGFubmVyLWEwMS9ydW50aW1lL3Jlc2VydmVkIiwiaW50ZW50IjoiNzYwMDExZGZiZDE1OTI3ZTIxMzMxMmM3MGI5ODZjNDlmZTZlZDZiNzkwNmMzNzEyOGQzMjE5OWRlMzc1MGM4NiJ9 -->
### 2026-11-07T00:42:55.555Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-15-code-owned-runtime.md","runID":"rex-whichkey-p15-coordinator-a01"},"eventID":"rex-whichkey-p18-planner-a01/runtime/reserved","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"75ff9ad4-a558-4295-8b20-855dafe77a8f"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-18-startup-speed.md","piSessionID":"rex-whichkey-p18-planner-a01","piSessionName":"p18-planner-a01","recordRef":"phases/rex-whichkey-18-startup-speed.md","rexSessionName":"rex-whichkey-main","rexWindowName":"p18-planner-a01","role":"planner","runID":"rex-whichkey-p18-planner-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"mtn-tasks","specHash":"sha256:c5dc3f0ed59c1982958b6ada9b24757589089555ad9b4db01407d08bd9dc2e08","stage":"reserved","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p18-planner-a01","windowName":"p18-planner-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p18-planner-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDAwOjQyOjU4LjgwOVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOC1wbGFubmVyLWEwMS9ydW50aW1lL2xhdW5jaGVkIiwiaW50ZW50IjoiMjE5ZDAxZTYwN2YzMTY1NGQ4YjY1MzQ2M2Y3NTVhYmQ4ZGVjNDU4YTI5YTFkZGMxNmRiZjYxMTBlMmUxMzFjOCJ9 -->
### 2026-11-07T00:42:58.809Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-15-code-owned-runtime.md","runID":"rex-whichkey-p15-coordinator-a01"},"eventID":"rex-whichkey-p18-planner-a01/runtime/launched","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"75ff9ad4-a558-4295-8b20-855dafe77a8f"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-18-startup-speed.md","piSessionID":"rex-whichkey-p18-planner-a01","piSessionName":"p18-planner-a01","provider":{"displayName":"p18-planner-a01","name":"rex","resources":{"block":"block:01a1b375-6201-74de-8bb7-36df60ebd304","endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","session":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","window":"window:01a1b375-6202-7781-bfcf-aa601ce6cb17"}},"recordRef":"phases/rex-whichkey-18-startup-speed.md","rexBlockID":"block:01a1b375-6201-74de-8bb7-36df60ebd304","rexSessionID":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","rexSessionName":"rex-whichkey-main","rexWindowID":"window:01a1b375-6202-7781-bfcf-aa601ce6cb17","rexWindowName":"p18-planner-a01","role":"planner","runID":"rex-whichkey-p18-planner-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"mtn-tasks","specHash":"sha256:c5dc3f0ed59c1982958b6ada9b24757589089555ad9b4db01407d08bd9dc2e08","stage":"launched","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p18-planner-a01","windowName":"p18-planner-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p18-planner-a01"}
<!-- mtn:end-entry:v1 -->
