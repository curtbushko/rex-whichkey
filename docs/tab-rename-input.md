# Tab and session rename input

## Current Go-owned flow

Accepted text-entry direction: `REXWHICHKEY-20261103-085109`.
Both built-in and JSON-configured `client.tab.rename` / `session.rename`
use a text-entry state in the existing WhichKey popup, not a native dialog,
second popup or special hardcoded `tr/wr` dispatch path.
Built-in keys remain Tabs → Rename (`tr`) and Sessions → Rename (`wr`).

On entry the form requests the captured Window/Session label. Loading/unavailable
context is shown inline, never in a native dialog. Edits and Escape work while
loading, but Enter is discarded until a retained ready publication exists. There
is no pending-submit replay. A ready label prefills an untouched field, initially
selected; it never overwrites edits, cursor movement, errors, partial UTF-8/escape/
paste input or confirmed text. Revisited forms use the invocation's shared label
snapshot, not a persistent cache:

- Typing replaces the selection; arrows/Home/End edit.
- Backspace/Delete remove grapheme clusters; Ctrl-U clears.
- Bounded bracketed paste is text data, never submission or mapping keys.
- Enter validates/confirms; Escape discards edits and returns to the existing
  submenu/page. A mapping's optional confirmation is still required.
- Names must be nonblank, control-free UTF-8, at most **4096 bytes and 128 Unicode
  characters**. An invalid/oversized initial presentation value is cleared,
  never silently truncated/submitted.

The same fixed rectangle, theme, status spacing, caret/scrolling and focus/resize
redraw apply. No app chooser or native control/font setting is changed.

## Retained resource and cleanup

[`RunDirectMenu`](../direct_menu.go) enables forms without eager label reads.
The required matching origin view supplies label/window/block presentation on
form demand, excluding the menu's own popup. [text_entry.go](../text_entry.go)
handles input; [menu_updates.go](../menu_updates.go) applies updates only on the
UI owner. The [in-process contract](launch-contract.md#in-process-menu-metadata-startup-ui-updates)
uses bounded detached results, not private files or child IPC. The scheduler
executes at most two jobs, including labels. Window/pane/rename presentation
reuses that matching origin view; sessions-list and hosts have independent
once-per-invocation results. See [provider scheduling](mapping-format.md#complete-provider-catalogue-and-expansion-boundary).

The single Go owner retains endpoint/session/block/window/layer, canonical
mapping, confirmation and ready publication. Selection text/bindings cannot
supply replacement action/ID/generation authority. Unknown, stale, malformed or
pending selection refuses. After worker join, tty restoration, exact visible
float detach/verification and exact original-focus restoration, Go rechecks
placement and calls typed `Window.Rename` or `Session.Rename` once. Text is one
label argument, never shell code. There is no parent/renderer, image copy,
lifetime pipe, private demand/update/result file or source-handle cleanup.

Accepted user1b permits the detached execution-host Block through dispatch until
Go exits/Rex deletes it. Go cannot observe its own deletion; coordinator must
inspect exact `session.list_blocks` after exit. A later cleanup failure cannot
suppress/undo an earlier rename. See [cleanup](launch-contract.md#fixed-owned-cleanup-inventory).

Resource rename does not need an implicit GUI app. Explicit client intent remains
strict and is revalidated, including its connection generation.
Missing/moved origin, invalid text, uncertain cleanup or changed explicit app
suppresses the effect. Lost acknowledgment/uncertain rename is not retried,
rolled back, redirected or reported as successful. Checks/mutation are not atomic.

## Migration and evidence

There is no loaded-Lua consumer feature, completion version, original-VM ticket,
callback registry or expiry negotiation in this path.
Configured rename takes empty args and follows action intent at any valid key.
Functions/callbacks/custom Lua actions are deferred, not silently converted into
rename. See [data catalogue](mapping-format.md).

Phase19 Task01/02 native summaries bind direct-source owned inline rename,
edit cancellation, default dismissal and verified pre-effect visible retirement/
post-exit host deletion. Earlier Phase16/18 results remain historical. New Task03
make-route runtime on reviewed combined source and rebuilt main remains a
coordinator gate; this leaf ran no product invocation. No physical-key/reload,
installed-profile, arbitrary user-resource rename or platform-wide result follows.
See [acceptance](acceptance.md) for revision-bound artifacts.

The old bridge-specific rename contract remains historical at
`cd5cbce818e07c264738ddcbc591ad49fdf0e9c8:docs/tab-rename-input.md`.
[Acceptance](acceptance.md) preserves prior evidence without treating it as a
current combined rename result.
