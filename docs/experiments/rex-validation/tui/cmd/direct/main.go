// Direct is a deliberately minimal terminal-I/O comparison, not a complete key parser.
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"example.com/whichkey-research/internal/probe"
	"golang.org/x/term"
)

func run() (result error) {
	fd := int(os.Stdin.Fd())
	original, err := term.MakeRaw(fd)
	if err != nil {
		return fmt.Errorf("raw mode: %w", err)
	}
	defer func() { result = errors.Join(result, term.Restore(fd, original)) }()
	if _, err := fmt.Print("\x1b[?1049h\x1b[?25l"); err != nil {
		return err
	}
	defer func() { _, err := fmt.Print("\x1b[?25h\x1b[?1049l"); result = errors.Join(result, err) }()
	m := probe.Menu{}
	updateSize := func() { m.Width, m.Height, _ = term.GetSize(fd) }
	updateSize()
	keys := make(chan byte)
	readErrors := make(chan error, 1)
	go func() {
		b := make([]byte, 1)
		for {
			if _, err := io.ReadFull(os.Stdin, b); err != nil {
				readErrors <- err
				return
			}
			keys <- b[0]
		}
	}()
	resized := make(chan os.Signal, 1)
	signal.Notify(resized, syscall.SIGWINCH)
	defer signal.Stop(resized)
	for {
		if _, err := fmt.Printf("\x1b[H\x1b[2J%s", m.View()); err != nil {
			return err
		}
		select {
		case b := <-keys:
			key := string(b)
			switch b {
			case 9:
				key = "tab"
			case 127, 8:
				key = "backspace"
			case 27:
				key = "esc"
			case 3:
				key = "ctrl+c"
			}
			m.Key(key)
		case <-resized:
			updateSize()
		case err := <-readErrors:
			return err
		}
		if m.Selection != "" {
			return probe.Finish(m)
		}
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
