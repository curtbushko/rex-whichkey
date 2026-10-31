---
workflowKind: prd
title: Rex WhichKey initial-release requirements
mtnCreateRetry:
  key: rex-whichkey-prd-registration-v1
  intent: 500ec2461fa8ce92618fb671a8a13b8155db6b3006c156f27f2f766f5ddee1a0
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-10-30T12:39:18.289Z'
dateModified: '2026-10-31T11:29:10.264Z'
tags:
  - task
type: task
id: 195cbbe0-b607-4419-9c86-cd9b70075568
---
# Current requirements authority

## Current direction — REXWHICHKEY-20261031-072107

User-directed replanning: which-key behavior is the product; Rex is a documented API dependency, not a research milestone. Stop further tests and experiments. Do not launch probes, benchmarks, test runners, GUI exercises, remote trials or temporary config/reload procedures. This revision schedules none, including at release. Existing results and untested limitations remain accurately recorded, not converted into passes. Product behavior and safety requirements remain; previous experimental/RFC-before-implementation barriers are superseded. Existing test files are preserved. This planning update does not execute implementation or approve provisional IF-* signatures.

## Impact on earlier requirements

Retain the functional intent and stable IDs REQ-001–REQ-015. This later explicit user correction supersedes earlier requirements to perform experiments, real-Rex/remote validation, mandatory test execution and a tested/fingerprint-approved final RFC before product planning. REQ-015 now tracks integrated delivery, documentation and honestly recorded limitations; it does not authorize another acceptance runner. REQ-003/005/008/010/011/012/013 do not require advance API/GUI/remote experiments. The previously final-only disconnected-origin test is no longer scheduled or mandatory under this stop instruction; it remains untested. No no-fallback behavior is dropped from implementation intent and no pass is invented.

Read published API contracts and implement practical menu integration. Preserve origin and cleanup safety, user configuration ownership, simple architecture, callback execution in Rex Lua and all other functional requirements. Exact proposed interfaces in the archived source are still proposals, not newly approved contracts.

## Preserved prior requirements and source

The text below is retained verbatim as historical requirements/source context. Where its workflow, testing or research gates conflict with the current amendment above, the later user instruction controls.

# Registered initial-release requirements

## Authority and interpretation

The user approved registration and the initial-release scope in this conversation by replying "yes" to the coordinator's four-point scope/registration request. This registers project rex-whichkey, assigns stable IDs, defers remote connection/Wake-on-LAN, and leaves all open proposals unapproved. Approval covers requirements only, not final interfaces, a delivery plan, or implementation dispatch. Manual approval is the default. The original source follows verbatim below for provenance; its draft/proposal labels remain in force for design details. This requirement index does not approve illustrative APIs or layout.

## Authorized workflow-recovery development

**✓ Agreed — REXWHICHKEY-20261030-121414 (Accepted for development and independent review only).** The user instructed "update the PRD to authorize that" in response to the proposal to develop and independently review a narrowly scoped evidence-preserving workflow-recovery feature. This extends the earlier investigation-only authority; it does not grant activation or use authority.

The coordinator may arrange development and independent review of the upstream workflow-tooling support needed to recover the malformed typed PRD approval evidence. The purpose is to unblock the existing delivery workflow without losing evidence or weakening its approval checks. This is a workflow prerequisite, not an additional rex-whichkey product requirement or expansion of REQ-001 through REQ-015.

Required acceptance constraints:

- Preserve the original malformed approval note, exact content bytes, native metadata and intent. Do not overwrite history, hand-edit managed evidence, replace the collection or silently reinterpret the original as valid.
- Limit recovery to an explicitly authorized target, bound to its record, event identity, native intent and exact content digest. Unrelated malformed, ambiguous, conflicting or unsupported evidence must continue to fail closed.
- Recovery alone must not grant document, design, plan, dispatch or completion approval. Any replacement approval must independently validate, identify its actual authority, match the current document fingerprint and retain its original limited scope.
- Require failing-first preservation, exact-binding, replay/idempotency and negative authorization tests, plus fresh independent review of the proposed tooling change. Detailed candidate mechanics and review tests remain in the project proposal and audit evidence; they are not finalized interfaces established by this PRD.

**Not authorized by this amendment:** installation or activation of the feature; applying it to this or another collection; migration, quarantine/reclassification of existing evidence, replacement approval or resolution of the registration blocker. Activation/use requires completed review and separate explicit user approval. No provider configuration changes, Rex restart/reload or user configuration mutation is authorized.

The existing registration blocker remains unresolved until a supported recovery is separately approved, applied and verified. Strict Unicode/input, graphical geometry/app routing, production security/lifecycle, remote-context and fingerprint-bound final-design approval gates remain intact. No implementation dispatch, task completion or final RFC approval is established by this amendment. Downstream document/plan fingerprints must be reassessed before future dispatch; old draft references must not be silently treated as current approvals.

## Requirements and acceptance criteria

### REQ-001 — Rewrite and process lifecycle
Implement a fresh, short-lived Go application with a minimal Rex Lua launcher. Invocation opens the menu; dismissal or completed dispatch terminates it. No daemon, reuse of the old architecture, hexagonal architecture, service hierarchy, or unnecessary directory layering. Verify lifecycle/end-to-end cleanup and review the simple, relatively flat package structure. Source: sections 1–3.

### REQ-002 — Typed Rex CLI client
Provide reusable concrete Client, Session, Window and Block types, with API identities/parameters grounded in published Rex schemas, resource methods and a generic API-call escape hatch. Centralize argument-array CLI execution, JSON decoding, targeting and errors; do not interpolate shell commands. Preserve session identity when blocks move. Verify schema fixtures, argument safety, invalid JSON, CLI failures, cancellation, timeouts and server/session selection. Exact signatures remain subject to design approval. Source: sections 3, 5, 9.

### REQ-003 — Lua configuration loading
Support inline init.lua configuration and an imported separate Lua file without integration/config module name collisions. Validate loading and precedence against real Rex; names and exact loader API remain open. Verify both configuration paths and documented precedence. Source: sections 6, 10b.

### REQ-004 — Mapping compilation and validation
Support full key sequences, implicit prefix groups, nested inherited attributes, descriptions, groups, icons, named actions and args; retain hidden-but-active entries, boolean/function conditions and function-valued strings. Support <Space> and <lt>; reject duplicate actions, group/action conflicts and unsupported key notation. Verify positive/negative mapping fixtures against the old functional baseline. Source: section 6.

### REQ-005 — Rex-runtime callbacks
Execute user-defined Lua callbacks in the originating Rex Lua runtime, with context and args; never serialize executable functions as ordinary JSON for Go execution. Verify a callback end to end. Launch/registration/dispatch transport and responsibility boundaries require approved design and validation before implementation planning. Source: sections 4, 6, 10a.

### REQ-006 — Menu navigation
Retain arbitrary-depth submenus, breadcrumbs, Esc/Backspace navigation, conditional q handling and Tab pagination. Verify state transitions and key behavior, including hidden active entries. Source: section 7.

### REQ-007 — Destructive-action confirmation
For confirm-marked actions, y confirms; cancellation does not dispatch. Verify named-action and callback confirmation/cancellation paths, including destructive Rex actions. Source: sections 6–7.

### REQ-008 — Geometry and rendering
Support top, bottom, left, right and center; adaptive columns, Unicode display widths, constrained geometry, paging, icons and help/footer. Validate native-layer clipping in the targeted Rex version rather than copying the old overlay workaround. Verify narrow/small terminal cases and all five positions with real-client visual/e2e checks. Source: sections 7, 10c.

### REQ-009 — Theme and terminal safety
Use available terminal foreground/background/palette, fallbacks and contrast handling; respect NO_COLOR and ASCII mode. Restore terminal state on normal dismissal and errors. Verify palette/fallback/no-color/ASCII fixtures plus terminal lifecycle behavior. Source: sections 7, 9.

### REQ-010 — Built-in actions and session parity
Retain sessions (new/switch/previous/next/rename/close), tabs (new/previous/next/rename/goto/close), split, directional focus, zoom, sidebar, command palette and pane close. Preserve old dynamic/session behavior by observing the old plugin, not assuming an enumerated-session UI. Verify each action against the actual API and real Rex integration; resolve exact dynamic-session acceptance before dependent task planning. Source: sections 3j, 7, 10e.

### REQ-011 — Origin context and focus
Target the originating Rex block/client/session as appropriate, restore focus before dispatch, and clean up popup resources after errors. Verify moved blocks, focus restoration, split/window operations, dismissal and error cleanup in real Rex. Source: sections 5, 7, 9.

### REQ-012 — Rex-known remote host listing
Discover and list only remote hosts already known to Rex; do not introduce a Tailscale inventory or independent registry. Validate the actual CLI/server discovery mechanism before planning its implementation; verify controlled outputs and actual Rex discovery. Listing does not authorize remote attach/bootstrap/WOL flows. Source: sections 2, 8, 10d.

### REQ-013 — Remote-server compatibility
Operate with Rex on another host without client-local absolute-path assumptions; preserve originating IDs and server/session selection. Verify remote-context targeting and configuration/launch behavior with real remote Rex where available. Unavailable infrastructure is an explicit unmet gate, not a mock-only pass. Source: sections 7–9.

### REQ-014 — Packaging and user configuration ownership
Provide Nix flake and Home Manager integration without owning or overwriting the user's Rex init.lua. Verify package build and Home Manager configuration alongside an existing user setup. Source: section 7.

### REQ-015 — Integration and release verification
Use controlled CLI fixtures and automated mapping/navigation/rendering tests, complemented by actual Rex end-to-end tests for popup creation/dismissal, focus, split/window operations and cleanup. Test remote context and packaging; do not substitute coverage percentages or mocks for real-server gates. Each implementation task includes its own tests; final gates run on the integrated target. Source: section 9.

## Deferred scope

Remote-host connection/attachment, reuse-or-create remote sessions, reachability protocols, relay-assisted Wake-on-LAN, SSH bootstrap/fallback and retries are deferred. They require separately approved requirements and design, not placeholder implementation phases.

## Mandatory design validation before implementation planning

An RFC is required because Lua-to-Go launch/config/callback boundaries and remote targeting are material public-interface/security concerns. Resolve source 10a–10g: callback transport/lifecycle; separate-file loading and precedence; native-layer geometry; known-host discovery; observed dynamic-session parity; concrete-package boundaries; and TUI framework versus direct I/O. Record the tested Rex version, CLI/API contracts and evidence, and obtain fingerprint-bound design approval. Do not treat source illustrations as contracts.

## Original source (verbatim)

# RFC: Rex WhichKey — Go Rewrite

**Status:** Draft — records accepted decisions and identifies unresolved details  
**Date:** 2026-10-27  
**Project:** rex-whichkey

## 1. Summary

Rex WhichKey will be rewritten from scratch in Go. The existing [`rex-whichkey.old`](https://github.com/curtbushko/rex-whichkey.old) repository defines the initial functional baseline, **not** the architecture or code to reuse. The implementation will be a short-lived Go process launched from a minimal Rex Lua integration, with a relatively flat project structure and no hexagonal architecture.

The project will offer a WhichKey-style contextual keyboard menu for Rex. It will use the Rex Server API through the `rex` CLI, preserve Lua-based keymap configuration and user-defined callbacks, and eventually support remote Rex sessions and Wake-on-LAN workflows.

## 2. Goals and scope

### Initial release goals

- Match the existing behavior of `rex-whichkey.old` as closely as practicable.
- Rewrite the implementation in Go without carrying over its architecture or technical debt.
- Launch the application on demand as a short-lived process; exit when dismissed or after dispatch.
- Provide a minimal Rex Lua launcher for registering the invocation and starting the Go process.
- Support both inline Lua configuration and a separate Lua configuration file.
- Preserve a WhichKey v3-inspired mapping specification, including named Rex actions and user-defined Lua callbacks.
- Provide a typed, reusable Go Rex client centered on concrete `Client`, `Session`, `Window`, and `Block` types, corresponding to Rex's resource model.
- Call the Rex CLI for Server API operations, using the published [Server API](https://www.superlogical.com/rex/docs/automate/server-api) and [CLI](https://www.superlogical.com/rex/docs/automate/cli) contracts.
- Preserve existing session-navigation features.
- List remote hosts already known to Rex rather than maintaining a separate host registry.

### Future direction

Support connecting to remote Rex sessions, reusing an existing session or creating one when absent. Later integration may add machine reachability checks, relay-assisted Wake-on-LAN, and SSH-based bootstrap or fallback as appropriate. This is a planned direction, not an approved detailed protocol or workflow.

### Non-goals

- Reusing the old Lua/Bash architecture, transport, rendering pipeline, or configuration snapshot protocol as requirements.
- Running a daemon or persistent coordinator for the initial application.
- Imposing hexagonal architecture, services, drivers/adapters directories, or an interface hierarchy.
- Maintaining a separate remote-host inventory when Rex already knows the hosts.
- Reimplementing the Rex server or replacing its existing API.

## 3. Accepted design decisions

The entries below capture explicit choices made during discussion. No retrospective historical timestamp IDs are assigned; discussion references are used only to organize this RFC. A formal decision register can assign new immutable IDs at the time it is actually established.

| Ref | Status | Decision |
| --- | --- | --- |
| 3a | ✓ Agreed | Rewrite Rex WhichKey from scratch in Go; use `rex-whichkey.old` as a functional reference, not an architectural basis. |
| 3b | ✓ Agreed | The Go program is short-lived, invoked when the popup is needed; no initial persistent service. |
| 3c | ✓ Agreed | A minimal Rex Lua integration is required to launch the Go executable. |
| 3d | ✓ Agreed | Keep the Go project simple and relatively flat; do not use hexagonal architecture. |
| 3e | ✓ Agreed | Use concrete Rex resource types — `Client`, `Session`, `Window`, `Block` — rather than a `*Service` hierarchy. |
| 3f | ✓ Agreed | Model the Go Rex client from Rex's Server API and map operations cleanly to calls through the Rex CLI. |
| 3g | ✓ Agreed | Use Lua for WhichKey configuration, with WhichKey-style mapping definitions. |
| 3h | ✓ Agreed | Support configuration directly in Rex `init.lua` and via an imported separate Lua file. |
| 3i | ✓ Agreed | Support both named Rex actions and user-defined Lua callbacks. |
| 3j | ✓ Agreed | Preserve the existing session functionality and list remote hosts available in Rex. |

## 4. Proposed architecture

This section describes implementation proposals consistent with the accepted constraints; these exact file boundaries and interfaces are **not yet accepted decisions**.

```text
Rex keybinding
    |
    v
Minimal Rex Lua launcher
    |
    v
Short-lived rex-whichkey Go process
    |-- configuration / mapping navigation
    |-- terminal input and rendering
    |-- action selection and confirmation
    `-- rex package (typed Rex API client via CLI)
    |
    v
Rex CLI / Server API
```

The launcher should register the Rex action and create/launch a popup terminal. It should not own the menu navigation or rendering. Executable Lua callbacks cannot be serialized as plain JSON; a small Rex-side action registration/dispatch mechanism is consequently expected, but its protocol and exact responsibility split require design and testing.

### Candidate layout

```text
rex-whichkey/
├── cmd/rex-whichkey/main.go
├── rex/
│   ├── rex.go
│   ├── session.go
│   ├── window.go
│   ├── block.go
│   ├── layer.go
│   └── layout.go
├── mapping.go
├── menu.go
├── action.go
├── render.go
├── theme.go
├── config.go
├── plugin.lua
├── nix/
└── go.mod
```

This tree is illustrative. Go package placement must remain valid (e.g., a `main` package must reside in an executable directory), and the final layout should minimize unnecessary directories and exported APIs.

## 5. Go Rex client design

### Principles

- Prefer ordinary concrete objects and methods (`Client`, `Session`, `Window`, `Block`) over `Client → SessionService → ...` chains.
- Model identities, API parameters, and response data using types grounded in Rex's published schemas.
- Use the `rex` CLI to call methods; avoid embedding shell scripts or interpolating untrusted input into shell commands.
- Keep CLI argument construction, JSON decoding, error handling, and server/session targeting in the `rex` package.
- Keep a low-level generic API-call escape hatch for methods not yet wrapped.
- Avoid introducing transport/service interfaces or code generation until a demonstrated need arises.
- Use small interfaces at call sites only where genuinely helpful for tests or composition.
- Keep session identity independent of temporary window membership, especially for blocks that may move.

Illustrative intended ergonomics (not finalized method signatures):

```go
client := rex.New()
sessions, err := client.Sessions(ctx)
if err != nil { return err }
windows, err := sessions[0].Windows(ctx)
if err != nil { return err }
blocks, err := windows[0].Blocks(ctx)
if err != nil { return err }
return blocks[0].Focus(ctx)
```

The package should be checked against the actual server method descriptions and CLI argument syntax. Do not infer request fields from convenient Go method names.

## 6. Configuration

### Agreed configuration model

- Lua is the user-facing configuration language.
- Both inline specification in `init.lua` and an imported configuration file are supported.
- Use a WhichKey v3-inspired list of mapping entries with full key sequences, groups, descriptions, icons, action names, callback functions, and action arguments.
- Preserve callback execution in the Rex Lua runtime. The Go process should never pretend a Lua function can be transmitted to Go as ordinary JSON and invoked there.

Illustrative syntax:

```lua
whichkey.setup({
  position = "center",
  spec = {
    { "w", group = "Sessions" },
    { "wn", "session.new", desc = "New session" },
    { "z", "pane.zoom", desc = "Zoom pane" },
    { "x", "pane.close", desc = "Close pane", confirm = true },
    {
      "h",
      function(ctx, args)
        -- Custom Rex Lua action
      end,
      desc = "Custom action",
    },
  },
})
```

The setup/module name and loader paths here are illustrative, not yet an API promise. To prevent import name collisions, the shipped integration module and user configuration file should have unambiguous names.

### Original mapping behavior to retain

The existing plugin's map compiler supports:

- Full key sequences and implicit prefix groups.
- Nested attribute scopes with inherited properties.
- Named actions and Lua callbacks.
- Conditional entries (`cond` boolean or function).
- Dynamic string-valued descriptions, group labels, and icons (functions).
- Hidden entries that still respond to keys.
- Named action arguments (`args`).
- Confirmation for destructive operations (`confirm`).
- Validation of duplicate action mappings, group/action conflicts, and unsupported key notation.
- `<Space>` and `<lt>` representations.

These are functional acceptance targets. Their new internal representation is unconstrained by the old implementation.

## 7. Initial feature-parity inventory

The inventory below is based on the public old repository's code and documentation, not on claims that the new implementation already works.

| Area | Initial functional requirement |
| --- | --- |
| Menu navigation | Arbitrary-depth submenus, breadcrumbs, Esc/Backspace navigation, conditional `q`, Tab pagination. |
| Destructive actions | Confirm with `y`; cancel without dispatch. |
| Positions | Top, bottom, left, right, and center. |
| Rendering | Terminal-oriented adaptive columns, Unicode display widths, constrained geometry, paging, icons, help/footer. |
| Theme | Respect terminal foreground/background/palette where available, fallbacks, contrast handling, `NO_COLOR`, ASCII mode. |
| Rex actions | Sessions (new/switch/previous/next/rename/close), tabs (new/previous/next/rename/goto/close), split, directional focus, zoom, sidebar, command palette, pane close. |
| Context safety | Act on the originating Rex block/client/session as appropriate; restore focus before dispatch. |
| Packaging | Nix flake and Home Manager integration without taking ownership of the user's Rex `init.lua`. |
| Remote compatibility | Correct behavior when Rex is running on another host; no client-local absolute path assumptions. |

The old implementation documents a floating-layer clipping problem on Rex Beta build 1101 and a full-window modal overlay workaround. This is **evidence to test**, not a decision to transplant that implementation. Native layer geometry and actual behavior must be validated against the Rex version targeted by the rewrite.

### Dynamic session behavior

The original configuration offers `session.switch` through Rex, not necessarily a dynamically expanded enumerated list of session names in the popup. The user requires preservation of current dynamic/session behavior; the exact UI parity should be tested against the old plugin rather than assumed from an action name.

## 8. Remote hosts and sessions

### Agreed

- Discover/list only the remote hosts Rex already knows about for the initial remote-host menu.
- Do not introduce an independent Tailscale host inventory for that purpose.

### Future direction (not yet detailed or accepted)

When a user selects a remote host, an eventual flow could:

1. Inspect server availability.
2. Show existing Rex sessions.
3. Attach to a selected session or create a session if appropriate.
4. If unreachable, optionally use a configured relay for Wake-on-LAN, following capabilities from the user's `tailssh` script.
5. Retry discovery/attachment and provide useful failure feedback.

The implementation must distinguish a reachable host from a reachable Rex service and from a specific existing session. The CLI's remote-server selection should remain the Rex client's primary path where possible; SSH/WOL details are unresolved and not required to recreate the old menu.

## 9. Testing and quality gates

Prioritize real integration and end-to-end behavior, not test coverage percentages.

- Test typed Rex calls against known API schemas and controlled CLI output.
- Use argument-array CLI invocation; verify escaping, JSON decoding, server/session selection, timeouts, cancellation, and errors.
- Exercise mapping parsing, inherited attributes, hidden keys, conditions, and callback references.
- Verify navigation, confirmations, pagination, Unicode width, themes, and terminal-state restoration.
- Run actual Rex integration tests for popup creation/dismissal, focus restoration, split/window operations, and cleanup after errors.
- Test remote-server context carefully, including the originating block/client IDs.
- Ensure Nix packaging and Home Manager configuration do not overwrite the user's existing Rex setup.

A mock CLI is useful for unit tests but is not a substitute for running the real Rex server where available.

## 10. Open proposals and validations

These items must not be interpreted as approved decisions.

| Ref | Status | Item | Resolution path |
| --- | --- | --- | --- |
| 10a | ◇ Proposal | Keep the minimal Lua callback dispatcher separate from all Go menu/navigation logic. | Specify the minimal launch/dispatch contract and test against Rex. |
| 10b | ? Validate | Exact API for loading a separate user mapping file while also permitting an inline `setup({ spec = ... })`. | Test Rex's `require`/`dofile` lookup and document precedence. |
| 10c | ? Validate | Whether native floating layers are now correctly clipped in target Rex versions. | Real-client visual/e2e test in all five positions. |
| 10d | ? Validate | Exact list/discovery mechanism for Rex-configured remote hosts in the CLI and Go API. | Inspect current `rex` CLI/server API output. |
| 10e | ? Validate | What “dynamic sessions” means beyond old `session.switch` behavior. | Compare interactive behavior of the Lua plugin with acceptance tests. |
| 10f | ◇ Proposal | Keep the Go Rex package in one package with concrete resource methods and no service hierarchy. | Review as implementation evolves; avoid unnecessary layers. |
| 10g | ? Validate | Whether the Go TUI should use a third-party framework or direct terminal I/O. | Build a minimal popup spike measuring startup and interaction. |

## 11. References

- [rex-whichkey.old](https://github.com/curtbushko/rex-whichkey.old) — current functional reference.
- [which-key.nvim](https://github.com/folke/which-key.nvim) — mapping syntax and interaction inspiration.
- [rexdiff/rex](https://github.com/rednafi/rexdiff/tree/master/rex) — Go examples of Rex CLI/API calls.
- [go-libghostty](https://github.com/mitchellh/go-libghostty) — preferred Go API/style reference.
- [Rex Server API](https://www.superlogical.com/rex/docs/automate/server-api) — authoritative API reference.
- [Rex CLI](https://www.superlogical.com/rex/docs/automate/cli) — CLI invocation reference.
- [tailssh](https://github.com/curtbushko/nixos-config/blob/main/modules/home/scripts/tailssh) — inspiration for future remote bootstrap/Wake-on-LAN.

---

**Decision authority:** This RFC documents existing explicit agreements. Its proposals and validation questions do not become accepted merely by inclusion in the RFC.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMwVDEyOjM5OjQ0LjU1NloiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9wcmQtYXBwcm92YWwvMSIsImludGVudCI6ImJmZjZhODY2Y2ZkOGUxMzIzYTAyY2UwYmEwNjM3MDEzY2Q4MjUyMDg1NTQ2MTY3YzIxNjIwOTkxNDM3OTJiMzIifQ -->
### 2026-10-30T12:39:44.556Z · approval · unspecified

{"schemaVersion":1,"eventId":"rex-whichkey/prd-approval/1","kind":"approval","recordRef":"prds/rex-whichkey.md","specHash":"sha256:2b57ebed7620171e9b1528c160dc374314926b89e021f1f920c77783638eccd8","attempt":1,"decision":"approved","approvedBy":"user","authority":"User in current conversation","requirements":["REQ-001","REQ-002","REQ-003","REQ-004","REQ-005","REQ-006","REQ-007","REQ-008","REQ-009","REQ-010","REQ-011","REQ-012","REQ-013","REQ-014","REQ-015"],"decisionEvidence":{"prompt":"Approve registration in .tasks for rex-whichkey; assign stable REQ IDs to initial-release requirements; defer remote connection/Wake-on-LAN; keep open proposals unapproved and require interface validation before implementation planning.","response":"yes"},"scope":"Initial-release requirements and registration only; no design or delivery-plan approval; manual approval policy"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMwVDE2OjQxOjIxLjA5NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS93b3JrZmxvdy1yZWNvdmVyeS1hdXRob3JpemF0aW9uLzEiLCJpbnRlbnQiOiJlZDI5YTY5Y2FkYjJhNTZhMTA0ZjBmZDUxYTIxZTYxMGJmOWNlODkyMjQ5OTVjODM5ZmNiMWNlNmE0ZmRlNGUzIn0 -->
### 2026-10-30T16:41:21.095Z · approval · unspecified

{
  "schemaVersion": 1,
  "eventId": "rex-whichkey/workflow-recovery-authorization/1",
  "kind": "approval",
  "recordRef": "prds/rex-whichkey.md",
  "specHash": "sha256:2179140873bca50f43022deea7d262ce148ba87a2052d296ecf7fc84ee6852db",
  "attempt": 1,
  "decision": "approved",
  "approvedBy": "user",
  "authority": "User in current coordinator conversation",
  "decisionEvidence": "User: update the PRD to authorize that. Immediate antecedent proposal 2a, REXWHICHKEY-20261030-121414: authorize development and independent review of narrowly scoped evidence-preserving workflow recovery; activation/use awaits review and explicit approval.",
  "decisionId": "REXWHICHKEY-20261030-121414",
  "scope": "Approve this PRD authorization amendment for development and independent review only. Original product requirements unchanged. Not installation, activation/use, evidence repair, migration, replacement approval, blocker resolution, final design/plan approval or implementation dispatch.",
  "recordedAt": "2026-10-30T16:41:20Z",
  "priorSpecHash": "sha256:2b57ebed7620171e9b1528c160dc374314926b89e021f1f920c77783638eccd8",
  "preservation": "Original malformed rex-whichkey/prd-approval/1 is unchanged. This append records new authority only; it cannot repair the old invalid event or make the PRD workflow snapshot valid."
}

<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMwVDIwOjUxOjIzLjU0N1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVjb3ZlcnkiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9wcmQtYXBwcm92YWwtcmVjb3ZlcnkvMSIsImludGVudCI6ImVkNmM3MGZiMDAyNmVhNzYyNjI1ZmE3MjA1YjRkNmYyMzQyODQ5ZDJlNGVhNDQxNDVmMTk0YmVmNmIxZWQ2ZGEifQ -->
### 2026-10-30T20:51:23.547Z · recovery · unspecified

{"approvedBy":"user","attempt":1,"collectionRoot":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.tasks","decision":"approved","decisionEvidence":"Current coordinator handoff explicitly instructs: Resume Phase 1 using that newly installed tooling and the existing narrowly scoped approval-recovery authorizations; use native guarded operations to preserve original approval history. User rebuilt the reviewed tooling and requested try phase1 again. Authority applies only to the named malformed rex-whichkey/prd-approval/1 and separate registration-only replacement, not unrelated approvals or tests.","disposition":"quarantine-malformed-approval","eventId":"rex-whichkey/prd-approval-recovery/1","kind":"recovery","recordRef":"prds/rex-whichkey.md","relatedDecisionId":"REXWHICHKEY-20261030-121414","replacementEventId":"rex-whichkey/prd-approval-replacement/1","schemaVersion":1,"scope":"Exact append-only quarantine of one malformed approval only; preserves original bytes/envelope, no approval by quarantine. Current and original fingerprints explicitly bound.","sourceRunId":"rex-whichkey-p01-coordinator-a03","specHash":"sha256:2179140873bca50f43022deea7d262ce148ba87a2052d296ecf7fc84ee6852db","target":{"contentDigest":"sha256:040e15c15bac643ad9cb5618863f0d747e2baebd583aee61422b047164e70227","eventId":"rex-whichkey/prd-approval/1","intent":"bff6a866cfd8e1323a02ce0ba0637013cd8252085546167c2162099143792b32","specHash":"sha256:2b57ebed7620171e9b1528c160dc374314926b89e021f1f920c77783638eccd8"}}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMwVDIwOjUxOjI4Ljk3MVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS9wcmQtYXBwcm92YWwtcmVwbGFjZW1lbnQvMSIsImludGVudCI6ImY0NTZlNzRiZTZiYTNlY2I2ZjM4MmYxYzk4ODRhMjNkMzA0OGQzOGQ1YzQ2ZDk2YTZkMmRkZDQzZGEzODg2ZDUifQ -->
### 2026-10-30T20:51:28.971Z · approval · unspecified

{"amendmentAuthorityEventId":"rex-whichkey/workflow-recovery-authorization/1","approvedBy":"user","attempt":1,"authority":"Original explicit user registration approval; separate approved development/review amendment; current exact-target approval-recovery instruction","decision":"approved","decisionEvidence":"Original registration request: approve registration in .tasks for rex-whichkey; assign stable REQ IDs; defer remote connection/Wake-on-LAN; keep open proposals unapproved and require interface validation before implementation planning. User response: yes. Current PRD amendment rex-whichkey/workflow-recovery-authorization/1 separately records user: update the PRD to authorize that, limited to development and independent review. Current handoff explicitly authorizes native guarded approval recovery using reviewed installed tooling and a separate replacement retaining original registration-only scope after user rebuilt config and requested try phase1 again.","eventId":"rex-whichkey/prd-approval-replacement/1","kind":"approval","priorSpecHash":"sha256:2b57ebed7620171e9b1528c160dc374314926b89e021f1f920c77783638eccd8","recordRef":"prds/rex-whichkey.md","recoveryEventId":"rex-whichkey/prd-approval-recovery/1","requirements":["REQ-001","REQ-002","REQ-003","REQ-004","REQ-005","REQ-006","REQ-007","REQ-008","REQ-009","REQ-010","REQ-011","REQ-012","REQ-013","REQ-014","REQ-015"],"schemaVersion":1,"scope":"Registration and unchanged REQ-001 through REQ-015 only; existing development/review amendment authority remains separately attributed. No final design/interface/framework adoption, project/phase plan, implementation dispatch, GUI/remote permission, test waiver, completion or blanket recovery authority.","sourceRunId":"rex-whichkey-p01-coordinator-a03","specHash":"sha256:2179140873bca50f43022deea7d262ce148ba87a2052d296ecf7fc84ee6852db"}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDExOjI0OjU0LjUzM1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiaGlzdG9yaWNhbCIsInJldHJ5S2V5IjoicmV4LXdoaWNoa2V5L3doaWNoa2V5LWZpcnN0L3ByZXNlcnZlL3ByZHNfX3JleC13aGljaGtleS5tZCIsImludGVudCI6IjMxYjQzYTQ1MTY4OGIyZDNiNmYwODA3MDk2NDY1YzA1M2Q5Y2FiMmNmNDNhZGNmMzFlOWNiZWY1MzQ0ZDkzOWEifQ -->
### 2026-10-31T11:24:54.533Z · historical · unspecified

{"schemaVersion":1,"eventId":"rex-whichkey/whichkey-first/preserve/prds__rex-whichkey.md","kind":"historical","recordRef":"prds/rex-whichkey.md","specHash":"sha256:2179140873bca50f43022deea7d262ce148ba87a2052d296ecf7fc84ee6852db","attempt":1,"content":"# Registered initial-release requirements\n\n## Authority and interpretation\n\nThe user approved registration and the initial-release scope in this conversation by replying \"yes\" to the coordinator's four-point scope/registration request. This registers project rex-whichkey, assigns stable IDs, defers remote connection/Wake-on-LAN, and leaves all open proposals unapproved. Approval covers requirements only, not final interfaces, a delivery plan, or implementation dispatch. Manual approval is the default. The original source follows verbatim below for provenance; its draft/proposal labels remain in force for design details. This requirement index does not approve illustrative APIs or layout.\n\n## Authorized workflow-recovery development\n\n**✓ Agreed — REXWHICHKEY-20261030-121414 (Accepted for development and independent review only).** The user instructed \"update the PRD to authorize that\" in response to the proposal to develop and independently review a narrowly scoped evidence-preserving workflow-recovery feature. This extends the earlier investigation-only authority; it does not grant activation or use authority.\n\nThe coordinator may arrange development and independent review of the upstream workflow-tooling support needed to recover the malformed typed PRD approval evidence. The purpose is to unblock the existing delivery workflow without losing evidence or weakening its approval checks. This is a workflow prerequisite, not an additional rex-whichkey product requirement or expansion of REQ-001 through REQ-015.\n\nRequired acceptance constraints:\n\n- Preserve the original malformed approval note, exact content bytes, native metadata and intent. Do not overwrite history, hand-edit managed evidence, replace the collection or silently reinterpret the original as valid.\n- Limit recovery to an explicitly authorized target, bound to its record, event identity, native intent and exact content digest. Unrelated malformed, ambiguous, conflicting or unsupported evidence must continue to fail closed.\n- Recovery alone must not grant document, design, plan, dispatch or completion approval. Any replacement approval must independently validate, identify its actual authority, match the current document fingerprint and retain its original limited scope.\n- Require failing-first preservation, exact-binding, replay/idempotency and negative authorization tests, plus fresh independent review of the proposed tooling change. Detailed candidate mechanics and review tests remain in the project proposal and audit evidence; they are not finalized interfaces established by this PRD.\n\n**Not authorized by this amendment:** installation or activation of the feature; applying it to this or another collection; migration, quarantine/reclassification of existing evidence, replacement approval or resolution of the registration blocker. Activation/use requires completed review and separate explicit user approval. No provider configuration changes, Rex restart/reload or user configuration mutation is authorized.\n\nThe existing registration blocker remains unresolved until a supported recovery is separately approved, applied and verified. Strict Unicode/input, graphical geometry/app routing, production security/lifecycle, remote-context and fingerprint-bound final-design approval gates remain intact. No implementation dispatch, task completion or final RFC approval is established by this amendment. Downstream document/plan fingerprints must be reassessed before future dispatch; old draft references must not be silently treated as current approvals.\n\n## Requirements and acceptance criteria\n\n### REQ-001 — Rewrite and process lifecycle\nImplement a fresh, short-lived Go application with a minimal Rex Lua launcher. Invocation opens the menu; dismissal or completed dispatch terminates it. No daemon, reuse of the old architecture, hexagonal architecture, service hierarchy, or unnecessary directory layering. Verify lifecycle/end-to-end cleanup and review the simple, relatively flat package structure. Source: sections 1–3.\n\n### REQ-002 — Typed Rex CLI client\nProvide reusable concrete Client, Session, Window and Block types, with API identities/parameters grounded in published Rex schemas, resource methods and a generic API-call escape hatch. Centralize argument-array CLI execution, JSON decoding, targeting and errors; do not interpolate shell commands. Preserve session identity when blocks move. Verify schema fixtures, argument safety, invalid JSON, CLI failures, cancellation, timeouts and server/session selection. Exact signatures remain subject to design approval. Source: sections 3, 5, 9.\n\n### REQ-003 — Lua configuration loading\nSupport inline init.lua configuration and an imported separate Lua file without integration/config module name collisions. Validate loading and precedence against real Rex; names and exact loader API remain open. Verify both configuration paths and documented precedence. Source: sections 6, 10b.\n\n### REQ-004 — Mapping compilation and validation\nSupport full key sequences, implicit prefix groups, nested inherited attributes, descriptions, groups, icons, named actions and args; retain hidden-but-active entries, boolean/function conditions and function-valued strings. Support <Space> and <lt>; reject duplicate actions, group/action conflicts and unsupported key notation. Verify positive/negative mapping fixtures against the old functional baseline. Source: section 6.\n\n### REQ-005 — Rex-runtime callbacks\nExecute user-defined Lua callbacks in the originating Rex Lua runtime, with context and args; never serialize executable functions as ordinary JSON for Go execution. Verify a callback end to end. Launch/registration/dispatch transport and responsibility boundaries require approved design and validation before implementation planning. Source: sections 4, 6, 10a.\n\n### REQ-006 — Menu navigation\nRetain arbitrary-depth submenus, breadcrumbs, Esc/Backspace navigation, conditional q handling and Tab pagination. Verify state transitions and key behavior, including hidden active entries. Source: section 7.\n\n### REQ-007 — Destructive-action confirmation\nFor confirm-marked actions, y confirms; cancellation does not dispatch. Verify named-action and callback confirmation/cancellation paths, including destructive Rex actions. Source: sections 6–7.\n\n### REQ-008 — Geometry and rendering\nSupport top, bottom, left, right and center; adaptive columns, Unicode display widths, constrained geometry, paging, icons and help/footer. Validate native-layer clipping in the targeted Rex version rather than copying the old overlay workaround. Verify narrow/small terminal cases and all five positions with real-client visual/e2e checks. Source: sections 7, 10c.\n\n### REQ-009 — Theme and terminal safety\nUse available terminal foreground/background/palette, fallbacks and contrast handling; respect NO_COLOR and ASCII mode. Restore terminal state on normal dismissal and errors. Verify palette/fallback/no-color/ASCII fixtures plus terminal lifecycle behavior. Source: sections 7, 9.\n\n### REQ-010 — Built-in actions and session parity\nRetain sessions (new/switch/previous/next/rename/close), tabs (new/previous/next/rename/goto/close), split, directional focus, zoom, sidebar, command palette and pane close. Preserve old dynamic/session behavior by observing the old plugin, not assuming an enumerated-session UI. Verify each action against the actual API and real Rex integration; resolve exact dynamic-session acceptance before dependent task planning. Source: sections 3j, 7, 10e.\n\n### REQ-011 — Origin context and focus\nTarget the originating Rex block/client/session as appropriate, restore focus before dispatch, and clean up popup resources after errors. Verify moved blocks, focus restoration, split/window operations, dismissal and error cleanup in real Rex. Source: sections 5, 7, 9.\n\n### REQ-012 — Rex-known remote host listing\nDiscover and list only remote hosts already known to Rex; do not introduce a Tailscale inventory or independent registry. Validate the actual CLI/server discovery mechanism before planning its implementation; verify controlled outputs and actual Rex discovery. Listing does not authorize remote attach/bootstrap/WOL flows. Source: sections 2, 8, 10d.\n\n### REQ-013 — Remote-server compatibility\nOperate with Rex on another host without client-local absolute-path assumptions; preserve originating IDs and server/session selection. Verify remote-context targeting and configuration/launch behavior with real remote Rex where available. Unavailable infrastructure is an explicit unmet gate, not a mock-only pass. Source: sections 7–9.\n\n### REQ-014 — Packaging and user configuration ownership\nProvide Nix flake and Home Manager integration without owning or overwriting the user's Rex init.lua. Verify package build and Home Manager configuration alongside an existing user setup. Source: section 7.\n\n### REQ-015 — Integration and release verification\nUse controlled CLI fixtures and automated mapping/navigation/rendering tests, complemented by actual Rex end-to-end tests for popup creation/dismissal, focus, split/window operations and cleanup. Test remote context and packaging; do not substitute coverage percentages or mocks for real-server gates. Each implementation task includes its own tests; final gates run on the integrated target. Source: section 9.\n\n## Deferred scope\n\nRemote-host connection/attachment, reuse-or-create remote sessions, reachability protocols, relay-assisted Wake-on-LAN, SSH bootstrap/fallback and retries are deferred. They require separately approved requirements and design, not placeholder implementation phases.\n\n## Mandatory design validation before implementation planning\n\nAn RFC is required because Lua-to-Go launch/config/callback boundaries and remote targeting are material public-interface/security concerns. Resolve source 10a–10g: callback transport/lifecycle; separate-file loading and precedence; native-layer geometry; known-host discovery; observed dynamic-session parity; concrete-package boundaries; and TUI framework versus direct I/O. Record the tested Rex version, CLI/API contracts and evidence, and obtain fingerprint-bound design approval. Do not treat source illustrations as contracts.\n\n## Original source (verbatim)\n\n# RFC: Rex WhichKey — Go Rewrite\n\n**Status:** Draft — records accepted decisions and identifies unresolved details  \n**Date:** 2026-10-27  \n**Project:** rex-whichkey\n\n## 1. Summary\n\nRex WhichKey will be rewritten from scratch in Go. The existing [`rex-whichkey.old`](https://github.com/curtbushko/rex-whichkey.old) repository defines the initial functional baseline, **not** the architecture or code to reuse. The implementation will be a short-lived Go process launched from a minimal Rex Lua integration, with a relatively flat project structure and no hexagonal architecture.\n\nThe project will offer a WhichKey-style contextual keyboard menu for Rex. It will use the Rex Server API through the `rex` CLI, preserve Lua-based keymap configuration and user-defined callbacks, and eventually support remote Rex sessions and Wake-on-LAN workflows.\n\n## 2. Goals and scope\n\n### Initial release goals\n\n- Match the existing behavior of `rex-whichkey.old` as closely as practicable.\n- Rewrite the implementation in Go without carrying over its architecture or technical debt.\n- Launch the application on demand as a short-lived process; exit when dismissed or after dispatch.\n- Provide a minimal Rex Lua launcher for registering the invocation and starting the Go process.\n- Support both inline Lua configuration and a separate Lua configuration file.\n- Preserve a WhichKey v3-inspired mapping specification, including named Rex actions and user-defined Lua callbacks.\n- Provide a typed, reusable Go Rex client centered on concrete `Client`, `Session`, `Window`, and `Block` types, corresponding to Rex's resource model.\n- Call the Rex CLI for Server API operations, using the published [Server API](https://www.superlogical.com/rex/docs/automate/server-api) and [CLI](https://www.superlogical.com/rex/docs/automate/cli) contracts.\n- Preserve existing session-navigation features.\n- List remote hosts already known to Rex rather than maintaining a separate host registry.\n\n### Future direction\n\nSupport connecting to remote Rex sessions, reusing an existing session or creating one when absent. Later integration may add machine reachability checks, relay-assisted Wake-on-LAN, and SSH-based bootstrap or fallback as appropriate. This is a planned direction, not an approved detailed protocol or workflow.\n\n### Non-goals\n\n- Reusing the old Lua/Bash architecture, transport, rendering pipeline, or configuration snapshot protocol as requirements.\n- Running a daemon or persistent coordinator for the initial application.\n- Imposing hexagonal architecture, services, drivers/adapters directories, or an interface hierarchy.\n- Maintaining a separate remote-host inventory when Rex already knows the hosts.\n- Reimplementing the Rex server or replacing its existing API.\n\n## 3. Accepted design decisions\n\nThe entries below capture explicit choices made during discussion. No retrospective historical timestamp IDs are assigned; discussion references are used only to organize this RFC. A formal decision register can assign new immutable IDs at the time it is actually established.\n\n| Ref | Status | Decision |\n| --- | --- | --- |\n| 3a | ✓ Agreed | Rewrite Rex WhichKey from scratch in Go; use `rex-whichkey.old` as a functional reference, not an architectural basis. |\n| 3b | ✓ Agreed | The Go program is short-lived, invoked when the popup is needed; no initial persistent service. |\n| 3c | ✓ Agreed | A minimal Rex Lua integration is required to launch the Go executable. |\n| 3d | ✓ Agreed | Keep the Go project simple and relatively flat; do not use hexagonal architecture. |\n| 3e | ✓ Agreed | Use concrete Rex resource types — `Client`, `Session`, `Window`, `Block` — rather than a `*Service` hierarchy. |\n| 3f | ✓ Agreed | Model the Go Rex client from Rex's Server API and map operations cleanly to calls through the Rex CLI. |\n| 3g | ✓ Agreed | Use Lua for WhichKey configuration, with WhichKey-style mapping definitions. |\n| 3h | ✓ Agreed | Support configuration directly in Rex `init.lua` and via an imported separate Lua file. |\n| 3i | ✓ Agreed | Support both named Rex actions and user-defined Lua callbacks. |\n| 3j | ✓ Agreed | Preserve the existing session functionality and list remote hosts available in Rex. |\n\n## 4. Proposed architecture\n\nThis section describes implementation proposals consistent with the accepted constraints; these exact file boundaries and interfaces are **not yet accepted decisions**.\n\n```text\nRex keybinding\n    |\n    v\nMinimal Rex Lua launcher\n    |\n    v\nShort-lived rex-whichkey Go process\n    |-- configuration / mapping navigation\n    |-- terminal input and rendering\n    |-- action selection and confirmation\n    `-- rex package (typed Rex API client via CLI)\n    |\n    v\nRex CLI / Server API\n```\n\nThe launcher should register the Rex action and create/launch a popup terminal. It should not own the menu navigation or rendering. Executable Lua callbacks cannot be serialized as plain JSON; a small Rex-side action registration/dispatch mechanism is consequently expected, but its protocol and exact responsibility split require design and testing.\n\n### Candidate layout\n\n```text\nrex-whichkey/\n├── cmd/rex-whichkey/main.go\n├── rex/\n│   ├── rex.go\n│   ├── session.go\n│   ├── window.go\n│   ├── block.go\n│   ├── layer.go\n│   └── layout.go\n├── mapping.go\n├── menu.go\n├── action.go\n├── render.go\n├── theme.go\n├── config.go\n├── plugin.lua\n├── nix/\n└── go.mod\n```\n\nThis tree is illustrative. Go package placement must remain valid (e.g., a `main` package must reside in an executable directory), and the final layout should minimize unnecessary directories and exported APIs.\n\n## 5. Go Rex client design\n\n### Principles\n\n- Prefer ordinary concrete objects and methods (`Client`, `Session`, `Window`, `Block`) over `Client → SessionService → ...` chains.\n- Model identities, API parameters, and response data using types grounded in Rex's published schemas.\n- Use the `rex` CLI to call methods; avoid embedding shell scripts or interpolating untrusted input into shell commands.\n- Keep CLI argument construction, JSON decoding, error handling, and server/session targeting in the `rex` package.\n- Keep a low-level generic API-call escape hatch for methods not yet wrapped.\n- Avoid introducing transport/service interfaces or code generation until a demonstrated need arises.\n- Use small interfaces at call sites only where genuinely helpful for tests or composition.\n- Keep session identity independent of temporary window membership, especially for blocks that may move.\n\nIllustrative intended ergonomics (not finalized method signatures):\n\n```go\nclient := rex.New()\nsessions, err := client.Sessions(ctx)\nif err != nil { return err }\nwindows, err := sessions[0].Windows(ctx)\nif err != nil { return err }\nblocks, err := windows[0].Blocks(ctx)\nif err != nil { return err }\nreturn blocks[0].Focus(ctx)\n```\n\nThe package should be checked against the actual server method descriptions and CLI argument syntax. Do not infer request fields from convenient Go method names.\n\n## 6. Configuration\n\n### Agreed configuration model\n\n- Lua is the user-facing configuration language.\n- Both inline specification in `init.lua` and an imported configuration file are supported.\n- Use a WhichKey v3-inspired list of mapping entries with full key sequences, groups, descriptions, icons, action names, callback functions, and action arguments.\n- Preserve callback execution in the Rex Lua runtime. The Go process should never pretend a Lua function can be transmitted to Go as ordinary JSON and invoked there.\n\nIllustrative syntax:\n\n```lua\nwhichkey.setup({\n  position = \"center\",\n  spec = {\n    { \"w\", group = \"Sessions\" },\n    { \"wn\", \"session.new\", desc = \"New session\" },\n    { \"z\", \"pane.zoom\", desc = \"Zoom pane\" },\n    { \"x\", \"pane.close\", desc = \"Close pane\", confirm = true },\n    {\n      \"h\",\n      function(ctx, args)\n        -- Custom Rex Lua action\n      end,\n      desc = \"Custom action\",\n    },\n  },\n})\n```\n\nThe setup/module name and loader paths here are illustrative, not yet an API promise. To prevent import name collisions, the shipped integration module and user configuration file should have unambiguous names.\n\n### Original mapping behavior to retain\n\nThe existing plugin's map compiler supports:\n\n- Full key sequences and implicit prefix groups.\n- Nested attribute scopes with inherited properties.\n- Named actions and Lua callbacks.\n- Conditional entries (`cond` boolean or function).\n- Dynamic string-valued descriptions, group labels, and icons (functions).\n- Hidden entries that still respond to keys.\n- Named action arguments (`args`).\n- Confirmation for destructive operations (`confirm`).\n- Validation of duplicate action mappings, group/action conflicts, and unsupported key notation.\n- `<Space>` and `<lt>` representations.\n\nThese are functional acceptance targets. Their new internal representation is unconstrained by the old implementation.\n\n## 7. Initial feature-parity inventory\n\nThe inventory below is based on the public old repository's code and documentation, not on claims that the new implementation already works.\n\n| Area | Initial functional requirement |\n| --- | --- |\n| Menu navigation | Arbitrary-depth submenus, breadcrumbs, Esc/Backspace navigation, conditional `q`, Tab pagination. |\n| Destructive actions | Confirm with `y`; cancel without dispatch. |\n| Positions | Top, bottom, left, right, and center. |\n| Rendering | Terminal-oriented adaptive columns, Unicode display widths, constrained geometry, paging, icons, help/footer. |\n| Theme | Respect terminal foreground/background/palette where available, fallbacks, contrast handling, `NO_COLOR`, ASCII mode. |\n| Rex actions | Sessions (new/switch/previous/next/rename/close), tabs (new/previous/next/rename/goto/close), split, directional focus, zoom, sidebar, command palette, pane close. |\n| Context safety | Act on the originating Rex block/client/session as appropriate; restore focus before dispatch. |\n| Packaging | Nix flake and Home Manager integration without taking ownership of the user's Rex `init.lua`. |\n| Remote compatibility | Correct behavior when Rex is running on another host; no client-local absolute path assumptions. |\n\nThe old implementation documents a floating-layer clipping problem on Rex Beta build 1101 and a full-window modal overlay workaround. This is **evidence to test**, not a decision to transplant that implementation. Native layer geometry and actual behavior must be validated against the Rex version targeted by the rewrite.\n\n### Dynamic session behavior\n\nThe original configuration offers `session.switch` through Rex, not necessarily a dynamically expanded enumerated list of session names in the popup. The user requires preservation of current dynamic/session behavior; the exact UI parity should be tested against the old plugin rather than assumed from an action name.\n\n## 8. Remote hosts and sessions\n\n### Agreed\n\n- Discover/list only the remote hosts Rex already knows about for the initial remote-host menu.\n- Do not introduce an independent Tailscale host inventory for that purpose.\n\n### Future direction (not yet detailed or accepted)\n\nWhen a user selects a remote host, an eventual flow could:\n\n1. Inspect server availability.\n2. Show existing Rex sessions.\n3. Attach to a selected session or create a session if appropriate.\n4. If unreachable, optionally use a configured relay for Wake-on-LAN, following capabilities from the user's `tailssh` script.\n5. Retry discovery/attachment and provide useful failure feedback.\n\nThe implementation must distinguish a reachable host from a reachable Rex service and from a specific existing session. The CLI's remote-server selection should remain the Rex client's primary path where possible; SSH/WOL details are unresolved and not required to recreate the old menu.\n\n## 9. Testing and quality gates\n\nPrioritize real integration and end-to-end behavior, not test coverage percentages.\n\n- Test typed Rex calls against known API schemas and controlled CLI output.\n- Use argument-array CLI invocation; verify escaping, JSON decoding, server/session selection, timeouts, cancellation, and errors.\n- Exercise mapping parsing, inherited attributes, hidden keys, conditions, and callback references.\n- Verify navigation, confirmations, pagination, Unicode width, themes, and terminal-state restoration.\n- Run actual Rex integration tests for popup creation/dismissal, focus restoration, split/window operations, and cleanup after errors.\n- Test remote-server context carefully, including the originating block/client IDs.\n- Ensure Nix packaging and Home Manager configuration do not overwrite the user's existing Rex setup.\n\nA mock CLI is useful for unit tests but is not a substitute for running the real Rex server where available.\n\n## 10. Open proposals and validations\n\nThese items must not be interpreted as approved decisions.\n\n| Ref | Status | Item | Resolution path |\n| --- | --- | --- | --- |\n| 10a | ◇ Proposal | Keep the minimal Lua callback dispatcher separate from all Go menu/navigation logic. | Specify the minimal launch/dispatch contract and test against Rex. |\n| 10b | ? Validate | Exact API for loading a separate user mapping file while also permitting an inline `setup({ spec = ... })`. | Test Rex's `require`/`dofile` lookup and document precedence. |\n| 10c | ? Validate | Whether native floating layers are now correctly clipped in target Rex versions. | Real-client visual/e2e test in all five positions. |\n| 10d | ? Validate | Exact list/discovery mechanism for Rex-configured remote hosts in the CLI and Go API. | Inspect current `rex` CLI/server API output. |\n| 10e | ? Validate | What “dynamic sessions” means beyond old `session.switch` behavior. | Compare interactive behavior of the Lua plugin with acceptance tests. |\n| 10f | ◇ Proposal | Keep the Go Rex package in one package with concrete resource methods and no service hierarchy. | Review as implementation evolves; avoid unnecessary layers. |\n| 10g | ? Validate | Whether the Go TUI should use a third-party framework or direct terminal I/O. | Build a minimal popup spike measuring startup and interaction. |\n\n## 11. References\n\n- [rex-whichkey.old](https://github.com/curtbushko/rex-whichkey.old) — current functional reference.\n- [which-key.nvim](https://github.com/folke/which-key.nvim) — mapping syntax and interaction inspiration.\n- [rexdiff/rex](https://github.com/rednafi/rexdiff/tree/master/rex) — Go examples of Rex CLI/API calls.\n- [go-libghostty](https://github.com/mitchellh/go-libghostty) — preferred Go API/style reference.\n- [Rex Server API](https://www.superlogical.com/rex/docs/automate/server-api) — authoritative API reference.\n- [Rex CLI](https://www.superlogical.com/rex/docs/automate/cli) — CLI invocation reference.\n- [tailssh](https://github.com/curtbushko/nixos-config/blob/main/modules/home/scripts/tailssh) — inspiration for future remote bootstrap/Wake-on-LAN.\n\n---\n\n**Decision authority:** This RFC documents existing explicit agreements. Its proposals and validation questions do not become accepted merely by inclusion in the RFC.\n","provenance":"Exact description before user-directed implementation-first replanning REXWHICHKEY-20261031-072107; previous evidence remains untouched.","gateEligible":false}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDExOjI5OjA4LjIyOVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoicmVjb3ZlcnkiLCJyZXRyeUtleSI6InJleC13aGljaGtleS93aGljaGtleS1maXJzdC9wcmQtcmVjb3ZlcnktcmVhc3Nlc3NtZW50LzEiLCJpbnRlbnQiOiI4OTAyMzE2MWU5MDBkZjAyNmQ4NGE3YWNmODNkN2VmZDAxNDI2Nzc3NmI2YTBkZjA0ZDliNDE5NTdhNThhYjRmIn0 -->
### 2026-10-31T11:29:08.229Z · recovery · unspecified

{"approvedBy":"user","attempt":2,"collectionRoot":"/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main/.tasks","decision":"approved","decisionEvidence":"User explicitly instructed stop further tests/experiments, prioritize which-key behavior and update the plans. Reassess the existing exact malformed-approval quarantine after that authorized PRD amendment; no new target or recovery disposition.","disposition":"quarantine-malformed-approval","eventId":"rex-whichkey/whichkey-first/prd-recovery-reassessment/1","kind":"recovery","recordRef":"prds/rex-whichkey.md","relatedDecisionId":"REXWHICHKEY-20261030-121414","replacementEventId":"rex-whichkey/whichkey-first/prd-amendment/1","schemaVersion":1,"scope":"Preserve the identical original malformed event, intent and digest; advance only its current-document binding after explicit user-directed plan correction. Approval is independently recorded for the current amendment.","specHash":"sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1","supersedesRecoveryEventId":"rex-whichkey/prd-approval-recovery/1","target":{"contentDigest":"sha256:040e15c15bac643ad9cb5618863f0d747e2baebd583aee61422b047164e70227","eventId":"rex-whichkey/prd-approval/1","intent":"bff6a866cfd8e1323a02ce0ba0637013cd8252085546167c2162099143792b32","specHash":"sha256:2b57ebed7620171e9b1528c160dc374314926b89e021f1f920c77783638eccd8"}}
<!-- mtn:end-entry:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTEwLTMxVDExOjI5OjEwLjI1N1oiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS93aGljaGtleS1maXJzdC9wcmQtYW1lbmRtZW50LzEiLCJpbnRlbnQiOiI0NjFjZDRhZTI0YTA2ZGYzYjEzMjdiZjc3ODFhZGQxZThiNGZjMGViZTI3YmI1MWNkNWM0YjI4NDUxMWY5ZWU2In0 -->
### 2026-10-31T11:29:10.257Z · approval · unspecified

{"approvedBy":"user","attempt":1,"decision":"approved","decisionEvidence":"User explicitly instructed: Do not do any more tests and experiments. The core of this program is the whichkey behaviour. rex is just an api and easy to understand. Follow-up: upate the plans.","decisionId":"REXWHICHKEY-20261031-072107","eventId":"rex-whichkey/whichkey-first/prd-amendment/1","kind":"approval","recordRef":"prds/rex-whichkey.md","requirements":["REQ-001","REQ-002","REQ-003","REQ-004","REQ-005","REQ-006","REQ-007","REQ-008","REQ-009","REQ-010","REQ-011","REQ-012","REQ-013","REQ-014","REQ-015"],"schemaVersion":1,"scope":"Record the user-directed correction to planning/research/verification policy while preserving functional requirements. No candidate public interface, test pass, implementation or release completion is approved.","specHash":"sha256:7ea834417c1b09142a3f8f9877ae6fc34b4500d7a8841d7460edb74ed97553d1"}
<!-- mtn:end-entry:v1 -->
