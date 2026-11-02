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
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T17:33:39.158Z'
dateModified: '2026-11-02T17:34:38.259Z'
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
