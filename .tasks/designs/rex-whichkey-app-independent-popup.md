---
workflowKind: design
title: Native popup display independent of app target
mtnCreateRetry:
  key: rex-whichkey-app-independent-design-v1
  intent: 607b2da4a10f85f5f31cf192fa32e315dcf029e782f69e1eded6da206c1ddd0b
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T16:14:28.529Z'
dateModified: '2026-11-02T16:14:33.891Z'
tags:
  - task
type: task
id: 650ea70a-859e-4a64-bd98-92e4ae26cdff
completedDate: '2026-11-02'
---
# Native popup display independent of graphical app selection

## User authority and requirement
User explicitly says please fix after explanation that app targeting is incorrectly coupled to popup launch. REQ-023: the basic native popup must open, navigate and dismiss using only the captured execution endpoint/session/block, with no requirement for exactly one live graphical app or a native app action registry. An optional syntactically valid exact client selector must not require a live app merely to display/dismiss the native menu. Only an executable native selection that needs a graphical app resolves and validates its target; zero/several eligible apps or a stale explicit selector must refuse that selection safely and identify action targeting rather than falsely reporting a popup launch prerequisite. Never guess first/last/label/current-focus/terminal-owner as invocation provenance.

## Refined binding boundary
Native basic mappings are immutable canonical Go-owned defaults. No app action executes on display/navigation/Escape. Route selection queries server configuration registration independently of GUI selection. A missing implicit default enters native mode without app lookup. Configured registered default or explicit --action retains existing original-Lua-VM action context binding, app eligibility/generation preflight,30second authenticated one-shot capability expiry and no uncertain fallback/retry. No new Lua VM/schema or configuration changes.

For native selection, retain chosen mapping/args/confirmation from trusted parent snapshot, resolve requested/unique live graphical app only when needed, capture full ID and connection generation then recheck them plus exact endpoint/session/block before dispatch. Validate only selected app action compatibility then; unsupported other actions must not prevent displaying a menu. No stale action execution or silent client retargeting. Syntax errors in endpoint/session/block/client options can still fail early.

## Native resource lifetime and cleanup
Distinguish exact server/session/block/resource liveness from GUI attachment. Native idle and dismissal use resource liveness only, preserving no idle timeout and FIFO parent lease; GUI visibility/connection ambiguity must not cancel it. Cleanup must close only its owned layer/block/private files before attempting native action target resolution/dispatch, and remain possible without an app. Preserve signal/origin-loss cleanup, bounded RPC/dispatch operations, exact layer membership checks and no cascading/window/session/global cleanup. Do not loosen configured decoded handoff validation through an empty native client field or untrusted lifetime flag. Focus remains captured-origin server focus, not an arbitrary app's current focus.

## Acceptance evidence and limits
Fresh independent source review and exact combined Go/Taskfile compilation required. Coordinator runs actual bare binary/make run in this Linux Rex session, root/submenu/navigation/Escape/exit0/exact owned cleanup, recording live app counts without requiring a singleton. Also run deliberate valid-but-nonexistent explicit client selector: popup must display and dismiss successfully; selecting one non-destructive app mapping with that nonexistent target must fail only after selection with zero action dispatch and owned resources already cleaned. This targeted refusal check is authorized by the user-requested moved guard; no actual destructive/native action dispatch, general tests/TDD/race suite, config/settings reload, new app connection/daemon, remote or Darwin trial. Do not disconnect user apps to fabricate a zero/multiple-app condition. Source-reviewed zero/multiple behavior and actual observed app counts must remain distinct. Server captures are not a remote GUI presentation certificate. No push.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE2OjE0OjMyLjk4MloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9hcHAtaW5kZXBlbmRlbnQvZGVzaWduLWFwcHJvdmFsLzEiLCJpbnRlbnQiOiI0YjQ1NTY1ZTllMmFjYjJhYWU0MGE2NTVjNTY5MTVmMTk4YzNhYTE1NTQzOTU4MDQxZTE3ZTgxOTk5YTIyMzFjIn0 -->
### 2026-11-02T16:14:32.982Z · approval · unspecified

{"approvedBy":"explicit user please fix after decoupled native launch explanation; delegated bounded implementation","attempt":1,"decision":"approved","decisionEvidence":"User explicitly requests fixing launch-time graphical app targeting after explanation that basic display needs only server/session/block. Native display/escape independent of app count or requested target liveness, native action binding deferred and still strict, configured original-VM binding unchanged. Targeted actual UI startup/dismissal and nonexistent-target non-destructive refusal checks authorized; no general tests, destructive dispatch, config changes or app disconnections.","decisionRegisterEntry":{"decision":"Native basic popup display independent of app target; validate native client only when needed for selected executable action. Preserve configured binding and exact action safety.","id":"REXWHICHKEY-20261102-111430","rationale":"Launch-time GUI selection incorrectly prevented a server-side terminal menu from opening; app ambiguity concerns action routing, not display.","state":"Accepted"},"eventId":"rex-whichkey/app-independent/design-approval/1","kind":"approval","recordRef":"designs/rex-whichkey-app-independent-popup.md","requirements":["REQ-023"],"schemaVersion":1,"specHash":"sha256:7215d04e0b71c6cc6a8a593372ef03cec57ae6aeb16907a81f1f9407b5781b4c"}
<!-- mtn:end-entry:v1 -->
