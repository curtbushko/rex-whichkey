# Tab rename text-entry prototype

## Accepted flow

**REXWHICHKEY-20261103-085109 — Accepted.** The user explicitly approved a
text-entry state inside the existing popup rather than another popup.

In the built-in native menu, **Tabs → Rename** opens an input screen in the same
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

`native-input.json` is private native-popup presentation data. It contains
`initial` for tab rename and `session_initial` for session rename, and is removed
with the other owned transport files. The native result may carry a `text` field
only for the retained `tr` or `wr` rename selection. The parent strips that field before the
canonical strict Completion decoder and retains the validated name privately.
This does not extend Mapping/Config JSON, Completion v1 JSON or the Lua bridge.
Configured completion rejects native form data rather than dropping it.

## Scope and evidence

This prototype is enabled only for the **built-in native menu**. Configured Lua
menus and callbacks keep their existing authenticated completion protocol and
Rex-native tab rename action. Session rename now reuses this screen, and Switch
Session uses our dynamic submenu; see [session UI](session-owned-ui.md). The
command palette retains Rex's UI.

Compilation and source/format/whitespace checks are separate from interactive
verification. No unit/general/race suites, UI trials, extra worker windows,
configuration reloads or font/settings changes were performed for this change,
following the user's stop-testing instruction.
