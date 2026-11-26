package rex

import "context"

// PresentedSize distinguishes a live client-sized terminal from Rex's initial
// 80x24 grid. That bootstrap grid says nothing about native window geometry.
// Only the exact popup is read; no origin or other client's size is substituted.
func (l Layer) PresentedSize(ctx context.Context, columns, rows int) (bool, error) {
	if err := l.valid(); err != nil {
		return false, err
	}
	var size struct {
		Columns int    `json:"columns"`
		Rows    int    `json:"rows"`
		Owner   string `json:"owner"`
		Desired []struct {
			ClientID string `json:"client_id"`
			Columns  int    `json:"columns"`
			Rows     int    `json:"rows"`
		} `json:"desired"`
	}
	if err := l.client.Call(ctx, "com.superlogical.terminal.size", map[string]any{
		fieldSession: l.client.session, fieldBlock: l.blockID, fieldArgs: map[string]any{},
	}, &size); err != nil {
		return false, err
	}
	if columns <= 0 || rows <= 0 || size.Columns != columns || size.Rows != rows || !resourceID(size.Owner, "client") {
		return false, nil
	}
	for _, desired := range size.Desired {
		if desired.ClientID == size.Owner && desired.Columns == columns && desired.Rows == rows {
			return true, nil
		}
	}
	return false, nil
}
