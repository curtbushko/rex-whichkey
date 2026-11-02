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
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T19:37:56.676Z'
dateModified: '2026-11-02T19:48:20.078Z'
tags:
  - task
type: task
id: 86700679-fce7-4107-a478-13d1f07f3c08
projects:
  - '[[phases/rex-whichkey-13-fixed-navigation.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-12-01-first-frame-fit.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-02'
---
# Keep popup bounds stable through submenu navigation

REQ-029; read approved designs/rex-whichkey-fixed-navigation.md. Own popup.go, popup_size.go (retirement), terminal.go, popup_initial.go, README.md, docs/appearance.md, docs/launch-contract.md and docs/acceptance.md. Existing ignored .trash is available for preserving retired source; do not commit it. No render/layout/config/schema/Lua/gate/test/dependency/.tasks/user settings changes without precise ownership amendment.

Remove popup production live sizing state/callback, preferred-height hint publisher/file plumbing and unused private terminal observer rather than leaving a dormant resize path. Popup creation still calls precomputed initial geometry/terminal seed once. Navigation/focus/paging/idle must never mutate layer bounds; larger menus fit/paginate in current cells. Font/window SIGWINCH and focus-in still redraw normally but not resize layer. Keep generic Rex public Resize/Bounds/TerminalSize/geometry adapters as appropriate; do not edit/delete those outside scope. Move unused popup_size.go to ignored .trash (preserve artifact) if all references removed.

Retain exact origin geometry and validated measured/fallback initial compact size. Account conservatively in cells for root spacing without later correction: preceding .mtn-tmp/popup-initial-fit/runtime/rex-whichkey-p12-t01-builder-a02/result.json showed initial0.1105/root later0.1312. Keep initial root small and properly seeded, no giant all-submenu maximum, fixed pixel-font guess or cached calibration. Do not fabricate live measurement from seed. Preserve native app-independent display/escape/no-idle/FIFO, configured original-VM30second authenticated expiry/generation, exact retained selection/confirmation/target validation, bounded resource liveness watcher and owned cleanup before dispatch. No uncertain retries, new service or geometry IPC.

Docs update old per-submenu-grow/shrink claims as superseded by explicit latest user correction, preserve historical real evidence versus current fixed bounds. Explain initial estimate, possible natural external terminal geometry changes/pagination, status spacers and unresolved native title/dock chrome. Do not claim project complete or supported native hiding option.

Leaf source/format/whitespace only: NO heavy/Go/compiler-backed lint/build, tests/TDD/general/unit/race, product/GUI/API/font/config/remote experiments or installs. Coordinator compiles exact reviewed combined source and checks promptly observed compact first bounds/seed, exact same rectangle across root/Tabs/paging/back/idle, root spacer rows and Escape/exit0/owned cleanup. No native/destructive selection. Source-only git -c core.hooksPath=/dev/null commits; freeze head/tree, no managed writes/agents/integration/push. Full protocol-v1 build to payloadFile and MTN_REX_RESULT_FILE; validation array actual passes, skipped gates limitations; final visible pointer only.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE5OjQ4OjIwLjA2N1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxMy10MDEtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiNzA0ODFiNTA4MGRkYjY2NjI3N2ZlN2M2YmMyYTFjOWY3N2JiOTcwNGEzMmRiOWZmYmM4ZGNlZGRlNDljMzJmOCJ9 -->
### 2026-11-02T19:48:20.067Z · summary · unspecified

{"acceptedScope":"Keep popup bounds fixed through navigation","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p13-t01-reviewer-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/popup-fixed-navigation/rex-whichkey-p13-t01-builder-a01/integration/integration.json",".mtn-tmp/popup-fixed-navigation/runtime/rex-whichkey-p13-t01-builder-a01/result.json"],"attempt":1,"eventId":"rex-whichkey-p13-t01-builder-a01/summary/1","files":["README.md","docs/acceptance.md","docs/appearance.md","docs/launch-contract.md","popup.go","popup_initial.go","popup_size.go","terminal.go"],"kind":"summary","limitations":["Compact first bounds/status spacers/exactly unchanged root-submenu-back-idle rectangle/Escape/cleanup actually checked. Native title/dock controls remain unresolved, no user deferral/project completion/settings changes or remote GUI certificate.","No project tests/TDD, destructive actions, configured Lua callbacks, remote/Darwin or configuration changes/reloads; fresh same-family review."],"recordRef":"tasks/rex-whichkey-13-01-fixed-navigation.md","requirements":["REQ-029"],"schemaVersion":1,"specHash":"sha256:ef8bc1c6f61dd447a38797e0d1ef78efd2463709666295614893ace2110f8173","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"eb9d2c719d8cd8f6fdf871c0a255f5a2e6f303d5","targetTree":"2a8f3b1ee09076a8abd232cb161715c37800f073","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"},{"command":"task build; actual compact initial layer, EXACT same rectangle across root/Tabs/paging/back/idle, root spacers, Escape/exit0/owned cleanup","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
