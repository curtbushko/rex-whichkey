package whichkey

import (
	"context"
	"encoding/json"
	"math"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	lua "github.com/yuin/gopher-lua"
)

const (
	cacheHomeFallback = "home fallback"
	cacheAbsent       = "missing"
)

func TestDirectMaximumHeightIsValidated(t *testing.T) {
	const uuid = "00000000-0000-0000-0000-000000000001"
	options := DirectOptions{OriginSession: "session:" + uuid, OriginBlock: "block:" + uuid,
		OriginWindow: "window:" + uuid, OriginLayer: "layer:" + uuid}
	t.Setenv("REX_SESSION", "")
	t.Setenv("REX_BLOCK", "block:00000000-0000-0000-0000-000000000002")
	for _, maximum := range []float64{0, .2, 1} {
		options.MaximumHeight = maximum
		_, _, err := directInvocation(context.Background(), options)
		require.NoError(t, err)
	}
	for _, maximum := range []float64{-.1, 1.1, math.NaN(), math.Inf(1)} {
		options.MaximumHeight = maximum
		_, _, err := directInvocation(context.Background(), options)
		require.ErrorContains(t, err, "--max-height")
	}
}

func TestGeometryCachePathAndAtomicReplacement(t *testing.T) {
	home, cache := t.TempDir(), t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_CACHE_HOME", "")
	require.Equal(t, filepath.Join(home, ".cache", "rex-whichkey", "geometry.json"), geometryCachePath())
	t.Setenv("XDG_CACHE_HOME", cache)
	path := filepath.Join(cache, "rex-whichkey", "geometry.json")
	require.Equal(t, path, geometryCachePath())
	entry := geometryCacheEntry{Version: 1, Width: .8, MaximumHeight: .2, Height: .12, Config: "/tmp/menu.json", ASCII: true}
	require.NoError(t, writeGeometryCache(path, entry))
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	var saved geometryCacheEntry
	require.NoError(t, json.Unmarshal(data, &saved))
	require.Equal(t, entry, saved)
	info, err := os.Stat(path)
	require.NoError(t, err)
	require.Equal(t, os.FileMode(0600), info.Mode().Perm())
	entry.Height = .1
	require.NoError(t, writeGeometryCache(path, entry))
	data, err = os.ReadFile(path)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(data, &saved))
	require.Equal(t, .1, saved.Height)
	files, err := os.ReadDir(filepath.Dir(path))
	require.NoError(t, err)
	require.Len(t, files, 1, "atomic writes must not leave temporary files")
	t.Setenv("XDG_CACHE_HOME", "relative/cache")
	require.Empty(t, geometryCachePath(), "relative cache roots must not write under the caller's cwd")
}

func TestLuaGeometryCacheReadsGoHintsAndFallsBack(t *testing.T) {
	for _, scenario := range []string{"hit", cacheAbsent, "corrupt", "oversized", "width", "maximum", "config", "ascii", "schema mismatch", "negative", "over cap", cacheHomeFallback, "relative root"} {
		t.Run(scenario, func(t *testing.T) {
			t.Setenv("HOME", t.TempDir())
			t.Setenv("XDG_CACHE_HOME", t.TempDir())
			if scenario == cacheHomeFallback {
				t.Setenv("XDG_CACHE_HOME", "")
			}
			path := geometryCachePath()
			entry := geometryCacheEntry{Version: 1, Width: .8, MaximumHeight: .2, Height: .12, Config: "/tmp/menu.json"}
			require.NoError(t, writeGeometryCache(path, entry))
			want := .2
			switch scenario {
			case "hit", cacheHomeFallback:
				want = entry.Height
			case cacheAbsent:
				path = filepath.Join(t.TempDir(), "absent", "geometry.json")
				t.Setenv("XDG_CACHE_HOME", filepath.Dir(filepath.Dir(path)))
			case "width":
				entry.Width = .7
			case "maximum":
				entry.MaximumHeight = .3
			case "config":
				entry.Config = "/tmp/other.json"
			case "ascii":
				entry.ASCII = true
			case "schema mismatch":
				entry.Version = 9
			case "negative":
				entry.Height = -.1
			case "over cap":
				entry.Height = .3
			case "relative root":
				t.Setenv("XDG_CACHE_HOME", "relative")
			}
			if scenario != cacheAbsent {
				data, err := json.Marshal(entry)
				require.NoError(t, err)
				if scenario == "corrupt" {
					data = []byte(`not json`)
				}
				if scenario == "oversized" {
					data = append(data, make([]byte, 4097)...)
				}
				require.NoError(t, os.WriteFile(path, data, 0600))
			}
			state := lua.NewState()
			defer state.Close()
			require.NoError(t, state.DoFile("tests/lua/geometry_launch.lua"))
			decoder := state.NewTable()
			state.SetField(decoder, "decode", state.NewFunction(func(state *lua.LState) int {
				var value any
				if err := json.Unmarshal([]byte(state.CheckString(1)), &value); err != nil {
					state.RaiseError("invalid JSON")
				}
				state.Push(geometryLuaValue(state, value))
				return 1
			}))
			state.SetField(state.GetGlobal("rex"), "json", decoder)
			require.NoError(t, state.DoFile("lua/rex_whichkey.lua"))
			require.NoError(t, state.DoString(`launch_action.run(launch_context, {config="/tmp/menu.json"})`))
			layer := state.GetGlobal("captured_layer").(*lua.LTable)
			bounds := layer.RawGetString("bounds").(*lua.LTable)
			require.InDelta(t, want, float64(bounds.RawGetInt(4).(lua.LNumber)), 1e-9)
			require.InDelta(t, (1-want)/2, float64(bounds.RawGetInt(2).(lua.LNumber)), 1e-9)
			block := state.GetGlobal("captured_block").(*lua.LTable)
			options := block.RawGetString("options").(*lua.LTable)
			command := options.RawGetString("command").(*lua.LTable)
			var argv []string
			command.ForEach(func(_ lua.LValue, value lua.LValue) { argv = append(argv, value.String()) })
			require.Contains(t, argv, "--max-height")
			require.Contains(t, argv, "0.2", "cached launch height must not replace the configured growth cap")
		})
	}
}

func geometryLuaValue(state *lua.LState, value any) lua.LValue {
	switch value := value.(type) {
	case map[string]any:
		table := state.NewTable()
		for key, member := range value {
			table.RawSetString(key, geometryLuaValue(state, member))
		}
		return table
	case string:
		return lua.LString(value)
	case float64:
		return lua.LNumber(value)
	case bool:
		return lua.LBool(value)
	default:
		return lua.LNil
	}
}

func TestGeometryCacheFailedReplacementCleansTemporaryFile(t *testing.T) {
	directory := t.TempDir()
	path := filepath.Join(directory, "geometry.json")
	require.NoError(t, os.Mkdir(path, 0700))
	entry := geometryCacheEntry{Version: 1, Width: .8, MaximumHeight: .2, Height: .1}
	require.Error(t, writeGeometryCache(path, entry))
	files, err := os.ReadDir(directory)
	require.NoError(t, err)
	require.Len(t, files, 1, "failed replacement must not leak temporary files")
	require.True(t, files[0].IsDir())
}

func TestGeometryCacheWriteRejectsInvalidHintsAndDoesNotFollowFinalSymlink(t *testing.T) {
	cache := t.TempDir()
	path := filepath.Join(cache, "geometry.json")
	entry := geometryCacheEntry{Version: 1, Width: .8, MaximumHeight: .2, Height: .3}
	require.Error(t, writeGeometryCache(path, entry))
	_, err := os.Stat(path)
	require.ErrorIs(t, err, os.ErrNotExist)
	victim := filepath.Join(cache, "unrelated")
	require.NoError(t, os.WriteFile(victim, []byte("unchanged"), 0600))
	require.NoError(t, os.Symlink(victim, path))
	entry.Height = .1
	require.NoError(t, writeGeometryCache(path, entry))
	data, err := os.ReadFile(victim)
	require.NoError(t, err)
	require.Equal(t, "unchanged", string(data))
}
