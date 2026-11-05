package rex

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"unicode"
	"unicode/utf8"
)

const maxExecutableBytes = 512 << 20

// CLIOptions are execution-host inputs, never paths supplied by menu text or a
// viewing client. Anchor is the trusted file-loader anchor from the application
// bootstrap, relative only to its unchanged execution-host working directory.
// An explicitly present empty executable/anchor is invalid, not absent.
type CLIOptions struct {
	Selection      *CLISelection
	Executable     *string
	Anchor         *string
	LayoutRevision int
}

// FileIdentity binds a logical installation link to one canonical regular file.
// These fields may travel only in the bootstrap's authenticated private receipt;
// a serialized path or matching digest alone is not execution authority.
type FileIdentity struct {
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

// CLISelection is one frozen execution-host selection. New and Hosts copy it,
// so subsequent caller changes cannot redirect those consumers. Reuse requires
// the original private receipt, not a newly nominated path after detachment.
// Environment provenance is deliberately not a terminal-context assertion.
type CLISelection struct {
	Executable FileIdentity   `json:"executable"`
	Provenance string         `json:"provenance"`
	OS         string         `json:"os"`
	Arch       string         `json:"arch"`
	Version    CLIVersion     `json:"version"`
	Layout     []FileIdentity `json:"layout,omitempty"`
}

// CLIVersion records read-only CLI metadata. Missing API revision fields mean
// zero. Release strings are audit data, not CLI/server equality requirements.
type CLIVersion struct {
	Version        string `json:"version"`
	APIRevision    uint64 `json:"api_revision"`
	APIMinRevision uint64 `json:"api_min_revision"`
}

var (
	errCLIInput         = errors.New("rex: invalid explicit CLI source; supply an absolute execution-host executable or reinstall the integration; not retried")
	errCLIAbsent        = errors.New("rex: no execution-host Rex CLI source; install the integration with its licensed dependency or use a Rex terminal; not retried")
	errCLIPath          = errors.New("rex: PATH must contain only nonempty absolute execution-host directories and one unambiguous Rex executable; not retried")
	errCLILayout        = errors.New("rex: invalid application runtime dependency layout; reinstall the integration for this execution host; not retried")
	errCLIChanged       = errors.New("rex: frozen CLI or runtime asset identity is invalid or changed; start a new invocation after checking the installation; not retried")
	errCLICompatibility = errors.New("rex: unsupported CLI/API metadata; update the execution-host integration and Rex software through normal installation; not retried")
)

// ResolveCLI selects once in precedence order: trusted frozen input, internal
// explicit intent, environment preference, applicable normal runtime dependency,
// then unambiguous absolute PATH. It only runs read-only version metadata after
// selection. An invalid source never permits a lower-precedence candidate.
func ResolveCLI(ctx context.Context, options CLIOptions) (*CLISelection, error) {
	if ctx == nil {
		return nil, errors.New("rex: CLI discovery requires a context")
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, discoveryTimeout)
	defer cancel()
	if options.Selection != nil {
		if options.Executable != nil || options.Anchor != nil || options.LayoutRevision != 0 {
			return nil, errCLIInput
		}
		selection := *options.Selection
		selection.Layout = append([]FileIdentity(nil), selection.Layout...)
		if err := selection.validate(ctx); err != nil {
			return nil, err
		}
		return &selection, nil
	}
	selection := &CLISelection{OS: runtime.GOOS, Arch: runtime.GOARCH}
	var source string
	var err error
	switch {
	case options.Executable != nil:
		if !validAbsolutePath(*options.Executable) {
			return nil, errCLIInput
		}
		source, selection.Provenance = *options.Executable, "internal-explicit"
	case os.Getenv("REX_BIN_DIR") != "":
		directory := os.Getenv("REX_BIN_DIR")
		if !validAbsolutePath(directory) {
			return nil, errCLIInput
		}
		source, selection.Provenance = filepath.Join(directory, "rex"), "execution-env"
	default:
		var applicable bool
		source, selection.Layout, applicable, err = layoutDependency(ctx, options)
		if err != nil {
			return nil, err
		}
		if applicable {
			selection.Provenance = "layout-dependency"
		} else {
			source, err = pathExecutable(ctx)
			if err != nil {
				return nil, err
			}
			selection.Provenance = "path"
		}
	}
	selection.Executable, err = inspectFile(ctx, source, true)
	if err != nil {
		if selection.Provenance == "layout-dependency" {
			return nil, errCLILayout
		}
		return nil, errCLIInput
	}
	output, err := runSelected(ctx, selection, []string{"--autostart=false", "version", "--json"}, nil, 64<<10)
	if err != nil {
		return nil, errors.Join(errCLICompatibility, ctx.Err())
	}
	selection.Version, err = decodeVersion(output)
	if err != nil {
		return nil, err
	}
	return selection, nil
}

// metadataArgs never inherits endpoint intent from a later environment value.
func (c *Client) metadataArgs(arguments ...string) []string {
	argv := []string{"--server", c.server, "--autostart=false", "--timeout", "2s"}
	return append(argv, arguments...)
}

// checkMethod accepts only the concrete API shapes used by this application.
// This is deliberately not a JSON Schema interpreter: references/combinators or
// missing properties are unsupported evidence, not guessed compatibility.
func (c *Client) checkMethod(ctx context.Context, method string) error {
	var arguments, results map[string]string
	switch method {
	case "server.version":
		results = map[string]string{"version": "string"}
	case "server.status":
		results = map[string]string{"instance_id": "string", "config.hash": "string", "config.generation": "integer"}
	case "session.list":
		results = map[string]string{"sessions": "array"}
	case "session.list_clients":
		arguments = map[string]string{"session_id": "string"}
		results = map[string]string{"clients": "array"}
	case "session.describe_block":
		arguments = map[string]string{"session_id": "string", "block_id": "string"}
		results = map[string]string{"block": "object"}
	case "session.focus_block":
		arguments = map[string]string{"session_id": "string", "block_id": "string"}
	case "session.resize_layer":
		arguments = map[string]string{"session_id": "string", "layer_id": "string", "bounds": "object"}
	case "session.new_layer":
		arguments = map[string]string{"session_id": "string", "window_id": "string", "bounds": "object", "focus": "boolean", "layout": "object"}
		results = map[string]string{"layer_id": "string", "block_ids": "array"}
	case "client.perform":
		arguments = map[string]string{"client_id": "string", "action": "string", "args": "object", "timeout_ms": "integer"}
	default:
		return errCLICompatibility
	}
	data, err := runSelected(ctx, c.selection, c.metadataArgs("api", "describe", "--json", "--", method), nil, 64<<10)
	if err != nil {
		return errors.Join(errCLICompatibility, ctx.Err())
	}
	entry, err := metadataObject(data)
	if err != nil || !compatibleSchema(entry["args_schema"], arguments, true) || !compatibleSchema(entry["output_schema"], results, false) {
		return errCLICompatibility
	}
	return nil
}

func compatibleSchema(data json.RawMessage, fields map[string]string, input bool) bool {
	// Known no-argument methods may omit their input schema; this does not
	// permit an absent result schema or an unknown method.
	if input && len(fields) == 0 && (len(data) == 0 || string(data) == "null") {
		return true
	}
	schema, err := metadataObject(data)
	if err != nil || !schemaType(schema, "object") {
		return false
	}
	if input {
		var required []string
		if raw, present := schema["required"]; present {
			if err := json.Unmarshal(raw, &required); err != nil {
				return false
			}
		}
		for _, field := range required {
			if _, known := fields[field]; !known {
				return false
			}
		}
	}
	for path, kind := range fields {
		if !schemaHasField(schema, strings.Split(path, "."), kind) {
			return false
		}
	}
	return true
}

func schemaHasField(schema map[string]json.RawMessage, path []string, kind string) bool {
	properties, err := metadataObject(schema["properties"])
	if err != nil {
		return false
	}
	field, err := metadataObject(properties[path[0]])
	if err != nil {
		return false
	}
	if len(path) == 1 {
		return schemaType(field, kind)
	}
	return schemaType(field, "object") && schemaHasField(field, path[1:], kind)
}

func schemaType(schema map[string]json.RawMessage, kind string) bool {
	for _, keyword := range []string{"$ref", "allOf", "anyOf", "oneOf", "not", "if", "then", "else"} {
		if _, present := schema[keyword]; present {
			return false
		}
	}
	var single string
	if err := json.Unmarshal(schema["type"], &single); err == nil {
		return single == kind
	}
	var kinds []string
	if err := json.Unmarshal(schema["type"], &kinds); err != nil {
		return false
	}
	found := false
	for _, value := range kinds {
		if value != "null" && value != kind {
			return false
		}
		found = found || value == kind
	}
	return found
}

func validAbsolutePath(path string) bool {
	return filepath.IsAbs(path) && validPath(path)
}

func validPath(path string) bool {
	if path == "" || len(path) > 64<<10 || !utf8.ValidString(path) || strings.ContainsRune(path, '\\') || strings.IndexFunc(path, unicode.IsControl) >= 0 {
		return false
	}
	for _, component := range strings.Split(path, string(os.PathSeparator)) {
		if component == ".." {
			return false
		}
	}
	return true
}

func layoutDependency(ctx context.Context, options CLIOptions) (string, []FileIdentity, bool, error) {
	anchor := options.Anchor
	if anchor == nil {
		if options.LayoutRevision != 0 {
			return "", nil, false, errCLILayout
		}
		self, err := os.Executable()
		if err != nil {
			return "", nil, false, errCLILayout
		}
		if !strings.HasSuffix(filepath.ToSlash(self), "/rex_whichkey/runtime/rex-whichkey") {
			return "", nil, false, nil
		}
		candidate := filepath.Join(filepath.Dir(self), "anchor.lua")
		anchor = &candidate
	} else if options.LayoutRevision != 1 {
		return "", nil, true, errCLILayout
	}
	logical := filepath.ToSlash(*anchor)
	if !validPath(*anchor) || (logical != "rex_whichkey/runtime/anchor.lua" && !strings.HasSuffix(logical, "/rex_whichkey/runtime/anchor.lua")) {
		return "", nil, true, errCLILayout
	}
	// Abs uses the process's real cwd, never terminal PWD/HOME. The caller may
	// supply a relative anchor only under the standard-loader unchanged-cwd
	// contract; Go cannot establish provenance for an arbitrary Lua loader.
	absolute, err := filepath.Abs(*anchor)
	if err != nil {
		return "", nil, true, errCLILayout
	}
	anchorFile, err := inspectFile(ctx, absolute, false)
	if err != nil {
		return "", nil, true, errCLILayout
	}
	helper, err := inspectFile(ctx, filepath.Join(filepath.Dir(absolute), "rex-whichkey"), true)
	if err != nil {
		return "", nil, true, errCLILayout
	}
	self, err := os.Executable()
	if err != nil {
		return "", nil, true, errCLILayout
	}
	selfFile, err := inspectFile(ctx, self, true)
	if err != nil || helper.Device != selfFile.Device || helper.Inode != selfFile.Inode || helper.SHA256 != selfFile.SHA256 {
		return "", nil, true, errCLILayout
	}
	return filepath.Join(filepath.Dir(absolute), "deps", "rex"), []FileIdentity{anchorFile, helper}, true, nil
}

func pathExecutable(ctx context.Context) (string, error) {
	path := os.Getenv("PATH")
	if path == "" {
		return "", errCLIAbsent
	}
	var selected *FileIdentity
	// Validate the entire list before accepting a candidate, not just the
	// prefix before the first hit. Bound pathological environment inputs.
	if len(path) > 64<<10 || len(strings.Split(path, string(os.PathListSeparator))) > 256 {
		return "", errCLIPath
	}
	for _, directory := range strings.Split(path, string(os.PathListSeparator)) {
		if !validAbsolutePath(directory) {
			return "", errCLIPath
		}
		candidate := filepath.Join(directory, "rex")
		if _, err := os.Lstat(candidate); errors.Is(err, os.ErrNotExist) {
			continue
		} else if err != nil {
			return "", errCLIPath
		}
		file, err := inspectFile(ctx, candidate, true)
		if err != nil {
			return "", errCLIPath
		}
		if selected != nil && (selected.Device != file.Device || selected.Inode != file.Inode) {
			return "", errCLIPath
		}
		if selected == nil {
			selected = &file
		}
	}
	if selected == nil {
		return "", errCLIAbsent
	}
	return selected.Source, nil
}

func (selection *CLISelection) validate(ctx context.Context) error {
	if selection == nil || selection.OS != runtime.GOOS || selection.Arch != runtime.GOARCH {
		return errCLIChanged
	}
	switch selection.Provenance {
	case "internal-explicit", "execution-env", "terminal-env", "layout-dependency", "path":
	default:
		return errCLIChanged
	}
	if selection.Version.Version == "" || selection.Version.APIMinRevision > selection.Version.APIRevision ||
		(selection.Provenance == "layout-dependency" && len(selection.Layout) != 2) ||
		(selection.Provenance != "layout-dependency" && len(selection.Layout) != 0) {
		return errCLIChanged
	}
	return selection.checkFiles(ctx)
}

func (selection *CLISelection) checkFiles(ctx context.Context) error {
	if selection == nil {
		return errCLIChanged
	}
	if err := checkFile(ctx, selection.Executable, true); err != nil {
		return err
	}
	for index, file := range selection.Layout {
		if err := checkFile(ctx, file, index == 1); err != nil {
			return err
		}
	}
	return nil
}

func checkFile(ctx context.Context, expected FileIdentity, executable bool) error {
	actual, err := inspectFile(ctx, expected.Source, executable)
	if err != nil || actual != expected {
		return errors.Join(errCLIChanged, ctx.Err())
	}
	return nil
}

func inspectFile(ctx context.Context, source string, executable bool) (identity FileIdentity, resultErr error) {
	if !validAbsolutePath(source) || !safeParents(source) {
		return identity, errCLIChanged
	}
	links, err := installationLinks(ctx, source)
	if err != nil {
		return identity, errCLIChanged
	}
	canonical, err := filepath.EvalSymlinks(source)
	if err != nil || !validAbsolutePath(canonical) || !safeParents(canonical) {
		return identity, errCLIChanged
	}
	file, err := os.OpenFile(canonical, os.O_RDONLY|syscall.O_NONBLOCK|syscall.O_NOFOLLOW, 0)
	if err != nil {
		return identity, errCLIChanged
	}
	defer func() {
		if err := file.Close(); err != nil {
			resultErr = errCLIChanged
		}
	}()
	before, err := file.Stat()
	if err != nil || !before.Mode().IsRegular() || before.Size() <= 0 || before.Size() > maxExecutableBytes ||
		before.Mode().Perm()&0022 != 0 || (executable && before.Mode().Perm()&0111 == 0) {
		return identity, errCLIChanged
	}
	stat, ok := before.Sys().(*syscall.Stat_t)
	if !ok || (stat.Uid != 0 && stat.Uid != uint32(os.Geteuid())) {
		return identity, errCLIChanged
	}
	digest := sha256.New()
	buffer := make([]byte, 128<<10)
	var count int64
	for {
		if err := ctx.Err(); err != nil {
			return identity, err
		}
		n, readErr := file.Read(buffer)
		count += int64(n)
		if count > maxExecutableBytes {
			return identity, errCLIChanged
		}
		if _, err := digest.Write(buffer[:n]); err != nil {
			return identity, errCLIChanged
		}
		if errors.Is(readErr, io.EOF) {
			break
		}
		if readErr != nil {
			return identity, errCLIChanged
		}
	}
	after, err := file.Stat()
	linked, linkErr := os.Stat(source)
	resolved, resolveErr := filepath.EvalSymlinks(source)
	afterLinks, linksErr := installationLinks(ctx, source)
	if err != nil || linkErr != nil || resolveErr != nil || linksErr != nil || afterLinks != links || resolved != canonical || !os.SameFile(before, after) ||
		!os.SameFile(after, linked) || count != before.Size() || before.Size() != after.Size() ||
		before.Mode() != after.Mode() || before.ModTime() != after.ModTime() {
		return identity, errCLIChanged
	}
	return FileIdentity{Source: source, Path: canonical, SHA256: hex.EncodeToString(digest.Sum(nil)), LinksSHA256: links,
		Device: uint64(stat.Dev), Inode: stat.Ino, Size: count, ModifiedNS: before.ModTime().UnixNano(),
		Mode: uint32(before.Mode()), Owner: stat.Uid}, nil
}

// Include logical directory links and intermediate link targets, not only the
// final file. Replacing an HM namespace link with even an alias to the same file
// must invalidate the receipt instead of silently changing installation identity.
func installationLinks(ctx context.Context, source string) (string, error) {
	digest := sha256.New()
	remaining := filepath.Clean(source)
	for links := 0; links < 256; links++ {
		if err := ctx.Err(); err != nil {
			return "", err
		}
		if !safeParents(remaining) {
			return "", errCLIChanged
		}
		components := strings.Split(strings.TrimPrefix(remaining, string(os.PathSeparator)), string(os.PathSeparator))
		prefix := string(os.PathSeparator)
		found := false
		for index, component := range components {
			prefix = filepath.Join(prefix, component)
			info, err := os.Lstat(prefix)
			if err != nil {
				return "", errCLIChanged
			}
			if info.Mode()&os.ModeSymlink == 0 {
				continue
			}
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok || (stat.Uid != 0 && stat.Uid != uint32(os.Geteuid())) {
				return "", errCLIChanged
			}
			target, err := os.Readlink(prefix)
			if err != nil {
				return "", errCLIChanged
			}
			if _, err := fmt.Fprintf(digest, "%q:%d:%d:%d:%d:%d:%q\n", prefix, stat.Dev, stat.Ino, stat.Uid, info.Size(), info.ModTime().UnixNano(), target); err != nil {
				return "", errCLIChanged
			}
			if !filepath.IsAbs(target) {
				target = filepath.Join(filepath.Dir(prefix), target)
			}
			remaining = filepath.Join(append([]string{target}, components[index+1:]...)...)
			found = true
			break
		}
		if !found {
			return hex.EncodeToString(digest.Sum(nil)), nil
		}
	}
	return "", errCLIChanged
}

func safeParents(path string) bool {
	for directory := filepath.Dir(path); ; directory = filepath.Dir(directory) {
		info, err := os.Stat(directory)
		if err != nil || !info.IsDir() {
			return false
		}
		stat, ok := info.Sys().(*syscall.Stat_t)
		if !ok || (stat.Uid != 0 && stat.Uid != uint32(os.Geteuid())) {
			return false
		}
		// Root-owned sticky temporary directories do not let unrelated users
		// replace a trusted child. Other group/world writers are unsupported.
		if info.Mode().Perm()&0022 != 0 && !(stat.Uid == 0 && info.Mode()&os.ModeSticky != 0) {
			return false
		}
		if filepath.Dir(directory) == directory {
			return true
		}
	}
}
