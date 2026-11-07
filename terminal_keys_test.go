//go:build darwin || linux

package whichkey

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRenameValidationDiagnostics(t *testing.T) {
	for _, test := range []struct{ input, problem string }{
		{"", "name cannot be empty"},
		{"\n", "name must be valid text without control characters"},
		{strings.Repeat("x", renameTextLimit+1), "name exceeds 4096 bytes"},
	} {
		err := validateRenameText(test.input)
		require.EqualError(t, err, test.problem)
	}
}

func TestTerminalKeyFraming(t *testing.T) {
	for _, test := range []struct {
		name, input   string
		entry         bool
		keys          []string
		paste         string
		focus, cancel bool
	}{
		{name: "unicode", input: "é界🙂", keys: []string{"é", "界", "🙂"}},
		{name: "alt unicode swallowed", input: "\x1béx", keys: []string{"x"}},
		{name: "invalid continuation reprocessed", input: "\xc2x", keys: []string{"x"}},
		{name: "reports swallowed", input: "\x1b[12;3R\x1bOPx", keys: []string{"x"}},
		{name: "legacy mouse unicode", input: "\x1b[M é界x", keys: []string{"x"}},
		{name: "string unicode ST continuation", input: "\x1bP\xc2\x9cx\x1b\\z", keys: []string{"z"}},
		{name: "paste is not keys", input: "\x1b[200~q\ny\x03\x1b[201~z", keys: []string{"z"}},
		{name: "entry paste", entry: true, input: "\x1b[200~é\n\x1b[201~", paste: "é\n"},
		{name: "entry edits", entry: true, input: "\x1b[D\x1b[3~\x1b[7~\x1b[8~", keys: []string{"\x1b[D", "\x1b[3~", "\x1b[H", "\x1b[F"}},
		{name: "focus", input: "\x1b[I", focus: true},
		{name: "parameterized focus ignored", input: "\x1b[1I"},
		{name: "cancel", input: "\x03", cancel: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			decoder := terminalKeys{textEntry: test.entry}
			var keys []string
			var canceled, focused bool
			var pasted string
			for _, b := range []byte(test.input) {
				key, cancel := decoder.feed(b)
				if key != "" {
					keys = append(keys, key)
				}
				canceled = canceled || cancel
				focused = focused || decoder.focusIn
				if decoder.pasteReady {
					pasted = string(decoder.pasteText)
				}
			}
			assert.Equal(t, test.keys, keys)
			assert.Equal(t, test.paste, pasted)
			assert.Equal(t, test.cancel, canceled)
			assert.Equal(t, test.focus, focused)
		})
	}
}

func TestTerminalPasteBoundAndEscape(t *testing.T) {
	decoder := terminalKeys{textEntry: true}
	for _, b := range []byte("\x1b[200~" + strings.Repeat("x", renameTextLimit+1) + "\x1b[201~") {
		decoder.feed(b)
	}
	assert.True(t, decoder.pasteReady)
	assert.True(t, decoder.pasteOverflow)
	assert.LessOrEqual(t, len(decoder.pasteText), renameTextLimit+len("\x1b[201~"))
	assert.Equal(t, terminalKeyGround, decoder.state)
	key, cancel := decoder.feed('\x1b')
	assert.Empty(t, key)
	assert.False(t, cancel)
	assert.Equal(t, terminalKeyEscape, decoder.state)
}

func TestCompletionUnicodeSurrogates(t *testing.T) {
	for _, input := range []string{`"\ud800"`, `"\udc00"`, `"\ud800\u0041"`, `"\ud800x"`, `null`} {
		_, err := completionString([]byte(input))
		require.Error(t, err, input)
	}
	for _, test := range []struct{ input, want string }{
		{`"\ud83d\ude42"`, "🙂"}, {`"\ufffd"`, "�"}, {`"\\ud800"`, `\ud800`}, {`"界"`, "界"},
	} {
		got, err := completionString([]byte(test.input))
		require.NoError(t, err)
		assert.Equal(t, test.want, got)
	}
}

func TestThemeColorChannels(t *testing.T) {
	for _, test := range []struct {
		input string
		want  themeRGB
	}{
		{"#abc", themeRGB{170, 187, 204}},
		{"#112233", themeRGB{17, 34, 51}},
		{"#ffff00008000", themeRGB{255, 0, 128}},
		{"rgb:f/00/8000", themeRGB{255, 0, 128}},
	} {
		got, ok := parseThemeColor(test.input)
		require.True(t, ok, test.input)
		assert.Equal(t, test.want, got)
	}
	for _, input := range []string{"", "#", "#1234", "#ggg", "rgb:1/2", "rgb:/0/f", "rgb:10000/0/0", "rgb:+f/0/0", "rgb:0_f/0/0"} {
		_, ok := parseThemeColor(input)
		require.False(t, ok, input)
	}
}

func TestPaletteFramingAndPartialReplies(t *testing.T) {
	var decoder PaletteDecoder
	var keys []byte
	var messages []TerminalPaletteMessage
	input := "界\x1b]11;rgb:0/0/0\x1b\\x\x1b]10;#fff\ay\x1b]0;title\xc2\x9cq\x1b\\z"
	for _, b := range []byte(input) {
		more, updates := decoder.Feed([]byte{b})
		keys = append(keys, more...)
		messages = append(messages, updates...)
	}
	assert.Equal(t, "界xyz", string(keys))
	require.Len(t, messages, 2)
	assert.Equal(t, "rgb:0/0/0", messages[1].Palette.Background)
	assert.Equal(t, "#fff", messages[1].Palette.Foreground)
	decoder.Feed([]byte("\x1b]10;" + strings.Repeat("x", maxPaletteReply+1)))
	assert.Empty(t, decoder.FlushEscape())
	keys, messages = decoder.Feed([]byte("\ax"))
	assert.Equal(t, []byte("x"), keys)
	assert.Empty(t, messages)
	decoder.Feed([]byte("\x1b"))
	assert.Equal(t, []byte{0x1b}, decoder.FlushEscape())
}
