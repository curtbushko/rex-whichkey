# Data menu configuration (version 1)

## Implementation status

`LoadMenuConfig(explicit *string) (Config, error)` implements this public data
contract in the flat Go package. It reads files and returns an independent,
validated configuration; it never invokes Rex, queries inventories, evaluates
Lua, writes configuration or expands providers.

The public entrypoint uses this loader, expands enabled providers once before
popup sizing/PTY handoff, and dispatches retained selections through Go's Rex CLI
resources after cleanup. It never queries Lua runtime/launch registrations.
Source wiring and static builds do not by themselves demonstrate a running menu;
coordinator observation of the public candidate is still required. Legacy
`DecodeConfig`, callback wire validation and `launchOutput` remain isolated for
unchanged source tests, not as an alternative public configuration/control path.

## File selection

Go starts with the current built-ins, then applies exactly one selected JSON
options file. In priority order:

1. A present `--config PATH` (represented by non-nil `explicit`). An empty path is
   an error. Relative paths are relative to the process working directory.
2. Nonempty `REX_WHICHKEY_CONFIG` (also permits a relative path).
3. `$XDG_CONFIG_HOME/rex-whichkey/config.json` when `XDG_CONFIG_HOME` is set.
   It must be an absolute, nonempty directory path; an explicitly empty or
   relative value is an error.
4. `$HOME/.config/rex-whichkey/config.json` when `XDG_CONFIG_HOME` is absent.

Unused lower-priority sources are not validated or read. Only a missing implicit
file (3/4) means built-ins. A dangling leaf symlink is an error, as are missing
explicit/environment files, unreadable files, malformed data and unknown versions.
An unavailable home directory is an error. No file is created automatically.

Normal user/Nix symlinks to regular files are supported. The opened descriptor
must be regular, and at most 4 MiB is accepted (including whitespace). Directories,
FIFOs and devices are rejected without reading them; opening uses nonblocking
mode so FIFO replacement does not wait for a writer. This is a normal data-file
boundary, not a private transport: no ownership, private-mode or no-symlink rule
is imposed. It does not promise isolation from hostile same-account file edits.

## Root object and replacement semantics

Required member: `"version": 1`, exactly the integer scalar `1`, not `1.0`, `1e0`,
`"1"`, a boolean or null. Optional members:

| Member | Type | Meaning |
| --- | --- | --- |
| `position` | string | `top`, `bottom`, `left`, `right`, `center` |
| `title` | string | Menu title; `""` suppresses it |
| `icon` | string | Menu icon; `""` suppresses it |
| `icons` | object | Only `breadcrumb` and `group`, both strings |
| `spec` | array | Ordered flat full-sequence mappings |

Omission retains the corresponding built-in value. A present `icons` replaces
**both** markers: missing members within that object become empty strings.
A present `spec` replaces the **whole** binding list, not individual keys;
`"spec": []` deliberately clears it. Explicit empty strings and false booleans
are data, never replaced by truthy/default merging. An empty position remains an
invalid position. `{"version":1}` preserves all built-ins.

Unknown or differently cased structural members are rejected by exact name,
including Unicode case-fold lookalikes. Duplicate members (including escaped
spellings of the same name), null anywhere, malformed UTF-8, unpaired escaped
UTF-16 surrogates, trailing JSON and nesting beyond 64 levels are rejected.
Legitimate Unicode, including paired surrogate escapes, is allowed. Display text
cannot contain terminal controls. JSON argument/data keys are case-sensitive;
no generic lowercasing changes arbitrary data. Built-in argument objects have
the narrow allowlists below, not an arbitrary Rex argument escape hatch.

## Mapping object

| Member | Type | Meaning |
| --- | --- | --- |
| `keys` | nonempty string | Required complete sequence, e.g. `wn` |
| `desc` | string | Action description; nonblank for actions |
| `group` | string | Explicit group, including `""`; mutually exclusive with action |
| `icon` | string | Row icon |
| `action` | string | One supported Go action/intent below |
| `args` | object | Only the action-specific members below |
| `hidden` | boolean | Hide from display **without disabling dispatch** |
| `confirm` | boolean | Require the menu's explicit `y` confirmation |
| `provider` | string | Named non-executable dynamic group below |
| `cond` | boolean | False disables display/dispatch; omission means enabled |

These are the only mapping fields. `callback_ref`, hooks, functions, executable
paths/commands, nested executable specs and custom Lua action names are not
supported. Conditions, labels and icons are literal values, not expressions.
Only actions may carry args or confirmation. A plain group/metadata row is not
executable. Providers additionally require `group` and forbid action, args and
confirmation. Flags apply to that record; the flat file has no implicit Lua
metadata inheritance/compiler.

`ParseKeys` splits literal Unicode tokens; only `<Space>` and `<lt>` notation is
supported (notation is case-insensitive). Esc, Backspace, Tab and other controls
are reserved. Missing prefixes are implicit groups. First-seen sibling order
follows file order. Duplicate normalized sequences and action/submenu conflicts
are errors **even for disabled records**. Providers reserve every descendant
sequence, including against disabled static children. `session.switch` also
reserves its subtree because it is a Go sessions-menu intent.

Conflict validation normalizes each full key once per pass, sorts complete
canonical paths and checks adjacent strict prefixes. It never builds or hashes
every growing key prefix, so a single long key takes linear key-processing work
rather than quadratic prefix hashing. The same validator runs from the first
configuration resolution and for disabled public records and generated children.
Sorting does not reorder the returned spec or diagnostics: duplicate/invalid
records and action/provider conflicts are still reported in source order.
No additional key-length or token-count limit is imposed.

## Complete public action catalogue

`args` may be omitted, meaning `{}`. Unknown argument members, wrong types and
nulls are errors. There is no configurable client/server/executable nomination.

| Action | Accepted args | Dispatch intent |
| --- | --- | --- |
| `session.new` | `{}` | New session |
| `session.previous` | `{}` | Previous session |
| `session.next` | `{}` | Next session |
| `session.switch` | `{}` | Open Go-owned sessions list; never native/Lua picker |
| `session.rename` | `{}` | Inline text entry; no preconfigured name/dialog args |
| `session.close` | `{}` | Close retained origin session; **confirm must be true** |
| `session.select` | required `session_id`, optional `window_id`, both strings | Select exact destination session/window |
| `client.tab.new` | `{}` | Create a Window in the current Session, then exact activation |
| `client.tab.previous` | `{}` | Previous tab |
| `client.tab.next` | `{}` | Next tab |
| `client.tab.rename` | `{}` | Inline text entry for retained Window |
| `client.tab.close` | `{}` | Close retained Window; **confirm must be true** |
| `pane.split` | required `direction`: `"right"` or `"down"` | Split in that direction |
| `pane.close` | `{}` | Close retained Block; **confirm must be true** |
| `client.sidebar.toggle` | `{}` | Toggle sidebar |
| `client.palette.toggle` | `{}` | Toggle command palette |

`session.select` IDs must be full typed `session:UUID` / `window:UUID`, with
36-character UUID bodies (hex digits and canonical hyphen positions), never
labels, prefixes or indices. The loader checks shape without querying Rex.
The consumer must check current destination existence and that an optional
Window belongs to that exact Session before any selected effect. Configuration
cannot choose a replacement origin/app or connection generation.

Destructive actions without `confirm:true` are rejected, not silently dispatched
or silently rewritten. Their retained confirmation is rechecked by the dispatch
owner. Other executable actions may request confirmation too; `session.switch`
cannot, because opening a list is not an executable selection.

## Complete provider catalogue and expansion boundary

Markers take no args/options, action or confirm; all require `group`.

| Provider | Generated meaning and retained identity |
| --- | --- |
| `sessions` | Select retained Session ID via `session.select` |
| `windows` | Select retained Window ID plus its source Session ID via exact `session.select` |
| `panes` | Focus retained Block through the Go resource method after placement checks |
| `known-hosts` | Execution-host Rex inventory **listing only**; no connect/SSH/action |

All four are opt-in. The only existing default dynamic intent is `ws`, adapted
by the consumer from `session.switch` to the sessions list. The loader leaves the
current built-in record intact. It adds no other default provider row.

The Go parent expands enabled markers once before sizing/PTY handoff, preserving
source order and hidden display flags. Sessions use one bounded session snapshot;
windows/panes list the origin Session's resources. Each resource provider rejects
more than 512 resources, without truncation; all batches and the combined transport
remain bounded and validated. Full IDs stay in retained targets and displayed rows.
Stale/moved resources fail rather than retargeting by label, focus or position.
Generated keys identify retained choices, not resource authority. Known-host
details remain non-executable, including when inventory is unavailable; remote
connect/management is not implied.

`ValidateMenuProviderChildren(marker, children)` is the structural batch contract:

- The marker must be a valid enabled provider (or `session.switch` intent).
- At most 1024 rows, including pagination/groups, and 4 MiB of total key bytes.
- Every child is a strict normalized descendant of that marker, never the
  marker itself or another subtree. Children cannot add providers/callbacks.
- Duplicate and action/prefix conflicts are rejected, including disabled rows.
- Known-host rows cannot have actions.

The helper validates generated structure, **not** target authority or arbitrary
internal action dispatch. Public file actions still use the closed catalogue;
Go-only generated pane intents are not public action names. The consumer must
also validate the combined expanded menu for cross-batch collisions, retain
resource-specific action data itself, and perform freshness/confirmation checks.
Provider queries, expansion, pagination and dispatch belong to the Go launch
consumer, not the file loader. Supported notation is canonicalized in one retained
snapshot before expansion, so case-insensitive key notation cannot break dispatch.
Inline rename follows retained action intent at any configured key, with optional
confirmation and validated private text; it never opens a native rename dialog.

## Built-ins and migration

The current built-in menu is preserved exactly: Sessions (`wn/ws/wh/wl/wr/wx`),
Tabs (`tn/th/tl/tr/tx`), `v` split right, `s` split down, `x` close pane, `e`
sidebar and `c` command palette, with existing labels/icons/confirmation flags.
There is no restored Zoom, directional focus, numbered tabs, app chooser or
default host row.

See `examples/my_whichkey.json` for a whole-spec replacement with literal groups,
a hidden active binding, destructive confirmation and opt-in provider markers.
To retain built-ins, omit `spec` rather than copying partial overrides.

Migrate old `M.new`/`setup`/`register_launch` menu data by flattening literal key
sequences into this JSON object. Do not copy Lua source, function-valued
conditions/labels, callbacks or callback references. Those require a future
separately designed hook feature, not this schema. The minimal keybinding and
package migration belong to later Phase 16 work; no user configuration has been
rewritten or reloaded by implementing this contract.
