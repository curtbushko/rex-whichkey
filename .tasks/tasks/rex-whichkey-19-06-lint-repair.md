---
identity:
  kind: task
  taskID: rex-whichkey-19-06-lint-repair
  uuid: 9be113f9-7278-4e40-a907-e2ecf9703d6f
structure:
  order: 6
scope:
  requirements:
    - REQ-001
    - REQ-002
    - REQ-003
    - REQ-004
    - REQ-006
    - REQ-008
    - REQ-009
    - REQ-010
    - REQ-011
    - REQ-012
    - REQ-013
    - REQ-014
    - REQ-015
    - REQ-016
    - REQ-017
    - REQ-025
    - REQ-028
    - REQ-029
  documents:
    - ref: designs/rex-whichkey-lua-float-launch.md
      hash: 'sha256:f7e38d49287de319753732b09fe472a40dfadcef81e5245e2deac208e968ab70'
    - ref: designs/rex-whichkey-quality-gates.md
      hash: 'sha256:f3120762f7075c7c767e5f39552ba9c2b38b2e213d33d2472d07f0bc0c1cbdb2'
implementation:
  language: go
  ownership:
    paths:
      - README.md
      - action.go
      - cmd
      - config.go
      - configured_rename_test.go
      - configured_rename_wire_test.go
      - controller_legacy_support_test.go
      - data_config.go
      - defaults.go
      - direct_lifecycle_darwin_test.go
      - direct_menu.go
      - docs/acceptance.md
      - docs/appearance.md
      - hosts.go
      - launch.go
      - launch_legacy_support_test.go
      - launch_output_test.go
      - layout.go
      - lua_context_test.go
      - mapping.go
      - menu.go
      - menu_updates.go
      - native.go
      - native_app_owner.go
      - native_resources.go
      - native_sessions.go
      - native_targets.go
      - popup.go
      - popup_legacy_support_test.go
      - popup_updates.go
      - providers.go
      - pty_darwin_test.go
      - render.go
      - render_entry.go
      - rex
      - terminal.go
      - terminal_input.go
      - terminal_interaction_darwin_test.go
      - terminal_keys.go
      - terminal_keys_test.go
      - terminal_modes.go
      - terminal_poll_darwin.go
      - terminal_poll_darwin_test.go
      - terminal_poll_linux.go
      - terminal_poll_test.go
      - terminal_setup.go
      - text_entry.go
      - theme.go
      - theme_palette.go
    sharedPaths: []
    resources: []
    interfaces:
      - P19-QUALITY-GATES
  gates:
    - ID: whitespace
      command: git diff --check
      required: true
    - ID: build
      command: task build
      required: true
    - ID: tests
      command: task test
      required: true
    - ID: full-lint
      command: task lint
      required: true
lifecycle:
  state: complete
  sourceEventID: rex-whichkey-p19-t06/summary/3
  changedAt: '2026-11-07T18:06:38.176Z'
title: Repair full lint pipeline and preserve tested behavior
mtnCreateRetry:
  key: rex-whichkey-19-06-lint-repair-v1
  intent: f58c2b9e295931b61cb893b74ab9a8501469e4474ad5fc5e51d260e4495bb7c8
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T16:00:07.807Z'
dateModified: '2026-11-07T18:06:39.195Z'
tags:
  - task
type: task
id: 156d2b1f-0a29-4d31-bdc9-e983bf627fb9
projects:
  - '[[phases/rex-whichkey-19-lua-float-launch.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-19-05-test-repair.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-07'
---
After delivered Task05 test repair, fix ALL current full task lint diagnostics by code changes, not rules. Baseline53 covers contextcheck/goconst/gocyclo/govet-shadow/nestif/prealloc/revive/staticcheck; exact current log .mtn-tmp/p19/post-delivery-gates/lint.log, don't assume fixed count or inherited-only. Once golangci is green run lint:ai too as part of full task lint, fix any further real findings. User explicitly always tests/lint after changes and fix the problems: no diagnostic-only exemption at final acceptance.

TDD/regression checks mandatory: preserve real terminal input parser behavior (UTF8, escape, paste, timeouts/key handling), raw-mode restoration/error context, menu/render/palette/readiness/rename edits, bounded metadata/retained exact original/own/destination/GUI-generation validation/confirm/no retries/lifecycle. Split complex code into meaningful private functions with actual production callers, not token helpers, interfaces, compatibility stubs, new framework/Cobra/Viper/hexagonal layers or dependencies. Constants/comments/shadow/error strings corrected consistently without altering user-facing contract silently; keep strong assertions and add regression tests for behavioral changes. Archived historical support remains test-only. Defaults/native chrome1a deferral and detached host1b rule settled, don't reopen.

Required task build, task test (unit+race), FULL task lint all GREEN after source changes and on final head. Also architecture check only if existing config. No editing/suppressing lint config/Taskfile/hook/test checks, no deleted/skipped/weakened assertions, no count-as-pass on tool failure. Existing automated sandbox fixtures/mocks/parser/FIFO tests and TDD authorized by user's correction; no live leaf Rex changes/profile/config/install/reload/remote/deploy/push/benchmark/global cleanup/new deps/pins. Coordinator can resolve concrete toolchain mismatch using permitted existing tools, not weaken checks. Separate dirty flake.nix preserved.14CPU/6GiB shared/Go5GiB, one gate; pinned existing Go1.26.8/lint2.15, no nix develop/shellHook.

Literal ownership root Go/Go tests and rex/cmd Go files as listed. No unrelated docs/legacy migration. Report exact extra path before edit if needed. Fresh independent exact revision combined Go review must run ALL tests/full lint, zero blockers/no test exceptions, then coordinator normal combined/main gates and ACTUAL make run/menu/cancel/owned effect/cleanup. Commit owned source hooksPath=/dev/null; return actual native build commands/results/source IDs and caller list. No green claim until all pipeline stages pass; no unsupported chrome/profile/platform/global coverage claim.

## Independent review coverage corrections001/002

Builder1 all build/unit/race/full vet/golangci/AI gates green atc7c3156229145f75b87d3dbf8223fb3152d17d93, but independent fresh reviewer1 requests persistent behavioral tests through real current RunTerminal input/render pipeline and current RunDirectMenu/directSelection/retireDirectPopup dispatch order. Add isolated controlling-PTY tests and executable CLI fixtures as detailed in native review rex-whichkey-p19-t06/review/1. Must cover selection/dismissal; fragmentedUTF8/editing/paste; loneEscape50ms versus reports/sequences; palette interleaving and resize/fixed-rect pagination retaining state; feasible I/O/cancellation suppression/restoration. Current direct lifecycle must prove retirement/focus checks before exactly one effect; zero effects on detach uncertainty, retirement/focus failure, changed original placement/GUI generation, dismissal/cancel; uncertain mutations once/no retry/fallback. No live Rex in leaf tests, no framework/interfaces/dormant production injectionhooks, no dependency/pin/rule/assertion weakening; not every impossible OS close failure must be injected. Three precise Darwin PTY/test source paths authorized as routine regression engineering (platform limitation explicit, no t.Skip). Actual tests/full pipeline rerun, fresh entire original Task06 diff independent review then coordinator combined/main gates/real smoke remain. Original historical42 assertions unchanged. Review1 negative retained, not overruled by lint green.


<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDE4OjA2OjM2LjI5NloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDYvc3VtbWFyeS8zIiwiaW50ZW50IjoiZDI4MjEzYTRlNzk5Mjk0ZTg5YmIxOWEyNjRlNWZmMTliZjg4YTA0YWM0ZTk3OGY1YjgwYTM2NDc5MDhlZTlkYSJ9 -->
### 2026-11-07T18:06:36.296Z · summary · unspecified

{"acceptedScope":"Fixed all53 golangci and2 AI diagnostics by meaningful production-called refactors/linear host keys/explicit goroutine arguments/context/errors/style without weakening rules/assertions. Current RunTerminal and RunDirectMenu real controllingPTY/executableCLI regressions cover input/paste/palette/resize/pages/cancel/restoration/cleanup-before-once-effect and failure/no-retry. Five temporary mutation checks red then restored. Fresh whole-source review2 and exact-docs3 review approved. Combined/main standard build/unit/race/FULL lint0 and actual same-code makerun/ownedrename/cancel/cleanup verified. Current docs and prior prematureclosure history reconciled; historicalrecords/accepted chrome+host boundary unchanged.","approvalSummary":["User explicit always tests AND full lint after change/fixproblems supersedes carried prior restriction; no tests waiver","Initial c7 all-gatesgreen but independent review1 NEGATIVE001002 preserved; actual current orchestration coverage added, five red sensitivity mutations removed, unchangedproductionGo source verified","Fix2 cacf whole-diff independently approved; actual combined/main allgates+sourceexecution verified","Docs3 only3docs, fresh exacthead review and strict runtime/test/configsource equality toapproved+observed source2, final combined/main all normal gates0; no fabricated postdocs live binarylaunch"],"attempt":3,"cleanup":{"allNewPopupHostsAbsentAndOwnedOriginClosedBaselineRestored":true,"eligibleStoppedDeliveredWorktreesRemoved":true,"officialFacade":"Native done verified after existing immutable helper cleanupjq defect; exact eligible cleanup manually scoped","recoveryRetained":["p19-t05-builder-a01 blocked","p19-t06-builder-a01 source c7 negative review linked","p19-t06-reviewer-a01 negative001002","all earlier negative/blocked resources"],"successfulAutoWindowsAbsentVerified":true},"eventID":"rex-whichkey-p19-t06/summary/3","files":["direct_menu.go","terminal.go","terminal_input.go","terminal_keys.go","terminal_modes.go","pty_darwin_test.go","terminal_interaction_darwin_test.go","direct_lifecycle_darwin_test.go","hosts.go","providers.go","README.md","docs/acceptance.md","docs/appearance.md",".mtn-tmp/p19/quality-fix/task6-fix2-build-result.json",".mtn-tmp/p19/quality-fix/task6-review-result.json",".mtn-tmp/p19/quality-fix/task6-review2-result.json",".mtn-tmp/p19/quality-fix/task6-docs3-build-result.json",".mtn-tmp/p19/quality-fix/task6-docs3-review-result.json",".mtn-tmp/p19/quality-fix/task6-public-evidence.json",".mtn-tmp/p19/quality-fix/task6-integration3.json"],"kind":"summary","limitations":["Docs3 is only3docs; actualsource2 executable observed, no fabricated post-docs binary launch; all runtime/test/config blobs unchanged verified","Same-family independent exactsource/code and exactdocs reviews, not cross-family","Darwin only, no deployment/all-OS/all-failure-path claim; accepted native defaults/hostthroughdispatch/postexitlimit unchanged"],"recordRef":"tasks/rex-whichkey-19-06-lint-repair.md","requirements":["REQ-001","REQ-002","REQ-003","REQ-004","REQ-006","REQ-008","REQ-009","REQ-010","REQ-011","REQ-012","REQ-013","REQ-014","REQ-015","REQ-016","REQ-017","REQ-025","REQ-028","REQ-029"],"schemaVersion":1,"sourceIntegrationEventID":"rex-whichkey-p19-t06/integration/2","specHash":"sha256:379d1d8b72d6bad7a8769209eb6dc3a306acf69fead21c28db94ce7b6529f91a","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"61caadcf5ddbb1f77d5df39fe13c58f280aba2b7","targetTree":"2a48a94568738e8fa6213d6d6416ab37da2ef2bf","toolchain":{"ai":"SAME pinned1.0.0 source/deps compiler-only private rebuildGo1.26; originalNAR hash unchanged. Existing flake already compileraligned; installed oldprofile unchanged.","fullPipeline":["vet","golangci","AI"],"go":"1.26.8","golangci":"2.15.0","rulesConfigTaskHooksDepsPinsUnchanged":true},"verification":[{"artifacts":[".mtn-tmp/p19/quality-fix/task6-docs3-combined-build.log",".mtn-tmp/p19/quality-fix/task6-docs3-combined-tests.log",".mtn-tmp/p19/quality-fix/task6-docs3-combined-lint.log",".mtn-tmp/p19/quality-fix/task6-docs3-main-build.log",".mtn-tmp/p19/quality-fix/task6-docs3-main-tests.log",".mtn-tmp/p19/quality-fix/task6-docs3-main-lint.log"],"command":"Final exact-docs reviewed combined AND MAIN task build, task test(unit+race), FULL task lint(vet/golangci/AI)","exitCode":0,"result":"pass"},{"artifacts":[".mtn-tmp/p19/quality-fix/task6-docs3-observed-runtime-test-source-identity.log",".mtn-tmp/p19/quality-fix/task6-docs3-reviewed-source-identity.log"],"command":"Byte-identical runtime/test/Lua/Task/Make/Nix/package/config/hooks/deps/pins to independently whole-diff approved AND actual observed source2; final docs equal fresh review3","result":"pass"},{"artifact":".mtn-tmp/p19/quality-fix/task6-public-evidence.json","command":"Actual unchanged-source runtime proof: candidate/main make run/menu/SessionsIDs/currentmarker/ownedrenameonce/editedcancel/dismiss/exit0/everyhostabsence/exactfocus+stty+primaryscreen/baseline cleanup","result":"pass"},{"artifacts":[".mtn-tmp/p19/quality-fix/task6-docs3-whitespace.log",".mtn-tmp/p19/quality-fix/task6-lint-config-final.sha256",".mtn-tmp/p19/quality-fix/flake-after-docs3.patch"],"command":"Whitespace, final unchanged lint configuration and separate flake patch, target-base ancestry","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
