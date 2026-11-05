package whichkey

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfiguredRenameWireRejectsUntrustedFields(t *testing.T) {
	prefix := `{"version":1,"kind":"complete","capability":"` + strings.Repeat("a", 64) + `","selection":`
	for _, selection := range []string{
		`{"keys":"tr","confirmed":false,"text":null}`,
		`{"keys":"tr","confirmed":false,"text":42}`,
		`{"keys":"tr","confirmed":false,"text":"Name","text":"Other"}`,
		`{"keys":"tr","confirmed":false,"text":"Name","window_id":"replacement"}`,
		`{"keys":"tr","confirmed":true,"text":"Name"}`,
		`{"keys":"tr","confirmed":false,"text":"Name","callback_ref":"cb:1:1"}`,
		`{"keys":"tr","confirmed":false,"text":"\ud800"}`,
		`{"keys":"tr","confirmed":false,"text":"bad\u0000name"}`,
		`{"keys":"x","confirmed":false,"text":"Name"}`,
	} {
		var completion Completion
		require.Error(t, json.Unmarshal([]byte(prefix+selection+`}`), &completion), "%s", selection)
	}
}
