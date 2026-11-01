# Server-side setup and options

`mapping.go` and [the mapping standard](../docs/mapping-format.md) own all option,
replacement and mapping rules. Lua consumes that standard; callbacks stay in the
original Rex configuration VM, not in Go or a separate script runtime.

## One-time user-owned setup and ordinary launch

Install the matching helper, Rex CLI and `rex_whichkey` Lua modules on the Rex
server host as described in the [README](../README.md). Make the modules available
to that runtime's module lookup, and both executables available on the terminal
and server process's PATH. In your own server-side `init.lua`:

```lua
local popup = require("rex_whichkey")
popup.register_completion("whichkey_complete")
local menu = popup.new(require("rex_whichkey.defaults"))
menu.setup()
menu.register_launch("whichkey")
```

Setup must succeed. Completion is registered once in the same VM; its name must
be distinct from the launch action. No socket literal is required for the normal
Linux/macOS default-server case. This code registers the real configuration
action, not a keybinding. Normal Rex keys/palette can invoke it with their actual
context; existing direct `menu.launch(ctx)` remains available.

With client **Remote Control** enabled, invoke from the originating Rex terminal:

```sh
rex-whichkey
```

The command captures full `REX_SESSION`/`REX_BLOCK`, then uses an exact `--client`
or present `REX_CLIENT`, or requires exactly one eligible attached client in that
session. Missing/stale/ambiguous origin is an error, never a focus/current-client
fallback. Use `--help` for public flags. To use another launch name, replace the
last setup line with `menu.register_launch("my_whichkey")` and invoke
`rex-whichkey --action my_whichkey`. Launch names use lowercase ASCII letters,
digits and underscores only. No user-written JSON or private mode is needed;
standalone `menu` is a low-level JSON UI, not this floating-popup launch path.
Accepted activation is not proof of display or callback execution. Do not retry
an uncertain activation/action.

Endpoint precedence is present explicit override > nonempty execution-host
`REX_SERVER` > bounded `rex --autostart=false --timeout 2s whoami --json` native
Linux/macOS discovery. Go validates/captures the concrete endpoint; Rex owns its
platform path/runtime/home rules. No viewing-Mac paths, guessed filenames,
server autostart or alternate socket probing are used. The normal default server
must already be reachable. Missing runtime/home context, malformed endpoints or
unsupported identity output fails with actionable guidance.

For a custom/shared server, Lua's optional `popup.new(defaults, endpoint)` binds
an immutable preparation override; `menu.setup` does not change it. Supply a
matching public `--server ENDPOINT` or environment on the terminal caller. The
public flag is not forwarded to Lua preparation: both canonical results must
agree. For example `/run/rex/server.sock` is Rex's documented **shared-server**
socket, not a personal Linux/macOS default. An absolute socket path, strict Unix
URI or numeric-port `host:port` is routing data, never opaque `ctx.server` or a
mutable host label. See the [launch contract](../docs/launch-contract.md) for
exact precedence/validation and the [README troubleshooting](../README.md#troubleshooting-and-safety)
for module/PATH, client, endpoint, registration and uncertain-activation failures.

You own `init.lua` and ordinary Rex configuration administration. Nothing here
edits or reloads it, creates a daemon or installs a separate Lua VM.

## Inline and sibling-module configuration

In your own Rex `init.lua`, inline options have this form:

```lua
local inline = {
  position = "center",
  title = "", -- explicit empty title, not a request for a default
  icons = { group = ">" }, -- breadcrumb becomes empty: whole-table replacement
  spec = {
    { "w", group = "Sessions" },
    { "wn", "session.new", desc = "New session" },
  },
}
```

Replace the initial `menu.setup()` with `menu.setup(inline)` to use that table.
Alternatively put [`my_whichkey.lua`](my_whichkey.lua) beside the **server's**
`init.lua` and use `menu.setup({}, "my_whichkey")`, or
`menu.setup(inline, "my_whichkey")` to overlay inline options. Rex documents
sibling-module lookup with `require("my_whichkey")`.
The installed integration namespace is `rex_whichkey`; user modules must have
another name. The internal loader accepts dotted module names such as
`my_whichkey` or `my_plugins.whichkey_options`, never client-local file paths.
It uses the server runtime's `require` without changing `package.path`, the
working directory or `REX_CONFIG`. Install the integration's `lua/` directory in
that same runtime's module search path using packaging, not by naming a local
workstation checkout in a remote configuration.

You own `init.lua`; nothing in these examples writes or reloads it. Ordinary
`require` caching applies. The loader neither clears `package.loaded` nor reloads
modules behind the user's back. A module must return an options table; missing
modules, syntax errors and modules returning other values are errors, not empty
configurations. As with ordinary Lua configuration, imported files are trusted
executable code. Their arbitrary external side effects cannot be rolled back.

## Internal composition reference (not user wiring)

The public setup above already supplies the real compiler and registry through
`lua/rex_whichkey.lua`; users do not need to assemble this internal loader.
`require("rex_whichkey.config")` supplies these helpers without requiring an
interpreter or callback VM in Go:

- `resolve(defaults, imported, inline)` selects and detaches raw options. It
  validates configuration fields and the outer spec list, **not mapping
  semantics**. It does not evaluate functions or install a configuration.
- `load(module_name)` loads and checks an imported options table through Rex's
  server-side `require`. Integration namespace names and file paths are rejected.
- `new(defaults, prepare)` constructs a loader with `setup(inline, module_name)`
  and `current()` functions (dot calls). `defaults` must contain all five resolved
  fields from the defaults owner, including the built-in spec; this loader
  invents neither another set of defaults nor any built-in actions.

The public composer supplies `mapping.compile(candidate.spec)` as the real
`prepare(candidate)` raw-mapping compiler. It validates and compiles according
to the mapping standard, throws on invalid mappings, and returns a non-nil/
non-false compiled template. There is no no-op validator or missing-module
success path. Compilation does not invoke action callbacks or evaluate dynamic
labels/conditions at setup time; those belong to invocation evaluation in the
originating Rex runtime.

`setup` imports, resolves and compiles before publishing. It returns detached raw
options and the compiler-owned template. `current()` returns the last successful
raw options and template, or nil before any successful setup. Templates must be
treated as immutable by their owner. Input and returned raw option tables are
copied; functions keep their original runtime/closure identity. Cyclic tables and
metatables are rejected. Reentrant setup fails instead of interleaving installs.

Precedence is **defaults < imported < inline**. Each setup starts again from the
original defaults. `spec` and `icons` replace whole values rather than merging:
inline `spec = {}` clears imported/default bindings; inline `icons = {}` clears
both markers. Absent/nil values retain the lower-priority value. Wrong types,
unknown fields and invalid display text fail, including when found in a supplied
lower-priority options table. Overridden specs never execute dynamic functions.

A failed import, resolution or compilation leaves the last successful loader
configuration/template installed. Invocation callback lifetimes, replacement
invalidation, dynamic evaluation and dispatch are the corresponding integration
owners' responsibilities; this module does not claim to implement them.

## Source references and verification limits

- [Rex configuration: server location and sibling modules](https://www.superlogical.com/rex/docs/customize/config)
- [Rex Lua reference: Lua 5.1 and standard `require`](https://www.superlogical.com/rex/docs/reference/lua)

These published contracts informed the loader. This documentation revision uses
source inspection/whitespace checks only: no example execution, tests, compilation,
Rex config check/reload, live API, remote or GUI exercises. Earlier coordinator
compilation and independent source reviews are recorded separately in the
[delivery inventory](../docs/acceptance.md). They do not establish runtime
compatibility, Linux/macOS certification or a working installed launcher.
