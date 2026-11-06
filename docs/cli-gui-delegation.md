# CLI callers and GUI targets

Decision: **REXWHICHKEY-20261104-172117 — Accepted**. Separating CLI activation
identity from GUI routing refines the old caller/app equality rule, not capability,
origin, connection, configuration or lifetime checks. See the
[decision register](../.tasks/designs/rex-whichkey-decision-register.md).

## Actual caller is not a nominated app

`rex --client APP do whichkey` selects APP for routing; it does not impersonate
its connection. The launch action and later completion can each have a fresh CLI
`ctx.client_id`. Neither must equal APP or the earlier CLI caller. Freeze the
real opaque `ctx.server`, session/block and actual caller for each activation;
never rewrite that context to match nomination data. Key/palette launch has no
public nomination: its actual app supplies both caller and GUI target.

Public Go selects an explicit full `--client`/present `REX_CLIENT`, or the
originating pane's documented GUI owner. Lua receives that exact app ID and
`connected_at` separately from the public runtime locator. It constructs the
complete origin from actual context and direct session APIs, checking unique
window/layer/block placement and the nominated app's exact attachment, `app`
kind, nonlost state and connection generation. No sole/first/focused app fallback,
label lookup or new app after failure is permitted. Recheck after every yield.

The private Go preflight authenticates local state, live lease, receipt, consumer
and that complete origin **before** returning compatibility data, reserving a
ticket or evaluating user mappings. Public nomination alone is not authority.
Constructor endpoint intent is immutable and must agree with the already frozen
public endpoint; it cannot retarget transport or replace the actual caller.

## Original VM and single-use completion

- Read-only `whichkey_runtime` advertises the loaded menu/config generation and
  features before popup creation. Go protocol/layout 1 and negotiated completion
  2 are internal formats, not Rex APIs. Build IDs need not match. Missing/old
  bridges reject before activating an old menu, never fall back or hot-reload.
- Closures, original args, guards, confirmation and immutable mapping/resource
  snapshots stay with table tickets in the original configuration VM. Go supplies
  a random capability; no callback is serialized or evaluated in another VM.
- Go-owned receipt and relay identities remain opaque byte-exact strings across
  Lua 5.1. Their timestamps/device/inode integers may exceed 2^53; parsing and
  reserializing them would change authority. Completion binds the exact final
  receipt hash, not a reconstructed Lua object. See the authoritative
  [mapping/transport standard](mapping-format.md), not a second schema here.
- Fresh completion strictly decodes and authenticates capability, exact receipt,
  server/session/block, retained state and original expiry before claiming.
  Claim occurs **before** yielding guards/APIs/callbacks. Wrong/malformed/foreign
  input cannot consume an unrelated entry; a claimed failure is consumed forever.
- Revalidate exact GUI connection, placement, menu/config generations, retained
  keys/ref/confirmation/text and guards. Focus only a still-valid captured origin.
  Dotted client actions use explicit `client.perform` with retained app/context/
  args; configuration actions use `rex.invoke` in the original VM. No ambient
  queue on a transient CLI replaces this targeting.
- Callbacks receive logical routing context: `client_id` is the retained GUI,
  `request_client_id` the actual completion caller. That table does not fabricate
  a Rex activation. Explicit callback GUI effects should target the retained app.
  Custom args/conditions/confirmation and rename-key callbacks remain unchanged.

## Startup, cleanup and limitations

Public Go prepares and starts its own selected component; the installed helper is
only a finite file relay. Key launch uses its normal anchored helper. Both use
finite private handoffs and Go's accepted/EOF, then positive proceed/EOF startup
barriers. Acceptance is not observed popup display or callback completion. Once
start may have succeeded, no failed/lost acknowledgment authorizes retry, killing
the uncertain child or deleting its workspace. Malformed/unauthenticated handoffs
may also retain artifacts; expiry revokes authority without proving file removal.

Only exact owned popup/files/handles are cleaned, before selected dispatch;
uncertain close/focus/cleanup suppresses effects. See the source-linked
[fixed owned inventory](launch-contract.md#fixed-owned-cleanup-inventory) for
popup and rendezvous/preparation files and the original-Lua release boundary.
Default tab/session rename uses validated text and retained resource APIs after
cleanup, never a native dialog,
discarded name or success inferred from a generic action acknowledgment. The
original 30-second configured lifetime is never renewed (64 pending entries,
five-second teardown reserve); native interaction retains its lifetime pipe and
no idle deadline. See [runtime bootstrap](runtime-bootstrap.md) and
[launch contract](launch-contract.md).

App IDs select routing, not authentication or proof of a physical keypress.
Rex transport authentication and app Remote Control remain in force. Read-only
checks/focus/effects are not atomic; external moves, setup and disconnects can
race the last check. No retry repairs that uncertainty. Execution-host files are
required; a remote viewing app or endpoint does not supply them.

This describes source-reviewed/compiled intent, not observed runtime verification.
Caller/app separation, compatible unequal builds, endpoint disagreement, custom
callbacks, rename, reconnect/expiry, lost acknowledgment and cleanup-failure cases
are **UNEXECUTED source scenarios** here. No live CLI/API/menu probes, tests,
reloads, configuration changes or deployment were performed for this cutover.
