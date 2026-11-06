---
identity:
  taskID: rex-whichkey-16-03-launch-package-docs
  kind: task
  uuid: 4357adda-da01-48e3-8e39-d764c7dff3a4
structure:
  order: 3
scope:
  requirements:
    - REQ-001
    - REQ-003
    - REQ-013
    - REQ-014
    - REQ-015
    - REQ-016
    - REQ-017
  documents:
    - ref: designs/rex-whichkey-cli-only-menu.md
      hash: 'sha256:23add062dac4326fe35e7e178a6a39f4c4d7631b4c8e164feb988b9aba07e21b'
  dependencies:
    - ref: tasks/rex-whichkey-16-02-go-cli-menu.md
      relation: finish-to-start
implementation:
  language: lua
  gates:
    - ID: whitespace
      command: git diff --check
      required: true
    - ID: nix-parse-only
      command: >-
        nix-instantiate --parse nix/package.nix >/dev/null && nix-instantiate
        --parse nix/home-manager.nix >/dev/null && nix-instantiate --parse
        flake.nix >/dev/null
      required: true
    - ID: nix-format-only
      command: alejandra --check nix/package.nix nix/home-manager.nix flake.nix
      required: true
    - ID: combined-build
      command: task build
      required: true
    - ID: coordinator-observed-combined-public-menu-dismissal-cleanup
      command: task run
      required: true
  ownership:
    paths:
      - lua/rex_whichkey.lua
      - lua/rex_whichkey/config.lua
      - lua/rex_whichkey/evaluate.lua
      - lua/rex_whichkey/mapping.lua
      - lua/rex_whichkey/dispatch.lua
      - lua/rex_whichkey/defaults/init.lua
      - lua/rex_whichkey/defaults/sessions.lua
      - lua/rex_whichkey/defaults/tabs.lua
      - lua/rex_whichkey/defaults/panes.lua
      - lua/rex_whichkey/defaults/app.lua
      - lua/rex_whichkey/defaults/hosts.lua
      - lua/rex_whichkey/runtime/anchor.lua
      - tests/fixtures/phase15-lua
      - tests/lua/action_context.lua
      - tests/lua/cli_delegation.lua
      - tests/lua/configured_rename.lua
      - examples/rex-whichkey-keybinding.lua
      - examples/my_whichkey.lua
      - examples/README.md
      - nix/package.nix
      - nix/home-manager.nix
      - flake.nix
      - README.md
      - docs/actions.md
      - docs/launch-contract.md
      - docs/runtime-bootstrap.md
      - docs/remote-compatibility.md
      - docs/cli-gui-delegation.md
      - docs/session-owned-ui.md
      - docs/tab-rename-input.md
      - docs/appearance.md
      - docs/acceptance.md
      - tests/packaging/home-manager/README.md
    sharedPaths:
      - docs/mapping-format.md
      - examples/my_whichkey.json
    resources:
      - 'build-output:bin'
      - coordinator-heavy-gate-slot
      - coordinator-owned-public-menu-observation-origin
    interfaces:
      - CLI-ONLY-LAUNCH
      - ORDINARY-PACKAGE-LAYOUT
lifecycle:
  state: complete
  sourceEventID: rex-whichkey-p16-t03/dispatch-release/1
  changedAt: '2026-11-06T22:56:59.620Z'
title: >-
  Ship minimal keybinding and ordinary package with migration and combined app
  validation
mtnCreateRetry:
  key: rex-whichkey-16-03-launch-package-docs-create-v1
  intent: c736d4036603549d61af15280c2b29b84bf95b85d6d4175d96de3db2345f0338
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-06T18:42:00.300Z'
dateModified: '2026-11-06T22:57:51.941Z'
tags:
  - task
type: task
id: 3fa36f79-07ef-4a73-a092-d108be5aa828
projects:
  - '[[phases/rex-whichkey-16-cli-only-menu.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-16-02-go-cli-menu.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-06'
---
Consume task2's running public CLI and task1's documented data schema. Reduce lua/rex_whichkey.lua to a minimal launch function used by a documented rex.bind function keybinding in examples/rex-whichkey-keybinding.lua. Use documented rex.exec.start argument-array execution of installed rex-whichkey, detach=true, no action-duration wait or command timeout; pass only exact ctx.session_id/ctx.block_id/ctx.client_id as origin/app intent using process environment. Go owns validation, endpoint/CLI discovery and menu lifecycle. No rex.action/register_launch/register_completion, metadata query, rex.call/session query/mutation, Lua JSON config compiler, registry, callback, completion, anchor/path search or helper process controller. Do not copy/stringify ctx.server or assume terminal-only environment is inherited correctly. Throw an honest immediate launch error; a returned detached process handle proves only start, not displayed menu. The function returns true for handled binding. Keep key spelling an example, never install/alter the user's binding. No physical key/reload result is claimed without observation and separate permission.
Retire the listed obsolete Lua config/evaluate/mapping/dispatch/default/anchor assets from the shipped production tree. Preserve old tests and assertions: three tests/lua scripts explicitly prepend lua paths and require legacy modules, so isolate the original legacy Lua tree under tests/fixtures/phase15-lua and change only those scripts' loader search-path plumbing to that historical fixture namespace. Preserve fixture behavior/assertions and all Go tests; no harness execution and no test deletion/skip/weakening. Fixtures must be excluded from shipped assets and cannot be loaded by the minimal production launcher. Retain any historical test limitations honestly instead of promising tests were run. Retire examples/my_whichkey.lua executable-options example in favor of task1's JSON data example, with migration instructions; do not silently discard a user's old custom data or attempt automatic Lua conversion. Explain unsupported callbacks/functions as deferred, not silently translated or accepted.
Update nix/package.nix to install an ordinary real bin/rex-whichkey and only minimal launcher/data/docs. If a legitimate optional licensed CLI dependency is present, install its symlink at libexec/rex-whichkey/rex for Go's code-owned discovery. Remove the old real-image-inside-Lua-tree, anchor, runtime/deps and generated-helper layout; no wrapper, runtime Nix resolver, policy flags or ldflags. Preserve upstream pins, Linux loader/PAM/terminfo/SSH assets, Darwin signature and truthful Intel-macOS artifact absence; retain doCheck=false/doInstallCheck=false. Update Home Manager to link only the needed launch module, no obsolete recursive module tree, init.lua, generated data config or whole Rex directory takeover. Preserve conflict safety and opt-in package/configDirectory semantics; remove obsolete namespace linkage from newly generated configuration without activating it. Update flake descriptions, README, examples and every listed active launch/config/action/bootstrap/remote/rename/appearance/session-owned/acceptance and packaging guide. Mark legacy sections historical rather than presenting old APIs/registrations as prerequisites, and preserve historical evidence/commits. Document migration, current defaults, data path/precedence, provider/action semantics, simple launch lifetime, endpoint/GUI-vs-caller safety, discovery and exact owned cleanup. Do not rewrite historical docs/rex-whichkey.prd or .tasks history.
Own source review of minimal Lua/Nix/docs and integration consistency; Nix gates are parse/format/whitespace only, no flake check/evaluation/build/activation or indirect tests. Final combined source requires normal task build and coordinator-observed actual public app/menu, ordinary dismissal and exact popup/terminal/private-artifact cleanup after review. Include one benign Go-owned dynamic selection or inline rename only on an explicitly coordinator-owned disposable resource under actual-run authority; never mutate a real user tab/session merely for evidence. Record precisely which actions/entrypoints were exercised and which remain source-reviewed only. The keymapping may be source-reviewed/documented without config deployment; do not block public app validation on an unapproved user reload. Task completion includes accurate docs and combined observations, not a disconnected verification task or an assertion that Phase15/unrelated phases are fixed.
<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA2VDIyOjU3OjUxLjkzM1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxNi10MDMvc3VtbWFyeS8xIiwiaW50ZW50IjoiZDRmZjFiNGU0ZjRiMjQxNjZkMDYyOGJhYjI1YzUxNGU2ZTc3NmE3Mzc3ZDg0MjFmNjc1NTM1OWJiNGI3NjgwNSJ9 -->
### 2026-11-06T22:57:51.933Z · summary · unspecified

{"acceptedScope":"Minimal launch-only Lua/function-binding example, ordinary real public binary/optional CLI dependency package layout, safe single-file Home Manager integration, version1 JSON migration and accurate active guides. Old Lua subsystem is byte-identical historical test fixtures with loader-only script changes. Combined reviewed source actually ran normally with owned inline tab rename, dynamic menu/navigation, Escape dismissal, exit0 and exact cleanup.","actualRuntime":{"artifacts":[".mtn-tmp/p16/task3-rename-menu-root.txt",".mtn-tmp/p16/task3-rename-entry.txt",".mtn-tmp/p16/task3-rename-typed-entry.txt",".mtn-tmp/p16/task3-rename-after-session.json",".mtn-tmp/p16/task3-rename-exit.json",".mtn-tmp/p16/task3-rename-parent-process.json",".mtn-tmp/p16/task3-rename-parent-modes.json",".mtn-tmp/p16/task3-rename-after-owned-window-cleanup.json",".mtn-tmp/p16/task3-dismiss-root.txt",".mtn-tmp/p16/task3-dismiss-dynamic-sessions.txt",".mtn-tmp/p16/task3-dismiss-after-session.json",".mtn-tmp/p16/task3-dismiss-exit.json",".mtn-tmp/p16/task3-dismiss-parent-process.json",".mtn-tmp/p16/task3-dismiss-parent-modes.json",".mtn-tmp/p16/task3-dismiss-after-owned-window-cleanup.json"],"binarySHA256":"9a6b24747382db84d759da323e0cb281672c65d5567efa961aca1d0f732fe979","candidateCommit":"c46705bfe0b0664ccb27d9cbc41fcf0443e0b816","command":"Ordinary task run on combined reviewed candidate in two explicitly owned Rex windows, same existing session; no direct private entrypoint","completedAtUTC":"2026-11-06T22:55:58Z","dismissal":{"created":{"block_ids":["block:01a1b312-77b1-776a-b622-59ffda64f7a6"],"revision":224,"window_id":"window:01a1b312-77b2-78dd-a0b6-9da0d474e4bc"},"exactPopupBlockLayerPrivateDirCleanup":true,"exitCode":0,"observed":"Root, Sessions/Switch actual Go dynamic inventory; Backspace twice; Escape root dismissal","originFocusAndPrimaryTerminalRestored":true,"ownedValidationWindowClosed":true,"resources":{"label":"Window d474e4bc","layers":[{"blocks":["block:01a1b312-77b1-776a-b622-59ffda64f7a6"],"kind":"tiled","layer_id":"layer:01a1b312-77b2-78e1-ba62-8a12ceb97c77"},{"blocks":["block:01a1b312-807b-71b5-9fa3-b0b0e5ceb5e2"],"kind":"floating","layer_id":"layer:01a1b312-807c-711d-a4aa-2f988362f574"}],"window_id":"window:01a1b312-77b2-78dd-a0b6-9da0d474e4bc"}},"kind":"actual-public-runtime-evidence","limitations":["No physical keybinding/user config/reload or installed package activation. Lua/Nix source review and Nix parse/format only.","Actual Darwin arm64 default public menu and one owned inline tab-rename only; no other action/platform/remote trial or custom config probe.","No prohibited tests/harnesses/benchmarks, no unrelated resources mutated or closed."],"normalPublicArgs":[],"preservedBinary":".mtn-tmp/p16/task3-observed-rex-whichkey","rename":{"created":{"block_ids":["block:01a1b311-43e4-7900-9de5-651e2945d2d9"],"revision":217,"window_id":"window:01a1b311-43e5-7b85-bdca-c3015f9bbc6d"},"exactPopupBlockLayerPrivateDirCleanup":true,"exitCode":0,"observed":"WhichKey Tabs/Rename inline form showed exact new text; Enter returned through Go retained resource dispatch; exact owned window label changed to p16-owned-cli-rename before owned-window closure","originFocusAndPrimaryTerminalRestored":true,"ownedValidationWindowClosed":true,"resources":{"label":"Window 5f9bbc6d","layers":[{"blocks":["block:01a1b311-43e4-7900-9de5-651e2945d2d9"],"kind":"tiled","layer_id":"layer:01a1b311-43e5-7b89-965f-0fca95f4ee5c"},{"blocks":["block:01a1b311-4e36-7441-a240-e2af5546f235"],"kind":"floating","layer_id":"layer:01a1b311-4e37-74ed-99ea-7e4532b11bee"}],"window_id":"window:01a1b311-43e5-7b85-bdca-c3015f9bbc6d"},"typedText":"p16-owned-cli-rename"},"schemaVersion":1},"approvalSummary":["User accepted CLI-only REXWHICHKEY-20261106-114935 and start-it execution; no hooks or live config/deployment authority.","Bounded current Phase16 task plan approved/1; delivered Go contracts/native task prerequisites verified.","Fresh independent same-family combined review rex-whichkey-p16-t03/review/1 approved exact build1head5002fe7/tree62db248; coordinator actual combined runtime passed before target advance."],"attempt":1,"cleanup":{"allRuntimeReservationsReleased":true,"facade":"complete invoked/retried; native/protocol complete. Immutable cleanup jq syntax defect persists, so exact clean delivered no-consumer worktrees were removed manually; no tool/schema changes.","removedCleanDeliveredWorktrees":["p16-t03-builder-a01","p16-t03-tasks-a01","p16-t03-integration-a01"],"retainedTask3Resources":[],"successfulWorkerWindowsConfirmedAbsent":true},"completionEventID":"rex-whichkey-p16-t03/completion/1","eventID":"rex-whichkey-p16-t03/summary/1","files":["lua/rex_whichkey.lua","examples/rex-whichkey-keybinding.lua","examples/my_whichkey.json","nix/package.nix","nix/home-manager.nix","flake.nix","README.md","docs/launch-contract.md","docs/mapping-format.md","tests/fixtures/phase15-lua",".mtn-tmp/p16/task3-result.json",".mtn-tmp/p16/task3-review-result.json",".mtn-tmp/p16/task3-integration.json",".mtn-tmp/p16/task3-public-runtime-evidence.json",".mtn-tmp/p16/task3-observed-rex-whichkey",".mtn-tmp/p16/task3-post-cleanup-worktrees.txt"],"followUp":"Phase16 final gates must rebuild/run main ordinary public binary and validate native nonempty graph/coverage before phase completion. Nonblocking installed Markdown link portability REX-P16-T03-001 is recorded, not a runtime blocker.","integrationEventID":"rex-whichkey-p16-t03/integration/1","kind":"summary","limitations":["No user config/profile/package activation/reload or physical keybinding test. Minimal Lua validated against published API/source, not executed in user VM.","Nix parse/format only, no Nix evaluation/build/activation or Linux/Intel/remote runtime claim.","No tests/TDD/race/Lua/parser/FIFO harnesses, benchmarks or custom-config probes; all existing assertions/standards preserved.","Fresh independent same-family review, not cross-family evidence. Global lint remains inherited failing baseline from Go task; not rerun or claimed green here.","Nonblocking installed offline Markdown-link portability follow-up REX-P16-T03-001 remains.","Prior Phase15 failed original bridge history and unrelated native incomplete scope remain unchanged."],"nativeEvidenceCommit":"467b5f2bc1a9cf5c946b7db47a72328861ad1124","recordRef":"tasks/rex-whichkey-16-03-launch-package-docs.md","requirements":["REQ-001","REQ-003","REQ-013","REQ-014","REQ-015","REQ-016","REQ-017"],"schemaVersion":1,"specHash":"sha256:cf1e98a592fa7bdc3e0bc6f84097fce34fb222395a2957dc43955c57bd976f60","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"c46705bfe0b0664ccb27d9cbc41fcf0443e0b816","targetTree":"344076aeef07eccef439115af4684e400c23c116","verification":[{"artifact":".mtn-tmp/p16/task3-integration-build.log","command":"Exact combined candidate task build","result":"pass"},{"artifact":".mtn-tmp/p16/task3-integration-nix-parse.log","command":"Combined candidate Nix parse/alejandra format/whitespace only","result":"pass"},{"artifact":".mtn-tmp/p16/task3-integration-source-identity.txt","command":"Reviewed46-path exact source identity, unchanged standards, clean candidate and target-base recheck before fast-forward; reviewed head ancestor of main","result":"pass"},{"artifact":".mtn-tmp/p16/task3-public-runtime-evidence.json","command":"Ordinary combined public task run: actual menus/dynamic navigation, inline rename of exact disposable owned tab with observed label, ordinary Escape dismissal, two exit0 outcomes and exact popup/block/private-dir/terminal/owned-window cleanup","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
