# Experimental x/ansi copy

Upstream: https://github.com/charmbracelet/x/tree/ansi/v0.11.7/ansi
Module: github.com/charmbracelet/x/ansi v0.11.7
Go module sum: h1:kzv1kJvjg2S3r9KHo8hDdHFQLEqn4RBCb39dAYC84jI=
Copied from the installed read-only module cache without modifying that cache.
All upstream files, tests, and MIT license are retained.

Original parser_decode.go SHA-256:
8f52f1d95595bd3f00527a2afbfcb682bbe3b5a9bd538b8b3e8fe5c2e3a5b9fd
LICENSE SHA-256:
0baa4a5942a801d93f4050d2bc0039d8c037105fcd335230d118eec99a11782b

The only upstream source change is the printable-ASCII branch in
parser_decode.go: preserve the ASCII fast path only before another ASCII byte
or EOF; otherwise use the existing FirstGraphemeCluster helper with the caller's
width mode. Controls and escapes retain their existing boundaries.

Only the enclosing disposable module replaces this dependency. Root dependencies,
shared cache, and Go 1.25.5 toolchain are unchanged. This is a correction experiment,
not an adopted fork or final terminal-framework decision. The buffer and input
layers have independent acceptance requirements.
