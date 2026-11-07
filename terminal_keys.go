//go:build darwin || linux

package whichkey

import (
	"unicode"
	"unicode/utf8"
)

// terminalKeys consumes non-OSC input after PaletteDecoder. Only complete literal
// Unicode tokens and the menu's three navigation keys can reach HandleKey.
// Terminal reports, function/mouse keys and Alt chords are not bindings. Text
// entry additionally accepts explicit editing keys and bounded paste as data,
// never replaying pasted control bytes as mappings or submission keys.
// Complete focus-in reports separately request a repaint; focus-out
// does nothing. State is bounded even for unterminated or oversized strings.
type terminalKeys struct {
	focusIn       bool // set only for the byte completing an unparameterized CSI I
	state         uint8
	sequence      [32]byte
	sequenceN     int
	runeBytes     [utf8.UTFMax]byte
	runeN         int
	altRune       bool
	stringUTF8    int
	pasteMatch    int
	mouseChars    int
	mouseUTF8     int
	textEntry     bool
	pasteForm     bool
	pasteText     []byte
	pasteReady    bool
	pasteOverflow bool
}

const (
	terminalKeyGround uint8 = iota
	terminalKeyEscape
	terminalKeyCSI
	terminalKeySS3
	terminalKeyString
	terminalKeyStringEscape
	terminalKeyIntermediate
	terminalKeyPaste
	terminalKeyMouse
)

func (d *terminalKeys) feed(b byte) (key string, cancel bool) {
	d.focusIn, d.pasteReady = false, false
	switch d.state {
	case terminalKeyMouse:
		return d.feedMouse(b)
	case terminalKeyPaste:
		return d.feedPaste(b)
	case terminalKeyString, terminalKeyStringEscape:
		return d.feedString(b)
	}
	if b == 0x03 || b == 0x1c || b == 0x1a {
		return "", true // raw Ctrl-C, Ctrl-\\, Ctrl-Z cancel without dispatch
	}
	switch d.state {
	case terminalKeyCSI, terminalKeySS3:
		return d.feedSequence(b)
	case terminalKeyIntermediate:
		if b < 0x20 || b > 0x2f {
			d.state = terminalKeyGround
		}
		return "", false
	case terminalKeyEscape:
		return d.feedEscape(b)
	}
	if d.runeN > 0 {
		return d.feedRune(b)
	}
	return d.feedGround(b)
}

func (d *terminalKeys) feedMouse(b byte) (string, bool) {
	// Legacy CSI M has three trailing coordinates, unlike SGR mouse
	// reports whose coordinates precede the final M/m. Consume UTF-8
	// coordinates too, so a queued report cannot choose a mapping.
	if d.mouseUTF8 > 0 && b >= 0x80 && b <= 0xbf {
		d.mouseUTF8--
		if d.mouseUTF8 > 0 {
			return "", false
		}
	} else {
		d.mouseUTF8 = 0
		switch {
		case b >= 0xc2 && b <= 0xdf:
			d.mouseUTF8 = 1
		case b >= 0xe0 && b <= 0xef:
			d.mouseUTF8 = 2
		case b >= 0xf0 && b <= 0xf4:
			d.mouseUTF8 = 3
		}
		if d.mouseUTF8 > 0 {
			return "", false
		}
	}
	d.mouseChars++
	if d.mouseChars == 3 {
		d.state = terminalKeyGround
	}
	return "", false
}

func (d *terminalKeys) feedPaste(b byte) (string, bool) {
	const end = "\x1b[201~"
	if d.pasteForm {
		if len(d.pasteText) < renameTextLimit+len(end) {
			d.pasteText = append(d.pasteText, b)
		} else {
			d.pasteOverflow = true
		}
	}
	if b != end[d.pasteMatch] {
		d.pasteMatch = 0
		if b == end[0] {
			d.pasteMatch = 1
		}
		return "", false
	}
	d.pasteMatch++
	if d.pasteMatch != len(end) {
		return "", false
	}
	d.state, d.pasteMatch = terminalKeyGround, 0
	if d.pasteForm {
		d.pasteReady = true
		if !d.pasteOverflow {
			d.pasteText = d.pasteText[:len(d.pasteText)-len(end)]
		}
	}
	return "", false
}

func (d *terminalKeys) feedString(b byte) (string, bool) {
	continuation := d.stringContinuation(b)
	if b == 0x9c && !continuation || b == 0x18 || b == 0x1a || d.state == terminalKeyStringEscape && b == '\\' {
		d.state = terminalKeyGround
	} else if b == 0x1b {
		d.state = terminalKeyStringEscape
	} else {
		d.state = terminalKeyString
	}
	return "", false
}

func (d *terminalKeys) feedSequence(b byte) (string, bool) {
	switch {
	case b == 0x1b:
		d.state = terminalKeyEscape
	case b == 0x18:
		d.state = terminalKeyGround
	case b >= 0x40 && b <= 0x7e:
		return d.finishSequence(b), false
	default:
		if d.sequenceN < len(d.sequence) {
			d.sequence[d.sequenceN] = b
			d.sequenceN++
		}
	}
	return "", false
}

func (d *terminalKeys) stringContinuation(b byte) bool {
	continuation := d.stringUTF8 > 0 && b >= 0x80 && b <= 0xbf
	if continuation {
		d.stringUTF8--
	} else {
		d.stringUTF8 = 0
		switch {
		case b >= 0xc2 && b <= 0xdf:
			d.stringUTF8 = 1
		case b >= 0xe0 && b <= 0xef:
			d.stringUTF8 = 2
		case b >= 0xf0 && b <= 0xf4:
			d.stringUTF8 = 3
		}
	}
	return continuation
}

func (d *terminalKeys) finishSequence(b byte) string {
	// xterm mode 1004 focus is unparameterized CSI I/O, never SS3 or key data.
	// https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h3-FocusIn_FocusOut
	csi := d.state == terminalKeyCSI
	d.focusIn = csi && b == 'I' && d.sequenceN == 0
	d.state = terminalKeyGround
	if csi && b == '~' && d.sequenceN == 3 && string(d.sequence[:3]) == "200" {
		d.state = terminalKeyPaste
		d.pasteForm, d.pasteOverflow = d.textEntry, false
		d.pasteText = d.pasteText[:0]
		return ""
	}
	if csi && b == 'M' && d.sequenceN == 0 {
		d.state, d.mouseChars, d.mouseUTF8 = terminalKeyMouse, 0, 0
		return ""
	}
	if d.textEntry {
		return d.editingKey(b)
	}
	return ""
}

func (d *terminalKeys) editingKey(b byte) string {
	sequence := string(d.sequence[:d.sequenceN])
	if sequence == "" {
		switch b {
		case 'C', 'D', 'H', 'F':
			return "\x1b[" + string(b)
		}
	}
	if b == '~' {
		switch sequence {
		case "1", "7":
			return "\x1b[H"
		case "4", "8":
			return "\x1b[F"
		case "3":
			return terminalDeleteKey
		}
	}
	return ""
}

func (d *terminalKeys) feedEscape(b byte) (string, bool) {
	d.state = terminalKeyGround
	switch b {
	case '[':
		d.state, d.sequenceN = terminalKeyCSI, 0
	case 'O':
		d.state, d.sequenceN = terminalKeySS3, 0
	case 'P', 'X', '^', '_', ']':
		d.state, d.stringUTF8 = terminalKeyString, 0
	case 0x1b:
		d.state = terminalKeyEscape
		return "\x1b", false
	default:
		if b >= 0x20 && b <= 0x2f {
			d.state = terminalKeyIntermediate
		} else if b >= 0xc2 && b <= 0xf4 {
			// Consume an entire Alt+Unicode chord without releasing its tail.
			d.runeN, d.altRune = 1, true
			d.runeBytes[0] = b
		}
	}
	return "", false
}

func (d *terminalKeys) feedRune(b byte) (string, bool) {
	if b < 0x80 || b > 0xbf {
		d.runeN, d.altRune = 0, false
		return d.feed(b)
	}
	d.runeBytes[d.runeN] = b
	d.runeN++
	if !utf8.FullRune(d.runeBytes[:d.runeN]) {
		return "", false
	}
	r, size := utf8.DecodeRune(d.runeBytes[:d.runeN])
	d.runeN = 0
	if d.altRune {
		d.altRune = false
		return "", false
	}
	if r == utf8.RuneError && size == 1 || unicode.IsControl(r) {
		return "", false
	}
	return string(r), false
}

func (d *terminalKeys) feedGround(b byte) (string, bool) {
	switch b {
	case 0x1b:
		d.state = terminalKeyEscape
	case 0x9b:
		d.state, d.sequenceN = terminalKeyCSI, 0
	case 0x90, 0x98, 0x9d, 0x9e, 0x9f:
		d.state, d.stringUTF8 = terminalKeyString, 0
	case '\t', '\b', 0x7f:
		return string(b), false
	case '\r', '\n', 0x15:
		if d.textEntry {
			return string(b), false
		}
	default:
		if b >= 0x20 && b < 0x7f {
			return string(b), false
		}
		if b >= 0xc2 && b <= 0xf4 {
			d.runeBytes[0], d.runeN = b, 1
		}
	}
	return "", false
}
