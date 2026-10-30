# Correction attempt 3: required gates GREEN; awaiting independent review

Run `rex-whichkey-p01-t01-builder-a03`, dispatch
`rex-whichkey-p01-t01/dispatch/3`, build event `rex-whichkey-p01-t01/build/3`.
Authorization remains Accepted `REXWHICHKEY-20261030-175720`. Findings 005–008
are ordinary corrective work within that authorization, not new product choices.
The upstream test changes are semantic corrections, not skips or waivers.

## Identity, ownership and preserved history

Started at `03aad5196717ae28c3c2524b3bfab7888577988b`. Parent main remains
`d0ac58193614de2b9fbc45b3333b0850ec6c65b5`. This attempt changes only files
beneath `docs/experiments/rex-validation/tui`. `VALIDATION.md` (A01) and
`VALIDATION-A02.md` remain byte-identical history. No old worktree or parent
artifact was changed. The exact committed head/tree are returned in build/3.

Pinned copies, licenses, module/checksum files, root dependencies, Taskfiles,
toolchain and lint policy are unchanged. No shared module-cache source edits,
downloads, agents, Rex/GUI/provider/remote actions, managed-record writes,
merges, pushes or integration. No framework adoption or completion authority.

## Failing-first evidence and corrections

Logs are retained outside Git in this run's unique directory:

`/var/folders/r3/zcjrn53x42n1zybs9hx__h_w0000gn/T/rex-a03.KdLKzV`

Before implementation edits, new regressions ran with:

```sh
go test -count=1 -timeout=90s -run 'TestStyledWidthChangingMarks|TestStyledMarkControls|TestStyledMarksUnicodeNotControls|TestReaderRepeatedOSCTimeouts|TestReaderOSCOverflowResync|TestReaderOSCShutdown' ./internal/probe
```

Result: expected RED, exit 1 (`regression-red.log`). It exposes width=1 and
missing placeholders after detached VS16/keycaps in GraphemeWidth, wrong
right-edge layout, executable controls in cells, and overflow payload keys.
The repeated-timeout cancellation and delayed valid-reply tests already pass
under A02 retention; they are strengthened semantic coverage, not claimed new
A02 failures. A test-only int/uint8 attribute assertion mismatch was also present
in this first run and was corrected without changing the expected attribute.

From the local UV module, before implementation edits:

```sh
go test -mod=readonly -count=1 -run TestScannerOSCResources .
```

Result: expected RED, exit 1 (`resource-red.log`): pending OSC consumed zero
bytes instead of transferring ownership to bounded incremental state.

### T01-REVIEW-005: contradictory upstream timeout assertion

The exact old row uses unterminated OSC 11, then `abc`, `x`, `x`, `x`, `x` over
60 ms gaps, and expects UnknownEvent plus seven payload keys. That historical
failure is preserved in A02/build/2/review/2; partial expiration is not restored.

The row now starts with incomplete CSI `ESC[999` and expects that exact
UnknownEvent, keeping its seven keys, 60 ms gaps and eight-event expectation.
New upstream rows retain the exact original OSC/payload chunks until CAN/SUB,
then require only separately sent `y`. Delayed BEL and split-ST rows require
exact BackgroundColorEvent plus separate `y`. Live-reader probe tests explicitly
assert silence between every chunk/time gap, exact color, no extra events and
following-y delivery. Standalone Escape and shutdown regressions remain active.

### T01-REVIEW-006: bounded incremental OSC retention

Established OSC transfers to scanner-owned state; reads consume only new bytes.
The buffer's capacity is bounded at 64 KiB including framing. On overflow it
enters bounded discard until BEL, 8-bit ST, split 7-bit ST, CAN or SUB. Payload
cannot enter key lookup or Win32 deserialization. Following normal input is
parsed again after termination. This is a conservative internal experiment
limit, not a final public interface. No age-based expiration or delayed-reply
waiver was introduced. The reader timer is stopped on exit.

Tests cover one-byte growth, sizes 65,535/65,536/65,537 with both OSC introducers,
1 MiB growth, capacity, allocation, zero-allocation idle scans, overflow with
key-table and Win32-looking payload, every resynchronizer, subsequent valid
reply, cancellation and shutdown before/after overflow. Existing 100 ms delayed
BEL/ST/8-bit-ST and all-position fragmentation tests remain active.

Final resource diagnostic (`final-resource-details.log`): the 1 MiB fragmented
payload allocates 16,384 additional bytes after initial scanner allocation;
retained length/capacity after overflow is 0/65,536. The test uses a generous
16 MiB linear-allocation ceiling and exact capacity checks, not a timing claim.

### T01-REVIEW-007: width-changing detached marks

Placement is deferred until the next base or line boundary. Detached marks
extend the cell's original bytes/metadata and recompute its selected-mode width
before placement; cursor advancement, placeholders, wrapping, tail clipping and
final-row checks use the completed extent. Decoder sequences are not merged
across escapes. Keycaps and hearts with VS16 are width 2 in GraphemeWidth and
width 1 in WcWidth. Tests cover SGR/OSC-8 BEL/ST, base/following attributes,
Lines, right-edge clip/wrap, final-row exhaustion, outside sentinels and redraw.

### T01-REVIEW-008: Unicode content is not executable control content

The zero-width attachment path accepts Unicode marks and ZWNJ/ZWJ, not generic
zero-width terminal sequences. Supported SGR/OSC-8 stay metadata. Unsupported
erase/cursor/OSC/control operations are ignored rather than replayed in cells.
The earlier tab/backspace byte-acceptance assertions were replaced with strict
safety, neighboring-codepoint preservation and geometry assertions. Tests cover
BEL, tab, backspace, erase/cursor CSI, title/clipboard OSC, ESC save-cursor, NUL,
DEL and DCS, including leading controls/marks, actual newline boundaries and
redraw. No normalization or synthetic visible fallback was added.

### Original regression families 001–004

All remain active and GREEN: detached/leading/trailing marks in both modes,
complete sub-area bounds and placeholders/redraw, all valid OSC fragmentation
and 100 ms ST delays with following-y selection, and the actual-newline fixture
preserving row-zero `x` plus row-one marks and `y`. The unsafe control-byte
assertions were corrected, not used to relax these Unicode requirements.

## Final required gates

All heavy commands were sequential in the assigned exclusive 5 GiB slot, using
Go 1.25.5, `GOTOOLCHAIN=local GOMAXPROCS=4 GOFLAGS=-p=2 GOPROXY=off`.
Candidate/root gates use the existing cache. Dependency-module gates reuse the
approved isolated cache, without downloads or pin changes:

`GOMODCACHE=/var/folders/r3/zcjrn53x42n1zybs9hx__h_w0000gn/T/rex-a02-gomodcache.oPQANm`

From the nested module, all exit 0:

```sh
task build
task test
task lint
WK_STRICT_UNICODE=1 go test -count=1 -v -run TestUnicodeObservation ./internal/probe
go test -race -count=1 -timeout=90s ./...
go test -race -count=3 -timeout=90s -run 'TestStyled|TestReader|TestFragmentedDelayedST' ./internal/probe
```

Logs: `final-nested-{build,test,lint,race}.log`, `final-strict.log`,
`final-focused-race3.log`. Safe unique captures and complete-footer waiting are
unchanged. From the worktree root, `task build`, `task test` (unit/race) and
`task lint` all exit 0 (`final-root-*.log`). Root tests do not traverse nested
modules. The existing `.githooks/pre-commit` also exits 0 without bypass,
rerunning root build/lint/unit/race (`final-hooks.log`); the commit invokes it
again. Shellcheck/shfmt on unchanged `restore.sh` pass. There is no
`.go-arch-lint.yml`; the accepted flat experiment scope is unchanged.

From `third_party/ansi`, both full suites exit 0, including sixel:

```sh
go test -mod=readonly -count=1 ./...
go test -mod=readonly -race -count=1 ./...
```

Logs: `final-ansi-full.log`, `final-ansi-race.log`. Both local dependency modules'
`go mod verify` commands exit 0 (`final-{ansi,uv}-verify.log`).

From `third_party/ultraviolet`, all final full repetitions exit 0:

```sh
go test -mod=readonly -count=50 -timeout=300s ./...
go test -mod=readonly -modfile="$E/effective.mod" -count=50 -timeout=300s ./...
go test -mod=readonly -race -count=10 -timeout=180s ./...
go test -mod=readonly -modfile="$E/effective.mod" -race -count=10 -timeout=180s ./...
```

Here `E` is the unique evidence directory above. `effective.mod`/`effective.sum`
copy UV's unchanged pins/sums; the only modfile edit adds an absolute replacement
of x/ansi to this worktree's `third_party/ansi`. The default UV modfile uses
original pinned ANSI v0.11.7. Logs: `final-uv-{original,effective}-{full50,race10}.log`.
Earlier single full/plain/race and 50/plain + 10/race repetitions in both
configurations also passed after the Alt-key correction; their logs are retained.

Review/2 observed one transient poll-cancellation failure, not reproduced in its
50 candidate/50 original focused repetitions. No such recurrence occurred in
this attempt's full repetitions. Its root cause remains unconfirmed; these
passes do not erase it or establish an unconditional single-failure history.

## Preserved intermediate failures

- `regression-red.log` and `resource-red.log`: expected failing-first results.
- `focused-green1.log`, exit 1: first focused run preceded binary build, so both
  delayed-ST PTY subtests lacked `tea`; the attribute int/uint8 test mismatch also
  remained. Building with the existing Taskfile and correcting only the expected
  value's type resolved these harness issues (`focused-green2.log`, exit 0).
- `uv-original-plain1.log`, exit 1: incremental parsing initially lost the
  immediate Alt+] / Alt+' ambiguity in existing `TestReadInput`. The scanner now
  preserves the existing decoder path when ESC follows the introducer before
  OSC data. No upstream assertion was changed to hide this regression. All
  subsequent full suites pass.
- Initial administrative evidence extraction fed mtn's text banner to jq and
  failed parsing; banner removal extracted the complete exact build/2/review/2
  payloads read-only. No managed records were changed.

## Rebuilt binary identity and limits

Both binaries were rebuilt from final source; PTY tests report their hashes and
Go build information. SHA-256:

- direct: `75dbd208019bf2df119d0cddd7bed1106f765588f7b77cd14818b4647f01e1f5`
- tea: `ac1f247da40a74bd7acac2961706d6b906630ec561b9fad53e14a088aebc345a`

Tea's build information confirms Go 1.25.5, Bubble Tea v2.0.8, exact pinned UV
and ANSI versions, and both nested local replacements. Current patched-source
hashes and unchanged license identity are in `third_party/ultraviolet/PROVENANCE.md`.
Binaries, logs, caches and managed records are excluded from the commit.

GREEN gates mean ready for fresh independent review, not integration/completion
or final framework adoption. Terminal bytes and buffer tests are not real-client
glyph-geometry proof. Base-less mark conservation does not decide final visible
ownership/fallback policy. Direct remains a diagnostic, not a qualified candidate.
The pre-existing intersecting-wide-cell area behavior and OSC-looking bracketed
paste behavior reported in review/2 are not claimed fixed by this correction.
