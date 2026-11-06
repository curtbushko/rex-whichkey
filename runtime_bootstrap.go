package whichkey

// Private popup transport only. The former Lua consumer, receipt, relay and
// detached controller protocol is retired; public startup has no registry route.
import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
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
	runtimeAssets     = "rex-whichkey/go-popup/1;data-menu/1;frozen-cli/1"
)

var (
	errRuntimeLayout        = errors.New("whichkey: popup image or installation identity is invalid or changed")
	errRuntimeCompatibility = errors.New("whichkey: frozen CLI/endpoint or private popup contract is invalid or changed")
	errRuntimePreparation   = errors.New("whichkey: private popup preparation failed")
	errRuntimeRelay         = errors.New("whichkey: private popup file identity or access failed")
	errRuntimeStartup       = errors.New("whichkey: private popup startup or result failed")
)

func runtimeHash(data []byte) string {
	digest := sha256.Sum256(data)
	return hex.EncodeToString(digest[:])
}

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

// Private image receipts and resolved popup snapshots check exact field names
// before struct decoding; JSON strings containing metadata remain opaque.
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

// runtimeFile records only application-owned components. Rex discovery and its
// proprietary dependency identity remain owned by rex.CLISelection.
type runtimeFile struct {
	Source      string `json:"source"`
	Path        string `json:"path"`
	SHA256      string `json:"sha256"`
	LinksSHA256 string `json:"links_sha256"`
	Device      uint64 `json:"device"`
	Inode       uint64 `json:"inode"`
	Size        int64  `json:"size"`
	ModifiedNS  int64  `json:"modified_ns"`
	Mode        uint32 `json:"mode"`
	Owner       uint32 `json:"owner"`
}

type runtimeImage struct {
	identity runtimeFile
	file     *os.File
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

func runtimePath(path string) bool {
	if path == "" || len(path) > 4096 || validateText(path) != nil || strings.Contains(path, "\\") {
		return false
	}
	for _, part := range strings.Split(filepath.ToSlash(path), "/") {
		if part == ".." {
			return false
		}
	}
	return true
}

// Link metadata is frozen too. Namespace symlinks are supported but updates
// invalidate this invocation even when another link reaches the same inode.
func runtimeLinks(source string) (string, error) {
	type link struct {
		Path   string `json:"path"`
		Device uint64 `json:"device"`
		Inode  uint64 `json:"inode"`
		Target string `json:"target"`
	}
	var links []link
	remaining := filepath.Clean(source)
	for range 256 {
		parts := strings.Split(strings.TrimPrefix(remaining, string(filepath.Separator)), string(filepath.Separator))
		path := string(filepath.Separator)
		followed := false
		for index, part := range parts {
			path = filepath.Join(path, part)
			info, err := os.Lstat(path)
			if err != nil {
				return "", errRuntimeLayout
			}
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok || stat.Uid != 0 && stat.Uid != uint32(os.Geteuid()) {
				return "", errRuntimeLayout
			}
			if info.Mode()&os.ModeSymlink == 0 {
				if info.Mode().Perm()&0022 != 0 && info.Mode()&os.ModeSticky == 0 {
					return "", errRuntimeLayout
				}
				continue
			}
			target, err := os.Readlink(path)
			if err != nil {
				return "", errRuntimeLayout
			}
			links = append(links, link{path, uint64(stat.Dev), stat.Ino, target})
			if !filepath.IsAbs(target) {
				target = filepath.Join(filepath.Dir(path), target)
			}
			remaining = filepath.Join(append([]string{target}, parts[index+1:]...)...)
			followed = true
			break
		}
		if !followed {
			data, err := json.Marshal(links)
			if err != nil {
				return "", errRuntimeLayout
			}
			return runtimeHash(data), nil
		}
	}
	return "", errRuntimeLayout
}

func captureRuntimeFile(ctx context.Context, source string, executable bool) (_ *runtimeImage, outcome error) {
	if !runtimePath(source) || !filepath.IsAbs(source) || ctx.Err() != nil {
		return nil, errRuntimeLayout
	}
	canonical, err := filepath.EvalSymlinks(source)
	if err != nil {
		return nil, errRuntimeLayout
	}
	links, err := runtimeLinks(source)
	if err != nil {
		return nil, err
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
	if err != nil {
		return nil, errRuntimeLayout
	}
	stat, err := trustedRuntimeFile(info, executable)
	if err != nil {
		return nil, err
	}
	hasher := sha256.New()
	size, err := io.Copy(hasher, io.LimitReader(runtimeSourceReader{ctx, file}, runtimeImageLimit+1))
	if err != nil || !runtimeSourceUnchanged(ctx, file, source, canonical, links, info, size) {
		return nil, errRuntimeLayout
	}
	identity := runtimeFile{Source: source, Path: canonical, SHA256: hex.EncodeToString(hasher.Sum(nil)), LinksSHA256: links,
		Device: uint64(stat.Dev), Inode: stat.Ino, Size: size, ModifiedNS: info.ModTime().UnixNano(), Mode: uint32(info.Mode()), Owner: stat.Uid}
	return &runtimeImage{identity: identity, file: file}, nil
}

func trustedRuntimeFile(info os.FileInfo, executable bool) (*syscall.Stat_t, error) {
	if !info.Mode().IsRegular() || info.Mode().Perm()&0022 != 0 || info.Size() <= 0 || info.Size() > runtimeImageLimit || executable && info.Mode().Perm()&0111 == 0 {
		return nil, errRuntimeLayout
	}
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok || stat.Uid != 0 && stat.Uid != uint32(os.Geteuid()) {
		return nil, errRuntimeLayout
	}
	return stat, nil
}

func runtimeSourceUnchanged(ctx context.Context, file *os.File, source, canonical, links string, before os.FileInfo, size int64) bool {
	after, err := file.Stat()
	linked, linkErr := os.Stat(source)
	path, pathErr := filepath.EvalSymlinks(source)
	afterLinks, linksErr := runtimeLinks(source)
	if err != nil || linkErr != nil || pathErr != nil || linksErr != nil || ctx.Err() != nil {
		return false
	}
	return path == canonical && links == afterLinks && os.SameFile(before, after) && os.SameFile(after, linked) &&
		size == before.Size() && size == after.Size() && before.Mode() == after.Mode() && before.ModTime().Equal(after.ModTime())
}

func captureSelf(ctx context.Context) (*runtimeImage, error) {
	source, err := os.Executable()
	if err != nil {
		return nil, errRuntimeLayout
	}
	return captureRuntimeFile(ctx, source, true)
}

func checkRuntimeFile(ctx context.Context, expected runtimeFile, executable bool) error {
	image, err := captureRuntimeFile(ctx, expected.Source, executable)
	if err != nil {
		return err
	}
	closeErr := image.file.Close()
	if image.identity != expected || closeErr != nil {
		return errRuntimeLayout
	}
	return nil
}

func (image *runtimeImage) copy(ctx context.Context, directory string) (outcome error) {
	if checkErr := checkRuntimeFile(ctx, image.identity, true); checkErr != nil {
		return checkErr
	}
	if _, err := image.file.Seek(0, io.SeekStart); err != nil {
		return errRuntimePreparation
	}
	file, err := os.OpenFile(filepath.Join(directory, runtimeImageName), os.O_WRONLY|os.O_CREATE|os.O_EXCL|unix.O_NOFOLLOW, 0700)
	if err != nil {
		return errRuntimePreparation
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	hasher := sha256.New()
	size, err := io.Copy(io.MultiWriter(file, hasher), io.LimitReader(runtimeSourceReader{ctx, image.file}, runtimeImageLimit+1))
	if err != nil || size != image.identity.Size || hex.EncodeToString(hasher.Sum(nil)) != image.identity.SHA256 {
		return errRuntimePreparation
	}
	if checkErr := file.Sync(); checkErr != nil {
		return errRuntimePreparation
	}
	return checkRuntimeFile(ctx, image.identity, true)
}

func verifyRuntimeCopy(ctx context.Context, directory string, expected runtimeFile) error {
	root, err := openPopupDirectory(directory)
	if err != nil {
		return errRuntimeStartup
	}
	if checkErr := root.Close(); checkErr != nil {
		return errRuntimeStartup
	}
	image, err := captureRuntimeFile(ctx, filepath.Join(directory, runtimeImageName), true)
	if err != nil {
		return errRuntimeStartup
	}
	stat, statErr := image.file.Stat()
	closeErr := image.file.Close()
	if statErr != nil || closeErr != nil || image.identity.SHA256 != expected.SHA256 || image.identity.Size != expected.Size || image.identity.Owner != uint32(os.Geteuid()) || stat.Mode().Perm() != 0700 {
		return errRuntimeStartup
	}
	owner, ok := stat.Sys().(*syscall.Stat_t)
	if !ok || owner.Nlink != 1 {
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
	request := r
	if request.image == nil {
		return errRuntimeStartup
	}
	if checkErr := request.image.copy(ctx, directory); checkErr != nil {
		return checkErr
	}
	if checkErr := verifyRuntimeCopy(ctx, directory, request.image.identity); checkErr != nil {
		return checkErr
	}
	if request.selection == nil || request.selection.Capabilities == nil || request.selection.Capabilities.Endpoint != request.Server {
		return errRuntimeCompatibility
	}
	imageReceipt, err := json.Marshal(popupImageReceipt{request.image.identity, runtimeHash([]byte(runtimeAssets)), request.selection, request.Server})
	if err != nil || len(imageReceipt) > runtimeLimit {
		return errRuntimeStartup
	}
	return writePopupFile(directory, "runtime.json", imageReceipt)
}

func openVerifiedPopupDirectory(ctx context.Context, directory string) (*os.Root, *rex.CLISelection, error) {
	root, err := openRuntimeDirectory(directory)
	if err != nil {
		return nil, nil, err
	}
	selection, checkErr := verifyPopupImage(ctx, root, directory)
	if checkErr != nil {
		return nil, nil, errors.Join(checkErr, root.Close())
	}
	return root, selection, nil
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

type popupImageReceipt struct {
	Component   runtimeFile       `json:"component"`
	AssetSHA256 string            `json:"asset_sha256"`
	CLI         *rex.CLISelection `json:"cli"`
	Endpoint    string            `json:"endpoint"`
}

func verifyPopupImage(ctx context.Context, root *os.Root, directory string) (_ *rex.CLISelection, outcome error) {
	file, err := root.OpenFile("runtime.json", os.O_RDONLY|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if err != nil {
		return nil, errRuntimeStartup
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	if statErr := checkPopupSlot(file, false); statErr != nil {
		return nil, errRuntimeStartup
	}
	var expected popupImageReceipt
	if checkErr := readRuntime(file, &expected, runtimeLimit); checkErr != nil || expected.AssetSHA256 != runtimeHash([]byte(runtimeAssets)) ||
		expected.CLI == nil || expected.CLI.Capabilities == nil || expected.Endpoint != expected.CLI.Capabilities.Endpoint {
		return nil, errRuntimeStartup
	}
	selection, err := rex.ResolveCLI(ctx, rex.CLIOptions{Selection: expected.CLI})
	if err != nil {
		return nil, errors.Join(errRuntimeCompatibility, err)
	}
	if checkErr := checkRuntimeFile(ctx, expected.Component, true); checkErr != nil {
		return nil, checkErr
	}
	image, err := captureSelf(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { outcome = errors.Join(outcome, image.file.Close()) }()
	if image.identity.Path != filepath.Join(directory, runtimeImageName) || image.identity.SHA256 != expected.Component.SHA256 || image.identity.Size != expected.Component.Size {
		return nil, errRuntimeStartup
	}
	return selection, verifyRuntimeCopy(ctx, directory, expected.Component)
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
