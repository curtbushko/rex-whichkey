---
workflowKind: phase
workflowOrder: 13
workflowRequirements:
  - REQ-029
workflowDocuments:
  - ref: designs/rex-whichkey-fixed-navigation.md
    hash: 'sha256:7fbb1596666f9acf0ddb56ecf00ff29c09c8ae26781577268138e6a4ca05177e'
title: 'Phase13: Stable popup navigation'
mtnCreateRetry:
  key: rex-whichkey-phase13-v1
  intent: 3b1be4c5362d98bd3d22378a8033e292fa0aa1aaf08c07f8183aa2eb79ae66f7
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T19:37:53.010Z'
dateModified: '2026-11-02T19:38:38.034Z'
tags:
  - task
type: task
id: e8ef574b-9d5e-4a88-a73b-66ae42d1e4be
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy:
  - uid: '[[phases/rex-whichkey-12-first-frame-fit.md]]'
    reltype: FINISHTOSTART
---
# Phase13: Stable popup navigation

Explicit latest user correction REQ029: keep initial popup rectangle stable across submenu entries, use existing pagination rather than live resize. Preserve completed phase12 first-frame/spacers and native/configured safety; retire unnecessary live geometry IPC. One source/docs task depends on phase12, fresh review/combined compilation/actual identical-bound root/submenu/paging/back/idle/Escape check. No tests/settings/fonts changes/native selections or remote/Darwin trials. REQ026 native title/dock capability remains blocked; no full project completion or push.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE5OjM4OjA0Ljc3OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicGxhbiIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L2ZpeGVkLW5hdmlnYXRpb24vcGhhc2UtcGxhbi8xIiwiaW50ZW50IjoiYjVhMzZhMGZmZjAyY2Y3MzBlNDMzMzNjN2I3YzI1YzBhMzc1MTc0NzEyOGJlMmVlYWM1YWI1MjcyOGFmMDg3YyJ9 -->
### 2026-11-02T19:38:04.779Z · plan · unspecified

{"approvedBy":"explicit user correction; delegated decomposition","attempt":1,"coverage":{"REQ-029":["tasks/rex-whichkey-13-01-fixed-navigation.md"]},"decision":"approved","decisionEvidence":"One coherent source/docs task after completed first-frame scope; frozen review and actual identical-bound navigation gate before completion.","documents":[{"hash":"sha256:7fbb1596666f9acf0ddb56ecf00ff29c09c8ae26781577268138e6a4ca05177e","ref":"designs/rex-whichkey-fixed-navigation.md"}],"eventId":"rex-whichkey/fixed-navigation/phase-plan/1","kind":"plan","recordRef":"phases/rex-whichkey-13-fixed-navigation.md","records":[{"ref":"tasks/rex-whichkey-13-01-fixed-navigation.md","specHash":"sha256:ef8bc1c6f61dd447a38797e0d1ef78efd2463709666295614893ace2110f8173"}],"schemaVersion":1,"specHash":"sha256:2660682d6a3d7f52a00745012de6b45549c888675e0469841b8947991fc9acc2"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE5OjM4OjM4LjAyMloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcnVudGltZS9maXhlZC1uYXZpZ2F0aW9uIiwiaW50ZW50IjoiN2JmZDBlNzY0MGJkZmVlZGI5ZGYzZTM3ZTg5NDk3NDdlNTI5NWU5Y2U0Nzg3ZGJhNWIxMGU2MzNlOWI0Nzg1NyJ9 -->
### 2026-11-02T19:38:38.022Z · runtime · unspecified

{"attempt":1,"continuationEvidence":"Same live coordinator process/runId/session/window implements explicit fixed-navigation correction. No restart/new coordinator; earlier first-frame delivery preserved.","eventId":"rex-whichkey-p03-coordinator-a01/runtime/fixed-navigation","handoff":true,"host":{"endpoint":"unix:///run/user/1000/rex/server.sock","hostId":"cf7ceda2-6cd7-468c-9360-3b7d68de6f91","instanceId":"0f6bfb2f-6739-42b6-a8cb-b226b8a8ec14"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","piSessionId":"rex-whichkey-p03-coordinator-a01","piSessionName":"coordinator","recordRef":"phases/rex-whichkey-13-fixed-navigation.md","rexBlockId":"block:01a1927e-9c44-7693-8d92-6ec72cd57f67","rexSessionId":"session:01a1927e-9c44-7614-b3bc-0de3fd331a9d","rexSessionName":"rex-whichkey-main","rexWindowId":"window:01a1927e-9c45-7176-82af-1a472289ecc4","rexWindowName":"coordinator","role":"coordinator","runId":"rex-whichkey-p03-coordinator-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-team","specHash":"sha256:2660682d6a3d7f52a00745012de6b45549c888675e0469841b8947991fc9acc2","stage":"launched","thinking":"high","transportDir":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p03-coordinator-a01","windowName":"coordinator","worktree":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->
