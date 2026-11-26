package rex

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPresentedSizeRequiresExactLivePopupGrid(t *testing.T) {
	for _, test := range []struct {
		name, output string
		ready        bool
	}{
		{"bootstrap", `{"columns":80,"rows":24,"owner":null,"desired":[]}`, false},
		{"live", `{"columns":80,"rows":24,"owner":"client:00000000-0000-0000-0000-000000000001","desired":[{"client_id":"client:00000000-0000-0000-0000-000000000001","columns":80,"rows":24}]}`, true},
		{"different owner", `{"columns":80,"rows":24,"owner":"client:00000000-0000-0000-0000-000000000002","desired":[{"client_id":"client:00000000-0000-0000-0000-000000000001","columns":80,"rows":24}]}`, false},
		{"stale pty", `{"columns":80,"rows":10,"owner":"client:00000000-0000-0000-0000-000000000001","desired":[{"client_id":"client:00000000-0000-0000-0000-000000000001","columns":80,"rows":10}]}`, false},
		{"missing", `{}`, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			client, _ := apiFixtureClient(t, test.output, 0)
			t.Setenv("GORACE", os.Getenv("GORACE")+" atexit_sleep_ms=0")
			view, origin, own := directSnapshot(t)
			client.session = view.session.ID
			view.session.client = client
			for i := range view.layers {
				for j := range view.layers[i].blocks {
					view.layers[i].blocks[j].client = client
				}
			}
			origin.client = client
			layer, _, err := view.DirectPopup(own.ID, origin)
			require.NoError(t, err)
			ready, err := layer.PresentedSize(context.Background(), 80, 24)
			require.NoError(t, err)
			require.Equal(t, test.ready, ready)
		})
	}
}
