---
workflowKind: phase
workflowOrder: 9
workflowRequirements:
  - REQ-021
  - REQ-022
workflowDocuments:
  - ref: designs/rex-whichkey-tab-return.md
    hash: 'sha256:a616c835e37588e3c673b289849b40573db242126626bcdc1597a9c2a76607c9'
title: 'Phase09: Native popup tab return'
mtnCreateRetry:
  key: rex-whichkey-phase09-v1
  intent: bd2997b0827c959b84c20447a09ed5c26902a37669a7f796410c0b3da0ccf3be
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T15:22:57.739Z'
dateModified: '2026-11-02T15:23:34.206Z'
tags:
  - task
type: task
id: a6c6e0e8-3f79-4336-bd8d-7ba9935c6546
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy:
  - uid: '[[phases/rex-whichkey-08-popup-polish.md]]'
    reltype: FINISHTOSTART
---
# Phase09: Tab-return reliability

User correction: the popup stops showing when returning to its tab. Implement REQ-021 removing Zoom defaults and REQ-022 native tab-return/repaint reliability from approved designs/rex-whichkey-tab-return.md, preserving prior delivery and configured authenticated expiry. One bounded implementation task owns Go/Lua/default/docs interaction, with fresh mixed source review and coordinator-owned compilation plus actual owned-tab round-trip/navigation/Escape/cleanup check. Do not claim source/server capture proves remote graphical presentation. No tests/TDD/general suites, destructive/native selections, config reloads/settings edits, new client-selection policy, remote or Darwin trials. Same coordinator/session, serialized integration, no push or global cleanup.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE1OjIzOjA4LjQ0OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicGxhbiIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3RhYi1yZXR1cm4vcGhhc2UtcGxhbi8xIiwiaW50ZW50IjoiNWRhNjc5MzY1MzVmMDE0OGMxN2UzMWRkZmYzYzlkMDBkOThjNGRmNTdmNmRkMzIyOWU2NzYxZGIwYThmMjY5OSJ9 -->
### 2026-11-02T15:23:08.449Z · plan · unspecified

{"approvedBy":"explicit user scope; coordinator bounded decomposition","attempt":1,"coverage":{"REQ-021":["tasks/rex-whichkey-09-01-tab-return.md"],"REQ-022":["tasks/rex-whichkey-09-01-tab-return.md"]},"decision":"approved","decisionEvidence":"One literal ownership task covers Go/Lua default removal and foreground return lifecycle/docs with independent mixed source review and coordinator compilation/actual owned-tab runtime check; no source changes accepted yet.","documents":[{"hash":"sha256:a616c835e37588e3c673b289849b40573db242126626bcdc1597a9c2a76607c9","ref":"designs/rex-whichkey-tab-return.md"}],"eventId":"rex-whichkey/tab-return/phase-plan/1","kind":"plan","recordRef":"phases/rex-whichkey-09-tab-return.md","records":[{"ref":"tasks/rex-whichkey-09-01-tab-return.md","specHash":"sha256:dac01498d83ae0c4fb515581056d97877eac4a4497674cf55a5acf8e92760622"}],"schemaVersion":1,"specHash":"sha256:09e27ea5551f6f2f7417863af724d156049413d7452b7d608d111ee7ab61e7b8"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE1OjIzOjM0LjE5NFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcnVudGltZS90YWItcmV0dXJuIiwiaW50ZW50IjoiZjI4MjBhNGUyZjRkZGFkNDRhNGQ1ZmViYTI0MjNlZmE3NTNmZWViOGVmZDMwYWFhZGI2ZDZiMzc1NDU2MjcyMSJ9 -->
### 2026-11-02T15:23:34.194Z · runtime · unspecified

{"attempt":1,"continuationEvidence":"Same live coordinator process/runId/session/window resumes user-corrected Phase09; no replacement/restart. Display coordinator retained.","eventId":"rex-whichkey-p03-coordinator-a01/runtime/tab-return","handoff":true,"host":{"endpoint":"unix:///run/user/1000/rex/server.sock","hostId":"cf7ceda2-6cd7-468c-9360-3b7d68de6f91","instanceId":"0f6bfb2f-6739-42b6-a8cb-b226b8a8ec14"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","piSessionId":"rex-whichkey-p03-coordinator-a01","piSessionName":"coordinator","recordRef":"phases/rex-whichkey-09-tab-return.md","rexBlockId":"block:01a1927e-9c44-7693-8d92-6ec72cd57f67","rexSessionId":"session:01a1927e-9c44-7614-b3bc-0de3fd331a9d","rexSessionName":"rex-whichkey-main","rexWindowId":"window:01a1927e-9c45-7176-82af-1a472289ecc4","rexWindowName":"coordinator","role":"coordinator","runId":"rex-whichkey-p03-coordinator-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-team","specHash":"sha256:09e27ea5551f6f2f7417863af724d156049413d7452b7d608d111ee7ab61e7b8","stage":"launched","thinking":"high","transportDir":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p03-coordinator-a01","windowName":"coordinator","worktree":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->
