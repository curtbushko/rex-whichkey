---
workflowKind: task
workflowLanguage: docs
workflowOrder: 2
workflowRequirements:
  - REQ-030
workflowDocuments:
  - ref: designs/rex-whichkey-action-audit.md
    hash: 'sha256:2d2d664eb4bbc5828b2f0d6d4a98346becfbbd03d90e8f9abc5a9874edc8b7a8'
  - ref: designs/rex-whichkey-selected-window-splits.md
    hash: 'sha256:1d5e6297352b97b84e3ba7730e586130b4103ddb4ad32ed03d35b6ac2ad6e4f2'
workflowOwnership:
  paths:
    - docs/actions.md
  sharedPaths: []
  resources: []
  interfaces:
    - native-action-audit-2
title: Verify resource actions and document matrix
mtnCreateRetry:
  key: rex-whichkey-phase14-task2-v1
  intent: 09ff9cce860438cc474f709a2134cfe2836b2dbfd36a2c810aedbf7e2c33a891
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T20:16:04.253Z'
dateModified: '2026-11-02T20:38:19.628Z'
tags:
  - task
type: task
id: 5fde9b62-cab4-4c1d-9561-296296780fe1
projects:
  - '[[phases/rex-whichkey-14-action-audit.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-14-01-action-target.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-14-04-selected-window-splits.md]]'
    reltype: FINISHTOSTART
---
# Audit every native default honestly and document observed effects

REQ030, depends on completed late-targeting repair and later selected-window split correction. User explicitly requests verification of all25. Actual evidence disproved original planning assumption that shared server active_window_id certifies selected-app tab navigation. It does not (tab-navigation-before/after/probe.json); current app tab state is client-local/unobservable through installed API. Later selected-window design and observability-correction.md refine earlier18/7 partition without waiving any verification.

Coordinator native resource-runtime/frozen-a01 verified effects for six bindings v s x tx wn wx; tn exited0 but no new server window/session was observed, so UNVERIFIED rather than passing or conclusively broken. Native selected-window s/v must be rerun after newest explicit user correction; original launcher-only split evidence is historical, not desired selected-window proof. Other19 bindings tn th tl tr t1..t9 ws wh wl wr e c require native app observation; registration/RPC receipts do not count as behavioral pass.

Builder documents coordinator's frozen final-source evidence in docs/actions.md: all25 binding/action/args/contracts on both live apps, expected result, actual source revision/observations/status/evidence, strict distinction verified effect versus acknowledged/unobserved/pending. Source audit confirms matrix matches native defaults/Lua defaults and directions correct. No fake full all-action/phase/project completion; GUI-only task remains blocked separately and native chrome remains blocked. No live trials/build/tests/compiler-heavy checks by leaves; coordinator only disposable IDs, never user/worker/coordinator resources or settings/fonts/config. Full frozen independent review/combined compilation for documentation completion, no remote GUI certificate.
