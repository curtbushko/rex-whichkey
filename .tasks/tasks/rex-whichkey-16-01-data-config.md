---
identity:
  taskID: rex-whichkey-16-01-data-config
  kind: task
  uuid: 564fdaa2-a66d-4280-90c1-f930da78bf4a
structure:
  order: 1
scope:
  requirements:
    - REQ-003
    - REQ-004
    - REQ-006
    - REQ-007
    - REQ-010
    - REQ-012
  documents:
    - ref: designs/rex-whichkey-cli-only-menu.md
      hash: 'sha256:23add062dac4326fe35e7e178a6a39f4c4d7631b4c8e164feb988b9aba07e21b'
  dependencies: []
implementation:
  language: go
  gates:
    - ID: owned-format
      command: test -z "$(gofmt -l data_config.go config.go mapping.go defaults.go)"
      required: true
    - ID: whitespace
      command: git diff --check
      required: true
    - ID: build
      command: task build
      required: true
    - ID: static-type-review
      command: 'task lint:vet'
      required: true
    - ID: lint-baseline-diagnostics
      command: task lint
      required: false
  ownership:
    paths:
      - data_config.go
      - config.go
      - mapping.go
      - defaults.go
      - docs/mapping-format.md
      - examples/my_whichkey.json
    sharedPaths: []
    resources:
      - 'build-output:bin'
      - coordinator-heavy-gate-slot
    interfaces:
      - DATA-MENU-CONFIG
      - GO-ACTION-PROVIDER-CATALOGUE
lifecycle:
  state: planned
  sourceEventID: null
  changedAt: '2026-11-06T18:41:55.840Z'
title: Define and implement data-only menu configuration and current-default adapter
mtnCreateRetry:
  key: rex-whichkey-16-01-data-config-create-v1
  intent: 75ac71a6c8edef247140c17e074f891f865242e6b84b415c2ad379774ddf2ae4
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-06T18:41:56.479Z'
dateModified: '2026-11-06T18:47:32.334Z'
tags:
  - task
type: task
id: 63ae3859-2488-4f2d-be7b-32f477e3d198
projects:
  - '[[phases/rex-whichkey-16-cli-only-menu.md]]'
---
Implement DATA-MENU-CONFIG v1 in the existing flat Go package, using encoding/json and existing ConfigOptions/ResolveConfig/Mapping/ParseKeys validation rather than a new framework or Lua evaluator. Add LoadMenuConfig(explicit *string) (Config, error) in data_config.go. Resolve built-in nativeDefaults first, then one selected JSON options file. Selection is present --config > nonempty REX_WHICHKEY_CONFIG > $XDG_CONFIG_HOME/rex-whichkey/config.json, or $HOME/.config/rex-whichkey/config.json when XDG_CONFIG_HOME is absent. Explicit empty paths and invalid relative XDG_CONFIG_HOME fail; only a missing implicit default file means defaults. Explicit/environment-selected missing files, unreadable/malformed files and unknown versions fail clearly. Never write configuration. File schema requires version:1, permits optional position/title/icon/icons/spec, and uses the existing flat full-sequence Mapping data fields except callback_ref. Omission retains defaults; explicit empty strings/false remain data; spec and icons replace whole values; spec:[] clears bindings. Bound input at 4 MiB; reject duplicate/noncanonical/unknown fields, null, invalid Unicode, executable/function/callback/hook fields, unsupported provider/action forms and key/prefix conflicts, including disabled entries. Keep cond boolean-only and hidden-but-active lookup. Produce a concrete action/provider argument catalogue and a JSON migration example in docs/mapping-format.md and examples/my_whichkey.json. Public supported action vocabulary is current defaults plus explicit session.select; session.switch is Go menu intent, not a Lua or native picker action. Providers sessions/windows/panes/known-hosts are named, non-executable group markers with reserved subtrees; known-hosts stays listing-only. Document generated retained identities, action-specific arguments and forced confirmation for destructive resource operations. Preserve defaults exactly as defaults.go currently supplies them: Sessions/Tabs, v/s/x/e/c; ws becomes a Go session list. Do not restore Zoom, directional focus, numbered tabs, app chooser or a default host row. Keep provider implementation for task2, but implement its data validation/catalogue now. Preserve existing legacy Config/Completion/CallbackRef helpers needed by unchanged tests and the transitional caller; isolate them from the new loader instead of widening public JSON to accept callbacks or breaking compilation. This is a real loader/default adapter deliverable, not an approval-only contract task. Consumer wiring follows in task2. Own checks: formatting, whitespace, static/type review and normal task build, no tests/TDD/harnesses. Review rejection/default/empty/Unicode/conflict/action/provider paths by source inspection, not experiments; record unexecuted behavior honestly.