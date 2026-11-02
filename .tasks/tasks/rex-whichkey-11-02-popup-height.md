---
workflowKind: task
workflowLanguage: go
workflowOrder: 2
workflowRequirements:
  - REQ-025
workflowDocuments:
  - ref: designs/rex-whichkey-compact-presentation.md
    hash: 'sha256:10a72840e340a8bb53e744309fd518d564a9aac00a73364c3a554da06d05388f'
workflowOwnership:
  paths:
    - popup.go
    - terminal.go
    - popup_size.go
    - rex/layer.go
    - rex/block.go
    - rex/terminal.go
  sharedPaths: []
  resources: []
  interfaces:
    - owned-popup-adaptive-height
title: Fit actual popup layer to current content height
mtnCreateRetry:
  key: rex-whichkey-phase11-task2-v1
  intent: 8113a0a12a58f4e4c7c221e9673fbb2bf2844d2cfdfbe286233d4d3f43d7fa08
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T17:33:44.685Z'
dateModified: '2026-11-02T17:33:50.206Z'
tags:
  - task
type: task
id: e58fa5bc-47f0-41a6-9ec0-2f4f340dceb6
projects:
  - '[[phases/rex-whichkey-11-compact-presentation.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-11-01-render-frame.md]]'
    reltype: FINISHTOSTART
---
# Fit actual owned popup layer to current menu height

REQ-025, consumes approved/integrated task11-01 MenuFrame.PreferredHeight (natural current-width rows including all status/padding independent of viewport height). Own popup.go, terminal.go, popup_size.go (new), rex/layer.go, rex/block.go and rex/terminal.go (new) only. Read actual producer and approved designs/rex-whichkey-compact-presentation.md plus captured native API metadata. No config/mapping/schema/CLI flag/dependency/test/gate/managed/user-config changes.

Make actual layer height content-driven rather than old fixed60percent/fullheight viewport. Child actual renderer must publish current natural rows initially and on submenu/back/page/confirmation/width/font/resize changes; parent alone mutates exact owned layer. Add internal observer/publication plumbing preserving existing RunTerminal/RunPopupTerminal entry APIs unless a precise ownership amendment is needed. Preferred size hints are bounded data-only in existing private0700 transport, atomic writes, validated size/type/bounds/nofollow, influence only owned geometry. Add hint/temp names to exact cleanup. Don't let children replace config/action/args/identity or renew configured expiry/native parent lease.

Use installed native terminal size metadata (actual rows/columns and optional cell pixel dimensions) and normalized layer bounds, not guessed fixed font pixels. Respect original max rectangle/available viewport, safe finite arithmetic/clamping/position anchors, transient zero/missing sizes, and convergence/rounding. Root/submenu must grow as well as shrink, so don't derive preferred natural height solely from already clipped frame. Retain same layer/block, no creation retry/replacement, no GUI app lookup, focus steal or setting/font modifications. Parent existing bounded liveness watcher may refresh geometry; no extra daemon/framework. Avoid oscillation or unconditional repeated resize when unchanged. Original frame placement should not remain vertically centered within a tall box after fitting.

Preserve app-independent native display/cancellation/resource liveness, explicit app targeting only for native action, configured original-VM30second authenticated expiry/generation, FIFO parent cancellation, bounded operations/cleanup and exact-owned-only action-before-cleanup prohibition. Geometry failures/transients need safe behavior without silently dispatching after uncertain cleanup. Task3 hiding Rex controls is BLOCKED; do not invent chrome flags or global changes. Native API Resize already checks owned membership; do not weaken.

Leaf only formatting/static/whitespace/source metadata research; NO heavy/build/compiler-backed lint, tests/TDD/general/unit/race suites, product/GUI/API/remote experiments, user configuration/font/settings changes/reloads or push. Coordinator compiles exact combined source; final docs task gate actually measures compact root/submenu/back bounds/rows/status frames/stability/Escape cleanup in existing Linux Rex, not fonts/settings/destructive actions. Source-only git -c core.hooksPath=/dev/null commits, freeze head/tree, no integration/agents/mtn writes. Full protocol-v1 build copied to MTN_REX_RESULT_FILE, validation array only actual passes, skipped gates limitations; concise visible pointer.
