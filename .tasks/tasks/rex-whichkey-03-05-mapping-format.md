---
workflowKind: task
workflowLanguage: go
workflowOrder: 5
workflowRequirements:
  - REQ-003
  - REQ-004
  - REQ-005
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - mapping.go
    - docs/mapping-format.md
  sharedPaths: []
  resources: []
  interfaces:
    - whichkey-format-standard
title: Define Go-owned mapping and configuration standard
mtnCreateRetry:
  key: rex-whichkey-user-format-producer-20261031-v1
  intent: 679843623660b6b84cf23519a30120f46b876c2f502ad85652220443e956a099
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-31T19:57:50.563Z'
dateModified: '2026-11-01T00:49:01.993Z'
tags:
  - task
type: task
id: b8465e20-4fed-4cec-b307-bd1808a8bfc6
projects:
  - '[[phases/rex-whichkey-03-lua-mappings.md]]'
completedDate: '2026-10-31'
---
# Go-owned mapping/configuration standard

## Authority and scope
See projects/rex-whichkey.md for current user authority, implementation-first priorities, independent source-review policy and no-project-tests/no-experiments limits. Golang owns the shared format and standard. This small Phase3 producer prevents an accidental whole-Phase4-menu barrier on configuration; it does not move Lua configuration or callbacks into another VM.

## Deliverable and ownership
Own mapping.go and docs/mapping-format.md only. mapping.go contains the authoritative Go data definitions and validation rules for the accepted configuration/mapping values. docs/mapping-format.md explains that same standard, its exact source paths and minimal Lua consumption. Keep a flat package layout; no generic services, ports/adapters or separate consumer schemas.

## Functional criteria
- Define one Go-owned source of truth covering the existing configuration/default/precedence/replacement behavior and key/group/inheritance/hidden/condition/confirmation/named-action/argument/callback-reference semantics required by REQ-003/004/005/013.
- Distinguish data from executable Lua functions. Callback execution and invocation identity remain in the originating Rex Lua runtime; do not create another callback VM or claim a reflection mapper solves cross-process transport.
- State valid/invalid values and ownership/lifetime constraints sufficiently for the Lua config/compiler/evaluation/dispatch and Go menu/popup consumers to implement existing requirements against these exact files.
- Existing illustrative signatures are not approved public interfaces. Escalate only a genuinely consequential unapproved public field/transport/setup contract or dependency choice before adopting it. Accepted Go ownership is not reopened.

## Consumers and prerequisites
No Phase1 research or Phase4 menu prerequisite. Phase3 Lua tasks, Phase4 navigation/layout/popup and Phase5 default/host-menu consumers depend on this producer and refer to mapping.go and docs/mapping-format.md, rather than copied schema prose. Authoritative functional scope remains prds/rex-whichkey.md at its accepted fingerprint.

## Verification and completion
Independent Go source review and graph/ownership/requirement checks are required. Do not run project tests, experiments, probes or remote trials under current project authority; record those execution limitations honestly rather than claiming passes or reinstating TDD. Package README research is not selection, compatibility verification or delivery. Coordinator integrates only reviewed authorized source, commits each actually completed phase on main, and retains concise completed task/phase delivery summaries with commit references and limitations; metadata-only checkpoints are not implementation delivery.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDAwOjQ5OjAxLjk4MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy10MDUtYnVpbGRlci1hMDIvc3VtbWFyeS8yIiwiaW50ZW50IjoiY2I1YjBjNTU1ZGFhZTFiNmM3ZGE2ZjNhM2VjZWMwZjZmZWNlZjZhYjkxZmM5NDY2NmI0OTIzOTYxZTVlNmE2ZiJ9 -->
### 2026-11-01T00:49:01.981Z · summary · unspecified

{"acceptedScope":"Define Go-owned mapping and configuration standard","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p03-t05-reviewer-a02/review/2 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p03-t05-builder-a02/integration/integration.json"],"attempt":2,"eventId":"rex-whichkey-p03-t05-builder-a02/summary/2","files":["docs/mapping-format.md","mapping.go"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-03-05-mapping-format.md","requirements":["REQ-003","REQ-004","REQ-005","REQ-013"],"schemaVersion":1,"specHash":"sha256:d269b66266505131a8ea3c09d564d2630184052333391a6a8cb06c31235057ca","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"25ba58a071608d28a19af0882a96a262627b5619","targetTree":"fb4e8d067a34dc335e2798a7896417dd57c66c31","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
