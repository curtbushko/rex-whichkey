# Correction attempt 2: blocked on an unchanged upstream assertion

Run: `rex-whichkey-p01-t01-builder-a02`. Dispatch:
`rex-whichkey-p01-t01/dispatch/2`. Authorization:
`REXWHICHKEY-20261030-175720`. The earlier `VALIDATION.md` is preserved unchanged
as attempt-1 history, not the current dependency description.

## Source and scope

Started at `4e3d16b765781fb5366de1d03f9b147ccc2d452e`; target main remains
`d0ac58193614de2b9fbc45b3333b0850ec6c65b5`. All authored changes and the licensed
102-file exact pinned UV copy are within this nested directory. The only changed
upstream UV files are `styled.go`, `decoder.go` and `terminal_reader.go`; all
upstream tests and module pins remain unchanged. The retained x/ansi correction
is unchanged. Root dependencies, toolchain, Taskfiles and lint configurations are
unchanged. No agents, Rex resources, GUI, remote/provider actions, managed-record
writes, integration or push were performed. No normalization or mode-2027 waiver.

## Failing-first evidence and finding coverage

- T01-REVIEW-004: first replaced the literal backslash-n with an actual newline,
  asserted `x` on row zero and both marks plus `y` on row one, and preserved the
  strict mark assertion. Before any UV copy/change, pinned UV failed by rendering
  `x\ny` without marks in both width modes. Evidence:
  `rex-a02-fixture-red.KNj5dI`.
- T01-REVIEW-001: expanded SGR and OSC-8 BEL/ST tests with trailing marks, following
  text, following newline, narrow/wide preceding bases, attributes and wide
  placeholders in both modes. Leading/wholly base-less marks require byte
  conservation only, not an invented visible fallback. Extra control tests check
  retained BEL/tab/backspace bytes and actual newline boundaries. Red evidence:
  `rex-a02-regressions-red.T7URHS` and
  `rex-a02-pinned-red.w3YQnN/controls-red.log`.
- T01-REVIEW-002: full-cell sub-area tests cover clip/wrap, final-row exhaustion,
  too-wide cells/tails, outside sentinels including a wide cell and placeholder,
  overwrite and clearing on redraw. Original clipping assertions remain active.
  Red evidence: `rex-a02-regressions-red.T7URHS`.
- T01-REVIEW-003: original every-position BEL/ST reader tests and 100 ms ST test
  remain active. Added every established OSC split at 100 ms, 8-bit ST coverage,
  bounded standalone Escape, and delayed-ST PTY selection tests. Every incomplete
  reply must emit no event, a completed reply must be a BackgroundColorEvent, and
  its first following key must be separately sent `y`. PTY tests require choose,
  not cancellation. Red evidence: `rex-a02-regressions-red.T7URHS`.

All these task regressions now pass. One intermediate GREEN attempt still failed
both too-wide-area cases (`rex-a02-regressions-green.WO68Ic`): the following `x`
wrapped after an unplaceable wide cell. The subsequent correction preserves the
clipped-line state until an actual line boundary. No assertion was weakened.

To reproduce the UV RED against the original pin without changing source or
retained evidence, from this nested directory:

```sh
original="$(mktemp -d "${TMPDIR:-/tmp}/rex-a02-red.XXXXXX")"
cp go.mod "$original/original.mod"
cp go.sum "$original/original.sum"
go mod edit -modfile="$original/original.mod" -dropreplace=github.com/charmbracelet/ultraviolet
go test -modfile="$original/original.mod" -count=1 -run 'TestStyled|TestReader' ./internal/probe
```

The absolute temporary modfile still resolves the retained relative ANSI
replacement from this working directory. Removing both replacements additionally
reproduces the original contiguous-cluster failure documented in `VALIDATION.md`.

## Current checks

All commands were sequential in the exclusive heavy slot, with Go 1.25.5,
`GOTOOLCHAIN=local GOMAXPROCS=4 GOFLAGS=-p=2`. Normal candidate/root checks used
`GOPROXY=off` and the existing cache without source edits.

From this nested directory:

```sh
task build
task test
task lint
WK_STRICT_UNICODE=1 go test -count=1 -v -run TestUnicodeObservation ./internal/probe
go test -race -count=1 -timeout=90s ./...
```

All pass. The unique `t.TempDir` capture and complete-footer boundary remain
unchanged. PTY launches log rebuilt binary hashes and effective dependency build
information, including both local replacements. Root `task build`, `task test`
(unit and race), and `task lint` pass independently. The existing commit hook runs
the root gates without bypass; root gates do not traverse nested modules.
No `.go-arch-lint.yml` exists; the accepted flat experimental scope is retained.

Gate logs: `rex-a02-gates.RpyxOr` beneath the runtime temporary directory. The first
root invocation accidentally moved one directory above the worktree: all three
commands returned Taskfile-not-found (exit 100). Corrected invocations from the
actual worktree pass (`root-corrected-*.log`); the failed logs are preserved.
A checksum command also reported the absent root go.sum; no file was created.

## Full dependency suites: one unresolved semantic conflict

The full local ANSI suite, including previously unavailable sixel tests, passes:

```sh
cd third_party/ansi
go test -mod=readonly -race -count=1 ./...
```

The full copied UV suite fails exactly
`TestSplitSequences/unknown_sequence_after_timeout` in unchanged `key_test.go`.
It expects an incomplete OSC at 50 ms to emit UnknownEvent, then `abcxxxx` as seven
keys. The required established-OSC retention instead produces no premature event.
Changing the decoder alone to UnknownEvent would restore the delayed data-loss
bug. There is no special case for this fixture or timer-dependent waiver.

This failure occurs with both effective dependency environments:

1. `cd third_party/ultraviolet; go test -mod=readonly -race -count=1 ./...` uses
   original pinned x/ansi v0.11.7. Log: `rex-a02-uv-original-ansi.wqADxo`.
2. A unique alternative UV modfile copies its unchanged go.mod/go.sum and adds an
   absolute replacement of x/ansi to `../ansi`; then
   `go test -mod=readonly -modfile="$effective/effective.mod" -race -count=1 ./...`
   uses the same ANSI implementation as the candidate. Log:
   `rex-a02-uv-effective.ytzUU7/full-race.log`.

All other UV packages pass. This unchanged assertion was neither skipped nor
edited. Required full dependency suites are not GREEN, so the outcome is blocked,
not ready-for-review, despite the nested regression gates passing. The parent
must resolve this test/requirement conflict explicitly; this leaf has no waiver,
completion, or integration authority.

## Isolated resolution and checksums

Only dependency-suite checks used the run-specific isolated cache:

`/var/folders/r3/zcjrn53x42n1zybs9hx__h_w0000gn/T/rex-a02-gomodcache.oPQANm`

Initial resolution used `GOMODCACHE` set to that path,
`GOPROXY=https://proxy.golang.org GOSUMDB=sum.golang.org`, with the exact existing
locked versions. Subsequent effective-ANSI UV checks used `GOPROXY=off` against
the same isolated cache. `go mod verify` in UV reports all modules verified.
No module pins or sums were changed and no global cleanup was performed.

The newly resolved bitset remains v1.24.4:

- Module sum: `h1:95H15Og1clikBrKr/DuzMXkQzECs1M6hhoGXLwLQOZE=`
- go.mod sum: `h1:7hO7Gc7Pp1vODcmWvKMRA9BNmbv6a/7QIWpPxHddWR8=`
- Proxy ZIP SHA-256:
  `f0677752f880ed313a2c9c055e4d2bd2408f228e12ac7c69ecdb396a00fe0827`

UV original source/license checksums are in `third_party/ultraviolet/PROVENANCE.md`.
Patched UV source SHA-256:

- styled.go: `cecd79f4929d478d28b8a3eded61fbdd1fff15bd32b7f50ca5f27442865af893`
- decoder.go: `279b168f49fd30f8e6535ccef2a3f38104637c68abc403e51c2e965288543a8a`
- terminal_reader.go: `6845ac5524232441cac30ff348a8fe666c23e7c07a9933ba6ac668446e9e4574`

Terminal bytes and buffer assertions are not proof of real-client glyph geometry.
Leading-mark conservation does not settle final public ownership/display policy.
The original direct-I/O program remains a diagnostic, not a qualified candidate.
