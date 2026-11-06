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
      - rex/endpoint.go
      - rex/exec.go
      - rex/hosts.go
      - rex/rex.go
      - rex/runtime.go
      - rex/session.go
    sharedPaths: []
    resources: []
    interfaces:
      - RUNTIME-CLI-DISCOVERY
lifecycle:
  state: complete
  sourceEventID: rex-whichkey-p15-t02-reviewer-a04/release-runtime-launched/1
  changedAt: '2026-11-05T17:36:23.336Z'
title: Centralize code-owned execution-host Rex CLI discovery
mtnCreateRetry:
  key: rex-whichkey-15-02-cli-discovery/create/1
  intent: 8bf1b6988caf4597df074a03fc5a9634835ee67d390da83deb5ee21c9151135b
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-05T00:40:15.121Z'
dateModified: '2026-11-05T17:36:36.576Z'
tags:
  - task
type: task
id: ea821f54-2565-4998-8e94-5bf3d3f84673
projects:
  - '[[phases/rex-whichkey-15-code-owned-runtime.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-15-01-bootstrap-contract.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-05'
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
Own rex/exec.go, rex/endpoint.go, rex/rex.go, rex/hosts.go and, only if required by that contract, rex/runtime.go. In rex/session.go, own only (*Client).resourceCommand’s existing run call: pass the already-frozen selection to that execution boundary. No other change to that file is authorized. Other resource wrappers remain inputs, not an expansion of this task. Do not edit root launch/controller code, Lua, packaging, project Taskfile/lint/hooks, external user configuration or managed .tasks records. Return a minimal interface handoff for the downstream Go bootstrap consumer.

## Review and acceptance scenarios

Require independent Go spec/quality source review, including actual non-test call sites for every added resolver/interface and no semantic dead code. Describe supported full/minimal PATH, trusted runtime metadata, explicit executable intent, missing/malformed sources, cancellation, CLI failures, remote execution-host separation and endpoint preservation. Keep existing regression tests intact; tests/TDD execution remain unperformed unless separately authorized. The current policy permits source/format review and compilation, not project test runners, runtime probes, live menu actions, user-config changes or deployment. Record the build actually performed and all unverified cases; do not alter lint/Taskfile standards to obtain a pass.

## Coverage

Phase criteria 2 and 6; complementary producer for criteria 1 and 3. Requirements REQ-002, REQ-011, REQ-013, REQ-017. This task does not claim CLI/keybinding bootstrap or packaged/Lua compatibility is delivered before downstream integration.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA1VDE3OjM2OjM2LjU2M1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxNS10MDIvc3VtbWFyeS80IiwiaW50ZW50IjoiYjk1ZjgyYjczZmQ4ZGZmMjNhN2FmZTA2OTc2NTQ2NGRiOGYwMjYyZGJjNWVkZTMxNWNmZGFhMzcyNmY5NjlkYiJ9 -->
### 2026-11-05T17:36:36.563Z · summary · unspecified

{"acceptedScope":"Complementary core producer: one execution-host Rex CLI selection and immutable executable/API evidence, shared client/Hosts/native endpoint discovery, bounded fixed-shape compatibility and safe diagnostics, plus exact resource-runner selection propagation. Root bootstrap/Lua/assets remain downstream work.","approvalSummary":["User explicitly authorized Phase15 and bounded ownership refinements; current plan approved/4.","Fresh same-family combined Go review rex-whichkey-p15-t02/review/4 approved 2b1124a373d813bd96945c6a91d8c5186ff0420a with zero blocking findings; not cross-family review."],"artifactRefs":[".mtn-tmp/p15-reconcile/task02-a04-result.json",".mtn-tmp/p15-reconcile/task02-a04-review-result.json",".mtn-tmp/p15-reconcile/task02-a04-integration-build.log"],"attempt":4,"eventID":"rex-whichkey-p15-t02/summary/4","files":["rex/endpoint.go","rex/exec.go","rex/hosts.go","rex/rex.go","rex/runtime.go","rex/session.go"],"integrationEventID":"rex-whichkey-p15-t02/integration/4","kind":"summary","limitations":["No project tests/TDD/race/Lua/runtime/API/menu/rename probes or experiments; source/static/host compilation is not runtime verification.","Full-project lint did not pass; independent candidate/baseline diagnostics remain explicit. Existing settings/checks unchanged; no blocking new owned findings in current review.","Only bounded fixed application schema semantics are supported; unsupported validation constraints fail closed, not a general JSON Schema compatibility promise.","Root duplicate discovery, coherent Go bootstrap, original-VM Lua consumer and normal dependency assets are not delivered by this producer task.","Cleanup facade has a jq runtime-guard compile defect; all workers independently confirmed stopped and owned worktrees retained safely pending cleanup planning.","Unreviewed source checkpoint 30a4439 is an input, not this task delivery."],"recordRef":"tasks/rex-whichkey-15-02-cli-discovery.md","requirements":["REQ-002","REQ-011","REQ-013","REQ-017"],"reviewEventIDs":["rex-whichkey-p15-t02/review/4"],"schemaVersion":1,"specHash":"sha256:0a892bfc8c548dca6a880379e746c45f2feef637c3b8ad7a1447ab4c9d1dda96","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"2b1124a373d813bd96945c6a91d8c5186ff0420a","targetTree":"ce8ca1b0281a56331b3600953f2312c0d7fdc7f1","verification":[{"command":"GOMAXPROCS=14 GOMEMLIMIT=5GiB task build (pinned Go1.26.8, combined candidate)","exitCode":0,"result":"pass"},{"command":"git diff TARGET_BASE CANDIDATE --check; exact reviewed tree/scope/cleanliness checks","exitCode":0,"result":"pass"},{"command":"unchanged target-base/source checks; git merge --ff-only CANDIDATE; git merge-base --is-ancestor CANDIDATE main","exitCode":0,"result":"pass"}]}
<!-- mtn:end-entry:v1 -->
