package whichkey

// Genuine historical completion-wire implementation retained for unchanged rename
// assertions and popup_legacy_support_test.go, not a production dispatch path.
import (
	"encoding/json"
	"errors"
	"strings"
	"unicode/utf8"
)

// CompletionVersion identifies the data-only post-return bridge format.
const CompletionVersion = 1

// CompletionReceiptVersion is negotiated separately from the strict legacy wire.
const CompletionReceiptVersion = 2

// Completion kinds select one retained action or release the invocation.
const (
	CompletionComplete = "complete"
	CompletionCancel   = "cancel"
)

// Completion carries only an invocation capability and a selection or cancellation.
// Origin, actions, arguments, callbacks and configuration identity stay retained
// in Rex Lua. Validation is not authentication or proof of a live origin.
type Completion struct {
	Version     int                  `json:"version"`
	Capability  string               `json:"capability"`
	ReceiptHash string               `json:"receipt_hash,omitempty"`
	Kind        string               `json:"kind"`
	Selection   *CompletionSelection `json:"selection,omitempty"`

	// Private child-result bookkeeping. Only Selection.Text crosses the Lua bridge.
	textInput *string
}

// CompletionSelection identifies a retained mapping. Confirmed records y at that
// mapping's prompt; it never overrides the retained mapping's confirmation flag.
type CompletionSelection struct {
	Keys        string `json:"keys"`
	CallbackRef string `json:"callback_ref,omitempty"`
	Confirmed   bool   `json:"confirmed"`
	// Text supplies only a bounded label for a retained default rename intent.
	Text *string `json:"text,omitempty"`
}

// ValidateInvocationCapability checks the canonical representation, not authority.
func ValidateInvocationCapability(capability string) error {
	if len(capability) != 64 {
		return errors.New("whichkey: capability must be 64 lowercase hex characters")
	}
	for _, c := range capability {
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') {
			return errors.New("whichkey: capability must be 64 lowercase hex characters")
		}
	}
	return nil
}

// Validate checks message structure. It does not check the Lua registry, expiry,
// origin, replay state, connection or configuration generation.
func (c Completion) Validate() error {
	if err := c.validateVersion(); err != nil {
		return err
	}
	if err := ValidateInvocationCapability(c.Capability); err != nil {
		return err
	}
	switch c.Kind {
	case CompletionCancel:
		if c.Selection != nil {
			return errors.New("whichkey: cancellation must not carry a selection")
		}
	case CompletionComplete:
		if c.Selection == nil {
			return errors.New("whichkey: completion requires a selection")
		}
		canonical, err := NormalizeKeys(c.Selection.Keys)
		if err != nil || canonical != c.Selection.Keys {
			return errors.New("whichkey: completion keys must be a canonical full sequence")
		}
		if err := validateText(c.Selection.CallbackRef); err != nil {
			return errors.New("whichkey: invalid completion callback reference")
		}
		if c.Selection.CallbackRef != "" && strings.TrimSpace(c.Selection.CallbackRef) == "" {
			return errors.New("whichkey: completion callback reference must not be blank")
		}
		if err := validateCompletionText(c.Selection); err != nil {
			return err
		}
	default:
		return errors.New("whichkey: invalid completion kind")
	}
	return nil
}

func (c Completion) validateVersion() error {
	switch c.Version {
	case CompletionVersion:
		if c.ReceiptHash != "" {
			return errors.New("whichkey: version 1 completion cannot carry a receipt")
		}
	case CompletionReceiptVersion:
		if ValidateInvocationCapability(c.ReceiptHash) != nil {
			return errors.New("whichkey: version 2 completion requires its exact receipt hash")
		}
	default:
		return errors.New("whichkey: unsupported completion version")
	}
	return nil
}

func validateCompletionText(selection *CompletionSelection) error {
	if selection.Text == nil {
		return nil
	}
	if (selection.Keys != "tr" && selection.Keys != "wr") || selection.CallbackRef != "" || selection.Confirmed {
		return errors.New("whichkey: text requires an inline rename selection")
	}
	if utf8.RuneCountInString(*selection.Text) > 128 {
		return errors.New("whichkey: resource rename label exceeds Rex's documented 128-character limit")
	}
	return validateRenameText(*selection.Text)
}

// ValidateSelection checks identity and confirmation against a retained, evaluated
// snapshot. Hidden actions remain selectable. Lua must still resolve its OWN
// retained snapshot and recheck conditions, authority and origin before dispatch.
// A valid cancellation needs no mapping and dispatches nothing.
func (c Completion) ValidateSelection(mappings []Mapping) error {
	if err := c.Validate(); err != nil {
		return err
	}
	if c.Kind == CompletionCancel {
		return nil
	}
	if err := ValidateMappings(mappings); err != nil {
		return err
	}
	for _, mapping := range mappings {
		if !mapping.Enabled() || (mapping.Action == "" && mapping.CallbackRef == "") {
			continue
		}
		canonical, err := NormalizeKeys(mapping.Keys)
		if err != nil {
			return err
		}
		if canonical != c.Selection.Keys {
			continue
		}
		if mapping.CallbackRef != c.Selection.CallbackRef {
			return errors.New("whichkey: completion callback reference does not match selection")
		}
		if mapping.Confirm != c.Selection.Confirmed {
			return errors.New("whichkey: completion confirmation does not match retained requirement")
		}
		if nativeRenameMapping(mapping) != (c.Selection.Text != nil) {
			return errors.New("whichkey: completion text does not match retained rename intent")
		}
		return nil
	}
	return errors.New("whichkey: completion does not select an enabled retained action")
}

// MarshalJSON emits a validated data-only message. Capability-bearing output is
// secret transport data, never a log or a user-visible diagnostic.
func (c Completion) MarshalJSON() ([]byte, error) {
	if c.textInput != nil {
		return nil, errors.New("whichkey: native form data cannot use the configured completion wire format")
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	type message Completion
	return json.Marshal(message(c))
}

// UnmarshalJSON rejects duplicate/unknown fields, nulls, missing required fields,
// invalid UTF-8/surrogates, coercions and trailing values, including in selection.
// The receiver is unchanged on failure. Registry authentication remains mandatory.
func (c *Completion) UnmarshalJSON(data []byte) error {
	if c == nil {
		return errors.New("whichkey: completion receiver is nil")
	}
	fields, err := completionObject(data, "version", "capability", "receipt_hash", "kind", "selection")
	if err != nil {
		return err
	}
	var candidate Completion
	if decodeErr := json.Unmarshal(fields["version"], &candidate.Version); decodeErr != nil {
		return errors.New("whichkey: completion version must be an integer")
	}
	if raw, present := fields["receipt_hash"]; present {
		if candidate.Version != CompletionReceiptVersion {
			return errors.New("whichkey: receipt extension requires negotiated completion version 2")
		}
		if candidate.ReceiptHash, err = completionString(raw); err != nil {
			return err
		}
	}
	if candidate.Capability, err = completionString(fields["capability"]); err != nil {
		return err
	}
	if candidate.Kind, err = completionString(fields["kind"]); err != nil {
		return err
	}
	if raw, present := fields["selection"]; present {
		candidate.Selection, err = decodeCompletionSelection(raw)
		if err != nil {
			return err
		}
	}
	if validationErr := candidate.Validate(); validationErr != nil {
		return validationErr
	}
	*c = candidate
	return nil
}

func decodeCompletionSelection(data []byte) (*CompletionSelection, error) {
	fields, err := completionObject(data, "keys", "callback_ref", "confirmed", "text")
	if err != nil {
		return nil, err
	}
	value := new(CompletionSelection)
	if value.Keys, err = completionString(fields["keys"]); err != nil {
		return nil, err
	}
	if raw, present := fields["callback_ref"]; present {
		if value.CallbackRef, err = completionString(raw); err != nil {
			return nil, err
		}
		if value.CallbackRef == "" {
			return nil, errors.New("whichkey: omit an absent completion callback reference")
		}
	}
	if decodeErr := json.Unmarshal(fields["confirmed"], &value.Confirmed); decodeErr != nil {
		return nil, errors.New("whichkey: completion confirmed must be an explicit boolean")
	}
	if raw, present := fields["text"]; present {
		text, parseErr := completionString(raw)
		if parseErr != nil {
			return nil, parseErr
		}
		value.Text = &text
	}
	return value, nil
}

// nativeRenameMapping preserves the historical Completion test contract. The
// supported Go menu uses action-based inlineRenameMapping, not legacy keys.
func nativeRenameMapping(mapping Mapping) bool {
	return (mapping.Keys == "tr" || mapping.Keys == "wr") && !mapping.Confirm && inlineRenameMapping(mapping)
}
