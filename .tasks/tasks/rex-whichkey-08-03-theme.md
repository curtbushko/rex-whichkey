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
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T13:17:06.669Z'
dateModified: '2026-11-02T13:17:07.585Z'
tags:
  - task
type: task
id: dade01f8-73a6-4feb-97c2-e386134c09b4
projects:
  - '[[phases/rex-whichkey-08-popup-polish.md]]'
---
# Give keys, icons and labels distinct visible semantic colors

REQ-020. Own theme.go, render.go and terminal.go only. Read helios-design, its color/accessibility guidance, actual theme palette/decoder/render span wiring, and .mtn-tmp/menu-polish/neovim-which-key.nix. The user-linked example uses yellow icons plus named cyan/orange/green/blue category accents; it does not specify exact key/description RGB. Do not invent an exact hex-copy claim or a second Lua/Nix schema.

Implement a genuine icon semantic role in rendering instead of ThemePrimary for icons, and distinct readable key/label/group treatments. Default icon accent yellow, key accent cyan/blue, primary/group labels readable and visibly styled. Fix current all-plain output when background/palette replies are absent: emit appropriate terminal-native ANSI semantic fallback immediately without waiting for OSC, preserving unknown terminal background/default readable text. As real palette metadata arrives, use compatible readable colors/contrast rather than collapsing every role to foreground. No terminal palette assignments or guessed dark background. Honor NO_COLOR, explicit no-color, ASCII decoration handling, control sanitization, width/clipping and terminal reset/restoration. Never use color alone for eligibility or confirmation.

No new mapping fields, dependencies, generic theme framework, user config/settings changes or lifecycle/default changes. Builder formatting/static/source checks only, no heavy/test/TDD/product/GUI/API/config/remote trials. Fresh Go/TUI reviewer must read helios-design and trace actual Theme.Role -> render spans -> terminal writes, unknown/partial/full palette, late OSC and no-color paths, clipping/alignment and defensive sanitization. Coordinator compiles actual combined revision; final assembly will capture a real colored VT/HTML frame in Rex and navigation/dismissal. A color helper with no real caller or a plain text capture alone is insufficient color proof. No worker agents/integration/managed writes or push. Full protocol JSON transport and concise final pointer.
