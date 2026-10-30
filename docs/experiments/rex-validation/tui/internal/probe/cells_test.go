package probe_test

import (
	"fmt"
	"testing"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func requireCell(t *testing.T, buf uv.ScreenBuffer, x, y int, content string, width int) {
	t.Helper()
	c := buf.CellAt(x, y)
	require.NotNil(t, c)
	assert.Equal(t, content, c.Content, "cell (%d,%d)", x, y)
	assert.Equal(t, width, c.Width, "cell (%d,%d)", x, y)
}

func TestStyledCells(t *testing.T) {
	for _, mode := range []ansi.Method{ansi.WcWidth, ansi.GraphemeWidth} {
		t.Run(fmt.Sprint(mode), func(t *testing.T) {
			buf := uv.NewScreenBuffer(10, 3)
			buf.Method = mode
			uv.NewStyledString("界 e\u0301\u0327x").Draw(buf, buf.Bounds())
			requireCell(t, buf, 0, 0, "界", 2)
			requireCell(t, buf, 1, 0, "", 0)
			requireCell(t, buf, 2, 0, " ", 1)
			requireCell(t, buf, 3, 0, "e\u0301\u0327", 1)
			requireCell(t, buf, 4, 0, "x", 1)
			// Redraw must clear old cells and the second half of a wide glyph.
			uv.NewStyledString("ab").Draw(buf, buf.Bounds())
			requireCell(t, buf, 0, 0, "a", 1)
			requireCell(t, buf, 1, 0, "b", 1)
			requireCell(t, buf, 3, 0, " ", 1)
			buf.SetCell(0, 0, uv.NewCell(mode, "界"))
			buf.SetCell(1, 0, uv.NewCell(mode, "x"))
			requireCell(t, buf, 0, 0, " ", 1)
			requireCell(t, buf, 1, 0, "x", 1)
		})
	}
}

func TestStyledWrapClip(t *testing.T) {
	for _, mode := range []ansi.Method{ansi.WcWidth, ansi.GraphemeWidth} {
		for _, wrap := range []bool{false, true} {
			t.Run(fmt.Sprint(mode)+map[bool]string{false: "/clip", true: "/wrap"}[wrap], func(t *testing.T) {
				buf := uv.NewScreenBuffer(6, 3)
				buf.Method = mode
				text := uv.NewStyledString("e\u0301界x")
				text.Wrap = wrap
				text.Draw(buf, uv.Rect(1, 1, 2, 2))
				requireCell(t, buf, 1, 1, "e\u0301", 1)
				requireCell(t, buf, 0, 0, " ", 1)
				if wrap {
					requireCell(t, buf, 1, 2, "界", 2)
					requireCell(t, buf, 2, 2, "", 0)
				} else {
					requireCell(t, buf, 2, 1, " ", 1)
				}
			})
		}
	}
}

func TestStyledIsolatedMarks(t *testing.T) {
	for _, mode := range []ansi.Method{ansi.WcWidth, ansi.GraphemeWidth} {
		for _, text := range []string{"\u0301\u0327", "\u0301\u0327x", "x\n\u0301\u0327y", "\u0301\u0327\nx", "x\n\u0301\u0327"} {
			t.Run(fmt.Sprintf("%d/%q", mode, text), func(t *testing.T) {
				buf := uv.NewScreenBuffer(8, 3)
				buf.Method = mode
				uv.NewStyledString(text).Draw(buf, buf.Bounds())
				// No placement policy is invented here: at minimum, retain the marks.
				assert.Contains(t, buf.Render(), "\u0301\u0327")
				if text == "x\n\u0301\u0327y" {
					requireCell(t, buf, 0, 0, "x", 1)
					assert.Contains(t, buf.Line(1).String(), "y")
					assert.Contains(t, buf.Line(1).String(), "\u0301\u0327")
					assert.NotContains(t, buf.Line(0).String(), "y")
				} else if text == "\u0301\u0327x" {
					assert.Contains(t, buf.Render(), "x")
				}
			})
		}
	}
}

// These are strict boundary requirements, not a diagnostic pass for data loss.
// A combining mark separated from its base by styling still needs an owner.
func TestStyledDetachedMarks(t *testing.T) {
	for _, mode := range []ansi.Method{ansi.WcWidth, ansi.GraphemeWidth} {
		for _, boundary := range []string{"\x1b[31m", "\x1b[0m", "\x1b]8;;https://example.test\x1b\\", "\x1b]8;;https://example.test\a", "\x1b]8;;\x1b\\"} {
			for _, suffix := range []string{"x", "", "\nq"} {
				for _, base := range []string{"e", "界"} {
					t.Run(fmt.Sprintf("%d/%q/%q/%q", mode, boundary, suffix, base), func(t *testing.T) {
						buf := uv.NewScreenBuffer(8, 2)
						buf.Method = mode
						uv.NewStyledString(base+boundary+"\u0301\u0327"+suffix).Draw(buf, buf.Bounds())
						width := mode.StringWidth(base)
						requireCell(t, buf, 0, 0, base+"\u0301\u0327", width)
						if width == 2 {
							requireCell(t, buf, 1, 0, "", 0)
						}
						if suffix == "x" {
							requireCell(t, buf, width, 0, "x", 1)
						} else if suffix == "\nq" {
							requireCell(t, buf, 0, 1, "q", 1)
						}
						// Styling after the base must not replace the base's attributes.
						assert.True(t, buf.CellAt(0, 0).Style.IsZero())
						assert.True(t, buf.CellAt(0, 0).Link.IsZero())
					})
				}
			}
		}
	}
}
