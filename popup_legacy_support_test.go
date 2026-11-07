package whichkey

// Original Completion-wire support for unchanged historical rename assertions.
// No production entry reads or writes private popup files.
import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"syscall"

	"golang.org/x/sys/unix"
)

const popupMessageLimit = menuConfigLimit + (64 << 10)

var (
	errRuntimeRelay   = errors.New("whichkey: private popup file identity or access failed")
	errRuntimeStartup = errors.New("whichkey: private popup startup or result failed")
)

func openPopupDirectory(directory string) (*os.Root, error) {
	info, err := os.Lstat(directory)
	if err != nil || !filepath.IsAbs(directory) || !info.IsDir() || info.Mode().Perm() != 0700 {
		return nil, errors.New("whichkey: popup directory must be private and absolute")
	}
	owner, ok := info.Sys().(*syscall.Stat_t)
	if !ok || owner.Uid != uint32(os.Geteuid()) {
		return nil, errors.New("whichkey: popup directory has the wrong owner")
	}
	root, err := os.OpenRoot(directory)
	if err != nil {
		return nil, errors.New("whichkey: cannot open private popup directory")
	}
	return root, nil
}

func checkPopupSlot(file *os.File, fifo bool) error {
	info, err := file.Stat()
	if err != nil || info.Mode().Perm() != 0600 || (fifo && info.Mode()&os.ModeNamedPipe == 0) || (!fifo && !info.Mode().IsRegular()) {
		return errRuntimeRelay
	}
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok || stat.Uid != uint32(os.Geteuid()) || stat.Nlink != 1 {
		return errRuntimeRelay
	}
	return nil
}

func readPopupResult(directory, capability string, config Config) (message Completion, outcome error) {
	root, err := openPopupDirectory(directory)
	if err != nil {
		return Completion{}, err
	}
	defer func() { outcome = errors.Join(outcome, root.Close()) }()
	file, err := root.OpenFile("result.json", os.O_RDONLY|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if err != nil {
		return Completion{}, errors.New("whichkey: cannot open popup result")
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	if statErr := checkPopupSlot(file, false); statErr != nil {
		return Completion{}, errRuntimeStartup
	}
	data, err := io.ReadAll(io.LimitReader(file, popupMessageLimit+1))
	if err != nil || len(data) > popupMessageLimit {
		return Completion{}, errors.New("whichkey: popup result exceeds transport bound")
	}
	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		return Completion{Version: CompletionVersion, Capability: capability, Kind: CompletionCancel}, nil
	}
	plain, textInput, err := nativeTextSelection(data, config.textEntry)
	if err != nil {
		return Completion{}, err
	}
	// Wrap the raw selection for the authoritative strict Completion decoder;
	// decoding CompletionSelection alone would accept missing false flags.
	wrapped, err := json.Marshal(struct {
		Version    int             `json:"version"`
		Capability string          `json:"capability"`
		Kind       string          `json:"kind"`
		Selection  json.RawMessage `json:"selection"`
	}{CompletionVersion, capability, CompletionComplete, plain})
	if err != nil || json.Unmarshal(wrapped, &message) != nil {
		return Completion{}, errors.New("whichkey: popup returned an invalid retained selection")
	}
	if message.Selection != nil {
		message.Selection.Text = textInput
	}
	if err := message.ValidateSelection(config.Spec); err != nil {
		return Completion{}, errors.New("whichkey: popup returned an invalid retained selection")
	}
	message.textInput = textInput
	return message, nil
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
