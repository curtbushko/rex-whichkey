-- Launch only: Go owns configuration, Rex discovery, interaction and cleanup.
local M = {}

function M.launch(ctx)
    ctx = ctx or {}
    local process, err = rex.exec.start({ "rex-whichkey" }, {
        detach = true,
        -- env overlays Rex's environment; false removes an unknown identity.
        env = {
            REX_SESSION = ctx.session_id or false,
            REX_BLOCK = ctx.block_id or false,
            REX_CLIENT = ctx.client_id or false,
        },
    })
    if not process then
        error("rex_whichkey: could not start rex-whichkey: " .. (err or "unknown start error"))
    end
    -- A process handle acknowledges start, not menu display or later success.
    return true
end

return M
