# Ultraviolet correction experiment

This is a local copy of `github.com/charmbracelet/ultraviolet` at exactly
`v0.0.0-20260703014108-f5a850f9c2b7`, commit `f5a850f9c2b7`, from
https://github.com/charmbracelet/ultraviolet. It is authorized only for this
short-lived nested experiment by REXWHICHKEY-20261030-175720. No product fork,
framework adoption, public policy or root-module replacement is established.

At the A02 import, the 102 upstream files, including all upstream tests and the
MIT license, were copied from the existing Go module cache. Before editing,
`diff -qr` confirmed an identical copy. The upstream module and checksum files
remain unchanged. At the A02 snapshot only `styled.go`, `decoder.go` and
`terminal_reader.go` differed; this document was additional. A03's current
source/test differences are listed below. Existing upstream lint directives are copied verbatim, not new
suppressions. No shared cache source was edited.

## Original identity

- Module sum: `h1:3FmWoGNWK4STvqg0O0Aeav2T7rodWJAPeF0QpH+8gFw=`
- go.mod sum: `h1:f/jRa757WUmaOZrbPspXymbg/GnbF+rwe4OLsG7aXYo=`
- Cached upstream ZIP SHA-256:
  `50b02b46f486e24d807a9784b68968743b1e502b09d08d9fb9b71484bbc50c1f`
- Unmodified LICENSE SHA-256:
  `8a77c755a1d1fbdc932e86f0449549121820b6f4cc1234b12fcfc8d38013c1de`

## Correction boundary (historical A02)

`styled.go` preserves zero-width bytes instead of overwriting them. Detached
marks extend the preceding emitted base cell, retaining that cell's attributes
and wide placeholders. Newlines reset this association. Leading bytes stay in
input order in the existing temporary content, or as zero-width content when
there is no following printable cell. This is minimal byte conservation, not a
choice of a visible fallback, normalized text, or final base-less ownership or
display policy. No dotted circle, replacement glyph, or Unicode negotiation
waiver is introduced. Complete cell extents, including truncation tails, must
fit within the drawing area; exhausted rows cannot spill into neighboring cells.

`decoder.go` distinguishes an established incomplete OSC from an ambiguous
Escape key and retains a final ESC that may begin ST. `terminal_reader.go` keeps
that incomplete OSC buffered across Escape timeouts, except inside bracketed
paste, whose existing handling is unchanged. Standalone Escape keeps its bounded
ambiguity timeout. Unterminated established OSC remains buffered until a
terminator, cancellation or reader shutdown; this experiment does not introduce
a new size/age policy.

## Deliberately preserved upstream failure (historical A02)

The copied `key_test.go` is unchanged. Its
`TestSplitSequences/unknown_sequence_after_timeout` expects an incomplete OSC to
expire after 50 ms and later bytes to become keys. That conflicts with established
OSC retention across a 100 ms gap. Full upstream suites fail this assertion with
both original x/ansi and the local ANSI correction. It is not skipped, rewritten,
or treated as a pass. See `../../VALIDATION-A02.md`. This prevents a
ready-for-review outcome despite passing the task's nested regression gates.

## Current A03 correction

`../../VALIDATION-A03.md` supersedes the preceding A02 implementation/gate
status, without changing `../../VALIDATION.md` or `../../VALIDATION-A02.md`.
This is the same accepted experiment authorization, not a new public interface.

- `styled.go` delays placement of a base until its detached marks have been
  associated. Width is recalculated with the selected method before applying
  cell/tail bounds, wrapping and final-row constraints. SGR/OSC-8 remain metadata;
  only Unicode marks and ZWNJ/ZWJ enter the zero-width content path. Unsupported
  executable controls are ignored, not embedded in cells. Original codepoints
  are retained without normalization or a synthetic visible base.
- `terminal_reader.go` transfers established OSC into incremental scanner state.
  It retains at most **64 KiB including introducer and terminator**, using a
  fixed-capacity buffer. Once oversized, it discards until BEL, 8-bit ST, split
  7-bit ST, CAN or SUB; discarded payload cannot become keys or Escape. This is
  a conservative **internal experiment limit**, not a final public interface.
  It scans new bytes once, decodes a bounded complete reply once, and does no
  OSC-prefix copy/reparse on idle timeouts. Win32 deserialization/key lookup
  cannot interpret discarded payload; normal input resumes after termination.
  Standalone Escape/Alt+] ambiguity and explicit cancellation remain supported.
  The reader's timeout timer is stopped on exit. The A02 `decoder.go` is unchanged.
- `key_test.go` is intentionally corrected, not skipped or waived. The old
  eight-event timeout row now uses incomplete CSI `ESC[999`, preserving all
  `abcxxxx` chunks, 60 ms delays and eight-event expectation. Additional rows
  retain the exact old OSC/payload through timeouts until CAN/SUB and require
  only following `y`; delayed BEL/split-ST require exact color plus separate `y`.
- New `osc_resource_test.go` and `osc_limit_test.go` check incremental consumption,
  bounded allocation/capacity, boundary sizes, discard/resynchronization, idle
  work and post-OSC input. They follow the copied module's native test style;
  no assertion-library dependency, module pin or checksum was added.
  Nested probe tests additionally check live-reader timing/shutdown, all four
  original regression families, changed widths, metadata, safe output and redraw.

Current SHA-256:

- styled.go: `7c2af3a02f205b6d34365829ce0901efb6e10931db37007b7b0e3fd9781f4f64`
- decoder.go: `279b168f49fd30f8e6535ccef2a3f38104637c68abc403e51c2e965288543a8a`
- terminal_reader.go: `08a58ea4073f779221c1815ce7b0fd2451ac8e4608cced012c2816091f5b046c`
- key_test.go: `1b615bf750e0414172f993ecbba0f5f1256f6ab8051450a2092d33b39a8f03ed`
- osc_resource_test.go: `ba9dcf9abfe72e582071e1d28d0fcb75c432dda0637a7c163b3e9a122dde9923`
- osc_limit_test.go: `d1186505cf96b05817d4da1a4f5a0053f4e7dd64772c8ff62b512c409ab563f6`

License, upstream go.mod/go.sum, x/ansi copy and root policies are unchanged.
Full UV suites now pass with original pinned ANSI and effective local ANSI,
including 50 full non-race and 10 full race repetitions **per environment** on
final source. This does not explain or erase review/2's earlier transient poll
cancellation failure; no recurrence was observed in these repetitions.
