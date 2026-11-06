package whichkey

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"reflect"
	"runtime"
	"slices"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"

	"golang.org/x/sys/unix"

	"temp/rex"
)

const (
	runtimeProtocol     = 1
	runtimeArgsField    = "args"
	runtimeLayout       = 1
	runtimeLimit        = 8 << 20
	runtimeImageLimit   = 512 << 20
	runtimeImageName    = "runtime-image"
	runtimePreflight    = "preflight"
	runtimeBoundFile    = "bound.json"
	runtimeNoAutostart  = "--autostart=false"
	runtimeServerFlag   = "--server"
	runtimeLeaseFile    = "lease"
	runtimeRequestFile  = "request.json"
	runtimeMenuFile     = "menu.json"
	runtimePreparedFile = "prepared.json"
	runtimeTimeoutFlag  = "--timeout"
	runtimeStageTimeout = "3s"
	runtimeJSONFlag     = "--json"
	runtimeAccepted     = "accepted"
	// This embedded identity describes Go-owned formats, not the installed Lua.
	runtimeAssets = "rex-whichkey/runtime/1;config/1;completion/2;selection-text-resource-rename/1;relay-string/1;startup-ack/1;pty-cli/1"
)

var (
	errRuntimeLayout        = errors.New("whichkey: helper/layout discovery failed; install the normal execution-host integration assets through the installer; not retried")
	errRuntimeCompatibility = errors.New("whichkey: runtime contract not established; inspect the named contract before a new invocation; no activation, retry or native fallback")
	errRuntimeOrigin        = errors.New("whichkey: runtime origin or generation changed; relaunch from the original Rex terminal; not retried")
	errRuntimeEndpoint      = errors.New("whichkey: constructor endpoint intent is invalid or conflicts with the retained public endpoint; use matching constructor and public launch endpoints; not retried")
	errRuntimeRelay         = errors.New("whichkey: private rendezvous failed; run on the server execution host with shared private files and current integration software; no retry or native fallback")
	errRuntimePreparation   = errors.New("whichkey: runtime preparation failed; check the execution-host integration and original invocation deadline; not retried")
	errRuntimeStartup       = errors.New("whichkey: controller/child startup failed or uncertain; inspect the original invocation before relaunching; not retried")
)

func runtimeFeatures() []string {
	return []string{"original-vm", "separate-caller-target", "single-use", "generation-bound", "retained-menu", "public-file-relay", "completion-receipt-v2", "selection-text-resource-rename"}
}

// runtimeConsumer describes the actual loaded registry, not files on disk. Build
// IDs are audit identities; compatibility is the protocol and feature contract.
type runtimeConsumer struct {
	ProtocolMin      int      `json:"protocol_min"`
	ProtocolMax      int      `json:"protocol_max"`
	Build            string   `json:"consumer_build"`
	Features         []string `json:"features"`
	Launch           string   `json:"launch_action"`
	Completion       string   `json:"completion_action"`
	MenuGeneration   string   `json:"menu_generation"`
	ServerInstance   string   `json:"server_instance"`
	ConfigGeneration uint64   `json:"config_generation"`
	ConfigHash       string   `json:"config_hash"`
	LayoutRevision   int      `json:"layout_revision"`
}

func runtimeName(value string) bool {
	return value != "" && len(value) <= 256 && strings.Trim(value, "abcdefghijklmnopqrstuvwxyz0123456789_") == ""
}

func runtimeString(value string) bool {
	return value != "" && len(value) <= 1024 && strings.TrimSpace(value) != "" && validateText(value) == nil
}

func (consumer runtimeConsumer) validate(action string, selection *rex.CLISelection) error {
	if consumer.ProtocolMin < 1 || consumer.ProtocolMin > runtimeProtocol || consumer.ProtocolMax < runtimeProtocol ||
		consumer.ProtocolMax < consumer.ProtocolMin || consumer.LayoutRevision != runtimeLayout ||
		consumer.Launch != action || !runtimeName(action) || !runtimeName(consumer.Completion) {
		return errRuntimeCompatibility
	}
	for _, value := range []string{consumer.Build, consumer.MenuGeneration, consumer.ServerInstance, consumer.ConfigHash} {
		if !runtimeString(value) {
			return errRuntimeCompatibility
		}
	}
	if checkErr := validateRuntimeFeatures(consumer.Features); checkErr != nil {
		return checkErr
	}
	return consumer.checkServer(selection)
}

func validateRuntimeFeatures(features []string) error {
	if len(features) > 64 {
		return errRuntimeCompatibility
	}
	seen := make(map[string]bool, len(features))
	for _, feature := range features {
		if !slices.Contains(runtimeFeatures(), feature) || seen[feature] {
			return errRuntimeCompatibility
		}
		seen[feature] = true
	}
	for _, required := range runtimeFeatures()[:7] {
		if !seen[required] {
			return errRuntimeCompatibility
		}
	}
	return nil
}

func negotiatedRuntimeFeatures(consumer runtimeConsumer) []string {
	features := runtimeFeatures()
	if !slices.Contains(consumer.Features, features[7]) {
		return features[:7]
	}
	return features
}

func (consumer runtimeConsumer) checkServer(selection *rex.CLISelection) error {
	if selection == nil || selection.Capabilities == nil {
		return errRuntimeCompatibility
	}
	server := selection.Capabilities.Server
	if consumer.ServerInstance != server.InstanceID || consumer.ConfigGeneration != server.Generation || consumer.ConfigHash != server.ConfigHash {
		return errRuntimeOrigin
	}
	return nil
}

// UnmarshalJSON requires every advertisement field, including a real integral
// configuration generation (zero is not a default for a missing field).
func (consumer *runtimeConsumer) UnmarshalJSON(data []byte) error {
	fields, err := completionObject(data, "protocol_min", "protocol_max", "consumer_build", "features", "launch_action", "completion_action", "menu_generation", "server_instance", "config_generation", "config_hash", "layout_revision")
	if err != nil || len(fields) != 11 {
		return errRuntimeCompatibility
	}
	type wire runtimeConsumer
	var candidate wire
	if checkErr := json.Unmarshal(data, &candidate); checkErr != nil {
		return errRuntimeCompatibility
	}
	slices.Sort(candidate.Features)
	*consumer = runtimeConsumer(candidate)
	return nil
}

// Strings retain exact JSON bytes across Go/Lua. Neither Lua nor a relay
// reconstructs a competing receipt serialization to compute this digest.
type runtimeReceipt struct {
	Protocol            int                  `json:"protocol"`
	Features            []string             `json:"features"`
	AssetSHA256         string               `json:"asset_sha256"`
	Component           runtimeFile          `json:"component"`
	Layout              *runtimeInstallation `json:"layout,omitempty"`
	Relay               *runtimeInstallation `json:"relay,omitempty"`
	CLI                 *rex.CLISelection    `json:"cli"`
	Endpoint            string               `json:"endpoint"`
	Consumer            runtimeConsumer      `json:"consumer"`
	Origin              runtimeOrigin        `json:"origin"`
	Expires             int64                `json:"expires"`
	Workspace           string               `json:"workspace,omitempty"`
	ConfigurationSHA256 string               `json:"configuration_sha256,omitempty"`
	OS                  string               `json:"os"`
	Arch                string               `json:"arch"`
}

type runtimeOrigin struct {
	Session     string `json:"session_id"`
	Block       string `json:"block_id"`
	Window      string `json:"window_id"`
	Layer       string `json:"layer_id"`
	Caller      string `json:"caller_id"`
	Client      string `json:"client_id"`
	ConnectedAt string `json:"connected_at"`
}

func (origin *runtimeOrigin) UnmarshalJSON(data []byte) error {
	fields, err := completionObject(data, "session_id", "block_id", "window_id", "layer_id", "caller_id", "client_id", "connected_at")
	if err != nil || len(fields) != 7 {
		return errRuntimeOrigin
	}
	type wire runtimeOrigin
	var candidate wire
	if checkErr := json.Unmarshal(data, &candidate); checkErr != nil {
		return errRuntimeOrigin
	}
	*origin = runtimeOrigin(candidate)
	return nil
}

func (origin runtimeOrigin) validate(requireCaller bool) error {
	if !popupResourceID(origin.Session, "session") || !popupResourceID(origin.Block, "block") ||
		!popupResourceID(origin.Window, "window") || (origin.Layer != "" && !popupResourceID(origin.Layer, "layer")) ||
		!popupResourceID(origin.Client, "client") || !popupConnectionGeneration(origin.ConnectedAt) ||
		(requireCaller && !popupResourceID(origin.Caller, "client")) {
		return errRuntimeOrigin
	}
	return nil
}

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
	return checkRuntimeValue(json.NewDecoder(bytes.NewReader(data)), false)
}

func checkRuntimeValue(decoder *json.Decoder, args bool) error {
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
		if checkErr := checkRuntimeValue(decoder, childArgs); checkErr != nil {
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

// All runtime consumers check decoded field identity before struct decoding,
// including nested receipts, slot payloads, startup messages and PTY receipts.
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

func writeRuntime(output io.Writer, value any, limit int) error {
	if file, ok := output.(*os.File); ok {
		info, err := file.Stat()
		if err != nil || info.Mode()&os.ModeCharDevice != 0 {
			return errRuntimePreparation
		}
	}
	data, err := json.Marshal(value)
	if err != nil || len(data) > limit {
		return errRuntimePreparation
	}
	if _, err := output.Write(data); err != nil {
		return errRuntimePreparation
	}
	return nil
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

// runtimeInstallation is the fixed normal layout, never an executable nomination.
type runtimeInstallation struct {
	Revision   int         `json:"revision"`
	Anchor     runtimeFile `json:"anchor"`
	Helper     runtimeFile `json:"helper"`
	Dependency runtimeFile `json:"dependency"`
}

func captureRuntimeAsset(ctx context.Context, path string, executable bool) (runtimeFile, error) {
	image, err := captureRuntimeFile(ctx, path, executable)
	if err != nil {
		return runtimeFile{}, err
	}
	if checkErr := image.file.Close(); checkErr != nil {
		return runtimeFile{}, errRuntimeLayout
	}
	return image.identity, nil
}

func captureInstallation(ctx context.Context, anchor string, revision int, self runtimeFile) (*runtimeInstallation, error) {
	if !runtimePath(anchor) || revision != runtimeLayout ||
		(anchor != "rex_whichkey/runtime/anchor.lua" && !strings.HasSuffix(filepath.ToSlash(anchor), "/rex_whichkey/runtime/anchor.lua")) {
		return nil, errRuntimeLayout
	}
	absolute, err := filepath.Abs(anchor) // actual unchanged process cwd, never PWD
	if err != nil {
		return nil, errRuntimeLayout
	}
	anchorFile, err := captureRuntimeAsset(ctx, absolute, false)
	if err != nil {
		return nil, err
	}
	helper, err := captureRuntimeAsset(ctx, filepath.Join(filepath.Dir(absolute), "rex-whichkey"), true)
	if err != nil || helper.Path != self.Path || helper.SHA256 != self.SHA256 || helper.Device != self.Device || helper.Inode != self.Inode {
		return nil, errRuntimeLayout
	}
	dependency, err := captureRuntimeAsset(ctx, filepath.Join(filepath.Dir(absolute), "deps", "rex"), true)
	if err != nil {
		return nil, err
	}
	return &runtimeInstallation{runtimeLayout, anchorFile, helper, dependency}, nil
}

func (installation *runtimeInstallation) check(ctx context.Context) error {
	if installation == nil || installation.Revision != runtimeLayout {
		return errRuntimeLayout
	}
	for _, entry := range []struct {
		file       runtimeFile
		executable bool
	}{{installation.Anchor, false}, {installation.Helper, true}, {installation.Dependency, true}} {
		if checkErr := checkRuntimeFile(ctx, entry.file, entry.executable); checkErr != nil {
			return checkErr
		}
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

func (receipt runtimeReceipt) check(ctx context.Context, caller bool) error {
	if checkErr := receipt.validate(caller); checkErr != nil {
		return checkErr
	}
	if checkErr := checkRuntimeFile(ctx, receipt.Component, true); checkErr != nil {
		return checkErr
	}
	if receipt.Layout != nil {
		if checkErr := receipt.Layout.check(ctx); checkErr != nil {
			return checkErr
		}
	}
	if receipt.Relay != nil {
		if checkErr := receipt.Relay.check(ctx); checkErr != nil {
			return checkErr
		}
	}
	_, err := rex.ResolveCLI(ctx, rex.CLIOptions{Selection: receipt.CLI})
	return err
}

func (receipt runtimeReceipt) validate(caller bool) error {
	if receipt.Protocol != runtimeProtocol || receipt.AssetSHA256 != runtimeHash([]byte(runtimeAssets)) ||
		receipt.OS != runtime.GOOS || receipt.Arch != runtime.GOARCH || !reflect.DeepEqual(receipt.Features, negotiatedRuntimeFeatures(receipt.Consumer)) {
		return errRuntimeCompatibility
	}
	if checkErr := receipt.validateEndpoint(); checkErr != nil {
		return checkErr
	}
	now := time.Now()
	if receipt.Expires <= now.Unix() || receipt.Expires > now.Add(InvocationLifetimeSeconds*time.Second).Unix() {
		return errRuntimePreparation
	}
	if checkErr := receipt.Origin.validate(caller); checkErr != nil {
		return checkErr
	}
	return receipt.Consumer.validate(receipt.Consumer.Launch, receipt.CLI)
}

func (receipt runtimeReceipt) validateEndpoint() error {
	endpoint, err := rex.CanonicalEndpoint(receipt.Endpoint)
	if err != nil || endpoint != receipt.Endpoint || receipt.CLI == nil || receipt.CLI.Capabilities == nil || receipt.CLI.Capabilities.Endpoint != endpoint {
		return errRuntimeCompatibility
	}
	return nil
}

func encodeReceipt(receipt runtimeReceipt) (string, string, error) {
	data, err := json.Marshal(receipt)
	if err != nil || len(data) > runtimeLimit {
		return "", "", errRuntimePreparation
	}
	return string(data), runtimeHash(data), nil
}

func decodeReceipt(data, hash string) (runtimeReceipt, error) {
	var receipt runtimeReceipt
	if ValidateInvocationCapability(hash) != nil || runtimeHash([]byte(data)) != hash || decodeRuntime([]byte(data), &receipt, runtimeLimit) != nil {
		return receipt, errRuntimeCompatibility
	}
	return receipt, nil
}

func runtimeClient(ctx context.Context, receipt runtimeReceipt) (*rex.Client, error) {
	return rex.NewContext(ctx, rex.Options{Selection: receipt.CLI, Server: receipt.Endpoint, Session: receipt.Origin.Session, Timeout: 3 * time.Second})
}

func (receipt runtimeReceipt) checkOrigin(ctx context.Context) error {
	client, err := runtimeClient(ctx, receipt)
	if err != nil {
		return errors.Join(errRuntimeCompatibility, err)
	}
	request := popupRequest{Server: receipt.Endpoint, SessionID: receipt.Origin.Session, BlockID: receipt.Origin.Block, ClientID: receipt.Origin.Client, ConnectedAt: receipt.Origin.ConnectedAt}
	block, err := request.originBlock(ctx, client)
	if err != nil || block.WindowID != receipt.Origin.Window || block.LayerID != receipt.Origin.Layer {
		return errRuntimeOrigin
	}
	return nil
}

// RuntimeDescribe is read-only and emits no ticket/capability. For key launches
// the fixed installed helper captures the CLI, endpoint and component before Lua
// reserves anything. The caller must compare the original VM's direct metadata.
func RuntimeDescribe(ctx context.Context, input io.Reader, output io.Writer) (outcome error) {
	var request struct {
		Anchor         string          `json:"anchor"`
		LayoutRevision int             `json:"layout_revision"`
		Descriptor     string          `json:"descriptor,omitempty"`
		Server         *string         `json:"server,omitempty"`
		Consumer       runtimeConsumer `json:"consumer"`
		Origin         runtimeOrigin   `json:"origin"`
	}
	if checkErr := readRuntime(input, &request, popupMessageLimit); checkErr != nil {
		return checkErr
	}
	if request.Descriptor != "" {
		data, err := json.Marshal(runtimeRelayRequest{Operation: runtimePreflight, Descriptor: request.Descriptor, Anchor: request.Anchor, LayoutRevision: request.LayoutRevision, Consumer: request.Consumer, Origin: request.Origin, Server: request.Server})
		if err != nil {
			return errRuntimeCompatibility
		}
		return RuntimeRelay(ctx, bytes.NewReader(data), output)
	}
	image, err := captureSelf(ctx)
	if err != nil {
		return err
	}
	defer func() { outcome = errors.Join(outcome, image.file.Close()) }()
	layout, err := captureInstallation(ctx, request.Anchor, request.LayoutRevision, image.identity)
	if err != nil {
		return err
	}
	selection, err := rex.ResolveCLI(ctx, rex.CLIOptions{Anchor: &request.Anchor, LayoutRevision: request.LayoutRevision})
	if err != nil {
		return err
	}
	endpoint, err := rex.ResolveEndpoint(ctx, request.Server, selection)
	if err != nil {
		return err
	}
	client, err := rex.NewContext(ctx, rex.Options{Selection: selection, Server: endpoint, Session: request.Origin.Session, Timeout: 3 * time.Second})
	if err != nil {
		return err
	}
	receipt := runtimeReceipt{Protocol: runtimeProtocol, Features: negotiatedRuntimeFeatures(request.Consumer), AssetSHA256: runtimeHash([]byte(runtimeAssets)), Component: image.identity, Layout: layout,
		CLI: client.Selection(), Endpoint: endpoint, Consumer: request.Consumer, Origin: request.Origin, Expires: time.Now().Unix() + InvocationLifetimeSeconds, OS: runtime.GOOS, Arch: runtime.GOARCH}
	if checkErr := receipt.check(ctx, true); checkErr != nil {
		return checkErr
	}
	if checkErr := receipt.checkOrigin(ctx); checkErr != nil {
		return checkErr
	}
	data, hash, err := encodeReceipt(receipt)
	if err != nil {
		return err
	}
	return writeRuntime(output, struct {
		ProtocolMin    int      `json:"protocol_min"`
		ProtocolMax    int      `json:"protocol_max"`
		Build          string   `json:"producer_build"`
		Features       []string `json:"features"`
		LayoutRevision int      `json:"layout_revision"`
		Receipt        string   `json:"receipt"`
		ReceiptHash    string   `json:"receipt_hash"`
	}{runtimeProtocol, runtimeProtocol, image.identity.SHA256, runtimeFeatures(), runtimeLayout, data, hash}, runtimeLimit)
}

// runtimeCommand is a bounded runner for the few existing CLI facilities not
// exposed as resource wrappers. It reuses/rechecks selection; it never resolves
// a path or endpoint policy, or returns private child diagnostics.
func runtimeCommand(ctx context.Context, selection *rex.CLISelection, scope launchOrigin, arguments []string, input io.Reader, output io.Writer) error {
	if selection == nil || selection.Capabilities == nil || selection.Capabilities.Endpoint != scope.Server || !popupResourceID(scope.SessionID, "session") || !popupResourceID(scope.BlockID, "block") {
		return errRuntimeCompatibility
	}
	if _, err := rex.ResolveCLI(ctx, rex.CLIOptions{Selection: selection}); err != nil {
		return err
	}
	command := exec.CommandContext(ctx, selection.Executable.Path, arguments...)
	for _, entry := range os.Environ() {
		name, _, _ := strings.Cut(entry, "=")
		switch name {
		case "REX_SERVER", "REX_CLIENT", "REX_SESSION", "REX_WINDOW", "REX_BLOCK":
			continue
		}
		command.Env = append(command.Env, entry)
	}
	command.Env = append(command.Env, "REX_SESSION="+scope.SessionID, "REX_BLOCK="+scope.BlockID)
	command.Stdin, command.Stdout = input, output
	diagnostics := &launchOutput{remaining: popupMessageLimit}
	command.Stderr = diagnostics
	command.WaitDelay = 250 * time.Millisecond
	if checkErr := command.Run(); checkErr != nil {
		if diagnostics.legacyClientBinding {
			return errRuntimeCompatibility
		}
		return errRuntimeStartup
	}
	return nil
}

func runtimeAction(ctx context.Context, selection *rex.CLISelection, origin launchOrigin, action string, input []byte, output io.Writer) error {
	arguments := []string{runtimeServerFlag, origin.Server, runtimeNoAutostart, runtimeTimeoutFlag, runtimeStageTimeout}
	if origin.ClientID != "" {
		arguments = append(arguments, "--client", origin.ClientID)
	}
	arguments = append(arguments, "do", action, "--session", origin.SessionID, "--block", origin.BlockID, "--args", "-", runtimeJSONFlag)
	return runtimeCommand(ctx, selection, origin, arguments, bytes.NewReader(input), output)
}

func queryRuntimeConsumer(ctx context.Context, selection *rex.CLISelection, origin launchOrigin, action string) (runtimeConsumer, error) {
	input, err := json.Marshal(struct {
		Min    int    `json:"protocol_min"`
		Max    int    `json:"protocol_max"`
		Action string `json:"action"`
	}{runtimeProtocol, runtimeProtocol, action})
	if err != nil {
		return runtimeConsumer{}, errRuntimeCompatibility
	}
	output := &launchBuffer{remaining: popupMessageLimit}
	query, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if checkErr := runtimeAction(query, selection, origin, "whichkey_runtime", input, output); checkErr != nil {
		// runtimeCommand already suppresses child diagnostics; only its safe
		// cause and the bounded query's cancellation may cross this boundary.
		return runtimeConsumer{}, fmt.Errorf("whichkey: whichkey_runtime action call failed: %w", errors.Join(errRuntimeCompatibility, checkErr, query.Err()))
	}
	var consumer runtimeConsumer
	if checkErr := decodeRuntime(bytes.TrimSpace(output.buffer.Bytes()), &consumer, popupMessageLimit); checkErr != nil {
		return consumer, fmt.Errorf("whichkey: whichkey_runtime reply-schema decode failed: %w", errRuntimeCompatibility)
	}
	if checkErr := consumer.validate(action, selection); checkErr != nil {
		return consumer, fmt.Errorf("whichkey: whichkey_runtime consumer validation failed: %w", errors.Join(errRuntimeCompatibility, checkErr))
	}
	return consumer, nil
}

// prepareRuntime is shared by public startup and fixed-helper key preparation.
// The only executable copied is this selected process's captured open image.
type popupPreparation struct {
	Version     int    `json:"version"`
	Capability  string `json:"capability"`
	Executable  string `json:"executable"`
	Server      string `json:"server"`
	Receipt     string `json:"receipt"`
	ReceiptHash string `json:"receipt_hash"`
	Workspace   string `json:"workspace"`
	Expires     int64  `json:"expires"`
}

// Keep the standalone Config boundary unchanged. Runtime-carried Config adds
// exact protocol names and scalar Unicode checks, but preserves arbitrary args
// keys at every depth and delegates all Config presence/type rules to its owner.
func decodeRuntimeConfig(input io.Reader) (Config, error) {
	const limit = 4 << 20
	data, err := io.ReadAll(io.LimitReader(input, limit+1))
	if err != nil || len(data) > limit {
		return Config{}, errRuntimeCompatibility
	}
	if checkErr := checkRuntimeJSON(bytes.TrimSpace(data), limit); checkErr != nil {
		return Config{}, checkErr
	}
	return DecodeConfig(bytes.NewReader(data))
}

func validateRuntimeConfig(configuration string, consumer runtimeConsumer) error {
	config, err := decodeRuntimeConfig(strings.NewReader(configuration))
	if err != nil {
		return errRuntimePreparation
	}
	for _, mapping := range config.Spec {
		if mapping.Enabled() && nativeRenameMapping(mapping) && !slices.Contains(consumer.Features, "selection-text-resource-rename") {
			return errRuntimeCompatibility
		}
	}
	return nil
}

func prepareRuntime(ctx context.Context, image *runtimeImage, receipt runtimeReceipt, configuration string, workspace string) (popupPreparation, error) {
	if checkErr := receipt.check(ctx, true); checkErr != nil {
		return popupPreparation{}, checkErr
	}
	if image.identity != receipt.Component || !time.Unix(receipt.Expires, 0).Add(-5*time.Second).After(time.Now()) {
		return popupPreparation{}, errRuntimePreparation
	}
	if checkErr := validateRuntimeConfig(configuration, receipt.Consumer); checkErr != nil {
		return popupPreparation{}, checkErr
	}
	if checkErr := receipt.checkOrigin(ctx); checkErr != nil {
		return popupPreparation{}, checkErr
	}
	if checkErr := image.copy(ctx, workspace); checkErr != nil {
		return popupPreparation{}, checkErr
	}
	if checkErr := verifyRuntimeCopy(ctx, workspace, receipt.Component); checkErr != nil {
		return popupPreparation{}, checkErr
	}
	capability, err := NewInvocationCapability()
	if err != nil {
		return popupPreparation{}, errRuntimePreparation
	}
	receipt.Workspace, receipt.ConfigurationSHA256 = workspace, runtimeHash([]byte(configuration))
	data, hash, err := encodeReceipt(receipt)
	if err != nil {
		return popupPreparation{}, err
	}
	return popupPreparation{runtimeProtocol, capability, filepath.Join(workspace, runtimeImageName), receipt.Endpoint, data, hash, workspace, receipt.Expires}, nil
}

func prepareRuntimeInput(ctx context.Context, input io.Reader, output io.Writer, server *string) (outcome error) {
	var request struct {
		Version       int    `json:"version"`
		Configuration string `json:"configuration"`
		Receipt       string `json:"receipt"`
		ReceiptHash   string `json:"receipt_hash"`
	}
	if checkErr := readRuntime(input, &request, runtimeLimit); checkErr != nil || request.Version != runtimeProtocol {
		return errors.New("whichkey: preparation requires the negotiated runtime receipt; remove obsolete execution.helper/execution.rex_bin_dir wiring and update the integration through its normal lifecycle")
	}
	receipt, err := decodeReceipt(request.Receipt, request.ReceiptHash)
	if err != nil {
		return err
	}
	if server != nil {
		endpoint, endpointErr := rex.CanonicalEndpoint(*server)
		if endpointErr != nil || endpoint != receipt.Endpoint {
			return errRuntimeOrigin
		}
	}
	if receipt.Layout == nil || receipt.Relay != nil {
		return errRuntimeLayout
	}
	image, err := captureSelf(ctx)
	if err != nil {
		return err
	}
	defer func() { outcome = errors.Join(outcome, image.file.Close()) }()
	workspace, err := newRuntimeDirectory("rex-whichkey-runtime-")
	if err != nil {
		return errRuntimePreparation
	}
	retained := false
	defer func() {
		if !retained {
			outcome = errors.Join(outcome, removeRuntimeWorkspace(workspace))
		}
	}()
	prepared, err := prepareRuntime(ctx, image, receipt, request.Configuration, workspace)
	if err != nil {
		return err
	}
	if checkErr := writeRuntime(output, prepared, runtimeLimit); checkErr != nil {
		return checkErr
	}
	retained = true // Lua owns cancellation until the controller has started.
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

func removeRuntimeWorkspace(directory string) error {
	root, err := openPopupDirectory(directory)
	if err != nil {
		return errRuntimeRelay
	}
	var outcome error
	for _, name := range []string{runtimeRequestFile, runtimeMenuFile, runtimePreparedFile, runtimeBoundFile, runtimeLeaseFile, runtimeImageName} {
		if checkErr := root.Remove(name); checkErr != nil && !errors.Is(checkErr, os.ErrNotExist) {
			outcome = errors.Join(outcome, errRuntimeRelay)
		}
	}
	if checkErr := root.Close(); checkErr != nil {
		outcome = errors.Join(outcome, errRuntimeRelay)
	}
	if checkErr := os.Remove(directory); checkErr != nil {
		outcome = errors.Join(outcome, errRuntimeRelay)
	}
	return outcome
}

// Rendezvous files are created by the public owner, never by a relay. A locator
// names only that fixed workspace; it does not nominate a file or executable.
type runtimeLocator struct {
	Version       int    `json:"version"`
	Workspace     string `json:"workspace"`
	RequestSHA256 string `json:"request_sha256"`
	ReceiptHash   string `json:"receipt_hash"`
}

type runtimeSlotIdentity struct {
	Name   string `json:"name"`
	Device uint64 `json:"device"`
	Inode  uint64 `json:"inode"`
}

type runtimeDescriptor struct {
	Version     int                   `json:"version"`
	Marker      string                `json:"marker"`
	Workspace   string                `json:"workspace"`
	Receipt     string                `json:"receipt"`
	ReceiptHash string                `json:"receipt_hash"`
	Slots       []runtimeSlotIdentity `json:"slots"`
}

type runtimeEnvelope struct {
	Stage       string          `json:"stage"`
	ReceiptHash string          `json:"receipt_hash"`
	Payload     json.RawMessage `json:"payload"`
}

type runtimeMenu struct {
	Configuration string              `json:"configuration"`
	Consumer      runtimeConsumer     `json:"consumer"`
	Origin        runtimeOrigin       `json:"origin"`
	Relay         runtimeInstallation `json:"relay"`
	Server        *string             `json:"server,omitempty"`
}

type runtimeBinding struct {
	ReceiptHash string `json:"receipt_hash"`
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

func runtimeSlotStat(file *os.File, name string, fifo bool) (runtimeSlotIdentity, error) {
	info, err := file.Stat()
	if err != nil || info.Mode().Perm() != 0600 || (fifo && info.Mode()&os.ModeNamedPipe == 0) || (!fifo && !info.Mode().IsRegular()) {
		return runtimeSlotIdentity{}, errRuntimeRelay
	}
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok || stat.Uid != uint32(os.Geteuid()) || stat.Nlink != 1 {
		return runtimeSlotIdentity{}, errRuntimeRelay
	}
	return runtimeSlotIdentity{name, uint64(stat.Dev), stat.Ino}, nil
}

func openRuntimeSlot(root *os.Root, name string, flags int, expected []runtimeSlotIdentity) (*os.File, error) {
	file, err := root.OpenFile(name, flags|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if err != nil {
		return nil, errRuntimeRelay
	}
	identity, err := runtimeSlotStat(file, name, name == runtimeLeaseFile)
	if err != nil {
		return nil, errors.Join(errRuntimeRelay, file.Close())
	}
	for _, entry := range expected {
		if entry == identity {
			return file, nil
		}
	}
	return nil, errors.Join(errRuntimeRelay, file.Close())
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

func readRuntimeSlot(ctx context.Context, root *os.Root, name string, expected []runtimeSlotIdentity) (data []byte, outcome error) {
	file, err := openRuntimeSlot(root, name, os.O_RDONLY, expected)
	if err != nil {
		return nil, err
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	if checkErr := runtimeLock(ctx, file, false); checkErr != nil {
		return nil, checkErr
	}
	// Closing releases the lock even if a read fails. Never retain a partial
	// JSON value after the writer's lock has been released.
	data, err = io.ReadAll(io.LimitReader(file, runtimeLimit+1))
	if err != nil || len(data) > runtimeLimit {
		return nil, errRuntimeRelay
	}
	return data, nil
}

func publishRuntimeSlot(ctx context.Context, root *os.Root, name, hash string, payload any, slots []runtimeSlotIdentity, lease *os.File) (outcome error) {
	stage := strings.TrimSuffix(name, ".json")
	data, err := json.Marshal(payload)
	if err != nil {
		return errRuntimeRelay
	}
	data, err = json.Marshal(runtimeEnvelope{stage, hash, data})
	if err != nil || len(data) > runtimeLimit {
		return errRuntimeRelay
	}
	file, err := openRuntimeSlot(root, name, os.O_RDWR, slots)
	if err != nil {
		return err
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	if checkErr := runtimeLock(ctx, file, true); checkErr != nil {
		return checkErr
	}
	info, err := file.Stat()
	if err != nil || info.Size() != 0 || ctx.Err() != nil {
		return errRuntimeRelay // write once; never truncate or replace anything
	}
	if lease != nil {
		if checkErr := runtimeLeaseAlive(lease); checkErr != nil {
			return checkErr
		}
	}
	if _, err := file.Write(data); err != nil {
		return errRuntimeRelay
	}
	if checkErr := file.Sync(); checkErr != nil {
		return errRuntimeRelay
	}
	return nil
}

func decodeRuntimeSlot(data []byte, name, hash string, payload any) error {
	var envelope runtimeEnvelope
	if checkErr := decodeRuntime(data, &envelope, runtimeLimit); checkErr != nil || envelope.Stage != strings.TrimSuffix(name, ".json") || envelope.ReceiptHash != hash {
		return errRuntimeRelay
	}
	if checkErr := decodeRuntime(envelope.Payload, payload, runtimeLimit); checkErr != nil {
		return errRuntimeRelay
	}
	return nil
}

func runtimeLeaseAlive(lease *os.File) error {
	var buffer [1]byte
	_, err := unix.Read(int(lease.Fd()), buffer[:])
	if errors.Is(err, unix.EAGAIN) || errors.Is(err, unix.EWOULDBLOCK) {
		return nil // EOF, data and every other error mean owner loss/invalidity.
	}
	return errRuntimeRelay
}

func waitRuntimeSlot(ctx context.Context, root *os.Root, name, hash string, slots []runtimeSlotIdentity, lease *os.File, payload any) error {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()
	for {
		if lease != nil {
			if checkErr := runtimeLeaseAlive(lease); checkErr != nil {
				return checkErr
			}
		}
		data, err := readRuntimeSlot(ctx, root, name, slots)
		if err != nil {
			return err
		}
		if len(data) != 0 {
			return decodeRuntimeSlot(data, name, hash, payload)
		}
		select {
		case <-ctx.Done():
			return errRuntimeRelay
		case <-ticker.C:
		}
	}
}

func runtimeSlotsEmpty(ctx context.Context, root *os.Root, slots []runtimeSlotIdentity, names ...string) error {
	for _, name := range names {
		data, err := readRuntimeSlot(ctx, root, name, slots)
		if err != nil || len(data) != 0 {
			return errRuntimeRelay
		}
	}
	return nil
}

func createRuntimeSlots(directory string) (slots []runtimeSlotIdentity, writer *os.File, outcome error) {
	root, err := openRuntimeDirectory(directory)
	if err != nil {
		return nil, nil, err
	}
	defer func() { outcome = errors.Join(outcome, root.Close()) }()
	for _, name := range []string{runtimeRequestFile, runtimeMenuFile, runtimePreparedFile, runtimeBoundFile} {
		identity, slotErr := createRuntimeSlot(root, name)
		if slotErr != nil {
			return nil, nil, slotErr
		}
		slots = append(slots, identity)
	}
	if checkErr := unix.Mkfifo(filepath.Join(directory, runtimeLeaseFile), 0600); checkErr != nil {
		return nil, nil, errRuntimeRelay
	}
	reader, err := root.OpenFile(runtimeLeaseFile, os.O_RDONLY|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if err != nil {
		return nil, nil, errRuntimeRelay
	}
	identity, statErr := runtimeSlotStat(reader, runtimeLeaseFile, true)
	if statErr != nil {
		return nil, nil, errors.Join(statErr, reader.Close())
	}
	slots = append(slots, identity)
	writer, err = openRuntimeSlot(root, runtimeLeaseFile, os.O_WRONLY, slots)
	closeErr := reader.Close()
	if err != nil || closeErr != nil {
		if writer != nil {
			return nil, nil, errors.Join(errRuntimeRelay, writer.Close())
		}
		return nil, nil, errRuntimeRelay
	}
	return slots, writer, nil
}

func createRuntimeSlot(root *os.Root, name string) (identity runtimeSlotIdentity, outcome error) {
	file, err := root.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL|unix.O_NOFOLLOW, 0600)
	if err != nil {
		return identity, errRuntimeRelay
	}
	defer func() { outcome = errors.Join(outcome, file.Close()) }()
	return runtimeSlotStat(file, name, false)
}

// The installed relay only performs these finite transitions. Lua must perform
// preflight before reservation, menu after one evaluation, and bound only after
// binding the exact prepared receipt/capability in the original registry.
type runtimeRelayRequest struct {
	Operation      string          `json:"operation"`
	Descriptor     string          `json:"descriptor"`
	Anchor         string          `json:"anchor"`
	LayoutRevision int             `json:"layout_revision"`
	Consumer       runtimeConsumer `json:"consumer"`
	Origin         runtimeOrigin   `json:"origin"`
	Server         *string         `json:"server,omitempty"`
	Relay          string          `json:"relay,omitempty"`
	Configuration  string          `json:"configuration,omitempty"`
	ReceiptHash    string          `json:"receipt_hash,omitempty"`
}

// The opaque preflight identity freezes both the installed relay and the exact
// constructor intent, including absence. Lua returns these bytes unchanged;
// only Go canonicalizes the intent to compare with the public receipt endpoint.
type runtimeRelayIdentity struct {
	Installation runtimeInstallation `json:"installation"`
	Server       *string             `json:"server,omitempty"`
}

func checkRuntimeEndpointIntent(endpoint string, intent *string) error {
	if intent == nil {
		return nil
	}
	canonical, err := rex.CanonicalEndpoint(*intent)
	if err != nil || canonical != endpoint {
		return errRuntimeEndpoint
	}
	return nil
}

func readRuntimeDescriptor(ctx context.Context, root *os.Root, locator runtimeLocator) (runtimeDescriptor, runtimeReceipt, error) {
	var descriptor runtimeDescriptor
	var receipt runtimeReceipt
	// Bootstrap the request's identity by checking the owned one-link slot. The
	// request digest then binds the entire fixed identity manifest.
	file, err := root.OpenFile(runtimeRequestFile, os.O_RDONLY|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if err != nil {
		return descriptor, receipt, errRuntimeRelay
	}
	identity, err := runtimeSlotStat(file, runtimeRequestFile, false)
	closeErr := file.Close()
	if err != nil || closeErr != nil {
		return descriptor, receipt, errRuntimeRelay
	}
	data, err := readRuntimeSlot(ctx, root, runtimeRequestFile, []runtimeSlotIdentity{identity})
	if err != nil || runtimeHash(data) != locator.RequestSHA256 || decodeRuntimeSlot(data, runtimeRequestFile, locator.ReceiptHash, &descriptor) != nil {
		return descriptor, receipt, errRuntimeRelay
	}
	if !descriptor.matches(locator, identity) {
		return descriptor, receipt, errRuntimeRelay
	}
	receipt, err = decodeReceipt(descriptor.Receipt, descriptor.ReceiptHash)
	if err != nil || !receipt.initial() {
		return descriptor, receipt, errRuntimeRelay
	}
	return descriptor, receipt, receipt.check(ctx, false)
}

func (descriptor runtimeDescriptor) matches(locator runtimeLocator, identity runtimeSlotIdentity) bool {
	return descriptor.Version == runtimeProtocol && descriptor.Workspace == locator.Workspace && descriptor.ReceiptHash == locator.ReceiptHash &&
		ValidateInvocationCapability(descriptor.Marker) == nil && validRuntimeSlots(descriptor.Slots) && slices.Contains(descriptor.Slots, identity)
}

func (receipt runtimeReceipt) initial() bool {
	return receipt.Workspace == "" && receipt.ConfigurationSHA256 == "" && receipt.Relay == nil && receipt.Origin.Caller == ""
}

func validRuntimeSlots(slots []runtimeSlotIdentity) bool {
	if len(slots) != 5 {
		return false
	}
	seen := make(map[string]bool, 5)
	for _, slot := range slots {
		if seen[slot.Name] || slot.Inode == 0 {
			return false
		}
		switch slot.Name {
		case runtimeRequestFile, runtimeMenuFile, runtimePreparedFile, runtimeBoundFile, runtimeLeaseFile:
		default:
			return false
		}
		seen[slot.Name] = true
	}
	return true
}

func sameRuntimeOrigin(expected, actual runtimeOrigin) bool {
	if actual.validate(true) != nil || expected.Caller != "" && expected.Caller != actual.Caller {
		return false
	}
	expected.Caller = actual.Caller
	return expected == actual
}

func sameRuntimeConsumer(expected, actual runtimeConsumer) bool {
	return reflect.DeepEqual(expected, actual)
}

func validateRelayMenu(ctx context.Context, base runtimeReceipt, menu runtimeMenu) error {
	if !sameRuntimeConsumer(base.Consumer, menu.Consumer) || !sameRuntimeOrigin(base.Origin, menu.Origin) {
		return errRuntimeOrigin
	}
	if checkErr := checkRuntimeEndpointIntent(base.Endpoint, menu.Server); checkErr != nil {
		return checkErr
	}
	if checkErr := validateRuntimeConfig(menu.Configuration, menu.Consumer); checkErr != nil {
		return checkErr
	}
	return menu.Relay.check(ctx)
}

func preparedRuntimeMatches(ctx context.Context, base runtimeReceipt, menu runtimeMenu, prepared popupPreparation, workspace string) error {
	if prepared.Version != runtimeProtocol || ValidateInvocationCapability(prepared.Capability) != nil || prepared.Workspace != workspace || prepared.Executable != filepath.Join(workspace, runtimeImageName) || prepared.Server != base.Endpoint || prepared.Expires != base.Expires {
		return errRuntimeRelay
	}
	receipt, err := decodeReceipt(prepared.Receipt, prepared.ReceiptHash)
	if err != nil {
		return err
	}
	expected := base
	expected.Origin, expected.Relay = menu.Origin, &menu.Relay
	expected.Workspace, expected.ConfigurationSHA256 = workspace, runtimeHash([]byte(menu.Configuration))
	data, hash, err := encodeReceipt(expected)
	if err != nil || data != prepared.Receipt || hash != prepared.ReceiptHash {
		return errRuntimeRelay
	}
	if checkErr := receipt.check(ctx, true); checkErr != nil {
		return checkErr
	}
	return verifyRuntimeCopy(ctx, workspace, receipt.Component)
}

// RuntimeRelay never executes a nominated program, creates a caller-named file,
// evaluates a mapping, or returns a secret to the public action's caller.
func RuntimeRelay(parent context.Context, input io.Reader, output io.Writer) (outcome error) {
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()
	request, locator, err := decodeRuntimeRelay(input)
	if err != nil {
		return err
	}
	root, err := openRuntimeDirectory(locator.Workspace)
	if err != nil {
		return err
	}
	defer func() { outcome = errors.Join(outcome, root.Close()) }()
	descriptor, receipt, err := readRuntimeDescriptor(ctx, root, locator)
	if err != nil {
		return err
	}
	stage, cancelStage := context.WithDeadline(ctx, time.Unix(receipt.Expires, 0).Add(-5*time.Second))
	defer cancelStage()
	lease, err := openRuntimeSlot(root, runtimeLeaseFile, os.O_RDONLY, descriptor.Slots)
	if err != nil {
		return err
	}
	defer func() { outcome = errors.Join(outcome, lease.Close()) }()
	if checkErr := runtimeLeaseAlive(lease); checkErr != nil {
		return checkErr
	}
	// The action's GUI nomination is untrusted. Compare it with the private
	// receipt before checkOrigin revalidates that exact live app/generation;
	// neither preflight nor later transitions may choose a replacement target.
	if !sameRuntimeConsumer(receipt.Consumer, request.Consumer) || !sameRuntimeOrigin(receipt.Origin, request.Origin) {
		return errRuntimeOrigin
	}
	if checkErr := checkRuntimeEndpointIntent(receipt.Endpoint, request.Server); checkErr != nil {
		return checkErr
	}
	image, err := captureSelf(stage)
	if err != nil {
		return err
	}
	defer func() { outcome = errors.Join(outcome, image.file.Close()) }()
	installation, err := captureInstallation(stage, request.Anchor, request.LayoutRevision, image.identity)
	if err != nil {
		return err
	}
	if checkErr := checkRuntimeRelayInstallation(request, installation); checkErr != nil {
		return checkErr
	}
	receipt.Origin = request.Origin
	if checkErr := receipt.checkOrigin(stage); checkErr != nil {
		return checkErr
	}
	if checkErr := runtimeLeaseAlive(lease); checkErr != nil {
		return checkErr
	}
	return performRuntimeRelay(stage, root, lease, descriptor, receipt, request, *installation, output)
}

func checkRuntimeRelayInstallation(request runtimeRelayRequest, installation *runtimeInstallation) error {
	if request.Operation == runtimePreflight {
		return nil
	}
	var retained runtimeRelayIdentity
	expected := runtimeRelayIdentity{*installation, request.Server}
	encoded, err := json.Marshal(expected)
	if err != nil || decodeRuntime([]byte(request.Relay), &retained, popupMessageLimit) != nil ||
		request.Relay != string(encoded) || !reflect.DeepEqual(retained, expected) {
		return errRuntimeCompatibility
	}
	return nil
}

func decodeRuntimeRelay(input io.Reader) (runtimeRelayRequest, runtimeLocator, error) {
	var request runtimeRelayRequest
	var locator runtimeLocator
	if checkErr := readRuntime(input, &request, runtimeLimit); checkErr != nil {
		return request, locator, checkErr
	}
	if checkErr := decodeRuntime([]byte(request.Descriptor), &locator, popupMessageLimit); checkErr != nil {
		return request, locator, errRuntimeRelay
	}
	if locator.Version != runtimeProtocol || ValidateInvocationCapability(locator.RequestSHA256) != nil || ValidateInvocationCapability(locator.ReceiptHash) != nil {
		return request, locator, errRuntimeRelay
	}
	return request, locator, nil
}

func performRuntimeRelay(stage context.Context, root *os.Root, lease *os.File, descriptor runtimeDescriptor, receipt runtimeReceipt, request runtimeRelayRequest, installation runtimeInstallation, output io.Writer) error {
	switch request.Operation {
	case runtimePreflight:
		if request.Configuration != "" || request.ReceiptHash != "" || request.Relay != "" || runtimeSlotsEmpty(stage, root, descriptor.Slots, runtimeMenuFile, runtimePreparedFile, runtimeBoundFile) != nil {
			return errRuntimeRelay
		}
		relay, err := json.Marshal(runtimeRelayIdentity{installation, request.Server})
		if err != nil || len(relay) > popupMessageLimit {
			return errRuntimeCompatibility
		}
		return writeRuntime(output, struct {
			Version        int      `json:"version"`
			ProtocolMin    int      `json:"protocol_min"`
			ProtocolMax    int      `json:"protocol_max"`
			Build          string   `json:"producer_build"`
			Features       []string `json:"features"`
			LayoutRevision int      `json:"layout_revision"`
			ReceiptHash    string   `json:"receipt_hash"`
			Relay          string   `json:"relay"`
		}{runtimeProtocol, runtimeProtocol, runtimeProtocol, installation.Helper.SHA256, runtimeFeatures(), runtimeLayout, descriptor.ReceiptHash, string(relay)}, popupMessageLimit)
	case "menu":
		return relayRuntimeMenu(stage, root, lease, descriptor, receipt, request, installation, output)
	case "bound":
		return relayRuntimeBound(stage, root, lease, descriptor, receipt, request, installation, output)
	default:
		return errRuntimeRelay
	}
}

func relayRuntimeMenu(ctx context.Context, root *os.Root, lease *os.File, descriptor runtimeDescriptor, receipt runtimeReceipt, request runtimeRelayRequest, installation runtimeInstallation, output io.Writer) error {
	if request.ReceiptHash != "" || runtimeSlotsEmpty(ctx, root, descriptor.Slots, runtimeMenuFile, runtimePreparedFile, runtimeBoundFile) != nil {
		return errRuntimeRelay
	}
	menu := runtimeMenu{request.Configuration, request.Consumer, request.Origin, installation, request.Server}
	if checkErr := validateRelayMenu(ctx, receipt, menu); checkErr != nil {
		return checkErr
	}
	if checkErr := publishRuntimeSlot(ctx, root, runtimeMenuFile, descriptor.ReceiptHash, menu, descriptor.Slots, lease); checkErr != nil {
		return checkErr
	}
	var prepared popupPreparation
	if checkErr := waitRuntimeSlot(ctx, root, runtimePreparedFile, descriptor.ReceiptHash, descriptor.Slots, lease, &prepared); checkErr != nil {
		return checkErr
	}
	base, err := decodeReceipt(descriptor.Receipt, descriptor.ReceiptHash)
	if err != nil || preparedRuntimeMatches(ctx, base, menu, prepared, descriptor.Workspace) != nil {
		return errRuntimeRelay
	}
	if checkErr := receipt.checkOrigin(ctx); checkErr != nil {
		return checkErr
	}
	if checkErr := runtimeLeaseAlive(lease); checkErr != nil {
		return checkErr
	}
	return writeRuntime(output, prepared, runtimeLimit)
}

func relayRuntimeBound(ctx context.Context, root *os.Root, lease *os.File, descriptor runtimeDescriptor, receipt runtimeReceipt, request runtimeRelayRequest, installation runtimeInstallation, output io.Writer) error {
	if request.Configuration != "" || ValidateInvocationCapability(request.ReceiptHash) != nil || runtimeSlotsEmpty(ctx, root, descriptor.Slots, runtimeBoundFile) != nil {
		return errRuntimeRelay
	}
	var menu runtimeMenu
	var prepared popupPreparation
	if checkErr := waitRuntimeSlot(ctx, root, runtimeMenuFile, descriptor.ReceiptHash, descriptor.Slots, lease, &menu); checkErr != nil {
		return checkErr
	}
	retained := runtimeRelayIdentity{menu.Relay, menu.Server}
	expected := runtimeRelayIdentity{installation, request.Server}
	if checkErr := validateRelayMenu(ctx, receipt, menu); checkErr != nil || !reflect.DeepEqual(retained, expected) || menu.Origin != request.Origin {
		return errRuntimeRelay
	}
	if checkErr := waitRuntimeSlot(ctx, root, runtimePreparedFile, descriptor.ReceiptHash, descriptor.Slots, lease, &prepared); checkErr != nil {
		return checkErr
	}
	base, err := decodeReceipt(descriptor.Receipt, descriptor.ReceiptHash)
	if err != nil || preparedRuntimeMatches(ctx, base, menu, prepared, descriptor.Workspace) != nil || prepared.ReceiptHash != request.ReceiptHash {
		return errRuntimeRelay
	}
	if checkErr := receipt.checkOrigin(ctx); checkErr != nil {
		return checkErr
	}
	if checkErr := runtimeLeaseAlive(lease); checkErr != nil {
		return checkErr
	}
	if checkErr := publishRuntimeSlot(ctx, root, runtimeBoundFile, descriptor.ReceiptHash, runtimeBinding{prepared.ReceiptHash}, descriptor.Slots, lease); checkErr != nil {
		return checkErr
	}
	return writeRuntime(output, struct {
		Version int `json:"version"`
	}{runtimeProtocol}, popupMessageLimit)
}

// runtimePublicOwner holds only preparation/relay handles. The detached
// controller owns its verified copy and finite stdin, not this FIFO lease.
// Clearing before close makes even a failed close a non-retryable transition.
type runtimePublicOwner struct {
	root  *os.Root
	lease *os.File
}

func (owner *runtimePublicOwner) close() error {
	var closeErr error
	if owner.root != nil {
		root := owner.root
		owner.root = nil
		closeErr = errors.Join(closeErr, root.Close())
	}
	if owner.lease != nil {
		lease := owner.lease
		owner.lease = nil
		closeErr = errors.Join(closeErr, lease.Close())
	}
	if closeErr != nil {
		return errRuntimeRelay
	}
	return nil
}

// retire closes the public caller's captured source once its verified private
// copy is ready. LaunchPopup's fallback defer must not close it again afterward.
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

func (owner *runtimePublicOwner) startController(ctx context.Context, image *runtimeImage, prepared popupPreparation) (bool, error) {
	// Binding and activation must be certain before retiring these handles.
	// A failed close suppresses startup; no uncertain close is repeated.
	if closeErr := errors.Join(owner.close(), image.retire()); closeErr != nil {
		return false, closeErr
	}
	return startRuntimeController(ctx, prepared)
}

func launchPublicRuntime(parent context.Context, origin launchOrigin, placement rex.Block, consumer runtimeConsumer) (outcome error) {
	started := time.Now()
	expires := started.Unix() + InvocationLifetimeSeconds
	ctx, stop := signalRuntimeContext(parent, time.Unix(expires, 0).Add(-5*time.Second))
	defer stop()
	base := runtimeReceipt{Protocol: runtimeProtocol, Features: negotiatedRuntimeFeatures(consumer), AssetSHA256: runtimeHash([]byte(runtimeAssets)), Component: origin.image.identity, CLI: origin.selection, Endpoint: origin.Server, Consumer: consumer,
		Origin: runtimeOrigin{Session: origin.SessionID, Block: origin.BlockID, Window: placement.WindowID, Layer: placement.LayerID, Client: origin.ClientID, ConnectedAt: origin.ConnectedAt}, Expires: expires, OS: runtime.GOOS, Arch: runtime.GOARCH}
	if checkErr := base.check(ctx, false); checkErr != nil {
		return checkErr
	}
	workspace, err := newRuntimeDirectory("rex-whichkey-runtime-")
	if err != nil {
		return errRuntimeRelay
	}
	transferred := false
	var pending *popupRequest
	defer func() { outcome = errors.Join(outcome, cleanupPublicRuntime(ctx, workspace, transferred, pending)) }()
	slots, writer, err := createRuntimeSlots(workspace)
	// Slot creation can return a live writer with a directory-close error.
	// Adopt it before inspecting that error so every acquired lease is retired.
	owner := runtimePublicOwner{lease: writer}
	defer func() { outcome = errors.Join(outcome, owner.close()) }()
	if err != nil {
		return err
	}
	root, err := openRuntimeDirectory(workspace)
	if err != nil {
		return err
	}
	owner.root = root
	descriptor, arguments, err := publishRuntimeRequest(ctx, root, base, workspace, slots)
	if err != nil {
		return err
	}

	// Exactly one action runs while this process services private files. Join
	// it on every exit; uncertain activation never starts a controller.
	activation, finishActivation := context.WithCancel(ctx)
	defer finishActivation()
	activated := make(chan error, 1)
	go func() {
		activated <- runtimeAction(activation, origin.selection, origin, consumer.Launch, arguments, &launchOutput{remaining: popupMessageLimit})
	}()
	joined := false
	defer func() {
		if !joined {
			finishActivation()
			outcome = errors.Join(outcome, <-activated)
		}
	}()
	prepared, err := preparePublicRendezvous(ctx, root, descriptor, base, origin.image)
	if err != nil {
		return err
	}
	cancellation, err := runtimeControllerRequest(ctx, prepared)
	if err != nil {
		return err
	}
	pending = &cancellation
	select {
	case activationErr := <-activated:
		joined = true
		if activationErr != nil {
			return errRuntimeRelay
		}
	case <-ctx.Done():
		return errRuntimeRelay
	}
	final, err := decodeReceipt(prepared.Receipt, prepared.ReceiptHash)
	if err != nil {
		return errors.Join(errRuntimeOrigin, err)
	}
	if checkErr := final.checkOrigin(ctx); checkErr != nil {
		return errors.Join(errRuntimeOrigin, checkErr)
	}
	if !time.Unix(expires, 0).Add(-5 * time.Second).After(time.Now()) {
		return errRuntimeOrigin
	}
	// Both relay stages and activation are finished. The controller never
	// reads the public lease; caller return cannot revoke its finite handoff.
	transferred, err = owner.startController(ctx, origin.image, prepared)
	return err
}

func signalRuntimeContext(parent context.Context, deadline time.Time) (context.Context, context.CancelFunc) {
	ctx, stop := signal.NotifyContext(parent, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	bounded, cancel := context.WithDeadline(ctx, deadline)
	return bounded, func() { cancel(); stop() }
}

func preparePublicRendezvous(ctx context.Context, root *os.Root, descriptor runtimeDescriptor, base runtimeReceipt, image *runtimeImage) (popupPreparation, error) {
	stage, cancel := context.WithTimeout(ctx, 3*time.Second)
	var menu runtimeMenu
	err := waitRuntimeSlot(stage, root, runtimeMenuFile, descriptor.ReceiptHash, descriptor.Slots, nil, &menu)
	cancel()
	if err != nil {
		return popupPreparation{}, err
	}
	if checkErr := validateRelayMenu(ctx, base, menu); checkErr != nil {
		return popupPreparation{}, checkErr
	}
	if checkErr := runtimeSlotsEmpty(ctx, root, descriptor.Slots, runtimePreparedFile, runtimeBoundFile); checkErr != nil {
		return popupPreparation{}, checkErr
	}
	final := base
	final.Origin, final.Relay = menu.Origin, &menu.Relay
	stage, cancel = context.WithTimeout(ctx, 3*time.Second)
	prepared, err := prepareRuntime(stage, image, final, menu.Configuration, descriptor.Workspace)
	cancel()
	if err != nil {
		return popupPreparation{}, err
	}
	if checkErr := publishRuntimeSlot(ctx, root, runtimePreparedFile, descriptor.ReceiptHash, prepared, descriptor.Slots, nil); checkErr != nil {
		return popupPreparation{}, checkErr
	}
	stage, cancel = context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	var binding runtimeBinding
	if checkErr := waitRuntimeSlot(stage, root, runtimeBoundFile, descriptor.ReceiptHash, descriptor.Slots, nil, &binding); checkErr != nil || binding.ReceiptHash != prepared.ReceiptHash {
		return popupPreparation{}, errRuntimeRelay
	}
	return prepared, nil
}

func runtimeControllerRequest(ctx context.Context, prepared popupPreparation) (_ popupRequest, outcome error) {
	receipt, err := decodeReceipt(prepared.Receipt, prepared.ReceiptHash)
	if err != nil {
		return popupRequest{}, err
	}
	root, err := openRuntimeDirectory(prepared.Workspace)
	if err != nil {
		return popupRequest{}, err
	}
	defer func() { outcome = errors.Join(outcome, root.Close()) }()
	// Only the public path starts its own controller. Key launch has already
	// retained configuration in its original VM and supplies it on stdin.
	file, err := root.OpenFile(runtimeMenuFile, os.O_RDONLY|unix.O_NOFOLLOW|unix.O_NONBLOCK, 0)
	if err != nil {
		return popupRequest{}, errRuntimeRelay
	}
	identity, err := runtimeSlotStat(file, runtimeMenuFile, false)
	closeErr := file.Close()
	if err != nil || closeErr != nil {
		return popupRequest{}, errRuntimeRelay
	}
	data, err := readRuntimeSlot(ctx, root, runtimeMenuFile, []runtimeSlotIdentity{identity})
	if err != nil {
		return popupRequest{}, err
	}
	var envelope runtimeEnvelope
	if checkErr := decodeRuntime(data, &envelope, runtimeLimit); checkErr != nil || envelope.Stage != "menu" {
		return popupRequest{}, errRuntimeRelay
	}
	var menu runtimeMenu
	if checkErr := decodeRuntime(envelope.Payload, &menu, runtimeLimit); checkErr != nil || receipt.ConfigurationSHA256 != runtimeHash([]byte(menu.Configuration)) {
		return popupRequest{}, errRuntimeRelay
	}
	return popupRequest{Configuration: menu.Configuration, Capability: prepared.Capability, Server: receipt.Endpoint, SessionID: receipt.Origin.Session, BlockID: receipt.Origin.Block, ClientID: receipt.Origin.Client, ConnectedAt: receipt.Origin.ConnectedAt, Completion: receipt.Consumer.Completion, Expires: receipt.Expires, Receipt: prepared.Receipt, ReceiptHash: prepared.ReceiptHash, Workspace: prepared.Workspace, retained: &receipt}, nil
}

func startRuntimeController(ctx context.Context, prepared popupPreparation) (bool, error) {
	request, err := runtimeControllerRequest(ctx, prepared)
	if err != nil {
		return false, err
	}
	return startRuntimeRequest(ctx, request, nil)
}

// Start success transfers ownership even if the subsequent finite handshake
// fails. Neither caller cancellation nor a lost acknowledgment authorizes retry,
// killing an uncertain child, or removing its workspace.
func startRuntimeRequest(parent context.Context, request popupRequest, beforeProceed func() error) (bool, error) {
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()
	receipt, err := decodeReceipt(request.Receipt, request.ReceiptHash)
	if err != nil || !request.matchesReceipt(receipt) || receipt.check(ctx, true) != nil ||
		verifyRuntimeCopy(ctx, request.Workspace, receipt.Component) != nil || ctx.Err() != nil {
		return false, errRuntimeStartup
	}
	data, err := json.Marshal(request)
	if err != nil || len(data) > runtimeLimit {
		return false, errRuntimeStartup
	}
	reader, writer, err := os.Pipe()
	if err != nil {
		return false, errRuntimeStartup
	}
	ackReader, ackWriter, err := os.Pipe()
	if err != nil {
		return false, errors.Join(errRuntimeStartup, reader.Close(), writer.Close())
	}
	proceedReader, proceedWriter, err := os.Pipe()
	if err != nil {
		return false, errors.Join(errRuntimeStartup, reader.Close(), writer.Close(), ackReader.Close(), ackWriter.Close())
	}
	// The child owns its original expiry, not the finite caller's context.
	command := exec.CommandContext(context.WithoutCancel(ctx), filepath.Join(request.Workspace, runtimeImageName), "controller-child")
	command.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	command.Stdin = reader                                    // no parent-owned exec copy goroutine after return
	command.ExtraFiles = []*os.File{ackWriter, proceedReader} // fd 3/4; no capability in argv/env
	startErr := ctx.Err()
	if startErr == nil {
		startErr = command.Start()
	}
	if startErr != nil {
		return false, errors.Join(errRuntimeStartup, reader.Close(), writer.Close(), ackReader.Close(), ackWriter.Close(), proceedReader.Close(), proceedWriter.Close())
	}
	return true, sendRuntimeHandoff(ctx, reader, writer, ackReader, ackWriter, proceedReader, proceedWriter, command.Process, data, beforeProceed)
}

// exec inherits blocking descriptors. Reopen a duplicated, nonblocking pipe as
// a pollable os.File so deadlines work on both Darwin and Linux. The original
// wrapper is closed/invalidated before the duplicate can be returned.
func runtimePollablePipe(file *os.File) (*os.File, error) {
	if file == nil {
		return nil, errRuntimeStartup
	}
	info, err := file.Stat()
	if err != nil || info.Mode()&os.ModeNamedPipe == 0 {
		return nil, errors.Join(errRuntimeStartup, file.Close())
	}
	fd, dupErr := unix.FcntlInt(file.Fd(), unix.F_DUPFD_CLOEXEC, 0)
	closeErr := file.Close()
	if dupErr != nil {
		return nil, errRuntimeStartup
	}
	if nonblockErr := unix.SetNonblock(fd, true); closeErr != nil || nonblockErr != nil {
		return nil, errors.Join(errRuntimeStartup, unix.Close(fd))
	}
	return os.NewFile(uintptr(fd), "runtime-pipe"), nil
}

type runtimeStartupAck struct {
	Version int    `json:"version"`
	Status  string `json:"status"`
}

func sendRuntimeHandoff(ctx context.Context, reader, writer, ackReader, ackWriter, proceedReader, proceedWriter *os.File, process *os.Process, data []byte, beforeProceed func() error) (outcome error) {
	defer func() {
		if err := errors.Join(proceedWriter.Close(), retireRuntimeStartup(&ackReader, &process), ctx.Err()); err != nil {
			outcome = errRuntimeStartup
		}
	}()
	// Closing our duplicate writer is necessary to observe the child's EOF.
	closeErr := errors.Join(reader.Close(), ackWriter.Close(), proceedReader.Close())
	deadline, _ := ctx.Deadline() // startRuntimeRequest always supplies a bound.
	deadlineErr := errors.Join(writer.SetWriteDeadline(deadline), ackReader.SetReadDeadline(deadline), proceedWriter.SetWriteDeadline(deadline))
	pipes := []*os.File{writer, ackReader, proceedWriter}
	stop := context.AfterFunc(ctx, func() {
		// Wake blocked pipe I/O on cancellation, without killing/retrying a child.
		for _, pipe := range pipes {
			if err := pipe.SetDeadline(time.Now()); err != nil {
				continue // A retired endpoint already wakes its blocked I/O.
			}
		}
	})
	defer stop()
	var writeErr error
	if errors.Join(closeErr, deadlineErr, ctx.Err()) == nil {
		_, writeErr = writer.Write(data)
	}
	writerErr := writer.Close() // EOF is part of the finite handoff.
	if errors.Join(closeErr, deadlineErr, writeErr, writerErr, ctx.Err()) != nil {
		return errRuntimeStartup
	}
	var ack runtimeStartupAck
	if readRuntime(ackReader, &ack, popupMessageLimit) != nil || ack.Version != runtimeProtocol || ack.Status != runtimeAccepted || ctx.Err() != nil {
		return errRuntimeStartup
	}
	if checkErr := finishRuntimeStartup(ctx, &ackReader, &process, beforeProceed); checkErr != nil {
		return checkErr
	}
	// Without this positive marker plus EOF the child must not enter its
	// lifecycle, even if it already sent an accepted acknowledgment.
	return writeRuntime(proceedWriter, runtimeStartupAck{runtimeProtocol, "proceed"}, popupMessageLimit)
}

func retireRuntimeStartup(reader **os.File, process **os.Process) error {
	var closeErr error
	if *reader != nil {
		file := *reader
		*reader = nil
		closeErr = file.Close()
	}
	if *process != nil {
		child := *process
		*process = nil
		closeErr = errors.Join(closeErr, child.Release())
	}
	return closeErr
}

func finishRuntimeStartup(ctx context.Context, reader **os.File, process **os.Process, beforeProceed func() error) error {
	if err := errors.Join(retireRuntimeStartup(reader, process), ctx.Err()); err != nil {
		return errRuntimeStartup
	}
	if beforeProceed != nil {
		if err := beforeProceed(); err != nil {
			return errRuntimeStartup
		}
	}
	return ctx.Err()
}

func popupControllerPipes(input, startup, proceed *os.File) (*os.File, *os.File, *os.File, error) {
	reader, err := runtimePollablePipe(input)
	if err != nil {
		return nil, nil, nil, errors.Join(err, closeRuntimePipe(&startup), closeRuntimePipe(&proceed))
	}
	writer, err := runtimePollablePipe(startup)
	if err != nil {
		return nil, nil, nil, errors.Join(err, reader.Close(), closeRuntimePipe(&proceed))
	}
	gate, err := runtimePollablePipe(proceed)
	if err != nil {
		return nil, nil, nil, errors.Join(err, reader.Close(), writer.Close())
	}
	return reader, writer, gate, nil
}

func closeRuntimePipe(pipe **os.File) error {
	if *pipe == nil {
		return nil
	}
	file := *pipe
	*pipe = nil
	return file.Close()
}

func reportRuntimeLauncher(output io.Writer) error {
	writeErr := writeRuntime(output, runtimeStartupAck{runtimeProtocol, runtimeAccepted}, popupMessageLimit)
	if closer, ok := output.(io.Closer); ok {
		return errors.Join(writeErr, closer.Close())
	}
	return writeErr
}

func rejectRuntimeStartup(ctx context.Context, output **os.File) error {
	if *output == nil {
		return nil
	}
	file := *output
	*output = nil
	return reportRuntimeStartup(ctx, file, false)
}

func waitRuntimeProceed(ctx context.Context, pipe *os.File) error {
	deadline, _ := ctx.Deadline()
	if err := pipe.SetReadDeadline(deadline); err != nil {
		return errors.Join(errRuntimeStartup, pipe.Close())
	}
	var message runtimeStartupAck
	readErr := readRuntime(pipe, &message, popupMessageLimit)
	closeErr := pipe.Close()
	if errors.Join(readErr, closeErr, ctx.Err()) != nil || message.Version != runtimeProtocol || message.Status != "proceed" {
		return errRuntimeStartup
	}
	return nil
}

func acknowledgePopupStartup(ctx context.Context, request popupRequest, output *os.File) (outcome error) {
	accepted := false
	defer func() { outcome = errors.Join(outcome, reportRuntimeStartup(ctx, output, accepted)) }()
	if checkErr := request.retained.checkOrigin(ctx); checkErr != nil {
		return checkErr
	}
	if checkErr := request.retained.checkLoadedConsumer(ctx); checkErr != nil {
		return checkErr
	}
	if ctx.Err() != nil || !time.Unix(request.Expires, 0).Add(-5*time.Second).After(time.Now()) {
		return errRuntimeStartup
	}
	accepted = true
	return nil
}

// reportRuntimeStartup closes the child's acknowledgment endpoint before any
// popup effect. Success means finite input and immediate startup checks passed,
// not that a UI was displayed or a later retained action ran.
func reportRuntimeStartup(ctx context.Context, output *os.File, accepted bool) error {
	status := "rejected"
	if accepted {
		status = runtimeAccepted
	}
	deadline := time.Now().Add(3 * time.Second)
	if bound, ok := ctx.Deadline(); ok && bound.Before(deadline) {
		deadline = bound
	}
	deadlineErr := output.SetWriteDeadline(deadline)
	var writeErr error
	if deadlineErr == nil {
		writeErr = writeRuntime(output, runtimeStartupAck{runtimeProtocol, status}, popupMessageLimit)
	}
	if err := errors.Join(deadlineErr, writeErr, output.Close()); err != nil {
		return errRuntimeStartup
	}
	return nil
}

func publishRuntimeRequest(ctx context.Context, root *os.Root, base runtimeReceipt, workspace string, slots []runtimeSlotIdentity) (runtimeDescriptor, []byte, error) {
	data, hash, err := encodeReceipt(base)
	if err != nil {
		return runtimeDescriptor{}, nil, err
	}
	marker, err := NewInvocationCapability()
	if err != nil {
		return runtimeDescriptor{}, nil, errRuntimeRelay
	}
	descriptor := runtimeDescriptor{runtimeProtocol, marker, workspace, data, hash, slots}
	if checkErr := publishRuntimeSlot(ctx, root, runtimeRequestFile, hash, descriptor, slots, nil); checkErr != nil {
		return descriptor, nil, checkErr
	}
	requestData, err := readRuntimeSlot(ctx, root, runtimeRequestFile, slots)
	if err != nil {
		return descriptor, nil, err
	}
	locator, err := json.Marshal(runtimeLocator{runtimeProtocol, workspace, runtimeHash(requestData), hash})
	if err != nil {
		return descriptor, nil, errRuntimeRelay
	}
	// Lua needs the retained GUI nomination to enter authenticated preflight.
	// These are routing hints, not authority or the fresh action's CLI caller;
	// RuntimeRelay compares them with this receipt before returning any reply.
	arguments, err := json.Marshal(struct {
		Runtime     string `json:"runtime"`
		ClientID    string `json:"client_id"`
		ConnectedAt string `json:"connected_at"`
	}{string(locator), base.Origin.Client, base.Origin.ConnectedAt})
	if err != nil || len(arguments) > popupMessageLimit {
		return descriptor, nil, errRuntimeRelay
	}
	return descriptor, arguments, nil
}

func (r popupRequest) preparePopupImage(ctx context.Context, directory string) error {
	request := r
	if request.image == nil {
		return errRuntimeStartup
	}
	if request.retained != nil {
		if checkErr := request.retained.check(ctx, true); checkErr != nil {
			return checkErr
		}
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

func (receipt runtimeReceipt) checkPlacement(block rex.Block) error {
	if block.WindowID != receipt.Origin.Window || block.LayerID != receipt.Origin.Layer {
		return errRuntimeOrigin
	}
	return nil
}

func (receipt runtimeReceipt) checkLoadedConsumer(ctx context.Context) error {
	origin := launchOrigin{Server: receipt.Endpoint, SessionID: receipt.Origin.Session, BlockID: receipt.Origin.Block, ClientID: receipt.Origin.Client}
	consumer, err := queryRuntimeConsumer(ctx, receipt.CLI, origin, receipt.Consumer.Launch)
	if err != nil {
		return errors.Join(errRuntimeCompatibility, err)
	}
	if !sameRuntimeConsumer(consumer, receipt.Consumer) {
		return errRuntimeCompatibility
	}
	return nil
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
	if _, statErr := runtimeSlotStat(file, "result.json", false); statErr != nil {
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

func cleanupPublicRuntime(ctx context.Context, workspace string, transferred bool, pending *popupRequest) error {
	if transferred {
		return nil
	}
	if err := removeRuntimeWorkspace(workspace); err != nil {
		return err
	}
	if pending == nil {
		return nil
	}
	return cancelPreparedRuntime(ctx, *pending)
}

func cancelPreparedRuntime(parent context.Context, request popupRequest) error {
	if request.retained == nil || !time.Now().Before(time.Unix(request.Expires, 0)) {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.WithoutCancel(parent), 3*time.Second)
	defer cancel()
	return request.complete(ctx, request.retained.CLI, Completion{Version: CompletionVersion, Capability: request.Capability, Kind: CompletionCancel})
}

// RuntimeRelease removes only the fixed owned key-preparation workspace. Lua
// calls it only before attempting controller start; uncertain starts are never
// retried or treated as permission to revoke a possibly running controller.
func RuntimeRelease(ctx context.Context, input io.Reader) (outcome error) {
	var prepared popupPreparation
	if checkErr := readRuntime(input, &prepared, runtimeLimit); checkErr != nil {
		return checkErr
	}
	receipt, err := decodeReceipt(prepared.Receipt, prepared.ReceiptHash)
	if err != nil || prepared.Version != runtimeProtocol || ValidateInvocationCapability(prepared.Capability) != nil || receipt.Workspace != prepared.Workspace || receipt.Relay != nil || receipt.Layout == nil {
		return errRuntimePreparation
	}
	image, err := captureSelf(ctx)
	if err != nil {
		return err
	}
	defer func() { outcome = errors.Join(outcome, image.file.Close()) }()
	if image.identity != receipt.Component || verifyRuntimeCopy(ctx, receipt.Workspace, receipt.Component) != nil {
		return errRuntimePreparation
	}
	return removeRuntimeWorkspace(receipt.Workspace)
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
	if _, statErr := runtimeSlotStat(file, "runtime.json", false); statErr != nil {
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

// validateRuntime authenticates structure and frozen identities before any popup.
// Private stdin is the handoff boundary, not a public executable selector.
func (r popupRequest) matchesReceipt(receipt runtimeReceipt) bool {
	return receipt.Workspace == r.Workspace && receipt.ConfigurationSHA256 == runtimeHash([]byte(r.Configuration)) &&
		receipt.Endpoint == r.Server && receipt.Expires == r.Expires && receipt.Consumer.Completion == r.Completion &&
		receipt.Origin.Session == r.SessionID && receipt.Origin.Block == r.BlockID && receipt.Origin.Client == r.ClientID && receipt.Origin.ConnectedAt == r.ConnectedAt
}

func (r *popupRequest) validateRuntime(ctx context.Context) (outcome error) {
	request := r
	receipt, err := decodeReceipt(request.Receipt, request.ReceiptHash)
	if err != nil {
		return err
	}
	if !request.matchesReceipt(receipt) {
		return errRuntimeCompatibility
	}
	if checkErr := verifyRuntimeCopy(ctx, request.Workspace, receipt.Component); checkErr != nil {
		return checkErr
	}
	image, err := captureSelf(ctx)
	if err != nil {
		return err
	}
	if image.identity.Path != filepath.Join(request.Workspace, runtimeImageName) || image.identity.SHA256 != receipt.Component.SHA256 || image.identity.Size != receipt.Component.Size {
		return errors.Join(errRuntimeStartup, image.file.Close())
	}
	request.retained, request.image = &receipt, image
	if checkErr := receipt.check(ctx, true); checkErr != nil {
		return checkErr
	}
	if checkErr := receipt.checkLoadedConsumer(ctx); checkErr != nil {
		return checkErr
	}
	return receipt.checkOrigin(ctx)
}
