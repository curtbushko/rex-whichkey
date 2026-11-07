-- Shared launch-only module. Go owns JSON menu data, interaction and effects.
--
-- Installed binding (no executable nomination):
--   local whichkey = require("rex_whichkey")
--   rex.bind("ctrl+space", function(ctx) return whichkey.launch(ctx) end)
-- Optional launch data: { width=.8, height=.3, config="/path/menu.json",
--                         ascii=false, no_color=false }.
--
-- Development, after `task build` in this checkout, without init.lua reload:
--   rex do "$PWD/lua/rex_whichkey.lua" --action rex_whichkey_launch \
--     development_directory="$PWD/bin/$(go env GOOS)-$(go env GOARCH)"
-- Use explicit -s SESSION -b BLOCK and -C CLIENT when launching for an owned
-- origin from another terminal. The action receives those as ctx, never from
-- menu data. The trusted source/build caller supplies its absolute bin/GOOS-GOARCH
-- directory as data, scoped to this one terminal's path_prepend. Omit it for
-- installed launches using ordinary PATH. Lua does not discover files or run Go.
-- The directory is one PATH entry, not an executable path, callback or command.
-- The terminal always starts the fixed program rex-whichkey, shell=none.
--
-- Native chrome uses Rex defaults. Terminal palette matching is Go's job.
-- exit.on_completion=true with a zero quick-failure threshold lets Rex delete
-- even a failed/detached host after Go exits. A launch result is not UI readiness
-- or proof of eventual deletion; observe session.list_blocks after completion.
local module_name = ...
local M = {}

local function resource_id(value, kind)
    if type(value) ~= "string" or value:sub(1, #kind + 1) ~= kind .. ":" then
        return false
    end
    local uuid = value:sub(#kind + 2)
    return #uuid == 36 and uuid:sub(9, 9) == "-" and uuid:sub(14, 14) == "-"
        and uuid:sub(19, 19) == "-" and uuid:sub(24, 24) == "-"
        and uuid:gsub("%x", "") == "----"
end

local function required_id(value, kind)
    if not resource_id(value, kind) then
        error("rex_whichkey: invocation requires a full " .. kind .. " ID", 0)
    end
    return value
end

local function text(value, name)
    if type(value) ~= "string" or value == "" or value:find("[%z\1-\31\127]") then
        error("rex_whichkey: invalid " .. name, 0)
    end
    return value
end

local function fraction(value, default, name)
    if value == nil then return default end
    if type(value) ~= "number" or not (value > 0 and value <= 1) then
        error("rex_whichkey: " .. name .. " must be a finite fraction in (0, 1]", 0)
    end
    return value
end

local function exact_origin(session, block)
    local view, err = rex.session.view{ session_id = session }
    if not view then error("rex_whichkey: original view unavailable: " .. (err or "unknown error"), 0) end
    if view.session_id ~= session or type(view.windows) ~= "table" then
        error("rex_whichkey: original session view is invalid", 0)
    end
    local seen, found = {}, nil
    for _, window in ipairs(view.windows) do
        local window_id = required_id(window.window_id, "window")
        if seen[window_id] or type(window.layers) ~= "table" then
            error("rex_whichkey: invalid original window inventory", 0)
        end
        seen[window_id] = true
        for _, layer in ipairs(window.layers) do
            local layer_id = required_id(layer.layer_id, "layer")
            if seen[layer_id] or type(layer.blocks) ~= "table" then
                error("rex_whichkey: invalid original layer inventory", 0)
            end
            seen[layer_id] = true
            for _, item in ipairs(layer.blocks) do
                local id = required_id(item.block_id, "block")
                if seen[id] then error("rex_whichkey: duplicate original block placement", 0) end
                seen[id] = true
                if id == block then found = { window = window_id, layer = layer_id } end
            end
        end
    end
    if not found then error("rex_whichkey: invocation block is not placed; no focused fallback", 0) end
    return found
end

function M.launch(ctx, options)
    if type(ctx) ~= "table" then error("rex_whichkey: original invocation context is required", 0) end
    if options == nil then options = {} end
    if type(options) ~= "table" then error("rex_whichkey: launch options must be data", 0) end
    local allowed = { width=true, height=true, config=true, ascii=true, no_color=true,
        development_directory=true }
    for key in pairs(options) do
        if not allowed[key] then error("rex_whichkey: unsupported launch option " .. tostring(key), 0) end
    end
    -- Copy original intent before any discovery or creation. Popup REX_* IDs
    -- are deliberately not supplied here; Rex sets its own execution identity.
    local session = required_id(ctx.session_id, "session")
    local block = required_id(ctx.block_id, "block")
    local client = ctx.client_id
    if client ~= nil then client = required_id(client, "client") end
    local width = fraction(options.width, .8, "width")
    local height = fraction(options.height, .3, "height")
    for _, key in ipairs({ "ascii", "no_color" }) do
        if options[key] ~= nil and type(options[key]) ~= "boolean" then
            error("rex_whichkey: " .. key .. " must be boolean", 0)
        end
    end
    local config = options.config
    if config ~= nil then config = text(config, "config path") end
    local path
    if options.development_directory ~= nil then
        local directory = text(options.development_directory, "development directory")
        if directory:sub(1, 1) ~= "/" or directory:find(":", 1, true) then
            error("rex_whichkey: development_directory must be one absolute PATH directory", 0)
        end
        path = { directory }
    end
    local origin = exact_origin(session, block)
    local command = { "rex-whichkey", "direct", "--origin-session", session,
        "--origin-block", block, "--origin-window", origin.window, "--origin-layer", origin.layer }
    if client then command[#command+1] = "--client"; command[#command+1] = client end
    if config then command[#command+1] = "--config"; command[#command+1] = config end
    if options.ascii then command[#command+1] = "--ascii" end
    if options.no_color then command[#command+1] = "--no-color" end
    local created, err = rex.session.new_layer{
        session_id = session,
        window_id = origin.window,
        bounds = { (1-width)/2, (1-height)/2, width, height },
        focus = true,
        layout = rex.layout.block{
            flavor = "com.superlogical.terminal.shell",
            options = { command = command, shell = "none", path_prepend = path,
                exit = { on_completion = true, quick_exit_threshold_ms = 0 } },
        },
    }
    if not created then
        error("rex_whichkey: float creation failed or is uncertain; not retried: " .. (err or "unknown error"), 0)
    end
    required_id(created.layer_id, "layer")
    if type(created.block_ids) ~= "table" or #created.block_ids ~= 1 then
        error("rex_whichkey: incomplete float identity; creation is not retried", 0)
    end
    required_id(created.block_ids[1], "block")
    return true, created -- Preserve the function-binding handled result.
end

-- Both rex do FILE --action and canonical require use this implementation.
-- This launch-only action needs no installed lookup or user configuration reload;
-- it is not a menu callback registry.
rex.action{
    name = "rex_whichkey_launch",
    title = "Which-key menu",
    args = { width="number?", height="number?", config="string?", ascii="boolean?",
        no_color="boolean?", development_directory="string?" },
    run = function(ctx, args)
        local _, created = M.launch(ctx, args)
        return created
    end,
}

-- require passes the module name as the chunk vararg. Only canonical require
-- gets the function table; rex do serializes script results, so return none there.
if module_name == "rex_whichkey" then return M end
