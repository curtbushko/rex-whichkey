package rex

import (
	"context"
	"encoding/json"
	"errors"
)

// Actions reads built-in metadata on exactly one app, never a server Lua
// registry. The shared frozen argument-array boundary owns every CLI process.
func (c *Client) Actions(ctx context.Context, clientID string) (map[string]string, error) {
	if c == nil || c.executable == "" || c.server == "" || c.timeout <= 0 || ctx == nil || !resourceID(clientID, "client") {
		return nil, errors.New("rex: exact app ID required")
	}
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	argv := []string{flagServer, c.server, flagNoAutostart, flagTimeout, c.timeout.String(), "--client", clientID, "actions", flagJSON}
	data, err := run(ctx, c.executable, argv, c.selection)
	if err != nil {
		return nil, err
	}
	return decodeActionMetadata(data)
}

func decodeActionMetadata(data []byte) (map[string]string, error) {
	if len(data) > 1<<20 {
		return nil, errCLICompatibility
	}
	fields, err := metadataObject(data)
	if err != nil {
		return nil, errCLICompatibility
	}
	var entries []json.RawMessage
	if json.Unmarshal(fields["actions"], &entries) != nil || entries == nil || len(entries) > 4096 {
		return nil, errCLICompatibility
	}
	actions := make(map[string]string, len(entries))
	for _, raw := range entries {
		entry, err := metadataObject(raw)
		if err != nil {
			return nil, errCLICompatibility
		}
		var name, kind string
		if json.Unmarshal(entry["name"], &name) != nil || json.Unmarshal(entry["kind"], &kind) != nil || !metadataText(name) || !metadataText(kind) {
			return nil, errCLICompatibility
		}
		if _, exists := actions[name]; exists {
			return nil, errCLICompatibility
		}
		actions[name] = kind
	}
	return actions, nil
}

// Perform sends one built-in client action. Callers retain and validate the app,
// generation, origin and destination before entering this effect boundary.
func (c *Client) Perform(ctx context.Context, clientID, action string, args any, sessionID, blockID string) error {
	if !resourceID(clientID, "client") || !metadataText(action) {
		return errors.New("rex: invalid client action target")
	}
	request := map[string]any{"client_id": clientID, "action": action, fieldArgs: args, "origin": "cli", "timeout_ms": 3000}
	if sessionID != "" || blockID != "" {
		if !resourceID(sessionID, "session") || !resourceID(blockID, "block") {
			return errors.New("rex: invalid client action context")
		}
		request["session_id"], request["block_id"] = sessionID, blockID
	}
	return c.Call(ctx, "client.perform", request, nil)
}
