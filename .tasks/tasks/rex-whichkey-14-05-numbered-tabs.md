---
workflowKind: task
workflowLanguage: go
workflowOrder: 5
workflowRequirements:
  - REQ-030
workflowDocuments:
  - ref: designs/rex-whichkey-numbered-tabs.md
    hash: 'sha256:12ac821b41273e253346ac9c5f0776ec7e836a9fdba1e3287a197b9bd392a90f'
workflowOwnership:
  paths:
    - native.go
    - native_tabs.go
    - README.md
    - docs/launch-contract.md
    - docs/acceptance.md
  sharedPaths: []
  resources: []
  interfaces:
    - native-numbered-tab-intent
title: Repair native numbered tab routing
mtnCreateRetry:
  key: rex-whichkey-phase14-task05-v1
  intent: 567148d0bec4ea385b831ea92804ecf3ba68c1dbb63b66345ae8f2335b0e841c
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T21:57:19.106Z'
dateModified: '2026-11-02T21:57:20.955Z'
tags:
  - task
type: task
id: 2289fe92-71e9-40e9-a4bf-dcf915d2a577
projects:
  - '[[phases/rex-whichkey-14-action-audit.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-14-04-selected-window-splits.md]]'
    reltype: FINISHTOSTART
---
# Repair native numbered tab selection through exact window IDs

REQ030 bounded bugfix, depends on completed Task14-04. Read latest designs/rex-whichkey-numbered-tabs.md and original native action/selected-window designs, newest supplement wins. Actual real-menu t1..t9 fresh target matrix all acknowledge but fail expected display geometry; Previous and direct native last-tab positive controls work. Don't guess index0/string or claim NativeRex repaired. Implement ONLY Go-owned native client.tab.goto intent translation to documented session.select {session_id:origin.SessionID,window_id:exactParentResolvedWindowID}. Canonical defaults and Lua/configured dispatch unchanged.

Use original retained numeric Args strictly one-based/default1..9; no child action/arg authority. Fetch ordered origin-session windows after all UI cleanup, resolve index once/freeze ID, bound all operations, refuse nonexistent index or removed/moved window rather than fallback. Recheck exact origin/app attachment/connection generation and resolved window membership immediately before one dispatch. Metadata check must include actual session.select native registration; leave unsupported action behavior safe. Prefer existing typed Rex client/session/window APIs, not duplicate view schema/new dependency. Separate unscoped client for actual explicit session.select if generic scope would inject duplicate/context; never ambiguous current client/server. Parent-authoritative mapping intent compilation explicitly documented; original completion/confirmation remains immutable. Never retry uncertain API failure.

Preserve native split current selected app context fromTask04, all other destructive/configured targets, late target picker/cancel/no idle/FIFO/whole signals, fixed geometry/initial sizing, originalVM capability. Source ownership native.go, native_tabs.go(new private helper if useful), README.md, docs/launch-contract.md, docs/acceptance.md. No generic Rex adapter/defaultschema/Lua/Taskfile changes.

Leaves source/format/whitespace only: NO builds/compiler/heavy lint/TDD/general/unit/race/RPC/nativeGUI/API/config/fonts/platform trials or agents. Coordinator runs frozen reviewed combined build and real-menu all9 fresh-hidden-target exact-app rendering checks, bounded refusal/cleanup, not screenshot/pixel certificate. Task14-03 full all25 still incomplete. Source ready only doesn'tcountbehaviorpass. Full protocol-v1 build schemaVersion/eventId/recordRef/specHash/attempt/dispatchEventId/outcome ready-for-review/baseCommit/head/tree, validation [{command,result:"pass"}] actual checks, limitations skipped. Commit scoped source with hooksdisabled; no managed writes/integration. Copy FULL object to MTN_REX_RESULT_FILE; final concise pointer.
