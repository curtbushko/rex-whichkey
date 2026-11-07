---
identity:
  kind: task
  taskID: rex-whichkey-19-02-direct-lua-lifetime
  uuid: 0179e681-6c80-4986-afed-82327b72a636
structure:
  order: 2
scope:
  requirements:
    - REQ-001
    - REQ-002
    - REQ-003
    - REQ-004
    - REQ-009
    - REQ-010
    - REQ-011
    - REQ-013
    - REQ-016
    - REQ-017
    - REQ-025
    - REQ-028
    - REQ-029
  documents:
    - ref: designs/rex-whichkey-lua-float-launch.md
      hash: 'sha256:f7e38d49287de319753732b09fe472a40dfadcef81e5245e2deac208e968ab70'
implementation:
  language: go
  ownership:
    paths:
      - cmd/controller_legacy_support_test.go
      - cmd/main.go
      - controller_legacy_support_test.go
      - direct_menu.go
      - examples/rex-whichkey-keybinding.lua
      - launch.go
      - launch_legacy_support_test.go
      - lua/rex_whichkey.lua
      - menu_updates.go
      - native.go
      - native_app_owner.go
      - native_resources.go
      - native_sessions.go
      - native_targets.go
      - native_text_entry.go
      - popup.go
      - popup_initial.go
      - popup_legacy_support_test.go
      - popup_updates.go
      - providers.go
      - rex/block.go
      - rex/controller_legacy_support_test.go
      - rex/direct_popup.go
      - rex/layer.go
      - rex/session.go
      - runtime_bootstrap.go
      - runtime_legacy_support_test.go
      - terminal.go
    sharedPaths: []
    resources: []
    interfaces:
      - P19-INPROCESS-MENU
      - P19-FILTERED-ORIGIN-SNAPSHOT
      - P19-LUA-ORIGIN-ARGV
      - P19-DIRECT-LIFETIME
      - P19-SHARED-LUA-ENTRY
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
  state: reviewing
  sourceEventID: rex-whichkey-p19-t02-reviewer-a01/runtime/launched
  changedAt: '2026-11-07T08:40:01.844Z'
title: 'Phase19.2: Lua float and controller-free direct Go lifecycle'
mtnCreateRetry:
  key: rex-whichkey-p19/task-create/direct-lua-lifetime/1
  intent: 8de180373eb195a56756314a9bca5e7f6d6c65f00a2718af593dc568256b0248
status: in-progress
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T04:48:23.279Z'
dateModified: '2026-11-07T08:40:02.449Z'
tags:
  - task
type: task
id: 1b2c1e8e-409b-4e80-8f7a-132bb7326bd0
projects:
  - '[[phases/rex-whichkey-19-lua-float-launch.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-19-01-inprocess-menu.md]]'
    reltype: FINISHTOSTART
---
After delivered task1 direct Go/Lua vertical slice, retire the OLD ordinary Go launch controller/renderer/image-copy/lease/result/update production plane. Wire the public direct entry intentionally; no no-argument/caller ambiguity may treat ordinary original pane as owned popup. Preserve task1 direct lifecycle/current behavior. Lua owns native float creation/sizing/default-chrome policy, Go menu/operations/palette only. No launching Go parent, hidden helper, renamed supervisor, registry/relay, private IPC/copy path still used by normal shared Lua flow. No Go origin-geometry startup sizing.

Trace real callers before removal. Keep live strict JSON/Unicode, concrete Rex resources/current app targeting/data validation/provider/rename/terminal code. Preserve ALL historical assertions/test files; where old controller symbols are necessary to compile unchanged tests, isolate genuine original support into approved *_test.go files, not production stubs. Do not remove unrelated legacy assets/dependencies or claim tests ran. Owned files only; add a precise literal support-file ownership request to coordinator if genuinely needed.

All original target/generation/freshness/confirm/no-retry and accepted lifecycle semantics retained: visible float/tty/workers/focus before selected effects, detached host retires after Go exit under Rex. No all-block-before-effects claim. Current callbacks/executable Lua/hooks remain deferred. Go normal build/vet/format/whitespace and fresh independent source/dead-production-path review required; optional inherited lint diagnostic. No tests/TDD/race/Lua/parser/FIFO harnesses/benchmark/environment/user config/profile/reload/remote/deploy/push.

Coordinator actual shared-Lua public run on combined source: direct program/no Go parent/copy/private transport, root/dynamic/inline exact owned rename, cancel/dismiss, terminal/focus/visible and detached cleanup. Actual final make run is task3 integration criterion, not a private ACK or compile substitute. Preserve separate dirty flake.nix, quality gates/history.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA0OjQ4OjQwLjYxNloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYmxvY2tlciIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDIvbGlmZXRpbWUtYmxvY2tlci8xIiwiaW50ZW50IjoiMmY0ZTkzNmE4MWVhMGU3YWE1ZDA4ZDVkODRiMzM5YTkwM2M2NmQ4OTY3NDM3MjZmZWE0Y2YyZjI0YjBiNTQxNSJ9 -->
### 2026-11-07T04:48:40.616Z · blocker · unspecified

{"affectedRefs":["tasks/rex-whichkey-19-02-direct-lua-lifetime.md","tasks/rex-whichkey-19-03-source-run-delivery.md"],"attempt":1,"decisionID":"REXWHICHKEY-20261106-234820","eventID":"rex-whichkey-p19-t02/lifetime-blocker/1","evidence":".mtn-tmp/p19/lifetime-assessment.json; installed close/detach/exit schemas; current docs/launch-contract.md cleanup inventory.","kind":"blocker","reason":"Documented detach removes visible float but retains executing block through dispatch. Existing accepted all-owned-block cleanup-before-effect and suppress-on-cleanup-failure cannot be preserved by the same process; close kills Go. No permitted same-process host retirement API found.","recordRef":"tasks/rex-whichkey-19-02-direct-lua-lifetime.md","requiredDecision":"Explicitly accept visible-layer/TTY/worker/focus cleanup before effect with detached execution host retired after Go exits (post-exit cleanup cannot suppress prior effect), OR retain strict all-resource pre-effect rule and supply a supported same-process capability/approve a different architecture.","schemaVersion":1,"specHash":"sha256:94903530c10ca9ae9af67231b30b1927c1c2363eea6b3d2e483e811883c2415d"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA3OjIxOjQwLjY3NFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVzdW1lIiwicmV0cnlLZXkiOiJyZXgtd2hpY2hrZXktcDE5LXQwMi9saWZldGltZS1ibG9ja2VyLzEvcmVzdW1lLzIiLCJpbnRlbnQiOiJkYTk1MWIyMWJkNTgzZTcxNGE4NjE3NmYxZDgwMjA4NDE2YzY4MmViY2E5NTdkODUwMjA5NGE4ZDY4NTZiMGZkIn0 -->
### 2026-11-07T07:21:40.674Z · resume · unspecified

{"attempt":2,"blockerEventID":"rex-whichkey-p19-t02/lifetime-blocker/1","decisionEvidence":"1a - yes, accept it for now; 1b - yes","eventID":"rex-whichkey-p19-t02/lifetime-blocker/1/resume/2","kind":"resume","recordRef":"tasks/rex-whichkey-19-02-direct-lua-lifetime.md","resolution":"User explicitly approves1a default chrome deferral and1b host-through-dispatch exception. Replanned current-used Go/Lua vertical slice provides real consumer with widened literal ownership; no dormant API or unapproved ordering change. Prior source attempt stopped/no edits; original blocker history preserved.","schemaVersion":1,"specHash":"sha256:31a2ab209e370f04195026c90db164569768106bcb65d0dc92d07b18f4661ab5"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA4OjMwOjI4LjM5NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiZGlzcGF0Y2giLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMTktdDAyL2Rpc3BhdGNoLzEiLCJpbnRlbnQiOiJhMmZkNzY5ODA2Yzc3YzU3OGEyNGYyOGJkMjAwMDM5YjUwYjQ4ZTBlODJiY2ViZTBhYmM2ZjM2ODFiYzNhNmQ1In0 -->
### 2026-11-07T08:30:28.395Z · dispatch · unspecified

{"allowedCommands":["owned source/format/whitespace","task build","task lint:vet","optional full unchanged lint diagnostic","hooks-disabled owned commit"],"allowedScope":{"interfaces":["P19-INPROCESS-MENU","P19-FILTERED-ORIGIN-SNAPSHOT","P19-LUA-ORIGIN-ARGV","P19-DIRECT-LIFETIME","P19-SHARED-LUA-ENTRY"],"paths":["cmd/controller_legacy_support_test.go","cmd/main.go","controller_legacy_support_test.go","direct_menu.go","examples/rex-whichkey-keybinding.lua","launch.go","launch_legacy_support_test.go","lua/rex_whichkey.lua","menu_updates.go","native.go","native_app_owner.go","native_resources.go","native_sessions.go","native_targets.go","native_text_entry.go","popup.go","popup_initial.go","popup_legacy_support_test.go","popup_updates.go","providers.go","rex/block.go","rex/controller_legacy_support_test.go","rex/direct_popup.go","rex/layer.go","rex/session.go","runtime_bootstrap.go","runtime_legacy_support_test.go","terminal.go"],"resources":[],"sharedPaths":[]},"attempt":1,"baseCommit":"2522b39b5b109a5fccc8fc42ef51886540aba5b4","branch":"work/p19-t02-builder-a01","constraints":["Task01 delivered/current summary/ancestry verified; accepted host-through-dispatch and default native chrome deferral unchanged","Remove production controller/copy/IPC; genuine old support only test-local, assertions unchanged","No tests/harnesses/benchmark/live leaf actions/environment hooks/config/deploy/push; coordinator alone mutates/integrates/runs UI","Separate flake.nix preserved; Task03 normal make/package pending"],"eventID":"rex-whichkey-p19-t02/dispatch/1","kind":"dispatch","limits":{"builders":1,"cpu":14,"goMemoryLimit":"5GiB","heavyGates":1,"integrations":1,"reviewers":1,"totalMemoryGiB":6},"recordRef":"tasks/rex-whichkey-19-02-direct-lua-lifetime.md","runID":"rex-whichkey-p19-t02-builder-a01","schemaVersion":1,"specHash":"sha256:be14766bf6966140bfda84c6daf31f7fc09bfa6b4eb11025b3ac9b4bd92b4cc2","targetBranch":"main","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t02-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA4OjMwOjM3LjY3MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDItYnVpbGRlci1hMDEvcnVudGltZS9yZXNlcnZlZCIsImludGVudCI6Ijc5ZjRhMWEwZThmMjNiYWRkZDZhZDNiNjBhYTU2ZjY0ZDZjMjhhZDY2YTM4OTNkYWYzNTc1OTJjOTA2MzJhMzgifQ -->
### 2026-11-07T08:30:37.671Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-15-code-owned-runtime.md","runID":"rex-whichkey-p15-coordinator-a01"},"dispatchEventID":"rex-whichkey-p19-t02/dispatch/1","eventID":"rex-whichkey-p19-t02-builder-a01/runtime/reserved","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"75ff9ad4-a558-4295-8b20-855dafe77a8f"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-19-lua-float-launch.md","piSessionID":"rex-whichkey-p19-t02-builder-a01","piSessionName":"p19-t02-builder-a01","recordRef":"tasks/rex-whichkey-19-02-direct-lua-lifetime.md","rexSessionName":"rex-whichkey-main","rexWindowName":"p19-t02-builder-a01","role":"builder","runID":"rex-whichkey-p19-t02-builder-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-builder","specHash":"sha256:be14766bf6966140bfda84c6daf31f7fc09bfa6b4eb11025b3ac9b4bd92b4cc2","stage":"reserved","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p19-t02-builder-a01","windowName":"p19-t02-builder-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t02-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA4OjMwOjQxLjM2N1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDItYnVpbGRlci1hMDEvcnVudGltZS9sYXVuY2hlZCIsImludGVudCI6ImVkOGYzZWUzZjBiNGE0MTc3ZjNlMjUzMjA4NmQwNjkwZDllOGM1OTZjYjk3YTVlZjNhZGM2Y2RlZjc5NjRiMDUifQ -->
### 2026-11-07T08:30:41.367Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-15-code-owned-runtime.md","runID":"rex-whichkey-p15-coordinator-a01"},"dispatchEventID":"rex-whichkey-p19-t02/dispatch/1","eventID":"rex-whichkey-p19-t02-builder-a01/runtime/launched","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"75ff9ad4-a558-4295-8b20-855dafe77a8f"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-19-lua-float-launch.md","piSessionID":"rex-whichkey-p19-t02-builder-a01","piSessionName":"p19-t02-builder-a01","provider":{"displayName":"p19-t02-builder-a01","name":"rex","resources":{"block":"block:01a1b521-9491-7572-a41a-0ef76405fcd9","endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","session":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","window":"window:01a1b521-9492-79e7-ac3b-1c77b9b2268e"}},"recordRef":"tasks/rex-whichkey-19-02-direct-lua-lifetime.md","rexBlockID":"block:01a1b521-9491-7572-a41a-0ef76405fcd9","rexSessionID":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","rexSessionName":"rex-whichkey-main","rexWindowID":"window:01a1b521-9492-79e7-ac3b-1c77b9b2268e","rexWindowName":"p19-t02-builder-a01","role":"builder","runID":"rex-whichkey-p19-t02-builder-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-builder","specHash":"sha256:be14766bf6966140bfda84c6daf31f7fc09bfa6b4eb11025b3ac9b4bd92b4cc2","stage":"launched","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p19-t02-builder-a01","windowName":"p19-t02-builder-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t02-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA4OjM5OjQ4LjYxNFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYnVpbGQiLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMTktdDAyL2J1aWxkLzEiLCJpbnRlbnQiOiI3NjE2NTEyNjExM2Q3M2JlNjliNTU2MjQyMDA4MDZlOTNmNThlMTJkNTBkNzZkNjQ0YjVhN2YwNTVkMmM1ZTgyIn0 -->
### 2026-11-07T08:39:48.614Z · build · unspecified

{"artifactRoot":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p19-t02-builder-a01","attempt":1,"baseCommit":"2522b39b5b109a5fccc8fc42ef51886540aba5b4","branch":"work/p19-t02-builder-a01","commitPolicy":"Owned source committed with core.hooksPath=/dev/null; no integration or mtn mutation. Clean worktree frozen for review.","commits":["578c40bb6712a46d4f72bf611dc732d401a44ecb"],"coordinator":{"phaseRef":"phases/rex-whichkey-15-code-owned-runtime.md","runID":"rex-whichkey-p15-coordinator-a01"},"coverage":{"requirements":["REQ-001","REQ-002","REQ-003","REQ-004","REQ-009","REQ-010","REQ-011","REQ-013","REQ-016","REQ-017","REQ-025","REQ-028","REQ-029"],"status":"Source/static/build evidence only; preserved delivered Task01 behavior is not a new runtime acceptance result.","tests":"All pre-existing test/assertion file blobs unchanged; only genuine original helpers moved into approved new test-support files. No tests/assertions added or executed."},"diagnosticValidation":[{"command":"task lint","detail":"Optional unchanged full-lint diagnostic omitted; inherited full-lint failures are not claimed resolved.","required":false,"result":"not-run"}],"dispatchEventID":"rex-whichkey-p19-t02/dispatch/1","eventID":"rex-whichkey-p19-t02/build/1","files":{"created":["launch_legacy_support_test.go","popup_legacy_support_test.go"],"deleted":["native_text_entry.go","popup_initial.go","runtime_bootstrap.go"],"modified":["cmd/main.go","direct_menu.go","launch.go","menu_updates.go","native.go","popup.go","popup_updates.go","providers.go","rex/layer.go","terminal.go"]},"filesBasis":"All paths versus dispatch base with rename detection disabled. Git alternatively reports native_text_entry.go -> popup_legacy_support_test.go as R050.","head":"578c40bb6712a46d4f72bf611dc732d401a44ecb","historicalSupport":{"deletedRatherThanPreserved":"Controller/lifetime/image/update-slot code, private snapshot decoders and readNativeTextSpec have no remaining live or historical-test callers; deleted rather than copied wholesale into tests. Deleted files were moved to ignored .trash/p19-controller locally.","launch_legacy_support_test.go":"Original launchOutput type and Write implementation moved unchanged for launch_output_test.go.","liveHelpers":"Public strict JSON/Unicode/Completion validation in data_config.go and mapping.go remains unchanged. Real retained-selection, metadata, provider, rename, native resource/action, terminal palette/input and direct-lifetime helpers remain production code.","popup_legacy_support_test.go":"Original openPopupDirectory, checkPopupSlot, readPopupResult, writePopupMenuResult and nativeTextSelection implementations, with their original required bound/error constants, retained only for unchanged configured rename assertions. No duplicate production implementations or invented stubs.","scope":"Only already-owned support paths used. No additional helper-path ownership request was needed. Unrelated generic Rex geometry/terminal/wait adapters and legacy Lua fixtures were not blanket-removed; no menu startup sizing caller invokes origin Geometry."},"integrationChecks":[{"check":"Fresh independent exact-head/tree review, especially removed production reachability and original test-helper disposition; combined-source normal build and vet before main advance.","status":"pending"},{"check":"Coordinator actually invokes this shared Lua script/action on reviewed combined source, observes one direct current-build Go menu, dynamic readiness, exact owned inline rename, edit cancel and dismissal.","status":"pending"},{"check":"Observe visible retirement/TTY/workers/exact original focus before effects and session.list_blocks after Go exit for detached-host cleanup; no compilation or creation-ACK substitution.","status":"pending"},{"check":"Task03 delivers and actually verifies make/source-run/package route; native chrome customization remains explicitly user-deferred.","status":"pending"}],"kind":"build","limitations":["No new runtime success, independent-review approval, integration or phase completion claimed.","Accepted user1a retains Rex default native chrome; REQ-026 customization is deferred. Accepted user1b permits the detached Go host during selected effects until Go exits; final cleanup cannot suppress an earlier effect.","Go cannot verify its own post-exit deletion; coordinator must inspect actual block inventory. Uncertain pre-effect ownership/detach/TTY/focus cleanup suppresses actions without retry or fallback.","No mtn writes, integration, agent/window/session creation or resume, config/profile/install/activation/reload/restart, Nix discovery/hooks, remote/deploy/push, dependency/pin changes or global cleanup performed."],"meaningfulCurrentCallers":["rex do ABSOLUTE/lua/rex_whichkey.lua --action rex_whichkey_launch invokes the shared source-script action. Canonical require('rex_whichkey').launch uses the same M.launch; only canonical require receives the function table.","M.launch captures original context IDs, resolves exact origin placement and creates one centered float with fixed rex-whichkey direct argv, shell=none, optional terminal-scoped development_directory PATH prepend, and Rex exit cleanup policy.","cmd.run -> runDirect -> RunDirectMenu -> SessionSnapshot.DirectPopup validates distinct original/own IDs and sole floating ownership, filtering the exact own popup from provider snapshots/counts/targets/current markers.","RunTerminal -> directMenuState.poll -> bounded metadataJobs retains provider/rename readiness and exact targets in-process before UI publication. Existing own-PTY sizing, SIGWINCH, palette, Unicode/editing and confirmation paths remain.","RunTerminal joins workers and restores/closes TTY before directSelection's deferred retireDirectPopup -> Layer.Detach verifies visible retirement/detached placement and restores/revalidates exact original focus/placement/app generation.","RunDirectMenu validates retained selection and dispatchNativeSelection invokes existing exact resource/app CLI effects once after successful pre-effect cleanup. Uncertain cleanup suppresses effects without retry; Rex owns detached-host deletion after Go exit.","cmd.run -> runMenuUI -> LoadMenuConfig -> RunTerminal -> WriteMenuResult remains explicitly in-terminal/result-only. Bare invocation and unknown/legacy popup commands return errors before config, CLI lookup, terminal or Rex effects."],"notRun":["Prohibited tests/unit/race/TDD/Lua/parser/FIFO harnesses, benchmarks, fake servers/delays and broad experiments.","Prohibited leaf live Rex actions, public UI execution, actual make run and physical keybinding checks."],"outcome":"ready-for-review","phaseRef":"phases/rex-whichkey-19-lua-float-launch.md","preservation":"Original assertion files, .tasks, hooks, Taskfile build/test/lint definitions, dependencies/pins, Lua source/example, direct popup ownership/filtering helper, shared live validation/UI helpers and package paths unchanged. Parent's separate dirty flake.nix remains untouched/outside commit. See preservation.log, final-preservation.log, commit.log and source-identity.log.","publicEntry":{"action":"rex_whichkey_launch","builtSourceCommand":"rex --server \"$ENDPOINT\" --client \"$CLIENT\" do --session \"$SESSION\" --block \"$BLOCK\" /Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t02-builder-a01/lua/rex_whichkey.lua --action rex_whichkey_launch development_directory=/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t02-builder-a01/bin/darwin-arm64 --json","directArgv":"rex-whichkey direct --origin-session SESSION --origin-block BLOCK --origin-window WINDOW --origin-layer LAYER [--client CLIENT] [--config PATH] [--ascii] [--no-color]","directOwnershipBlob":"545a88eb79888f00822f2f9833183013705d9f47","installedBinding":"local whichkey = require('rex_whichkey'); rex.bind('ctrl+space', function(ctx) return whichkey.launch(ctx, {width=.8, height=.3}) end)","preparation":"Coordinator-only invocation supplied, not executed by leaf. Select exact owned ENDPOINT/SESSION/BLOCK and optional original CLIENT; omit --client entirely when absent. For combined-source verification, rebuild there and substitute that same checkout root in both script and bin-directory arguments.","script":"lua/rex_whichkey.lua","scriptBlob":"44fee5d965abb396f18e1da838d4f156a28ef986","sourceCallerExample":"rex do \"$PWD/lua/rex_whichkey.lua\" --action rex_whichkey_launch development_directory=\"$PWD/bin/$(go env GOOS)-$(go env GOARCH)\"","status":"Working Task01 Lua script/action and direct ownership implementation preserved unchanged. This revision requires fresh coordinator runtime verification. Current Taskfile run still invokes the bare binary; Task03 owns make/source-run/package routing. Native chrome customization is deferred, not delivered."},"recordRef":"tasks/rex-whichkey-19-02-direct-lua-lifetime.md","removedCallerChain":"cmd.run -> runLaunch -> LaunchPopup -> launchNativePopup -> nativePopupSelection -> runPopupLifecycle -> copied-image popup child/RunNativePopupTerminal/runPopupTerminal. Removed that chain, its watcher/controller/transport/image helpers, and unused Go Layer creation/resize/close/wait/bounds methods. No alternate launcher, supervisor or production compatibility wrapper remains.","role":"builder","runID":"rex-whichkey-p19-t02-builder-a01","schemaVersion":1,"skill":"go-builder","specHash":"sha256:be14766bf6966140bfda84c6daf31f7fc09bfa6b4eb11025b3ac9b4bd92b4cc2","summary":"Removed the ordinary production Go popup launcher/controller, renderer child entry, self-image copying, lifetime FIFO, result/update file IPC and Go startup sizing. Explicit direct and result-only menu modes remain; bare/ambiguous/legacy popup invocation fails without launch or detach. Required source/build/vet gates passed; no new runtime claim.","targetBranch":"main","tree":"98d7af67b07d738857a850ed11a717e01918e606","validation":[{"artifact":"source-format.log","command":"test -z \"$(gofmt -l ./*.go ./rex/*.go ./cmd/*.go)\"","exitCode":0,"result":"pass"},{"artifact":"whitespace.log","command":"git diff --check","exitCode":0,"result":"pass"},{"artifact":"build.log","command":"task build","exitCode":0,"result":"pass"},{"artifact":"vet.log","command":"task lint:vet","exitCode":0,"result":"pass"},{"artifact":"final-preservation.log","command":"git diff --check 2522b39b5b109a5fccc8fc42ef51886540aba5b4 HEAD; git diff --cached --check","exitCode":0,"result":"pass"}],"validationEnvironment":{"GOMAXPROCS":14,"GOMEMLIMIT":"5GiB","GOTOOLCHAIN":"local","Go":"go1.26.8 darwin/arm64","PATHPrefix":"/nix/store/dx7xl254pz6c99rmslmaz8hxzfich3m7-go-1.26.8/bin","execution":"Required build then vet sequentially; no environment activation."},"worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t02-builder-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA4OjM5OjU3LjI1M1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDItcmV2aWV3ZXItYTAxL3J1bnRpbWUvcmVzZXJ2ZWQiLCJpbnRlbnQiOiI0MDgyOGU4N2MzODk4N2RhZWZkNjNkYzI2NDUyNGQ4ZWZkMjY1NzgwMmQ5Yzg4MjA1YmE0NzAwNTBlY2Q0Y2MyIn0 -->
### 2026-11-07T08:39:57.253Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-15-code-owned-runtime.md","runID":"rex-whichkey-p15-coordinator-a01"},"dispatchEventID":"rex-whichkey-p19-t02/dispatch/1","eventID":"rex-whichkey-p19-t02-reviewer-a01/runtime/reserved","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"75ff9ad4-a558-4295-8b20-855dafe77a8f"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-19-lua-float-launch.md","piSessionID":"rex-whichkey-p19-t02-reviewer-a01","piSessionName":"p19-t02-reviewer-a01","recordRef":"tasks/rex-whichkey-19-02-direct-lua-lifetime.md","rexSessionName":"rex-whichkey-main","rexWindowName":"p19-t02-reviewer-a01","role":"reviewer","runID":"rex-whichkey-p19-t02-reviewer-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-reviewer","specHash":"sha256:be14766bf6966140bfda84c6daf31f7fc09bfa6b4eb11025b3ac9b4bd92b4cc2","stage":"reserved","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p19-t02-reviewer-a01","windowName":"p19-t02-reviewer-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t02-reviewer-a01"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA4OjQwOjAwLjk0MFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDItcmV2aWV3ZXItYTAxL3J1bnRpbWUvbGF1bmNoZWQiLCJpbnRlbnQiOiIwNjE3YjkxNGM3MTAyZjZiOWI1Y2FlMDlmOGEyNTVkZDA1ZGNiMWVlNzczMjUxYzY4MjUxNzZlZWY2NDUyMjMwIn0 -->
### 2026-11-07T08:40:00.940Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-15-code-owned-runtime.md","runID":"rex-whichkey-p15-coordinator-a01"},"dispatchEventID":"rex-whichkey-p19-t02/dispatch/1","eventID":"rex-whichkey-p19-t02-reviewer-a01/runtime/launched","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"75ff9ad4-a558-4295-8b20-855dafe77a8f"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-19-lua-float-launch.md","piSessionID":"rex-whichkey-p19-t02-reviewer-a01","piSessionName":"p19-t02-reviewer-a01","provider":{"displayName":"p19-t02-reviewer-a01","name":"rex","resources":{"block":"block:01a1b52a-1e5b-7b5e-a1b7-c2659b3fd0bb","endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","session":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","window":"window:01a1b52a-1e5d-70b3-b47e-3973b7715714"}},"recordRef":"tasks/rex-whichkey-19-02-direct-lua-lifetime.md","rexBlockID":"block:01a1b52a-1e5b-7b5e-a1b7-c2659b3fd0bb","rexSessionID":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","rexSessionName":"rex-whichkey-main","rexWindowID":"window:01a1b52a-1e5d-70b3-b47e-3973b7715714","rexWindowName":"p19-t02-reviewer-a01","role":"reviewer","runID":"rex-whichkey-p19-t02-reviewer-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-reviewer","specHash":"sha256:be14766bf6966140bfda84c6daf31f7fc09bfa6b4eb11025b3ac9b4bd92b4cc2","stage":"launched","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p19-t02-reviewer-a01","windowName":"p19-t02-reviewer-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p19-t02-reviewer-a01"}
<!-- mtn:end-entry:v1 -->
