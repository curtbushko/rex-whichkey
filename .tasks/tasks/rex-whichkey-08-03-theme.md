---
workflowKind: task
workflowLanguage: go
workflowOrder: 3
workflowRequirements:
  - REQ-020
workflowDocuments:
  - ref: designs/rex-whichkey-popup-polish.md
    hash: 'sha256:ce42e479546a1181ce61138f0f5d37c88e5ed18ae3b6c39987fca90d89aee8f0'
workflowOwnership:
  paths:
    - theme.go
    - render.go
    - terminal.go
  sharedPaths: []
  resources: []
  interfaces:
    - semantic-menu-color-roles
title: Color keys icons and labels with usable fallback
mtnCreateRetry:
  key: rex-whichkey-phase08-task-3-v1
  intent: 87d283237c5fd4f6581eb67d82a874a120be8e7216bdd969cb47258192c98fa4
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T13:17:06.669Z'
dateModified: '2026-11-02T13:40:14.706Z'
tags:
  - task
type: task
id: dade01f8-73a6-4feb-97c2-e386134c09b4
projects:
  - '[[phases/rex-whichkey-08-popup-polish.md]]'
completedDate: '2026-11-02'
---
# Give keys, icons and labels distinct visible semantic colors

REQ-020. Own theme.go, render.go and terminal.go only. Read helios-design, its color/accessibility guidance, actual theme palette/decoder/render span wiring, and .mtn-tmp/menu-polish/neovim-which-key.nix. The user-linked example uses yellow icons plus named cyan/orange/green/blue category accents; it does not specify exact key/description RGB. Do not invent an exact hex-copy claim or a second Lua/Nix schema.

Implement a genuine icon semantic role in rendering instead of ThemePrimary for icons, and distinct readable key/label/group treatments. Default icon accent yellow, key accent cyan/blue, primary/group labels readable and visibly styled. Fix current all-plain output when background/palette replies are absent: emit appropriate terminal-native ANSI semantic fallback immediately without waiting for OSC, preserving unknown terminal background/default readable text. As real palette metadata arrives, use compatible readable colors/contrast rather than collapsing every role to foreground. No terminal palette assignments or guessed dark background. Honor NO_COLOR, explicit no-color, ASCII decoration handling, control sanitization, width/clipping and terminal reset/restoration. Never use color alone for eligibility or confirmation.

No new mapping fields, dependencies, generic theme framework, user config/settings changes or lifecycle/default changes. Builder formatting/static/source checks only, no heavy/test/TDD/product/GUI/API/config/remote trials. Fresh Go/TUI reviewer must read helios-design and trace actual Theme.Role -> render spans -> terminal writes, unknown/partial/full palette, late OSC and no-color paths, clipping/alignment and defensive sanitization. Coordinator compiles actual combined revision; final assembly will capture a real colored VT/HTML frame in Rex and navigation/dismissal. A color helper with no real caller or a plain text capture alone is insufficient color proof. No worker agents/integration/managed writes or push. Full protocol JSON transport and concise final pointer.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDEzOjQwOjE0LjY5NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwOC10MDMtYnVpbGRlci1hMDEvc3VtbWFyeS8xIiwiaW50ZW50IjoiMDY5OTljOWM5ZDRhMjAxODY2MzZiYWYwNzI2YjA3OWE2YTg3ODgzOWMyNDIzNGE1NjEyOGY0NGM1NDkxNjEyZCJ9 -->
### 2026-11-02T13:40:14.695Z · summary · unspecified

{"acceptedScope":"Color keys icons and labels with usable fallback","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p08-t03-reviewer-a01/review/1 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/menu-polish/rex-whichkey-p08-t03-builder-a01/integration/integration.json"],"attempt":1,"eventId":"rex-whichkey-p08-t03-builder-a01/summary/1","files":["render.go","theme.go"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-08-03-theme.md","requirements":["REQ-020"],"schemaVersion":1,"specHash":"sha256:a495094e2fb1c8a0638831f1b057f6f9c5594574a3621daa689067158d1db016","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"3e9a1fd19a6a930fac293966d0ee7b4114aab2dc","targetTree":"0550e40e9efa97ce41fadd81e13913e041c3be31","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
