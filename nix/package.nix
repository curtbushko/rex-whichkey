{
  lib,
  buildGoModule,
  fetchurl,
  stdenv,
  stdenvNoCC,
  patchelf,
  unzip,
  glibc,
  pam,
  rex-cli ? null,
}: let
  # Official immutable package source inspected by the runtime contract. These
  # native artifact pins come from its sources.json, not a system-package guess.
  rexSources = {
    x86_64-linux = {
      archive = "rex-linux-amd64.tar.gz";
      sha256 = "b7ee40e9f8ce199ad9992a7601536cd03f5d25166f63c06e6dc1cf866b773a5a";
    };
    aarch64-linux = {
      archive = "rex-linux-arm64.tar.gz";
      sha256 = "250e05847fedd78d6ca850a222f0db74038070150e3b502c165b111b3f286284";
    };
    aarch64-darwin = {
      archive = "rex-darwin-arm64.zip";
      sha256 = "dfc3c5d57c9637ac21f41a7fa425a453fed64d68152b90772afb2d815b028977";
    };
  };
  rexPackageSource = fetchurl {
    url = "https://pkgs.superlogical.com/rex/nix/preview/0.1.0-preview.13.tar.gz";
    sha256 = "dc5297710269f1eff712b265465c851f2ddc953a61d4310f1bb81bf9d543ac12";
  };
  system = stdenv.hostPlatform.system;
  dependency =
    if rex-cli != null
    then rex-cli
    else if builtins.hasAttr system rexSources
    then let
      source = rexSources.${system};
    in
      stdenvNoCC.mkDerivation {
        pname = "rex";
        version = "0.1.0-preview.13";
        src = fetchurl {
          url = "https://blob.superlogical.com/download/rex/preview/1b6654811346488b696ded8f3c7e6701c57242f1/${source.archive}";
          inherit (source) sha256;
        };
        sourceRoot = ".";
        nativeBuildInputs = lib.optional stdenv.hostPlatform.isLinux patchelf ++ lib.optional (lib.hasSuffix ".zip" source.archive) unzip;
        # Preserve the released Darwin signature and the Linux dlopen PAM rpath.
        dontStrip = true;
        dontPatchELF = true;
        doCheck = false;
        doInstallCheck = false;
        postUnpack = ''
          mkdir -p "$TMPDIR/rex-package-source"
          tar -xzf ${rexPackageSource} -C "$TMPDIR/rex-package-source"
        '';
        installPhase = ''
          runHook preInstall
          install -Dm755 rex "$out/bin/rex"
          for name in x/xterm-rex 78/xterm-rex r/rex 72/rex; do
            install -Dm644 "$TMPDIR/rex-package-source/rex/xterm-rex.bin" "$out/share/terminfo/$name"
          done
          install -Dm644 "$TMPDIR/rex-package-source/rex/50-rex.conf" "$out/share/rex/ssh_config.d/50-rex.conf"
          runHook postInstall
        '';
        # Upstream requires rpath first, interpreter second (patchelf#674).
        postInstall = lib.optionalString stdenv.hostPlatform.isLinux ''
          patchelf --set-rpath "${lib.getLib glibc}/lib:${lib.getLib pam}/lib" "$out/bin/rex"
          patchelf --set-interpreter "$(cat ${stdenv.cc}/nix-support/dynamic-linker)" "$out/bin/rex"
        '';
        meta = {
          description = "Official proprietary Rex CLI dependency";
          homepage = "https://superlogical.com";
          license = lib.licenses.unfree;
          sourceProvenance = [lib.sourceTypes.binaryNativeCode];
          platforms = builtins.attrNames rexSources;
          mainProgram = "rex";
        };
      }
    # Upstream publishes no Intel macOS artifact in this pinned source. Keep the
    # standalone binary available there; do not invent a restricted-PATH asset.
    else null;
in
  buildGoModule {
    pname = "rex-whichkey";
    version = "0-unstable";

    # Keep workflow transport and historical experiment modules out of the source.
    src = lib.cleanSourceWith {
      src = ../.;
      filter = path: type: let
        relative = lib.removePrefix (toString ../. + "/") (toString path);
      in
        (type == "directory" && builtins.elem relative ["cmd" "rex" "lua" "docs"])
        || lib.hasPrefix "cmd/" relative
        || lib.hasPrefix "rex/" relative
        || lib.hasPrefix "lua/" relative
        || builtins.elem relative ["go.mod" "go.sum" "LICENSE" "docs/mapping-format.md" "docs/runtime-bootstrap.md"]
        || (!lib.hasInfix "/" relative && lib.hasSuffix ".go" relative);
    };

    vendorHash = "sha256-Y29DnvGzSEIrKl4wZStCRsmpF44QpR3Lfzd9GdxzhlY=";
    subPackages = ["cmd"];

    # Current delivery policy permits no project or runtime tests.
    doCheck = false;
    doInstallCheck = false;

    postInstall =
      ''
        mv "$out/bin/cmd" "$out/bin/rex-whichkey"
        mkdir -p "$out/share/rex-whichkey/lua"
        cp -R lua/. "$out/share/rex-whichkey/lua/"
        # Layout revision 1 is declared by the shipped anchor, not generated Lua.
        ln -s ../../../../../bin/rex-whichkey "$out/share/rex-whichkey/lua/rex_whichkey/runtime/rex-whichkey"
        install -Dm644 docs/mapping-format.md "$out/share/doc/rex-whichkey/mapping-format.md"
        install -Dm644 docs/runtime-bootstrap.md "$out/share/doc/rex-whichkey/runtime-bootstrap.md"
        install -Dm644 LICENSE "$out/share/licenses/rex-whichkey/LICENSE"
      ''
      + lib.optionalString (dependency != null) ''
        mkdir -p "$out/share/rex-whichkey/lua/rex_whichkey/runtime/deps"
        ln -s ${dependency}/bin/rex "$out/share/rex-whichkey/lua/rex_whichkey/runtime/deps/rex"
      '';

    meta = {
      description = "Which-key menu and opt-in server-side Lua integration for Rex";
      homepage = "https://github.com/curtbushko/rex-whichkey";
      license = lib.licenses.mit;
      mainProgram = "rex-whichkey";
      platforms = lib.platforms.linux ++ lib.platforms.darwin;
    };
  }
