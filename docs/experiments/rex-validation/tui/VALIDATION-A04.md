# Correction attempt 4: required gates GREEN; awaiting independent review

Run `rex-whichkey-p01-t01-builder-a04`, dispatch
`rex-whichkey-p01-t01/dispatch/4`, build event `rex-whichkey-p01-t01/build/4`.
Accepted scope remains `REXWHICHKEY-20261030-175720`. This corrects
T01-REVIEW-009 from review/3; it does not adopt a framework or complete the task.

## Identity and preserved boundaries

Start: `8107b1e46583e468997ba45a8f1231af948b47aa`.
Parent main: `d0ac58193614de2b9fbc45b3333b0850ec6c65b5`, unchanged.
Final committed head/tree are returned in build/4. Changes are confined to this
nested directory: one production file (`third_party/ultraviolet/styled.go`),
one new probe test, this validation document and UV provenance.

Historical `VALIDATION.md`, `VALIDATION-A02.md`, `VALIDATION-A03.md` and previous
logs remain untouched. Pinned copies/licenses, module/checksum files, root
source/dependencies, Taskfiles, toolchain, lint policy and module-cache source
are unchanged. No agents, Rex/session/window/GUI/provider/remote actions,
managed-record writes, integration, merges or pushes occurred.

## Failing first and minimal correction

Evidence directory (logs/modfiles are not committed):

`/var/folders/r3/zcjrn53x42n1zybs9hx__h_w0000gn/T/rex-a04.VM2DbO`

Before any implementation changes, both initial and expanded final test tables
failed with exit 1 (`regression-red.log`, `final-regression-red.log`):

```sh
go test -count=1 -v -timeout=30s -run 'TestStyledUnicodeMetadata|TestStyledUnicodeControlBytes' ./internal/probe
go test -count=1 -timeout=30s -run 'TestStyledUnicode' ./internal/probe
```

They reproduce the exact England flag tag loss in GraphemeWidth across SGR and
OSC-8 BEL/ST, including tag partition points and wrapped-cell geometry. Zero-width
space/word-joiner comparisons expose the same overly restrictive classification.
The new negative cases additionally expose invalid UTF-8 bytes entering cells
in both modes. C0/C1 and existing ANSI-operation safety comparisons already pass
on the starting revision; those passes are controls, not claimed new failures.

`styled.go` now distinguishes genuine Unicode content from executable controls
using UTF-8 validity and standard-library `unicode.IsControl`, not an extender
whitelist. Supported SGR/OSC-8 remain metadata. Invalid bytes within a decoded
positive-width cluster are removed without synthesizing U+FFFD; neighboring valid
marks survive and the existing selected width method is reapplied. A genuinely
encoded U+FFFD remains intact. There is no Unicode normalization, substitute base,
forced mode, width-algorithm change or scanner redesign.

The first focused post-correction run passes (`focused-green1.log`). All prior
001–008 regression families remain active and GREEN, including the actual-newline
fixture, strict decomposed Unicode, complete sub-area cell extents and redraw,
width-changing marks, unsafe controls, all OSC fragmentation/delay boundaries,
standalone Escape, cancellation, overflow discard/resynchronization and shutdown.
The conservative internal 64 KiB scanner and its tests are byte-identical to A03.

### New regression coverage

`internal/probe/unicode_metadata_test.go` keeps the complete England sequence
explicit as U+1F3F4 plus U+E0067/U+E0062/U+E0065/U+E006E/U+E0067/U+E007F.
Tests assert exact bytes through Lines and Draw, not only rendered appearance.
They cover SGR and OSC-8 BEL/ST, each tag partition point, both modes, original
base bold/link/id versus following foreground/link/id, placeholders, trailing
and newline-following text, sub-area clipping/wrapping, final-row exhaustion,
outside sentinels and redraw. Contiguous and metadata-separated cases are active.

Pinned WcWidth assigns the detached tag run width 1, so it remains a separate
cell; GraphemeWidth joins the zero-width tags to the width-2 base. The assertions
retain this existing mode difference rather than forcing grapheme geometry.
ZWNJ/ZWJ, CGJ, VS16, zero-width space, word joiner and encoded U+FFFD provide
Unicode comparison cases. Negative tests cover C0, raw/encoded C1, DEL, invalid
UTF-8 and unsupported CSI/OSC/DCS while preserving neighboring marks and text.

## Final required checks

All heavy commands ran sequentially in the assigned exclusive 5 GiB slot:
Go 1.25.5 darwin/arm64, `GOTOOLCHAIN=local GOMAXPROCS=4 GOFLAGS=-p=2 GOPROXY=off`,
with `GOMEMLIMIT=4GiB`. No downloads, new profiles or cache-source changes.
Dependency suites reuse the approved isolated locked module cache:

`/var/folders/r3/zcjrn53x42n1zybs9hx__h_w0000gn/T/rex-a02-gomodcache.oPQANm`

Candidate/root gates use the existing default cache. Every final check below
exits 0; exact statuses are also retained in `gates.tsv`.

From this nested directory:

```sh
task build
task test
task lint
WK_STRICT_UNICODE=1 go test -count=1 -v -run TestUnicodeObservation ./internal/probe
go test -race -count=1 -timeout=90s ./...
go test -count=20 -timeout=90s -run 'TestStyledUnicode|TestStyledWidthChangingMarks|TestStyledMarkControls|TestStyledMarksUnicodeNotControls|TestStyledWrapClip|TestStyledAreaExhaustion' ./internal/probe
go test -race -count=3 -timeout=90s -run 'TestStyled|TestReader|TestFragmentedDelayedST' ./internal/probe
```

Logs: `final-nested-{build,test,lint,race}.log`, `final-strict.log`,
`final-unicode20.log`, `final-focused-race3.log`.

Also from this directory, a temporary `original-ansi.mod`/`.sum` copies unchanged
nested module files, drops only the local ANSI replacement and points UV to this
worktree's absolute local path. New regressions pass with original ANSI too:

```sh
go test -mod=readonly -modfile="$E/original-ansi.mod" -count=20 -timeout=90s -run 'TestStyledUnicode' ./internal/probe
go test -mod=readonly -modfile="$E/original-ansi.mod" -race -count=3 -timeout=90s -run 'TestStyledUnicode' ./internal/probe
```

Logs: `final-unicode-original20.log`, `final-unicode-original-race3.log`.
`E` denotes the evidence directory above, not a committed module file.

From the worktree root, `task build`, `task test` (unit and race), `task lint`
and `.githooks/pre-commit` pass (`final-root-*.log`, `final-hooks.log`). The commit
runs the existing hook again without bypass. Root checks do not traverse nested
modules. `shellcheck` and `shfmt -d` on unchanged `restore.sh` pass
(`final-shellcheck.log`, `final-shfmt.log`). No `.go-arch-lint.yml` exists;
the accepted flat experimental structure is unchanged.

Full ANSI normal/race suites, including sixel, pass independently in both the
original pinned module in the isolated cache and the effective local copy:

```sh
go test -mod=readonly -count=1 ./...
go test -mod=readonly -race -count=1 ./...
```

Logs: `final-ansi-{original,effective}-{full,race}.log`. Original ANSI tests only
read that pinned module's source; no cache source was edited.

From the local UV directory, all full suites pass:

```sh
go test -mod=readonly -count=1 -timeout=180s ./...
go test -mod=readonly -race -count=1 -timeout=180s ./...
go test -mod=readonly -modfile="$E/effective.mod" -count=1 -timeout=180s ./...
go test -mod=readonly -modfile="$E/effective.mod" -race -count=1 -timeout=180s ./...
go test -mod=readonly -count=3 -v -run '^TestScannerOSC' .
```

`effective.mod`/`.sum` copy unchanged UV pins/checksums with only an absolute
replacement pointing ANSI to this worktree's local copy. The default UV modfile
uses original ANSI v0.11.7. Logs: `final-uv-{original,effective}-{full,race}.log`,
`final-uv-resource.log`. ANSI and UV `go mod verify` both pass
(`final-{ansi,uv}-verify.log`). A03's 50 full plain/10 full race repetitions per
ANSI environment remain preserved, not relabeled as A04 runs. As authorized for
this narrow styled-only correction, those full repetitions are not rerun.

## Failure history and limits

No intermediate implementation or final gate failure occurred after the minimal
correction. The two failing-first runs remain retained, not included among final
passing checks. Initial diagnostic log extraction using `head` under pipefail
returned nonzero on early pipe closure; subsequent extraction read the retained
log successfully. An initial read of nonexistent `ansi/grapheme.go` returned
ENOENT; inspection located the existing helper in `parser_decode.go`. Neither
administrative error modified source or test results.

A01–A03 and review/3 histories remain authoritative for earlier RED/intermediate
failures. In particular, review/2's transient poll-cancellation failure has an
unconfirmed cause; passing A03/A04 repetitions do not erase it. It did not recur
in A04. Pre-existing intersecting-wide-cell area behavior and OSC-looking
bracketed paste are not claimed fixed.

Rebuilt binary SHA-256 (ignored, not committed):

- direct: `5accbb7ab13e4634a2986b4f940b01561342c935aa800cd04d878a9614e431fa`
- tea: `75e1e8760be4ea4657b8eea1f0b328a1f7dddcbe0aa6828351d53110abe0d067`

PTY tests log rebuilt identities and complete-footer captures. Retained
`direct-build-info.log`/`tea-build-info.log` confirm Go 1.25.5 and Tea's exact
pinned dependencies plus both local replacements. Current source hashes and
unchanged license identity are in UV provenance, which now also documents the
historical omitted upstream `.golangci.yml` without changing root lint policy.

GREEN means ready for fresh independent review only. Terminal bytes and buffer
assertions are not real-client glyph-geometry proof. No GUI qualification or
visible base-less mark ownership/fallback policy was established. Direct remains
diagnostic. Coordinator-only integration and task-completion authority are not
part of this delivery.
