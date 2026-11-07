---
identity:
  kind: design
  designID: rex-whichkey-quality-gates
  uuid: f2f45239-b83b-4995-855f-d64a9410b6bf
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
    - REQ-026
    - REQ-028
    - REQ-029
  documents: []
structure: {}
lifecycle:
  state: approved
  sourceEventID: rex-whichkey/quality-gates/approval/1
  changedAt: '2026-11-07T15:58:51.233Z'
title: Mandatory tests and full lint quality gates
mtnCreateRetry:
  key: rex-whichkey-quality-gates-v1
  intent: e81881f22b44b525d4f80fba6329a10d79f22c7389d6174bfb8cc08375990a57
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T15:58:44.893Z'
dateModified: '2026-11-07T15:58:52.247Z'
tags:
  - task
type: task
id: 285a694e-f180-4ef7-ac78-030190954398
---
# Mandatory quality gates and Phase19 remediation

## Authority and correction

The user explicitly corrected the execution policy: "You should always run the tests and linting after a change", then authorized "fix the problems". This current instruction supersedes every carried-forward no-unit/race/TDD/Lua/parser/FIFO-test-harness clause for this repository. Do not infer or insist that the user approved skipping current tests. Preserve the earlier execution record as history, not present policy.

Build, automated unit/race suites and the FULL existing lint pipeline must run after changes; do not present go vet, source review or actual make run as a replacement. Phase19 is reopened until the verification problems are fixed and the final combined/main build/test/lint gates pass. No lint/config/hook/test weakening, suppressions, skipped tests, deleted assertions, dummy compatibility objects or fake passing receipts. The new directive authorizes existing automated fixtures/mocks/sandbox harnesses and TDD; it does not authorize live remote trials, user configuration/profile edits, install/activation/Rex restart/reload, dependency/pin updates, benchmarks, global cleanup or unrelated source rewrites.

## Observed baseline and accepted repair scope

At e867a3b, normal task build passes. task test:unit and task test:race both fail four groups: TestConfiguredRenameDispatch, TestCLIDelegation, TestLaunchOriginPreflightDiagnostics and TestRexActionContexts. Errors include archived Phase15 Lua fixture GUI-target/context and helper/layout source-discovery mismatches. No race detector diagnostic was seen; a failing race-suite assertion is not itself a data race. task lint fails with53 diagnostics, so lint:ai has not been reached by the aggregate target. Logs .mtn-tmp/p19/post-delivery-gates are actual evidence; do not assume inherited versus introduced failures without tracing.

Repair tests/fixtures faithfully: retain security/identity/context/rename/cancellation assertions and scenarios, correct stale isolated harness setup/source loading as required, never relax the archived implementation's guards or restore that control plane into production. Preserve current direct Lua-launch/Go-menu behavior. Add targeted current-production regression coverage where the change requires it; no dormant test-only production hooks.

Repair all current lint diagnostics by fixing code, particularly complexity/nesting, error/context handling, shadowing, constants and comments. Run the existing AI lint stage too after golangci passes; no omitted pipeline stage can be credited. Refactors must preserve terminal parsing/Unicode/paste/escape/mode restoration, palette/menu/defaults, bounded async metadata/retained target/readiness/rename edits, JSON validation and exact original/own/GUI/destination/generation/no-retry lifetime behavior. No Cobra/Viper/hexagonal/framework/dependency rewrite. Lua controls native float sizing/launch; Go handles menus/palette/actions. User1a native chrome defaults/customization deferral and1b detached host through dispatch/post-exit retirement remain settled and unchanged.

## Serial implementation and delivery

Keep completed Task01-04 source/review/runtime histories. Their earlier no-tests delivery is historical, not quality acceptance under this correction. Add Task05 faithful test repair and Task06 complete lint/code-quality repair, native dependencies05 then06, and final real make run smoke on reviewed combined/main source. Task05 must get build/unit/race green, run/report full lint with the observed baseline only until Task06 removes it; this staging is not a final lint waiver. Task06 and the phase require task build, task test and task lint fully green, plus architecture gate if an existing config requires it. Rerun tests/lint after every source change. Never alter Taskfile quality commands, ignored .golangci.yml, .go-ai-lint.yml, hooks, assertions or deps to obtain green.

The same sole coordinator owns all .tasks/Git integration. Isolated committed-base leaves, fresh revision-bound independent review, speculative combined checks before target advance, real normal app behavior verification and eligible-successful-resource cleanup. Preserve failed recovery resources and separate dirty flake.nix Darwin fix. One builder/reviewer/heavy gate/integration,14CPU/6GiB shared/Go5GiB, reuse existing pinned Go1.26.8/lint2.15 tools; do not enter nix develop or trigger shellHook. Unrelated old phases remain untouched.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDE1OjU4OjUwLjA5NFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9xdWFsaXR5LWdhdGVzL2FwcHJvdmFsLzEiLCJpbnRlbnQiOiJhOTUyNjEwNWI1MTI4NzQ1ODU5MDUwMDY1YjI1NGU5ZjBlNDM1MzllM2IzZWMxYjUzZDhkYTU4NWY1MGU5ZTBhIn0 -->
### 2026-11-07T15:58:50.094Z · approval · unspecified

{"approvalScope":"Mandatory test/full lint correction and bounded repair of actual failures; supersedes all earlier test prohibitions, preserves architecture/safety/native chrome and lifetime decisions; no new deployment/profile/config authority","approvedBy":"user","attempt":1,"decision":"approved","decisionEvidence":"You should always run the tests and linting after a change; fix the problems","decisionID":"REXWHICHKEY-20261107-105542","eventID":"rex-whichkey/quality-gates/approval/1","kind":"approval","recordRef":"designs/rex-whichkey-quality-gates.md","schemaVersion":1,"specHash":"sha256:f3120762f7075c7c767e5f39552ba9c2b38b2e213d33d2472d07f0bc0c1cbdb2"}
<!-- mtn:end-entry:v1 -->
