---
identity:
  kind: task
  taskID: rex-whichkey-15-02-cli-discovery
  uuid: 93f98f39-060a-40a2-ad3f-1a788c085d59
structure:
  order: 2
  parentRef: phases/rex-whichkey-15-code-owned-runtime.md
scope:
  requirements:
    - REQ-002
    - REQ-011
    - REQ-013
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
      - rex/exec.go
      - rex/endpoint.go
      - rex/rex.go
      - rex/hosts.go
      - rex/runtime.go
    sharedPaths: []
    resources: []
    interfaces:
      - RUNTIME-CLI-DISCOVERY
lifecycle:
  state: planned
  sourceEventID: null
  changedAt: '2026-11-05T00:40:14.130Z'
title: Centralize code-owned execution-host Rex CLI discovery
mtnCreateRetry:
  key: rex-whichkey-15-02-cli-discovery/create/1
  intent: 8bf1b6988caf4597df074a03fc5a9634835ee67d390da83deb5ee21c9151135b
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-05T00:40:15.121Z'
dateModified: '2026-11-05T00:40:26.485Z'
tags:
  - task
type: task
id: ea821f54-2565-4998-8e94-5bf3d3f84673
projects:
  - '[[phases/rex-whichkey-15-code-owned-runtime.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-15-01-bootstrap-contract.md]]'
    reltype: FINISHTOSTART
---
# Centralize code-owned execution-host Rex CLI discovery

## Deliverable and contract input

Implement RUNTIME-CLI-DISCOVERY from the reviewed docs/runtime-bootstrap.md produced by tasks/rex-whichkey-15-01-bootstrap-contract.md. Work only in the existing concrete rex package; do not introduce domain/ports/adapters, a daemon, direct socket transport or a general dependency-resolution framework.

## Acceptance criteria

1. Provide one bounded, code-owned Rex CLI selection/validation path with explicit provenance on the execution/server host. Ordinary supported operation must not require a user-maintained rex_bin_dir constructor argument or repaired GUI-server PATH. The exact supported sources and precedence come from the producer contract; do not guess home/.local/bin, Nix store prefixes, a viewing Mac's filesystem or undocumented paths.
2. Existing rex.New, Hosts and endpoint discovery use the shared live resolver rather than duplicating path policy. Expose the minimum reusable contract needed by the Go launcher; avoid speculative unused APIs or a second resolver for popup code.
3. Keep endpoint selection separate from executable identity: preserve explicit server/client intent, native CLI endpoint discovery and exact captured endpoint targeting. Do not reinterpret opaque Lua ctx.server or retry another server/executable after an uncertain operation.
4. Validate malformed/empty/unavailable/ambiguous or stale identity sources according to the agreed contract before mutations. Use argument arrays, bounded output/cancellation and fixed safe diagnostics; never expose credentials or capability-bearing output.
5. Keep genuine unsupported-context failures distinguishable from normal discovery. A fallback that silently ignores explicit intent or assumes a sole/current app is not acceptable.

## Ownership and dependencies

Prerequisite: tasks/rex-whichkey-15-01-bootstrap-contract.md.
Own rex/exec.go, rex/endpoint.go, rex/rex.go, rex/hosts.go and, only if required by that contract, rex/runtime.go. Other resource wrappers are inputs, not an expansion of this task. Do not edit root launch/controller code, Lua, packaging, project Taskfile/lint/hooks, external user configuration or managed .tasks records. Return a minimal interface handoff for the downstream Go bootstrap consumer.

## Review and acceptance scenarios

Require independent Go spec/quality source review, including actual non-test call sites for every added resolver/interface and no semantic dead code. Describe supported full/minimal PATH, trusted runtime metadata, explicit executable intent, missing/malformed sources, cancellation, CLI failures, remote execution-host separation and endpoint preservation. Keep existing regression tests intact; tests/TDD execution remain unperformed unless separately authorized. The current policy permits source/format review and compilation, not project test runners, runtime probes, live menu actions, user-config changes or deployment. Record the build actually performed and all unverified cases; do not alter lint/Taskfile standards to obtain a pass.

## Coverage

Phase criteria 2 and 6; complementary producer for criteria 1 and 3. Requirements REQ-002, REQ-011, REQ-013, REQ-017. This task does not claim CLI/keybinding bootstrap or packaged/Lua compatibility is delivered before downstream integration.
