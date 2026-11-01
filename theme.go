package whichkey

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// TerminalPalette contains xterm OSC 10, 11 and 4 color values. Empty or
// malformed values are unavailable, not black. Only the base ANSI palette is
// needed for menu accents; no terminal palette is ever changed.
type TerminalPalette struct {
	Foreground string
	Background string
	Colors     [16]string
}

// TerminalPaletteQuery requests the default foreground, background and base ANSI
// palette using xterm OSC 10, 11 and 4. These are read-only queries, not palette
// assignments. The terminal owner may write this once after installing its single
// input reader and PaletteDecoder; it must not wait for replies before accepting
// keys or rendering with an empty TerminalPalette. Replies may arrive separately
// or late, and the decoder must remain in the input path after any query timeout.
// The caller may skip querying altogether when color output is disabled.
// See https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h3-Operating-System-Commands.
const TerminalPaletteQuery = "\x1b]10;?\x1b\\" + "\x1b]11;?\x1b\\" +
	"\x1b]4;0;?;1;?;2;?;3;?;4;?;5;?;6;?;7;?;8;?;9;?;10;?;11;?;12;?;13;?;14;?;15;?\x1b\\"

// ThemeRole describes meaning rather than a fixed terminal color. Callers must
// retain labels and key hints: color alone must never convey an action or error.
type ThemeRole uint8

const (
	ThemePrimary ThemeRole = iota
	ThemeKey
	ThemeGroup
	ThemeMuted
	ThemeCritical
)

type themeRGB struct{ r, g, b uint8 }

// Theme formats menu text without owning terminal modes or input. Its zero value
// uses terminal defaults and emits no color sequences.
type Theme struct {
	foreground *themeRGB
	background *themeRGB
	accents    [4]*themeRGB
	noColor    bool
	ascii      bool
}

// NewTheme resolves a terminal snapshot. A nonempty NO_COLOR overrides color
// output even when noColor is false. Unknown backgrounds retain terminal defaults
// instead of guessing a dark theme. ASCII affects decorations, not user labels.
// Rebuild from each TerminalPaletteMessage.Palette to apply partial or late
// replies while retaining the same noColor/ascii configuration.
func NewTheme(palette TerminalPalette, noColor, ascii bool) Theme {
	t := Theme{noColor: noColor || os.Getenv("NO_COLOR") != "", ascii: ascii}
	background, ok := parseThemeColor(palette.Background)
	if !ok || t.noColor {
		return t
	}
	t.background = &background
	foreground, ok := parseThemeColor(palette.Foreground)
	if !ok || themeContrast(foreground, background) < 4.5 {
		foreground = themeContrastFallback(background)
	}
	t.foreground = &foreground

	// Prefer the user's blue/cyan, magenta, neutral and red palette entries.
	// Even subdued help text must remain readable at normal text contrast.
	candidates := [4][]int{{4, 12, 6, 14}, {5, 13}, {8, 7, 15}, {1, 9}}
	for role, indexes := range candidates {
		t.accents[role] = t.foreground
		for _, index := range indexes {
			color, valid := parseThemeColor(palette.Colors[index])
			if valid && themeContrast(color, background) >= 4.5 {
				t.accents[role] = &color
				break
			}
		}
	}
	return t
}

// Style colors plain menu text and resets only the foreground/background that it
// sets. Control characters are removed before styling, including in NO_COLOR
// mode, so a label cannot inject SGR, OSC or cursor-control sequences.
func (t Theme) Style(role ThemeRole, text string) string {
	text = strings.Map(func(r rune) rune {
		if unicode.IsControl(r) {
			return -1
		}
		return r
	}, text)
	if t.noColor || t.foreground == nil || t.background == nil || text == "" {
		return text
	}
	foreground := t.foreground
	if role >= ThemeKey && role <= ThemeCritical {
		foreground = t.accents[role-ThemeKey]
	}
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm\x1b[48;2;%d;%d;%dm%s\x1b[39;49m",
		foreground.r, foreground.g, foreground.b,
		t.background.r, t.background.g, t.background.b, text)
}

// Decoration chooses an application-owned glyph or its ASCII equivalent. Pass
// an empty fallback for optional icons. User-provided descriptions should instead
// go directly to Style so ASCII mode does not destroy their text.
func (t Theme) Decoration(glyph, fallback string) string {
	if !t.ascii {
		return glyph
	}
	return strings.Map(func(r rune) rune {
		if r < 0x20 || r > 0x7e {
			return -1
		}
		return r
	}, fallback)
}

func parseThemeColor(value string) (themeRGB, bool) {
	var components []string
	switch {
	case strings.HasPrefix(value, "rgb:"):
		components = strings.Split(value[4:], "/")
	case strings.HasPrefix(value, "#"):
		digits := len(value) - 1
		if digits == 0 || digits%3 != 0 || digits > 12 {
			return themeRGB{}, false
		}
		width := digits / 3
		components = []string{value[1 : 1+width], value[1+width : 1+2*width], value[1+2*width:]}
	default:
		return themeRGB{}, false
	}
	if len(components) != 3 {
		return themeRGB{}, false
	}
	var rgb [3]uint8
	for i, component := range components {
		if len(component) < 1 || len(component) > 4 {
			return themeRGB{}, false
		}
		// Terminal color components contain only hexadecimal digits.
		for _, digit := range component {
			if !(digit >= '0' && digit <= '9' || digit >= 'a' && digit <= 'f' || digit >= 'A' && digit <= 'F') {
				return themeRGB{}, false
			}
		}
		value, err := strconv.ParseUint(component, 16, 16)
		if err != nil {
			return themeRGB{}, false
		}
		maximum := uint64(1)<<(4*len(component)) - 1
		rgb[i] = uint8((value*255 + maximum/2) / maximum)
	}
	return themeRGB{rgb[0], rgb[1], rgb[2]}, true
}

// Relative luminance uses the WCAG sRGB transfer function, not perceived
// brightness in gamma-encoded RGB. This also works for light terminal themes.
func themeLuminance(color themeRGB) float64 {
	linear := func(channel uint8) float64 {
		value := float64(channel) / 255
		if value <= 0.04045 {
			return value / 12.92
		}
		return math.Pow((value+0.055)/1.055, 2.4)
	}
	return 0.2126*linear(color.r) + 0.7152*linear(color.g) + 0.0722*linear(color.b)
}

func themeContrast(a, b themeRGB) float64 {
	first, second := themeLuminance(a), themeLuminance(b)
	if first < second {
		first, second = second, first
	}
	return (first + 0.05) / (second + 0.05)
}

func themeContrastFallback(background themeRGB) themeRGB {
	black, white := themeRGB{}, themeRGB{255, 255, 255}
	if themeContrast(black, background) >= themeContrast(white, background) {
		return black
	}
	return white
}

// TerminalPaletteMessage is terminal metadata, never a key or Escape/cancel.
// Palette is a complete snapshot, including earlier valid partial replies. The
// terminal owner passes it to NewTheme and uses the result for subsequent renders;
// text and application-owned glyphs use Style and Decoration respectively.
type TerminalPaletteMessage struct {
	Palette TerminalPalette
}

// PaletteDecoder separates OSC replies from keyboard input before Escape-key
// handling. Feed must receive bytes in order from the terminal's single reader.
// Once OSC starts it has no key timeout: delayed, malformed and oversized replies
// are consumed through their terminator, never replayed as mapping keys.
// The zero value is ready to use. It does not issue queries or alter terminal state.
type PaletteDecoder struct {
	palette       TerminalPalette
	state         uint8
	body          []byte
	discard       bool
	utf8Remaining int
}

const (
	paletteGround uint8 = iota
	paletteEscape
	paletteOSC
	paletteOSCEscape
	maxPaletteReply = 4096
)

// Feed returns ordinary input separately from parsed palette messages. Unknown
// OSC commands and invalid colors are ignored; they never enter the key stream.
func (d *PaletteDecoder) Feed(data []byte) (keys []byte, messages []TerminalPaletteMessage) {
	for _, b := range data {
		switch d.state {
		case paletteGround:
			if d.continuation(b) {
				keys = append(keys, b)
				continue
			}
			switch b {
			case 0x1b:
				d.state = paletteEscape
			case 0x9d:
				d.state = paletteOSC
			default:
				keys = append(keys, b)
			}
		case paletteEscape:
			if b == ']' {
				d.state = paletteOSC
			} else {
				keys = append(keys, 0x1b)
				d.state = paletteGround
				// Reprocess the byte as ordinary input, including a second ESC
				// or the start of a UTF-8 character.
				moreKeys, moreMessages := d.Feed([]byte{b})
				keys = append(keys, moreKeys...)
				messages = append(messages, moreMessages...)
			}
		case paletteOSC, paletteOSCEscape:
			// A C1-looking continuation byte in an OSC title is not ST.
			// Unknown OSC is consumed with the same framing as palette OSC.
			continuation := d.continuation(b)
			switch {
			case b == 0x18 || b == 0x1a:
				d.resetOSC()
			case b == 0x07 || b == 0x9c && !continuation || d.state == paletteOSCEscape && b == '\\':
				if !d.discard && d.applyReply() {
					messages = append(messages, TerminalPaletteMessage{Palette: d.palette})
				}
				d.resetOSC()
			case b == 0x1b:
				d.state = paletteOSCEscape
			default:
				if d.state == paletteOSCEscape || len(d.body) == maxPaletteReply {
					d.discard = true
				}
				d.state = paletteOSC
				if !d.discard {
					d.body = append(d.body, b)
				}
			}
		}
	}
	return keys, messages
}

// FlushEscape releases a lone ESC when the caller's normal key-ambiguity timer
// expires. It deliberately leaves an established OSC sequence untouched.
func (d *PaletteDecoder) FlushEscape() []byte {
	if d.state == paletteEscape {
		d.state = paletteGround
		return []byte{0x1b}
	}
	return nil
}

func (d *PaletteDecoder) resetOSC() {
	d.state = paletteGround
	d.body = d.body[:0]
	d.discard = false
	d.utf8Remaining = 0
}

// Track UTF-8 framing so valid multibyte user text cannot be mistaken for an
// eight-bit OSC introducer or terminator. Byte validity remains the key decoder's
// responsibility; no input bytes are rewritten here.
func (d *PaletteDecoder) continuation(b byte) bool {
	if d.utf8Remaining > 0 && b >= 0x80 && b <= 0xbf {
		d.utf8Remaining--
		return true
	}
	d.utf8Remaining = 0
	if b >= 0xc2 && b <= 0xdf {
		d.utf8Remaining = 1
	} else if b >= 0xe0 && b <= 0xef {
		d.utf8Remaining = 2
	} else if b >= 0xf0 && b <= 0xf4 {
		d.utf8Remaining = 3
	}
	return false
}

func (d *PaletteDecoder) applyReply() bool {
	fields := strings.Split(string(d.body), ";")
	if len(fields) == 2 && (fields[0] == "10" || fields[0] == "11") {
		if _, valid := parseThemeColor(fields[1]); !valid {
			return false
		}
		if fields[0] == "10" {
			d.palette.Foreground = fields[1]
		} else {
			d.palette.Background = fields[1]
		}
		return true
	}
	if len(fields) < 3 || fields[0] != "4" || len(fields)%2 != 1 {
		return false
	}
	changed := false
	for i := 1; i+1 < len(fields); i += 2 {
		index, err := strconv.Atoi(fields[i])
		if err != nil || index < 0 || index >= len(d.palette.Colors) {
			continue
		}
		if _, valid := parseThemeColor(fields[i+1]); valid {
			d.palette.Colors[index] = fields[i+1]
			changed = true
		}
	}
	return changed
}
