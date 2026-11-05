---
identity:
  kind: task
  taskID: rex-whichkey-15-01-bootstrap-contract
  uuid: 375ec893-c440-4e42-89b2-2d900f14cc18
structure:
  order: 1
  parentRef: phases/rex-whichkey-15-code-owned-runtime.md
scope:
  requirements:
    - REQ-002
    - REQ-003
    - REQ-005
    - REQ-011
    - REQ-013
    - REQ-014
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
  language: docs
  gates:
    - ID: whitespace
      command: git diff --check
      required: true
  ownership:
    paths:
      - docs/runtime-bootstrap.md
    sharedPaths: []
    resources: []
    interfaces:
      - RUNTIME-CLI-DISCOVERY
      - RUNTIME-BOOTSTRAP-COMPATIBILITY
lifecycle:
  state: planned
  sourceEventID: null
  changedAt: '2026-11-05T00:40:11.037Z'
title: Specify the supported code-owned bootstrap contract
mtnCreateRetry:
  key: rex-whichkey-15-01-bootstrap-contract/create/1
  intent: 3af8e2436bf21b7e80ee254a479f7dc1b9b21fd209b07d6738d32174a5e343f5
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-05T00:40:12.030Z'
dateModified: '2026-11-05T00:40:13.073Z'
tags:
  - task
type: task
id: 9da80940-08a4-4935-aab3-451ddc69be33
projects:
  - '[[phases/rex-whichkey-15-code-owned-runtime.md]]'
---
# Specify the supported code-owned bootstrap contract

## Deliverable and inputs

Produce docs/runtime-bootstrap.md as the concrete implementation contract for Phase 15's accepted responsibility boundary. Read the parent phase and designs/rex-whichkey-runtime-ownership.md through mtn, plus the current public launcher, Rex CLI resolver, configured Lua launcher/dispatcher, controller and packaging source. This is bounded source/API-documentation work, not a new research experiment or permission to run the app.

## Acceptance criteria

1. Trace ordinary public CLI launch, key/palette launch, helper preparation, controller/PTY self-spawn and configured completion. Name every executable/asset identity source and who owns it; distinguish real CLI callers from retained GUI targets and execution/server host from viewing host.
2. Resolve a documented, supported bootstrap under an app-launched restricted PATH without user helper/rex_bin_dir literals, Nix runtime substitutions, hardcoded user directories, alternate-server probing or a second callback VM. Define producer/consumer interfaces RUNTIME-CLI-DISCOVERY and RUNTIME-BOOTSTRAP-COMPATIBILITY, including precedence, bounds, failure stages and compatibility/capability semantics. Do not merely move the existing literal into Go.
3. Define how a local rebuilt executable and packaged executable obtain coherent helper/controller/child/asset selection without user-maintained version pairs. Preserve registered customization instead of silently bypassing it when incompatible; record any unavoidable supported-platform limitation precisely.
4. Define safe compatibility behavior before popup mutation and before rename dispatch; include default tab/session rename text and original-VM callback preservation. Keep endpoint/origin/generation/capability/expiry/confirmation/consume-before-dispatch/owned-cleanup/no-retry invariants explicit.
5. Record the source revision and reconcile which existing working-tree repairs are inputs, not delivered Phase 15 work. Do not overwrite/stash/commit other writers' changes or silently migrate legacy project metadata. Return any necessary managed-design update as a coordinator mutation proposal, not a worker edit to .tasks.
6. Supply a boundary table for the downstream Go resolver, Go bootstrap, Lua consumer and packaging/docs tasks. State which old constructor fields/examples become obsolete and how upgrade compatibility works without taking over user configuration.

## Scope, dependencies and review

Own docs/runtime-bootstrap.md only. Read other source; do not implement production changes. No task prerequisites. This task is the real contract producer for the remaining tasks, not an approval of an already-chosen bootstrap algorithm. Use supported documentation/source, preserve the flat Go application and one rex package, and require independent combined documentation/security-boundary source review.

If no documented mechanism can meet the accepted restricted-PATH/original-VM/code-owned boundary, report the exact missing capability and safe within-scope alternatives before marking this task complete. Do not invent API capabilities or restore init.lua/Nix path workarounds. A material requirement/security/public-contract departure still needs user approval; routine within-scope choices do not.

## Acceptance scenarios and verification policy

Describe traceable cases for public versus key/palette launch, restricted/missing/ambiguous PATH, local versus installed helpers, stale/unsupported protocol, malformed identity, rename text, custom callbacks, cancellation, cleanup and no retry. These are falsifiable source-review/TDD scenarios, not permission to execute them. Existing no-project-tests/no-experiments restrictions remain. No test runners, probes, live Rex actions, config reloads, package activation or deployment are authorized by task creation. Preserve existing tests; record unperformed checks honestly. Independent review is evidence, not a command gate. Only metadata/source/whitespace checks are required for this documentation task.

## Coverage

Parent criteria 1–8: this task fixes the interfaces and responsibility boundaries; downstream tasks implement them. Requirements REQ-002, REQ-003, REQ-005, REQ-011, REQ-013, REQ-014, REQ-016, REQ-017. Completion is an integrated reviewed contract, not a claimed working runtime.
