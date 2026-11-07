package whichkey

import "io"

// launchOutput is isolated historical diagnostic compatibility for unchanged
// tests. No supported public launch or Rex execution path constructs it.
type launchOutput struct {
	remaining           int
	legacyMatch         int
	legacyClientBinding bool
}

func (output *launchOutput) Write(data []byte) (int, error) {
	if len(data) > output.remaining {
		return 0, io.ErrShortWrite
	}
	output.remaining -= len(data)
	// Retain only a prefix length and a boolean, never process output (which
	// may contain capabilities). The literal has no overlapping proper prefix.
	const legacy = "rex_whichkey: launch origin client_id does not match the actual action context"
	for _, char := range data {
		if char != legacy[output.legacyMatch] {
			output.legacyMatch = 0
		}
		if char == legacy[output.legacyMatch] {
			output.legacyMatch++
			if output.legacyMatch == len(legacy) {
				output.legacyClientBinding = true
				output.legacyMatch = 0
			}
		}
	}
	return len(data), nil
}
