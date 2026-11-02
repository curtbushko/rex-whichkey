---
workflowKind: phase
workflowOrder: 8
workflowRequirements:
  - REQ-018
  - REQ-019
  - REQ-020
workflowDocuments:
  - ref: designs/rex-whichkey-popup-polish.md
    hash: 'sha256:ce42e479546a1181ce61138f0f5d37c88e5ed18ae3b6c39987fca90d89aee8f0'
title: 'Phase08: Popup polish'
mtnCreateRetry:
  key: rex-whichkey-phase08-popup-polish-v1
  intent: 65c5b228bf90a9f32bb24ba089b8253d6666e70c8d175757c14fc2ad1db6cb09
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T13:16:59.384Z'
dateModified: '2026-11-02T13:19:27.788Z'
tags:
  - task
type: task
id: 6e214899-4476-48a7-90be-ff12fb70e837
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy:
  - uid: '[[phases/rex-whichkey-07-usability.md]]'
    reltype: FINISHTOSTART
---
# Phase08: Popup polish

## Accepted user changes

REQ-018 normal native popup stays open while idle and expected configured expiry is not a child failure; REQ-019 remove Focus and Rex-known-hosts from native/Lua defaults; REQ-020 distinct visible key/icon/label colors inspired by the linked Neovim WhichKey configuration. See designs/rex-whichkey-popup-polish.md for actual scope, retained safety and source-versus-runtime evidence.

## Acceptance

1. A native popup remains visible and usable past the old30-second bound until selection/Escape/cancellation, while individual operations/cleanup remain bounded. Existing configured Lua capability expiry remains fail-closed and dismisses cleanly.
2. Both delivered default sources omit Focus/fh/fj/fk/fl and automatic known-hosts/h. Origin focus restoration, user mappings and explicitly configured host provider remain supported; no inventory acquisition without an explicit provider.
3. Keys, icons and labels use real separate color roles; icons default to yellow, keys cyan/blue, labels retain readable primary/group styling. A terminal without OSC palette/background replies still receives usable semantic ANSI color; no forced background/palette mutation. Honor NO_COLOR, no-color/ASCII options and alignment/sanitization.
4. Docs describe final defaults/lifetimes/colors and accurate runtime/compile evidence. Actual final native Linux/Rex popup idles at least35seconds, shows the reduced colored menu, navigates/dismisses with exit0 and closes only owned resources. No destructive action, config reload, general tests or remote/Darwin claims.
5. Independently source-review every task and compile combined candidates. Final docs/assembly depends on all producers; complete/compact native records only after actual runtime gate and commit phase delivery on main.

## Dependencies and execution

Uses delivered Phase07 native popup and canonical Go mapping/provider/terminal infrastructure. Do not reopen or renumber historical phases/tasks. Four ownership-bounded tasks: lifecycle, defaults, theme/render, final docs/assembly. First three are source-disjoint; two builders/two reviewers, one compile/runtime gate and serialized integration. Same existing coordinator/Rex session, no duplicate process. User scope and bounded technical authority authorize decomposition/implementation, no routine approval pause. Preserve original-VM callback/security guarantees and don't modify user configuration/settings or test/lint/Taskfile/hook definitions.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDEzOjE5OjAzLjg0MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicGxhbiIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3BvcHVwLXBvbGlzaC9waGFzZS1wbGFuLzEiLCJpbnRlbnQiOiIwNmM5MjhhYjcxNmU5YmRhZjgxYjhmOTdmNThiZGE0MGZlODkxYmRlNDM5NjE0MTUwMjkwNThkNjU2Nzc3MzRkIn0 -->
### 2026-11-02T13:19:03.841Z · plan · unspecified

{"approvedBy":"user requested scope; coordinator task decomposition under delegation","attempt":1,"coverage":{"REQ-018":["tasks/rex-whichkey-08-01-lifetime.md","tasks/rex-whichkey-08-04-docs.md"],"REQ-019":["tasks/rex-whichkey-08-02-defaults.md","tasks/rex-whichkey-08-04-docs.md"],"REQ-020":["tasks/rex-whichkey-08-03-theme.md","tasks/rex-whichkey-08-04-docs.md"]},"decision":"approved","decisionEvidence":"User requests timeout/default/color corrections. Four literal ownership tasks separate three independent producers from final docs/runtime assembly; max2 builders/reviewers and one heavy gate. Native idle removal does not weaken configured expiry. No general tests or config changes authorized.","documents":[{"hash":"sha256:ce42e479546a1181ce61138f0f5d37c88e5ed18ae3b6c39987fca90d89aee8f0","ref":"designs/rex-whichkey-popup-polish.md"}],"eventId":"rex-whichkey/popup-polish/phase-plan/1","kind":"plan","recordRef":"phases/rex-whichkey-08-popup-polish.md","records":[{"ref":"tasks/rex-whichkey-08-01-lifetime.md","specHash":"sha256:eb23f711d55a1496beb05e725256bc7163378f109618c30fee11ce5a484bdf49"},{"ref":"tasks/rex-whichkey-08-02-defaults.md","specHash":"sha256:5b38acfb7f05dc4362b40e8d94198858da08423ba27e249df031dc3c574d1313"},{"ref":"tasks/rex-whichkey-08-03-theme.md","specHash":"sha256:a495094e2fb1c8a0638831f1b057f6f9c5594574a3621daa689067158d1db016"},{"ref":"tasks/rex-whichkey-08-04-docs.md","specHash":"sha256:d489e0bc7508abf668a3b063539b0d828f38d957dba53e9e87283392324974d2"}],"schemaVersion":1,"specHash":"sha256:baf78cecbd7f4271b318cae4cb0bccf51e850a035096b539d55b3e6a6e27802a"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDEzOjE5OjI3Ljc3NloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicnVudGltZSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcnVudGltZS9wb3B1cC1wb2xpc2giLCJpbnRlbnQiOiI5M2Q2MzQ0OWNhMjhjZGQ0NGYyOWJhMjk4MGZhYTgxMGIyOWY4MWFhMzE1ZGVhNDQ5MGRjOTBlNzMzNDZjOGYxIn0 -->
### 2026-11-02T13:19:27.776Z · runtime · unspecified

{"attempt":1,"continuationEvidence":"Same live coordinator process and immutable runId starts user-requested Phase08 with existing session/window; no restart/duplicate. Display remains coordinator, phase association separately recorded.","eventId":"rex-whichkey-p03-coordinator-a01/runtime/popup-polish","handoff":true,"host":{"endpoint":"unix:///run/user/1000/rex/server.sock","hostId":"cf7ceda2-6cd7-468c-9360-3b7d68de6f91","instanceId":"0f6bfb2f-6739-42b6-a8cb-b226b8a8ec14"},"kind":"runtime","model":"openai-codex/gpt-6.1-sol","piSessionId":"rex-whichkey-p03-coordinator-a01","piSessionName":"coordinator","recordRef":"phases/rex-whichkey-08-popup-polish.md","rexBlockId":"block:01a1927e-9c44-7693-8d92-6ec72cd57f67","rexSessionId":"session:01a1927e-9c44-7614-b3bc-0de3fd331a9d","rexSessionName":"rex-whichkey-main","rexWindowId":"window:01a1927e-9c45-7176-82af-1a472289ecc4","rexWindowName":"coordinator","role":"coordinator","runId":"rex-whichkey-p03-coordinator-a01","schemaVersion":1,"sessionName":"rex-whichkey-main","skill":"go-team","specHash":"sha256:baf78cecbd7f4271b318cae4cb0bccf51e850a035096b539d55b3e6a6e27802a","stage":"launched","thinking":"high","transportDir":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/rex/rex-whichkey-p03-coordinator-a01","windowName":"coordinator","worktree":"/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main"}
<!-- mtn:end-entry:v1 -->
