# Execution-host and remote-context limits

## What runs where

Install/run `rex-whichkey` and a compatible licensed `rex` CLI on the
**execution/server host** in the originating Rex context. A graphical Mac viewing
a Linux session supplies neither Linux binaries nor configuration/socket paths.
No remote deployment, attach, SSH, Wake-on-LAN, binary copying or connection
management is implemented.

Direct terminal and minimal keybinding launch use the same ordinary Go entrypoint.
The binding starts the installed binary via Rex's execution-host PATH and forwards
exact session/block/client context IDs, explicitly clearing missing inherited IDs.
It does not query metadata, stringify `ctx.server`, choose viewing-host paths or
load executable menu data. There is no original-VM/private-file relay dependency.

Go owns an invocation-local PTY child and private files on its execution host.
A network endpoint or matching pathname alone does not establish that host's
runtime/files/origin authority. Use the actual originating execution-host terminal
or its keybinding; cross-host public invocation is not certified.

## Frozen CLI and endpoint

See [discovery](runtime-bootstrap.md) for the exact precedence:
present valid `REX_BIN_DIR` > applicable ordinary package
`libexec/rex-whichkey/rex` > unambiguous absolute PATH, after trusted internal
selection where applicable. Missing optional dependency permits PATH;
broken applicable entries or invalid higher intent do not. No runtime Nix
resolver, anchor, helper path, shell/store/home scan or alternate binary retry.

Endpoint selection is present `--server` > nonempty `REX_SERVER` >
bounded no-autostart native Linux/macOS Rex `whoami`. Explicit empty/malformed
intent fails. Accepted values are concrete Unix paths/URIs or numeric-port
host:port, not mutable host labels or opaque Lua server objects.
Rex, not application path guesses, owns platform runtime/home/XDG discovery.
The CLI/image/link/content and endpoint/server contracts remain frozen through
child wait, cleanup and dispatch.

The pinned package offers official licensed CLI artifacts for x86_64/aarch64
Linux and aarch64 macOS, not Intel macOS. Standalone PATH remains supported with
JSON/default menus; a legitimate `rex-cli` override can supply the optional
dependency. No platform output declaration or source parse proves runtime support,
license permission or a deployed server environment.

## Origin and app identity

Full session/block IDs are mandatory. Go captures exact Window/Layer placement,
rechecks it, and refuses moved/missing/ambiguous origins. GUI intent is independent:
explicit full `--client`/present `REX_CLIENT` stays strict; otherwise app actions
retain the pane's validated live app owner and generation. No focused/first/sole
app or session fallback exists. Resource-only display/rename/close need no implicit
GUI app. App selectors never bypass Rex authentication or Remote Control.
[CLI/GUI delegation](cli-gui-delegation.md) gives the details.

## Lifetime, cleanup and effects

The menu has no idle expiry or original Lua action lifetime. Its Go parent watches
resource liveness and a private pipe ties the PTY to that parent. Tab switches alone
do not cancel it. RPC checks/dispatch and cancellation-independent cleanup remain
bounded. After terminal restoration, exact owned popup/origin-focus/private-file/
pipe/source-handle cleanup must finish before any selected operation.

The parent retains target IDs/actions; a key or rename text cannot replace them.
Session/window destinations and pane placement are revalidated. New Tab creates
one source-session Window and activates its returned ID; a failed activation
preserves it, without retry/rollback. Inline session/tab rename uses exact resource
methods, not a native dialog. Unknown creation identity or uncertain cleanup
suppresses effects and may leave resources. Never close a whole user window/session,
search labels or recursively delete unrelated paths to compensate.
See the [fixed owned inventory](launch-contract.md#fixed-owned-cleanup-inventory).

Checks and mutations are not atomic against external changes. SIGKILL/server loss,
hostile same-account edits and lost acknowledgments are not guaranteed recoverable.
No reconnect retargeting, renewed ticket, callback replay or uncertain retry exists.

## Data, host inventory and migration

Go reads one version-1 JSON file after built-ins; no Lua execution, function labels/
conditions, custom Lua action names, callbacks or hooks are supported.
User files/init.lua/bindings remain user-owned; no automatic Lua conversion,
silent unsupported-data loss, reload/restart or activation is performed.
[Examples](../examples/README.md) describe manual migration.

Only an enabled `known-hosts` provider asks Rex for execution-host/home inventory,
once per invocation. Neither the viewer nor `--server` selects that inventory.
Rex may synchronize/contact its home; listing is not a no-network or reachability
guarantee. Rows/warnings are non-executable, not connect/SSH commands.
There is no default hosts row.

## Evidence and historical boundary

Task2's normal public menu/navigation/dismissal/exact cleanup was observed by the
coordinator on Darwin arm64 at `7bdbda2`; it is not a remote/Linux trial or proof
of selected action effects, JSON probes, physical keys or Nix/HM activation.
Final combined source/runtime validation remains a coordinator gate.
No general minimum tested Rex build, all-platform compatibility or post-crash
cleanup guarantee is claimed. [Acceptance](acceptance.md) separates evidence.

The Phase15 bridge/anchor/receipt/completion remote guidance is historical at
`cd5cbce818e07c264738ddcbc591ad49fdf0e9c8:docs/remote-compatibility.md`,
not an active dependency or permission to deploy it. History is preserved;
Phase15 and unrelated phases are not completed by this cutover.
