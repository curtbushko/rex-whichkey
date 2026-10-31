package rex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// InventoryOptions selects the Rex executable on the host running this process.
// Unlike Options, it has no server or session target: rex hosts ignores --server
// and REX_SERVER when choosing the servers that own its effective inventory.
type InventoryOptions struct {
	// Executable follows the same execution-host resolution rules as Options.
	Executable string
	// Timeout bounds the CLI process. Zero uses ten seconds; an earlier context
	// deadline still takes precedence.
	Timeout time.Duration
}

// HostInventory is the effective result of rex hosts --json, not hosts.list.
// It belongs to the machine executing the CLI and its configured home, which
// need not be the originating graphical client's inventory. Rex performs the
// home/local merge and resolves duplicate labels; this package does not redo it.
type HostInventory struct {
	Home     InventoryHome      `json:"home"`
	Hosts    []HostEntry        `json:"hosts"`
	Warnings []InventoryWarning `json:"warnings"`
}

// InventoryHome retains the CLI's home status and cached-home metadata without
// inventing a server-API schema for this CLI-only record. Values remain JSON so
// timestamps and future status fields survive decoding without conversion or
// loss. A successful listing does not mean the home is currently reachable.
type InventoryHome map[string]json.RawMessage

// InventoryWarning retains a structured CLI warning, including its message and
// any ownership/source metadata. Warnings are data, not execution failures, and
// must be presented alongside the inventory rather than silently discarded.
type InventoryWarning map[string]json.RawMessage

// HostEntry preserves the nested record and its inventory ownership. From is
// the server the record came from; it can be empty when home is this machine.
// Held means the change has not yet been sent to its owning home server.
type HostEntry struct {
	Host HostRecord `json:"host"`
	From string     `json:"from"`
	Held bool       `json:"held"`
}

// HostRecord is a Rex-known host. ID identifies the inventory record (host:<uuid>);
// HostID, when known, identifies the machine/account reported by server.status.
// They are different identities. Label is a mutable display/CLI target name.
// Endpoints are retained in Rex's order, never probed by this package.
type HostRecord struct {
	ID        string   `json:"id"`
	Label     string   `json:"label"`
	Endpoints []string `json:"endpoints"`
	HostID    string   `json:"host_id,omitempty"`
	Source    string   `json:"source,omitempty"`
	AddedFrom string   `json:"added_from,omitempty"`
	UpdatedAt string   `json:"updated_at,omitempty"`
}

// Hosts reads the execution host's effective Rex inventory once. It is separate
// from Client.Call because a Client's fixed server/session cannot retarget this
// command. It does not add hosts, enumerate Tailscale, check endpoints, attach to
// a host, or modify configuration. Rex itself may contact its configured home
// and forward changes it previously held, as documented by rex hosts --help.
func Hosts(ctx context.Context, options InventoryOptions) (HostInventory, error) {
	if ctx == nil {
		return HostInventory{}, errors.New("rex hosts: a context is required")
	}
	if err := ctx.Err(); err != nil {
		return HostInventory{}, fmt.Errorf("rex hosts: %w", err)
	}
	if options.Timeout < 0 {
		return HostInventory{}, errors.New("rex hosts: timeout must not be negative")
	}
	if options.Timeout == 0 {
		options.Timeout = defaultTimeout
	}
	executable, err := resolveExecutable(options.Executable)
	if err != nil {
		return HostInventory{}, fmt.Errorf("rex hosts: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, options.Timeout)
	defer cancel()
	// No --server, session selector, or API call: these do not select the
	// effective inventory. Do not start a missing local server for a menu.
	output, err := run(ctx, executable, []string{
		"--autostart=false", "--timeout", options.Timeout.String(),
		"hosts", "--json",
	})
	if err != nil {
		return HostInventory{}, fmt.Errorf("rex hosts: %w", err)
	}
	inventory, err := decodeHostInventory(output)
	if err != nil {
		return HostInventory{}, fmt.Errorf("rex hosts decode inventory: %w", err)
	}
	return inventory, nil
}

func decodeHostInventory(data []byte) (HostInventory, error) {
	// Pointers distinguish an actual empty inventory from missing/null fields
	// and reject raw hosts.list responses, whose hosts are not nested entries.
	var result *struct {
		Home     InventoryHome      `json:"home"`
		Hosts    *[]HostEntry       `json:"hosts"`
		Warnings []InventoryWarning `json:"warnings"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return HostInventory{}, err
	}
	if result == nil || result.Home == nil || result.Hosts == nil {
		return HostInventory{}, errors.New("expected an effective inventory object with home and hosts")
	}
	for index, entry := range *result.Hosts {
		if strings.TrimSpace(entry.Host.ID) == "" || strings.TrimSpace(entry.Host.Label) == "" {
			return HostInventory{}, fmt.Errorf("hosts[%d].host must contain a nonempty id and label", index)
		}
	}
	for index, warning := range result.Warnings {
		if warning == nil {
			return HostInventory{}, fmt.Errorf("warnings[%d] must be an object, not null", index)
		}
	}
	return HostInventory{
		Home:     result.Home,
		Hosts:    *result.Hosts,
		Warnings: result.Warnings,
	}, nil
}
