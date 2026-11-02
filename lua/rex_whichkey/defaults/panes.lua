-- Ordered spec fragment; mapping.go/docs/mapping-format.md own its format.
-- The dispatcher restores the authenticated origin before queueing these client
-- actions. Importing defaults must never split, focus or close a pane itself.
return {
    -- These are pane.split action directions, not CLI --split values: the
    -- downward action takes "down", whereas the CLI uses "below".
    { "v", "pane.split", desc = "Split right", icon = "", args = { direction = "right" } },
    { "s", "pane.split", desc = "Split down", icon = "", args = { direction = "down" } },
    -- Shared menu/dispatch confirmation requires y; cancellation queues nothing.
    { "x", "pane.close", desc = "Close pane", icon = "󰅖", confirm = true },
}
