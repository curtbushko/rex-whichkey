---
identity:
  kind: task
  taskID: rex-whichkey-18-01-path-startup
  uuid: 7bda23ca-a3d6-4f50-91b3-f69ecdf42dfe
structure:
  order: 1
scope:
  requirements:
    - REQ-001
    - REQ-002
    - REQ-011
    - REQ-013
    - REQ-014
    - REQ-015
    - REQ-016
    - REQ-017
  documents:
    - ref: designs/rex-whichkey-startup-path.md
      hash: 'sha256:c7a6d0d736fbbb3925173391431a05eb2c8c65273c2a81e2e5e7d0ac23e1c85c'
    - ref: designs/rex-whichkey-cli-only-menu.md
      hash: 'sha256:23add062dac4326fe35e7e178a6a39f4c4d7631b4c8e164feb988b9aba07e21b'
implementation:
  language: go
  ownership:
    paths:
      - rex/runtime.go
      - rex/exec.go
      - rex/rex.go
      - rex/endpoint.go
      - rex/hosts.go
      - runtime_bootstrap.go
      - launch.go
      - native.go
      - popup.go
      - terminal.go
      - hosts.go
      - nix/package.nix
      - nix/home-manager.nix
      - README.md
      - docs/runtime-bootstrap.md
      - docs/launch-contract.md
      - docs/remote-compatibility.md
      - tests/packaging/home-manager/README.md
    sharedPaths: []
    resources: []
    interfaces:
      - STARTUP-PATH-LOOKUP
      - STARTUP-NO-DURABILITY-SYNC
  gates:
    - ID: whitespace
      command: git diff --check
      required: true
    - ID: source-format
      command: test -z "$(gofmt -l ./*.go ./rex/*.go ./cmd/*.go)"
      required: true
    - ID: normal-build
      command: task build
      required: true
    - ID: go-vet
      command: 'task lint:vet'
      required: true
    - ID: full-lint-diagnostic
      command: task lint
      required: false
lifecycle:
  state: complete
  sourceEventID: rex-whichkey-p18-t01/summary/1
  changedAt: '2026-11-07T01:26:51.424Z'
title: '18-01: Simplify PATH and private-image startup'
mtnCreateRetry:
  key: rex-whichkey-p18/create/18-01/v1
  intent: 7d07c5f1b57b7f35cfa5bf29406cb4c039b7bdf81b36df2f9088cdd6af726a51
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T00:54:22.508Z'
dateModified: '2026-11-07T01:26:52.048Z'
tags:
  - task
type: task
id: e6f3df2c-41cb-4a2a-88d3-dffeae7edaed
projects:
  - '[[phases/rex-whichkey-18-startup-speed.md]]'
completedDate: '2026-11-06'
---
Implement only PATH/CLI and invocation-image startup simplification. Replace ResolveCLI's REX_BIN_DIR/layout/unambiguous-absolute-PATH policy with ordinary execution-host executable lookup and clear absence errors. Reuse the selected PATH result within the invocation; do not repeatedly rediscover, hash or negotiate it. Remove production binary/link/content receipts, version/API/schema admission, checkMethod startup admission and server.status/config-generation before/after wrappers across constructors, resource execution and popup-child consumption. Retain bounded argument-array execution, explicit canonical endpoint/no-autostart semantics, cancellation, strict actual result decoding, action-specific built-in metadata after selection, sanitized honest uncertainty and no retries/fallbacks. Preserve useful decoding helpers used by actual action results; do not confuse them with rejected schema negotiation. No new public executable nomination flag or path authority.

Simplify capture/copy/child bootstrap coherently: retain the existing invocation-private copy and parent/renderer split unless concrete implementation evidence warrants an in-scope simplification; removal is NOT required. Remove executable rehash/receipt validation throughout capture/copy/child startup and the copy-only file.Sync. Complete ordinary copy/write/close before start, preserve bounded image copying, private directory/fixed-slot ownership/no-follow checks, result integrity, lifetime pipe and exact source-handle retirement. Do not blanket-remove commitPopupResult's result Sync, locks, JSON bounds or lifecycle checks. Update terminal/popup consumers of Capabilities/receipt endpoint fields in the same change so no unsafe temporary fallback or fake compatibility object is introduced. Retain endpoint independently as invocation context.

Packaging: replace the existing optional private libexec Rex symlink with ordinary bin/rex exposure of the SAME already pinned optional dependency, where available; no wrapper, runtime fallback, dependency/pin/vendor change, flake.nix edit or activation. Update README, runtime/bootstrap/launch/remote and Home Manager package guidance, including Intel macOS limitations and running-server PATH inheritance. Preserve source-only installation caveats. Lua stays unchanged and launch-only. Own only the listed literal paths; overlapping later ownership is strictly serial. Preserve all tests/assertions/Taskfile/hooks/lint settings. If retired helper references actually obstruct compilation, inspect minimal safe test-only isolation with the coordinator; do not add production runtime stubs, discard assertions or claim tests ran.

Acceptance/review: trace the public launch and every CLI subprocess to show rejected validation is absent, actual result/identity checks remain and effect invocation remains once-only. Independently review the changed source and exact owned cleanup. Run owned-file gofmt/whitespace, normal task build and task lint:vet under the pinned budget; inspect Nix changes as source, without evaluation/build/activation or a packaging harness. Optional unchanged full lint is diagnostic only; supplied Phase16 baseline is 62 diagnostics, not a green gate.

COORDINATOR FIRST INTEGRATION CHECK, BEFORE ANY PHASE18 SOURCE INTEGRATION: build the unmodified committed baseline at 9634246b7e4edbc8bed59210466cac492ddf2f0f in an appropriate clean baseline checkout, preserving the parent's separate flake.nix edit. Run the ordinary public binary with no private modes from a real explicitly owned Rex origin. Record exact public command/binary/source/tree, host/PATH/endpoint, actual full origin IDs, config/defaults, GUI/caller intent, terminal/font/cell context and observation method. From public process launch observe separately first visible static frame, required inventory-data readiness and relevant action/rename readiness; if readiness coincides, say so. Capture observations/timestamps with uncertainty and limits, not invented precision. Observe normal dismissal/restoration/exact popup/private cleanup. This is coordinator-owned normal app observation, not a planner/builder run or custom timing harness. Baseline is a hard pre-integration ordering requirement; an API error or absent menu cannot be replaced by help/compile/private ACK. Do not start candidate integration before the baseline artifact exists. After integration/build, run the same ordinary public path/context to verify PATH startup, real display/navigation/dismissal/cancellation and exact cleanup; record intermediate observations without claiming full asynchronous completion or an SLA. Only the sole coordinator integrates and publishes evidence; do not launch agents or mutate records from a builder/reviewer leaf.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDAxOjI2OjUwLjUxNFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOC10MDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiOWM0YWE0MTM5Y2NkMzhkZjhmZjBmYmQ0Mjk0ODNlZjEzNDg4OTg3YTNiODA0NjI0ZjUwZTgzZDE5ZTBlZDYxMSJ9 -->
### 2026-11-07T01:26:50.514Z · summary · unspecified

{"acceptedScope":"Normal PATH-only Rex lookup; no executable receipts/hashing/status/version/API-schema admission wrappers in ordinary parent/child/resource paths. Endpoint and actual result/identity/once-only cleanup preserved. Copied-image durability Sync removed, result Sync/lifecycle intact. Same optional pinned Rex exposed in bin, active PATH docs. Reviewed combined public binary actually displayed root/dynamic/inline rename form, canceled/dismissed exit0 with exact cleanup and no REX_BIN_DIR.","approvalSummary":["User startup decision REXWHICHKEY-20261106-191113 and explicit finish-all execution/bounded task planning authority","Current approved Phase18 task plan1 and exact design fingerprints","Fresh same-family combined review1 approved exact build676728c/tree dbd4981; coordinator combined build/vet/source identity and normal public runtime before target advance"],"attempt":1,"cleanup":{"facade":"Native/protocol completion verified; immutable complete cleanup jq syntax defect persists. Official complete attempted, exact clean stopped ancestor worktrees manually removed after trusted results and source/evidence preservation.","removedCleanDeliveredWorktrees":["p18-t01-builder-a01","p18-t01-reviewer-a01","p18-t01-integration-a01"],"runtimeReservationsReleased":true,"successfulWorkerWindowsConfirmedAbsent":true},"eventID":"rex-whichkey-p18-t01/summary/1","files":["README.md","docs/launch-contract.md","docs/remote-compatibility.md","docs/runtime-bootstrap.md","hosts.go","nix/home-manager.nix","nix/package.nix","popup.go","rex/exec.go","rex/hosts.go","rex/rex.go","rex/runtime.go","runtime_bootstrap.go","terminal.go","tests/packaging/home-manager/README.md",".mtn-tmp/p18/task1-build-result.json",".mtn-tmp/p18/task1-review-result.json",".mtn-tmp/p18/task1-builder-evidence",".mtn-tmp/p18/task1-reviewer-evidence",".mtn-tmp/p18/task1-integration.json",".mtn-tmp/p18/task1-public-evidence.json",".mtn-tmp/p18/baseline-public-evidence.json",".mtn-tmp/p18/task1-observed-rex-whichkey"],"kind":"summary","limitations":["Full unchanged lint62 inherited failures/AI lint not reached, not global green","No tests/harnesses/benchmarks or user config/deploy/reload/remote/platform claims","Async work remains later tasks","First-frame captures baseline~2.74s/candidate~0.80s are coarse CLI/poll upper bounds, not exact render timings or final async comparison; candidate intentionally cleared REX_BIN_DIR.","Nix source formatter/source review only, no package activation or platform runtime claim."],"recordRef":"tasks/rex-whichkey-18-01-path-startup.md","requirements":["REQ-001","REQ-002","REQ-011","REQ-013","REQ-014","REQ-015","REQ-016","REQ-017"],"schemaVersion":1,"specHash":"sha256:47c41416c0c0e9066134535413aed3968955c3a216e48ea8361d793a99e630bf","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"cc3a320dbbf3ef05bbb959f71feed4f3dae1bced","targetTree":"13154b9dc62f55c95ebfd61f1998efd06a1bba03","verification":[{"artifacts":[".mtn-tmp/p18/task1-integration-build.log",".mtn-tmp/p18/task1-integration-vet.log"],"command":"Combined normal task build and task lint:vet, pinned14CPU/5GiB","result":"pass"},{"artifact":".mtn-tmp/p18/task1-candidate-source-identity.log","command":"Reviewed owned-source equality, combined whitespace, frozen source and unchanged target-base recheck; reviewed head ancestor of main","result":"pass"},{"artifact":".mtn-tmp/p18/task1-public-evidence.json","command":"Actual ordinary public PATH-only root/dynamic/inline rename form/cancel/dismissal exit0/terminal restoration/exact popup/private cleanup and owned-window closure","result":"pass"},{"artifact":".mtn-tmp/p18/baseline-public-evidence.json","command":"Required real public baseline recorded BEFORE source integration","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
