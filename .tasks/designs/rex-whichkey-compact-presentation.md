---
workflowKind: design
title: Compact popup presentation and native chrome limitation
mtnCreateRetry:
  key: rex-whichkey-compact-design-v1
  intent: aec157a936c7f0f53dab95f1c4d12e671b5141d84ad62803d5d2fdef5c0f7cb9
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T17:33:32.899Z'
dateModified: '2026-11-02T17:33:38.247Z'
tags:
  - task
type: task
id: dd070c0e-d54b-4f9a-9207-cbbd3cb4fe74
completedDate: '2026-11-02'
---
# Compact popup presentation

## User requirements and authority
User requests removal of Page1/1, content-fitting popup height (font/rows considered), hiding Rex pane/tab/window controls on popup, and title/key hints separated from menu as status-like borders. Existing delegated bounded implementation authority applies; no global settings/init reload or remote/Darwin/test-suite authority is inferred.

REQ-024: Remove the single-page Page1/1 indicator and redundant footer row. Preserve real multi-page navigation/indicator only when needed. Put title/breadcrumbs and concise key hints in visually separated full-width top/bottom status-like borders, with no distracting literal Keys: prefix in menu content. Preserve semantic key/icon/label colors, NO_COLOR/ASCII, Unicode grapheme safety, confirmation warnings and readable tiny-viewport degradation. Border/status rows yield to content when necessary.

REQ-025: The actual owned floating layer must shrink/grow in height to fit current visible menu content/status rows, including submenu changes and terminal/font-driven cell geometry changes. It is not enough to center a small text frame inside the old60percent layer. Use measured terminal cells/available geometry, not guessed fixed font pixels; cap to available maximum, preserve position anchoring and pagination. Native display remains app-independent; configured Lua expiry/action binding unchanged. Avoid endless resize feedback, fixed repaint timer, new daemon or global font/terminal setting changes.

REQ-026: Hide native Rex pane/tab/window controls for this popup only. Do not fake it by changing global client settings, shrinking content over controls, or sending undocumented flags. This accepted request is currently BLOCKED: installed session.new_layer/resize_layer schemas have only bounds/layout/focus and shell options_schema additionalProperties=false has no chrome flag; current connected app registry exposes no control/header/toolbar/chrome action, and published Lua layout docs expose only block flavor/options. Evidence .mtn-tmp/popup-presentation/{new-layer-api.json,resize-layer-api.json,terminal-creator.json,client-actions.json,lua-reference.md}. No documented per-popup capability was found. Keep requirement/task open, request a supported Rex capability/update or explicit user deferral, never mark phase/project fully complete while unresolved.

## Accepted bounded technical direction
Existing flat Go package/config schema remain. Renderer exposes MenuFrame.PreferredHeight int: natural terminal-row height for current menu and current width INCLUDING all status/border/padding rows, independent of current height clipping and pagination. Current frame Height/Lines remain clipped. Geometry consumer uses that exact contract, not duplicate menu measurement.

Popup child may publish bounded data-only preferred-row hints through existing private owned transport, atomically and only on change. Parent alone owns Rex Layer.Resize and authenticates exact owned layer/block membership. Scope hints can influence only clamped owned geometry, never mapping/selection/action/capability or other resources. Use actual current terminal rows/cell geometry for feedback, cap to original max rectangle, skip invalid/zero/transient measurements, preserve top/bottom/center/side anchoring, converge without a resize storm. Support native/configured lifetime without weakening either. Existing resource watcher can bound geometry refresh; do not add a service. Cleanup includes every new owned temporary/hint file; never unlink unrelated files or follow symlinks. No app lookup for native geometry.

## Delivery/verification
Task1 renderer producer; task2 geometry consumer depends on1; task3 native chrome remains explicitly blocked without invented implementation; task4 docs/assembly depends on1/2 only and documents chrome limitation, not fulfillment. Fresh independent source review and exact combined Go/Taskfile compilation per task. Leaf agents formatting/static/whitespace only, no product/GUI/API experiments/build/test suites. Coordinator runs actual Linux Rex normal menu/submenu/back/Escape with recorded native layer bounds, terminal rows/cell geometry, full VT/HTML/status frames, compact/stable root/submenu sizing and exact owned cleanup; preserve prior app-independent startup even with multiple apps. Do not change fonts/user settings or invoke native/destructive mapping actions to certify sizing. Test tiny/multipage/NO_COLOR/ASCII paths by source review under existing test prohibition, distinguish unexecuted cases. Remote app chrome cannot be certified from server captures. No push/global cleanup. Partial delivery must state blocked REQ-026 and must not publish full phase/project completion.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE3OjMzOjM3LjMzNloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9wcmVzZW50YXRpb24vZGVzaWduLWFwcHJvdmFsLzEiLCJpbnRlbnQiOiJiNzQzM2EzMTg0ZjdjZDRiZGMxMzQwMDg5NmMyY2FkNmVkYWNlNTllMjRkYzljODU0MGRjNjBhOTc3YWFlNmY4In0 -->
### 2026-11-02T17:33:37.336Z · approval · unspecified

{"approvedBy":"explicit user presentation requests; coordinator bounded design/decomposition","attempt":1,"decision":"approved","decisionEvidence":"User requests no Page1/1, compact content/font-aware height, native controls hidden and separated title/key-hint status borders. Coordinator chooses natural-row renderer contract and bounded private parent-owned geometry adaptation. Native controls capability absent from documented installed schemas/actions; requirement stays open, not waived. No global settings/font changes/tests/destructive/remote authority inferred.","decisionRegisterEntry":{"decision":"Compact status-border rendering and content-fitting owned layer using actual cells; retain native controls hiding as explicitly blocked requirement, no invented flags or full completion.","id":"REXWHICHKEY-20261102-123334","state":"Accepted"},"eventId":"rex-whichkey/presentation/design-approval/1","kind":"approval","recordRef":"designs/rex-whichkey-compact-presentation.md","requirements":["REQ-024","REQ-025","REQ-026"],"schemaVersion":1,"specHash":"sha256:10a72840e340a8bb53e744309fd518d564a9aac00a73364c3a554da06d05388f"}
<!-- mtn:end-entry:v1 -->
