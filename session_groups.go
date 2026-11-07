package whichkey

import (
	"fmt"
	"strings"
)

const sessionHostShortcuts = "abcdefghijkmoprstuvwxyz123456789"

type sessionMenuHost struct {
	name, key string
	entries   []providerRow
}

// Host labels group presentation only. Each leaf keeps its own session and
// endpoint; grouping or shortcut changes never supply dispatch identity.
func groupSessionHosts(entries []providerRow) []sessionMenuHost {
	hosts := []sessionMenuHost{{name: "This Mac", key: "l"}}
	indices := map[string]int{"": 0}
	var notices []providerRow
	for _, entry := range entries {
		if entry.host == "" && entry.mapping.Action == "" {
			notices = append(notices, entry)
			continue
		}
		index, exists := indices[entry.host]
		if !exists {
			index = len(hosts)
			indices[entry.host] = index
			hosts = append(hosts, sessionMenuHost{name: hostMenuText(entry.host)})
		}
		hosts[index].entries = append(hosts[index].entries, entry)
	}
	if len(notices) > 0 {
		hosts = append(hosts, sessionMenuHost{name: "Inventory notices", key: "0", entries: notices})
	}
	return hosts
}

func sessionHostKey(name string, used map[byte]bool) string {
	for _, candidates := range []string{strings.ToLower(name), sessionHostShortcuts} {
		for index := range len(candidates) {
			key := candidates[index]
			if strings.ContainsRune(sessionHostShortcuts, rune(key)) && !used[key] {
				used[key] = true
				return string(key)
			}
		}
	}
	return ""
}

// Reserve l for this machine, q for dismissal, n for overflow and 0 for notices.
// Hosts beyond the mnemonic key space remain reachable in ordinary More groups.
func sessionHostPaths(hosts []sessionMenuHost) []Mapping {
	used := map[byte]bool{'l': true, 'q': true, 'n': true, '0': true}
	var overflow []int
	for index := range hosts {
		if hosts[index].key != "" {
			continue
		}
		hosts[index].key = sessionHostKey(hosts[index].name, used)
		if hosts[index].key == "" {
			overflow = append(overflow, index)
		}
	}
	var more []Mapping
	add := func(keys, label string) { more = append(more, Mapping{Keys: keys, Group: &label}) }
	if len(overflow) > 0 {
		add("n", "More hosts")
	}
	for index, host := range overflow {
		hosts[host].key = hostMenuListKey("n", index, len(overflow), "hosts", add)
	}
	return more
}

func reserveSessionHost(budget *menuKeyBudget, prefixBytes int, host sessionMenuHost) error {
	if err := budget.reserve(prefixBytes, len(host.key)); err != nil {
		return err
	}
	prefixBytes += len(host.key)
	if len(host.entries) == 0 {
		return budget.reserve(prefixBytes, 1)
	}
	for index := range host.entries {
		if _, err := budget.reserveList(prefixBytes, index); err != nil {
			return err
		}
	}
	return nil
}

func sessionHostLabel(host sessionMenuHost) string {
	if host.key == "0" {
		return host.name
	}
	count := 0
	for _, entry := range host.entries {
		if entry.mapping.Action == menuSessionSelect {
			count++
		}
	}
	return fmt.Sprintf("%s (%d)", host.name, count)
}

func sessionProviderRows(marker Mapping, entries []providerRow) ([]Mapping, error) {
	hosts := groupSessionHosts(entries)
	more := sessionHostPaths(hosts)
	budget := menuKeyBudget{rowsLeft: 1024, bytesLeft: menuConfigLimit}
	for _, group := range more {
		if err := budget.reserve(len(marker.Keys), len(group.Keys)); err != nil {
			return nil, err
		}
	}
	for _, host := range hosts {
		if err := reserveSessionHost(&budget, len(marker.Keys), host); err != nil {
			return nil, err
		}
	}
	rows := make([]Mapping, 0, 1024-budget.rowsLeft)
	for _, host := range hosts {
		label := sessionHostLabel(host)
		keys := marker.Keys + host.key
		rows = append(rows, Mapping{Keys: keys, Group: &label, Hidden: marker.Hidden, breadcrumbLabel: host.name})
		childMarker := marker
		childMarker.Keys = keys
		children, err := resourceProviderRows(childMarker, host.entries)
		if err != nil {
			return nil, err
		}
		rows = append(rows, children...)
	}
	for _, group := range more {
		group.Keys = marker.Keys + group.Keys
		group.Hidden = marker.Hidden
		rows = append(rows, group)
	}
	return rows, nil
}
