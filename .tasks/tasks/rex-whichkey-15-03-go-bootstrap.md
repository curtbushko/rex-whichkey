---
identity:
  kind: task
  taskID: rex-whichkey-15-03-go-bootstrap
  uuid: 4d5d24f5-6838-4340-b783-8ac646ab8791
structure:
  order: 3
  parentRef: phases/rex-whichkey-15-code-owned-runtime.md
scope:
  requirements:
    - REQ-002
    - REQ-011
    - REQ-013
    - REQ-016
    - REQ-017
  documents:
    - ref: designs/rex-whichkey-runtime-ownership.md
      hash: 'sha256:253d40dc6204148a88924f7c860ad092f86f3b541a618330f112c6202b3aaa65'
    - ref: prds/rex-whichkey.md
      hash: 'sha256:97614db611a6f35a44d0db32ccf2bf4d4854d6ffea3dfd5375959e59259f01ff'
    - ref: prds/rex-whichkey-usability.md
      hash: 'sha256:966ef70c264360b4820e5faf510496166e91f0effeae479ab7400a3792de88e9'
implementation:
  language: go
  gates:
    - ID: build
      command: task build
      required: true
  ownership:
    paths:
      - cmd/main.go
      - launch.go
      - popup.go
      - native_text_entry.go
      - mapping.go
      - runtime_bootstrap.go
    sharedPaths:
      - docs/mapping-format.md
    resources: []
    interfaces:
      - RUNTIME-BOOTSTRAP-COMPATIBILITY
lifecycle:
  state: planned
  sourceEventID: null
  changedAt: '2026-11-05T00:40:17.230Z'
title: Implement coherent Go-owned startup and compatibility
mtnCreateRetry:
  key: rex-whichkey-15-03-go-bootstrap/create/1
  intent: d114128a11bf062d0b51ad73e7b6962d3c27e72298e5e771cbbe32fd62aba531
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-05T00:40:18.214Z'
dateModified: '2026-11-05T00:40:27.511Z'
tags:
  - task
type: task
id: b3aeecac-e9a4-4855-83db-6345b6789c02
projects:
  - '[[phases/rex-whichkey-15-code-owned-runtime.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-15-02-cli-discovery.md]]'
    reltype: FINISHTOSTART
---
# Implement coherent Go-owned startup and compatibility

## Deliverable and inputs

Implement the Go side of RUNTIME-BOOTSTRAP-COMPATIBILITY from docs/runtime-bootstrap.md using the integrated RUNTIME-CLI-DISCOVERY producer. The application owns its runtime identity and coherent components; Nix and user Lua configuration do not select helper binaries or pair component versions.

## Acceptance criteria

1. Route the public entrypoint and internal preparation/controller/PTY paths through the code-owned runtime contract. Ordinary invocation requires no helper/rex_bin_dir constructor data, repository/store literals or user version coordination. Use the actual running executable for self-spawn where the agreed contract calls for it; do not defer a local rebuild to an arbitrary older installed helper because a Lua action exists.
2. Deliver the authoritative Go-owned compatibility/capability data and bounded validation required by the Lua consumer. Existing original-VM custom mappings/callbacks remain selected according to the accepted contract; incompatibility or uncertain activation cannot silently switch to the native menu or another app/server.
3. Establish component compatibility before opening a transient popup or accepting a selection that an older dispatcher will drop. Reject malformed, unsupported, ambiguous and stale identities with fixed stage-specific diagnostics, retaining no arbitrary private child output. Do not expose capabilities in argv, filenames, logs or errors.
4. Propagate coherent executable/asset identities through preparation, detached controller, terminal child and fresh completion. Preserve exact endpoint/session/block origin, separate caller/GUI identities, configuration/connection generation, original expiry, private owned transport, consume-once and cleanup-before-dispatch ordering.
5. Preserve the inline tab/session rename path and its typed text/resource binding. Missing or incompatible rename support cannot fall through to Rex's dialog or report successful completion without dispatch. Do not serialize callbacks or reconstruct their registry in a different Lua VM.
6. Keep ordinary explicit server/client intent and public help behavior compatible. Remove required duplicate root CLI-path policy in favor of the resolver producer; introduce no speculative or disconnected exported interfaces.

## Ownership and dependencies

Prerequisite: tasks/rex-whichkey-15-02-cli-discovery.md. The earlier contract task is transitively required.
Own cmd/main.go, launch.go, popup.go, native_text_entry.go, mapping.go and runtime_bootstrap.go if the agreed contract needs that new file. Reserve docs/mapping-format.md as an exclusive shared path for Go-owned contract documentation. Do not change Lua consumers, rex producer internals, packaging, rendering/terminal behavior, Taskfile/lint/hooks or user init.lua. Any necessary additional file ownership must be replanned rather than assumed. Publish the minimal stable consumer handoff before the dependent Lua task starts.

## Review and acceptance scenarios

Require independent Go spec/quality and mixed-boundary source review. Describe direct/local versus packaged startup, stale installed registered components, unsupported capability/rename fields, cancellation at every handoff, wrong origin, secret-safe errors, replay/expiry and failed/uncertain dispatch. These are acceptance/TDD scenarios only: the phase's no-tests/no-experiments policy remains. Preserve existing tests and do not weaken their assertions or quality configuration. Only authorized formatting/static/compilation checks may run during execution; no test runners, live Rex/remote trials, config reloads or deployment are implied. Record unverified behavior explicitly.

## Coverage

Phase criteria 1–4, 6 and 7 on the Go side; requirements REQ-002, REQ-011, REQ-013, REQ-016, REQ-017. Go-only integration does not complete the phase or certify the Lua consumer before its own reviewed integration.
