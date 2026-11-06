{
  config,
  lib,
  pkgs,
  ...
}: let
  cfg = config.programs.rex-whichkey;
in {
  options.programs.rex-whichkey = {
    enable = lib.mkEnableOption "the opt-in Rex which-key package and server-side Lua modules";

    package = lib.mkOption {
      type = lib.types.package;
      default = pkgs.callPackage ./package.nix {};
      defaultText = lib.literalExpression "pkgs.callPackage ./package.nix {}";
      description = ''
        Package providing bin/rex-whichkey and share/rex-whichkey/lua.
        The flake module defaults to this flake's package for the host platform.
        Includes normal layout-1 helper/anchor assets and, where published,
        the licensed execution-host Rex CLI dependency. Installation does not
        select an endpoint, configure a Rex service or reload the server.
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
        Only the reserved rex_whichkey.lua and rex_whichkey module paths are
        linked; existing user mappings must use a different module name.
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
      "${cfg.configDirectory}/rex_whichkey" = {
        source = "${cfg.package}/share/rex-whichkey/lua/rex_whichkey";
        recursive = false;
        force = false;
      };
    };
  };
}
