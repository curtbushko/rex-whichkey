---
workflowKind: design
title: Self-contained native default popup startup
mtnCreateRetry:
  key: rex-whichkey-native-startup-design-v1
  intent: ae364a99fe9013e50ed55a259aa6ca807c49dd133bad5cd826b7082365577ac5
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T01:18:32.804Z'
dateModified: '2026-11-02T01:18:38.099Z'
tags:
  - task
type: task
id: 19736962-52fc-4c4f-ab1d-d789085d83c5
completedDate: '2026-11-01'
---
# Self-contained default popup startup

## Context, user authority and observed failure

REQ-016 requires just running rex-whichkey to show the popup; the user explicitly rejects further manual diagnosis/configuration of an already-working Rex session and authorizes launching/running it here. A real invocation in an owned Rex terminal exited1. Read-only live server/client action registries have no whichkey/whichkey_complete, while exactly one connected app is attached. Thus the current default caller requires custom action registration and does not deliver the requested basic experience. The user's session is healthy; do not modify/reload its configuration.

The actual terminal renderer was separately launched with explicit diagnostic data, displayed root and Tabs submenu and dismissed via Escape with exit0. This is renderer/navigation evidence only, not proof that the normal floating popup/default actions work. Saved live evidence is .mtn-tmp/live-launch/. Native action registry, session.new_layer schema and block creator metadata are saved there for implementation, not guessed contracts.

## Chosen bounded refinement

Keep the existing configured original-VM path for registered/custom Lua mappings and callbacks. Remove custom registration as a prerequisite for the basic built-in popup: when the default whichkey launch action is absent, the normal Go entrypoint starts an invocation-scoped Go-owned popup using canonical Mapping/Config and native Rex actions. A provided --action remains an explicit configured-action request and fails with actionable registration information if absent. Prefer the configured action when present; never fall back/retry native execution after a configured activation was attempted or became uncertain.

The Go caller captures exact terminal session/block, native endpoint and attached live app/generation as already delivered. It can remain the short-lived controller while the floating PTY child runs; no server configuration action or new background daemon is needed. Reuse existing owned-layer, private transport, renderer, selection and cleanup mechanisms. Resolve defaults as actual Go-owned Mapping data (no alternative schema), matching delivered built-in keys and native actions supported by the matching client contract. Host provider remains informational. No fake/demo menu or private command requirement.

Keep an immutable parent snapshot of eligible actions/args. The child returns selection identity/confirmation only; map that identity to the retained trusted mapping, never accept replacement executable/action/args from child data. Enforce one selection/attempt, explicit origin/generation checks before mutation/dispatch, destructive confirmation and bounded lifetime. Restore terminal/close exact owned popup and remove private files BEFORE invoking a selected native action, so closing the origin cannot kill its own cleanup. Invoke native actions once with argv and exact endpoint/client/session/block; no focused/current-client or alias fallback. Cancellation performs no action. No new Lua VM is created; configured Lua callbacks continue exclusively through the unchanged original-runtime authenticated bridge.

Avoid duplicated service frameworks or full-controller forks: share/refactor existing lifecycle helpers where practical without weakening the Lua path. Parent/child config/result files remain invocation-owned 0700/0600, bounded, no secret logs or broad cleanup. No user init.lua/PATH/system setting edits, config reload/restart, package activation or global cleanup.

## Diagnostics and compatibility

Distinguish native basic startup from optional configured Lua integration in public help/docs. Surface safe stage-specific failures rather than discarding every Rex error into the current generic prerequisite list; never expose arbitrary callback output/capabilities/credentials. Existing menu/prepare/controller/popup modes, custom --action, constructor socket overrides, native endpoint defaults and registered Lua launch flow remain compatible. Do not silently choose another execution route after uncertain activation.

## Verification

The user's latest instruction authorizes a bounded real launch/navigation/dismissal of rex-whichkey in this existing Rex session. Coordinator owns these checks and records exact resources/output/exit/cleanup; do not modify/reload user configuration or execute destructive menu actions. Unit/race/TDD/general test runners, unrelated experiments, remote/Darwin/host trials remain prohibited. Builders format/static-check only; a fresh independent reviewer covers both execution routes, snapshot/origin/confirmation/privacy/lifecycle and documentation. Coordinator compiles combined candidate, performs the targeted actual popup smoke check, fixes concrete failures with new review, then commits. No source-only or renderer-only result may be called a verified normal popup launch.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDAxOjE4OjM3LjE5OFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS91c2FiaWxpdHkvbmF0aXZlLWRlZmF1bHQtZGVzaWduLzEiLCJpbnRlbnQiOiJiMzgzYjcyYzM0MzdlMGE3YWVhZDRiZjdlMTA0ZmE1NTQyNzM4MGVkNDM4MzNjYWE1ZmU5ZmU0MDIyZDMzZjlmIn0 -->
### 2026-11-02T01:18:37.198Z · approval · unspecified

{"approvedBy":"coordinator under user-authorized direct-startup implementation and delegated technical choices","attempt":1,"decision":"approved","decisionEvidence":"User asks just launching the binary to show a popup, rejects more manual setup/diagnosis and explicitly authorizes actual launch here. Live reproduction proves the normal route depends on absent custom actions despite a healthy attached app. Basic Go/native startup removes that prerequisite without editing working Rex configuration; original-VM configured callbacks remain compatible. Targeted live menu launch/navigation/dismissal is now authorized; general project tests/destructive/config/remote probes remain prohibited.","decisionRegisterEntry":{"decision":"Use self-contained Go/native default popup when default Lua launch action is unregistered; preserve explicit/registered original-VM action route and no uncertain fallback.","id":"REXWHICHKEY-20261101-201834","previousState":"Open","refines":"REXWHICHKEY-20261101-101502","state":"Accepted"},"eventId":"rex-whichkey/usability/native-default-design/1","kind":"approval","recordRef":"designs/rex-whichkey-native-startup.md","schemaVersion":1,"specHash":"sha256:3c49a792bd2302bae6253174ab5d992dbf45cd9b557f53dd32e539bae8ce6f7c"}
<!-- mtn:end-entry:v1 -->
