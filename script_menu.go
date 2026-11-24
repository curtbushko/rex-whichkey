package whichkey

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"time"
	"unicode"

	"github.com/rivo/uniseg"
	"golang.org/x/sys/unix"
	"golang.org/x/term"

	"temp/rex"
)

// Selection authority stays inside directSelection until this script-only path
// completes. Every error path consumes it. Built-ins keep exact detach semantics.
func runDirectScript(ctx context.Context, client *rex.Client, origin launchOrigin, placement rex.Block, requested string, popup rex.Layer, mapping Mapping, options DirectOptions) (bool, error) {
	action, err := decodeScriptAction(mapping.Args)
	if err != nil {
		return false, err
	}
	if _, preflightErr := scriptPreflight(ctx, client, origin, placement, requested, popup, action); preflightErr != nil {
		return false, acknowledgeScriptError(ctx, options, preflightErr)
	}
	if minimizeErr := popup.Minimize(ctx); minimizeErr != nil {
		return false, minimizeErr
	}
	if focusErr := restoreDirectOrigin(ctx, client, origin, placement, requested); focusErr != nil {
		return false, focusErr
	}
	prepared, err := scriptPreflight(ctx, client, origin, placement, requested, popup, action)
	if err != nil {
		return false, restoreScriptError(ctx, client, origin, placement, requested, popup, options, err)
	}
	if hiddenErr := popup.VerifyMinimized(ctx); hiddenErr != nil {
		return false, hiddenErr
	}
	if focusErr := verifyDirectOriginFocus(ctx, client, origin, placement, requested); focusErr != nil {
		return false, focusErr
	}
	started, err := startScript(ctx, prepared)
	if started {
		// Never restore focus after exec. Even a cleanup failure must not steal it
		// back, reopen the menu, claim rollback or invite an automatic retry.
		cleanup, done := context.WithTimeout(context.WithoutCancel(ctx), 9*time.Second)
		defer done()
		if detachErr := popup.Detach(cleanup); detachErr != nil {
			return true, fmt.Errorf("script started; menu retirement failed or is uncertain; inspect manually, do not retry: %w", detachErr)
		}
		return true, err
	}
	if ctx.Err() != nil {
		return false, ctx.Err()
	}
	return false, restoreScriptError(ctx, client, origin, placement, requested, popup, options, err)
}

func scriptPreflight(ctx context.Context, client *rex.Client, origin launchOrigin, placement rex.Block, requested string, popup rex.Layer, action scriptAction) (preparedScript, error) {
	ctx, done := context.WithTimeout(ctx, 3*time.Second)
	defer done()
	if _, _, err := nativeResourceOrigin(ctx, client, origin, placement, requested); err != nil {
		return preparedScript{}, err
	}
	host, err := client.NativeScriptHost(ctx)
	if err != nil {
		return preparedScript{}, err
	}
	if hostErr := verifyScriptHost(ctx, host); hostErr != nil {
		return preparedScript{}, hostErr
	}
	own, err := popup.Block().Inspect(ctx)
	if err != nil || own.WindowID != popup.WindowID || own.LayerID != popup.ID {
		return preparedScript{}, errors.New("script not started: exact owned menu placement changed")
	}
	cwd, err := client.ScriptCWD(ctx, host, placement, own)
	if err != nil {
		return preparedScript{}, err
	}
	return prepareScript(action, cwd)
}

func restoreScriptError(ctx context.Context, client *rex.Client, origin launchOrigin, placement rex.Block, requested string, popup rex.Layer, options DirectOptions, problem error) error {
	// Certainty of nonexecution permits only this same retained error UI. Origin
	// and explicit app generation must still match before focusing the old layer.
	if err := verifyDirectOriginFocus(ctx, client, origin, placement, requested); err != nil {
		return errors.Join(problem, err)
	}
	if err := popup.RestoreForError(ctx); err != nil {
		return errors.Join(problem, err)
	}
	return acknowledgeScriptError(ctx, options, problem)
}

func acknowledgeScriptError(ctx context.Context, _ DirectOptions, problem error) error {
	if problem == nil {
		return errors.New("whichkey: script startup failed without diagnostic")
	}
	return errors.Join(problem, runScriptAcknowledgement(ctx, problem))
}

// Error acknowledgement owns only a temporary TTY. It has no menu, selected
// mapping, metadata worker, config loader or dispatch path to accidentally retry.
func runScriptAcknowledgement(ctx context.Context, problem error) (outcome error) {
	if err := ctx.Err(); err != nil {
		return err
	}
	fd, err := unix.Open("/dev/tty", unix.O_RDWR|unix.O_NOCTTY|unix.O_NONBLOCK|unix.O_CLOEXEC, 0)
	if err != nil {
		return fmt.Errorf("whichkey: open startup-error terminal: %w", err)
	}
	defer func() { outcome = errors.Join(outcome, unix.Close(fd)) }()
	original, err := term.GetState(fd)
	if err != nil {
		return err
	}
	defer func() { outcome = errors.Join(outcome, term.Restore(fd, original)) }()
	if _, rawErr := term.MakeRaw(fd); rawErr != nil {
		return rawErr
	}
	var display terminalDisplay
	defer func() { outcome = errors.Join(outcome, display.restore(ctx, fd)) }()
	if setupErr := display.setup(ctx, fd); setupErr != nil {
		return setupErr
	}
	display.viewportUsed = true
	width, height, err := term.GetSize(fd)
	if err != nil {
		return err
	}
	if drawErr := writeTerminal(ctx, fd, scriptErrorFrame(problem, width, height)); drawErr != nil {
		return drawErr
	}
	return waitScriptAcknowledgement(ctx, fd)
}

func scriptErrorFrame(problem error, width, height int) string {
	width, height = max(1, width), max(1, height)
	cause := strings.TrimPrefix(problem.Error(), "script not started: ")
	recovery := "Check origin host/account/cwd and trusted config."
	var diagnostic *scriptStartupProblem
	if errors.As(problem, &diagnostic) {
		cause, recovery = diagnostic.cause, diagnostic.recovery
	}
	var lines []string
	if height >= 4 {
		lines = append(lines, "Script not started")
	}
	// Reserve recovery and close instructions before allocating space to the
	// cause. Full paths remain in the returned error, not ahead of the repair.
	budget := height - 1 - len(lines)
	var repair []string
	if budget >= 2 {
		repair = scriptDiagnosticLines(recovery, width, max(1, budget/2))
	}
	lines = append(lines, scriptDiagnosticLines(cause, width, budget-len(repair))...)
	lines = append(lines, repair...)
	var output strings.Builder
	output.WriteString("\x1b[39;49m\x1b[2J")
	for row, text := range lines {
		fmt.Fprintf(&output, "\x1b[%d;1H%s", row+1, scriptClipLine(text, width))
	}
	fmt.Fprintf(&output, "\x1b[%d;1H%s", height, scriptClipLine("Esc/q: close; no retry", width))
	return output.String()
}

func scriptDiagnosticLines(message string, width, limit int) []string {
	if limit <= 0 {
		return nil
	}
	message = strings.Map(func(r rune) rune {
		if unicode.IsControl(r) {
			return ' '
		}
		return r
	}, message)
	var lines []string
	var line strings.Builder
	lineWidth := 0
	graphemes := uniseg.NewGraphemes(message)
	for graphemes.Next() {
		part := graphemes.Str()
		partWidth := graphemes.Width()
		if lineWidth+partWidth > max(1, width-2) {
			lines = append(lines, line.String())
			if len(lines) == limit {
				return lines
			}
			line.Reset()
			lineWidth = 0
		}
		line.WriteString(part)
		lineWidth += partWidth
	}
	return append(lines, line.String())
}

func scriptClipLine(text string, width int) string {
	var output strings.Builder
	used := 0
	graphemes := uniseg.NewGraphemes(text)
	for graphemes.Next() {
		if used+graphemes.Width() > width {
			break
		}
		output.WriteString(graphemes.Str())
		used += graphemes.Width()
	}
	return output.String()
}

func waitScriptAcknowledgement(ctx context.Context, fd int) error {
	poll := []unix.PollFd{{Fd: int32(fd), Events: unix.POLLIN}}
	var input terminalKeys
	last := time.Now()
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		if _, err := pollTerminal(poll, terminalPollInterval); err != nil {
			if errors.Is(err, unix.EINTR) {
				continue
			}
			return err
		}
		if poll[0].Revents&(unix.POLLHUP|unix.POLLERR|unix.POLLNVAL) != 0 {
			return errors.New("whichkey: startup error acknowledgement disconnected")
		}
		if input.state == terminalKeyEscape && time.Since(last) >= terminalEscapeDelay {
			return nil
		}
		if poll[0].Revents&unix.POLLIN == 0 {
			continue
		}
		var buffer [4096]byte
		n, err := unix.Read(fd, buffer[:])
		if err != nil {
			if terminalRetry(err) {
				continue
			}
			return err
		}
		if n == 0 {
			return errors.New("whichkey: startup error acknowledgement disconnected")
		}
		last = time.Now()
		acknowledged, inputErr := scriptAcknowledgementKeys(&input, buffer[:n])
		if inputErr != nil || acknowledged {
			return inputErr
		}
	}
}

func scriptAcknowledgementKeys(input *terminalKeys, data []byte) (bool, error) {
	for _, b := range data {
		key, cancel := input.feed(b)
		if cancel {
			return false, context.Canceled
		}
		if key == "q" || key == "\x1b" {
			return true, nil
		}
	}
	return false, nil
}
