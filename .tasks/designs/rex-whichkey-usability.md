---
workflowKind: design
title: Usability startup and native endpoint discovery
mtnCreateRetry:
  key: rex-whichkey-usability-native-discovery-design-v1
  intent: 5d87654a3f3fe05fe044f230b8063c61c7a7e831fd0ebf02d2e3e75bff88f580
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-01T15:15:01.061Z'
dateModified: '2026-11-01T15:15:06.393Z'
tags:
  - task
type: task
id: a76623c9-87b5-4248-87aa-d5dc3d3bdad8
completedDate: '2026-11-01'
---
# Usability startup and native endpoint discovery

## Status and authority

Accepted implementation direction under the user's explicit Phase07 request and previously delegated bounded technical choices. Requirements are REQ-016/REQ-017 in prds/rex-whichkey-usability.md. The earlier mandatory-explicit-endpoint decision is refined only as to common-case setup; actual opaque action context, explicit overrides and origin safety remain. This document is intent, not source delivery or runtime verification.

## Problem and chosen design

The binary currently exposes only internal modes and resolved-JSON terminal mode. The new public root route must launch the real delivered popup. Exact native macOS/Linux path algorithms were not available in inspected source/docs, so duplicating a speculative directory/filename table was rejected. The installed Rex CLI already owns its platform-specific default connection resolution. Reuse that authority instead of building a second resolver for its paths.

### canonical-popup-endpoint

One Go policy in rex/endpoint.go resolves a present explicit endpoint override first, then a nonempty execution-host REX_SERVER, otherwise invokes the installed Rex CLI with bounded context, argument-array execution, autostart disabled and `whoami --json`. The static CLI help documents identity/server JSON. Existing saved orchestration identity includes a concrete `endpoint`; that is evidence of available transport representation, not a product runtime acceptance test. Reject unsupported/malformed response, missing concrete endpoint and failure clearly. Require Linux/macOS for automatic defaults; unsupported systems need a clearly supported explicit path or error according to retained build scope.

Do not hard-code undocumented paths, probe several sockets, start servers, use host labels, inspect graphical-client files, stringify opaque ctx.server, or use a default/focused client. On the no-override/no-REX_SERVER branch ensure inherited selector variables cannot redirect native default discovery; platform runtime/home environment remains owned by Rex. Normalize accepted absolute Unix paths to canonical unix URIs if supported; preserve strict existing URI/address validation and fail on present empty/malformed overrides.

The resolver returns one concrete endpoint. All later application/Rex clients, controller handoffs, focus and completion calls target it explicitly. A connection failure does not authorize selecting another server. Native discovery contacts only the normal default server as an ordinary read-only part of product startup, not a new service or discovery scan. No agent runtime probe is authorized by this design.

### private-prepare-server-result

`prepare` continues decoding actual resolved configuration and returning captured capability/executable data; it adds canonical `server` and accepts an optional explicit server override. Preserve old PreparePopup(input,output) source callers through a wrapper/additive API. Lua `popup.new(defaults)` becomes valid; a supplied endpoint stays an immutable explicit override. The real preparation call consumes Go's result and forwards its concrete server to the existing detached controller. Opaque actual ctx is never replaced or populated from this endpoint. Existing empty controller handoffs still fail closed; controller itself does not ambient-default.

### public-original-vm-launch-origin

No-argument/default public CLI calls a flat LaunchPopup implementation. Capture exact full REX_SESSION and REX_BLOCK, resolve endpoint once and verify placement. Use explicit full --client/REX_CLIENT if supplied, otherwise inspect only that captured session and require exactly one eligible attached full client ID plus connection generation. Zero/multiple/lost/ambiguous matches fail with instructions; never delegate missing client resolution to Rex's current/focused/default app selection. All public context flags must be explicit and validated if offered.

Activate the installed public configuration action (default whichkey; public --action may select a user-registered alternative) exactly once through bounded argv execution with explicit endpoint/session/block/client and origin JSON via stdin. No Lua script or second VM. Public help describes normal startup and prerequisites; legacy menu/prepare/controller/popup modes remain compatible.

Lua menu.register_launch(name) registers a genuine run(ctx,args) action at user-owned configuration load time. Its optional Go-owned origin-envelope string is decoded and compared to the actual ctx, canonical endpoint and live exact connection. Then it calls the existing menu.launch path with that original ctx unchanged, rechecking expected connection/realm across yielding preparation calls. Palette/key calls with genuine context can omit the envelope. Existing menu.launch(ctx), completion registration, snapshot evaluation, one-shot capability, generation/expiry/confirmation and owned-only finite cleanup stay intact. Mapping.go remains the single mapping standard; a private launch DTO is Go-owned, not a new Lua schema.

## Prerequisites and user experience

Normal one-time installation/registration in the Rex server's existing user-owned configuration is still required. The helper and matching Rex CLI/modules must be installed on the execution host. Client Remote Control policy applies. The user runs `rex-whichkey` inside a Rex terminal; ambiguous attached clients require an explicit full selector. Missing prerequisites produce short actionable errors rather than usage-only success, guessed context or silent fallback. Do not auto-edit/reload init.lua or claim outside-Rex/unknown-context startup works.

## Alternatives and lineage

- Exact copied platform path algorithms: superseded because readable authoritative algorithms were unavailable and duplicate ownership is unnecessary.
- Mandatory constructor socket literal: refined by user-requested platform defaults; explicit override and separate opaque identity still apply.
- Bare CLI demo/default JSON or a second Lua VM: rejected; not the real action/callback popup.
- Implicit current/focused client or server retry: rejected; wrong-origin risk.
- New daemon/transport framework or dependency adoption: rejected as unnecessary scope.

## Delivery and verification

Three serial tasks: endpoint producer plus real Go/Lua consumer; no-argument original-VM startup plus caller; docs/help/final assembled source review. Each build receives a fresh mixed/source reviewer and coordinator combined-target compilation with exact Git proof. Current policy forbids project tests/TDD, experiments, product execution, GUI/live/remote trials, config reloads and destructive/global cleanup. No implementation/runtime pass is implied here. Preserve gates/tests and use hook-disabled implementation/metadata commits; no push.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDE1OjE1OjA1LjQ4MloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS91c2FiaWxpdHkvbmF0aXZlLWRpc2NvdmVyeS1kZXNpZ24tYXBwcm92YWwvMSIsImludGVudCI6ImI2M2ViZTQxMTcxYmQ3ZjRlNTk4NGZiOTBmMjgzMmRmNTEzOWIxYjU4N2NlYzU5MGFkMDM5ODhiOGNlNTIxNDAifQ -->
### 2026-11-01T15:15:05.482Z · approval · unspecified

{"approvedBy":"coordinator under explicit user delegation","attempt":1,"decision":"approved","decisionEvidence":"User authorized implementation and delegated bounded design choices. Installed Rex whoami help documents server identity JSON; retained orchestration identity has a concrete endpoint. Reuse native platform discovery rather than guess inaccessible path algorithms, then freeze explicit endpoint and preserve original-runtime validation. New source and runtime are not verified by approval.","decisionRegisterEntry":{"decision":"Delegate platform endpoint discovery to installed Rex CLI; add real zero-argument original-VM popup activation with explicit captured origin.","id":"REXWHICHKEY-20261101-101502","previousState":"Open","refines":"REXWHICHKEY-20261101-015342","state":"Accepted","supersedes":"REXWHICHKEY-20261101-100106"},"eventId":"rex-whichkey/usability/native-discovery-design-approval/1","kind":"approval","recordRef":"designs/rex-whichkey-usability.md","schemaVersion":1,"specHash":"sha256:c8553dd4ac48fcd34a43d7cb3cd842bb18c65310d0fc6133cf1b6e3f196f1b8a"}
<!-- mtn:end-entry:v1 -->
