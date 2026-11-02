---
workflowKind: task
workflowLanguage: go
workflowOrder: 6
workflowRequirements:
  - REQ-030
workflowDocuments:
  - ref: designs/rex-whichkey-client-perform.md
    hash: 'sha256:94636573720ef04607582a3b78a83f2484a1533224f1356fbac92f5fb8131c51'
workflowOwnership:
  paths:
    - native.go
    - README.md
    - docs/launch-contract.md
    - docs/acceptance.md
  sharedPaths: []
  resources: []
  interfaces:
    - native-client-perform-dispatch
title: Dispatch native actions with explicit client.perform
mtnCreateRetry:
  key: rex-whichkey-phase14-task06-v1
  intent: 37ab29ad1be4490acb490d304407ebbcb893cef491362daeb98f9b16e3ae292d
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T22:13:48.440Z'
dateModified: '2026-11-02T22:13:50.259Z'
tags:
  - task
type: task
id: 5cada17b-86a6-441b-b54c-691a5b81151e
projects:
  - '[[phases/rex-whichkey-14-action-audit.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-14-05-numbered-tabs.md]]'
    reltype: FINISHTOSTART
---
# Use explicit client.perform for native actions

REQ030 bugfix, depends on Task14-05 completed. Exact latest design designs/rex-whichkey-client-perform.md. Observed exact CLI do NewTab no-op vs direct documented client.perform same app/ctx createswindow, same executable not version skew. Repair native.go only dispatch mechanism and accurate docs. Use a separate unscoped exact-server short-timeout Rex client for explicit client.perform. Preserve original action and json.RawMessage args, origincli/timeout; for ALL non-split actions include captured origin session_id/block_id as previous CLI intended, but pane.split MUST OMIT both so current chosen app focused pane remains destination. Don't silently broaden destructive actions, remove generation/origin checks, change numbered-tab exact-ID adapter, mutable child authority, configuredoriginalVM/expiry, late picker/noIdle/FIFO/whole signals, fixed dimensions/initial spacers. All UI/private files/leases gone before one dispatch, fresh budget only after choice. Selected action registration checked; exact origin/app generation recheck immediately before request; nil result discard validatesJSON; sanitize error and no uncertainty retry/CLI fallback.

Allowed native.go, README.md, docs/launch-contract.md, docs/acceptance.md. Existing native_tabs.go adapter remains unchanged unless compilation demands owned expansion (ask before exceeding). Remove dead CLI dispatch wrappers/imports if unused but metadata registry still invokes Rex executable legitimately. No unsupported API guessed/action args or CLI env overrides userworkaround. No new framework/deps/schema/config changes.

Leaves source/format/whitespace only, no compiler/build/lint-heavy/tests/TDD/general/unit/race/RPC/nativeUI/config/fonts/platform trials or managed writes/agents. Coordinator compiles frozen reviewed combined candidate and real native menu v/s/tn/x/tx/wn/wx on exact owned resources; all7 resulting effects mandatory before completion, source review not behaviorproof. Scope6alive full25 still incomplete unless later genuinelyverified; chrome remains blocked. Full protocol-v1 build head/tree/base/task/spec/attempt/dispatchEventId/outcome ready-for-review; validation ARRAY actual sourcepasses, skipped gates limitations; source commit hooksdisabled; copyFULL to MTN_REX_RESULT_FILE, final pointer only.
