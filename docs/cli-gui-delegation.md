# CLI callers and GUI targets

## Active CLI-only boundary

An execution-host process identity and a graphical app identity are distinct.
The accepted caller/app separation correction
`REXWHICHKEY-20261104-172117` remains relevant to routing; its old bridge mechanics
are historical. Current CLI-only direction is `REXWHICHKEY-20261106-114935`.

The public binary requires full origin `REX_SESSION/REX_BLOCK`. Present
`--client` overrides present `REX_CLIENT`; either must be an exact full
`client:UUID`, never empty, a label or a transient CLI chosen as an app.
The minimal keybinding passes exact `ctx.client_id` as intent and clears
REX_CLIENT when absent. Go independently verifies it. There is no
`rex --client APP do whichkey` custom-action route or caller impersonation.

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

The same Go parent owns both direct and keybinding startup, retained menu/action
snapshot, popup lifecycle and exact cleanup before effects.
Lua neither queries/captures connection generations nor owns a registry, completion
action, ticket or callback VM. Its detached start handle is not display evidence.
There is no menu idle expiry, even though individual operations/cleanup are bounded.

Moved origins, reconnects, missing permissions, uncertain cleanup/effects or
changed CLI/server identity fail safely. Inspection and mutation are separate APIs,
not atomic. No retry, label/focus substitute, alternate server or rollback repairs
an uncertain effect. [Launch cleanup](launch-contract.md#fixed-owned-cleanup-inventory)
and [remote limits](remote-compatibility.md) apply.

## Evidence and historical boundary

Task2's coordinator observed public default menu/navigation/dismissal/cleanup,
not every selected app action, CLI-vs-GUI case or physical key.
Final combined candidate observation remains pending; no config reload/deploy is
needed or authorized by this source guide.

The former original-VM/preflight/receipt/completion contract is preserved in Git at
`cd5cbce818e07c264738ddcbc591ad49fdf0e9c8:docs/cli-gui-delegation.md`.
It is not an active prerequisite. Functions/hooks/callbacks remain deferred.
See [acceptance](acceptance.md) and [data/action catalogue](mapping-format.md).
