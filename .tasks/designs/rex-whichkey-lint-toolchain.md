---
identity:
  kind: design
  designID: rex-whichkey-lint-toolchain
  uuid: 6e396526-906a-4dc2-8c81-c18ffd30bbd1
lifecycle:
  state: approved
  sourceEventID: rex-whichkey-p21/design/authority/1
  changedAt: '2026-11-24T13:50:12.565Z'
title: Reproducible lint-toolchain compatibility
mtnCreateRetry:
  key: rex-whichkey-p21-design-v1
  intent: 706ab6562636c37ff56a4fe9a7dbe4a87bd50aed6afcb4042137fc1ec33cda7a
status: open
priority: normal
recurrenceAnchor: scheduled
occurrenceMaterialization: manual
occurrenceNextTrigger: completion
dateCreated: '2026-11-24T13:50:07.469Z'
dateModified: '2026-11-24T13:50:13.186Z'
tags:
  - task
type: task
id: 9bc01e37-2ad7-4760-80c0-5fc717ed3573
---
# Reproducible lint toolchain compatibility

## Authority and problem
User reported golangci-lint 2.8.0 built with Go1.25.5 rejecting this repository's Go1.26.8 target, then explicitly approved: "yes, fix flake.nix". This authorizes the bounded maintenance repair and its tests/reviews/integration, not broad dependency upgrades, profile activation or weakening quality rules. Earlier lint passes used compatible scoped tools; they did not verify the default development environment. Phase20 behavior stays settled/completed.

## Accepted outcomes
- REQ-ENV-001: The declared default dev shell provides golangci-lint and go-ai-lint compatible with the application's still-pinned Go1.26.8. Full existing task lint succeeds using those declared tools, not private .mtn-tmp binaries, a hardcoded store path or a global/profile workaround.
- REQ-ENV-002: Tool sources/build inputs are reproducibly pinned. Preserve application Go version/checksum, Go dependencies, lint configuration/rules, Taskfile/hooks, package/runtime behavior and existing shellHook behavior. Prefer a focused tool override; no broad nixpkgs/input refresh. A narrowly necessary lint-only source/input/lock pin is within this authorized repair. Never run the shellHook's automatic git pull/config-copy operations during verification; no user profile/config edits, installation, activation, push or GC.
- REQ-ENV-003: Add a small committed lint-toolchain regression check using the same Go/lint derivations selected by the dev shell. It must exercise a Go1.26.8-targeted module and both lint tools, detecting the original incompatibility rather than only inspecting version strings. Evaluate existing four-system outputs, build the host check, and run task build, task test (unit+race) and the FULL unchanged task lint pipeline under the actual declared dev-shell environment without executing shellHook.

## Bounded implementation plan
One maintainer-owned vertical task owns package alignment plus its regression and verification. Inspect locked package definitions, compiler provenance and tool minimum Go versions first. golangci-lint2.15.0 is the already-verified compatible candidate, not a separately user-mandated version. Ensure go-ai-lint's real builder/toolchain is compatible as well; an overlay's displayed Go version alone is not proof. Do not merely bypass the version guard. Keep sources/vendor hashes fixed and reuse existing caches; no nested persistent Nix profiles.

Preserve all existing shellHook commands unexecuted. Inspect a generated development environment (for example supported print-dev-env JSON) and apply only its declared environment to gates, never evaluating shellHook. Copy unchanged shared lint data files into isolated worktrees only as needed; no overwrites of the user's main config files. A committed fixture may contain only incidental minimal Go source; meaningful new Bash check logic receives its own code/test reviewers alongside separate Nix code/test and maintenance review.

## Verification and limitations
Require genuine failure evidence for the old Go1.25-built linter and green actual tool execution against Go1.26.8. The regression check and real repository gates must consume the default selected derivations, not an unused alternate package set. Host builds are actual runtime evidence; other-system evaluation is not runtime certification. Keep logs/private env transports excluded, use one heavy gate at a time within CPU14/memory3GiB, Nix max-jobs1/cores2, and no automatic global cleanup.

<!-- mtn:evidence:v1 -->

<!-- mtn:entry:v1:eyJ0aW1lc3RhbXAiOiIyMDI2LTExLTI0VDEzOjUwOjExLjg2NVoiLCJhdXRob3IiOiJ1bnNwZWNpZmllZCIsInJvbGUiOiJjb29yZGluYXRvciIsImNhdGVnb3J5IjoiYXBwcm92YWwiLCJyZXRyeUtleSI6InJleC13aGljaGtleS1wMjEvZGVzaWduL2F1dGhvcml0eS8xIiwiaW50ZW50IjoiYTYxYTYxYjQ0NGU4NDBmYWNmODc0MGFjZTNjYzhlODMyNGRkZmUzMjc4MGM0MzdlZTc0N2Q5NWE1ODNkOWY1OCJ9 -->
### 2026-11-24T13:50:11.865Z · approval · unspecified

{"approvedBy":"User explicit flake repair authorization; coordinator bounded maintenance plan","attempt":1,"decision":"approved","decisionEvidence":"User reported Go1.25-built golangci-lint rejecting Go1.26.8, then explicitly said yes, fix flake.nix after recommendation to align compatible linters reproducibly. No separate approval of low-level implementation details is claimed.","eventID":"rex-whichkey-p21/design/authority/1","kind":"approval","recordRef":"designs/rex-whichkey-lint-toolchain.md","requirements":["REQ-ENV-001","REQ-ENV-002","REQ-ENV-003"],"schemaVersion":1,"specHash":"sha256:5d3ea3bebc1b203988b65e69ac1c0838025d7287509c668817fd4aa9e4a8e606"}
<!-- mtn:end-entry:v1 -->
