---
identity:
  kind: task
  taskID: rex-whichkey-15-04-lua-bridge
  uuid: f170d8a5-b2a0-470c-a6e1-d6bcb0bccbe1
structure:
  order: 4
  parentRef: phases/rex-whichkey-15-code-owned-runtime.md
scope:
  requirements:
    - REQ-003
    - REQ-005
    - REQ-011
    - REQ-013
    - REQ-016
  documents:
    - ref: designs/rex-whichkey-runtime-ownership.md
      hash: 'sha256:253d40dc6204148a88924f7c860ad092f86f3b541a618330f112c6202b3aaa65'
    - ref: prds/rex-whichkey.md
      hash: 'sha256:97614db611a6f35a44d0db32ccf2bf4d4854d6ffea3dfd5375959e59259f01ff'
    - ref: prds/rex-whichkey-usability.md
      hash: 'sha256:966ef70c264360b4820e5faf510496166e91f0effeae479ab7400a3792de88e9'
implementation:
  language: lua
  gates:
    - ID: whitespace
      command: git diff --check
      required: true
  ownership:
    paths:
      - lua/rex_whichkey.lua
      - lua/rex_whichkey/dispatch.lua
    sharedPaths: []
    resources: []
    interfaces: []
lifecycle:
  state: planned
  sourceEventID: null
  changedAt: '2026-11-05T00:40:20.327Z'
title: Migrate original-VM Lua launch and completion to code-owned wiring
mtnCreateRetry:
  key: rex-whichkey-15-04-lua-bridge/create/1
  intent: b8c817dfda389daec8ed52b6eb0e287de267da243b18eb1dce884cc50adf4b02
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-05T00:40:21.315Z'
dateModified: '2026-11-05T00:40:28.541Z'
tags:
  - task
type: task
id: 7c440263-ea5c-493f-8d45-6e01d08c8019
projects:
  - '[[phases/rex-whichkey-15-code-owned-runtime.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-15-03-go-bootstrap.md]]'
    reltype: FINISHTOSTART
---
# Migrate original-VM Lua launch and completion to code-owned runtime wiring

## Deliverable and producer contract

Consume the integrated Go-owned runtime/compatibility contract from tasks/rex-whichkey-15-03-go-bootstrap.md and docs/runtime-bootstrap.md. Lua remains minimal configuration/registration/original-VM dispatch glue, not an alternative format owner, a path-configuration surface or a second callback VM.

## Acceptance criteria

1. Ordinary menu construction, registration and key/palette launch do not require execution.helper, execution.rex_bin_dir, socket/store/repository literals or Nix-generated substitutions in user init.lua. The supported code-owned bootstrap works in the restricted-PATH context specified by the producer contract; do not guess a filesystem location or fabricate a Rex API capability.
2. Use the authoritative producer's compatibility/identity data before helper launch and completion. Handle supported component updates without user-maintained binary/Lua pairs. Unsupported or malformed data fails closed with fixed diagnostics rather than a flashing popup, dropped selection, native-menu fallback or a retry.
3. Retain closures, condition guards, original args, snapshot and capability registry in the originating persistent server VM. Do not replace real activation context with a nominated app ID or deserialize executable Lua into Go. Each actual caller is frozen within its own activation; the GUI target remains separately retained and validated.
4. Preserve exact server/session/block and rename resource identity, app attachment/generation, configuration/expiry checks, one-shot claim before yielding effects, bounded pending state, cancellation and privacy. Existing original-VM callbacks and custom mappings are not bypassed or rewritten.
5. Default tr/wr selections use WhichKey's input page and validated text to rename the retained tab/session once. A stale/incompatible text contract must not discard the name, open Rex's dialog, choose another resource or report an unperformed rename. Custom callback/argument mappings remain their original actions.
6. Retire or safely handle obsolete execution-path constructor inputs according to the producer migration contract without automatically editing or reloading user configuration. Do not move their values into generated Lua or a hidden Nix runtime resolver.

## Ownership and dependencies

Prerequisite: tasks/rex-whichkey-15-03-go-bootstrap.md; its Go schema/compatibility producer must be available and stable first.
Own lua/rex_whichkey.lua and lua/rex_whichkey/dispatch.lua. Treat config/mapping/evaluate/defaults and all Go files as consumers/inputs to inspect, not unrestricted write scope. If the contract genuinely requires changing another Lua module, request a coordinator ownership adjustment with exact rationale before edits. Never hand-edit managed records or external init.lua/Nix configuration.

## Review and acceptance scenarios

Language is Lua, not an implicitly Go task. Require independent combined source review covering the real Rex activation boundary and Go producer contract. Review supported public/key/palette routes, restricted PATH, wrong or missing identities, component mismatch, rename text, retained callbacks, reentry/replay, setup replacement, cancellation, expiry and uncertain RPCs. These are acceptance/TDD scenarios, not permission to run tests, Lua harnesses, live Rex actions, probes, temporary configuration reloads or remote trials. Preserve existing tests; use only verification actually authorized at execution time and record unperformed checks. Source review must not claim real app/runtime effects.

## Coverage

Phase criteria 1–4, 6 and 7 on the original-VM side; requirements REQ-003, REQ-005, REQ-011, REQ-013 and REQ-016. No phase completion until the combined Go/Lua/package revision has independent source review and required authorized integration checks.
