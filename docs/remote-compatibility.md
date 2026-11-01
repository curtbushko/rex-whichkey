# Remote-context compatibility

## Status and scope

This describes the delivered source paths for REQ-003 (configuration loading),
REQ-005 (Rex-runtime callbacks), REQ-011 (origin/focus/cleanup) and REQ-013
(remote-server compatibility). It is not a remote acceptance result. No project
tests, compilation, live API/GUI probes, remote trials or configuration reloads
were run for this documentation task. Remote execution, disconnect/reconnect,
moved-block behavior and cleanup remain runtime-unverified.

The application is a short-lived Go controller and popup terminal, launched by
minimal Lua in the originating Rex configuration runtime. It adds no daemon,
independent host registry or transport framework. Listing Rex-known hosts does
not implement remote connection, attachment, bootstrap, Wake-on-LAN or SSH.

## Execution host, paths and configuration

The graphical client may be on a different machine. The delivered popup path
expects the helper, controller and popup child to execute on the **Rex server
host**, with access to the same server-local executable and private temporary
files. It does not copy binaries, modules or configuration from the graphical
client or translate client-local absolute paths.

| Boundary | Delivered resolution |
| --- | --- |
| Lua integration and user options | `lua/rex_whichkey.lua` requires the installed `rex_whichkey.*` modules in the server runtime. `lua/rex_whichkey/config.lua` loads an optional user module with server-side `require`, including files beside that server's `init.lua`. It does not rewrite `package.path` or infer a client home directory. |
| User module name | A dotted Lua module name, not an absolute/relative filename. The loader rejects the `rex_whichkey` integration namespace and its submodules for user configuration. The imported module must return an options table. |
| Initial helper | Lua uses argument-array `rex.exec({ "rex-whichkey", "prepare" }, ...)`, so the server process's command environment/PATH must resolve the installed helper. A graphical client's PATH does not install it on the server. |
| Controller and popup | `PreparePopup` in `popup.go` returns `os.Executable()` through captured output; Lua starts that server-local executable as a detached controller. The controller uses its own `os.Executable()` for the popup terminal command. No client-discovered binary path is substituted. |
| Rex CLI | The generic `rex` package resolves explicit `Options.Executable`, otherwise `REX_BIN_DIR/rex`, otherwise `rex` from the execution host's PATH. A configured but unusable executable/directory produces an error, not fallback to another copy. Popup composition separately requires `REX_BIN_DIR` to be absolute when set and uses the resolved CLI for its client and completion calls. |

The popup terminal command is an argv list with `shell = "none"`; Go CLI calls
also use argv lists, not shell interpolation. No interactive shell/profile is
needed to interpret the popup command. The server must already have a suitable
command environment, executable permissions and module installation. PATH lookup
alone does not establish that the selected CLI version matches the server.

Client/block actions queued by the CLI-origin completion activation also require
**Remote Control** to be enabled in the originating graphical client's settings.
The published [Custom Actions contract](https://www.superlogical.com/rex/docs/customize/actions),
under “Calling Other Actions”, says the app accepts CLI actions only with that
setting enabled. Such selections can therefore be rejected when it is disabled,
even with valid endpoint/origin IDs. This client policy is separate from Rex
server transport authentication and which-key capability validation. Its behavior
remains runtime-unverified here; this documentation task changes no settings.

Configuration precedence is **defaults < imported module < inline options**.
Present `spec` and `icons` values replace those whole values rather than merging
bindings/markers. Repeated setup starts from the baseline; candidate loading and
compilation must succeed before replacement, and failure preserves the installed
configuration. Go's `DecodeConfig` in `config.go` accepts bounded, resolved,
data-only JSON; it does not load Lua files or execute callbacks. The canonical
fields, mapping rules and completion data belong to [`mapping.go`](../mapping.go)
and [the mapping standard](mapping-format.md), not a second Lua schema.

The user owns `init.lua`. Installation/registration is explicit opt-in, using the
initialization fragment in the mapping standard: load the integration/defaults,
register the completion configuration action at load time, construct the menu
with this server's fixed endpoint, call setup and register the launch action.
Nothing here writes or reloads that file. Imported modules are trusted Lua code;
this loader is not a configuration sandbox or a hot-reload mechanism.

## Endpoint selection is not runtime identity

`popup.new(defaults, server_endpoint)` requires an explicit constructor-bound
endpoint for **the server hosting that Lua runtime**. It is not a mapping option
and setup cannot replace it. The authoritative Go validator accepts an absolute
`unix:///...` socket address or `host:port` (including bracketed IPv6), with a
nonzero numeric port. Mutable host labels, missing endpoints and automatic
conversion of `ctx.server` are not accepted popup bindings. An absolute socket
path is local to the execution host, not to the graphical client.

The Lua `ctx.server` value is an opaque runtime identity. The dispatcher retains
it and compares it with `rawequal` on the fresh completion action. It is never
serialized or stringified into a CLI address. The separately configured endpoint
is routing data, not replacement authority for that retained identity.

Before creating resources or arming deferred focus/cancellation, the controller
uses that exact endpoint to find the captured session, placed block and client
connection generation. Missing/mismatched origin data fails read-only at this
stage. The endpoint must still be supplied correctly by the server owner; these
checks are not endpoint discovery, a new authentication protocol or proof of
remote deployment. Transport access/authentication remains Rex's responsibility.

`rex.New` has a broader reusable CLI contract: it captures `Options.Server`, or
`REX_SERVER` when the option is empty, and refuses an absent server. Calls retain
that selector and pass `--server` and `--autostart=false`; they do not silently
start or select a local server. The popup supplies its constructor-bound endpoint
explicitly, so it does not use that environment fallback. Generic API calls pass
an explicit configured session when present; they do not default to
`REX_SESSION`. `Session`/`Window`/`Block` resources retain their captured server and
stable session IDs rather than following environment changes or mutable labels.

## Origin propagation and post-return dispatch

1. The original action supplies full `session_id`, `block_id` and `client_id` plus
   opaque `ctx.server`. `dispatch.lua` reserves an invocation before evaluation,
   inspects placement and the originating client's `connected_at` connection
   generation, and retains callbacks, arguments, guards and a mapping snapshot in
   that same server Lua VM. Missing/ambiguous identities fail; another current
   client is not selected instead. Only `connection_state == "lost"` is treated
   as disconnected; missing/unknown state is accepted only with the required
   matching client and nonempty generation data.
2. Go preparation validates resolved configuration and generates a capability
   with the Go-owned `crypto/rand` helper. Lua binds it to the retained ticket
   before starting the controller. The ticket and executable functions never
   leave Lua. Capability-bearing preparation output is captured, and the private
   controller handoff travels on stdin, not argv or terminal output. The original
   action returns without waiting for a selection; `detach = true` lets the
   bounded controller outlive that action.
3. The controller receives the explicit endpoint, captured session/block/client
   IDs and connection generation. It creates one floating terminal in the
   origin's currently resolved window. Its private directory is mode 0700 and
   its two reserved files are mode 0600; the popup child reads configuration and writes a
   selection there, but does not receive the capability or Lua ticket. The
   directory path/deadline in popup argv are not dispatch authority.
4. After successful terminal restoration, result validation and owned-popup
   cleanup, the controller rechecks the origin and restores server focus. It
   activates the installed **configuration action**, using explicit `--server`,
   `--client`, `--session` and `--block` targets. Completion JSON goes through
   stdin (`--args -`), with child output/diagnostics discarded. This is not a
   dispatch script in a separate Lua VM or an activation from the popup block.
5. The fresh action must match the retained opaque server and exact IDs. The
   dispatcher strictly decodes the Go-owned completion data, claims the entry
   once, checks expiry/configuration/connection state, resolves the retained
   selection and confirmation, rechecks guards and restores origin focus. After
   a post-focus placement/connection check it calls the original callback with
   retained origin context and args, invokes a custom named action with
   `rex.invoke`, or queues a dotted client/block action with `rex.client.queue`
   on that actual activation. Copying fields into a Lua table is not a substitute
   for targeting the real action context.

Selection cannot supply replacement action names, arguments or executable Lua.
Configuration replacement invalidates the previous registry. The Go-owned limits
are 30 seconds from reservation and 64 pending entries per registry, including
in-progress work; redraws do not renew the lifetime. Lua expiry is lazy on
registry operations, and backwards wall-clock movement revokes the generation.
Malformed/unknown/wrong-origin requests cannot consume unrelated entries;
authenticated claimed entries are consumed even when later checks or execution
fail. Uncertain completion is not retried or followed by a competing cancellation.
Authenticated cancellation releases the invocation without executing a selection
or conditions; the Lua cancellation handler itself does not change focus.

## Moved blocks, focus and owned cleanup

`rex/session.go` resolves human-readable session targets once; session and window
resources retain stable IDs. `rex/block.go` treats window/layer membership as a
snapshot: inspection finds the block in its captured session's current layout, and focus
names the stable block/session, not an old window. A move **within that session**
can therefore be followed; a missing/detached block or a move outside the captured
session is an error, not permission to search other sessions or focus a fallback.
Server block/window focus is not a request to switch an arbitrary graphical
client to a different session.

`rex/layer.go` resolves the origin's current window when creating a floating
terminal and captures the returned owned layer/block IDs. Cleanup closes only
that returned terminal, even if moved. It refuses cascading `close_layer`, window,
session, label-based or replacement-resource cleanup. If the layer remains after
closing its terminal, it reports an error instead of closing other occupants.

The controller leaves five seconds of the original lifetime for UI teardown and
completion, and has a separately bounded five-second deferred cleanup context
that survives caller cancellation. Dismissal/errors attempt owned cleanup,
validated origin focus restoration and authenticated cancellation when completion
has not already been attempted. Temporary cleanup unlinks only `config.json`,
`result.json` and the empty invocation directory; it is not recursive deletion.
This is bounded best effort, not a guarantee after process/server termination.

## Rex-known inventory has different targeting

`rex/hosts.go` intentionally runs `rex hosts --json`, **not** `hosts.list` through
the origin client. Its effective home/local inventory belongs to the machine
executing that CLI and its configured Rex home. `--server`, `REX_SERVER` and the
origin session/client do not select its inventory. In the delivered popup path
this normally means the server host, not the graphical client's known hosts.
Rex performs its own merge, cached-home handling, label resolution and held-change
forwarding; the menu does not implement a separate registry. The CLI may itself
contact home and forward previously held changes, so listing is not a claim of
no network activity or no Rex-managed synchronization.

`hosts.go`/`terminal.go` expand only enabled `known-hosts` provider markers, once
per invocation with a three-second inventory bound. No marker means no read.
Host IDs remain distinct from mutable labels and machine `host_id`; warnings and
home metadata are retained, and listing failure becomes a visible warning, not
successful empty inventory. All generated rows are informational/non-executable.
Endpoints are data, not reachability results, and selecting a row neither changes
the popup endpoint nor attaches to a remote host.

## Residual limits and source ownership

- **Deployment/API compatibility (02-01, 03-01, 04-05):** server-host module/PATH
  setup, shared executable/temp filesystem access and compatible installed Rex
  APIs are prerequisites, not verified deployment results. The published Server
  API describes version-dependent methods; this document establishes no minimum
  tested Rex build or cross-platform certification. The delivered popup code
  uses Unix process/filesystem primitives. Generic client support for a remote
  selector does not make client-host controller launch or arbitrary cross-host
  endpoint substitution a supported popup deployment.
- **Races (02-03, 03-04, 04-05):** lookup/create and inspect/focus/execute are
  separate Rex operations. A concurrent external move, focus change, disconnect
  or configuration change can race the final check/side effect. Rechecks and
  explicit targeting reduce exposure but are not an atomic transaction. Callback
  side effects cannot be rolled back by a later expiry/error check.
- **Uncertain mutations and termination (02-03, 04-05):** creation may succeed
  before a timeout/decode error hides its identity. Without the returned owned
  terminal ID, broad cleanup is deliberately refused and a resource may remain.
  Network failure, SIGKILL or server exit can also prevent cleanup/completion.
  No retry, scavenging daemon or label search compensates for this. Lua's lazy
  expiry retains bounded state until a later registry operation or VM teardown.
- **Trust and identity (03-04, 03-05, 04-05):** the implementation relies on Rex
  supplying a consistent opaque server identity and a meaningful `connected_at`
  generation for a client across activations. These semantics, real queued
  client-action routing and disconnected-origin rejection have not been exercised
  here. Private transport protects against other users under normal filesystem
  permissions; it is not isolation from trusted user Lua, the same OS account or
  privileged processes. Confirmation is the trusted popup's assertion, not
  cryptographic proof of a physical keypress.
- **Inventory (02-04):** successful listing, cached metadata and endpoint strings
  do not establish host/service reachability or origin-server equivalence.

These are implementation limitations and runtime-unverified behavior, not passes
for REQ-003/005/011/013. Runtime fixes belong to the source owners above; this
file owns documentation only. Existing tests are preserved, and this document
does not schedule tests, remote trials, reloads or changes to user configuration.

## Source and contract references

- [`lua/rex_whichkey.lua`](../lua/rex_whichkey.lua),
  [`config.lua`](../lua/rex_whichkey/config.lua) and
  [`dispatch.lua`](../lua/rex_whichkey/dispatch.lua): server loading, launch and
  original-VM completion authority.
- [`popup.go`](../popup.go), [`config.go`](../config.go) and
  [`cmd/main.go`](../cmd/main.go): process composition, decoding and private
  transport; [`mapping.go`](../mapping.go) and [mapping-format.md](mapping-format.md)
  remain the shared-data standard.
- [`rex/rex.go`](../rex/rex.go), [`exec.go`](../rex/exec.go),
  [`session.go`](../rex/session.go), [`window.go`](../rex/window.go),
  [`block.go`](../rex/block.go), [`layer.go`](../rex/layer.go) and
  [`hosts.go`](../rex/hosts.go): execution and resource/inventory targeting.
- Published Rex [CLI](https://www.superlogical.com/rex/docs/automate/cli),
  [configuration](https://www.superlogical.com/rex/docs/customize/config),
  [custom actions](https://www.superlogical.com/rex/docs/customize/actions),
  [Lua API](https://www.superlogical.com/rex/docs/reference/lua),
  [Lua scripts](https://www.superlogical.com/rex/docs/automate/lua-scripts) and
  [Server API](https://www.superlogical.com/rex/docs/automate/server-api) contracts:
  server-local configuration, argv execution/detach, explicit CLI targeting,
  configuration versus script actions, the client's Remote Control policy for
  CLI-origin actions and version-dependent API documentation.
  Reading these contracts is source evidence, not a live-server verification.
