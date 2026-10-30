package uv

import (
	"runtime"
	"strings"
	"testing"
)

// Keep the copied module's native assertion style; its locked dependencies do
// not include a test assertion library.
func TestScannerOSCResources(t *testing.T) {
	scanner := newEventScanner()
	n, events := scanner.scanEvents([]byte("\x1b]999;"), false)
	if n != 6 || len(events) != 0 {
		t.Fatalf("pending OSC must transfer ownership: n=%d events=%v", n, events)
	}
	chunk := []byte(strings.Repeat("a", 4096))
	var before, after runtime.MemStats
	runtime.ReadMemStats(&before)
	for range 256 {
		n, events = scanner.scanEvents(chunk, false)
		if n != len(chunk) || len(events) != 0 {
			t.Fatalf("payload must be consumed without keys: %d %v", n, events)
		}
		if len(scanner.osc) > maxPendingOSC || cap(scanner.osc) > maxPendingOSC {
			t.Fatal("pending OSC exceeded the internal memory bound")
		}
	}
	runtime.ReadMemStats(&after)
	// A generous linear-allocation ceiling, not a timing benchmark. The previous
	// repeated-prefix path allocated ~281 MB for this 1 MiB input in review.
	delta := after.TotalAlloc - before.TotalAlloc
	t.Logf("1 MiB OSC allocation: %d bytes; retained len/cap: %d/%d", delta, len(scanner.osc), cap(scanner.osc))
	if delta > 16<<20 {
		t.Fatalf("1 MiB OSC allocated %d bytes", delta)
	}
	allocs := testing.AllocsPerRun(100, func() {
		n, events = scanner.scanEvents(nil, true)
		if n != 0 || len(events) != 0 {
			t.Fatal("idle timeout emitted input")
		}
	})
	if allocs != 0 {
		t.Fatalf("idle pending OSC allocated %v times", allocs)
	}
	n, events = scanner.scanEvents([]byte("\x1b"), false)
	if n != 1 || len(events) != 0 {
		t.Fatal("split ST leaked Escape")
	}
	n, events = scanner.scanEvents([]byte("\\y"), false)
	if n != 2 || len(events) != 1 || events[0] != (KeyPressEvent{Code: 'y', Text: "y"}) {
		t.Fatalf("overflow resync: %d %#v", n, events)
	}
}
