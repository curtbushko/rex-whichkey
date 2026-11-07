package rex

import (
	"encoding/json"
	"errors"
)

// decodeBounds requires all four finite, positive-size window coordinates.
// Layer APIs also accept the older object representation; layout.rect uses an
// array. Missing or null coordinates must not turn into plausible zero values.
func decodeBounds(data json.RawMessage) (Bounds, error) {
	var coordinates []*float64
	if err := json.Unmarshal(data, &coordinates); err != nil {
		var object struct {
			X *float64 `json:"x"`
			Y *float64 `json:"y"`
			W *float64 `json:"w"`
			H *float64 `json:"h"`
		}
		if err := json.Unmarshal(data, &object); err != nil {
			return Bounds{}, errors.New("rex: malformed window rectangle")
		}
		coordinates = []*float64{object.X, object.Y, object.W, object.H}
	}
	if len(coordinates) != 4 {
		return Bounds{}, errors.New("rex: missing window rectangle")
	}
	for _, value := range coordinates {
		if value == nil {
			return Bounds{}, errors.New("rex: incomplete window rectangle")
		}
	}
	bounds := Bounds{X: *coordinates[0], Y: *coordinates[1], W: *coordinates[2], H: *coordinates[3]}
	if err := bounds.valid(); err != nil {
		return Bounds{}, err
	}
	return bounds, nil
}

// layoutRectangles extracts leaf rectangles directly: session.view defines them
// in WINDOW coordinates, including floating leaves. Multiplying by layer bounds
// again would underestimate both split and floating origins. Invalid or absent
// layout metadata is unavailable geometry, not a full-window measurement.
func layoutRectangles(data json.RawMessage) (map[string]Bounds, error) {
	rectangles := make(map[string]Bounds)
	nodes := 0
	var visit func(json.RawMessage, int) error
	visit = func(data json.RawMessage, depth int) error {
		nodes++
		if depth > 64 || nodes > 8192 {
			return errors.New("rex: layout geometry exceeds traversal bound")
		}
		var node struct {
			Rect    json.RawMessage `json:"rect"`
			BlockID string          `json:"block_id"`
			Split   *struct {
				Before json.RawMessage `json:"before"`
				After  json.RawMessage `json:"after"`
			} `json:"split"`
		}
		if err := json.Unmarshal(data, &node); err != nil {
			return err
		}
		rect, err := decodeBounds(node.Rect)
		if err != nil {
			return err
		}
		if node.Split != nil {
			if node.BlockID != "" {
				return errors.New("rex: ambiguous layout geometry")
			}
			if err := visit(node.Split.Before, depth+1); err != nil {
				return err
			}
			return visit(node.Split.After, depth+1)
		}
		_, duplicate := rectangles[node.BlockID]
		if !resourceID(node.BlockID, "block") || duplicate {
			return errors.New("rex: invalid or duplicate geometry block")
		}
		rectangles[node.BlockID] = rect
		return nil
	}
	if err := visit(data, 0); err != nil {
		return nil, err
	}
	return rectangles, nil
}
