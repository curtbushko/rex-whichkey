//go:build darwin || linux

package whichkey

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// Rollback is armed only after each save completes. Never restore a previous
// invocation's save slot after a partial save.
type terminalDisplay struct {
	cursorSaved, modesSaved, viewportUsed bool
}

func (display *terminalDisplay) setup(ctx context.Context, fd int) error {
	if err := writeTerminal(ctx, fd, "\x1b7"); err != nil {
		return err
	}
	display.cursorSaved = true
	if err := writeTerminal(ctx, fd, "\x1b[?7;25;1004;2004s"); err != nil {
		return err
	}
	display.modesSaved = true
	return writeTerminal(ctx, fd, terminalSetup)
}

func (display terminalDisplay) restore(ctx context.Context, fd int) error {
	var cleanup strings.Builder
	// CAN aborts incomplete saves, mutations, redraws or palette queries.
	cleanup.WriteByte('\x18')
	if display.viewportUsed {
		cleanup.WriteString("\x1b[39;49m\x1b[2J")
	}
	if display.cursorSaved {
		cleanup.WriteString("\x1b8")
	}
	if display.modesSaved {
		cleanup.WriteString("\x1b[?7;25;1004;2004r")
	}
	// Preserve caller values, but allow bounded cleanup after input cancellation.
	// Termios restoration still runs if the terminal is lost/non-reading.
	cleanupCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 250*time.Millisecond)
	defer cancel()
	if err := writeTerminal(cleanupCtx, fd, cleanup.String()); err != nil {
		return fmt.Errorf("whichkey: restore display: %w", err)
	}
	return nil
}
