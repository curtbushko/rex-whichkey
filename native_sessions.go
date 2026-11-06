package whichkey

import (
	"context"
	"errors"

	"temp/rex"
)

const maxSessionMenuEntries = 512

// validateSessionChoice checks the full retained destination, including optional
// window membership. No assumption about the configured key or label remains.
func validateSessionChoice(ctx context.Context, client *rex.Client, mapping Mapping) error {
	if mapping.Action != menuSessionSelect || mapping.CallbackRef != "" {
		return errors.New("whichkey: invalid retained session choice")
	}
	if err := validateMenuSessionArgs(mapping.Args); err != nil {
		return err
	}
	fields, err := completionObject(mapping.Args, menuSessionID, menuWindowID)
	if err != nil {
		return err
	}
	id, err := completionString(fields[menuSessionID])
	if err != nil {
		return err
	}
	session, err := client.Session(ctx, id)
	if err != nil {
		return errors.New("whichkey: chosen session is no longer available; not retargeted")
	}
	if raw, exists := fields[menuWindowID]; exists {
		windowID, idErr := completionString(raw)
		if idErr != nil {
			return idErr
		}
		if _, windowErr := session.Window(ctx, windowID); windowErr != nil {
			return errors.New("whichkey: chosen window left its retained session; not retargeted")
		}
	}
	return nil
}
