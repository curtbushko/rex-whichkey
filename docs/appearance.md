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

## Fixed Lua rectangle and own-PTY fitting

[`lua/rex_whichkey.lua`](../lua/rex_whichkey.lua) creates one centered float with
finite launch `width`/`height` fractions in `(0,1]`, default `.8/.3`:
`{(1-width)/2, (1-height)/2, width, height}`. Shared source/script/require launch
uses this same implementation. Source `make run` builds current Go and supplies
its absolute target-bin directory; installed launch uses ordinary PATH.
See [launch options](launch-contract.md#lua-presentation-and-creation).

Go does not measure origin geometry, choose native content-height bounds or seed
a hypothetical window grid. The direct menu starts at row zero and fits its
**own PTY** cells. Larger menus paginate; status/spacer behavior above is retained.
`MenuFrame.PreferredHeight` remains a natural terminal-row measurement, not a
native layer sizing command. JSON `position` is still validated but does not
choose outer bounds in direct mode; standalone UI/result positioning remains.

Submenu/back/paging/metadata/focus/idle never resize or replace the layer.
External terminal/window/font changes can alter own-PTY cells; SIGWINCH and
focus-in repaint/reflow retained content without application-driven outer resize.
No idle menu expiry, repaint timer, font/settings mutation or geometry feedback
loop is introduced. Palette replies are metadata, not keyboard actions.
Captured placement/liveness checks remain distinct from selected app-generation
checks. This is not pixel-perfect shrink wrapping, GUI pixel provenance or a
first-size-final guarantee. Historical Go-controller sizing observations below
are not evidence for this Lua sizing route.

### Native chrome: accepted defaults, customization deferred (REQ-026)

User1a explicitly accepted Rex defaults **for now** and deferred per-float native
titlebar/border/corner/shadow customization. It is not implemented or credited as
REQ-026 completion. The inspected Rex API has no supported per-float knobs for
that customization; this is not a claim about all future Rex versions.

No guessed/ignored API fields, global settings changes, clipped controls, font
changes or ANSI/text borders substitute for native chrome support. Go's terminal
palette matching below is separate. Task04 records the accepted policy and final
revision-bound evidence; it is not a native-feature implementation. Earlier
blocked statements/API observations remain historical in
[acceptance](acceptance.md#compact-popup-presentation-req-024025-and-blocked-req-026).

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

The direct Go process retains one menu and theme throughout `RunTerminal`. Its
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
  rendering runs in the Rex-created server-host terminal, not the caller's
  process. Lua does not forward caller `NO_COLOR` or supply an environment
  override to `session.new_layer`; do not assume a new caller-side export changes
  that terminal. Its effective environment is managed
  by Rex. Changing a terminal variable does not change an already-running
  server's environment.
- Public Lua launch `no_color=true` / `ascii=true` explicitly supplies Go direct
  `--no-color` / `--ascii`; the require binding accepts corresponding booleans.
  Source PATH wrapper: `make run -- no_color=true ascii=true`; GNU make:
  `make run ARGS='no_color=true ascii=true'`. Standalone UI/result `menu` also
  accepts Go flags. Bare Go/root flags and private `popup` mode are not launchers.
  ASCII mode changes decorations/optional icons, not user labels or color policy.
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

Phase19 Task01/02 native summaries bind direct-source menu/palette captures,
real own-PTY resize/reflow without outer-bound change, owned rename/cancel/dismissal
and pre-effect visible/post-exit host cleanup. The observed terminal resize had
no GUI owner; it is not a physical GUI drag or pixel measurement. New Task03
make-route checks on reviewed combined source and rebuilt main remain coordinator
gates. No physical-key/reload, installed-profile, Nix build/activation, every-
palette/geometry/flags or all-platform result is claimed by this source update.

The following observations are **historical and revision-bound**, not current
combined-candidate certification. Legacy configured-Lua/callback references
describe those earlier revisions, not supported current setup.

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
