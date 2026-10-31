package rex

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

// Window is a tab in one stable session on one selected server. Fields describe
// a snapshot; operations use its privately captured IDs, not mutable labels or
// the process's current window, session, block or client.
type Window struct {
	ID string
	// Label is empty after unlabeled creation; Inspect reads the server's default.
	Label          string
	FocusedBlockID string
	Active         bool

	client   *Client
	windowID string
}

// sessionView is the documented session inspect shape. Layouts and block data
// are deliberately not modeled here: these operations need only tab metadata.
type sessionView struct {
	ID             string       `json:"session_id"`
	Label          string       `json:"label"`
	ActiveWindowID string       `json:"active_window_id"`
	Windows        []windowView `json:"windows"`
}

type windowView struct {
	ID             string `json:"window_id"`
	Label          string `json:"label"`
	FocusedBlockID string `json:"focused_block_id"`
}

func (v sessionView) validate() error {
	if !resourceID(v.ID, "session") {
		return errors.New("rex session inspect: missing or invalid session_id")
	}
	if v.Windows == nil {
		return errors.New("rex session inspect: missing or null windows array")
	}
	seen := make(map[string]bool, len(v.Windows))
	for _, window := range v.Windows {
		if !resourceID(window.ID, "window") || seen[window.ID] {
			return fmt.Errorf("rex session inspect: invalid or duplicate window_id %q", window.ID)
		}
		if window.FocusedBlockID != "" && !resourceID(window.FocusedBlockID, "block") {
			return fmt.Errorf("rex session inspect: invalid focused_block_id %q", window.FocusedBlockID)
		}
		seen[window.ID] = true
	}
	if v.ActiveWindowID != "" && !seen[v.ActiveWindowID] {
		return fmt.Errorf("rex session inspect: active_window_id %q is not in windows", v.ActiveWindowID)
	}
	return nil
}

// Windows inspects the session and returns tabs in server-provided tab order.
// An empty collection is returned as a non-nil empty slice.
func (s Session) Windows(ctx context.Context) ([]Window, error) {
	if err := s.valid(); err != nil {
		return nil, err
	}
	var view sessionView
	if err := s.client.resourceCommand(ctx, &view, "session", "inspect", "--json", "--", s.client.session); err != nil {
		return nil, err
	}
	if err := view.validate(); err != nil {
		return nil, err
	}
	if view.ID != s.client.session {
		return nil, fmt.Errorf("rex session inspect: requested %q, received %q", s.client.session, view.ID)
	}
	windows := make([]Window, 0, len(view.Windows))
	for _, item := range view.Windows {
		window := s.windowResource(item.ID, item.Label)
		window.FocusedBlockID = item.FocusedBlockID
		window.Active = item.ID == view.ActiveWindowID
		windows = append(windows, window)
	}
	return windows, nil
}

// Window inspects a window by its full stable ID within this session. It never
// derives the owning session from transient window membership or a local client.
func (s Session) Window(ctx context.Context, id string) (Window, error) {
	if !resourceID(id, "window") {
		return Window{}, errors.New("rex: a full window ID is required")
	}
	windows, err := s.Windows(ctx)
	if err != nil {
		return Window{}, err
	}
	for _, window := range windows {
		if window.ID == id {
			return window, nil
		}
	}
	return Window{}, fmt.Errorf("rex: window %q not found in session %q", id, s.client.session)
}

// NewWindow creates a default shell tab, focuses it in this server session, then
// applies label when nonempty. It does not ask a client app to display the session.
// An error may follow successful creation; a nonzero returned Window retains its
// identity. Do not blindly retry creation. Inspect refreshes its metadata.
func (s Session) NewWindow(ctx context.Context, label string) (Window, error) {
	if err := s.valid(); err != nil {
		return Window{}, err
	}
	if strings.ContainsRune(label, '\x00') {
		return Window{}, errors.New("rex: window label must not contain NUL")
	}
	var result struct {
		ID       string   `json:"window_id"`
		BlockIDs []string `json:"block_ids"`
	}
	// As with session creation, do not put a label after the creation command's
	// -- separator: Rex would treat it as the program to run.
	if err := s.client.resourceCommand(ctx, &result, "window", "new", "--session", s.client.session, "--json"); err != nil {
		return Window{}, err
	}
	if !resourceID(result.ID, "window") {
		return Window{}, errors.New("rex window new: missing or invalid window_id (creation may have succeeded)")
	}
	window := s.windowResource(result.ID, "")
	window.Active = true
	// Preserve the created window's identity even when other response data is
	// malformed, so the caller can inspect it instead of creating another tab.
	if len(result.BlockIDs) != 1 || !resourceID(result.BlockIDs[0], "block") {
		return window, errors.New("rex window new: invalid block_ids (creation may have succeeded)")
	}
	window.FocusedBlockID = result.BlockIDs[0]
	if label != "" {
		if err := window.Rename(ctx, label); err != nil {
			return window, fmt.Errorf("rex: created window %q but could not label it: %w", result.ID, err)
		}
		window.Label = label
	}
	return window, nil
}

// NextWindow focuses the next server tab, optionally wrapping at the end.
func (s Session) NextWindow(ctx context.Context, wrap bool) error {
	return s.focusAdjacentWindow(ctx, "next", wrap)
}

// PreviousWindow focuses the previous server tab, optionally wrapping at the end.
func (s Session) PreviousWindow(ctx context.Context, wrap bool) error {
	return s.focusAdjacentWindow(ctx, "prev", wrap)
}

func (s Session) focusAdjacentWindow(ctx context.Context, direction string, wrap bool) error {
	if err := s.valid(); err != nil {
		return err
	}
	args := []string{"window", "focus", "--session", s.client.session}
	if wrap {
		args = append(args, "--wrap")
	}
	args = append(args, "--", direction)
	return s.client.resourceCommand(ctx, nil, args...)
}

// Inspect refreshes a window using its stable session and window IDs.
func (w Window) Inspect(ctx context.Context) (Window, error) {
	if err := w.valid(); err != nil {
		return Window{}, err
	}
	return w.Session().Window(ctx, w.windowID)
}

// Session returns the captured owning session. Its label is not cached here.
func (w Window) Session() Session {
	if w.client == nil {
		return Session{}
	}
	return w.client.sessionResource(w.client.session, "")
}

// Server returns the captured server selector, not a client-local address.
func (w Window) Server() string { return w.Session().Server() }

// Focus changes server session focus; it does not switch a client app's session.
func (w Window) Focus(ctx context.Context) error {
	if err := w.valid(); err != nil {
		return err
	}
	return w.client.resourceCommand(ctx, nil, "window", "focus", "--session", w.client.session, "--", w.windowID)
}

// Rename changes this tab's label. Inspect refreshes the snapshot afterward.
func (w Window) Rename(ctx context.Context, label string) error {
	if err := w.valid(); err != nil {
		return err
	}
	if strings.ContainsRune(label, '\x00') {
		return errors.New("rex: window label must not contain NUL")
	}
	return w.client.resourceCommand(ctx, nil, "window", "rename", "--session", w.client.session, "--", w.windowID, label)
}

// Close explicitly closes this window and every block in it. It never performs
// broad cleanup of other windows or any automatic cleanup on process exit.
func (w Window) Close(ctx context.Context) error {
	if err := w.valid(); err != nil {
		return err
	}
	return w.client.resourceCommand(ctx, nil, "window", "close", "--session", w.client.session, "--", w.windowID)
}

func (s Session) windowResource(id, label string) Window {
	return Window{ID: id, Label: label, client: s.client, windowID: id}
}

func (w Window) valid() error {
	if err := w.Session().valid(); err != nil {
		return err
	}
	if !resourceID(w.windowID, "window") {
		return errors.New("rex: window must be obtained from a Session")
	}
	return nil
}
