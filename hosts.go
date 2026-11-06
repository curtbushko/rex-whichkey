package whichkey

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"temp/rex"
)

// resolveHostMenu expands only enabled, explicit known-hosts markers. All markers
// share one bounded execution-host snapshot; neither --server nor popup/origin
// identities select that inventory. No marker means no executable resolution or
// inventory I/O, including for an explicitly empty/replaced user spec.
func resolveHostMenu(ctx context.Context, config Config, selection *rex.CLISelection) (Config, error) {
	if err := config.Validate(); err != nil {
		return Config{}, err
	}
	requested := false
	for _, mapping := range config.Spec {
		if mapping.Enabled() && mapping.Provider == "known-hosts" {
			requested = true
			break
		}
	}
	if !requested {
		return config, nil
	}
	// The endpoint remains bound inside Selection's API receipt; it does not
	// retarget rex hosts, whose execution-host/home inventory semantics remain.
	inventory, readErr := rex.Hosts(ctx, rex.InventoryOptions{Selection: selection, Timeout: 3 * time.Second})
	if err := ctx.Err(); err != nil {
		return Config{}, err // Invocation cancellation/expiry is not a listing warning.
	}
	var expanded []Mapping
	for _, mapping := range config.Spec {
		if !mapping.Enabled() || mapping.Provider != menuProviderHosts {
			expanded = append(expanded, mapping)
			continue
		}
		rows, expandErr := expandedHostRows(mapping, inventory, readErr)
		if expandErr != nil {
			return Config{}, expandErr
		}
		expanded = append(expanded, rows...)
		bounded := config
		bounded.Spec = expanded
		if err := menuSnapshotBound(bounded); err != nil {
			return Config{}, err
		}
	}
	config.Spec = expanded
	if err := config.Validate(); err != nil {
		return Config{}, fmt.Errorf("whichkey: expanded host menu: %w", err)
	}
	return config, nil
}

func expandedHostRows(mapping Mapping, inventory rex.HostInventory, readErr error) ([]Mapping, error) {
	keys, err := NormalizeKeys(mapping.Keys)
	if err != nil {
		return nil, err
	}
	mapping.Keys = keys // Standalone UI callers may not have a canonical snapshot yet.
	var rows []Mapping
	if readErr == nil {
		var formatErr error
		rows, formatErr = HostMenuMappings(mapping.Keys, inventory)
		if formatErr != nil {
			return nil, formatErr
		}
	} else {
		budget := menuKeyBudget{rowsLeft: 1024, bytesLeft: menuConfigLimit}
		for range 2 {
			if err := budget.reserve(len(mapping.Keys), 1); err != nil {
				return nil, err
			}
		}
		label := "Warning: inventory unavailable (read failed, timed out or invalid); not retried"
		scope := "Inventory: execution host and its configured Rex home; reachability unknown"
		rows = []Mapping{mapping, {Keys: mapping.Keys + "i", Group: &scope}, {Keys: mapping.Keys + "w", Group: &label}}
	}
	root := mapping
	root.Provider = ""
	rows[0] = root
	for index := 1; index < len(rows); index++ {
		rows[index].Hidden = mapping.Hidden
	}
	if err := ValidateMenuProviderChildren(mapping, rows[1:]); err != nil {
		return nil, err
	}
	return rows, nil
}

// Reserve exact full-key lengths before formatting any host rows. Prefixes are
// canonical bytes; all generated suffix tokens are single-byte ASCII. Include
// the root as in the host formatter's original row/key bound.
func hostMenuKeyBudget(prefixBytes int, inventory rex.HostInventory) error {
	budget := menuKeyBudget{rowsLeft: 1025, bytesLeft: menuConfigLimit}
	if err := budget.reserve(prefixBytes, 0); err != nil {
		return err
	}
	// Inventory, home, warnings and hosts headings.
	for range 4 {
		if err := budget.reserve(prefixBytes, 1); err != nil {
			return err
		}
	}
	if err := budget.reserveFields(prefixBytes+1, len(inventory.Home)); err != nil {
		return err
	}
	if len(inventory.Warnings) == 0 {
		if err := budget.reserve(prefixBytes, 2); err != nil {
			return err
		}
	}
	for index, warning := range inventory.Warnings {
		keysBytes, err := budget.reserveList(prefixBytes+1, index)
		if err != nil {
			return err
		}
		if err := budget.reserve(keysBytes, 1); err != nil {
			return err
		}
		if err := budget.reserveFields(keysBytes+1, len(warning)); err != nil {
			return err
		}
	}
	if len(inventory.Hosts) == 0 {
		return budget.reserve(prefixBytes, 2)
	}
	for index, entry := range inventory.Hosts {
		keysBytes, err := budget.reserveList(prefixBytes+1, index)
		if err != nil {
			return err
		}
		if err := budget.reserveHostDetails(keysBytes, entry); err != nil {
			return err
		}
	}
	return nil
}

func (budget *menuKeyBudget) reserveHostDetails(keysBytes int, entry rex.HostEntry) error {
	// ID, label, owner, held status and endpoints heading.
	for range 5 {
		if err := budget.reserve(keysBytes, 1); err != nil {
			return err
		}
	}
	for _, value := range []string{entry.Host.HostID, entry.Host.Source, entry.Host.AddedFrom, entry.Host.UpdatedAt} {
		if value != "" {
			if err := budget.reserve(keysBytes, 1); err != nil {
				return err
			}
		}
	}
	return budget.reserveFields(keysBytes+1, len(entry.Host.Endpoints))
}

// Fields and endpoints share the same paginated shape, including an empty row.
func (budget *menuKeyBudget) reserveFields(prefixBytes, count int) error {
	if count == 0 {
		return budget.reserve(prefixBytes, 1)
	}
	for index := range count {
		if _, err := budget.reserveList(prefixBytes, index); err != nil {
			return err
		}
	}
	return nil
}

// HostMenuMappings presents one effective rex.Hosts snapshot beneath prefix.
// resolveHostMenu inserts these non-executable records before NewMenu, rejecting
// collisions rather than overwriting user mappings. Every record is a group:
// opening a host only shows its details, never selects an action or emits a
// bridge completion.
//
// The caller obtains the snapshot on the execution/server host using rex.Hosts,
// not a retargeted Client or the graphical client's filesystem. This formatter
// performs no I/O, refresh, inventory merge, endpoint probe or host mutation.
func HostMenuMappings(prefix string, inventory rex.HostInventory) ([]Mapping, error) {
	prefix, err := NormalizeKeys(prefix)
	if err != nil {
		return nil, fmt.Errorf("whichkey: host menu prefix: %w", err)
	}
	if inventory.Home == nil {
		return nil, errors.New("whichkey: host menu requires effective inventory home data")
	}
	// Reject impossible prefixes/counts before even cloning the sort input.
	if len(prefix) > menuConfigLimit || len(inventory.Hosts) > 1024 {
		return nil, errors.New("whichkey: host provider exceeds row/key bounds; not truncated")
	}
	// IDs, not mutable/duplicate labels, determine display order. Do not merge
	// records or sort the caller's slice; Rex already owns inventory resolution.
	hosts := slices.Clone(inventory.Hosts)
	slices.SortStableFunc(hosts, func(a, b rex.HostEntry) int {
		if order := strings.Compare(a.Host.ID, b.Host.ID); order != 0 {
			return order
		}
		return strings.Compare(a.From, b.From)
	})
	inventory.Hosts = hosts // Admission must use the same pagination order as formatting.
	if err := hostMenuKeyBudget(len(prefix), inventory); err != nil {
		return nil, err
	}

	var mappings []Mapping
	add := func(keys, label string) {
		mappings = append(mappings, Mapping{Keys: keys, Group: &label})
	}
	add(prefix, "Rex-known hosts (listing only)")
	add(prefix+"i", "Inventory: execution host and its configured Rex home; not selected by --server")
	add(prefix+"h", "Home status (Rex-reported; not a reachability check)")
	if err := hostMenuFields(prefix+"h", inventory.Home, add); err != nil {
		return nil, fmt.Errorf("whichkey: host menu home: %w", err)
	}

	add(prefix+"w", fmt.Sprintf("Warnings (%d)", len(inventory.Warnings)))
	if len(inventory.Warnings) == 0 {
		add(prefix+"w0", "No inventory warnings")
	}
	for index, warning := range inventory.Warnings {
		if warning == nil {
			return nil, fmt.Errorf("whichkey: host menu warning %d is null", index+1)
		}
		keys := hostMenuListKey(prefix+"w", index, len(inventory.Warnings), "warnings", add)
		label, err := hostMenuJSON(warning)
		if err != nil {
			return nil, fmt.Errorf("whichkey: host menu warning %d: %w", index+1, err)
		}
		add(keys, "Warning: "+label)
		add(keys+"d", "Warning details")
		if err := hostMenuFields(keys+"d", warning, add); err != nil {
			return nil, fmt.Errorf("whichkey: host menu warning %d details: %w", index+1, err)
		}
	}

	add(prefix+"l", fmt.Sprintf("Known hosts (%d); open details only", len(hosts)))
	if len(hosts) == 0 {
		add(prefix+"l0", "No Rex-known hosts in this effective inventory")
	}
	for index, entry := range hosts {
		if strings.TrimSpace(entry.Host.ID) == "" || strings.TrimSpace(entry.Host.Label) == "" {
			return nil, fmt.Errorf("whichkey: host menu record %d needs an inventory ID and label", index+1)
		}
		keys := hostMenuListKey(prefix+"l", index, len(hosts), "hosts", add)
		label := hostMenuText(entry.Host.Label) + " [" + hostMenuText(entry.Host.ID) + "]"
		if entry.Held {
			label += " [held: not yet sent to home]"
		}
		add(keys, label)
		add(keys+"i", "Inventory ID: "+hostMenuText(entry.Host.ID))
		add(keys+"l", "Label: "+hostMenuText(entry.Host.Label))
		owner := "not separately reported by Rex (empty from)"
		if entry.From != "" {
			owner = hostMenuText(entry.From)
		}
		add(keys+"o", "Inventory owner: "+owner)
		held := "no"
		if entry.Held {
			held = "yes; change not yet sent to its owning home server"
		}
		add(keys+"h", "Held: "+held)
		if entry.Host.HostID != "" {
			add(keys+"m", "Machine/account ID (not inventory ID): "+hostMenuText(entry.Host.HostID))
		}
		if entry.Host.Source != "" {
			add(keys+"s", "Source: "+hostMenuText(entry.Host.Source))
		}
		if entry.Host.AddedFrom != "" {
			add(keys+"a", "Added from: "+hostMenuText(entry.Host.AddedFrom))
		}
		if entry.Host.UpdatedAt != "" {
			add(keys+"u", "Updated at: "+hostMenuText(entry.Host.UpdatedAt))
		}
		add(keys+"e", fmt.Sprintf("Endpoints (%d); not probed", len(entry.Host.Endpoints)))
		if len(entry.Host.Endpoints) == 0 {
			add(keys+"e0", "No endpoints reported")
		}
		for endpointIndex, endpoint := range entry.Host.Endpoints {
			endpointKeys := hostMenuListKey(keys+"e", endpointIndex, len(entry.Host.Endpoints), "endpoints", add)
			add(endpointKeys, hostMenuText(endpoint))
		}
	}
	if err := ValidateMappings(mappings); err != nil {
		return nil, fmt.Errorf("whichkey: host menu: %w", err)
	}
	return mappings, nil
}

// Home/warning fields are CLI-owned opaque data. Preserve all of them instead
// of guessing status enums, hiding unfamiliar warnings or treating cached data
// as a successful live-home check. Sorting also avoids map iteration ordering.
func hostMenuFields(prefix string, fields map[string]json.RawMessage, add func(string, string)) error {
	keys := make([]string, 0, len(fields))
	for key := range fields {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	if len(keys) == 0 {
		add(prefix+"0", "No additional metadata reported")
	}
	for index, key := range keys {
		text, err := hostMenuJSON(fields[key])
		if err != nil {
			return fmt.Errorf("field %q: %w", key, err)
		}
		fieldKeys := hostMenuListKey(prefix, index, len(keys), "metadata fields", add)
		add(fieldKeys, hostMenuText(key)+": "+text)
	}
	return nil
}

// Menu keys are single input tokens, not numbered-list strings. Explicit More
// groups keep item 10 from becoming a child of item 1. Within a snapshot each
// level offers 1-9 and n, leaving q and the usual navigation keys available.
func hostMenuListKey(prefix string, index, total int, noun string, add func(string, string)) string {
	list := prefix + strings.Repeat("n", index/9)
	if index > 0 && index%9 == 0 {
		add(list, fmt.Sprintf("More %s (%d-%d)", noun, index+1, total))
	}
	return list + strconv.Itoa(index%9+1)
}

func hostMenuJSON(value any) (string, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return "", fmt.Errorf("encode inventory metadata: %w", err)
	}
	return hostMenuText(string(data)), nil
}

// Unlike user mapping text, inventory text is external CLI data. Escape unsafe
// values visibly rather than dropping a whole listing or silently losing bytes.
func hostMenuText(text string) string {
	if !utf8.ValidString(text) || validateText(text) != nil {
		return strconv.QuoteToGraphic(text)
	}
	return text
}
