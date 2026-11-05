---
identity:
  kind: design
  uuid: 04e931a8-bf4f-4dfc-850e-fa020407ac02
scope:
  requirements:
    - REQ-002
    - REQ-003
    - REQ-005
    - REQ-011
    - REQ-013
    - REQ-014
    - REQ-016
    - REQ-017
lifecycle:
  state: approved
  sourceEventID: rex-whichkey/phase15-runtime-authority/1
  changedAt: '2026-11-05T01:00:11.219Z'
title: Code-owned runtime discovery and compatibility
mtnCreateRetry:
  key: rex-whichkey-runtime-ownership-design/create/1
  intent: 6fc0121415bec89b63d66bc7e1497f6cb6ade0a5726eb3f7020e2150492bf73e
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-05T00:33:01.681Z'
dateModified: '2026-11-05T01:00:11.838Z'
tags:
  - task
type: task
id: 32f131c5-948d-4c2d-9188-ecdd5cc16524
---
# Code-owned runtime discovery and compatibility

## Authority and scope

REXWHICHKEY-20261104-193124 — Accepted responsibility correction, captured now rather than asserted as the exact time of the earlier discussion. The user explicitly rejected runtime wiring in both init.lua and Nix: "it shouldn't even be in nix at all... it should all be in code", then requested "create a new phase to fix this". This approves the responsibility boundary and phase creation, not an unselected bootstrap algorithm, implementation completion, new runtime trials or a metadata migration.

The normal launch path must not require the user to specify helper paths, Rex CLI directories or matching binary/Lua versions. Runtime discovery and compatibility belong to application code. This refines REQ-002 (CLI ownership), REQ-003/005 (configuration and original-VM callbacks), REQ-011 (origin safety), REQ-013 (execution-host correctness), REQ-014 (packaging/configuration ownership), and REQ-016/017 (ordinary launch and native endpoint discovery). Existing REQ IDs are retained; no parallel requirements registry is introduced.

## Context and verified source facts

The public Go launcher and the server's configured Lua launcher can execute different helpers. The current Lua constructor accepts execution.helper and execution.rex_bin_dir; external declarative configuration supplies absolute paths. A local build can update the public caller while the registered action continues to use an older installed helper or dispatcher. The registered action currently takes precedence when present. Suppressed detached-controller/completion errors can turn a component mismatch into a flashing popup or a rename that is never applied.

The current configured rename path additionally requires compatible helper and Lua selection-text handling. Existing working-tree repairs are not an integrated or phase-complete baseline. Preserve them and reconcile the actual source revision before implementation; this design does not certify their runtime behavior.

## Accepted responsibility boundary

1. Application code owns selecting/validating the execution-host Rex CLI, selecting its own helper/controller/PTY child, and handling runtime component compatibility. Moving a literal from init.lua to Go or generating it in Nix is not dynamic discovery.
2. Nix/Home Manager may build and install the executable and normal assets. They must not implement runtime discovery, inject executable paths into Lua constructors, substitute installation paths into launch logic, or select runtime behavior on the application's behalf.
3. User init.lua owns only keybindings and optional mappings/callback configuration/registration. No helper, rex_bin_dir, socket literal, Nix-store path or repository-specific binary path is required for ordinary operation. Explicit server/client intent selectors remain distinct supported user choices.
4. Ordinary local development and packaged launch must select coherent runtime components without asking the user to pair a new executable with an old Lua module, repeatedly refresh a flake to repair runtime wiring, or copy source files into live configuration. Installing an actual software update remains normal packaging, not runtime discovery.
5. Preserve custom mappings and original-server-VM callbacks, exact endpoint/session/block identity, separately validated GUI targets, bounded ownership/cleanup, private capability transport, single-use completion, expiry and no uncertain retries. Do not mask incompatibility by silently bypassing a configured menu or reverting rename to Rex's dialogs.

## Implementation work to resolve in Phase 15

Inspect the existing Go/Lua boundary and documented Rex APIs/source. Choose the smallest supported code-owned bootstrap and compatibility mechanism. Reuse current-executable identity for self-spawn where appropriate and execution-host runtime metadata/environment/PATH only where its provenance is established. Do not guess a user's home directories, Nix store prefixes or remote client filesystem paths. A new Lua VM is not a replacement for the retained original callback registry.

The exact persistent-VM bootstrap, discovery under an app-launched restricted PATH, and asset/protocol compatibility strategy are implementation questions, not facts established by this record. Resolve them from supported contracts. If no safe supported mechanism satisfies the accepted boundary, report the precise missing capability rather than restoring path literals in configuration. Do not introduce a daemon, new socket transport, arbitrary code injection, automatic user-config edits/reloads or server restarts.

## Observable acceptance examples

- A normal public launch and keybinding launch need no configured helper or CLI-directory literals; user configuration contains only customization/registration.
- Rebuilding the local executable cannot inadvertently launch an older helper through a registered action. Packaged installation has the same code-owned behavior without Nix runtime substitution.
- A supported app-launched restricted-PATH context is resolved through supported code-owned provenance, not a hardcoded workaround. Missing or ambiguous required information fails before popup mutation with a bounded actionable diagnostic.
- Version/component mismatch cannot silently flash a popup, discard rename text, dispatch another action or require the user to maintain component pairings. No capability-bearing child output is exposed.
- Default tab/session rename remains a WhichKey page and completes against the retained resource; custom callbacks are preserved. Cancellation and uncertain failures never retry or redirect.
- Remote execution resolves server-host binaries/assets rather than a viewing Mac's paths. No new remote attach/bootstrap/WOL work is included.

## Verification and planning boundary

This is a documentation/phase-creation request. It authorizes no implementation dispatch, live configuration changes, deployment, test runners or new experiments. The project's existing no-tests/no-experiments policy remains unless the user separately changes it. Record the examples above as source-review/acceptance scenarios, not as passed tests or permission to run them. Preserve existing tests and quality configuration. Future implementation requires independent source review and may use already-authorized formatting/static/compilation checks; record actual results and remaining runtime uncertainty separately.

Reference requirements: prds/rex-whichkey.md and prds/rex-whichkey-usability.md. Historical startup choice: designs/rex-whichkey-native-startup.md. Phase: phases/rex-whichkey-15-code-owned-runtime.md. No existing records, phase identities or historical approval evidence are migrated or overwritten by this addition.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTA1VDAxOjAwOjEwLjQ5NFoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9waGFzZTE1LXJ1bnRpbWUtYXV0aG9yaXR5LzEiLCJpbnRlbnQiOiJjOWZmZjE2YzhlNzgyOGE5ZTU3NzM4YTZhNDgyOWExOTQyMTc3NDZjYWI5ZGFlZGIzM2NhYjEwNjA2NjcyNWQ4In0 -->
### 2026-11-05T01:00:10.494Z · approval · unspecified

{"approvedBy":"user, explicit execution instruction conveyed by authorized handoff","attempt":1,"authorityRef":".mtn-tmp/runtime-ownership/team/handoff.md","decision":"approved","decisionEvidence":"User: you are now the coordinator. use go-team to implement the phases. Handoff explicitly narrows execution to Phase 15 and its five fingerprinted tasks, retaining responsibility correction REXWHICHKEY-20261104-193124 and no-project-tests/no-experiments restrictions.","eventID":"rex-whichkey/phase15-runtime-authority/1","kind":"approval","recordRef":"designs/rex-whichkey-runtime-ownership.md","requirements":["REQ-002","REQ-003","REQ-005","REQ-011","REQ-013","REQ-014","REQ-016","REQ-017"],"schemaVersion":1,"scope":"Code-owned runtime responsibility and scoped implementation; no preselected bootstrap algorithm, unrelated phase reopening, historical verification, source checkpoint or deployment approval.","specHash":"sha256:253d40dc6204148a88924f7c860ad092f86f3b541a618330f112c6202b3aaa65"}
<!-- mtn:end-entry:v1 -->
