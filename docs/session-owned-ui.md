# Go-owned session switch and rename screens

The direct public binary and minimal function keybinding use the same Go menu.
JSON configuration can place the supported intents at any valid full key.
There is no native session picker, app chooser, extra pane or navigation resize.

## Sessions → Switch

The default `ws` / `session.switch` is a Go sessions-provider adapter.
Enabled `sessions` markers use the same expansion. Visible enabled markers request
one bounded list through the selected Rex CLI only after the static first draw;
hidden enabled markers wait for explicit prefix entry. All markers share one
in-flight attempt and terminal result, including failures:

- Full typed IDs lead labels; the origin is marked `[current]`.
- Keys 1–9 select rows; `n` opens More groups. Ordinary Tab paging,
  Escape/Backspace and focus redraw retain fixed bounds.
- Empty inventory has a non-executable message; read failure has an explicit
  warning, not a successful empty inventory or picker fallback.
- More than 512 resources is an error, not truncation. Expansion batches and
  combined transport have further validated bounds.
- Loading is non-executable and does not buffer keys/Enter. Redraw or re-entering
  a prefix does not retry/refresh. Start a new invocation for a new snapshot.

Destination IDs remain in the parent's retained mapping, not child-supplied
authority. After exact cleanup, Go checks destination membership, original
placement and retained app/connection/owner and sends one unscoped
`client.perform session.select` with exact destination args. No outer origin
session/block override substitutes for the destination. A removed session,
reconnected app or uncertain effect refuses without retry or label/index fallback.

Optional `windows` and `panes` wait for first prefix entry and share the already
captured origin session-view presentation: Windows retain exact session/window
selection; Blocks retain exact placement and focus only after fresh revalidation.
Neither is a default menu row.
`known-hosts` remains listing-only. See [providers](mapping-format.md#complete-provider-catalogue-and-expansion-boundary).

## Sessions → Rename

Default `wr` / `session.rename` opens the existing popup's text-entry page,
prefilled from the captured Session label. It follows the same Go resource/text
flow as tab rename, including custom keys, optional confirmation and strict
nonblank UTF-8/4096-byte/128-character validation.
Escape returns without mutation; Enter validates. After complete cleanup, Go
revalidates retained placement and calls `Session.Rename` once.
No implicit GUI is needed; explicit app intent remains strict.

There is no original-Lua completion/receipt/callback route, custom executable
configuration, refreshed expiry or native-dialog fallback.
[Rename details](tab-rename-input.md) and the
[owned cleanup inventory](launch-contract.md#fixed-owned-cleanup-inventory)
apply to both session and tab rename.

## Evidence and history

Phase16 final public menu/navigation/dismissal/cleanup and an exact owned inline
rename were actually verified on Darwin arm64. Phase18 task2 additionally observed
loading edits, ready preservation, cancel without rename, revisit/owned rename
and exact cleanup. These runs do not certify final async provider behavior or
session selection. Final combined source review and coordinator public
first-frame/data-ready/action-ready, effect/cancel/cleanup checks remain pending;
physical keybinding, installed profiles and config reload are not established.

Historical bridge-era session guidance is preserved in Git at
`cd5cbce818e07c264738ddcbc591ad49fdf0e9c8:docs/session-owned-ui.md`.
It is not an active prerequisite. See [acceptance](acceptance.md) for exact scope.
