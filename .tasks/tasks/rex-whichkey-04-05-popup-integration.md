---
workflowKind: task
workflowLanguage: go
workflowOrder: 5
workflowRequirements:
  - REQ-001
  - REQ-005
  - REQ-007
  - REQ-008
  - REQ-011
  - REQ-013
workflowDocuments:
  - ref: prds/rex-whichkey.md
    hash: 'sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1'
workflowOwnership:
  paths:
    - cmd/main.go
    - cmd/main_test.go
    - config.go
    - config_test.go
    - lua/rex_whichkey.lua
    - popup.go
    - tests/integration/popup
  sharedPaths: []
  resources: []
  interfaces:
    - popup-launch-selection
title: 'Wire short-lived popup launch, selection and cleanup'
mtnCreateRetry:
  key: rex-whichkey/draft-task/04-05-popup-integration/1
  intent: a01287e2e4f673506cf505b15f463871777c77ed2da0b7a45f291635965e1b1c
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T13:56:33.418Z'
dateModified: '2026-11-01T07:04:13.865Z'
tags:
  - task
type: task
id: 5f1a721b-86f5-4524-8ed3-e46951b8d2f9
projects:
  - '[[phases/rex-whichkey-04-menu-ui.md]]'
blockedBy:
  - uid: '[[tasks/rex-whichkey-04-04-terminal-lifecycle.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-03-04-invocation-dispatch.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-02-03-block-layers.md]]'
    reltype: FINISHTOSTART
  - uid: '[[tasks/rex-whichkey-03-05-mapping-format.md]]'
    reltype: FINISHTOSTART
completedDate: '2026-11-01'
---
# Product implementation task

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. Provisional IF-* signatures remain unapproved; current execution and completion policy is recorded in projects/rex-whichkey.md.

## Deliverable and scope

Wire the short launcher and terminal-owned Go process to launch/config/selection contracts. Launcher returns promptly; Go owns menu navigation and exits on dismissal or dispatch.


## Functional criteria

REQ-001, REQ-005, REQ-007, REQ-008, REQ-011, REQ-013
Popup survives beyond the configuration action lifetime, carries exact origin identities, executes one validated selection and cleans only owned resources on cancellation/errors. Implement all five graphical positions and focus restoration; moved origins resolve safely. Independent source review must trace layout, focus and cleanup paths; graphical/runtime behavior remains unexecuted, not passed. No background daemon or user init.lua takeover.


## Ownership

Allowed source/test edits: cmd/main.go, cmd/main_test.go, config.go, config_test.go, lua/rex_whichkey.lua, tests/integration/popup. Paths are literal reservations; changes outside them require coordinator replanning. Shared dependency/flake files serialize; no lint/Taskfile weakening or architecture layering.


## Inputs and prerequisites

Use the current PRD and published Rex API/CLI contracts. Read documentation/source as needed; implement only the integration the menu needs. Preserve the flat short-lived Go application, minimal Lua integration, concrete Rex resource types, argument-array execution, origin identity, bounded owned-resource cleanup and user init.lua ownership. No daemon, service hierarchy, new transport framework, research harness or prerequisite final RFC. Unknown API behavior is a concrete implementation issue, not an open-ended research gate.

Use native product-task dependencies, not Phase1/final-RFC gates. Proposed IF-* names in historical drafts are not adopted public contracts. Prioritize mapping and menu tasks before integration-only work.

## Execution and completion

No tests or experiments are scheduled or authorized by this plan revision. Do not silently reinstate old TDD/live-Rex/final-matrix commands. Source review and documentation may describe implementation and residual risks, but cannot claim runtime verification. Preserve source and necessary current state; completed tasks/phases retain concise delivery summaries. Only actual delivered implementation counts as product progress; cancelled research does not count as successful delivery. Independent source review is required before delivery; no runtime verification is claimed. Follow the current execution and completion policy in projects/rex-whichkey.md.

## Authoritative format input

Consume mapping.go and docs/mapping-format.md from tasks/rex-whichkey-03-05-mapping-format.md after its native producer prerequisite is delivered. These Go-owned files define the shared mapping/configuration standard; do not define or copy an independent Lua/Nix/consumer schema. Executable configuration and callbacks remain in the originating Rex Lua runtime. Follow the existing functional requirements; raise a consequential new public interface or dependency choice for user approval instead of silently selecting it.

## Remote implementation and mixed-language review

Own server-side popup executable/config resolution and propagation of exact originating server/session/block/client identities in config.go and cmd/main.go, with only required launcher glue in lua/rex_whichkey.lua. Use the execution host's installed executable/module paths, not a client-local absolute path; 03-01 owns user module loading, 03-04 owns callback/action validation and Phase 2 owns CLI/resource targeting. Preserve no-fallback origin behavior and cleanup only resources returned as owned. 05-05 documents these source paths; it cannot supply missing implementation.

Keep workflowLanguage go for this mixed task. Independent source review must explicitly cover lua/rex_whichkey.lua and its loading/argv/origin/lifetime behavior alongside all owned Go files and the Lua config/dispatch boundary. No Lua work may be omitted because the primary route is Go; no Go TDD/runtime gates are authorized.


## Explicit producer/consumer delivery boundary
Composition owns final source wiring of 04-03 theme producer, 04-02 rendering and 04-04 terminal palette/input lifecycle. Require actual semantic Style/Decoration callers, ordered palette-reply decoding before key handling and metadata application in the assembled popup source. No token callers or claims of graphical/runtime passes; Phase4 cannot complete before this wiring is integrated and source-reviewed.


## Explicit frame producer and loop/composition consumers
Own real executable composition invoking the delivered terminal render/input loop, not an unused helper. The assembled source must reach RenderMenu for initial/navigation/resize/theme refresh, consume frame placements/Lines and pass frame.Pages into Menu.HandleKey. Verify these actual callers along with palette decoder ordering and semantic theme use in the frozen combined source before Phase4 completion. No graphical/runtime execution or test passes may be inferred.


## Accepted bridge and current implementation authority
The user approved bridge decision REXWHICHKEY-20261031-220348 and delegated subsequent bounded technical choices. Consume revised Go standard and Lua dispatch producers only after actual delivery. Implement prompt-return detached Go popup lifecycle and authenticated post-return completion/cancellation using the same original server Lua VM and explicitly captured caller identities, not synchronous waiting or weakened context guards. Go owns shared message/capability data; Lua callbacks/tickets remain private in the originating VM. Wire a user-owned explicit completion-action registration fragment without taking over init.lua. Preserve actual five-position render/input/palette/pages composition, server-local executable/module resolution, origin focus and owned-only cleanup. Go argument-array execution, private/bounded transport artifacts and explicit fail-closed API validation are required; no daemon, Go-Lua VM, deferred attach/WOL/SSH feature or product trials. Partial source 00fe2599ef12ccbc784ed10e992b450a453a6c4a is preserved for a NEW reviewed attempt after producers are delivered; it is not already approved or integrated. Current authority: .mtn-tmp/remaining-delivery/bridge-resume/accepted-authority.md.


## Bounded application lifecycle layout
popup.go is an additional literal owned path for flat application lifecycle helpers, keeping cmd/main.go composition and config.go decoding focused. Prefer a simple invocation-bounded Go controller using documented detached execution and a separate Go popup/PTY child when needed, so closing the UI or invoking an action that closes the origin cannot kill its own cleanup/completion caller. This is not a persistent daemon/service framework; native close-on-exit primitives may simplify it if actually documented. No configuration-Lua timers/events are assumed. Validate real API/argv/context contracts from source/docs; do not fabricate callback hooks or silently retarget. Preserve exact owned returned resource IDs, finite lifetime, private transport artifacts, cancellation and independent current source review. The approved bridge and current delegated authority cover these bounded implementation choices; no new approval round is required.


## Accepted delegated endpoint binding — REXWHICHKEY-20261101-015342
Eliminate the unsupported ctx.server-to-CLI-endpoint conversion requirement. Require an explicit immutable server-owned endpoint as a separate constructor argument, e.g. popup.new(defaults, server_endpoint). This is launcher transport configuration, not a new raw mapping-schema field. The existing Go-owned popup handoff server field and Go endpoint validator remain authoritative. Preserve the ORIGINAL actual ctx object/opaque ctx.server unchanged for issue/bind/completion and all dispatcher checks; never overwrite/copy endpoint or claimed identity into ctx, remove origin guards, or infer an endpoint from opaque data. Missing/invalid/mismatched endpoint fails closed; no host label/current/local/environment fallback or discovery. Before creating any resource/mutation, Go must validate the captured exact session/block/client and connection generation through that explicit endpoint. Keep all registry config/nonce/single-use/expiry/snapshot/confirmation safety and existing owned-only finite controller cleanup. Document the required server-side binding and unexecuted runtime behavior accurately. No automatic ctx.server conversion or verified routing claim is permitted. This bounded public setup choice was selected under the explicit user delegation, not a request for another approval or a claim that undocumented APIs exist.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDA3OjA0OjEzLjg1NFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5Ijoic3VtbWFyeSIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5LXAwNC10MDUtYnVpbGRlci1hMDMvc3VtbWFyeS8zIiwiaW50ZW50IjoiMzk3NWIyZTJjZmVkZGI5NmI5ZDA0ZmExNzcxMTkxMTkyYTcwOTM1MmQwYzkzNzYwMmRjMDFjZGIwNmNmYTUyOCJ9 -->
### 2026-11-01T07:04:13.854Z · summary · unspecified

{"acceptedScope":"Wire short-lived popup launch, selection and cleanup","approvalSummary":["User explicitly requests all remaining approved phases under corrected plan.","Independent source review rex-whichkey-p04-t05-reviewer-a03/review/3 APPROVED current build."],"artifactRefs":["/home/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.mtn-tmp/remaining-delivery/rex-whichkey-p04-t05-builder-a03/integration/integration.json"],"attempt":3,"eventId":"rex-whichkey-p04-t05-builder-a03/summary/3","files":["cmd/main.go","cmd/main_test.go","config.go","lua/rex_whichkey.lua","popup.go"],"kind":"summary","limitations":["No project tests, TDD, experiments or live/remote trials; source review/compilation are not runtime verification.","Fresh same-family reviewer context, not cross-family verification."],"recordRef":"tasks/rex-whichkey-04-05-popup-integration.md","requirements":["REQ-001","REQ-005","REQ-007","REQ-008","REQ-011","REQ-013"],"schemaVersion":1,"specHash":"sha256:d0a34791f1138e58723f03849fa274e1392749030418d88ca6ccde09a09efe64","status":"complete","summaryVersion":1,"targetBranch":"main","targetCommit":"8f2a196e50fd58848540d6ba0ca6b192385b7c30","targetTree":"38b51a2b64eb33e7030aac609e525e23b5e486f1","verification":[{"command":"GOTOOLCHAIN=local GOMAXPROCS=16 go build ./... (combined candidate)","result":"pass"},{"command":"Reviewed changed-file contents unchanged; git diff --check; target commit/tree and ancestry verified","result":"pass"}]}
<!-- mtn:end-entry:v1 -->
