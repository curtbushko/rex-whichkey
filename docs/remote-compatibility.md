# Execution-host and remote-context limits

## What runs where

Shared Lua and `rex-whichkey`/compatible licensed `rex` belong on the
**execution/server host**, not merely the graphical viewer. Lua creates one
native floating terminal and starts fixed `rex-whichkey direct` there; Go owns
menu data, providers, palette, own-PTY sizing and CLI effects. No Go launching
parent, renderer child, copied image, private transport or remote deployment,
attach, SSH, Wake-on-LAN or connection management is implemented.

Installed require binding and explicit script action use the same shared file.
Source `make run` builds the current checkout, invokes that file with explicit
origin `-s/-b` and supplies its absolute target-bin directory for a terminal-scoped
PATH prepend. It does not require global registration, installed stale modules,
init.lua edits or reload. Installed launch omits the development directory.
See [source commands](../README.md#run-this-checkout).

A matching remote pathname/network endpoint alone does not establish that host's
program/environment/origin authority. Use the actual originating execution-host
context. Cross-host public launch is not certified.

## CLI, endpoint and environment

Go selects `rex` once via ordinary execution-host PATH. There is no `REX_BIN_DIR`,
private executable override, Nix runtime lookup, anchor, wrapper, hash/status/
version/schema preflight or alternate binary retry. The installed server's PATH
must already contain both program names; package/profile changes do not update
an existing process environment. Source launch changes only its new terminal's
PATH for the known current-checkout `rex-whichkey` name.

Go endpoint policy is present direct `--server` > nonempty rendering-host
`REX_SERVER` > bounded no-autostart native Linux/macOS `whoami`. The shared Lua
action has no server data option and never stringifies `ctx.server`; source
`rex do` and the created terminal retain their independent execution-host endpoint
policies. Do not treat a viewer socket or caller-only export as terminal authority.
Concrete Unix addresses/numeric-port host:port are supported, not mutable Go host
labels. No socket guess, autostart, reconnect retargeting or uncertain retry occurs.
See [runtime discovery](runtime-bootstrap.md).

The pinned package declares Linux/macOS outputs, with official CLI artifacts for
x86_64/aarch64 Linux and aarch64 macOS, **not Intel macOS**. A legitimate `rex-cli`
override or normal PATH CLI can supply the dependency. Pins/licensing/loader/assets
are unchanged. Static parse or a declared output does not certify installation,
platform runtime, license permission or server environment.

## Original resource and GUI intent

Source run requires originating full `REX_SESSION/REX_BLOCK`; absence fails rather
than allowing `rex do` to pick the focused block. A present original `REX_CLIENT`
is forwarded with `-C`, not converted from a CLI caller to a GUI ID. The Lua
context supplies exact original session/block/client intent before creation;
Lua resolves original Window/Layer placement. Go receives all four original IDs
separately from popup `REX_*` identity and validates its distinct sole-block float.
Moved/missing/ambiguous placement fails, never retargets.

Go never uses popup `REX_CLIENT` as original intent. Explicit original client is
strict and must be a validated graphical app; a CLI connection is not one.
Without it, app actions may use only the original pane's validated live attached
app owner/connection generation. Display/resource rename/close need no implicit
GUI app; explicit invalid intent still rejects. No current/first/sole-app fallback
or chooser exists. Rex authentication and Remote Control remain required.
See [CLI/GUI delegation](cli-gui-delegation.md).

## Lifetime, cleanup and effects

The direct menu has no idle expiry or original-Lua action lifetime. Go watches
captured placement/liveness, joins metadata jobs and restores tty modes before
cleanup. Tab changes alone do not dismiss. Lua fixes fractional outer bounds;
Go fits only its own PTY/SIGWINCH, not the origin's GUI pixel geometry.
Native chrome uses Rex defaults for now; user1a deferred customization/REQ-026.

Exact sole-popup detach/visible retirement, detached identity and exact original
focus/placement/app generation must be verified before effects. Uncertain
pre-effect cleanup suppresses dispatch. Under accepted user1b, the detached host
Block may remain through the once-only effect until Go exits/Rex deletes it.
On-completion/zero-quick-failure exit policy is not proof of actual host deletion;
Go cannot observe its own post-exit cleanup. A later failure cannot suppress an
already-issued effect. Coordinator runtime evidence must inspect exact
`session.list_blocks` after exit. See [cleanup](launch-contract.md#fixed-owned-cleanup-inventory).

Targets/actions remain Go-retained; keys/text/labels do not replace them.
Destinations/placement/generations are rechecked. New Tab creates one exact
source-session Window and activates its returned ID; failed activation preserves
it without rollback/duplicate creation. Unknown/uncertain creation or cleanup may
leave resources. Never compensate by label cleanup, closing unrelated windows or
recursive scavenging. Checks and mutations are separate, not atomic.
SIGKILL/server loss is not guaranteed recoverable.

## Data, migration and evidence

Go reads strict version-1 JSON/defaults; no executable Lua menu data, callbacks,
functions/custom actions or hooks are supported. User files/init.lua/bindings stay
user-owned. Migration is manual and must preserve unsupported data, not silently
drop it. No conversion, reload/restart, activation or deployment is performed.
See [examples](../examples/README.md).

Only an enabled `known-hosts` provider asks Rex for execution-host/home inventory,
once per invocation. The viewer and `--server` do not select that inventory.
Rex may synchronize/contact its home; listing is not a no-network/reachability
guarantee. Rows are informational, not connect/SSH commands; no default hosts row.

Phase19 tasks1/2 have revision-bound direct-source Darwin observations; final
Task03 make-route combined/rebuilt-main runs remain coordinator gates. No remote,
Linux, installed-profile, physical-key or universal failure-branch coverage is
claimed. [Acceptance](acceptance.md) preserves earlier evidence without certifying
the current package/runtime. Phase15 bridge/anchor/completion and Phase16/18
controller mechanics are historical, not active dependencies.
