---
workflowKind: task
workflowLanguage: docs
workflowOrder: 3
workflowRequirements:
  - REQ-030
workflowDocuments:
  - ref: designs/rex-whichkey-action-audit.md
    hash: 'sha256:2d2d664eb4bbc5828b2f0d6d4a98346becfbbd03d90e8f9abc5a9874edc8b7a8'
  - ref: designs/rex-whichkey-selected-window-splits.md
    hash: 'sha256:1d5e6297352b97b84e3ba7730e586130b4103ddb4ad32ed03d35b6ac2ad6e4f2'
workflowOwnership:
  paths:
    - docs/gui-action-verification.md
  sharedPaths: []
  resources: []
  interfaces:
    - native-action-audit-3
title: Verify seven native GUI-only action outcomes
mtnCreateRetry:
  key: rex-whichkey-phase14-task3-v1
  intent: 3b03b5f0d4d2b1e22b21b2a2334d54f215d29bb407ce36a6c17911328047b56e
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T20:16:08.762Z'
dateModified: '2026-11-02T20:38:15.981Z'
tags:
  - task
type: task
id: b8310b3e-9e3b-4e4a-8cdc-a94e48f9b3eb
projects:
  - '[[phases/rex-whichkey-14-action-audit.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-14-02-resource-audit.md]]'
    reltype: FINISHTOSTART
---
# Verify nineteen native GUI-local action outcomes

REQ030 remains incomplete until all25 verified. After real resource probes, six resource effects can be observed (v/s/x/tx/wn/wx);19 remaining bindings need genuine GUI observability: tn, th/tl/tr/t1..t9, ws/wh/wl/wr, e/c. This corrects unverified initial18-resource/7-UI planning assumption, not scope waiver. Evidence .mtn-tmp/action-audit/observability-correction.md: client.tab.goto null result leaves server active_window_id unchanged; server field is not chosen app's selected tab. Native tn exited0 but no server window or session appeared; may be native UI/client-local tab, not certified working or conclusively broken.

Need actual selected tab/session, sidebar/palette, picker/rename modal and completed rename/new-tab behavior. client.inspect exposes metadata/actions/attached IDs only; terminal send-key drives block PTY, not native app dialogs. Investigate documented capability before pass. If none, BLOCKED recovery approved genuine GUI observation/control or explicit observed user evidence, not invented screenshot/action/args, not stale server focus/palette state or RPC acknowledgment. Do not open unclosable prompts or pretend user deferred requirement. No general tests/settings/font/config/reload/remote/Darwin trials, no destructive user resources. Depends on resource documentation; reserved ownership docs/gui-action-verification.md for future real evidence. Preserve independent native chrome blocker. Source repair can complete while this and entire phase/project stay incomplete.
