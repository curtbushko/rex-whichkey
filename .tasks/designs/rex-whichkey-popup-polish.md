---
workflowKind: design
title: 'Popup polish: lifetime, defaults and colors'
mtnCreateRetry:
  key: rex-whichkey-popup-polish-design-v1
  intent: d80953cf83b71ce123bac337505230a837d82dd2ea2d6c5debca7f85a1532b5c
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-02T13:16:53.141Z'
dateModified: '2026-11-02T13:16:58.476Z'
tags:
  - task
type: task
id: 8c70281c-448d-48e2-aa40-357106924f5f
completedDate: '2026-11-02'
---
# Popup lifetime, default-menu reduction and appearance

## User scope and authority

User reports idle popup failure, asks whether a timeout exists, explicitly removes Focus and Rex-known hosts, and requests distinct key/icon/text colors inspired by https://github.com/curtbushko/neovim-flake/blob/main/config/plugins/which-key.nix. Existing bounded technical implementation authority applies. Current source native.go creates a 30-second request; popup.go reserves five seconds and reports expected child timeout as a generic exit failure. The downloaded Neovim source names mostly yellow icons and cyan/orange/green/blue category accents; it does not define an exact key/description hex palette.

## REQ-018 — Normal native popup lifetime

The normal Go/native foreground popup stays open while idle until the user selects/dismisses, parent cancellation occurs or a real terminal/origin failure requires cleanup. Do not turn an ordinary idle wait into 'popup child did not exit successfully'. Remove the fixed native UI deadline rather than merely extending it. Operations such as RPC/preflight/cleanup/dispatch retain bounded timeouts; parent cancellation, terminal/process loss and owned-only cleanup remain handled.

Keep the configured original-Lua-VM bridge's authenticated one-shot/config-generation/30-second expiry protection. Native indefinite interaction must not make expired Lua capabilities usable, renew them from untrusted UI data or weaken completion/context guards. If the configured bridge reaches its intentional expiry, treat it as graceful cancellation/expiry without executing a selection or emitting the misleading generic child-failed message. Document the native/configured lifetime distinction accurately. No daemon or orphan indefinite controller: native caller remains a foreground invocation tied to its child/cancellation.

## REQ-019 — Reduced default menu

Remove the Focus group and its fh/fj/fk/fl entries from both Go native defaults and Lua-composed defaults. This does not remove internal origin-focus restoration/safety or user-configured focus actions. Remove the default h/Rex-known-hosts provider entry and stop importing its fragment by default; absence must avoid inventory requests. Preserve optional user-explicit known-hosts provider support and customized mappings. Do not delete unrelated code or remove session/tab navigation. Keep Go/Lua defaults consistent.

## REQ-020 — Colored key, icon and label roles

Give displayed keys, icons and labels distinct semantic color handling, inspired by the linked named-color WhichKey configuration. Use yellow as the default icon accent, cyan/blue key accent and readable primary/group label treatments; do not claim exact RGB values not present in that file. Add a real icon role instead of styling icons as ordinary text. Provide immediate ANSI color fallback when Rex does not supply OSC background/palette replies, then apply readable terminal palette colors when available. Do not leave the whole menu unstyled merely because a background reply is absent. Keep NO_COLOR/explicit no-color, ASCII behavior, readable contrast and terminal state restoration. Do not overwrite the terminal's palette or assume a dark background; unknown-background styling must retain terminal background/default readable text.

Canonical mapping schema remains Go-owned and unchanged unless an actual need is approved. Prefer semantic style roles over adding a separate Lua/Nix color schema or a generic theme framework. User-facing descriptions/icons/keys are still sanitized and clipped by cell width; SGR must not break alignment/navigation or color alone convey eligibility/confirmation.

## Scope, rollout and verification

Four tasks: lifetime (Go), reduced defaults (Go/Lua), semantic appearance (Go), then docs/combined assembly. First three have distinct files and may run within current two-builder budget; final docs depends on all. Preserve current native default startup, optional configured callbacks, actual origin/app membership/generation, single selection, confirmation, private files, no uncertain retry and owned-only cleanup. No user Rex configuration/settings edits/reloads, daemon, new dependency, host features, push or global cleanup.

User's standing targeted runtime authorization plus these reported popup issues permits coordinator verification in the existing Linux Rex session: actual normal `make run`/binary popup, idle for at least35 seconds beyond the old deadline, confirm it remains visible/alive, inspect colored VT/HTML spans and absent default groups, navigate Tabs and Escape-dismiss with exit0/exact cleanup. Do not execute destructive/native menu actions or configured Lua callbacks for the check. Darwin/remote behavior, general unit/race/TDD/test suites and unrelated probes remain unauthorized. Builders/reviewers static/format/source checks only; coordinator compiles exact combined candidates and performs the bounded actual UI gate before final completion. Treat source-only and observed runtime evidence separately.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAyVDEzOjE2OjU3LjU3MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9wb3B1cC1wb2xpc2gvZGVzaWduLWFwcHJvdmFsLzEiLCJpbnRlbnQiOiJmNjE1YTBjN2Y4NTVkZDI1MjliOTk2OTBhMWE4M2MwODA2NTk2OTI5MDIwYzJlZDRiMGExZjk4MmI2NjlmNTAyIn0 -->
### 2026-11-02T13:16:57.571Z · approval · unspecified

{"approvedBy":"user requested scope; coordinator bounded implementation choices under delegation","attempt":1,"decision":"approved","decisionEvidence":"User reports idle popup exit, requests removing Focus/known-hosts and separate colors inspired by linked Neovim configuration. Coordinator chooses idle-until-dismiss native interaction, retains configured capability expiry, removes only default groups and uses semantic ANSI/palette colors. Standing targeted live UI verification is extended to the reported idle/color/default behavior; no general test, destructive action or config/remote probe authority inferred.","decisionRegisterEntry":{"decision":"Polish native popup idle lifetime, trim defaults and add visible semantic key/icon/label color with no-color/contrast safeguards.","id":"REXWHICHKEY-20261102-081654","previousState":"Open","state":"Accepted"},"eventId":"rex-whichkey/popup-polish/design-approval/1","kind":"approval","recordRef":"designs/rex-whichkey-popup-polish.md","requirements":["REQ-018","REQ-019","REQ-020"],"schemaVersion":1,"specHash":"sha256:ce42e479546a1181ce61138f0f5d37c88e5ed18ae3b6c39987fca90d89aee8f0"}
<!-- mtn:end-entry:v1 -->
