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
  - ref: designs/rex-whichkey-native-startup.md
    hash: 'sha256:3c49a792bd2302bae6253174ab5d992dbf45cd9b557f53dd32e539bae8ce6f7c'
title: 'Phase07: Usability'
mtnCreateRetry:
  key: rex-whichkey-phase07-usability-v1
  intent: ae9f3baa27cca9ed342f10f90d5169d9dc0d5648285767f2f805a2e4258ddc07
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-01T14:52:26.658Z'
dateModified: '2026-11-02T01:36:39.008Z'
tags:
  - task
type: task
id: a6773f97-4e50-4960-ae86-9c96af731a54
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy:
  - uid: '[[phases/rex-whichkey-06-packaging-release.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-01'
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


## Final-audit refinement
Require durable accurate user-facing delivery-status wording: actual source review/Go/Nix compilation must not be reported as unexecuted, and unexecuted tests/runtime trials must not be called passes. Final README status must not self-declare pending review after completion. This is scoped to the existing usage/assembly task and does not reopen endpoint/direct-launch implementation.


## User-reported client-selection correction
Normal startup inside a Rex terminal with an attached graphical app must not require REX_CLIENT solely because CLI/automation connections are also attached. Select the live graphical app(s) through authoritative native classification and exact captured-session membership. Distinguish no-app/multiple-app/invalid-selected-client errors; keep all origin/generation safeguards and no runtime trial policy. Actual user connection counts remain unobserved; this is a verified source-selection mismatch, not a claim that the app is absent.


## Actual runtime launch correction
User now explicitly authorizes bounded real binary launch/navigation/dismissal in the existing Rex session and rejects manual setup changes to a working Rex installation. Actual invocation exited1 because custom whichkey actions were absent. REQ-016 basic default startup must not require those registrations: implement a self-contained Go/native floating popup while preserving registered/explicit original-Lua-VM routes. See designs/rex-whichkey-native-startup.md. Do not modify/reload working Rex configuration; unit/TDD/general/remote/destructive trials remain prohibited. Coordinator must record actual normal floating-popup display/navigation/Escape dismissal and owned cleanup before completing this correction.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDAxOjM2OjM4Ljk5NloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcGhhc2UwNy9zdW1tYXJ5LzQiLCJpbnRlbnQiOiJhYTA2NzU5M2NjYzRjNzBjMjY2OWEwNTU5NDgzNDIyN2MxYmQwNjdhNjZhYmNhZGRmMTliY2ViOWVjZmYwZDUxIn0 -->
### 2026-11-02T01:36:38.996Z · summary · unspecified

{"acceptedScope":"Phase07: Usability","approvalSummary":["User explicitly authorizes Phase07 usability, task decomposition/implementation and per-phase commits under retained bounded technical authority; native platform discovery design is approved.","All exact native child summaries bind current independent source approval and delivered target commit/tree; ancestry verified."],"artifactRefs":[".mtn-tmp/live-launch/rex-whichkey-p07-t05-builder-a01/result.json",".mtn-tmp/usability/phase07-final-r4/validate.json",".mtn-tmp/usability/phase07-final-r4/worktrees.txt"],"attempt":4,"coverage":{"REQ-016":["tasks/rex-whichkey-07-02-direct-launch.md","tasks/rex-whichkey-07-03-usage-assembly.md","tasks/rex-whichkey-07-04-app-client-selection.md","tasks/rex-whichkey-07-05-native-default-popup.md"],"REQ-017":["tasks/rex-whichkey-07-01-endpoint-policy.md","tasks/rex-whichkey-07-02-direct-launch.md","tasks/rex-whichkey-07-03-usage-assembly.md","tasks/rex-whichkey-07-05-native-default-popup.md"]},"eventId":"rex-whichkey-p03-coordinator-a01/phase07/summary/4","files":["README.md","cmd/main.go","defaults.go","docs/acceptance.md","docs/launch-contract.md","docs/remote-compatibility.md","launch.go","lua/rex_whichkey.lua","native.go","popup.go","rex/endpoint.go","rex/layer.go"],"kind":"summary","limitations":["No unit/race/TDD/general suites, destructive native actions, configured Lua callbacks, remote/Darwin trials or configuration reloads. One bounded user-authorized normal Linux/Rex launch/navigation/dismissal/owned cleanup was actually observed; broader runtime compatibility remains unverified.","Independent fresh contexts used the same authenticated model family, not cross-family review.","Prior transferred coordinator process status remains unverified; ownership transfer did not assert remote/local stoppage.","No global cleanup or push; retained worktrees/artifacts remain available."],"recordRef":"phases/rex-whichkey-07-usability.md","records":["tasks/rex-whichkey-07-01-endpoint-policy.md","tasks/rex-whichkey-07-02-direct-launch.md","tasks/rex-whichkey-07-03-usage-assembly.md","tasks/rex-whichkey-07-04-app-client-selection.md","tasks/rex-whichkey-07-05-native-default-popup.md"],"requirements":["REQ-016","REQ-017"],"schemaVersion":1,"specHash":"sha256:b225dc6a390dcfbc34135af5f1feb16413b123458ff07b9b08bef515a98540bb","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"9f9b903e7cb8a5ed526a54a506aee408e8f11fd0","targetTree":"c124f58cbd1fa10c9f65795c786928d14b5c14fe","verification":[{"command":"Actual user-authorized normal native Linux/Rex floating popup, Tabs navigation, Escape dismissal, parent exit0 and exact owned popup cleanup","result":"pass"},{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./...; make build (combined delivered revision and refreshed CLI artifact)","result":"pass"},{"command":"git diff --check; native graph/schema validation; exact completed child scope and Git delivery identities/ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
