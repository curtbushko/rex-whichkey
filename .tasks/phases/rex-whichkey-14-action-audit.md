---
workflowKind: phase
workflowOrder: 14
workflowRequirements:
  - REQ-030
  - REQ-031
workflowDocuments:
  - ref: designs/rex-whichkey-action-audit.md
    hash: 'sha256:2d2d664eb4bbc5828b2f0d6d4a98346becfbbd03d90e8f9abc5a9874edc8b7a8'
title: 'Phase14: Native action audit and late targeting'
mtnCreateRetry:
  key: rex-whichkey-phase14-v1
  intent: 78526c197c066e6d45f239adde9c381d8850ec6ab6e9aefc13e222b42c0a7a73
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T20:15:57.031Z'
dateModified: '2026-11-02T20:19:34.358Z'
tags:
  - task
type: task
id: 7f4f23e3-0ce0-4aab-9dbb-6f732e7b064b
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy:
  - uid: '[[phases/rex-whichkey-13-fixed-navigation.md]]'
    reltype: FINISHTOSTART
---
# Phase14: Native actions and late app targeting

Latest user asks verify every action, specifically broken split down. REQ031 repair shared ambiguity via explicit late app choice, not wrong direction changes/mandatory flags/arbitrary targeting. REQ030 verify all25 native defaults by behavior. Depend on completed fixed-navigation phase13. One source repair, dependent disposable resource audit/evidence documentation (18 resource bindings), separate GUI-only verification (7 bindings), blocked if native GUI observability remains unavailable. Do not silently count RPC receipts as visual behavior. Exact approved design governs source/runtime scope and disposable-only action cleanup authority. Keep phase/project incomplete if GUI checks unresolved; REQ026 chrome blocker remains independently unresolved. No generic tests, settings/config/font changes or push.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDIwOjE2OjQ1Ljg3NloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicGxhbiIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L2FjdGlvbi1hdWRpdC9waGFzZS1wbGFuLzEiLCJpbnRlbnQiOiIwYWU3YjhhN2NlNWZlMTgzZDY2NTk3M2VkNDEzNjMxODY1YzkwZmZmMjc5NjBlMWYyZmRhZTZjMGMzMDE1MGJjIn0 -->
### 2026-11-02T20:16:45.876Z · plan · unspecified

{"approvedBy":"Explicit all-action verification and delegated decomposition","attempt":1,"coverage":{"REQ-030":["tasks/rex-whichkey-14-02-resource-audit.md","tasks/rex-whichkey-14-03-gui-action-verification.md"],"REQ-031":["tasks/rex-whichkey-14-01-action-target.md"]},"decision":"approved","decisionEvidence":"Source repair after completed fixed navigation; dependent resource audit18 and distinct GUI-only7 verification. Missing GUI observation is blocker, not pass or silent requirement waiver.","documents":[{"hash":"sha256:2d2d664eb4bbc5828b2f0d6d4a98346becfbbd03d90e8f9abc5a9874edc8b7a8","ref":"designs/rex-whichkey-action-audit.md"}],"eventId":"rex-whichkey/action-audit/phase-plan/1","kind":"plan","recordRef":"phases/rex-whichkey-14-action-audit.md","records":[{"ref":"tasks/rex-whichkey-14-01-action-target.md","specHash":"sha256:96e3d8f88d25c2635ed78d7ca7de32eff427d5b07798c755451b9fdce4022e1e"},{"ref":"tasks/rex-whichkey-14-02-resource-audit.md","specHash":"sha256:b621b5b66ba40eba979553ae608dfa32b50aa1da4ff8baaa7ef4b2c998b06490"},{"ref":"tasks/rex-whichkey-14-03-gui-action-verification.md","specHash":"sha256:85f4b6e1dfbcf4e5735383551d7eb33ec07c1bb7c65dd5f2b4a53371c074ecf7"}],"schemaVersion":1,"specHash":"sha256:6132d927b1d5f34e423883d519b2da44374ebf01a6dc14e1c2190cf13316684a"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDIwOjE3OjIwLjU4NloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcnVudGltZS9hY3Rpb24tYXVkaXQiLCJpbnRlbnQiOiJjZjBmMjYxMDg4ZGI1MjU0NTEyMjEyMzc0Y2IyOTE1ZDNkNTI3N2EyMjExMzRkY2Q5MzA1MmVlYTg3ZjBiOWY1In0 -->
### 2026-11-02T20:17:20.586Z · runtime · unspecified

{"attempt":1,"continuationEvidence":"Same live coordinator process/runId/session/window handles latest all-action verification; no coordinator restart.","eventId":"rex-whichkey-p03-coordinator-a01/runtime/action-audit","handoff":true,"host":{"endpoint":"unix:///run/user/1000/rex/server.sock","hostId":"cf7ceda2-6cd7-468c-9360-3b7d68de6f91","instanceId":"0f6bfb2f-6739-42b6-a8cb-b226b8a8ec14"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","piSessionId":"rex-whichkey-p03-coordinator-a01","piSessionName":"coordinator","recordRef":"phases/rex-whichkey-14-action-audit.md","rexBlockId":"block:01a1927e-9c44-7693-8d92-6ec72cd57f67","rexSessionId":"session:01a1927e-9c44-7614-b3bc-0de3fd331a9d","rexSessionName":"rex-whichkey-main","rexWindowId":"window:01a1927e-9c45-7176-82af-1a472289ecc4","rexWindowName":"coordinator","role":"coordinator","runId":"rex-whichkey-p03-coordinator-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-team","specHash":"sha256:6132d927b1d5f34e423883d519b2da44374ebf01a6dc14e1c2190cf13316684a","stage":"launched","thinking":"high","transportDir":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p03-coordinator-a01","windowName":"coordinator","worktree":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDIwOjE5OjM0LjM0NloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicGxhbiIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L2FjdGlvbi1hdWRpdC9waGFzZS1wbGFuLzIiLCJpbnRlbnQiOiIyZjY3YTE0NTE3YTRjM2Q0ZDJiMDIxOTU1OGQwYjczYTkzNDVjNzVhMjg3ZDI4OTJlOTFjZTg4ZjRkNDFjYjgyIn0 -->
### 2026-11-02T20:19:34.346Z · plan · unspecified

{"approvedBy":"Explicit all-action verification and delegated decomposition","attempt":1,"coverage":{"REQ-030":["tasks/rex-whichkey-14-02-resource-audit.md","tasks/rex-whichkey-14-03-gui-action-verification.md"],"REQ-031":["tasks/rex-whichkey-14-01-action-target.md"]},"decision":"approved","decisionEvidence":"Source repair after completed fixed navigation; dependent resource audit18 and distinct GUI-only7 verification. Missing GUI observation is blocker, not pass or silent requirement waiver. GUI verification owns its future concrete evidence document; semantics unchanged.","documents":[{"hash":"sha256:2d2d664eb4bbc5828b2f0d6d4a98346becfbbd03d90e8f9abc5a9874edc8b7a8","ref":"designs/rex-whichkey-action-audit.md"}],"eventId":"rex-whichkey/action-audit/phase-plan/2","kind":"plan","recordRef":"phases/rex-whichkey-14-action-audit.md","records":[{"ref":"tasks/rex-whichkey-14-01-action-target.md","specHash":"sha256:96e3d8f88d25c2635ed78d7ca7de32eff427d5b07798c755451b9fdce4022e1e"},{"ref":"tasks/rex-whichkey-14-02-resource-audit.md","specHash":"sha256:b621b5b66ba40eba979553ae608dfa32b50aa1da4ff8baaa7ef4b2c998b06490"},{"ref":"tasks/rex-whichkey-14-03-gui-action-verification.md","specHash":"sha256:89d37457c7f33042dbf5f62f87d0f3726fd10c4b1cf9d118060533dad8f084cb"}],"schemaVersion":1,"specHash":"sha256:6132d927b1d5f34e423883d519b2da44374ebf01a6dc14e1c2190cf13316684a"}
<!-- mtn:end-entry:v1 -->
