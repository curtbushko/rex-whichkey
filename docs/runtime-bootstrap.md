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
libexec/rex-whichkey/rex                 # optional symlink to licensed bin/rex
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
`rex-cli` override can supply `bin/rex`; otherwise ordinary standalone PATH
discovery remains available there too, with or without a JSON config.

Fixtures, obsolete Lua modules/defaults/runtime directories, workflow state and
experiments are excluded from the package source/install surface. No binary lives
under Lua, and no runtime/deps/generated-helper layout is installed.
Home Manager links only the launch module with conflict safety. It does not
generate menu data, init.lua, bindings or an endpoint policy; see the
[packaging source guide](../tests/packaging/home-manager/README.md).

## CLI executable discovery

[rex/runtime.go](../rex/runtime.go) freezes one execution-host CLI:

1. Trusted internal frozen selection (the PTY child reuses it, never rediscovers).
2. Valid internal executable intent where used by the concrete client API;
   there is **no public executable/path flag**.
3. Present `REX_BIN_DIR`, an absolute directory containing `rex`.
   Present-empty is invalid, not absent.
4. Applicable package dependency at `../libexec/rex-whichkey/rex`, relative to
   the **canonical public executable** when its parent directory is named `bin`.
5. Unambiguous execution-host PATH candidate.

Absent optional package directories/leaf allow standalone PATH. Dangling links,
unreadable/malformed applicable entries or invalid higher-priority intent fail
without trying lower candidates. PATH must consist solely of nonempty absolute
directories; empty/relative/cwd entries reject even after an earlier candidate.
Multiple links to the same device/inode are allowed, distinct candidates reject.
No shell search, viewing-host path, store scan or repository inference occurs.

Go freezes canonical file/link/content/platform identity, bounded read-only
`version --json` and the consumed API metadata contracts. Release-string equality
is not required and no invented minimum Rex version is asserted. Replacement,
changed links or incompatible metadata invalidate this invocation; no fallback
to another binary, endpoint or Lua route is attempted. All Rex subprocesses use
the shared argument-array [rex/exec.go](../rex/exec.go) boundary.

The Lua start itself uses the server process's ordinary PATH for **rex-whichkey**.
The packaged optional CLI dependency helps Go find **rex**, not the launch binary.
No claim is made that a running Rex server inherited a new profile or terminal env.

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
It copies its verified image into an invocation-owned 0700 directory, passes
resolved menu data, a frozen CLI/image receipt, rename presentation and a parent
lifetime pipe, then creates the Rex terminal. Those are private implementation
files, not install-time helpers, public config or a callback channel.
The child returns keys/confirmation/text only; the parent retains resource/action
authority and dispatches once after exact cleanup. There is no idle expiry.

The [fixed cleanup inventory](launch-contract.md#fixed-owned-cleanup-inventory)
names all files/resources. Root-scoped fixed-name removal, no-follow/owner/mode
checks and image/link/content validation remain. No guarantee after SIGKILL,
hostile same-account changes or server loss, atomic resource transaction or
uncertain-effect retry is implied.

## Evidence and history

Task2's integrated Go source was actually observed through normal public invocation
on Darwin arm64 at `7bdbda2`; [acceptance](acceptance.md) identifies that limited
menu/navigation/dismissal/cleanup result. This package/launcher revision is source
work with parse/format/whitespace/Go compilation gates, not an installed Nix run.
Final combined reviewed runtime remains a coordinator gate.

Phase15's original-VM advertisement, protocol/layout anchor, completion receipts,
finite startup barriers and public relay are **historical**, not startup
prerequisites. The earlier document remains in Git at
`cd5cbce818e07c264738ddcbc591ad49fdf0e9c8:docs/runtime-bootstrap.md`.
Old Lua source is preserved solely as test fixtures. Hooks/callbacks/functions
are deferred, with no new schema or compatibility stubs. No user configuration,
pin, dependency, profile, running server or legacy task history is migrated.
