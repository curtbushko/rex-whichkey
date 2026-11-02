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
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T15:23:01.367Z'
dateModified: '2026-11-02T15:35:27.070Z'
tags:
  - task
type: task
id: 35f162bd-51fa-4ec7-aced-52e0954b068c
projects:
  - '[[phases/rex-whichkey-09-tab-return.md]]'
completedDate: '2026-11-02'
---
# Remove default Zoom and restore native popup on tab return

REQ-021/REQ-022; read approved designs/rex-whichkey-tab-return.md and exact original user correction quoted there. Own only defaults.go, lua/rex_whichkey/defaults/panes.lua, terminal.go, popup.go, README.md, docs/appearance.md, docs/launch-contract.md and docs/acceptance.md. Do not modify rex API wrappers without a precise coordinator ownership amendment, nor native.go/launch.go, user init/config, tests/gates/dependencies or .tasks. One concern is popup interaction reliability with default footprint cleanup; no new CLI/client-selection policy.

Remove z/Zoom from both real default generators and document current bindings without Zoom, retaining explicit custom pane.zoom support. Diagnose native popup tab-return. Existing coor evidence .mtn-tmp/tab-recovery/reproduction shows same native process/layer/content surviving simple server and app tab goto round-trip, so do not claim a crash or attachment-loss root cause without evidence. No native layer persistence option exists in captured new_layer API. Read actual terminal parser/palette/RunTerminal and popup lifetime callers; do not implement a token helper, unconditional redraw timer or duplicate layer.

A genuine foreground/focus-in redraw is recommended: standard DECF focus-reporting mode1004 saved/enabled/restored as part of terminal ownership, complete focus-in events redraw retained menu immediately without navigating or dispatching. Handle fragmented/multiple/OSC-interleaved focus events correctly, incomplete sequences must not leak keys, focus-out must not cancel or steal focus, non-focus escape/mouse/paste/CSI protection unchanged. Preserve width/zero-size/re-resize/NO_COLOR/ASCII/palette behavior and terminal restoration. Research source/docs (not live experiments) when uncertain. Any idle-liveness adjustment must preserve captured app binding for executable selection, configured expiry and safe cleanup; don't conflate user absence with origin deletion. Avoid unrelated refactor and report limitations honestly if actual GUI disappearance cannot be source-proven.

Update exact docs to distinguish implemented focus-in repaint, real no-argument caller, retained configured expiry and verification lineage. No premature full GUI/platform certificate or claims unsupported focus semantics guarantee a native app layer renderer. Acceptance must be independently traced through actual input decoder -> redraw -> renderer and canonical defaults; fresh reviewer covers mixed Go/Lua/docs and all criteria. Builder formats/static/whitespace only: no tests/TDD/general unit/race suites, Go/lint-backed builds, actual product/GUI/API/config/remote experiments, user config edits/reloads, or push. Coordinator owns compilation and bounded actual native tab round-trip/focus/nav/Escape/cleanup, recording remaining visual uncertainty.

Use implementation-only git -c core.hooksPath=/dev/null commits, freeze clean source head/tree, no managed writes/integration/agents. FULL protocol-v1 result to payloadFile and MTN_REX_RESULT_FILE; validation MUST be an ARRAY of actual passing {command,result:"pass"} objects, skipped checks go in limitations, not fabricated pass. Final concise pointer only; machine result must remain full protocol object including schemaVersion/eventId/kind/recordRef.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE1OjM1OjI3LjA1OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwOS10MDEtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiZmQ5MTE2MGU4NGYzOTQwMWY1MWQyYzZkMjE1OTg2YWQxZmNkNTNjOGM3M2M3OTFhYTY0MzhkMDI3OTJiOWY3YiJ9 -->
### 2026-11-02T15:35:27.059Z · summary · unspecified

{"acceptedScope":"Remove default Zoom and restore popup on tab return","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p09-t01-reviewer-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/tab-recovery/rex-whichkey-p09-t01-builder-a01/integration/integration.json",".mtn-tmp/tab-recovery/runtime/rex-whichkey-p09-t01-builder-a01/result.json"],"attempt":1,"eventId":"rex-whichkey-p09-t01-builder-a01/summary/1","files":["README.md","defaults.go","docs/acceptance.md","docs/appearance.md","docs/launch-contract.md","lua/rex_whichkey/defaults/panes.lua","popup.go","terminal.go"],"kind":"summary","limitations":["Native app tab-round-trip/root/submenu/focus-in repaint/dismissal/exit0/owned cleanup checked in Linux Rex. Server snapshots are not proof of remote app presentation or reproduction of the original graphical disappearance; no general runtime/platform certificate.","No project tests/TDD, destructive actions, configured Lua callbacks, remote/Darwin or configuration changes/reloads; fresh same-family review."],"recordRef":"tasks/rex-whichkey-09-01-tab-return.md","requirements":["REQ-021","REQ-022"],"schemaVersion":1,"specHash":"sha256:dac01498d83ae0c4fb515581056d97877eac4a4497674cf55a5acf8e92760622","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"4e03865091daff8a96144f878f959efddcc4d8b5","targetTree":"4cd9f3b8fbd551b872dd3069dddf78897621f5a2","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"},{"command":"task build; actual native popup without Zoom, root/submenu native-app tab-away/back and focus-in redraw, navigation/Escape/exit0/exact owned cleanup in existing Linux Rex session","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
