---
workflowKind: task
workflowLanguage: go
workflowOrder: 1
workflowRequirements:
  - REQ-023
workflowDocuments:
  - ref: designs/rex-whichkey-app-independent-popup.md
    hash: 'sha256:7215d04e0b71c6cc6a8a593372ef03cec57ae6aeb16907a81f1f9407b5781b4c'
workflowOwnership:
  paths:
    - launch.go
    - native.go
    - popup.go
    - README.md
    - docs/launch-contract.md
    - docs/acceptance.md
  sharedPaths: []
  resources: []
  interfaces:
    - native-display-action-binding-boundary
title: Decouple native popup display from app targeting
mtnCreateRetry:
  key: rex-whichkey-phase10-task01-v1
  intent: db2412d1c7cd97502b0371ba8175dd1f5c79d53e9098a48123938b8c3fe5ec57
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T16:14:38.525Z'
dateModified: '2026-11-02T16:28:54.799Z'
tags:
  - task
type: task
id: 8d222913-cf67-4dda-8b22-a2ef9b754608
projects:
  - '[[phases/rex-whichkey-10-app-independent-popup.md]]'
completedDate: '2026-11-02'
---
# Decouple native popup display from app-target validation

REQ-023; read approved designs/rex-whichkey-app-independent-popup.md. Own only launch.go, native.go, popup.go, README.md, docs/launch-contract.md, docs/acceptance.md. No cmd/options/mapping/config/schema/Lua VM/default changes, tests/gates/deps, .tasks or user configuration. Parent project/repo /home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main; read only canonical parent .tasks.

Implement genuine app-independent native launch, not just catch/swallow several-app error or choose first/last/focused/terminal owner. LaunchPopup must validate exact execution endpoint/session/block and syntax, decide native versus configured route from server action metadata BEFORE requiring GUI identity. Native display must not call session.list_clients or selected-client actions registry. Pass optional requested selector as trusted Go-owned native intent, not a fabricated connection generation. Native defaults already provide trusted immutable mappings and should render regardless of app/other-action availability.

Separate native resource liveness from app binding across actual popup lifecycle, wait watcher, focus, cancel and cleanup. Native idle/escape must remain usable with zero/several GUI apps and preserve no timeout/parent FIFO. Owned-only layer/files cleanup must remain possible when graphical clients unavailable and must precede action-target validation/dispatch. Keep wrong-endpoint read-only failure and exact captured session/block/resource identity checks. Existing configured decoded handoffs stay strict: same VM, prelaunch app/generation, original authenticated expiry, no native-policy bypass, retry or uncertain route fallback.

After native retained selection/confirmation is validated, cancellation returns without GUI lookup. Executable app selection then resolves optional exact selector or unique eligible live app, captures full generation, checks selected action compatibility (not every other default) and revalidates exact app generation/origin immediately before one dispatch. No caller/child-supplied action/args, stale reconnect, silent retargeting or uncertain retry. Ambiguous/unavailable target produces an ACTION-targeting refusal, not claim native menu cannot launch. Never execute if target selection or cleanup fails.

Docs must describe native display versus native action targeting versus configured launch binding truthfully, including deferred explicit-selector liveness and unknown/zero/several apps. Remove obsolete launch-time graphical-app prerequisite statements without erasing historical scope/limitations. No temporal pending-review checklists or retrospective runtime claims.

Leaf checks format/static/whitespace only; NO product execution/live API/GUI/config/remote probes, heavy/build/lint-backed compilation, tests/TDD/general/unit/race suites or installs. Coordinator owns exact combined Taskfile/Go compilation and bounded actual native UI checks: normal startup/nav/Escape, nonexistent exact selector startup/Escape, then one non-destructive selection with that nonexistent target must refuse after selection/cleanup, never dispatch. No worker integration/managed writes/agents or push. git -c core.hooksPath=/dev/null for scoped source commits; freeze clean head/tree. Full protocol-v1 build to payloadFile AND MTN_REX_RESULT_FILE (not pointer); validation array of actual passing {command,result:"pass"}; skipped gates in limitations. Final visible reply concise pointer only.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE2OjI4OjU0Ljc4OFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxMC10MDEtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiNjI5YzQxNThhYjc3YjNkNWVlZWEzNzMzYTBmNDlhY2ZkNTk5MGIxNzQzZmFiMjlmMTY5NGEyNjQ0MTUzOWRiNSJ9 -->
### 2026-11-02T16:28:54.788Z · summary · unspecified

{"acceptedScope":"Decouple native popup display from app targeting","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p10-t01-reviewer-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/startup-decoupling/rex-whichkey-p10-t01-builder-a01/integration/integration.json",".mtn-tmp/startup-decoupling/runtime/rex-whichkey-p10-t01-builder-a01/result.json"],"attempt":1,"eventId":"rex-whichkey-p10-t01-builder-a01/summary/1","files":["README.md","docs/acceptance.md","docs/launch-contract.md","launch.go","native.go","popup.go"],"kind":"summary","limitations":["Basic popup and deliberate nonexistent-client-selector popup startup/dismissal plus selected-action refusal after owned cleanup were actually checked. No arbitrary app selection, destructive dispatch or general runtime/remote-GUI certificate.","No project tests/TDD, destructive actions, configured Lua callbacks, remote/Darwin or configuration changes/reloads; fresh same-family review."],"recordRef":"tasks/rex-whichkey-10-01-app-independent-popup.md","requirements":["REQ-023"],"schemaVersion":1,"specHash":"sha256:1b007c064cc0f5f60819251a8d259450baa76641ac643f286ba116a11683046d","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"8a5f388444ae371a15ea89c4d54085d9f80180eb","targetTree":"bad2da66d266758ee52ac925dc9da3b8bd273565","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"},{"command":"task build; actual normal native popup and nonexistent-client-selector popup navigation/Escape/exit0/owned cleanup; nonexistent-target non-destructive selection refused only after UI cleanup without action dispatch","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
