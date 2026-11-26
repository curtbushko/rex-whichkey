package rex

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	resizeFloating   = "floating"
	resizeTiled      = "tiled"
	resizeBlocks     = "blocks"
	resizeSuccess    = "success"
	resizeKind       = "kind"
	resizeUncertain  = "uncertain"
	resizeUnverified = "unverified"
)

func resizeView(t *testing.T, bounds Bounds, problem string) json.RawMessage {
	t.Helper()
	view, origin, own := directSnapshot(t)
	kind := resizeFloating
	blocks := []any{map[string]string{fieldBlock: own.ID}}
	var geometry any = bounds
	switch problem {
	case "shared":
		blocks = append(blocks, map[string]string{fieldBlock: "block:00000000-0000-0000-0000-000000000009"})
	case resizeTiled:
		kind = resizeTiled
	case "missing geometry":
		geometry = nil
	}
	data, err := json.Marshal(map[string]any{fieldSession: view.session.ID, "active_window_id": origin.WindowID,
		"windows": []any{map[string]any{"window_id": origin.WindowID, "focused_block_id": own.ID,
			"layers": []any{
				map[string]any{fieldLayer: origin.LayerID, resizeKind: resizeTiled, resizeBlocks: []any{map[string]string{fieldBlock: origin.ID}}},
				map[string]any{fieldLayer: own.LayerID, resizeKind: kind, fieldBounds: geometry, resizeBlocks: blocks},
			}}}})
	require.NoError(t, err)
	return data
}

func TestLayerResizeExactOwnershipAndReadback(t *testing.T) {
	for _, scenario := range []string{resizeSuccess, "shared", resizeTiled, "missing geometry", "changed geometry", resizeUnverified, resizeUncertain} {
		t.Run(scenario, func(t *testing.T) {
			before := Bounds{X: .1, Y: .4, W: .8, H: .2}
			after := Bounds{X: .1, Y: .45, W: .8, H: .1}
			initial := resizeView(t, before, "")
			preflight := resizeView(t, before, scenario)
			readback := resizeView(t, after, "")
			if scenario == "changed geometry" {
				preflight = readback
			}
			if scenario == resizeUnverified {
				readback = initial
			}
			sequence, err := json.Marshal([]json.RawMessage{initial, preflight, json.RawMessage(`{}`), readback})
			require.NoError(t, err)
			client, logPath := apiFixtureClient(t, string(initial), 0)
			t.Setenv("GORACE", os.Getenv("GORACE")+" atexit_sleep_ms=0")
			t.Setenv("REX_WHICHKEY_API_OUTPUT_SEQUENCE", string(sequence))
			view, origin, own := directSnapshot(t)
			client.session = view.session.ID
			snapshot, err := client.SessionSnapshot(context.Background(), client.session)
			require.NoError(t, err)
			placedOrigin, err := snapshot.Block(origin.ID)
			require.NoError(t, err)
			layer, _, err := snapshot.DirectPopup(own.ID, placedOrigin)
			require.NoError(t, err)
			captured, available := layer.InitialBounds()
			require.True(t, available)
			require.Equal(t, before, captured)
			// Exported display IDs must not redirect a native mutation.
			layer.ID, layer.BlockID = "other", "other"
			if scenario == resizeUncertain {
				t.Setenv("REX_WHICHKEY_API_FAIL_AT", "3")
			}
			err = layer.Resize(context.Background(), before, after)
			if scenario == resizeSuccess {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
			data, err := os.ReadFile(logPath)
			require.NoError(t, err)
			mutations := strings.Count(string(data), "session.resize_layer")
			if scenario == resizeSuccess || scenario == resizeUnverified || scenario == resizeUncertain {
				require.Equal(t, 1, mutations)
				calls := make([][]string, 0, 4)
				for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
					var argv []string
					require.NoError(t, json.Unmarshal([]byte(line), &argv))
					calls = append(calls, argv)
				}
				var request map[string]json.RawMessage
				require.NoError(t, json.Unmarshal([]byte(calls[2][len(calls[2])-1]), &request))
				require.JSONEq(t, `"`+own.LayerID+`"`, string(request[fieldLayer]))
				require.JSONEq(t, `{"x":0.1,"y":0.45,"w":0.8,"h":0.1}`, string(request[fieldBounds]))
				require.NotContains(t, request, fieldFocus)
			} else {
				require.Zero(t, mutations, "unverified ownership must prevent resize")
			}
		})
	}
}
