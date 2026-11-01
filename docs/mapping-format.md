# Mapping and configuration standard

`mapping.go`, in the root Go package `whichkey` (module `temp`), owns these data
values and validation rules. This document explains that source, not a parallel
Lua schema. It covers REQ-003/004/005/013 and the confirmation data used by REQ-007.
No Lua interpreter, reflection mapper, dependency, launch transport or new public
setup API is selected here.

## Configuration, defaults and replacement

`Config` is resolved invocation data. Its fields and JSON names are `position`,
`title`, `icon`, `icons` and `spec`. `ConfigOptions` represents optional overrides;
its pointers distinguish absence from explicit empty values. Lua `nil` means
absent, not false, empty string or an empty table. Unknown options and wrong Lua
value types are errors, not silently ignored extensions.

- `position`: exactly `top`, `bottom`, `left`, `right` or `center`.
- `title`, `icon`: strings; an explicit empty string is valid.
- `icons`: a table/object containing string `breadcrumb` and `group` markers.
  An empty marker suppresses it. An explicit icons table replaces the whole table;
  omitted markers in that replacement become empty strings. Unknown marker names
  and non-string values are errors.
- `spec`: an ordered list of mappings/scopes in Lua; an ordered list of evaluated
  `Mapping` records in Go. An explicitly empty list replaces all default bindings.
  A nil/empty resolved Go slice is an empty menu, not a request to load defaults.

`DefaultConfig()` supplies display defaults: center position, title `Rex`, icon
``, breadcrumb `»` and group marker `+`. It intentionally does not define built-in
actions. The built-in-action owner supplies the default spec before resolving user
options; no fabricated action or undefined future Go symbol is required here.

`ResolveConfig(defaults, imported, inline)` implements fieldwise precedence:
**defaults < imported values < inline values**. A present `spec` or `icons`
replaces that entire value, never appending bindings or deep-merging markers.
For example, inline `spec = {}` clears imported and built-in bindings, while an
absent inline spec retains the imported spec. An explicitly empty title/icon does
not fall back. Invalid positions, including `""`, are errors.

An imported Lua module returns options in the same format as inline options. It
must have a name distinct from the installed integration module. Module lookup
and paths belong to the Rex server host, never to a client-local absolute path.
The user owns `init.lua`; this standard neither rewrites it nor chooses a new
loader/module-name contract. Loading/compilation errors must be reported rather
than treated as an empty configuration.

A setup replacement starts from defaults and its supplied imported/inline values,
not from a previous setup's accumulated bindings. Compile and validate a candidate
before atomically publishing it. A repeated successful setup replaces the previous
configuration; a failed setup retains the last working configuration. Choose the
winning raw values before evaluating dynamic mappings, so overridden specs do not
execute functions. Go's resolver mirrors this precedence on already evaluated
data; it does not load Lua. It returns detached slices, argument bytes and optional
values, leaving inputs unchanged. Treat installed configurations as immutable.

## Lua input and inheritance

The user-facing notation is a WhichKey-inspired ordered list:

```lua
{
  { "w", group = "Sessions" },
  { "wn", "session.new", desc = "New session" },
  {
    confirm = true,
    { "x", "pane.close", desc = "Close pane" },
  },
  { "h", function(ctx, args) --[[ runs only in Rex Lua ]] end,
    desc = "Custom action", args = { name = "example" } },
}
```

This is a spec value, not a promise of a particular `setup` signature. Each entry
is either a mapping whose first positional value is a **full key sequence**, or
an attribute scope containing child entries. Nested scopes do not concatenate key
prefixes. A mapping's second positional value is absent, a nonblank named action
string, or a Lua callback function. Other positional values, malformed/sparse
lists, mixed child/mapping forms and unsupported attribute names are errors.

Allowed attributes are `desc`, `group`, `icon`, `cond`, `hidden`, `confirm`, `args`.
Each child inherits the nearest explicitly supplied value of each attribute;
child values override that attribute only. In particular, `false` and `""` are
real overrides, not absence. `args` replaces the whole inherited argument table.
The key sequence and action/callback are positional and never inherited.

- `desc`, `group`, `icon`: absent, string, or zero-argument Lua function returning
  a string (or nil for absence). Evaluate in Rex Lua. Action descriptions must
  resolve to a nonblank string. Empty group strings still declare a group.
- `cond`: an absent raw condition or true enables the entry; false disables it.
  A zero-argument Lua function may return true, false or nil; false and nil both
  disable the entry. A nil function result is not an absent raw condition and is
  not an evaluation error. Other return types are errors. Evaluate scope
  conditions before walking their children. A child's condition overrides an
  inherited condition, but cannot revive children of a scope already skipped
  because its condition evaluated to false or nil.
- `hidden`, `confirm`: absent or boolean, default false. Hidden means not drawn,
  **not** inactive. Confirmation applies to named actions and callbacks equally.
- `args`: absent or a data-only object table. Object keys are strings; nested
  objects, dense arrays, strings, finite numbers and booleans are supported.
  Functions, userdata, threads, cyclic tables and ambiguous/mixed/sparse tables
  are invalid. Empty args means an object, not a JSON array or null. Nil/absent
  args becomes an empty argument object at dispatch. Args are not Lua source or
  shell fragments; action-specific argument checks belong to the action owner.

Evaluate conditions and dynamic strings once per invocation, not every redraw.
Do not invoke an action callback during loading, compilation or evaluation.
Evaluation/type errors abort the candidate invocation; do not silently show a
partially evaluated menu or execute a fallback action. Disabled entries do not
create prefixes or claim active keys. Display-only filtering must not remove
hidden bindings from lookup or dispatch.

## Go mapping records and keys

`Mapping` is the evaluated data form after scope inheritance, not a representation
of executable functions. Its JSON names come directly from `mapping.go`:

| Field | Meaning |
| --- | --- |
| `keys` | Nonempty full key sequence. |
| `desc`, `icon` | Resolved strings; omitted means empty. |
| `group` | Optional resolved label; absent differs from explicit empty string. |
| `action` | Nonblank Rex action name, or absent. |
| `callback_ref` | Opaque Rex-runtime-owned callback reference, or absent. |
| `args` | Optional UTF-8 JSON object, never null/array/function. |
| `hidden`, `confirm` | Booleans, default false. |
| `cond` | Optional evaluated boolean; absent means true. A nil Lua function result becomes false, not absent. Never Lua source. |

An actionable record has exactly one of `action` and `callback_ref`, a nonblank
`desc`, and no `group`. Group/prefix records have neither executable target and
must not carry `args` or `confirm`; the compiler discards those action-only
attributes from group records after applying inheritance to their child actions.
All text fields must be valid UTF-8 without terminal control characters. This
also prevents labels/icons from injecting escape sequences into rendering.

`ParseKeys` returns literal Unicode code-point input tokens, not UTF-8 bytes or
display columns. `<Space>` and `<lt>` are the only supported bracket forms and
are case-insensitive. Literal spaces and `<Space>` are the same key. Literal
less-than must be written `<lt>`. `NormalizeKeys` emits canonical `<Space>` and
`<lt>` spellings. For example, `w<SPACE>x` normalizes to `w<Space>x` and consumes
three input tokens. `<C-x>`, `<Tab>`, unclosed bracket notation, invalid UTF-8 and
literal control characters (including Esc, Backspace, Tab, newline and DEL) are
invalid. `q` is a normal mapping key; its dismissal fallback is a navigation rule.
This tokenizer does not claim grapheme-cluster or display-width measurement.

The Lua compiler preserves first-seen sibling order and creates implicit prefix
groups for full sequences. Explicit groups can decorate those prefixes. Repeated
metadata-only declarations may coalesce into the same node, with later supplied
text metadata taking precedence; two enabled actions at one normalized sequence
are always invalid, even when they name the same action. An action cannot also
be a group or a prefix of any other active entry. These conflicts are independent
of declaration order. Functions/conditions cannot be used to bypass final active
mapping validation.

Output at most one enabled compiled record per normalized sequence, coalescing
metadata in Lua before handing it to Go. Missing prefix records are implicit
submenus. A record without a target is a group, even if currently empty; its label
falls back from nonempty `group` to nonempty `desc` to the final key token's label
(`Space` for space, literal character otherwise). No depth limit is introduced.

`Config.Validate` and `ValidateMappings` implement data validation in Go. They
check structural data even for disabled records, but only enabled records claim
keys and create prefix conflicts. `Mapping.Enabled()` is the common eligibility
rule. Consumers may omit disabled entries entirely after validation. When a
condition function returns nil, the Lua evaluator must either omit the disabled
record or retain it with `Cond` pointing to false (`"cond": false`), never with
`Cond == nil` or an omitted `cond` field, which would enable it. Raw Lua validation
and coalescing must not depend on Go executing Lua. Go struct methods do not
validate unknown JSON fields or coercions performed by an external decoder;
a boundary decoder must reject unknown fields and invalid field types, including
null for required strings/booleans, rather than relying on Go zero values.

## Callback ownership, lifetime and dispatch

`callback_ref` is data, not a function, file path, chunk of Lua, action name or an
authorization capability by itself. Its spelling is opaque to Go. The originating
Rex Lua runtime retains the function and original arguments. No Go callback VM,
reflection mapper or JSON encoding can reproduce its closure or runtime context.
Only data-only evaluated values reach the menu.

The dispatch owner must bind references and selected mappings to the particular
invocation, originating server/client/session/block and current runtime/config
lifetime. This file does not specify a wire envelope, launcher arguments, registry
API or capability transport. Those are separate integration work, not invented
fields on user mappings. Independent invocations must not overwrite each other's
callback identity. Reload/replacement invalidates stale selections/references;
dismissal, completion and errors release invocation-owned references and resources.

Dispatch resolves the selected record from the retained originating invocation,
not from caller-supplied replacement code/action/args. Recheck conditions in Rex
Lua before execution; newly false/nil-returning or erroring conditions, stale
references, wrong origin, consumed selections or a lost originating runtime fail
closed. Never resolve a reference against another runtime or fall back to the
current local server/client/session. Hidden entries follow the same eligibility and safety rules.
Confirmation requires `y` in the menu and must be validated by dispatch before
executing a confirm-marked action; cancellation dispatches nothing. Restore origin
focus before dispatch. Invoke Lua callbacks as `callback(ctx, args)` there, with
origin context and retained args (or an empty object), never popup context. Do not
automatically retry an uncertain destructive operation.

## Ownership and verification limits

The consumers are the Lua config/compiler/evaluator/dispatcher, Go menu/popup and
built-in/default-menu code. They reference `mapping.go` and this document instead
of maintaining separate schemas. At this producer revision only these two files
are delivered; a documented consumer responsibility is not a claim that its
implementation or transport already exists. No tests, compilation, runtime Rex,
remote, reload or graphical experiments were run under current project authority.
Source formatting and diff checks do not establish runtime compatibility.
