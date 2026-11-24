---
identity:
  kind: phase
  phaseID: rex-whichkey-20-background-scripts
  uuid: 1cb3d13c-3303-40a3-8bdc-db4657f3c018
structure:
  order: 20
scope:
  requirements:
    - REQ-SCRIPT-001
    - REQ-SCRIPT-002
    - REQ-SCRIPT-003
    - REQ-SCRIPT-004
  documents:
    - ref: designs/rex-whichkey-background-scripts.md
      hash: 'sha256:1a7cb4ce61a8d2723f031e956b35b9a3e758285a41152d04c6d63debf2e99eb9'
    - ref: designs/rex-whichkey-quality-gates.md
      hash: 'sha256:f3120762f7075c7c767e5f39552ba9c2b38b2e213d33d2472d07f0bc0c1cbdb2'
  dependencies: []
lifecycle:
  state: incomplete
  sourceEventID: rex-whichkey-p20/plan/1
  changedAt: '2026-11-24T02:59:09.884Z'
title: 'Phase 20: Background shell scripts on pane host'
mtnCreateRetry:
  key: rex-whichkey-phase20-background-scripts-v1
  intent: f9139ae027cb9325a9dd94b5cfbc3bb0cf0582188d53e917efceed55337369c8
status: in-progress
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-24T02:34:18.674Z'
dateModified: '2026-11-24T03:01:57.710Z'
tags:
  - task
type: task
id: 46ba0afe-9963-4d5e-a955-96f2f19bba2d
projects:
  - '[[projects/rex-whichkey.md]]'
---
# Phase 20 — Background shell scripts on the originating pane host

## Authority and agreed requirements
The user explicitly requested: "create a phase, break down the tasks and implement the feature" after agreeing that shell scripts run in the background wherever the originating pane lives, use that pane's current working directory, and must report an error without launching if the script does not exist on that host. This is bounded execution authority for decomposition, routine implementation choices, independent reviews, integration and completion, not permission for unrelated scope. Do not ask again for routine phase/task approval. The project authorizes maintainable bounded technical choices. Current Phase19 quality-gate correction supersedes older no-tests/no-experiments clauses: run tests and lint after changes and fix actual failures; preserve standards and assertions.

## Requirements and numbered acceptance criteria
1. REQ-SCRIPT-001: User JSON configuration can map a menu key to running a shell script, while preserving existing built-in actions and strict validation. Design and document a narrow script action with explicit script path and literal argument vector; no implicit shell interpolation of data, arbitrary Lua callbacks or generic Rex API escape hatch.
2. REQ-SCRIPT-002: Dispatch executes in the background on the host of the originating pane, including remote origins, using that pane's current working directory. No visible new pane/tab, focus takeover, blocking the menu on script completion, typing into the existing pane, copying scripts between hosts or fallback to another host. Retain exact origin identity and existing once-only/confirmation/cleanup guarantees. Verify actual Rex APIs/source before choosing implementation; never assume the menu renderer is necessarily the origin host.
3. REQ-SCRIPT-003: Resolve relative script paths from the originating pane's working directory on that execution host. Check script existence there before launching; a missing script produces an actionable error and no execution. Reject unsuitable script targets and report startup/permission/working-directory/host errors without hidden fallback. Scripts remain trusted user executable configuration; document the change from data-only trust and execution-host environment semantics.
4. REQ-SCRIPT-004: Existing config/default actions/menu result-only behavior remain compatible. Add focused automated regression coverage for config validation, literal path/argument handling, background lifetime, correct host/cwd, missing scripts and no-effect error paths. Run required build, tests/race and lint using existing supported tooling. Update user docs and a usable config example, explicitly explaining paths must exist on the pane host.

## Planning and implementation
Capture a concise fingerprinted requirements/design record for this materially consequential execution boundary through mtn. Research relevant code and published Rex CLI/API/source to settle a minimal within-scope mechanism. Decompose into ownership-bounded tasks each owning its own tests and implementation; expected seams are configuration/action contract, native background dispatch with host/cwd preflight, and documentation/integrated verification. Actual task boundaries may combine tightly coupled code to avoid an unused producer or duplicated schema. Add explicit task prerequisites; numbering is not a barrier. No dependency on already-completed earlier phases solely for numbering. Preserve prior phase/task identities and unrelated records; no collection migration.

Every new task needs language review pairs for materially changed languages and explicit architecture/security/usability/maintenance trigger rationales. Architecture, security and usability apply to this feature. Require fresh independent reviews of frozen exact revisions, serial speculative integration and combined target verification. Preserve real RED/GREEN and verification evidence; do not claim live remote execution from mocks/source inspection. Ask only if actual capabilities cannot satisfy accepted host/background requirements without a material scope change.

## Boundaries
Flat short-lived Go application and minimal existing Lua launch integration; no daemon/framework, dependency churn, user config/profile edits, reload, deployment, push or global cleanup. Commit reviewed delivered scope on main using existing project policy. Keep native origin/focus checks and configuration ownership. Continue until delivered or concretely blocked. Successful owned task worktree/window cleanup is authorized under the workflow; failed/dirty/unrelated resources must be retained.

## References and dependencies
Parent projects/rex-whichkey.md. Source authority is the current user conversation, not approval inferred from silence. Existing design designs/rex-whichkey-lua-float-launch.md describes current runtime; designs/rex-whichkey-quality-gates.md describes restored test/lint authority. Phase19 is completed historical context, not a new whole-phase prerequisite. No unrelated open phase is included in this execution selection.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAyOjM1OjEzLjI2OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC1jb29yZGluYXRvci1hMDEvcnVudGltZS9yZXNlcnZlZCIsImludGVudCI6ImU4MzZjMGEzZTBkMDgzZDYzNWEyZDUxYTg0YzU5Yzg0MjNmOTk4ZTI1Y2E4MzUzNGY1MjMwOTBmMWUzNmZhODQifQ -->
### 2026-11-24T02:35:13.269Z · runtime · unspecified

{"attempt":1,"eventID":"rex-whichkey-p20-coordinator-a01/runtime/reserved","handoff":true,"host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"03e1c49e-f1a3-4d8e-9f55-7d5f0f4ef38e"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","piSessionID":"rex-whichkey-p20-coordinator-a01","piSessionName":"coordinator","recordRef":"phases/rex-whichkey-20-background-scripts.md","rexSessionName":"rex-whichkey-main","rexWindowName":"coordinator","role":"coordinator","runID":"rex-whichkey-p20-coordinator-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-team","specHash":"sha256:ef6617e767290fd88977acf8cce006d65f9547629b74c5406118c5f891a90e5d","stage":"reserved","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-coordinator-a01","windowName":"coordinator","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAyOjM1OjE4LjY5MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC1jb29yZGluYXRvci1hMDEvcnVudGltZS9sYXVuY2hlZCIsImludGVudCI6ImZiNGZiNjZjNmJmMDE4OWY2MmJlNjk0MzM5MjQwMjczYmUyNmNjNjllNTdlZTc2NzE3ZDdkNGU4NzE4OWZkNTgifQ -->
### 2026-11-24T02:35:18.691Z · runtime · unspecified

{"attempt":1,"eventID":"rex-whichkey-p20-coordinator-a01/runtime/launched","handoff":true,"host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"03e1c49e-f1a3-4d8e-9f55-7d5f0f4ef38e"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","piSessionID":"rex-whichkey-p20-coordinator-a01","piSessionName":"coordinator","provider":{"displayName":"coordinator","name":"rex","resources":{"block":"block:01a20b68-50b0-7d2f-9c65-5c557c2f82a0","endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","session":"session:01a20b68-50b0-7c48-81eb-9d5947ecb20d","window":"window:01a20b68-50c3-7109-b54c-88235572ac44"}},"recordRef":"phases/rex-whichkey-20-background-scripts.md","rexBlockID":"block:01a20b68-50b0-7d2f-9c65-5c557c2f82a0","rexSessionID":"session:01a20b68-50b0-7c48-81eb-9d5947ecb20d","rexSessionName":"rex-whichkey-main","rexWindowID":"window:01a20b68-50c3-7109-b54c-88235572ac44","rexWindowName":"coordinator","role":"coordinator","runID":"rex-whichkey-p20-coordinator-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-team","specHash":"sha256:ef6617e767290fd88977acf8cce006d65f9547629b74c5406118c5f891a90e5d","stage":"launched","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-coordinator-a01","windowName":"coordinator","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAyOjM3OjE5LjYzMloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC1wbGFubmVyLWEwMS9ydW50aW1lL3Jlc2VydmVkIiwiaW50ZW50IjoiMjU4ZjM4YTNlNTA0YWE1ODcyYzUxNTkyZTJjM2FkM2Y5ZmJiNGU3NDFmZTA4ODJiMjcyNjJkZjc1MzJmNDVjMiJ9 -->
### 2026-11-24T02:37:19.632Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-20-background-scripts.md","runID":"rex-whichkey-p20-coordinator-a01"},"eventID":"rex-whichkey-p20-planner-a01/runtime/reserved","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"03e1c49e-f1a3-4d8e-9f55-7d5f0f4ef38e"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-20-background-scripts.md","piSessionID":"rex-whichkey-p20-planner-a01","piSessionName":"p20-planner-a01","recordRef":"phases/rex-whichkey-20-background-scripts.md","rexSessionName":"rex-whichkey-main","rexWindowName":"p20-planner-a01","role":"planner","runID":"rex-whichkey-p20-planner-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"product-manager","specHash":"sha256:ef6617e767290fd88977acf8cce006d65f9547629b74c5406118c5f891a90e5d","stage":"reserved","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-planner-a01","windowName":"p20-planner-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAyOjM3OjI1LjAyNloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC1wbGFubmVyLWEwMS9ydW50aW1lL2xhdW5jaGVkIiwiaW50ZW50IjoiNDBmMmRkMjkxNzgzYmU3MTQ1Mzk5NjU1NjUwNDYxMTZjNzY0MjhlMGNlMmZkOGVjNTI4MTBjODFmMzczNjJmYiJ9 -->
### 2026-11-24T02:37:25.026Z · runtime · unspecified

{"attempt":1,"coordinator":{"phaseRef":"phases/rex-whichkey-20-background-scripts.md","runID":"rex-whichkey-p20-coordinator-a01"},"eventID":"rex-whichkey-p20-planner-a01/runtime/launched","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"03e1c49e-f1a3-4d8e-9f55-7d5f0f4ef38e"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-20-background-scripts.md","piSessionID":"rex-whichkey-p20-planner-a01","piSessionName":"p20-planner-a01","provider":{"displayName":"p20-planner-a01","name":"rex","resources":{"block":"block:01a20b6a-3e30-7148-9338-fe41175a5d26","endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","session":"session:01a20b68-50b0-7c48-81eb-9d5947ecb20d","window":"window:01a20b6a-3e40-72ea-b1f1-9b1483c6194c"}},"recordRef":"phases/rex-whichkey-20-background-scripts.md","rexBlockID":"block:01a20b6a-3e30-7148-9338-fe41175a5d26","rexSessionID":"session:01a20b68-50b0-7c48-81eb-9d5947ecb20d","rexSessionName":"rex-whichkey-main","rexWindowID":"window:01a20b6a-3e40-72ea-b1f1-9b1483c6194c","rexWindowName":"p20-planner-a01","role":"planner","runID":"rex-whichkey-p20-planner-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"product-manager","specHash":"sha256:ef6617e767290fd88977acf8cce006d65f9547629b74c5406118c5f891a90e5d","stage":"launched","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-planner-a01","windowName":"p20-planner-a01","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAyOjUwOjM1LjkzN1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC1wbGFubmVyLWEwMi9ydW50aW1lL3Jlc2VydmVkIiwiaW50ZW50IjoiNDQ4YzAwMjhiMmFjNGNkMGZkZDcyZjY0ZmM2Y2E5ZTE4NzFkYjAyMjY5MGY0YWQxYjk3NmJhNmNkYjEzM2ZjYSJ9 -->
### 2026-11-24T02:50:35.937Z · runtime · unspecified

{"attempt":2,"coordinator":{"phaseRef":"phases/rex-whichkey-20-background-scripts.md","runID":"rex-whichkey-p20-coordinator-a01"},"eventID":"rex-whichkey-p20-planner-a02/runtime/reserved","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"03e1c49e-f1a3-4d8e-9f55-7d5f0f4ef38e"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-20-background-scripts.md","piSessionID":"rex-whichkey-p20-planner-a02","piSessionName":"p20-planner-a02","recordRef":"phases/rex-whichkey-20-background-scripts.md","rexSessionName":"rex-whichkey-main","rexWindowName":"p20-planner-a02","role":"planner","runID":"rex-whichkey-p20-planner-a02","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"code-architect","specHash":"sha256:ef6617e767290fd88977acf8cce006d65f9547629b74c5406118c5f891a90e5d","stage":"reserved","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-planner-a02","windowName":"p20-planner-a02","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAyOjUwOjQxLjQ0MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC1wbGFubmVyLWEwMi9ydW50aW1lL2xhdW5jaGVkIiwiaW50ZW50IjoiZmVkYWZjMTc3YzEyMGNlNWE2M2ZkYWNiMDViYmQxYzdiYjY2M2QxNzgxMzI1ZWMyNzAzNmFkMzY4N2UyMWExMCJ9 -->
### 2026-11-24T02:50:41.441Z · runtime · unspecified

{"attempt":2,"coordinator":{"phaseRef":"phases/rex-whichkey-20-background-scripts.md","runID":"rex-whichkey-p20-coordinator-a01"},"eventID":"rex-whichkey-p20-planner-a02/runtime/launched","host":{"endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","hostID":"37ba14dc-71d4-468a-a829-32d70294f3e6","instanceID":"03e1c49e-f1a3-4d8e-9f55-7d5f0f4ef38e"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","phaseRef":"phases/rex-whichkey-20-background-scripts.md","piSessionID":"rex-whichkey-p20-planner-a02","piSessionName":"p20-planner-a02","provider":{"displayName":"p20-planner-a02","name":"rex","resources":{"block":"block:01a20b76-64f4-7635-9cd7-23a1960f6b62","endpoint":"unix:///Users/curtbushko/Library/Application%20Support/rex/server.sock","session":"session:01a20b68-50b0-7c48-81eb-9d5947ecb20d","window":"window:01a20b76-6506-7121-8c6f-00d61b46191a"}},"recordRef":"phases/rex-whichkey-20-background-scripts.md","rexBlockID":"block:01a20b76-64f4-7635-9cd7-23a1960f6b62","rexSessionID":"session:01a20b68-50b0-7c48-81eb-9d5947ecb20d","rexSessionName":"rex-whichkey-main","rexWindowID":"window:01a20b76-6506-7121-8c6f-00d61b46191a","rexWindowName":"p20-planner-a02","role":"planner","runID":"rex-whichkey-p20-planner-a02","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"code-architect","specHash":"sha256:ef6617e767290fd88977acf8cce006d65f9547629b74c5406118c5f891a90e5d","stage":"launched","thinking":"high","transportDir":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-planner-a02","windowName":"p20-planner-a02","worktree":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAyOjU2OjA5LjA0MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVsZWFzZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC1wbGFubmVyLWEwMS9yZWxlYXNlL3Jlc2VydmVkIiwiaW50ZW50IjoiYmNiYjZjMDMwOThiYTVjOWYyNWVkNjA3N2NhMGM5ZjA3MTY5NDU1YzUzMTk0Y2ZkZDUxNDM2ZTg1YzgwOGU5ZSJ9 -->
### 2026-11-24T02:56:09.041Z · release · unspecified

{"attempt":1,"dispatchEventID":"rex-whichkey-p20-planner-a01/runtime/reserved","eventID":"rex-whichkey-p20-planner-a01/release/reserved","kind":"release","outcome":"finished","processEvidence":{"exitCode":0,"exitFile":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-planner-a01/exit-code","resultFile":".mtn-tmp/p20/planner-result.json","resultValidatedBy":"mtn-rex result","windowClosure":"owned block no longer found after successful launcher result; auto-closed"},"recordRef":"phases/rex-whichkey-20-background-scripts.md","runID":"rex-whichkey-p20-planner-a01","schemaVersion":1,"specHash":"sha256:ef6617e767290fd88977acf8cce006d65f9547629b74c5406118c5f891a90e5d"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAyOjU2OjE2LjY5OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVsZWFzZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC1wbGFubmVyLWEwMS9yZWxlYXNlL2xhdW5jaGVkIiwiaW50ZW50IjoiMDVjMjBiZGVkYmI4NmMxYmY2MTJkOTVlZjAyMWFiYjVkOGEyZDkxODgwZDJlY2Y0ZDc2ZGM0YzIzMmEzY2U4NyJ9 -->
### 2026-11-24T02:56:16.699Z · release · unspecified

{"attempt":1,"dispatchEventID":"rex-whichkey-p20-planner-a01/runtime/launched","eventID":"rex-whichkey-p20-planner-a01/release/launched","kind":"release","outcome":"finished","processEvidence":{"exitCode":0,"exitFile":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-planner-a01/exit-code","resultFile":".mtn-tmp/p20/planner-result.json","resultValidatedBy":"mtn-rex result","windowClosure":"owned block no longer found after successful launcher result; auto-closed"},"recordRef":"phases/rex-whichkey-20-background-scripts.md","runID":"rex-whichkey-p20-planner-a01","schemaVersion":1,"specHash":"sha256:ef6617e767290fd88977acf8cce006d65f9547629b74c5406118c5f891a90e5d"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAyOjU4OjUwLjIyN1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVsZWFzZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC1wbGFubmVyLWEwMi9yZWxlYXNlL3Jlc2VydmVkIiwiaW50ZW50IjoiMmIyZTZhMTMxOTNkZDdjZGMwODdjNmRjNDc0NmZiYTBlYTQzMTgwMzI4NzljNzIwM2QxODQxMzJkMDZlOTQyMiJ9 -->
### 2026-11-24T02:58:50.227Z · release · unspecified

{"attempt":2,"dispatchEventID":"rex-whichkey-p20-planner-a02/runtime/reserved","eventID":"rex-whichkey-p20-planner-a02/release/reserved","kind":"release","outcome":"finished","processEvidence":{"exitCode":0,"exitFile":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-planner-a02/exit-code","resultFile":".mtn-tmp/p20/architecture-result.json","resultValidatedBy":"mtn-rex result"},"recordRef":"phases/rex-whichkey-20-background-scripts.md","runID":"rex-whichkey-p20-planner-a02","schemaVersion":1,"specHash":"sha256:ef6617e767290fd88977acf8cce006d65f9547629b74c5406118c5f891a90e5d"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAyOjU4OjU3LjgzMloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVsZWFzZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC1wbGFubmVyLWEwMi9yZWxlYXNlL2xhdW5jaGVkIiwiaW50ZW50IjoiMDFhY2JlMDE0YmQ4ZTBiOTI4OGY3ZjA4MzFiYjlmOGYxMjhkOWUxNmNlOWFmN2QwZmY0OGUwMDBiYTI4MWFjMCJ9 -->
### 2026-11-24T02:58:57.832Z · release · unspecified

{"attempt":2,"dispatchEventID":"rex-whichkey-p20-planner-a02/runtime/launched","eventID":"rex-whichkey-p20-planner-a02/release/launched","kind":"release","outcome":"finished","processEvidence":{"exitCode":0,"exitFile":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p20-planner-a02/exit-code","resultFile":".mtn-tmp/p20/architecture-result.json","resultValidatedBy":"mtn-rex result"},"recordRef":"phases/rex-whichkey-20-background-scripts.md","runID":"rex-whichkey-p20-planner-a02","schemaVersion":1,"specHash":"sha256:ef6617e767290fd88977acf8cce006d65f9547629b74c5406118c5f891a90e5d"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAyOjU5OjA4LjcxN1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicGxhbiIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAyMC9wbGFuLzEiLCJpbnRlbnQiOiJkYWI1YWU5NmZjNTBjM2Q5ZTA2OGRiZGRhZjllODI4OTE2YTVmOGM2NjkzOTY0NGZjMDE2OTgzYjMxMjhmZjFkIn0 -->
### 2026-11-24T02:59:08.717Z · plan · unspecified

{"approvedBy":"Coordinator operating under explicit user bounded execution authority","attempt":1,"coverage":[{"records":["tasks/rex-whichkey-20-01-background-scripts.md"],"requirement":"REQ-SCRIPT-001"},{"records":["tasks/rex-whichkey-20-01-background-scripts.md"],"requirement":"REQ-SCRIPT-002"},{"records":["tasks/rex-whichkey-20-01-background-scripts.md"],"requirement":"REQ-SCRIPT-003"},{"records":["tasks/rex-whichkey-20-01-background-scripts.md"],"requirement":"REQ-SCRIPT-004"}],"decision":"approved","decisionEvidence":"User explicitly: create a phase, break down the tasks and implement the feature. Current handoff authorizes bounded decomposition, reviews and integration; no separate user approval of technical details is claimed.","decompositionRationale":"Single coherent vertical task is smaller/safer than separate parser producer, remote backend or docs-only tasks: real script consumer, literal contract, host/cwd proof and error lifetime are tightly coupled. No unused producers or false phase barriers.","dependencyGraph":[],"documents":[{"hash":"sha256:1a7cb4ce61a8d2723f031e956b35b9a3e758285a41152d04c6d63debf2e99eb9","ref":"designs/rex-whichkey-background-scripts.md"},{"hash":"sha256:f3120762f7075c7c767e5f39552ba9c2b38b2e213d33d2472d07f0bc0c1cbdb2","ref":"designs/rex-whichkey-quality-gates.md"}],"eventID":"rex-whichkey-p20/plan/1","impact":"Adds governing document fingerprints and concrete bounded plan without changing accepted requirements. Single vertical task avoids unconsumed producers; no false historical-phase prerequisites.","initialPhaseSpecHash":"sha256:ef6617e767290fd88977acf8cce006d65f9547629b74c5406118c5f891a90e5d","kind":"plan","limits":{"builders":2,"cpu":14,"gates":1,"integrations":1,"memoryGiB":3,"reviewers":2},"planningEvidence":[".mtn-tmp/p20/planner-result.json",".mtn-tmp/p20/architecture-result.json"],"recordRef":"phases/rex-whichkey-20-background-scripts.md","records":[{"dependencies":[],"ref":"tasks/rex-whichkey-20-01-background-scripts.md","requirements":["REQ-SCRIPT-001","REQ-SCRIPT-002","REQ-SCRIPT-003","REQ-SCRIPT-004"],"specHash":"sha256:9f2ef6e168a87f200c462f922af39383e525968aca2f0222c2c22e4d6f94d783"}],"requirements":["REQ-SCRIPT-001","REQ-SCRIPT-002","REQ-SCRIPT-003","REQ-SCRIPT-004"],"reviewKeys":["go:code","go:test","bash:code","bash:test","architecture","security","usability"],"reviewPlanRationale":"Go owns executable config and actual dispatch; nontrivial embedded shell fixtures get independent Bash code/test seats. User documentation is assessed by usability plus Go public-contract review, not an artificial standalone docs producer. Architecture/security/usability triggered; maintenance not triggered.","schemaVersion":1,"sourceBaseline":{"commit":"fb92eae","validation":[{"command":"task build","evidence":".mtn-tmp/p20/baseline/build.log","result":"pass"},{"command":"task test","evidence":".mtn-tmp/p20/baseline/test.log","result":"pass"},{"command":"task lint","evidence":".mtn-tmp/p20/baseline/lint.log","result":"pass"}]},"specHash":"sha256:87545cff84389f3bd1cdef242ff6825cdff7b7f5280ec90fca511221b07312f0"}
<!-- mtn:end-entry:v1 -->
