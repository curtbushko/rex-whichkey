# CLI callers and GUI targets

Decision: **REXWHICHKEY-20261104-172117 — Accepted**. The user explicitly approved
separating CLI activation identity from GUI routing for launch and completion
("yes, please fix"). This refines the previous requirement that every fresh
activation's `ctx.client_id` equal the originating app ID; it does not weaken the
remaining capability, origin, connection, configuration or lifetime checks.
The authoritative decision index is
[the decision register](../.tasks/designs/rex-whichkey-decision-register.md).

## Why the previous contract failed

`rex --client APP do whichkey` does not impersonate APP's connection. The server
can supply the CLI connection as `ctx.client_id`. A later completion command has
another fresh CLI connection. Requiring those IDs to equal APP prevented public
launch and would also reject completion from the controller.

## Current design

- Preserve the real action context. Freeze its caller ID during each activation
  and recheck it after yielding operations. Never rewrite it to the app ID.
- Public Go launch resolves the app from the originating pane's GUI owner or an
  explicit full client selector. Lua validates the nominated target as a live
  graphical app attached to the exact originating session, with the captured
  `session.list_clients.connected_at` value. Key/palette launch uses its actual
  app caller. A CLI action without a GUI binding fails; there is no sole-app or
  focused-pane fallback.
- Retain GUI identity/generation separately from the launcher caller. The
  controller handoff still carries the GUI target, not the ephemeral CLI ID.
- A fresh completion must present the private capability in the original server
  Lua VM and match the exact server/session/block. Its actual caller must be a
  full client ID and remain unchanged throughout that activation, but need not
  equal the launcher or GUI client. Completion data cannot nominate another app.
- Claim before any yielding completion work. Revalidate target attachment,
  graphical kind, connection timestamp, origin placement, mapping/confirmation,
  guards, configuration generation and original expiry. Errors consume claimed
  invocations. Unknown capabilities and wrong origins cannot consume another
  invocation. Cancellation, expiry, replacement and replay remain fail-closed.
- Focus the retained origin, then dispatch dotted client actions with documented
  `client.perform`, explicitly supplying app/session/block and retained args.
  Check `performed_by`; refused, failed, timed-out or uncertain requests are not
  retried or redirected. Configuration action names still use `rex.invoke` in
  the original VM. No mutable global `rex` override changes ambient authority.
- Callbacks receive the logical origin context: `client_id` is the retained GUI
  target and `request_client_id` is the actual completion caller. This table is
  routing information, not a fabricated Rex activation. Callbacks needing explicit
  GUI effects should use `rex.call("client.perform", ...)` with that target.
  Nested configuration actions keep Rex's real caller context; CLI `--client`
  remains the selector for any client actions returned by those actions.

## Security and limitations

The capability authorizes one retained invocation within Rex's authenticated
execution boundary; possession is not proof of a physical GUI keypress. App IDs
are routing selectors, not authentication credentials. Remote Control remains
Rex's responsibility. Selecting an app never bypasses its permission checks.

The existing connection binding uses Rex's session attachment timestamp and lost
state. It cannot detect a reconnect that Rex does not expose as a changed/lost
attachment. Focus, validation and execution are not atomic, so an external change
can still race the final check. No automatic retry resolves that uncertainty.

No extra daemon, client impersonation, callback serialization or alternate-server
fallback is introduced. Older Lua modules paired with the updated Go launcher
produce fixed reinstall guidance when their legacy client-ID mismatch is detected;
raw child diagnostics and capability data remain suppressed.
