{
  config,
  lib,
  pkgs,
  ...
}: let
  cfg = config.programs.rex-whichkey;
in {
  options.programs.rex-whichkey = {
    enable = lib.mkEnableOption "the Rex which-key package and launch-only Lua module";

    package = lib.mkOption {
      type = lib.types.package;
      default = pkgs.callPackage ./package.nix {};
      defaultText = lib.literalExpression "pkgs.callPackage ./package.nix {}";
      description = ''
        Package providing the real bin/rex-whichkey and
        share/rex-whichkey/lua/rex_whichkey.lua launch module.
        The flake module defaults to this flake's package for the host platform.
        Exposes the same optional licensed execution-host CLI as bin/rex
        where published. The pinned source has no Intel macOS CLI artifact;
        provide a legitimate rex on the execution-host PATH there.
        Go uses ordinary PATH lookup and independent endpoint policy, without
        binary receipts or startup version/API/schema negotiation. The running
        Rex server must already inherit PATH for both rex-whichkey and rex;
        installing a package does not update an existing process environment.
        Installation does not configure data, bindings or a Rex service,
        select an endpoint, or reload the server.
      '';
    };

    configDirectory = lib.mkOption {
      type = lib.types.str;
      default = "${config.xdg.configHome}/rex";
      defaultText = lib.literalExpression ''config.xdg.configHome + "/rex"'';
      description = ''
        Directory containing the server's user-owned Rex init.lua, inside the
        Home Manager user's home. For a custom REX_CONFIG, select its parent
        directory here. This does not set REX_CONFIG or modify init.lua.
        Only the reserved rex_whichkey.lua launch module is linked.
        No module directory, init.lua or menu data is managed.
      '';
    };
  };

  config = lib.mkIf cfg.enable {
    home.packages = [cfg.package];

    # Link only the integration namespace, never the enclosing Rex directory.
    # Leave conflict handling to Home Manager; do not force unmanaged files away.
    home.file = {
      "${cfg.configDirectory}/rex_whichkey.lua" = {
        source = "${cfg.package}/share/rex-whichkey/lua/rex_whichkey.lua";
        force = false;
      };
    };
  };
}
