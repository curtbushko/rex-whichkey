package whichkey

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"golang.org/x/sys/unix"
)

// Text presentation is private data, separate from public menu options.
const nativeTextFile = "native-input.json"

func readNativeTextSpec(root *os.Root) (spec *textEntrySpec, outcome error) {
	file, err := root.OpenFile(nativeTextFile, os.O_RDONLY|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.New("whichkey: cannot read native text-entry presentation")
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	if statErr := checkPopupSlot(file, false); statErr != nil {
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

// The child carries bounded name data; the parent retains action and resource.
func writePopupMenuResult(output io.Writer, event MenuEvent, config Config) error {
	if event.provider != nil {
		if event.Selection == nil || event.Dismissed || event.Selection.textInput != nil || event.update != nil {
			return errors.New("whichkey: invalid provider result binding")
		}
		return json.NewEncoder(output).Encode(struct {
			Keys      string         `json:"keys"`
			Confirmed bool           `json:"confirmed"`
			Provider  *renameBinding `json:"provider"`
		}{event.Selection.mapping.Keys, event.Selection.confirmed, event.provider})
	}
	if event.Selection == nil || event.Selection.textInput == nil {
		if event.update != nil {
			return errors.New("whichkey: update binding requires inline rename text")
		}
		return WriteMenuResult(output, event)
	}
	mapping := event.Selection.Mapping()
	if config.textEntry == nil || event.Dismissed || !inlineRenameMapping(mapping) || event.Selection.Confirmed() != mapping.Confirm {
		return errors.New("whichkey: unexpected native text-entry result")
	}
	if err := validateRenameText(*event.Selection.textInput); err != nil {
		return err
	}
	return json.NewEncoder(output).Encode(struct {
		Keys      string         `json:"keys"`
		Confirmed bool           `json:"confirmed"`
		Text      string         `json:"text"`
		Update    *renameBinding `json:"update,omitempty"`
	}{mapping.Keys, event.Selection.Confirmed(), *event.Selection.textInput, event.update})
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
	raw, supplied := fields["text"]
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
