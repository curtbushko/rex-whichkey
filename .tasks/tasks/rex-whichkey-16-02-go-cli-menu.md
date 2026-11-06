---
identity:
  taskID: rex-whichkey-16-02-go-cli-menu
  kind: task
  uuid: 5809a0d4-ace5-4402-9200-04e935eff696
structure:
  order: 2
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
    - REQ-016
    - REQ-017
  documents:
    - ref: designs/rex-whichkey-cli-only-menu.md
      hash: 'sha256:23add062dac4326fe35e7e178a6a39f4c4d7631b4c8e164feb988b9aba07e21b'
  dependencies:
    - ref: tasks/rex-whichkey-16-01-data-config.md
      relation: finish-to-start
implementation:
  language: go
  gates:
    - ID: owned-format
      command: >-
        test -z "$(git ls-files -z -- '*.go' ':!:*_test.go' | xargs -0 gofmt
        -l)"
      required: true
    - ID: whitespace
      command: git diff --check
      required: true
    - ID: build
      command: task build
      required: true
    - ID: static-type-review
      command: 'task lint:vet'
      required: true
    - ID: lint-baseline-diagnostics
      command: task lint
      required: false
    - ID: coordinator-observed-public-menu-dismissal-cleanup
      command: task run
      required: true
  ownership:
    paths:
      - action.go
      - menu.go
      - launch.go
      - native.go
      - native_app_owner.go
      - native_resources.go
      - native_sessions.go
      - native_targets.go
      - native_text_entry.go
      - popup.go
      - popup_transport.go
      - providers.go
      - runtime_bootstrap.go
      - terminal.go
      - text_entry.go
      - hosts.go
      - cmd/main.go
      - rex/rex.go
      - rex/exec.go
      - rex/runtime.go
      - rex/endpoint.go
      - rex/block.go
      - rex/session.go
      - rex/window.go
      - rex/layer.go
      - rex/geometry.go
      - rex/terminal.go
      - rex/hosts.go
      - rex/actions.go
      - rex/wait.go
    sharedPaths:
      - config.go
      - mapping.go
      - defaults.go
      - docs/mapping-format.md
      - data_config.go
    resources:
      - 'build-output:bin'
      - coordinator-heavy-gate-slot
      - coordinator-owned-public-menu-observation-origin
    interfaces:
      - CLI-ONLY-LAUNCH
      - REX-CLI-EXECUTION
      - GO-MENU-LIFECYCLE
      - GO-ACTION-PROVIDER-CATALOGUE
lifecycle:
  state: complete
  sourceEventID: rex-whichkey-p16-t02/dispatch-release/2
  changedAt: '2026-11-06T22:27:32.710Z'
title: Run the ordinary Go CLI-only menu with dynamic targets and real actions
mtnCreateRetry:
  key: rex-whichkey-16-02-go-cli-menu-create-v1
  intent: 1aa07d9252a08eb35486f217eba2a5a69c8c6e6aeb2f68d987d6dd936c9cec6c
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-06T18:41:58.393Z'
dateModified: '2026-11-06T22:28:58.546Z'
tags:
  - task
type: task
id: 148e98d8-5909-4ad3-967c-601aa579efa9
projects:
  - '[[phases/rex-whichkey-16-cli-only-menu.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-16-01-data-config.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-06'
---
Consume task1's data contract and deliver a REAL RUNNING public rex-whichkey binary, not a new private native entry or compile-only milestone. Wire cmd/main.go/runLaunch -> LaunchPopup -> Go-loaded configuration -> existing native popup/terminal/menu lifecycle -> retained selection -> real Go rex dispatch. Public no-argument and --config/--server/--client/--ascii/--no-color invocation use this one path; remove --action routing and ordinary prepare/controller/controller-child/runtime-describe/runtime-relay/runtime-release/expiry modes. A private PTY child mode remains legitimate internal Go plumbing, never user wiring. No server registry lookup, whichkey_runtime, completion action, advertisement, public relay, Lua action invocation or legacy action fallback selects or gates startup. Preserve help without effects and safe diagnostics. Keep useful standalone RunTerminal behavior; if retaining the low-level menu command, use the data-only contract and clearly separate UI/result use from actual public popup dispatch, never accept callback config. Generalize launchNativePopup to receive task1 config rather than always replacing it with nativeDefaults. Extend nativeSessionMenu's retained snapshot/provider handling to arbitrary configured prefixes, bounded sessions/windows/panes inventories and exact target dispatch; reuse Sessions, Session.Windows, Block/placement inspection and host pagination/status formatting. Add only narrow production-used helpers in providers.go and rex/actions.go/rex/wait.go as needed. windows/panes providers are opt-in, not resurrected default tab/focus lists. Preserve explicit labels, visibility, order, warning/empty submenus, full IDs and the no-truncation session bound; use bounded results and revalidate selected membership/placement before effects. Expand Go providers before popup sizing/snapshot so parent and child agree; do not leave host expansion independently altering the executable retained snapshot. known-hosts remains execution-host/home inventory, no attach/probe or --server retargeting. Remove ws/tn-only dispatch assumptions where custom keys must work. Preserve WhichKey inline tab/session rename, full text validation and parent-retained target; adapt form intent instead of relying on tr/wr keys for custom data mappings. No native rename dialog or dropped text.
Preserve exact origin session/block/window/layer, endpoint and separately intended GUI identity/generation. Lua key context supplies exact resource/app IDs only; Go validates/captures connection context, never treats a transient CLI caller or ctx.server as app/endpoint authority. Without explicit app intent, native display remains independent of GUI eligibility and app actions use the launching pane's validated owner; no app chooser/current focus fallback. Before dispatch revalidate origin placement, explicit intent and connection generation; restore origin and finish terminal, owned popup, private files/lifetime pipe and captured-image cleanup first. Retain confirmed resource closes and resource label rename via concrete resource methods. New Tab remains exactly one Session.NewWindow in the source session followed by exact-ID session.select activation on the retained app, with partial-created IDs preserved on failure and no rollback/retry. Other supported native client actions use bounded kind=client metadata and explicit client.perform; never delegate a named Lua action. Retain exact session/window destination arguments and split/session-select context-field exceptions. All Rex processes, including actions metadata and long wait, move behind shared rex argument-array execution; replace runtimeCommand's duplicate Rex exec path with concrete rex methods. Keep operation deadlines, parent cancellation and independent bounded cleanup, but remove InvocationLifetimeSeconds/receipt/Lua-action idle expiry from active menus. Long terminal waits must not inherit the short RPC timeout.
Retire runtime_bootstrap.go's consumer/receipt/relay/controller registration subsystem only after extracting its generic image capture/copy/verify/retire, strict JSON, private mode/owner/no-follow/slot checks, root-scoped fixed-name removal, result commit and frozen-CLI handoff helpers into popup_transport.go or a reduced existing file. popup.go/native_text_entry.go/terminal.go currently depend on these; deletion without moving live consumers is unacceptable. Keep existing Go test symbols/Completion validation and launchOutput compatibility where needed, explicitly isolated from public config/control paths; these are historical source compatibility, not supported hooks. No standards/test assertion weakening. Preserve rendering/theme/layout/initial-size and OS poll implementations unchanged unless a concrete owned lifecycle adaptation needs the existing caller files.
Make CLI discovery code-owned and independent of Lua anchor: frozen input > valid execution-host REX_BIN_DIR > applicable ordinary package dependency at ../libexec/rex-whichkey/rex relative to canonical public bin > validated unambiguous absolute PATH. An absent optional package dependency permits standalone PATH; an applicable broken entry or invalid higher-priority intent fails, never falls through. No public arbitrary executable flag, generated paths, wrapper or ldflags. Preserve endpoint precedence and pinned schema admission fixes (bounds tuple/anyOf and closed initial_windows envelope). Source review must trace every new symbol into the normal public path. Own formatting/static/type/build checks and, after independent source review, coordinator-observed normal public menu, navigation, dismissal and exact owned cleanup on the combined candidate. Builders request coordinated earlier observation if needed; they do not run product trials themselves. Do not complete this task when only compilation/private startup works. Record exact origin, binary, revision and observations; no config/reload is needed for direct public validation.
<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA2VDIyOjI4OjU4LjUzOFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxNi10MDIvc3VtbWFyeS8yIiwiaW50ZW50IjoiZDkyOGZiZDU1N2E4MDMzMmRiZDU3YjhkNzQxZWU4ZTdkYzk0MWU5YjIyMjQ0MWEyZmQ2ODY3ZDUwN2RhMTEzYyJ9 -->
### 2026-11-06T22:28:58.538Z · summary · unspecified

{"acceptedScope":"Go CLI-only ordinary public startup/data loader/providers/retained resource dispatch/inline forms/lifecycle/shared Rex CLI boundary; obsolete Lua runtime registry/completion/relay/controller path retired. Reviewed source actually built/run through normal task run, real root/Sessions/Switch menus observed, Backspace/Escape navigation and exact owned cleanup verified. Packaging/launcher/docs final Task3 and final combined observation remain separate.","actualRuntime":{"artifacts":[".mtn-tmp/p16/task2-public-created.json",".mtn-tmp/p16/task2-public-menu-root.txt",".mtn-tmp/p16/task2-public-menu-sessions.txt",".mtn-tmp/p16/task2-public-menu-session-choices.txt",".mtn-tmp/p16/task2-public-menu-returned-root.txt",".mtn-tmp/p16/task2-public-exit-wait.json",".mtn-tmp/p16/task2-public-after-dismiss-session.json",".mtn-tmp/p16/task2-public-parent-final-process.json",".mtn-tmp/p16/task2-public-parent-final-modes.json",".mtn-tmp/p16/task2-public-after-owned-window-cleanup.json"],"binary":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p16-t02-integration-a02/bin/darwin-arm64/rex-whichkey","binarySHA256":"7e2dd251be2f5847d45e772b6e4be01e59ed71b21aecf5655d44238fd28b4a62","candidateCommit":"7bdbda2b1e1e9d55ced9d26cc03204d61f505c00","cleanup":{"appRemovedExactPopupLayerAndBlock":true,"appRemovedPrivateDirectories":true,"coordinatorClosedOnlyOwnedValidationWindowAfterExit":true,"noUnrelatedResourcesClosed":true,"originFocusRestored":true,"parentExitedZero":true,"parentScreenPrimary":true},"command":"Normal task run in a new explicitly coordinator-owned Rex terminal/window in the existing execution session; no private entrypoint invoked directly","completedAtUTC":"2026-11-06T22:25:53Z","kind":"actual-public-runtime-evidence","limitations":["Actual normal direct public invocation on Darwin arm64 only.","No physical keybinding/reload/deploy, custom JSON probes, remote/Linux trial or non-owned effect.","No action selection/rename/close/split/new-tab effect exercised in this task; source reviewed, final Task3 may use benign owned inline rename.","No tests/harnesses/benchmarks; source/static checks and real app observation only."],"normalPublicArgs":[],"observed":["Root Rex menu with current reduced defaults","Sessions submenu","Go-built Switch session submenu with full IDs, current marker and actual two-session inventory","Backspace navigation to root","Escape ordinary dismissal","Public process exit0 and primary parent terminal screen"],"ownedResources":{"originBlock":"block:01a1b2f4-d767-738a-bde0-05caefae5542","popupBlock":"block:01a1b2f4-e1a4-7259-b680-9916b866b4d0","popupLayer":"layer:01a1b2f4-e1a5-74b3-a572-eb895f859624","privateDirectories":["/var/folders/r3/zcjrn53x42n1zybs9hx__h_w0000gn/T/rex-whichkey-816828524"],"session":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","validationWindow":"window:01a1b2f4-d768-7d23-a5b2-89fad13276e2"},"schemaVersion":1},"approvalSummary":["User accepted CLI-only REXWHICHKEY-20261106-114935 and start-it execution; hooks deferred.","Coordinator bounded Phase16 task plan approved/1 under explicit execution authority; current governing design/phase scope preserved.","Fresh independent same-family combined review rex-whichkey-p16-t02/review/2 approved exact build2eaec24a/tree44dc91b after fixing REX-P16-T02-001. Actual coordinator public runtime then passed; no compile-only completion."],"attempt":2,"cleanup":{"allRuntimeReservationsReleased":true,"facade":"complete invoked/retried; native/protocol complete. Immutable cleanup jq syntax defect persists; exact clean delivered successful worktrees removed with source/branch preservation and no live consumers. Failed reviewer kept.","removedCleanDeliveredWorktrees":["p16-t02-builder-a01","p16-t02-builder-a02","p16-t02-reviewer-a02","p16-t02-integration-a02"],"retainedNegativeReview":{"head":"14eaabad3b996bfc7cea9b19f724a911eae07ff1","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"75ff9ad4-a558-4295-8b20-855dafe77a8f"},"provider":{"displayName":"p16-t02-reviewer-a01","name":"rex","resources":{"block":"block:01a1b2d9-8d25-7616-b8fc-7bfcaf620ade","endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","session":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","window":"window:01a1b2d9-8d27-70c7-8a41-ff2a043821d1"}},"reason":"Stopped negative-review resource preserved, no failed/global cleanup authorization.","rexBlockID":"block:01a1b2d9-8d25-7616-b8fc-7bfcaf620ade","rexSessionID":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","rexWindowID":"window:01a1b2d9-8d27-70c7-8a41-ff2a043821d1","runID":"rex-whichkey-p16-t02-reviewer-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p16-t02-reviewer-a01"},"successfulWorkerWindowsConfirmedAbsent":true},"completionEventID":"rex-whichkey-p16-t02/completion/2","eventID":"rex-whichkey-p16-t02/summary/2","files":["cmd/main.go","launch.go","native.go","providers.go","popup.go","runtime_bootstrap.go","rex/actions.go","rex/wait.go","menu.go","hosts.go","docs/mapping-format.md",".mtn-tmp/p16/task2-a02-result.json",".mtn-tmp/p16/task2-a02-review-result.json",".mtn-tmp/p16/task2-integration.json",".mtn-tmp/p16/task2-public-runtime-evidence.json",".mtn-tmp/p16/task2-observed-rex-whichkey",".mtn-tmp/p16/task2-post-cleanup-worktrees.txt"],"followUp":"Task3 minimal launch-only Lua/ordinary package/migration docs and final combined actual public build/run, including benign owned inline rename where safe. Physical keybinding/user config/profile reload not authorized.","integrationEventID":"rex-whichkey-p16-t02/integration/2","kind":"summary","limitations":["Optional full lint failed62 inherited diagnostics; no global lint pass or AI-lint completion claimed.","No unit/race/TDD/Lua/parser/FIFO harnesses, benchmarks, custom config probes or remote/Linux trial.","Only actual normal default public menu/navigation/dismissal/popup/terminal/private cleanup exercised at Task2; selected real action effects not yet exercised.","Fresh same-family review, not cross-family verification.","Failed reviewer1 exact window/worktree retained; all Git branches/history/private evidence kept.","Phase15 failed original bridge integration/native incomplete status remains unchanged; no user config or installed assets modified."],"nativeEvidenceCommit":"1e8af3a509ac67d801437a2c4453799262a4d5ac","preservedObservedBinaryArtifact":".mtn-tmp/p16/task2-observed-rex-whichkey","recordRef":"tasks/rex-whichkey-16-02-go-cli-menu.md","requirements":["REQ-001","REQ-002","REQ-003","REQ-004","REQ-006","REQ-007","REQ-008","REQ-009","REQ-010","REQ-011","REQ-012","REQ-013","REQ-016","REQ-017"],"schemaVersion":1,"specHash":"sha256:d87f3737b1fbde061a559655f63ac2ea7d1bf14bd5ed580560ce69a14171a1b2","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"7bdbda2b1e1e9d55ced9d26cc03204d61f505c00","targetTree":"65c9a47ed3fc876e78f1324434b6a91152e69add","verification":[{"artifact":".mtn-tmp/p16/task2-integration-build.log","command":"Exact combined candidate task build","result":"pass"},{"artifact":".mtn-tmp/p16/task2-integration-vet.log","command":"Combined candidate task lint:vet and owned gofmt/whitespace","result":"pass"},{"artifact":".mtn-tmp/p16/task2-integration-source-identity.txt","command":"Reviewed21-file source identity, unchanged standards, clean candidate, unchanged target base, then fast-forward and reviewed-head target ancestry","result":"pass"},{"artifact":".mtn-tmp/p16/task2-public-runtime-evidence.json","command":"Normal public task run: observed menu/dynamic sessions/navigation/Backspace/Escape, exit0, primary terminal, exact owned popup/block/private-dir cleanup and owned validation-window closure","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
