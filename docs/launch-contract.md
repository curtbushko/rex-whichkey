# Popup launch, endpoint and preparation contract

## Ordinary startup and one-time setup

After the one-time setup below, run inside the originating Rex terminal:

```sh
rex-whichkey
```

No arguments and no hand-authored JSON are needed: this activates a real Rex
floating popup, not a greeting, demo or standalone terminal menu.
`rex-whichkey --help` (or `-h`) describes the public `rex-whichkey [options]`
entrypoint without contacting Rex. Public options are:

- `--server ENDPOINT`: a concrete execution-host endpoint override.
- `--client FULL_ID`: an exact `client:UUID` attached to this terminal's session.
- `--action NAME`: an installed configuration launch action, default `whichkey`.
  Names contain only lowercase ASCII letters, digits and underscores.

The root parser rejects unexpected positional arguments, unknown flags, missing
flag values and an explicitly empty action before Rex calls. An explicitly empty
server/client is invalid, not an instruction to default. Existing private `menu`,
`prepare`, `controller` and `popup` modes remain compatible; they are not the
ordinary startup interface. Root launch does not read resolved configuration
from the terminal's stdin or run the standalone menu in place of a Rex popup.

Install the matching server-host helper, Rex CLI and Lua modules first. In the
server's **user-owned** `init.lua`, perform ordinary Rex configuration setup once:

```lua
local popup = require("rex_whichkey")
popup.register_completion("whichkey_complete")
local menu = popup.new(require("rex_whichkey.defaults"))
menu.setup() -- or menu.setup(inline_options, "my_whichkey")
menu.register_launch("whichkey") -- omitted name also defaults to whichkey
```

This registers a genuine `rex.action` whose `run(ctx, args)` receives the actual
fresh context in the original server configuration Lua VM. The named action also
works from keys/palette without an origin envelope; existing direct
`menu.launch(ctx)` calls remain supported. Completion and launch names must be
distinct; duplicate module launch registration fails. Use another registered name
with public `--action NAME`. The constructor's optional second endpoint override
remains available and must agree with the public caller's resolved endpoint.
For example, a deployment actually using Rex's documented shared socket can use
`popup.new(defaults, "/run/rex/server.sock")` and launch with
`rex-whichkey --server /run/rex/server.sock`. That is an explicit shared-server
example, **not** the personal Linux/macOS default. The public flag is not forwarded
as a Lua constructor override: preparation resolves separately in the server
process's environment. Matching canonical endpoints are required; a flag cannot
retarget a runtime hosted elsewhere. Changing a terminal's environment does not
change an already-running server's environment. No keybinding is installed;
bind the registered action through your existing Rex configuration if desired.
Modules never edit or reload `init.lua`. Remote Control must permit CLI activation
on the selected client; installing/registering these pieces is not automatic.

Missing helper/CLI/modules, registration, successful setup or Remote Control can
make activation fail. The public caller reports fixed prerequisite guidance,
without forwarding raw child stdout/stderr or private handoffs. Nonzero/timeout/
output-limit failures are uncertain and never retried. A zero exit status means
**activation was accepted**, not proof that a popup appeared or a callback ran.

## Exact terminal origin and client selection

[`LaunchPopup`](../launch.go) captures execution-host `REX_SESSION` and `REX_BLOCK`
as full typed `session:UUID` and `block:UUID` IDs. Missing/malformed IDs fail with
instructions to run in the original Rex terminal. There is no focus/current-block
fallback, session label lookup, ID suffix or graphical-client filesystem lookup.
Resolve one canonical endpoint by the policy below, then use it explicitly for
all read-only preflight and activation calls.

Client precedence is present `--client`, then present `REX_CLIENT`, otherwise a
relationship lookup of **only the captured session's attached clients**. Present
empty, label/suffix, malformed and duplicate IDs fail. Every nonlost attachment
must have a nonblank UTF-8 `connected_at` generation of at most 1024 bytes without
control characters. Rex's documented missing/unknown connection state counts as
connected; only `lost` is disconnected. Without a selector, require exactly one
eligible attachment. Rex terminals are not assumed to export `REX_CLIENT`.
Zero/several eligible attachments fail with instructions to supply an exact full
attached client ID. A supplied selector must match exactly one eligible
attachment; no global or arbitrary-first client selection is used.

The launch route reuses `popupRequest.originBlock` and `popupOriginConnection`:
recheck the selected connection generation, inspect the exact session and require
the exact block to be uniquely placed there. The delivered Rex resource client
rejects malformed/duplicate window/layer/block placement. A wrong endpoint or
stale origin fails before named-action activation, without changing focus or
creating resources. The Lua action and detached controller independently validate
again rather than treating preflight as enduring authority.

## Go-owned launch origin envelope

`launch.go` owns a small fixed `launchOrigin` DTO, separate from mapping options
and the authenticated completion protocol. Go validates each field before strict
`encoding/json` serialization; a single bounded object is emitted, never unknown,
missing, null or duplicate fields. Its five string fields are:

| Field | Meaning and validation |
| --- | --- |
| `server` | Already canonical concrete endpoint; no ambient defaulting in the envelope |
| `session_id` | Captured full typed session ID |
| `block_id` | Captured full typed block ID, uniquely placed in that session |
| `client_id` | Exact full typed attached client ID |
| `connected_at` | Captured nonlost connection generation, compared byte-for-byte, never renewed |

The envelope is at most 64 KiB, encoded as the string argument `origin` inside the
outer JSON argument object. Only stdin carries it. The caller makes one installed
configuration-action activation using an argument array equivalent to:

```text
rex --server CANONICAL_ENDPOINT --client FULL_CLIENT --autostart=false --timeout 10s do ACTION --session FULL_SESSION --block FULL_BLOCK --args -
```

There is no shell, `do -e`, Lua file, second VM, fake/token-only command, daemon or
uncertain retry. Preflight/activation share a 20-second context; individual
preflight CLI calls are bounded to three seconds and activation to ten seconds,
with earlier cancellation/deadlines taking precedence. Activation stdout/stderr
share a discarded 64 KiB total output budget and a 250 ms pipe shutdown allowance.
Errors are fixed safe text; top-level diagnostics quote errors against terminal
escape injection. No private completion/capability data is exposed.

`menu.register_launch` declares `origin` as an optional string, rejects other
arguments, and passes the **actual ctx unchanged** to `menu.launch`. Lua consumes
only the five documented fields, rejecting missing/unknown fields, wrong types,
malformed full IDs, blank generation and oversized envelope data. It compares
session/block/client with actual ctx and checks live `connected_at` **before**
issuing/evaluating a reservation. Opaque `ctx.server` is retained by identity,
never stringified, serialized or replaced by `server`.

Across yielding issue/evaluation, connection inspection, preparation and bind
calls, the launcher rechecks actual IDs/opaque server identity, connection and
configuration generation and original finite lifetime. Before controller start,
Go's canonical preparation `server` must exactly equal the expected envelope
endpoint. A detected mismatch cancels any newly issued reservation without
starting a controller or dispatching/focusing a selected action. The action also
rechecks after the yielding controller-start call and cancels on failure. As in
the retained bridge, no API provides an atomic check-and-start/focus/dispatch
transaction: external disconnects/moves can race a final check. The controller
therefore performs its own exact-origin checks before mutation and completion
still requires a live authenticated retained reservation. No failed reservation
is reopened; no selected action is retried.

This envelope is transport context, **not** executable Lua, a capability,
authority to fabricate ctx, or a second mapping schema. Other callers supplying
it cannot change the actual action activation or bypass retained registry checks.

## Complete production route (source trace)

1. [`cmd/main.go`](../cmd/main.go) parses no arguments/public flags through
   `runLaunch`, then calls `LaunchPopup`, not private-mode usage or a JSON demo.
2. [`launch.go`](../launch.go) resolves/fixes endpoint, validates the exact
   session/block/client connection and invokes the installed named CONFIG action.
3. [`lua/rex_whichkey.lua`](../lua/rex_whichkey.lua) registers that action in the
   original VM; its real `run(ctx,args)` calls `menu.launch(ctx, optional_origin)`.
   [`config.lua`](../lua/rex_whichkey/config.lua) setup previously resolves
   defaults/import/inline options transactionally and compiles the raw mappings.
   [`defaults/init.lua`](../lua/rex_whichkey/defaults/init.lua) supplies the shipped
   menu; no new hard-coded demo or mapping schema is involved.
4. [`dispatch.lua`](../lua/rex_whichkey/dispatch.lua) `issue` reserves and evaluates
   the retained mapping snapshot in the original VM. The launcher encodes actual
   resolved configuration, captures private Go `prepare` output, validates endpoint
   and origin again, binds the secret capability, then starts the detached finite
   controller with a private stdin handoff.
5. [`popup.go`](../popup.go) `RunPopupController` verifies the frozen origin before
   mutation, creates the owned popup terminal layer and invokes `RunPopupTerminal`.
   [`terminal.go`](../terminal.go) `RunTerminal` uses real `RenderMenu` redraws and
   `HandleKey` navigation/selection; terminal cleanup succeeds before result commit.
6. The independent controller closes only its owned popup, revalidates/focuses the
   exact origin and sends authenticated fresh completion to the installed action.
   The original VM checks one-shot capability, connection/configuration generation,
   snapshot, expiry, confirmation and focus; original closures remain there and
   named/client actions use the actual completion activation. No popup-origin or
   fabricated-context dispatch is substituted.

## Execution-host routing

`rex.CanonicalEndpoint` in [`rex/endpoint.go`](../rex/endpoint.go) is the
popup's single Go-owned transport validator. `rex.ResolveEndpoint(ctx, explicit)`
selects one endpoint in this order:

1. A present explicit override, including an explicitly empty value (which fails
   validation). No invalid override falls through to another source.
2. A nonempty execution-host `REX_SERVER`, validated by the same rules. A malformed
   environment value fails rather than trying the platform default.
3. On Linux or macOS only, the installed matching Rex CLI's native default:
   `rex --autostart=false --timeout 2s whoami --json`. No `--server` is supplied.
   The command is an argv invocation, not a shell command. Its context is bounded
   to two seconds, with a 250 ms process-pipe shutdown allowance and a 64 KiB
   stdout bound. Child stdin is closed and stderr is discarded.

An empty `REX_SERVER` permits native discovery, unlike an explicitly empty
constructor/public override. A nonempty malformed environment value never falls
through. For custom/shared deployments, set matching concrete overrides or
execution-host environments on both public caller and server-side preparation.

The native branch removes inherited `REX_SERVER`, `REX_CLIENT`, `REX_SESSION`,
`REX_WINDOW` and `REX_BLOCK` selectors from the child environment. It leaves Rex's
native home, runtime, authentication and XDG environment intact. The platform
socket-directory/filename algorithms belong to Rex, not rex-whichkey: there is
no copied path table, viewing-client filesystem lookup, socket scan, explicit
directory creation or server autostart in the resolver. This branch is an ordinary
read-only identity query during product startup, not a second service. It requires
a reachable default server and valid Rex runtime/home context on that host.

The identity response must be one UTF-8 JSON object with a top-level string
`endpoint`. Other identity fields are ignored; duplicate top-level fields,
invalid/missing endpoints, oversized output and CLI failures are rejected.
Errors are sanitized and actionable, never raw identity JSON or child diagnostics.
There is no retry, alternate socket or alternate server after failure. Other
platforms require an explicit endpoint or `REX_SERVER` rather than automatic
native discovery; the product's existing platform build constraints still apply.

## Concrete endpoints, not identities

Accepted forms are an absolute Unix socket path (normalized by adding `unix://`),
a strict absolute `unix:///...` URI, or `host:port`, including bracketed IPv6.
Ports must be decimal, nonzero and fit 16 bits. Unix URIs must have an empty host,
a non-root absolute path, no user/query/fragment and no encoded-path ambiguity.
Invalid UTF-8, whitespace, controls, mutable Rex host labels and missing addresses
are rejected. Socket paths retain their spelling; no filesystem or symlink lookup
is used to rewrite them. This validation retains the previous popup URI/address
checks and adds absolute-path input normalization. Transport authentication is
still Rex's responsibility.

After resolution, consumers pass the concrete endpoint explicitly. The controller
calls `CanonicalEndpoint` directly, **not** `ResolveEndpoint`: an absent controller
server is always invalid and can never acquire a default. Generic `rex.New` keeps
its existing broader behavior; the popup supplies its validated server explicitly.
A routable endpoint is not the opaque Lua `ctx.server`. That value stays unchanged
in the original action context and remains subject to the dispatcher's original-VM
identity checks.

## Captured preparation API

The existing Go `PreparePopup(input, output)` remains a wrapper for callers without
an override. `PreparePopupWithServer(ctx, input, output, server *string)` adds
cancellation and a present/absent override. The private CLI forms are:

- `rex-whichkey prepare`
- `rex-whichkey prepare --server ENDPOINT`

Both read the actual resolved Go configuration on stdin, validate it through
`DecodeConfig`, require the server-host Rex executable, resolve the endpoint, and
only then generate an invocation capability. Captured stdout is one JSON object:

| Field | Meaning |
| --- | --- |
| `capability` | Go-generated 64-character lowercase hexadecimal single-use invocation secret |
| `executable` | Absolute path of this rex-whichkey executable on the execution host |
| `server` | Concrete canonical endpoint selected by the policy above |

This response is private plumbing, not ordinary user output. A terminal output
file is refused; the launcher must capture stdout. Capability/configuration and
controller/completion data stay off argv and logs. Discovery errors cannot include
private handoff data. `REX_BIN_DIR`, if supplied to popup preparation, must be an
absolute execution-host directory containing the matching Rex CLI; otherwise the
installed helper uses that host's PATH. No client-local executable is discovered.

## Actual Lua consumer and retained safety

`popup.new(defaults)` is valid. `popup.new(defaults, endpoint)` retains an optional
immutable constructor-only override outside mapping options. `menu.setup` cannot
replace it. Lua checks only the optional argument's type; Go owns transport syntax
and precedence. The real `menu.launch(ctx)` preparation argv includes `--server`
only when that constructor argument was supplied. An empty supplied string is not
silently treated as absent.

Lua validates the preparation response's capability/executable/server shape and
forwards **Go's returned `server`**, not the constructor spelling or `ctx.server`,
in the existing detached controller's private stdin handoff. It still reserves a
ticket before evaluating mappings, binds after yielding preparation through the
registry's generation/connection rechecks, and cancels the ticket on launch failure.
Preparation failures, including invalid response data, reach the caller as fixed
safe guidance: check the execution-host rex-whichkey helper, matching Rex CLI,
running server/runtime environment and endpoint override or `REX_SERVER`. Lua
still discards child stderr and never forwards child stdout or arbitrary error
objects to diagnostics. Later launch failures retain the generic cancellation
message; neither failure path retries or selects a fallback server.
The original `ctx`, opaque server identity, session/block/client identities and
connection generation are not synthesized from routing data.

The controller still checks exact placement and connection generation through the
explicit endpoint before arming mutations or cleanup. One-shot capability, expiry,
configuration generation, retained snapshot, confirmation, focus revalidation,
original-VM callbacks, owned-only bounded cleanup and no uncertain action retries
are unchanged. This does not install a daemon, copy callbacks into a new Lua VM,
select a focused/default client, or edit/reload the user's `init.lua`. The helper,
Rex CLI, original-server Lua modules, explicit user-owned action registration and
app Remote Control policy remain prerequisites.

## Source evidence and verification limits

The published [CLI contract](https://www.superlogical.com/rex/docs/automate/cli)
documents execution-host `REX_SERVER`, `REX_BIN_DIR`, explicit server selectors and
machine-readable JSON. The published [shared-server instructions](https://www.superlogical.com/rex/docs/linux/shared-server)
use absolute socket paths with `-S`. Installed Rex
`0.0.1792609031-gc465ec` read-only `whoami --help` documents identity/server/location
JSON and `--autostart`/`--timeout`. The coordinator-retained identity fixture has a
top-level concrete `endpoint`; this consumer requires that shape and rejects
unsupported shapes rather than guessing field alternatives. These are source/help
contracts, not a claim that every CLI version has been exercised.

No exact undocumented Linux/macOS path formula is asserted here. The earlier
explicit-only endpoint decision REXWHICHKEY-20261101-015342 was refined by accepted
REXWHICHKEY-20261101-101502 under delegated authority to meet REQ-017. Native CLI
discovery replaces the draft duplicated-path proposal REXWHICHKEY-20261101-100106
(Superseded); explicit overrides, opaque identity separation and fail-closed
origin safeguards remain. This document reflects that choice, not a new approval
or rewrite of historical evidence.

Endpoint task 07-01 and launch task 07-02 have recorded independent mixed Go/Lua
APPROVED source reviews and coordinator `go build ./...`/Git integration evidence
at `d1164e7` and `69078fc`, respectively; see the [delivery inventory](acceptance.md)
for exact identities and artifact refs. The six-file usage documentation received
independent source-consistency review and was integrated at `1af28738`; the
inventory also records coordinator Go compilation and Nix package builds of the
delivered source. These are supplied revision-bound results, not commands run by
this documentation builder or approval of later corrections. No unrecorded
`task build` result or runtime acceptance is inferred.

No project tests/TDD, example/product execution, new live identity queries,
GUI/remote trials or configuration reloads were performed for this revision.
Formatting, compilation and source checks do not verify runtime or Darwin
behavior. The complete route above is a source trace, not a claimed popup,
Remote Control or callback pass. Private preparation is not the normal entrypoint.
