---
workflowKind: prd
title: 'Usability: direct launch and platform socket defaults'
mtnCreateRetry:
  key: rex-whichkey-usability-requirements-v1
  intent: 395e93c56b1a9ac6af57e0db140ef92c2348f97a687babd8e629e215a8540295
status: done
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-01T14:52:20.455Z'
dateModified: '2026-11-01T15:15:00.147Z'
tags:
  - task
type: task
id: e13c2135-a806-4dd5-9bcd-8703adc6e83d
completedDate: '2026-11-01'
---
# Usability requirements

## Approval and context

The user explicitly requests one new Usability phase, a task breakdown, and implementation: (1) launching rex-whichkey with no arguments should run the popup menu; (2) provide platform-dependent default sockets for Linux and macOS. Existing delegated authority covers bounded technical implementation choices. These requirements extend, rather than renumber or replace, prds/rex-whichkey.md REQ-001–REQ-015.

## REQ-016 — Direct launch

With the normal Rex integration prerequisites installed, invoking `rex-whichkey` with no arguments launches the real default popup/menu path without manually supplying internal subcommands, private handoffs, or resolved JSON. It must use the actual delivered navigation/render/defaults/callback path, not a greeting, usage-only placeholder, token caller or hard-coded demo. Existing `menu`, `prepare`, `controller`, and `popup` plumbing remains compatible. Ordinary command help explains the supported user entrypoint and prerequisite failures. Missing Rex/configuration/origin prerequisites fail clearly and safely; do not fabricate success or silently target whichever client/block is focused. Do not modify or reload the user's init.lua automatically.

## REQ-017 — Platform socket defaults

Supply Linux and macOS default endpoints by delegating native platform discovery to the installed Rex CLI and capturing its concrete endpoint, preserving an explicit endpoint override. Ordinary invocation and server-side Lua composition must no longer require a socket literal in the common platform-default case. Resolve on the execution/server host, not a viewing Mac. Do not confuse opaque `ctx.server` identity with a routable endpoint. Missing required home/runtime context, unsupported platforms, malformed explicit values, or origin mismatch must produce actionable errors, never a mutable alias/current-client fallback. Do not duplicate or guess Rex's undocumented platform path algorithms. The single Go-owned policy uses a present explicit override, then a nonempty execution-host REX_SERVER, then bounded `rex --autostart=false whoami --json` native discovery. Validate/capture the returned concrete endpoint and target it explicitly afterward; no multi-socket probing, autostart or runtime trials.

## Retained constraints and verification

Keep executable callbacks in the original Rex configuration Lua VM. Preserve authenticated single-use completion, connection/config-generation/expiry/snapshot/confirmation safeguards and owned-only bounded cleanup. No new daemon/transport framework, second Lua VM, host attach/WOL/SSH scope, init.lua takeover, configuration reload, push or global/destructive cleanup. Existing no-project-tests/no-TDD/no-experiments/no-GUI/live/remote-probe policy remains: formatting, static source review and compilation only. All runtime behavior stays unverified unless separately authorized later. Preserve tests/gates rather than weakening them; do not change Taskfile/lint/hook definitions to satisfy checks.

## Success evidence

Trace no-argument parsing to an actual production popup caller; trace platform endpoint selection and explicit-override precedence through actual Go and Lua consumers; independently source-review complete new builds; compile combined integration candidates; update user instructions to distinguish terminal UI from Rex floating popup and remove obsolete mandatory endpoint boilerplate. Record skipped/runtime gates as limitations, never passes.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDE0OjUyOjI0Ljg0N1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS91c2FiaWxpdHkvcmVxdWlyZW1lbnRzLWFwcHJvdmFsLzEiLCJpbnRlbnQiOiJiOTY3MTU2NmQzYTAyMzQyN2Q5MDBiMjhjYTdlOWZmNWMyY2U5OWU2ZTc5NzQzYmEzMDNjYTU0ZWYzODg2MWY3In0 -->
### 2026-11-01T14:52:24.847Z · approval · unspecified

{"approvedBy":"user","attempt":1,"decision":"approved","decisionEvidence":"User explicitly requested a new usability phase, task breakdown and implementation: just launching rex-whichkey should run the popup menu, with platform default sockets on Linux/macOS. Existing user-delegated bounded technical authority and retained execution/safety constraints apply; exact platform conventions/caller implementation remain source validations, not assumed facts.","eventId":"rex-whichkey/usability/requirements-approval/1","kind":"approval","recordRef":"prds/rex-whichkey-usability.md","requirements":["REQ-016","REQ-017"],"schemaVersion":1,"specHash":"sha256:11c2a7d29a1596d9c117d2c8f5495db615debec70eb65a56156a4525273eb258"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTAxVDE1OjE0OjU5LjI0MloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS91c2FiaWxpdHkvbmF0aXZlLWRpc2NvdmVyeS1yZXF1aXJlbWVudC1yZWZpbmVtZW50LzEiLCJpbnRlbnQiOiI4MjI1NmFkMzdlNjMyYTRiNjczYmNjYTA0NTA0ZTVhM2VjNzhiMmVkOWZkNGMyMjBkNmEyOGNlMjFhMmNkZDU2In0 -->
### 2026-11-01T15:14:59.242Z · approval · unspecified

{"approvedBy":"coordinator under explicit user delegation","attempt":2,"decision":"approved","decisionEvidence":"User goals remain no-argument popup and Linux/macOS defaults. Coordinator selects delegation to installed Rex native platform discovery, eliminating an unnecessary unverified duplicate path algorithm. Explicit override and origin safety retained; no runtime pass or personal user approval of this detail claimed.","eventId":"rex-whichkey/usability/native-discovery-requirement-refinement/1","kind":"approval","recordRef":"prds/rex-whichkey-usability.md","requirements":["REQ-016","REQ-017"],"schemaVersion":1,"specHash":"sha256:c2df9b9404ce7f2d649eed40d5c219f7098e9389c6407159b6ae35eddb7746c5"}
<!-- mtn:end-entry:v1 -->
