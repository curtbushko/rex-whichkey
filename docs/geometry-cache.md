# Cached startup height

Accepted decision: **REXWHICHKEY-20261126-103524** in the
[decision register](decisions.md). The user requested reuse of the calculated
height to reduce the visible maximum-height-to-content-height startup jump.

## Location and format

The execution account writes one disposable JSON hint:

- `$XDG_CACHE_HOME/rex-whichkey/geometry.json` when XDG_CACHE_HOME is nonempty.
- Otherwise `$HOME/.cache/rex-whichkey/geometry.json`, including on macOS.

Relative cache roots or an unavailable absolute HOME disable caching rather than
writing under the working directory. This is not menu configuration or a required
installation asset. Deleting the cache returns to cold-start measurement.

Version 1 stores only `version`, `width`, `maximum_height`, `height`, `config`
(explicit launch path or the empty string), and `ascii`. There are no resource IDs,
server/client selectors, executable paths, actions or cleanup authority. It is
JSON data, never executable Lua. A single last-writer-wins entry intentionally
avoids a database, lock service or unbounded cache growth.

## Startup and fitting

The Lua launcher reads at most 4097 bytes from the local account's regular cache
file and rejects data larger than 4096 bytes. Missing/unreadable JSON, decode
failures, unsupported versions, mismatched width/maximum/explicit-config/ASCII,
and nonpositive, nonnumeric or over-cap heights all use the normal configured
maximum. A valid hit changes only initial height and vertical centering; width,
original invocation context and all ownership checks remain unchanged.

Lua always passes the configured `--max-height` separately to Go. Cached compact
bounds therefore do not become the submenu growth cap. Legacy launchers omitting
this flag retain their launch-bounds cap and do not populate this cache. Update
the Lua asset and Go binary together to enable caching; installed bindings already
loaded into a Rex process are not updated by a source build alone.

Go caches only a settled, presented **root menu**, not the last submenu,
confirmation or rename form. Its own terminal grid must match an actual sizing
owner, and pending/failed resizes cannot seed the cache. A clipped root at the
configured maximum may cache that maximum. Unchanged measurements are written
once per invocation. Returning from a submenu can refresh the root hint.

Files are written via a same-directory 0600 temporary file and atomic rename;
new cache directories use 0700. The final file's symlink target is not overwritten.
Concurrent invocations may replace each other's hints, but cannot partially
interleave JSON. Write failures and cache-acquisition failures are best effort:
they never suppress menu actions or change cleanup behavior.

## Limits and trust boundary

This is a startup hint, not a first-size-final promise. The first presented run
still measures normally. Screen, font, window size, implicit config environment
or menu-content changes can require adjustment and refresh; explicit launch
width/maximum/config-path/ASCII changes cause a cache miss. There is no screen-ID
registry or content hashing. Lua and Go must use the same account/cache filesystem
for a warm hit; a CLI script running elsewhere may miss. Cross-host behavior is
not certified.

The user-owned cache directory is an ordinary local-filesystem trust boundary,
not a sandbox for hostile filesystem objects. The reader bounds JSON bytes and
validates hint values, while live Rex snapshots retain all mutation authority.
No cached resource identity, app routing, external executable or arbitrary code
is trusted, and cache failures do not authorize retries of uncertain mutations.

## Verification

Tests cover paths, permissions, atomic replacement/temp cleanup, final-symlink
behavior, Go-written hints consumed by the Lua launcher, malformed/oversized and
mismatched hints, root-only settling/presentation, best-effort write failures and
preservation of submenu growth beyond cached startup height. Native Rex Lua
lookup was exercised with isolated synthetic cache data and mocked resource
creation, not an actual warm-start GUI certification. See [acceptance](acceptance.md)
for build/runtime evidence and its limits.
