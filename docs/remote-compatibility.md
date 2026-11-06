# Remote-context compatibility

## Scope and evidence

This describes source intent for REQ-003/005/011/013/016/017, not a remote acceptance
result. The reviewed/compiled runtime, resolver and original-VM consumer do not
establish observed popup, callback, rename, disconnect or platform behavior. All
installation/upgrade/endpoint/callback/rename cases below are **UNEXECUTED source
scenarios**. This packaging/docs cutover performed only permitted static checks;
no builds, tests, runtime probes, remote trials, reloads, activation or deployment.

Authoritative contracts: [runtime bootstrap](runtime-bootstrap.md) and
[Go-owned mapping/completion formats](mapping-format.md). No parallel wire schema,
remote host registry, daemon, socket client, callback VM or bootstrap service is
introduced.

## Execution host is not the viewing host

The execution host must provide the selected processes, normal module assets,
licensed Rex CLI and shared invocation-private files accessible to the original
server configuration VM. A graphical Mac viewing a Linux server supplies none of
those Linux paths. No client-home/app-bundle/store/repository guesses, binary
copying to a remote host, automatic attach, WOL or SSH workflow is implemented.

- **Server-host terminal:** public Go selects itself for preparation/controller/
  PTY, freezes its Rex CLI and endpoint, validates origin and uses the original VM
  only for configured customization. An installed helper relays finite private
  file messages; it never substitutes its own popup component.
- **Key/palette:** the original VM loads the normal file-backed module anchor and
  starts only its fixed sibling helper. Go validates layout revision 1 and
  `runtime/deps/rex`, so restricted server PATH needs no constructor paths or
  presumed configuration-action `REX_BIN_DIR` injection.
- **Cross-host public caller:** a network endpoint and matching pathname do not
  establish shared execution files. Configured launch requires authenticated
  owned private rendezvous state and live lease; no shared state rejects before
  reservation/evaluation/popup. Run in the execution-host Rex terminal or use its
  key action instead. Native public launch also requires the originating
  execution-host terminal endpoint, not a conflicting remote override.

The complete installed Lua namespace must use the standard file loader. Absolute
anchors and standard-loader relative anchors with unchanged server cwd are
supported; synthetic/preloaded/custom-loader labels, mixed or shadowed trees and
unsupported layouts reject. Home Manager's namespace symlinks are supported.
An installation pinned directly to immutable old modules remains pinned; code
never scans for a newer replacement. Installer/user-owned components must have
no untrusted writers. This is not hostile-same-account isolation or a portable
atomic execute-by-file-descriptor guarantee.

The package includes the normal anchor/helper and official pinned proprietary CLI
on x86_64/aarch64 Linux and aarch64 macOS, preserving upstream loader/libraries/
signature/assets. Upstream has no Intel macOS artifact in that source: standalone
terminal/PATH use is retained, but Intel restricted-PATH key bootstrap is not
claimed. Applicable license/unfree-policy constraints remain; no redistribution
right or minimum tested Rex release follows from the pin. See [installation](../README.md#install-on-the-rex-server-host).

## One frozen CLI and endpoint

Go's concrete `rex` package owns selection: trusted frozen/internal input,
nonempty execution-host `REX_BIN_DIR`, applicable validated fixed-relative normal
dependency, otherwise unambiguous absolute PATH. Terminal-child documentation does
not promise that the server configuration VM receives `REX_BIN_DIR`. Invalid
explicit preferences, broken layouts, relative/cwd/empty PATH entries and distinct
executable candidates fail, not fall through or run probes until one works.

Before endpoint operations Go freezes canonical file/content/link identities,
platform and read-only CLI/server API-schema evidence. Controller, PTY, resource,
inventory and completion calls reuse that selection; no post-detach environment
rediscovery or PATH repair occurs. Replaced/removed components invalidate the
invocation, never select a second copy. `version --json` is retained as bounded
metadata, not an invented release-shape/API-version promise; required consumed
method schemas establish compatibility. CLI/server release strings need not match.

Endpoint precedence is present public `--server` or key constructor intent >
nonempty `REX_SERVER` > bounded no-autostart native Rex Linux/macOS discovery.
Rex owns socket/runtime/home/XDG rules. Empty explicit intent or malformed values
fail without alternate endpoints, scans, directory creation or server autostart.
Accepted addresses are concrete absolute Unix paths/URIs or numeric nonzero-port
`host:port` (including bracketed IPv6); opaque Lua `ctx.server` is not an address.
Generic typed resources retain exact endpoint/session IDs rather than labels or
later environment values.

`M.new(defaults, endpoint)` is optional immutable intent outside setup options.
Key describe applies its precedence. Public Go already selected its endpoint:
a present constructor string must canonically agree, after private/origin
authentication, and stays byte-identical across preflight/menu/bound. Absence does
not request rediscovery; empty/malformed/conflicting or changed intent rejects
before authority/popup, never silently ignored or retargeted. The shared socket
`/run/rex/server.sock` is an explicit deployment example, not a platform default.
This preserves explicit routing while refining the old mandatory-explicit setup
REXWHICHKEY-20261101-015342. It grants no execution-host authority.

## Original VM, actual callers and compatibility

Read-only `whichkey_runtime` advertises actual loaded menu/config/server generation
and supported features before ticket/evaluation. Protocol/layout 1 and configured
completion 2 are Go-owned internal interfaces, not undocumented Rex APIs. Required
text/resource-rename support is negotiated before popup. Compatible public,
helper and loaded-Lua build IDs may differ. Missing/old bridges reject before
calling an old launch action; no fallback, callback serialization or hot reload.

Public GUI nomination remains distinct from each real CLI action caller. Explicit
`--client`/present `REX_CLIENT` is strict; otherwise public Go binds the pane's exact
GUI owner, not a sole/first/focused app. Key launch uses its actual app. Original
Lua freezes actual opaque server/session/block/caller and independently checks
exact nominated app membership/kind/connection and unique window/layer placement.
The public preflight authenticates private request digest/manifest/receipt/live
lease and all origin fields before returning a reply or reserving a ticket.
Recheck after every yield; no refreshed connection, app or target replaces it.
See [CLI/GUI delegation](cli-gui-delegation.md), accepted correction
REXWHICHKEY-20261104-172117.

Closures, args, guards, confirmation, immutable mapping/resource snapshots and
table tickets remain in that original VM. Go-generated random capability and
private handoff stay on captured stdin/pipes or 0700/0600 owned files, never argv,
environment, logs or action results. Go receipt and relay strings are opaque and
byte-exact across Lua 5.1; large timestamps/inode/device numbers must not be
reencoded through doubles. Fresh completion authenticates the final receipt/hash,
real caller/server/session/block, expiry and retained state before a single-use
claim **before** yielding work. Claimed errors consume authority; malformed or
foreign data cannot consume unrelated entries. Callback GUI routing context is
not fabricated ambient action authority. Remote Control/authentication remain
Rex's policies; no selector/capability bypasses them.

## Lifetime, ownership, cleanup and dispatch

Configured work retains its original 30-second lifetime, 64 pending entries and
five-second teardown reserve. Stages are bounded; no yield, redraw, receipt,
relay or completion renews expiry. Generation replacement, backwards time,
reconnect/detach or moved origin rejects/revokes. Native interaction separately
uses its parent lifetime pipe and no idle deadline.

Both entry routes require finite private controller input, accepted acknowledgment
plus EOF and positive proceed plus EOF before popup effects. Startup handles are
retired before effects. Acceptance is not visible display or callback completion.
Before start only authenticated owned cleanup is permitted. Once `Start` may have
succeeded, ownership remains with the possible child: no uncertain retry, kill or
workspace removal. Authentic complete expired pre-child input can authorize only
its fixed owned cleanup after receipt/copy/self-image authentication; malformed/
partial/foreign input cannot authorize nominated deletion and may leave artifacts.
Expiry revokes registry authority, not proof of artifact removal.

Only exact returned owned popup resources and fixed owned files/handles are
cleaned; never cascade to a layer/window/session or search by label. Restore only
a still-valid captured origin. Cleanup and source-handle retirement finish before
selected dispatch, including actions that close their origin. Unknown creation
identity or uncertain close/focus/file removal suppresses effects and may leave
resources. Never retry an uncertain close or selected mutation.

Original Lua rechecks receipt, generation, original expiry, GUI connection,
placement, retained keys/ref/confirmation/text and guards. Dotted client actions
use explicit retained-target `client.perform`; named configuration actions use
original-VM `rex.invoke`; callbacks receive retained args and logical origin.
Default `tr`/`wr` uses WhichKey text and exact retained window/session resource
rename once after cleanup, with nonblank control-free UTF-8, 4096-byte and
128-character limits. Missing text support never opens a native dialog, drops
the name or reports false success. Custom callbacks/args/confirmation retain
normal semantics. See [rename](tab-rename-input.md) and [session UI](session-owned-ui.md).

Inspection/focus/mutation are separate APIs, not atomic against external changes.
Generic resources may inspect a block's current placement within its session;
invocation-authorized dispatch requires the frozen placement and rejects a move.
No cross-session search, reconnect retargeting, uncertain completion retry or
competing cancellation compensates for lost acknowledgment.

## Inventory, installation and residual limitations

Enabled `known-hosts` markers alone acquire Rex's execution-host/home inventory,
once per invocation through the frozen CLI and endpoint receipt. Neither the
viewing app nor `--server` chooses that inventory. Rex may contact home/forward
held changes itself; listing is not a reachability or no-network claim. Generated
rows are informational/non-executable, with explicit warnings on failure, not
remote attach or a successful empty list. No marker means no inventory lookup.

User `init.lua`, custom menu modules and keybindings remain user-owned. Whole-value
spec/icons replacement, defaults < import < inline precedence, transactional setup
and callback semantics follow the [mapping standard](mapping-format.md). Normal
installation is not runtime discovery: no Nix source substitution, runtime path
injection, wrappers or mandatory consumer settings. Present obsolete
`execution.helper`/`execution.rex_bin_dir` fields fail clearly; users remove them.
Install updates normally and, if needed, load the integration through the user's
normal configuration lifecycle. No auto reload/restart or manual version pairing.

No minimum runtime-tested Rex version, remote deployment certification, hostile
same-account protection, atomic resource transaction or guaranteed cleanup after
SIGKILL/server loss is claimed. Compatible unequal builds, source-only public use,
restricted-PATH keys, HM symlink upgrade, explicit endpoints, custom callbacks,
rename, moved origins, disconnect/expiry/replay and uncertain starts/effects are
source scenarios only. Source owners are `runtime_bootstrap.go`, `launch.go`,
`popup.go`, `rex/runtime.go`, `rex/endpoint.go` and the normal Lua integration;
packaging/docs do not verify or redefine those protocols.
