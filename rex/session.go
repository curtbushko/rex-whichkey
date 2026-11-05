package rex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Session is a session on one selected Rex server. Its fields are snapshots;
// operations retain the server and stable session ID captured when it was read.
// Creating a session does not select it in a client app.
type Session struct {
	ID string
	// Label is empty after unlabeled creation; Inspect reads the server's default.
	Label string

	client *Client
}

// Sessions lists sessions on the selected server, independently of c's session.
func (c *Client) Sessions(ctx context.Context) ([]Session, error) {
	if c == nil {
		return nil, errors.New("rex: client is required")
	}
	unscoped := *c
	unscoped.session = ""
	var result struct {
		Sessions []struct {
			ID    string `json:"session_id"`
			Label string `json:"label"`
		} `json:"sessions"`
	}
	if err := unscoped.Call(ctx, "session.list", nil, &result); err != nil {
		return nil, err
	}
	if result.Sessions == nil {
		return nil, errors.New("rex session.list: missing or null sessions array")
	}
	sessions := make([]Session, 0, len(result.Sessions))
	seen := make(map[string]bool, len(result.Sessions))
	for _, item := range result.Sessions {
		if !resourceID(item.ID, "session") || seen[item.ID] {
			return nil, fmt.Errorf("rex session.list: invalid or duplicate session_id %q", item.ID)
		}
		seen[item.ID] = true
		sessions = append(sessions, c.sessionResource(item.ID, item.Label))
	}
	return sessions, nil
}

// Session inspects target (an ID, label, position or ID suffix). An empty target
// uses c's explicitly configured session, never the process's REX_SESSION.
func (c *Client) Session(ctx context.Context, target string) (Session, error) {
	if c == nil {
		return Session{}, errors.New("rex: client is required")
	}
	if target == "" {
		target = c.session
	}
	if strings.TrimSpace(target) == "" || strings.ContainsRune(target, '\x00') {
		return Session{}, errors.New("rex: session target is required and must not contain NUL")
	}
	// Let Rex resolve human-readable targets once, then retain the returned ID.
	var view sessionView
	if err := c.resourceCommand(ctx, &view, "session", "inspect", "--json", "--", target); err != nil {
		return Session{}, err
	}
	if err := view.validate(); err != nil {
		return Session{}, err
	}
	if resourceID(target, "session") && view.ID != target {
		return Session{}, fmt.Errorf("rex session inspect: requested %q, received %q", target, view.ID)
	}
	return c.sessionResource(view.ID, view.Label), nil
}

// NewSession creates a session with its default shell window, then applies label
// when nonempty. The CLI constructs the terminal layout: no undocumented
// session.create fields are synthesized. An error may follow successful creation;
// a nonzero returned Session retains its identity. Do not blindly retry creation.
func (c *Client) NewSession(ctx context.Context, label string) (Session, error) {
	if strings.ContainsRune(label, '\x00') {
		return Session{}, errors.New("rex: session label must not contain NUL")
	}
	var result struct {
		ID             string `json:"session_id"`
		InitialWindows []struct {
			BlockIDs []string `json:"block_ids"`
		} `json:"initial_windows"`
	}
	// In creation commands, -- starts a terminal command, not a label. Rename
	// afterward so even a label beginning with '-' cannot become a flag or a
	// program to launch in the new terminal.
	if err := c.resourceCommand(ctx, &result, "session", "new", "--json"); err != nil {
		return Session{}, err
	}
	if !resourceID(result.ID, "session") {
		return Session{}, errors.New("rex session new: missing or invalid session_id (creation may have succeeded)")
	}
	session := c.sessionResource(result.ID, "")
	if len(result.InitialWindows) != 1 || len(result.InitialWindows[0].BlockIDs) != 1 ||
		!resourceID(result.InitialWindows[0].BlockIDs[0], "block") {
		return session, errors.New("rex session new: invalid initial_windows (creation may have succeeded)")
	}
	if label != "" {
		if err := session.Rename(ctx, label); err != nil {
			return session, fmt.Errorf("rex: created session %q but could not label it: %w", result.ID, err)
		}
		session.Label = label
	}
	return session, nil
}

// Inspect refreshes this session's metadata without following a mutable label.
func (s Session) Inspect(ctx context.Context) (Session, error) {
	if err := s.valid(); err != nil {
		return Session{}, err
	}
	return s.client.Session(ctx, s.client.session)
}

// Server returns the captured server selector, including for remote sessions.
func (s Session) Server() string {
	if s.client == nil {
		return ""
	}
	return s.client.server
}

// Rename changes the server resource's label, not the selected client session.
// Call Inspect to refresh this value's snapshot after renaming it.
func (s Session) Rename(ctx context.Context, label string) error {
	if err := s.valid(); err != nil {
		return err
	}
	if strings.ContainsRune(label, '\x00') {
		return errors.New("rex: session label must not contain NUL")
	}
	return s.client.resourceCommand(ctx, nil, "session", "rename", "--", s.client.session, label)
}

// Close explicitly kills this session and all of its windows and blocks.
// It is not automatic cleanup and must only be called for an intended target.
func (s Session) Close(ctx context.Context) error {
	if err := s.valid(); err != nil {
		return err
	}
	return s.client.resourceCommand(ctx, nil, "session", "kill", "--", s.client.session)
}

func (c *Client) sessionResource(id, label string) Session {
	scoped := *c
	scoped.session = id
	return Session{ID: id, Label: label, client: &scoped}
}

func (s Session) valid() error {
	if s.client == nil || !resourceID(s.client.session, "session") {
		return errors.New("rex: session must be obtained from a Client")
	}
	return nil
}

// resourceCommand uses documented resource CLI commands where the published
// static docs do not specify API request layouts. In particular, session new and
// window new own terminal layout construction. Keep their flags as argv values,
// not guessed JSON fields or shell strings. The common run helper still owns
// process execution and diagnostic/context errors.
func (c *Client) resourceCommand(ctx context.Context, result any, args ...string) error {
	if c == nil || c.executable == "" || c.server == "" || c.timeout <= 0 {
		return errors.New("rex: client must be constructed with New")
	}
	if ctx == nil {
		return errors.New("rex: resource command requires a context")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("rex resource command: %w", err)
	}
	argv := []string{"--server", c.server, "--autostart=false", "--timeout", c.timeout.String()}
	argv = append(argv, args...)
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	output, err := run(ctx, c.executable, argv, c.selection)
	if err != nil {
		return fmt.Errorf("rex resource command: %w", err)
	}
	if result != nil {
		if err := json.Unmarshal(output, result); err != nil {
			return fmt.Errorf("rex resource command decode (operation may have succeeded): %w", err)
		}
	}
	return nil
}

// resourceID accepts only full resource IDs, not labels or transient positions.
func resourceID(id, kind string) bool {
	prefix := kind + ":"
	if !strings.HasPrefix(id, prefix) {
		return false
	}
	uuid := strings.TrimPrefix(id, prefix)
	if len(uuid) != 36 {
		return false
	}
	for i, b := range []byte(uuid) {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			if b != '-' {
				return false
			}
			continue
		}
		if !(b >= '0' && b <= '9' || b >= 'a' && b <= 'f' || b >= 'A' && b <= 'F') {
			return false
		}
	}
	return true
}
