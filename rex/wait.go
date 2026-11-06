package rex

import (
	"context"
	"errors"
)

// WaitTerminal waits for one exact terminal without applying the short RPC
// deadline to its event stream. Cancellation still kills/joins the CLI process;
// --timeout continues to bound each connection/nonstreaming operation.
func (b Block) WaitTerminal(ctx context.Context) error {
	if err := b.valid(); err != nil {
		return err
	}
	if ctx == nil {
		return errors.New("rex: terminal wait requires a context")
	}
	c := b.client
	_, err := run(ctx, c.executable, []string{flagServer, c.server, flagNoAutostart, flagTimeout, c.timeout.String(), "wait", "--session", c.session, "--", b.blockID}, c.selection)
	return err
}
