package whichkey

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"temp/rex"
)

const maxSessionMenuEntries = 512

// nativeSessionMenu expands only the built-in switch intent into a frozen,
// parent-retained canonical spec. Children return keys, never destination IDs.
// A failed inventory leaves an informative empty submenu, not guessed targets.
func nativeSessionMenu(parent context.Context, client *rex.Client, origin launchOrigin, config Config) (Config, string, error) {
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()
	sessions, inventoryErr := client.Sessions(ctx)
	currentName := ""
	for _, session := range sessions {
		if session.ID == origin.SessionID {
			currentName = session.Label
		}
	}
	title := "Switch session"
	var children []Mapping
	if inventoryErr != nil {
		title += " (list unavailable; reopen to retry)"
	} else if len(sessions) == 0 {
		title += " (no sessions)"
	} else if len(sessions) > maxSessionMenuEntries {
		title += " (too many sessions; list not truncated)"
	} else {
		shortIDs := make(map[string]int, len(sessions))
		for _, session := range sessions {
			shortIDs[sessionMenuID(session.ID)]++
		}
		addGroup := func(keys, label string) {
			children = append(children, Mapping{Keys: keys, Group: &label})
		}
		for index, session := range sessions {
			keys := hostMenuListKey("ws", index, len(sessions), "sessions", addGroup)
			label := []rune(hostMenuText(session.Label))
			if len(label) > 80 {
				label = append(label[:80], []rune("...")...)
			}
			name := string(label)
			if strings.TrimSpace(name) == "" {
				name = "Unnamed session"
			}
			shortID := sessionMenuID(session.ID)
			if shortIDs[shortID] != 1 {
				shortID = session.ID
			}
			marker := ""
			if session.ID == origin.SessionID {
				marker = "[current] "
			}
			// Identity first: clipping a long or duplicate label cannot hide
			// which resource this row represents. Full IDs stay in retained args.
			name = "[" + shortID + "] " + marker + name
			args, err := json.Marshal(map[string]string{"session_id": session.ID})
			if err != nil {
				return Config{}, "", errors.New("whichkey: cannot retain session choice")
			}
			children = append(children, Mapping{Keys: keys, Desc: name, Action: "session.select", Args: args})
		}
	}
	expanded := make([]Mapping, 0, len(config.Spec)+len(children))
	for _, mapping := range config.Spec {
		if mapping.Keys == "ws" && mapping.Action == "session.switch" && mapping.CallbackRef == "" && mapping.Enabled() && !mapping.Confirm {
			expanded = append(expanded, Mapping{Keys: "ws", Group: &title, Icon: mapping.Icon})
			expanded = append(expanded, children...)
		} else {
			expanded = append(expanded, mapping)
		}
	}
	config.Spec = expanded
	if err := config.Validate(); err != nil {
		return Config{}, "", err
	}
	return config, currentName, nil
}

func sessionMenuID(id string) string {
	if len(id) > 8 {
		return id[len(id)-8:]
	}
	return id
}

// validateSessionChoice inspects the full retained destination ID on the same
// server. Removed sessions refuse rather than resolving a label or new position.
func validateSessionChoice(ctx context.Context, client *rex.Client, mapping Mapping) error {
	if !strings.HasPrefix(mapping.Keys, "ws") || mapping.Action != "session.select" || mapping.CallbackRef != "" {
		return errors.New("whichkey: invalid retained session choice")
	}
	fields, err := completionObject(mapping.Args, "session_id")
	if err != nil {
		return err
	}
	id, err := completionString(fields["session_id"])
	if err != nil || !popupResourceID(id, "session") {
		return errors.New("whichkey: invalid retained destination session ID")
	}
	if _, err := client.Session(ctx, id); err != nil {
		return errors.New("whichkey: chosen session is no longer available; not retargeted")
	}
	return nil
}
