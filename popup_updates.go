package whichkey

import (
	"context"
	"errors"
	"time"

	"temp/rex"
)

// Exactly two in-process form kinds, each resolved once per invocation.
const (
	renameStateReady       = "ready"
	renameStateUnavailable = "unavailable"
)

type renameDemand struct {
	Invocation string `json:"invocation"`
	Kind       string `json:"kind"`
	Load       int    `json:"load"`
	Keys       string `json:"keys"`
}

// The selection binding omits demand keys: the menu already retains full keys.
// In-process metadata supplies readiness, never action or target authority.
type renameBinding struct {
	Invocation string `json:"invocation"`
	Kind       string `json:"kind"`
	Load       int    `json:"load"`
}

func (d renameDemand) binding() renameBinding {
	return renameBinding{Invocation: d.Invocation, Kind: d.Kind, Load: d.Load}
}

type renameUpdate struct {
	Demand  renameDemand `json:"demand"`
	State   string       `json:"state"`
	Initial string       `json:"initial"`
	Problem string       `json:"problem"`
}

func renameKind(mapping Mapping) int {
	if inlineRenameMapping(mapping) {
		switch mapping.Action {
		case menuTabRename:
			return 0
		case menuSessionRename:
			return 1
		}
	}
	return -1
}

func renameKindName(kind int) string {
	if kind == 0 {
		return "tab"
	}
	return "session"
}

type renameLabels struct {
	initial [2]string
	problem [2]string
	err     error
}

// This worker receives no mutable config/menu/targets. It returns value-only
// presentation from an immutable view, never executes an effect or publishes UI.
func loadRenameLabels(ctx context.Context, client *rex.Client, request popupRequest) renameLabels {
	view := request.snapshot
	if view == nil {
		read, done := context.WithTimeout(ctx, 3*time.Second)
		snapshot, err := client.SessionSnapshot(read, request.SessionID)
		done()
		if err != nil {
			return renameLabels{err: errors.New("whichkey: captured rename context unavailable")}
		}
		view = &snapshot
	}
	if err := ctx.Err(); err != nil {
		return renameLabels{err: err}
	}
	block, err := view.Block(request.BlockID)
	if err != nil || view.Session().ID != request.SessionID || request.placement == nil ||
		block.WindowID != request.placement.WindowID || block.LayerID != request.placement.LayerID {
		return renameLabels{err: errors.New("whichkey: captured rename placement unavailable")}
	}
	window, err := view.Window(block.WindowID)
	if err != nil {
		return renameLabels{err: err}
	}
	labels := renameLabels{initial: [2]string{window.Label, view.Session().Label}}
	for kind, value := range labels.initial {
		if len(value) > renameTextLimit || validateText(value) != nil {
			labels.initial[kind], labels.problem[kind] = "", "Enter a new name"
		}
	}
	return labels
}
