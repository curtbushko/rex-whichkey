---
workflowKind: design
title: Native action audit and explicit late target choice
mtnCreateRetry:
  key: rex-whichkey-action-audit-design-v1
  intent: 5c5e5f0e38b0e66ddf336374d95dadf574e6b1d84d7ba34b496e3223f7677c9b
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T20:15:50.805Z'
dateModified: '2026-11-02T20:15:56.130Z'
tags:
  - task
type: task
id: 371de676-2811-44f5-baf0-66d1554c3594
completedDate: '2026-11-02'
---
# Native action targeting and behavioral action audit

## User scope and verified fault
User requests verification of every action, specifically split down, and suspects others. REQ-030: audit all25 executable native default bindings against installed contracts AND behavior; resource-changing actions require resulting-state evidence, UI-only receipt is not a pass. REQ-031: native actions work without mandatory --client when several app targets exist by explicit user target choice at ACTION time, not arbitrary selection or blocking menu display. Existing bounded technical delegation authorizes the safe chooser implementation; no security weakening or global settings change.

Actual fresh reproduction .mtn-tmp/action-audit/split-reproduction: bare s selects Split down but fails before dispatch because both work/m4-pro apps attached. Same fresh menu with explicit work app exits0 and creates verified below-origin vertical split. Both installed action schemas explicitly accept direction=down; do not change it to guessed CLI below. Both action registries currently match. Current generic guard affects all native defaults, not just s.

## Bounded targeting repair
Retain initial normal native popup app-independent and fixed bounds through navigation, compact seed/spacers and one-shot confirmation/selection. Only after executable selection resolve target intent: exact explicit selector remains strict/no fallback; zero apps refuses safely; one app binds current full ID/generation; several apps offer a short explicit target picker rather than guessing first/last/focused/terminal owner. Existing canonical Go Config/Mapping/menu and owned popup lifecycle can implement a separate action-target prompt AFTER original menu cleanup; it is not a submenu resizing loop. Canceling target prompt dispatches nothing. Do not add a new bidirectional service/schema/framework, saved app preference or mandatory environment variable.

Target choices retain private Go-owned full ID/connection generation snapshot; labels only display, duplicate labels disambiguated with IDs, untrusted/control labels sanitized. Child returns only a canonical choice resolved against retained parent choices, never replacement action/args. Original retained action/args/confirmation must survive target choice unchanged. Revalidate exact origin/session/app attachment/generation and selected action compatibility immediately before one native dispatch; reconnect/choice loss fails, never silent retarget/uncertain retry. All popup/target UI/files/leases cleaned BEFORE executing original action, including origin-closing actions. Action dispatch deadline starts after user choice, not across indefinite UI. Preserve signal cancellation through entire native workflow and configured original-Lua-VM route/expiry/capability semantics unchanged.

## Verification authority and isolation
Latest user explicitly authorizes exercising default actions, including bounded close/new/rename/navigation on ONLY automation-created disposable resources. This refines earlier prohibition on destructive-action trials solely for these owned empty fixtures. Never close/rename user/coordinator/worker panes or sessions, change config/fonts/settings, reload Rex or try remote/Darwin/system/destructive features. Project tests/TDD/unit/race/general suites remain prohibited. Leaf source-only; coordinator performs actual menu-driven action trials with pinned endpoint/app/session/block identities, app metadata and resource before/after/effect/cleanup. Separate audit sessions are allowed solely as disposable action fixtures, not a second coordinator or worker session. Cleanup exact allowlisted IDs after proof, no global session/label kill.

Audit25: wn/ws/wh/wl/wr/wx; tn/th/tl/tr/tx/t1..t9; v/s/x/e/c. Resource cases18: new/close session; new/previous/next/goto1..9/close tab; right/down split/close pane. UI-only7: session picker, previous/next session selected view, session/tab rename dialog, sidebar, palette. Current client.inspect exposes IDs/metadata/attachments/actions, no selected-session/dialog/sidebar state; terminal send-key does not operate native app dialogs. Investigate documented observable mechanisms before calling these pass; do not open unclosable prompts or guess argument names. If still unobservable, keep dedicated UI verification task BLOCKED, document invoked/schema-only versus verified/missing, and do not claim every action works or full phase/project completion. Native title/dock REQ026 remains independently blocked.

One repair task, dependent evidence documentation task, separate blocked GUI-only verification task if current observability gap persists. Fresh frozen review, combined compile, actual bare-menu Split down through explicit target picker plus cancellation/stale selector and exact cleanup before repair completion. Coordinator then runs full disposable resource matrix, records both-app contract differences and unverified GUI cases, refreshes local binary and checkpoints partial delivery honestly. No push.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDIwOjE1OjU1LjIyMloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9hY3Rpb24tYXVkaXQvZGVzaWduLWFwcHJvdmFsLzEiLCJpbnRlbnQiOiI0MmE3NjI1M2E1MWFmNjNmY2RkMzZjZTY3MjM1OGU2ODA1NDBiZjIyYzA5NmI2NGY3MGE3MTk4NTQ3OTRiZjM4In0 -->
### 2026-11-02T20:15:55.222Z · approval · unspecified

{"approvedBy":"Latest explicit user verify every action including split down; previously delegated bounded implementation choices","attempt":1,"decision":"approved","decisionEvidence":"Real bare split-down fails because multiple app targets; explicit app succeeds with correct down arg. Safe explicit late app choice implements action usability without guessing/mandatory flags. Disposable-only native action trials authorized by latest verification request; prior no-general-tests/settings/fonts/platform/user-resource constraints remain. No user claim that UI-only observation is available or all actions pass.","decisionRegisterEntry":{"authority":"Bounded implementation choice under prior explicit delegation, not a claim the user personally chose this UI design","decision":"Late explicit app chooser for ambiguous native action target; retain immutable original action/args/confirmation, strict explicit selectors, exact generation/origin checks and cleanup before one dispatch. Verify resource effects and keep GUI-only tests blocked until observable.","id":"REXWHICHKEY-20261102-151552","rationale":"Shared refusal blocks all native actions with two apps. Never select an arbitrary device or hide unverified GUI behavior.","state":"Accepted"},"eventId":"rex-whichkey/action-audit/design-approval/1","kind":"approval","recordRef":"designs/rex-whichkey-action-audit.md","requirements":["REQ-030","REQ-031"],"schemaVersion":1,"specHash":"sha256:2d2d664eb4bbc5828b2f0d6d4a98346becfbbd03d90e8f9abc5a9874edc8b7a8"}
<!-- mtn:end-entry:v1 -->
