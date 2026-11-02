# Menu appearance

The renderer assigns separate semantic roles to keys, icons and labels. These
roles live in [`theme.go`](../theme.go), not in a second Lua/Nix theme schema.
[`render.go`](../render.go) applies them to actual menu spans;
[`terminal.go`](../terminal.go) writes the styled frames and updates the theme
when terminal palette metadata arrives.

## Compact status borders and paging

[`render.go`](../render.go) and [`layout.go`](../layout.go) separate three regions:
a full-width top status border containing title/icon/breadcrumbs, the menu entries,
and a full-width bottom status border containing concise key hints. Status text
shares its row with the horizontal rule (`─`, or `-` in ASCII mode). One blank
spacer separates content from each status border when room permits. No `Keys:`
prefix or separate page-footer row is generated. These are terminal text borders,
not native Rex pane/tab/window controls.

A single-page menu has no page counter or Tab hint. Multiple pages retain a
`current/total` counter and Tab navigation in the bottom border where width permits.
Paging shares that row rather than allocating another row that could itself
cause pagination. The renderer picks shorter complete hints before grapheme
clipping; at narrow widths Escape takes priority over decoration/counter text.
Hidden and off-page bindings remain active, and Tab still wraps through pages.

Content takes priority: page capacity and page count are established without
spacers, then separation uses rows left after the selected page's visible content.
Partially filled final pages can therefore have spacers even when full pages
cannot. Bottom-hint separation takes the first spare row; both spacers appear
when two rows remain. Three available rows permit both status borders, two permit
content plus bottom hints, and one permits content alone. Confirmation replaces
the entries with `y=yes / other=cancel: …` in the critical role, so
confirmation/cancellation instructions remain on the content row even without
status borders, subject to width clipping. A zero-height viewport draws nothing
but does not mutate navigation. Semantic roles and text warnings remain intact
without color; ASCII mode changes decorations, not user labels.

## Actual layer height and measurement limits

`MenuFrame.PreferredHeight` is the natural terminal-row count at the current
width: all current entries (or one confirmation/empty message) plus both status
borders and both ideal spacer rows. It is independent of viewport-height clipping,
pagination and selected page. `Height`/`Lines` still describe only the constrained
visible frame. The geometry controller consumes this exact producer value, not
a second menu measurement or the already-clipped frame height.

Before creating the popup, [`popup_initial.go`](../popup_initial.go) calls this
same renderer at an estimated popup width. The parent computes a content-sized
rectangle and sends terminal startup columns/rows in the **single creation RPC**,
instead of opening at the position maximum or using the server's unrelated 80×24
default. It uses the exact origin leaf's window-relative rectangle and validated
current-owner terminal measurement, never the active window or a leaf rectangle
multiplied by floating-layer bounds. Split/floating origins keep their actual
scale; documented zoomed origins use the window rectangle. Placement changes
during sampling or before creation invalidate the measured snapshot.

Initial sizing is an estimate: normalize origin bounds by measured rows/columns,
reserve two cells for insets, and round height upward by half a row. Optional
valid paired cell pixels seed terminal metadata, not font preferences or viewer
measurement provenance. New-pane font/native chrome can differ from the origin.
When geometry is missing, malformed or unmeasured, use a hypothetical 100-column,
40-row window grid and cap initial height at 30% (or the smaller position cap).
The seeded rows are the natural count capped to that estimated available space;
large menus paginate. This compact fallback is policy, not measured GUI geometry.
No hidden/offscreen first layer, recreation, cache or app binding is involved.

The popup child then publishes preferred rows and its actual terminal columns/rows
only when those values change. The existing liveness watcher refines the **same
owned floating layer**, through `Layer.Resize`, for startup estimation differences
and later submenu/back or viewer changes. Late shrinking is not the startup
sizing strategy. Popup frames start at row zero during settling; standalone `RunTerminal` positioning is unchanged. See the
[launch contract](launch-contract.md#private-preferred-row-hints-and-owned-layer-sizing)
for private-file validation, ownership and cleanup.

Sizing uses current normalized layer bounds and applied terminal rows whose
current owner has exactly one matching, timestamped desired-size report. Seeded
80×24 defaults alone are not measurements. Missing provenance pauses adaptation
without canceling native display; zero or child/server-disagreeing dimensions do
not drive a resize. Optional cell-pixel dimensions help detect geometry changes;
applied rows suffice when pixels are absent. No fixed font pixels are guessed and
no font preference is changed.
Owner, width, row/cell-geometry and natural-content changes trigger reevaluation.

The original position rectangle caps height: top/bottom use at most 40% of the
window height, center 60%, and left/right 100%. Width and horizontal placement
stay unchanged. Shrinking/growing preserves the original top edge, bottom edge,
or vertical center (center/side positions). Menus too tall for that cap paginate.
Half-row headroom and a one-row deadband tolerate rounding, so fitting is not a
promise of pixel-tight native borders. Newer owner reports and stable samples
bound successive corrections; missing fresh reports after a resize,
nonconvergence or uncertain resize fail closed rather than looping or dispatching
a selection. Report timestamps
are not native layer-application acknowledgements. There is no idle repaint
timer, replacement popup, daemon, app-action lookup or global font/settings edit.

### Native controls remain blocked (REQ-026)

The requested per-popup hiding of native Rex pane/tab/window controls is **not
implemented and not deferred by the user**. Retained `session.new_layer` and
`session.resize_layer` schemas expose bounds/layout/focus, not a chrome switch;
the shell creator rejects unknown options. The captured connected-app action
registry and published Lua layout reference provide no supported per-popup
control/header/toolbar suppression. This is a limitation of the inspected Rex
capabilities, not proof that no future Rex version can support it.

No undocumented flag, global setting, font change or clipped-over control is a
substitute. A supported Rex capability/update or explicit user deferral is needed;
Phase11/project completion remains blocked by this requirement. The
[acceptance inventory](acceptance.md#compact-popup-presentation-req-024025-and-blocked-req-026)
indexes the API evidence and distinguishes measured server layer/terminal frames
from native GUI chrome, including remote clients.

## Roles and inspiration

| Content | Role | Preferred treatment |
| --- | --- | --- |
| Displayed keys | `ThemeKey` | Cyan, then blue palette accents |
| Menu/header icons, breadcrumb separators, group markers | `ThemeIcon` | Yellow accent |
| Ordinary labels and title | `ThemePrimary` | Readable body foreground |
| Group and breadcrumb labels | `ThemeGroup` | Magenta accent |
| Help, paging and clipping markers | `ThemeMuted` | Readable neutral |
| Confirm-marked action labels and `[confirm]` text; confirmation prompt/footer | `ThemeCritical` | Red accent, plus explicit text |

Inventory warnings from the optional host provider are non-executable group rows,
not confirmation content. Their labels use `ThemeGroup`, not `ThemeCritical`:
[`hosts.go`](../hosts.go) creates them as groups, and
[`render.go`](../render.go) styles group labels accordingly. Warning text alone
does not select a critical role.

The inspiration is the user's
[Neovim WhichKey configuration](https://github.com/curtbushko/neovim-flake/blob/main/config/plugins/which-key.nix).
The retained source snapshot is `.mtn-tmp/menu-polish/neovim-which-key.nix` in the
authoritative checkout. It names mostly yellow icons with cyan/orange/green/blue
category accents; it does **not** prescribe exact key/description hex colors.
Rex WhichKey borrows the yellow/cyan emphasis, not a pixel-identical palette or
per-item color configuration. Mapping `icon` remains a string; no `icon.color`,
hex-color setup option or imported Neovim highlight schema is supported.

## Immediate fallback and palette replies

The first frame does not wait for OSC replies. Without a valid background reply,
keys use terminal-native ANSI cyan, icons yellow, groups magenta and critical text
red. Body/help text use the terminal's default foreground; no background is set
and no dark theme is assumed. These accents follow the user's terminal palette,
so their contrast cannot be measured while background metadata is unavailable.
A missing or malformed reply does not make the whole menu plain.

The terminal sends read-only OSC 10/11/4 queries for foreground, background and
the base ANSI palette. It never assigns palette entries. Partial and late replies
are decoded as metadata, not keystrokes or Escape dismissal, and subsequent
frames use the updated snapshot without blocking navigation.

Once valid background metadata is available, the theme uses the reported
foreground if readable, otherwise a contrasting neutral. It prefers compatible
palette accents and adjusts missing/low-contrast accents toward a readable
neutral to reach the source's 4.5:1 contrast threshold. Styling uses that reported
background, not a guessed one. Individual roles can converge under extreme
palette/contrast constraints; distinct roles are not a promise of fixed RGB
values or universally distinct hues.

## Foreground return

The real popup child retains one menu and theme throughout `RunTerminal`. Its
terminal ownership now includes saving, enabling and restoring xterm focus-reporting
mode 1004 alongside wrap, cursor visibility and bracketed paste. The single input
reader passes bytes through `PaletteDecoder`, then `terminalKeys`; only a complete,
unparameterized `CSI I` marks focus-in. It requests `drawTerminal` -> `RenderMenu`
without `HandleKey`, preserving submenu, page and pending confirmation. Fragmented
reports retain decoder state across reads; multiple reports and interleaved OSC
replies stay ordered. Focus-out (`CSI O`) and other terminal reports are consumed
without navigation, dismissal or an action. An incomplete CSI never times out
into mapping keys.

Repaint uses the current viewport size and theme, including partial palette
metadata and no-color/ASCII settings. It does not query the palette again, change
Rex focus, recreate a popup or run an idle redraw timer. See xterm's published
[focus protocol](https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h3-FocusIn_FocusOut).
This source path depends on terminal focus-report delivery; it cannot by itself
prove that a native app restores/displays its floating layer on tab return.

## Opt-outs, width and terminal safety

- A **nonempty** `NO_COLOR` in the execution environment disables theme color
  output and skips palette queries **in the rendering process**. Normal popup
  rendering runs in the Rex-created server-host terminal child, not the calling
  terminal's process. The launchers do not forward the caller's `NO_COLOR` or
  supply an environment override to `session.new_layer`; do not assume a new
  caller-side export changes that child. Its effective environment is managed
  by Rex. Changing a terminal variable does not change an already-running
  server's environment.
- Private `menu` and `popup` modes also accept `--no-color` and `--ascii`.
  They are not public root-launch flags or Lua setup fields; the delivered
  launchers do not forward them. ASCII mode replaces application decorations
  with ASCII equivalents and omits optional icons, but does not transliterate
  user labels or itself disable colors.
- Control characters are removed before text measurement/styling, including
  in no-color mode. The renderer clips plain grapheme/cell-width spans before
  adding SGR, so color escapes do not consume layout columns. Actual font/glyph
  widths and tiny-viewport geometry still depend on the client.
- Keys, labels, group markers, `[confirm]` text and confirmation prompts carry
  meaning without color. A destructive selection still requires `y`; color
  never replaces eligibility checks or confirmation.
- Each styled span resets its changed colors. The terminal owner separately
  restores display modes, cursor and termios before a result can be committed.
  I/O/restoration failures are errors, not successful selections; cleanup is
  bounded best effort, not a guarantee after SIGKILL or server/terminal loss.

## Evidence

The [compact-presentation inventory](acceptance.md#compact-popup-presentation-req-024025-and-blocked-req-026)
records supplied producer review/compilation and coordinator measurements at
`3c68b784`, including actual root/Tabs/back normalized bounds, terminal rows and
VT/HTML status frames. That one font/viewport observation is not a font-change,
five-position, multipage, tiny-viewport, ASCII or no-color runtime matrix, nor
native-controls-removal proof. Documentation leaves run static checks only;
revision-bound assembly outcomes belong to task 11-04 and its phase record.

[Theme task 08-03](../.tasks/tasks/rex-whichkey-08-03-theme.md) records fresh source
review and coordinator compilation at `3e9a1fd`; those are not graphical
observations. The [assembly record](../.tasks/tasks/rex-whichkey-08-04-docs.md) and
[acceptance inventory](acceptance.md) index the coordinator's frozen-candidate
35-second idle/default/color/navigation/Escape gate, including actual VT and HTML
frames. A plain text screenshot or this source description alone proves no color
pass. The gate covers one existing Linux Rex session, not all palettes, fonts,
Darwin/remote deployments or configured Lua callbacks. No runtime check is run by
the documentation builder/reviewer. Phase09's focus-in repaint has no new runtime
result here; earlier Phase07/08 observations do not verify that later path. The
[tab-return inventory](acceptance.md#tab-return-and-default-zoom-removal-req-021022)
separates retained server/process evidence from native GUI visibility.
