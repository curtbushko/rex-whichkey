---
workflowKind: task
workflowLanguage: go
workflowOrder: 6
workflowRequirements:
  - REQ-030
workflowDocuments:
  - ref: designs/rex-whichkey-client-perform.md
    hash: 'sha256:94636573720ef04607582a3b78a83f2484a1533224f1356fbac92f5fb8131c51'
  - ref: designs/rex-whichkey-new-tab-intent.md
    hash: 'sha256:2f49a61c9959851b3a7c0bd45654d9079dc97d88d0d64177bad16150ed7ebee6'
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
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T22:13:48.440Z'
dateModified: '2026-11-02T22:42:19.768Z'
tags:
  - task
type: task
id: 5cada17b-86a6-441b-b54c-691a5b81151e
projects:
  - '[[phases/rex-whichkey-14-action-audit.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-14-05-numbered-tabs.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-02'
---
# Use explicit client.perform for native actions

REQ030 bugfix, depends on Task14-05 completed. Exact latest design designs/rex-whichkey-client-perform.md. Observed exact CLI do NewTab no-op vs direct documented client.perform same app/ctx createswindow, same executable not version skew. Repair native.go only dispatch mechanism and accurate docs. Use a separate unscoped exact-server short-timeout Rex client for explicit client.perform. Preserve original action and json.RawMessage args, origincli/timeout; for ALL non-split actions include captured origin session_id/block_id as previous CLI intended, but pane.split MUST OMIT both so current chosen app focused pane remains destination. Don't silently broaden destructive actions, remove generation/origin checks, change numbered-tab exact-ID adapter, mutable child authority, configuredoriginalVM/expiry, late picker/noIdle/FIFO/whole signals, fixed dimensions/initial spacers. All UI/private files/leases gone before one dispatch, fresh budget only after choice. Selected action registration checked; exact origin/app generation recheck immediately before request; nil result discard validatesJSON; sanitize error and no uncertainty retry/CLI fallback.

Allowed native.go, README.md, docs/launch-contract.md, docs/acceptance.md. Existing native_tabs.go adapter remains unchanged unless compilation demands owned expansion (ask before exceeding). Remove dead CLI dispatch wrappers/imports if unused but metadata registry still invokes Rex executable legitimately. No unsupported API guessed/action args or CLI env overrides userworkaround. No new framework/deps/schema/config changes.

Leaves source/format/whitespace only, no compiler/build/lint-heavy/tests/TDD/general/unit/race/RPC/nativeUI/config/fonts/platform trials or managed writes/agents. Coordinator compiles frozen reviewed combined candidate and real native menu v/s/tn/x/tx/wn/wx on exact owned resources; all7 resulting effects mandatory before completion, source review not behaviorproof. Scope6alive full25 still incomplete unless later genuinelyverified; chrome remains blocked. Full protocol-v1 build head/tree/base/task/spec/attempt/dispatchEventId/outcome ready-for-review; validation ARRAY actual sourcepasses, skipped gates limitations; source commit hooksdisabled; copyFULL to MTN_REX_RESULT_FILE, final pointer only.


## Latest accepted refinement (takes precedence over generic original-action dispatch for tn)

# Native New Tab intent via explicit owned creation (refines transport-only repair)

REQ030 all-action functionality, user continue and delegated bounded implementation. Runtime candidate a01 compiled/reviewed but FAILED tn resulting-effect gate:6of7 effects only. Do not merge as fixed or relabel failed gate. Direct API and CLI behavior varied with context; same executable confirmed, tracing original candidate shows actual client.perform client.tab.new {} with correct exact app/session/block/origin. Real interactive zsh caller with valid OSC PWD also no new tab. Therefore transport-only hypothesis insufficient; do not blame missing GUI, executable version, guessed args or shell fixture without evidence.

Refine only native client.tab.new intent to documented explicit resource adapter: parent validates original retained action/empty arguments, captures original source session, requires chosen app actual session.select action registration before mutating, resolves scoped session, rechecks origin/app generation immediately before one Session.NewWindow default terminal creation. Freeze returned exact ID. Recheck original origin/chosen generation/window membership before one client.perform session.select with explicit session_id/window_id args. Existing typed Rex.Session.NewWindow and Window.Inspect are authoritative; no duplicate native schema or unsupported API. All menu/chooser UI/files/leases gone before FIRST resource mutation. Parent alone derives creation/focus intent; child never replaces action/args. No duplicate create, retry, uncertain fallback to client.tab.new or compensating automatic close. If creation result uncertain, refuse/report possible effect without retry. If creation succeeded but activation fails, preserve created tab and return a sanitized actionable error including its exact ID, not fake rollback or silent recreation. Original destructive targets and split selected-app context untouched; configured Lua/native default schema unchanged. This deliberately translates Go built-in New Tab intent rather than claiming native client.tab.new executed; document honestly as for numeric adapter.

Do not expand to other UI features. Explicit client.perform transport repair for remaining native actions stays, same original action/args and context (except split), no idle/signal/cleanup/auth regressions. Original mandatory all7 resource-effect gate retained; real tn must create exactly one tab in captured own session. Source-only leaves, no test/build/runtime or managed writes. Fresh attempt/review/combined gate mandatory; earlier failed a01/artifacts retained. Remaining native modal verification and chrome still incomplete.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDIyOjQyOjE5Ljc1NloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxNC10MDYtYnVpbGRlci1hMDIvc3VtbWFyeS8yIiwiaW50ZW50IjoiM2ZkNmE0ZGQ0OTIxYWM5Yzk1YjY4YmQwNGUxOWUwMDJjY2U2N2I4MDk1MTAxZDc3MWQ0OTM5ODI4ZGIyMWNjMSJ9 -->
### 2026-11-02T22:42:19.756Z · summary · unspecified

{"acceptedScope":"Dispatch native actions with explicit client.perform","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p14-t06-reviewer-a02/review/2 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/action-audit/rex-whichkey-p14-t06-builder-a02/integration/integration.json",".mtn-tmp/action-audit/resource-runtime/rex-whichkey-p14-t06-builder-a02/result.json"],"attempt":2,"eventId":"rex-whichkey-p14-t06-builder-a02/summary/2","files":["README.md","docs/acceptance.md","docs/launch-contract.md","native.go"],"kind":"summary","limitations":["Actual v/s/tn/x/tx/wn/wx all7 resource outcomes checked after explicit transport change, not all25 or native GUI pixels; no general suites/user-resource/settings/config/font/remoteOS/callback trials."],"recordRef":"tasks/rex-whichkey-14-06-client-perform.md","requirements":["REQ-030"],"schemaVersion":1,"specHash":"sha256:df8d9e1f72774a8a4a6476a2fbc58517dfe0a8cfc6e288bbbef8e6025ecc16e4","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"f28531ec574ed3011d680e71d42979f84fc2d54b","targetTree":"6d8b503c4bec5bb6dd79b4d437ee01cdb4e1173e","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"},{"command":"task build; actual real menu v/s/tn/x/tx/wn/wx all7 resource effects on exact own fixtures, UI cleanup before action","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
