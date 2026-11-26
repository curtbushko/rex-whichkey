# Decision register

This register tracks new decisions from the current collaboration. Earlier
managed decision/phase records and historical acceptance entries retain their
own identity and evidence; this file does not reconstruct missing history.

| ID | State | Decision | Rationale and relationships |
| --- | --- | --- | --- |
| REXWHICHKEY-20261126-092545 | Accepted | Keep fractional popup width; fit current menu height plus padding, including submenu transitions, with configured height as the maximum and pagination beyond it. | User explicitly accepted after reporting excess whitespace on large screens. Supersedes the fixed-height/navigation-no-resize policy for the current direct popup, not historical verification. See [design](content-sized-popup.md). |
| REXWHICHKEY-20261126-102939 | Superseded | Proposed calculating initial compact bounds before creating the popup to avoid the visible maximum-to-content launch jump. | This proposal was not accepted. The user chose cached calculated height instead; replaced by REXWHICHKEY-20261126-103524. |
| REXWHICHKEY-20261126-103524 | Accepted | Save calculated startup geometry in `$XDG_CACHE_HOME/rex-whichkey/geometry.json`, falling back to `~/.cache/rex-whichkey/geometry.json`, and reuse it as a disposable initial sizing hint. | User explicitly requested implementation after approving the location. Refines REXWHICHKEY-20261126-092545 without changing the configured maximum or submenu fitting. ID assigned at implementation capture, after approval; it is not an exact approval timestamp. See [cache design](geometry-cache.md). |
