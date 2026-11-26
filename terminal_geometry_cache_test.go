//go:build darwin || linux

package whichkey

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"temp/rex"

	"github.com/stretchr/testify/require"
)

func TestGeometryRecorderOnlySavesPresentedSettledRoot(t *testing.T) {
	for _, scenario := range []string{"root", "submenu", "pending", "wrong rows", "unpresented", "query failure", "canceled cache read"} {
		t.Run(scenario, func(t *testing.T) {
			var saved []float64
			recorder := popupGeometryRecorder{save: func(height float64) { saved = append(saved, height) }}
			sizing := newPopupSizing(rex.Bounds{X: .1, Y: .45, W: .8, H: .1})
			sizing.maximum.H = .2
			demand := popupSizingDemand{rows: 6, root: true}
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			switch scenario {
			case "submenu":
				demand.root = false
			case "pending":
				sizing.waiting = true
			case "wrong rows":
				demand.rows = 20
			case "canceled cache read":
				cancel()
			}
			presented := func(context.Context, int, int) (bool, error) {
				if scenario == "query failure" {
					return false, errors.New("size unavailable")
				}
				return scenario != "unpresented", nil
			}
			recorder.remember(ctx, sizing, demand, 80, 6, presented)
			recorder.nextCheck = time.Time{}
			recorder.remember(ctx, sizing, demand, 80, 6, presented)
			if scenario == "root" {
				require.Equal(t, []float64{.1}, saved, "unchanged root is written once")
			} else {
				require.Empty(t, saved)
			}
		})
	}
}

func TestGeometryCacheWriteFailureDoesNotFailMenuSizing(t *testing.T) {
	path := filepath.Join(t.TempDir(), "geometry.json")
	require.NoError(t, os.Mkdir(path, 0700))
	attempted := make(chan struct{}, 1)
	settings := popupSizingSettings{maximumHeight: .2, save: func(height float64) {
		_ = writeGeometryCache(path, geometryCacheEntry{Version: 1, Width: .8, MaximumHeight: .2, Height: height})
		attempted <- struct{}{}
	}}
	job := startPopupSizing(context.Background(), rex.Bounds{X: .1, Y: .45, W: .8, H: .1},
		func() (int, int, error) { return 80, 6, nil }, sizingPresented,
		func(context.Context, rex.Bounds, rex.Bounds) error {
			t.Error("a correct warm-start height must not resize")
			return nil
		}, settings)
	t.Cleanup(func() { _ = job.stop() })
	job.requestGeometry(6, true)
	select {
	case <-attempted:
	case <-time.After(2 * time.Second):
		t.Fatal("settled root was not offered to cache")
	}
	require.NoError(t, job.stop(), "cache filesystem errors must not suppress menu actions")
}

func TestCachedInitialHeightKeepsConfiguredGrowthCap(t *testing.T) {
	sizing := newPopupSizing(rex.Bounds{X: .1, Y: .45, W: .8, H: .1})
	sizing.maximum.H = .2
	_, changed := sizing.fit(6, 80, 6)
	require.False(t, changed, "a correct warm-start height needs no resize")
	bounds, changed := sizing.fit(20, 80, 6)
	require.True(t, changed)
	require.Equal(t, .2, bounds.H, "longer submenus retain the original height cap")
}
