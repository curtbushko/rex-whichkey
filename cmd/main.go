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

// run composes the configured controller, native foreground child and standalone menu.
// Secret handoff/completion data travels on pipes, never command-line arguments.
// Configured callbacks stay in the original Rex VM; native defaults use exact
// captured Rex activation after owned-popup and private-file cleanup.
func run(ctx context.Context, args []string, input io.Reader, output, diagnostics io.Writer) error {
	if len(args) == 0 {
		return runLaunch(ctx, args, output)
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
		native := flags.Bool("native", false, "require a foreground parent's private lifetime pipe instead of a deadline")
		ascii := flags.Bool("ascii", false, "use ASCII decorations")
		noColor := flags.Bool("no-color", false, "disable terminal colors")
		if err := flags.Parse(args[1:]); err != nil {
			return err
		}
		if flags.NArg() != 0 {
			return errors.New("whichkey: popup does not accept positional arguments")
		}
		nativeSet, deadlineSet := false, false
		flags.Visit(func(option *flag.Flag) {
			switch option.Name {
			case "native":
				nativeSet = true
			case "deadline":
				deadlineSet = true
			}
		})
		if nativeSet {
			if !*native || deadlineSet {
				return errors.New("whichkey: popup requires either its native lifetime pipe or its original deadline, not both")
			}
			return whichkey.RunNativePopupTerminal(ctx, *directory, *noColor, *ascii)
		}
		expires, err := time.Parse(time.RFC3339, *deadline)
		if err != nil {
			return errors.New("whichkey: popup requires its original deadline")
		}
		return whichkey.RunPopupTerminal(ctx, *directory, expires, *noColor, *ascii)
	case "menu":
	default:
		return runLaunch(ctx, args, output)
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

const publicHelp = `usage: rex-whichkey [options]

Open the basic floating menu from this Rex terminal, without custom registration.
A registered default whichkey action instead uses your original-VM Lua menu.

Options:
  --server ENDPOINT  Execution-host socket path, unix URI or host:port override
  --client FULL_ID   Exact attached client override (otherwise REX_CLIENT or the
                     launching pane's GUI owner for all app actions)
  --action NAME      Require this registered Lua launch action (no native fallback)
  --help, -h         Show help without contacting Rex

Requires full REX_SESSION/REX_BLOCK from the originating Rex terminal, matching
execution/server-host helper and Rex CLI. App UI actions additionally require an
attached graphical app and its Remote Control permission. Native pane/tab/session
closes and built-in tab/session rename use captured resource IDs without an app
chooser. Built-in Switch lists sessions in our popup; rename uses text entry. Explicit client selectors remain strict. Lua customization additionally needs installed modules,
setup and launch/completion registration. Nothing edits/reloads configuration.
The native caller waits until selection/dismissal, cancellation or origin loss;
there is no native idle timeout. Configured launch returns after activation and
retains its authenticated 30-second expiry. Failed/uncertain activation is never
retried or switched to native.

Endpoint: --server > nonempty REX_SERVER > Rex native Linux/macOS default.
Missing, stale or ambiguous origin context fails; no focused/current-client or
alternate-server fallback. Activation is not proof that the popup was displayed.
See docs/launch-contract.md for setup and routing details.
`

func runLaunch(ctx context.Context, args []string, output io.Writer) error {
	flags := flag.NewFlagSet("rex-whichkey", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	server := flags.String("server", "", "execution-host endpoint override")
	client := flags.String("client", "", "exact attached client ID")
	action := flags.String("action", "whichkey", "registered configuration launch action")
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			if _, writeErr := io.WriteString(output, publicHelp); writeErr != nil {
				return fmt.Errorf("write help: %w", writeErr)
			}
		}
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("whichkey: expected rex-whichkey [options]; use --help for normal launch and setup")
	}
	if *action == "" {
		return errors.New("whichkey: --action must name the registered configuration launch action")
	}
	options := whichkey.LaunchOptions{}
	flags.Visit(func(option *flag.Flag) {
		switch option.Name {
		case "server":
			options.Server = server
		case "client":
			options.Client = client
		case "action":
			options.Action = *action
		}
	})
	// Parsing, including help/errors, finishes before any Rex lookup or call.
	return whichkey.LaunchPopup(ctx, options)
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
