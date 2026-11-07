package rex

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Re-execution is an isolated CLI boundary, not a Rex server or live action.
func TestMain(m *testing.M) {
	if logPath := os.Getenv("REX_WHICHKEY_API_FIXTURE"); logPath != "" {
		os.Exit(runAPIFixture(logPath))
	}
	os.Exit(m.Run())
}

func runAPIFixture(logPath string) int {
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return 90
	}
	err = json.NewEncoder(file).Encode(os.Args[1:])
	closeErr := file.Close()
	if err != nil || closeErr != nil {
		return 91
	}
	if _, writeErr := os.Stdout.WriteString(os.Getenv("REX_WHICHKEY_API_OUTPUT")); writeErr != nil {
		return 92
	}
	code, err := strconv.Atoi(os.Getenv("REX_WHICHKEY_API_EXIT"))
	if err != nil {
		return 93
	}
	return code
}

func apiFixtureClient(t *testing.T, output string, code int) (*Client, string) {
	t.Helper()
	executable, err := os.Executable()
	require.NoError(t, err)
	logPath := t.TempDir() + "/calls.jsonl"
	t.Setenv("REX_WHICHKEY_API_FIXTURE", logPath)
	t.Setenv("REX_WHICHKEY_API_OUTPUT", output)
	t.Setenv("REX_WHICHKEY_API_EXIT", strconv.Itoa(code))
	client, err := NewContext(context.Background(), Options{Selection: &CLISelection{Path: executable},
		Server: "unix:///tmp/isolated-api.sock", Session: "captured-session", Timeout: 5 * time.Second})
	require.NoError(t, err)
	return client, logPath
}

func TestCallExactArgumentsAndResult(t *testing.T) {
	client, logPath := apiFixtureClient(t, `{"value":7}`, 0)
	var result struct {
		Value int `json:"value"`
	}
	require.NoError(t, client.Call(context.Background(), "--user-method", map[string]string{"name": "a; $(not a shell)"}, &result))
	require.Equal(t, 7, result.Value)
	data, err := os.ReadFile(logPath)
	require.NoError(t, err)
	var argv []string
	require.NoError(t, json.Unmarshal(data, &argv))
	require.Equal(t, []string{"--server", client.server, "--autostart=false", "--timeout", "5s", "api", "call", "--session", client.session,
		"--", "--user-method", `{"name":"a; $(not a shell)"}`}, argv)
}

func TestCallFailuresNeverRetry(t *testing.T) {
	for _, test := range []struct {
		name, output string
		code         int
	}{
		{"invalid result", `not json`, 0},
		{"uncertain failure", `{"ignored":true}`, 1},
	} {
		t.Run(test.name, func(t *testing.T) {
			client, logPath := apiFixtureClient(t, test.output, test.code)
			err := client.Call(context.Background(), "mutate", nil, nil)
			require.ErrorContains(t, err, "not retried")
			data, readErr := os.ReadFile(logPath)
			require.NoError(t, readErr)
			require.Equal(t, 1, strings.Count(string(data), "\n"), "uncertain effects must not be retried")
		})
	}
}

func TestCallValidationPreventsExecution(t *testing.T) {
	client, logPath := apiFixtureClient(t, `{}`, 0)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	require.ErrorIs(t, client.Call(ctx, "mutate", nil, nil), context.Canceled)
	var nilResult *string
	for _, result := range []any{nilResult, "not a pointer", 7} {
		require.Error(t, client.Call(context.Background(), "mutate", nil, result))
	}
	for _, method := range []string{"", " ", "bad\x00method"} {
		require.Error(t, client.Call(context.Background(), method, nil, nil))
	}
	for _, arguments := range []any{[]string{"not an object"}, json.RawMessage(`null`), make(chan int)} {
		require.Error(t, client.Call(context.Background(), "mutate", arguments, nil))
	}
	_, err := os.Stat(logPath)
	require.ErrorIs(t, err, os.ErrNotExist)
}
