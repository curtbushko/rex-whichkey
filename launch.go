package whichkey

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"temp/rex"
)

// launchOrigin retains original invocation authority separately from the popup.
type launchOrigin struct {
	Server      string `json:"server"`
	SessionID   string `json:"session_id"`
	BlockID     string `json:"block_id"`
	ClientID    string `json:"client_id"`
	ConnectedAt string `json:"connected_at"`
	selection   *rex.CLISelection
}

func popupConnectionGeneration(value string) bool {
	return len(value) <= 1024 && utf8.ValidString(value) && strings.TrimSpace(value) != "" &&
		strings.IndexFunc(value, unicode.IsControl) < 0
}
