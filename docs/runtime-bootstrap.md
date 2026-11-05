# Code-owned runtime bootstrap contract

## 1. Status, authority and evidence

**Phase 15 implementation contract candidate; source-reviewed intent, not a
working-runtime claim.** Downstream release requires independent combined
specification/security review and coordinator integration. This task implements
no Go, Lua, packaging or consumer configuration changes.

Authority: accepted responsibility correction `REXWHICHKEY-20261104-193124`,
`designs/rex-whichkey-runtime-ownership.md`, and Phase 15. Discovery and
compatibility belong to application code. Packaging installs normal assets;
user configuration registers keys/menus and retains customization. No new
architectural framework, daemon, socket client, callback VM, arbitrary script
injection, automatic configuration edit/reload/restart or server replacement is
part of this contract.

Inspected application baseline:
`664f8fb7e4f0a7f0a7b9eb8b943dd144d92cb742`. Its ancestor
`30a4439302ade4e08487d176ba6823206ddc326d` preserves pre-existing endpoint/PTY,
GUI/caller, configured-rename and packaging repairs **without source review or
Phase 15 delivery**. Those repairs are inputs, not this task's accomplishments or
test evidence. Attempt 1's frozen, uncommitted document and the independent
non-gating diagnosis are historical inputs, not approved algorithms or workflow
authority. No legacy managed records are migrated.

### Supported basis, including the restricted-PATH resolution

- The **complete introduction** to [Rex Lua][lua] promises Lua 5.1's full standard
  library, including `io`, `os` and `require`; only `io.popen` and `os.execute`
  are disabled. Use `rex.exec` for commands. A long-page outline omitting this
  introduction is not evidence that `debug`, loaders or file I/O are absent.
- [Configuration modules][modules] and standard Lua file-backed loading provide
  the module-relative anchor below. `debug.getinfo(..., "S").source` is a file
  label, not an executable authorization API. Synthetic loaders are not trusted
  merely because their label starts with `@`.
- [Terminal PATH][path] guarantees `REX_BIN_DIR` for **terminal children**. Neither
  an environment variable observed in this worker terminal nor `os.getenv` proves
  that configuration actions receive that injection. This contract makes no such
  claim. `rex.exec`'s documented PATH lookup cannot find an absent dependency.
- A normal application-owned CLI dependency is available from Rex's **official
  published package source**, not inferred from an app-bundle path. Inspected
  `https://pkgs.superlogical.com/rex/nix/preview/0.1.0-preview.13.tar.gz`:
  SHA256 `dc5297710269f1eff712b265465c851f2ddc953a61d4310f1bb81bf9d543ac12`.
  Its `sources.json` identifies Rex commit
  `1b6654811346488b696ded8f3c7e6701c57242f1` and native artifacts for
  `aarch64-darwin`, `aarch64-linux`, and `x86_64-linux`. Its `flake.nix` installs
  `$out/bin/rex`; on macOS it preserves released bytes/signature, on Linux it
  supplies the normal loader/library dependencies. This versioned source is
  stronger evidence for Apple-silicon availability than the Linux-only summary
  in [the installation page][nix-rex]. No binary was downloaded or executed.
- Rex is proprietary (`meta.license = lib.licenses.unfree`). Install its dependency
  through the official upstream distribution under the user's applicable terms;
  do not vendor its bytes into this repository, relabel them MIT, or publish a
  bundled binary/cache closure without redistribution rights. An ordinary local
  dependency link is not a grant to redistribute Rex. This contract requires no
  new server service, Rex Home Manager module, account configuration or activation.

Thus restricted-PATH launch uses an explicitly installed **normal dependency
asset**, selected by Go at runtime. It does not depend on a new upstream
introspection API, a configuration-action environment guarantee, runtime Nix
queries, constructor injection or manually repairing server PATH. The cited
release is evidence of an available distribution mechanism, not a demand that
users pair that exact release with WhichKey or their server.

## 2. Baseline trace and target identity selection

| Stage | Inspected baseline | Target owner and identity |
| --- | --- | --- |
| Public CLI | `cmd/main.go` calls `LaunchPopup`; current process is caller-selected. | Go selects its own executable at entry. A registered menu cannot replace it with the installed popup helper. |
| Rex CLI / endpoint | `rex/exec.go` resolves explicit executable, `REX_BIN_DIR`, PATH; `popupRexExecutable` duplicates resolution; endpoint discovery resolves again. | One `rex` package selection, frozen before endpoint discovery, then passed unchanged through all consumers (§4). |
| Public routing | `launch.go` reads action metadata; absent implicit default permits native; registered/default or explicit action retains configured intent. | Read-only original-VM compatibility query precedes configured activation. Failure never becomes native fallback. |
| Key/palette | Original configuration VM calls `menu.launch(ctx)`. | Fixed module-relative helper initiates Go preparation. There is no public executable to nominate. |
| Configured menu | `require`, `setup` and `dispatch.new` retain Lua code/registry. | Same VM and exact installed menu/registry; negotiate its actual loaded capabilities, not the currently installed Lua files' version. |
| Preparation | Lua runs configured `helper prepare`, optionally forwards `rex_bin_dir`; Go returns own executable, endpoint and capability. | Selected Go component validates config, CLI, endpoint, compatibility and rename support before issuing authority. Public uses public Go; key uses anchored Go. |
| Controller / PTY | Controller rediscovers CLI; `os.Executable()` selects popup child. | Controller and child use the frozen selected Go image; never redo PATH/asset/endpoint selection after detach. |
| Completion | Controller invokes configuration completion action via fresh CLI; `dispatch.lua` claims original registry. | Frozen Rex CLI and endpoint; same original-VM consumer/menu receipt. Actual completion caller remains separate from retained GUI target. |
| Installation | `nix/package.nix` copies Lua; HM links only reserved module namespace. | Extend that normal asset tree (§3); no generated paths in Lua, environment wrappers or runtime Nix resolver. |

The **execution host** is where the selected processes and server's configuration
VM can access the invocation's files. A viewing Mac is never a source of Linux
paths. Explicit `--server` changes transport intent, not filesystem locality.
Public configured launch without shared execution-host files fails before popup;
run the public executable on the execution host or use its key/palette action.
Native public launch requires the documented originating Rex-terminal context
and its captured endpoint; an explicit endpoint differing from that origin does
not authorize spawning this machine's executable remotely. Reject absent or
contradictory execution-host provenance before constructing the native popup.
No remote attach/bootstrap or alternate-server probing is added.

## 3. Normal asset layout and trust boundary

The reserved `rex_whichkey` integration namespace contains:

```text
rex_whichkey/runtime/anchor.lua
rex_whichkey/runtime/rex-whichkey
rex_whichkey/runtime/deps/rex
```

`anchor.lua` is a shipped file-backed module returning the layout revision `1`
and its own source label captured while that file loads. Its sibling helper is
the application binary, installed once or linked to the same installed binary
as the public entrypoint. `deps/rex` is a normal dependency link to the official
host-platform Rex package's `bin/rex`. Keep upstream runtime dependencies and the
macOS signature intact. These names are code-owned constants, **not** generated
absolute strings. A package manager may make the link, just as it installs other
assets; it neither chooses a live endpoint nor decides compatibility or fallback.
Non-Nix installers must provide the same relative layout and licensed dependency.

The existing HM namespace links can expose this subtree without editing
`init.lua`. Do not strip `share/lua` suffixes, walk parents looking for a prefix,
search the store, inspect a viewing client's application bundle, or guess home,
repository, profile or system directories.

### Anchor validation and supported loader bounds

1. Use the normal file loader for the reserved anchor module, loaded through the
   same trusted integration namespace as `rex_whichkey`. Capture the label once
   at load; it must be `@` followed by a nonempty file path ending in the exact
   `rex_whichkey/runtime/anchor.lua` components, with no NUL/control bytes or `..`
   components. Synthetic `=...`, string chunks, preload/custom loader anchors,
   missing source metadata and cross-platform path syntax are unsupported.
2. Absolute labels are accepted. A relative label is supported only with the
   standard file loader's unchanged server working directory: the same relative
   sibling is passed to `rex.exec`, whose default cwd is Rex's own. Do not resolve
   it against `PWD`, a user's terminal cwd or `os.getenv("HOME")`. If this loader/
   cwd contract cannot be established, reject; do not guess an absolute path.
3. Append only the fixed sibling helper name; do not accept a helper argument
   from an action or constructor. Initially executing that sibling has the same
   installation trust as loading the executable Lua module itself. Source labels
   are not a sandbox against malicious same-user configuration/loaders. Require
   the opted-in integration tree and dependencies to be installer/user-owned,
   without untrusted writers; do not claim to authenticate arbitrary Lua code.
4. The helper canonicalizes its executable, anchor and dependency with Go file
   APIs, opens/checks regular executable files and verifies the declared layout.
   Reject dangling/cyclic symlinks, nonregular assets, unexpected helper targets,
   unsupported layout revisions, mismatched namespace modules, ownership or
   unsafe write permissions. The helper must not execute another candidate to
   see whether it works. An install-owned dependency symlink may leave the
   WhichKey tree; its canonical target and identity are frozen, not guessed.
5. Symlinked namespace directories are supported, including HM's normal links.
   Preserve the logical file-loader anchor for future launches; canonicalize a
   fresh snapshot at each launch. An atomic namespace update may select new
   assets for the **next** invocation, never change an open invocation. A loader
   pinned directly to an immutable old installation is explicitly pinned; code
   does not scan for its replacement. Normal next configuration load selects the
   new installation; a public rebuilt caller still uses its own component.
6. One loaded integration bundle owns the reserved namespace and one compatibility
   action. Duplicate registration or modules resolved from different trees fails
   registration/preflight. Do not choose by PATH order or silently reload modules
   to repair shadowing. User mapping modules live outside the reserved namespace.

At launch Go freezes canonical paths, executable content hashes, file identities,
platform, embedded asset digest and protocol/features. Copies of the selected
WhichKey executable used for re-exec live in the invocation-owned private workspace;
verify the copy against the selected open source before use. Never copy/modify the
proprietary Rex dependency: retain its canonical package path and identity, and
recheck before spawning. A replacement/removal or changed symlink after selection
aborts; no second candidate is selected. Package-managed immutable targets avoid
normal upgrade races. Rechecks are not an atomic OS execute-by-descriptor guarantee
on every platform; hostile same-user replacement is outside the trusted-install
boundary, not a security property supplied by `os.Executable()`.

Local development's public binary need not live beside installed Lua. It embeds
its own Go-owned format/asset identity and uses the already-loaded compatible Lua
bridge only for original-VM customization. It does not load a second callback VM
or defer popup preparation to an installed helper. Bare `go build`/the existing
Taskfile remain valid; no generated user paths or special Nix rebuild is needed.
An arbitrary standalone binary outside an asset layout still needs a documented
terminal `REX_BIN_DIR` or usable execution-host PATH for its initial Rex CLI.

## 4. RUNTIME-CLI-DISCOVERY

**Producer:** existing concrete `rex` package, not a new dependency-injection
framework. **Consumers:** native endpoint discovery, public metadata/activation,
key preparation, controller, resource calls, PTY-related calls and completion.

Inputs are execution-host environment and the validated application layout, plus
an optional already-frozen internal selection. No public helper-directory option
or Lua runtime constructor is added. Ordered selection:

1. An already-frozen selection supplied by trusted Go composition is reused and
   checked, never resolved again. Existing `rex.Options.Executable` remains an
   internal explicit selection; malformed/present-but-unusable intent fails.
2. Nonempty execution-host `REX_BIN_DIR`, when available, selects its absolute
   `rex`. Its **documented provenance is terminal launch only**. Configuration
   action environments may supply it as an ordinary environment preference, but
   neither its presence nor compatibility is assumed; validate the candidate.
3. The validated fixed-relative `runtime/deps/rex` dependency. This is the
   restricted-PATH configuration-action path and the packaged fallback when the
   terminal variable is absent. An installed-but-broken asset is an installation
   failure, not permission to continue searching.
4. Only when neither environment preference nor applicable dependency layout is
   present: execution-host PATH. Reject empty/relative/cwd entries; canonicalize
   candidates and deduplicate aliases to one file. Distinct executable candidates
   are ambiguous, not a reason to select whichever appears first. Missing PATH
   produces a fixed discovery diagnostic. No login shell or external `which`.

Invalid explicit preferences do not fall through. Freeze one candidate before
any endpoint operation. Output is an internal selection containing absolute
canonical executable, content/file identity, provenance (`terminal-env`,
`execution-env`, `layout-dependency`, `path`, or trusted internal reuse), platform
and validated CLI/API capabilities. Provenance and compatibility are separate:
a matching version string cannot make a viewing-host path server-local.

Use documented read-only CLI facilities (`version --json`, `whoami --json`,
`api describe`/server metadata as needed), never a mutating capability probe.
Go owns decoding the supported release shapes and required-method schemas; unknown
or missing evidence is incompatible, not guessed success. `server.version`'s
absent API revision fields mean zero, not latest. WhichKey feature negotiation
below is separate from Rex API revision compatibility. Do not require CLI/server
release strings to be identical when the required schemas are compatible; doing
so would reintroduce user-maintained version pairing. No installed/server
minimum version has been runtime verified by this documentation task.

**Endpoint precedence, independently:** present explicit `--server` (or retained
server-owned endpoint intent) > nonempty `REX_SERVER` > native selected-CLI
`--autostart=false --timeout 2s whoami --json` discovery on macOS/Linux. Present
empty explicit intent is invalid. Validate/canonicalize a concrete absolute Unix
socket path/URI or host:port, never a mutable host label or opaque Lua `ctx.server`.
Native discovery strips ambient client/session/window/block/server selectors but
retains Rex runtime/home/authentication environment. Bound it to two seconds and
64 KiB. Capture exactly one endpoint; every subsequent call passes it explicitly
with autostart disabled. No alternate endpoints, default fallback after failure,
or manual platform socket-path reconstruction.

Read server identity/configuration through the original VM and this CLI endpoint;
require matching server `instance_id`, configuration hash/generation and exact
resource placement. `host_id` equality only proves shared files between the
servers reporting it; it is not user authorization, an endpoint, or proof that a
remote CLI process runs there. The private rendezvous in §5 establishes shared
execution-host access for public configured launch instead of asserting locality
from `--server` or a pathname alone.

**Platform bound:** the inspected upstream asset distribution covers Apple-silicon
macOS and both listed Linux architectures. It does not establish an Intel macOS
asset. Existing terminal/PATH operation there is not removed; restricted-PATH key
launch without an available licensed dependency fails before popup. Do not claim
Intel restricted-PATH support or fabricate an artifact. Expanding that matrix
needs upstream distribution evidence, not a new path heuristic.

## 5. RUNTIME-BOOTSTRAP-COMPATIBILITY

**Producer:** flat Go application owns all transport/format semantics. **Consumer:**
minimal Lua registration/retention in the original configuration VM. These are
new internal interfaces to implement, not existing Rex APIs.

### Read-only original-VM advertisement

Loading the integration module registers one `whichkey_runtime` configuration
action at load time. `menu.register_launch` associates each launch name with its
actual retained registry and completion handler in that module. Query arguments:
`{ protocol_min: 1, protocol_max: 1, action: NAME }`. Strictly bound the request
and response to 64 KiB. Names follow existing lowercase/digit/underscore syntax.

The response has: `protocol_min`, `protocol_max`, `consumer_build`, `features`,
`launch_action`, `completion_action`, `menu_generation`, `server_instance`,
`config_generation`, `config_hash`, and `layout_revision`. Require types,
nonempty bounded strings, integral ordered ranges, unique features, exact names
and all mandatory fields. Unknown mandatory semantics or disjoint ranges fail;
optional extensions require a negotiated protocol revision, not permissive
interpretation of unknown fields. Choose the highest common protocol (initially
1). The action neither reserves a ticket nor evaluates user conditions, invokes
callbacks, creates files/popups or returns capabilities/handles. It advertises
**loaded code**, not presumed features from disk/version strings.

Required features are `original-vm`, `separate-caller-target`, `single-use`,
`generation-bound`, `retained-menu`, `public-file-relay`,
`completion-receipt-v2`, and `selection-text-resource-rename` when the retained
template can contain a built-in rename. `completion-v1` alone does not imply text
support. Advertise a feature
only if the registered launch and completion paths both implement it. The Go
producer advertises the same understood range/features from a captured internal
`runtime-describe` command for key launch; no authority is produced by that
command. Consumer/producer build IDs are audit identities, not equality locks:
compatible builds may differ while preserving the same loaded custom closures.

### Two entry paths, one selected popup component

**Key/palette:** the original VM validates actual `ctx` and GUI connection, runs
only its fixed anchored helper's `runtime-describe`, and compares the actual
loaded consumer's features. For this entry path the read-only command also returns
the frozen CLI/endpoint/component receipt after discovery and metadata checks;
Lua compares its server instance/configuration with direct original-VM metadata.
It passes that receipt back on private preparation stdin, where Go revalidates
rather than rediscovers it. Thus that same selected Go image serves preparation/
controller/child, and CLI/endpoint failures precede any reservation or popup. A
replaced anchor/helper invalidates this launch.

**Public:** the running Go binary selects itself and the Rex CLI, queries the
read-only advertisement, and keeps the native route only if the implicit default
is truly absent. It must not send a caller-supplied executable path for Lua to
execute. Instead, the selected public Go process performs preparation and starts
its own selected controller. Original-VM menu data and the private capability
cross a finite **file rendezvous**, using the fixed installed helper only as a
bounded transport relay, not as the popup preparer/controller:

1. Public Go creates its existing-style 0700 invocation workspace, a live owner
   lease and precreated 0600 regular transport slots. Its descriptor contains the
   protocol/selection receipt, canonical endpoint, exact origin/GUI generation,
   workspace locator and deadline; no bearer capability. Invoke the registered
   launch action once, with descriptor on CLI stdin. Keep its activation pending
   while servicing the bounded rendezvous; no second Lua VM or script evaluation.
2. Original Lua validates actual caller/session/block and expected GUI binding,
   rechecks the advertisement's menu/config generation, then calls the anchored
   `runtime-relay` with descriptor and actual retained context on captured stdin.
   The relay checks local file ownership, no-follow regular files, one-link
   transport slots, private modes, live lease, protocol marker and exact matching
   origin/endpoint receipt. Fixed names only; no arbitrary filename/code/command
   field. It must not create/truncate an unrelated caller-named file. No valid
   local workspace/lease means unsupported cross-host launch, before reservation.
3. Only after this preflight does Lua issue the original registry ticket and
   evaluate conditions once. Relay writes the resolved Go-owned Config and the
   retained receipt into the `menu` slot. Public Go validates that data and uses
   **its own** preparation code to generate the random capability and immutable
   selected component/CLI/endpoint receipt; it writes the private `prepared` slot.
4. Relay reads `prepared`, returns it only on captured stdout to original Lua;
   Lua verifies the receipt, binds its retained ticket, then relay writes `bound`.
   The public process requires both matching `bound` and certain successful
   activation before starting the selected self-controller. The action returns
   no secret, ticket, process handle or arbitrary callback result.
5. Timeout, lease loss, generation/identity change, malformed/duplicate transition
   or uncertain activation revokes/cancels what is still unclaimed and closes the
   workspace without creating a popup. Never repeat activation, relay transition
   or a possibly successful spawn. Installed relay compatibility is checked too;
   it cannot silently run an older popup on the public caller's behalf.

Rendezvous layout is fixed: `request.json` (public writes once), `menu.json`
(relay writes once), `prepared.json` (public writes once), `bound.json` (relay
writes once), plus the FIFO `lease`. The request contains a fresh random local
marker; its digest in the action descriptor must match the local file, preventing
an accidentally coincident path on another host from being accepted. Each JSON
slot uses `{stage, receipt_hash, payload}`; the stage must equal its filename's
stage, the receipt hash must match this invocation, and payload is respectively
the launch descriptor, resolved configuration, preparation, or binding receipt.
An initially empty slot means not yet published. Writers take an exclusive Unix
file lock for their one bounded write; readers take a shared lock, so a partially
written JSON value cannot be interpreted as a state transition. Repeated read-only
waiting is bounded by the same stage/deadline and lease, not a retry of an effect.
Unexpected nonempty contents or a second publication aborts. File ownership,
permissions and identity checks apply again on every open.

The relay is a finite internal command in the existing executable, not a service,
new public CLI dependency framework or alternate Rex transport. It only handles
fixed private files; it cannot execute a nominated program. Use the same
no-follow/owned-directory discipline as existing popup transport (`os.Root` and
Unix file checks), strict stage ordering and bounded reads. The lease is an owned mode-0600 FIFO: public Go retains its write-capable
endpoint; the relay uses nonblocking reads and treats EOF as owner loss, never as
success. Do not substitute a PID-only check. A locator alone is not authority. Matching private local state establishes filesystem access;
`server.whoami` principal data is not repurposed as authorization. This is a
trusted same-account installation boundary, not isolation from malicious code
already running as that account.

### Receipts, bounds and pre-popup ordering

The private Go-owned receipt binds: selected protocol/features; popup executable
path/content/file identity and embedded asset digest; relay/layout identity when
used; frozen Rex executable identity/provenance; canonical endpoint; server
instance/config hash/generation; launch/completion names; retained menu generation;
full session/block/window placement; actual activation caller; distinct GUI target
and `connected_at`; and original expiry. The controller handoff includes this
receipt alongside existing configuration/capability fields. It never takes paths
or action implementations from selected menu text. Completion carries its receipt
binding with the existing Go-owned completion data; Lua compares with the retained
entry. Do not silently extend the existing version-1 strict completion decoder:
negotiate the bootstrap receipt extension explicitly before either side uses it.
Go's Config/Mapping/Completion definitions remain format owners. The new completion
wire version is 2, with exact fields `version`, `capability`, `receipt_hash`,
`kind`, and optional `selection`; selection keeps the existing strict
keys/ref/confirmation/text semantics. `receipt_hash` is the SHA256 of the exact
bounded receipt bytes retained at binding, not a separately reconstructed Lua
JSON serialization. No version-1 message is reinterpreted as version 2.

Ordering is: discovery/describe -> endpoint/server/origin validation -> compatibility
of producer, relay and loaded consumer -> ticket/evaluation -> preparation/bind ->
controller validation -> rename presentation validation -> popup creation. Recheck
retained identities after every yield. No popup is used as a compatibility probe.
Missing/old negotiation support is a compatibility error **before** calling an old
launch action; there is no attempt to hot-upgrade arbitrary old loaded closures.

Retain existing bounds: 64 KiB small transport/completion messages, 8 MiB complete
handoff/config transport, 64 pending Lua invocations, 30-second configured lifetime,
three-second command/preparation stages within that total, five seconds reserved
for teardown, and detached-controller command timeout 35 seconds. The rendezvous
uses those bounds, never a fresh renewable 30 seconds per step. A slow preparation
cannot consume the teardown reserve and then open a popup. `rex.exec` timeout
sends SIGTERM and later SIGKILL; it is not instantaneous revocation. Native display
keeps its existing no-idle-timeout/lifetime-pipe policy, selected only by trusted Go
routing, never by empty fields in a decoded configured request.

## 6. Dispatch, rename and failure invariants

- Preserve actual Lua `ctx` and its opaque server identity without stringifying
  or replacing it. Public activation/completion CLI callers are not the retained
  app. Explicit `--client`/`REX_CLIENT` nominates an exact full app ID, not caller
  impersonation; validate app kind, session attachment and `connected_at` again.
- Retain closures, args, guards, confirmation requirements, table tickets and
  immutable selection snapshots only in the original VM. Generation replacement,
  expiry, backwards clock, reconnect or moved origin revokes/suppresses authority.
  Never serialize a callback or re-evaluate the menu in another VM.
- Keep random capabilities on finite stdin/captured pipes or invocation-owned
  private files only. No argv/environment/action result/log diagnostics contain
  them. Bound child output; report fixed stage-specific messages, not raw child
  or user-code exceptions. Component paths/hashes are not bearer capabilities.
- Strict bounded completion decoding precedes lookup. Authenticate exact origin,
  receipt, expiry and capability before claiming; malformed/foreign data cannot
  consume an unrelated entry. Claim **before** any yielding guard/API/callback;
  after claim, failure or lost acknowledgment consumes authority permanently.
  Confirmed selection keys/ref/text must match the retained snapshot. No uncertain
  retry, reopened popup, second completion, native fallback or alternate target.
- Default `tr`/`client.tab.rename` and `wr`/`session.rename`, without callback,
  custom args or confirmation override, use the existing **WhichKey text page**.
  Preserve captured initial tab/session labels and strict nonblank control-free
  UTF-8 text (existing 4096-byte transport bound); never truncate/discard text.
  Before dispatch, validate the resource API's documented label limit as well;
  the transport limit is not a promise that Rex accepts every such label.
- Text supplies only a label. After cleanup and a final generation/connection/
  placement check, original Lua calls `session.set_window_label` with the retained
  window/session or `session.set_label` with the retained session. Moving the
  origin into another tab cannot retarget rename. Native Go uses its corresponding
  typed resource. Missing text support fails before popup, never a Rex dialog.
  A custom callback on those keys remains exactly that callback.
- Close only the exact known owned popup and remove only owned files; restore
  only a still-valid captured origin. Finish cleanup before retained dispatch,
  which may close its own origin. Unknown creation identity or uncertain close,
  focus or transport cleanup suppresses selected dispatch; no cascading cleanup.
  External moves/disconnects can race read-only checks; there is no atomic
  focus-and-execute transaction, and no successful API acknowledgment alone
  proves the intended visible resource effect.

Diagnostics distinguish: helper/layout discovery, CLI discovery, endpoint discovery,
compatibility, origin validation, private rendezvous, preparation, controller/child
startup, cleanup and dispatch. Every failure names a bounded recovery action without
exposing private child output. Never tell users to repair constructor paths or pair
binary/Lua versions as a routine runtime fix.

## 7. Downstream boundaries and normal upgrades

| Task | Owns implementation of | Must not do |
| --- | --- | --- |
| 02: Go resolver | §4, one frozen Rex CLI selection and native endpoint policy in the concrete `rex` package; required read-only schema checks. | Runtime Nix/host directory search, alternate-server probes, new client/DI architecture, or trusting terminal documentation as configuration-action provenance. |
| 03: Go bootstrap | Go-owned describe/receipt/relay definitions, public self-selection and private rendezvous, key preparation, immutable self-spawn selection, pre-popup compatibility/text checks and controller diagnostics. | Execute a nominated remote/caller path, bypass a registered custom menu, expose capabilities, add a daemon/socket transport or change public launch into a version-pairing tool. |
| 04: Lua consumer | File-backed anchor, at-load advertisement, loaded-menu generation/features, minimal relay orchestration, original tickets/closures, separate callers/app and exact retained-resource rename. | Invent a mapping schema, become the CLI resolver, execute caller-nominated code, load a second callback VM or auto-reload configuration. |
| 05: Packaging/docs | Normal relative runtime/helper/dependency assets, direct upstream dependency installation/link, reserved namespace links, help/examples/support matrix and migration guidance. | Inject absolute runtime strings, resolve live servers, wrap runtime PATH, activate consumer configuration, vendor unlicensed Rex bytes or move user mappings. |

`execution.helper`, `execution.rex_bin_dir` and their environment forwarding are
obsolete runtime-wiring fields. Target code rejects present legacy wiring with a
fixed migration diagnostic; do not silently prefer or ignore it. Existing public
`--server`, `--client`, `--action`, optional server-endpoint constructor intent,
`setup`, custom mappings/callbacks and launch/completion registration remain.
Removing obsolete entries from user-owned configuration is user work, never an
automatic migration. Normal examples require only registration/customization.

Install binary/assets/dependency together through the ordinary installer. Nix
only builds/installs those assets; Go chooses and validates runtime use. No repeated
flake refresh is a launch repair. Compatible **already-loaded** Lua remains usable
with a rebuilt public Go binary; it retains its original callbacks and negotiates
features rather than requiring matching release IDs. Compatible installed relay
changes are likewise permitted after preflight. Unsupported/missing old bridges
reject before popup and explain that the integration software must be updated and
loaded through the user's normal configuration lifecycle. Installation cannot
replace arbitrary loaded closures; no automatic reload or manual selection of
matching component versions is promised or required. A normal reload/restart that
the user independently chooses is distinct from runtime version pairing.

No managed-design mutation is required to implement these interfaces within the
accepted boundary. This document is not a new approval record or permission to
weaken unsupported-platform/security constraints.

## 8. Falsifiable scenarios (all unexecuted)

| Scenario | Required outcome / observable failure boundary |
| --- | --- |
| Public implicit default absent | Native route uses current Go component; no app required merely for display; Escape cleans only owned resources. |
| Public registered/default or explicit custom action | Current public Go prepares/controller/child; original-VM menu and callbacks retained; installed relay never substitutes its popup implementation. |
| Key/palette with empty or restricted app PATH | File-backed anchor starts helper; fixed-relative official CLI dependency works without environment/constructor repair. |
| Terminal reordered PATH | Valid terminal `REX_BIN_DIR` wins; bad explicit environment preference fails without trying a different copy. |
| Missing/relative/cwd/ambiguous PATH | Fixed discovery error; no guessed directories, endpoint changes, popup or retries. |
| Local rebuild versus installed old Lua/relay | Compatible feature receipt works across build IDs; missing/unknown consumer or relay fails before popup, not after text selection. |
| HM symlink update during invocation | Current frozen selection unchanged; replacement invalidates rather than redirects; next launch can resolve new normal assets. |
| Relative/synthetic/shadowed/duplicate module | Only declared standard-loader/cwd case works; other contexts fail registration/preflight, not suffix guessing or silent module replacement. |
| Remote viewer or public caller without shared workspace | Viewing-host paths cannot authorize execution or arbitrary file writes; reject before reservation/popup; key/server-host public launch remains separate. |
| Malformed endpoint/server/origin/receipt | Reject before mutation; opaque ctx is never an endpoint or substitute caller. |
| Default tab/session rename | WhichKey text reaches only exact retained resource API after cleanup; no dialog/discarded text or false success. |
| Custom callback on rename keys | Same original closure/args/guards/confirmation; no synthesized default rename. |
| Missing text feature or stale completion handler | Compatibility failure before popup; completion version 1 alone is insufficient evidence. |
| Setup/reload, reconnect, expiry/backwards clock or moved tab | Reject/revoke at generation/placement checks; no renewed lifetime or target substitution. |
| Cancellation, forged/duplicate completion, callback failure or lost reply | No unrelated reservation consumed; authenticated claim is single-use; no uncertain retry or cancellation masking a possible dispatch. |
| Unknown creation ID, failed owned close/focus/file removal or killed caller | Bounded owned cleanup only; lease loss suppresses startup/dispatch; no cascading close or reopening. |
| Intel macOS without CLI source | Explicit unsupported restricted-PATH dependency diagnostic; no invented download or app-bundle scan. |

Only metadata/source/document inspection and `git diff --check` are admitted here.
No tests, TDD/race runners, builds, Lua/API/menu/rename probes, remote trials,
benchmarks, configuration reloads, Nix evaluation/activation or deployment were
performed. Existing tests, Taskfile, lint configuration and hooks are unchanged.
The official Lua manual fetch timed out; standard-library availability is directly
supported by Rex's complete fetched introduction, while the diagnostic's manual
references remain references, not a new successful network verification.

## References

Local baseline: [launcher](../launch.go), [CLI](../cmd/main.go),
[resolver](../rex/exec.go), [endpoint](../rex/endpoint.go),
[controller/transport](../popup.go), [rename preparation](../native_text_entry.go),
[format owner](../mapping.go), [original-VM launcher](../lua/rex_whichkey.lua),
[dispatch](../lua/rex_whichkey/dispatch.lua), [package](../nix/package.nix),
[HM namespace](../nix/home-manager.nix), [flake](../flake.nix).
Existing [launch](launch-contract.md), [mapping](mapping-format.md) and
[CLI/GUI](cli-gui-delegation.md) documents describe baseline context; this document
supersedes only their runtime-wiring guidance when implemented, not their safety
invariants or historical verification status.

[lua]: https://www.superlogical.com/rex/docs/reference/lua.md
[modules]: https://www.superlogical.com/rex/docs/customize/config#splitting-configuration
[path]: https://www.superlogical.com/rex/docs/automate/cli#path
[nix-rex]: https://www.superlogical.com/rex/docs/linux/installation#nix-and-nixos

Further supported API/CLI references: [actions](https://www.superlogical.com/rex/docs/customize/actions),
[exec](https://www.superlogical.com/rex/docs/reference/lua#rexexecargv-opts),
[exec.start](https://www.superlogical.com/rex/docs/reference/lua#rexexecstartargv-opts),
[server.status](https://www.superlogical.com/rex/docs/reference/api#serverstatus),
[server.version](https://www.superlogical.com/rex/docs/reference/api#serverversion),
[version JSON](https://www.superlogical.com/rex/docs/reference/cli#rex-version),
[whoami JSON](https://www.superlogical.com/rex/docs/reference/cli#rex-whoami),
[global flags](https://www.superlogical.com/rex/docs/reference/cli#global-options),
[server start environment](https://www.superlogical.com/rex/docs/reference/cli#rex-server-start),
[Lua 5.1 loaders](https://www.lua.org/manual/5.1/manual.html#pdf-package.loaders),
[Lua 5.1 debug metadata](https://www.lua.org/manual/5.1/manual.html#pdf-debug.getinfo).
