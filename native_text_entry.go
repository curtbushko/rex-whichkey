package whichkey

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"

	"golang.org/x/sys/unix"

	"temp/rex"
)

// Keep the private filename compatible with native popups; configured popups
// now use the same input presentation without changing Config JSON.
const nativeTextFile = "native-input.json"

func prepareConfiguredRename(ctx context.Context, request popupRequest, config Config, selection *rex.CLISelection) (Config, error) {
	needed := false
	for _, mapping := range config.Spec {
		needed = needed || mapping.Enabled() && nativeRenameMapping(mapping)
	}
	if !needed {
		return config, nil
	}
	ctx, cancel := context.WithDeadline(ctx, time.Unix(request.Expires, 0))
	defer cancel()
	client, err := rex.NewContext(ctx, rex.Options{Selection: selection, Server: request.Server,
		Session: request.SessionID, Timeout: 3 * time.Second})
	if err != nil {
		return Config{}, errors.New("whichkey: cannot construct rename presentation client")
	}
	session, err := client.Session(ctx, request.SessionID)
	if err != nil {
		return Config{}, errors.New("whichkey: cannot inspect rename session label")
	}
	block, err := session.Block(ctx, request.BlockID)
	if err != nil || request.retained == nil {
		return Config{}, errors.New("whichkey: cannot inspect the retained rename origin")
	}
	if placementErr := request.retained.checkPlacement(block); placementErr != nil {
		return Config{}, errors.New("whichkey: cannot inspect the retained rename origin")
	}
	window, err := session.Window(ctx, block.WindowID)
	if err != nil {
		return Config{}, errors.New("whichkey: cannot inspect rename tab label")
	}
	config.textEntry = &textEntrySpec{Initial: window.Label, SessionInitial: session.Label}
	if len(config.textEntry.Initial) > renameTextLimit || validateText(config.textEntry.Initial) != nil {
		config.textEntry.Initial = ""
	}
	if len(config.textEntry.SessionInitial) > renameTextLimit || validateText(config.textEntry.SessionInitial) != nil {
		config.textEntry.SessionInitial = ""
	}
	return config, nil
}

func readNativeTextSpec(root *os.Root) (spec *textEntrySpec, outcome error) {
	file, err := root.OpenFile(nativeTextFile, os.O_RDONLY|unix.O_NOFOLLOW, 0)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.New("whichkey: cannot read native text-entry presentation")
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	if _, statErr := runtimeSlotStat(file, nativeTextFile, false); statErr != nil {
		return nil, errors.New("whichkey: invalid native text-entry file")
	}
	data, err := io.ReadAll(io.LimitReader(file, popupMessageLimit+1))
	if err != nil || len(data) > popupMessageLimit {
		return nil, errors.New("whichkey: native text-entry presentation exceeds transport bound")
	}
	fields, err := completionObject(data, "initial", "session_initial")
	if err != nil {
		return nil, err
	}
	initial, err := completionString(fields["initial"])
	if err != nil {
		return nil, err
	}
	sessionInitial, err := completionString(fields["session_initial"])
	if err != nil {
		return nil, err
	}
	return &textEntrySpec{Initial: initial, SessionInitial: sessionInitial}, nil
}

// The child result carries only bounded name data; controller and Lua retain
// the action and resource identity independently.
func writePopupMenuResult(output io.Writer, event MenuEvent, config Config) error {
	if event.Selection == nil || event.Selection.textInput == nil {
		return WriteMenuResult(output, event)
	}
	mapping := event.Selection.Mapping()
	if config.textEntry == nil || event.Dismissed || !nativeRenameMapping(mapping) || event.Selection.Confirmed() {
		return errors.New("whichkey: unexpected native text-entry result")
	}
	if err := validateRenameText(*event.Selection.textInput); err != nil {
		return err
	}
	return json.NewEncoder(output).Encode(struct {
		Keys      string `json:"keys"`
		Confirmed bool   `json:"confirmed"`
		Text      string `json:"text"`
	}{mapping.Keys, false, *event.Selection.textInput})
}

// Strip only an authorized native rename form's text before the canonical strict
// Completion decoder. Parent-retained mappings still validate keys/confirmation;
// this data never supplies an action, arguments, app or resource identity.
func nativeTextSelection(data []byte, spec *textEntrySpec) ([]byte, *string, error) {
	if spec == nil {
		if _, err := completionObject(data, "keys", "callback_ref", "confirmed"); err != nil {
			return nil, nil, err // No input page was authorized for this child.
		}
		return data, nil, nil
	}
	fields, err := completionObject(data, "keys", "callback_ref", "confirmed", "text")
	if err != nil {
		return nil, nil, err
	}
	keys, err := completionString(fields["keys"])
	if err != nil {
		return nil, nil, err
	}
	raw, supplied := fields["text"]
	if keys != "tr" && keys != "wr" {
		if supplied {
			return nil, nil, errors.New("whichkey: text input does not match the retained form")
		}
		return data, nil, nil
	}
	if !supplied {
		return data, nil, nil // The retained mapping decides whether text is required.
	}
	value, err := completionString(raw)
	if err != nil {
		return nil, nil, err
	}
	if validationErr := validateRenameText(value); validationErr != nil {
		return nil, nil, validationErr
	}
	delete(fields, "text")
	plain, err := json.Marshal(fields)
	return plain, &value, err
}
