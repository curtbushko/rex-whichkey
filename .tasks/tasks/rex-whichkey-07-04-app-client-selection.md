---
workflowKind: task
workflowLanguage: go
workflowOrder: 4
workflowRequirements:
  - REQ-016
workflowDocuments:
  - ref: prds/rex-whichkey-usability.md
    hash: 'sha256:c2df9b9404ce7f2d649eed40d5c219f7098e9389c6407159b6ae35eddb7746c5'
  - ref: designs/rex-whichkey-usability.md
    hash: 'sha256:c8553dd4ac48fcd34a43d7cb3cd842bb18c65310d0fc6133cf1b6e3f196f1b8a'
workflowOwnership:
  paths:
    - popup.go
    - launch.go
    - README.md
    - docs/launch-contract.md
    - docs/remote-compatibility.md
  sharedPaths: []
  resources: []
  interfaces:
    - origin-app-client-selection
title: Select the attached graphical Rex app automatically
mtnCreateRetry:
  key: rex-whichkey-usability-app-client-selection-v1
  intent: 52b349189bc14ffdb98f2b7c757341bfbfc503d336d60011fb875e9c89cf639b
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-01T19:52:47.641Z'
dateModified: '2026-11-01T20:01:20.280Z'
tags:
  - task
type: task
id: 2136a023-e84a-48ba-bc95-60453d2e9368
projects:
  - '[[phases/rex-whichkey-07-usability.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-07-02-direct-launch.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-01'
---
# Select the attached graphical Rex app, not automation connections

## User report and verified facts

User runs inside a Rex terminal and communicates through an attached graphical Rex client, yet no-argument startup reports zero/several eligible clients. User supplied full REX_SESSION, REX_BLOCK and matching REX_BIN_DIR; REX_CLIENT and REX_SERVER are absent. Do not ask user to configure REX_CLIENT for the ordinary one-app case. Do not persist or repeat unrelated token environment values.

Verified current popup.go popupOriginConnection decodes every session.list_clients record and counts all nonlost connections without app classification. Installed matching Rex `client ls --help` explicitly says command-line clients are hidden in normal text but included in JSON's full server response. Existing server-side configuration source uses client.list with kinds={"app"}; use authoritative native classification/filtering, not invented labels/platform heuristics. Actual user-side counts/response records have NOT been inspected; do not assert missing GUI or exact root-cause cardinality beyond verified source mismatch.

## Requirements and ownership

REQ-016 (normal no-argument popup) plus retained REQ-017 origin/endpoint scope. Literal owned files: popup.go, launch.go, README.md, docs/launch-contract.md, docs/remote-compatibility.md. Expand only if the coordinator approves a concrete source need. No dispatcher/Lua VM, dependency, mapping-schema, Taskfile/lint/hook, managed-record or user init.lua changes.

## Acceptance

1. Select only actual graphical application clients attached to the exact captured origin session. Use documented native client classification/filtering (e.g. server client.list kinds app) and intersect with exact session attachments if native filtered scoping is unavailable. Confirm actual request/response contract from static source/docs/help; do not invent field names or treat label/name/platform as client kind. Global app list alone is insufficient: preserve exact session membership.
2. Ignore CLI, automation and other non-app connections for automatic candidate cardinality and explicit selector eligibility. A single attached live app succeeds even when other CLI connections exist. Do not force the user to export REX_CLIENT or bypass validation; absent REX_CLIENT is normal. Still reject genuinely ambiguous multiple graphical apps, zero live graphical apps, stale/wrong explicit IDs, malformed/duplicate identity data and connection-generation changes. Explicit requested client must identify a live attached app.
3. Keep full typed IDs, selected endpoint, connection generation, original session/block placement and preflight-before-mutation. Preserve fail-closed no-focused/current-client fallback, one-shot/completion/config/expiry/confirmation/original Lua-VM safeguards, bounded child handling and no uncertain retries. Reuse the corrected resolver for launch/controller checks consistently rather than only special-casing the initial root route.
4. Distinguish errors for no attached live app, several graphical apps, and invalid/not-attached supplied selector. State the actionable cause without leaking credentials/capabilities/private payloads. Do not present all cases as zero-or-several; do not blame absence of REX_CLIENT. Do not report a popup shown merely on activation acknowledgement.
5. Update affected user/contract documentation with graphical-app classification and scoped membership, no mandatory client ID in the normal single-app case, and honest source/compile versus runtime status. Preserve native socket defaults/overrides and all other existing behavior.

## Verification and delivery

No project tests/TDD/experiments/product execution/GUI/live API/remote trials/config reloads/dependency installation. Read-only published docs/source and installed CLI help are allowed; actual product client enumeration/activation by the agent remains prohibited under retained policy. Builder static checks/formatting only; coordinator heavy compilation after fresh independent complete Go/source-doc review. Existing tests/gates must not change. Compile combined candidate and rebuild development artifact before finalization. No source/managed writes outside owned worktree, push, other agents or integration by worker. Record any genuinely unavailable contract precisely, never guess it or repeat a routine approval request. Return full protocol JSON to payloadFile and MTN_REX_RESULT_FILE, concise visible pointer only.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDIwOjAxOjIwLjI2OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNy10MDQtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiMDFlN2NkYzMyMzdlODgzZTU3Mzk5MDIwZWI2NjI3OTIyMmI4NTYyYjM5ZDVjN2FjMjI5Mzc3MGVhNWY5NDExMCJ9 -->
### 2026-11-01T20:01:20.269Z · summary · unspecified

{"acceptedScope":"Select the attached graphical Rex app automatically","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p07-t04-reviewer-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/usability/rex-whichkey-p07-t04-builder-a01/integration/integration.json"],"attempt":1,"eventId":"rex-whichkey-p07-t04-builder-a01/summary/1","files":["README.md","docs/launch-contract.md","docs/remote-compatibility.md","launch.go","popup.go"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-07-04-app-client-selection.md","requirements":["REQ-016"],"schemaVersion":1,"specHash":"sha256:e828a2f445dd277935f112adc9e12bdeefd2ea6d2688003a9a9785a3f4a8d93e","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"609bc2ad09d91d395a0826aa50a89b7c574b699b","targetTree":"d740b34194913d948abbeab7da3930cff8303933","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
