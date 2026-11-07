package whichkey

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
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

type menuProviderUpdates struct {
	base       Config
	invocation string
	wanted     [4]*renameDemand
	sent       [4]bool
	ready      [4]*providerUpdate
}

// Called only after the first successful draw, or after explicit prefix entry.
// Hidden markers remain callable but never trigger the sessions background load.
func (m *Menu) demandProviders(firstDraw bool) {
	if m.providers == nil || m.Done() {
		return
	}
	for _, marker := range m.providers.base.Spec {
		kind := providerKind(marker)
		if !marker.Enabled() || kind < 0 || m.providers.wanted[kind] != nil {
			continue
		}
		background := firstDraw && kind == 0 && !marker.Hidden
		entered := !firstDraw && (m.Prefix() == marker.Keys || strings.HasPrefix(m.Prefix(), marker.Keys))
		if background || entered {
			m.providers.wanted[kind] = &renameDemand{Invocation: m.providers.invocation, Kind: providerKinds[kind], Load: 1, Keys: marker.Keys}
		}
	}
}

func (m *Menu) pollUpdates(transport terminalRenameTransport, localJobs *metadataJobs) (bool, error) {
	if localJobs != nil {
		return m.pollLocalProviders(localJobs)
	}
	renameChanged, err := m.pollRenameUpdates(transport)
	if err != nil {
		return false, err
	}
	providerChanged, err := m.pollProviderUpdates(transport)
	return renameChanged || providerChanged, err
}

func (m *Menu) pollLocalProviders(jobs *metadataJobs) (bool, error) {
	for kind, demand := range m.providers.wanted {
		jobs.wanted[kind] = demand != nil
	}
	changed := false
	select {
	case result := <-jobs.results:
		jobs.active--
		if err := jobs.ctx.Err(); err != nil {
			return false, err
		}
		demand := m.providers.wanted[result.kind]
		if demand == nil {
			return false, errRuntimeStartup
		}
		update, _, err := buildProviderUpdate(m.providers.base, *demand, result)
		if err != nil {
			return false, err
		}
		if err := m.applyProviderUpdate(result.kind, update); err != nil {
			return false, err
		}
		changed = true
	default:
	}
	jobs.start()
	return changed, nil
}

func (m *Menu) pollProviderUpdates(transport terminalRenameTransport) (bool, error) {
	changed := false
	for kind, demand := range m.providers.wanted {
		if demand == nil || m.providers.ready[kind] != nil {
			continue
		}
		if !m.providers.sent[kind] {
			data, err := json.Marshal(demand)
			if err != nil {
				return false, err
			}
			_, sent, err := transferRenameSlot(transport.root, providerDemandFiles[kind], data)
			if err != nil {
				return false, err
			}
			m.providers.sent[kind] = sent
			if !sent {
				continue
			}
		}
		data, ready, err := transferRenameSlot(transport.root, providerUpdateFiles[kind], nil)
		if err != nil {
			return false, err
		}
		if !ready {
			continue
		}
		var update providerUpdate
		if err := decodeRuntime(data, &update, popupMessageLimit); err != nil {
			return false, err
		}
		if err := m.applyProviderUpdate(kind, update); err != nil {
			return false, err
		}
		changed = true
	}
	return changed, nil
}

// Rebuild/validate off to the side, then rebind the existing interaction state
// in one UI-loop operation. Decoder state lives outside Menu and is untouched;
// form value/cursor/touched/error and confirmed text remain the same objects.
func (m *Menu) applyProviderUpdate(kind int, update providerUpdate) error {
	if m.Done() || m.providers.wanted[kind] == nil || m.providers.ready[kind] != nil {
		return errors.New("whichkey: provider update after terminal state")
	}
	if err := validateProviderUpdate(m.providers.base, *m.providers.wanted[kind], update); err != nil {
		return err
	}
	ready := m.providers.ready
	ready[kind] = &update
	config, err := providerConfig(m.providers.base, ready)
	if err != nil {
		return err
	}
	next, err := NewMenu(config.Spec)
	if err != nil {
		return err
	}
	if err := m.rebindMenu(next); err != nil {
		return err
	}
	m.providers.ready = ready
	return nil
}

func menuNodeAt(root *menuNode, keys string) *menuNode {
	if keys == "" {
		return root
	}
	tokens, err := ParseKeys(keys)
	if err != nil {
		return nil
	}
	node := root
	for _, token := range tokens {
		node = node.lookup[token]
		if node == nil {
			return nil
		}
	}
	return node
}

func (m *Menu) rebindMenu(next *Menu) error {
	current := menuNodeAt(next.current, m.Prefix())
	if current == nil {
		return errors.New("whichkey: provider update changed active prefix")
	}
	var pending, entry *menuNode
	if m.pending != nil {
		pending = menuNodeAt(next.current, m.pending.keys)
		if pending == nil {
			return errors.New("whichkey: provider update changed confirmation")
		}
	}
	if m.entry != nil {
		entry = menuNodeAt(next.current, m.entry.node.keys)
		if entry == nil {
			return errors.New("whichkey: provider update changed active form")
		}
	}
	m.current, m.pending = current, pending
	if m.entry != nil {
		m.entry.node = entry
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
