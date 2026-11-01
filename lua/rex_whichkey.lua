-- Short server-runtime composition. mapping.go/docs/mapping-format.md own the
-- data; config/mapping/dispatch retain all executable Lua in this original VM.
-- No package.path rewrite, client-local executable, timer or suspended action.
local config = require("rex_whichkey.config")
local mapping = require("rex_whichkey.mapping")
local dispatch = require("rex_whichkey.dispatch")
local M = {}
local completion_action

local function fail(message)
    error("rex_whichkey: " .. message, 0)
end

local function encode(value)
    local text, err = rex.json.encode(value)
    if err ~= nil or type(text) ~= "string" then
        fail("cannot encode invocation data")
    end
    return text
end

-- Rex encodes empty Lua tables as objects. Encode the ordered spec explicitly as
-- an array, including [], without changing args' legitimate empty-object shape.
-- Field spelling is consumed from Config; no alternative Lua validator/schema.
local function configuration(options, records)
    local values = {}
    for i, record in ipairs(records) do
        values[i] = encode(record)
    end
    local display = encode({
        position = options.position,
        title = options.title,
        icon = options.icon,
        icons = options.icons,
    })
    return display:sub(1, -2) .. ',"spec":[' .. table.concat(values, ",") .. "]}"
end

local function endpoint(value)
    -- This is explicit server-owned transport configuration, NOT ctx.server.
    -- The latter remains Rex's opaque identity, untouched in the original ctx.
    -- Go is the authoritative endpoint validator; this only rejects obviously
    -- missing/label-only bindings without guessing or discovering an endpoint.
    return type(value) == "string"
        and not value:find("[%z%s]")
        and (value:match("^unix:///.") ~= nil or value:match("^.+:%d+$") ~= nil)
end

local function connection(ctx)
    local result, err = rex.session.list_clients({ session_id = ctx.session_id })
    if err ~= nil or type(result) ~= "table" or type(result.clients) ~= "table" then
        fail("cannot inspect originating client")
    end
    local matches, generation = 0, nil
    for _, client in ipairs(result.clients) do
        if client.client_id == ctx.client_id then
            matches = matches + 1
            if client.connection_state == "lost" then
                fail("originating client connection was lost")
            end
            generation = client.connected_at
        end
    end
    if matches ~= 1 or type(generation) ~= "string" or generation == "" then
        fail("originating client connection is missing or ambiguous")
    end
    return generation
end

-- Explicit at-load registration in the user-owned init.lua; nothing here writes
-- or reloads it. Use the SAME installed module/VM for launch and completion:
--   local popup = require("rex_whichkey")
--   popup.register_completion("whichkey_complete")
--   local menu = popup.new(defaults, "unix:///absolute/server/socket")
--   -- Supply THIS server's fixed endpoint, not ctx.server or a host label.
--   -- The defaults owner supplies the resolved baseline; setup cannot replace
--   -- the transport binding. There is no environment/current-server fallback.
--   menu.setup(inline_options, optional_user_module_name)
--   rex.action{ name="whichkey", title="Which key", run=function(ctx) menu.launch(ctx) end }
-- The completion action is NOT a script file or a popup-origin action.
function M.register_completion(name)
    if completion_action ~= nil then
        fail("completion action is already registered")
    end
    dispatch.register_completion(name)
    completion_action = name
end

-- Defaults are supplied by the defaults owner, not invented named actions here.
-- Setup uses the delivered transactional loader/compiler and only then replaces
-- the installed invocation registry. Failure retains the last working setup.
function M.new(defaults, server_endpoint)
    if not endpoint(server_endpoint) then
        fail("an explicit server-owned endpoint is required")
    end
    -- Strings are immutable; retain this constructor-only binding separately
    -- from mapping options and the actual action context for the menu's lifetime.
    local loader = config.new(defaults, function(candidate)
        return mapping.compile(candidate.spec)
    end)
    local registry, installed
    local menu = {}

    function menu.setup(inline, module_name)
        local candidate, template = loader.setup(inline, module_name)
        local replacement = dispatch.new(template)
        registry, installed = replacement, candidate
    end

    function menu.launch(ctx)
        if completion_action == nil then
            fail("register the completion configuration action before launching")
        end
        if registry == nil then
            fail("setup must succeed before launching")
        end
        if type(ctx) ~= "table" then
            fail("the original action context is required")
        end
        -- Capture this generation/options before any call can yield to setup.
        -- issue reserves before evaluation; bind rechecks that same generation
        -- and connection after every intervening external preparation call.
        local generation, options = registry, installed
        local started = os.time()
        if type(started) ~= "number" then
            fail("wall clock is unavailable")
        end
        local records, ticket = generation.issue(ctx)
        local ok = pcall(function()
            local connected_at = connection(ctx)
            local resolved = configuration(options, records)
            -- rex.exec resolves the installed helper on THIS server's PATH.
            -- Only captured stdout contains the random capability. Never return
            -- it to the invoking action, print it, or propagate child errors.
            local prepared, prepare_err = rex.exec({ "rex-whichkey", "prepare" }, {
                timeout = 3,
                stdin = resolved,
                stderr = false,
            })
            if prepare_err ~= nil or type(prepared) ~= "table" or not prepared.ok then
                fail("server-side invocation preparation failed")
            end
            local secret, decode_err = rex.json.decode(prepared.stdout)
            if
                decode_err ~= nil
                or type(secret) ~= "table"
                or type(secret.capability) ~= "string"
                or #secret.capability ~= 64
                or secret.capability:find("[^0-9a-f]")
                or type(secret.executable) ~= "string"
                or secret.executable:sub(1, 1) ~= "/"
                or secret.executable:find("%z")
            then
                fail("invalid server-side invocation preparation")
            end
            generation.bind(ticket, ctx, secret.capability)
            local handoff = encode({
                configuration = resolved,
                capability = secret.capability,
                server = server_endpoint,
                session_id = ctx.session_id,
                block_id = ctx.block_id,
                client_id = ctx.client_id,
                connected_at = connected_at,
                completion_action = completion_action,
                -- Consume Go-owned InvocationLifetimeSeconds, not a renewable
                -- per-redraw timeout. Reservation began no earlier than started.
                expires = started + 30,
            })
            -- No waiting for selection here. stdin is a finite string; Rex closes
            -- it after writing. detach is the documented action-lifetime escape.
            -- The bounded Go controller outlives the action and popup terminal;
            -- it alone owns popup cleanup and the fresh completion activation.
            local handle, launch_err = rex.exec.start({ secret.executable, "controller" }, {
                stdin = handoff,
                detach = true,
                timeout = 35,
                stdout = false,
                stderr = false,
            })
            if launch_err ~= nil or handle == nil then
                fail("cannot start detached popup controller")
            end
        end)
        if not ok then
            generation.cancel(ticket)
            -- Neither raw process diagnostics nor arbitrary user-code errors can
            -- expose a capability-bearing handoff. No fallback or retry.
            fail("popup launch failed; invocation cancelled")
        end
        -- Never return a ticket, capability, handle or callback to the caller.
    end

    return menu
end

return M
