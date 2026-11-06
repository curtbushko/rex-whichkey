# Historical Phase15 Lua source fixtures

The Lua files here are byte-for-byte copies of the production Lua tree at
cd5cbce818e07c264738ddcbc591ad49fdf0e9c8, before the CLI-only launcher reduction.

Only the package.path prefixes in tests/lua/action_context.lua,
tests/lua/cli_delegation.lua and tests/lua/configured_rename.lua were redirected
here. Existing assertions and behavior were not changed. These historical scripts
are not current public-launch coverage and were not executed for this cutover.

This tree is excluded from package assets and never required by the production
launcher. Its registrations/callbacks/anchor are historical, not supported hooks
or migration instructions. Do not edit fixtures to make current behavior pass.
