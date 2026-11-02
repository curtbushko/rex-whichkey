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
title: Verify seven native GUI-only action outcomes
mtnCreateRetry:
  key: rex-whichkey-phase14-task3-v1
  intent: 3b03b5f0d4d2b1e22b21b2a2334d54f215d29bb407ce36a6c17911328047b56e
status: in-progress
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T20:16:08.762Z'
dateModified: '2026-11-02T21:26:22.303Z'
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
# Verify nineteen native GUI-local action outcomes

REQ030 remains incomplete until all25 verified. After real resource probes, six resource effects can be observed (v/s/x/tx/wn/wx);19 remaining bindings need genuine GUI observability: tn, th/tl/tr/t1..t9, ws/wh/wl/wr, e/c. This corrects unverified initial18-resource/7-UI planning assumption, not scope waiver. Evidence .mtn-tmp/action-audit/observability-correction.md: client.tab.goto null result leaves server active_window_id unchanged; server field is not chosen app's selected tab. Native tn exited0 but no server window or session appeared; may be native UI/client-local tab, not certified working or conclusively broken.

Need actual selected tab/session, sidebar/palette, picker/rename modal and completed rename/new-tab behavior. client.inspect exposes metadata/actions/attached IDs only; terminal send-key drives block PTY, not native app dialogs. Investigate documented capability before pass. If none, BLOCKED recovery approved genuine GUI observation/control or explicit observed user evidence, not invented screenshot/action/args, not stale server focus/palette state or RPC acknowledgment. Do not open unclosable prompts or pretend user deferred requirement. No general tests/settings/font/config/reload/remote/Darwin trials, no destructive user resources. Depends on resource documentation; reserved ownership docs/gui-action-verification.md for future real evidence. Preserve independent native chrome blocker. Source repair can complete while this and entire phase/project stay incomplete.

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
