package whichkey

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"unicode"

	"github.com/rivo/uniseg"
	"github.com/stretchr/testify/require"
)

const (
	scriptPermission = "permission"
	scriptNonregular = "nonregular"
	scriptRepairCWD  = "Check the origin pane's foreground cwd."
	scriptRepairExec = "Check cwd, executable permission and shebang interpreter."
)

func TestScriptConfig(t *testing.T) {
	for _, args := range []string{`{"path":"./build Ω;$(touch nope)","args":["","a b","$HOME","*","\"'","Ω"]}`, `{"path":"/opt/task"}`, `{"path":"a\nb","args": [ "\n\t", "\ud83d\ude00" ]}`} {
		options, err := decodeMenuOptions([]byte(`{"version":1,"spec":[{"keys":"x","desc":"Run","action":"script.run","confirm":true,"args":` + args + `}]}`))
		require.NoError(t, err)
		defaults, err := nativeDefaults()
		require.NoError(t, err)
		config, err := resolveMenuConfig(defaults, &options)
		require.NoError(t, err)
		target, err := decodeScriptAction(config.Spec[0].Args)
		require.NoError(t, err)
		var expected struct {
			Path string
			Args []string
		}
		require.NoError(t, json.Unmarshal([]byte(args), &expected))
		require.Equal(t, expected.Path, target.path)
		require.Equal(t, expected.Args, target.args)
	}
	for _, args := range []string{`{}`, `null`, `{"path":""}`, `{"path":null}`, `{"Path":"x"}`, `{"path":"x","path":"y"}`, `{"path":"x","args":null}`, `{"path":"x","args":"x"}`, `{"path":"x","args":[1]}`, `{"path":"x","args":[null]}`, `{"path":"x","args":["\u0000"]}`, `{"path":"x\u0000"}`, `{"path":"\ud800"}`, `{"path":"x","shell":"bash"}`, `{"path":"x","Args":[]}`} {
		t.Run(args, func(t *testing.T) { _, err := decodeScriptAction([]byte(args)); require.Error(t, err) })
	}
}

func TestScriptLoadAndStandaloneAreEffectFree(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	data := []byte(`{"version":1,"spec":[{"keys":"z","desc":"Never execute here","action":"script.run","args":{"path":"/definitely/missing/script","args":["; touch nope"]}}]}`)
	require.NoError(t, os.WriteFile(path, data, 0600))
	config, err := LoadMenuConfig(&path)
	require.NoError(t, err, "loading does not inspect target or require Rex")
	menu, err := NewMenu(config.Spec)
	require.NoError(t, err)
	event := menu.HandleKey("z", 1)
	var result bytes.Buffer
	require.NoError(t, WriteMenuResult(&result, event))
	require.JSONEq(t, `{"keys":"z","confirmed":false}`, result.String(), "standalone returns keys, never executable authority")
}

func TestScriptErrorAcknowledgement(t *testing.T) {
	for _, width := range []int{1, 10, 80} {
		frame := scriptErrorFrame(errors.New("missing target Ω\x1b[31m\nno retry"), width, 12)
		cursor := regexp.MustCompile(`\x1b\[[0-9]+;1H`)
		for _, line := range cursor.Split(strings.TrimPrefix(frame, "\x1b[39;49m\x1b[2J"), -1) {
			require.LessOrEqual(t, uniseg.StringWidth(line), width)
			require.NotContains(t, line, "\x1b", "external diagnostics cannot inject terminal escapes")
		}
	}
	for _, keys := range []string{"xyy", "\x1b[200~qxyy\x1b[201~", "\x1b[I"} {
		var input terminalKeys
		ack, err := scriptAcknowledgementKeys(&input, []byte(keys))
		require.NoError(t, err)
		require.False(t, ack, "mapping, confirmation, paste and focus events are not acknowledgements")
	}
	var input terminalKeys
	ack, err := scriptAcknowledgementKeys(&input, []byte("q"))
	require.NoError(t, err)
	require.True(t, ack)
}

// Read only the explicitly positioned acknowledgement frame, not the returned
// error or stderr emitted after closing it. Bounds include every drawn row.
func scriptDrawnRows(t *testing.T, frame string, width, height int) []string {
	t.Helper()
	start := strings.LastIndex(frame, "\x1b[39;49m\x1b[2J")
	require.NotEqual(t, -1, start)
	frame = frame[start+len("\x1b[39;49m\x1b[2J"):]
	cursor := regexp.MustCompile(`\x1b\[([0-9]+);1H`)
	positions := cursor.FindAllStringSubmatchIndex(frame, -1)
	rows := make([]string, height)
	for index, position := range positions {
		row, err := strconv.Atoi(frame[position[2]:position[3]])
		require.NoError(t, err)
		require.GreaterOrEqual(t, row, 1)
		require.LessOrEqual(t, row, height)
		end := len(frame)
		if index+1 < len(positions) {
			end = positions[index+1][0]
		}
		text := frame[position[1]:end]
		require.False(t, strings.ContainsFunc(text, unicode.IsControl), "only renderer-owned control sequences are allowed")
		require.LessOrEqual(t, uniseg.StringWidth(text), width)
		rows[row-1] = text
	}
	return rows
}

func TestScriptShortErrorFrame(t *testing.T) {
	cwd := filepath.Join(t.TempDir(), strings.Repeat("long-origin-Ω-", 12))
	require.NoError(t, os.MkdirAll(cwd, 0700))
	_, problem := prepareScript(scriptAction{path: "./scripts/report.sh"}, cwd)
	require.Error(t, problem)
	for _, height := range []int{5, 7} {
		t.Run(strconv.Itoa(height), func(t *testing.T) {
			rows := scriptDrawnRows(t, scriptErrorFrame(problem, 64, height), 64, height)
			drawn := strings.Join(rows, " ")
			require.Contains(t, drawn, "no such file or directory")
			require.Contains(t, drawn, "Check the literal path in the origin cwd.")
			require.Contains(t, drawn, "Esc/q: close; no retry")
		})
	}
}

func TestScriptShortDiagnosticFailures(t *testing.T) {
	cwd := filepath.Join(t.TempDir(), strings.Repeat("long-origin-Ω-", 12))
	require.NoError(t, os.MkdirAll(cwd, 0700))
	path := filepath.Join(cwd, "script")
	require.NoError(t, os.WriteFile(path, []byte("#!/bin/sh\n"), 0600))
	_, permission := prepareScript(scriptAction{path: path}, cwd)
	_, nonregular := prepareScript(scriptAction{path: cwd}, cwd)
	_, missingCWD := prepareScript(scriptAction{path: path}, filepath.Join(cwd, "missing"))
	_, fileCWD := prepareScript(scriptAction{path: path}, path)
	started, latePermission := startScript(context.Background(), preparedScript{path: path, cwd: cwd})
	require.False(t, started)
	require.NoError(t, os.Chmod(path, 0700))
	require.NoError(t, os.WriteFile(path, []byte("not a shebang\n"), 0700))
	started, badFormat := startScript(context.Background(), preparedScript{path: path, cwd: cwd})
	require.False(t, started)
	started, lateCWD := startScript(context.Background(), preparedScript{path: path, cwd: filepath.Join(cwd, "missing")})
	require.False(t, started)
	for name, test := range map[string]struct {
		problem         error
		cause, recovery string
	}{
		scriptPermission:  {permission, "Target is not executable", "Set executable permission and a valid shebang."},
		scriptNonregular:  {nonregular, "regular executable script", "Choose a regular file, not a directory or device."},
		"missing cwd":     {missingCWD, "Original cwd: no such file or directory", scriptRepairCWD},
		"file cwd":        {fileCWD, "Original cwd is not a directory", scriptRepairCWD},
		"late permission": {latePermission, "direct exec: permission denied", scriptRepairExec},
		"bad format":      {badFormat, "direct exec: exec format error", scriptRepairExec},
		"late cwd":        {lateCWD, "direct exec: no such file or directory", scriptRepairExec},
	} {
		t.Run(name, func(t *testing.T) {
			require.Error(t, test.problem)
			require.Contains(t, test.problem.Error(), cwd, "caller keeps the full path")
			for _, height := range []int{5, 7} {
				drawn := strings.Join(scriptDrawnRows(t, scriptErrorFrame(fmt.Errorf("caller: %w", test.problem), 64, height), 64, height), " ")
				require.Contains(t, drawn, test.cause)
				require.Contains(t, drawn, test.recovery)
				require.Contains(t, drawn, "Esc/q: close; no retry")
			}
		})
	}
	require.ErrorIs(t, missingCWD, os.ErrNotExist)
	require.ErrorIs(t, latePermission, os.ErrPermission)
	var original *os.PathError
	require.ErrorAs(t, lateCWD, &original)
	require.Equal(t, path, original.Path)
}

func TestScriptErrorFrameBounds(t *testing.T) {
	require.Equal(t, "e\u0301", scriptClipLine("e\u0301界👩‍💻", 1))
	require.Equal(t, "e\u0301界", scriptClipLine("e\u0301界👩‍💻", 3))
	require.Equal(t, "e\u0301界👩‍💻", scriptClipLine("e\u0301界👩‍💻", 5))
	// Both the structured and generic paths must sanitize external diagnostics.
	problem := errors.New("Ω界e\u0301👩‍💻\x1b[31m\r\n\t" + strings.Repeat("long", 200))
	for _, diagnostic := range []error{problem, scriptProblem(problem, problem.Error(), problem.Error())} {
		for _, width := range []int{0, 1, 2, 10, 64} {
			for _, height := range []int{0, 1, 2, 3, 5, 7} {
				scriptDrawnRows(t, scriptErrorFrame(diagnostic, width, height), max(1, width), max(1, height))
			}
		}
	}
}

func TestScriptPreflight(t *testing.T) {
	dir := t.TempDir()
	executable := filepath.Join(dir, "script Ω; $HOME")
	require.NoError(t, os.WriteFile(executable, []byte("#!/bin/sh\nexit 0\n"), 0700))
	for _, path := range []string{executable, filepath.Base(executable)} {
		prepared, err := prepareScript(scriptAction{path: path, args: []string{"", scriptLiteralHome, "*"}}, dir)
		require.NoError(t, err)
		require.Equal(t, executable, prepared.path)
		require.Equal(t, dir, prepared.cwd)
	}
	require.NoError(t, os.Symlink(executable, filepath.Join(dir, "link")))
	_, err := prepareScript(scriptAction{path: "link"}, dir)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(dir, "not-executable"), []byte("#!/bin/sh\n"), 0600))
	require.NoError(t, os.Symlink(filepath.Join(dir, "absent"), filepath.Join(dir, "dangling")))
	for _, path := range []string{"absent", ".", "not-executable", "dangling", "~/$HOME", "/dev/null"} {
		_, err := prepareScript(scriptAction{path: path}, dir)
		require.Error(t, err, path)
	}
	for _, cwd := range []string{"", "relative", filepath.Join(dir, "absent"), executable} {
		_, err := prepareScript(scriptAction{path: executable}, cwd)
		require.Error(t, err, cwd)
	}
}
