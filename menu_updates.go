package whichkey

import (
	"encoding/json"
	"errors"
	"os"
)

// This context value carries an owned root, not a reader goroutine. Only the
// terminal UI loop calls poll, changes form readiness, or draws its result.
type terminalRenameKey struct{}

type terminalRenameTransport struct {
	root       *os.Root
	invocation string
}

type menuRenameUpdates struct {
	invocation string
	wanted     [2]*renameDemand
	sent       [2]bool
	ready      [2]*renameUpdate
}

func (u *menuRenameUpdates) demand(node *menuNode) {
	kind := renameKind(node.mapping)
	if kind < 0 || u.wanted[kind] != nil {
		return
	}
	u.wanted[kind] = &renameDemand{Invocation: u.invocation, Kind: renameKindName(kind), Load: 1, Keys: node.keys}
}

func (m *Menu) renameReady(node *menuNode) bool {
	if m.updates == nil {
		return true // Standalone/historical presentation has no private update path.
	}
	kind := renameKind(node.mapping)
	return kind >= 0 && m.updates.ready[kind] != nil && m.updates.ready[kind].State == renameStateReady
}

func (m *Menu) renameStatus() string {
	if m.entry == nil || m.updates == nil {
		return ""
	}
	kind := renameKind(m.entry.node.mapping)
	if kind < 0 {
		return ""
	}
	update := m.updates.ready[kind]
	if update == nil {
		return "Loading rename context; Enter disabled; Esc: back"
	}
	if update.State != renameStateReady {
		return update.Problem
	}
	return ""
}

// No update feeds or resets terminal decoders. The caller separately marks a
// form touched when partial input/paste is retained by those decoders.
func (m *Menu) pollRenameUpdates(transport terminalRenameTransport) (bool, error) {
	changed := false
	for kind, demand := range m.updates.wanted {
		if demand == nil || m.updates.ready[kind] != nil {
			continue
		}
		if !m.updates.sent[kind] {
			sent, err := sendRenameDemand(transport.root, kind, *demand)
			if err != nil {
				return false, err
			}
			m.updates.sent[kind] = sent
			if !sent {
				continue
			}
		}
		update, err := readRenameUpdate(transport.root, kind, *demand)
		if err != nil {
			return false, err
		}
		if update != nil {
			m.applyRenameUpdate(kind, *update)
			changed = true
		}
	}
	return changed, nil
}

func sendRenameDemand(root *os.Root, kind int, demand renameDemand) (bool, error) {
	data, err := json.Marshal(demand)
	if err != nil {
		return false, err
	}
	_, sent, err := transferRenameSlot(root, renameDemandFiles[kind], data)
	return sent, err
}

func readRenameUpdate(root *os.Root, kind int, demand renameDemand) (*renameUpdate, error) {
	data, ready, err := transferRenameSlot(root, renameUpdateFiles[kind], nil)
	if err != nil || !ready {
		return nil, err
	}
	var update renameUpdate
	if err := decodeRuntime(data, &update, popupMessageLimit); err != nil {
		return nil, err
	}
	if err := validateRenameUpdate(update, demand); err != nil {
		return nil, err
	}
	return &update, nil
}

func validateRenameUpdate(update renameUpdate, demand renameDemand) error {
	if update.Demand != demand || update.State != renameStateReady && update.State != renameStateUnavailable {
		return errors.New("whichkey: invalid or stale rename presentation update")
	}
	if len(update.Initial) > renameTextLimit || validateText(update.Initial) != nil ||
		len(update.Problem) > 128 || validateText(update.Problem) != nil {
		return errors.New("whichkey: invalid rename presentation text")
	}
	if update.State == renameStateUnavailable && (update.Initial != "" || update.Problem == "") {
		return errors.New("whichkey: invalid unavailable rename presentation")
	}
	return nil
}

func (m *Menu) applyRenameUpdate(kind int, update renameUpdate) {
	m.updates.ready[kind] = &update
	if m.entry != nil && renameKind(m.entry.node.mapping) == kind && !m.entry.touched &&
		m.entry.problem == "" && m.pendingText == nil && update.State == renameStateReady {
		m.entry.value, m.entry.cursor, m.entry.selected = update.Initial, len(update.Initial), true
		m.entry.problem = update.Problem
	}
}
