---
workflowKind: task
workflowLanguage: go
workflowOrder: 3
workflowRequirements:
  - REQ-009
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - theme.go
    - theme_test.go
    - testdata/theme
  sharedPaths: []
  resources: []
  interfaces:
    - menu-theme
title: 'Implement terminal palette, fallbacks and accessible output'
mtnCreateRetry:
  key: rex-whichkey/draft-task/04-03-theme/1
  intent: ab2e6c567e0b75332fc5eb4d91e2c84f6dc7b375de302fd27e08cfa419b355da
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:29.409Z'
dateModified: '2026-11-01T00:49:32.008Z'
tags:
  - task
type: task
id: 83e020be-7ac6-416b-a38b-e970cb2f5848
projects:
  - '[[phases/rex-whichkey-04-menu-ui.md]]'
completedDate: '2026-10-31'
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Implement foreground/background/palette resolution and contrast fallbacks, NO_COLOR and ASCII behavior according to the terminal contract.


## Functional criteria

REQ-009
Output behavior: NO_COLOR emits no color SGR; ASCII avoids non-ASCII decorations; absent/malformed palette data falls back safely. Delayed OSC palette replies are terminal messages rather than Escape/cancel input.


## Ownership

Allowed source/test edits: theme.go, theme_test.go, testdata/theme. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.


## Explicit producer/consumer delivery boundary
This is the independently reviewable theme/terminal-palette producer, owning theme.go and its listed paths. Deliver usable palette query/reply decoding, metadata-to-theme interpretation, semantic styles and accessible ASCII decorations without undefined future symbols. Production invocation of these APIs belongs to downstream source owners, not token callers or out-of-scope theme edits: 04-04 routes ordered terminal input through PaletteDecoder before key handling, acquires palette metadata and applies it; 04-02 routes rendered text/decorations through semantic Style/Decoration; 04-05 composes the application. Producer completion is not a claim that those consumers are already integrated or runtime-verified. Phase4 completion requires their actual production wiring and independent source review. Preserve every accepted palette/focus/accessibility behavior; no execution probes or tests are authorized.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDAwOjQ5OjMxLjk5NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNC10MDMtYnVpbGRlci1hMDIvc3VtbWFyeS8yIiwiaW50ZW50IjoiYmYzNGI2ZDBhYjJiOWMyMzhkMjRiMTU4Y2Y1MWRkYjNiMDYxMGY3NDNlZWQ0Mzc3NDYzMDA4NDZmNzQ1Yjg3NiJ9 -->
### 2026-11-01T00:49:31.995Z · summary · unspecified

{"acceptedScope":"Implement terminal palette, fallbacks and accessible output","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p04-t03-reviewer-a02/review/2 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p04-t03-builder-a02/integration/integration.json"],"attempt":2,"eventId":"rex-whichkey-p04-t03-builder-a02/summary/2","files":["theme.go"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-04-03-theme.md","requirements":["REQ-009"],"schemaVersion":1,"specHash":"sha256:0930fff81840797161d2bd2d279314b1eb5621469a1843fe83c9b3b31cde017a","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"f71fd4710a283165ea5b5a091ad7104ff18124de","targetTree":"5fe9049f5d11a330efa7351e766bdbbe7321270d","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
