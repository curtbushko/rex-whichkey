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

// Only an explicit direct invocation can retire an owned Lua-created float.
// Bare or ambiguous use never grants authority over the caller's terminal.
func run(ctx context.Context, args []string, output, diagnostics io.Writer) error {
	if len(args) == 0 {
		return errors.New("whichkey: choose direct or menu; use --help for the shared Lua launcher")
	}
	switch args[0] {
	case "direct":
		return runDirect(ctx, args[1:], output)
	case "menu":
		return runMenuUI(ctx, args[1:], output, diagnostics)
	case "help", "--help", "-h":
		if len(args) != 1 {
			return errors.New("whichkey: help accepts no additional arguments")
		}
		if _, err := io.WriteString(output, publicHelp); err != nil {
			return fmt.Errorf("write help: %w", err)
		}
		return flag.ErrHelp
	default:
		return errors.New("whichkey: expected direct or menu; use --help; no Go popup launcher is available")
	}
}

const directHelp = `usage: rex-whichkey direct --origin-session ID --origin-block ID --origin-window ID --origin-layer ID [options]

Run inside the sole terminal of a Lua-created float, not an ordinary pane.
Original IDs are fixed invocation data, separate from popup REX_* identities.
Options: --client FULL_ID --server ENDPOINT --config PATH --ascii --no-color
No nested popup, parent process, private IPC or init reload is used.
The shared public Lua action is rex_whichkey_launch; run lua/rex_whichkey.lua.
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

const publicHelp = `usage: rex-whichkey direct [options]
       rex-whichkey menu [--config PATH] [--ascii] [--no-color]

The shared Lua launcher creates the float and starts one direct Go menu:
  rex do /absolute/checkout/lua/rex_whichkey.lua --action rex_whichkey_launch development_directory=/absolute/checkout/bin/GOOS-GOARCH
Build the checkout first. See lua/rex_whichkey.lua for the exact source command.
Installed bindings use the same shared Lua module with ordinary PATH.

direct requires explicit original session/block/window/layer IDs and must run
inside the distinct sole terminal of a Lua-created float. Use direct --help.
menu is an in-terminal UI/result-only mode; it performs no selected action.
Bare invocation and legacy popup/private-parent flags are rejected without effects.
There is no Go float launcher, copied image, child renderer or private IPC.

Native chrome uses Rex defaults; Go matches only the terminal palette.
Visible float/TTY/workers/original focus are retired or restored before effects.
Rex deletes the detached host after Go exits; Go cannot verify its own deletion.
No focused/current-client fallback, configuration edits, reload or idle expiry.
See docs/mapping-format.md for validated version 1 JSON menu data.
`

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
