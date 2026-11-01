# Release delivery inventory and limitations

## Scope and evidence boundary

This is a **source-based delivery inventory, not an acceptance-test result**.
The original Phase06 inventory described baseline commit
`4980e1f95a96731a1baee445562864c7e6e771bd` (tree
`68b72d7d8bdb2b4ccd99a7312c4bb9d3e772b2bb`), plus documentation task 06-03,
subsequently reviewed and integrated at `40c4f53` as recorded below. Those completed
phase/task identities and historical limitations are preserved.

The Phase07 documentation revision describes the integrated endpoint/launch
baseline `69078fcee3dc0bbd9598e1b58e33920934e82497` (tree
`82dc432f715295a3b82829be8be509e83aa0197a`), plus the six Markdown files owned by
task 07-03. It does not approve its own review, complete task 07-03/Phase07 or
establish functional acceptance. Final assembled-target compilation and Git
integration remain coordinator gates.

Requirements come from `.tasks/prds/rex-whichkey.md`, with stable REQ-001–REQ-015.
`.tasks/prds/rex-whichkey-usability.md` adds REQ-016/017 without renumbering or
reopening completed phases02–06.

The later direction **REXWHICHKEY-20261031-072107** supersedes earlier experimental,
TDD, live-Rex and final acceptance-matrix execution requirements, not product
behavior or safety intent. REQ-015 now records integrated delivery, documentation
and honest limitations. The authenticated post-return bridge follows accepted
**REXWHICHKEY-20261031-220348**; the separate constructor-bound endpoint follows
**REXWHICHKEY-20261101-015342**, selected under delegated implementation authority.
The latter is now **Refined** by accepted **REXWHICHKEY-20261101-101502** for the
user-requested default endpoints: native Rex discovery replaces mandatory explicit
setup, retaining overrides, opaque-context separation and fail-closed origin
checks. The duplicated-path planning proposal **REXWHICHKEY-20261101-100106** is
**Superseded**, not retroactively approved or implemented. This reflects the
coordinator's recorded lineage under delegated authority, not a new decision by
this documentation task. Historical draft IF-* signatures are not adopted contracts.

**Not executed under current delivery authority:** project tests, TDD, experiments,
Nix package builds/checks, Home Manager evaluation/activation, graphical checks,
live API probes, remote trials, disconnected-origin exercises and configuration
reloads. No graphical-client or remote-test-server resource is reserved. No new
acceptance runner or runtime verification is scheduled here. Existing test files
and historical evidence are preserved, not converted into current passes.

The new documentation work uses source inspection and whitespace checks only;
compilation is not authorized for its leaf builder. Earlier coordinator compilation
is recorded separately below. Reading source/docs, compilation and independent
source review do not establish runtime compatibility, installed configuration
success, exact graphical geometry or end-to-end behavior.

## Requirement-to-source inventory

“Implemented source” means the named files contain the path described; it does
not mean that the requirement passed a runtime check. Review references in the
next section identify the independently reviewed deliveries. Limits L1–L8 apply
as indicated and across the assembled application; L9 additionally covers the
Phase07 startup/default-endpoint paths.

| Requirement | Implemented source and integration | Outstanding verification / limitations |
| --- | --- | --- |
| REQ-001 — rewrite/lifecycle | [`cmd/main.go`](../cmd/main.go) composes root package `whichkey` (module remains `temp`); [`popup.go`](../popup.go) owns one bounded detached controller and PTY child; [`lua/rex_whichkey.lua`](../lua/rex_whichkey.lua) launches and returns. No persistent daemon or service hierarchy. Reviews 04-04/04-05. | Launch, prompt return, exit and resource/terminal cleanup are runtime-unverified. The controller is detached but invocation-bounded, not a service. L1/L3. |
| REQ-002 — typed CLI client | [`rex/rex.go`](../rex/rex.go), [`exec.go`](../rex/exec.go), [`session.go`](../rex/session.go), [`window.go`](../rex/window.go), [`block.go`](../rex/block.go), [`layer.go`](../rex/layer.go): concrete Client/Session/Window/Block, generic Call, argv execution, JSON/error handling, deadlines and captured targets. Reviews 02-01–02-04. | No current fixture, cancellation, invalid-JSON or live API tests were run. Timeout/decode failure does not prove a mutation had no effect. L1/L2/L3. |
| REQ-003 — Lua loading | [`lua/rex_whichkey/config.lua`](../lua/rex_whichkey/config.lua), [`lua/rex_whichkey.lua`](../lua/rex_whichkey.lua), [`config.go`](../config.go): server-side module/inline loading and strict resolved-data decoding. Defaults < imported < inline; spec/icons replace whole values. Review 03-01 and composition 04-05. | Actual module lookup, PATH and inline/import deployment remain unexecuted. Ordinary require caching applies; imports are trusted executable Lua. L1/L4. |
| REQ-004 — mappings | [`mapping.go`](../mapping.go) and [mapping-format.md](mapping-format.md) own shared data/validation; [`mapping.lua`](../lua/rex_whichkey/mapping.lua) compiles full sequences, implicit groups, nested inheritance, conflicts and `<Space>`/`<lt>`; [`evaluate.lua`](../lua/rex_whichkey/evaluate.lua) resolves conditions/dynamic strings and retains hidden active entries. Reviews 03-02/03-03/03-05, provider revision 05-04. | No new positive/negative fixtures executed. Lua consumes the Go-owned standard, not a parallel schema or Go callback VM. L1/L4. |
| REQ-005 — runtime callbacks | [`evaluate.lua`](../lua/rex_whichkey/evaluate.lua), [`dispatch.lua`](../lua/rex_whichkey/dispatch.lua), [`mapping.go`](../mapping.go), [`popup.go`](../popup.go) and Lua launcher: callbacks/args/tickets remain in the original server Lua VM; a fresh origin-targeted configuration action handles authenticated one-shot completion. Reviews 03-03/03-04/03-05/04-05. | No end-to-end callback run. Opaque server identity, connection-generation semantics and actual queued-client routing remain runtime-unverified. L2/L3/L4. |
| REQ-006 — navigation | [`menu.go`](../menu.go), [`action.go`](../action.go), [`render.go`](../render.go), [`terminal.go`](../terminal.go): arbitrary-depth groups, breadcrumbs, Esc/Backspace back/root dismissal, unmapped-q dismissal and Tab paging; hidden/off-page mappings stay active. Reviews 04-01/04-02/04-04/04-05. | No state-transition or interactive tests executed under current authority. L1/L5. |
| REQ-007 — confirmation | [`menu.go`](../menu.go) accepts exactly y at the prompt; any other key cancels without fallthrough. [`mapping.go`](../mapping.go), [`popup.go`](../popup.go) and [`dispatch.lua`](../lua/rex_whichkey/dispatch.lua) validate confirmation against the retained mapping for named actions and callbacks. Reviews 04-01/03-05/03-04/04-05. | No confirmation/cancellation/destructive-action runtime exercise. Confirmation is an authenticated popup assertion, not cryptographic proof of a physical keypress. L2/L4. |
| REQ-008 — geometry/rendering | [`layout.go`](../layout.go)/[`render.go`](../render.go) constrain cell-based columns, paging, grapheme-safe clipping, breadcrumbs/icons/help; [`popup.go`](../popup.go) selects native top/bottom/left/right/center bounds. [`terminal.go`](../terminal.go) consumes actual frame placements/Lines and passes Pages to key handling. Reviews 04-02/04-04/04-05. | All five positions, native floating-layer clipping, resize and tiny-terminal geometry remain graphically unverified. No old full-window overlay workaround is adopted. L1/L5. |
| REQ-009 — theme/terminal safety | [`theme.go`](../theme.go) supplies palette decoding, semantic contrast-aware styles, default/fallback colors, NO_COLOR and ASCII decorations; [`render.go`](../render.go) uses styles/decorations. [`terminal.go`](../terminal.go) orders palette metadata before keys and restores terminal state before returning a result. Reviews 04-03/04-02/04-04/04-05. | Real palette behavior, glyph widths and restoration were not exercised. Unknown background retains terminal defaults; ASCII does not transliterate user labels. L1/L3/L5. |
| REQ-010 — built-ins/session parity | [`defaults/sessions.lua`](../lua/rex_whichkey/defaults/sessions.lua), [`tabs.lua`](../lua/rex_whichkey/defaults/tabs.lua), [`panes.lua`](../lua/rex_whichkey/defaults/panes.lua), [`app.lua`](../lua/rex_whichkey/defaults/app.lua), [`init.lua`](../lua/rex_whichkey/defaults/init.lua) and dispatcher provide session/tab lifecycle/navigation, split/focus/zoom/close, sidebar and palette. Reviews 05-01–05-03/05-04 and 03-04. | Each actual Rex action, native rename/picker interaction and old-plugin runtime parity remain unexecuted. Session switch delegates Rex's picker, not an enumerated popup session list. L1/L2. |
| REQ-011 — origin/focus/cleanup | [`rex/block.go`](../rex/block.go)/[`layer.go`](../rex/layer.go), [`dispatch.lua`](../lua/rex_whichkey/dispatch.lua), [`popup.go`](../popup.go), Lua launcher: captured server/session/block/client/connection generation, current placement lookup, origin focus and returned-owned-terminal cleanup. Reviews 02-03/03-04/04-05 and documentation 05-05. | Moved-block, reconnect, focus, error/dismissal cleanup and disconnected-origin behavior are runtime-unverified, not passes. No local/current-client fallback or broad cascading cleanup is intended. L2/L3/L4. |
| REQ-012 — known hosts | [`rex/hosts.go`](../rex/hosts.go) decodes `rex hosts --json`; [`hosts.go`](../hosts.go)/[`terminal.go`](../terminal.go), [`defaults/hosts.lua`](../lua/rex_whichkey/defaults/hosts.lua) expand explicit known-hosts provider groups once per invocation. IDs, labels, home metadata and warnings remain distinct; generated rows are non-executable. Reviews 02-04/05-04. | No current discovery or reachability trial. A failed listing produces a warning, not successful empty inventory. Execution-host inventory differs from origin-server selection. L6. |
| REQ-013 — remote compatibility | Captured targeting in [`rex/`](../rex), server-local loading/launch in Lua/config/popup paths, and [remote-compatibility.md](remote-compatibility.md). Reviews 02-01–02-04/03-01/03-04/04-05/05-05. | No remote deployment/trial. Helper/controller/child must share the Rex server host's executable and temporary filesystem; client-local installation does not install remotely. L1/L2/L3/L4/L6. |
| REQ-014 — packaging/config ownership | [`flake.nix`](../flake.nix), [`nix/package.nix`](../nix/package.nix), [`nix/home-manager.nix`](../nix/home-manager.nix) package the executable/Lua and opt-in sibling links. User init/mapping files are not managed; links are non-forcing. Reviews 06-01/06-02. | No Nix package build/evaluation, Home Manager activation or existing-file byte-preservation trial. Declared four-system outputs are not certification. L7. |
| REQ-015 — integrated delivery/docs | [README](../README.md), this inventory, [mapping standard](mapping-format.md), [remote compatibility](remote-compatibility.md), [configuration examples](../examples/README.md) and [Home Manager cases](../tests/packaging/home-manager/README.md) describe the delivered source and retained limits. Prior deliveries, including 06-03, have recorded reviews below; current 07-03 documentation still requires independent review. | No acceptance runner, final GUI matrix, remote/disconnected-origin test or inferred functional pass. Historical evidence is not current execution. L8. |

### Usability extension: REQ-016/017

| Requirement | Implemented source and integration | Outstanding verification / limitations |
| --- | --- | --- |
| REQ-016 — ordinary popup launch | [`cmd/main.go`](../cmd/main.go) routes no arguments/public flags to [`launch.go`](../launch.go) `LaunchPopup`; it captures terminal session/block, selects an exact or unique eligible attached client/generation, preflights placement and activates the installed named action once. [`lua/rex_whichkey.lua`](../lua/rex_whichkey.lua) `menu.register_launch` uses the actual original-VM ctx, then the delivered defaults/evaluation/prepare/controller/terminal/completion route. Review 07-02; integrated `69078fc`. [Launch contract](launch-contract.md) traces the full path. | No installed no-argument run, popup display, Remote Control or callback exercise. Source reachability/compilation is not a runtime pass. Missing/stale/ambiguous origins fail; activation acceptance is not display success. L1–L5/L9. |
| REQ-017 — native platform endpoints | [`rex/endpoint.go`](../rex/endpoint.go) owns explicit override > nonempty execution-host REX_SERVER > bounded no-autostart native Linux/macOS Rex identity discovery, capturing/validating a concrete endpoint. [`popup.go`](../popup.go) preparation and the public caller consume it; Lua permits `popup.new(defaults)` and forwards preparation's canonical server. Private controller handoffs never ambient-default. Review 07-01, later composition review 07-02; integrated `d1164e7`/`69078fc`. | No live default discovery, custom-server or Darwin runtime exercise. Exact native socket algorithms belong to Rex; no guessed path formula or platform certification. CLI identity shape requires compatible installed Rex; no autostart, scan or alternate-server retry. L1/L2/L4/L9. |

### Phase07 intent, source traces and remaining gates

The accepted usability PRD/design define the intended product path; they are not
implementation evidence by themselves. The following traces were inspected in
the actual integrated baseline:

1. `cmd/main.go` `run`/`runLaunch` -> `LaunchPopup` -> one explicitly targeted
   configuration action -> `menu.register_launch` `run(ctx,args)` -> `menu.launch`.
   Public launch consumes no user JSON and never substitutes low-level `menu`.
2. Lua `config.new`/`mapping.compile` and shipped `defaults/init.lua` ->
   `dispatch.issue`/evaluation -> real `prepare` -> `PreparePopupWithServer` ->
   `rex.ResolveEndpoint` -> captured capability/executable/server -> `dispatch.bind`
   -> detached `RunPopupController`. The actual opaque ctx and original closures
   remain in the server VM; the launch envelope only constrains actual origin.
3. Controller exact-origin preflight -> owned floating terminal ->
   `RunPopupTerminal` -> `RunTerminal` -> `RenderMenu`/`HandleKey` -> terminal
   restoration/result commit -> owned popup cleanup/origin revalidation/focus ->
   authenticated fresh completion -> retained original-VM callback, `rex.invoke`
   or `rex.client.queue`, with unchanged replay/expiry/generation/snapshot/
   confirmation safeguards.

These are source-consistency observations, not a new independent review of this
documentation. The supplied coordinator records below establish earlier producer
review/compilation/integration only. Task 07-03 still needs fresh independent
review of all six Markdown changes against the integrated Go/Lua path, followed
by coordinator `task build` on the exact final combined candidate under its gate
allocation, reviewed-diff/commit/tree/ancestry verification and actual main
integration. The coordinator alone records phase delivery/completion. No final
`task build`, integration of this revision or Phase07 completion is claimed here;
no runtime gate or example execution is scheduled.

## Recorded independent source-review outcomes

These are **reported coordinator delivery records**, not reviews performed by this
documentation builder. Each listed outcome is APPROVED for its frozen source
revision and subsequently recorded as integrated. The reviewer contexts were
fresh but same-family; cross-family verification was unavailable. Later changes
to shared files are represented by their later review row, especially 05-04.
No approval here means runtime acceptance.

Task shorthand `PP-TT` refers to the exact linked record in the authoritative
parent `.tasks` collection. Commit abbreviations identify integrated source, not
worker heads or runtime binaries. Full identities and artifact pointers remain
in each record's compact delivery summary.

| Task record | Recorded APPROVED review event | Integrated commit |
| --- | --- | --- |
| [02-01 CLI execution](../.tasks/tasks/rex-whichkey-02-01-cli-execution.md) | `rex-whichkey-p02-t01-reviewer-a01/review/1` | `020a31e` |
| [02-02 resources](../.tasks/tasks/rex-whichkey-02-02-resources.md) | `rex-whichkey-p02-t02-reviewer-a01/review/1` | `3ec92e9` |
| [02-03 blocks/layers](../.tasks/tasks/rex-whichkey-02-03-block-layers.md) | `rex-whichkey-p02-t03-reviewer-a02/review/2` | `fb7ed12` |
| [02-04 inventory](../.tasks/tasks/rex-whichkey-02-04-host-inventory.md) | `rex-whichkey-p02-t04-reviewer-a01/review/1` | `4335256` |
| [03-01 Lua loading](../.tasks/tasks/rex-whichkey-03-01-lua-config.md) | `rex-whichkey-p03-t01-tasks-a01/review/1` | `6109fff` |
| [03-02 mapping compiler](../.tasks/tasks/rex-whichkey-03-02-mapping-keys.md) | `rex-whichkey-p03-t02-tasks-a01/review/1` | `39ca52f` |
| [03-03 evaluation](../.tasks/tasks/rex-whichkey-03-03-dynamic-mappings.md) | `rex-whichkey-p03-t03-tasks-a01/review/1` | `44f22c1` |
| [03-04 dispatch](../.tasks/tasks/rex-whichkey-03-04-invocation-dispatch.md) | `rex-whichkey-p03-t04-tasks-a02/review/2` | `0a0ac89` |
| [03-05 Go-owned standard](../.tasks/tasks/rex-whichkey-03-05-mapping-format.md) | `rex-whichkey-p03-t05-reviewer-a03/review/3` | `55da14c` |
| [04-01 navigation](../.tasks/tasks/rex-whichkey-04-01-navigation.md) | `rex-whichkey-p04-t01-reviewer-a02/review/2` | `2d612d6` |
| [04-02 layout](../.tasks/tasks/rex-whichkey-04-02-layout.md) | `rex-whichkey-p04-t02-reviewer-a03/review/3` | `43c62df` |
| [04-03 theme](../.tasks/tasks/rex-whichkey-04-03-theme.md) | `rex-whichkey-p04-t03-reviewer-a02/review/2` | `f71fd47` |
| [04-04 terminal](../.tasks/tasks/rex-whichkey-04-04-terminal-lifecycle.md) | `rex-whichkey-p04-t04-reviewer-a03/review/3` | `0d8fc4c` |
| [04-05 popup composition](../.tasks/tasks/rex-whichkey-04-05-popup-integration.md) | `rex-whichkey-p04-t05-reviewer-a03/review/3` | `8f2a196` |
| [05-01 sessions](../.tasks/tasks/rex-whichkey-05-01-session-actions.md) | `rex-whichkey-p05-t01-tasks-a01/review/1` | `647b29b` |
| [05-02 tabs](../.tasks/tasks/rex-whichkey-05-02-tab-actions.md) | `rex-whichkey-p05-t02-tasks-a01/review/1` | `c8c810f` |
| [05-03 panes/application](../.tasks/tasks/rex-whichkey-05-03-pane-actions.md) | `rex-whichkey-p05-t03-tasks-a01/review/1` | `7525b2c` |
| [05-04 host menu/default composer](../.tasks/tasks/rex-whichkey-05-04-host-menu.md) | `rex-whichkey-p05-t04-reviewer-a02/review/2` | `cfba4cf` |
| [05-05 remote documentation](../.tasks/tasks/rex-whichkey-05-05-remote-context.md) | `rex-whichkey-p05-t05-tasks-a02/review/2` | `cddf437` |
| [06-01 Nix package](../.tasks/tasks/rex-whichkey-06-01-nix-package.md) | `rex-whichkey-p06-t01-tasks-a01/review/1` | `908a339` |
| [06-02 Home Manager](../.tasks/tasks/rex-whichkey-06-02-home-manager.md) | `rex-whichkey-p06-t02-tasks-a01/review/1` | `4980e1f` |
| [06-03 release documentation](../.tasks/tasks/rex-whichkey-06-03-release-acceptance.md) | `rex-whichkey-p06-t03-tasks-a01/review/1` | `40c4f53` |
| [07-01 endpoint policy](../.tasks/tasks/rex-whichkey-07-01-endpoint-policy.md) | `rex-whichkey-p07-t01-reviewer-a02/review/2` | `d1164e7` |
| [07-02 ordinary launch](../.tasks/tasks/rex-whichkey-07-02-direct-launch.md) | `rex-whichkey-p07-t02-reviewer-a01/review/1` | `69078fc` |

The 02-03 summary explicitly records cleanup-race finding `P02-T03-R1` resolved
before its approved second review; it retains refusal of unsafe remaining-layer
removal. The later final approvals do not relabel earlier blocked attempts as
successful. Endpoint task 07-01's earlier actionable-error finding
`REXWK-P07-T01-R1` was resolved before its approved second review; the earlier
blocked attempt is not relabeled a pass. The current 07-03 documentation candidate
is outside this prior-review ledger and is not self-approved by inclusion here.

### Historical compilation versus current checks

The Phase 2 summaries record coordinator `go build ./rex` on their combined
integration candidates. Later task summaries record coordinator
`GOTOOLCHAIN=local GOMAXPROCS=16 go build ./...` plus source/identity checks on
combined candidates, including 06-03 and the two Phase07 producer records below.
Those are **supplied earlier compilation records**, not commands rerun by this
task. They do not compile Lua, evaluate Nix, exercise existing tests or prove that
the packaged program starts correctly. No Nix package-build success is inferred
from Go compilation or Nix source review.

For Phase07, the authoritative completed-task summaries and these coordinator
integration events report `GOTOOLCHAIN=local GOMAXPROCS=16 go build ./...` passed
on each combined candidate, unchanged reviewed file contents, `git diff --check`
and verified target identity/ancestry:

| Coordinator integration event | Integrated commit / tree | Retained artifact relative to parent checkout |
| --- | --- | --- |
| `rex-whichkey-p07-t01-builder-a02/integration/2` | `d1164e75028760421446a926421d50750ddade28` / `07f7faabe9248e41c94a6d9e05ede573301273cc` | `.mtn-tmp/usability/rex-whichkey-p07-t01-builder-a02/integration/integration.json` |
| `rex-whichkey-p07-t02-builder-a01/integration/1` | `69078fcee3dc0bbd9598e1b58e33920934e82497` / `82dc432f715295a3b82829be8be509e83aa0197a` | `.mtn-tmp/usability/rex-whichkey-p07-t02-builder-a01/integration/integration.json` |

The corresponding fresh mixed Go/Lua review artifacts are
`.mtn-tmp/usability/rex-whichkey-p07-t01-reviewer-a02/review.json` and
`.mtn-tmp/usability/rex-whichkey-p07-t02-reviewer-a01/review.json`. They approve
frozen worker trees matching the integrated trees above. These ignored transport
files are coordinator-retained evidence, not installed package files; durable
compact summaries remain in the linked authoritative task records. Neither
producer record asserts a final Phase07 `task build` pass or runtime acceptance.
The documentation leaf uses source inspection and whitespace checks only; its
actual static results and exact revision are returned in its build payload,
separate from these historical results.

### Preserved earlier experiments

The nested [terminal experiment](experiments/rex-validation/tui/VALIDATION.md) and
[attempt A02](experiments/rex-validation/tui/VALIDATION-A02.md),
[A03](experiments/rex-validation/tui/VALIDATION-A03.md),
[A04](experiments/rex-validation/tui/VALIDATION-A04.md) documents retain earlier
failures, corrective work and recorded unit/PTY/race/root-gate results. They
predate the stop instruction and describe an isolated diagnostic module, not the
assembled product. A04's green gates meant readiness for independent review only,
not GUI qualification, framework adoption or release acceptance. Its preserved
limits include an earlier unexplained transient poll-cancellation failure,
intersecting-wide-cell behavior and OSC-looking bracketed paste not claimed fixed
there. Those historical experiment limits/results must not be silently asserted
as either confirmed defects or passes in the different current product paths.
Root commands did not traverse that nested module. No historical command was
rerun or scheduled by this documentation task; cancelled research is not product
delivery.

## Outstanding limitations

- **L1 — deployment/API/platform:** no minimum tested Rex build or Linux/macOS
  runtime certification. The flake declares four systems; popup/terminal code
  uses Unix facilities. Compatible installed APIs, module lookup, server PATH,
  executable permissions and shared server-local temporary filesystem are
  prerequisites, not verified deployment outcomes. CLI-origin client/block
  actions require the graphical client's Remote Control setting. Configuration
  changes remain user-owned; no init.lua rewrite, automatic reload or restart.
- **L2 — origin/races:** source targets the captured server/session/block/client
  and `connected_at` generation, following block moves within the captured
  session, not searching other sessions. Missing/ambiguous/lost/changed origins
  fail closed. Inspect/focus/execute are separate Rex operations, not an atomic
  transaction; external moves, disconnects, focus or configuration changes can
  race side effects after the last check. Disconnected-origin rejection and real
  graphical-client action routing remain untested.
- **L3 — cleanup/termination/uncertainty:** only the returned owned terminal is
  closed; a remaining layer causes an error rather than cascading closure.
  Creation can succeed before timeout/decoding loses its resource identity.
  Network/server loss, SIGKILL/SIGSTOP and terminal loss can prevent cleanup.
  Terminal restoration is bounded best effort, with the documented transient
  Darwin PENDIN caveat, not a byte-for-byte platform comparison. The controller
  reserves five seconds of its original lifetime for UI teardown/completion and
  has bounded deferred cleanup. No retry of uncertain mutations/completion,
  label-based scavenging, recursive cleanup or compensating daemon is provided.
- **L4 — trust/lifetime:** source binds a Go crypto/rand capability, retained
  mapping/confirmation, original runtime context, config and connection
  generations to one invocation. Lifetime is 30 seconds from reservation, with
  at most 64 pending/in-progress entries per registry. Lua expiry is lazy on
  registry operations; backwards wall-clock changes invalidate the registry.
  Private pipes/files are not isolation from the same OS account, trusted user
  Lua or privileged processes. Capabilities and original tickets must not be
  logged/shared. Callback side effects cannot be rolled back. Rex's opaque
  server identity and connection-generation semantics remain unexercised.
- **L5 — terminal/geometry:** direct I/O, palette replies, grapheme clipping,
  resize, all five positions and native-layer clipping are source-reviewed only.
  Real fonts, terminal widths, Unicode presentation, palette/contrast and tiny
  viewports remain graphically unverified. ASCII flags exist in low-level menu/
  popup modes but are not forwarded by the delivered Lua launcher; a nonempty
  NO_COLOR environment value is the available launcher-path color opt-out.
- **L6 — inventory/deferred remote scope:** listing reads the execution host's
  effective Rex home inventory, not inventory selected by popup `--server`.
  Rex may contact home and forward held changes. Empty/failed/stale listings and
  endpoints are not reachability results. Rows cannot execute/attach or retarget
  the popup. No independent registry, remote attachment/connection, bootstrap,
  Wake-on-LAN, SSH, retries or reachability protocol is delivered.
- **L7 — packaging/user files:** package recipes have `doCheck = false` and
  `doInstallCheck = false` under current policy, not passing package tests.
  Home Manager defaults disabled, links only the reserved namespace and never
  forces unmanaged files. Existing-file preservation, conflicts and enabled/
  disabled activation have not been exercised. Whole-directory Rex symlinks need
  user-owned layout integration. Removing the module does not remove user-added
  require/setup wiring; users must account for that themselves.
- **L8 — delivery is not acceptance:** recorded reviewed/integrated tasks,
  coordinator compilation and this inventory do not assert a working installed
  release. This task still needs independent source review and coordinator
  integration. There is no runtime acceptance pass, coverage claim or outstanding
  instruction to execute a final matrix under the current stop policy.
- **L9 — ordinary launch/native defaults:** REQ-016/017 are implemented and the
  producer revisions have recorded review/compilation/integration, not live
  startup or platform verification. A reachable already-running default server,
  valid execution-host Rex runtime/home settings and compatible CLI identity JSON
  are required for native discovery. Present invalid overrides fail rather than
  falling through; controller handoffs never default. Terminal and server-side
  preparation must resolve the same canonical endpoint for custom/shared setups.
  Full captured session/block and an exact/unique eligible attached client with
  unchanged connection generation are mandatory; no focus/current-client/server
  substitution is authorized. No-argument activation may be accepted without
  popup display, and uncertainty is not retried. Current documentation still
  needs independent review and the coordinator's final assembled `task build`/
  integration gate; skipped tests/runtime gates are not passing usability criteria.

## Usage and published contracts

The [README](../README.md) covers installation layout, explicit user-owned
registration, imported/inline configuration, default keys and troubleshooting.
[launch-contract.md](launch-contract.md) details ordinary no-argument popup launch,
exact client/origin selection and native/default/custom endpoint routing;
[mapping-format.md](mapping-format.md) is the Go-owned shared standard;
[remote-compatibility.md](remote-compatibility.md) details endpoint/runtime identity,
server-local execution and cleanup boundaries. The
[Home Manager examples](../tests/packaging/home-manager/README.md) are source-review
cases, not an activation runner.

Published Rex references used by the implementation/documentation:
[CLI](https://www.superlogical.com/rex/docs/automate/cli),
[Server API](https://www.superlogical.com/rex/docs/automate/server-api),
[configuration](https://www.superlogical.com/rex/docs/customize/config),
[custom actions](https://www.superlogical.com/rex/docs/customize/actions),
[Lua API](https://www.superlogical.com/rex/docs/reference/lua) and
[Lua scripts](https://www.superlogical.com/rex/docs/automate/lua-scripts).
These contract references are not live-server observations or new verification
claims.
