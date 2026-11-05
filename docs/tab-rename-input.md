# Tab and session rename input

## Accepted flow

**REXWHICHKEY-20261103-085109 — Accepted.** The user explicitly approved a
text-entry state inside the existing popup rather than another popup.

In both the native menu and the registered Lua defaults, **Tabs → Rename** opens an input screen in the same
layer and fixed rectangle. The current server tab label is prefilled. The field
starts selected: typing replaces it; arrows/Home/End allow editing instead.

- **Enter:** validate and confirm the name.
- **Escape:** discard the edit and return to the existing Tabs submenu/page.
- **Backspace/Delete:** remove whole grapheme clusters.
- **Ctrl-U:** clear the field.
- Bracketed paste is bounded text data, never submission or menu keys.
- Names must be nonblank UTF-8 without control characters, at most 4096 bytes.
  Existing names beyond that limit are not silently truncated and submitted;
  the prototype starts an empty field requiring a new name.

The input uses the existing semantic colors, status spacers, focus-in/resize
redraw and terminal restoration. A software caret stays visible while long input
scrolls horizontally. No additional pane/header, native rename dialog, app
chooser, popup resize or GUI setting change is involved.

## Resource and completion ownership

The parent freezes the originating server/session/pane/tab placement before
showing the menu and reads that exact tab's label for presentation. The rename
form receives the initial name, not authority to choose its resource. Session
picker choices are separately retained as canonical parent-owned mappings.

On Enter, terminal restoration must succeed before the child commits a result.
The parent validates the retained `tr` selection and bounded name, then completes
popup close, origin focus restoration, private-file and lifetime-pipe cleanup.
Only afterward does it revalidate the frozen resource/placement and call typed
`Window.Rename` with the name as one CLI argument. No shell evaluation, app action,
replacement target, uncertain retry or label-based lookup is used. Explicit app
overrides retain their strict checks but no app binding is required by default.

`native-input.json` remains the compatible private presentation filename for
both routes. It contains `initial` for tab rename and `session_initial` for session
rename and is removed with other owned transport files. Mapping/Config JSON stays
unchanged. The typed Completion v1 selection now permits optional `text` only for
the retained empty-argument `tr`/`client.tab.rename` or `wr`/`session.rename` intent,
without callback/confirmation. Those intents require text; other actions reject it.
The controller validates the child result before forwarding this name data.

## Scope and evidence

The earlier prototype was native-only. The user's configured-menu parity
correction extends the same input state to registered Lua defaults. Lua retains
the reservation's exact tab/session, validates text, guards, capability, origin
and connection, then invokes `session.set_window_label` or `session.set_label`
after popup cleanup. A pane moved to another tab cannot redirect rename. Custom
callbacks, replacement actions and nonempty retained arguments are not
intercepted; incoming text cannot turn them into resource renames. Failed or
uncertain rename requests are consumed and never retried.

Registered WhichKey defaults no longer open Rex rename dialogs. Rex's own command
palette is unchanged. Native Switch Session remains its separate dynamic submenu;
see [session UI](session-owned-ui.md).

Current regression coverage includes both menu pages, prefilled names, Escape,
name transport/validation, retained mapping/resource binding, custom callbacks,
moved origins, uncertainty and replay. These are unit/race checks, not a claim
that live tab/session rename was exercised.

Historical prototype evidence: compilation/source checks only were performed
for the original native-only delivery under the stop-testing instruction. The
current parity regression checks above are distinct from that historical record;
no live rename, configuration reload or settings change was performed here.
