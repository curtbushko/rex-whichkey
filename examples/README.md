# Server-side Lua options

`mapping.go` and [the mapping standard](../docs/mapping-format.md) own all option,
replacement and mapping rules. These examples are options, not a new public
launcher/setup signature. This task supplies the internal loader only; it does
not register a Rex action, install a keybinding or launch the Go menu.

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

Alternatively put [`my_whichkey.lua`](my_whichkey.lua) beside the **server's**
`init.lua`. Rex documents sibling-module lookup with `require("my_whichkey")`.
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

## Internal composition contract

`require("rex_whichkey.config")` supplies these internal helpers for the Lua
composition owner, without requiring an interpreter or callback VM in Go:

- `resolve(defaults, imported, inline)` selects and detaches raw options. It
  validates configuration fields and the outer spec list, **not mapping
  semantics**. It does not evaluate functions or install a configuration.
- `load(module_name)` loads and checks an imported options table through Rex's
  server-side `require`. Integration namespace names and file paths are rejected.
- `new(defaults, prepare)` constructs a loader with `setup(inline, module_name)`
  and `current()` functions (dot calls). `defaults` must contain all five resolved
  fields from the defaults owner, including the built-in spec; this loader
  invents neither another set of defaults nor any built-in actions.

The composition owner must supply a real `prepare(candidate)` raw-mapping
compiler. It must validate and compile according to the mapping standard, throw
on invalid mappings, and return a non-nil/non-false compiled template. There is
no optional no-op validator or missing-module success path. Compilation must
not invoke action callbacks or evaluate dynamic labels/conditions at setup time;
those belong to invocation evaluation in the originating Rex runtime. The loader
has no dependency on an as-yet-installed compiler module or public setup wrapper.

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

These published contracts informed the loader. No tests, compilation, Rex config
check/reload, live API, remote or GUI exercises were performed. Source formatting
and review do not establish runtime compatibility or a working installed launcher.
