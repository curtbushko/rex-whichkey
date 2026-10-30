// Package probe contains disposable research helpers, not a proposed application API.
package probe

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Menu struct {
	Page      int
	Width     int
	Height    int
	Selection string
	Failed    bool
}

func (m *Menu) Key(key string) {
	switch key {
	case "tab":
		m.Page++
	case "backspace":
		if m.Page > 0 {
			m.Page--
		}
	case "y":
		m.Selection = "choose"
	case "esc", "q", "ctrl+c":
		m.Selection = "cancel"
	case "e":
		m.Failed = true
		m.Selection = "cancel"
	}
}

func (m Menu) View() string {
	label := "界 e\u0301"
	if _, ascii := os.LookupEnv("WK_ASCII"); ascii {
		label = "ASCII"
	}
	return fmt.Sprintf("RESEARCH READY PAGE=%d SIZE=%dx%d %s\r\nTab page; Backspace back; y select; q/Esc cancel; e fail", m.Page, m.Width, m.Height, label)
}

func DispatchArgs(endpoint, session, block, id, selection string) ([]string, error) {
	payload, err := json.Marshal(struct {
		ID        string `json:"id"`
		Selection string `json:"selection"`
		Confirmed bool   `json:"confirmed"`
	}{id, selection, selection == "choose"})
	if err != nil {
		return nil, fmt.Errorf("marshal selection: %w", err)
	}
	return []string{"-S", endpoint, "--autostart=false", "do", "-s", session, "-b", block, "wk_probe_dispatch", "--args", string(payload), "--json"}, nil
}

func Finish(m Menu) error {
	if m.Failed {
		return fmt.Errorf("intentional research failure")
	}
	endpoint := os.Getenv("WK_CALLBACK_ENDPOINT")
	if endpoint == "" {
		_, err := fmt.Fprintln(os.Stderr, "RESEARCH RESULT "+m.Selection)
		return err
	}
	args, err := DispatchArgs(endpoint, os.Getenv("WK_ORIGIN_SESSION"), os.Getenv("WK_ORIGIN_BLOCK"), os.Getenv("WK_INVOCATION"), m.Selection)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	output, err := exec.CommandContext(ctx, os.Getenv("WK_REX"), args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("dispatch: %w: %s", err, output)
	}
	if err := os.WriteFile(os.Getenv("WK_CALLBACK_OUTPUT"), output, 0600); err != nil {
		return fmt.Errorf("write callback result: %w", err)
	}
	return nil
}
