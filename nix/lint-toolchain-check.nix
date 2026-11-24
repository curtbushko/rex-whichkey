{
  pkgs,
  toolchain,
}:
pkgs.runCommand "lint-toolchain" {
  nativeBuildInputs = [toolchain.go toolchain.golangciLint toolchain.goAiLint];
  # Keep this offline and bounded, including analyzer subprocesses.
  GOTOOLCHAIN = "local";
  GOPROXY = "off";
  GOSUMDB = "off";
  GOMAXPROCS = "2";
  GOMEMLIMIT = "2GiB";
} ''
  #!/usr/bin/env bash
  set -euo pipefail
  : "''${out:?Nix must provide the output path}"
  export HOME="$TMPDIR/home" GOCACHE="$TMPDIR/go-cache" GOPATH="$TMPDIR/go"
  export GOLANGCI_LINT_CACHE="$TMPDIR/golangci-cache"
  mkdir -p "$HOME" fixture
  cd fixture

  # Use the exact default-shell derivations, not an inherited compatible prefix.
  test "$(command -v go)" = '${toolchain.go}/bin/go'
  test "$(command -v golangci-lint)" = '${toolchain.golangciLint}/bin/golangci-lint'
  test "$(command -v go-ai-lint)" = '${toolchain.goAiLint}/bin/go-ai-lint'
  test "$(go env GOVERSION)" = 'go${toolchain.go.version}'
  printf 'module example.com/lint-toolchain\n\ngo ${toolchain.go.version}\n' > go.mod
  cat > .golangci.yml << 'YAML'
  version: "2"
  linters:
    default: none
    enable: [govet]
  YAML
  cat > fixture.go << 'GO'
  package fixture

  import "fmt"

  func Message() string { return fmt.Sprintf("%s", "ready") }
  GO
  go test ./...
  golangci-lint run --config .golangci.yml ./...
  go-ai-lint -paniclint ./...

  # Known diagnostics prove both analyzers really load/type-check this Go module.
  cat > fixture.go << 'GO'
  package fixture

  import "fmt"

  func Message() string { return fmt.Sprintf("%d", "wrong") }
  func Fail() { panic("unexpected") }
  GO
  if golangci-lint run --config .golangci.yml ./... > golangci.log 2>&1; then
    echo 'golangci-lint missed the printf diagnostic' >&2
    exit 1
  fi
  grep -F '(govet)' golangci.log
  if go-ai-lint -paniclint ./... > ai.log 2>&1; then
    echo 'go-ai-lint missed the panic diagnostic' >&2
    exit 1
  fi
  grep -F 'AIL090: panic in library code' ai.log
  # Raw diagnostics contain temporary build paths; retain only a stable receipt.
  # The assertions above still print the matched diagnostics in the build log.
  mkdir -p "$out"
  printf '%s\n' \
    'Clean fixture: go test, golangci-lint and go-ai-lint passed' \
    'Injected diagnostics: govet and AIL090 detected' > "$out/passed"
''
