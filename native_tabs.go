package whichkey

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"temp/rex"
)

// nativeNumberedTab compiles parent-retained t1..t9 intent, not child-provided
// action/args. The caller has cleaned both UIs, validated the original completion
// and checked the chosen app's actual session.select registration.
func nativeNumberedTab(ctx context.Context, client *rex.Client, bound popupRequest, executable string, mapping Mapping) error {
	fields, err := completionObject(mapping.Args, "index")
	var index int
	if err != nil || json.Unmarshal(fields["index"], &index) != nil || index < 1 || index > 9 ||
		mapping.Action != "client.tab.goto" || mapping.Keys != "t"+strconv.Itoa(index) {
		return errors.New("whichkey: native numbered tab refused: invalid retained one-based tab intent")
	}

	session, err := client.Session(ctx, bound.SessionID)
	if err != nil {
		return errors.New("whichkey: native numbered tab refused: cannot inspect the captured origin session")
	}
	windows, err := session.Windows(ctx)
	if err != nil {
		return errors.New("whichkey: native numbered tab refused: cannot inspect origin session tab order")
	}
	if index > len(windows) {
		return errors.New("whichkey: native numbered tab refused: requested tab does not exist in the origin session")
	}
	// Resolve the position only once. Reordering cannot substitute another tab;
	// later validation follows this exact ID within its captured session.
	window := windows[index-1]
	unscoped, err := rex.New(rex.Options{Executable: executable, Server: bound.Server,
		Timeout: 3 * time.Second})
	if err != nil {
		return errors.New("whichkey: native numbered tab refused: cannot construct the captured endpoint client")
	}

	if _, err := bound.originBlock(ctx, client); err != nil {
		return errors.New("whichkey: native numbered tab refused: origin or chosen app connection changed after popup cleanup")
	}
	if _, err := window.Inspect(ctx); err != nil {
		return errors.New("whichkey: native numbered tab refused: resolved tab is missing or no longer in the origin session")
	}
	// Unscoped client.perform avoids ambient or injected context overrides. Only
	// these parent-resolved args select a destination; origin still authorizes it.
	if err := unscoped.Call(ctx, "client.perform", map[string]any{
		"client_id": bound.ClientID, "action": "session.select",
		"args":   map[string]string{"session_id": bound.SessionID, "window_id": window.ID},
		"origin": "cli", "timeout_ms": 3000,
	}, nil); err != nil {
		return errors.New("whichkey: native numbered tab failed or is uncertain; check app Remote Control and action compatibility; not retried")
	}
	return nil
}
