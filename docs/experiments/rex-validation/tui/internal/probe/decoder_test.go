package probe_test

import (
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeClusters(t *testing.T) {
	for _, tc := range []struct {
		name, cluster, suffix string
		wc, grapheme          int
	}{
		{"ascii", "a", "bc", 1, 1},
		{"decomposed", "e\u0301", "x", 1, 1},
		{"multiple", "e\u0301\u0327", "a\u0308", 1, 1},
		{"trailing", "e\u0301", "", 1, 1},
		{"isolated", "\u0301", "", 0, 0},
		{"leading", "\u0301\u0327", "e", 0, 0},
		{"cjk", "界", "e\u0301", 2, 2},
		{"keycap", "1\ufe0f\u20e3", "x", 1, 2},
		{"vs16", "\u2764\ufe0f", "x", 1, 2},
		{"zwj", "👩\u200d💻", "x", 2, 2},
		{"control-boundary", "e", "\n\u0301", 1, 1},
		{"sgr-boundary", "e", "\x1b[31m\u0301", 1, 1},
		{"osc-boundary", "e", "\x1b]8;;https://example.test\x1b\\\u0301", 1, 1},
	} {
		for _, mode := range []string{"wc", "grapheme"} {
			t.Run(tc.name+"/"+mode, func(t *testing.T) {
				input := tc.cluster + tc.suffix
				wantWidth := tc.wc
				ds, db := ansi.DecodeSequenceWc[string], ansi.DecodeSequenceWc[[]byte]
				if mode == "grapheme" {
					ds, db = ansi.DecodeSequence[string], ansi.DecodeSequence[[]byte]
					wantWidth = tc.grapheme
				}
				seq, width, n, state := ds(input, ansi.NormalState, nil)
				assert.Equal(t, tc.cluster, seq)
				assert.Equal(t, wantWidth, width)
				assert.Equal(t, len(tc.cluster), n)
				assert.Equal(t, ansi.NormalState, state)
				require.LessOrEqual(t, n, len(input))
				assert.Equal(t, tc.suffix, input[n:])
				bseq, bwidth, bn, bstate := db([]byte(input), ansi.NormalState, nil)
				assert.Equal(t, []byte(tc.cluster), bseq)
				assert.Equal(t, wantWidth, bwidth)
				assert.Equal(t, len(tc.cluster), bn)
				assert.Equal(t, ansi.NormalState, bstate)
				require.LessOrEqual(t, bn, len(input))
				assert.Equal(t, []byte(tc.suffix), []byte(input)[bn:])
			})
		}
	}
}

func TestDecodeState(t *testing.T) {
	for _, mode := range []string{"wc", "grapheme"} {
		t.Run(mode, func(t *testing.T) {
			ds, db := ansi.DecodeSequenceWc[string], ansi.DecodeSequenceWc[[]byte]
			if mode == "grapheme" {
				ds, db = ansi.DecodeSequence[string], ansi.DecodeSequence[[]byte]
			}
			for _, fragment := range []struct {
				input string
				state byte
			}{{"\x1b[", ansi.PrefixState}, {"\x1b[31", ansi.ParamsState}} {
				seq, width, n, state := ds(fragment.input, ansi.NormalState, nil)
				require.Equal(t, fragment.input, seq)
				require.Zero(t, width)
				require.Equal(t, len(fragment.input), n)
				require.Equal(t, fragment.state, state)
				seq, width, n, state = ds("me\u0301", state, nil)
				require.Equal(t, "m", seq)
				require.Zero(t, width)
				require.Equal(t, 1, n)
				require.Equal(t, ansi.NormalState, state)
				bseq, bwidth, bn, bstate := db([]byte(fragment.input), ansi.NormalState, nil)
				require.Equal(t, []byte(fragment.input), bseq)
				require.Zero(t, bwidth)
				require.Equal(t, len(fragment.input), bn)
				require.Equal(t, fragment.state, bstate)
				bseq, bwidth, bn, bstate = db([]byte("me\u0301"), bstate, nil)
				require.Equal(t, []byte("m"), bseq)
				require.Zero(t, bwidth)
				require.Equal(t, 1, bn)
				require.Equal(t, ansi.NormalState, bstate)
			}
		})
	}
}

func TestDecodeBoundaries(t *testing.T) {
	for _, mode := range []string{"wc", "grapheme"} {
		for _, control := range []string{"\n", "\r", "\t", "\x1b[31m", "\x1b[0m", "\x1b]8;;https://example.test\a", "\x1b]8;;\x1b\\"} {
			t.Run(mode+"/"+control, func(t *testing.T) {
				ds, db := ansi.DecodeSequenceWc[string], ansi.DecodeSequenceWc[[]byte]
				if mode == "grapheme" {
					ds, db = ansi.DecodeSequence[string], ansi.DecodeSequence[[]byte]
				}
				input := control + "e\u0301x"
				seq, width, n, state := ds(input, ansi.NormalState, ansi.NewParser())
				require.Equal(t, control, seq)
				require.Zero(t, width)
				require.Equal(t, len(control), n)
				require.Equal(t, ansi.NormalState, state)
				seq, width, next, state := ds(input[n:], state, nil)
				assert.Equal(t, "e\u0301", seq)
				assert.Equal(t, 1, width)
				require.Equal(t, 3, next)
				assert.Equal(t, "x", input[n+next:])
				assert.Equal(t, ansi.NormalState, state)
				bseq, bwidth, bn, bstate := db([]byte(input), ansi.NormalState, ansi.NewParser())
				assert.Equal(t, []byte(control), bseq)
				assert.Zero(t, bwidth)
				assert.Equal(t, n, bn)
				assert.Equal(t, ansi.NormalState, bstate)
			})
		}
	}
}
