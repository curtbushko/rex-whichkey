# Unreviewed historical checkpoint — test fixture only

HISTORICAL ONLY. This is not production code, a reviewed delivery, or evidence
that the old controller/callback design is supported by the current application.

The complete 11-file `lua/` tree was extracted byte-for-byte from preservation
checkpoint `30a4439302ade4e08487d176ba6823206ddc326d`:

```sh
git archive 30a4439302ade4e08487d176ba6823206ddc326d lua |
  tar -x -C tests/fixtures/legacy-context-checkpoint --strip-components=1
git ls-tree -r 30a4439302ade4e08487d176ba6823206ddc326d lua
```

`MANIFEST.tsv` identifies the source commit/tree and historical-only status in
comment headers, followed by that last command's exact output: mode, object type,
Git blob ID and original source path. Strip the `lua/` prefix to find the fixture
file; `git hash-object` of each file must equal its recorded blob ID. The source
commit is the checkpoint above, not the later incompatible `phase15-lua` archive.
That separate archive remains unchanged.

## Matching historical regression contract

These harnesses and their Go driver first appeared in the same checkpoint:

| Original source | Git blob at the checkpoint |
| --- | --- |
| `lua_context_test.go` | `c4d6e2a2dae2f8b0e640f6ee8b1542ce05baac94` |
| `tests/lua/action_context.lua` | `c383422d20074835de30f6ab509810168ed2cd2e` |
| `tests/lua/cli_delegation.lua` | `57cc865411c1c77e36f6d0e926b118fcbdecac50` |
| `tests/lua/configured_rename.lua` | `58392d4f4f9efb1eedfab2407c4eba11205cacd9` |

The Go driver remains byte-identical. Each Lua harness differs from its original
only in its first-line `package.path`, now pointing here. All test statements,
42 scenarios, assertions, side-effect counts and negative/security cases remain
unchanged. No current-production regression has been redirected to old code.

This snapshot supplies the original `execution.helper` / `rex_bin_dir` launch
options, field-specific origin diagnostics, and `registry.issue` / `registry.bind`
contract used by those harnesses. It does not have the later file-source/layout
discovery requirement, so GopherLua needs no source-name override or staged fake
installation. No compatibility objects, fake success replies, guard changes or
production hooks were added.

Current runtime behavior remains owned by the production `lua/` and Go sources;
passing these historical tests establishes only the preserved historical contract.
