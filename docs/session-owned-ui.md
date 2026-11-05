# Session switch and rename screens

The **built-in native menu** uses our existing popup for both entries. No extra
Rex pane, popup, header, app chooser or navigation resize is created. Configured
Lua menus keep their existing authenticated protocol and native Rex actions.

## Sessions → Switch

Each invocation reads `session.list` through the existing typed Rex client. The
parent expands the logical `ws` entry into an ordinary canonical submenu:

- Session names are loaded from Rex, not hardcoded.
- Rows show a short identifying ID before the label and mark the origin
  session `[current]`. Short-ID collisions use full IDs. Labels are display-only.
- Keys 1–9 select rows; `n` opens More groups when needed. Existing Tab paging,
  Escape/Backspace navigation, focus redraw and fixed bounds remain in effect.
- The list is a snapshot at menu opening. Reopen the menu to refresh it.
- Unavailable/empty inventory is shown without inventing choices. More than 512
  sessions produces an explicit unavailable list rather than silent truncation.

Every child row's full destination ID is stored in the parent's retained Mapping
args. The terminal returns only a normal canonical key selection; no new result
field, child-supplied destination, label lookup or positional retargeting is used.
After full popup/transport cleanup, the parent revalidates the exact destination,
origin placement and pane-owner/explicit app binding. It sends one unscoped
`client.perform session.select` with the retained destination ID as an argument.
There is no outer origin session/block override to replace that destination.
Removed sessions or changed bindings refuse without retry or fallback. This uses
Rex's action API to activate the selected session, not Rex's native session picker.

## Sessions → Rename

The same text-entry state as tab rename opens with the origin session label
prefilled from the snapshot. The header identifies **Rename session**.

Typing replaces the initially selected value; arrows, Home/End, Backspace/Delete,
Ctrl-U and bounded bracketed paste edit it. Enter validates/commits; Escape
returns to Sessions without changes. Names follow the same nonblank UTF-8,
no-control-character, 4096-byte policy as tab rename.

After terminal and popup/transport cleanup, the parent revalidates its frozen
session/pane/tab placement and invokes typed `Session.Rename` once. No app action
or native rename dialog is used. An explicit app override is still validated;
otherwise resource rename needs no GUI app binding. The private presentation
file adds `session_initial`; private `text` data is allowed only for retained
`wr` or `tr` rename selections. Config/Mapping JSON remains unchanged. Registered
Lua defaults now use the same rename form; typed Completion v1 selection allows
bounded `text` only for the retained default rename intent. The Lua dispatcher
revalidates the invocation and uses `session.set_label` or
`session.set_window_label`, never a GUI dialog. Custom callbacks remain ordinary
callbacks. See [rename input](tab-rename-input.md) for the updated wire contract.

## Evidence boundary

Historical native-only delivery used build/source checks under the stop-testing
instruction. The subsequent registered-rename parity correction has focused
unit/race regression coverage described in [rename input](tab-rename-input.md).
Neither record claims live rename verification or configuration/settings changes.
