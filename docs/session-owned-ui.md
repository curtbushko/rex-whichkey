# Go-owned session switch and rename screens

The direct public binary and minimal function keybinding use the same Go menu.
JSON configuration can place the supported intents at any valid full key.
There is no native session picker, app chooser, extra pane or navigation resize.

## Sessions → Switch

The default `ws` / `session.switch` is a Go sessions-provider adapter.
Enabled `sessions` markers use the same expansion. One bounded invocation snapshot
is read through the frozen Rex CLI before sizing/PTY handoff:

- Full typed IDs lead labels; the origin is marked `[current]`.
- Keys 1–9 select rows; `n` opens More groups. Ordinary Tab paging,
  Escape/Backspace and focus redraw retain fixed bounds.
- Empty inventory has a non-executable message; read failure has an explicit
  warning, not a successful empty inventory or picker fallback.
- More than 512 resources is an error, not truncation. Expansion batches and
  combined transport have further validated bounds.
- Redraw does not refresh inventory. Reopen to obtain a new snapshot.

Destination IDs remain in the parent's retained mapping, not child-supplied
authority. After exact cleanup, Go checks destination membership, original
placement and retained app/connection/owner and sends one unscoped
`client.perform session.select` with exact destination args. No outer origin
session/block override substitutes for the destination. A removed session,
reconnected app or uncertain effect refuses without retry or label/index fallback.

Optional `windows` lists origin-session Windows with exact session/window
selection; `panes` lists origin-session Blocks and focuses the retained Block
only after placement revalidation. Neither is a default menu row.
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

The coordinator actually observed Task2 root → Sessions → Switch with full IDs
and current marker, Backspace to root, Escape, exit0 and exact cleanup at
`7bdbda2` on Darwin arm64. It did **not** select a session or rename a resource.
Final combined reviewed runtime and any benign owned selected effect remain
pending; physical keybinding and config reload are not established.

Historical bridge-era session guidance is preserved in Git at
`cd5cbce818e07c264738ddcbc591ad49fdf0e9c8:docs/session-owned-ui.md`.
It is not an active prerequisite. See [acceptance](acceptance.md) for exact scope.
