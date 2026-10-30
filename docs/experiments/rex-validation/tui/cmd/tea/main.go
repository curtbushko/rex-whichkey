package main

import (
	"errors"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"example.com/whichkey-research/internal/probe"
	"golang.org/x/term"
)

type model struct{ menu probe.Menu }

func (m model) Init() tea.Cmd { return nil }
func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyPressMsg:
		m.menu.Key(msg.String())
	case tea.WindowSizeMsg:
		m.menu.Width = msg.Width
		m.menu.Height = msg.Height
	}
	if m.menu.Selection != "" {
		return m, tea.Quit
	}
	return m, nil
}
func (m model) View() tea.View {
	view := tea.NewView(m.menu.View())
	view.AltScreen = true
	return view
}
func run() (outcome error) {
	fd := int(os.Stdin.Fd())
	original, err := term.GetState(fd)
	if err != nil {
		return fmt.Errorf("capture terminal: %w", err)
	}
	// Keep lifecycle restoration owned by the short-lived application as well.
	// The comparison harness separately ignores only Darwin's transient PENDIN bit.
	defer func() { outcome = errors.Join(outcome, term.Restore(fd, original)) }()
	result, err := tea.NewProgram(model{}).Run()
	if err != nil {
		return fmt.Errorf("TUI: %w", err)
	}
	m, ok := result.(model)
	if !ok {
		return fmt.Errorf("unexpected research model %T", result)
	}
	return probe.Finish(m.menu)
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
