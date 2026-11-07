package whichkey

import (
	"errors"
	"strings"
)

type menuRenameUpdates struct {
	wanted [2]*renameDemand
	ready  [2]*renameUpdate
}

func (u *menuRenameUpdates) demand(node *menuNode) {
	kind := renameKind(node.mapping)
	if kind < 0 || u.wanted[kind] != nil {
		return
	}
	u.wanted[kind] = &renameDemand{Kind: renameKindName(kind), Load: 1, Keys: node.keys}
}

func (m *Menu) renameReady(node *menuNode) bool {
	if m.updates == nil {
		return true // Standalone/historical presentation has no asynchronous rename context.
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
	base   Config
	wanted [4]*renameDemand
	ready  [4]*providerUpdate
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
			m.providers.wanted[kind] = &renameDemand{Kind: providerKinds[kind], Load: 1, Keys: marker.Keys}
		}
	}
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
			return false, errors.New("whichkey: unexpected local inventory result")
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
