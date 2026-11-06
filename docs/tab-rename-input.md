# Tab and session rename input

## Accepted flow

**REXWHICHKEY-20261103-085109 — Accepted.** Rename is a text-entry state in the
existing WhichKey popup, not another popup or Rex's native rename dialog.
Both the Go native menu and registered Lua defaults use it for **Tabs → Rename**
(`tr`) and **Sessions → Rename** (`wr`). The captured resource's current label
prefills the field, initially selected so typing replaces it.

- Enter validates/confirms; Escape discards edits and returns to the existing
  submenu/page.
- Arrows/Home/End edit; Backspace/Delete remove whole grapheme clusters; Ctrl-U
  clears. Bounded bracketed paste is text data, not submission or mapping keys.
- Names must be nonblank, control-free UTF-8, at most **4096 bytes and 128 Unicode
  characters** (Rex's documented resource-label limit). The larger transport bound
  does not relax the API limit. An initial label beyond the presentation byte
  bound or containing invalid text is cleared, not silently truncated/submitted;
  the user must supply a valid replacement.

The same layer, fixed rectangle, theme, status spacing and focus/resize redraw
remain. A software caret and horizontal scrolling handle long text. No native
controls, global GUI setting, extra pane/header or app chooser is changed.

## Compatibility before popup

Go validates rename presentation/resource support and the loaded consumer's
`selection-text-resource-rename` feature **before** popup creation, then rechecks
loaded-menu/origin state after yields. Protocol/layout 1 and configured completion
2 are Go-owned; completion 1 alone is not text-feature evidence. Compatible helper,
public and loaded-Lua build IDs may differ. Unsupported consumers fail before
popup rather than opening a native dialog, dropping text or reporting false
success. See [runtime bootstrap](runtime-bootstrap.md) and the authoritative
[mapping/completion standard](mapping-format.md) for exact formats.

The private presentation asset is `native-input.json` for both routes (`initial`
and `session_initial`). It supplies labels, not authority or replacement IDs.
No user mapping field or constructor path is introduced.

## Exact retained-resource ownership

The parent freezes server/session/block and unique window/layer placement before
presentation. Lua retains its exact original window/session resources, mapping
snapshot, callback/args/guards/confirmation and ticket in the original VM. Default
inline rename applies only to retained `tr`/`client.tab.rename` or
`wr`/`session.rename` with no callback, nonempty args or confirmation override.
Custom callbacks, replacement actions, conditions, args and confirmation keep
ordinary mapping semantics; incoming text cannot turn them into default rename.

Successful terminal restoration precedes result commit. The controller validates
keys/confirmation/text against its snapshot, closes only the exact owned popup,
restores only a valid origin and retires private transport/source handles **before**
dispatch. Any uncertain close, focus or file/handle cleanup suppresses the effect.
The complete source-linked [fixed owned inventory](launch-contract.md#fixed-owned-cleanup-inventory)
includes the rename presentation and private image/receipt, not only config/result.
Native Go then revalidates retained placement and invokes typed `Window.Rename`
or `Session.Rename` once; default resource rename needs no app lookup. Explicit
client intent still receives its strict checks.

Configured completion authenticates capability, exact opaque receipt/hash,
actual fresh caller/server/session/block, original expiry and generation before
single-use claim. It claims before yielding guards/API work and rechecks retained
GUI connection/placement and state. Original Lua calls `session.set_window_label`
with its retained window/session or `session.set_label` with its retained session,
once. Text is only the validated label; a moved origin cannot redirect rename.
The callback VM is never reconstructed and Go identities are never reserialized
through Lua 5.1 numbers. The original expiry is never renewed.

A failed/uncertain rename or lost acknowledgment consumes claimed authority;
there is no retry, replacement target, dialog fallback or rollback inferred from
an error. Read-only checks and mutation are not atomic against external changes.
Possible controller start also transfers workspace ownership irreversibly;
uncertain startup or malformed unauthenticated handoff may leave artifacts, not
permission for broad deletion. See [launch contract](launch-contract.md).

## Evidence boundary

These are source-reviewed/compiled implementation intentions, **not observed
runtime rename verification**. Default tab/session text, label limits, custom
callbacks/args/confirmation, moved origins, missing features, unequal compatible
builds, expiry/replay, cleanup failure and uncertain mutation remain **UNEXECUTED
source scenarios** for this cutover. No tests, Lua harnesses, live rename/API/menu
probes, configuration reloads or settings changes were performed. Earlier native
prototype observations do not verify this combined revision. Session switching
has its separate [session UI contract](session-owned-ui.md).
