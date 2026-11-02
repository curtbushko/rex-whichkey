package rex

import (
	"context"
	"fmt"
	"time"
)

// TerminalSize includes the applied size, which may still be the server's
// unmeasured starting size. Owner and ReportedAt are populated only when exactly
// one current owner's desired report matches the applied positive rows/columns.
// They establish measurement provenance, never app-action targeting authority.
// Cell dimensions are optional for geometry consumers; no font is inferred.
type TerminalSize struct {
	Columns      uint32    `json:"columns"`
	Rows         uint32    `json:"rows"`
	CellWidthPx  uint32    `json:"cell_width_px"`
	CellHeightPx uint32    `json:"cell_height_px"`
	Owner        string    `json:"-"`
	ReportedAt   time.Time `json:"-"`
}

// TerminalSize reads only the captured block in the captured session. It neither
// selects a graphical app nor claims terminal ownership or changes its size.
func (b Block) TerminalSize(ctx context.Context) (TerminalSize, error) {
	if err := b.valid(); err != nil {
		return TerminalSize{}, err
	}
	// The captured terminal size schema documents owner plus desired entries
	// (client_id, columns, rows, updated_at). Nonzero defaults alone are not a
	// viewer measurement, nor is another client's pending desired size.
	var report struct {
		TerminalSize
		Owner   *string `json:"owner"`
		Desired []struct {
			ClientID  *string `json:"client_id"`
			Columns   uint32  `json:"columns"`
			Rows      uint32  `json:"rows"`
			UpdatedAt string  `json:"updated_at"`
		} `json:"desired"`
	}
	if err := b.client.resourceCommand(ctx, &report, "block", "call", "--session", b.client.session,
		"--block", b.blockID, "--", "size"); err != nil {
		return TerminalSize{}, fmt.Errorf("rex: read owned terminal size: %w", err)
	}
	size := report.TerminalSize
	if report.Owner == nil || !resourceID(*report.Owner, "client") || size.Columns == 0 || size.Rows == 0 ||
		size.Columns > 65535 || size.Rows > 65535 {
		return size, nil
	}
	matches := 0
	var reportedAt time.Time
	for _, desired := range report.Desired {
		if desired.ClientID == nil || *desired.ClientID != *report.Owner {
			continue
		}
		matches++
		stamp, err := time.Parse(time.RFC3339Nano, desired.UpdatedAt)
		if err != nil || stamp.IsZero() || desired.Columns != size.Columns || desired.Rows != size.Rows {
			return size, nil
		}
		reportedAt = stamp
	}
	if matches == 1 {
		size.Owner, size.ReportedAt = *report.Owner, reportedAt
	}
	return size, nil
}
