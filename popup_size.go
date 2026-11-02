package whichkey

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"math"
	"os"
	"syscall"

	"golang.org/x/sys/unix"

	"temp/rex"
)

const (
	popupSizeLimit       = 512
	popupRowLimit        = 1 << 20
	popupSizeSettleLimit = 4 // Existing watcher ticks, not a new timer or a mutation retry.
)

// These hints grant no action or lifetime authority. The child reports the
// renderer's natural height even when its current terminal clips the menu.
type popupSizeHint struct {
	PreferredRows int `json:"preferred_rows"`
	Columns       int `json:"columns"`
	Rows          int `json:"rows"`
}

func (h popupSizeHint) valid() bool {
	return h.PreferredRows > 0 && h.PreferredRows <= popupRowLimit &&
		h.Columns > 0 && h.Columns <= 65535 && h.Rows > 0 && h.Rows <= 65535
}

func popupSizePublisher(root *os.Root) func(MenuFrame, int, int) error {
	var previous popupSizeHint
	return func(frame MenuFrame, columns, rows int) error {
		hint := popupSizeHint{min(frame.PreferredHeight, popupRowLimit), columns, rows}
		if !hint.valid() || hint == previous {
			return nil // Retain the last usable hint through zero-size transitions.
		}
		data, err := json.Marshal(hint)
		if err != nil {
			return errors.New("whichkey: cannot encode popup size hint")
		}
		// A single terminal loop owns this exact temporary name. Exclusive
		// creation never truncates or follows an unexpected file or symlink.
		file, err := root.OpenFile("size.tmp", os.O_WRONLY|os.O_CREATE|os.O_EXCL|unix.O_NOFOLLOW, 0600)
		if err != nil {
			return errors.New("whichkey: cannot create popup size hint")
		}
		_, writeErr := file.Write(data)
		if err := errors.Join(writeErr, file.Close()); err != nil {
			return errors.New("whichkey: cannot write popup size hint")
		}
		if err := root.Rename("size.tmp", "size.json"); err != nil {
			return errors.New("whichkey: cannot publish popup size hint")
		}
		previous = hint
		return nil
	}
}

func readPopupSize(directory string) (hint popupSizeHint, outcome error) {
	root, err := openPopupDirectory(directory)
	if err != nil {
		return hint, err
	}
	defer func() { outcome = errors.Join(outcome, root.Close()) }()
	// NONBLOCK prevents a replaced FIFO/device from blocking before Stat.
	file, err := root.OpenFile("size.json", os.O_RDONLY|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if errors.Is(err, os.ErrNotExist) {
		return hint, nil // The first complete frame has not been published yet.
	}
	if err != nil {
		return hint, errors.New("whichkey: cannot open popup size hint")
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	info, err := file.Stat()
	if err != nil {
		return hint, errors.New("whichkey: cannot inspect popup size hint")
	}
	// An opened prior hint may have zero links after the child's next atomic
	// rename. It remains a valid immutable snapshot; hard-linked files do not.
	owner, ok := info.Sys().(*syscall.Stat_t)
	if !info.Mode().IsRegular() || info.Mode().Perm() != 0600 || info.Size() <= 0 || info.Size() > popupSizeLimit ||
		!ok || owner.Uid != uint32(os.Geteuid()) || owner.Nlink > 1 {
		return hint, errors.New("whichkey: invalid private popup size file")
	}
	data, err := io.ReadAll(io.LimitReader(file, popupSizeLimit+1))
	if err != nil || len(data) > popupSizeLimit || !json.Valid(data) {
		return hint, errors.New("whichkey: invalid popup size data")
	}
	if err := checkConfigJSON(json.NewDecoder(bytes.NewReader(data)), false); err != nil {
		return hint, errors.New("whichkey: invalid popup size object")
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&hint); err != nil || !hint.valid() {
		return popupSizeHint{}, errors.New("whichkey: invalid popup size dimensions")
	}
	return hint, nil
}

type popupSizeSample struct {
	bounds rex.Bounds
	size   rex.TerminalSize
	hint   popupSizeHint
}

// sameGeometry excludes the report timestamp: clients may refresh an unchanged
// desired size while their applied geometry is already stable.
func (s popupSizeSample) sameGeometry(other popupSizeSample) bool {
	return s.bounds == other.bounds && s.hint == other.hint && s.size.Owner == other.size.Owner &&
		s.size.Columns == other.size.Columns && s.size.Rows == other.size.Rows &&
		s.size.CellWidthPx == other.size.CellWidthPx && s.size.CellHeightPx == other.size.CellHeightPx
}

// popupSizer runs only in the parent's existing liveness watcher. Row feedback
// measures normalized units per actual cell, including fixed native insets
// without guessing font pixels or claiming to hide native controls.
type popupSizer struct {
	maximum    rex.Bounds
	position   string
	previous   popupSizeSample
	have       bool
	unmeasured bool
	unit       float64
	steps      int
	direction  int
	pending    bool
	requested  rex.Bounds
	settled    popupSizeSample
	settling   bool
	waits      int
}

func (s *popupSizer) refresh(ctx context.Context, directory string, layer rex.Layer) error {
	hint, err := readPopupSize(directory)
	if err != nil || !hint.valid() {
		s.settling = false
		return err
	}
	bounds, err := layer.Bounds(ctx)
	if err != nil {
		return err
	}
	size, err := layer.Block().TerminalSize(ctx)
	if err != nil {
		return err
	}
	// Headless/unclaimed terminals retain nonzero seeded sizes. Never learn a
	// row ratio or mutate from them. Losing a viewer pauses adaptation without
	// canceling native display or discarding a pending request's freshness
	// barrier. The next measured epoch must relearn its cell scale.
	if size.Owner == "" || size.ReportedAt.IsZero() {
		s.unmeasured, s.settling, s.waits = true, false, 0
		return nil
	}
	if s.pending {
		s.waits++
		if s.waits > popupSizeSettleLimit {
			return errors.New("whichkey: popup geometry did not settle; no selected action dispatched")
		}
	}
	// Wait for the child and native terminal to agree during layout/font
	// changes. Cell pixels can be absent; actual applied rows suffice.
	if size.Rows == 0 || size.Columns == 0 || uint32(hint.Columns) != size.Columns || uint32(hint.Rows) != size.Rows {
		s.settling = false
		return nil
	}
	current := popupSizeSample{bounds, size, hint}
	if s.pending {
		if bounds != s.requested || !size.ReportedAt.After(s.previous.size.ReportedAt) {
			s.settling = false
			return nil // Never compound a correction using the pre-resize report.
		}
		// A newer owner report plus two agreed geometry samples provides bounded
		// settling, not a native layer-application generation. A fresh same-row
		// result still advances correction across rounding/minimum-row limits.
		// Missing fresh reports fail closed rather than resending a mutation.
		if !s.settling || !current.sameGeometry(s.settled) {
			s.settled, s.settling = current, true
			return nil
		}
	} else if s.have && !s.unmeasured && current.sameGeometry(s.previous) {
		// Keep the freshness baseline current even when geometry is unchanged.
		s.previous = current
		return nil // Stable fitted/capped geometry needs no further mutation.
	}
	geometryChanged := !s.have || s.unmeasured || size.Owner != s.previous.size.Owner ||
		size.Columns != s.previous.size.Columns || size.CellWidthPx != s.previous.size.CellWidthPx ||
		size.CellHeightPx != s.previous.size.CellHeightPx ||
		!s.pending && bounds == s.previous.bounds && size.Rows != s.previous.size.Rows
	if geometryChanged {
		s.unit = bounds.H / float64(size.Rows)
	} else if s.pending && bounds.H != s.previous.bounds.H && size.Rows != s.previous.size.Rows {
		// The slope between two measured rectangles removes native insets.
		// Clamp noisy/rounded observations; direction reversal below damps
		// the next correction without permanently shrinking the cell estimate.
		unit := (bounds.H - s.previous.bounds.H) / (float64(size.Rows) - float64(s.previous.size.Rows))
		if unit > 0 && !math.IsInf(unit, 0) && !math.IsNaN(unit) {
			s.unit = max(s.unit/2, min(s.unit*2, unit))
		}
	}
	if geometryChanged || hint.PreferredRows != s.previous.hint.PreferredRows {
		s.steps, s.direction = 0, 0
	}
	s.previous, s.have, s.pending, s.unmeasured = current, true, false, false
	s.settling, s.waits = false, 0
	rows, preferred := float64(size.Rows), float64(hint.PreferredRows)
	// Half a row of headroom and a one-row deadband cover pixel flooring.
	if rows >= preferred && rows <= preferred+1 {
		return nil
	}
	direction := 1
	if rows > preferred {
		direction = -1
	}
	correction := (preferred + 0.5 - rows) * s.unit
	if s.direction != 0 && direction != s.direction {
		correction /= 2
	}
	height := min(s.maximum.H, max(min(s.unit, s.maximum.H), bounds.H+correction))
	if math.IsNaN(height) || math.IsInf(height, 0) || height <= 0 {
		return errors.New("whichkey: invalid measured popup geometry")
	}
	next := s.maximum
	next.H = height
	switch s.position {
	case "top":
		// Keep the original top edge.
	case "bottom":
		next.Y += s.maximum.H - height
	default: // Center and side popups keep their original vertical center.
		next.Y += (s.maximum.H - height) / 2
	}
	// Clamp floating point edge rounding before native bounds validation.
	next.Y = min(next.Y, 1-next.H)
	if math.Abs(next.H-bounds.H) < 1e-9 && math.Abs(next.Y-bounds.Y) < 1e-9 && next.X == bounds.X && next.W == bounds.W {
		return nil // Already capped, including menus that must paginate.
	}
	if s.steps >= 12 {
		return errors.New("whichkey: popup geometry did not converge; no selected action dispatched")
	}
	// Resize rechecks exact owned membership immediately before mutation.
	// An error cancels the lifecycle; no uncertain mutation is retried.
	if err := layer.Resize(ctx, next); err != nil {
		return err
	}
	s.steps++
	s.direction, s.pending, s.requested = direction, true, next
	return nil
}
