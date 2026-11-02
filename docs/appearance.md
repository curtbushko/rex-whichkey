# Menu appearance

The renderer assigns separate semantic roles to keys, icons and labels. These
roles live in [`theme.go`](../theme.go), not in a second Lua/Nix theme schema.
[`render.go`](../render.go) applies them to actual menu spans;
[`terminal.go`](../terminal.go) writes the styled frames and updates the theme
when terminal palette metadata arrives.

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

[Theme task 08-03](../.tasks/tasks/rex-whichkey-08-03-theme.md) records fresh source
review and coordinator compilation at `3e9a1fd`; those are not graphical
observations. The [assembly record](../.tasks/tasks/rex-whichkey-08-04-docs.md) and
[acceptance inventory](acceptance.md) index the coordinator's frozen-candidate
35-second idle/default/color/navigation/Escape gate, including actual VT and HTML
frames. A plain text screenshot or this source description alone proves no color
pass. The gate covers one existing Linux Rex session, not all palettes, fonts,
Darwin/remote deployments or configured Lua callbacks. No runtime check is run by
the documentation builder/reviewer.
