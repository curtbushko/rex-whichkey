package whichkey

import (
	"context"
	"errors"
	"time"

	"temp/rex"
)

// Remote reads run sequentially inside the existing bounded sessions job. One
// slow host cannot consume its whole deadline; failures remain visible choices
// without actions, rather than hiding healthy local sessions or retrying reads.
func readSessionRows(ctx context.Context, client *rex.Client, currentID string) ([]providerRow, error) {
	sessions, err := client.Sessions(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := appendSessionRows(nil, sessions, "", currentID)
	if err != nil {
		return nil, err
	}
	inventory, err := rex.Hosts(ctx, rex.InventoryOptions{Selection: client.Selection(), Timeout: 3 * time.Second})
	if err != nil {
		return appendSessionNotice(rows, "Remote host inventory unavailable; not retried this invocation")
	}
	if len(inventory.Hosts) > maxSessionMenuEntries {
		return nil, errProviderOverflow
	}
	if len(inventory.Warnings) > 0 {
		warning, encodeErr := hostMenuJSON(inventory.Warnings)
		if encodeErr != nil {
			return nil, encodeErr
		}
		rows, err = appendSessionNotice(rows, "Remote host inventory warnings: "+warning)
		if err != nil {
			return nil, err
		}
	}
	seen := map[string]bool{client.Server(): true}
	for _, entry := range inventory.Hosts {
		host := entry.Host.Label
		remote, endpointErr := sessionHostClient(ctx, client, entry.Host)
		if endpointErr != nil {
			rows, err = appendSessionWarning(rows, host)
			if err != nil {
				return nil, err
			}
		} else if !seen[remote.Server()] {
			seen[remote.Server()] = true
			rows, err = appendRemoteSessions(ctx, rows, remote, host)
			if err != nil {
				return nil, err
			}
		}
		if len(rows) > maxSessionMenuEntries {
			return nil, errProviderOverflow
		}
	}
	return rows, nil
}

func sessionHostClient(ctx context.Context, origin *rex.Client, host rex.HostRecord) (*rex.Client, error) {
	if len(host.Endpoints) == 0 {
		return nil, errors.New("whichkey: remote host has no endpoint")
	}
	// Freeze the endpoint from this inventory, never the mutable host label.
	// No endpoint failover follows an unsuccessful read in this invocation.
	return rex.NewContext(ctx, rex.Options{Selection: origin.Selection(), Server: host.Endpoints[0], Timeout: 3 * time.Second})
}

func appendRemoteSessions(ctx context.Context, rows []providerRow, client *rex.Client, host string) ([]providerRow, error) {
	read, cancel := context.WithTimeout(ctx, 750*time.Millisecond)
	defer cancel()
	sessions, err := client.Sessions(read)
	if err != nil {
		return appendSessionWarning(rows, host)
	}
	if len(sessions) == 0 {
		rows, err = appendSessionNotice(rows, "No sessions in this snapshot")
		if err == nil {
			rows[len(rows)-1].host = host
		}
		return rows, err
	}
	return appendSessionRows(rows, sessions, host, "")
}

func appendSessionRows(rows []providerRow, sessions []rex.Session, host, currentID string) ([]providerRow, error) {
	if len(sessions) > maxSessionMenuEntries-len(rows) {
		return nil, errProviderOverflow
	}
	seen := make(map[string]bool, len(rows)+len(sessions))
	for _, row := range rows {
		seen[row.id] = true
	}
	for _, session := range sessions {
		// session.select addresses a UUID in the original app. If two servers
		// advertise the same UUID, no host-qualified action exists to disambiguate.
		if seen[session.ID] {
			return nil, errors.New("whichkey: session identity occurs on multiple hosts; not retargeted")
		}
		seen[session.ID] = true
		mapping, err := sessionRowMapping(session.ID, "")
		if err != nil {
			return nil, err
		}
		label := session.Label
		if session.ID == currentID {
			label = "[current] " + label
		}
		rows = append(rows, providerRow{id: session.ID, label: label, host: host, mapping: mapping, session: &session})
	}
	return rows, nil
}

func appendSessionWarning(rows []providerRow, host string) ([]providerRow, error) {
	rows, err := appendSessionNotice(rows, host+" unavailable; not retried this invocation")
	if err == nil {
		rows[len(rows)-1].host = host
	}
	return rows, err
}

func appendSessionNotice(rows []providerRow, label string) ([]providerRow, error) {
	if len(rows) >= maxSessionMenuEntries {
		return nil, errProviderOverflow
	}
	label = hostMenuText(label)
	return append(rows, providerRow{label: label, mapping: Mapping{Group: &label}}), nil
}
