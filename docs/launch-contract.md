# Popup endpoint and preparation contract

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

No exact undocumented Linux/macOS path formula is asserted here. This change has
not been compiled or run by its builder: project tests/TDD, product execution,
live identity queries, GUI/remote trials and configuration reloads are prohibited.
Formatting and source checks do not verify runtime or Darwin behavior. Ordinary
no-argument launch wiring and public help are separate Phase07 delivery work; the
private preparation mode is not the normal user entrypoint.
