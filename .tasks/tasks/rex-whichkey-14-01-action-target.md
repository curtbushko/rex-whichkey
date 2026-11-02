---
workflowKind: task
workflowLanguage: go
workflowOrder: 1
workflowRequirements:
  - REQ-031
workflowDocuments:
  - ref: designs/rex-whichkey-action-audit.md
    hash: 'sha256:2d2d664eb4bbc5828b2f0d6d4a98346becfbbd03d90e8f9abc5a9874edc8b7a8'
workflowOwnership:
  paths:
    - native.go
    - popup.go
    - native_targets.go
    - README.md
    - docs/launch-contract.md
    - docs/acceptance.md
  sharedPaths: []
  resources: []
  interfaces:
    - native-action-audit-1
title: Repair ambiguous native action target
mtnCreateRetry:
  key: rex-whichkey-phase14-task1-v1
  intent: b3371edada4d6463a137a325dec2cdfa91b7254d1aa0bea7e9b86abd4bcc637d
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T20:16:00.654Z'
dateModified: '2026-11-02T20:33:01.328Z'
tags:
  - task
type: task
id: 4912e807-2cc3-4139-bea0-a60e8d5006f1
projects:
  - '[[phases/rex-whichkey-14-action-audit.md]]'
completedDate: '2026-11-02'
---
# Repair shared native action targeting with explicit late app choice

Implement REQ031 per designs/rex-whichkey-action-audit.md. User's latest request is verify every action incl broken split down. Reproduction proves mapping direction=down is correct; bare selection fails because two GUI apps. Do not change correct pane args or require manual --client. Add safe explicit app target picker at action time when several eligible apps, using existing canonical menu/popup lifecycle, no new service/protocol/daemon or live sizing loop. Basic menu/Escape stay independent of graphical clients.

Strict explicit selector stays strict (missing/stale must fail, not chooser fallback). Single target binds exact ID/generation. Multiple targets retained in Go-owned snapshot, display safe labels+full IDs as needed, chooser Completion validated against retained config; only selected app identity comes from choice, not replacement original action/args/confirmation. New target prompt after original popup cleanup is permitted (purposeful action routing, NOT submenu sizing). Cancel -> exit0/no action. All UI/files/leases cleaned before original native dispatch. Zero app safely refuses native GUI action.

Critical lifecycle: runPopupWithLifetime currently wraps finish in a short operation context. Do NOT run an indefinite chooser in that context. Capture first retained selection during finish, then subsequent picker with native unbounded interaction and fresh short dispatch budget afterward. Keep signal cancellation across whole workflow including gaps/lookups/dispatch (outer native signal context or equally robust ownership); no leaking child on signal. After choice revalidate exact captured session/block/app generation and selected action metadata, then exactly one dispatch. Never infer action target from terminal owner/focus/first/labels; no retries on uncertain activation. Preserve configured original VM route and lifetime/auth, native FIFO/parent lease, popup initial fit/spacers/fixed rectangle, validation/confirmation, cleanup/error behavior.

Allowed source: native.go, popup.go, native_targets.go (new private helper), README.md, docs/launch-contract.md, docs/acceptance.md. Refactor eligible GUI enumeration without weakening existing popupOriginConnection validation (configured path must retain original semantics). Existing generic public Rex APIs/Mapping schema not owned. Keep code straightforward and no duplication of VM/schema/action registry.

Leaves source/format/whitespace checks only. NO tests/TDD/general/unit/race, builds/heavy/compiler lint, native/GUI/RPC/action/config/font/platform experiments or installs. Coordinator owns compilation and runtime gates. Stage retired code only in ignored .trash if retirement needed. Source-only validation array of actual passes; skipped checks limitations.

Acceptance: normal native display/cancel never preflights GUI target; exact optional selector preserved; multiple apps produce explicit safe choice after action selected; canceled target no action; same original action+args+confirmation retained; chosen ID/generation/origin rechecked; no guessing/caching/retry; all owned cleanup before one action; indefinite choice/whole signal safety; docs no more blanket bare-action failure guidance and no fake complete all-action/UI certification. Actual coordinator gate bare s -> chooser -> exact work app -> real vertical split below origin, chooser cancellation no new pane, strict invalid explicit client refuses, resource/file cleanup and fixed menu navigation. Split direction currently correct; no gratuitous mapping changes. Native chrome stays blocked; source repair does NOT prove seven GUI-only behaviors.

Full protocol-v1 build result to MTN_REX_RESULT_FILE; exact base/head/tree/attempt/task/spec, changes and criteria evidence, validation actual array, limitations. Commit only source with git -c core.hooksPath=/dev/null. No collection mutation/integration/subagents. Visible final concise pointer only.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDIwOjMzOjAxLjMxNloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxNC10MDEtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiYWQxZTRhOTVhMTc1ZTVmOTE2NzQzZDM1ODhmNWFkNzBkNjA2YmEyNzQ1ZjczMWY3M2FmZTE3ODdjY2JhMTQ1MyJ9 -->
### 2026-11-02T20:33:01.316Z · summary · unspecified

{"acceptedScope":"Repair ambiguous native action target","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p14-t01-reviewer-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/action-audit/rex-whichkey-p14-t01-builder-a01/integration/integration.json",".mtn-tmp/action-audit/runtime/rex-whichkey-p14-t01-builder-a01/result.json"],"attempt":1,"eventId":"rex-whichkey-p14-t01-builder-a01/summary/1","files":["README.md","docs/acceptance.md","docs/launch-contract.md","native.go","native_targets.go","popup.go"],"kind":"summary","limitations":["Late app choice/bare split-down/cancel/strict selector/owned cleanup checked on disposable resources only, not all25 verified. GUI-only7 action outcomes and native title/dock remain unresolved; no full phase/project completion.","No project tests/TDD, user-resource mutation, configured Lua callbacks, remote/Darwin or config/settings/fonts changes/reloads; fresh same-family review."],"recordRef":"tasks/rex-whichkey-14-01-action-target.md","requirements":["REQ-031"],"schemaVersion":1,"specHash":"sha256:96e3d8f88d25c2635ed78d7ca7de32eff427d5b07798c755451b9fdce4022e1e","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"75a849cff9628d720ba74715491a32386afbfffa","targetTree":"67eae0de5c8413c31b6e6293f5730b248af1c2d4","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"},{"command":"task build; actual late app picker/bare split-down below origin, target cancel/no action, strict missing selector, exact owned cleanup","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
