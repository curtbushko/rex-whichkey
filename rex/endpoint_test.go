package rex_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"temp/rex"
)

const macOSAddress = "unix:///Users/me/Library/Application%20Support/rex/server.sock"

func TestCanonicalEndpoint(t *testing.T) {
	for _, test := range []struct{ name, input, want string }{
		{"macOS CLI URI", macOSAddress, macOSAddress},
		{"macOS socket path", "/Users/me/Library/Application Support/rex/server.sock", macOSAddress},
		{"Linux socket", "/run/user/1000/rex/server.sock", "unix:///run/user/1000/rex/server.sock"},
		{"TCP", "localhost:1234", "localhost:1234"},
		{"qualified remote", "studio.example", "studio.example"},
		{"IPv6", "[::1]:1234", "[::1]:1234"},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := rex.CanonicalEndpoint(test.input)
			require.NoError(t, err)
			require.Equal(t, test.want, got)
			again, err := rex.CanonicalEndpoint(got)
			require.NoError(t, err)
			require.Equal(t, got, again)
		})
	}
	for _, input := range []string{"", "host-label", ".bad.example", "bad..example", "-bad.example", "bad-.example", "studio.example/path", "studio.example\n", "localhost:0", "localhost:abc", "host name:1234", "unix://host/path", "unix:///", "unix:///tmp/socket?x=1", "unix:///tmp/socket#frag", "unix:///tmp/%00socket", "unix:///tmp/%0asocket", "unix:///tmp/%zz", "/tmp/socket\n"} {
		t.Run(input, func(t *testing.T) {
			_, err := rex.CanonicalEndpoint(input)
			require.Error(t, err)
		})
	}
}

func TestResolveEndpointInheritsRexCLIAddress(t *testing.T) {
	t.Setenv("REX_SERVER", macOSAddress)
	got, err := rex.ResolveEndpoint(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, macOSAddress, got)
}
