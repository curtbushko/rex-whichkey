# Unnecessary-parts audit: Phase17

## 1. Scope, authority and inspected identities

This is an audit/report, not a deletion list or a performance study. Only this
file changes. No removal, assertion adaptation, dependency update, status fix,
configuration migration or runtime redesign is approved or implemented here.
The author is builder `rex-whichkey-p17-t01-builder-a01`, not an independent
reviewer. Fresh frozen-revision review and coordinator integration remain pending.

The governing task is `tasks/rex-whichkey-17-01-unnecessary-parts-audit.md`,
specification `sha256:efc85700f7bc5613ecb7c423eb19628f095c5669dec4a5e30b203c90829d74fb`.
The accepted CLI-only design is parent collection
`.tasks/designs/rex-whichkey-cli-only-menu.md`, semantic hash
`sha256:23add062dac4326fe35e7e178a6a39f4c4d7631b4c8e164feb988b9aba07e21b`.
Its accepted direction `REXWHICHKEY-20261106-114935` governs: flat, short-lived
Go, data configuration, concrete argument-array Rex CLI operations, and
launch-only Lua. Hooks/callbacks and executable REQ-004 portions remain deferred;
REQ-003's Lua loading mechanism is superseded, and REQ-005 is not claimed.
Later Phase18 startup decisions are separate work, not an audit implementation.
“Required” below describes current consumers, not a veto on that separately
accepted future scope.

| Identity | Exact value / meaning |
| --- | --- |
| Required inventory baseline | Commit `9634246b7e4edbc8bed59210466cac492ddf2f0f`; tree `452d7579e12ccf7fb55dac1ba4454f7d88f58019` |
| Actual inspected checkout | `/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/p17-t01-builder-a01`; branch `work/p17-t01-builder-a01` |
| Actual initial HEAD / builder base | `cef25463b4dd58c99228d8f9b857ed3554a3164b`; tree `8ebed6c0d2c567024e2a5659c8f769802a6022ee` |
| Baseline-to-base differences | Only `.tasks/phases/rex-whichkey-17-unnecessary-parts-audit.md`, `.tasks/phases/rex-whichkey-18-startup-speed.md`, `.tasks/tasks/rex-whichkey-17-01-unnecessary-parts-audit.md`; `git diff --quiet BASELINE HEAD -- ':!.tasks'` succeeded |
| Parent inspected read-only | `/Users/curtbushko/workspace/github.com/curtbushko/rex-whichkey/main`; HEAD was the same `cef25463b4dd58c99228d8f9b857ed3554a3164b` |
| Parent local differences | Modified Phase17 phase/task records and **uncommitted `flake.nix`**, described separately below; none copied, edited or staged |
| Author's delivery identity | Exact final report commit/tree and changed-file list are in the build protocol result; the report cannot embed its own containing commit hash |

Unless another revision is stated, a path, line range, declaration, function or
heading below means that source at the baseline above, byte-identical to the
inspected base outside `.tasks`. Symbol/heading references are revision-stable
locators, not claims about current HEAD after future work. Parent `.mtn-tmp`
references are local recorded artifacts, not tracked or installed assets.

### Method and bounded exclusions

Inspected tracked-file inventory, declarations/checksums, source callers,
serialization and resource boundaries, platform tags, shell/Task/hook text,
Nix source/install filters, production Lua, examples, documentation, retained
Go tests, Lua loader paths, fixture imports, experiment modules/provenance and
specific parent Phase16 evidence. Source/Git/local-file reads and whitespace
checks only: no Task, Go, Nix, shellHook, hook, product, test or harness execution.
Reading commands in a file is not running them. No environment entry, reload,
activation, deployment, remote/API/GUI interaction or Rex resource change occurred.

Searches establish bounded repository observations, never external non-use.
For candidate code, production paths were followed from `cmd/main.go`, not just
searched by name. JSON tags and marshal/unmarshal methods matter even without
an explicit call spelling; `rex.Client.Call` uses reflection on result pointers.
No `go:embed` or `go:generate` directives were found in active root/cmd/rex Go
sources; assets have explicit file/install consumers instead. Platform tags and
nested module boundaries are considered below. This is not a compiler-derived
whole-program reachability proof, module-graph recomputation or binary analysis.

Excluded from exhaustive review: every vendored experiment implementation/test,
all historical task events, external repositories/websites, personal editor
settings, user init/config/profiles, private module-cache internals and live
platform/package behavior. Inspected dependency cache files are specified below;
the cache was neither modified nor executed. Prior evidence artifacts were read,
not regenerated. No missing external fact is treated as deletion permission.

## 2. Classification vocabulary and consumer rules

Each matrix row has exactly one of these six classifications:

| Classification | Meaning in this report |
| --- | --- |
| **required production** | A traced supported runtime, configuration, library or installation contract consumes it; retain |
| **required build/quality tooling** | A declared build, reproducibility, gate or update workflow consumes it; retain standards and inputs |
| **optional developer convenience** | A real declared developer entrypoint exists but the supported application/gates do not require that entrypoint; external/personal use may still exist |
| **historical/test-only** | Retained assertions, fixture loaders, diagnostic modules or historical evidence consume it; this is not safe-to-delete evidence |
| **confirmed unreferenced/redundant** | A narrowly demonstrated redundancy, not an inference from zero name hits; only the repeated ignore rule below qualifies |
| **unresolved** | Missing external/public/editor/graph/behavior evidence prevents a disposition; retain pending that evidence |

For the matrices, **P** = production, **B** = build/quality, **T** = tests/fixtures,
**D** = docs/examples/package consumers, **E** = editor/external use. “No traced
P caller” is scoped to this checkout and entrypoint analysis, not global non-use.
For unexported historical helpers E is not a supported public API; for exported
Go declarations E remains unknown. No row permits weakening assertions or lint.

## 3. Tooling, dependency and package disposition matrix

### 3.1 Shell tools and gate wiring

| Candidate and exact evidence | Classification | Consumers, disposition, risk and smallest remaining action |
| --- | --- | --- |
| `flake.nix:87` `docker` devShell package | unresolved | B: shell provides it, but root `Taskfile.yml`, package recipe and hooks do not invoke it; no Dockerfile/container entrypoint appears in the tracked inventory. P/T/D: no traced dependency in active paths. E: personal/editor/container workflows unknown. Retain. Before proposing removal, ask the user about external workflows and review any supplied callers; zero hits is insufficient. No runtime/build benefit is claimed. |
| `flake.nix:81-83,96` `make-wrapper`, a `make` script doing `exec ${pkgs.go-task}/bin/task "$@"` | optional developer convenience | B: real argv-preserving alias to Task, not GNU Make and not a Makefile parser. Gates call Task directly. D: historical `docs/actions.md` evidence index includes normal `make run`; that observation depended on this compatibility entrypoint. E: external make habits/editor commands unknown. Retain; any later removal must preserve historical commands and document direct Task substitution, with explicit user acceptance of the entrypoint change. |
| `flake.nix:93` Python `pre-commit`; `.pre-commit-config.yaml:1-23` local system-language hooks | optional developer convenience | B: optional framework invokes `task build`, `task lint`, `task test` with no filename filtering. Native hook does not invoke Python pre-commit. E: manual framework runs/installations unknown. Retain while choosing whether both entrypoints are wanted; removal would require editing package/config/documentation together, never dropping the three standards. No Python application source is implied by the package. |
| `.githooks/pre-commit:1-11` direct Task hook | required build/quality tooling | B: `set -euo pipefail`, Task availability check, then build/lint/test. `flake.nix:107-109` selects it via local Git configuration. T: test means unit **and race**. E: actual installed hook settings were not changed. Retain intact. This audit bypasses it only per commit invocation because the user prohibited those suites; bypass is not a gate pass or policy removal. |
| `flake.nix:91` `gotools` | unresolved | B: explicitly supplied tool bundle, but gates spell `go vet`, `golangci-lint`, `go-ai-lint`, not a gotools executable. Generated golangci config enables goimports formatting; that does not prove it shells out to this bundle. E: editor use of tools such as goimports is unknown, with no tracked `.vscode`/`.idea` config. Retain; identify exact binaries/editor consumers before a separately approved trim. |
| `flake.nix` `goVersion`, overlay, source hash; `go.mod:3` `go 1.26.8` | required build/quality tooling | B: devShell uses overlay Go; flake package overrides buildGoModule with it; Task build/install/vet/test/tidy use Go. P: compilation of root `whichkey`, `temp/rex`, `cmd`. T: same module test compilation. D: package exports. Retain version/hash/locks; no toolchain update or downgrade proposed. |
| `flake.nix:90` `go-task`; `Taskfile.yml` tasks | required build/quality tooling | B: direct and Python hooks, make wrapper, build/run/install/lint/test/tidy callers. Root variables execute `go env` and Git commands when Task is invoked, so even `task --list` was not run here. `clean` deletes bin; tidy mutates modules. Retain; no audit invocation, including ostensibly harmless discovery. |
| `flake.nix:92`, `Taskfile.yml:54-57` `golangci-lint` | required build/quality tooling | B: `lint` → `lint:vet` → `lint:golangci` → `lint:ai`; golangci consumes explicit `.golangci.yml`. Parent generated config has version2, readonly modules, enabled error/context/unused/style rules and existing exclusions. Retain all standards, not just the tool name. Prior failure is not justification to remove a linter or suppress findings. |
| `flake.nix:31-43,94` `go-ai-lint` source/pin/vendorHash; `Taskfile.yml:59-77` | required build/quality tooling | B: Task's final lint stage requests twelve analyzers (concurrency, context, defer, error, goroutine, init, interface, naming, option, panic, slicemap, string). `.go-ai-lint.yml` is copied for tool configuration, but Task supplies no explicit config flag: actual upstream discovery/CLI-rule interaction was not executed or verified. Retain pin, configuration and analyzer flags. Phase16 full lint stopped before this stage; do not claim it passed. |
| `flake.nix:6,80,95,111-112`; `.gitignore:46-48`; `flake.lock` shared-configs node; `renovate.json` | required build/quality tooling | B: locked `golang-shared-configs` `09ee23161f5df536e77985493757f44896d39847` supplies all-configs; shell copies both ignored lint files. Parent copies were read: golangci version2 and AI config version1 with include/exclude/rule settings. Leaf has no copies. Renovate enables Nix and automerge scheduling for this input. E: actual bot service/use unknown. Retain input/lock/rules/update declaration; entry-time copying is separable from the rules' necessity. |
| `flake.nix:99-113` shellHook; `.envrc:1` `use flake` | optional developer convenience | B/E: entering the configured shell can pull main, mutate hooksPath and overwrite ignored lint configs; direnv may trigger entry. P: not product startup; package does not require shellHook. These are source-observed side effects, not executed observations. Retain for now; any later explicit-entry alternative must preserve lint availability/standards and clarify user-owned local files and hook choice. |
| `.gitignore:16,32` duplicate `*.test` pattern | confirmed unreferenced/redundant | B/E: Git ignore semantics repeat the identical rule with no intervening negation; either retained occurrence matches the same test binaries. P/T: no assertions or executable code affected by the repeated line itself. This confirms only textual/semantic redundancy of the second rule, not that test binaries or tests are unnecessary. A later separately approved housekeeping change could drop one copy/comment and review ignore behavior; untouched here. |

**Entry-time side effects are not quality requirements.** The shellHook tests
whether the branch is `main`, then runs `git pull --quiet || true`; failure is
ignored. If `.githooks` exists it executes
`git config core.hooksPath .githooks 2>/dev/null || true`. Both `cp -f` commands
are unconditional. Thus entering a development environment can update the
checkout, set repository hook behavior and overwrite local ignored files.
Nothing in this audit executes those commands or recommends deleting lint rules.
Moving or making these effects opt-in would be a separate tooling change needing
explicit scope, consumer checks and validation permission.

### 3.2 Dependencies, checksums and transitive consumers

The root module is **`temp`**, with packages `whichkey`, `temp/rex` and executable
`cmd`; no framework/service hierarchy or embedded Lua production engine is
introduced. Root `go.mod` and `go.sum` are distinct from experiment module files.

| Candidate / declaration | Classification | Actual consumers and bounded disposition |
| --- | --- | --- |
| `github.com/rivo/uniseg v0.4.7` (`go.mod`, `go.sum`) | required production | P: `render.go` width/clipping and `text_entry.go` grapheme cursor/editing/display call `uniseg` directly. D: appearance and rename contracts. B/T: root build/tests consume module. Retain; Unicode behavior cannot be traded away as “cleanup.” |
| `golang.org/x/sys v0.36.0`, `golang.org/x/term v0.35.0` | required production | P: `terminal.go` raw terminal/state, `terminal_poll_{darwin,linux}.go`, `popup.go`, `native_text_entry.go`, `runtime_bootstrap.go` Unix file/poll/pipe operations; T: terminal poll tests. Cached x/term v0.35.0 go.mod explicitly requires x/sys v0.36.0. Retain direct and transitive role; Darwin/Linux branches are not interchangeable dead files. |
| `github.com/stretchr/testify v1.11.1` | historical/test-only | T: `require` imports in launch-output, configured-rename/wire, Lua context, poll and endpoint tests. Root Task test unit/race paths compile these; Nix project checks are disabled, not a deletion signal. Cached `require/require.go` imports assert; assert uses spew/difflib/YAML. Retain unchanged assertions and dependency even though no tests ran. |
| `github.com/yuin/gopher-lua v1.1.2` | historical/test-only | T: `lua_context_test.go` creates/closes states, loads all three `tests/lua` scripts, makes protected calls and checks successes/errors/replay. P: launcher is ordinary Lua loaded by Rex, **not** this Go evaluator. Retain while these preserved tests exist. Removing it alone breaks their compile/import contract. |
| `go.mod` indirect spew v1.1.1, difflib v1.0.0, yaml.v3 v3.0.1 | historical/test-only | T: cached testify `assert/assertions.go` imports spew/difflib and assert/yaml; default YAML wrapper imports yaml.v3. Root's indirect declarations and sums thus have positive consumers, not merely vestigial names. Retain; no tidy/pruning run. |
| `go.sum` check.v1 `v0.0.0-20161208181325-20d25e280405` entries | historical/test-only | T/B: cached yaml.v3 v3.0.1 go.mod explicitly requires this test dependency; checksum history is not a direct product import. Retain checksum evidence. A later dependency-maintenance task, if authorized, must let the graph/toolchain determine required sums instead of hand-pruning. |
| Other module metadata edges (testify → objx; gopher-lua → readline/logex/test and older x/sys) | unresolved | Cached dependency go.mod files declare them; no direct root import was found and root sums do not include all of them. These are upstream module-level edges, not proof the selected root package set links them. Exact pruned package graph was not computed. Retain declarations/checksums as-is; don't add or remove pins based on this inventory. |
| Root `go.sum`, `flake.lock`, `nix/package.nix` vendorHash/source/artifact hashes | required build/quality tooling | B: reproducible module/vendor and Nix inputs; flake root uses `nixpkgs_2` `be5afa0fcb31f0a96bf9ecba05a516c66fcd8114`, shared-configs uses `nixpkgs` `aa290c9891fa4ebe88f8889e59633d20cc06a5f2`. Two nixpkgs nodes have distinct incoming edges, not demonstrated duplication. Retain every lock/hash; no evaluation/download/update performed. |

Cached metadata inspected read-only under
`/Users/curtbushko/go/pkg/mod/cache/download/`: exact
`github.com/stretchr/testify/@v/v1.11.1.mod`,
`github.com/yuin/gopher-lua/@v/v1.1.2.mod`,
`golang.org/x/term/@v/v0.35.0.mod`, and
`gopkg.in/yaml.v3/@v/v3.0.1.mod`. Assertion source imports were inspected under
`github.com/stretchr/testify@v1.11.1/{require,assert,mock}`. These local cache
observations support the stated edges; they are not a fresh authenticated module
verification, complete transitive closure, tidy result or test pass.

### 3.3 Packaging, installed assets and platforms

| Surface / exact source | Classification | Consumer trace, risk and retention / follow-up |
| --- | --- | --- |
| `flake.nix` `packages`, `overlays.default`, `homeManagerModules` | required production | D/B: packages exports both `rex-whichkey` and default; HM exports named/default module and sets host-platform package default. Four declared systems: x86_64/aarch64 Linux/Darwin. External flake consumers are unknown; preserve output names and Go selection. Declarations do not certify those platforms. |
| `nix/package.nix` `rexSources`, `rexPackageSource`, `dependency` | required production | D/P: pinned proprietary preview13 CLI for x86_64/aarch64 Linux and aarch64 Darwin; optional `rex-cli` override; **no Intel Darwin artifact**. Installs Rex binary, four terminfo paths and SSH config; Linux patchelf glibc/PAM rpath/interpreter and Darwin signature-preserving no-strip policy remain. These assets have external Rex/SSH/terminal consumers even without root Go imports. No license/dependency/native-asset removal proposed. |
| `nix/package.nix` `src = lib.cleanSourceWith`, `subPackages`, `postInstall` | required build/quality tooling | B/D: source filter admits root `.go` (including `_test.go`), cmd/rex trees, exact minimal Lua/examples, top-level docs Markdown, modules/LICENSE/README. It omits `.tasks`, workflow transport, tests tree and experiment descendants. `subPackages=["cmd"]`, checks/install checks false. Build source inclusion is not runtime reachability; exclusion from installed assets is not safe deletion. Preserve filters and disabled-check policy, not a test-pass claim. |
| Installed `bin/rex-whichkey`, optional `libexec/rex-whichkey/rex` | required production | P/D: normal image plus optional licensed CLI link consumed by Go discovery (`rex/runtime.go` `layoutDependency`). Recipe renames cmd; no generated wrapper or Lua-tree executable. Actual installed/profile behavior was not exercised in Phase17. Retain public paths; changes would affect execution-host users. |
| `lua/rex_whichkey.lua` `M.launch`; `nix/home-manager.nix` `home.file` | required production | P/D: package installs only this launch module; HM opt-in adds package and one non-forcing file link beside user init.lua, with configurable package/directory. It does not own init/data/binding/whole directory or reload Rex. Lua calls detached argv start, passes/clears exact context env, raises immediate start errors, returns true. Retain minimal integration; physical-key/module-loading/deployed PATH remains unverified. |
| `examples/{README.md,my_whichkey.json,rex-whichkey-keybinding.lua}` | optional developer convenience | D: explicitly installed under share/rex-whichkey/examples, linked by README/data docs. P: not auto-loaded defaults; user may explicitly select/adapt JSON or binding. Example JSON replaces entire spec and has opt-in providers; key is not deployed. E: external copies unknown. Retain migration/examples; changes need contract/doc consistency and no automatic user-data rewrite. |
| `README.md`, top-level `docs/*.md`, `LICENSE` | required production | D: installed docs and license have real distribution consumers. This new top-level audit Markdown will also be included by existing docs filter/copy if packaged later; no packaging edit/build is made. Preserve license/history and documentation availability; address status/link defects only in a later docs change. |
| `terminal.go`, `terminal_poll_darwin.go`, `terminal_poll_linux.go` build tags | required production | P/B: terminal is Darwin-or-Linux; `terminal.go:228,319` calls platform-specific `pollTerminal`. Darwin uses `unix.Select` for the controlling `/dev/tty` (`terminal_poll_darwin.go:10-14,38`), rejecting negative or `>= unix.FD_SETSIZE` descriptors with `unix.EBADF` (`:23-24`); Linux calls `unix.Poll` (`terminal_poll_linux.go:8`). Matching poll tests have tags too; `terminal_poll_darwin_test.go:16-20` preserves the descriptor-range assertions. Runtime files use Unix interfaces. Windows support cannot be inferred from portable-looking declarations or copied experiment Windows files. Retain both implementations and unchanged assertions: terminal I/O consumes platform-specific readiness, including Darwin's descriptor-range guard. No tests, cross-build or new runtime verification here. |

**Parent-only Darwin NuGet diff — unresolved, preserve separately.** Read-only
`git -C main diff -- flake.nix` showed a new local `preCommit` binding using
`pkgs.pre-commit.overrideAttrs` with
`pkgs.lib.optionalAttrs pkgs.stdenv.hostPlatform.isDarwin { dontConfigureNuget = true; }`,
and replacement of `pre-commit` by `preCommit` in devShell packages. Its comment
attributes the problem to a retained .NET test dependency/NuGet setup unrelated
to building Python pre-commit. That is the parent's stated rationale, **not** a
reproduced derivation diagnosis or verified fix. It is absent from baseline and
leaf; no locks or other package declarations are changed by the displayed diff.
It changes optional tooling construction, not product cleanup or permission to
remove pre-commit. Preserve uncommitted bytes; do not copy, stage, incorporate or
revert it. The parent owns any separate Darwin build validation/commit decision.

## 4. Production versus retained compatibility code

### 4.1 Positive production trace (retain)

The active path is `cmd/main.go` `run` → `runLaunch` → `LaunchPopup`
(`launch.go:43`) → `launchNativePopup` (`native.go:34`). Configuration is
`LoadMenuConfig` → strict `decodeMenuOptions`/`menuObject`/`checkMenuJSON`
(`data_config.go`), defaults and canonical snapshot (`defaults.go`, `providers.go`).
The parent validates exact origin, captures CLI/endpoint, expands providers,
retains targets and rename labels, then invokes `nativePopupSelection` →
`runPopupLifecycle`. No legacy Lua registration gates or redirects this route.

`runPopupLifecycle` (`popup.go:134`) creates the private transport/PTY and waits;
its result is **`readNativePopupResult`** (`popup.go:431`), which accepts keys,
confirmation and optional text and resolves them through `retainedNativeEvent`
against the retained mapping. Cleanup precedes capture/dispatch, and
`launchNativePopup` additionally retires its source-image handle before effects.
The parent calls `dispatchNativeSelection`; inline rename/close/generated pane
focus use concrete resources; app actions use `dispatchNativeApp`, validated
built-in metadata and `rex.Client.Perform`. New Tab uses one `Session.NewWindow`
then activates its exact ID, preserving it if activation is uncertain.

The private child path is `cmd/main.go` `runPopupChild` →
`RunNativePopupTerminal` → `runPopupTerminal` (`popup.go:552-636`) →
`decodePopupConfig` (`providers.go:53`) → `readNativeTextSpec` → `RunTerminal` →
`commitPopupResult` → `writePopupMenuResult`. `menu` is a separate UI/result-only
entrypoint using `LoadMenuConfig`, `RunTerminal`, `WriteMenuResult`, not resource
dispatch or an executable-Lua loader. That makes `WriteMenuResult` live even
though its comment mentions a retained Lua invocation.

| Required item / exact symbols | Classification | Positive consumers / important non-removal boundary |
| --- | --- | --- |
| `mapping.go` `Config`, `ConfigOptions`, `Mapping`, validation/key helpers; `defaults.go`, `action.go`, `menu.go`, `layout.go`, `render.go`, `theme.go`, `text_entry.go` except historical predicate below | required production | P: loader/config snapshot, menu trie/navigation/confirmation, rendering, terminal/paste/rename. T: rename and other retained tests also call shared functions. D: public data/action/provider/appearance contracts. JSON tags/raw args and clone behavior are data consumers, not evidence of callback support. Retain shared validation, hidden/condition/confirmation and Unicode behavior. |
| `mapping.go:579` `completionObject`, `:617` `completionString`, `validateText`, key normalization; `runtime_bootstrap.go` strict JSON/Unicode helpers | required production | P: `data_config.go` menuObject/checkMenuJSON/action/session args, `native_resources.go`, `native_sessions.go`, `native.go` New Tab empty args, `native_text_entry.go` labels, `popup.go` live result decoder. `runtimeUnicode` calls completionString; readRuntime → decodePopupConfig/receipt parsing. Names are historical, consumers are live. Removing the whole Completion-adjacent region would break strict duplicate/unknown/null/UTF-8/surrogate boundaries. |
| `popup.go:42` `popupResourceID`; `rex/session.go:198` `resourceID` | required production | P: origin/selector/provider validation and typed Rex resource/client action paths respectively. Each package owns a different boundary; similar UUID checks are not confirmed redundant. External users of exported typed resources unknown. Retain shape checks and fresh membership/placement checks; names alone are not authority. |
| `popup.go` transport/file/FIFO/origin/cleanup helpers, `native_text_entry.go` `readNativeTextSpec`/`writePopupMenuResult`, `runtime_bootstrap.go` image/directory/receipt/result helpers | required production | P: traced parent/private-child flow above; private files, lease and retained authority remain necessary **as implemented**, not evidence that the old callback bridge survives. T: historical helpers share slot/directory validation too. Do not delete generic Go transport with legacy Completion readers. No redesign of copies/hashes/checks/synchronization/lifecycle is proposed. |
| `providers.go`, `hosts.go`, `native_sessions.go`, `native_app_owner.go`, `native_targets.go`, `native_resources.go`, `popup_initial.go` | required production | P: nativeSessionMenu reads sessions/windows/panes, resolveHostMenu → rex.Hosts only for enabled host markers; prepareNativeMenu records labels; exact targets and app generation/owner rechecked after cleanup. Popup geometry seeds one rectangle. D: opt-in providers/examples still consume nondefault behavior. Hosts are listing-only, not removable because absent from default root. |
| `rex/{rex,exec,runtime,endpoint,session,window,block,layer,geometry,terminal,wait,actions,hosts}.go` | required production | P: concrete client/resources and shared argv subprocess boundary; no second socket/RPC/Lua client. `Client.Call` validates reflected result pointers, serializes args and decodes even discarded JSON; JSON methods can be invoked indirectly. `Actions` queries built-in app metadata, `Perform` uses client.perform. `Block.WaitTerminal` is the live long wait; geometry/terminal seed are used before popup creation. Public library methods not called by this app are not proven removable; E unknown. Retain generic Call/resource APIs and existing safety behavior. |

### 4.2 Legacy candidates and their actual consumers

| Candidate / stable locator | Classification | Consumers, retention/removal rationale, risk, adaptations and smallest next validation |
| --- | --- | --- |
| `mapping.go:312-575` Completion constants/types/methods; `CompletionSelection`, versions1/2, complete/cancel, invocation lifetime30/max pending64 | historical/test-only | T: `configured_rename_test.go` constructs/round-trips Completion, validates retained selection/text/confirmation; `configured_rename_wire_test.go` invokes `UnmarshalJSON` indirectly through encoding/json and asserts rejection of null/type/duplicate/unknown/text/callback/surrogate/control/keys cases. `readPopupResult` supplies legacy messages. P: supported lifecycle does not construct Completion. Exported API/JSON external consumers unknown. Retain; any later retirement must preserve equivalent assertion intent, decide API compatibility and keep shared strict helpers in place. No automatic expiry/deferred-hooks requirement follows from these constants. |
| `mapping.go` `NewInvocationCapability`, `ValidateInvocationCapability`, `InvocationCapabilityMatches` | unresolved | T/internal: Completion.Validate validates shape; versions use same validation for receipt hashes. No active public-launch caller of constructor/matcher was found after following entrypoints, and no current local test directly calls those two exported functions. E: exported API consumers unknown. Serialization/internal relationships make the whole family nontrivial. Retain until explicit public API/compatibility decision and focused authorized validation; zero direct hits alone never authorize deletion. |
| `config.go:17-101` `DecodeConfig`, `checkConfigJSON` | unresolved | P: **not** public file loader or private snapshot decoder; those are LoadMenuConfig and decodePopupConfig. No caller of DecodeConfig appears in current root/cmd/rex/tests source, despite its comment claiming a transitional caller/compatibility tests and mapping docs' general test-retention description. checkConfigJSON is called by DecodeConfig and itself. E: DecodeConfig is exported; external library consumers unknown. Retain. Correct the inaccurate caller description in a separately scoped change, then decide compatibility/deprecation and needed legacy decoder coverage, without weakening live decoding. |
| `launch.go:103` `launchOutput`, `Write` | historical/test-only | T: `launch_output_test.go` constructs it, checks every split point in the fixed old client-mismatch diagnostic, no unrelated-match signal, and short-write budget. P: LaunchPopup uses captured image/CLI, not this writer; actual bounded metadata output uses `native.go` launchBuffer / rex boundedOutput. E: unexported. Retain assertions; retirement needs an explicit decision whether historical diagnostic compatibility is still a requirement and how to preserve evidence, not silent test deletion. |
| `popup.go:506` `readPopupResult`; `native_text_entry.go:68` `nativeTextSelection` | historical/test-only | T: configured_rename_test calls reader; reader validates private file/slot, strips authorized text through nativeTextSelection, wraps raw selection for json.Unmarshal(Completion), validates against retained config. It shares directory/slot primitives with production. P: runPopupLifecycle calls readNativePopupResult instead, never this reader. E: unexported. Retain; later tests would need to retain old wire assertions or deliberately separate historical compatibility from current wire tests under explicit approval. |
| `text_entry.go:22` `nativeRenameMapping` | historical/test-only | T/internal: Completion.ValidateSelection calls it, binding historical `tr/wr`, no-confirm/empty args rules; rename authority tests assert differing actions/callbacks/custom args are rejected. P: Menu.HandleKey, retainedNativeEvent, nativeResourceRename and dispatch use **inlineRenameMapping** instead, allowing supported action-based rename at configured keys. Retain legacy predicate while tests depend on it; never replace live predicate with hardcoded keys as cleanup. |
| `Mapping.CallbackRef`, `Selection` copy/accessors and `config.go` `WriteMenuResult` | required production | P: shared Mapping validation/clone/selection and live UI/result writer still carry/check the field; public decoder allowlist rejects callback_ref and active dispatch refuses callbacks. T: historical callback tests use the same data model and writer. E: exported shape unknown. Retain shared surface; unused executable callback support must not be inferred from a serialized field, nor all of config.go deleted with legacy DecodeConfig. Any future field/API pruning needs schema/test/public-contract approval. |
| `cmd/main_test.go` `greeting`, `TestGreeting` | historical/test-only | T: greeting exists only in the test file to preserve original Hello, World scaffold assertion; cmd production run/main is real menu code. Root test gates include it, Nix source filter admits it but checks are off. E: not public. Retain the assertion under current instructions; low product relevance is not authorization to discard it. |
| Other root/cmd/rex `_test.go`, including `terminal_poll*_test.go`, `rex/endpoint_test.go` | historical/test-only | T/B: Task test unit/race includes active module packages; poll tests are platform-tagged, endpoint tests retain canonicalization/discovery assertions. Some test behavior remains relevant production regression intent; this classification describes the files' consumers, not that all tests are obsolete. Retain all assertions/fixtures, no skips or expected-result changes; execution permission remains absent. |

### 4.3 Lua dependency and fixture chain

`lua_context_test.go` imports gopher-lua and testify, then:

- `TestConfiguredRenameDispatch` loads `tests/lua/configured_rename.lua`, checks
  tab/session/callback success and missing/blank/control/custom/injected/moved/
  uncertain failures, then rejects replay.
- `TestCLIDelegation` loads `tests/lua/cli_delegation.lua` for app/key/confirmation/
  callback/cancel/origin/capability/generation/expiry/reload/error scenarios, with
  protected calls into bad_completion/finish_test when present.
- `TestLaunchOriginPreflightDiagnostics` and `TestRexActionContexts` load
  `tests/lua/action_context.lua` for origin mismatches and server-context kinds.

All three scripts' **line1** prepend
`tests/fixtures/phase15-lua/?.lua` and `?/init.lua` to `package.path`.
`action_context.lua:42` and `cli_delegation.lua:68` require rex_whichkey;
configured_rename.lua:2-3 requires mapping/dispatch. The archived root module
requires runtime.anchor/config/mapping/dispatch; dispatch requires evaluate and
mapping; evaluate requires mapping. Archived defaults/init composes fragment
names dynamically, so literal-only searches miss imports. Some archive assets
are retained as a coherent historical snapshot even where no current script
visits them. That distinction does not justify breaking the snapshot.

**Classification: historical/test-only** for `lua_context_test.go`,
`tests/lua/{action_context,cli_delegation,configured_rename}.lua` and the complete
`tests/fixtures/phase15-lua` tree. The README states provenance
`cd5cbce818e07c264738ddcbc591ad49fdf0e9c8:lua/`. Read-only Git-blob comparisons
confirmed all **12 archived Lua files** are byte-identical to their corresponding
original paths, including config/evaluate/mapping/dispatch, five defaults plus
init, root module and runtime/anchor. No fixtures/assertions were executed or
modified. These tests exercise archived bridge contracts, **not current minimal
launcher coverage**. Removing gopher-lua, the loader scripts or snapshot would
break real retained consumers. Future retirement needs a user decision on
historical regression obligations, explicit assertion-preservation/adaptation
scope and separately authorized validation, not restoration of a production
callback requirement.

### 4.4 Historical experiment modules and docs provenance

**Classification: historical/test-only** for
`docs/experiments/rex-validation/tui/` and its nested replacements. Its root
`go.mod` is `example.com/whichkey-research`, Go1.25.5, not module `temp`.
`cmd/{direct,tea}/main.go` import internal/probe; Tea imports
`charm.land/bubbletea/v2 v2.0.8`. Its own Taskfile builds direct/tea and runs
nested tests/vet; `restore.sh` uses WK_BINARY and stty to compare terminal modes.
These are real historical tool/test consumers, not active product entrypoints.
Nothing was run here.

The nested module declares pty/testify/x/term and its own transitive graph/sums,
with replacements for `github.com/charmbracelet/x/ansi v0.11.7` and pinned
`github.com/charmbracelet/ultraviolet v0.0.0-20260703014108-f5a850f9c2b7`
to `third_party/{ansi,ultraviolet}`. Each replacement has its own go.mod/go.sum,
license and PROVENANCE.md; ANSI is Go1.24.2, UV Go1.25.0. Its sums also record
colorprofile, term/termios/windows, displaywidth/uax29, go-runewidth, cancelreader,
terminfo, colorful, x/exp/sync/sys, test helpers and YAML. They belong to this
separate experiment graph, not root product requirements. Copied platform files,
fixtures and upstream workflows have nested/historical consumers and cannot be
judged by the current Darwin root alone.

Root Task `./...` follows root module boundaries, not these nested modules;
package source filter excludes their descendants. Exclusion from package/runtime
therefore has positive structural evidence, but **does not mean delete them**.
`docs/acceptance.md` “Preserved earlier experiments” links VALIDATION/A02/A03/A04.
The provenance documents preserve the local experimental corrections, upstream
identity/licenses, changed upstream assertions and their rationale. A04 records
historical green checks awaiting review, not framework adoption or GUI/release
certification; earlier failures and unexplained poll-cancellation history remain.
Git's latest subtree change is `1deb79a9bcbc297875b855e317a26dc61aff7bf3`.
No historical suite result is promoted to a current root pass.

`docs/rex-whichkey.prd` is a historical draft/RFC source, last changed at
`d0ac58193614de2b9fbc45b3333b0850ec6c65b5`. Its “Agreed” Lua configuration/
callback prose and illustrative launcher architecture are not current approval
of executable Lua; the accepted CLI-only design explicitly discusses this
provenance. It is outside package top-level `*.md` copying because its suffix
is `.prd`. Preserve it and managed PRD history; a later documentation index can
label/link the original rather than silently rewriting it as current agreement.

## 5. Documentation and status audit

### 5.1 What Phase16 actually recorded

Read parent `.mtn-tmp/p16/phase16-summary.json` (`rex-whichkey-p16/summary/1`,
with milestone `rex-whichkey-p16/milestone/1`) and
`final-main-public-runtime-evidence.json`. They record the final normal public
binary on Darwin arm64 at commit
`9fa9310b1398672fdc0f78823d96cff0f6ec0292`, tree
`51327ec9528c6060ebb34fb7cc1710ac57cbe91b`, binary SHA256
`76be2e1ee915a91409ade98d09f41f2a9298289fdce385b5dbdde6a36bc40e3d`.
The recorded observations are root menu, Go dynamic Sessions → Switch inventory
with exact IDs/current marker, Backspace, Escape, exit0, primary terminal and
exact owned popup/block/private-directory cleanup, followed by closure of only
the disposable validation window.

Available supporting files read directly include `final-main-menu-root.txt`,
`final-main-menu-dynamic-sessions.txt`, `final-main-exit.json` (exit_code0) and
`final-main-parent-modes.json` (primary screen and restored modes). The evidence
index also points to after-session/process/owned-window cleanup captures; the
index's full cleanup conclusion is recorded coordinator evidence, not a new
observation or an independently reproduced live result by this author.

`task3-public-runtime-evidence.json` separately records public ordinary invocation
at `c46705bfe0b0664ccb27d9cbc41fcf0443e0b816`, binary SHA256
`9a6b24747382db84d759da323e0cb281672c65d5567efa961aca1d0f732fe979`, with an
actual inline rename of exact owned disposable window
`window:01a1b311-43e5-7b85-bdca-c3015f9bbc6d` to
`p16-owned-cli-rename`, exit0, focus/primary restoration and exact cleanup.
It also records a separate root/dynamic-navigation/dismissal run. Git comparison
confirmed **no non-.tasks tracked differences** from that candidate to final
Phase16 runtime commit, or from final runtime commit to inspected audit base.
That supports source continuity; it is not a new binary run or proof that all
installed/configured variants behave alike.

The summary records final build/vet and Nix syntax/format checks, **not a Nix
package build/installation or full test run**. Optional full lint failed:
`final-main-optional-lint.log:204` says **62 issues**, and the summary records
exit201 with AI lint not reached. These are prior coordinator results, not gates
run by this leaf. No global tests/lint green claim is justified.

Still unverified: physical keybinding, minimal Lua in an actually deployed Rex
configuration, package/HM evaluation/build/activation/conflict behavior for this
cutover, user profile/PATH migration, JSON/provider/action matrix, selected
session effects, session rename and other current action effects, Linux/Intel
Darwin/remote runtime, crash/error/moved-origin matrix and all terminal/theme/
geometry variants. Prior user reports and older action trials remain evidence
for their own revisions, not this entire matrix. Native chrome REQ-026 remains
unresolved. Phase15's failed original bridge runtime and unrelated incomplete
phases are not retroactively completed by Phase16 replacement or Phase17 audit.

### 5.2 Documentation disposition matrix (recommendations only)

| Candidate / exact locator | Classification | Confirmed finding, affected readers, bounded adaptation and remaining validation |
| --- | --- | --- |
| `README.md:8-14`; `docs/acceptance.md:10-17,61-75` | required production | Active status still says final combined review/runtime pending and presents only Task2 prerequisite observation. This is stale relative to the later Phase16 summary/runtime. Recommend a separately approved evidence-only status refresh: identify final revision/binary and owned rename evidence, preserve Task2 provenance, enumerate unverified key/config/platform/lint areas. Do not simply replace “pending” with “all passed.” These files are installed. |
| `docs/acceptance.md:55-59` “Allowed leaf gates”; `:61` “Pending final combined coordinator gate” | required production | Historical Task3 execution instructions sit in active present tense; they are not Phase17 permission to build/Nix/run. Recommend relabeling them as Task3 revision-scoped gates/recorded outcome, then linking final Phase16 evidence. No additional run or old test requirement is scheduled by this report. |
| `examples/README.md:88-93`; `docs/actions.md:66-74`; `docs/appearance.md:207-217`; `docs/{launch-contract,runtime-bootstrap,cli-gui-delegation,remote-compatibility,session-owned-ui,tab-rename-input}.md` evidence/history headings | required production | Repeated Task2-only/final-pending phrasing has drifted. Rename guide still says owned rename result pending. Preserve truthful Task2 scope but link one canonical later evidence ledger; status synchronization should not rewrite routing/schema contracts or imply physical-key/HM/remote success. Example's own binding/JSON remains unverified despite direct-menu success. |
| `docs/mapping-format.md:13-19`; `config.go` DecodeConfig comment | unresolved | Docs group legacy DecodeConfig with unchanged source-test retention; actual current source has no caller for this exported decoder. Other legacy test consumers are real. A later comment/doc correction should distinguish uncalled exported compatibility surface from positively test-consumed wire helpers. External imports remain unknown; no deletion based on prose. |
| Repeated migration/path/action explanations across README, examples, data/launch/discovery/remote/session/rename docs | required production | Deliberate audience overlap and cross-links are real consumers, not confirmed useless duplication. Drift already affects status. Recommend central status/evidence link and concise cross-references while retaining each guide's relevant contract and installed readability. Review links and wording after any separately approved consolidation; do not remove contract details merely because repeated. |
| `docs/acceptance.md:94-933` historical details; `docs/actions.md:76-466` historical action audit | historical/test-only | Explicit details wrappers label earlier routes/counts/pending statements as historical, including failed/partial attempts, old app chooser/numbered tabs and original-VM guidance. These are not active instructions; preserve verbatim revision/provenance and negative outcomes. Any future split/archive needs stable source links and explicit history preservation, not blanket status search/replace. |
| `docs/rex-whichkey.prd`, experiment VALIDATION/provenance docs, original paths in historical acceptance tables | historical/test-only | Historical consumers are established above. Some historical links now point at retired production Lua paths because the prose predates fixture relocation; annotate a new historical index/commit locator if approved, rather than pretending original path means live dependency. Do not relabel inherited Agreed prose as present user approval. |
| Installed-relative links in README/docs/examples and historical evidence | unresolved | `nix/package.nix` copies README/docs to share/doc/rex-whichkey, examples to separate share/rex-whichkey/examples, no root Go/Nix/tests/.tasks/.mtn-tmp/experiment trees. README's `docs/...` and intra-doc links resolve, but README `examples/...` and `tests/packaging/...`, docs `../source.go` / `../nix/...` / `../examples/...`, and examples `../docs/...` do not match installed layout. Historical experiment/task/artifact links are absent too. Source-relative links can be valid in Git yet fail installed. Existing finding REX-P16-T03-001 remains. Choose source permalink versus intentional installed cross-tree links/packaged guide in a later doc scope; inspect constructed paths without executing product, and separately authorize any package change. |

A status correction is proposed only inside this report; **none of the audited
documents was edited**. Installed-link limitations do not warrant copying tests,
workflow/private evidence, source trees or historical experiments into the package
without an explicit package/documentation decision.

## 6. Bounded follow-up package and approvals remaining

**◇ Proposal 6a–6d — Open, `REXWHICHKEY-20261106-194853`.** This is one
recommended follow-up package with separable groups, not an accepted cleanup
plan or a request to pause this authorized audit. The ID was assigned during
this audit in America/Toronto, not retrospectively to earlier decisions.
The user/coordinator must explicitly authorize the applicable future scope.
No removal, adaptation or claimed performance benefit follows from this report.

| Group | Bounded future scope / retention rationale | Required adaptation, risk and smallest remaining approval or validation |
| --- | --- | --- |
| 6a — evidence/status and installed-link docs | Refresh the exact active status paragraphs listed in §5 against Phase16 ledger; label old gates correctly; preserve all history. Prefer one canonical evidence source over duplicating status. | Explicit later docs-edit scope. Decide installed-link strategy before packaging/source-link changes. Source/Git/link inspection can establish path consistency; no test/build/runtime permission implied. Do not publish a new pass or mark failed phases complete. |
| 6b — developer entrypoints | Ask about Docker/gotools/editor use and whether make alias/Python pre-commit remain useful. Assess shellHook entry-time effects separately from lint rules. | User choice and exact ownership for any tooling change; preserve direct Task build/lint/unit/race standards and config bytes. A hook-removal/consolidation change affects external workflows and must not be slipped into docs work. Parent Darwin NuGet fix stays separate and uncommitted unless parent explicitly handles it. |
| 6c — legacy compatibility and historical regression obligations | Decide exported Completion/capability/DecodeConfig compatibility and which historical tests/snapshot must remain. Prefer retention until obligations and external consumers are known. Keep all production helpers identified in §4.1. | Explicit API/test-retention decision, authorized assertion-preserving adaptation plan and validation budget **before** any removal. Existing rename/diagnostic/Lua assertions cannot simply disappear or be weakened to allow deletion. Current no-test policy makes behavior validation unavailable here; that is a remaining condition, not a fabricated blocker to the report. |
| 6d — isolated low-risk housekeeping / historical index | The repeated `.gitignore` rule is the sole confirmed redundancy. Optionally add a historical-source index without altering original PRD/experiment/failed evidence. | Separate tiny owned change and source/Git review; tests/history remain intact. No global cache/worktree/history cleanup or archive deletion. No module-lock pruning based on sum-only names. |

**? Validate:** external/editor/public library consumers and deployed environments
remain unknown. Resolve by user-supplied workflow information or separately
authorized local inspection, not by assuming non-use or launching experiments.
No performance, asynchronous/deferred-call/cache, executable-copy/hash/server-check
or safety-policy redesign is contained in any group. Phase18 remains separate.

## 7. Criterion and requirement traceability

### Acceptance criteria 1–8

| Criterion | Report evidence and honest boundary |
| --- | --- |
| 1 — bound inventory | §1 exact baseline/tree, actual checkout/base/tree and metadata-only difference; §§3–4 declarations/locks/tools/root/cmd/rex/Lua/package/assets/tests/experiments; §3.3 parent-only diff and exclusions |
| 2 — consumer tracing / six classifications | §2 vocabulary; matrices §§3–5 supply paths/symbols/line or revision-stable locators, positive callers, build/test/docs/package/editor limits; serialization/reflection/platform/embed/module boundaries explicitly considered |
| 3 — developer tools / side effects | §3.1 Docker, make wrapper, optional Python versus required direct Task hook, gotools uncertainty, required Go/Task/lint/shared-config chain and entry-time pull/hooksPath/copies |
| 4 — legacy versus live consumers | §§4.1–4.4 exact Completion/capability/DecodeConfig/launchOutput/readPopupResult/nativeTextSelection/nativeRenameMapping trace, live helpers, twelve-file fixture identity, gopher-lua script chain and nested experiment provenance |
| 5 — current and historical docs | §5 stale active status and old gate prose, duplicated explanations, historical wrappers/provenance and installed path limits; actual Phase16 public menu/navigation/owned rename versus unverified scope |
| 6 — disposition / follow-ups | All matrices include retention rationale, consumers/contracts/platforms, risks/adaptations and next validation; §6 four bounded groups, open proposal identity and explicit remaining approvals, no cleanup approval implied |
| 7 — preservation | Only this report is changed; parent diff inspected not copied; no source/tool/dependency/lock/test/fixture/history/config/status fix, environment entry, performance work or prohibited execution; final Git scope check binds this claim |
| 8 — traceability / verification / fresh review | This section and §8 ledger identify seven REQs, actual checks and limitations. Author review is **not** independent acceptance; fresh reviewer must inspect frozen report and referenced source, with findings/revision recorded by coordinator before completion. |

### Seven scoped REQs (maintenance traceability, not new functionality)

| Stable ID | Audit coverage / retained interpretation |
| --- | --- |
| REQ-001 — rewrite/process lifecycle | §§1,4 preserve flat short-lived public Go path/minimal Lua and distinguish required private PTY transport from retired bridge; §5 recorded dismissal/cleanup only, no daemon/new architecture |
| REQ-002 — typed CLI client | §§3.3,4.1 trace Client/Session/Window/Block, generic Call/reflection/JSON/argv and concrete callers; no shell interpolation, alternate control plane or API removal approved |
| REQ-003 — configuration | §§1,4 trace LoadMenuConfig data path versus legacy DecodeConfig/archive. Lua loading mechanism superseded by accepted CLI-only design; unsupported old executable data not silently discarded |
| REQ-004 — mappings | §4 retains shared validation/keys/conditions/hidden/confirmation and action-based rename; executable/function/callback portions deferred, not inferred supported from legacy types/tests |
| REQ-014 — packaging/config ownership | §3.3 package/HM exports, native pins/platforms/one non-forcing launch link, examples/docs filters, user state boundaries; §5 installed-link/deployment limitations |
| REQ-015 — delivery/evidence | §§5,7–8 exact prior observed revisions versus current static audit; preserve failed history, lint failure and prohibited tests; independent review/integration remain coordinator responsibilities |
| REQ-016 — direct launch | §§4.1,5 normal no-argument public path and recorded menu/navigation/cleanup; minimal binding reaches same path in source but physical key/config/profile deployment unverified; no old registry prerequisite reinstated |

REQ-005 is deliberately **not** covered as implemented callbacks/hooks. No
renumbering, retrospective approval of original illustrative APIs, or completion
of unrelated requirements is asserted.

## 8. Verification ledger and residual limits

| Check actually performed by this author | Result / limit |
| --- | --- |
| Read worker contract and exact task through read-only mtn-workflow task; compare dispatch/task semantic hash | Match; no mtn mutation or other worker/coordinator action |
| Git status/rev-parse/ls-files and baseline-to-base path/content diff | Initial leaf clean; exact identities in §1; non-task content identical to baseline |
| Read-only source/config/docs/Git inspection and bounded rg caller/import/platform searches | Findings in matrices; source reasoning, not compiler/static-analyzer or runtime proof |
| Read-only parent flake diff and specific Phase16 local evidence | Local Darwin fix preserved separately; prior runtime/lint outcomes attributed, not rerun |
| Twelve archived Lua files compared with original cd5cbce Git blobs using cmp | All identical; no Lua evaluation or assertions executed |
| Non-.tasks Git comparison of recorded owned-rename candidate → final Phase16 runtime → audit base | No differences; source continuity only, not a new run |
| Required `git diff --check` on report change, plus staged/final base-to-head whitespace/scope inspection | Required author gate; exact final commands/results/commit/tree in build result |
| RED / GREEN / unit / race / TDD / Lua / parser / FIFO harness / benchmarks | **Not run — user restriction. No test passes claimed.** |
| Task/build/lint/Nix/normal run/environment entry/shellHook/hooks/config/reload/deploy/live/remote | **Not run by author — user restriction.** Historical Phase16 results are separate in §5. |
| Independent frozen-report/source review | Pending fresh reviewer via coordinator; this report and author inspection do not satisfy it |

The author will commit only `docs/unnecessary-parts-audit.md`, using per-command
`git -c core.hooksPath=/dev/null commit` as explicitly instructed, then freeze
the checkout and return exact base/head/tree/files and actual gate outcomes.
That local builder commit is not target integration, the coordinator's phase
commit, review approval or task/phase completion. No user/editor/platform unknown
has been converted into non-use, a pass, an accepted design or deletion authority.
