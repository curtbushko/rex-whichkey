---
workflowKind: task
workflowLanguage: go
workflowOrder: 1
workflowRequirements:
  - REQ-027
  - REQ-028
workflowDocuments:
  - ref: designs/rex-whichkey-first-frame-fit.md
    hash: 'sha256:e40a82729585e2a35de60126c4a3771c4c651d92e287f55c2808ff32d244e777'
workflowOwnership:
  paths:
    - render.go
    - layout.go
    - popup.go
    - popup_size.go
    - popup_initial.go
    - rex/layer.go
    - rex/block.go
    - rex/terminal.go
    - rex/geometry.go
    - README.md
    - docs/appearance.md
    - docs/launch-contract.md
    - docs/acceptance.md
  sharedPaths: []
  resources: []
  interfaces:
    - popup-initial-frame-geometry
title: Space status bars and pre-size initial popup frame
mtnCreateRetry:
  key: rex-whichkey-phase12-task01-v1
  intent: 68522f9d6d43608ad07c2746d2b5a6472d9cf6c81e33b808a92c4b7a3811361e
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T19:10:54.296Z'
dateModified: '2026-11-02T19:10:57.967Z'
tags:
  - task
type: task
id: e69c2e1d-755d-40c1-83af-31599f7830e1
projects:
  - '[[phases/rex-whichkey-12-first-frame-fit.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-11-01-render-frame.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-11-02-popup-height.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-11-04-appearance-docs.md]]'
    reltype: FINISHTOSTART
---
# Add status spacers and size the popup before first draw

REQ-027/028; read approved designs/rex-whichkey-first-frame-fit.md plus current captured API/rect/terminal metadata in .mtn-tmp/popup-initial-fit. Own render.go, layout.go, popup.go, popup_size.go, popup_initial.go (new), rex/layer.go, rex/block.go, rex/terminal.go, rex/geometry.go (new), README.md, docs/appearance.md, docs/launch-contract.md, docs/acceptance.md only. Preserve existing APIs where practical, flat Go module/canonical config; no user config, tests/gates/dependencies/mapping/Lua changes or managed writes.

Insert a blank spacer row between content and each top/bottom status border when sufficient room. Include ideal spacers in MenuFrame.PreferredHeight, dropping blank space before content/essential status hints on short viewports; no restored Page1/1 or extra footer row. Update comments/docs no-blank-padding claims. Preserve renderer widths/sanitization/NO_COLOR/ASCII/confirm/paging safety.

Compute initial layer rectangle and terminal startup rows/columns from renderer natural first-frame rows, actual origin placement/window-relative rect and validated current origin terminal measurement before NewLayer, rather than creating popupBounds maximum first. Captured session.view explicitly says leaf layout.rect is WINDOW-relative and floating leaf rect equals layer bounds; no extra multiplication by layer. Seed documented shell options.size rows/cols and optional cell metadata in the same creation RPC, so initial text cannot be centered in unrelated default80x24. Keep caller-compatible NewLayer API/wrapper or narrowly named options method; never claim current seed values as live measured geometry. Reuse strict native geometry parser/provenance where applicable, validate finite/zero/oversized/clamped dimensions and split/floating/moved origin placements; no global active-window guesses.

Original max bounds remain adaptation cap, not initial rectangle. Keep top/bottom/center/side anchors. Origin and new popup fonts/chrome may differ: document initial estimate and use conservative cell-sized padding/rounding, not magic fixed pixel fonts. Missing/unmeasured origin gets compact bounded fallback and native menu still works without app binding; no startup requirement of one app, no fetching selected-app actions. Existing measured adaptation handles later menu/font/view changes, only small startup correction should remain. No invisible/offscreen-first layer, delayed show/recreate/retry or cached cross-invocation settings. Exactly one creation attempted, owned-only resize/cleanup and private geometry hint validation unchanged. Preserve configured original-VM expiry/capabilities, FIFO parent and cleanup-before-action dispatch.

Native title/dock controls are still REQ026 BLOCKED outside this task; no guessed flags/global workaround or claim bars remove them. Docs must clearly distinguish this persistent gap, estimated precreation size, live later fitting and actual versus unexecuted verification. No self-staling review status or false GUI pixel certificate.

Leaf only format/static/whitespace/read-only source/docs/help research. NO tests/TDD/general/unit/race suites, compiler-backed/heavy/Go/lint builds, product/GUI/API/font/config/remote experiments or installs. Coordinator compiles exact reviewed candidate and actual promptly observed FIRST layer bounds/initial terminal geometry/status spacers/settled size/submenu/back/Escape/cleanup. No native/destructive selection trial. Source-only hook-disabled commits, clean frozen head/tree, no agents/integration/mtn writes/push. Full protocol-v1 build at payloadFile and MTN_REX_RESULT_FILE, validation array only real pass checks, skipped work limitations; concise final visible pointer. Do not complete phase11/project while native chrome gap remains.
