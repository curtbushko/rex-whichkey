---
workflowKind: phase
workflowOrder: 12
workflowRequirements:
  - REQ-027
  - REQ-028
workflowDocuments:
  - ref: designs/rex-whichkey-first-frame-fit.md
    hash: 'sha256:e40a82729585e2a35de60126c4a3771c4c651d92e287f55c2808ff32d244e777'
title: 'Phase12: First-frame compact fit'
mtnCreateRetry:
  key: rex-whichkey-phase12-v1
  intent: fcdb50de1a43cba096b727ec28988b023c76b4038910e58f518fc0c41e5b7769
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T19:10:51.539Z'
dateModified: '2026-11-02T19:11:39.983Z'
tags:
  - task
type: task
id: c258ac4f-0f45-4211-880f-a4577d07c5f5
projects:
  - '[[projects/rex-whichkey.md]]'
---
# Phase12: First-frame compact sizing

User requests status/menu spacer rows and no visible oversized popup before fitting. Implement REQ027/028 in one coherent source/docs task with exact precreation geometry/terminal seed, frozen independent review and actual early-first-layer/spacing/adaptation/cleanup check. Depend on completed phase11 renderer/height/docs tasks, not its whole milestone: REQ026 native title/dock controls remains blocked separately. No fake hide/global settings, tests, font changes, destructive/native selections or remote/Darwin trials. Same coordinator/session. Phase12 can complete on actual supported scope; phase11 and project cannot while chrome remains unresolved.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE5OjExOjA0LjIxM1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicGxhbiIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L2ZpcnN0LWZyYW1lL3BoYXNlLXBsYW4vMSIsImludGVudCI6IjcxMWI2OTI3YTNkN2YyNjgzMTRlY2RjOGEzNWI4MmFjZThlY2ExMDc0ZGZhYjkyYmRmYWMwNGI4ZDlkOTRhY2MifQ -->
### 2026-11-02T19:11:04.213Z · plan · unspecified

{"approvedBy":"explicit user corrections; bounded coordinator decomposition","attempt":1,"coverage":{"REQ-027":["tasks/rex-whichkey-12-01-first-frame-fit.md"],"REQ-028":["tasks/rex-whichkey-12-01-first-frame-fit.md"]},"decision":"approved","decisionEvidence":"One coherent source/docs task depends on delivered phase11 renderer/height/docs, not whole blocked phase11 milestone. Native chrome stays separately blocked. Fresh review/combined compilation/actual early-first-layer gate before completion.","documents":[{"hash":"sha256:e40a82729585e2a35de60126c4a3771c4c651d92e287f55c2808ff32d244e777","ref":"designs/rex-whichkey-first-frame-fit.md"}],"eventId":"rex-whichkey/first-frame/phase-plan/1","kind":"plan","recordRef":"phases/rex-whichkey-12-first-frame-fit.md","records":[{"ref":"tasks/rex-whichkey-12-01-first-frame-fit.md","specHash":"sha256:e8f02cbea4ce9c6c7f70162261c1343b74f29d53b664b13e08fe1df2330e4f49"}],"schemaVersion":1,"specHash":"sha256:88d07bd9eb72c870ccd9e95bf4ff1bc4053eb474c8cce2daa7fbbc2138339243"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE5OjExOjM5Ljk3MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcnVudGltZS9maXJzdC1mcmFtZSIsImludGVudCI6ImRjMjYzY2YwN2MwN2E0YjU3ZmM2YTZlYjhiMzJiNThmODFiYjY2ODcwOGQwYWYyN2EzNDA2MjcyNzIxNzExYmUifQ -->
### 2026-11-02T19:11:39.971Z · runtime · unspecified

{"attempt":1,"continuationEvidence":"Same live coordinator process/runId/session/window schedules independent Phase12 despite unrelated phase11 chrome blocker; no restart or new coordinator.","eventId":"rex-whichkey-p03-coordinator-a01/runtime/first-frame","handoff":true,"host":{"endpoint":"unix:///run/user/1000/rex/server.sock","hostId":"cf7ceda2-6cd7-468c-9360-3b7d68de6f91","instanceId":"0f6bfb2f-6739-42b6-a8cb-b226b8a8ec14"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","piSessionId":"rex-whichkey-p03-coordinator-a01","piSessionName":"coordinator","recordRef":"phases/rex-whichkey-12-first-frame-fit.md","rexBlockId":"block:01a1927e-9c44-7693-8d92-6ec72cd57f67","rexSessionId":"session:01a1927e-9c44-7614-b3bc-0de3fd331a9d","rexSessionName":"rex-whichkey-main","rexWindowId":"window:01a1927e-9c45-7176-82af-1a472289ecc4","rexWindowName":"coordinator","role":"coordinator","runId":"rex-whichkey-p03-coordinator-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-team","specHash":"sha256:88d07bd9eb72c870ccd9e95bf4ff1bc4053eb474c8cce2daa7fbbc2138339243","stage":"launched","thinking":"high","transportDir":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p03-coordinator-a01","windowName":"coordinator","worktree":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->
