//go:build darwin || linux

package whichkey

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"golang.org/x/sys/unix"
)

// One UI goroutine owns parsing, retained menu state and redraw ordering.
type terminalMenu struct {
	config         Config
	menu           *Menu
	direct         *directMenuState
	palette        PaletteDecoder
	input          terminalKeys
	theme          Theme
	frame          MenuFrame
	noColor, ascii bool
}

func newTerminalMenu(ctx context.Context, config Config, noColor, ascii bool) (*terminalMenu, error) {
	// Local validation/placeholder construction performs no provider I/O.
	base, err := canonicalMenuSnapshot(config)
	if err != nil {
		return nil, err
	}
	config, err = providerConfig(base, [4]*providerUpdate{})
	if err != nil {
		return nil, err
	}
	menu, err := NewMenu(config.Spec)
	if err != nil {
		return nil, err
	}
	menu.textSpec = config.textEntry
	menu.providers = &menuProviderUpdates{base: base}
	direct, inProcess := ctx.Value(terminalDirectKey{}).(*directMenuState)
	if inProcess {
		if direct == nil || direct.client == nil || direct.request.snapshot == nil || config.textEntry == nil {
			return nil, errors.New("whichkey: invalid in-process menu context")
		}
		menu.updates = &menuRenameUpdates{}
		direct.config = config
	}
	return &terminalMenu{config: config, menu: menu, direct: direct,
		theme: NewTheme(TerminalPalette{}, noColor, ascii), noColor: noColor, ascii: ascii}, nil
}

func (ui *terminalMenu) draw(ctx context.Context, fd int) error {
	frame, err := drawTerminal(ctx, fd, ui.menu, ui.config, ui.theme)
	if err == nil {
		ui.frame = frame
	}
	return err
}

func (ui *terminalMenu) handleKeys(ctx context.Context, fd int, data []byte) (MenuEvent, error) {
	for _, b := range data {
		if err := ctx.Err(); err != nil {
			return MenuEvent{}, err
		}
		ui.input.textEntry = ui.menu.entry != nil
		key, cancel := ui.input.feed(b)
		if cancel {
			return MenuEvent{}, context.Canceled
		}
		if ui.input.pasteReady && ui.menu.entry != nil {
			ui.insertPaste()
		} else if key == "" && !ui.input.focusIn {
			continue
		}
		if key != "" {
			event := ui.menu.HandleKey(key, ui.frame.Pages)
			if ui.menu.Done() {
				return event, nil
			}
		}
		// Focus-in repaints without HandleKey so pending confirmation survives.
		// Refresh Pages before another key in this same read.
		if err := ui.draw(ctx, fd); err != nil {
			return MenuEvent{}, err
		}
	}
	return MenuEvent{}, nil
}

func (ui *terminalMenu) insertPaste() {
	if ui.input.pasteOverflow {
		ui.menu.entry.touched = true
		ui.menu.entry.problem = "Name exceeds 4096 bytes"
		return
	}
	ui.menu.insertText(string(ui.input.pasteText))
}

func (ui *terminalMenu) refreshUpdates(ctx context.Context, fd int, jobs *metadataJobs) error {
	if ui.menu.entry != nil && (ui.input.state != terminalKeyGround || ui.input.runeN > 0 || ui.palette.state == paletteEscape) {
		ui.menu.entry.touched = true // Preserve partial UTF-8/escape/paste intent.
	}
	var changed bool
	var err error
	if ui.direct != nil {
		changed, err = ui.direct.poll(ui.menu, jobs)
	} else {
		changed, err = ui.menu.pollLocalProviders(jobs)
	}
	if err != nil {
		return err
	}
	if changed {
		return ui.draw(ctx, fd)
	}
	return nil
}

func (ui *terminalMenu) startMetadata(ctx context.Context) *metadataJobs {
	ui.menu.demandProviders(true)
	if ui.direct != nil {
		return newMetadataJobs(ctx, ui.direct.request, ui.direct.client)
	}
	// Standalone is result-only; it cannot fabricate native origin/authority.
	return newMetadataJobs(ctx, popupRequest{}, nil)
}

func (ui *terminalMenu) refresh(ctx context.Context, fd int, resize <-chan os.Signal, jobs *metadataJobs) error {
	if err := ui.refreshUpdates(ctx, fd, jobs); err != nil {
		return err
	}
	select {
	case <-resize:
		return ui.draw(ctx, fd)
	default:
		return nil
	}
}

func (ui *terminalMenu) run(ctx context.Context, fd int, resize <-chan os.Signal, jobs *metadataJobs) (MenuEvent, error) {
	poll := []unix.PollFd{{Fd: int32(fd), Events: unix.POLLIN}}
	lastInput := time.Now()
	for {
		if err := ctx.Err(); err != nil {
			return MenuEvent{}, err
		}
		if err := ui.refresh(ctx, fd, resize, jobs); err != nil {
			return MenuEvent{}, err
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
			if time.Since(lastInput) < terminalEscapeDelay {
				continue
			}
			event, err := ui.flushEscape(ctx, fd)
			if err != nil || ui.menu.Done() {
				return event, err
			}
			continue
		}
		event, receivedAt, err := ui.read(ctx, fd)
		if err != nil || ui.menu.Done() {
			return event, err
		}
		if !receivedAt.IsZero() {
			lastInput = receivedAt
		}
	}
}

func (ui *terminalMenu) flushEscape(ctx context.Context, fd int) (MenuEvent, error) {
	// Drain queued bytes before resolving Escape, including bytes arriving during
	// redraw. Established OSC/CSI/paste never time out into confirmation keys.
	event, err := ui.handleKeys(ctx, fd, ui.palette.FlushEscape())
	if err != nil || ui.menu.Done() {
		return event, err
	}
	if ui.input.state != terminalKeyEscape {
		return MenuEvent{}, nil
	}
	ui.input.state = terminalKeyGround
	event = ui.menu.HandleKey("\x1b", ui.frame.Pages)
	if ui.menu.Done() {
		return event, nil
	}
	return MenuEvent{}, ui.draw(ctx, fd)
}

func (ui *terminalMenu) read(ctx context.Context, fd int) (MenuEvent, time.Time, error) {
	var buffer [4096]byte
	n, err := unix.Read(fd, buffer[:])
	if err != nil {
		if terminalRetry(err) {
			return MenuEvent{}, time.Time{}, nil
		}
		return MenuEvent{}, time.Time{}, fmt.Errorf("whichkey: read terminal: %w", err)
	}
	if n == 0 {
		return MenuEvent{}, time.Time{}, io.EOF
	}
	// Escape ambiguity is measured from receipt, not completion of redraws.
	receivedAt := time.Now()
	event, err := ui.consume(ctx, fd, buffer[:n])
	return event, receivedAt, err
}

func (ui *terminalMenu) consume(ctx context.Context, fd int, data []byte) (MenuEvent, error) {
	for i := range data {
		// Byte order preserves keys on both sides of palette replies in one read.
		keys, messages := ui.palette.Feed(data[i : i+1])
		for _, message := range messages {
			ui.theme = NewTheme(message.Palette, ui.noColor, ui.ascii)
			if err := ui.draw(ctx, fd); err != nil {
				return MenuEvent{}, err
			}
		}
		event, err := ui.handleKeys(ctx, fd, keys)
		if err != nil || ui.menu.Done() {
			return event, err
		}
	}
	return MenuEvent{}, nil
}
