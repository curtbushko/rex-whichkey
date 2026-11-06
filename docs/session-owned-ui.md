# Session switch and rename screens

The built-in native menu uses the existing fixed WhichKey popup for both screens.
No extra Rex pane, header, app chooser or navigation resize is created. Registered
Lua defaults also use WhichKey's rename text page; configured **Switch Session**
retains its named-action/custom-menu semantics, not the native Go inventory adapter.

## Sessions → Switch (native)

Each invocation reads `session.list` through the frozen execution-host Rex client.
The parent expands `ws` into an ordinary canonical submenu:

- Rows show identifying IDs before labels and mark the origin `[current]`.
  Short-ID collisions use full IDs; labels are presentation only.
- Keys 1–9 select rows; `n` opens More groups. Normal Tab paging, Escape/Backspace,
  focus redraw and fixed bounds apply.
- The menu-opening snapshot is not refreshed during redraw. Unavailable/empty
  inventory is explicit; more than 512 sessions is unavailable, not truncated.

Full destination IDs remain in parent-retained Mapping args. The terminal returns
canonical keys, not a replacement destination. After exact popup/transport/source
cleanup, the parent revalidates destination, origin placement and pane-owner or
explicit app binding, including its connection generation. It sends one unscoped
`client.perform session.select` with the retained destination in args and no outer
origin session/block override. This uses Rex's resource-selected action, not its
native picker. Removed sessions, changed bindings and uncertain effects refuse
without retry or fallback. A remote viewer does not change the execution host.

## Sessions → Rename (both routes)

`wr` opens **Rename session**, prefilled from the captured session label. Typing
replaces the selected value; arrows/Home/End, grapheme deletion, Ctrl-U and bounded
paste edit. Enter validates/confirms; Escape returns without mutation. Text must
be nonblank control-free UTF-8, at most 4096 bytes **and 128 Unicode characters**.

Go validates resource/text support before popup. Configured Lua must advertise
`selection-text-resource-rename` with its real loaded menu generation; completion
1 alone is insufficient. The negotiated configured completion is 2, binding the
exact opaque Go receipt/hash without a Lua numeric reserialization. Compatible
build IDs need not match. Missing support rejects before popup, never opens a
native dialog, discards the label or treats generic action acceptance as rename.

After terminal restoration and exact owned-popup/file/handle cleanup, native Go
revalidates frozen placement and calls typed `Session.Rename` once. Default
resource rename needs no GUI app binding; explicit client intent remains strict.
Configured completion authenticates and claims the original-VM invocation before
yielding checks, validates the retained selection/guards/connection/generation
and original expiry, then invokes `session.set_label` for the retained session.
Tab rename similarly uses `session.set_window_label` for its retained window.
Text supplies no resource ID. Moved origins cannot retarget the operation.
Custom callbacks, args, conditions and confirmation are not intercepted.

Failure after claim consumes authority, even if the API acknowledgment is lost.
No uncertain retry, alternate target, fallback dialog or renewed expiry is
allowed. Cleanup uncertainty suppresses effects. Possible controller start keeps
workspace ownership with the child; lost acknowledgment or unauthenticated input
can retain artifacts and does not authorize broad deletion. The source-linked
[fixed owned cleanup inventory](launch-contract.md#fixed-owned-cleanup-inventory)
covers popup and rendezvous/preparation files plus the original-Lua release boundary.

See [rename input](tab-rename-input.md), [runtime bootstrap](runtime-bootstrap.md)
and the authoritative [mapping standard](mapping-format.md) for shared semantics.

## Evidence boundary

Source-reviewed/compiled intent is not observed runtime verification. Native
inventory selection, configured/native rename, label limits, callbacks, moved
origins, compatibility, disconnect/expiry/replay and uncertain cleanup/mutation
are **UNEXECUTED source scenarios** here. No tests, live CLI/API/menu/rename probes,
reloads, activation or settings changes were performed for this cutover; earlier
native-only evidence does not certify this combined revision.
