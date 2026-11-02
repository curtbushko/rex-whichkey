---
workflowKind: task
workflowLanguage: docs
workflowOrder: 3
workflowRequirements:
  - REQ-030
workflowDocuments:
  - ref: designs/rex-whichkey-action-audit.md
    hash: 'sha256:2d2d664eb4bbc5828b2f0d6d4a98346becfbbd03d90e8f9abc5a9874edc8b7a8'
  - ref: designs/rex-whichkey-selected-window-splits.md
    hash: 'sha256:1d5e6297352b97b84e3ba7730e586130b4103ddb4ad32ed03d35b6ac2ad6e4f2'
workflowOwnership:
  paths:
    - docs/gui-action-verification.md
  sharedPaths: []
  resources: []
  interfaces:
    - native-action-audit-3
title: Verify remaining native UI interactions
mtnCreateRetry:
  key: rex-whichkey-phase14-task3-v1
  intent: 3b03b5f0d4d2b1e22b21b2a2334d54f215d29bb407ce36a6c17911328047b56e
status: in-progress
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T20:16:08.762Z'
dateModified: '2026-11-02T22:58:30.919Z'
tags:
  - task
type: task
id: b8310b3e-9e3b-4e4a-8cdc-a94e48f9b3eb
projects:
  - '[[phases/rex-whichkey-14-action-audit.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-14-02-resource-audit.md]]'
    reltype: FINISHTOSTART
---
# Finish remaining native UI interaction verification

REQ030 all25 still required. Latest API-driven research corrects blanket observer/screenshot blocker: native app reporting fresh target geometry, ordered resources and focus/input controls can verify functional behavior. Source5 explicit-ID numbered intent and Source6 newtab/resource transport repairs are separate reviewed code scopes; Task14-02 refreshed 25-row matrix documents actual evidence. Derive actual verified count from final-source gates, anticipated21 but no assumedpass.

Outstanding interactions requiring genuine evidence: ws native session picker, wr native session rename including successful own-session rename, tr native tab rename including own-tab rename, c native paletteopen/close. Source-only, registration, performed_by acknowledgments, unrelated popupFocus transitions or unlabelled multi-app focus bytes are insufficient. Existing palette direct+menu toggles/focus traces are evidence of invocation but not yet exact UI state. Existing rename attempt typed audit/ESC into terminal logger, did NOT change label; pane.send_key documented terminalinput, NOT modal keystrokeAPI. No invented textargs/getters/screenshot/actions, no clipboard/debug export/userinit/client-script takeover, nativeOS or remoteMac trials/config/fonts/reloads without new explicit authority. All experiments confined to exact own resources; never blind input/split probes when native context could leave fixtures. Investigate documented API path; if genuinely absent record specific four-case capability gap/recovery, not blanket19 cannotbeverified or requirementwaiver.

Depends on refreshed resource/audit document Task14-02. Reserved source ownership docs/gui-action-verification.md only for future final interaction proof; coor owns actual bounded experiments, leaves no builds/tests/nativeactions. Do not complete task/all25/phase/project while any bindingoutcome unresolved. Independent chromeREQ026 remains blocked separately. Existing accepted scope/history preserved and explicit latercorrections take precedence.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDIwOjUyOjU2LjAyMFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYmxvY2tlciIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L2FjdGlvbi1hdWRpdC9uYXRpdmUtZ3VpLW9ic2VydmFiaWxpdHktYmxvY2tlci8xIiwiaW50ZW50IjoiODZjYTc0ZTE0ZGJlOGQxNGE0ZTcwY2Y4N2RkY2U2NTc1NjBmMDRkNGE5YWMzNmRlOWUzOGY1Y2Y2ZTQ0ZmZhOCJ9 -->
### 2026-11-02T20:52:56.020Z · blocker · unspecified

{"affectedRefs":["tasks/rex-whichkey-14-03-gui-action-verification.md","phases/rex-whichkey-14-action-audit.md","projects/rex-whichkey.md"],"artifactRefs":[".mtn-tmp/action-audit/api-list.json",".mtn-tmp/action-audit/client-inspect-api.json",".mtn-tmp/action-audit/client-key-api.json",".mtn-tmp/action-audit/tab-navigation-before.json",".mtn-tmp/action-audit/tab-navigation-after.json",".mtn-tmp/action-audit/tab-navigation-probe.json",".mtn-tmp/action-audit/resource-runtime/integrated-selected-source/result.json",".mtn-tmp/action-audit/observability-correction.md"],"attempt":1,"eventId":"rex-whichkey/action-audit/native-gui-observability-blocker/1","kind":"blocker","reason":"Nineteen native GUI outcomes remain behaviorally unverified. Installed client.inspect exposes metadata/actions/attached IDs only; client.key drives a block PTY/Lua bindings, not native app dialogs. A real tab.goto probe returned null without changing shared server active_window_id, so that field cannot certify app tab state. Native tn acknowledged with exit0 but produced no new server window/session, leaving actual app behavior unknown. No safe documented native UI observation/input mechanism for tab/session selected state, new-tab/picker/rename/sidebar/palette. Do not count schema/receipt/server survival as visible behavior or claim every action works.","recordRef":"tasks/rex-whichkey-14-03-gui-action-verification.md","recovery":"Provide approved genuine native-app observation/control capability or explicit observed user evidence for all19 remaining binding outcomes; never invent screenshot/actions/args or waive request silently. Six resource effects and current-window splits can complete independently.","schemaVersion":1,"specHash":"sha256:27c75d06e39f01153b986fa112aefc430a6526bd6b5ada8dd2395479c5d7f74b"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDIxOjI2OjE4LjY3N1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L2FjdGlvbi1hdWRpdC9hcGktb2JzZXJ2YXRpb24tcmVzZWFyY2gvMSIsImludGVudCI6ImRhOWU1N2QxN2U0YTI4NDEwNGRiNDU3Y2NlZDA5OTcyODQ3Yjk0OThjOTZhMzkyNDAxYjExNDMxYWI3ODU1NjgifQ -->
### 2026-11-02T21:26:18.677Z · historical · unspecified

{"artifactRefs":[".mtn-tmp/action-audit/api-capability-investigation/findings.md",".mtn-tmp/action-audit/api-capability-investigation/api-list.json",".mtn-tmp/action-audit/api-capability-investigation/tab-routing-comparison/result.json",".mtn-tmp/action-audit/api-capability-investigation/selected-tab-probe/control-probe-response.json"],"attempt":1,"content":"User correctly challenges blanket need for native GUI access. Full API/Lua/client scripting/data/event contracts investigated. Documented pane.send_key plus owned receiver provides positive-control observable current-focus behavior without screenshot. Real menu t1 and direct native tab.goto variants acknowledged but subsequent markers reached none of owned receivers, unlike explicit session.select control. This is unresolved routing/focus/scope discrepancy, NOT a pass and NOT proof19 actions impossible. Stop further blind input/split probes until actual selected-context semantics established. Earlier six resource effects remain; all25 request incomplete. No config/script takeover/clipboard/remote OS changes.","eventId":"rex-whichkey/action-audit/api-observation-research/1","gateEligible":false,"kind":"historical","provenance":"Coordinator full installed contracts/published references and actual owned input-receiver trials, including direct positive controls and failed expectation; no GUI state or source fix inferred","recordRef":"tasks/rex-whichkey-14-03-gui-action-verification.md","schemaVersion":1,"specHash":"sha256:27c75d06e39f01153b986fa112aefc430a6526bd6b5ada8dd2395479c5d7f74b"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDIxOjI2OjIxLjM3OFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVzdW1lIiwicmV0cnlLZXkiOiJyZXgtd2hpY2hrZXkvYWN0aW9uLWF1ZGl0L2FwaS1vYnNlcnZhdGlvbi1yZXNlYXJjaC1yZXN1bWUvMSIsImludGVudCI6ImE3MDFmOWZhZDU1ZDBjMzlhZjIwYWYyMjExMGY2ZmFhZjQ4YmZhODI4MWI3YTE5NTMzNjJjMDI0MjZjZTEyNDAifQ -->
### 2026-11-02T21:26:21.378Z · resume · unspecified

{"approvedBy":"User explicitly challenges API limitation; coordinator corrects overbroad blocker, no requirement waiver","artifactRefs":[".mtn-tmp/action-audit/api-capability-investigation/findings.md",".mtn-tmp/action-audit/api-capability-investigation/research-event.json"],"attempt":1,"blockerEventId":"rex-whichkey/action-audit/native-gui-observability-blocker/1","decisionEvidence":"Remove blanket19-require-GUI-access blocker: indirect functional probes are possible. Resume investigation, NOT verified task/all25 completion. Named tab/focus discrepancy and native modal outcomes still unresolved; no new action passes claimed.","eventId":"rex-whichkey/action-audit/api-observation-research-resume/1","kind":"resume","recordRef":"tasks/rex-whichkey-14-03-gui-action-verification.md","schemaVersion":1,"specHash":"sha256:27c75d06e39f01153b986fa112aefc430a6526bd6b5ada8dd2395479c5d7f74b"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDIyOjU4OjMwLjkwOVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L2FjdGlvbi1hdWRpdC9hcGkyMS1wYXJ0aWFsLXVpLWNoZWNrcG9pbnQvMSIsImludGVudCI6ImVlNTI4NTYxMWJiMDk1ZWFmNTM3MWYxNGE1NmU2YjFkMDgyYzIwNjE4NzI5ZDRjODRkYzg5Y2JjOGYzZTk0OTQifQ -->
### 2026-11-02T22:58:30.909Z · historical · unspecified

{"artifactRefs":["docs/actions.md",".mtn-tmp/action-audit/api-capability-investigation/current-action-matrix.json",".mtn-tmp/action-audit/api-capability-investigation/independent-diagnosis.md",".mtn-tmp/action-audit/resource-runtime/rex-whichkey-p14-t06-builder-a02/result.json",".mtn-tmp/action-audit/api-capability-investigation/final-cleanup/remaining-sessions.json",".mtn-tmp/action-audit/api-capability-investigation/final/validate.json",".mtn-tmp/popup-fixed-navigation/runtime/final-api21-make-run/result.json",".mtn-tmp/action-audit/selected-runtime/final-api21/result.json"],"attempt":1,"content":"Continued API verification now21/25 binding effects verified. Native selected-window s/v retained; numbered t1..t9 repaired through frozen exact-ID intent, actual9 fresh-target rendering checks. NewTab transport-only a01 FAILED unchanged7-effect gate, preserved; a02 explicit typed default-window creation+exact activation passed all7 resource effects, no uncertain retry/rollback. th/tl interior,wh/wl,sidebar real-menu API effects verified. Earlier key probes are NOT reliable selected-tab oracle; passive fresh app rendering calibrated positive. Native last/previous work; unsupported wrap/constant-row assertions corrected with raw failures retained. Four ws/wr/tr/c interactions remain individually unverified, not blanket19 require-screenshots blocker or fullall25 pass. Task03 stays inprogress. Native chromeREQ026 separatelyblocked. Finalmain Go/Taskfile/Nix build, actual make-run fixed geometry and selected-B vslauncherA splits freshlychecked. All disposable test sessions reconciled with IDs; original4sessions retained; artifacts/worktrees preserved.","deliveredRepairs":["REQ-031","REQ-032","REQ-030 numbered/NewTab subscopes"],"eventId":"rex-whichkey/action-audit/api21-partial-ui-checkpoint/1","gateEligible":false,"kind":"historical","limitations":["API acknowledgments/focus bytes do not certify native picker/rename/palette state; terminal key API typed audit into ownTTY, not native dialog; no fake completed rename or undocumented args.","No general/unit/race/TDD suites, user resources destroyed, config/fonts/settings/reload/clipboard/remoteOS or configured callback trials. No new API/framework/daemon or push."],"openRequirements":["REQ-030 remaining native interactions","REQ-026 chrome"],"provenance":"Coordinator actual API/menu/resource/rendering/focus evidence plus independent source reviews, exact source and final combined compiler gates; no GUI-pixel/native modal proof inferred","recordRef":"tasks/rex-whichkey-14-03-gui-action-verification.md","schemaVersion":1,"sourceCommit":"32f6a160c4cca613b17b9a97e2c922e19ed31c08","sourceTree":"008b8072783d2e5f617c14694d25adc432178638","specHash":"sha256:69cfe1efeec139ac092682446933a44ed1f42d2d3ce9dfa0d3d8db8cbce7ee5f","unverifiedBindings":["ws","wr","tr","c"],"verifiedBindings":21}
<!-- mtn:end-entry:v1 -->
