---
workflowKind: task
workflowLanguage: go
workflowOrder: 1
workflowRequirements:
  - REQ-021
  - REQ-022
workflowDocuments:
  - ref: designs/rex-whichkey-tab-return.md
    hash: 'sha256:a616c835e37588e3c673b289849b40573db242126626bcdc1597a9c2a76607c9'
workflowOwnership:
  paths:
    - defaults.go
    - lua/rex_whichkey/defaults/panes.lua
    - terminal.go
    - popup.go
    - README.md
    - docs/appearance.md
    - docs/launch-contract.md
    - docs/acceptance.md
  sharedPaths: []
  resources: []
  interfaces:
    - native-popup-focus-return
title: Remove default Zoom and restore popup on tab return
mtnCreateRetry:
  key: rex-whichkey-phase09-task01-v1
  intent: 6f164e041f6fd5dd4135c10068d17de8533569ec255361bf865f6d416fe30466
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T15:23:01.367Z'
dateModified: '2026-11-02T15:23:02.284Z'
tags:
  - task
type: task
id: 35f162bd-51fa-4ec7-aced-52e0954b068c
projects:
  - '[[phases/rex-whichkey-09-tab-return.md]]'
---
# Remove default Zoom and restore native popup on tab return

REQ-021/REQ-022; read approved designs/rex-whichkey-tab-return.md and exact original user correction quoted there. Own only defaults.go, lua/rex_whichkey/defaults/panes.lua, terminal.go, popup.go, README.md, docs/appearance.md, docs/launch-contract.md and docs/acceptance.md. Do not modify rex API wrappers without a precise coordinator ownership amendment, nor native.go/launch.go, user init/config, tests/gates/dependencies or .tasks. One concern is popup interaction reliability with default footprint cleanup; no new CLI/client-selection policy.

Remove z/Zoom from both real default generators and document current bindings without Zoom, retaining explicit custom pane.zoom support. Diagnose native popup tab-return. Existing coor evidence .mtn-tmp/tab-recovery/reproduction shows same native process/layer/content surviving simple server and app tab goto round-trip, so do not claim a crash or attachment-loss root cause without evidence. No native layer persistence option exists in captured new_layer API. Read actual terminal parser/palette/RunTerminal and popup lifetime callers; do not implement a token helper, unconditional redraw timer or duplicate layer.

A genuine foreground/focus-in redraw is recommended: standard DECF focus-reporting mode1004 saved/enabled/restored as part of terminal ownership, complete focus-in events redraw retained menu immediately without navigating or dispatching. Handle fragmented/multiple/OSC-interleaved focus events correctly, incomplete sequences must not leak keys, focus-out must not cancel or steal focus, non-focus escape/mouse/paste/CSI protection unchanged. Preserve width/zero-size/re-resize/NO_COLOR/ASCII/palette behavior and terminal restoration. Research source/docs (not live experiments) when uncertain. Any idle-liveness adjustment must preserve captured app binding for executable selection, configured expiry and safe cleanup; don't conflate user absence with origin deletion. Avoid unrelated refactor and report limitations honestly if actual GUI disappearance cannot be source-proven.

Update exact docs to distinguish implemented focus-in repaint, real no-argument caller, retained configured expiry and verification lineage. No premature full GUI/platform certificate or claims unsupported focus semantics guarantee a native app layer renderer. Acceptance must be independently traced through actual input decoder -> redraw -> renderer and canonical defaults; fresh reviewer covers mixed Go/Lua/docs and all criteria. Builder formats/static/whitespace only: no tests/TDD/general unit/race suites, Go/lint-backed builds, actual product/GUI/API/config/remote experiments, user config edits/reloads, or push. Coordinator owns compilation and bounded actual native tab round-trip/focus/nav/Escape/cleanup, recording remaining visual uncertainty.

Use implementation-only git -c core.hooksPath=/dev/null commits, freeze clean source head/tree, no managed writes/integration/agents. FULL protocol-v1 result to payloadFile and MTN_REX_RESULT_FILE; validation MUST be an ARRAY of actual passing {command,result:"pass"} objects, skipped checks go in limitations, not fabricated pass. Final concise pointer only; machine result must remain full protocol object including schemaVersion/eventId/kind/recordRef.
