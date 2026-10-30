package probe_test

import (
	"fmt"
	"testing"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStyledWidthChangingMarks(t *testing.T) {
	for _, mode := range []ansi.Method{ansi.WcWidth, ansi.GraphemeWidth} {
		for _, cluster := range []struct{ base, marks string }{{"1", "\ufe0f\u20e3"}, {"❤", "\ufe0f"}} {
			for _, boundary := range []string{"\x1b[31m", "\x1b]8;;https://example.test\a", "\x1b]8;;https://example.test\x1b\\"} {
				t.Run(fmt.Sprintf("%d/%s/%q", mode, cluster.base, boundary), func(t *testing.T) {
					text := "\x1b[1m" + cluster.base + boundary + cluster.marks + "x"
					width := 1
					if mode == ansi.GraphemeWidth {
						width = 2
					}
					buf := uv.NewScreenBuffer(8, 3)
					buf.Method = mode
					uv.NewStyledString(text).Draw(buf, buf.Bounds())
					requireCell(t, buf, 0, 0, cluster.base+cluster.marks, width)
					assert.Equal(t, uint8(uv.AttrBold), buf.CellAt(0, 0).Style.Attrs)
					assert.True(t, buf.CellAt(0, 0).Link.IsZero())
					if width == 2 {
						requireCell(t, buf, 1, 0, "", 0)
					}
					requireCell(t, buf, width, 0, "x", 1)
					if boundary == "\x1b[31m" {
						assert.Equal(t, ansi.Red, buf.CellAt(width, 0).Style.Fg)
					} else {
						assert.Equal(t, "https://example.test", buf.CellAt(width, 0).Link.URL)
					}
					lines := uv.NewStyledString(text).Lines(mode)
					require.Len(t, lines, 1)
					require.Len(t, lines[0], 2)
					assert.Equal(t, width, lines[0][0].Width)
					assert.Equal(t, cluster.base+cluster.marks, lines[0][0].Content)
					uv.NewStyledString("ab").Draw(buf, buf.Bounds())
					requireCell(t, buf, 0, 0, "a", 1)
					requireCell(t, buf, 1, 0, "b", 1)
				})
				for _, wrap := range []bool{false, true} {
					for _, rows := range []int{1, 2} {
						t.Run(fmt.Sprintf("edge/%d/%s/%q/%t/%d", mode, cluster.base, boundary, wrap, rows), func(t *testing.T) {
							buf := uv.NewScreenBuffer(6, 4)
							buf.Method = mode
							for y := 0; y < 4; y++ {
								for x := 0; x < 6; x++ {
									buf.SetCell(x, y, uv.NewCell(mode, "#"))
								}
							}
							area := uv.Rect(1, 1, 2, rows)
							text := uv.NewStyledString("a" + cluster.base + boundary + cluster.marks + "x")
							text.Wrap = wrap
							text.Draw(buf, area)
							requireCell(t, buf, 1, 1, "a", 1)
							if mode == ansi.WcWidth {
								requireCell(t, buf, 2, 1, cluster.base+cluster.marks, 1)
							} else if wrap && rows == 2 {
								requireCell(t, buf, 2, 1, " ", 1)
								requireCell(t, buf, 1, 2, cluster.base+cluster.marks, 2)
								requireCell(t, buf, 2, 2, "", 0)
							} else {
								requireCell(t, buf, 2, 1, " ", 1)
							}
							for y := 0; y < 4; y++ {
								for x := 0; x < 6; x++ {
									if !uv.Pos(x, y).In(area) {
										requireCell(t, buf, x, y, "#", 1)
									}
								}
							}
							uv.NewStyledString("ab").Draw(buf, area)
							requireCell(t, buf, 1, 1, "a", 1)
							requireCell(t, buf, 2, 1, "b", 1)
						})
					}
				}
			}
		}
	}
}

func TestStyledMarksUnicodeNotControls(t *testing.T) {
	for _, mode := range []ansi.Method{ansi.WcWidth, ansi.GraphemeWidth} {
		for _, text := range []string{"e\x1b[31m\u0301\u0327x", "👩\x1b[31m\u200dx"} {
			lines := uv.NewStyledString(text).Lines(mode)
			assert.Contains(t, lines[0].String(), "x")
			assert.NotContains(t, lines[0][0].Content, "\x1b")
			if text[0] == 'e' {
				assert.Equal(t, "e\u0301\u0327", lines[0][0].Content)
			} else {
				assert.Equal(t, "👩\u200d", lines[0][0].Content)
			}
		}
	}
}
