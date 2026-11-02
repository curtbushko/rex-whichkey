---
workflowKind: task
workflowLanguage: go
workflowOrder: 1
workflowRequirements:
  - REQ-029
workflowDocuments:
  - ref: designs/rex-whichkey-fixed-navigation.md
    hash: 'sha256:7fbb1596666f9acf0ddb56ecf00ff29c09c8ae26781577268138e6a4ca05177e'
workflowOwnership:
  paths:
    - popup.go
    - popup_size.go
    - terminal.go
    - popup_initial.go
    - README.md
    - docs/appearance.md
    - docs/launch-contract.md
    - docs/acceptance.md
  sharedPaths: []
  resources: []
  interfaces:
    - fixed-popup-navigation-bounds
title: Keep popup bounds fixed through navigation
mtnCreateRetry:
  key: rex-whichkey-phase13-task01-v1
  intent: eec7ff99a037e85303f2ee5a994c31f6732f322f3790ea49cf0539d8a32d8d43
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T19:37:56.676Z'
dateModified: '2026-11-02T19:37:58.514Z'
tags:
  - task
type: task
id: 86700679-fce7-4107-a478-13d1f07f3c08
projects:
  - '[[phases/rex-whichkey-13-fixed-navigation.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-12-01-first-frame-fit.md]]'
    reltype: FINISHTOSTART
---
# Keep popup bounds stable through submenu navigation

REQ-029; read approved designs/rex-whichkey-fixed-navigation.md. Own popup.go, popup_size.go (retirement), terminal.go, popup_initial.go, README.md, docs/appearance.md, docs/launch-contract.md and docs/acceptance.md. Existing ignored .trash is available for preserving retired source; do not commit it. No render/layout/config/schema/Lua/gate/test/dependency/.tasks/user settings changes without precise ownership amendment.

Remove popup production live sizing state/callback, preferred-height hint publisher/file plumbing and unused private terminal observer rather than leaving a dormant resize path. Popup creation still calls precomputed initial geometry/terminal seed once. Navigation/focus/paging/idle must never mutate layer bounds; larger menus fit/paginate in current cells. Font/window SIGWINCH and focus-in still redraw normally but not resize layer. Keep generic Rex public Resize/Bounds/TerminalSize/geometry adapters as appropriate; do not edit/delete those outside scope. Move unused popup_size.go to ignored .trash (preserve artifact) if all references removed.

Retain exact origin geometry and validated measured/fallback initial compact size. Account conservatively in cells for root spacing without later correction: preceding .mtn-tmp/popup-initial-fit/runtime/rex-whichkey-p12-t01-builder-a02/result.json showed initial0.1105/root later0.1312. Keep initial root small and properly seeded, no giant all-submenu maximum, fixed pixel-font guess or cached calibration. Do not fabricate live measurement from seed. Preserve native app-independent display/escape/no-idle/FIFO, configured original-VM30second authenticated expiry/generation, exact retained selection/confirmation/target validation, bounded resource liveness watcher and owned cleanup before dispatch. No uncertain retries, new service or geometry IPC.

Docs update old per-submenu-grow/shrink claims as superseded by explicit latest user correction, preserve historical real evidence versus current fixed bounds. Explain initial estimate, possible natural external terminal geometry changes/pagination, status spacers and unresolved native title/dock chrome. Do not claim project complete or supported native hiding option.

Leaf source/format/whitespace only: NO heavy/Go/compiler-backed lint/build, tests/TDD/general/unit/race, product/GUI/API/font/config/remote experiments or installs. Coordinator compiles exact reviewed combined source and checks promptly observed compact first bounds/seed, exact same rectangle across root/Tabs/paging/back/idle, root spacer rows and Escape/exit0/owned cleanup. No native/destructive selection. Source-only git -c core.hooksPath=/dev/null commits; freeze head/tree, no managed writes/agents/integration/push. Full protocol-v1 build to payloadFile and MTN_REX_RESULT_FILE; validation array actual passes, skipped gates limitations; final visible pointer only.
