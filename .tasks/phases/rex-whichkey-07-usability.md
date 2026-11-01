---
workflowKind: phase
workflowOrder: 7
workflowRequirements:
  - REQ-016
  - REQ-017
workflowDocuments:
  - ref: prds/rex-whichkey-usability.md
    hash: 'sha256:c2df9b9404ce7f2d649eed40d5c219f7098e9389c6407159b6ae35eddb7746c5'
  - ref: designs/rex-whichkey-usability.md
    hash: 'sha256:c8553dd4ac48fcd34a43d7cb3cd842bb18c65310d0fc6133cf1b6e3f196f1b8a'
title: 'Phase07: Usability'
mtnCreateRetry:
  key: rex-whichkey-phase07-usability-v1
  intent: ae9f3baa27cca9ed342f10f90d5169d9dc0d5648285767f2f805a2e4258ddc07
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-01T14:52:26.658Z'
dateModified: '2026-11-01T15:16:38.449Z'
tags:
  - task
type: task
id: a6773f97-4e50-4960-ae86-9c96af731a54
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy:
  - uid: '[[phases/rex-whichkey-06-packaging-release.md]]'
    reltype: FINISHTOSTART
---
# Phase07: Usability

## Accepted scope

User explicitly requests this new phase, task decomposition and implementation. REQ-016: no-argument rex-whichkey launches the real popup/menu instead of a scaffold greeting or internal-command usage. REQ-017: Linux/macOS platform socket defaults remove common-case mandatory socket arguments while retaining explicit overrides and actual-origin safety. See prds/rex-whichkey-usability.md; existing requirements remain in prds/rex-whichkey.md.

## Acceptance criteria

1. A real production no-argument startup path launches the delivered menu/popup with normal integration prerequisites; it needs neither manually authored resolved JSON nor direct invocation of private modes. Preserve existing internal subcommands and original Lua runtime callbacks.
2. Canonical Go-owned endpoint resolution delegates Linux/macOS path conventions to the installed Rex CLI's native platform default discovery with deterministic override precedence; explicit overrides work and the actual returned endpoint is captured before further calls. The Lua launcher can omit its endpoint in the normal platform-default case without copying/guessing opaque ctx.server fields.
3. CLI and Lua consumers use the same Go-owned endpoint policy. Validate captured session/block/client/connection and server realm before mutation; ambiguity/missing origin fails safely and clearly. No ambient focused/current-client fallback, daemon, duplicate VM or user init takeover.
4. User help/README/examples explain ordinary launch, platform defaults and overrides, existing setup and Remote Control prerequisites, and actionable missing-context failures. Displayed public usage does not present private controllers as the normal start command.
5. Each task receives fresh independent source review, combined-target compilation and actual Git integration proof. Phase completion is committed on main. No project tests/TDD/experiments/GUI/live/remote trials/config reloads; static checks and compilation are accurately labeled.

## Planning and implementation

Task Manager proposes viable original-runtime no-argument activation and a small ownership-bounded graph. Coordinator selected native Rex CLI endpoint discovery instead of reimplementing undocumented platform path algorithms: explicit override > nonempty REX_SERVER > bounded no-autostart whoami JSON; validate returned endpoint and keep every subsequent call explicit. No exact platform filename assumption is required. Prefer three coherent concerns: endpoint policy plus Go/Lua consumption; zero-argument launch wiring; documentation/final assembly. Split or serialize overlapping files/interfaces honestly. Do not add approval pauses for routine choices already delegated; unsupported capabilities are genuine factual blockers, not permission to guess.

## Dependencies and scope impact

Consumes delivered phases03–06 sources and their existing integration; it extends the completed project without reopening or renumbering their immutable task identities. The old explicit-only endpoint choice REXWHICHKEY-20261101-015342 must be refined/superseded only to the extent required by the user's new platform-default requirement; preserve its no-opaque-context-conversion, explicit-override and fail-closed safety rationale. No existing implementation or runtime success is implied by adding this phase.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDE0OjU0OjAzLjg3M1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcnVudGltZS91c2FiaWxpdHktY29udGludWF0aW9uIiwiaW50ZW50IjoiZTE2OGMwODEwZWI0OTc1Njk3OTI0NzAwNzEyMzVlMDU0YTI0NDEwMmVjOWQ0YTJlMzUyNWEyNTJmZTFiMDA0MiJ9 -->
### 2026-11-01T14:54:03.873Z · runtime · unspecified

{"attempt":1,"continuationEvidence":"Same live coordinator process, exact immutable runId and existing Rex session/window/block. User explicitly requested new Phase07 after prior delivery; original startup phase/run identity retained, phase association is separate. No process or conversation created/restarted.","eventId":"rex-whichkey-p03-coordinator-a01/runtime/usability-continuation","handoff":true,"host":{"endpoint":"unix:///run/user/1000/rex/server.sock","hostId":"cf7ceda2-6cd7-468c-9360-3b7d68de6f91","instanceId":"0f6bfb2f-6739-42b6-a8cb-b226b8a8ec14"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","piSessionId":"rex-whichkey-p03-coordinator-a01","piSessionName":"coordinator","recordRef":"phases/rex-whichkey-07-usability.md","rexBlockId":"block:01a1927e-9c44-7693-8d92-6ec72cd57f67","rexSessionId":"session:01a1927e-9c44-7614-b3bc-0de3fd331a9d","rexSessionName":"rex-whichkey-main","rexWindowId":"window:01a1927e-9c45-7176-82af-1a472289ecc4","rexWindowName":"coordinator","role":"coordinator","runId":"rex-whichkey-p03-coordinator-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-team","specHash":"sha256:d4bb433b44f4fc23d5834cb5abbab105e3e21ad6cb4f8546193b8f3722a50ec3","stage":"launched","thinking":"high","transportDir":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p03-coordinator-a01","windowName":"coordinator","worktree":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDE0OjU0OjEwLjE2MFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNy1wbGFubmVyLWEwMS9ydW50aW1lL3Jlc2VydmVkIiwiaW50ZW50IjoiYzhkMmZkZWVlNGYzNjBkZTM1M2EzM2RkY2IxOGVhZWIxYjZlMTg2MGRlM2VjNDk1YmFiMDY4ZGM2NGM5OGExNiJ9 -->
### 2026-11-01T14:54:10.160Z · runtime · unspecified

{"attempt":1,"coordinatorPhaseRef":"phases/rex-whichkey-07-usability.md","eventId":"rex-whichkey-p07-planner-a01/runtime/reserved","host":{"endpoint":"unix:///run/user/1000/rex/server.sock","hostId":"cf7ceda2-6cd7-468c-9360-3b7d68de6f91","instanceId":"0f6bfb2f-6739-42b6-a8cb-b226b8a8ec14"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","parentRunId":"rex-whichkey-p03-coordinator-a01","phaseRef":"phases/rex-whichkey-07-usability.md","piSessionId":"rex-whichkey-p07-planner-a01","piSessionName":"p07-planner-a01","recordRef":"phases/rex-whichkey-07-usability.md","rexSessionName":"rex-whichkey-main","rexWindowName":"p07-planner-a01","role":"planner","runId":"rex-whichkey-p07-planner-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"mtn-tasks","specHash":"sha256:d4bb433b44f4fc23d5834cb5abbab105e3e21ad6cb4f8546193b8f3722a50ec3","stage":"reserved","thinking":"high","transportDir":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p07-planner-a01","windowName":"p07-planner-a01","worktree":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDE0OjU0OjEyLjgzOFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNy1wbGFubmVyLWEwMS9ydW50aW1lL2xhdW5jaGVkIiwiaW50ZW50IjoiMzRkMzhkZWY0OThiZGU1ZWQyYzg5YTU1YThkMmFjOTEzMGU1ODZhNWQxNjhiOTQ3ZjY4YTE5MDQwNjFkNzA4OCJ9 -->
### 2026-11-01T14:54:12.838Z · runtime · unspecified

{"attempt":1,"coordinatorPhaseRef":"phases/rex-whichkey-07-usability.md","eventId":"rex-whichkey-p07-planner-a01/runtime/launched","host":{"endpoint":"unix:///run/user/1000/rex/server.sock","hostId":"cf7ceda2-6cd7-468c-9360-3b7d68de6f91","instanceId":"0f6bfb2f-6739-42b6-a8cb-b226b8a8ec14"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","parentRunId":"rex-whichkey-p03-coordinator-a01","phaseRef":"phases/rex-whichkey-07-usability.md","piSessionId":"rex-whichkey-p07-planner-a01","piSessionName":"p07-planner-a01","recordRef":"phases/rex-whichkey-07-usability.md","rexBlockId":"block:01a1979a-8b2c-7da1-9b1f-3da805f99d71","rexSessionId":"session:01a1927e-9c44-7614-b3bc-0de3fd331a9d","rexSessionName":"rex-whichkey-main","rexWindowId":"window:01a1979a-8b2d-7a3d-9ba4-88dfdcd03270","rexWindowName":"p07-planner-a01","role":"planner","runId":"rex-whichkey-p07-planner-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"mtn-tasks","specHash":"sha256:d4bb433b44f4fc23d5834cb5abbab105e3e21ad6cb4f8546193b8f3722a50ec3","stage":"launched","thinking":"high","transportDir":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p07-planner-a01","windowName":"p07-planner-a01","worktree":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDE1OjE0OjQ3LjcwOFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYmxvY2tlciIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNy1wbGFubmVyLWEwMS9CMSIsImludGVudCI6ImNmMjM3ODhmNTQ4NmM0MTZjNDQ1ODlhYjAyN2RhOTljNDFlNDdmOTVmMWQxMmU4MGU5ZDMwOTI1NTlkZTZiMTQifQ -->
### 2026-11-01T15:14:47.708Z · blocker · unspecified

{"attempt":1,"eventId":"rex-whichkey-p07-planner-a01/B1","kind":"blocker","proposalRef":"REXWHICHKEY-20261101-100106","reason":"Exact native socket algorithms were unavailable in inspected source/docs. Coordinator must choose a supported source contract instead of guessing. Detailed frozen planning evidence is in .mtn-tmp/usability/planner-plan.json and planner-source-contracts.md.","recordRef":"phases/rex-whichkey-07-usability.md","schemaVersion":1,"specHash":"sha256:d4bb433b44f4fc23d5834cb5abbab105e3e21ad6cb4f8546193b8f3722a50ec3"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDE1OjE0OjUwLjM3OFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVsZWFzZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNy1wbGFubmVyLWEwMS9yZWxlYXNlLXJlc2VydmVkLzEiLCJpbnRlbnQiOiI2MDYyYWQyMmRmZmU0OTFmZTI2MmYxMTBhNWI0YzQyYzY2YzI0MDBiMjljODMzYzViMjE2NWNkODkyYWQ2Njg1In0 -->
### 2026-11-01T15:14:50.378Z · release · unspecified

{"attempt":1,"dispatchEventId":"rex-whichkey-p07-planner-a01/runtime/reserved","eventId":"rex-whichkey-p07-planner-a01/release-reserved/1","kind":"release","outcome":"finished","processEvidence":{"exitCode":0,"exitCodeFile":".mtn-tmp/rex/rex-whichkey-p07-planner-a01/exit-code","resultFile":".mtn-tmp/usability/planner-plan.json"},"recordRef":"phases/rex-whichkey-07-usability.md","schemaVersion":1,"specHash":"sha256:d4bb433b44f4fc23d5834cb5abbab105e3e21ad6cb4f8546193b8f3722a50ec3"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDE1OjE0OjUzLjA0M1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVsZWFzZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNy1wbGFubmVyLWEwMS9yZWxlYXNlLWxhdW5jaGVkLzEiLCJpbnRlbnQiOiI1ODU2MTYyNzQ2MDJjMGRlMmY2ODYxNTk0MDNiMGRlOTEyN2RjMzdlYjk4NDdlMGI0YmExYjE2NzI2ODU2ODM5In0 -->
### 2026-11-01T15:14:53.043Z · release · unspecified

{"attempt":1,"dispatchEventId":"rex-whichkey-p07-planner-a01/runtime/launched","eventId":"rex-whichkey-p07-planner-a01/release-launched/1","kind":"release","outcome":"finished","processEvidence":{"exitCode":0,"exitCodeFile":".mtn-tmp/rex/rex-whichkey-p07-planner-a01/exit-code","resultFile":".mtn-tmp/usability/planner-plan.json"},"recordRef":"phases/rex-whichkey-07-usability.md","schemaVersion":1,"specHash":"sha256:d4bb433b44f4fc23d5834cb5abbab105e3e21ad6cb4f8546193b8f3722a50ec3"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDE1OjE2OjE2LjM0MFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicGxhbiIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3VzYWJpbGl0eS9waGFzZS1wbGFuLzEiLCJpbnRlbnQiOiI2ZTlhMzQxM2MyYTFlMDlkM2U1YjMzNzQ1Yjg2ZDlhNzAwMDc0NzU2YTk4MDI3OTA4ZmM3YmRjNGRjMmZkOGJlIn0 -->
### 2026-11-01T15:16:16.340Z · plan · unspecified

{"approvedBy":"coordinator under explicit user task-decomposition/implementation authority","attempt":1,"coverage":{"REQ-016":["tasks/rex-whichkey-07-02-direct-launch.md","tasks/rex-whichkey-07-03-usage-assembly.md"],"REQ-017":["tasks/rex-whichkey-07-01-endpoint-policy.md","tasks/rex-whichkey-07-02-direct-launch.md","tasks/rex-whichkey-07-03-usage-assembly.md"]},"criteria":{"1":["tasks/rex-whichkey-07-02-direct-launch.md","tasks/rex-whichkey-07-03-usage-assembly.md"],"2":["tasks/rex-whichkey-07-01-endpoint-policy.md"],"3":["tasks/rex-whichkey-07-01-endpoint-policy.md","tasks/rex-whichkey-07-02-direct-launch.md"],"4":["tasks/rex-whichkey-07-02-direct-launch.md","tasks/rex-whichkey-07-03-usage-assembly.md"],"5":["tasks/rex-whichkey-07-01-endpoint-policy.md","tasks/rex-whichkey-07-02-direct-launch.md","tasks/rex-whichkey-07-03-usage-assembly.md"]},"decision":"approved","decisionEvidence":"User explicitly requests the phase and work. Native resolver design resolves the source-information blocker without guessing paths; three ownership-bounded serial tasks cover the actual Go/Lua path and documentation. No project tests/runtime passes or further user agreement invented.","documents":[{"hash":"sha256:c2df9b9404ce7f2d649eed40d5c219f7098e9389c6407159b6ae35eddb7746c5","ref":"prds/rex-whichkey-usability.md"},{"hash":"sha256:c8553dd4ac48fcd34a43d7cb3cd842bb18c65310d0fc6133cf1b6e3f196f1b8a","ref":"designs/rex-whichkey-usability.md"}],"eventId":"rex-whichkey/usability/phase-plan/1","kind":"plan","recordRef":"phases/rex-whichkey-07-usability.md","records":[{"ref":"tasks/rex-whichkey-07-01-endpoint-policy.md","specHash":"sha256:04c3bc95c1b164f2d02e6d500995fb8d01e7a6630369deebb8ec1105d437dc3d"},{"ref":"tasks/rex-whichkey-07-02-direct-launch.md","specHash":"sha256:408dd4478e572880d5b6fe9d0c3354876e5d583f47569f398c3c14d40e60af6b"},{"ref":"tasks/rex-whichkey-07-03-usage-assembly.md","specHash":"sha256:6a20f5f77ec91c0b496a65a0e6946c035ccd78920606cf2401b367289c819a46"}],"schemaVersion":1,"specHash":"sha256:c8917e4b87c3bced79a2d66bedd6d34b900209a1a4f83d5c7d98fc29a26f753b"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDE1OjE2OjE5LjAyNVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVzdW1lIiwicmV0cnlLZXkiOiJyZXgtd2hpY2hrZXkvdXNhYmlsaXR5L25hdGl2ZS1kaXNjb3ZlcnktcmVzdW1lLzEiLCJpbnRlbnQiOiJkNTU3MDA0Y2M0MWEwN2I5ZDQ1MzgwNjE0NGQ0NjMwMmIyMWQwYjZkYjI0MGRmYmIyYTk0OWRmMTU4Y2JkM2VkIn0 -->
### 2026-11-01T15:16:19.025Z · resume · unspecified

{"attempt":1,"blockerEventId":"rex-whichkey-p07-planner-a01/B1","eventId":"rex-whichkey/usability/native-discovery-resume/1","kind":"resume","recordRef":"phases/rex-whichkey-07-usability.md","resolution":"Approved native endpoint design delegates platform path algorithms to installed Rex whoami JSON with autostart disabled, validates returned concrete endpoint and preserves explicit targeting. Exact undocumented formulas no longer need duplication. Tasks/docs and approved manifests reflect the new contract. This unblocks source implementation, not runtime acceptance.","schemaVersion":1,"specHash":"sha256:c8917e4b87c3bced79a2d66bedd6d34b900209a1a4f83d5c7d98fc29a26f753b"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDE1OjE2OjM4LjQzN1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcnVudGltZS91c2FiaWxpdHktaW1wbGVtZW50YXRpb24iLCJpbnRlbnQiOiI0NjVmNjQ1MDhhODkwZGQ3OTdhMmIwZWQxMmQwMDc5ZWNmYzdhZDM5M2Q0MzhhYzkwMmY5MGUzYjZhZWQyNDExIn0 -->
### 2026-11-01T15:16:38.437Z · runtime · unspecified

{"attempt":1,"continuationEvidence":"Same live coordinator process, exact immutable runId and existing Rex session/window/block. User explicitly requested new Phase07 after prior delivery; original startup phase/run identity retained, phase association is separate. No process or conversation created/restarted.","eventId":"rex-whichkey-p03-coordinator-a01/runtime/usability-implementation","handoff":true,"host":{"endpoint":"unix:///run/user/1000/rex/server.sock","hostId":"cf7ceda2-6cd7-468c-9360-3b7d68de6f91","instanceId":"0f6bfb2f-6739-42b6-a8cb-b226b8a8ec14"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","piSessionId":"rex-whichkey-p03-coordinator-a01","piSessionName":"coordinator","recordRef":"phases/rex-whichkey-07-usability.md","rexBlockId":"block:01a1927e-9c44-7693-8d92-6ec72cd57f67","rexSessionId":"session:01a1927e-9c44-7614-b3bc-0de3fd331a9d","rexSessionName":"rex-whichkey-main","rexWindowId":"window:01a1927e-9c45-7176-82af-1a472289ecc4","rexWindowName":"coordinator","role":"coordinator","runId":"rex-whichkey-p03-coordinator-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-team","specHash":"sha256:c8917e4b87c3bced79a2d66bedd6d34b900209a1a4f83d5c7d98fc29a26f753b","stage":"launched","thinking":"high","transportDir":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p03-coordinator-a01","windowName":"coordinator","worktree":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->
