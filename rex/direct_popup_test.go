package rex

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func directSnapshot(t *testing.T) (SessionSnapshot, Block, Block) {
	t.Helper()
	client := &Client{session: "session:00000000-0000-0000-0000-000000000001"}
	session := client.sessionResource(client.session, "Origin session")
	window := session.windowResource("window:00000000-0000-0000-0000-000000000002", "Origin window")
	window.Active = true
	origin := Block{ID: "block:00000000-0000-0000-0000-000000000003", WindowID: window.ID,
		LayerID: "layer:00000000-0000-0000-0000-000000000004", client: session.client,
		blockID: "block:00000000-0000-0000-0000-000000000003"}
	own := Block{ID: "block:00000000-0000-0000-0000-000000000005", WindowID: window.ID,
		LayerID: "layer:00000000-0000-0000-0000-000000000006", client: session.client,
		blockID: "block:00000000-0000-0000-0000-000000000005"}
	window.FocusedBlockID = own.ID
	view := SessionSnapshot{session: session, windows: []Window{window}, layers: []placedLayer{
		{id: origin.LayerID, kind: "tiling", blocks: []Block{origin}},
		{id: own.LayerID, kind: "floating", blocks: []Block{own}},
	}}
	return view, origin, own
}

func TestDirectPopupFiltersOnlyOwnedFloat(t *testing.T) {
	view, origin, own := directSnapshot(t)
	layer, filtered, err := view.DirectPopup(own.ID, origin)
	require.NoError(t, err)
	require.Equal(t, own.LayerID, layer.ID)
	require.Equal(t, own.ID, layer.Block().ID)
	require.Equal(t, []Block{origin}, filtered.Blocks())
	require.Equal(t, origin.ID, filtered.Windows()[0].FocusedBlockID)
	require.True(t, filtered.Windows()[0].Active)
	require.Equal(t, own.ID, view.Windows()[0].FocusedBlockID, "filtered markers must not mutate raw ownership evidence")
	require.Len(t, view.Blocks(), 2)
	_, err = filtered.Block(own.ID)
	require.Error(t, err)
}

func TestDirectPopupRejectsOwnershipAndPlacementChanges(t *testing.T) {
	for _, name := range []string{"same block", "tiled caller", "shared float", "moved origin", "different session", "missing own"} {
		t.Run(name, func(t *testing.T) {
			view, origin, own := directSnapshot(t)
			ownID := own.ID
			switch name {
			case "same block":
				ownID = origin.ID
			case "tiled caller":
				view.layers[1].kind = "tiling"
			case "shared float":
				view.layers[1].blocks = append(view.layers[1].blocks, origin)
			case "moved origin":
				origin.WindowID = "window:00000000-0000-0000-0000-000000000007"
			case "different session":
				origin.client = &Client{session: "session:00000000-0000-0000-0000-000000000008"}
			case "missing own":
				view.layers = view.layers[:1]
			}
			layer, filtered, err := view.DirectPopup(ownID, origin)
			require.Error(t, err)
			require.Empty(t, layer.ID)
			require.Empty(t, filtered.Blocks())
		})
	}
}
