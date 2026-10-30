package probe_test

import (
	"fmt"
	"strings"
	"testing"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// These Cf codepoints encode England, not executable C0/C1 terminal controls.
const englandBase = "\U0001f3f4"
const englandTags = "\U000e0067\U000e0062\U000e0065\U000e006e\U000e0067\U000e007f"

var unicodeMetadata = []struct{ name, sequence string }{
	{"SGR", "\x1b[31m"},
	{"OSC8-BEL", "\x1b]8;id=next;https://next.test\a"},
	{"OSC8-ST", "\x1b]8;id=next;https://next.test\x1b\\"},
}

func TestStyledUnicodeMetadata(t *testing.T) {
	for _, mode := range []ansi.Method{ansi.WcWidth, ansi.GraphemeWidth} {
		for _, sample := range []struct {
			name, base, extender string
			baseWidth            int
			wcSeparate           bool
		}{
			{"England", englandBase, englandTags, 2, true},
			{"ZWNJ", "e", "\u200c", 1, false},
			{"ZWJ", "e", "\u200d", 1, false},
			{"CGJ", "e", "\u034f", 1, false},
			{"VS16", "\u2764", "\ufe0f", 1, false},
			{"ZWSP", "e", "\u200b", 1, false},
			{"WordJoiner", "e", "\u2060", 1, true},
			{"EncodedReplacement", "e", "\ufffd", 1, true},
		} {
			for _, metadata := range unicodeMetadata {
				for _, suffix := range []string{"", "x", "\nq"} {
					t.Run(fmt.Sprintf("%d/%s/%s/%q", mode, sample.name, metadata.name, suffix), func(t *testing.T) {
						cluster := sample.base + sample.extender
						prefix := "\x1b[1m\x1b]8;id=base;https://base.test\x1b\\"
						text := prefix + sample.base + metadata.sequence + sample.extender + suffix
						width := sample.baseWidth
						if sample.name == "VS16" && mode == ansi.GraphemeWidth {
							width = 2
						}
						want := []uv.Cell{{Content: cluster, Width: width}}
						// Pinned WcWidth assigns the detached tag run/word joiner
						// width 1. Conserve it without forcing grapheme geometry.
						if mode == ansi.WcWidth && sample.wcSeparate || sample.name == "EncodedReplacement" {
							want = []uv.Cell{{Content: sample.base, Width: width}, {Content: sample.extender, Width: 1}}
							width++
						}
						lines := uv.NewStyledString(text).Lines(mode)
						require.NotEmpty(t, lines)
						firstCells := len(want)
						if suffix == "x" {
							want = append(want, uv.Cell{Content: "x", Width: 1})
						}
						require.Len(t, lines[0], len(want))
						buf := uv.NewScreenBuffer(10, 2)
						buf.Method = mode
						uv.NewStyledString(text).Draw(buf, buf.Bounds())
						var plain strings.Builder
						x := 0
						for i, expected := range want {
							assert.Equal(t, []byte(expected.Content), []byte(lines[0][i].Content))
							assert.Equal(t, expected.Width, lines[0][i].Width)
							requireCell(t, buf, x, 0, expected.Content, expected.Width)
							assert.Equal(t, lines[0][i], *buf.CellAt(x, 0))
							plain.WriteString(lines[0][i].Content)
							assert.Equal(t, uint8(uv.AttrBold), lines[0][i].Style.Attrs)
							if i == 0 || metadata.name == "SGR" {
								assert.Equal(t, "https://base.test", lines[0][i].Link.URL)
								assert.Equal(t, "id=base", lines[0][i].Link.Params)
							} else {
								assert.Equal(t, "https://next.test", lines[0][i].Link.URL)
								assert.Equal(t, "id=next", lines[0][i].Link.Params)
							}
							if i == 0 || metadata.name != "SGR" {
								assert.Nil(t, lines[0][i].Style.Fg)
							} else {
								assert.Equal(t, ansi.Red, lines[0][i].Style.Fg)
							}
							for col := 1; col < expected.Width; col++ {
								requireCell(t, buf, x+col, 0, "", 0)
							}
							x += expected.Width
						}
						if suffix == "\nq" {
							require.Len(t, lines, 2)
							assert.Equal(t, "q", lines[1].String())
							requireCell(t, buf, 0, 1, "q", 1)
							plain.WriteString("\nq")
						} else {
							require.Len(t, lines, 1)
						}
						assert.Equal(t, []byte(cluster+suffix), []byte(plain.String()))
						assert.Contains(t, ansi.Strip(buf.Render()), cluster)
						if suffix == "x" {
							requireCell(t, buf, width, 0, "x", 1)
							assert.Equal(t, "x", lines[0][firstCells].Content)
						}
						uv.NewStyledString("ab").Draw(buf, buf.Bounds())
						requireCell(t, buf, 0, 0, "a", 1)
						requireCell(t, buf, 1, 0, "b", 1)
						assert.NotContains(t, buf.Render(), sample.extender)
					})
				}
			}
		}
	}
}

func TestStyledUnicodeTagGeometry(t *testing.T) {
	for _, mode := range []ansi.Method{ansi.WcWidth, ansi.GraphemeWidth} {
		for _, metadata := range unicodeMetadata {
			for _, wrap := range []bool{false, true} {
				for _, rows := range []int{1, 2} {
					t.Run(fmt.Sprintf("%d/%s/%t/%d", mode, metadata.name, wrap, rows), func(t *testing.T) {
						buf := uv.NewScreenBuffer(6, 4)
						buf.Method = mode
						for y := 0; y < 4; y++ {
							for x := 0; x < 6; x++ {
								buf.SetCell(x, y, uv.NewCell(mode, "#"))
							}
						}
						area := uv.Rect(1, 1, 3, rows)
						text := uv.NewStyledString("ab" + englandBase + metadata.sequence + englandTags + "x")
						text.Wrap = wrap
						text.Draw(buf, area)
						requireCell(t, buf, 1, 1, "a", 1)
						requireCell(t, buf, 2, 1, "b", 1)
						requireCell(t, buf, 3, 1, " ", 1)
						if wrap && rows == 2 {
							content := englandBase + englandTags
							following := "x"
							if mode == ansi.WcWidth {
								content, following = englandBase, englandTags
							}
							requireCell(t, buf, 1, 2, content, 2)
							requireCell(t, buf, 2, 2, "", 0)
							requireCell(t, buf, 3, 2, following, 1)
						} else {
							assert.NotContains(t, buf.Render(), englandBase)
						}
						for y := 0; y < 4; y++ {
							for x := 0; x < 6; x++ {
								if !uv.Pos(x, y).In(area) {
									requireCell(t, buf, x, y, "#", 1)
								}
							}
						}
						uv.NewStyledString("abc").Draw(buf, area)
						requireCell(t, buf, 1, 1, "a", 1)
						requireCell(t, buf, 2, 1, "b", 1)
						requireCell(t, buf, 3, 1, "c", 1)
						assert.NotContains(t, buf.Render(), englandBase)
						assert.NotContains(t, buf.Render(), englandTags)
					})
				}
			}
		}
	}
}

func TestStyledUnicodeTagPartitions(t *testing.T) {
	tags := []rune(englandTags)
	for _, mode := range []ansi.Method{ansi.WcWidth, ansi.GraphemeWidth} {
		for _, metadata := range unicodeMetadata {
			for split := 0; split <= len(tags); split++ {
				t.Run(fmt.Sprintf("%d/%s/%d", mode, metadata.name, split), func(t *testing.T) {
					text := englandBase + string(tags[:split]) + metadata.sequence + string(tags[split:]) + "x"
					lines := uv.NewStyledString(text).Lines(mode)
					require.Len(t, lines, 1)
					assert.Equal(t, []byte(englandBase+englandTags+"x"), []byte(lines[0].String()))
					buf := uv.NewScreenBuffer(10, 1)
					buf.Method = mode
					uv.NewStyledString(text).Draw(buf, buf.Bounds())
					assert.Contains(t, ansi.Strip(buf.Render()), englandBase+englandTags+"x")
					if mode == ansi.GraphemeWidth {
						requireCell(t, buf, 0, 0, englandBase+englandTags, 2)
						requireCell(t, buf, 1, 0, "", 0)
						requireCell(t, buf, 2, 0, "x", 1)
					}
				})
			}
		}
	}
}

func TestStyledUnicodeControlBytes(t *testing.T) {
	// Invalid bytes must not become cell content or synthesized replacement
	// glyphs. A legitimately encoded U+FFFD is covered separately above.
	controls := []string{"\x80", "\x9c", "\xff", "\xc0", "\xc2", "\xed\xa0\x80", "\x1b[2J", "\x1b]0;title\a", "\x1bPdata\x1b\\"}
	for r := rune(0); r <= 0x9f; r++ {
		if r < 0x20 && r != '\n' && r != '\r' && r != '\x1b' || r >= 0x7f {
			controls = append(controls, string(r))
		}
	}
	for _, mode := range []ansi.Method{ansi.WcWidth, ansi.GraphemeWidth} {
		for _, control := range controls {
			t.Run(fmt.Sprintf("%d/%x", mode, control), func(t *testing.T) {
				text := "e\x1b[31m" + control + "\u0301x"
				lines := uv.NewStyledString(text).Lines(mode)
				require.Len(t, lines, 1)
				assert.Equal(t, "e\u0301x", lines[0].String())
				buf := uv.NewScreenBuffer(5, 1)
				buf.Method = mode
				uv.NewStyledString(text).Draw(buf, buf.Bounds())
				requireCell(t, buf, 0, 0, "e\u0301", 1)
				requireCell(t, buf, 1, 0, "x", 1)
				assert.NotContains(t, buf.Render(), control)
				assert.NotContains(t, buf.Render(), "\ufffd")
			})
		}
	}
}
