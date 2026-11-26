# Content-sized direct popup

Accepted decision: **REXWHICHKEY-20261126-092545**, indexed in the
[decision register](decisions.md). The user's approval includes submenu resizing.
This replaces the current direct popup's fixed-height navigation behavior;
historical fixed-bounds decisions/evidence are not rewritten as new results.

## Behavior and implementation

Lua still creates exactly one centered floating terminal. `width` is fractional
window width; `height` is the maximum fractional height (defaults `.8` / `.2`).
A matching [cached root height](geometry-cache.md) seeds the initial rectangle;
the configured maximum is passed to Go separately and remains the submenu cap.
Go retains its width, centers its height, and fits the same native layer to the
current menu's natural terminal rows, including the footer, optional heading and
their spacers, plus one outer blank row above the menu. Instructions occupy the
last row, with no bottom padding. Status rows are borderless; an absent heading
reserves no additional space. Padding yields to content/navigation in tiny viewports.
Root, submenu, back, confirmation, rename and provider updates use the same path.
Long menus paginate at the cap; Tab does not shrink a partially filled final page.
Within a direct popup, content stays at the top and instructions anchor to
the final row, with a reserved separation row when space permits. Any extra
viewport rows stay between content and instructions. This does not inflate the
preferred height or cached startup hint.
No JSON schema change, extra process, focus operation or native-chrome change is
introduced. Standalone `menu` remains result-only and never resizes Rex resources.

`RenderMenu.PreferredHeight` supplies natural rows at the actual PTY width.
`terminal_sizing.go` coalesces updates in a single worker, independently of input.
It confirms that the exact popup has a live sizing owner whose desired grid
matches the own-PTY grid. Rex's initial ownerless 80x24 terminal is not a measure
of the native window. Missing launch bounds disables fitting; ownerless or stale
grids wait without mutation, preserving the launch cap until presented.

`popup_sizing.go` estimates fractional row height from the own popup, then uses
two acknowledged bounds/row samples to account for native chrome outside the PTY.
A half-row allowance avoids clipping at cell boundaries. One spare row is allowed;
corrections wait for PTY acknowledgement and are capped at six per content/width
change. Screen/font changes reflow and recalibrate. Native minimum sizing can
prevent an exact fit; bounded settling, not pixel-perfect geometry, is promised.
A cold or stale-cache first rectangle can be larger while the own terminal is
measured; a valid warm hint reduces that initial adjustment.

`rex.Layer.Resize` validates exact sole-block ownership, placement, normal state
and expected bounds immediately before calling documented `session.resize_layer`
once, then verifies fresh bounds. Public display IDs cannot retarget it. It never
sets focus or layer order. A manual geometry change or uncertain mutation stops
the menu rather than chasing or retrying it. The worker cancels/joins before tty
restoration and popup retirement. A failed in-flight resize suppresses any selected
effect; cancelling an unfinished read-only size query does not.

## Alternatives and tradeoffs

- A smaller fixed fraction still wastes space on large displays and can clip on
  small ones. A fixed row cap alone leaves short submenus oversized.
- Sizing from the origin pane breaks with splits, different font settings or
  display ownership. Only the popup's own presented terminal is used.
- Using only a proportional row estimate ignores native chrome. A second own-PTY
  acknowledgement calibrates that overhead without inventing native pixel APIs.
- Synchronous CLI resizing blocks keyboard input. A joined worker keeps keyboard
  handling independent while preserving cleanup ordering and error propagation.

## Verification scope

Automated tests cover fitting/shrink/grow/cap, PTY acknowledgement and rounding,
screen-width changes, worker coalescing/join/error paths, live sizing ownership,
exact resize arguments, stale/shared/tiled/missing geometry, readback and uncertain
mutation without retry. Existing lifecycle tests retain their own scope. Live
checks and gate results belong to [acceptance](acceptance.md), not this design.
No installed-profile, physical-key, remote, all-platform or pixel-perfect claim
follows from implementing this design.
