# Tab and session rename input

## Current Go-owned flow

Accepted text-entry direction: `REXWHICHKEY-20261103-085109`.
Both built-in and JSON-configured `client.tab.rename` / `session.rename`
use a text-entry state in the existing WhichKey popup, not a native dialog,
second popup or special hardcoded `tr/wr` dispatch path.
Built-in keys remain Tabs → Rename (`tr`) and Sessions → Rename (`wr`).

The exact captured Window/Session label prefills the field, initially selected:

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

[prepareNativeMenu](../native.go) captures labels before popup creation;
`native-input.json` is private presentation data, not public menu schema or
replacement identity. [text_entry.go](../text_entry.go) handles input;
[native_text_entry.go](../native_text_entry.go) validates private text/results.

The Go parent retains endpoint/session/block/window/layer, canonical menu mapping
and confirmation. The terminal returns keys and authorized bounded text; it
cannot supply an action, callback, replacement ID or app generation.
After terminal restoration, exact popup close, valid origin-focus restoration,
pipe/private-file removal and image-handle retirement, the parent rechecks
original placement and calls typed `Window.Rename` or `Session.Rename` once.
Text is one label argument, never shell code. See
[cleanup inventory](launch-contract.md#fixed-owned-cleanup-inventory).

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

Task2's coordinator observed menu/navigation/dismissal/cleanup on Darwin arm64,
**not rename**. Final combined runtime may exercise a benign inline rename only
on an explicitly coordinator-owned disposable resource; that result is pending,
not supplied by this source guide. No tests, probes, physical key/reload,
user-resource rename or deployment were performed by this leaf.

The old bridge-specific rename contract remains historical at
`cd5cbce818e07c264738ddcbc591ad49fdf0e9c8:docs/tab-rename-input.md`.
[Acceptance](acceptance.md) preserves prior evidence without treating it as a
current combined rename result.
