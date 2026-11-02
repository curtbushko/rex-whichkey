---
workflowKind: task
workflowLanguage: docs
workflowOrder: 3
workflowRequirements:
  - REQ-030
workflowDocuments:
  - ref: designs/rex-whichkey-action-audit.md
    hash: 'sha256:2d2d664eb4bbc5828b2f0d6d4a98346becfbbd03d90e8f9abc5a9874edc8b7a8'
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
dateModified: '2026-11-02T20:19:29.930Z'
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
# Verify seven native GUI-only action outcomes

REQ030 completion needs actual observable behavior for ws(session picker), wh/wl(previous/next selected session view), wr/tr(session/tab rename dialog and successful rename), e(sidebar), c(palette). RPC acknowledgement or registration is not behavior proof. Current documented client.inspect only returns metadata/actions/attached session IDs, no chosen session/dialog/sidebar state; terminal key injection controls block TTY, not native app dialogs. Investigate installed documented API/capability evidence. If no native app observation/control exists, mark BLOCKED with exact gap, recovery approved GUI observation/control or explicit observed user evidence; do not open unclosable modal dialogs or invent hidden args/screenshots/actions. Do not claim user deferred requirement. Depends on completed resource matrix. No code ownership until recovery approved; no settings/font/config/remote/Darwin/destructive trials. Root chrome blocker stays unrelated/open. No whole phase/project completion until verified or explicit user changes scope.
