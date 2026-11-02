---
workflowKind: task
workflowLanguage: go
workflowOrder: 5
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
workflowOwnership:
  paths:
    - native.go
    - defaults.go
    - launch.go
    - popup.go
    - cmd/main.go
    - rex/layer.go
    - README.md
    - docs/launch-contract.md
    - docs/acceptance.md
  sharedPaths: []
  resources: []
  interfaces:
    - self-contained-native-popup
    - configured-native-launch-selection
title: Launch a real default popup without custom action setup
mtnCreateRetry:
  key: rex-whichkey-native-default-popup-v1
  intent: 0d663fa6658689a2ac21c4868a82e2068b47ad810f60bef73b284196701055c0
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T01:18:41.735Z'
dateModified: '2026-11-02T01:35:28.479Z'
tags:
  - task
type: task
id: 3502034b-aa1e-4c94-a8ad-f183bcbda81a
projects:
  - '[[phases/rex-whichkey-07-usability.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-07-04-app-client-selection.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-01'
---
# Launch the basic popup without custom Rex action registration

REQ-016 and REQ-017. Read approved supplemental native-startup design and .mtn-tmp/live-launch/ evidence. Actual normal binary invocation failed, although one graphical app is connected, because server/client registries lack whichkey/whichkey_complete. Renderer diagnostic actually displayed/navigated/dismissed successfully, but this is not a normal popup pass. User explicitly requests running the actual binary here, not more manual commands or changes to the working Rex setup. Do not question the user's working session or require invented exported REX_CLIENT/REX_SERVER.

## Owned source

native.go (new flat root helper), defaults.go (new Go data baseline), launch.go, popup.go, cmd/main.go, rex/layer.go, README.md, docs/launch-contract.md, docs/acceptance.md. Existing lifecycle/default types/helpers may be reused/refactored within these literal paths. No other source, dependency, mapping schema, Lua/dispatcher, Taskfile/lint/hook, user init/config or managed-record edits. Existing configured original-Lua-VM route must stay compatible.

## Acceptance

1. With running Rex and installed binary, bare rex-whichkey opens the REAL floating popup without custom whichkey/completion registration. Use the delivered native endpoint resolver and exact captured session/block/live app/generation. No demo/usage-only/standalone-terminal substitute, host/current-client guess, new daemon/framework/VM, remote bootstrap or user config edit/reload. A short-lived native CLI controller may wait for its child; callback-controller original prompt-return contract remains unchanged.
2. Preserve optional configured-Lua support: if the default launch action is genuinely registered, keep that original-VM path. A public explicit --action must remain a configured-action request; missing registration is actionable. Use actual read-only action metadata, bound output/context and exact server/client. Never switch to native mode after an activation was attempted/uncertain. Distinguish unregistered default from failed configured launch, not an exception-matching retry.
3. Define genuine Go-owned built-in Mapping/Config data using existing canonical types and native actions. Match existing defaults' keys/intent (sessions/tabs/panes/application/host inventory), destructive confirmation and provider semantics. Read actual .mtn-tmp/live-launch/client-actions.json/native-action-names.json native contracts and existing Lua fragments for behavior; do not invent action names/args or a new schema. No Lua callback reference may be executed by Go. Configured callbacks remain in the original VM; native default actions execute through explicitly scoped Rex native activation.
4. Reuse/shared lifecycle with existing controller/PTY child where practical. Keep private bounded transport and immutable retained parent mapping snapshot; selection returns only identity/confirmation, no action/args replacement. At most one selection/native dispatch. Validate origin/generation before every mutation; close the exact owned popup, restore terminal and unlink owned private files BEFORE dispatch, since a selected pane/tab close may kill its caller. Dismissal sends no action. Never broad-delete layer/window/session, follow labels, retry uncertain creation/action, leak capability/credentials or leave unbounded processes.
5. Keep existing menu/prepare/controller/popup modes, explicit overrides, Go native socket discovery, registered Lua default/custom action behavior and user setup semantics compatible. Native controller must be an actual production caller, not duplicate dormant symbols. No hard-coded platform sockets or mandatory missing environment variables.
6. Improve diagnostics safely: stage-specific actionable errors for actual known failures, not the generic discarded reason list; never print arbitrary configured callback errors, private capabilities or credentials. Docs distinguish self-contained basic defaults from optional Lua customization and actual tested versus untested behavior. Do not claim successful popup display on RPC acceptance alone.

## Verification authority

Builder may format and statically inspect only, no heavy slot or product probes; coordinator owns actual compilation and the newly user-authorized bounded live launch/navigation/Escape-dismissal in this existing Rex session. General project tests/TDD/race suites, unrelated experiments, remote/Darwin probes, configuration changes/reloads, installs and destructive menu action tests remain prohibited. Live schema/registry evidence is already available in .mtn-tmp/live-launch/new-layer-api.json, block-types.json, actions.json and client-actions.json. If a needed actual API contract is unavailable, report exact need; do not guess.

A fresh independent reviewer must inspect both configured/native routing, default-action data, retained snapshot/confirmation/origin/privacy/lifetime and owned-only cleanup, help/docs and real caller reachability. Coordinator must compile the exact combined candidate then actually launch that candidate, capture the floating menu/submenu, dismiss via Escape, verify zero parent exit and exact owned-popup cleanup BEFORE native task completion. A source-only or renderer-only check is not this runtime gate. Fix observed failures as new reviewed attempts; no fake pass or broad cleanup. Implementation commits use git -c core.hooksPath=/dev/null, no push or other agents.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDAxOjM1OjI4LjQ2N1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNy10MDUtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiNjNhZmU4OGRjM2VhM2ZmNDc1NmQyOTBkYTlkOWFjMGI5YjJhNzRiMTc2MGEyMTIxYmVhYzEzOWMzNDkyMGIzNSJ9 -->
### 2026-11-02T01:35:28.467Z · summary · unspecified

{"acceptedScope":"Launch a real default popup without custom action setup","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p07-t05-reviewer-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/usability/rex-whichkey-p07-t05-builder-a01/integration/integration.json",".mtn-tmp/live-launch/rex-whichkey-p07-t05-builder-a01/result.json"],"attempt":1,"eventId":"rex-whichkey-p07-t05-builder-a01/summary/1","files":["README.md","cmd/main.go","defaults.go","docs/acceptance.md","docs/launch-contract.md","launch.go","native.go","popup.go","rex/layer.go"],"kind":"summary","limitations":["Actual user-authorized Linux/Rex normal popup launch/navigation/dismissal/exit0/owned cleanup observed; not a general runtime or platform certificate.","No project tests/TDD, destructive actions, configured Lua callbacks, remote/Darwin or configuration changes/reloads; fresh same-family review."],"recordRef":"tasks/rex-whichkey-07-05-native-default-popup.md","requirements":["REQ-016","REQ-017"],"schemaVersion":1,"specHash":"sha256:230c39ea9b0d36eeda17d17502751e53ed9c3816a0ce3f7cc86b1be0d83ca811","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"9f9b903e7cb8a5ed526a54a506aee408e8f11fd0","targetTree":"c124f58cbd1fa10c9f65795c786928d14b5c14fe","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"},{"command":"task build; actual bare native popup launch, Tabs navigation, Escape dismissal, exit0 and exact owned-popup cleanup in existing Linux Rex session","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
