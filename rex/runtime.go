package rex

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"unicode"
	"unicode/utf8"
)

const (
	maxExecutableBytes = 512 << 20
	fieldLayout        = "layout"
	fieldWindow        = "window_id"
	schemaInteger      = "integer"
	schemaMinItems     = "minItems"
	schemaMaxItems     = "maxItems"
	schemaMinimum      = "minimum"
	schemaMaximum      = "maximum"
	schemaConst        = "const"
	schemaEnum         = "enum"
)

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
	Executable   FileIdentity     `json:"executable"`
	Provenance   string           `json:"provenance"`
	OS           string           `json:"os"`
	Arch         string           `json:"arch"`
	VersionJSON  string           `json:"version_json"`
	Layout       []FileIdentity   `json:"layout,omitempty"`
	Capabilities *APICapabilities `json:"capabilities,omitempty"`
}

// CLIVersion records the documented server.version result. Missing API revision
// fields mean zero. Release strings are audit data, not equality requirements.
type CLIVersion struct {
	Version        string `json:"rex_version"`
	APIRevision    uint64 `json:"api_revision"`
	APIMinRevision uint64 `json:"api_min_revision"`
}

var (
	errCLIInput         = errors.New("rex: invalid explicit CLI source; inspect the execution-host executable source before a new invocation; not retried")
	errCLIAbsent        = errors.New("rex: no execution-host Rex CLI source; install the integration with its licensed dependency or use a Rex terminal; not retried")
	errCLIPath          = errors.New("rex: PATH must contain only nonempty absolute execution-host directories and one unambiguous Rex executable; not retried")
	errCLILayout        = errors.New("rex: invalid application runtime dependency layout; inspect the execution-host dependency assets before a new invocation; not retried")
	errCLIChanged       = errors.New("rex: frozen CLI, runtime asset or server/config identity is invalid or changed; inspect any possibly completed operation before a new invocation; not retried")
	errCLICompatibility = errors.New("rex: CLI/API contract not established; activation refused, not retried; inspect the named metadata contract and any possibly completed operation before a new invocation")
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
		selection := options.Selection.clone()
		if err := selection.validate(ctx); err != nil {
			return nil, err
		}
		return selection, nil
	}
	selection, err := selectCLI(ctx, options)
	if err != nil {
		return nil, err
	}
	output, err := runSelected(ctx, selection, []string{flagNoAutostart, "version", flagJSON}, nil, 64<<10)
	if err != nil {
		return nil, errors.Join(compatibilityFailure("CLI version", "command failed"), err, ctx.Err())
	}
	// The version command's JSON fields are not specified by the CLI docs.
	// Retain its bounded nonempty object verbatim, not an invented CLI revision
	// range. Successful server metadata calls establish transport compatibility.
	version, err := metadataObject(output)
	if err != nil || len(version) == 0 {
		return nil, compatibilityFailure("CLI version", "nonempty object")
	}
	selection.VersionJSON = string(output)
	return selection, nil
}

// selectCLI fixes the source and installation identity before running metadata.
func selectCLI(ctx context.Context, options CLIOptions) (*CLISelection, error) {
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
			selection.Provenance = layoutProvenance
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
		if selection.Provenance == layoutProvenance {
			return nil, errCLILayout
		}
		return nil, errCLIInput
	}
	return selection, nil
}

// metadataArgs never inherits endpoint intent from a later environment value.
func (c *Client) metadataArgs(arguments ...string) []string {
	argv := make([]string, 0, 5+len(arguments))
	argv = append(argv, flagServer, c.server, flagNoAutostart, flagTimeout, "2s")
	return append(argv, arguments...)
}

// checkMethod reads the immutable receipt, never re-fetches a descriptor during
// a resource call. The execution boundary separately rechecks server identity.
func (c *Client) checkMethod(method string) error {
	if c.selection.Capabilities != nil {
		for _, entry := range c.selection.Capabilities.Methods {
			if entry.Name == method {
				return nil
			}
		}
	}
	return errCLICompatibility
}

func outputFailure(data json.RawMessage, required methodContract) string {
	// Deliberately discarded outputs need no shape.
	if len(required.output) == 0 && (len(data) == 0 || strings.TrimSpace(string(data)) == schemaNull) {
		return ""
	}
	schema, err := metadataObject(data)
	if err != nil || !schemaType(schema, schemaObject) {
		return required.name + ": object schema/type"
	}
	for _, path := range sortedKeys(required.output) {
		if !schemaHasField(schema, strings.Split(path, "."), required.output[path]) {
			return required.name + "." + path + ": consumed field schema/type"
		}
	}
	return ""
}

// Request paths describe fixed producers, not invented server properties.
// Diagnostics use only these app-owned paths and fixed constraint names.
func requestFailure(schema map[string]json.RawMessage, kind string, fields, supplied map[string]string, path string) string {
	if constraint := unsupportedRequest(schema, kind, path); constraint != "" {
		return path + ": " + constraint
	}
	if len(schema["type"]) != 0 && !schemaType(schema, kind) {
		return path + ": type"
	}
	if constraint := valuesFailure(schema, requestValues(path, kind)); constraint != "" {
		return path + ": " + constraint
	}
	switch kind {
	case schemaObject:
		return requestObjectFailure(schema, fields, supplied, path)
	case schemaArray:
		return requestArrayFailure(schema, fields, supplied, path)
	default:
		if boundsItems(path) {
			if constraint := boundsRangeFailure(schema); constraint != "" {
				return path + ": " + constraint
			}
		}
		return ""
	}
}

// This is intentionally not a general JSON Schema evaluator. Only bounds have
// a proved fixed length and numeric range; other ranges/unknown constraints
// still fail closed. Never include an unknown server keyword in a diagnostic.
func unsupportedRequest(schema map[string]json.RawMessage, kind, path string) string {
	for _, keyword := range sortedKeys(schema) {
		if requestAnnotation(keyword) || keyword == "type" {
			continue
		}
		switch keyword {
		case "properties", "required", "additionalProperties":
			if kind == schemaObject {
				continue
			}
		case "items":
			if kind == schemaArray {
				continue
			}
		case schemaMinItems, schemaMaxItems:
			if boundsRequest(path) {
				continue
			}
		case schemaMinimum, schemaMaximum:
			if boundsItems(path) {
				continue
			}
		case schemaEnum, schemaConst:
			if kind != schemaObject && kind != schemaArray {
				continue
			}
		}
		return constraintName(keyword)
	}
	return ""
}

func constraintName(keyword string) string {
	switch keyword {
	case "type", "properties", "required", "additionalProperties", "items",
		schemaMinItems, schemaMaxItems, schemaMinimum, schemaMaximum, "exclusiveMinimum",
		"exclusiveMaximum", schemaEnum, schemaConst, "pattern", "format", "minLength",
		"maxLength", "allOf", "anyOf", "oneOf", "not", "$ref", "prefixItems",
		"additionalItems", "uniqueItems", "contains", "multipleOf", "minProperties",
		"maxProperties", "patternProperties", "propertyNames", "dependencies",
		"dependentRequired", "dependentSchemas", "if", "then", "else":
		return keyword
	default:
		return "unsupported constraint"
	}
}

func sortedKeys[T any](values map[string]T) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func boundsRequest(path string) bool {
	return path == "session.new_layer.bounds" || path == "session.resize_layer.bounds"
}

func boundsItems(path string) bool {
	return path == "session.new_layer.bounds.[]" || path == "session.resize_layer.bounds.[]"
}

// Bounds.valid proves every emitted coordinate lies in [0,1]. These closed
// endpoints are conservative for all four positions; tighter item constraints
// are not silently waived, even if some individual positions might satisfy them.
func boundsRangeFailure(schema map[string]json.RawMessage) string {
	for _, keyword := range []string{schemaMinimum, schemaMaximum} {
		if data, present := schema[keyword]; present {
			limit, ok := boundedSchemaNumber(data)
			if !ok || (keyword == schemaMinimum && limit.Sign() > 0) ||
				(keyword == schemaMaximum && limit.Cmp(big.NewRat(1, 1)) < 0) {
				return keyword
			}
		}
	}
	return ""
}

// Compare schema limits exactly: float rounding near 0 or 1 could admit a
// constraint that excludes a valid coordinate. Bound decimal/exponent work
// independently of the descriptor's byte cap; unsupported numbers fail closed.
func boundedSchemaNumber(data json.RawMessage) (*big.Rat, bool) {
	number := strings.TrimSpace(string(data))
	if len(number) == 0 || len(number) > 128 {
		return nil, false
	}
	if index := strings.IndexAny(number, "eE"); index >= 0 {
		exponent, err := strconv.Atoi(number[index+1:])
		if err != nil || exponent < -1024 || exponent > 1024 {
			return nil, false
		}
	}
	return new(big.Rat).SetString(number)
}

func requestArrayFailure(schema map[string]json.RawMessage, fields, supplied map[string]string, path string) string {
	for _, keyword := range []string{schemaMinItems, schemaMaxItems} {
		if data, present := schema[keyword]; present {
			var limit uint64
			if strings.TrimSpace(string(data)) == schemaNull || json.Unmarshal(data, &limit) != nil ||
				(keyword == schemaMinItems && limit > 4) || (keyword == schemaMaxItems && limit < 4) {
				return path + ": " + keyword
			}
		}
	}
	itemKind, known := fields["[]"]
	if !known {
		return path + ": items"
	}
	itemPath := path + ".[]"
	return requestPropertyFailure(schema["items"], itemKind, childFields(fields, "[]"), childFields(supplied, "[]"), itemPath)
}

func requestPropertyFailure(data json.RawMessage, kind string, fields, supplied map[string]string, path string) string {
	if len(data) == 0 || strings.TrimSpace(string(data)) == "true" {
		return ""
	}
	schema, err := metadataObject(data)
	if err != nil {
		return path + ": schema"
	}
	return requestFailure(schema, kind, fields, supplied, path)
}

func requestAnnotation(keyword string) bool {
	switch keyword {
	case "title", "description", "default", "examples", "$comment", "deprecated", "readOnly", "writeOnly":
		return true
	default:
		return false
	}
}

// Only string/boolean values fixed by non-test producers are enumerated.
// Finite numeric sets and restrictions on invocation-dependent strings are
// unsupported; the separately proved bounds ranges are not finite sets.
// Variable booleans use both values, not a guessed default.
func requestValues(path, kind string) []any {
	switch path {
	case "session.new_layer.layout.block.flavor":
		return []any{"com.superlogical.terminal.shell"}
	case "session.new_layer.layout.block.options.shell":
		return []any{"none"}
	case "session.new_layer.layout.block.options.exit.on_completion":
		return []any{false}
	case "session.new_layer.focus":
		return []any{true}
	case "client.perform.origin":
		return []any{"cli"}
	case "block.method.method":
		return []any{"size"}
	default:
		if kind == schemaBoolean {
			return []any{false, true}
		}
		return nil
	}
}

func valuesFailure(schema map[string]json.RawMessage, values []any) string {
	if data, present := schema[schemaConst]; present {
		var constant any
		if len(values) == 0 || json.Unmarshal(data, &constant) != nil {
			return schemaConst
		}
		for _, value := range values {
			if !reflect.DeepEqual(value, constant) {
				return schemaConst
			}
		}
	}
	if data, present := schema[schemaEnum]; present {
		var allowed []any
		if len(values) == 0 || json.Unmarshal(data, &allowed) != nil || len(allowed) == 0 {
			return schemaEnum
		}
		for _, value := range values {
			if !containsValue(allowed, value) {
				return schemaEnum
			}
		}
	}
	return ""
}

func containsValue(allowed []any, value any) bool {
	for _, candidate := range allowed {
		if reflect.DeepEqual(candidate, value) {
			return true
		}
	}
	return false
}

func requestObjectFailure(schema map[string]json.RawMessage, fields, supplied map[string]string, path string) string {
	if reason := requiredFailure(schema["required"], fields, supplied, path); reason != "" {
		return reason
	}
	properties := make(map[string]json.RawMessage)
	if data, present := schema["properties"]; present {
		var err error
		properties, err = metadataObject(data)
		if err != nil {
			return path + ": properties"
		}
	}
	additional := schema["additionalProperties"]
	if opaqueRequest(path) {
		// CLI/user-owned children are not absent. Require all possible names
		// and values to remain unconstrained, not merely optional.
		if !unconstrainedAdditional(additional) {
			return path + ": additionalProperties"
		}
		for _, property := range properties {
			if !unconstrainedRequest(property) {
				return path + ": properties"
			}
		}
		return ""
	}
	for _, name := range sortedKeys(fields) {
		if strings.Contains(name, ".") {
			continue
		}
		property, declared := properties[name]
		if !declared {
			// JSON Schema defaults to allowing absent property declarations.
			// Schema-valued additionalProperties constrains those values;
			// false closes the object even for optional producer fields.
			property = additional
			if strings.TrimSpace(string(property)) == "false" {
				return path + "." + name + ": additionalProperties"
			}
		}
		if reason := requestPropertyFailure(property, fields[name], childFields(fields, name), childFields(supplied, name), path+"."+name); reason != "" {
			return reason
		}
	}
	return ""
}

func unconstrainedAdditional(data json.RawMessage) bool {
	return len(data) == 0 || unconstrainedRequest(data)
}

func opaqueRequest(path string) bool {
	switch path {
	case "client.perform.args", "block.method.args", "session.new_window.layout", "session.create.initial_windows.[]":
		return true
	default:
		return false
	}
}

func unconstrainedRequest(data json.RawMessage) bool {
	if strings.TrimSpace(string(data)) == "true" {
		return true
	}
	schema, err := metadataObject(data)
	if err != nil {
		return false
	}
	for keyword := range schema {
		if !requestAnnotation(keyword) {
			return false
		}
	}
	return true
}

func childFields(fields map[string]string, name string) map[string]string {
	children := make(map[string]string)
	for path, kind := range fields {
		if child, found := strings.CutPrefix(path, name+"."); found {
			children[child] = kind
		}
	}
	return children
}

func requiredFailure(data json.RawMessage, fields, supplied map[string]string, path string) string {
	if len(data) == 0 {
		return ""
	}
	var required []string
	if err := json.Unmarshal(data, &required); err != nil || required == nil {
		return path + ": required array"
	}
	seen := make(map[string]bool, len(required))
	for _, field := range required {
		if _, known := fields[field]; !known || strings.Contains(field, ".") {
			return path + ": required unproved field"
		}
		if seen[field] {
			return path + ": required uniqueness"
		}
		if _, guaranteed := supplied[field]; !guaranteed {
			// The name is safe only after membership in the app-owned map.
			return path + "." + field + ": required"
		}
		seen[field] = true
	}
	return ""
}

func schemaHasField(schema map[string]json.RawMessage, path []string, kind string) bool {
	if path[0] == "[]" {
		item, err := metadataObject(schema["items"])
		if err != nil || !schemaType(schema, schemaArray) {
			return false
		}
		if len(path) == 1 {
			return schemaType(item, kind)
		}
		return schemaType(item, schemaObject) && schemaHasField(item, path[1:], kind)
	}
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
	parentKind := schemaObject
	if path[1] == "[]" {
		parentKind = schemaArray
	}
	return schemaType(field, parentKind) && schemaHasField(field, path[1:], kind)
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
		if value != schemaNull && value != kind {
			return false
		}
		found = found || value == kind
	}
	return found
}

// APICapabilities is an invocation-local receipt, not a process-global cache.
// Descriptors are exact bounded server-help entries, retained for root receipt
// binding. Only documented input_schema/output_schema entries are supported.
// This proves the structural operations below, not arbitrary JSON Schema support
// or compatibility with a release that has not supplied this evidence.
type APICapabilities struct {
	Endpoint        string           `json:"endpoint"`
	SelectionSHA256 string           `json:"selection_sha256"`
	Server          ServerIdentity   `json:"server"`
	Version         CLIVersion       `json:"version"`
	Methods         []MethodEvidence `json:"methods"`
	SHA256          string           `json:"sha256"`
}

// ServerIdentity is the minimum freshness binding. Host IDs are intentionally
// absent: equality is neither authorization nor execution-host provenance.
type ServerIdentity struct {
	InstanceID string `json:"instance_id"`
	ConfigHash string `json:"config_hash"`
	Generation uint64 `json:"generation"`
}

// MethodEvidence retains the documented descriptor without mutable JSON maps.
type MethodEvidence struct {
	Name string `json:"name"`
	JSON string `json:"json"`
}

type methodContract struct {
	name          string
	input, output map[string]string
}

// These are the actual non-test Call methods and the API operations behind the
// documented resource CLI wrappers. CLI layout construction and CLI result
// formatting stay with Rex: their API output schemas are retained, not assumed
// identical to CLI JSON. This is a fixed subset, not a schema framework.
func requiredMethods() []methodContract {
	return []methodContract{
		{"server.version", nil, map[string]string{"rex_version": schemaString}},
		{"server.status", nil, map[string]string{"instance_id": schemaString, "config.hash": schemaString, "config.generation": schemaInteger}},
		{"session.list", nil, map[string]string{"sessions": schemaArray, "sessions.[].session_id": schemaString, "sessions.[].label": schemaString}},
		{"session.list_clients", map[string]string{fieldSession: schemaString}, map[string]string{"clients": schemaArray, "clients.[].client_id": schemaString, "clients.[].connected_at": schemaString, "clients.[].connection_state": schemaString, "clients.[].info.kind": schemaString}},
		{"session.describe_block", map[string]string{fieldSession: schemaString, fieldBlock: schemaString}, map[string]string{"block.block_id": schemaString, "block.window_id": schemaString, "block.placement": schemaString, "block.owner": schemaString}},
		{"session.focus_block", map[string]string{fieldSession: schemaString, fieldBlock: schemaString}, nil},
		{"session.resize_layer", map[string]string{fieldSession: schemaString, "layer_id": schemaString, fieldBounds: schemaArray, "bounds.[]": schemaNumber}, nil},
		{"session.new_layer", map[string]string{fieldSession: schemaString, fieldWindow: schemaString, fieldBounds: schemaArray, "bounds.[]": schemaNumber, fieldFocus: schemaBoolean,
			fieldLayout: schemaObject, "layout.block": schemaObject, "layout.block.flavor": schemaString,
			"layout.block.options": schemaObject, "layout.block.options.command": schemaArray, "layout.block.options.command.[]": schemaString,
			"layout.block.options.shell": schemaString, "layout.block.options.exit": schemaObject, "layout.block.options.exit.on_completion": schemaBoolean,
			"layout.block.options.size": schemaObject, "layout.block.options.size.columns": schemaInteger, "layout.block.options.size.rows": schemaInteger,
			"layout.block.options.size.cell_width_px": schemaInteger, "layout.block.options.size.cell_height_px": schemaInteger}, map[string]string{"layer_id": schemaString, "block_ids": schemaArray, "block_ids.[]": schemaString}},
		{"client.perform", map[string]string{"client_id": schemaString, "action": schemaString, "args": schemaObject, "timeout_ms": schemaInteger, "origin": schemaString, fieldSession: schemaString, fieldBlock: schemaString}, nil},
		{"session.view", map[string]string{fieldSession: schemaString}, nil},
		{"session.create", map[string]string{"initial_windows": schemaArray, "initial_windows.[]": schemaObject}, nil},
		{"session.set_label", map[string]string{fieldSession: schemaString, fieldLabel: schemaString}, nil},
		{"session.destroy", map[string]string{fieldSession: schemaString}, nil},
		{"session.new_window", map[string]string{fieldSession: schemaString, fieldLayout: schemaObject}, nil},
		{"session.focus_window", map[string]string{fieldSession: schemaString, fieldWindow: schemaString}, nil},
		{"session.focus_next_window", map[string]string{fieldSession: schemaString, "wrap": schemaBoolean}, nil},
		{"session.focus_previous_window", map[string]string{fieldSession: schemaString, "wrap": schemaBoolean}, nil},
		{"session.set_window_label", map[string]string{fieldSession: schemaString, fieldWindow: schemaString, fieldLabel: schemaString}, nil},
		{"session.close_window", map[string]string{fieldSession: schemaString, fieldWindow: schemaString}, nil},
		{"block.close", map[string]string{fieldSession: schemaString, fieldBlock: schemaString}, nil},
		{"block.method", map[string]string{fieldSession: schemaString, fieldBlock: schemaString, "method": schemaString, "args": schemaObject}, nil},
	}
}

func (selection *CLISelection) clone() *CLISelection {
	snapshot := *selection
	snapshot.Layout = append([]FileIdentity(nil), selection.Layout...)
	if selection.Capabilities != nil {
		receipt := *selection.Capabilities
		receipt.Methods = append([]MethodEvidence(nil), receipt.Methods...)
		snapshot.Capabilities = &receipt
	}
	return &snapshot
}

func (selection *CLISelection) fingerprint() (string, error) {
	snapshot := *selection
	snapshot.Capabilities = nil
	return metadataDigest(snapshot)
}

func metadataDigest(value any) (string, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return "", errCLICompatibility
	}
	digest := sha256.Sum256(data)
	return hex.EncodeToString(digest[:]), nil
}

func (receipt APICapabilities) fingerprint() (string, error) {
	receipt.SHA256 = ""
	return metadataDigest(receipt)
}

func (c *Client) captureCapabilities(ctx context.Context) error {
	before, err := c.selection.serverIdentity(ctx, c.server)
	if err != nil {
		return err
	}
	versionJSON, err := runSelected(ctx, c.selection, c.metadataArgs("api", "call", "--", "server.version"), nil, 64<<10)
	if err != nil {
		return errors.Join(compatibilityFailure("server.version", "command failed"), err)
	}
	version, err := decodeVersion(versionJSON)
	if err != nil {
		return err
	}
	// One bounded help snapshot avoids independently fetched descriptors spanning
	// generations. The CLI docs specify api list --json as help.all's document.
	data, err := runSelected(ctx, c.selection, c.metadataArgs("api", "list", flagJSON), nil, 8<<20)
	if err != nil {
		return errors.Join(compatibilityFailure("api list", "command failed"), err)
	}
	methods, err := captureMethods(data)
	if err != nil {
		return err
	}
	receipt := &APICapabilities{Endpoint: c.server, Server: before, Version: version, Methods: methods}
	receipt.SelectionSHA256, err = c.selection.fingerprint()
	if err != nil {
		return err
	}
	receipt.SHA256, err = receipt.fingerprint()
	if err != nil {
		return err
	}
	c.selection.Capabilities = receipt
	// Any change across capture invalidates the whole invocation, not just the
	// last descriptor. No new endpoint, generation or candidate is selected.
	return c.selection.checkServer(ctx)
}

func captureMethods(data []byte) ([]MethodEvidence, error) {
	help, err := metadataObject(data)
	if err != nil {
		return nil, compatibilityFailure("api list", "object")
	}
	var entries []json.RawMessage
	if decodeErr := json.Unmarshal(help["methods"], &entries); decodeErr != nil || len(entries) == 0 || len(entries) > 4096 {
		return nil, compatibilityFailure("api list", "methods array/limit")
	}
	byName := make(map[string]string, len(entries))
	for _, raw := range entries {
		entry, entryErr := metadataObject(raw)
		var name string
		if entryErr != nil || json.Unmarshal(entry["name"], &name) != nil || !metadataText(name) || byName[name] != "" {
			return nil, compatibilityFailure("api list", "method entry/name uniqueness")
		}
		byName[name] = string(raw)
	}
	var methods []MethodEvidence
	for _, required := range requiredMethods() {
		raw := byName[required.name]
		if reason := methodFailure(raw, required); reason != "" {
			return nil, compatibilityFailure("api list", reason)
		}
		methods = append(methods, MethodEvidence{Name: required.name, JSON: raw})
	}
	return methods, nil
}

// A failure contains only fixed stage/constraint names and app-owned contracts.
// Never format raw descriptors, unknown keys, child output or invocation data.
func compatibilityFailure(stage, reason string) error {
	return fmt.Errorf("%w (%s: %s)", errCLICompatibility, stage, reason)
}

func methodFailure(raw string, required methodContract) string {
	if len(raw) == 0 {
		return required.name + ": missing method"
	}
	if len(raw) > 64<<10 {
		return required.name + ": descriptor limit"
	}
	entry, err := metadataObject([]byte(raw))
	var name string
	if err != nil || json.Unmarshal(entry["name"], &name) != nil || name != required.name {
		return required.name + ": descriptor"
	}
	if reason := inputFailure(entry["input_schema"], required); reason != "" {
		return "input " + reason
	}
	if reason := outputFailure(entry["output_schema"], required); reason != "" {
		return "output " + reason
	}
	return ""
}

func inputFailure(data json.RawMessage, method methodContract) string {
	if len(method.input) == 0 && (len(data) == 0 || strings.TrimSpace(string(data)) == schemaNull) {
		return ""
	}
	schema, err := metadataObject(data)
	if err != nil {
		return method.name + ": schema"
	}
	return requestFailure(schema, schemaObject, method.input, suppliedFields(method), method.name)
}

func suppliedFields(method methodContract) map[string]string {
	supplied := make(map[string]string, len(method.input))
	for key, kind := range method.input {
		supplied[key] = kind
	}
	switch method.name {
	case "client.perform":
		delete(supplied, fieldSession)
		delete(supplied, fieldBlock)
	case "block.method":
		delete(supplied, "args")
	case "session.new_layer":
		// Unsized creation omits size; sized creation always supplies rows and
		// columns, but omits zero pixel dimensions. Required is checked at each
		// parent even when its optional child has a known conditional shape.
		delete(supplied, "layout.block.options.size")
		delete(supplied, "layout.block.options.size.cell_width_px")
		delete(supplied, "layout.block.options.size.cell_height_px")
	}
	return supplied
}

func (selection *CLISelection) validateCapabilities() error {
	receipt := selection.Capabilities
	endpoint, err := CanonicalEndpoint(receipt.Endpoint)
	if err != nil || endpoint != receipt.Endpoint || !metadataText(receipt.Server.InstanceID) || !metadataText(receipt.Server.ConfigHash) ||
		!metadataText(receipt.Version.Version) || receipt.Version.APIMinRevision > receipt.Version.APIRevision {
		return errCLIChanged
	}
	identity, err := selection.fingerprint()
	if err != nil || identity != receipt.SelectionSHA256 {
		return errCLIChanged
	}
	digest, err := receipt.fingerprint()
	if err != nil || digest != receipt.SHA256 {
		return errCLIChanged
	}
	required := requiredMethods()
	if len(required) != len(receipt.Methods) {
		return errCLIChanged
	}
	for index, method := range required {
		entry := receipt.Methods[index]
		if entry.Name != method.name {
			return compatibilityFailure("frozen receipt", method.name+": method order")
		}
		if reason := methodFailure(entry.JSON, method); reason != "" {
			return compatibilityFailure("frozen receipt", reason)
		}
	}
	return nil
}

func metadataText(value string) bool {
	return value != "" && len(value) <= 256 && utf8.ValidString(value) && strings.IndexFunc(value, unicode.IsControl) < 0
}

func (selection *CLISelection) serverIdentity(ctx context.Context, endpoint string) (ServerIdentity, error) {
	ctx, cancel := context.WithTimeout(ctx, endpointTimeout)
	defer cancel()
	data, err := runSelected(ctx, selection, []string{flagServer, endpoint, flagNoAutostart, flagTimeout, "2s", "api", "call", "--", "server.status"}, nil, 64<<10)
	if err != nil {
		return ServerIdentity{}, errors.Join(compatibilityFailure("server.status", "command failed"), err)
	}
	object, err := metadataObject(data)
	if err != nil {
		return ServerIdentity{}, compatibilityFailure("server.status", "object")
	}
	config, err := metadataObject(object["config"])
	if err != nil {
		return ServerIdentity{}, compatibilityFailure("server.status", "config object")
	}
	var identity ServerIdentity
	if json.Unmarshal(object["instance_id"], &identity.InstanceID) != nil || !metadataText(identity.InstanceID) {
		return ServerIdentity{}, compatibilityFailure("server.status", "instance_id text")
	}
	if json.Unmarshal(config["hash"], &identity.ConfigHash) != nil || !metadataText(identity.ConfigHash) {
		return ServerIdentity{}, compatibilityFailure("server.status", "config.hash text")
	}
	if json.Unmarshal(config["generation"], &identity.Generation) != nil || strings.TrimSpace(string(config["generation"])) == schemaNull {
		return ServerIdentity{}, compatibilityFailure("server.status", "config.generation unsigned integer")
	}
	return identity, nil
}

func (selection *CLISelection) checkServer(ctx context.Context) error {
	if selection == nil || selection.Capabilities == nil {
		return errCLIChanged
	}
	actual, err := selection.serverIdentity(ctx, selection.Capabilities.Endpoint)
	if err != nil {
		return err
	}
	if actual != selection.Capabilities.Server {
		return errCLIChanged
	}
	return nil
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

func runtimeAnchor(options CLIOptions) (*string, error) {
	if options.Anchor != nil {
		if options.LayoutRevision != 1 {
			return nil, errCLILayout
		}
		return options.Anchor, nil
	}
	if options.LayoutRevision != 0 {
		return nil, errCLILayout
	}
	self, err := os.Executable()
	if err != nil {
		return nil, errCLILayout
	}
	if !strings.HasSuffix(filepath.ToSlash(self), "/rex_whichkey/runtime/rex-whichkey") {
		return nil, nil
	}
	candidate := filepath.Join(filepath.Dir(self), "anchor.lua")
	return &candidate, nil
}

func layoutDependency(ctx context.Context, options CLIOptions) (string, []FileIdentity, bool, error) {
	anchor, err := runtimeAnchor(options)
	if err != nil || anchor == nil {
		return "", nil, false, err
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
	case "internal-explicit", "execution-env", "terminal-env", layoutProvenance, "path":
	default:
		return errCLIChanged
	}
	version, err := metadataObject([]byte(selection.VersionJSON))
	if err != nil || len(version) == 0 || len(selection.VersionJSON) > 64<<10 ||
		(selection.Provenance == layoutProvenance && len(selection.Layout) != 2) ||
		(selection.Provenance != layoutProvenance && len(selection.Layout) != 0) {
		return errCLIChanged
	}
	if err := selection.checkFiles(ctx); err != nil {
		return err
	}
	if selection.Capabilities != nil {
		if err := selection.validateCapabilities(); err != nil {
			return err
		}
		return selection.checkServer(ctx)
	}
	return nil
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
		if closeErr := file.Close(); closeErr != nil {
			resultErr = errCLIChanged
		}
	}()
	before, err := file.Stat()
	if err != nil {
		return identity, errCLIChanged
	}
	stat, ok := trustedFile(before, executable)
	if !ok {
		return identity, errCLIChanged
	}
	digest, count, err := fileDigest(ctx, file)
	if err != nil {
		return identity, err
	}
	if !unchangedFile(ctx, file, source, canonical, links, before, count) {
		return identity, errCLIChanged
	}
	return FileIdentity{Source: source, Path: canonical, SHA256: digest, LinksSHA256: links,
		Device: uint64(stat.Dev), Inode: stat.Ino, Size: count, ModifiedNS: before.ModTime().UnixNano(),
		Mode: uint32(before.Mode()), Owner: stat.Uid}, nil
}

func trustedFile(info os.FileInfo, executable bool) (*syscall.Stat_t, bool) {
	if !info.Mode().IsRegular() || info.Size() <= 0 || info.Size() > maxExecutableBytes ||
		info.Mode().Perm()&0022 != 0 || (executable && info.Mode().Perm()&0111 == 0) {
		return nil, false
	}
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok || (stat.Uid != 0 && stat.Uid != uint32(os.Geteuid())) {
		return nil, false
	}
	return stat, true
}

func fileDigest(ctx context.Context, file *os.File) (string, int64, error) {
	digest := sha256.New()
	buffer := make([]byte, 128<<10)
	var count int64
	for {
		if err := ctx.Err(); err != nil {
			return "", 0, err
		}
		n, readErr := file.Read(buffer)
		count += int64(n)
		if count > maxExecutableBytes {
			return "", 0, errCLIChanged
		}
		if _, err := digest.Write(buffer[:n]); err != nil {
			return "", 0, errCLIChanged
		}
		if errors.Is(readErr, io.EOF) {
			return hex.EncodeToString(digest.Sum(nil)), count, nil
		}
		if readErr != nil {
			return "", 0, errCLIChanged
		}
	}
}

// The open file, logical link chain and canonical target must still agree after
// hashing. A metadata/content race invalidates selection instead of choosing again.
func unchangedFile(ctx context.Context, file *os.File, source, canonical, links string, before os.FileInfo, count int64) bool {
	after, err := file.Stat()
	linked, linkErr := os.Stat(source)
	resolved, resolveErr := filepath.EvalSymlinks(source)
	afterLinks, linksErr := installationLinks(ctx, source)
	return err == nil && linkErr == nil && resolveErr == nil && linksErr == nil && afterLinks == links && resolved == canonical &&
		os.SameFile(before, after) && os.SameFile(after, linked) && count == before.Size() && before.Size() == after.Size() &&
		before.Mode() == after.Mode() && before.ModTime().Equal(after.ModTime())
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
		if info.Mode().Perm()&0022 != 0 && (stat.Uid != 0 || info.Mode()&os.ModeSticky == 0) {
			return false
		}
		if filepath.Dir(directory) == directory {
			return true
		}
	}
}
