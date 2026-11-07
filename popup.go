package whichkey

import (
	"context"
	"errors"
	"strings"

	"temp/rex"
)

// popupRequest carries immutable original resource and metadata context within
// this process. It contains no launch command, transport or mutable controller.
type popupRequest struct {
	Server      string
	SessionID   string
	BlockID     string
	ClientID    string
	ConnectedAt string
	placement   *rex.Block
	snapshot    *rex.SessionSnapshot
	selection   *rex.CLISelection
}

func popupResourceID(id, kind string) bool {
	value := strings.TrimPrefix(id, kind+":")
	if value == id || len(value) != 36 {
		return false
	}
	for i, char := range value {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			if char != '-' {
				return false
			}
		} else if !strings.ContainsRune("0123456789abcdefABCDEF", char) {
			return false
		}
	}
	return true
}

// popupOriginConnection uses Rex's native info.kind classification on the
// captured session's attachments: only app is eligible, never CLI/automation or
// an unknown kind. This keeps classification and membership in one response.
// Callers must supply the exact pane-owner or explicit binding. This validator
// never chooses a client from list order or prompts for one. Missing/unknown
// state is connected per Rex's contract; only lost is disconnected.
func popupOriginConnection(ctx context.Context, client *rex.Client, sessionID, requested string) (string, string, error) {
	if !popupResourceID(requested, "client") {
		return "", "", errors.New("whichkey: exact graphical app binding is missing or malformed")
	}
	connections, err := popupAppConnections(ctx, client, sessionID)
	if err != nil {
		return "", "", err
	}
	for _, connection := range connections {
		if requested == connection.ID {
			return connection.ID, connection.ConnectedAt, nil
		}
	}
	return "", "", errors.New("whichkey: bound client is not a live graphical app attached to the origin session")
}

// originBlock checks the captured app generation and exact resource origin.
// Resource-only display/restoration uses resourceOriginBlock instead.
func (r popupRequest) originBlock(ctx context.Context, client *rex.Client) (rex.Block, error) {
	_, generation, err := popupOriginConnection(ctx, client, r.SessionID, r.ClientID)
	if err != nil {
		return rex.Block{}, err
	}
	if generation != r.ConnectedAt {
		return rex.Block{}, errors.New("whichkey: origin client connection changed; relaunch from the original terminal")
	}
	block, err := r.resourceOriginBlock(ctx, client)
	if err != nil {
		return rex.Block{}, err
	}
	return block, nil
}

// resourceOriginBlock checks only the captured endpoint/session/block. Native
// display, focus and liveness do not grant authority to execute app actions.
// An executable app action additionally checks originBlock; absent app fields
// never select a replacement graphical target.
func (r popupRequest) resourceOriginBlock(ctx context.Context, client *rex.Client) (rex.Block, error) {
	session, err := client.Session(ctx, r.SessionID)
	if err != nil {
		return rex.Block{}, errors.New("whichkey: origin session is unavailable; relaunch inside the original Rex terminal")
	}
	block, err := session.Block(ctx, r.BlockID)
	if err != nil {
		return rex.Block{}, errors.New("whichkey: origin block is missing, detached or ambiguously placed in REX_SESSION; relaunch inside the original Rex terminal")
	}
	if r.placement != nil && (block.WindowID != r.placement.WindowID || block.LayerID != r.placement.LayerID) {
		return rex.Block{}, errors.New("whichkey: origin placement changed; not retargeted")
	}
	return block, nil
}

// The retained record, never the result's labels/action/args, defines execution.
func retainedNativeEvent(config Config, keys string, confirmed bool, text *string) (MenuEvent, error) {
	canonical, err := NormalizeKeys(keys)
	if err != nil || canonical != keys {
		return MenuEvent{}, errors.New("whichkey: popup result requires canonical keys")
	}
	for _, mapping := range config.Spec {
		if mapping.Keys != keys {
			continue
		}
		if !mapping.Enabled() || mapping.Action == "" || mapping.CallbackRef != "" || mapping.Confirm != confirmed {
			return MenuEvent{}, errors.New("whichkey: popup result does not match a retained Go action")
		}
		if inlineRenameMapping(mapping) != (text != nil) {
			return MenuEvent{}, errors.New("whichkey: popup text does not match retained rename intent")
		}
		if text != nil {
			if config.textEntry == nil {
				return MenuEvent{}, errors.New("whichkey: no inline rename form was retained")
			}
			if textErr := validateRenameText(*text); textErr != nil {
				return MenuEvent{}, textErr
			}
		}
		return MenuEvent{Selection: &Selection{mapping: cloneMenuMapping(mapping), confirmed: confirmed, textInput: text}}, nil
	}
	return MenuEvent{}, errors.New("whichkey: popup selection missing from retained snapshot")
}
