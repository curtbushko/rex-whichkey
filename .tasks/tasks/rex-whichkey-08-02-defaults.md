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
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T13:17:04.837Z'
dateModified: '2026-11-02T13:17:05.763Z'
tags:
  - task
type: task
id: 65b044ae-abca-4ee1-a70c-b218666409ba
projects:
  - '[[phases/rex-whichkey-08-popup-polish.md]]'
---
# Remove Focus and automatic known-hosts from defaults

REQ-019. Own defaults.go, lua/rex_whichkey/defaults/panes.lua and lua/rex_whichkey/defaults/init.lua only. Remove the f/Focus group and fh/fj/fk/fl default mappings in native Go and composed Lua defaults. Remove automatic h/Rex-known-hosts provider construction/import. Keep session/tab/app/split/zoom/close defaults, order and confirmation behavior. Do not remove origin focus restoration or underlying user-defined focus support, and do not delete optional known-host provider implementation or the optional hosts fragment: users may still explicitly configure it. No default marker means no inventory lookup via the real terminal resolver.

Read actual default generators/aggregation and canonical Mapping rules, not an independent Lua schema. Keep Go/Lua default footprint consistent and no default action references dead/invented symbols. Do not write README/docs here (assembly owns them), native.go lifetime, theme/render or unrelated source. No Taskfile/lint/hook/dependency/user config/managed edits. Builder formats Go/Lua/static/whitespace only, no tests/TDD/heavy builds/product/GUI/API/remote trials/config reloads/push. Fresh mixed source reviewer must explicitly cover Go AND Lua default outputs and retained optional-provider/user-spec behavior. Coordinator compiles combined candidate; final assembly verifies live root lacks removed groups and still navigates Tabs/dismisses. No worker integration/agents. Return full protocol JSON to designated files and concise visible pointer.
