package whichkey

// Private popup transport only. The former Lua consumer, receipt, relay and
// detached controller protocol is retired; public startup has no registry route.
import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"

	"golang.org/x/sys/unix"

	"temp/rex"
)

const (
	runtimeArgsField  = "args"
	runtimeLimit      = 8 << 20
	runtimeImageLimit = 512 << 20
	runtimeImageName  = "runtime-image"
)

var (
	errRuntimeLayout        = errors.New("whichkey: cannot capture a bounded regular popup source image")
	errRuntimeCompatibility = errors.New("whichkey: selected CLI/endpoint or private popup contract is invalid")
	errRuntimePreparation   = errors.New("whichkey: private popup preparation failed")
	errRuntimeRelay         = errors.New("whichkey: private popup file identity or access failed")
	errRuntimeStartup       = errors.New("whichkey: private popup startup or result failed")
)

// Runtime field names are lowercase ASCII, not merely Unicode lowercase:
// encoding/json also matches struct fields through Unicode SimpleFold. Restrict
// decoded names before typed decoding so DisallowUnknownFields is exact. User
// args are case-sensitive data maps; opaque JSON strings are never rewritten.
func checkRuntimeJSON(data []byte, limit int) error {
	if len(data) == 0 || len(data) > limit || !utf8.Valid(data) || !json.Valid(data) || data[0] != '{' {
		return errRuntimeCompatibility
	}
	if !runtimeUnicode(data) {
		return errRuntimeCompatibility
	}
	return checkRuntimeValue(json.NewDecoder(bytes.NewReader(data)), false, 0)
}

func checkRuntimeValue(decoder *json.Decoder, args bool, depth int) error {
	if depth > 64 {
		return errRuntimeCompatibility
	}
	token, err := decoder.Token()
	if err != nil || token == nil {
		return errRuntimeCompatibility
	}
	delimiter, container := token.(json.Delim)
	if !container {
		return nil
	}
	seen := make(map[string]bool)
	for decoder.More() {
		childArgs := args
		if delimiter == '{' {
			name, fieldErr := runtimeJSONField(decoder, seen, args)
			if fieldErr != nil {
				return fieldErr
			}
			childArgs = args || name == runtimeArgsField
		}
		if checkErr := checkRuntimeValue(decoder, childArgs, depth+1); checkErr != nil {
			return checkErr
		}
	}
	_, err = decoder.Token()
	return err
}

func runtimeJSONField(decoder *json.Decoder, seen map[string]bool, args bool) (string, error) {
	token, err := decoder.Token()
	name, ok := token.(string)
	if err != nil || !ok || seen[name] {
		return "", errRuntimeCompatibility
	}
	if !args && (name == "" || strings.Trim(name, "abcdefghijklmnopqrstuvwxyz0123456789_") != "") {
		return "", errRuntimeCompatibility
	}
	seen[name] = true
	return name, nil
}

// Private invocation context and resolved popup snapshots check exact field
// names before struct decoding; JSON strings containing metadata remain opaque.
func decodeRuntime(data []byte, value any, limit int) error {
	if checkErr := checkRuntimeJSON(data, limit); checkErr != nil {
		return checkErr
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if checkErr := decoder.Decode(value); checkErr != nil {
		return errRuntimeCompatibility
	}
	return nil
}

// JSON validity alone accepts unpaired surrogate escapes. Reuse Completion's
// scalar string boundary for every name/value before typed decoding.
func runtimeUnicode(data []byte) bool {
	for index := 0; index < len(data); index++ {
		if data[index] != '"' {
			continue
		}
		start := index
		index++
		for index < len(data) {
			if data[index] == '\\' {
				index += 2
				continue
			}
			if data[index] == '"' {
				break
			}
			index++
		}
		if index >= len(data) {
			return false
		}
		if _, err := completionString(data[start : index+1]); err != nil {
			return false
		}
	}
	return true
}

func readRuntime(input io.Reader, value any, limit int) error {
	data, err := io.ReadAll(io.LimitReader(input, int64(limit)+1))
	if err != nil || len(data) > limit {
		return errRuntimeCompatibility
	}
	return decodeRuntime(bytes.TrimSpace(data), value, limit)
}

// runtimeImage holds the public image open until exact owned cleanup. It is
// copied from this handle, never reopened through an installation receipt.
type runtimeImage struct {
	file *os.File
	size int64
}

type runtimeSourceReader struct {
	ctx  context.Context
	file *os.File
}

func (reader runtimeSourceReader) Read(data []byte) (int, error) {
	if err := reader.ctx.Err(); err != nil {
		return 0, err
	}
	return reader.file.Read(data)
}

func captureSelf(ctx context.Context) (_ *runtimeImage, outcome error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	source, err := os.Executable()
	if err != nil {
		return nil, errRuntimeLayout
	}
	canonical, err := filepath.EvalSymlinks(source)
	if err != nil {
		return nil, errRuntimeLayout
	}
	file, err := os.OpenFile(canonical, os.O_RDONLY|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if err != nil {
		return nil, errRuntimeLayout
	}
	defer func() {
		if outcome != nil {
			outcome = errors.Join(outcome, file.Close())
		}
	}()
	info, err := file.Stat()
	if err != nil || !info.Mode().IsRegular() || info.Size() <= 0 || info.Size() > runtimeImageLimit || info.Mode().Perm()&0111 == 0 {
		return nil, errRuntimeLayout
	}
	return &runtimeImage{file: file, size: info.Size()}, nil
}

func (image *runtimeImage) copy(ctx context.Context, directory string) (outcome error) {
	if image == nil || image.file == nil || image.size <= 0 || image.size > runtimeImageLimit {
		return errRuntimePreparation
	}
	if _, err := image.file.Seek(0, io.SeekStart); err != nil {
		return errRuntimePreparation
	}
	root, err := openRuntimeDirectory(directory)
	if err != nil {
		return err
	}
	defer func() { outcome = errors.Join(outcome, root.Close()) }()
	file, err := root.OpenFile(runtimeImageName, os.O_WRONLY|os.O_CREATE|os.O_EXCL|unix.O_NOFOLLOW, 0700)
	if err != nil {
		return errRuntimePreparation
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	size, err := io.Copy(file, io.LimitReader(runtimeSourceReader{ctx, image.file}, runtimeImageLimit+1))
	if err != nil || size != image.size || ctx.Err() != nil {
		return errRuntimePreparation
	}
	// Keep fixed-slot ownership/mode checks, not executable receipts. Deferred
	// closes must succeed before process start; power-loss Sync is unnecessary.
	return checkPrivateImage(root)
}

func checkPrivateImage(root *os.Root) (outcome error) {
	file, err := root.OpenFile(runtimeImageName, os.O_RDONLY|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if err != nil {
		return errRuntimeStartup
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	info, err := file.Stat()
	if err != nil || !info.Mode().IsRegular() || info.Mode().Perm() != 0700 || info.Size() <= 0 || info.Size() > runtimeImageLimit {
		return errRuntimeStartup
	}
	owner, ok := info.Sys().(*syscall.Stat_t)
	if !ok || owner.Uid != uint32(os.Geteuid()) || owner.Nlink != 1 {
		return errRuntimeStartup
	}
	return nil
}

func newRuntimeDirectory(prefix string) (string, error) {
	directory, err := os.MkdirTemp("", prefix)
	if err != nil {
		return "", errRuntimePreparation
	}
	canonical, err := filepath.EvalSymlinks(directory)
	if err != nil {
		return "", errors.Join(errRuntimePreparation, os.Remove(directory))
	}
	return canonical, nil
}

func openRuntimeDirectory(directory string) (*os.Root, error) {
	before, err := os.Lstat(directory)
	if err != nil {
		return nil, errRuntimeRelay
	}
	root, err := openPopupDirectory(directory)
	if err != nil {
		return nil, errRuntimeRelay
	}
	after, err := root.Stat(".")
	if err != nil || !os.SameFile(before, after) || after.Mode().Perm() != 0700 {
		return nil, errors.Join(errRuntimeRelay, root.Close())
	}
	return root, nil
}

func checkPopupSlot(file *os.File, fifo bool) error {
	info, err := file.Stat()
	if err != nil || info.Mode().Perm() != 0600 || (fifo && info.Mode()&os.ModeNamedPipe == 0) || (!fifo && !info.Mode().IsRegular()) {
		return errRuntimeRelay
	}
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok || stat.Uid != uint32(os.Geteuid()) || stat.Nlink != 1 {
		return errRuntimeRelay
	}
	return nil
}

func (image *runtimeImage) retire() error {
	file := image.file
	image.file = nil
	if file == nil {
		return nil
	}
	if err := file.Close(); err != nil {
		return errRuntimeStartup
	}
	return nil
}

func (r popupRequest) preparePopupImage(ctx context.Context, directory string) error {
	endpoint, err := rex.CanonicalEndpoint(r.Server)
	if err != nil || endpoint != r.Server || r.selection == nil {
		return errRuntimeCompatibility
	}
	selection, err := rex.ResolveCLI(ctx, rex.CLIOptions{Selection: r.selection})
	if err != nil {
		return errors.Join(errRuntimeCompatibility, err)
	}
	if checkErr := r.image.copy(ctx, directory); checkErr != nil {
		return checkErr
	}
	invocation, err := json.Marshal(popupRuntime{CLI: selection, Endpoint: endpoint, Invocation: r.invocationID})
	if err != nil || len(invocation) > runtimeLimit {
		return errRuntimeStartup
	}
	return writePopupFile(directory, "runtime.json", invocation)
}

func openVerifiedPopupDirectory(ctx context.Context, directory string) (*os.Root, *popupRuntime, error) {
	root, err := openRuntimeDirectory(directory)
	if err != nil {
		return nil, nil, err
	}
	invocation, checkErr := readPopupRuntime(ctx, root, directory)
	if checkErr != nil {
		return nil, nil, errors.Join(checkErr, root.Close())
	}
	return root, invocation, nil
}

func commitPopupResult(ctx context.Context, root *os.Root, event MenuEvent, config Config) (outcome error) {
	output := &launchBuffer{remaining: popupMessageLimit}
	if err := writePopupMenuResult(output, event, config); err != nil {
		return err
	}
	file, err := root.OpenFile("result.json", os.O_WRONLY|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if err != nil {
		return errRuntimeStartup
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	if statErr := checkPopupSlot(file, false); statErr != nil {
		return errRuntimeStartup
	}
	if checkErr := runtimeLock(ctx, file, true); checkErr != nil {
		return checkErr
	}
	info, err := file.Stat()
	if err != nil || info.Size() != 0 {
		return errRuntimeStartup
	}
	if _, writeErr := file.Write(output.buffer.Bytes()); writeErr != nil {
		return errRuntimeStartup
	}
	if syncErr := file.Sync(); syncErr != nil {
		return errRuntimeStartup
	}
	return nil
}

// popupRuntime is private invocation context, not binary/API evidence. The
// endpoint is independent of the selected PATH executable.
type popupRuntime struct {
	CLI        *rex.CLISelection `json:"cli"`
	Endpoint   string            `json:"endpoint"`
	Invocation string            `json:"invocation"`
}

func readPopupRuntime(ctx context.Context, root *os.Root, directory string) (_ *popupRuntime, outcome error) {
	file, err := root.OpenFile("runtime.json", os.O_RDONLY|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if err != nil {
		return nil, errRuntimeStartup
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	if statErr := checkPopupSlot(file, false); statErr != nil {
		return nil, errRuntimeStartup
	}
	var invocation popupRuntime
	if checkErr := readRuntime(file, &invocation, runtimeLimit); checkErr != nil || invocation.CLI == nil || len(invocation.Invocation) < 26 || len(invocation.Invocation) > 64 ||
		strings.Trim(invocation.Invocation, "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567") != "" {
		return nil, errRuntimeStartup
	}
	endpoint, err := rex.CanonicalEndpoint(invocation.Endpoint)
	if err != nil || endpoint != invocation.Endpoint {
		return nil, errRuntimeCompatibility
	}
	invocation.CLI, err = rex.ResolveCLI(ctx, rex.CLIOptions{Selection: invocation.CLI})
	if err != nil {
		return nil, errors.Join(errRuntimeCompatibility, err)
	}
	// A private child must run from its invocation's fixed, owned image slot.
	// Do not hash it or reopen/revalidate the parent's original installation.
	self, err := os.Executable()
	if err != nil || self != filepath.Join(directory, runtimeImageName) {
		return nil, errRuntimeStartup
	}
	if checkErr := checkPrivateImage(root); checkErr != nil {
		return nil, checkErr
	}
	return &invocation, nil
}

func runtimeLock(ctx context.Context, file *os.File, exclusive bool) error {
	operation := unix.LOCK_SH | unix.LOCK_NB
	if exclusive {
		operation = unix.LOCK_EX | unix.LOCK_NB
	}
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()
	for {
		if checkErr := unix.Flock(int(file.Fd()), operation); checkErr == nil {
			return nil
		} else if !errors.Is(checkErr, unix.EWOULDBLOCK) && !errors.Is(checkErr, unix.EAGAIN) {
			return errRuntimeRelay
		}
		select {
		case <-ctx.Done():
			return errRuntimeRelay
		case <-ticker.C:
		}
	}
}
