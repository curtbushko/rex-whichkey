-- Raw spec fragment under mapping.go/docs/mapping-format.md, not a second
-- configuration schema. The shared authenticated dispatcher queues these on the
-- originating client; no current-client lookup or Rex call runs during import.
return {
    { "e", "client.sidebar.toggle", desc = "Sidebar", icon = "󰙀" },
    { "c", "client.palette.toggle", desc = "Command palette", icon = "" },
}
