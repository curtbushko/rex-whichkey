package uv

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestScannerOSCSuffix(t *testing.T) {
	for _, end := range []string{"\a", "\x1b\a", "\x1b\x9c", "\x1b\x18", "\x1b\x1a"} {
		t.Run(end, func(t *testing.T) {
			scanner := newEventScanner()
			scanner.scanEvents([]byte("\x1b]999;"), false)
			_, events := scanner.scanEvents([]byte("ok"+end+"\x1b[0;0;121;1;0;1_"), false)
			want := []Event{KeyPressEvent{Code: 'y', Text: "y"}}
			if end == "\a" {
				want = append([]Event{UnknownOscEvent("\x1b]999;ok\a")}, want...)
			}
			if !reflect.DeepEqual(events, want) {
				t.Fatalf("post-OSC input: got %#v, want %#v", events, want)
			}
		})
	}
}

func TestScannerOSCLimit(t *testing.T) {
	for _, size := range []int{maxPendingOSC - 1, maxPendingOSC, maxPendingOSC + 1} {
		for _, intro := range []string{"\x1b]", "\x9d"} {
			t.Run(fmt.Sprintf("%x/%d", intro, size), func(t *testing.T) {
				scanner := newEventScanner()
				sequence := intro + "999;" + strings.Repeat("a", size-len(intro)-5) + "\a"
				if len(sequence) != size {
					t.Fatal("invalid fixture size")
				}
				var got []Event
				for i := 0; i < len(sequence); i++ {
					n, events := scanner.scanEvents([]byte(sequence[i:i+1]), false)
					// A lone 7-bit ESC/Alt+] needs the caller's tiny ambiguity buffer.
					if intro == "\x1b]" && i == 0 {
						n, events = scanner.scanEvents([]byte(sequence[:3]), false)
						if n != 3 || len(events) != 0 {
							t.Fatalf("intro transfer: %d %v", n, events)
						}
						i = 2
					} else if n != 1 {
						t.Fatalf("unconsumed payload: %d", n)
					}
					got = append(got, events...)
				}
				if size <= maxPendingOSC {
					if len(got) != 1 || got[0] != UnknownOscEvent(sequence) {
						t.Fatalf("bounded complete OSC was not preserved: %d events", len(got))
					}
				} else if len(got) != 0 {
					t.Fatalf("overflow leaked %d events", len(got))
				}
				n, events := scanner.scanEvents([]byte("y"), false)
				if n != 1 || len(events) != 1 || events[0] != (KeyPressEvent{Code: 'y', Text: "y"}) {
					t.Fatalf("resync failed: %d %#v", n, events)
				}
			})
		}
	}
}
