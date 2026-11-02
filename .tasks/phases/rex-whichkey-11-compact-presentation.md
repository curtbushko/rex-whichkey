---
workflowKind: phase
workflowOrder: 11
workflowRequirements:
  - REQ-024
  - REQ-025
  - REQ-026
workflowDocuments:
  - ref: designs/rex-whichkey-compact-presentation.md
    hash: 'sha256:10a72840e340a8bb53e744309fd518d564a9aac00a73364c3a554da06d05388f'
title: 'Phase11: Compact popup presentation'
mtnCreateRetry:
  key: rex-whichkey-phase11-v1
  intent: 9172ba9b4338588a4faac2117b870be97dcdd6d5416d8048f22e5fa1f0dc32eb
status: in-progress
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T17:33:39.158Z'
dateModified: '2026-11-02T18:41:40.924Z'
tags:
  - task
type: task
id: 52856da8-2e32-485e-8a7a-9c6f47fc27d5
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy:
  - uid: '[[phases/rex-whichkey-10-app-independent-popup.md]]'
    reltype: FINISHTOSTART
---
# Phase11: Compact popup presentation

Implement accepted REQ-024 status borders/conditional pagination and REQ-025 actual content-fitting layer with measured terminal geometry. REQ-026 hiding native controls is accepted but blocked by no documented per-popup Rex capability; no deferral is approved. One renderer producer, dependent geometry consumer, blocked chrome task and dependent docs/actual appearance assembly. Independent frozen reviews and combined compilation; coordinator bounded actual layer/root/submenu/cleanup check. No tests/general suites/settings/font changes/global chrome workaround, push or destructive cleanup. Same coordinator/session; phase/project remain incomplete while native chrome unresolved.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE3OjM0OjA0LjQ5NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicGxhbiIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3ByZXNlbnRhdGlvbi9waGFzZS1wbGFuLzEiLCJpbnRlbnQiOiI3MWYyNWFhZDBlYTgzZDhlNWY1NTZmYzkzNjIxYTBlODUxZTIwYzFlM2VhZDBmYzUzMWM1ZGY3NDNmMzA1MjI1In0 -->
### 2026-11-02T17:34:04.495Z · plan · unspecified

{"approvedBy":"explicit user scope; delegated coordinator decomposition","attempt":1,"coverage":{"REQ-024":["tasks/rex-whichkey-11-01-render-frame.md","tasks/rex-whichkey-11-04-appearance-docs.md"],"REQ-025":["tasks/rex-whichkey-11-02-popup-height.md","tasks/rex-whichkey-11-04-appearance-docs.md"],"REQ-026":["tasks/rex-whichkey-11-03-popup-chrome.md","tasks/rex-whichkey-11-04-appearance-docs.md"]},"decision":"approved","decisionEvidence":"Renderer -> geometry -> docs/actual gate; chrome task remains blocked independently. No full phase completion or user deferral inferred.","documents":[{"hash":"sha256:10a72840e340a8bb53e744309fd518d564a9aac00a73364c3a554da06d05388f","ref":"designs/rex-whichkey-compact-presentation.md"}],"eventId":"rex-whichkey/presentation/phase-plan/1","kind":"plan","recordRef":"phases/rex-whichkey-11-compact-presentation.md","records":[{"ref":"tasks/rex-whichkey-11-01-render-frame.md","specHash":"sha256:139ba9d08e8cf3dcb82eb28aff52d1aac0f80a4dc641f607de6a8fd9b63790e2"},{"ref":"tasks/rex-whichkey-11-02-popup-height.md","specHash":"sha256:defc1c3f5aa18077f971ad405b3ec196517f243388951f1f63503d1b89fa9afc"},{"ref":"tasks/rex-whichkey-11-03-popup-chrome.md","specHash":"sha256:00961722dcf9a5f2b742e8fcd91be861c00d825b1f0ed3de4a4f293486ca6250"},{"ref":"tasks/rex-whichkey-11-04-appearance-docs.md","specHash":"sha256:b83d836886acc660a6d2750d40df204f2604c2dcead59c6c0e40d7164ef60141"}],"schemaVersion":1,"specHash":"sha256:df29725a00be8007579da79d1dc9a7e882768879378b3f23f6b38009c7638d3b"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE3OjM0OjM4LjI0N1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcnVudGltZS9wcmVzZW50YXRpb24iLCJpbnRlbnQiOiI4ZGU0ODE5NmJkNzhhYTk5ODc3ZTBiZDEwMDhlZWJhMGMwZGRmNWNhMDgxZTA2YjQ3Mjg1Y2QxY2NhMmJjMmQzIn0 -->
### 2026-11-02T17:34:38.247Z · runtime · unspecified

{"attempt":1,"continuationEvidence":"Same live coordinator process/runId/session/window starts Phase11; no replacement/restart or identity rename. Chrome blocker does not authorize full completion.","eventId":"rex-whichkey-p03-coordinator-a01/runtime/presentation","handoff":true,"host":{"endpoint":"unix:///run/user/1000/rex/server.sock","hostId":"cf7ceda2-6cd7-468c-9360-3b7d68de6f91","instanceId":"0f6bfb2f-6739-42b6-a8cb-b226b8a8ec14"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","piSessionId":"rex-whichkey-p03-coordinator-a01","piSessionName":"coordinator","recordRef":"phases/rex-whichkey-11-compact-presentation.md","rexBlockId":"block:01a1927e-9c44-7693-8d92-6ec72cd57f67","rexSessionId":"session:01a1927e-9c44-7614-b3bc-0de3fd331a9d","rexSessionName":"rex-whichkey-main","rexWindowId":"window:01a1927e-9c45-7176-82af-1a472289ecc4","rexWindowName":"coordinator","role":"coordinator","runId":"rex-whichkey-p03-coordinator-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-team","specHash":"sha256:df29725a00be8007579da79d1dc9a7e882768879378b3f23f6b38009c7638d3b","stage":"launched","thinking":"high","transportDir":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p03-coordinator-a01","windowName":"coordinator","worktree":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE4OjQxOjQwLjkxNFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYmxvY2tlciIsInJldHJ5S2V5IjoicGhhc2VzL3JleC13aGljaGtleS0xMS1jb21wYWN0LXByZXNlbnRhdGlvbi5tZC9uYXRpdmUtY2hyb21lLWJsb2NrZXItY2hlY2twb2ludC8xIiwiaW50ZW50IjoiNzkyMTI5MjQ1ZGEwYTI0ZjNmMmRhZmFhMDlmZTJmODU1ODU4ZmQxOGI1YTc1ZGYwYTUyOGI1NTBlMzk4NjhhNCJ9 -->
### 2026-11-02T18:41:40.914Z · blocker · unspecified

{"actualVerification":[{"command":"Combined final Go/Taskfile and Nix compilation (tests/checks disabled), graph/schema validation and ancestry","result":"pass"},{"command":"Actual final main make run compact root/submenu/back measured bounds/rows/status frame/stability/Escape/owned cleanup","result":"pass"}],"affectedRefs":["tasks/rex-whichkey-11-03-popup-chrome.md","phases/rex-whichkey-11-compact-presentation.md","projects/rex-whichkey.md"],"artifactRefs":[".mtn-tmp/popup-presentation/runtime/final-make-run/result.json",".mtn-tmp/popup-presentation/partial-final/validate.json",".mtn-tmp/popup-presentation/chrome-blocker.json"],"attempt":1,"completedSourceRefs":["tasks/rex-whichkey-11-01-render-frame.md","tasks/rex-whichkey-11-02-popup-height.md","tasks/rex-whichkey-11-04-appearance-docs.md"],"deliveredRequirements":["REQ-024","REQ-025"],"eventId":"phases/rex-whichkey-11-compact-presentation.md/native-chrome-blocker-checkpoint/1","kind":"blocker","limitations":["No phase/project completion or user deferral. Native controls remain visible/unverified, not hidden by a global workaround.","One observed font/viewport context, no font/settings/config changes/reloads, project tests/TDD/general suites, native/destructive selections, configured Lua callback, remote or Darwin trials; server capture not remote GUI certificate."],"openRequirements":["REQ-026"],"reason":"Supported compact presentation work is reviewed/integrated/compiled and actual geometry/navigation/cleanup checked, but accepted REQ026 native pane/tab/window controls hiding remains blocked: installed API/options/actions provide no documented per-popup chrome control. No user deferral or full scope completion.","recordRef":"phases/rex-whichkey-11-compact-presentation.md","requiredRecovery":"Documented supported per-popup capability/approved Rex integration or explicit user deferral; resolve original chrome-task blocker and these parent blockers before further completion.","schemaVersion":1,"sourceCommit":"178d3f80c1d7efeab898e17ae07c883439570f66","sourceTree":"0b71f55d20052f408a1f9533dfe0ba4747b56969","specHash":"sha256:df29725a00be8007579da79d1dc9a7e882768879378b3f23f6b38009c7638d3b"}
<!-- mtn:end-entry:v1 -->
