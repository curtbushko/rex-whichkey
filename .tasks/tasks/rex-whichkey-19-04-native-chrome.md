---
identity:
  kind: task
  taskID: rex-whichkey-19-04-native-chrome
  uuid: 934a4988-4ad8-47b5-ac4b-edf7fc7c6256
structure:
  order: 4
scope:
  requirements:
    - REQ-026
    - REQ-014
    - REQ-015
  documents:
    - ref: designs/rex-whichkey-lua-float-launch.md
      hash: 'sha256:a3657116206eb455b158b9727099b7c3859882d18f759171205eb830dda163c8'
implementation:
  language: lua
  ownership:
    paths:
      - lua/rex_whichkey.lua
      - nix/package.nix
      - docs/appearance.md
      - docs/acceptance.md
    sharedPaths: []
    resources: []
    interfaces:
      - P19-SHARED-LUA-ENTRY
      - P19-NATIVE-CHROME
  gates:
    - ID: whitespace
      command: git diff --check
      required: true
lifecycle:
  state: blocked
  sourceEventID: rex-whichkey-p19-t04/chrome-blocker/1
  changedAt: '2026-11-07T04:48:44.931Z'
title: 'Phase19.4: Required native float chrome via verified Rex capability'
mtnCreateRetry:
  key: rex-whichkey-p19/task-create/native-chrome/1
  intent: c650c07508d11dc05b7c59b8bc32fdc5a3efa821338a0fb12be741cd713c64bc
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-07T04:48:27.141Z'
dateModified: '2026-11-07T04:48:45.558Z'
tags:
  - task
type: task
id: 6dadf60f-878a-4332-95c0-fe56f0cf3c79
projects:
  - '[[phases/rex-whichkey-19-lua-float-launch.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-19-03-source-run-delivery.md]]'
    reltype: FINISHTOSTART
---
REQUIRED AND BLOCKED by P19-NATIVE-CHROME. Current installed session.new_layer accepts only session_id/window_id/bounds/layout/focus; terminal creator options expose command/shell/cwd/size/theme/exit and related terminal data, not native titlebar/border/corners/shadow. options.theme is colors, not native chrome. No source implementation, ignored configuration knobs, guessed fields, ANSI substitutes, global settings/user init changes or external Rex edits are authorized as a workaround. Resolve only through a genuinely documented supported per-float capability available within authorized scope, or an explicit user change to this requirement; obtaining/updating Rex externally is not implicitly authorized. Once unblocked, extend the real Lua launch implementation and ordinary asset packaging to control all four native chrome aspects; source/schema review plus actual per-float visual runtime confirmation are required, independent of Go terminal palette/header evidence. Task19.3 is a serialization prerequisite solely for shared Lua/package/docs handoff, not a barrier into supported tasks. Until then retain criterion2 and REQ-026 as unmet and preserve old Phase11/14/15 blocker history unchanged; no implicit deferral, retrospective closure or full phase completion.
<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA3VDA0OjQ4OjQ0LjE0OVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYmxvY2tlciIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAxOS10MDQvY2hyb21lLWJsb2NrZXIvMSIsImludGVudCI6IjY1NmNjNWY5YzljY2I1MmJiNjdkY2NiYTM4Nzc3MGNlNTY5ZThkZWZmZjcxNWFhMGU1YTk3YzU2ZTI5ZjA5ZmUifQ -->
### 2026-11-07T04:48:44.149Z · blocker · unspecified

{"attempt":1,"decisionID":"REXWHICHKEY-20261106-234819","eventID":"rex-whichkey-p19-t04/chrome-blocker/1","evidence":[".mtn-tmp/p19/rex-api-list.json",".mtn-tmp/p19/rex-block-types.json","https://discord.com/channels/1546797889009881158/1566968591062929548"],"kind":"blocker","reason":"Installed Rex new_layer/terminal creation and published methods expose no per-float native titlebar/border/corners/shadow controls. Terminal theme is colors, not chrome; active Rex feature request confirms unresolved dependency.","recordRef":"tasks/rex-whichkey-19-04-native-chrome.md","requiredDecision":"Provide an actual supported Rex chrome capability/version within authorized access, or explicitly defer/change chrome requirement. Default chrome is not silently accepted and guessed/ignored flags are forbidden.","schemaVersion":1,"specHash":"sha256:b61c8b6635362677eaba352f4b01d34da26528a75031634721cebde8d9de6431"}
<!-- mtn:end-entry:v1 -->
