---
workflowKind: task
workflowLanguage: go
workflowOrder: 4
workflowRequirements:
  - REQ-032
workflowDocuments:
  - ref: designs/rex-whichkey-selected-window-splits.md
    hash: 'sha256:1d5e6297352b97b84e3ba7730e586130b4103ddb4ad32ed03d35b6ac2ad6e4f2'
workflowOwnership:
  paths:
    - native.go
    - README.md
    - docs/launch-contract.md
    - docs/acceptance.md
  sharedPaths: []
  resources: []
  interfaces:
    - native-selected-window-splits
title: Split chosen app current selected window
mtnCreateRetry:
  key: rex-whichkey-phase14-task04-v1
  intent: a04bccc3b58b818e52e50a33f522e4c4231658446a6865144650fb6cc6a1f654
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T20:38:16.882Z'
dateModified: '2026-11-02T20:38:18.717Z'
tags:
  - task
type: task
id: 691f0b64-4d0f-4e29-825c-5821070d795a
projects:
  - '[[phases/rex-whichkey-14-action-audit.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-14-01-action-target.md]]'
    reltype: FINISHTOSTART
---
# Native splits operate on chosen app's currently selected window

REQ032 explicit latest user correction. Read designs/rex-whichkey-selected-window-splits.md and existing action audit design with later supplement taking precedence. Current native.go always passes origin session/block to rex do, forcing splits into launcher origin even when chosen app selects another tab. Fix ONLY native pane.split v/right and s/down to use chosen app current selected tab/focused pane. Never split popup terminal, launch-origin unconditionally, another arbitrary app or server shared active_window_id.

Use documented client.perform with context omitted (block_id/session_id absent), via exact-server UNscoped Rex client or equally proven CLI environment isolation; generic scoped Client.Call injects origin session. Retain scoped origin/session/app-generation validation immediately before one dispatch; selected action registration, retained original args down/right and capability/confirmation unchanged. Do not silently alter destructive native close targets or configured originalVM/auth lifetime. Full cleanup BEFORE dispatch, short RPC context after user choice, signal cancellation/no-idle/FIFO, late explicit app choice and strict --client behavior intact. No retries/fallback on uncertain effects. Error returns sanitized. No new API guessed/dependency/service/schema.

Source ownership native.go, README.md, docs/launch-contract.md, docs/acceptance.md. No defaults arg changes, generic Rex adapter changes, Lua/config changes or UI sizing changes. Leaves source/format/whitespace only; NO tests/TDD/general/unit/race/compiler/build/heavy lint/RPC/native GUI/action/settings/fonts/platform trials. Coordinator compiles frozen reviewed combined candidate and verifies launcher A versus chosen app selected B, B two known panes, s/v split focused B pane in expected direction, A unchanged/other B anchor unchanged and exact owned cleanup. Also normal same-window split/late chooser/cancel/missing-selector gate. Server active_window_id NOT app-selected evidence. Resource audit and GUI19 still incomplete; do not claim every action/project done.

Full protocol-v1 build exact record/spec/attempt/base/head/tree/dispatch; outcome ready-for-review, validation [{command,result:"pass"}] actual source checks only; limitations skipped gates. Copy FULL object to MTN_REX_RESULT_FILE. Source commit hooks disabled, no collection mutations/integration/subagents. Final concise pointer.
