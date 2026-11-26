# Menu appearance

The renderer assigns separate semantic roles to keys, icons and labels. These
roles live in [`theme.go`](../theme.go), not in a second Lua/Nix theme schema.
[`render.go`](../render.go) applies them to actual menu spans;
[`terminal.go`](../terminal.go) owns terminal acquisition/restoration;
[`terminal_input.go`](../terminal_input.go) orders keys, metadata and redraws,
updating the theme when terminal palette replies arrive. Input framing lives in
[`terminal_keys.go`](../terminal_keys.go), and bounded display rollback in
[`terminal_modes.go`](../terminal_modes.go).

## Borderless heading, muted footer and paging

[`render.go`](../render.go) and [`layout.go`](../layout.go) render an optional
plain-text heading (title/icon/breadcrumbs), menu entries, and a plain-text footer.
There are no application-drawn border lines, including in ASCII mode. Direct
popups have one top blank row, a menu/footer separation row when space permits,
and instructions on the final row, with no blank padding below them. Standalone
frames retain padding at both edges. Padding counts toward fitted height and
yields to content/navigation in tiny viewports. An absent heading reserves no
additional row or spacer. Rex's native popup chrome is unchanged. The popup explicitly uses non-login
startup (`shell="none", login=false`) so no login banner flashes before the menu.

In direct popups the menu stays at the top. Extra viewport rows become blank
space between menu and instructions, not below them. Anchoring does not inflate
the natural height used by content fitting or the startup cache; removing the
bottom padding reduces the preferred popup height by one row.

The footer uses the terminal's muted neutral (ANSI gray before palette replies;
a readable palette neutral or foreground/background-derived color afterwards).
It does not fall back to the bright body foreground merely because ANSI gray
has low contrast. Known backgrounds retain at least 4.5:1 text contrast. The
right-hand `rex-whichkey` label appears when it fits without removing navigation,
paging or recovery hints. Confirmation/error hints retain their critical color;
branding stays muted. NO_COLOR remains plain text.

A single-page menu has no page counter or Tab hint. Multiple pages retain a
`current/total` counter and Tab navigation in the footer where width permits.
Paging shares that row rather than allocating another row that could itself
cause pagination. The renderer picks shorter complete hints before grapheme
clipping; at narrow widths Escape takes priority over decoration/counter text.
Hidden and off-page bindings remain active, and Tab still wraps through pages.

Direct popups reserve one menu/footer separation row before calculating page
capacity whenever content and navigation still fit, so full pages retain that
gap too. Heading separation uses spare rows. Standalone UI keeps its configured
positioning and compact frame: its spacers use only rows left after content.
Tiny viewports prioritize content and navigation over spacing. Three available
inner rows permit heading and footer, two permit content plus bottom hints,
and one permits content alone. Confirmation replaces
the entries with `y=yes / other=cancel: …` in the critical role, so
confirmation/cancellation instructions remain on the content row even without
status rows, subject to width clipping. A zero-height viewport draws nothing
but does not mutate navigation. Semantic roles and text warnings remain intact
without color; ASCII mode changes decorations, not user labels.

## Content-sized height and own-PTY fitting

[`lua/rex_whichkey.lua`](../lua/rex_whichkey.lua) creates one centered float with
finite launch `width`/`height` fractions in `(0,1]`, default `.8/.2`:
`{(1-width)/2, (1-initial_height)/2, width, initial_height}`. A matching cached
root height supplies `initial_height`; a miss uses the configured maximum.
Shared source/script/require launch
uses this same implementation. Source `make run` builds current Go and supplies
its absolute target-bin directory; installed launch uses ordinary PATH.
See [launch options](launch-contract.md#lua-presentation-and-creation).

Go retains that width and treats launch height as a cap. The current menu's
`MenuFrame.PreferredHeight` supplies its natural content rows plus the footer,
optional heading and their spacers. An independent worker fits the same verified layer to those rows,
including submenu/back and provider updates, using only its **own client-sized
PTY**. Longer menus paginate at the height cap. Screen/font changes reflow and
refit; focus/palette redraws alone do not resize unchanged content.

There is no origin-grid guess, replacement popup, focus change or font/settings
mutation. Unknown launch geometry disables fitting, and an unpresented terminal
retains the launch rectangle rather than using its bootstrap 80x24 grid.
Cell rounding can leave one spare row; native chrome remains visible and initial
size can settle after creation. JSON `position` still does not choose outer bounds
in direct mode; standalone UI/result positioning is unchanged.
See [sizing design and limits](content-sized-popup.md) and
[startup geometry cache](geometry-cache.md). Historical observations
below are not new evidence for this sizing route.

### Native chrome: accepted defaults, customization deferred (REQ-026)

User1a explicitly accepted Rex defaults **for now** and deferred per-float native
titlebar/border/corner/shadow customization. It is not implemented or credited as
REQ-026 completion. Rex's default native header/titlebar, border, corners and
shadow may remain visible. The inspected Rex API has no supported per-float knobs
for that customization; this is not a claim about all future Rex versions.

Lua owns native presentation policy, initial maximum bounds and launch. Go owns
content-height fitting and
the rendered menu header/footer, semantic terminal colors and palette matching;
those terminal spans do not style or remove native controls. No guessed/ignored
API fields, fake knobs, global Rex settings/user configuration changes, clipped
controls, font changes or painted borders substitute for native chrome support.
Task04 documents the accepted policy and supplied revision-bound observations,
not a native-feature implementation. The explicit user acceptance (`1a - yes,
accept it for now; 1b - yes`) refines the original blocker without deleting it;
see the [accepted design](../.tasks/designs/rex-whichkey-lua-float-launch.md).
Earlier blocked statements/API observations remain historical in
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
no GUI owner; it is not a physical GUI drag or pixel measurement.

Task03's `.mtn-tmp/p19/resume/task3-public-evidence.json` records actual PATH
`make run` on reviewed combined source and rebuilt main at `6564d07e`, with current
Go executable paths, default centered `.8/.3` bounds, root VT/HTML captures,
providers, owned rename and separate edited cancellation/dismissal. GNU Make's
custom-config `.75/.25` float had bounds `[.125,.375,.75,.25]`; its later popup
and original-window disappearance had an unobserved cause, not a controlled
cancellation pass. These are coordinator observations, not a new leaf runtime.
See [current acceptance](acceptance.md#phase19-accepted-policy-and-observed-make-route)
for exact source, commands, Rex IDs and artifacts.

User1b retains worker join, tty restoration, exact visible-float retirement and
original focus verification before effects; uncertainty suppresses dispatch.
Only the detached Go execution host may survive through the once-only effect
until Go exits/Rex retires it. Go cannot prove its own post-exit deletion, and
post-exit cleanup cannot suppress an already-issued effect. Exact `list_blocks`
checks prove only the observed cases, not universal cleanup.

Task06 refactors the Go renderer, palette decoder and terminal loop without
changing Lua sizing/native chrome. Fix2 `cacf7a9694c8ed912ff72d9f8ecfe86c31fa094c`
received fresh whole-diff APPROVED independent review, then integration into
observed main `7e6160f89c0025991712a751e51aa9b2abeaeeef`. Combined/main
`task build`, `task test` (unit/race) and full `task lint` (vet/golangci/compatible
AI) all exited 0. Automated regressions cover UTF-8/escape/report/paste/edit
framing, palette replies, pagination, confirmation, retained edits/providers,
bounded workers, and real current terminal/direct lifecycle orchestration in
isolated Darwin PTYs/executable CLI fixtures. The termios check excludes only
kernel-transient PENDIN; production restores the captured state unchanged.

Coordinator `task6-public-evidence.json` under `.mtn-tmp/p19/quality-fix/` records
actual candidate/main `make run`: default menu, full-ID/current-marker inventory,
one owned rename, separate edited cancel/root dismissal, native exit 0, exact
post-exit host absence and focus/tty/primary-screen restoration. Eight snapshots
showed detached host/no visible float with the old label before the rename;
these are case-specific ordering observations, not internal instrumentation.
The owned origin window was closed and the exact baseline restored. This docs-only
followup leaves approved/observed runtime and test source byte-identical, so those
observations apply without claiming a new post-doc live launch. Native completion
metadata and exact-doc-head review/combined-main gates remain coordinator followup,
not missing source2 runtime proof. See
[current verification](acceptance.md#current-verification-boundary).
No physical-key/reload, installed-profile, Nix build/activation, every-palette/
geometry/flags, timing SLA or all-platform result is claimed; terminal captures
are not native pixel/chrome certification.

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
