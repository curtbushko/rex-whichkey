package whichkey

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"golang.org/x/sys/unix"
)

const nativeTextFile = "native-input.json"

func readNativeTextSpec(root *os.Root) (spec *textEntrySpec, outcome error) {
	file, err := root.OpenFile(nativeTextFile, os.O_RDONLY|unix.O_NOFOLLOW, 0)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.New("whichkey: cannot read native text-entry presentation")
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	info, err := file.Stat()
	if err != nil || !info.Mode().IsRegular() || info.Mode().Perm() != 0600 {
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

// The extra text field belongs only to native private transport. Ordinary
// WriteMenuResult and the configured Completion/Lua wire format stay unchanged.
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
		return data, nil, nil // Configured decoding still rejects any extra field.
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
		return nil, nil, errors.New("whichkey: rename requires confirmed text input")
	}
	value, err := completionString(raw)
	if err != nil {
		return nil, nil, err
	}
	if err := validateRenameText(value); err != nil {
		return nil, nil, err
	}
	delete(fields, "text")
	plain, err := json.Marshal(fields)
	return plain, &value, err
}
