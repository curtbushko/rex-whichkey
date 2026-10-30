package probe_test

import (
	"fmt"
	"testing"
	"time"
)

func TestFragmentedDelayedST(t *testing.T) {
	reply := "\x1b]11;rgb:2222/3333/4444\x1b\\"
	for _, gap := range []time.Duration{10 * time.Millisecond, 100 * time.Millisecond} {
		t.Run(fmt.Sprint(gap), func(t *testing.T) {
			tr := startTerminal(t, "tea")
			tr.frame(t)
			tr.send(t, reply[:len(reply)-1])
			tr.quiet(t, gap)
			tr.send(t, reply[len(reply)-1:])
			tr.quiet(t, gap)
			tr.send(t, "y")
			tr.until(t, "RESEARCH RESULT choose")
			tr.finish(t, 0)
		})
	}
}
