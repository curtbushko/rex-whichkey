package whichkey

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"temp/rex"
)

// popupAppConnection is a private parent-owned exact binding, not mapping or
// child authority. Labels and list order never supply routing identity.
type popupAppConnection struct {
	ID          string
	ConnectedAt string
}

// popupAppConnections supplies the shared eligibility/attachment validation for
// a binding already derived from the pane owner or an explicit app selector.
func popupAppConnections(ctx context.Context, client *rex.Client, sessionID string) ([]popupAppConnection, error) {
	var result struct {
		Clients json.RawMessage `json:"clients"`
	}
	if err := client.Call(ctx, "session.list_clients", map[string]string{menuSessionID: sessionID}, &result); err != nil {
		return nil, errors.New("whichkey: cannot inspect origin session clients; check the endpoint and captured REX_SESSION")
	}
	var connections []struct {
		ID          string `json:"client_id"`
		ConnectedAt string `json:"connected_at"`
		State       string `json:"connection_state"`
		Info        struct {
			Kind string `json:"kind"`
		} `json:"info"`
	}
	// Null is Rex's empty list. Missing/malformed data remains an inspection
	// error, and every ID (including non-app/lost records) must be unambiguous.
	if err := json.Unmarshal(result.Clients, &connections); err != nil {
		return nil, errors.New("whichkey: origin session returned malformed client attachments")
	}
	seen := make(map[string]bool, len(connections))
	var eligible []popupAppConnection
	for _, connection := range connections {
		identity := strings.ToLower(connection.ID)
		if !popupResourceID(connection.ID, "client") || seen[identity] {
			return nil, errors.New("whichkey: origin session returned malformed or duplicate client IDs")
		}
		seen[identity] = true
		if connection.Info.Kind != "app" || connection.State == "lost" {
			continue
		}
		if !popupConnectionGeneration(connection.ConnectedAt) {
			return nil, errors.New("whichkey: origin session returned a client without a valid connection generation")
		}
		eligible = append(eligible, popupAppConnection{
			ID: connection.ID, ConnectedAt: connection.ConnectedAt,
		})
	}
	return eligible, nil
}
