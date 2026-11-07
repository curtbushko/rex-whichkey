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

	"golang.org/x/sys/unix"
	"golang.org/x/term"
)

const (
	terminalPollInterval = 20 // milliseconds; also bounds cancellation latency during I/O
	terminalEscapeDelay  = 50 * time.Millisecond
	terminalSetup        = "\x1b[?7l\x1b[?25l\x1b[?1004h\x1b[?2004h"
)

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
	ui, err := newTerminalMenu(ctx, config, noColor, ascii)
	if err != nil {
		return MenuEvent{}, err
	}
	resize := make(chan os.Signal, 1)
	signal.Notify(resize, syscall.SIGWINCH, syscall.SIGCONT)
	defer signal.Stop(resize)
	// A separate file description keeps nonblocking flags off stdin/stdout.
	// No background reader may survive terminal restoration.
	fd, err := unix.Open("/dev/tty", unix.O_RDWR|unix.O_NOCTTY|unix.O_NONBLOCK|unix.O_CLOEXEC, 0)
	if err != nil {
		return MenuEvent{}, fmt.Errorf("whichkey: open terminal: %w", err)
	}
	defer func() {
		if closeErr := unix.Close(fd); closeErr != nil {
			outcome = errors.Join(outcome, fmt.Errorf("whichkey: close terminal: %w", closeErr))
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
		// Restore even if raw-mode setup or display cleanup fails.
		if restoreErr := term.Restore(fd, original); restoreErr != nil {
			outcome = errors.Join(outcome, fmt.Errorf("whichkey: restore terminal: %w", restoreErr))
		}
	}()
	if _, rawErr := term.MakeRaw(fd); rawErr != nil {
		return MenuEvent{}, fmt.Errorf("whichkey: set raw terminal: %w", rawErr)
	}
	var display terminalDisplay
	defer func() { outcome = errors.Join(outcome, display.restore(ctx, fd)) }()
	if setupErr := display.setup(ctx, fd); setupErr != nil {
		return MenuEvent{}, setupErr
	}
	display.viewportUsed = true // Failed redraw may already have cleared/styled cells.
	if drawErr := ui.draw(ctx, fd); drawErr != nil {
		return MenuEvent{}, drawErr
	}
	// No demand/job is eligible before the static frame has been written.
	localJobs := ui.startMetadata(ctx)
	defer localJobs.stop() // Join before display/termios/descriptor restoration.
	if !noColor && os.Getenv("NO_COLOR") == "" {
		// Read-only startup acquisition; keys/rendering never wait for replies.
		if queryErr := writeTerminal(ctx, fd, TerminalPaletteQuery); queryErr != nil {
			return MenuEvent{}, queryErr
		}
	}
	return ui.run(ctx, fd, resize, localJobs)
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
