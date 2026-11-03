package whichkey

import "slices"

// Selection is a one-shot menu result, not authorization to execute an action.
// The dispatch owner must resolve it against the retained originating invocation,
// recheck conditions and confirmation, restore origin focus and execute in Rex.
// In particular, CallbackRef remains opaque data; Go never executes Lua callbacks.
// A zero Selection does not identify an action.
type Selection struct {
	mapping   Mapping
	confirmed bool
	textInput *string // Native form value, not a replacement Mapping or target.
}

// Mapping returns a detached copy of the selected, canonical-keyed record.
// An absent args value is returned as an empty JSON object for dispatch.
func (s Selection) Mapping() Mapping {
	mapping := cloneMenuMapping(s.mapping)
	if len(mapping.Args) == 0 {
		mapping.Args = []byte("{}")
	}
	return mapping
}

// Confirmed reports whether the user answered this mapping's prompt with y.
func (s Selection) Confirmed() bool { return s.confirmed }

func cloneMenuMapping(mapping Mapping) Mapping {
	mapping.Args = slices.Clone(mapping.Args)
	if mapping.Group != nil {
		group := *mapping.Group
		mapping.Group = &group
	}
	if mapping.Cond != nil {
		cond := *mapping.Cond
		mapping.Cond = &cond
	}
	return mapping
}
