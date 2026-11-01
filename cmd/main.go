package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	whichkey "temp"
)

// run composes the bounded controller, its PTY child and the standalone menu.
// Secret handoff/completion data travels on pipes, never command-line arguments.
// All action execution remains in the originating Rex configuration Lua VM.
func run(ctx context.Context, args []string, input io.Reader, output, diagnostics io.Writer) error {
	if len(args) == 0 {
		return errors.New("usage: rex-whichkey menu|prepare|controller|popup")
	}
	switch args[0] {
	case "prepare":
		flags := flag.NewFlagSet("rex-whichkey prepare", flag.ContinueOnError)
		flags.SetOutput(io.Discard)
		server := flags.String("server", "", "optional execution-host endpoint override")
		if err := flags.Parse(args[1:]); err != nil {
			return err
		}
		if flags.NArg() != 0 {
			return errors.New("whichkey: prepare accepts only an optional --server endpoint and private stdin configuration")
		}
		var override *string
		flags.Visit(func(option *flag.Flag) {
			if option.Name == "server" {
				override = server // Preserve the distinction between absent and empty.
			}
		})
		return whichkey.PreparePopupWithServer(ctx, input, output, override)
	case "controller":
		if len(args) != 1 {
			return errors.New("whichkey: controller accepts only its private stdin handoff")
		}
		return whichkey.RunPopupController(ctx, input)
	case "popup":
		flags := flag.NewFlagSet("rex-whichkey popup", flag.ContinueOnError)
		flags.SetOutput(io.Discard)
		directory := flags.String("directory", "", "invocation-owned private transport directory")
		deadline := flags.String("deadline", "", "original bounded invocation deadline")
		ascii := flags.Bool("ascii", false, "use ASCII decorations")
		noColor := flags.Bool("no-color", false, "disable terminal colors")
		if err := flags.Parse(args[1:]); err != nil {
			return err
		}
		if flags.NArg() != 0 {
			return errors.New("whichkey: popup does not accept positional arguments")
		}
		expires, err := time.Parse(time.RFC3339, *deadline)
		if err != nil {
			return errors.New("whichkey: popup requires its original deadline")
		}
		return whichkey.RunPopupTerminal(ctx, *directory, expires, *noColor, *ascii)
	case "menu":
	default:
		return errors.New("usage: rex-whichkey menu|prepare|controller|popup")
	}
	flags := flag.NewFlagSet("rex-whichkey menu", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	ascii := flags.Bool("ascii", false, "use ASCII decorations")
	noColor := flags.Bool("no-color", false, "disable terminal colors")
	if err := flags.Parse(args[1:]); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			if _, writeErr := fmt.Fprintln(diagnostics, "usage: rex-whichkey menu [--ascii] [--no-color] (resolved JSON configuration on stdin)"); writeErr != nil {
				return fmt.Errorf("write help: %w", writeErr)
			}
		}
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("whichkey: menu does not accept positional arguments")
	}
	config, err := whichkey.DecodeConfig(input)
	if err != nil {
		return err
	}
	// RunTerminal performs the real initial/navigation/resize/palette redraws,
	// consumes RenderMenu placements/Lines and passes frame.Pages to HandleKey.
	// It returns a selection only after all terminal cleanup has succeeded.
	event, err := whichkey.RunTerminal(ctx, config, *noColor, *ascii)
	if err != nil {
		return err
	}
	return whichkey.WriteMenuResult(output, event)
}

func main() {
	if err := run(context.Background(), os.Args[1:], os.Stdin, os.Stdout, os.Stderr); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return
		}
		// Decoder/I/O failures may contain external text. Quote it instead of
		// allowing a diagnostic to inject terminal escape sequences.
		if _, writeErr := fmt.Fprintf(os.Stderr, "rex-whichkey: %q\n", err.Error()); writeErr != nil {
			os.Exit(1)
		}
		os.Exit(1)
	}
}
