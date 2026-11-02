---
workflowKind: task
workflowLanguage: go
workflowOrder: 1
workflowRequirements:
  - REQ-018
workflowDocuments:
  - ref: designs/rex-whichkey-popup-polish.md
    hash: 'sha256:ce42e479546a1181ce61138f0f5d37c88e5ed18ae3b6c39987fca90d89aee8f0'
workflowOwnership:
  paths:
    - native.go
    - popup.go
    - cmd/main.go
  sharedPaths: []
  resources: []
  interfaces:
    - native-configured-popup-lifetime
title: Keep native popup open while idle
mtnCreateRetry:
  key: rex-whichkey-phase08-task-1-v1
  intent: 66de1a9ec9d67a47e3c3e78c515fbe27385c60345c3d9d1ed79eaf4dfd5ca57f
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T13:17:03.007Z'
dateModified: '2026-11-02T13:17:03.929Z'
tags:
  - task
type: task
id: 6fde0105-e5c9-4e32-a44f-899853173051
projects:
  - '[[phases/rex-whichkey-08-popup-polish.md]]'
---
# Correct native popup idle lifetime and configured expiry

REQ-018. Read current approved polish design, native.go, runPopupLifecycle/RunPopupTerminal in popup.go and popup flags in cmd/main.go. Current native request has Expires=now+30 and UI deadline5seconds earlier. User reports expected idle expiry as a child-exit failure. Implement actual indefinite native interaction, not a longer fixed timer or catch-all swallowed failure.

Own native.go, popup.go and cmd/main.go only. Keep the flat API and existing modes compatible; no dispatcher/Lua/config, tests/gates/dependencies or managed edits. Refactor lifetime policy explicitly so only trusted native foreground invocation can request no interaction deadline; configured decoded Lua handoffs must still enforce canonical30-second expiry, one-shot/config/origin security and never renew from user/child data. A zero/absent deadline must not silently relax configured completion. The child and parent need a genuine no-fixed-native-deadline path. Preserve cancellation/signal/terminal/process-loss, one action snapshot/confirmation, exact owned popup/private file cleanup before dispatch and bounded RPC/cleanup/action operations. Avoid orphan indefinite processes or a daemon.

Expected configured bridge expiry should return safe cancellation/expiry rather than generic popup child failure; no stale selection execution. Unexpected child nonzero/status failure remains an actionable error, not success. Native idle should remain usable beyond35seconds and dismiss normally. Existing explicit deadlines/private commands remain correctly validated. Do not remove actual origin focus restoration when the user asks to remove the Focus menu group.

Builder formats/static-checks only: no heavy slot, project tests/TDD/product execution/live/GUI/config/remote probes, gate changes or push. Fresh Go review must trace native versus configured lifetime invariants, cancellation/cleanup paths, expected versus unexpected exits and real callers. Coordinator compiles reviewed combined candidate; final assembly task performs the user-authorized actual35second native idle/navigation/Escape check before phase completion. No worker agents, managed writes or integration. Full protocol file plus concise final pointer per supplied policy.
