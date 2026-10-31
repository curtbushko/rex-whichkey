---
workflowKind: task
workflowLanguage: go
workflowOrder: 2
workflowRequirements:
  - REQ-006
  - REQ-008
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - layout.go
    - layout_test.go
    - render.go
    - render_test.go
    - testdata/render
  sharedPaths: []
  resources: []
  interfaces:
    - menu-layout
title: Implement constrained Unicode-aware layout and paging
mtnCreateRetry:
  key: rex-whichkey/draft-task/04-02-layout/1
  intent: 7268f8b6ef3d4ebd79893dee3b777c4a552bd616c12b2539e8133a51d35acd1d
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:27.393Z'
dateModified: '2026-10-31T19:58:31.188Z'
tags:
  - task
type: task
id: 003b96b8-24d9-4902-9e1f-5e3614f3a947
projects:
  - '[[phases/rex-whichkey-04-menu-ui.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-04-01-navigation.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-03-05-mapping-format.md]]'
    reltype: FINISHTOSTART
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Render menu state using the product terminal approach, adaptive columns and Unicode display/grapheme widths; compute stable paging with icons, breadcrumbs, help and footer.


## Functional criteria

REQ-006, REQ-008
Combining marks, wide characters and icons are preserved; no out-of-bounds cells for narrow/tiny/zero-resize transitions. All five positions have constrained layouts and stable Tab page navigation. Preserve original Unicode codepoints and correct display widths.


## Ownership

Allowed source/test edits: layout.go, layout_test.go, render.go, render_test.go, testdata/render. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Authoritative format input

Consume mapping.go and docs/mapping-format.md from tasks/rex-whichkey-03-05-mapping-format.md after its native producer prerequisite is delivered. These Go-owned files define the shared mapping/configuration standard; do not define or copy an independent Lua/Nix/consumer schema. Executable configuration and callbacks remain in the originating Rex Lua runtime. Follow the existing functional requirements; raise a consequential new public interface or dependency choice for user approval instead of silently selecting it.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDExOjI1OjI2Ljc0M1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3doaWNoa2V5LWZpcnN0L3ByZXNlcnZlL3Rhc2tzX19yZXgtd2hpY2hrZXktMDQtMDItbGF5b3V0Lm1kIiwiaW50ZW50IjoiOWE0OTBhNmM1ODU0M2RmNjVkMjA1ZGEyYTlhZDBhZDQyOTVhM2YxOTFiZTdkZWU2ZDYzMTQ2YjA1MTNmNjZmMSJ9 -->
### 2026-10-31T11:25:26.743Z · historical · unspecified

{"schemaVersion":1,"eventId":"rex-whichkey/whichkey-first/preserve/tasks__rex-whichkey-04-02-layout.md","kind":"historical","recordRef":"tasks/rex-whichkey-04-02-layout.md","specHash":"sha256:c4ce7a8a2e41e5ff71c79d6863db1485b82631fb98dfba5092092590d2622b38","attempt":1,"content":"# Draft task\n\nDRAFT ONLY: user approved creation of this decomposition, not final interfaces or implementation dispatch. Governing decision REXWHICHKEY-20261030-094737 is Accepted for draft creation. PRD evidence repair and fingerprint-bound design approval remain gates. Paths and contracts below are provisional planning reservations, not accepted public APIs; reassess them after the RFC and republish fingerprints before dispatch. Follow the accepted flat concrete Go architecture, not generic hexagonal defaults. Go workflow language denotes this Go project; Lua/Nix work also requires its specialist checks. Workers never edit .tasks; coordinator alone publishes design/approval evidence.\n\n## Deliverable and scope\nRender menu state using the accepted terminal approach, adaptive columns and Unicode display/grapheme widths; compute stable paging with icons, breadcrumbs, help and footer.\n\n## Requirement criteria\nREQ-006, REQ-008\nCombining marks, wide characters and icons are preserved; no out-of-bounds cells for narrow/tiny/zero-resize transitions. All five positions have constrained layouts and stable Tab page navigation. Strict Unicode tests must assert product success, not merely record observations.\n\n## Ownership\nAllowed source/test edits: layout.go, layout_test.go, render.go, render_test.go, testdata/render. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.\n\n## Inputs, outputs and prerequisites\nConsume the approved RFC and relevant frozen producer contracts. Output the behavior and tests above. Explicit task prerequisites: tasks/rex-whichkey-04-01-navigation.md. Check native parent-phase dependencies too; membership alone does not confer readiness.\n\n## TDD and verification\nRED/GREEN golden/property-style geometry fixtures and strict Unicode regression through actual renderer output; root quality gates. Each code change includes its own failing-first tests, GREEN evidence and refactor checks; do not split tests into a separately completed task. Root Go gates are task build, task test (unit and race), task lint. Lua requires its isolated real-Rex/config fixtures; Nix requires package/module evaluation.\n\n## Integration and completion\nRequire independent Go-project combined specification/quality review plus relevant Lua/Nix expertise. Coordinator integrates reviewed results serially, runs affected checks on the combined target and publishes evidence. Final-phase gates belong to the parent milestone; unavailable GUI/remote infrastructure remains unmet. No implementation, dispatch, completion or interface adoption is authorized by this draft.","provenance":"Exact description before user-directed implementation-first replanning REXWHICHKEY-20261031-072107; previous evidence remains untouched.","gateEligible":false}
<!-- mtn:end-entry:v1 -->
