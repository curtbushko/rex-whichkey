{
  description = "Rex which-key menu and server-side Lua integration";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

  inputs.golang-shared-configs.url = "github:curtbushko/golang-shared-configs";

  outputs = {
    self,
    nixpkgs,
    golang-shared-configs,
  }: let
    goVersion = "1.26.8";
    # Official source checksum; keep the toolchain independent of Nixpkgs patch updates.
    goSourceHash = "sha256-Tjm5jkL5RvoFrIvFtxh335fb23y7Gnd7VBZnrXEX/S4=";
    goAttribute = "go_${builtins.replaceStrings ["."] ["_"] (nixpkgs.lib.versions.majorMinor goVersion)}";

    supportedSystems = ["x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin"];
    forEachSupportedSystem = f:
      nixpkgs.lib.genAttrs supportedSystems (system:
        f {
          inherit system;
          pkgs = import nixpkgs {
            inherit system;
            config.allowUnfree = true;
            overlays = [self.overlays.default];
          };
        });

    # Build go-ai-lint from source
    go-ai-lint = {pkgs}:
      (pkgs.buildGoModule.override {go = pkgs.go;}) {
        pname = "go-ai-lint";
        version = "1.0.0";
        src = pkgs.fetchFromGitHub {
          owner = "curtbushko";
          repo = "go-ai-lint";
          rev = "v1.0.0";
          sha256 = "sha256-y2G7dTZqM/rEQaALu54bHigBeO1xxRIblBJ7QxOffW4=";
        };
        subPackages = ["cmd/go-ai-lint"];
        vendorHash = "sha256-zkXyXTEnMmBZnvzoq0UWKgzWZlyNRyQZCYAv+huZo0I=";
      };
  in {
    overlays.default = final: prev: {
      go = final.${goAttribute}.overrideAttrs (_: {
        version = goVersion;
        src = final.fetchurl {
          url = "https://go.dev/dl/go${goVersion}.src.tar.gz";
          hash = goSourceHash;
        };
      });
    };

    packages = forEachSupportedSystem ({pkgs, ...}: let
      rex-whichkey = pkgs.callPackage ./nix/package.nix {
        buildGoModule = pkgs.buildGoModule.override {go = pkgs.go;};
      };
    in {
      inherit rex-whichkey;
      default = rex-whichkey;
    });

    homeManagerModules = {
      rex-whichkey = {
        lib,
        pkgs,
        ...
      }: {
        imports = [./nix/home-manager.nix];
        programs.rex-whichkey.package = lib.mkDefault self.packages.${pkgs.stdenv.hostPlatform.system}.rex-whichkey;
      };
      default = self.homeManagerModules.rex-whichkey;
    };

    devShells = forEachSupportedSystem ({
      pkgs,
      system,
    }: let
      sharedConfigs = golang-shared-configs.packages.${system}.all-configs;
      make-wrapper = pkgs.writeShellScriptBin "make" ''
        exec ${pkgs.go-task}/bin/task "$@"
      '';
    in {
      default = pkgs.mkShell {
        packages = with pkgs; [
          docker
          # go (version is specified by overlay)
          go
          go-task
          gotools
          golangci-lint
          pre-commit
          (go-ai-lint {inherit pkgs;})
          sharedConfigs
          make-wrapper
        ];

        shellHook = ''
          # Auto-pull if on main branch
          if [ "$(git rev-parse --abbrev-ref HEAD 2>/dev/null)" = "main" ]; then
            echo "On main branch, pulling latest changes..."
            git pull --quiet || true
          fi

          # Enable native git pre-commit hook (build/lint/test)
          if [ -d .githooks ]; then
            git config core.hooksPath .githooks 2>/dev/null || true
          fi

          cp -f ${sharedConfigs}/.golangci.yml .golangci.yml
          cp -f ${sharedConfigs}/.go-ai-lint.yml .go-ai-lint.yml
        '';
      };
    });
  };
}
