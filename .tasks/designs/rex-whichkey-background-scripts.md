---
identity:
  kind: design
  designID: rex-whichkey-background-scripts
  uuid: 1008a512-fdc1-426c-803c-a1b7bbc840e6
lifecycle:
  state: approved
  sourceEventID: rex-whichkey-p20/design/authority/1
  changedAt: '2026-11-24T02:57:56.402Z'
title: Origin-host background script requirements and design
mtnCreateRetry:
  key: rex-whichkey-p20-design-create-v1
  intent: 798efe68ff4b6f6070fb9f5e3b8e5841e29648f5d3de5c72e14a40dd22a6d301
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-24T02:57:49.904Z'
dateModified: '2026-11-24T02:57:57.454Z'
tags:
  - task
type: task
id: aaf0b499-6ae2-44a5-815b-3908ed87640a
---
# Phase20 background scripts requirements/design

Status: proposal, not a managed record or approval. Proposal ID REXWHICHKEY-20261123-214535 (Open); script-error lifetime refinement REXWHICHKEY-20261123-214752 (Open). Prepared 2026-11-24T02:45:35Z, refined after coordinator supplement at 2026-11-24T02:47:52Z by rex-whichkey-p20-planner-a01.
Managed ref: designs/rex-whichkey-background-scripts.md. Parent: phases/rex-whichkey-20-background-scripts.md.

## 1. Authority
✓ Agreed user behavior, as recorded in the phase and repeated in this delegation: background scripts, no new visible script pane/tab or focus takeover, originating pane host and current cwd, origin-host existence check/error if absent, clear startup errors, no later-focused-pane targeting, copying/local fallback or arbitrary Lua callbacks. Phase records the user request “create a phase, break down the tasks and implement the feature”. This authorizes bounded technical decomposition/implementation, not explicit user approval of this proposed syntax, process mechanism or error UI. Current quality-gates design supersedes historical no-test clauses. Phase19 chrome deferral and visible/TTY/worker/focus cleanup-before-effect remain settled. Existing built-ins retain their exact detach-before-dispatch implementation. The script-only minimized-host refinement below addresses the newly required clear startup-error behavior without displaying a script pane.

## 2. Requirements and acceptance
- REQ-SCRIPT-001: Add one narrow version-1 JSON script action; preserve strict exact-member/type/Unicode/duplicate validation, existing defaults/actions/providers and confirmation. Proposed syntax: action `script.run`, args `{path: string, args?: string[]}`; no command string, shell selector, environment override, callback, generic Rex method or implicit interpolation. Loading and standalone `menu` remain effect-free. Accept literal spaces/metacharacters/Unicode/empty argument elements; reject empty path, NUL, wrong types, extra members and malformed JSON. Existing whole-spec replacement remains.
- REQ-SCRIPT-002: Execute once after retained-selection/confirmation, worker/TTY cleanup, verified menu invisibility and exact original-focus/placement checks. For scripts, the proposal verifies the exact existing layer is minimized before startup, keeping it only for certain-startup-failure recovery; successful startup retires it by exact detach and exits. Resolve only captured session/block/endpoint, never later focus. Prove executing process is on origin host/account before filesystem access or launch; mismatch/missing identity refuses. Query exact origin `process.foreground.cwd` immediately before dispatch; require an absolute known cwd. Start detached from menu lifetime and controlling PTY, standard streams disconnected, and return after OS startup, not script completion. No additional block/window, input injection, script copying or host fallback. Normal remote origins use the Go menu already running on that origin server; this is not remote execution by a local renderer.
- REQ-SCRIPT-003: Resolve relative paths against that cwd on that host, not config directory, popup cwd, home or PATH. Check resolved target exists and is a regular executable script before launch; normal symlinks may resolve to suitable targets. OS direct exec/shebang is proposed; invalid/missing interpreter, permission, missing cwd/host and exec failures must return actionable, sanitized startup errors with no retry. No shell fallback on ENOEXEC. Error must reach the normal user invocation, not merely a soon-deleted detached terminal. Do not promise atomic protection against trusted same-account script replacement or completion monitoring.
- REQ-SCRIPT-004: Own code and meaningful regression tests together; preserve built-ins/result-only mode and all Phase19 lifecycle assertions. Cover literal argv/path, exact origin host/cwd, mismatches/unavailable cwd, missing/nonregular/unexecutable targets, invalid interpreter, cancellation/cleanup failures, once-only start and survival after parent exit. Document trusted executable configuration, host-local paths, executable/shebang requirement, host-process environment (not an interactive pane shell's aliases/exported variables), disconnected stdio and no later completion reporting. Include a usable JSON example and safe isolated outcome evidence. Required gates: task build; task test (unit+race); full task lint (vet+golangci+AI); configured architecture gate if found. No rule/assertion/config/dependency weakening.

## 3. Bounded technical proposal
◇ Proposal 3a / REXWHICHKEY-20261123-214535: one coherent Go feature task, including parser, real dispatch consumer, process implementation, tests and docs. Extend concrete Rex metadata helpers only for consumed host/process fields. Keep flat application and launch-only Lua; no daemon, supervisor, new dependency or public generic execution surface.

Capture own-popup host identity while exact sole-block ownership is valid; compare fresh origin/own descriptions and native execution-host identity, preserving explicit endpoint/no-autostart. Own-block metadata alone or a default CLI endpoint is NOT sufficient proof of local execution: REX_SERVER can point elsewhere/fall back. Require trustworthy native local-host/account evidence; refuse unknown/mismatch. Existing launcher creates its terminal through session.new_layer on the addressed server. Co-location is an invariant to validate, not an assumption to skip.

Use Go's native direct process-start boundary with exact argv, absolute executable/cwd, detached Unix session and real file stdio handles. Native StartProcess/exec startup handshake can report exec/chdir/interpreter errors without waiting for completion. Own resource/reaping behavior explicitly; never let menu context cancellation kill an already accepted detached script. Check cancellation just before start; after an accepted start, report uncertainty without retry if later bookkeeping fails. Builders choose the smallest correct resource implementation.

## 4. Script-only error recovery and remaining validation
◇ Proposal 4a / REXWHICHKEY-20261123-214752: reuse the exact owned menu, not a script terminal. Assessed coordinator supplements planning-alternatives.txt and planning-error-equivalence.txt. Public invariants are invisible/TTY/worker/focus cleanup before effects, exact origin and no new script pane; detach is the current internal mechanism, not a new public feature requirement. Newly required clear startup errors justify an equivalent script-only mechanism; do not reopen settled built-in lifetimes or chrome. Installed `session.set_layer_state` supports minimized/normal, retains bounds/blocks, and returning to normal leaves focus unchanged unless explicitly requested. This gives a supported alternative to destroying the layer before a fallible start.

1. After selection/confirmation, perform host/cwd/path preflight while the existing menu remains available. Show failures in that same menu as an error-only acknowledgement, never an automatic retry or fresh config selection. No script is started.
2. Join workers, restore TTY, minimize the exact sole-owned floating layer, independently verify state=minimized and unchanged ownership, restore/verify exact origin placement/focus and explicit app generation, then do one native OS start. No script can run if these checks fail. Mere minimized-request ACK is not evidence of invisibility.
3. On certainly-failed startup only, revalidate the same layer/block/origin identities, restore that same layer to normal for error acknowledgement, and finally restore TTY/focus and detach it. No block/window creation, script output pane, later-focus retargeting, config reread, script retry, arbitrary Lua or notification command. Focus for error acknowledgement is limited to this owned UI, never successful execution.
4. On successful startup, detach that exact minimized menu host and exit promptly. Do not wait for or kill the script. A post-start detach failure cannot undo the accepted process; report uncertainty, do not claim no execution or retry. Keep the existing post-exit-retirement limitation honest.

History: P20-ERROR-SURFACE-001 was initially a credible blocking invisible-stderr concern: directSelection detaches before dispatch, cmd/main.go writes later errors to its own stderr and Lua deletes that block on exit. Restoring a detached block directly as a float was not supported (float_block rejects detached; new_layer creates blocks). After inspecting the coordinator supplement and set_layer_state, this is now a required implementation/review validation, not an unavoidable scope blocker. Preserve old built-in detach assertions and add script-specific minimized/no-visible-effect/error-recovery assertions; do not weaken tests to treat an arbitrary placed layer as cleanup.

? Validate 4b: actual minimized-layer invisibility, error readability/acknowledgement, process survival and native host proof need isolated same-revision evidence. Published schemas establish supported facilities, not an executed feature. If real evidence disproves this equivalent lifecycle, stop affected work and return the exact conflict before substituting an unsupported UI or changing the accepted visual/once-only/host boundary. User decisions needed now: none. Coordinator may record bounded technical delegation; neither proposal is explicitly user-approved in detail.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDAyOjU3OjU1LjIxNloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMjAvZGVzaWduL2F1dGhvcml0eS8xIiwiaW50ZW50IjoiYjM3OTc3YjI4ZmNjODIxYjMzYmNjYjMxNjZlNjJhMDc5YjkxMmE1NjI5MGE3NTAyZGUzODc2ZDUyM2IzNzUzZSJ9 -->
### 2026-11-24T02:57:55.216Z · approval · unspecified

{"approvedBy":"User requirements and explicit bounded execution authority; coordinator-selected implementation plan","attempt":1,"decision":"approved","decisionEvidence":"User explicitly required background/origin host/current cwd/existence error, then create a phase, break down the tasks and implement the feature. Current handoff authorizes bounded design, reviews, integration and completion without routine reapproval. This is not a claim of separate user approval of technical details.","eventID":"rex-whichkey-p20/design/authority/1","kind":"approval","planningEvidence":[".mtn-tmp/p20/planner-result.json",".mtn-tmp/p20/architecture-result.json"],"recordRef":"designs/rex-whichkey-background-scripts.md","requirements":["REQ-SCRIPT-001","REQ-SCRIPT-002","REQ-SCRIPT-003","REQ-SCRIPT-004"],"schemaVersion":1,"specHash":"sha256:1a7cb4ce61a8d2723f031e956b35b9a3e758285a41152d04c6d63debf2e99eb9","technicalAuthority":"Same-menu minimized script-only lifetime assessed as equivalent visible cleanup; built-ins retain detachment. User approval of additional scope is not inferred."}
<!-- mtn:end-entry:v1 -->
