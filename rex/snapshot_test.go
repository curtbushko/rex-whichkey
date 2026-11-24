package rex

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScriptLayerStateSnapshot(t *testing.T) {
	for _, state := range []string{"", "normal", "minimized", "unknown"} {
		t.Run(state, func(t *testing.T) {
			view, origin, own := directSnapshot(t)
			data, err := json.Marshal(map[string]any{"session_id": view.session.ID, "windows": []any{map[string]any{"window_id": origin.WindowID, "layers": []any{map[string]any{"layer_id": own.LayerID, "kind": "floating", "state": state, "blocks": []any{map[string]string{"block_id": own.ID}}}}}}})
			require.NoError(t, err)
			var wire sessionView
			require.NoError(t, json.Unmarshal(data, &wire))
			layers, err := view.session.snapshotLayers(wire)
			require.NoError(t, err)
			require.Equal(t, state, layers[0].state)
		})
	}
}
