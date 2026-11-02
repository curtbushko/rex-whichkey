---
workflowKind: design
title: First-frame compact sizing and status spacing
mtnCreateRetry:
  key: rex-whichkey-first-frame-design-v1
  intent: 56d89c0c6b9ab9d280ae14e57bed14299551f6c4449aa05e25729a0202bc4b58
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T19:10:45.242Z'
dateModified: '2026-11-02T19:10:50.632Z'
tags:
  - task
type: task
id: 33ab6cf3-7bcb-4026-b13e-0c021ab812db
completedDate: '2026-11-02'
---
# First-frame compact sizing and status spacing

User explicitly corrects visible large-then-small popup startup and requests spacer lines between menu content and status borders. User repeats native title/dock controls still visible; accepted REQ-026 remains unresolved in phase11, not waived or claimed fixed here.

REQ-027: One blank spacer row below top title/breadcrumb status border and one above bottom key-hint border when viewport permits. Natural preferred height includes both spacers; tiny viewports drop spacing before content or essential hints. Preserve single-page-indicator removal, real pagination, Unicode widths, ASCII/NO_COLOR, confirmation and menu state.

REQ-028: Calculate an initial content-sized normalized rectangle BEFORE session.new_layer. Never create the default native menu at old60percent height and rely on visible late shrinking. Seed the new terminal's initial columns/rows consistently so its first draw is not centered within an unrelated default80x24 terminal. Use exact origin window-relative block rect plus actual measured terminal rows/columns/cell geometry when available, canonical renderer natural rows and max-position bounds. Initial geometry is necessarily an estimate where native chrome/new-pane font differs; allow only bounded small measured corrections, do not claim pixel-perfect equivalence or change user fonts. Unmeasured/headless cases still open at a documented compact bounded fallback, not a fabricated live measurement or a GUI requirement. Preserve parent-owned later submenu/font adaptation, no GUI binding on display, strict native action targeting and configured lifetime/security/cleanup.

Verified native contracts: session.view layout.rect is ALREADY window-relative fractions, not layer-relative; captured floating bounds and leaf rect both [0.1,0.2,0.8,0.6]. Do not multiply them. Terminal.size provides rows/columns/cell pixels and owner/desired provenance. New shell options support size:{columns,rows,cell_width_px?,cell_height_px?}; values are startup terminal geometry, not a change to app font preferences. Actual popup font may differ from origin; explicit estimated cell-sized headroom/rounding can be used and documented, never guessed fixed font pixels. No hidden/offscreen/recreated layer trick, saved cross-invocation calibration/cache or new daemon/schema/framework.

REQ-026 recheck: current rex --help, API list/new_layer and client action registry expose no dedicated frameless popup/per-popup title/dock/header control; only pane.float. Native title/dock strip remains blocked by Rex capability. Do not invent layout flags, edit global settings/init, modify another repo or imply status bars hide native controls. Keep existing task11-03 and phase11/project blockers.

One coherent first-frame source/docs task depends on delivered phase11 renderer/height/docs tasks, NOT the incomplete phase11 milestone. This allows supported follow-up despite independent chrome blocker. Fresh independent frozen mixed Go/TUI/security/docs review; coordinator combined compilation and actual first-layer-snapshot plus terminal startup dimensions, root/spacer/status frames, later small stable fit/submenu/back, Escape/exit0/owned cleanup. Poll first layer promptly before one-second resize watcher, not only after final settling; large initial bounds are a failure even if eventually compact. Direct remote app pixels remain unobserved; startup bounds/seed/source ordering evidence is not a native chrome certificate. No general tests/TDD/unit/race suites, user fonts/settings/config reloads, destructive/native selections, remote/Darwin trial or push. Phase12 may complete if REQ027/028 pass; whole project/phase11 stay incomplete while REQ026 remains open.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE5OjEwOjQ5LjcyM1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9maXJzdC1mcmFtZS9kZXNpZ24tYXBwcm92YWwvMSIsImludGVudCI6IjczMDQ0ZDlmYjAyYzBmOWQ3ZTFkNmQ2ZmI0Y2M3ODBlNmU5ZjBiMjc3YWVkNzg3Mjk3ZTFiOTRmMTQ2YzE0ODQifQ -->
### 2026-11-02T19:10:49.723Z · approval · unspecified

{"approvedBy":"explicit user correction; delegated bounded implementation","attempt":1,"decision":"approved","decisionEvidence":"User requests spacer lines and rejects visible large-then-small startup, asks calculations before drawing. Initial natural-row/origin-geometry calculation and terminal seed required before creation, not only later adaptation. User repeats title/dock controls gap; existing REQ026 remains blocked without documented API, not waived. No tests/settings/fonts/destructive/remote authority inferred.","decisionRegisterEntry":{"decision":"Add content/status spacers and compute compact initial layer/terminal size before creation; later measured adaptation only refines. Preserve phase11 native chrome blocker and distinguish estimates from exact GUI pixels.","id":"REXWHICHKEY-20261102-141047","state":"Accepted"},"eventId":"rex-whichkey/first-frame/design-approval/1","kind":"approval","recordRef":"designs/rex-whichkey-first-frame-fit.md","requirements":["REQ-027","REQ-028"],"schemaVersion":1,"specHash":"sha256:e40a82729585e2a35de60126c4a3771c4c651d92e287f55c2808ff32d244e777"}
<!-- mtn:end-entry:v1 -->
