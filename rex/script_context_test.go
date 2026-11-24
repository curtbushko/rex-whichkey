package rex

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// terminal.process uses user_id; whoami's separate principal uses uid.
// Keep literal wire keys here so a renamed production field cannot also rename
// the fixture and silently preserve an incompatible decoder.
func TestScriptForegroundCWDWire(t *testing.T) {
	uid := os.Geteuid()
	for _, test := range []struct {
		name, output string
		valid        bool
	}{
		{"native user_id", fmt.Sprintf(`{"child":{"cwd":"/child","user_id":%d},"foreground":{"pid":123,"user_id":%d,"user":"fixture","name":"sh","path":"/bin/sh","invoked_path":"/bin/sh","argv0":"sh","cwd":"/origin"},"home":"/home"}`, uid, uid), true},
		{"missing user_id", `{"foreground":{"cwd":"/origin"}}`, false},
		{"wrong account", fmt.Sprintf(`{"foreground":{"cwd":"/origin","user_id":%d}}`, uid+1), false},
		{"obsolete uid only", fmt.Sprintf(`{"foreground":{"cwd":"/origin","uid":%d}}`, uid), false},
		{"obsolete uid cannot override wrong account", fmt.Sprintf(`{"foreground":{"cwd":"/origin","user_id":%d,"uid":%d}}`, uid+1, uid), false},
		{"null user_id", `{"foreground":{"cwd":"/origin","user_id":null}}`, false},
		{"string user_id", fmt.Sprintf(`{"foreground":{"cwd":"/origin","user_id":"%d"}}`, uid), false},
		{"no foreground", fmt.Sprintf(`{"child":{"cwd":"/child","user_id":%d}}`, uid), false},
		{"no foreground cwd", fmt.Sprintf(`{"child":{"cwd":"/child","user_id":%d},"foreground":{"user_id":%d}}`, uid, uid), false},
		{"relative cwd", fmt.Sprintf(`{"foreground":{"cwd":"relative","user_id":%d}}`, uid), false},
		{"NUL cwd", fmt.Sprintf(`{"foreground":{"cwd":"/origin\u0000","user_id":%d}}`, uid), false},
	} {
		t.Run(test.name, func(t *testing.T) {
			client, logPath := apiFixtureClient(t, test.output, 0)
			cwd, err := client.scriptForegroundCWD(context.Background(), Block{blockID: "captured-origin"})
			if test.valid {
				require.NoError(t, err)
				require.Equal(t, "/origin", cwd)
			} else {
				require.Error(t, err)
				require.Empty(t, cwd, "refusal must not expose a fallback cwd")
			}
			data, readErr := os.ReadFile(logPath)
			require.NoError(t, readErr)
			require.Equal(t, 1, strings.Count(string(data), "\n"), "read only the retained origin, without retries")
			var argv []string
			require.NoError(t, json.Unmarshal(data, &argv))
			require.Equal(t, []string{"--server", client.server, "--autostart=false", "--timeout", "5s", "api", "call", "--session", client.session,
				"--", "com.superlogical.terminal.process", `{"args":{},"block_id":"captured-origin","session_id":"captured-session"}`}, argv)
		})
	}
}

func TestNativeScriptHostIdentity(t *testing.T) {
	uid := os.Geteuid()
	host := ScriptHost{Endpoint: "unix:///tmp/native.sock", HostID: "native-machine-account", PID: 123}
	host.Principal.Kind = "peer"
	host.Principal.Transport = scriptUnixTransport
	host.Principal.UID = &uid
	data, err := json.Marshal(host)
	require.NoError(t, err)
	parsed, err := decodeNativeScriptHost(data)
	require.NoError(t, err)
	require.Equal(t, host, parsed)
	for _, field := range []string{"host", "pid", "uid", "principal kind", "transport", "endpoint"} {
		candidate := host
		switch field {
		case "host":
			candidate.HostID = ""
		case "pid":
			candidate.PID = 0
		case "uid":
			candidate.Principal.UID = nil
		case "principal kind":
			candidate.Principal.Kind = "tailscale"
		case "transport":
			candidate.Principal.Transport = "tcp"
		case "endpoint":
			candidate.Endpoint = "example.org:7297"
		}
		encoded, marshalErr := json.Marshal(candidate)
		require.NoError(t, marshalErr)
		_, decodeErr := decodeNativeScriptHost(encoded)
		require.Error(t, decodeErr, field)
	}
	for _, raw := range []string{`null`, `{}`, `{"endpoint":"unix:///tmp/a","endpoint":"unix:///tmp/b"}`, `{"host_id":42}`, `{"pid":1.5}`} {
		_, err := decodeNativeScriptHost([]byte(raw))
		require.Error(t, err)
	}
}
