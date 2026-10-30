package probe_test

import (
	"fmt"
	"testing"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/stretchr/testify/assert"
)

func TestStyledAreaExhaustion(t *testing.T) {
	for _, mode := range []ansi.Method{ansi.WcWidth, ansi.GraphemeWidth} {
		for _, tc := range []struct {
			name, text, tail string
			width, height    int
			wrap             bool
			rows             []string
		}{
			{"clip", "e\u0301界x", "", 2, 2, false, []string{"e\u0301 ", "  "}},
			{"wrap", "e\u0301界x", "", 2, 2, true, []string{"e\u0301 ", "界"}},
			{"last-row-clip", "ae\u0301界x", "", 3, 1, false, []string{"ae\u0301 "}},
			{"last-row-wrap", "ae\u0301界x", "", 3, 1, true, []string{"ae\u0301 "}},
			{"too-wide", "界x", "", 1, 2, true, []string{" ", " "}},
			{"tail", "ab界x", ".", 3, 1, false, []string{"ab."}},
			{"wide-tail", "abc界x", "界", 3, 1, false, []string{"a界"}},
			{"tail-too-wide", "ax", "界", 1, 1, false, []string{" "}},
		} {
			t.Run(fmt.Sprint(mode)+"/"+tc.name, func(t *testing.T) {
				buf := uv.NewScreenBuffer(9, 5)
				buf.Method = mode
				buf.Fill(uv.NewCell(mode, "#"))
				// Include a wide sentinel, with its placeholder, wholly outside the area.
				buf.SetCell(6, 2, uv.NewCell(mode, "外"))
				before := buf.Clone()
				area := uv.Rect(1, 1, tc.width, tc.height)
				ss := uv.NewStyledString(tc.text)
				ss.Wrap, ss.Tail = tc.wrap, tc.tail
				ss.Draw(buf, area)
				for y, row := range tc.rows {
					want := uv.NewScreenBuffer(tc.width, 1)
					want.Method = mode
					uv.NewStyledString(row).Draw(want, want.Bounds())
					for x := range tc.width {
						assert.Equal(t, want.CellAt(x, 0), buf.CellAt(x+1, y+1), "(%d,%d)", x+1, y+1)
					}
				}
				assertOutside := func() {
					for y := range buf.Height() {
						for x := range buf.Width() {
							if !uv.Pos(x, y).In(area) {
								assert.Equal(t, before.CellAt(x, y), buf.CellAt(x, y), "outside (%d,%d)", x, y)
							}
						}
					}
				}
				assertOutside()
				// Overwrite and redraw must not leave a wide placeholder behind.
				uv.NewStyledString("z").Draw(buf, area)
				requireCell(t, buf, 1, 1, "z", 1)
				for y := area.Min.Y; y < area.Max.Y; y++ {
					for x := area.Min.X; x < area.Max.X; x++ {
						if x != 1 || y != 1 {
							requireCell(t, buf, x, y, " ", 1)
						}
					}
				}
				assertOutside()
			})
		}
	}
}
