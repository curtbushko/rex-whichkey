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
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T13:16:59.384Z'
dateModified: '2026-11-02T14:13:45.860Z'
tags:
  - task
type: task
id: 6e214899-4476-48a7-90be-ff12fb70e837
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy:
  - uid: '[[phases/rex-whichkey-07-usability.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-02'
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

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE0OjEzOjQ1Ljg0OFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcGhhc2UwOC9zdW1tYXJ5LzEiLCJpbnRlbnQiOiI4ZTJmYmFkZDIzMjgyNjBhMjAxMmZhMmVkNjBhNDM2MjQyMzE2YjQ1OWExMTAwYjYwODA1ZmQ4NmU2MWJiNzFiIn0 -->
### 2026-11-02T14:13:45.848Z · summary · unspecified

{"acceptedScope":"Phase08: Popup polish","approvalSummary":["User requests popup timeout/default/color polish; Phase08 ownership/design and targeted35second native idle/color/navigation/dismissal verification are authorized under retained bounded technical authority.","All exact native child summaries bind current independent source approval and delivered target commit/tree; ancestry verified."],"artifactRefs":[".mtn-tmp/menu-polish/phase08-final/validate.json",".mtn-tmp/menu-polish/phase08-final/worktrees.txt",".mtn-tmp/menu-polish/runtime/combined-source-check-2/result.json"],"attempt":1,"coverage":{"REQ-018":["tasks/rex-whichkey-08-01-lifetime.md","tasks/rex-whichkey-08-04-docs.md"],"REQ-019":["tasks/rex-whichkey-08-02-defaults.md","tasks/rex-whichkey-08-04-docs.md"],"REQ-020":["tasks/rex-whichkey-08-03-theme.md","tasks/rex-whichkey-08-04-docs.md"]},"eventId":"rex-whichkey-p03-coordinator-a01/phase08/summary/1","files":["README.md","cmd/main.go","defaults.go","docs/acceptance.md","docs/appearance.md","docs/launch-contract.md","lua/rex_whichkey/defaults/init.lua","lua/rex_whichkey/defaults/panes.lua","native.go","popup.go","render.go","theme.go"],"kind":"summary","limitations":["No unit/race/TDD/general suites, destructive native actions, configured Lua callbacks, remote/Darwin trials or configuration reloads. One bounded user-authorized Linux/Rex idle35seconds/reduced-defaults/VT-color/navigation/dismissal/owned-cleanup check was observed; broader runtime compatibility remains unverified. Final documentation-only candidate could not repeat display after graphical app disconnected; prior actual UI observation is reused only for mechanically verified identical tracked runtime source, not claimed as a new-frame/full-tree final runtime pass.","Independent fresh contexts used the same authenticated model family, not cross-family review.","Prior transferred coordinator process status remains unverified; ownership transfer did not assert remote/local stoppage.","No global cleanup or push; retained worktrees/artifacts remain available."],"recordRef":"phases/rex-whichkey-08-popup-polish.md","records":["tasks/rex-whichkey-08-01-lifetime.md","tasks/rex-whichkey-08-02-defaults.md","tasks/rex-whichkey-08-03-theme.md","tasks/rex-whichkey-08-04-docs.md"],"requirements":["REQ-018","REQ-019","REQ-020"],"schemaVersion":1,"specHash":"sha256:baf78cecbd7f4271b318cae4cb0bccf51e850a035096b539d55b3e6a6e27802a","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"6090379edb0d43f7d56b01b81cc413ac9cc5fa5b","targetTree":"c98f2454d8e4d9288353b1dfd26e6944f71f82da","verification":[{"command":"Actual native Linux/Rex popup idle35seconds/reduced defaults/VT colors/Tabs/Escape/exit0/owned cleanup on identical tracked runtime source; documentation-only final candidate compiled, source equivalence verified","result":"pass"},{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./...; make build (combined delivered revision and refreshed CLI artifact)","result":"pass"},{"command":"git diff --check; native graph/schema validation; exact completed child scope and Git delivery identities/ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
