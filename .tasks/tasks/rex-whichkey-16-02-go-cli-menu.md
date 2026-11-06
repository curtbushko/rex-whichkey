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
  state: planned
  sourceEventID: null
  changedAt: '2026-11-06T18:41:57.749Z'
title: Run the ordinary Go CLI-only menu with dynamic targets and real actions
mtnCreateRetry:
  key: rex-whichkey-16-02-go-cli-menu-create-v1
  intent: 1aa07d9252a08eb35486f217eba2a5a69c8c6e6aeb2f68d987d6dd936c9cec6c
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-06T18:41:58.393Z'
dateModified: '2026-11-06T18:47:34.254Z'
tags:
  - task
type: task
id: 148e98d8-5909-4ad3-967c-601aa579efa9
projects:
  - '[[phases/rex-whichkey-16-cli-only-menu.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-16-01-data-config.md]]'
    reltype: FINISHTOSTART
---
Consume task1's data contract and deliver a REAL RUNNING public rex-whichkey binary, not a new private native entry or compile-only milestone. Wire cmd/main.go/runLaunch -> LaunchPopup -> Go-loaded configuration -> existing native popup/terminal/menu lifecycle -> retained selection -> real Go rex dispatch. Public no-argument and --config/--server/--client/--ascii/--no-color invocation use this one path; remove --action routing and ordinary prepare/controller/controller-child/runtime-describe/runtime-relay/runtime-release/expiry modes. A private PTY child mode remains legitimate internal Go plumbing, never user wiring. No server registry lookup, whichkey_runtime, completion action, advertisement, public relay, Lua action invocation or legacy action fallback selects or gates startup. Preserve help without effects and safe diagnostics. Keep useful standalone RunTerminal behavior; if retaining the low-level menu command, use the data-only contract and clearly separate UI/result use from actual public popup dispatch, never accept callback config. Generalize launchNativePopup to receive task1 config rather than always replacing it with nativeDefaults. Extend nativeSessionMenu's retained snapshot/provider handling to arbitrary configured prefixes, bounded sessions/windows/panes inventories and exact target dispatch; reuse Sessions, Session.Windows, Block/placement inspection and host pagination/status formatting. Add only narrow production-used helpers in providers.go and rex/actions.go/rex/wait.go as needed. windows/panes providers are opt-in, not resurrected default tab/focus lists. Preserve explicit labels, visibility, order, warning/empty submenus, full IDs and the no-truncation session bound; use bounded results and revalidate selected membership/placement before effects. Expand Go providers before popup sizing/snapshot so parent and child agree; do not leave host expansion independently altering the executable retained snapshot. known-hosts remains execution-host/home inventory, no attach/probe or --server retargeting. Remove ws/tn-only dispatch assumptions where custom keys must work. Preserve WhichKey inline tab/session rename, full text validation and parent-retained target; adapt form intent instead of relying on tr/wr keys for custom data mappings. No native rename dialog or dropped text.
Preserve exact origin session/block/window/layer, endpoint and separately intended GUI identity/generation. Lua key context supplies exact resource/app IDs only; Go validates/captures connection context, never treats a transient CLI caller or ctx.server as app/endpoint authority. Without explicit app intent, native display remains independent of GUI eligibility and app actions use the launching pane's validated owner; no app chooser/current focus fallback. Before dispatch revalidate origin placement, explicit intent and connection generation; restore origin and finish terminal, owned popup, private files/lifetime pipe and captured-image cleanup first. Retain confirmed resource closes and resource label rename via concrete resource methods. New Tab remains exactly one Session.NewWindow in the source session followed by exact-ID session.select activation on the retained app, with partial-created IDs preserved on failure and no rollback/retry. Other supported native client actions use bounded kind=client metadata and explicit client.perform; never delegate a named Lua action. Retain exact session/window destination arguments and split/session-select context-field exceptions. All Rex processes, including actions metadata and long wait, move behind shared rex argument-array execution; replace runtimeCommand's duplicate Rex exec path with concrete rex methods. Keep operation deadlines, parent cancellation and independent bounded cleanup, but remove InvocationLifetimeSeconds/receipt/Lua-action idle expiry from active menus. Long terminal waits must not inherit the short RPC timeout.
Retire runtime_bootstrap.go's consumer/receipt/relay/controller registration subsystem only after extracting its generic image capture/copy/verify/retire, strict JSON, private mode/owner/no-follow/slot checks, root-scoped fixed-name removal, result commit and frozen-CLI handoff helpers into popup_transport.go or a reduced existing file. popup.go/native_text_entry.go/terminal.go currently depend on these; deletion without moving live consumers is unacceptable. Keep existing Go test symbols/Completion validation and launchOutput compatibility where needed, explicitly isolated from public config/control paths; these are historical source compatibility, not supported hooks. No standards/test assertion weakening. Preserve rendering/theme/layout/initial-size and OS poll implementations unchanged unless a concrete owned lifecycle adaptation needs the existing caller files.
Make CLI discovery code-owned and independent of Lua anchor: frozen input > valid execution-host REX_BIN_DIR > applicable ordinary package dependency at ../libexec/rex-whichkey/rex relative to canonical public bin > validated unambiguous absolute PATH. An absent optional package dependency permits standalone PATH; an applicable broken entry or invalid higher-priority intent fails, never falls through. No public arbitrary executable flag, generated paths, wrapper or ldflags. Preserve endpoint precedence and pinned schema admission fixes (bounds tuple/anyOf and closed initial_windows envelope). Source review must trace every new symbol into the normal public path. Own formatting/static/type/build checks and, after independent source review, coordinator-observed normal public menu, navigation, dismissal and exact owned cleanup on the combined candidate. Builders request coordinated earlier observation if needed; they do not run product trials themselves. Do not complete this task when only compilation/private startup works. Record exact origin, binary, revision and observations; no config/reload is needed for direct public validation.