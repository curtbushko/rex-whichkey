package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	whichkey "temp"
)

// run has one public resource path and one private PTY child. The low-level menu
// command is UI/result-only; it never dispatches a resource operation.
func run(ctx context.Context, args []string, output, diagnostics io.Writer) error {
	if len(args) > 0 {
		switch args[0] {
		case "popup":
			return runPopupChild(ctx, args[1:])
		case "menu":
			return runMenuUI(ctx, args[1:], output, diagnostics)
		}
	}
	return runLaunch(ctx, args, output)
}

func runPopupChild(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet("rex-whichkey popup", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	directory := flags.String("directory", "", "private popup directory")
	native := flags.Bool("native", false, "parent lifetime pipe")
	ascii := flags.Bool("ascii", false, "ASCII decorations")
	noColor := flags.Bool("no-color", false, "disable colors")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 || !*native {
		return errors.New("whichkey: popup requires its private parent lifetime transport")
	}
	return whichkey.RunNativePopupTerminal(ctx, *directory, *noColor, *ascii)
}

func runMenuUI(ctx context.Context, args []string, output, diagnostics io.Writer) error {
	flags := flag.NewFlagSet("rex-whichkey menu", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	configPath := flags.String("config", "", "version 1 data configuration")
	ascii := flags.Bool("ascii", false, "ASCII decorations")
	noColor := flags.Bool("no-color", false, "disable colors")
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			_, writeErr := fmt.Fprintln(diagnostics, "usage: rex-whichkey menu [--config PATH] [--ascii] [--no-color] (UI/result only)")
			if writeErr != nil {
				return writeErr
			}
		}
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("whichkey: menu does not accept positional arguments")
	}
	var explicit *string
	flags.Visit(func(f *flag.Flag) {
		if f.Name == "config" {
			explicit = configPath
		}
	})
	config, err := whichkey.LoadMenuConfig(explicit)
	if err != nil {
		return err
	}
	event, err := whichkey.RunTerminal(ctx, config, *noColor, *ascii)
	if err != nil {
		return err
	}
	return whichkey.WriteMenuResult(output, event)
}

const publicHelp = `usage: rex-whichkey [options]

Open the Go-owned floating menu from the originating Rex terminal.
No Lua action registration or runtime bridge is used.

Options:
  --config PATH      Version 1 JSON menu data (otherwise environment/user default)
  --server ENDPOINT  Execution-host endpoint override
  --client FULL_ID   Exact graphical app intent (otherwise REX_CLIENT/pane owner)
  --ascii            Use ASCII decorations
  --no-color         Disable terminal colors
  --help, -h         Show help without effects

Requires full REX_SESSION/REX_BLOCK and a compatible execution-host Rex CLI.
Resource-only operations do not require a GUI. App actions require the retained
pane owner or explicit app, with a fresh connection and Remote Control permission.
Endpoint: --server > nonempty REX_SERVER > native execution-host default.
No focused/current-client fallback, Lua execution, or idle menu expiry.
Configuration is never edited or reloaded. See docs/mapping-format.md.
`

func runLaunch(ctx context.Context, args []string, output io.Writer) error {
	flags := flag.NewFlagSet("rex-whichkey", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	server := flags.String("server", "", "execution-host endpoint override")
	client := flags.String("client", "", "exact graphical app ID")
	config := flags.String("config", "", "version 1 data configuration")
	ascii := flags.Bool("ascii", false, "ASCII decorations")
	noColor := flags.Bool("no-color", false, "disable colors")
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			if _, writeErr := io.WriteString(output, publicHelp); writeErr != nil {
				return fmt.Errorf("write help: %w", writeErr)
			}
		}
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("whichkey: expected rex-whichkey [options]; use --help")
	}
	options := whichkey.LaunchOptions{ASCII: *ascii, NoColor: *noColor}
	flags.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "server":
			options.Server = server
		case "client":
			options.Client = client
		case "config":
			options.Config = config
		}
	})
	return whichkey.LaunchPopup(ctx, options)
}
func main() {
	if err := run(context.Background(), os.Args[1:], os.Stdout, os.Stderr); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return
		}
		if _, writeErr := fmt.Fprintf(os.Stderr, "rex-whichkey: %q\n", err.Error()); writeErr != nil {
			os.Exit(1)
		}
		os.Exit(1)
	}
}
