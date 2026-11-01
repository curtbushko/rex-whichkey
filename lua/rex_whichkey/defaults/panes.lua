-- Ordered spec fragment; mapping.go/docs/mapping-format.md own its format.
-- The dispatcher restores the authenticated origin before queueing these client
-- actions. Importing defaults must never split, focus or close a pane itself.
return {
    -- These are pane.split action directions, not CLI --split values: the
    -- downward action takes "down", whereas the CLI uses "below".
    { "v", "pane.split", desc = "Split right", icon = "", args = { direction = "right" } },
    { "s", "pane.split", desc = "Split down", icon = "", args = { direction = "down" } },
    { "f", group = "Focus", icon = "󱡁" },
    { "fh", "pane.focus", desc = "Left", icon = "", args = { direction = "left" } },
    { "fj", "pane.focus", desc = "Down", icon = "", args = { direction = "down" } },
    { "fk", "pane.focus", desc = "Up", icon = "", args = { direction = "up" } },
    { "fl", "pane.focus", desc = "Right", icon = "", args = { direction = "right" } },
    { "z", "pane.zoom", desc = "Zoom pane", icon = "󰊓" },
    -- Shared menu/dispatch confirmation requires y; cancellation queues nothing.
    { "x", "pane.close", desc = "Close pane", icon = "󰅖", confirm = true },
}
