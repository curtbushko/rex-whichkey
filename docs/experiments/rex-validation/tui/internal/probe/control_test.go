package probe_test

import (
	"fmt"
	"testing"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/stretchr/testify/assert"
)

func TestStyledMarkControls(t *testing.T) {
	for _, mode := range []ansi.Method{ansi.WcWidth, ansi.GraphemeWidth} {
		for _, control := range []string{"\a", "\t", "\b", "\x1b[2J", "\x1b[5C", "\x1b]0;unsafe title\a", "\x1b]52;c;eA==\x1b\\", "\x1b7", "\x00", "\x7f", "\x1bPignored\x1b\\"} {
			for _, suffix := range []string{"", "x"} {
				t.Run(fmt.Sprintf("%d/%q/%q", mode, control, suffix), func(t *testing.T) {
					buf := uv.NewScreenBuffer(8, 2)
					buf.Method = mode
					uv.NewStyledString("e"+control+"\u0301\u0327"+suffix).Draw(buf, buf.Bounds())
					// Unsupported terminal operations are not executable cell content.
					// Ignoring them must not lose the neighboring marks or move x.
					requireCell(t, buf, 0, 0, "e\u0301\u0327", 1)
					assert.NotContains(t, buf.Render(), control)
					lines := uv.NewStyledString(control + "\u0301e").Lines(mode)
					assert.Equal(t, "\u0301e", lines[0].String())
					if suffix != "" {
						requireCell(t, buf, 1, 0, "x", 1)
					}
				})
			}
		}
		for _, input := range []string{"\u0301\u0327\nx", "x\n\u0301\u0327", "e\n\u0301\u0327y"} {
			t.Run(fmt.Sprintf("%d/%q", mode, input), func(t *testing.T) {
				buf := uv.NewScreenBuffer(8, 2)
				buf.Method = mode
				uv.NewStyledString(input).Draw(buf, buf.Bounds())
				assert.Contains(t, buf.Render(), "\u0301\u0327")
				if input == "e\n\u0301\u0327y" {
					requireCell(t, buf, 0, 0, "e", 1)
					assert.Contains(t, buf.Line(1).String(), "\u0301\u0327y")
				} else if input == "\u0301\u0327\nx" {
					assert.Equal(t, "\u0301\u0327", buf.Line(0).String())
					requireCell(t, buf, 0, 1, "x", 1)
				} else {
					requireCell(t, buf, 0, 0, "x", 1)
					assert.Equal(t, "\u0301\u0327", buf.Line(1).String())
				}
				uv.NewStyledString("ab").Draw(buf, buf.Bounds())
				assert.NotContains(t, buf.Render(), "\u0301\u0327")
				requireCell(t, buf, 0, 0, "a", 1)
				requireCell(t, buf, 1, 0, "b", 1)
			})
		}
	}
}
