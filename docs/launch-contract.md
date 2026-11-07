# Shared Lua float and direct Go launch contract

## Public launch and source routing

[`lua/rex_whichkey.lua`](../lua/rex_whichkey.lua) is the single implementation for:

- Installed `require("rex_whichkey").launch(ctx, options)` in an optional
  function binding.
- `rex do /absolute/path/rex_whichkey.lua --action rex_whichkey_launch` with
  explicit original `-s SESSION -b BLOCK` and optional original `-C CLIENT`.
- Normal source `make run` / `task run`, which builds this checkout first and
  supplies its known absolute `bin/GOOS-GOARCH` as `development_directory`.

The script registers its action in this invocation; no installed global action,
user init.lua registration/reload or stale `require` lookup is needed for source
run. Only canonical require receives the module function table; ordinary script
execution returns no function value to serialize. There is no debug-source or
Go-environment discovery in Lua.

`make` on the current development PATH is a tiny exec-Task wrapper. Use
`make run -- width=0.8 height=0.3 config=/absolute/menu.json ascii=true no_color=true`.
Conventional GNU make reads [Makefile](../Makefile) and uses
`make run ARGS='width=0.8 height=0.3 config=/absolute/menu.json'` instead.
Both delegate the unchanged quality commands to [Taskfile.yml](../Taskfile.yml).
Run from the source checkout; `task build` retains the ordinary Go build command,
outputting the fixed `rex-whichkey` program under the selected GOOS-GOARCH.

The source task fails if originating `REX_SESSION` or `REX_BLOCK` is missing or
empty, before any launch; it always passes both `-s` and `-b` rather than allowing
Rex's focused-block default. A present `REX_CLIENT` is passed explicitly as `-C`;
empty intent fails, and Go validates full ID/app eligibility. Do not replace a
CLI caller's identity with a focused/sole GUI app. An absent original client
remains absent at the source route; it is not taken from popup environment.
Endpoint/CLI selection is independent; see [runtime discovery](runtime-bootstrap.md).

## Lua presentation and creation

`launch(ctx, options)` requires full original session/block IDs and validates an
optional original full client ID. It obtains the authoritative session view and
requires unique exact Block → Layer → Window placement, with no focus fallback.
Options are a closed data table:

| Option | Contract |
| --- | --- |
| `width`, `height` | Finite fractions `(0,1]`, defaults `.8`, `.3` |
| `config` | Optional nonempty control-free JSON path; prefer absolute |
| `ascii`, `no_color` | Optional booleans, forwarded as Go presentation flags |
| `development_directory` | Trusted source/build caller only: one absolute PATH directory, no colon; omit installed |

Lua centers bounds at `{(1-width)/2, (1-height)/2, width, height}` and creates
exactly one terminal in one new floating layer in the original window. It uses
`rex.layout.block`, `shell="none"`, and the fixed argument-array program:

```text
rex-whichkey direct --origin-session SESSION --origin-block BLOCK \
  --origin-window WINDOW --origin-layer LAYER [--client ORIGINAL_CLIENT] \
  [--config PATH] [--ascii] [--no-color]
```

Development prepends only the supplied directory to this terminal's PATH.
Lua does no executable nomination, source probes, copy or supervisor work.
Installed normal launch omits the directory and uses the server's ordinary PATH.
Native titlebar/border/corners/shadow use Rex defaults. User1a explicitly deferred
REQ-026 customization; terminal palette/text borders are not its implementation.

Creation is never retried on uncertainty. Returned full layer/block IDs describe
creation acknowledgment, **not** a drawn menu, successful action, Go exit or
cleanup. The binding returns `true` (plus the IDs); the script action returns
serializable IDs. `rex do` can exit0 while the Go menu is still running or later
fails. No menu idle timeout or Lua lifetime supervisor is introduced.

## Direct Go identity and startup

[`cmd/main.go`](../cmd/main.go) → [`RunDirectMenu`](../direct_menu.go) runs inside
the created terminal. Bare invocation/root launcher flags/legacy popup mode are
rejected; they cannot detach an ordinary caller pane. `menu` is UI/result-only,
not a popup/resource launcher.

Original IDs are explicit arguments captured **before creation**. The terminal's
`REX_SESSION/REX_BLOCK` identify only its own execution context. Go requires the
same session, a distinct own Block, and exact sole-block floating-layer placement
in the captured original window. Raw own-float validation is separate from the
filtered provider view; own popup is excluded from pane rows/counts/targets and
current markers describe the original source.

Go resolves `rex` once via execution-host PATH and independently retains a concrete
endpoint. No binary receipt/hash/status/version/schema preflight, Nix lookup or
alternate binary/server route runs. It captures original app eligibility and
connection generation, loads strict version-1 JSON/defaults and runs one terminal
UI owner. The static root is written before provider demands or label waits.
Go reads only its **own PTY** for cells/palette/SIGWINCH; it does not calculate
native bounds from origin geometry.

Explicit `--client` is strict original intent; popup `REX_CLIENT` is never read as
original intent. Without it, app operations may retain the original pane's exact
validated live graphical owner/generation. A CLI ID is not a GUI app. Display,
resource rename/close and pane focus need no implicit eligible app; explicit
invalid intent still rejects. No focused/first/sole-app selection or chooser is
used. [CLI/GUI delegation](cli-gui-delegation.md) describes this boundary.

## In-process menu metadata (STARTUP-UI-UPDATES)

This replaces the old private incremental rename/provider file transport. Internal
bounded request/publication values remain Go data, not IPC or a public schema:

- At most two metadata jobs and two result slots; one attempt/result per kind,
  including errors. Visible enabled sessions demand after first successful draw;
  hidden sessions and windows/panes/known-hosts demand on explicit prefix entry.
  Rename labels demand on form entry. Disabled/absent providers do no work.
- Window/pane/rename presentation reuses the validated matching origin view;
  sessions-list and hosts have independent once-per-invocation results. No cache
  substitutes for fresh pre-effect resource/app/destination checks.
- Workers return detached immutable results. The UI owner retains complete
  mappings/args/full targets before ready rows and applies complete batches without
  changing prefix/page/confirmation/edits/cursor/partial decoder state.
- Loading/error rows cannot execute. Rename is editable/cancellable while loading;
  Enter is discarded until ready, never queued/replayed. Ready labels initialize
  only untouched fields. Selections match retained ready publications; labels,
  indices and text do not grant replacement target authority.
- Existing 512-resource, child/key-byte/4-MiB, UTF-8/text and confirmation bounds
  remain. Re-entering does not retry/refresh. All workers cancel/join on exit.

See [data/providers](mapping-format.md), [rename](tab-rename-input.md) and
[session UI](session-owned-ui.md). There is no renderer child, runtime image,
lifetime FIFO, lease, result/update/demand files, callback registry or relay.

## Fixed navigation bounds and liveness

Lua chooses one fixed fractional rectangle. Menu navigation, background readiness,
paging/back/focus/idle do not resize or recreate it. Go fits/paginates inside its
current own-PTY cells; SIGWINCH/focus-in redraw do not alter outer bounds. This is
not content-height native shrink wrapping or proof of first-size-final GUI pixels.

The terminal has no idle expiry. Original resource placement/liveness, selected
app generation and operation deadlines remain guarded. Tab switching alone does
not cancel the menu. No original Lua VM, callback-ticket timeout or parent pipe
determines menu lifetime.

## Fixed owned cleanup inventory

After joining jobs and restoring/closing its tty, Go must complete bounded,
cancellation-independent cleanup before any selected effect:

1. Revalidate the exact returned sole terminal/floating layer and placed identity.
2. Call `session.detach_block` once on that exact Block, never close a user
   window/session or select by label. Verify layer retirement/no placed popup
   and exact detached terminal identity through `session.list_blocks`.
3. Revalidate original placement and captured app generation as applicable;
   restore original focus and verify exact window/focused Block. No replacement
   resource or current-focus fallback is permitted.
4. Match the selected retained mapping/publication, then perform fresh required
   target/app checks and issue at most one selected effect. Uncertain detach,
   tty/worker/focus/ownership cleanup suppresses dispatch; no mutation retry.

User1b explicitly permits the detached execution-host Block to remain through
that effect until Go exits or Rex deletes it. Lua sets
`exit={on_completion=true, quick_exit_threshold_ms=0}` so Rex can delete even a
failed/short-lived or detached terminal after exit. **Go cannot verify its own
post-exit deletion.** A later cleanup failure cannot suppress/undo an earlier
action; creation/visible retirement alone does not prove host deletion. Coordinator
runtime evidence must inspect exact `session.list_blocks` after Go exits.

No invocation-private files/image/source handle or child wait remain to clean.
Pre-effect checks and mutations are separate, not atomic against external changes.
SIGKILL, server/terminal loss or uncertain acknowledgments may leave resources.
No broad scavenging, rollback, uncertain retry or guaranteed post-crash cleanup
is claimed.

## Effects and evidence boundary

Go retains original action/resource authority. Session/window selection checks
exact destination membership and uses `client.perform session.select` without
outer origin overrides; splits omit them too. Other app actions retain original
context. New Tab creates one source-session Window and activates its returned
ID; failed activation preserves it, not rollback/duplicate creation. Rename/close
use exact resource methods after pre-effect cleanup. Authentication and Remote
Control remain Rex policy.

Current native Task01/02 summaries and `.mtn-tmp/p19/resume/task{1,2}-public-evidence.json`
in the coordinator checkout bind delivered direct-source observations. Actual
Task03 `make run` on reviewed combined source **and rebuilt main** remains a
coordinator gate; no leaf runtime was performed. It must establish actual
program/popup identity, no Go launching parent/copy/IPC, menu/provider/palette/
own-PTY behavior, owned rename/cancel/dismissal and pre-effect visible/post-exit
host cleanup. Compile/help/creation ACK cannot replace these checks. Task04
refreshes final current evidence. Installed-profile, physical-key, remote and
platform-wide behavior are not certified. [Acceptance](acceptance.md) preserves
older revision-bound evidence without treating old controller mechanics as current.
