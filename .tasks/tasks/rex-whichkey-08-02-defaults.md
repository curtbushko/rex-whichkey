---
workflowKind: task
workflowLanguage: go
workflowOrder: 2
workflowRequirements:
  - REQ-019
workflowDocuments:
  - ref: designs/rex-whichkey-popup-polish.md
    hash: 'sha256:ce42e479546a1181ce61138f0f5d37c88e5ed18ae3b6c39987fca90d89aee8f0'
workflowOwnership:
  paths:
    - defaults.go
    - lua/rex_whichkey/defaults/panes.lua
    - lua/rex_whichkey/defaults/init.lua
  sharedPaths: []
  resources: []
  interfaces:
    - default-menu-footprint
title: Remove Focus and known-hosts from defaults
mtnCreateRetry:
  key: rex-whichkey-phase08-task-2-v1
  intent: 2500ac0d591c766cd7194777dbe76f3b03dfa235a50ca308cc3989a51bf44f67
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T13:17:04.837Z'
dateModified: '2026-11-02T13:35:32.973Z'
tags:
  - task
type: task
id: 65b044ae-abca-4ee1-a70c-b218666409ba
projects:
  - '[[phases/rex-whichkey-08-popup-polish.md]]'
completedDate: '2026-11-02'
---
# Remove Focus and automatic known-hosts from defaults

REQ-019. Own defaults.go, lua/rex_whichkey/defaults/panes.lua and lua/rex_whichkey/defaults/init.lua only. Remove the f/Focus group and fh/fj/fk/fl default mappings in native Go and composed Lua defaults. Remove automatic h/Rex-known-hosts provider construction/import. Keep session/tab/app/split/zoom/close defaults, order and confirmation behavior. Do not remove origin focus restoration or underlying user-defined focus support, and do not delete optional known-host provider implementation or the optional hosts fragment: users may still explicitly configure it. No default marker means no inventory lookup via the real terminal resolver.

Read actual default generators/aggregation and canonical Mapping rules, not an independent Lua schema. Keep Go/Lua default footprint consistent and no default action references dead/invented symbols. Do not write README/docs here (assembly owns them), native.go lifetime, theme/render or unrelated source. No Taskfile/lint/hook/dependency/user config/managed edits. Builder formats Go/Lua/static/whitespace only, no tests/TDD/heavy builds/product/GUI/API/remote trials/config reloads/push. Fresh mixed source reviewer must explicitly cover Go AND Lua default outputs and retained optional-provider/user-spec behavior. Coordinator compiles combined candidate; final assembly verifies live root lacks removed groups and still navigates Tabs/dismisses. No worker integration/agents. Return full protocol JSON to designated files and concise visible pointer.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDEzOjM1OjMyLjk2MloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwOC10MDItYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiYmRiNjA0MDFiZjQ2YzU4M2IxYzZlZjZkNDg5OWU1MmZhMmYwNDJmMzdjM2EwMjk5ZDUzYTFhMzBkYTU0MWMxYyJ9 -->
### 2026-11-02T13:35:32.962Z · summary · unspecified

{"acceptedScope":"Remove Focus and known-hosts from defaults","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p08-t02-reviewer-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/menu-polish/rex-whichkey-p08-t02-builder-a01/integration/integration.json"],"attempt":1,"eventId":"rex-whichkey-p08-t02-builder-a01/summary/1","files":["defaults.go","lua/rex_whichkey/defaults/init.lua","lua/rex_whichkey/defaults/panes.lua"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-08-02-defaults.md","requirements":["REQ-019"],"schemaVersion":1,"specHash":"sha256:5b38acfb7f05dc4362b40e8d94198858da08423ba27e249df031dc3c574d1313","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"3ad78d627d6355c17731da53de6c710736f278e6","targetTree":"7a9972f5157393e53e20f42a340b16085702e7df","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
