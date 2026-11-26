//go:build darwin || linux

package whichkey

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"temp/rex"

	"github.com/stretchr/testify/require"
)

func sizingPresented(context.Context, int, int) (bool, error) { return true, nil }

func TestPopupSizingWorkerCoalescesAndJoins(t *testing.T) {
	maximum := rex.Bounds{X: .1, Y: .4, W: .8, H: .2}
	var rows atomic.Int32
	rows.Store(20)
	calls := make(chan rex.Bounds, 8)
	job := startPopupSizing(context.Background(), maximum,
		func() (int, int, error) { return 120, int(rows.Load()), nil },
		sizingPresented,
		func(_ context.Context, expected, bounds rex.Bounds) error {
			calls <- bounds
			rows.Store(6)
			return nil
		})
	t.Cleanup(func() { require.NoError(t, job.stop()) })
	for range 100 {
		job.request(50)
	}
	job.request(6)
	select {
	case bounds := <-calls:
		require.Less(t, bounds.H, maximum.H, "latest content wins")
	case <-time.After(2 * time.Second):
		t.Fatal("no resize")
	}
	require.NoError(t, job.stop())
	select {
	case <-job.done:
	default:
		t.Fatal("worker survived stop")
	}
}

func TestPopupSizingWorkerPreservesUncertainResizeOnStop(t *testing.T) {
	started := make(chan struct{})
	uncertain := errors.New("uncertain resize")
	job := startPopupSizing(context.Background(), rex.Bounds{X: .1, Y: .4, W: .8, H: .2},
		func() (int, int, error) { return 120, 20, nil },
		sizingPresented,
		func(ctx context.Context, _, _ rex.Bounds) error {
			close(started)
			<-ctx.Done()
			return uncertain
		})
	t.Cleanup(func() { _ = job.stop() })
	job.request(6)
	select {
	case <-started:
	case <-time.After(2 * time.Second):
		t.Fatal("resize did not start")
	}
	require.ErrorIs(t, job.stop(), uncertain, "a selected effect must not hide an in-flight resize failure")
}

func TestPopupSizingWorkerWaitsForPresentation(t *testing.T) {
	var ready atomic.Bool
	checked := make(chan struct{}, 8)
	resized := make(chan struct{}, 8)
	job := startPopupSizing(context.Background(), rex.Bounds{X: .1, Y: .4, W: .8, H: .2},
		func() (int, int, error) { return 80, 24, nil },
		func(context.Context, int, int) (bool, error) {
			checked <- struct{}{}
			return ready.Load(), nil
		},
		func(context.Context, rex.Bounds, rex.Bounds) error {
			resized <- struct{}{}
			return nil
		})
	t.Cleanup(func() { require.NoError(t, job.stop()) })
	job.request(6)
	select {
	case <-checked:
	case <-time.After(2 * time.Second):
		t.Fatal("presentation was not checked")
	}
	select {
	case <-resized:
		t.Fatal("bootstrap grid must not resize the native layer")
	default:
	}
	ready.Store(true)
	select {
	case <-resized:
	case <-time.After(2 * time.Second):
		t.Fatal("presented grid did not enable fitting")
	}
}

func TestPopupSizingWorkerReportsSizeFailure(t *testing.T) {
	failure := errors.New("terminal unavailable")
	job := startPopupSizing(context.Background(), rex.Bounds{X: .1, Y: .4, W: .8, H: .2},
		func() (int, int, error) { return 0, 0, failure },
		sizingPresented,
		func(context.Context, rex.Bounds, rex.Bounds) error {
			t.Error("must not resize without dimensions")
			return nil
		})
	job.request(6)
	select {
	case err := <-job.errors:
		require.ErrorIs(t, err, failure)
	case <-time.After(2 * time.Second):
		t.Fatal("missing sizing failure")
	}
	require.NoError(t, job.stop())
}
