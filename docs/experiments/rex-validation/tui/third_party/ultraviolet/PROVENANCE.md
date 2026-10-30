# Ultraviolet correction experiment

This is a local copy of `github.com/charmbracelet/ultraviolet` at exactly
`v0.0.0-20260703014108-f5a850f9c2b7`, commit `f5a850f9c2b7`, from
https://github.com/charmbracelet/ultraviolet. It is authorized only for this
short-lived nested experiment by REXWHICHKEY-20261030-175720. No product fork,
framework adoption, public policy or root-module replacement is established.

The 102 upstream files, including all upstream tests and the MIT license, were
copied from the existing Go module cache. Before editing, `diff -qr` confirmed
an identical copy. The upstream module and checksum files remain unchanged.
Only `styled.go`, `decoder.go` and `terminal_reader.go` differ; this document is
additional. Existing upstream lint directives are copied verbatim, not new
suppressions. No shared cache source was edited.

## Original identity

- Module sum: `h1:3FmWoGNWK4STvqg0O0Aeav2T7rodWJAPeF0QpH+8gFw=`
- go.mod sum: `h1:f/jRa757WUmaOZrbPspXymbg/GnbF+rwe4OLsG7aXYo=`
- Cached upstream ZIP SHA-256:
  `50b02b46f486e24d807a9784b68968743b1e502b09d08d9fb9b71484bbc50c1f`
- Unmodified LICENSE SHA-256:
  `8a77c755a1d1fbdc932e86f0449549121820b6f4cc1234b12fcfc8d38013c1de`

## Correction boundary

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

## Deliberately preserved upstream failure

The copied `key_test.go` is unchanged. Its
`TestSplitSequences/unknown_sequence_after_timeout` expects an incomplete OSC to
expire after 50 ms and later bytes to become keys. That conflicts with established
OSC retention across a 100 ms gap. Full upstream suites fail this assertion with
both original x/ansi and the local ANSI correction. It is not skipped, rewritten,
or treated as a pass. See `../../VALIDATION-A02.md`. This prevents a
ready-for-review outcome despite passing the task's nested regression gates.
