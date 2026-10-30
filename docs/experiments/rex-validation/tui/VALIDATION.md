# Bounded terminal correction experiment

Status: blocked, not a qualified candidate or framework adoption. This nested
module is disposable Phase1 evidence for REQ-008 and REQ-009, not product code.
The existing direct program is an unchanged diagnostic baseline. Its raw-byte
parser still treats OSC replies as Escape. Passing its explicitly diagnostic test
does not qualify it. Bubble Tea plus the local ANSI correction is also unqualified.

## Safe capture and identity

Captures use a fresh `t.TempDir` per artifact, never retained `/tmp` paths.
The reader waits for the complete footer, not READY. Every byte split of a frame,
including splits immediately after READY and inside UTF-8, is covered. The harness
RED was observed before replacing that boundary and fixed artifact path. Strict
Unicode RED was then reproduced against rebuilt, unchanged dependencies before
adding the local replacement. Each PTY launch logs the binary SHA-256, Go version,
and effective dependency build information. Callback destinations are not inherited
by tests. Failed and successful processes are reaped; the output reader is stopped.

The only dependency replacement is x/ansi v0.11.7 in this module; provenance and
license are in `third_party/ansi`. All other pins, root dependencies, Taskfiles,
lint configurations, and Go 1.25.5 are unchanged. No normalization or forced
mode-2027 negotiation is used.

## Reproduction

Run from this directory, sequentially in the allocated heavy-gate slot:

```sh
export GOTOOLCHAIN=local GOMAXPROCS=4 GOFLAGS=-p=2 GOPROXY=off
task build
task test
task lint
WK_STRICT_UNICODE=1 go test -count=1 -v -run TestUnicodeObservation ./internal/probe
go test -race -count=1 -timeout=90s ./...
```

The build, vet lint, decoder tables, connected-cell tests, strict decomposed-Unicode
PTY gate, and PTY lifecycle/navigation tests pass. Full test and race commands
remain RED for the strict buffer and input-reader requirements below; no skip or
expected-failure wrapper hides those failures. Passing the Unicode gate alone is
not candidate qualification. The root `task build`, `task test` (including race),
`task lint`, and existing pre-commit hook are additional independent checks; they
do not traverse this nested module.

To reproduce unchanged-dependency RED without editing the replacement or touching
retained evidence, use a unique temporary directory and alternative module file:

```sh
original="$(mktemp -d "${TMPDIR:-/tmp}/rex-task01-original.XXXXXX")"
cp go.mod "$original/original.mod"
cp go.sum "$original/original.sum"
go mod edit -modfile="$original/original.mod" -dropreplace=github.com/charmbracelet/x/ansi
go build -modfile="$original/original.mod" -o "$original/tea" ./cmd/tea
go build -modfile="$original/original.mod" -o "$original/direct" ./cmd/direct
WK_FIXTURE_BINARY_DIR="$original" WK_STRICT_UNICODE=1 go test -modfile="$original/original.mod" -count=1 -v -run TestUnicodeObservation ./internal/probe
go test -modfile="$original/original.mod" -count=1 -run 'TestDecode|TestStyledCells' ./internal/probe
```

The last two commands must fail: unchanged Tea loses U+0301; the ASCII shortcut
splits clusters and advances by one byte instead of the complete cluster length.
The same decoder tables and strict Unicode test pass with the local replacement.
Tables assert string and byte sequences, widths in both modes, consumed lengths,
suffixes, parser state, CJK, keycap/VS16, ZWJ, isolated/multiple marks, and control,
SGR, and OSC-8 boundaries. Buffer tests assert coordinates, wide placeholders,
area clipping/wrapping, overwrite, and redraw clearing.

## Remaining scope blockers

No additional dependency was forked. The pinned Ultraviolet implementation needs
separate authorization or an explicitly approved alternative before correction:

- `TestStyledDetachedMarks`: a mark after SGR or OSC-8 loses its base association
  in both width modes. `styled.go` resets the cell after storing a positive-width
  cluster, appends detached marks to that temporary cell, then overwrites them at
  the next printable cluster. Fixing the ANSI ASCII shortcut cannot join clusters
  across an escape boundary without breaking the decoder's sequence contract.
- `TestStyledIsolatedMarks`: leading U+0301/U+0327 before `x` disappears. A policy
  for ownership/display of isolated marks also needs confirmation; the test only
  requires conservation, not a newly invented glyph or normalization policy.
- `TestStyledWrapClip`: a wide glyph starting in the last column of a sub-area
  crosses its right boundary. The renderer checks the origin, not the full cell
  extent when no truncation tail is configured. The outside screen remains wide
  enough that the underlying buffer cannot enforce this sub-area boundary.
- `TestReaderOSC` and `TestReaderDelayedST`: ST split after its ESC causes
  `EventDecoder.parseOsc` in `decoder.go` to consume the prefix as cancelled; the
  following backslash becomes a key instead of completing a background reply.
  This fails even with a 10 ms gap, inside the default 50 ms Escape timeout.
  A 100 ms ST gap is also tested. BEL and every other split pass. The PTY check
  alone misleadingly passes because the leaked backslash is unbound, so the
  stronger event-level assertion is essential. No incomplete OSC may emit a key;
  after a completed response the first key must be the separately sent `y`.

A bounded Ultraviolet correction would need at least `styled.go` and `decoder.go`
plus tests/provenance. This exceeds the specifically permitted x/ansi experiment;
this attempt stops rather than silently adding that replacement or weakening tests.

## Limits

No GUI, Rex resource creation, remote/provider mutation, final interface decision,
performance-based framework choice, integration, or push occurred. Terminal bytes
and buffer assertions are not screenshots or proof of real-terminal glyph geometry.
Startup samples include PTY/shell overhead, are warm-local observations (n=20),
and now measure full content through the footer, so they are not interchangeable
with older READY-only timings. Preserve older distributions and artifacts.

The copied upstream ANSI main/iterm2/kitty tests pass locally; its supplemental
`go test ./...` cannot run sixel tests offline because the already-pinned
`github.com/bits-and-blooms/bitset v1.24.4` is absent from the shared cache. This
experiment neither downloads it nor changes that pin. The candidate does not
import sixel, and this limitation is not represented as an upstream full-suite pass.
