# Home Manager integration: source-review guide

This is documentation, not an evaluation/activation runner.
Nix checks for this cutover are parse/format/whitespace only; no Nix
evaluation/build, profile activation, Rex reload or installed keybinding result
is claimed.

## Import and opt in on the execution host

In the Home Manager configuration for the user running the Rex server:

```nix
{ inputs, ... }: {
  imports = [ inputs.rex-whichkey.homeManagerModules.default ];
  programs.rex-whichkey.enable = true;
}
```

`homeManagerModules.rex-whichkey` is the same module.
The flake module defaults to its host-platform package;
`programs.rex-whichkey.package` can override it with a compatible package.
Direct import of `nix/home-manager.nix` defaults to
`pkgs.callPackage ./package.nix {}`. Flake import retains pinned Go selection.

When enabled, the module adds the package to `home.packages` and links **only**
`rex_whichkey.lua` to `share/rex-whichkey/lua/rex_whichkey.lua`.
The default destination is `${config.xdg.configHome}/rex`, beside init.lua.
For a custom `REX_CONFIG`, choose its parent inside the user's home:

```nix
{ config, ... }: {
  programs.rex-whichkey = {
    enable = true;
    configDirectory = "${config.home.homeDirectory}/rex-server-config";
  };
}
```

This does not set REX_CONFIG, own the whole Rex directory, change Lua search paths,
generate init.lua/menu JSON or install a keybinding. The module relies on Rex's
documented sibling-module loading. It does not configure a viewing client or
deploy anything remotely.

## User-owned wiring and data

Only if/when you choose to change your init.lua:

```lua
local whichkey = require("rex_whichkey")
rex.bind("ctrl+space", function(ctx)
    return whichkey.launch(ctx)
end)
```

The key is an example, not an installed/reloaded/observed shortcut.
The launcher starts ordinary `rex-whichkey` through documented detached
argument-array execution. Rex's process PATH must already find both
`rex-whichkey` and `rex`; installing the package or updating a terminal profile
does not update the running server's inherited environment.
Go selects `rex` through ordinary execution-host PATH lookup, retains the
endpoint independently and validates exact context IDs. Binary receipts,
startup version/API/schema negotiation and status sandwiches are not required.
No Lua action/runtime/completion registration or anchor/helper path is required.

Go reads optional version-1 JSON from its
[documented selection](../../../docs/mapping-format.md#file-selection).
Home Manager does not import/convert user Lua, copy user data into the Nix store,
or install an automatic replacement config. Preserve old files; manually account
for literal data and deferred functions/hooks/callbacks before migration.
See [examples](../../../examples/README.md).

The real package image is `bin/rex-whichkey`; the same optional pinned licensed
CLI is exposed as ordinary `bin/rex`, with no private libexec lookup or wrapper.
The pinned source has no Intel macOS CLI artifact; provide a legitimate `rex`
on execution-host PATH there, or use a legitimate `rex-cli` override. A source-only
Go build installs neither that CLI nor the Lua/assets. Go, not Nix, owns lookup
and endpoint policy. No dependency, pin, profile or activation change is implied.
See [package layout](../../../docs/runtime-bootstrap.md).

## Preservation and migration checklist (source-only)

- Disabled (the default): `lib.mkIf cfg.enable` contributes no package/file links.
- Enabled: exactly the package and one launch-module link; no obsolete recursive
  namespace/module-directory link, binding, hook, service or activation script.
- The file retains `force=false`; normal Home Manager conflict handling applies.
  Do not force/delete unmanaged files to bypass a conflict.
- A pre-existing whole-Rex-directory link must be accounted for in your own
  declarative layout; this module never takes over init.lua or sibling files.
- Compared with the old module, newly generated configuration no longer declares
  the `rex_whichkey` directory link. Source removal is **not activation** and does
  not remove an existing live generation or authorize deletion of user assets.
- Update old user-owned require/setup/registration wiring deliberately through
  your own lifecycle. No automatic reload/restart or executable-Lua migration.
- Disabling/removing a future managed generation stops this module's contributions,
  not user-owned init.lua/menu data. Adjust your own require/binding if absent.

Parse and format checks do not prove enabled/disabled evaluation, conflict safety
on activation, file preservation, installed PATH, Lua loading or physical-key
launch. Task2's observed direct Go menu is separate evidence and does not validate
Home Manager deployment. Final combined coordinator runtime is still pending.

Historical two-link/bridge instructions are preserved at
`cd5cbce818e07c264738ddcbc591ad49fdf0e9c8:tests/packaging/home-manager/README.md`,
not active setup requirements.

References: [module](../../../nix/home-manager.nix),
[flake](../../../flake.nix), [Rex config](https://www.superlogical.com/rex/docs/customize/config),
[launch contract](../../../docs/launch-contract.md).
