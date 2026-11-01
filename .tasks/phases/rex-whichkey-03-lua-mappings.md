---
workflowKind: phase
workflowOrder: 1
workflowRequirements:
  - REQ-003
  - REQ-004
  - REQ-005
  - REQ-007
  - REQ-011
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
title: 'Phase 3: Lua integration and mapping compilation'
mtnCreateRetry:
  key: rex-whichkey/draft-phase/03-lua-mappings/1
  intent: 6c2ccbf6952c758b7bacc923a94233271905c1cda0137a9d0d376463cf2e1a80
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:55:55.471Z'
dateModified: '2026-11-01T02:06:38.022Z'
tags:
  - task
type: task
id: 4a761217-f5ea-4401-b5dc-159996e5ede3
projects:
  - '[[projects/rex-whichkey.md]]'
blockedBy: []
---
# Product implementation milestone

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable

Implement minimal Rex Lua configuration, full key-sequence compilation, implicit/nested groups, inherited attributes, dynamic labels/conditions, hidden active entries and named actions/callback references against the Go-owned standard in mapping.go and docs/mapping-format.md, produced by tasks/rex-whichkey-03-05-mapping-format.md. Mapping and configuration behavior comes first; executable callbacks stay in Rex Lua. Dispatch wiring follows the core menu and uses the documented API rather than an experimental protocol milestone.

## Requirement coverage

REQ-003, REQ-004, REQ-005, REQ-007, REQ-011, REQ-013. Child records retain their complementary functional responsibilities.

## Inputs and dependencies

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

No Phase1 or final-RFC prerequisite remains. Existing task-level product dependencies still apply; phase numbering does not create a dependency. Prior research is optional historical context, never an approval-required input.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Phase delivery policy

Follow projects/rex-whichkey.md: no project tests/builds/experiments/probes/remote trials; require independent source review. Commit each actually completed phase on main, separately from any metadata-only checkpoint. Completed tasks/phases retain concise actual-delivery summaries with reviewed files, requirements, limitations and commit pointers; active records retain necessary current state only. No proposed plan, source review alone without delivered implementation, successful worker exit or cancelled research establishes phase delivery. Coordinator uses supported native compaction at real completion, not a hand-edited history rewrite or duplicate archive.

## Producer, consumers and remote responsibilities

04-01 is the existing Go source-owning producer; its format/model work does not depend on Lua consumers. All Phase 3 tasks explicitly depend on that producer and use its files rather than copying schema prose or inventing a Lua-owned schema. Keep setup functions, dynamic functions and callback execution in the originating Rex Lua runtime. 03-01 owns server-side module resolution in lua/rex_whichkey/config.lua; 03-04 owns origin/server/session/client identity binding and callback dispatch in lua/rex_whichkey/dispatch.lua. Remote documentation only describes their delivered behavior.


## Go-owned standard producer
tasks/rex-whichkey-03-05-mapping-format.md owns mapping.go and docs/mapping-format.md as a small flat Go producer, independently of Phase4 menu completion. Lua configuration/compiler/evaluation/dispatch remain Lua consumers with real producer edges; no separately owned Lua schema or callback VM is introduced.

## Reopened integration blocker — REXWHICHKEY-20261031-220348 (Open, not approved)
The accepted popup consumer must return promptly and survive its initiating configuration action. The current dispatch producer instead requires the same live action/rawequal context and a 30-second lifetime; integration cannot satisfy both. This is new source evidence, not a new research gate. The dispatcher task and Phase3 are reopened; other delivered child source and prior c3a0b284/961d8cfd references remain accurate historical deliveries. Current target scope is still original-runtime/captured-origin safe post-return dispatch, not same-action blocking. Exact authentication/re-entry/public completion contract is unresolved: do not merely relax context guards, select a new bridge, or change the accepted lifetime requirement without the pending user decision. Details: .mtn-tmp/remaining-delivery/popup-contract-proposal.md and the blocked popup source-notes.md. No new contract or implementation approval is established by this scope reconciliation.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDAxOjM0OjIyLjY3NloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwMy1jb29yZGluYXRvci1hMDEvcGhhc2UwMy9zdW1tYXJ5LzEiLCJpbnRlbnQiOiJhYjdjMzY2NzMyZDUzZTAyOGY4NWE3ZDk4ZDg3YWQzM2E0NGNjODNmMzFlNWI0NDg0NGEwMTc3NzZkMzQyYWYzIn0 -->
### 2026-11-01T01:34:22.676Z · summary · unspecified

{"acceptedScope":"Phase 3: Lua integration and mapping compilation","approvalSummary":["User explicitly authorizes all remaining accepted phases3–6 and per-phase commits without routine confirmation.","All exact native child summaries bind current independent source approval and delivered target commit/tree; ancestry verified."],"artifactRefs":[".mtn-tmp/remaining-delivery/phase03-final/worktrees.txt",".mtn-tmp/remaining-delivery/phase03-final/validate.json"],"attempt":1,"coverage":{"REQ-003":["tasks/rex-whichkey-03-01-lua-config.md","tasks/rex-whichkey-03-05-mapping-format.md"],"REQ-004":["tasks/rex-whichkey-03-01-lua-config.md","tasks/rex-whichkey-03-02-mapping-keys.md","tasks/rex-whichkey-03-03-dynamic-mappings.md","tasks/rex-whichkey-03-05-mapping-format.md"],"REQ-005":["tasks/rex-whichkey-03-03-dynamic-mappings.md","tasks/rex-whichkey-03-04-invocation-dispatch.md","tasks/rex-whichkey-03-05-mapping-format.md"],"REQ-007":["tasks/rex-whichkey-03-03-dynamic-mappings.md","tasks/rex-whichkey-03-04-invocation-dispatch.md"],"REQ-011":["tasks/rex-whichkey-03-04-invocation-dispatch.md"],"REQ-013":["tasks/rex-whichkey-03-01-lua-config.md","tasks/rex-whichkey-03-04-invocation-dispatch.md","tasks/rex-whichkey-03-05-mapping-format.md"]},"eventId":"rex-whichkey-p03-coordinator-a01/phase03/summary/1","files":["docs/mapping-format.md","examples/README.md","examples/my_whichkey.lua","lua/rex_whichkey/config.lua","lua/rex_whichkey/dispatch.lua","lua/rex_whichkey/evaluate.lua","lua/rex_whichkey/mapping.lua","mapping.go"],"kind":"summary","limitations":["No project tests, TDD, experiments, GUI probes, live API/remote trials or config reloads; compilation/source review are not runtime verification.","Independent fresh contexts used the same authenticated model family, not cross-family review.","Prior transferred coordinator process status remains unverified; ownership transfer did not assert remote/local stoppage.","No global cleanup or push; retained worktrees/artifacts remain available."],"recordRef":"phases/rex-whichkey-03-lua-mappings.md","records":["tasks/rex-whichkey-03-01-lua-config.md","tasks/rex-whichkey-03-02-mapping-keys.md","tasks/rex-whichkey-03-03-dynamic-mappings.md","tasks/rex-whichkey-03-04-invocation-dispatch.md","tasks/rex-whichkey-03-05-mapping-format.md"],"requirements":["REQ-003","REQ-004","REQ-005","REQ-007","REQ-011","REQ-013"],"schemaVersion":1,"specHash":"sha256:73dcfb6a7fe83bec62cfceebb9386953ca635fae16095d747de5c2c79b073772","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"c3a0b28435d00e675876f7f846492ce4b480f6f9","targetTree":"e752514be4bc43b9a46588be859e866cc4244874","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined delivered revision)","result":"pass"},{"command":"git diff --check; native graph/schema validation; exact completed child scope and Git delivery identities/ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDAyOjA2OjM4LjAxMFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicGxhbiIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3BvcHVwLWxpZmV0aW1lLXJlY29uY2lsaWF0aW9uL3BoYXNlMy1wbGFuLzEiLCJpbnRlbnQiOiJlMzIzYzc2Y2NlNWYyMWQzYjI4M2Q0MWY0NTU4NjRjYmQzYTIzYzNhNGUzY2JkYTRlYjc5MmMxOWY4MzQ2ZGJmIn0 -->
### 2026-11-01T02:06:38.010Z · plan · unspecified

{"approvedBy":"user","attempt":2,"coverage":{"REQ-003":["tasks/rex-whichkey-03-01-lua-config.md","tasks/rex-whichkey-03-05-mapping-format.md"],"REQ-004":["tasks/rex-whichkey-03-01-lua-config.md","tasks/rex-whichkey-03-02-mapping-keys.md","tasks/rex-whichkey-03-03-dynamic-mappings.md","tasks/rex-whichkey-03-05-mapping-format.md"],"REQ-005":["tasks/rex-whichkey-03-03-dynamic-mappings.md","tasks/rex-whichkey-03-04-invocation-dispatch.md","tasks/rex-whichkey-03-05-mapping-format.md"],"REQ-007":["tasks/rex-whichkey-03-03-dynamic-mappings.md","tasks/rex-whichkey-03-04-invocation-dispatch.md"],"REQ-011":["tasks/rex-whichkey-03-04-invocation-dispatch.md"],"REQ-013":["tasks/rex-whichkey-03-01-lua-config.md","tasks/rex-whichkey-03-04-invocation-dispatch.md","tasks/rex-whichkey-03-05-mapping-format.md"]},"decision":"approved","decisionEvidence":"User authorizes continuing accepted phases and reconciling actual implementation issues. Existing prompt-return/original-runtime requirements require reopening this producer after a concrete consumer incompatibility; task ownership, requirements and document bindings remain. The proposed post-return security/interface contract is explicitly Open and blocked, not approved or implemented by this manifest.","dependencies":[{"prerequisites":["tasks/rex-whichkey-03-05-mapping-format.md"],"ref":"tasks/rex-whichkey-03-01-lua-config.md"},{"prerequisites":["tasks/rex-whichkey-03-01-lua-config.md","tasks/rex-whichkey-03-05-mapping-format.md"],"ref":"tasks/rex-whichkey-03-02-mapping-keys.md"},{"prerequisites":["tasks/rex-whichkey-03-02-mapping-keys.md","tasks/rex-whichkey-03-05-mapping-format.md"],"ref":"tasks/rex-whichkey-03-03-dynamic-mappings.md"},{"prerequisites":["tasks/rex-whichkey-03-03-dynamic-mappings.md","tasks/rex-whichkey-03-05-mapping-format.md"],"ref":"tasks/rex-whichkey-03-04-invocation-dispatch.md"},{"prerequisites":[],"ref":"tasks/rex-whichkey-03-05-mapping-format.md"}],"documents":[{"hash":"sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1","ref":"prds/rex-whichkey.md"}],"eventId":"rex-whichkey/popup-lifetime-reconciliation/phase3-plan/1","kind":"plan","limitations":["Metadata correction only; no implementation delivery or completed phase claimed.","No project tests or experiments performed; independent source review remains required for implementation.","Go-Lua packages remain researched candidates, not adopted dependencies; consequential public interfaces remain unapproved."],"recordRef":"phases/rex-whichkey-03-lua-mappings.md","records":[{"ref":"tasks/rex-whichkey-03-01-lua-config.md","specHash":"sha256:334b767966694c3f7940807254c3f04d071d8b10369f125331351aa93d4f1dfb"},{"ref":"tasks/rex-whichkey-03-02-mapping-keys.md","specHash":"sha256:ce55b445d2d930d8eca0d2b32cd57fca97a8034e6087cfedb70b45dec0fb3f6d"},{"ref":"tasks/rex-whichkey-03-03-dynamic-mappings.md","specHash":"sha256:cc1e5010665b1432299c5f8709d039cf697cd8a30f278abb06c09594c41b9fdc"},{"ref":"tasks/rex-whichkey-03-04-invocation-dispatch.md","specHash":"sha256:35ca71748243fe387179fbbfc6c80624edfce3fa2432eb3c12e95ffb6dd54dab"},{"ref":"tasks/rex-whichkey-03-05-mapping-format.md","specHash":"sha256:d269b66266505131a8ea3c09d564d2630184052333391a6a8cb06c31235057ca"}],"schemaVersion":1,"specHash":"sha256:0117a5cf37ae0c266644addbb8242c93a0ebf048354973564a48b1abed4c4536"}
<!-- mtn:end-entry:v1 -->
