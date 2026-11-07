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

// direct is the Lua-created, in-popup public entry. The no-argument controller
// remains transitional; menu is result-only and popup is its private PTY child.
func run(ctx context.Context, args []string, output, diagnostics io.Writer) error {
	if len(args) > 0 {
		switch args[0] {
		case "direct":
			return runDirect(ctx, args[1:], output)
		case "popup":
			return runPopupChild(ctx, args[1:])
		case "menu":
			return runMenuUI(ctx, args[1:], output, diagnostics)
		}
	}
	return runLaunch(ctx, args, output)
}

const directHelp = `usage: rex-whichkey direct --origin-session ID --origin-block ID --origin-window ID --origin-layer ID [options]

Run inside the sole terminal of a Lua-created float, not an ordinary pane.
Original IDs are fixed invocation data, separate from popup REX_* identities.
Options: --client FULL_ID --server ENDPOINT --config PATH --ascii --no-color
No nested popup, parent process, private IPC or init reload is used.
The shared public Lua launch entry is rex_whichkey_launch; see lua/rex_whichkey.lua.
`

func runDirect(ctx context.Context, args []string, output io.Writer) error {
	flags := flag.NewFlagSet("rex-whichkey direct", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	options := whichkey.DirectOptions{}
	flags.StringVar(&options.OriginSession, "origin-session", "", "original full session ID")
	flags.StringVar(&options.OriginBlock, "origin-block", "", "original full block ID")
	flags.StringVar(&options.OriginWindow, "origin-window", "", "original full window ID")
	flags.StringVar(&options.OriginLayer, "origin-layer", "", "original full layer ID")
	server := flags.String("server", "", "execution-host endpoint override")
	client := flags.String("client", "", "original exact app intent")
	config := flags.String("config", "", "version 1 JSON menu data")
	flags.BoolVar(&options.ASCII, "ascii", false, "ASCII decorations")
	flags.BoolVar(&options.NoColor, "no-color", false, "disable terminal colors")
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			if _, writeErr := io.WriteString(output, directHelp); writeErr != nil {
				return writeErr
			}
		}
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("whichkey: direct accepts no positional arguments")
	}
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
	return whichkey.RunDirectMenu(ctx, options)
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
This no-argument path retains the transitional Go controller.
The shared Lua launcher uses the public direct entry instead:
  rex do /absolute/checkout/lua/rex_whichkey.lua --action rex_whichkey_launch development=true
See the module's launch comments for installed and development arguments.

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
