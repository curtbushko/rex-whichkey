-- Ordered spec fragment consumed by the defaults owner. mapping.go and
-- docs/mapping-format.md own the format; this module only supplies tab bindings.
-- Named client actions use the dispatcher's authenticated origin/focus path,
-- never the popup's context. Loading this fragment performs no Rex operations.
local spec = {
    { "t", group = "Tabs", icon = "" },
    { "tn", "client.tab.new", desc = "New tab", icon = "" },
    { "th", "client.tab.previous", desc = "Previous", icon = "" },
    { "tl", "client.tab.next", desc = "Next", icon = "" },
    -- No replacement name: Rex opens its native rename interaction.
    { "tr", "client.tab.rename", desc = "Rename", icon = "󰏫" },
    -- Confirmation belongs to the shared menu/dispatch path. Cancellation never
    -- queues this action; do not close tabs here or during configuration loading.
    { "tx", "client.tab.close", desc = "Close tab", icon = "󰅖", confirm = true },
}

-- Preserve the old t1..t9 bindings. Rex's client.tab.goto index is one-based;
-- pass numbers as action data, not key strings, window IDs or shell fragments.
for index = 1, 9 do
    spec[#spec + 1] = {
        "t" .. index,
        "client.tab.goto",
        desc = "Go to tab " .. index,
        icon = "",
        args = { index = index },
    }
end

return spec
