# Mapping and configuration standard

`mapping.go`, in the root Go package `whichkey` (module `temp`), owns these data
values and validation rules. This document explains that source, not a parallel
Lua schema. It covers REQ-003/004/005/013 and the confirmation data used by REQ-007.
The post-return completion/cancellation data below implements accepted bridge
REXWHICHKEY-20261031-220348. No Lua interpreter, reflection mapper, dependency,
CLI invocation syntax or new public setup API is selected here.

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
actions. `rex_whichkey.defaults` supplies the default spec before resolving user
options, aggregating the delivered sessions, tabs, panes, application and known-hosts
fragments in that order. It returns the required resolved baseline table using
these Go-owned display defaults; no action runs or inventory is acquired on
import. No fabricated action or undefined future Go symbol is required.

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

Allowed attributes are `desc`, `group`, `icon`, `cond`, `hidden`, `confirm`, `args`
and `provider`. `provider` belongs only to a mapping, never an attribute scope;
it is not inherited. Each child inherits the nearest explicitly supplied value
of every other attribute; child values override that attribute only. In particular, `false` and `""` are
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
| `provider` | Absent/empty, or exactly `known-hosts` on a non-executable explicit group; see below. |
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

## Known-hosts provider and resolved defaults

`provider = "known-hosts"` is the one Go-owned provider value, not a generic
plugin mechanism or an action. In raw Lua, for example:

```lua
{ "h", group = "Rex-known hosts (listing only)", provider = "known-hosts" }
```

The marker must declare a group (including an explicitly empty group label) and
must not have an action, callback, args or true confirmation flag. Provider is a
literal string, not a function. Unknown values/types fail closed; empty is the
same as absence. An enabled marker reserves its entire subtree: any other enabled
mapping below it is an error, even if hidden or an informational group. Multiple
markers at disjoint prefixes share one inventory read. Disabled records still
require valid structure, but neither acquire inventory nor claim active keys.
Lua compilation/evaluation forward this canonical field and apply these rules;
they do not define another schema or create executable provider callbacks.

`RunTerminal` expands enabled markers before `NewMenu`, on both the direct menu
path and the popup child path. Only an explicit marker triggers one bounded
`rex.Hosts` call, with a three-second process timeout and the existing invocation
context/deadline; redraws do not refresh it. No marker means no inventory request
or inserted records, including an explicit empty or replaced user spec. The
popup child receives the original frozen `rex.CLISelection` and canonical endpoint
inside private `runtime.json`, verified before `RunTerminal` and passed through
its host-provider boundary into `rex.Hosts`. The endpoint remains part of the
selection's API receipt; it does not retarget the execution-host inventory.
Missing/malformed expected popup data fails closed, never selecting again from
the temporary child's environment/PATH. A standalone `menu` may make its own
initial selection only when an explicit provider marker needs inventory. The
provider replaces only each marker in source order, preserving its group label,
description, icon and visibility. Hidden is display-only: a hidden marker still
acquires inventory and opens its details when its key is typed; its generated
rows are informational groups, not inherited hidden bindings. The assembled
records are validated again; collisions are errors, never silent overwrites.

Inventory belongs to the execution host and its Rex-configured home, not to the
graphical client or to the origin endpoint selected with `--server`. Rex owns the
effective inventory merge, cached-home status, held changes and label resolution.
The menu preserves IDs separately from labels, structured warnings and opaque
home metadata, sorting host rows by ID (then owner) for stable duplicate-label
presentation. Empty inventory and warning lists are explicit. Endpoints are data,
not reachability results. No separate registry, endpoint probes, retargeting,
host mutation, attach, WOL or SSH behavior is added. The existing Rex inventory
command may itself contact home and forward previously held changes; the menu
does not implement or retry that behavior.

A failed, timed-out or invalid listing produces a persistent warning group, not
a successful empty inventory or a reachability claim. Raw CLI diagnostics are
not displayed. Invocation cancellation/expiry still aborts the menu. Every
generated row is non-executable and cannot produce a selection or bridge action.
The popup controller continues to validate selections against the original
retained mapping snapshot, and callbacks remain in the originating Lua VM.
Existing endpoint binding, origin checks, privacy, expiry and owned-only cleanup
are unchanged.

Opt-in initialization in the **user-owned** Rex configuration consumes the
resolved composer, with no mandatory socket in the Linux/macOS default case:

```lua
local popup = require("rex_whichkey")
local defaults = require("rex_whichkey.defaults")
popup.register_completion("whichkey_complete")
local menu = popup.new(defaults)
menu.setup() -- or menu.setup(inline_options, optional_user_module_name)
menu.register_launch("whichkey")
```

After one-time module/PATH/setup/registration and client Remote Control setup,
`rex-whichkey` in the originating Rex terminal activates the real floating popup;
standalone `menu` remains low-level resolved-JSON UI. The optional constructor
endpoint `popup.new(defaults, endpoint)` is immutable routing data outside mapping
options. Public `--server`/constructor override > nonempty execution-host
`REX_SERVER` > bounded no-autostart native Rex Linux/macOS discovery; Rex owns the
platform paths. Public and preparation endpoints must agree for custom/shared
servers. Never convert opaque `ctx.server` into an address or select a focused/
current client to compensate for missing/stale origin. See the
[launch contract](launch-contract.md) for endpoint validation, exact origin/client
selection, action-name overrides and actionable prerequisite failures.

This is an opt-in fragment, not an installed/reloaded user `init.lua`. Callbacks
remain in the original VM. Imported and inline `spec` overrides still replace the
whole default spec. No defaults are silently reinserted after setup; retaining
the listing in a custom spec requires retaining an explicit provider marker.
Mapping and retained-callback semantics remain unchanged. Code-owned startup now
requires the explicitly negotiated runtime receipt and completion version 2 below;
unsupported loaded Lua fails before its launch action is invoked. The Go-side
producer alone does not deliver the dependent Lua consumer or installed assets.

## Callback ownership, lifetime and dispatch

`callback_ref` is data, not a function, file path, chunk of Lua, action name or an
authorization capability by itself. Its spelling is opaque to Go. The originating
Rex Lua runtime retains the function and original arguments. No Go callback VM,
reflection mapper or JSON encoding can reproduce its closure or runtime context.
Only data-only evaluated values reach the menu.

The dispatch owner must bind references and selected mappings to the particular
invocation, originating server/client/session/block, client connection generation
and current runtime/config lifetime. `Completion` below is the authoritative
post-return wire data; launcher arguments, registration API and secure handoff
plumbing belong to the composition owner, not user mapping fields. Independent
invocations must not overwrite each other's callback identity. Reload/replacement
invalidates stale selections/references; dismissal, completion and errors release
invocation-owned references and resources.

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

## Canonical post-return bridge data

`Completion`, `CompletionSelection`, their JSON tags and validation in `mapping.go`
are the single shared standard for Go popup and server-Lua completion consumers.
They are internal integration data, not additional configuration attributes. The
original action may return after launch; a **fresh explicitly origin-targeted Rex
configuration action in the same server Lua VM** authenticates completion and
resolves the retained ticket. Do not require the original action's table identity
or a suspended original action. Conversely, copying context fields into a Lua
table does not change the ambient Rex action identity: composition must actually
target the captured server/session/block when invoking that fresh action. Its
actual caller client is distinct from the retained GUI target; app effects must
explicitly select that target. Never route through the popup, a current-client
default or a local-server fallback.

The strict legacy **version 1** JSON field sets are:

- Complete: required `version`, `capability`, `kind`, `selection`.
  `version` is the integer `1`; `kind` is the string `"complete"`.
- Cancel: required `version`, `capability`, `kind` only.
  `version` is `1`; `kind` is `"cancel"`. `selection` must be absent, not null.
- Selection: required string `keys` and boolean `confirmed`; optional string
  `callback_ref`. `keys` must already equal `NormalizeKeys(keys)`, not merely
  normalize to the selected sequence. `callback_ref` must be absent for a named
  action and must exactly match the retained opaque reference for a callback.
  An explicitly empty callback reference is invalid; omit it instead.
  `confirmed` must be present even when false; no strings, numbers or nulls.
- Selection may additionally contain string `text` only for inline rename:
  canonical `tr` + retained `client.tab.rename`, or `wr` + retained
  `session.rename`, without callback, confirmation or nonempty retained args.
  Those mappings require text; all other mappings reject it. Names must be
  nonblank UTF-8 without controls, at most 4096 transport bytes and 128 Unicode
  characters (the documented Rex session/window label limit). Text provides only a label,
  never a resource ID, action or replacement args. The original Lua registry
  retains and checks the tab/session identities and invokes the resource API
  after popup cleanup. Custom rename-key callbacks are not intercepted.
  Older decoders reject this optional field. Runtime negotiation must establish
  `selection-text-resource-rename` before showing an inline rename; build IDs need
  not match, and missing support never becomes a dialog or discarded text.

Complete example (the all-zero capability is illustrative, never generate it):

```json
{"version":1,"capability":"0000000000000000000000000000000000000000000000000000000000000000","kind":"complete","selection":{"keys":"wn","confirmed":false}}
```

Cancellation has no selection or confirmation:

```json
{"version":1,"capability":"0000000000000000000000000000000000000000000000000000000000000000","kind":"cancel"}
```

No action, args, Lua source, origin identifiers, expiry, confirmation requirement
or configuration/connection generation may be supplied as replacement authority
in these messages. Unknown fields, case aliases, duplicate fields (including
escaped equivalents), nulls, missing required fields, invalid UTF-8, unpaired
UTF-16 surrogate escapes, wrong types, unsupported versions/kinds and trailing
JSON values are invalid. Object field order, whitespace and valid JSON string
escaping are immaterial. Supplementary Unicode characters encoded as valid
surrogate pairs are supported. These are data encodings, never executable Lua.
Lua must enforce the same rules on raw JSON before a permissive decoder could
collapse duplicate fields or coerce nulls. Do not stringify unexpected values.

For Go, use `json.Unmarshal(data, &completion)` on one complete input value;
`Completion.UnmarshalJSON` enforces the boundary rules and leaves the receiver
unchanged on failure. If using a streaming decoder, independently require EOF
after its one message and bound transport input allocation. `json.Marshal` invokes
`Completion.MarshalJSON`, which validates structure and emits only the fields
above. Neither operation authenticates a caller. Direct Go struct construction
also requires `Validate()` or `ValidateSelection(snapshot)` before use. Do not
independently decode an untrusted selection into `CompletionSelection` and assume
its zero values prove required-field presence.

`Completion.ValidateSelection(snapshot)` checks the complete message against
validated, retained evaluated mappings: the selected normalized sequence must
identify an enabled actionable record, the callback reference must match exactly,
and `confirmed` must equal the retained `Mapping.Confirm`. Thus a confirm-marked
selection requires true and an ordinary action requires false. Hidden actions are
eligible; groups, missing/disabled records, mismatched references and spurious
confirmation are rejected. A structurally valid cancel needs no mapping and
selects nothing. The popup builds selection identity from `Selection.Mapping()`
and confirmation from `Selection.Confirmed()`, never copies that mapping's action
or args into a completion. The Lua dispatcher performs these checks independently
against its own snapshot and rechecks original condition guards before execution.
A boolean is an authenticated popup assertion of y, not cryptographic evidence of
a physical keypress or a way to change the retained confirmation requirement.

### Negotiated runtime receipt and completion version 2

`runtime_bootstrap.go` owns the private runtime formats; `mapping.go` owns
Completion. These are internal integration messages, not mapping options or
executable constructor parameters. The protocol is 1, normal asset layout is 1,
and the configured completion wire is **2**. Version 1 remains strict: supplying
`receipt_hash` with version 1 is invalid. Version 2 requires exactly `version`,
`capability`, `receipt_hash`, `kind`, plus `selection` only for `complete`.
Selection's strict keys/ref/confirmation/text semantics above are unchanged.
`receipt_hash` is 64 lowercase hexadecimal SHA256 characters over the **exact
receipt string bytes** returned by Go. Lua retains those bytes/hash; it must not
serialize its own competing receipt or renew the receipt's original expiry.
File identity timestamps such as `ModifiedNS` exceed Lua 5.1's exact integer
range (2^53): retain/pass the opaque Go receipt string, never decode/re-encode
it through Lua numbers. The preflight `relay` identity is likewise an opaque,
bounded Go JSON **string**, including all timestamps/device/inode integers. Retain
and return that exact string in both relay operations; never decode/re-encode it.
Go strictly decodes it and compares both exact bytes and the fresh full identity.
Parsed copies are inspection data, not receipt/relay bytes.

The loaded `whichkey_runtime` action is read-only. Its request is
`{protocol_min:1, protocol_max:1, action:NAME}` on `rex do ... --args - --json`.
The response requires all of `protocol_min`, `protocol_max`, `consumer_build`,
`features`, `launch_action`, `completion_action`, `menu_generation`,
`server_instance`, `config_generation`, `config_hash`, and `layout_revision`.
Names use lowercase letters/digits/underscores. Strings are nonblank,
control-free and at most 1024 bytes (action names at most 256). Ranges must
contain protocol 1; features must be unique and understood. Required features:
`original-vm`, `separate-caller-target`, `single-use`, `generation-bound`,
`retained-menu`, `public-file-relay`, `completion-receipt-v2`.
`selection-text-resource-rename` is additionally mandatory if the retained
resolved menu has a built-in inline rename. The consumer must advertise its
actual retained-template support before reservation. Build IDs are audit data,
not equality locks. Missing/old advertisement fails before calling an old launch
action; registration alone never authorizes an installed popup helper.

All commands below take **only captured stdin** and emit captured stdout (never
terminal output). Internal records are strict JSON objects: unknown/duplicate/
case-alias fields, nulls, coercions, bad Unicode and trailing data fail. Optional
fields are omitted, not null. CLI/API descriptors and version JSON inside the
Go-owned receipt remain JSON **strings**, preserving their exact bytes.

The Go runtime boundary checks decoded object names before typed decoding. Every
protocol struct tag uses lowercase ASCII letters, digits and underscores;
restricting protocol names to that alphabet makes `DisallowUnknownFields` an
exact-name check, rather than accepting `encoding/json`'s Unicode SimpleFold
aliases. Exact decoded duplicates (including escaped equivalents) and nulls are
rejected recursively, including slot payloads and nested CLI/file identities.
The popup handoff uses this same boundary, not an independent permissive decoder.
Runtime-carried Config adds the same name/Unicode checks before the existing
Config decoder, retaining its required fields, types, validation and 4 MiB bound.
Only `args` and its descendants are arbitrary case-sensitive data maps: mixed
case, Unicode and empty keys remain allowed there, including distinct keys that
would fold together as struct fields. String values are not restricted to ASCII.
Receipt, relay and embedded JSON strings are never normalized or reserialized;
their consumers check names when decoding their contents. The legacy standalone
`DecodeConfig` and user Lua mapping semantics are unchanged.

Source-review cases for this boundary (all **UNEXECUTED**, not parser fixtures):

| Input variation, otherwise valid | Required source outcome |
| --- | --- |
| Canonical `version:1`, or `ver\u0073ion:1`, with each required field once | Decode the same exact field name; ordinary type/version/authority checks still apply. Canonical `version:2` remains unsupported where protocol 1 is required. |
| `Version:1` or `ver\u017fion:1`, alone or beside `version:2` in either order | Reject before struct decoding; no alias can replace or conceal the unsupported version. |
| Repeated `version`, including `version` plus `ver\u0073ion` | Reject the duplicate decoded name, independent of order/value. |
| Long-s alias in `status`, `workspace`, nested file `size`, or a slot payload's field; ASCII unknown field | Reject across acknowledgment/proceed, preparation/release, describe/relay/locator, receipt/envelope and PTY-image consumers, not just the locator. Lowercase ASCII unknown names still fail typed decoding. |
| Runtime Config uses `\u017fpec` beside or instead of `spec`, or a Mapping uses `KeyS`/`key\u017f` | Reject protocol aliases before Config/Mapping struct decoding; no exemption outside `args`. |
| Mapping `args` contains distinct `Version`, `version`, `verſion`, `K`, `k`, `é` and empty keys, including inside nested objects/arrays | Preserve those legitimate distinct map keys and their Unicode values; only an exact decoded duplicate, null or otherwise invalid data fails existing data rules. |
| A valid Unicode title/label, supplementary surrogate pair, or opaque receipt/relay with integers above 2^53 | Preserve the original values/opaque bytes; no blanket Unicode rejection or Lua numeric round trip. |
| Invalid UTF-8, unpaired surrogate, null, wrong type, unknown field, trailing value, missing Config marker or oversized record | Existing strict/bounded rejection remains; no startup/dispatch or nominated cleanup authority follows from malformed input. |

**Key/palette consumer handoff**:

1. Run the fixed anchor sibling `runtime-describe` with `{anchor,
   layout_revision:1, consumer, origin}` and optional present `server` override.
   `anchor` is the validated logical file-loader path without `@`, relative only
   under the unchanged standard-loader/server-cwd contract. `origin` has required
   `session_id`, `block_id`, `window_id`, `layer_id` (empty only if absent),
   `caller_id` (actual action caller), `client_id` (distinct retained GUI target),
   and `connected_at`. IDs are full typed UUIDs; never substitute the GUI target
   for the actual caller. Compare original-VM metadata before/after every yield.
2. Describe returns `{protocol_min, protocol_max, producer_build, features,
   layout_revision, receipt, receipt_hash}`. It issues no capability. Go freezes
   the current helper image, normal layout, shared `rex.CLISelection`, required
   method evidence, canonical endpoint, server/config/menu generations, exact
   placement/GUI generation, platform, embedded format digest and original expiry.
3. After original-VM issue/evaluation, run the same fixed helper's `prepare` with
   `{version:1, configuration:CONFIG_JSON_STRING, receipt, receipt_hash}`.
   Its optional `--server` must agree with the receipt, never rediscover transport.
   Go rechecks the selected receipt, validates Config/rename support, copies its
   captured open source image into a new private workspace and generates a
   capability. Output is `{version:1, capability, executable, server, receipt,
   receipt_hash, workspace, expires}`. This final immutable receipt additionally
   binds the workspace and exact configuration bytes; retain the **returned**
   final bytes/hash. It does not change the selected source, origin or expiry.
4. Bind the original ticket to capability plus final receipt/hash and retained
   exact resources. Start the returned selected copy as `controller` with
   `rex.exec.start`, `detach=true`, captured stdout, suppressed stderr and
   finite private stdin `{configuration, capability, server, session_id,
   block_id, client_id, connected_at, completion_action, expires, receipt,
   receipt_hash, workspace}`. `controller` is a finite Go launcher, not the
   interactive child: use supported `h:wait(35)` and require a successful exit
   plus exactly `{version:1,status:"accepted"}` on captured stdout. A timeout,
   nonzero exit, malformed/missing acknowledgment or lost result is failed or
   uncertain startup, never permission to retry/remove its workspace. Do not
   return the handle/PID or stdout from the configuration action. Use returned
   `expires` unchanged and a 35-second command bound. No callback, ticket,
   closure or opaque `ctx.server` is serialized.
5. Before any attempted controller start, a bind/preparation failure releases the
   ticket and calls the fixed helper's `runtime-release` with the entire prepared
   output. After a possibly successful start, never retry, release its workspace
   or claim failure proves no controller exists. The controller owns its workspace
   after a valid handoff; original-VM expiry still revokes unclaimed authority.

**Public consumer handoff** (no caller executable is ever executed by Lua):

- Public Go selects itself and the CLI before endpoint discovery, queries the
  loaded advertisement, and activates the chosen original action **once**, with
  `{runtime:LOCATOR_JSON_STRING}` on stdin. The locator has `version:1`,
  `workspace`, `request_sha256`, `receipt_hash`; no capability. A public native
  launch is allowed only for a truly absent implicit default and the originating
  execution-host terminal endpoint. An explicit/custom/registered menu never
  falls back to native.
- Before issue/evaluation, the original action runs only its fixed anchored
  helper's `runtime-describe`, adding `descriptor:LOCATOR_JSON_STRING` to its
  describe input and omitting `server`. This path reuses the public frozen CLI;
  it does not choose another CLI/helper. Output is `{version:1, protocol_min,
  protocol_max, producer_build, features, layout_revision, receipt_hash, relay}`.
  Validate producer features/protocol and retain the exact opaque `relay`
  installation-identity string before issuing a ticket. Its Go-owned numeric
  fields must never pass through Lua number reserialization. The read-only
  `runtime-relay` operation `preflight` is the same finite preflight; it is not
  a popup or authority probe.
- After one original-VM evaluation, invoke fixed `runtime-relay` with
  `{operation:"menu", descriptor, anchor, layout_revision:1, consumer, origin,
  relay, configuration:CONFIG_JSON_STRING}`. It writes `menu.json`, waits for
  Go-owned public preparation, and returns the same preparation shape as the key
  path. Bind the ticket to those final receipt bytes/hash. Do **not** start the
  returned executable from Lua for this entry path.
- Then invoke fixed `runtime-relay` with `{operation:"bound", descriptor, anchor,
  layout_revision:1, consumer, origin, relay, receipt_hash:FINAL_HASH}` (omit
  configuration). It publishes binding and returns `{version:1}`. The action
  returns no secret/handle/result. Only after matching bound state **and certain
  activation success** does public Go retire its finite rendezvous root/FIFO
  handles and captured source handle, then start its selected controller. A
  retirement failure prevents startup and takes owned cleanup/cancellation;
  uncertain closes are never repeated. The controller uses its verified private
  image and finite stdin, not the public relay lease, so normal caller return
  does not revoke it. After any possibly successful start, workspace ownership
  stays with the controller, even if the finite handoff/acknowledgment fails. The
  public caller requires the private child startup acknowledgment before success;
  pipe acceptance or `Process.Release` alone is not acknowledgment. On
  failure before start, revoke what is still unclaimed; never repeat activation
  or a relay transition.
- Public Go precreates 0700 workspace / 0600 one-link no-follow regular
  `request.json`, `menu.json`, `prepared.json`, `bound.json`, plus mode-0600 FIFO
  `lease`. Only the public owner retains a write-capable lease endpoint. Relay
  EOF means owner loss, never success. Every fixed slot has
  `{stage, receipt_hash, payload}`; stage matches the basename without `.json`.
  Slot hashes bind the initial request receipt, while prepared/bound payloads
  carry the distinct final receipt hash including actual activation caller,
  relay identity, workspace and Config digest. Request SHA binds the fresh local
  marker and inode manifest. Writers lock exclusively and publish once without
  create/truncate; readers use shared locks. Reopened files recheck ownership,
  private modes, one-link identity and stage. No valid shared local state means
  unsupported cross-host launch **before** reservation, not a remote path guess.

Small advertisements, locators, preflight and completion are bounded to 64 KiB;
full Config/receipt/preparation/handoff and slot envelopes to 8 MiB total at each
boundary (including string escaping). Stages/relay commands are three seconds
within the original 30-second configured lifetime, with five seconds reserved
for teardown; no step renews expiry. Native interaction retains its lifetime
pipe and no idle deadline. Controller/PTY use verified private executable copies;
Rex's proprietary dependency is never copied. Sources are captured as open files
and hashed, copied from those handles and rechecked before execution. This assumes
trusted, stable installer/user-owned components: neither `os.Executable()` nor
path/content rechecks promise hostile-same-account isolation or atomic
execute-by-descriptor on every platform.

Both direct public spawn and the key's finite Go launcher use one Go-owned
handshake: bounded finite stdin, a sole inherited startup writer (fd 3), and strict
`{version:1,status:"accepted"|"rejected"}` followed by EOF. The child consumes
all input through EOF, closes stdin, authenticates receipt/image/CLI/endpoint,
checks the loaded consumer and exact origin, prepares rename presentation, then
rechecks origin/menu/expiry before accepting. The child closes/checks the startup
writer **before** popup effects; failed acknowledgment prevents its lifecycle.
The parent closes its duplicate writer, waits for the complete bounded reply and
EOF, closes its reader and releases the process. The key launcher also writes and
closes its captured accepted stdout at this barrier. Only then does Go send
`{version:1,status:"proceed"}` plus EOF over a separate finite fd-4 pipe. The
child requires that exact positive marker/EOF and successfully closes its reader
before entering the lifecycle. EOF without the marker is cancellation, not
permission to proceed; a failed parent acknowledgment/output close suppresses
that marker. No startup endpoint remains live during popup effects.
Three-second startup I/O/check bounds and caller cancellation never renew the original expiry. Inherited pipes
are converted to owned nonblocking/pollable descriptors so Darwin/Linux deadlines
work; all duplicates are retired. `Start` transfers ownership irrevocably: failure,
cancellation, rejected/lost acknowledgment, deadline or release/close error after
that point must not retry, kill an uncertain child or delete its workspace.
Before `Start`, only authenticated owned cleanup/cancellation is permitted. A
complete expired key handoff still rejects startup: the finite launcher uses the
same receipt/copy/self-image authentication as the child to adopt only its owned
workspace before the live-authority checks reject expiry. Its bounded failure
path retires the captured image and removes only fixed owned workspace entries;
it does not require a live popup, origin or callback to authenticate that local
cleanup, and sends no completion after expiry. Cleanup errors remain failures,
and uncertain closes are never repeated. Malformed/partial input or failed
receipt/copy/self-image authentication cannot authorize deletion of nominated
paths; such input may leave preparation artifacts. A possibly started detached
child remains the workspace owner even if its acknowledgment is lost. Lua must
not infer permission to release from a failed finite launcher result.
The key launcher returns success only after this handshake (an accepted stdout
object alone without successful launcher exit is insufficient). Neither
that object nor public success proves visible popup display or later dispatch.

Source-review scenarios for this ownership boundary (all **UNEXECUTED**):

| Input / transition | Required source outcome |
| --- | --- |
| Complete authentic key handoff expires before finite launcher decoding; its private copy and receipt still match | `decodePopupRequest` retains the request; `validateRuntime` adopts its verified image before rejecting expiry; launcher closes input once, retires the image and removes the exact owned workspace before returning failure. No child start, accepted stdout, popup or completion action. |
| Truncated/malformed handoff, mismatched receipt/hash/configuration/workspace, or foreign self-image/copy | No authenticated image is adopted, so launcher performs no nominated workspace deletion, child start or completion; any retained preparation is an explicit limitation, not successful cleanup. |
| Child `Start` succeeded but acknowledgment/launcher result is lost, rejected or late | Transferred ownership suppresses launcher workspace removal even on failure. Neither Go nor Lua retries the spawn or releases the possibly running child's workspace; expiry revokes registry authority without claiming artifact removal. |
| Authentic expired pre-child handoff encounters an image-close or owned-file removal failure | Return startup failure joined with the cleanup error; do not retry an uncertain close, report cleanup success or send stale cancellation/completion. |

Controller compatibility/placement/rename checks precede popup creation. After
rename and geometry yields, read-only `whichkey_runtime` revalidates the exact
retained loaded menu/features/generation immediately before layer creation,
during the joined lifecycle watcher and before completion/cancellation. Changed
or unavailable state suppresses startup/selected effects with owned cleanup, no
fallback/reload/retry. Each advertisement/completion activation freezes its own
actual caller; the old activation caller need not remain connected or equal a
fresh caller, and neither is the retained GUI target. These read-only checks are
not atomic locks against external setup/resource races.

Cleanup precedes fresh completion. The native route also retires its captured
source image after popup return and before any selected effect; close failure
suppresses dispatch, and fallback cleanup never repeats an uncertain close.
The **actual fresh CLI caller** is not the retained GUI or original activation
caller. Lua strictly decodes/authenticates before
lookup/claim, checks capability, final exact receipt/hash, expiry and retained
state, then consumes **before** any yielding guard/API/callback. A failed/uncertain
claim/effect/acknowledgment is not retryable. Default `tr`/`wr` text targets only
the retained window/session APIs, with the documented 128-character label bound;
custom callbacks/args/confirmation retain ordinary registry semantics.

Diagnostics are fixed by layout, CLI/endpoint, compatibility, origin, rendezvous,
preparation, startup, cleanup and dispatch stage. Private output is suppressed.
Obsolete `execution.helper` / `execution.rex_bin_dir` wiring is not a fallback.
The inspected pre-Phase-15 `cmd/main.go` rejects unknown positional modes at
`flag.NArg()` before `LaunchPopup`; that source establishes a safe unknown-mode
failure for that baseline, **not** for arbitrary historical helper versions.
Unsupported installed/loaded consumers must be updated through the normal user
lifecycle, never auto-reloaded or asked to pair equal build IDs.

This is a Go/source handoff, not phase/runtime verification. Direct versus
packaged startup, compatible unequal builds, missing text features, malformed
receipts, replay/expiry, cancellation at each transition, wrong/moved origins,
source upgrades, FIFO owner death, uncertain activation/start/dispatch, custom
callbacks and exact-resource rename remain **UNEXECUTED** scenarios. No tests,
Lua harnesses, product/API/menu/rename probes, remote trials or consumer reloads
were performed by this Go task.

### Capability and trusted retained state

`NewInvocationCapability()` uses Go `crypto/rand` for 32 random bytes and encodes
exactly 64 lowercase ASCII hex characters (256 bits). Abort invocation preparation
on RNG failure; never substitute a counter, timestamp, callback reference, Lua
`math.random`, empty token or other fallback. `ValidateInvocationCapability` checks
only that canonical spelling, not entropy or authorization.
`InvocationCapabilityMatches(presented, expected)` validates both spellings and
uses a constant-time comparison. Its expected value must come from trusted state,
not the incoming message. Do not case-fold or otherwise normalize capabilities.

The composition owner must generate the capability with this Go helper during
trusted invocation preparation, before exposing the popup, and associate it with
the issuing Lua registry entry. The helper is Go source, not a claim that Lua can
call it without process plumbing. The original Lua VM keeps the opaque ticket,
callbacks, original args and guards. The popup receives the capability, data-only
evaluated configuration and explicit origin routing data through a restricted
handoff. Routing data tells Go where to send completion; it is not replacement
authority for the registry's independently captured origin. No ticket, closure or
VM identity is serialized. Callback references are identifiers, not secrets or
standalone bearer authority. Keep capability-bearing handoffs/messages out of
logs, terminal output, diagnostics, filenames and publicly readable files. The
composition owner must restrict transport access and release temporary handoff
resources on every exit path; this producer does not implement that transport.

One live registry entry binds the capability to ALL of:

- The exact originating server/runtime identity and session/block IDs, plus an
  independently validated GUI target client ID. Separate CLI activation callers
  are not required to equal that app. Window placement is not identity: a block
  may move within its session.
- The captured client **connection generation**, not just a reused client ID
  (for the delivered Rex consumer this is the server-provided `connected_at`
  connection value). Missing, lost, detached, ambiguous or changed connection
  information fails closed; do not adopt the new generation on completion.
- The installed configuration generation and retained invocation snapshot,
  including callback references, original functions/args, condition guards and
  confirmation requirements. Never rebuild a snapshot from current configuration
  or trust an incoming mapping as the retained one.
- Reservation time, expiry and single-use state. `InvocationLifetimeSeconds` is
  30 elapsed wall-clock seconds, starting before evaluation/launch; completion,
  redraws and validation never renew it. At `now >= expires` it is expired.
  `MaxPendingInvocations` is 64 per installed registry, including reserved and
  in-progress entries. Reject a new reservation at capacity, never evict another
  invocation to make room. A backwards wall-clock jump invalidates the registry;
  unavailable/invalid time fails closed. Lua may expire lazily on each registry
  operation because its configuration runtime has no scheduled timer here.

Neither the capability string nor Go's `Validate`/`ValidateSelection` enforces
these bindings, elapsed time or replay protection. Those checks and resource
ownership are the server-Lua registry/composition consumers' responsibility.
Reserve entries before any yielding evaluation/launch work. On successful setup
replacement/reload/teardown, invalidate the old generation, including pending and
in-progress entries; failed candidate setup leaves the working generation intact.

### Completion, cancellation and races

For either kind, strictly decode first, look up the live capability and compare
the fresh action's server/session/block with the captured origin. Its actual
client ID is a fresh caller, not necessarily the retained GUI target; freeze that
caller within the activation. Claim its open state
**before** condition evaluation or any yielding call, then validate the current
client connection/config generation and expiry. Reentrant, expired, cancelled,
consumed, wrong-origin or stale-generation requests cannot execute. An
unknown/malformed capability must not consume an unrelated entry.
Once an authenticated entry is claimed, a later selection/check/API error still
consumes it; never reopen or retry it. Completion and cancellation race for the
same one-shot claim. Cancellation only releases the retained entry: no selected
action, condition/callback execution or origin focus change.

For completion, resolve identity and confirmation solely against the retained
snapshot, recheck condition guards, inspect the originating session's current
placement and exact client connection, then focus the stable origin block in that
session. Execute only in the fresh action genuinely targeted at that origin,
using retained action/args or `callback(ctx, args)` in the original Lua VM.
Dotted client actions use explicit `client.perform` targeting the retained GUI,
not a queue on the transient CLI. Named configuration actions use `rex.invoke`.
Callback `ctx.client_id` is the GUI routing target; `ctx.request_client_id` is
the actual completion caller. Neither field fabricates an ambient activation.
This refines the old app/caller equality requirement under accepted decision
REXWHICHKEY-20261104-172117; see [CLI/GUI delegation](cli-gui-delegation.md). Recheck
expiry, generation, registry state and origin after every yielding boundary and
immediately before side effects. If a configuration change invalidates a ticket
while evaluation returns a new invocation handle, release that handle as well.
Never fall back to another server/session/block/client or restart missing Lua
state. Release references on all success, cancellation, expiry and failure paths.

Rex calls are not an atomic inspect/focus/execute transaction; concurrent external
moves/disconnections may race execution after the last check. Do not describe the
bridge as an atomic transaction or a runtime-verified guarantee. Explicit targeting,
connection binding, fail-closed guards and no retry remain required despite that
API limitation. Transport/API invocation and registry integration belong to
03-04/04-05; defining these Go values does not deliver those consumers.

## Ownership and verification limits

The consumers are the Lua config/compiler/evaluator/dispatcher, Go menu/popup and
built-in/default-menu code. They reference `mapping.go` and this document instead
of maintaining separate schemas. The provider composition described above is
implemented in `hosts.go`, `terminal.go` and the Lua mapping/evaluation/defaults
consumers. Other consumer
responsibilities remain with their respective owners; documenting them is not
runtime verification. No tests, compilation, runtime Rex, remote, reload or
graphical experiments were run for this provider revision under current project
authority. Source formatting and diff checks do not establish runtime compatibility.
