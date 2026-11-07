# Ordinary package, source routing and Go runtime discovery

## Responsibility boundary

Shared [`lua/rex_whichkey.lua`](../lua/rex_whichkey.lua) owns exact-origin native
float creation/fractional sizing and starts fixed `rex-whichkey direct`, shell=none.
Go owns JSON/defaults/menu/providers/palette/own-PTY interaction and CLI effects.
There is no Go launching controller, copied image, private transport, helper-path
constructor, runtime Nix resolver, anchor, wrapper or generated configuration.
See the [launch contract](launch-contract.md) for public options and lifetime.

## Source run without installation

[Taskfile.yml](../Taskfile.yml) is authoritative. `task build` builds `./cmd` into
this root's `bin/GOOS-GOARCH/rex-whichkey`; `task run` checks explicit originating
`REX_SESSION/REX_BLOCK`, builds current source and executes:

```text
rex [-C ORIGINAL_REX_CLIENT] do ABSOLUTE_ROOT/lua/rex_whichkey.lua \
  -s ORIGINAL_REX_SESSION -b ORIGINAL_REX_BLOCK --action rex_whichkey_launch \
  [width=FRACTION height=FRACTION config=PATH ascii=BOOL no_color=BOOL] \
  development_directory=ABSOLUTE_ROOT/bin/GOOS-GOARCH
```

Origin IDs are required, never inferred from current focus outside Rex. A present
original client is forwarded, including strict rejection of empty/invalid/non-app
intent; popup environment is not the origin. No endpoint is inferred from a Lua
server object. Use the actual execution-host endpoint and original context,
not this worker's IDs or a remote viewer's paths.

On the current PATH, `make` is a tiny script execing Task: `make run -- width=0.8`.
GNU make uses [Makefile](../Makefile): `make run ARGS='width=0.8'`. Both build before
launch and preserve all Task quality definitions. See [README](../README.md) for
full commands. Source routing supplies the trusted absolute target directory as
data; Lua neither computes `go env`, probes source files nor accepts executable
code/program names. Only this terminal's normal PATH is prepended. No installed
stale module/binary, user init registration or reload is needed.

`make run`/`rex do` completion is the **launch command's** exit, not Go's eventual
menu status. Actual display/action/terminal exit and exact host deletion need
separate observation. No creation ACK, compile or help result proves them.

## Package and Home Manager assets

```text
bin/rex-whichkey                         # real Go executable
bin/rex                                # optional licensed CLI symlink
share/rex-whichkey/lua/rex_whichkey.lua   # same require module / script
share/rex-whichkey/examples/             # keybinding and version-1 JSON
share/doc/rex-whichkey/README.md
share/doc/rex-whichkey/docs/
share/licenses/rex-whichkey/LICENSE
```

[nix/package.nix](../nix/package.nix) retains upstream pins/vendor hash,
Linux loader/PAM/terminfo/SSH assets, Darwin signature preservation and
`doCheck=false`/`doInstallCheck=false`. Official CLI artifacts cover
x86_64/aarch64 Linux and aarch64 Darwin, not Intel macOS. A legitimate `rex-cli`
override or execution-host PATH CLI can supply `rex`; licensing still applies.
Source-only Go build supplies neither that dependency nor installed Lua/assets.
Fixtures/workflow/experiments are not installed. There is no binary under Lua.

[nix/home-manager.nix](../nix/home-manager.nix) already links only
`rex_whichkey.lua` beside user init.lua without forced conflict replacement and
contributes the package. It does not own init.lua, bindings, JSON, a service,
endpoint, full module directory or configuration lifecycle. No activation,
profile mutation, reload/restart or takeover is part of this source change.

Installed bindings and script actions use the same public implementation and
**omit `development_directory`**. The server must already find both program
names on PATH. Installing/updating a package or terminal profile does not update
a running process environment. Prefer explicit launch `config`/presentation data
rather than assuming caller-only exports reach the new terminal.

## CLI executable discovery

[`rex/runtime.go`](../rex/runtime.go) uses ordinary `exec.LookPath("rex")` once
on the execution host; relative-directory results become absolute. Standard Go
current-directory lookup safety applies. The first executable match wins; all
concrete clients/inventory readers reuse it for the invocation. No `REX_BIN_DIR`,
executable override, private libexec, store/home/repository scan, link/hash receipt,
startup version/API/schema check or `server.status` wrapper is supported.

Rex calls remain bounded cancellable argument-array processes with explicit
endpoint/no-autostart where applicable, strict actual-result decoding and sanitized
failure/uncertainty. Selected built-in action metadata is checked after selection,
not as a startup gate. Failure never permits another binary/server or retry of an
uncertain operation. There is no private PTY child to rediscover or reuse a path.

## Independent endpoint discovery

[`rex/endpoint.go`](../rex/endpoint.go) applies:

1. Present explicit Go direct `--server ENDPOINT`.
2. Nonempty rendering execution-host `REX_SERVER`.
3. On Linux/macOS, bounded native `rex --autostart=false --timeout 2s whoami --json`
   using the selected CLI.

The Lua action does not expose a `server` data option; source `rex do` connects
using the CLI's independent endpoint policy. The terminal then uses its own
execution environment and Go's policy above. Do not substitute `ctx.server`,
viewing-host paths, focused resources or a client ID for endpoint authority.
A caller-only new environment export is not proof of the server terminal's env.

An explicitly empty override fails. Go accepts concrete absolute Unix socket
paths/URIs or numeric nonzero-port host:port (including bracketed IPv6), not mutable
host labels. Native discovery removes inherited target selectors but preserves
HOME/XDG/runtime/home/auth settings. Rex owns platform discovery. There is no
socket guess/scan, directory creation, autostart or alternate-server retry; the
resolved endpoint remains explicit for the invocation.

## In-process lifetime and evidence

One Go UI owner retains complete target/action publications; at most two detached
metadata jobs execute. Window/pane/rename presentation reuses a matching origin
snapshot with own-popup filtering; sessions and hosts have separate one-shot
results. Fresh resource/app-generation/destination/cleanup checks remain live.
No private image, FIFO, result/update files, lease or parent-child protocol exists.

Workers join and tty modes restore before exact visible float detach/verification
and original-focus restoration. Those must succeed before a selected effect.
The accepted exception allows the detached execution-host Block through dispatch;
Rex's on-completion/zero-quick-failure exit policy retires it after Go exits.
Go cannot verify its own deletion, and a later failure cannot suppress an earlier
effect. Coordinator must observe exact `session.list_blocks` after exit, not just
absence of a floating layer. SIGKILL/server loss has no guaranteed cleanup.

Phase19 tasks1/2 summaries bind real direct-source runs. New Task03 make-route
observation on reviewed combined source and rebuilt main remains pending here;
Task04 refreshes final evidence. [Acceptance](acceptance.md) distinguishes source
checks from actual runtime. No installed-profile/physical-key/Nix build/platform
claim follows from these static package definitions.

Phase15 bridge/anchor/receipt/completion mechanics and Phase16/18 Go-controller
mechanics are historical, not startup prerequisites. Old assertions/source history
remain; no config, dependency, pin, profile or running server is migrated.
