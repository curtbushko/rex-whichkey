# CLI callers and GUI targets

## Active shared-Lua/direct-Go boundary

An execution-host process identity and a graphical app identity are distinct.
The accepted caller/app separation correction
`REXWHICHKEY-20261104-172117` remains relevant to routing; its old bridge mechanics
are historical. Current CLI-only direction is `REXWHICHKEY-20261106-114935`.

The shared Lua launcher requires full original session/block context and resolves
exact Window/Layer placement before creation. Source `make run` requires originating
`REX_SESSION/REX_BLOCK` and passes explicit `-s/-b`; a present original `REX_CLIENT`
is passed with `-C`. Missing origin fails rather than focused-block selection.
The binding uses exact `ctx.client_id` intent when supplied. A CLI caller is not
converted to a focused/sole app; full ID shape alone does not prove app eligibility.
Go independently verifies original intent and connection generation.

Lua passes original IDs as `direct --origin-*` arguments and optional `--client`.
Popup `REX_SESSION/REX_BLOCK` identify only the distinct execution terminal;
popup `REX_CLIENT` is **never** read as original app intent. Bare/root-flag Go
launch is rejected. The shared script-local `rex_whichkey_launch` action needs no
global user registration/reload; it is not caller impersonation or a menu callback.

Without explicit app intent, Go inspects the original pane's documented owner
and validates that exact live attached graphical `app` connection and
`connected_at`. No sole/first/focused-app fallback, remembered choice, app
chooser or list-order selection occurs. Pane ownership is routing context, not
proof of which device pressed a key or additional authentication.

## Display versus selected effects

Display, ordinary dismissal and origin focus require valid resource placement,
not an eligible implicit GUI app. Resource rename/close and generated pane focus
also work without implicit app eligibility. A supplied explicit app selector
remains strict, including for these operations.

For app actions, Go retains the binding before interaction and rechecks exact
connection generation, origin Window/Layer placement and (for implicit routing)
unchanged pane owner. Missing or changed identity refuses, never retargets.
Built-in action metadata must have `kind=client`; dispatch uses one explicit
`client.perform` through the same frozen CLI and endpoint. There is no
custom Lua fallback. Rex authentication and app Remote Control still apply.

Session/window selection validates exact destination membership and omits outer
origin session/block overrides; splits omit them too. Other app actions carry
captured origin context. New Tab creates one typed source-session Window then
activates that returned ID on the same bound app. Creation can change pane
ownership, so activation rechecks retained connection/placement, not a replacement
owner. Failed activation preserves the created tab for manual inspection.

## Lifecycle and uncertainty

The same shared Lua file creates one float for source/script/keybinding use;
one direct Go process owns menu/action state. No Go launching parent, renderer
child, copied image, private transport or supervisor remains. Lua resolves original
placement but does not capture app connection generations or run a callback VM.
Creation IDs/launch exit0 are not display, Go exit or cleanup evidence.

Go joins workers/restores tty, detaches/verifies the exact visible float and
restores/verifies original focus/placement before effects. User1b permits only the
detached execution-host Block through dispatch until Go exits/Rex deletes it.
Go cannot verify its own post-exit deletion; a later cleanup failure cannot
suppress an earlier effect. Coordinator must inspect exact `session.list_blocks`
after exit. There is no menu idle expiry; operations/cleanup remain bounded.

Moved origins, reconnects, missing permissions, uncertain cleanup/effects or
changed CLI/server identity fail safely. Inspection and mutation are separate APIs,
not atomic. No retry, label/focus substitute, alternate server or rollback repairs
an uncertain effect. [Launch cleanup](launch-contract.md#fixed-owned-cleanup-inventory)
and [remote limits](remote-compatibility.md) apply.

## Evidence and historical boundary

Phase19 Task01/02 native summaries bind actual reviewed direct-source menu,
provider/owned rename/cancel/dismissal and cleanup observations. The new Task03
make-route combined/rebuilt-main observation remains a coordinator gate, not a
leaf result. No every-app-action, all CLI-vs-GUI, physical-key or installed-profile
coverage is implied; no config reload/deploy is needed or authorized here.

The former original-VM/preflight/receipt/completion contract is preserved in Git at
`cd5cbce818e07c264738ddcbc591ad49fdf0e9c8:docs/cli-gui-delegation.md`.
It is not an active prerequisite. Functions/hooks/callbacks remain deferred.
See [acceptance](acceptance.md) and [data/action catalogue](mapping-format.md).
