-- Resolved baseline consumed by rex_whichkey.new(defaults, server_endpoint).
-- Display values follow Go DefaultConfig; fragments use the Go-owned mapping
-- standard. User setup/import precedence and whole-spec replacement stay in the
-- transactional config loader, never in this aggregation step.
local spec = {}
for _, name in ipairs({ "sessions", "tabs", "panes", "app", "hosts" }) do
    for _, entry in ipairs(require("rex_whichkey.defaults." .. name)) do
        spec[#spec + 1] = entry
    end
end

return {
    position = "center",
    title = "Rex",
    icon = "",
    icons = { breadcrumb = "»", group = "+" },
    spec = spec,
}
