{
  lib,
  buildGoModule,
}:
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
      || builtins.elem relative ["go.mod" "go.sum" "LICENSE" "docs/mapping-format.md"]
      || (!lib.hasInfix "/" relative && lib.hasSuffix ".go" relative);
  };

  vendorHash = "sha256-Y29DnvGzSEIrKl4wZStCRsmpF44QpR3Lfzd9GdxzhlY=";
  subPackages = ["cmd"];
  ldflags = ["-s" "-w"];

  # Current delivery policy permits compilation, not project or runtime tests.
  doCheck = false;
  doInstallCheck = false;

  postInstall = ''
    mv "$out/bin/cmd" "$out/bin/rex-whichkey"

    # The server-side module consumer opts into this path; never edit init.lua
    # or wrap the caller/client's Lua runtime or Rex executable.
    mkdir -p "$out/share/rex-whichkey/lua"
    cp -R lua/. "$out/share/rex-whichkey/lua/"
    install -Dm644 docs/mapping-format.md "$out/share/doc/rex-whichkey/mapping-format.md"
    install -Dm644 LICENSE "$out/share/licenses/rex-whichkey/LICENSE"
  '';

  meta = {
    description = "Which-key menu and opt-in server-side Lua integration for Rex";
    homepage = "https://github.com/curtbushko/rex-whichkey";
    license = lib.licenses.mit;
    mainProgram = "rex-whichkey";
    platforms = lib.platforms.linux ++ lib.platforms.darwin;
  };
}
