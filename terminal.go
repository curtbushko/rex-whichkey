//go:build darwin || linux

package whichkey

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"unicode"
	"unicode/utf8"

	"golang.org/x/sys/unix"
	"golang.org/x/term"

	"temp/rex"
)

const (
	terminalPollInterval = 20 // milliseconds; also bounds cancellation latency during I/O
	terminalEscapeDelay  = 50 * time.Millisecond
	terminalSetup        = "\x1b[?7l\x1b[?25l\x1b[?1004h\x1b[?2004h"
)

// A popup always supplies its selected CLI path and independent endpoint.
// Absence is reserved for the standalone menu's initial CLI capture.
type terminalRuntimeKey struct{}

type terminalRuntime struct {
	selection *rex.CLISelection
	endpoint  string
}

func resolveTerminalHostMenu(ctx context.Context, config Config) (Config, error) {
	var selection *rex.CLISelection
	if frozen, present := ctx.Value(terminalRuntimeKey{}).(terminalRuntime); present {
		endpoint, err := rex.CanonicalEndpoint(frozen.endpoint)
		if err != nil || endpoint != frozen.endpoint || frozen.selection == nil {
			return Config{}, errRuntimeCompatibility
		}
		// The parent expanded providers before sizing and image handoff.
		// Never let the child replace the retained executable snapshot.
		return config, nil
	}
	return resolveHostMenu(ctx, config, selection)
}

// RunTerminal owns one short-lived menu on the controlling popup terminal. The
// caller must give it exclusive terminal ownership, then dispatch its Selection
// only after a successful return. It does not close Rex resources or run actions.
// /dev/tty keeps terminal replies separate from configuration on stdin and the
// result on stdout. Linux and Darwin use the same ordered, nonblocking I/O loop.
//
// The popup owns the viewport: redraw clears it, then places RenderMenu's frame.
// Cursor/rendition and xterm wrap, visibility, focus-reporting and paste modes
// are saved/restored; termios is restored on selection, dismissal, cancellation,
// errors and caught HUP/INT/QUIT/TERM/PIPE/TSTP. TSTP cancels rather than suspending
// in raw mode.
// SIGKILL, SIGSTOP, process/terminal loss cannot run cleanup. Darwin may transiently
// expose PENDIN while restoring; restore the captured state unchanged, never mask
// other configured flags or claim a byte-for-byte platform comparison here.
func RunTerminal(ctx context.Context, config Config, noColor, ascii bool) (event MenuEvent, outcome error) {
	if err := ctx.Err(); err != nil {
		return MenuEvent{}, err
	}
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT,
		syscall.SIGTERM, syscall.SIGPIPE, syscall.SIGTSTP)
	defer stop()
	// Both the direct menu and popup child enter here on the execution/server
	// host. Resolve inventory once before menu construction or terminal mutation;
	// generated groups cannot add or replace a retained executable selection.
	config, err := resolveTerminalHostMenu(ctx, config)
	if err != nil {
		return MenuEvent{}, err
	}
	menu, err := NewMenu(config.Spec)
	if err != nil {
		return MenuEvent{}, err
	}
	menu.textSpec = config.textEntry
	transport, incremental := ctx.Value(terminalRenameKey{}).(terminalRenameTransport)
	if incremental {
		if transport.root == nil || config.textEntry == nil {
			return MenuEvent{}, errRuntimeStartup
		}
		menu.updates = &menuRenameUpdates{invocation: transport.invocation}
	}
	resize := make(chan os.Signal, 1)
	signal.Notify(resize, syscall.SIGWINCH, syscall.SIGCONT)
	defer signal.Stop(resize)

	// Open a separate file description: nonblocking flags must not change the
	// caller's stdin/stdout. No background reader can survive terminal restoration.
	fd, err := unix.Open("/dev/tty", unix.O_RDWR|unix.O_NOCTTY|unix.O_NONBLOCK|unix.O_CLOEXEC, 0)
	if err != nil {
		return MenuEvent{}, fmt.Errorf("whichkey: open terminal: %w", err)
	}
	defer func() {
		if err := unix.Close(fd); err != nil {
			outcome = errors.Join(outcome, fmt.Errorf("whichkey: close terminal: %w", err))
		}
		if outcome != nil {
			event = MenuEvent{}
		}
	}()
	original, err := term.GetState(fd)
	if err != nil {
		return MenuEvent{}, fmt.Errorf("whichkey: capture terminal: %w", err)
	}
	defer func() {
		// Restore termios even if raw-mode setup or display cleanup fails.
		if err := term.Restore(fd, original); err != nil {
			outcome = errors.Join(outcome, fmt.Errorf("whichkey: restore terminal: %w", err))
		}
	}()
	if _, err := term.MakeRaw(fd); err != nil {
		return MenuEvent{}, fmt.Errorf("whichkey: set raw terminal: %w", err)
	}

	cursorSaved, modesSaved, viewportUsed := false, false, false
	defer func() {
		var cleanup strings.Builder
		// CAN aborts an incomplete save, mutation, redraw or palette query.
		// Never restore a previous invocation's save slot after a partial save.
		cleanup.WriteByte('\x18')
		if viewportUsed {
			cleanup.WriteString("\x1b[39;49m\x1b[2J")
		}
		if cursorSaved {
			cleanup.WriteString("\x1b8")
		}
		if modesSaved {
			cleanup.WriteString("\x1b[?7;25;1004;2004r")
		}
		// Cleanup outlives canceled input, but a lost/non-reading terminal
		// must not hold the process forever. The earlier termios defer runs
		// afterwards regardless of this bounded, best-effort display write.
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
		err := writeTerminal(cleanupCtx, fd, cleanup.String())
		cancel()
		if err != nil {
			outcome = errors.Join(outcome, fmt.Errorf("whichkey: restore display: %w", err))
		}
	}()
	// Complete each save in its own write before arming its rollback. No
	// display mutations may start until both save sequences have completed.
	if err := writeTerminal(ctx, fd, "\x1b7"); err != nil {
		return MenuEvent{}, err
	}
	cursorSaved = true
	if err := writeTerminal(ctx, fd, "\x1b[?7;25;1004;2004s"); err != nil {
		return MenuEvent{}, err
	}
	modesSaved = true
	if err := writeTerminal(ctx, fd, terminalSetup); err != nil {
		return MenuEvent{}, err
	}

	var palette PaletteDecoder
	var input terminalKeys
	theme := NewTheme(TerminalPalette{}, noColor, ascii)
	draw := func() (MenuFrame, error) {
		return drawTerminal(ctx, fd, menu, config, theme)
	}
	viewportUsed = true // A failed redraw may already have cleared or styled cells.
	frame, err := draw()
	if err != nil {
		return MenuEvent{}, err
	}
	if !noColor && os.Getenv("NO_COLOR") == "" {
		// Startup acquisition, not a blocking probe: keys/rendering never wait
		// for a reply, and palette decoding stays installed until exit.
		if err := writeTerminal(ctx, fd, TerminalPaletteQuery); err != nil {
			return MenuEvent{}, err
		}
	}

	handleKeys := func(data []byte) (MenuEvent, error) {
		for _, b := range data {
			if err := ctx.Err(); err != nil {
				return MenuEvent{}, err
			}
			input.textEntry = menu.entry != nil
			key, cancel := input.feed(b)
			if cancel {
				return MenuEvent{}, context.Canceled
			}
			if input.pasteReady && menu.entry != nil {
				if input.pasteOverflow {
					menu.entry.touched = true
					menu.entry.problem = "Name exceeds 4096 bytes"
				} else {
					menu.insertText(string(input.pasteText))
				}
			} else if key == "" && !input.focusIn {
				continue
			}
			if key != "" {
				event := menu.HandleKey(key, frame.Pages)
				if menu.Done() {
					return event, nil
				}
			}
			// Focus-in repaints the retained menu without HandleKey: even a
			// pending confirmation must survive returning to this terminal.
			// Refresh Pages before handling another key in this same read.
			var err error
			frame, err = draw()
			if err != nil {
				return MenuEvent{}, err
			}
		}
		return MenuEvent{}, nil
	}
	refreshRename := func() error {
		if !incremental {
			return nil
		}
		if menu.entry != nil && (input.state != terminalKeyGround || input.runeN > 0 || palette.state == paletteEscape) {
			menu.entry.touched = true // Preserve partial UTF-8/escape/paste intent.
		}
		changed, updateErr := menu.pollRenameUpdates(transport)
		if updateErr != nil {
			return updateErr
		}
		if changed {
			frame, updateErr = draw()
		}
		return updateErr
	}
	poll := []unix.PollFd{{Fd: int32(fd), Events: unix.POLLIN}}
	var buffer [4096]byte
	lastInput := time.Now()
	for {
		if err := ctx.Err(); err != nil {
			return MenuEvent{}, err
		}
		if updateErr := refreshRename(); updateErr != nil {
			return MenuEvent{}, updateErr
		}
		select {
		case <-resize:
			frame, err = draw()
			if err != nil {
				return MenuEvent{}, err
			}
		default:
		}
		if _, err := pollTerminal(poll, terminalPollInterval); err != nil {
			if errors.Is(err, unix.EINTR) {
				continue
			}
			return MenuEvent{}, fmt.Errorf("whichkey: poll terminal input: %w", err)
		}
		if poll[0].Revents&(unix.POLLERR|unix.POLLHUP|unix.POLLNVAL) != 0 {
			return MenuEvent{}, errors.New("whichkey: terminal input disconnected")
		}
		if poll[0].Revents&unix.POLLIN == 0 {
			if time.Since(lastInput) >= terminalEscapeDelay {
				// Drain queued bytes before resolving Escape, including bytes
				// that arrived during redraw. Established OSC/CSI/paste never
				// time out into menu/confirmation keys.
				event, err := handleKeys(palette.FlushEscape())
				if err != nil || menu.Done() {
					return event, err
				}
				if input.state == terminalKeyEscape {
					input.state = terminalKeyGround
					event = menu.HandleKey("\x1b", frame.Pages)
					if menu.Done() {
						return event, nil
					}
					frame, err = draw()
					if err != nil {
						return MenuEvent{}, err
					}
				}
			}
			continue
		}
		n, err := unix.Read(fd, buffer[:])
		if err != nil {
			if terminalRetry(err) {
				continue
			}
			return MenuEvent{}, fmt.Errorf("whichkey: read terminal: %w", err)
		}
		if n == 0 {
			return MenuEvent{}, io.EOF
		}
		lastInput = time.Now()
		for i := 0; i < n; i++ {
			// Feeding in byte order preserves metadata/key ordering even when
			// a single read contains keys on both sides of a palette reply.
			keys, messages := palette.Feed(buffer[i : i+1])
			for _, message := range messages {
				theme = NewTheme(message.Palette, noColor, ascii)
				frame, err = draw()
				if err != nil {
					return MenuEvent{}, err
				}
			}
			event, err := handleKeys(keys)
			if err != nil || menu.Done() {
				return event, err
			}
		}
	}
}

func drawTerminal(ctx context.Context, fd int, menu *Menu, config Config, theme Theme) (MenuFrame, error) {
	width, height, err := term.GetSize(fd)
	if err != nil {
		return MenuFrame{}, fmt.Errorf("whichkey: terminal size: %w", err)
	}
	frame := RenderMenu(menu, config, theme, width, height)
	var output strings.Builder
	// The entire dedicated viewport is cleared so old geometry cannot survive
	// a smaller/relocated frame. Disabled autowrap prevents resize-time scrolling.
	output.WriteString("\x1b[39;49m\x1b[2J")
	for row, line := range frame.Lines {
		if row >= frame.Height || frame.Width <= 0 {
			break
		}
		fmt.Fprintf(&output, "\x1b[%d;%dH%s", frame.Y+row+1, frame.X+1, line)
	}
	if err := writeTerminal(ctx, fd, output.String()); err != nil {
		return MenuFrame{}, err
	}
	return frame, nil
}

func writeTerminal(ctx context.Context, fd int, text string) error {
	data := []byte(text)
	poll := []unix.PollFd{{Fd: int32(fd), Events: unix.POLLOUT}}
	for len(data) > 0 {
		if err := ctx.Err(); err != nil {
			return err
		}
		if _, err := pollTerminal(poll, terminalPollInterval); err != nil {
			if errors.Is(err, unix.EINTR) {
				continue
			}
			return fmt.Errorf("whichkey: poll terminal output: %w", err)
		}
		if poll[0].Revents&(unix.POLLERR|unix.POLLHUP|unix.POLLNVAL) != 0 {
			return errors.New("whichkey: terminal output disconnected")
		}
		if poll[0].Revents&unix.POLLOUT == 0 {
			continue
		}
		n, err := unix.Write(fd, data)
		if err != nil {
			if terminalRetry(err) {
				continue
			}
			return fmt.Errorf("whichkey: write terminal: %w", err)
		}
		if n == 0 {
			return io.ErrShortWrite
		}
		data = data[n:]
	}
	return nil
}

func terminalRetry(err error) bool {
	return errors.Is(err, unix.EINTR) || errors.Is(err, unix.EAGAIN) || errors.Is(err, unix.EWOULDBLOCK)
}

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
	if d.state == terminalKeyMouse {
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
	if d.state == terminalKeyPaste {
		const end = "\x1b[201~"
		if d.pasteForm {
			if len(d.pasteText) < renameTextLimit+len(end) {
				d.pasteText = append(d.pasteText, b)
			} else {
				d.pasteOverflow = true
			}
		}
		if b == end[d.pasteMatch] {
			d.pasteMatch++
			if d.pasteMatch == len(end) {
				d.state, d.pasteMatch = terminalKeyGround, 0
				if d.pasteForm {
					d.pasteReady = true
					if !d.pasteOverflow {
						d.pasteText = d.pasteText[:len(d.pasteText)-len(end)]
					}
				}
			}
		} else {
			d.pasteMatch = 0
			if b == end[0] {
				d.pasteMatch = 1
			}
		}
		return "", false
	}
	if d.state == terminalKeyString || d.state == terminalKeyStringEscape {
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
		if b == 0x9c && !continuation || b == 0x18 || b == 0x1a || d.state == terminalKeyStringEscape && b == '\\' {
			d.state = terminalKeyGround
		} else if b == 0x1b {
			d.state = terminalKeyStringEscape
		} else {
			d.state = terminalKeyString
		}
		return "", false
	}
	if b == 0x03 || b == 0x1c || b == 0x1a {
		return "", true // raw Ctrl-C, Ctrl-\\, Ctrl-Z cancel without dispatch
	}
	if d.state == terminalKeyCSI || d.state == terminalKeySS3 {
		switch {
		case b == 0x1b:
			d.state = terminalKeyEscape
		case b == 0x18:
			d.state = terminalKeyGround
		case b >= 0x40 && b <= 0x7e:
			// xterm mode 1004 reports CSI I/O, with no parameters. Do not
			// release any part as keys or treat SS3/parameterized I as focus.
			// https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h3-FocusIn_FocusOut
			d.focusIn = d.state == terminalKeyCSI && b == 'I' && d.sequenceN == 0
			paste := d.state == terminalKeyCSI && b == '~' && d.sequenceN == 3 && string(d.sequence[:3]) == "200"
			mouse := d.state == terminalKeyCSI && b == 'M' && d.sequenceN == 0
			d.state = terminalKeyGround
			if paste {
				d.state = terminalKeyPaste
				d.pasteForm, d.pasteOverflow = d.textEntry, false
				d.pasteText = d.pasteText[:0]
			} else if mouse {
				d.state, d.mouseChars, d.mouseUTF8 = terminalKeyMouse, 0, 0
			} else if d.textEntry {
				sequence := string(d.sequence[:d.sequenceN])
				if sequence == "" && (b == 'C' || b == 'D' || b == 'H' || b == 'F') {
					return "\x1b[" + string(b), false
				}
				if b == '~' {
					switch sequence {
					case "1", "7":
						return "\x1b[H", false
					case "4", "8":
						return "\x1b[F", false
					case "3":
						return "\x1b[3~", false
					}
				}
			}
		default:
			if d.sequenceN < len(d.sequence) {
				d.sequence[d.sequenceN] = b
				d.sequenceN++
			}
		}
		return "", false
	}
	if d.state == terminalKeyIntermediate {
		if b < 0x20 || b > 0x2f {
			d.state = terminalKeyGround
		}
		return "", false
	}
	if d.state == terminalKeyEscape {
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
	if d.runeN > 0 {
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
