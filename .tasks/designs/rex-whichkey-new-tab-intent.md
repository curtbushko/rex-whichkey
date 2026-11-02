---
workflowKind: design
title: Explicit native New Tab resource intent
mtnCreateRetry:
  key: rex-whichkey-new-tab-intent-design-v1
  intent: ab5ef110df307c6cce211e5cd7588c11a2ecf7d10b98f17a2555a6250e674a24
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T22:31:03.259Z'
dateModified: '2026-11-02T22:31:08.720Z'
tags:
  - task
type: task
id: c2afa6be-7c3c-4993-b77e-4e164382f603
completedDate: '2026-11-02'
---
# Native New Tab intent via explicit owned creation (refines transport-only repair)

REQ030 all-action functionality, user continue and delegated bounded implementation. Runtime candidate a01 compiled/reviewed but FAILED tn resulting-effect gate:6of7 effects only. Do not merge as fixed or relabel failed gate. Direct API and CLI behavior varied with context; same executable confirmed, tracing original candidate shows actual client.perform client.tab.new {} with correct exact app/session/block/origin. Real interactive zsh caller with valid OSC PWD also no new tab. Therefore transport-only hypothesis insufficient; do not blame missing GUI, executable version, guessed args or shell fixture without evidence.

Refine only native client.tab.new intent to documented explicit resource adapter: parent validates original retained action/empty arguments, captures original source session, requires chosen app actual session.select action registration before mutating, resolves scoped session, rechecks origin/app generation immediately before one Session.NewWindow default terminal creation. Freeze returned exact ID. Recheck original origin/chosen generation/window membership before one client.perform session.select with explicit session_id/window_id args. Existing typed Rex.Session.NewWindow and Window.Inspect are authoritative; no duplicate native schema or unsupported API. All menu/chooser UI/files/leases gone before FIRST resource mutation. Parent alone derives creation/focus intent; child never replaces action/args. No duplicate create, retry, uncertain fallback to client.tab.new or compensating automatic close. If creation result uncertain, refuse/report possible effect without retry. If creation succeeded but activation fails, preserve created tab and return a sanitized actionable error including its exact ID, not fake rollback or silent recreation. Original destructive targets and split selected-app context untouched; configured Lua/native default schema unchanged. This deliberately translates Go built-in New Tab intent rather than claiming native client.tab.new executed; document honestly as for numeric adapter.

Do not expand to other UI features. Explicit client.perform transport repair for remaining native actions stays, same original action/args and context (except split), no idle/signal/cleanup/auth regressions. Original mandatory all7 resource-effect gate retained; real tn must create exactly one tab in captured own session. Source-only leaves, no test/build/runtime or managed writes. Fresh attempt/review/combined gate mandatory; earlier failed a01/artifacts retained. Remaining native modal verification and chrome still incomplete.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDIyOjMxOjA3LjgxNloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9hY3Rpb24tYXVkaXQvbmV3LXRhYi1pbnRlbnQtYXBwcm92YWwvMSIsImludGVudCI6ImEyNjI1YmI3Y2NhOTNlNzcwY2UzMzcxYmQ4MWRkNTIwMTUwMTJhODNiOGM4ZjcwOWFhNDE4MDViMDAwMTVjNzMifQ -->
### 2026-11-02T22:31:07.816Z · approval · unspecified

{"approvedBy":"Explicit user all-action verification and continuing API work; delegated bounded bugfix choices","attempt":2,"decision":"approved","decisionEvidence":"a01 reviewed candidate failed actual tn gate despite explicit original RPC; retain failure. Use documented typed session.NewWindow then exact ID session.select for retained native NewTab intent, no blind retry/rollback, preserve successful-created tab if activation fails. Full all25 still required.","decisionRegisterEntry":{"authority":"Delegated bounded implementation","decision":"Native NewTab intent compiles to one explicit source-session default-window creation and one exact-app activation, with checks before each, no retries or uncertain fallback.","id":"REXWHICHKEY-20261102-173105","rationale":"Transport-only candidate did not meet unchanged7-effect gate.","refines":"REXWHICHKEY-20261102-171343","state":"Accepted"},"eventId":"rex-whichkey/action-audit/new-tab-intent-approval/1","kind":"approval","recordRef":"designs/rex-whichkey-new-tab-intent.md","relatedDecisions":[{"authority":"Delegated bounded implementation, not claim of user personally choosing wire transport","decision":"Native leaves use explicit documented client.perform, original args/action and explicit captured context except split; numbered exact-ID adapter/configuredLua unchanged.","id":"REXWHICHKEY-20261102-171343","rationale":"Require observable resource results, not opaque CLI success.","refinedBy":"REXWHICHKEY-20261102-173105","refinement":"Transport repair retained but native NewTab needs explicit resource adapter; no claim transport alone fixed it.","state":"Refined"}],"schemaVersion":1,"specHash":"sha256:2f49a61c9959851b3a7c0bd45654d9079dc97d88d0d64177bad16150ed7ebee6"}
<!-- mtn:end-entry:v1 -->
