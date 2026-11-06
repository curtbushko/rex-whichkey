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
  state: complete
  sourceEventID: rex-whichkey-p16-t01/dispatch-release/2
  changedAt: '2026-11-06T19:37:06.362Z'
title: Define and implement data-only menu configuration and current-default adapter
mtnCreateRetry:
  key: rex-whichkey-16-01-data-config-create-v1
  intent: 75ac71a6c8edef247140c17e074f891f865242e6b84b415c2ad379774ddf2ae4
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-06T18:41:56.479Z'
dateModified: '2026-11-06T19:39:40.086Z'
tags:
  - task
type: task
id: 63ae3859-2488-4f2d-be7b-32f477e3d198
projects:
  - '[[phases/rex-whichkey-16-cli-only-menu.md]]'
completedDate: '2026-11-06'
---
Implement DATA-MENU-CONFIG v1 in the existing flat Go package, using encoding/json and existing ConfigOptions/ResolveConfig/Mapping/ParseKeys validation rather than a new framework or Lua evaluator. Add LoadMenuConfig(explicit *string) (Config, error) in data_config.go. Resolve built-in nativeDefaults first, then one selected JSON options file. Selection is present --config > nonempty REX_WHICHKEY_CONFIG > $XDG_CONFIG_HOME/rex-whichkey/config.json, or $HOME/.config/rex-whichkey/config.json when XDG_CONFIG_HOME is absent. Explicit empty paths and invalid relative XDG_CONFIG_HOME fail; only a missing implicit default file means defaults. Explicit/environment-selected missing files, unreadable/malformed files and unknown versions fail clearly. Never write configuration. File schema requires version:1, permits optional position/title/icon/icons/spec, and uses the existing flat full-sequence Mapping data fields except callback_ref. Omission retains defaults; explicit empty strings/false remain data; spec and icons replace whole values; spec:[] clears bindings. Bound input at 4 MiB; reject duplicate/noncanonical/unknown fields, null, invalid Unicode, executable/function/callback/hook fields, unsupported provider/action forms and key/prefix conflicts, including disabled entries. Keep cond boolean-only and hidden-but-active lookup. Produce a concrete action/provider argument catalogue and a JSON migration example in docs/mapping-format.md and examples/my_whichkey.json. Public supported action vocabulary is current defaults plus explicit session.select; session.switch is Go menu intent, not a Lua or native picker action. Providers sessions/windows/panes/known-hosts are named, non-executable group markers with reserved subtrees; known-hosts stays listing-only. Document generated retained identities, action-specific arguments and forced confirmation for destructive resource operations. Preserve defaults exactly as defaults.go currently supplies them: Sessions/Tabs, v/s/x/e/c; ws becomes a Go session list. Do not restore Zoom, directional focus, numbered tabs, app chooser or a default host row. Keep provider implementation for task2, but implement its data validation/catalogue now. Preserve existing legacy Config/Completion/CallbackRef helpers needed by unchanged tests and the transitional caller; isolate them from the new loader instead of widening public JSON to accept callbacks or breaking compilation. This is a real loader/default adapter deliverable, not an approval-only contract task. Consumer wiring follows in task2. Own checks: formatting, whitespace, static/type review and normal task build, no tests/TDD/harnesses. Review rejection/default/empty/Unicode/conflict/action/provider paths by source inspection, not experiments; record unexecuted behavior honestly.
<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA2VDE5OjM5OjQwLjA3OFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxNi10MDEvc3VtbWFyeS8yIiwiaW50ZW50IjoiZWQ4M2FlZWZjNzA5ODQ2MmQzYjFjYzI0ZWI0YWE5NTAyYjE4NDkwMmMyNjA4M2E4NjFlNTc2YjM1MzVhNzg1YiJ9 -->
### 2026-11-06T19:39:40.078Z · summary · unspecified

{"acceptedScope":"Task1 producer portions only: strict v1 JSON data loader/source selection, unchanged reduced defaults, static builtin/provider contract and migration example. Shared conflict validator now sorts complete canonical paths without new key limits. Public CLI/provider dispatch/runtime completion is downstream, not delivered by this task.","approvalSummary":["User approved CLI-only direction REXWHICHKEY-20261106-114935 and then start it execution.","Coordinator authorized bounded plan rex-whichkey/phase16-task-plan/approved/1; current design-only governing scope and phase authority2 recorded without legacy migration.","Fresh same-family independent combined review rex-whichkey-p16-t01/review/2 approved exact build2 head7a87a7e/treeafe0ec9; source reviewer1 finding REX-P16-T01-001 fixed/reviewed, not erased."],"attempt":2,"cleanup":{"facade":"complete invoked and retried after delivery. Native/protocol completion confirmed. Its immutable cleanup jq syntax defect rejected all(.events[] as $event;...); eligible successful worktrees removed manually with exact scope/clean/source preservation checks, no tool/schema/standard change.","releasedAllRuntimeReservations":true,"removedCleanDeliveredWorktrees":["p16-t01-builder-a01","p16-t01-builder-a02","p16-t01-reviewer-a02","p16-t01-integration-a02"],"retainedNegativeReview":{"head":"253e537166dadf333639db496da1530d4249b88b","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"75ff9ad4-a558-4295-8b20-855dafe77a8f"},"process":"confirmed zero exit/no child or foreground process","provider":{"displayName":"p16-t01-reviewer-a01","name":"rex","resources":{"block":"block:01a1b240-ca9b-7b17-a92e-26891ab76023","endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","session":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","window":"window:01a1b240-ca9d-748f-ad47-cd0802861b55"}},"reason":"Negative review window/worktree retained per failed-resource policy; no authorization for failed/global resource deletion.","rexBlockID":"block:01a1b240-ca9b-7b17-a92e-26891ab76023","rexSessionID":"session:01a1a935-0b64-7c7b-8e83-5d35011bdee2","rexWindowID":"window:01a1b240-ca9d-748f-ad47-cd0802861b55","runID":"rex-whichkey-p16-t01-reviewer-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p16-t01-reviewer-a01"},"successfulWorkerWindowsConfirmedAbsent":true},"completionEventID":"rex-whichkey-p16-t01/completion/2","eventID":"rex-whichkey-p16-t01/summary/2","files":["data_config.go","config.go","mapping.go","defaults.go","docs/mapping-format.md","examples/my_whichkey.json",".mtn-tmp/p16/task1-a02-result.json",".mtn-tmp/p16/task1-a02-review-result.json",".mtn-tmp/p16/task1-integration.json",".mtn-tmp/p16/task1-post-cleanup-worktrees.txt"],"followUp":"Task2 must consume LoadMenuConfig and provider validators, canonicalize retained keys, preserve pure helper dependencies, address NewMenu cumulative-prefix allocation (review2 NB-001), and actually run the normal public menu after review. Task3 ships minimal launcher/package/docs. Neither is completed here.","integrationEventID":"rex-whichkey-p16-t01/integration/2","kind":"summary","limitations":["No tests/TDD/race/Lua/parser/FIFO harnesses, benchmarks or product invocation; actual app still not established by task1.","Optional independent full lint failed84 inherited diagnostics; lint:ai not reached. No full-lint pass claimed.","Normal compilation/static checks on Darwin arm64 only; no Linux/remote/platform runtime claim.","Fresh independent same-family review, not cross-family verification.","Negative reviewerA01 exact window/worktree retained; successful owned resources cleaned and all Git branches/history/private artifacts preserved.","Phase15 original runtime failure/incomplete status and user configuration remain unchanged."],"nativeEvidenceCommit":"df6df4efc8e5e736175fa0a04e9bd19ac3883d32","recordRef":"tasks/rex-whichkey-16-01-data-config.md","requirements":["REQ-003","REQ-004","REQ-006","REQ-007","REQ-010","REQ-012"],"schemaVersion":1,"specHash":"sha256:c941d8c673c675675aa3fdc5dfd66ce9b96a2e3eafc45f767dba189fb60b1f7d","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"4d675c2b23962dcfcb4f1138227e295a263182dc","targetTree":"ee15ffce9674754da4322e3f1732c38bba8a3061","verification":[{"artifact":".mtn-tmp/p16/task1-integration-build.log","command":"Combined speculative candidate task build (GOMAXPROCS14 GOMEMLIMIT5GiB)","result":"pass"},{"artifact":".mtn-tmp/p16/task1-integration-vet.log","command":"Combined candidate task lint:vet","result":"pass"},{"artifact":".mtn-tmp/p16/task1-integration-format.log","command":"Owned gofmt and target-base-to-candidate git diff --check","result":"pass"},{"artifact":".mtn-tmp/p16/task1-integration-source-identity.txt","command":"Exact six-file source identity versus reviewed build2, unchanged lint bytes, clean candidate and target-base recheck before fast-forward; reviewed head ancestor of main","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
