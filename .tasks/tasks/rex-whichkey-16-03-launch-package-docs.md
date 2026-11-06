---
identity:
  taskID: rex-whichkey-16-03-launch-package-docs
  kind: task
  uuid: 4357adda-da01-48e3-8e39-d764c7dff3a4
structure:
  order: 3
scope:
  requirements:
    - REQ-001
    - REQ-003
    - REQ-013
    - REQ-014
    - REQ-015
    - REQ-016
    - REQ-017
  documents:
    - ref: designs/rex-whichkey-cli-only-menu.md
      hash: 'sha256:23add062dac4326fe35e7e178a6a39f4c4d7631b4c8e164feb988b9aba07e21b'
  dependencies:
    - ref: tasks/rex-whichkey-16-02-go-cli-menu.md
      relation: finish-to-start
implementation:
  language: lua
  gates:
    - ID: whitespace
      command: git diff --check
      required: true
    - ID: nix-parse-only
      command: >-
        nix-instantiate --parse nix/package.nix >/dev/null && nix-instantiate
        --parse nix/home-manager.nix >/dev/null && nix-instantiate --parse
        flake.nix >/dev/null
      required: true
    - ID: nix-format-only
      command: alejandra --check nix/package.nix nix/home-manager.nix flake.nix
      required: true
    - ID: combined-build
      command: task build
      required: true
    - ID: coordinator-observed-combined-public-menu-dismissal-cleanup
      command: task run
      required: true
  ownership:
    paths:
      - lua/rex_whichkey.lua
      - lua/rex_whichkey/config.lua
      - lua/rex_whichkey/evaluate.lua
      - lua/rex_whichkey/mapping.lua
      - lua/rex_whichkey/dispatch.lua
      - lua/rex_whichkey/defaults/init.lua
      - lua/rex_whichkey/defaults/sessions.lua
      - lua/rex_whichkey/defaults/tabs.lua
      - lua/rex_whichkey/defaults/panes.lua
      - lua/rex_whichkey/defaults/app.lua
      - lua/rex_whichkey/defaults/hosts.lua
      - lua/rex_whichkey/runtime/anchor.lua
      - tests/fixtures/phase15-lua
      - tests/lua/action_context.lua
      - tests/lua/cli_delegation.lua
      - tests/lua/configured_rename.lua
      - examples/rex-whichkey-keybinding.lua
      - examples/my_whichkey.lua
      - examples/README.md
      - nix/package.nix
      - nix/home-manager.nix
      - flake.nix
      - README.md
      - docs/actions.md
      - docs/launch-contract.md
      - docs/runtime-bootstrap.md
      - docs/remote-compatibility.md
      - docs/cli-gui-delegation.md
      - docs/session-owned-ui.md
      - docs/tab-rename-input.md
      - docs/appearance.md
      - docs/acceptance.md
      - tests/packaging/home-manager/README.md
    sharedPaths:
      - docs/mapping-format.md
      - examples/my_whichkey.json
    resources:
      - 'build-output:bin'
      - coordinator-heavy-gate-slot
      - coordinator-owned-public-menu-observation-origin
    interfaces:
      - CLI-ONLY-LAUNCH
      - ORDINARY-PACKAGE-LAYOUT
lifecycle:
  state: planned
  sourceEventID: null
  changedAt: '2026-11-06T18:41:59.658Z'
title: >-
  Ship minimal keybinding and ordinary package with migration and combined app
  validation
mtnCreateRetry:
  key: rex-whichkey-16-03-launch-package-docs-create-v1
  intent: c736d4036603549d61af15280c2b29b84bf95b85d6d4175d96de3db2345f0338
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-06T18:42:00.300Z'
dateModified: '2026-11-06T18:47:36.402Z'
tags:
  - task
type: task
id: 3fa36f79-07ef-4a73-a092-d108be5aa828
projects:
  - '[[phases/rex-whichkey-16-cli-only-menu.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-16-02-go-cli-menu.md]]'
    reltype: FINISHTOSTART
---
Consume task2's running public CLI and task1's documented data schema. Reduce lua/rex_whichkey.lua to a minimal launch function used by a documented rex.bind function keybinding in examples/rex-whichkey-keybinding.lua. Use documented rex.exec.start argument-array execution of installed rex-whichkey, detach=true, no action-duration wait or command timeout; pass only exact ctx.session_id/ctx.block_id/ctx.client_id as origin/app intent using process environment. Go owns validation, endpoint/CLI discovery and menu lifecycle. No rex.action/register_launch/register_completion, metadata query, rex.call/session query/mutation, Lua JSON config compiler, registry, callback, completion, anchor/path search or helper process controller. Do not copy/stringify ctx.server or assume terminal-only environment is inherited correctly. Throw an honest immediate launch error; a returned detached process handle proves only start, not displayed menu. The function returns true for handled binding. Keep key spelling an example, never install/alter the user's binding. No physical key/reload result is claimed without observation and separate permission.
Retire the listed obsolete Lua config/evaluate/mapping/dispatch/default/anchor assets from the shipped production tree. Preserve old tests and assertions: three tests/lua scripts explicitly prepend lua paths and require legacy modules, so isolate the original legacy Lua tree under tests/fixtures/phase15-lua and change only those scripts' loader search-path plumbing to that historical fixture namespace. Preserve fixture behavior/assertions and all Go tests; no harness execution and no test deletion/skip/weakening. Fixtures must be excluded from shipped assets and cannot be loaded by the minimal production launcher. Retain any historical test limitations honestly instead of promising tests were run. Retire examples/my_whichkey.lua executable-options example in favor of task1's JSON data example, with migration instructions; do not silently discard a user's old custom data or attempt automatic Lua conversion. Explain unsupported callbacks/functions as deferred, not silently translated or accepted.
Update nix/package.nix to install an ordinary real bin/rex-whichkey and only minimal launcher/data/docs. If a legitimate optional licensed CLI dependency is present, install its symlink at libexec/rex-whichkey/rex for Go's code-owned discovery. Remove the old real-image-inside-Lua-tree, anchor, runtime/deps and generated-helper layout; no wrapper, runtime Nix resolver, policy flags or ldflags. Preserve upstream pins, Linux loader/PAM/terminfo/SSH assets, Darwin signature and truthful Intel-macOS artifact absence; retain doCheck=false/doInstallCheck=false. Update Home Manager to link only the needed launch module, no obsolete recursive module tree, init.lua, generated data config or whole Rex directory takeover. Preserve conflict safety and opt-in package/configDirectory semantics; remove obsolete namespace linkage from newly generated configuration without activating it. Update flake descriptions, README, examples and every listed active launch/config/action/bootstrap/remote/rename/appearance/session-owned/acceptance and packaging guide. Mark legacy sections historical rather than presenting old APIs/registrations as prerequisites, and preserve historical evidence/commits. Document migration, current defaults, data path/precedence, provider/action semantics, simple launch lifetime, endpoint/GUI-vs-caller safety, discovery and exact owned cleanup. Do not rewrite historical docs/rex-whichkey.prd or .tasks history.
Own source review of minimal Lua/Nix/docs and integration consistency; Nix gates are parse/format/whitespace only, no flake check/evaluation/build/activation or indirect tests. Final combined source requires normal task build and coordinator-observed actual public app/menu, ordinary dismissal and exact popup/terminal/private-artifact cleanup after review. Include one benign Go-owned dynamic selection or inline rename only on an explicitly coordinator-owned disposable resource under actual-run authority; never mutate a real user tab/session merely for evidence. Record precisely which actions/entrypoints were exercised and which remain source-reviewed only. The keymapping may be source-reviewed/documented without config deployment; do not block public app validation on an unapproved user reload. Task completion includes accurate docs and combined observations, not a disconnected verification task or an assertion that Phase15/unrelated phases are fixed.