---
workflowKind: phase
workflowOrder: 10
workflowRequirements:
  - REQ-023
workflowDocuments:
  - ref: designs/rex-whichkey-app-independent-popup.md
    hash: 'sha256:7215d04e0b71c6cc6a8a593372ef03cec57ae6aeb16907a81f1f9407b5781b4c'
title: 'Phase10: App-independent popup display'
mtnCreateRetry:
  key: rex-whichkey-phase10-v1
  intent: 5b41ac2d723b7c4cd6210942fd8f9ae2e3acfddc83de08ff7bb97db0e70fe29a
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T16:14:34.788Z'
dateModified: '2026-11-02T16:15:13.123Z'
tags:
  - task
type: task
id: 708c5036-24cb-422e-b6f5-20952329a818
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy:
  - uid: '[[phases/rex-whichkey-09-tab-return.md]]'
    reltype: FINISHTOSTART
---
# Phase10: App-independent native popup startup

Deliver REQ-023 from approved designs/rex-whichkey-app-independent-popup.md: remove app-target validation from basic menu display, retain it only where executable native selection requires a target, and preserve configured Lua binding/security. User explicitly requests this correction. One source/docs task, independent fresh review, coordinator exact combined compilation and bounded real normal/nonexistent-selector display/dismissal plus nonexistent-target refusal after cleanup. No broad tests, destructive action dispatch, configuration changes/reloads or arbitrary client selection. Existing coordinator/session/runtime identity and completed phases02–09 remain intact; no push/global cleanup.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE2OjE0OjQ1LjYxNVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicGxhbiIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L2FwcC1pbmRlcGVuZGVudC9waGFzZS1wbGFuLzEiLCJpbnRlbnQiOiJlZGQ0OTNkNzNlNDkzMDc1YjgyMTg4NjQ4MzVkY2YxMzY4MDEwYWJmYzI1ZmU0NDYxYzQ5MTA4YjNmNTFiOTg4In0 -->
### 2026-11-02T16:14:45.615Z · plan · unspecified

{"approvedBy":"explicit user correction; bounded coordinator decomposition","attempt":1,"coverage":{"REQ-023":["tasks/rex-whichkey-10-01-app-independent-popup.md"]},"decision":"approved","decisionEvidence":"One source/docs task with independent security/lifecycle review and exact combined compilation/actual native UI gate. Prior phases remain complete; no source pass inferred from approval.","documents":[{"hash":"sha256:7215d04e0b71c6cc6a8a593372ef03cec57ae6aeb16907a81f1f9407b5781b4c","ref":"designs/rex-whichkey-app-independent-popup.md"}],"eventId":"rex-whichkey/app-independent/phase-plan/1","kind":"plan","recordRef":"phases/rex-whichkey-10-app-independent-popup.md","records":[{"ref":"tasks/rex-whichkey-10-01-app-independent-popup.md","specHash":"sha256:1b007c064cc0f5f60819251a8d259450baa76641ac643f286ba116a11683046d"}],"schemaVersion":1,"specHash":"sha256:6bdf93feca5f655089585868e7bdd56b5177a4689b6dcbc76f7ac5d7bfc2d16d"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE2OjE1OjEzLjExMloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcnVudGltZS9hcHAtaW5kZXBlbmRlbnQiLCJpbnRlbnQiOiI4OTczMzJmNDE4Yzk0OTI0OTRmN2E5YjcyZThiNzJhZDhjNWQ4YjgyNmQ0ZGFjNjY2ZjQ4YmFiZWRmODM3ZGJjIn0 -->
### 2026-11-02T16:15:13.112Z · runtime · unspecified

{"attempt":1,"continuationEvidence":"Same live coordinator process/runId/session/window handles explicit Phase10 fix; no replacement/restart or identity rename.","eventId":"rex-whichkey-p03-coordinator-a01/runtime/app-independent","handoff":true,"host":{"endpoint":"unix:///run/user/1000/rex/server.sock","hostId":"cf7ceda2-6cd7-468c-9360-3b7d68de6f91","instanceId":"0f6bfb2f-6739-42b6-a8cb-b226b8a8ec14"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","piSessionId":"rex-whichkey-p03-coordinator-a01","piSessionName":"coordinator","recordRef":"phases/rex-whichkey-10-app-independent-popup.md","rexBlockId":"block:01a1927e-9c44-7693-8d92-6ec72cd57f67","rexSessionId":"session:01a1927e-9c44-7614-b3bc-0de3fd331a9d","rexSessionName":"rex-whichkey-main","rexWindowId":"window:01a1927e-9c45-7176-82af-1a472289ecc4","rexWindowName":"coordinator","role":"coordinator","runId":"rex-whichkey-p03-coordinator-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-team","specHash":"sha256:6bdf93feca5f655089585868e7bdd56b5177a4689b6dcbc76f7ac5d7bfc2d16d","stage":"launched","thinking":"high","transportDir":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p03-coordinator-a01","windowName":"coordinator","worktree":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->
