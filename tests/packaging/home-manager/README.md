# Home Manager integration: source-review cases

These are configuration examples and a source-review checklist, not an evaluation
or activation runner. No tests, Nix evaluation/build, Rex invocation or reload were
performed for this module.

## Import and opt in on the server host

Add the flake as an input to your own configuration, then import its module in the
**Home Manager configuration for the user running the Rex server**:

```nix
# Within your own Home Manager module, with inputs supplied by your configuration:
{ inputs, ... }: {
  imports = [ inputs.rex-whichkey.homeManagerModules.default ];
  programs.rex-whichkey.enable = true;
}
```

`homeManagerModules.rex-whichkey` is the same module. Neither importing the module
nor installing the package alone registers an action. The flake module selects
`packages.<host-system>.rex-whichkey`; `programs.rex-whichkey.package` can override
it with a compatible package containing the same executable/Lua layout. Directly
importing `nix/home-manager.nix` instead uses `pkgs.callPackage ./package.nix {}`.
The flake export is preferred for the project's pinned Go package selection.

When enabled, the module adds the executable to `home.packages` and links exactly
these two integration paths beside the server's `init.lua`:

- `rex_whichkey.lua` → the package's Lua entry point
- `rex_whichkey` → the package's Lua module directory (one directory symlink)

The default destination is `${config.xdg.configHome}/rex`, matching Rex's
`$XDG_CONFIG_HOME/rex/init.lua` / `~/.config/rex/init.lua` location on Linux and
macOS. For a server using a custom `REX_CONFIG`, select the directory containing
that file, inside the Home Manager user's home:

```nix
{ config, ... }: {
  programs.rex-whichkey = {
    enable = true;
    configDirectory = "${config.home.homeDirectory}/rex-server-config";
  };
}
```

This option does not set `REX_CONFIG`, create an init file, change the Lua search
path or configure a client machine. Rex's server runtime loads the installed
modules through its documented sibling-module `require` lookup.

## Explicit wiring in the user-owned init.lua

You choose whether and when to add wiring to **your own** server configuration.
The module generates none of this Lua and never reads, writes or imports your
existing mapping file into the Nix store. The following consumes the delivered
`lua/rex_whichkey.lua` and `lua/rex_whichkey/defaults/init.lua` interfaces:

```lua
local popup = require("rex_whichkey")
popup.register_completion("whichkey_complete")

-- Replace this example with this server's fixed, reachable socket endpoint.
-- Never use ctx.server, a client-local socket or an inferred current server.
local menu = popup.new(
    require("rex_whichkey.defaults"),
    "unix:///absolute/path/to/server.sock"
)

-- Optional sibling my_whichkey.lua remains user-owned, with its original bytes.
-- Use menu.setup({}) instead for built-ins only, or put inline options here.
menu.setup({}, "my_whichkey")

rex.action {
    name = "whichkey",
    title = "Which key",
    run = function(ctx)
        menu.launch(ctx)
    end,
}
```

Register completion once in the same server Lua VM as the launch action. Keep the
completion name unique among your actions, and bind the `whichkey` action through
your existing Rex keybinding configuration if desired. The installed namespace
`rex_whichkey` is reserved; user mapping modules must use a different name, such
as `my_whichkey`. Inline/import precedence remains **defaults < imported < inline**;
`spec` and `icons` replace whole tables, as defined by `docs/mapping-format.md`.

Both `rex-whichkey` and the existing `rex` CLI must be on the **server process's**
PATH. An already-running server need not inherit a newly installed user profile;
this module does not restart it or promise that its environment has changed.
Nothing is installed on a remote server by configuring only the local client.

## Disabled and preservation review

With the import retained, this is the default (explicitly shown):

```nix
{
  programs.rex-whichkey.enable = false;
}
```

Read `nix/home-manager.nix` and the small flake export without executing them:

- **Disabled:** the entire `home.packages` and `home.file` contribution is behind
  `lib.mkIf cfg.enable`. There are no files, hooks, services or environment changes
  contributed by the disabled module. Option defaults do not activate anything.
- **Enabled:** only the executable package and the two reserved module links are
  contributed. No `init.lua`, user mapping file, whole Rex directory, keybinding,
  action registration, `REX_CONFIG` or `LUA_PATH` setting is managed.
- Both links explicitly retain `force = false`; the directory link also retains
  `recursive = false`. Do not set force or delete/move conflicting files to make
  activation succeed. Existing unrelated sibling files are not module targets.
- If either reserved path already exists unmanaged, normal Home Manager conflict
  handling applies; the module contains no bypass or overwrite command. If your
  configuration already links the **entire** Rex directory, account for these
  sibling paths in your own declarative layout rather than replacing init.lua.
- There is no custom activation, `onChange`, reload/restart, service or launchd
  unit. Enabling does not launch the menu or mutate running server state.
- Removing/disabling the module in a later Home Manager generation removes its
  managed links/package contribution, not your init or mapping files. User-added
  require/setup wiring is still yours and must be adjusted before using Rex with
  the integration absent; disabling does not edit it for you.

This establishes source intent only. Disabled/enabled evaluation, conflict
handling, existing-file byte preservation on activation, server PATH inheritance,
Lua search compatibility and end-to-end invocation remain runtime-unverified.
Independent review is required; no evaluation or runtime success is claimed here.

## References

- [Rex configuration location and sibling-module loading](https://www.superlogical.com/rex/docs/customize/config)
- [Go-owned mapping and options standard](../../../docs/mapping-format.md)
- [Lua loader configuration semantics](../../../examples/README.md)
- [Packaged layout](../../../nix/package.nix)
