---
workflowKind: task
workflowLanguage: go
workflowOrder: 1
workflowRequirements:
  - REQ-017
workflowDocuments:
  - ref: prds/rex-whichkey-usability.md
    hash: 'sha256:c2df9b9404ce7f2d649eed40d5c219f7098e9389c6407159b6ae35eddb7746c5'
  - ref: designs/rex-whichkey-usability.md
    hash: 'sha256:c8553dd4ac48fcd34a43d7cb3cd842bb18c65310d0fc6133cf1b6e3f196f1b8a'
workflowOwnership:
  paths:
    - rex/endpoint.go
    - popup.go
    - cmd/main.go
    - lua/rex_whichkey.lua
    - docs/launch-contract.md
  sharedPaths: []
  resources: []
  interfaces:
    - canonical-popup-endpoint
    - private-prepare-server-result
projects:
  - '[[phases/rex-whichkey-07-usability.md]]'
status: open
priority: normal
title: Implement platform endpoint policy and real Lua consumption
blockedBy: []
mtnCreateRetry:
  key: rex-whichkey/usability/task07-01/v1
  intent: 9a7e04b23539efa38a438e4ffd8297cb85e0f0f5298ad357187d3797d43a50d6
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-01T15:15:17.098Z'
dateModified: '2026-11-01T15:15:17.098Z'
tags:
  - task
type: task
id: 30b5d1fe-6dbc-49ce-aed8-1b24a97546b8
---
# Implement Go-owned endpoint policy and its real Lua consumer

## Authority and inputs

REQ-017; Phase07 criteria 2/3/5. Read the current approved usability requirements and design (native refs in workflowDocuments), planner-source-contracts.md for retained source evidence, and current popup.go, rex/rex.go, cmd/main.go and lua/rex_whichkey.lua. Coordinator has resolved planner B1 by changing the implementation approach: delegate native OS path discovery to Rex rather than reproducing unavailable path algorithms. This is approved under explicit user-delegated technical authority, not an unresolved source/approval blocker. The design supersedes contradictory exact-path-table draft requirements.

## Literal ownership

rex/endpoint.go, popup.go, cmd/main.go, lua/rex_whichkey.lua, docs/launch-contract.md. Go plus substantive Lua/Markdown; mixed review mandatory. Flat package structure; no dependency, mapping-schema, Taskfile/lint/hook or dispatcher changes.

## Acceptance

1. Implement one small canonical Go resolver/validator in rex/endpoint.go. Resolve a present explicit override first, then nonempty execution-host REX_SERVER, otherwise use bounded argv-based `rex --autostart=false whoami --json` to delegate native Linux/macOS connection/socket resolution to the installed matching Rex CLI. No hard-coded/guessed OS filename table, socket scanning, server autostart, directory creation or graphical-client filesystem lookup. Keep Rex's native runtime/home environment under Rex ownership; prevent inherited target selectors from silently redirecting the platform-default branch. Validate/capture one concrete returned endpoint; malformed/missing JSON/endpoint or CLI failure produces an actionable sanitized error, no retries/fallback. Unsupported automatic-default platforms fail clearly. Distinguish absent override from explicitly empty invalid override.
2. Reuse strict existing URI/address validation for both resolution and controller handoffs without weakening controls/whitespace/port/absolute-path checks. Reject mutable labels; preserve accepted unix URI/host:port/bracketed IPv6 forms and canonicalize absolute Unix paths where supported by published Rex. Existing generic rex.New behavior can remain unchanged: all new popup callers receive a concrete endpoint explicitly. A missing controller server never defaults.
3. Extend actual prepare plumbing: optional explicit --server input, resolve on execution/server host, include canonical server in captured capability/executable preparation response. Preserve PreparePopup(input,output) callers through a backward-compatible wrapper/additive API and existing prepare no-argument parsing. Capabilities remain pipe-only, never argv/logs/errors. Native-discovery failures must not disclose private handoff data.
4. Make popup.new(defaults) valid while preserving an optional supplied immutable endpoint override. Remove Lua's duplicate policy guesses. The real preparation invocation passes an override only when supplied, validates response shape and uses Go's returned server in the existing detached controller handoff. Keep original ctx/opaque ctx.server unchanged; no caller identity copied from routing data.
5. Preserve issue-before-evaluation, bind/cancellation, config generation, client connection generation, placement checks, one-shot/expiry/snapshot/confirmation, original-VM callbacks, bounded owned cleanup and no uncertain action retries. Existing constructor calls with explicit endpoint keep working. No new daemon/VM/host feature or init.lua edit/reload.
6. docs/launch-contract.md describes actual API/response and override/native-discovery precedence, read-only identity query, prerequisites, canonical validation, failures and unexecuted runtime behavior. State platform algorithms belong to Rex; do not claim exact undocumented formulas or runtime/Darwin verification.

## Verification and delivery

No project tests/TDD/experiments/product execution/GUI/live API/remote trials/config reloads/dependency installation. Builder formats Go/Lua and performs whitespace/static source checks only; no heavy slot. Preserve tests/gates. Fresh review must cover Go and Lua real resolver -> prepare -> launcher -> explicit controller flow, native-discovery selector handling, override precedence and original dispatch safeguards. Coordinator compiles reviewed combined candidate and verifies exact source/Git ancestry before task completion. No stub/token caller, source/managed write outside scope, push or integration by worker. Full protocol build goes to dispatch payloadFile AND MTN_REX_RESULT_FILE; final visible handoff is a concise filename pointer.
