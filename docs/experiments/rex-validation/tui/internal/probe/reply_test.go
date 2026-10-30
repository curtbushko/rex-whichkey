package probe_test

import "testing"

func TestColorReplyIsNotEscapeKey(t *testing.T) {
	for _, name := range []string{"direct", "tea"} {
		t.Run(name, func(t *testing.T) {
			tr := startTerminal(t, name)
			tr.until(t, "RESEARCH READY")
			tr.send(t, "\x1b]11;rgb:2222/3333/4444\x07y")
			// The direct baseline intentionally lacks a terminal-response parser.
			// Record its limitation instead of pretending it is feature-equivalent.
			want := "choose"
			if name == "direct" {
				want = "cancel"
			}
			tr.until(t, "RESEARCH RESULT "+want)
			tr.finish(t, 0)
		})
	}
}
