package whichkey

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// WriteMenuResult encodes the UI-only menu command outcome, not direct dispatch.
// It never sends replacement action names or arguments. A dismissal is JSON null,
// not an empty selection or a request to execute a default action.
func WriteMenuResult(output io.Writer, event MenuEvent) error {
	if event.Dismissed && event.Selection == nil {
		if err := json.NewEncoder(output).Encode(nil); err != nil {
			return fmt.Errorf("whichkey: write dismissal: %w", err)
		}
		return nil
	}
	if event.Dismissed || event.Selection == nil {
		return errors.New("whichkey: menu returned no unambiguous outcome")
	}
	if event.Selection.textInput != nil {
		return errors.New("whichkey: native text entry requires its private result writer")
	}
	mapping := event.Selection.Mapping()
	if err := ValidateMappings([]Mapping{mapping}); err != nil {
		return err
	}
	if !mapping.Enabled() || mapping.Action == "" && mapping.CallbackRef == "" || mapping.Confirm && !event.Selection.Confirmed() {
		return errors.New("whichkey: menu returned an ineligible selection")
	}
	result := struct {
		Keys        string `json:"keys"`
		CallbackRef string `json:"callback_ref,omitempty"`
		Confirmed   bool   `json:"confirmed"`
	}{mapping.Keys, mapping.CallbackRef, event.Selection.Confirmed()}
	if err := json.NewEncoder(output).Encode(result); err != nil {
		return fmt.Errorf("whichkey: write selection: %w", err)
	}
	return nil
}
