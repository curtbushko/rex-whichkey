---
identity:
  kind: task
  taskID: rex-whichkey-18-02-incremental-ui
  uuid: 4f615659-e678-43f1-badf-b3edf4ce2c78
structure:
  order: 2
scope:
  requirements:
    - REQ-004
    - REQ-006
    - REQ-008
    - REQ-009
    - REQ-010
    - REQ-011
    - REQ-014
    - REQ-015
    - REQ-016
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
      - native.go
      - popup.go
      - runtime_bootstrap.go
      - popup_updates.go
      - menu_updates.go
      - terminal.go
      - menu.go
      - render.go
      - text_entry.go
      - native_text_entry.go
      - docs/runtime-bootstrap.md
      - docs/launch-contract.md
      - docs/tab-rename-input.md
    sharedPaths: []
    resources: []
    interfaces:
      - STARTUP-UI-UPDATES
      - REX-SESSION-VIEW-SNAPSHOT
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
  sourceEventID: rex-whichkey-p18-t02/summary/1
  changedAt: '2026-11-07T01:55:16.402Z'
title: '18-02: Add safe incremental UI and on-demand rename labels'
mtnCreateRetry:
  key: rex-whichkey-p18/create/18-02/v1
  intent: 5189028b138ebcc7e36a91c8cf5c929c94648937e966e34369fa750314a2f20e
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T00:54:24.450Z'
dateModified: '2026-11-07T01:55:17.078Z'
tags:
  - task
type: task
id: 2b732f6c-6c96-410b-84d7-223f0f3e7fe2
projects:
  - '[[phases/rex-whichkey-18-startup-speed.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-18-01-path-startup.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-06'
---
Produce the minimal coordinated controller/renderer update path and exercise it with real on-demand inline-rename presentation. Depend on task18-01's PATH-only invocation context, not legacy executable receipts. Before coding consumers, document the private update contract in docs/launch-contract.md and docs/runtime-bootstrap.md: invocation/load identity, canonical demand kind/prefix, bounded immutable presentation result, loading/unavailable/ready eligibility, parent-retained mapping/target publication before renderer availability, selection binding to the published state, terminal ownership and cancellation/error ordering. Exact channel/file/IPC encoding is engineering detail. Use the existing private invocation transport; no service, framework, registry, direct socket, Lua control path or future-hook scaffold. No fake/no-op producer. Transport and menu consumer land together with the real rename-label consumer so the task is an operational vertical slice, not an independent framework task.

Controller owns authoritative action/args/resource IDs and freezes the existing exact origin/app generation intent; the child supplies only bounded keys/confirmation/text plus minimal update identification if needed. A renderer-visible ready update cannot authorize a selection the parent has not retained. Unknown/stale/malformed/pending selection refuses; no label/list-order/child-supplied replacement target. Workers return detached immutable values to the controller owner; only the renderer UI loop mutates its live menu, status and redraw. No shared live Spec slices/maps/trie nodes, background effects or concurrent terminal writer. Bound workers, queues/messages and every read; no idle menu timeout. Dismissal/cancel/terminal loss/parent lifetime loss cancels all work, joins workers and transport readers/writers, rejects late results and completes bounded exact cleanup before any effect. Propagate join/transport/cleanup errors rather than allowing a selection after uncertain teardown. Extend the explicit fixed cleanup inventory for any added private slots; no recursive scavenging.

Remove eager label fetching in prepareNativeMenu. Resolve session/tab presentation on demand when its configured inline rename form is entered, using any matching invocation snapshot already available. Loading/error text is non-executable status, not a native rename dialog. The form remains cancellable and can accept edits while labels load; submission requires the relevant retained context to be ready. A late label may initialize an untouched form only; never overwrite edits, selection/cursor, text validation problems or pending confirmed text, and never execute queued Enter automatically when data becomes ready. Error/unavailable context blocks submission honestly. Preserve custom full rename keys/confirm behavior, limits, inline text semantics and canonical result validation. Back/partial-key prefix, page and confirmation state survive unrelated updates; updates do not feed or reset the partial terminal escape/UTF-8/paste decoder. Pending/not-ready action keys are not buffered/replayed into effects. Keep one initial popup rectangle and existing pagination/defaults/config schema.

Introduce only the small concrete Rex session-view snapshot helper needed to share the already-required origin view and label/window/block metadata; expose detached exact typed resources, not a generic cache/port layer. Preserve strict existing field/ID/placement validation and ordinary fresh Inspect methods. T3 will consume the same invocation snapshot for inventories. Existing non-label inventories may remain eager in this intermediate revision; document this honestly rather than claim final first-frame behavior. No background inventory implementation or broad resource rewrite is required here.

Acceptance: independent frozen source/concurrency/lifetime review must trace producer-before-consumer retention, single UI writer, bounded cancellation/join, stale-result rejection, unchanged authoritative result matching, rename editing and pre-effect safety. Run owned-file gofmt/whitespace, task build and task lint:vet with the pinned budget; no tests/harnesses. Coordinator combined public run must show real menu, enter an owned-resource rename form, edit/cancel while presentation is pending where ordinarily observable, revisit and complete one benign inline rename on an explicitly disposable owned tab/session, observe the exact renamed ID and cleanup-before-effect, plus ordinary cancellation/dismissal/restoration/private cleanup. Do not artificially delay/stub APIs to manufacture pending state; report any unobservable transition as a runtime limitation supported only by source review. Do not edit user config, reload, use a private ACK, or claim remaining eager inventories are asynchronous.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDAxOjU1OjE1LjUwNFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOC10MDIvc3VtbWFyeS8xIiwiaW50ZW50IjoiZjY0MTRiZmJmZmFiMzM0ZjQ2MDZjMjFjOWFhYmIyYTFiODc3MjM1Njg1NTUwNjIxNmUwMTVhYzllMDc1NzRiNCJ9 -->
### 2026-11-07T01:55:15.504Z · summary · unspecified

{"acceptedScope":"Minimal fixed-slot bounded controller/renderer rename demand/update transport, parent-retained selection binding and single UI writer; real on-demand inline rename context, untouched initialization/edit preservation/ready eligibility; concrete detached session snapshot helper, ordinary fresh Inspect unchanged. Cancel/join/lifetime/owned cleanup preserved. Inventories intentionally remain eager for task3.","approvalSummary":["User startup decision and start-next-go-team/finish-all/phase-commit authority","Current approved Phase18 serial plan and exact producer prerequisite task1 delivery","Fresh same-family combined review1 approved exact head13fe8ca/tree c2087b5; coordinator combined normal build/vet and two actual public runs passed before target advance"],"attempt":1,"cleanup":{"facade":"Official native/protocol completion verified; immutable cleanup jq syntax defect persists. Exact clean stopped delivered ancestor worktrees removed manually after evidence preservation.","removedCleanDeliveredWorktrees":["p18-t02-builder-a01","p18-t02-reviewer-a01","p18-t02-integration-a01"],"runtimeReservationsReleased":true,"successfulWorkerWindowsConfirmedAbsent":true},"eventID":"rex-whichkey-p18-t02/summary/1","files":["popup_updates.go","menu_updates.go","rex/snapshot.go","native.go","popup.go","runtime_bootstrap.go","terminal.go","menu.go","render.go","text_entry.go","native_text_entry.go","rex/block.go","rex/window.go","docs/runtime-bootstrap.md","docs/launch-contract.md","docs/tab-rename-input.md",".mtn-tmp/p18/task2-build-result.json",".mtn-tmp/p18/task2-review-result.json",".mtn-tmp/p18/task2-integration.json",".mtn-tmp/p18/task2-public-evidence.json",".mtn-tmp/p18/task2-observed-rex-whichkey",".mtn-tmp/rex/rex-whichkey-p18-t02-builder-a01",".mtn-tmp/rex/rex-whichkey-p18-t02-reviewer-a01"],"kind":"summary","limitations":["Inventories remain eager until task3; no full async completion","No unit/race/TDD/Lua/parser/FIFO harness/benchmark/fake server or remote/config/reload/deploy trial","Inherited full lint diagnostic remains failing; no standards/assertion changes","Fresh same-family independent source review, not cross-family","Actual pending form displayed Enter disabled and edited text survived ready publication. Not every replay/malformed/error/disconnect branch runtime-forced; no synthetic delay/test harness."],"recordRef":"tasks/rex-whichkey-18-02-incremental-ui.md","requirements":["REQ-004","REQ-006","REQ-008","REQ-009","REQ-010","REQ-011","REQ-014","REQ-015","REQ-016"],"schemaVersion":1,"specHash":"sha256:20ffc690aaef3b181e58c5353f9f658edd0dff80cf9f6b10cab1b30d9c1de15d","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"6c5c9ed67c8adde938a620cd86d55f777f665b18","targetTree":"8f7eebcc7b2b645c61972491c729e32c4802572c","verification":[{"artifacts":[".mtn-tmp/p18/task2-integration-build.log",".mtn-tmp/p18/task2-integration-vet.log"],"command":"Combined normal task build and task lint:vet with admitted pinned budget","result":"pass"},{"artifact":".mtn-tmp/p18/task2-candidate-source-identity.log","command":"Frozen reviewed owned-source equality, combined whitespace, target-base recheck and fast-forward; reviewed head ancestor of main","result":"pass"},{"artifact":".mtn-tmp/p18/task2-public-evidence.json","command":"Actual ordinary public pending-edit/ready text preservation, form cancel no effect, revisit exact owned tab rename, separate root dismissal; both exit0, terminal/focus and exact popup/private cleanup","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
