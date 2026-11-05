---
identity:
  kind: task
  taskID: rex-whichkey-15-05-packaging-docs
  uuid: c1ce9955-340f-44a1-987e-6581000c1493
structure:
  order: 5
  parentRef: phases/rex-whichkey-15-code-owned-runtime.md
scope:
  requirements:
    - REQ-003
    - REQ-005
    - REQ-014
    - REQ-016
  documents:
    - ref: designs/rex-whichkey-runtime-ownership.md
      hash: 'sha256:253d40dc6204148a88924f7c860ad092f86f3b541a618330f112c6202b3aaa65'
    - ref: prds/rex-whichkey.md
      hash: 'sha256:97614db611a6f35a44d0db32ccf2bf4d4854d6ffea3dfd5375959e59259f01ff'
    - ref: prds/rex-whichkey-usability.md
      hash: 'sha256:966ef70c264360b4820e5faf510496166e91f0effeae479ab7400a3792de88e9'
implementation:
  language: nix
  gates:
    - ID: nix-parse
      command: >-
        nix-instantiate --parse flake.nix && nix-instantiate --parse
        nix/package.nix && nix-instantiate --parse nix/home-manager.nix
      required: true
    - ID: whitespace
      command: git diff --check
      required: true
  ownership:
    paths:
      - nix/package.nix
      - nix/home-manager.nix
      - README.md
      - examples/my_whichkey.lua
      - docs/launch-contract.md
      - docs/cli-gui-delegation.md
      - docs/tab-rename-input.md
      - docs/session-owned-ui.md
      - docs/remote-compatibility.md
    sharedPaths:
      - flake.nix
    resources: []
    interfaces: []
lifecycle:
  state: planned
  sourceEventID: null
  changedAt: '2026-11-05T00:40:23.419Z'
title: Reduce packaging to installation and document code-owned runtime
mtnCreateRetry:
  key: rex-whichkey-15-05-packaging-docs/create/1
  intent: b8d16daba919285ecc2cea88f87d4aeee9b483b7e01b7cb6cc7f9a9e65e8ede5
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-05T00:40:24.405Z'
dateModified: '2026-11-05T00:40:29.576Z'
tags:
  - task
type: task
id: 5c9e50f6-10c8-4e88-9b5f-d30dae255613
projects:
  - '[[phases/rex-whichkey-15-code-owned-runtime.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-15-04-lua-bridge.md]]'
    reltype: FINISHTOSTART
---
# Reduce packaging to installation and document the code-owned runtime

## Deliverable and inputs

Update packaging and user-facing documentation after the coherent Go runtime and original-VM Lua consumer are integrated. Nix/Home Manager may build/install the executable and normal assets, but must not implement executable discovery, helper selection, protocol/version pairing or runtime behavior.

## Acceptance criteria

1. nix/package.nix, nix/home-manager.nix and flake.nix contain only appropriate build/install/declarative asset-linking responsibilities for this integration. No generated constructor paths, Nix-store/repository executable literals, source substitution, wrapper runtime resolver or ldflag runtime-path workaround is needed to make startup work.
2. Preserve normal optional integration assets and user ownership of init.lua. Do not overwrite external consumer configuration, activate a profile/system, change the Rex server environment or reload/restart anything. External examples are migration documentation, not authority to mutate another repository.
3. Public help/README/examples explain one supported code-owned startup model for direct/local development, packaged installation and optional Lua customization. init.lua examples contain only keybindings, mapping/callback setup and registration; no required helper/rex_bin_dir/socket literals or manual binary/Lua version pairing.
4. Explain that normal installation of a software update is distinct from runtime discovery; remove repeated-flake-refresh/copy-path instructions as a solution to normal launch. Describe genuine unsupported-context diagnostics without promising guessed automatic discovery or fabricating an observed runtime success.
5. Document preserved custom callbacks, separate actual caller/GUI target, safe completion, inline default tab/session rename and no uncertain fallback/retry. Synchronize docs/launch-contract.md, docs/cli-gui-delegation.md and rename/session UI documentation with the actual integrated code.
6. Provide a coordinator handoff mapping all eight parent criteria and REQ coverage to the integrated contract, resolver, Go bootstrap, Lua consumer and packaging/docs artifacts. Identify all remaining limitations and stale compatibility assumptions. Do not claim phase completion from a documentation edit, build or RPC acknowledgment.

## Ownership and dependencies

Prerequisite: tasks/rex-whichkey-15-04-lua-bridge.md; resolver/Go producer tasks are transitively required. This is a genuine final cutover/documentation dependency so examples cannot ship ahead of the new consumer.
Own nix/package.nix, nix/home-manager.nix, README.md, examples/my_whichkey.lua, docs/launch-contract.md, docs/cli-gui-delegation.md, docs/tab-rename-input.md, docs/session-owned-ui.md and docs/remote-compatibility.md. Reserve flake.nix as an exclusive shared path. Read docs/runtime-bootstrap.md and docs/mapping-format.md as stable producer contracts; changes to those interfaces require coordinator impact review, not a second independent schema. Do not change Taskfile/lint/hook standards or unrelated dependencies/locks to fix existing lint/toolchain problems.

## Review, checks and final-phase responsibility

Language is Nix with documentation in its delivery scope; require independent combined Nix/source/documentation review. Describe acceptance cases for package layout, source-only local development, optional integration assets, user-config preservation, relocatable/code-owned selection, mismatched old installs and safe migration instructions. No Nix activation, input update, live menu/rename trial, test runner, experiment or config reload is authorized by task creation.

At implementation time, perform only authorized parse/source/build checks. Existing doCheck/install-check policy is not silently changed. Keep unperformed runtime/TDD checks and existing lint/toolchain blockers explicit rather than hiding them or weakening standards. The coordinator, not this worker, owns the final combined-revision source review/check admission, requirement/criterion coverage, integration proof, metadata updates and any phase completion decision. No implementation has been scheduled by this task breakdown.

## Coverage

Phase criteria 1, 3, 5 and 8 plus combined criterion coverage reporting; requirements REQ-003, REQ-005, REQ-014 and REQ-016. This task removes packaging/configuration workarounds; it must not recreate the resolver in Nix.
