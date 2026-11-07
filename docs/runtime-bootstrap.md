# Ordinary package and Go-owned runtime discovery

## Current responsibility boundary

Go owns runtime policy; Nix installs ordinary assets and Lua starts
`rex-whichkey` by name. There is no anchor, helper-path constructor, wrapper,
ldflags runtime path, runtime resolver in Nix, custom Lua action registration or
configuration compiler. [Launch](launch-contract.md) and
[data configuration](mapping-format.md) describe the active contracts.

## Package layout

```text
bin/rex-whichkey                         # real Go image, not a Lua-tree symlink
bin/rex                                # optional symlink to licensed bin/rex
share/rex-whichkey/lua/rex_whichkey.lua   # only production Lua module
share/rex-whichkey/examples/             # keybinding + version-1 JSON
share/doc/rex-whichkey/README.md
share/doc/rex-whichkey/docs/
share/licenses/rex-whichkey/LICENSE
```

[nix/package.nix](../nix/package.nix) retains the upstream source/artifact pins,
Go vendor hash, Linux loader/PAM rpath, terminfo/SSH assets and Darwin signature
preservation. It keeps `doCheck=false` and `doInstallCheck=false`.
The licensed CLI is available from the pinned source for x86_64/aarch64 Linux and
aarch64 Darwin. **No Intel macOS CLI artifact is configured.** A legitimate
`rex-cli` override can supply `bin/rex`; otherwise provide a legitimate `rex`
on the execution-host PATH, with or without a JSON config. The optional
symlink exposes the same already pinned dependency, not a new CLI or wrapper.
A source-only Go build installs neither that CLI nor the Lua/assets; expose
the named Go binary and separately supplied CLI through ordinary PATH.

Fixtures, obsolete Lua modules/defaults/runtime directories, workflow state and
experiments are excluded from the package source/install surface. No binary lives
under Lua, and no runtime/deps/generated-helper layout is installed.
Home Manager links only the launch module with conflict safety. It does not
generate menu data, init.lua, bindings or an endpoint policy; see the
[packaging source guide](../tests/packaging/home-manager/README.md).

## CLI executable discovery

[rex/runtime.go](../rex/runtime.go) uses Go's ordinary `exec.LookPath("rex")`
on the execution host. The first executable match wins; a relative directory
result is made absolute once so later working-directory changes cannot retarget
it. Standard Go lookup safety (including rejection of an implicit current-directory
executable) applies. Absence produces a clear PATH/install error, without exposing
child diagnostics. There is no application-wide PATH-directory/uniqueness policy.

The parent, concrete clients, inventory readers and private PTY child reuse that
selected path. No `REX_BIN_DIR`, internal executable override, private libexec
search, public executable flag, viewing-host path or store/repository scan is
supported. There is no repeated executable discovery, binary/link/content receipt,
startup version/API/schema negotiation, method admission or before/after
`server.status`/config-generation wrapper. Trust now follows ordinary PATH/OS
policy; replacement detection by application hashing is intentionally retired.

All Rex subprocesses still use bounded, cancellable argument-array execution via
[rex/exec.go](../rex/exec.go), explicit endpoint/no-autostart where applicable,
strict actual result decoding and sanitized failure/uncertainty. Built-in action
metadata is fetched after selection, not as a startup gate. No execution error
permits another binary, endpoint or Lua route, or an uncertain-operation retry.

The Lua start uses the running server's PATH for **rex-whichkey**, and Go uses
its inherited PATH for **rex**. Package `bin/rex` exposure helps only when that
bin directory is on the launching process's PATH. Installing a package or changing
a terminal profile does not update a running server's inherited environment.

## Endpoint discovery is independent

[rex/endpoint.go](../rex/endpoint.go) applies:

1. Present public `--server ENDPOINT`.
2. Nonempty execution-host `REX_SERVER`.
3. On Linux/macOS, bounded native `rex --autostart=false --timeout 2s whoami --json`
   using the already selected CLI.

A present empty explicit override is an error. Addresses are concrete absolute
Unix socket paths/URIs or host:nonzero-numeric-port (including bracketed IPv6),
not mutable host labels, opaque `ctx.server` or viewing-machine paths.
Native discovery removes inherited Rex target selectors while preserving
HOME/XDG/runtime/home/authentication settings; Rex owns its platform algorithm.
No socket guess, filesystem scan, directory creation, autostart or alternate
server retry occurs. The selected endpoint is frozen and explicit thereafter.

## Invocation-private PTY transport

A public Go parent owns one interactive child, not a daemon/controller framework.
It opens its running image once and copies from that retained handle into an
invocation-owned 0700 directory. Ordinary bounded copy/write/close finishes
before the Rex terminal starts; the image copy has no durability Sync, hashing
or installation receipt. It passes resolved menu data, the selected CLI path,
an independent canonical endpoint, invocation ID, inline-form enablement and a
parent lifetime pipe.
Those are private implementation files, not install-time helpers, public config
or a callback channel.
The child returns keys/confirmation/text and minimal rename/provider publication
bindings only; the parent retains resource/action authority and dispatches once
after exact cleanup. There is no idle expiry.

The [private incremental contract](launch-contract.md#private-incremental-rename-contract-startup-ui-updates)
uses twelve fixed single-publication slots: two form demand/update pairs and four
provider demand/update pairs. Strict bounded JSON, invocation/kind/load identity,
canonical demand prefixes and nonblocking flock preserve complete publications.
The existing UI loop writes the static root first, then requests visible enabled
sessions in background; hidden sessions, windows/panes/hosts wait for explicit
enabled prefix entry. Labels wait for form entry. Absent/disabled providers do
nothing. One in-flight/terminal result per kind includes errors and is reused by
all matching markers without reopen retry.

One small invocation scheduler runs at most two metadata jobs with two result
slots. Workers return detached values; the controller alone retains complete
mapping/target batches before publishing; only the UI loop applies/redraws.
Updates preserve active prefix/page/confirmation/form edits/cursor and partial
terminal decoding. Loading/unavailable rows cannot execute; forms remain
editable/cancellable but cannot submit until ready. There is no buffered Enter.
Cancellation/selection/error/parent loss stops and joins jobs/controller and the
child lifetime watcher before result/cleanup permits effects. Errors propagate.
No background terminal writer, extra service or persistent cache exists.

`rex.SessionSnapshot` retains one validated exact origin session-view. Detached
`Session()`, `Window(fullID)`, `Block(fullID)`, `Windows()` and `Blocks()` serve
placement and same-read window/pane/rename presentation. This removes the prior
presentation-only `Client.Session + Session.Windows/Blocks` inspect pairs.
The sessions list and known hosts each have their own once-per-invocation result.
Endpoint/session/placement matching is captured, never retargeted from later
focus. Current `Inspect`, native origin, app owner/generation, destination,
pre-create geometry consistency, liveness and cleanup checks intentionally remain
fresh reads. Selected built-in action metadata remains off startup.

The [fixed cleanup inventory](launch-contract.md#fixed-owned-cleanup-inventory)
names all files/resources. Root-scoped fixed-name removal, no-follow/owner/mode
checks, bounded image copying, strict results/locks/result Sync and exact source
handle retirement remain. The child checks its fixed private image slot, not the
parent installation's content or links. No guarantee after SIGKILL,
hostile same-account changes or server loss, atomic resource transaction or
uncertain-effect retry is implied.

## Evidence and history

Phase16 delivery is complete with actual final public menu/navigation/owned
rename/dismissal/cleanup on Darwin arm64. Phase18 tasks1/2 add reviewed/integrated
public runs, including pending rename edits preserved on ready. The current
async source has format/whitespace/normal-build/vet gates; final combined source
review and comparable public timing/effect/cancel/cleanup remain coordinator
gates. [Acceptance](acceptance.md) records exact artifacts and limitations.
No installed Nix/profile, physical-keybinding or cross-platform result is implied.

Phase15's original-VM advertisement, protocol/layout anchor, completion receipts,
finite startup barriers and public relay are **historical**, not startup
prerequisites. The earlier document remains in Git at
`cd5cbce818e07c264738ddcbc591ad49fdf0e9c8:docs/runtime-bootstrap.md`.
Old Lua source is preserved solely as test fixtures. Hooks/callbacks/functions
are deferred, with no new schema or compatibility stubs. No user configuration,
pin, dependency, profile, running server or legacy task history is migrated.
