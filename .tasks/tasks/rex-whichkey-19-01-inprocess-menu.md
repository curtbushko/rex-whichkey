---
identity:
  kind: task
  taskID: rex-whichkey-19-01-inprocess-menu
  uuid: a92251c4-eb13-4f3e-bff7-2ebda111e7a3
structure:
  order: 1
scope:
  requirements:
    - REQ-001
    - REQ-003
    - REQ-004
    - REQ-006
    - REQ-008
    - REQ-009
    - REQ-010
    - REQ-011
    - REQ-012
    - REQ-025
    - REQ-028
    - REQ-029
  documents:
    - ref: designs/rex-whichkey-lua-float-launch.md
      hash: 'sha256:a3657116206eb455b158b9727099b7c3859882d18f759171205eb830dda163c8'
implementation:
  language: go
  ownership:
    paths:
      - terminal.go
      - menu_updates.go
      - providers.go
      - menu.go
      - text_entry.go
      - render.go
      - theme.go
      - rex/snapshot.go
      - direct_menu.go
    sharedPaths: []
    resources: []
    interfaces:
      - P19-INPROCESS-MENU
      - P19-FILTERED-ORIGIN-SNAPSHOT
      - P19-OWN-PTY
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
  state: blocked
  sourceEventID: rex-whichkey-p19-t01/release/1
  changedAt: '2026-11-07T04:56:19.464Z'
title: 'Phase19.1: In-process menu and own-PTY producer'
mtnCreateRetry:
  key: rex-whichkey-p19/task-create/inprocess-menu/1
  intent: 5bc882008f001efa9d437bbc293dc95ddaebfeb84d5f7e670e0c74bc3173787f
status: in-progress
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T04:48:21.351Z'
dateModified: '2026-11-07T04:56:20.088Z'
tags:
  - task
type: task
id: 148d116e-23cd-4431-9fa0-6568fe058ff6
projects:
  - '[[phases/rex-whichkey-19-lua-float-launch.md]]'
---
Implement concrete in-process menu support without activating a new launcher/effect path. Reuse bounded metadataJobs: at most two workers, five once-only jobs and bounded immutable results; only the UI loop changes menu/retained targets. Add a concrete origin-aware menu entry that returns the retained config, MenuEvent and exact provider targets after cancel/join and terminal restoration. Keep existing public/standalone and current private-child behavior compatible until task19.2 removes the latter; no placeholder stubs or new transport framework. Read actual own PTY size/palette, retain SIGWINCH, fixed float rectangle/pagination, rename editing/readiness, confirmations, Unicode, defaults/JSON, limits and theme behavior. Derive original current markers from retained original IDs, not popup focus. Provide a filtered post-create origin snapshot and provider inventories excluding the exact own popup block, including counts/limits/targets. Shared raw view may be used solely for own-popup validation. No Go float geometry RPCs on the new path. Safe to prepare/review/build independently; do not claim direct-launch/runtime lifecycle delivery from this source-only producer. Coordinator's later combined normal make run must verify this task's menu contribution, with evidence attributed back to this task, before phase runtime acceptance. Existing RunTerminal/production callers must meaningfully exercise any extracted core. No exported origin-aware API with only a future blocked-task/test caller may be called delivered; if safe current integration cannot meet that invariant, return the exact blocker rather than creating dormant scaffolding. Current public result-only menu must never dispatch resources.
<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA0OjUxOjExLjA0NloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiZGlzcGF0Y2giLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMTktdDAxL2Rpc3BhdGNoLzEiLCJpbnRlbnQiOiI2YmEzYTI3ODNlY2M4MDZmMGRhNTY4MTVmOTNmZDA4ZjIzOGZjZmQ3ZGM0YmEzYmJiYWMyNGE0Njc0NTgwMmE0In0 -->
### 2026-11-07T04:51:11.046Z · dispatch · unspecified

{"allowedCommands":["Current-used safe in-process core only; no blocked direct launch/effect activation","owned source/format/whitespace","task build","task lint:vet","optional unchanged lint diagnostic","per-command hooksPath=/dev/null owned commit"],"allowedScope":{"interfaces":["P19-INPROCESS-MENU","P19-FILTERED-ORIGIN-SNAPSHOT","P19-OWN-PTY"],"paths":["terminal.go","menu_updates.go","providers.go","menu.go","text_entry.go","render.go","theme.go","rex/snapshot.go","direct_menu.go"],"resources":[],"sharedPaths":[]},"attempt":1,"baseCommit":"c4a054d5f53136ec199b9845a983ef14821057c6","branch":"work/p19-t01-builder-a01","constraints":["Material lifetime exception and native chrome deferral remain unapproved","No dormant producer APIs/tests-only calls/stub framework; report concrete blocker if no meaningful current integration","All assertions/standards/config/history protected; separate flake.nix preserved","Coordinator alone writes/integrates/actual normal make run; no tests/harnesses/benchmark/live/config/reload/deploy"],"eventID":"rex-whichkey-p19-t01/dispatch/1","kind":"dispatch","limits":{"builders":1,"cpu":14,"gates":1,"goMemoryLimit":"5GiB","integrations":1,"reviewers":1,"totalMemoryGiB":6},"recordRef":"tasks/rex-whichkey-19-01-inprocess-menu.md","runID":"rex-whichkey-p19-t01-builder-a01","schemaVersion":1,"specHash":"sha256:c35544123e5829bfc193948c34e6069452501da2dd122754ccf0610de058c5b1","targetBranch":"main","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t01-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA0OjUxOjE5Ljc1NFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDEtYnVpbGRlci1hMDEvcnVudGltZS9yZXNlcnZlZCIsImludGVudCI6IjM1NGM3MTExYTNjZTA4YjY2NTY5MTI0NzVkZmIwMWIxODAyOGZhY2NkZjJjMzFkOGJjMjIwZGYyNjMwN2I2MDAifQ -->
### 2026-11-07T04:51:19.754Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-15-code-owned-runtime.md","runID":"rex-whichkey-p15-coordinator-a01"},"dispatchEventID":"rex-whichkey-p19-t01/dispatch/1","eventID":"rex-whichkey-p19-t01-builder-a01/runtime/reserved","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"75ff9ad4-a558-4295-8b20-855dafe77a8f"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-19-lua-float-launch.md","piSessionID":"rex-whichkey-p19-t01-builder-a01","piSessionName":"p19-t01-builder-a01","recordRef":"tasks/rex-whichkey-19-01-inprocess-menu.md","rexSessionName":"rex-whichkey-main","rexWindowName":"p19-t01-builder-a01","role":"builder","runID":"rex-whichkey-p19-t01-builder-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-builder","specHash":"sha256:c35544123e5829bfc193948c34e6069452501da2dd122754ccf0610de058c5b1","stage":"reserved","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p19-t01-builder-a01","windowName":"p19-t01-builder-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t01-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA0OjUxOjIzLjI3MloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDEtYnVpbGRlci1hMDEvcnVudGltZS9sYXVuY2hlZCIsImludGVudCI6IjQ3NzI2NzNhMmIyNjM5N2VkY2E0ZWQ2NjkxMGU0YzkxZGFlMDU4MWUwZDgyNjRjMTRkMmViMmQyMjJlOTBhNDAifQ -->
### 2026-11-07T04:51:23.272Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-15-code-owned-runtime.md","runID":"rex-whichkey-p15-coordinator-a01"},"dispatchEventID":"rex-whichkey-p19-t01/dispatch/1","eventID":"rex-whichkey-p19-t01-builder-a01/runtime/launched","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"75ff9ad4-a558-4295-8b20-855dafe77a8f"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-19-lua-float-launch.md","piSessionID":"rex-whichkey-p19-t01-builder-a01","piSessionName":"p19-t01-builder-a01","provider":{"displayName":"p19-t01-builder-a01","name":"rex","resources":{"block":"block:01a1b458-cdd8-7e29-a47f-fa46333c25ae","endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","session":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","window":"window:01a1b458-cdda-737a-8001-0fbc517c7966"}},"recordRef":"tasks/rex-whichkey-19-01-inprocess-menu.md","rexBlockID":"block:01a1b458-cdd8-7e29-a47f-fa46333c25ae","rexSessionID":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","rexSessionName":"rex-whichkey-main","rexWindowID":"window:01a1b458-cdda-737a-8001-0fbc517c7966","rexWindowName":"p19-t01-builder-a01","role":"builder","runID":"rex-whichkey-p19-t01-builder-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-builder","specHash":"sha256:c35544123e5829bfc193948c34e6069452501da2dd122754ccf0610de058c5b1","stage":"launched","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p19-t01-builder-a01","windowName":"p19-t01-builder-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t01-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA0OjU2OjA0LjgzNVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYnVpbGQiLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMTktdDAxL2J1aWxkLzEiLCJpbnRlbnQiOiI3ZTI1N2E5NTIxNzJhMDJmMTFiMzlhNTEzZTRmMTE0OTg2ZWUxZDc0ZWU5ZjY0ODBiNjY0MzM1M2QxNzBiZTU3In0 -->
### 2026-11-07T04:56:04.835Z · build · unspecified

{"attempt":1,"baseCommit":"c4a054d5f53136ec199b9845a983ef14821057c6","blocker":{"evidence":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p19-t01-builder-a01/source-inspection.txt","minimumCoordinatorAction":"Provide a narrowly authorized current-consumer handoff and literal ownership preserving private-child authority and all-owned-resource-before-effects ordering, or keep task1 blocked until task19.2's material cleanup conflict permits a safe consumer. Relevant existing consumer surfaces are popup.go, native.go, popup_updates.go and runtime_bootstrap.go; none were edited.","missingSafeAlternative":"An owned private-loop extraction alone cannot provide original-ID markers, a post-create exact-own-popup-filtered snapshot and a real retained-target consumer. A future-only exported entry is prohibited. Inferring origin from popup environment/focus or transferring parent authority to the child would violate current semantics; direct activation/detach remains prohibited.","reason":"RunTerminal's standalone caller is result-only and supplies no retained native origin. Its private-child caller supplies presentation transport and frozen CLI/endpoint, while the parent owns authoritative config/provider targets. Neither caller can meaningfully consume the requested origin-aware Config/MenuEvent/targets result within task1 ownership."},"branch":"work/p19-t01-builder-a01","commits":[],"dispatchEventID":"rex-whichkey-p19-t01/dispatch/1","eventID":"rex-whichkey-p19-t01/build/1","files":{"created":[],"deleted":[],"modified":[]},"head":"c4a054d5f53136ec199b9845a983ef14821057c6","integrationChecks":[{"check":"Independent review and combined-source build/vet of eventual implementation","status":"pending"},{"check":"Coordinator actual normal make run through Lua using the current checkout binary","status":"pending"},{"check":"Own-PTY palette/SIGWINCH/fixed viewport/pagination/Unicode/rename readiness/confirmation behavior","status":"pending"},{"check":"Original-ID current markers, exact own-popup exclusion, bounded jobs and retained exact targets","status":"pending"},{"check":"Cancel/join/restoration, all-owned-resource cleanup before effects and final resource/focus cleanup","status":"pending"}],"kind":"build","limitations":["No task1 implementation delivered; no new symbols or meaningfulCurrentCallers entries.","Tests/unit/race/TDD/Lua/parser/FIFO harnesses, benchmarks and live experiments were not run, as prohibited.","Optional full lint was not run; no current full-lint pass claimed and no lint configuration copied or changed.","No launcher activation, Rex actions, agents, integration, mtn writes or user configuration changes."],"meaningfulCurrentCallers":[],"outcome":"blocked","phaseRef":"phases/rex-whichkey-19-lua-float-launch.md","recordRef":"tasks/rex-whichkey-19-01-inprocess-menu.md","runID":"rex-whichkey-p19-t01-builder-a01","schemaVersion":1,"scopeLifetimePreservation":{"activePaths":"Unchanged controller, standalone result-only menu, private-child transport/authority, launch/effect behavior and historical cleanup guard.","phaseBlockers":"Task19.2 cleanup-order conflict and task19.4 unsupported native chrome remain unwaived. No detach exception, chrome substitute or controller-free delivery claimed.","protectedFiles":"No task records, tests/assertions, hooks, Taskfile, lint configuration, dependencies or flake.nix changed. Parent unrelated Darwin fix untouched.","trackedWorktree":"Clean; HEAD and tree unchanged; no commit fabricated."},"specHash":"sha256:c35544123e5829bfc193948c34e6069452501da2dd122754ccf0610de058c5b1","summary":"No safe current origin-aware producer consumer exists within literal ownership and unchanged private-child semantics. No source edits or commit; unchanged-source required gates passed.","tree":"53832c57cd4cfba436608d5c85f6b41eb6017da7","validation":[{"artifact":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p19-t01-builder-a01/whitespace.log","command":"git diff --check","detail":"Exit 0 on unchanged source.","result":"pass"},{"artifact":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p19-t01-builder-a01/source-format.log","command":"test -z \"$(gofmt -l ./*.go ./rex/*.go ./cmd/*.go)\"","detail":"Underlying gofmt exited 0 with empty output; equivalent empty-output check passed. No files reformatted.","result":"pass"},{"artifact":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p19-t01-builder-a01/build.log","command":"task build","detail":"Exit 0; unchanged-source darwin-arm64 build. Not implementation or runtime verification.","result":"pass"},{"artifact":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p19-t01-builder-a01/vet.log","command":"task lint:vet","detail":"Exit 0; unchanged-source go vet ./.... No test execution.","result":"pass"}],"validationEnvironment":{"GOMAXPROCS":14,"GOMEMLIMIT":"5GiB","GOTOOLCHAIN":"local","PATHPrefix":"/nix/store/dx7xl254pz6c99rmslmaz8hxzfich3m7-go-1.26.8/bin:/nix/store/858pfyzissfrlcspn1whpcvlcyhcdrn3-golangci-lint-2.15.0/bin","heavyGates":"Sequential; one gate at a time."},"worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t01-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA0OjU2OjA4LjMwOVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYmxvY2tlciIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDEvY29uc3VtZXItYmxvY2tlci8xIiwiaW50ZW50IjoiNzIzYzRjZTUxYzdiMTc0MDk0NDg4MGI3MWViMTAxY2RiZDlmNjU4NDY1YzljYzk1M2ExZDNjNGRjZjE0NzJlMyJ9 -->
### 2026-11-07T04:56:08.309Z · blocker · unspecified

{"attempt":1,"eventID":"rex-whichkey-p19-t01/consumer-blocker/1","evidence":".mtn-tmp/p19/task1-build-result.json","kind":"blocker","reason":"No meaningful current origin-aware retained config/target consumer within safe producer ownership without crossing blocked direct-lifetime boundary. No source edits/commit; do not add dormant future-only API to claim implementation.","recordRef":"tasks/rex-whichkey-19-01-inprocess-menu.md","relatedBlocker":"rex-whichkey-p19-t02/lifetime-blocker/1","requiredAction":"Resolve lifetime ordering decision for a safe direct consumer, then coordinator can replan/widen literal producer-consumer ownership without another routine approval. Do not infer permission to weaken cleanup or add a controller.","schemaVersion":1,"specHash":"sha256:c35544123e5829bfc193948c34e6069452501da2dd122754ccf0610de058c5b1"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA0OjU2OjExLjc3NFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVsZWFzZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDEtYnVpbGRlci1hMDEvcmVsZWFzZS9yZXNlcnZlZCIsImludGVudCI6IjlhNDk4NmMxN2ExNWExMzE4ZWVlMjE2OWI0MzIxZTQ0YWY5ZjczZDZhNTNlYzY2NWRkNjk2ZDBjNDYwM2JmZTQifQ -->
### 2026-11-07T04:56:11.774Z · release · unspecified

{"attempt":1,"dispatchEventID":"rex-whichkey-p19-t01-builder-a01/runtime/reserved","eventID":"rex-whichkey-p19-t01-builder-a01/release/reserved","evidence":"Trusted zero-exit BLOCKED build retrieved; no source edits/commit, HEAD/tree unchanged, no live writer. Blocked window/worktree retained for recovery, not successful-resource cleanup.","kind":"release","outcome":"finished","recordRef":"tasks/rex-whichkey-19-01-inprocess-menu.md","runID":"rex-whichkey-p19-t01-builder-a01","schemaVersion":1,"specHash":"sha256:c35544123e5829bfc193948c34e6069452501da2dd122754ccf0610de058c5b1"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA0OjU2OjE1LjE4MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVsZWFzZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDEtYnVpbGRlci1hMDEvcmVsZWFzZS9sYXVuY2hlZCIsImludGVudCI6ImI2ZTAyNjNkOTgyMmRhYmQyMDM1MzIxMDk0ZmZjOWFlYzJlNDQ5MWE5NmEwZjgyMTIxOTVhYmFlNTU2NzMwYjMifQ -->
### 2026-11-07T04:56:15.181Z · release · unspecified

{"attempt":1,"dispatchEventID":"rex-whichkey-p19-t01-builder-a01/runtime/launched","eventID":"rex-whichkey-p19-t01-builder-a01/release/launched","evidence":"Trusted zero-exit BLOCKED build retrieved; no source edits/commit, HEAD/tree unchanged, no live writer. Blocked window/worktree retained for recovery, not successful-resource cleanup.","kind":"release","outcome":"finished","recordRef":"tasks/rex-whichkey-19-01-inprocess-menu.md","runID":"rex-whichkey-p19-t01-builder-a01","schemaVersion":1,"specHash":"sha256:c35544123e5829bfc193948c34e6069452501da2dd122754ccf0610de058c5b1"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA0OjU2OjE4LjY1MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVsZWFzZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDEvcmVsZWFzZS8xIiwiaW50ZW50IjoiYmE5NjhiN2NkYTg4MTJmMTljY2FiYTZmZWIwMDc0NDJmYjlkNjI5MzNkMjRkOWFhMTM5OTcyMGZkMWExZDc4ZiJ9 -->
### 2026-11-07T04:56:18.651Z · release · unspecified

{"attempt":1,"dispatchEventID":"rex-whichkey-p19-t01/dispatch/1","eventID":"rex-whichkey-p19-t01/release/1","evidence":"Builder returned blocked after static inspection/build/vet on unchanged source; trusted exit0, no source writes. Required material decisions remain unresolved.","kind":"release","outcome":"stopped","recordRef":"tasks/rex-whichkey-19-01-inprocess-menu.md","schemaVersion":1,"specHash":"sha256:c35544123e5829bfc193948c34e6069452501da2dd122754ccf0610de058c5b1"}
<!-- mtn:end-entry:v1 -->
