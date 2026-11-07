---
identity:
  kind: task
  taskID: rex-whichkey-18-03-async-inventories
  uuid: 8a37278a-2bdb-405f-839c-e34d0978548f
structure:
  order: 3
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
implementation:
  language: go
  ownership:
    paths:
      - rex/snapshot.go
      - rex/session.go
      - rex/block.go
      - rex/window.go
      - providers.go
      - hosts.go
      - native.go
      - native_app_owner.go
      - native_targets.go
      - popup.go
      - popup_initial.go
      - popup_updates.go
      - menu_updates.go
      - terminal.go
      - menu.go
      - render.go
      - text_entry.go
      - native_text_entry.go
      - README.md
      - docs/runtime-bootstrap.md
      - docs/launch-contract.md
      - docs/mapping-format.md
      - docs/session-owned-ui.md
      - docs/tab-rename-input.md
      - docs/acceptance.md
    sharedPaths: []
    resources: []
    interfaces:
      - STARTUP-ASYNC-DATA
      - STARTUP-QUERY-REUSE
  gates:
    - ID: whitespace
      command: git diff --check
      required: true
    - ID: source-format
      command: test -z "$(gofmt -l ./*.go ./rex/*.go ./cmd/*.go)"
      required: true
    - ID: normal-build
      command: task build
      required: true
    - ID: go-vet
      command: 'task lint:vet'
      required: true
    - ID: full-lint-diagnostic
      command: task lint
      required: false
lifecycle:
  state: complete
  sourceEventID: rex-whichkey-p18-t03/summary/2
  changedAt: '2026-11-07T02:43:35.644Z'
title: '18-03: Draw first, load inventories asynchronously and share snapshots'
mtnCreateRetry:
  key: rex-whichkey-p18/create/18-03/v1
  intent: 31b209e1e8f8956b3b99d0b4d522d2507aabe73be1707b014aafb1b82ce2700c
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T00:54:26.393Z'
dateModified: '2026-11-07T02:43:36.275Z'
tags:
  - task
type: task
id: 14f843d7-f8c8-44dc-a717-058d6723cabd
projects:
  - '[[phases/rex-whichkey-18-startup-speed.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-18-02-incremental-ui.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-06'
---
Consume the integrated update/retention/snapshot contract to finish draw-first inventories and invocation read reuse. Static validated root and provider-group/loading placeholders must be actually drawn before any nonessential inventory or rename-label wait. Keep required endpoint/origin/placement, strict explicit app intent and original owner/generation capture before those actions can be eligible; no delayed implicit app renewal. Preserve local JSON/config/conflict errors, confirmations and exact cleanup. Move all remaining eager nativeSessionMenu/resolveHostMenu/standalone-terminal inventory calls off first-frame setup. The standalone menu remains UI/result-only, with the same on-demand known-host behavior and no newly invented resource-dispatch context.

Scheduling choice: sessions start asynchronously AFTER a successful first draw only when an enabled visible session-provider marker is present; a hidden enabled sessions marker loads on explicit prefix demand. Windows and panes load on first entry/demand of an enabled provider prefix. Known hosts load only on explicit enabled prefix demand, never automatically: listing-only rex hosts --json with existing execution-host/home semantics and warnings, no added probes/attach/remote trial. Rename labels remain task18-02 form-demand work. Disabled/condition-false/absent/unused providers schedule nothing; hidden still means callable, not disabled. Multiple markers of a kind share one in-flight attempt and terminal result, including errors; no repeated retry on reopen and no persistent cache. Different required independent reads may proceed concurrently under a small fixed bound (at most two metadata jobs); demand/result queues remain finite and invocation-local. Async overflow/conflict/bound failure cannot silently truncate into a plausible ready menu.

Reuse one validated origin session-view snapshot across startup consumers, window/pane rows and current session/tab labels when semantically the same read, rather than Client.Session + Session.Windows/Blocks/Window each reissuing session inspect. Session-list/known-host loads each have their own once-per-invocation snapshot; publish immutable results and corresponding exact targets before enabling child choices. Seed cached origin metadata only from the captured matching endpoint/session/placement, never a stale result from another invocation. Identify eliminated duplicate reads in review. Do NOT route current Inspect, nativeResourceOrigin, validatePaneAppOrigin, validateSessionChoice, validatePaneChoice, pre-create placement/geometry consistency, liveness or exact cleanup through the presentation cache: actual pre-effect origin, GUI generation/owner and destination freshness checks still read current state where required. Keep action metadata deferred until selection. No arbitrary expiration/refresh, endpoint retarget, uncertain retry or daemon.

Apply completed provider batches atomically through the existing owner/UI loop. Show loading, empty, warning/unavailable and ready distinctly; pending/error rows are not executable. Preserve full IDs, retained target associations, ordering/current markers, per-provider 512-resource limit, existing 1024-child/key-byte/4-MiB combined bounds, known-host limits/fields/warnings, hidden/condition semantics and More/pagination behavior. Do not resize an existing layer as results arrive. Preserve partial menu prefix/page/confirmation, active rename value/cursor and partial terminal decoding; no stale key reindexing may select a different resource. Reject updates after selection/dismissal or canceled load generation, and join all work/IPC before terminal restoration/result/owned cleanup/effect ordering permits completion. Integrate docs for async/error/readiness/demand/cache semantics and exact cleanup; update current acceptance evidence without overwriting historical Phase16 evidence. README's stale source-only status should reflect only actually verified current delivery, never fabricated full-platform/keybinding/installation coverage.

Acceptance: independent frozen source review traces first-draw ordering, all provider paths, zero unused loads, query reuse versus required freshness, immutable handoff/parent retention, one UI writer and every cancellation/join/error path. Run owned-file gofmt/whitespace, normal task build and task lint:vet; optional unchanged full lint reports actual inherited failures without suppressions. No tests/TDD/race/Lua/parser/FIFO harnesses, benchmark suites, synthetic servers/delays/hooks or new dependencies.

COORDINATOR FINAL INTEGRATION: on the reviewed combined source, repeat the exact ordinary public no-argument startup observation against task18-01's baseline using comparable owned origin, host/PATH/endpoint, config/defaults, GUI intent and terminal context. Record launch-to-first-visible-static-frame, required-data-ready and relevant-actions-ready separately with observation limits; no approved numeric SLA and no prior measured timing. Observe input/dismissal responsiveness during normally occurring loading, provider navigation during/after loads, full loaded identities/current markers/pagination, correct relevant rename behavior, one benign selected effect on an exact owned disposable resource, normal dismissal and cancellation, no duplicate dispatch, worker/child exit, terminal restoration, origin focus and exact returned popup/private-slot/directory cleanup before effects. Use an existing authorized immutable data config if available for nondefault providers; otherwise a scoped disposable owned data file outside user config is permitted by the existing data interface, but never user profile/init edits or reload. Keep the baseline-comparable default observation separate from optional-provider cases. Do not demand remote hosts, destructive user-resource trials or a custom benchmark/harness to exercise unavailable paths. Record actual CLI/API failures and unobservable loading/error branches honestly; compile/help/private ACK does not prove display or safety. Coordinator alone records runtime evidence, integrates and commits the completed phase, preserving flake.nix and independent Phase17 work. Failure of a required real-menu/effect/cleanup check leaves delivery pending rather than converting it to a source-only pass.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDAyOjQzOjM0LjY3OFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOC10MDMvc3VtbWFyeS8yIiwiaW50ZW50IjoiMzkzMjEwYjUyMjQzYjVmNTAxOWQyMDAyZjUwYmY0YTA2NGNlMGU0N2ZiNzRhYTc5MmY4OTczZDVlMzQwNWEyOCJ9 -->
### 2026-11-07T02:43:34.678Z · summary · unspecified

{"acceptedScope":"Actual draw-first static menu; bounded after-draw visible Sessions and on-demand window/pane/host/hidden data, invocation snapshot/in-flight/error reuse, parent-retained ready selections, one UI writer and joined cancellation/exact cleanup. Current evidence docs refreshed after real verification with all260 Go blobs unchanged. Reviewed combined and MAIN public runs, comparable default readiness/timing, opt-in providers/loading/pagination/fixed bounds/owned rename verified.","approvalSummary":["User REXWHICHKEY-20261106-191113 startup direction and explicit go-team finish-all/phase-commit execution authority","Current exact Phase18 serial task plan and delivered producer prerequisites","Fresh same-family review1 approved d60 Go implementation; actual combined candidate397 public evidence; docs-only refresh and fresh review2 approved exact fca99c6/tree8aef58c; all Go blobs identical and combined/main build/runtime passed before completion"],"attempt":2,"cleanup":{"facade":"Official native/protocol completion verified; immutable cleanup jq syntax defect persists. Exact clean stopped delivered ancestor worktrees manually removed; original observed candidate ancestry preserved with tree-identical merge. Referenced source/binaries/logs retained.","removedCleanDeliveredWorktrees":["p18-t03-builder-a01","p18-t03-reviewer-a01","p18-t03-builder-a02","p18-t03-reviewer-a02","p18-t03-integration-a01","p18-t03-integration-a02","p18-baseline-observation"],"runtimeReservationsReleased":true,"successfulWorkerWindowsConfirmedAbsent":true},"eventID":"rex-whichkey-p18-t03/summary/2","files":["README.md","docs/acceptance.md","docs/launch-contract.md","docs/mapping-format.md","docs/runtime-bootstrap.md","docs/session-owned-ui.md","docs/tab-rename-input.md","hosts.go","menu.go","menu_updates.go","native.go","native_text_entry.go","popup.go","popup_initial.go","popup_updates.go","providers.go","rex/snapshot.go","terminal.go",".mtn-tmp/p18/task3-build-result.json",".mtn-tmp/p18/task3-review-result.json",".mtn-tmp/p18/task3-doc-build-result.json",".mtn-tmp/p18/task3-doc-review-result.json",".mtn-tmp/p18/task3-integration.json",".mtn-tmp/p18/task3-public-evidence.json",".mtn-tmp/p18/paired-baseline-timing.json",".mtn-tmp/p18/paired-final-timing.json",".mtn-tmp/p18/final-main-root.txt",".mtn-tmp/p18/final-main-inventory.txt",".mtn-tmp/p18/final-main-observed-rex-whichkey",".mtn-tmp/p18/task3-preserved-worker-artifacts"],"kind":"summary","limitations":["Coarse single paired CLI/poll observed upper bounds2.69s baseline vs0.54s final, no exact render/internal readiness mean/SLA/benchmark","No prohibited tests/harnesses/benchmarks/fake delays/user config/reload/activation/deploy/push/remote/platform/physical-keybinding trials","Full unchanged lint62 inherited diagnostics/AI lint not reached, not global green","Hidden provider prefix accessed but rows intentionally hidden and no hidden destination dispatched; malformed/error/disconnect/overflow branches source-reviewed rather than artificially forced","Fresh independent same-family review; separate flake.nix change preserved"],"recordRef":"tasks/rex-whichkey-18-03-async-inventories.md","requirements":["REQ-001","REQ-002","REQ-003","REQ-004","REQ-006","REQ-008","REQ-009","REQ-010","REQ-011","REQ-012","REQ-013","REQ-014","REQ-015","REQ-016","REQ-017"],"schemaVersion":1,"specHash":"sha256:28a76651cc9de862e674e323386e831aa2a91069ce3e584b9df34e5f4b7dd3c6","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"765265884a1a94c4084d6be4075c89cf5af029a4","targetTree":"45131e6eb08b82ddfacc7689666a61db6e0ed756","verification":[{"artifacts":[".mtn-tmp/p18/task3-doc-candidate-source-equality.log",".mtn-tmp/p18/task3-doc-observed-code-equality.log",".mtn-tmp/p18/task3-doc-preserve-tree-equality.log"],"command":"Exact combined reviewed source/Go identity, tree-identical preservation merge, whitespace/target-base check/fast-forward and reviewed head ancestor","result":"pass"},{"artifacts":[".mtn-tmp/p18/task3-doc-integration-build.log",".mtn-tmp/p18/task3-doc-integration-vet.log",".mtn-tmp/p18/final-main-build.log",".mtn-tmp/p18/final-main-vet.log"],"command":"Final combined candidate and MAIN normal task build and task lint:vet","result":"pass"},{"artifact":".mtn-tmp/p18/task3-public-evidence.json","command":"Actual comparable baseline/final ordinary public first-frame and separate inventory/rename-ready observations, opt-in provider/loading/pagination/currentID/fixed-bounds/owned rename/cancel/cleanup","result":"pass"},{"artifacts":[".mtn-tmp/p18/final-candidate-root.txt",".mtn-tmp/p18/final-candidate-after.json",".mtn-tmp/p18/final-main-root.txt",".mtn-tmp/p18/final-main-inventory.txt",".mtn-tmp/p18/final-main-after.json",".mtn-tmp/p18/final-main-modes.json"],"command":"Final documentation-updated candidate AND rebuilt MAIN ordinary public root/dynamic navigation/root dismissal exit0, primary terminal/focus and exact owned floating/private cleanup; windows closed","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
