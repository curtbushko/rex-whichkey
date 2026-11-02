package whichkey

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"temp/rex"
)

// popupAppConnection is a private parent-owned target snapshot, not mapping or
// child authority. Label is display-only, decoded separately so malformed label
// metadata cannot change configured popupOriginConnection's eligibility rules.
type popupAppConnection struct {
	ID          string
	ConnectedAt string
	Label       json.RawMessage
}

// popupAppConnections shares the strict attachment validation used by configured
// binding and native dispatch. Only native action-time routing offers a chooser;
// popupOriginConnection still requires an exact selector or a single live app.
func popupAppConnections(ctx context.Context, client *rex.Client, sessionID string) ([]popupAppConnection, error) {
	var result struct {
		Clients json.RawMessage `json:"clients"`
	}
	if err := client.Call(ctx, "session.list_clients", map[string]string{"session_id": sessionID}, &result); err != nil {
		return nil, errors.New("whichkey: cannot inspect origin session clients; check the endpoint and captured REX_SESSION")
	}
	var connections []struct {
		ID          string `json:"client_id"`
		ConnectedAt string `json:"connected_at"`
		State       string `json:"connection_state"`
		Info        struct {
			Kind  string          `json:"kind"`
			Label json.RawMessage `json:"label"`
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
			ID: connection.ID, ConnectedAt: connection.ConnectedAt, Label: connection.Info.Label,
		})
	}
	return eligible, nil
}

// nativeActionTarget runs only after original selection/confirmation and cleanup.
// An empty returned identity means explicit cancellation, never implicit choice.
func nativeActionTarget(ctx context.Context, client *rex.Client, origin launchOrigin,
	executable, requested, description string,
) (popupAppConnection, error) {
	if err := ctx.Err(); err != nil {
		return popupAppConnection{}, err
	}
	if requested != "" {
		// Strict intent: a stale/non-app selector never falls back to a chooser,
		// even when another eligible app exists.
		id, generation, err := popupOriginConnection(ctx, client, origin.SessionID, requested)
		return popupAppConnection{ID: id, ConnectedAt: generation}, err
	}
	connections, err := popupAppConnections(ctx, client, origin.SessionID)
	if err != nil {
		return popupAppConnection{}, err
	}
	switch len(connections) {
	case 0:
		return popupAppConnection{}, errors.New("whichkey: origin session has no attached live graphical Rex app; connect the app to this session and relaunch")
	case 1:
		return connections[0], nil
	}

	config := DefaultConfig()
	config.Title = description + ": choose app"
	config.Icon = ""
	choices := make(map[string]popupAppConnection, len(connections))
	addGroup := func(keys, label string) {
		config.Spec = append(config.Spec, Mapping{Keys: keys, Group: &label})
	}
	for index, connection := range connections {
		// Reuse the canonical menu's 1-9/More groups, not ambiguous decimal
		// prefixes. Every choice shows a full ID, even when labels collide.
		keys := hostMenuListKey("", index, len(connections), "apps", addGroup)
		config.Spec = append(config.Spec, Mapping{
			Keys: keys, Desc: connection.ID + " (" + nativeTargetLabel(connection.Label) + ")",
			// This private leaf tag only makes a canonical selectable mapping.
			// It is never sent to Rex; finish returns keys into choices below.
			Action: "whichkey.select_app",
		})
		choices[keys] = connection
	}
	config, err = ResolveConfig(config, nil, nil)
	if err != nil {
		return popupAppConnection{}, errors.New("whichkey: cannot construct retained app choices")
	}
	// No short operation deadline surrounds this UI. The caller's outer native
	// signal context and this popup's FIFO still own its entire lifetime.
	message, err := nativePopupSelection(ctx, origin, executable, config)
	if err != nil {
		return popupAppConnection{}, err
	}
	if message.Kind == CompletionCancel {
		return popupAppConnection{}, nil
	}
	connection, present := choices[message.Selection.Keys]
	if !present {
		return popupAppConnection{}, errors.New("whichkey: selected app is missing from the retained target snapshot")
	}
	return connection, nil
}

// App-provided labels are hints, not identity. Keep display bounded and printable
// ASCII so control/bidi/format characters cannot spoof the adjacent full app ID.
func nativeTargetLabel(raw json.RawMessage) string {
	var label string
	if json.Unmarshal(raw, &label) != nil {
		return "Rex app"
	}
	var safe strings.Builder
	for _, char := range label {
		if safe.Len() >= 48 {
			break
		}
		if char < ' ' || char > '~' {
			char = '?'
		}
		safe.WriteRune(char)
	}
	if strings.TrimSpace(safe.String()) == "" {
		return "Rex app"
	}
	return safe.String()
}
