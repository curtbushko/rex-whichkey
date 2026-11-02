---
workflowKind: task
workflowLanguage: go
workflowOrder: 3
workflowRequirements:
  - REQ-026
workflowDocuments:
  - ref: designs/rex-whichkey-compact-presentation.md
    hash: 'sha256:10a72840e340a8bb53e744309fd518d564a9aac00a73364c3a554da06d05388f'
workflowOwnership:
  paths:
    - rex/layer.go
  sharedPaths: []
  resources: []
  interfaces:
    - native-popup-chrome
title: Hide native popup controls (blocked Rex capability)
mtnCreateRetry:
  key: rex-whichkey-phase11-task3-v1
  intent: 30a6c4ce2f3cfa103756c3479b3ece9007d6b8cee628ee70ad78442817315cc5
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T17:33:46.520Z'
dateModified: '2026-11-02T19:11:37.266Z'
tags:
  - task
type: task
id: f28f5e7b-277a-422d-8761-1c82f7a65479
projects:
  - '[[phases/rex-whichkey-11-compact-presentation.md]]'
---
# Hide native Rex popup controls (blocked capability)

REQ-026. User requests hiding pane/tab/window controls only on this popup. Own rex/layer.go only for a future documented per-popup capability. Current installed session.new_layer/resize_layer API schemas, terminal creator options_schema with additionalProperties=false, connected app actions, published Lua layout docs expose no documented per-popup chrome/header/toolbar option. Coordinator research artifacts .mtn-tmp/popup-presentation provide exact evidence. No source implementation/worker dispatch is authorized while this capability is unresolved.

Recover only with documented supported current Rex control or an explicitly approved Rex update/integration; alternatively user may explicitly defer this requirement, which is NOT assumed. Never send guessed flags, globally hide Rex headers, edit/reload user init/settings, modify unrelated Rex source or call destructive/native actions to guess capability. Keep incomplete with concrete blocker, no fabricated pass or full phase/project completion. This task does not block independently reviewed renderer/height improvements or truthful limitation documentation.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE3OjM0OjA4Ljk1MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYmxvY2tlciIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3ByZXNlbnRhdGlvbi9uYXRpdmUtY2hyb21lLWJsb2NrZXIvMSIsImludGVudCI6IjY4ZmJhM2NmM2Q2Y2M2NjVkNmFiNzRjN2RiYTBjMWYyYmVhYjg5ZDJhYWIwOGU1OWJhNjcyZWIyYjE3MDA0ODMifQ -->
### 2026-11-02T17:34:08.951Z · blocker · unspecified

{"affectedRefs":["tasks/rex-whichkey-11-03-popup-chrome.md","phases/rex-whichkey-11-compact-presentation.md","projects/rex-whichkey.md"],"artifactRefs":[".mtn-tmp/popup-presentation/new-layer-api.json",".mtn-tmp/popup-presentation/resize-layer-api.json",".mtn-tmp/popup-presentation/terminal-creator.json",".mtn-tmp/popup-presentation/client-actions.json",".mtn-tmp/popup-presentation/lua-reference.md"],"attempt":1,"eventId":"rex-whichkey/presentation/native-chrome-blocker/1","kind":"blocker","reason":"No documented per-popup native header/controls/chrome option in installed session.new_layer/resize_layer/terminal creator schemas or connected app actions and published Lua layout docs. Terminal options reject additional properties. Global settings/user init/Rex source changes are outside authority.","recordRef":"tasks/rex-whichkey-11-03-popup-chrome.md","recovery":"Supply documented supported per-popup Rex capability/update or explicit user requirement deferral. Do not invent flags, modify global settings or mark full phase/project complete.","schemaVersion":1,"specHash":"sha256:00961722dcf9a5f2b742e8fcd91be861c00d825b1f0ed3de4a4f293486ca6250"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDE5OjExOjM3LjI1NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L2ZpcnN0LWZyYW1lL2Nocm9tZS1yZWNoZWNrLzEiLCJpbnRlbnQiOiJhNzAwYWJjOWVjMTAxMDMwM2ZlODkwMGUzYzg2OTg1ZjRmOTZiZDFkNjE5NzNjM2E2Mzk3Yjk0ZWRkZGI3YmMzIn0 -->
### 2026-11-02T19:11:37.255Z · historical · unspecified

{"attempt":1,"content":"User repeats native title/dock strip visible. Rechecked current CLI help/API/new_layer/creators/client actions; only pane.float relates, no documented frameless/per-popup header/dock option found. Requirement still unimplemented, old blocker not resolved or deferred.","eventId":"rex-whichkey/first-frame/chrome-recheck/1","gateEligible":false,"kind":"historical","provenance":"Coordinator read-only installed API/client-action metadata recheck after user correction; artifacts .mtn-tmp/popup-initial-fit/{rex-help.txt,api-list.txt,new-layer-api.json,block-types.json,client-actions.json}.","recordRef":"tasks/rex-whichkey-11-03-popup-chrome.md","schemaVersion":1,"specHash":"sha256:00961722dcf9a5f2b742e8fcd91be861c00d825b1f0ed3de4a4f293486ca6250"}
<!-- mtn:end-entry:v1 -->
