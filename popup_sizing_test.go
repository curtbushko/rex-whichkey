package whichkey

import (
	"math"
	"testing"

	"temp/rex"

	"github.com/stretchr/testify/require"
)

func TestPopupSizingFitsContentAndSubmenus(t *testing.T) {
	maximum := rex.Bounds{X: .1, Y: .4, W: .8, H: .2}
	sizing := newPopupSizing(maximum)
	// Model native chrome consuming two rows, with terminal cells rounded down.
	rows := func() int { return int(math.Floor(sizing.bounds.H*100 - 2)) }
	fit := func(wanted int) {
		t.Helper()
		for range 8 {
			_, changed := sizing.fit(wanted, 120, rows())
			if !changed {
				break
			}
		}
		require.InDelta(t, maximum.X, sizing.bounds.X, 1e-9)
		require.InDelta(t, maximum.W, sizing.bounds.W, 1e-9)
		require.InDelta(t, .5, sizing.bounds.Y+sizing.bounds.H/2, 1e-9)
		require.LessOrEqual(t, sizing.bounds.H, maximum.H)
	}
	fit(6)
	require.GreaterOrEqual(t, rows(), 6)
	require.LessOrEqual(t, rows(), 7)
	rootHeight := sizing.bounds.H
	fit(15)
	require.Greater(t, sizing.bounds.H, rootHeight, "a longer submenu can grow")
	require.GreaterOrEqual(t, rows(), 15)
	fit(5)
	require.Less(t, sizing.bounds.H, rootHeight, "returning to short content shrinks")
	fit(100)
	require.Equal(t, maximum.H, sizing.bounds.H, "large menus page at the height cap")
}

func TestPopupSizingWaitsForPTYAndDoesNotChaseRounding(t *testing.T) {
	sizing := newPopupSizing(rex.Bounds{X: .1, Y: .4, W: .8, H: .2})
	_, changed := sizing.fit(6, 120, 20)
	require.True(t, changed)
	for range 20 {
		_, changed = sizing.fit(6, 120, 20)
		require.False(t, changed, "no repeated mutation while PTY acknowledgement is pending")
	}
	_, changed = sizing.fit(6, 120, 6)
	require.False(t, changed)
	_, changed = sizing.fit(6, 120, 7)
	require.False(t, changed, "one row of cell rounding is acceptable")
	_, changed = sizing.fit(6, 120, 0)
	require.False(t, changed, "zero-size transitions cannot supply geometry")
}

func TestPopupSizingReflowsAfterScreenResize(t *testing.T) {
	sizing := newPopupSizing(rex.Bounds{X: .1, Y: .4, W: .8, H: .2})
	_, changed := sizing.fit(6, 120, 20)
	require.True(t, changed)
	_, changed = sizing.fit(6, 120, 6)
	require.False(t, changed)
	_, changed = sizing.fit(12, 60, 6)
	require.True(t, changed, "narrower screen and more content rows can grow")
	require.LessOrEqual(t, sizing.bounds.H, .2)
}
