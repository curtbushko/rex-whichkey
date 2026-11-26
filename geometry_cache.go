package whichkey

import (
	"encoding/json"
	"errors"
	"math"
	"os"
	"path/filepath"
)

// Only disposable presentation hints cross the Lua/Go cache boundary. Never put
// resource IDs, app routing, endpoint selectors or executable paths in this file.
type geometryCacheEntry struct {
	Version       int     `json:"version"`
	Width         float64 `json:"width"`
	MaximumHeight float64 `json:"maximum_height"`
	Height        float64 `json:"height"`
	Config        string  `json:"config"`
	ASCII         bool    `json:"ascii"`
}

func geometryCachePath() string {
	root := os.Getenv("XDG_CACHE_HOME")
	if root == "" {
		home := os.Getenv("HOME")
		if !filepath.IsAbs(home) {
			return ""
		}
		root = filepath.Join(home, ".cache")
	}
	if !filepath.IsAbs(root) {
		return ""
	}
	return filepath.Join(root, "rex-whichkey", "geometry.json")
}

func geometryFraction(value float64) bool {
	return !math.IsNaN(value) && !math.IsInf(value, 0) && value > 0 && value <= 1
}

func writeGeometryCache(path string, entry geometryCacheEntry) error {
	if path == "" || entry.Version != 1 || !geometryFraction(entry.Width) ||
		!geometryFraction(entry.MaximumHeight) || !geometryFraction(entry.Height) || entry.Height > entry.MaximumHeight {
		return errors.New("whichkey: invalid geometry cache hint")
	}
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	if len(data) > 4096 {
		return errors.New("whichkey: geometry cache hint exceeds size limit")
	}
	if directoryErr := os.MkdirAll(filepath.Dir(path), 0700); directoryErr != nil {
		return directoryErr
	}
	file, err := os.CreateTemp(filepath.Dir(path), ".geometry-*")
	if err != nil {
		return err
	}
	defer func() { _ = os.Remove(file.Name()) }()
	_, writeErr := file.Write(data)
	if closeErr := errors.Join(writeErr, file.Close()); closeErr != nil {
		return closeErr
	}
	// Rename replaces the directory entry, not a symlink's target. Concurrent
	// launches are last-writer-wins: losing an optional hint cannot lose config.
	return os.Rename(file.Name(), path)
}
