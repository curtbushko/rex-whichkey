//go:build darwin || linux

package whichkey

import (
	"context"
	"fmt"
	"time"

	"temp/rex"

	"golang.org/x/term"
)

// Only the UI goroutine publishes preferred rows. One worker owns native resize
// calls; it is canceled/joined before tty restoration and popup retirement.
type popupSizingJob struct {
	wanted chan popupSizingDemand
	errors chan error
	done   chan struct{}
	cancel context.CancelFunc
}

type popupSizingDemand struct {
	rows int
	root bool
}

type popupSizingSettings struct {
	maximumHeight float64
	save          func(float64)
}

func (ui *terminalMenu) startSizing(ctx context.Context, fd int) {
	if ui.direct == nil {
		return // Standalone menu mode must not mutate Rex resources.
	}
	bounds, available := ui.direct.popup.InitialBounds()
	if !available {
		return // Unknown geometry never becomes a guessed window rectangle.
	}
	settings := popupSizingSettings{maximumHeight: bounds.H}
	entry := ui.direct.geometry
	if geometryFraction(entry.MaximumHeight) && entry.MaximumHeight >= bounds.H {
		settings.maximumHeight = entry.MaximumHeight
		entry.Width = bounds.W
		if path := geometryCachePath(); path != "" {
			settings.save = func(height float64) {
				entry.Height = height
				_ = writeGeometryCache(path, entry) // Disposable hints never block menu use.
			}
		}
	}
	ui.sizing = startPopupSizing(ctx, bounds,
		func() (int, int, error) { return term.GetSize(fd) }, ui.direct.popup.PresentedSize, ui.direct.popup.Resize, settings)
	ui.requestSizing()
}

func startPopupSizing(parent context.Context, bounds rex.Bounds, size func() (int, int, error),
	presented func(context.Context, int, int) (bool, error),
	resize func(context.Context, rex.Bounds, rex.Bounds) error, options ...popupSizingSettings,
) *popupSizingJob {
	ctx, cancel := context.WithCancel(parent)
	settings := popupSizingSettings{maximumHeight: bounds.H}
	if len(options) > 0 {
		settings = options[0]
	}
	job := &popupSizingJob{wanted: make(chan popupSizingDemand, 1), errors: make(chan error, 1), done: make(chan struct{}), cancel: cancel}
	go func() {
		defer close(job.done)
		if err := job.run(ctx, bounds, size, presented, resize, settings); err != nil {
			job.errors <- err
		}
	}()
	return job
}

func (ui *terminalMenu) requestSizing() {
	if ui.sizing != nil {
		_, confirming := ui.menu.Confirmation()
		root := ui.menu.Prefix() == "" && ui.menu.entry == nil && !confirming
		ui.sizing.requestGeometry(ui.frame.PreferredHeight, root)
	}
}

func (job *popupSizingJob) request(rows int) { job.requestGeometry(rows, false) }

func (job *popupSizingJob) requestGeometry(rows int, root bool) {
	demand := popupSizingDemand{rows: rows, root: root}
	select {
	case job.wanted <- demand:
		return
	default:
	}
	select {
	case <-job.wanted:
	default:
	}
	job.wanted <- demand
}

func (job *popupSizingJob) stop() error {
	if job == nil {
		return nil
	}
	job.cancel()
	<-job.done
	select {
	case err := <-job.errors:
		return err
	default:
		return nil
	}
}

func (job *popupSizingJob) run(ctx context.Context, bounds rex.Bounds, size func() (int, int, error),
	presented func(context.Context, int, int) (bool, error),
	resize func(context.Context, rex.Bounds, rex.Bounds) error, settings popupSizingSettings,
) error {
	sizing := newPopupSizing(bounds)
	sizing.maximum.H = settings.maximumHeight
	recorder := popupGeometryRecorder{save: settings.save}
	// Coalesce rapid navigation/provider updates. Read our own PTY rather than
	// guessing native pixels or blocking keyboard handling on a CLI request.
	tick := time.NewTicker(75 * time.Millisecond)
	defer tick.Stop()
	demand := popupSizingDemand{}
	var nextPresentationCheck time.Time
	for {
		select {
		case <-ctx.Done():
			return nil
		case demand = <-job.wanted:
		case <-tick.C:
			columns, rows, err := size()
			if err != nil {
				return fmt.Errorf("whichkey: popup sizing terminal dimensions: %w", err)
			}
			next := sizing
			bounds, changed := next.fit(demand.rows, columns, rows)
			if !changed {
				sizing = next
				recorder.remember(ctx, sizing, demand, columns, rows, presented)
				continue
			}
			if time.Now().Before(nextPresentationCheck) {
				continue
			}
			nextPresentationCheck = time.Now().Add(250 * time.Millisecond)
			ready, presentationErr := presented(ctx, columns, rows)
			if presentationErr != nil {
				if ctx.Err() != nil {
					return nil // Canceling a read-only size query is not an uncertain resize.
				}
				return presentationErr
			}
			if !ready {
				continue
			}
			if resizeErr := resize(ctx, sizing.bounds, bounds); resizeErr != nil {
				return resizeErr
			}
			sizing = next
		}
	}
}
