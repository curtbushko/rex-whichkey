-- Short server-runtime composition. mapping.go/docs/mapping-format.md own the
-- data; config/mapping/dispatch retain all executable Lua in this original VM.
-- No package.path rewrite, client-local executable, timer or suspended action.
local config = require("rex_whichkey.config")
local mapping = require("rex_whichkey.mapping")
local dispatch = require("rex_whichkey.dispatch")
local M = {}
local completion_action
local launch_names = {}

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

local function plain(value)
    return type(value) == "table" and getmetatable(value) == nil
end

local function resource(value, kind)
    if type(value) ~= "string" then
        return false
    end
    local prefix, a, b, c, d, e = value:match("^([a-z]+):(%x+)%-(%x+)%-(%x+)%-(%x+)%-(%x+)$")
    return prefix == kind and #a == 8 and #b == 4 and #c == 4 and #d == 4 and #e == 12
end

local function valid_generation(value)
    return type(value) == "string" and #value <= 1024 and value:find("%S") ~= nil and not value:find("%c")
end

-- Consume launch.go's five transport fields, not a mapping schema or ctx. The
-- endpoint's canonical meaning remains Go-owned and is checked against prepare.
local function launch_expectation(text)
    if type(text) ~= "string" or #text > 65536 then
        fail("launch origin must be bounded JSON string data")
    end
    local ok, expected, err = pcall(rex.json.decode, text)
    if not ok or err ~= nil or not plain(expected) then
        fail("invalid launch origin JSON")
    end
    local fields = { server = true, session_id = true, block_id = true, client_id = true, connected_at = true }
    local count = 0
    for key in pairs(expected) do
        if not fields[key] then
            fail("unexpected launch origin field")
        end
        count = count + 1
    end
    if
        count ~= 5
        or type(expected.server) ~= "string"
        or expected.server == ""
        or expected.server:find("%c")
        or not resource(expected.session_id, "session")
        or not resource(expected.block_id, "block")
        or not resource(expected.client_id, "client")
        or not valid_generation(expected.connected_at)
    then
        fail("launch origin requires endpoint, full session/block/client IDs and connected_at")
    end
    return expected
end

local function connection(ctx)
    local result, err = rex.session.list_clients({ session_id = ctx.session_id })
    if err ~= nil or type(result) ~= "table" or type(result.clients) ~= "table" then
        fail("cannot inspect originating client")
    end
    local matches, generation, seen = 0, nil, {}
    for _, client in ipairs(result.clients) do
        if not plain(client) or not resource(client.client_id, "client") or seen[client.client_id] then
            fail("originating session returned malformed or duplicate clients")
        end
        seen[client.client_id] = true
        if client.connection_state ~= nil and type(client.connection_state) ~= "string" then
            fail("originating session returned a malformed connection state")
        end
        if client.connection_state ~= "lost" and not valid_generation(client.connected_at) then
            fail("originating client has no valid connection generation")
        end
        if client.client_id == ctx.client_id then
            matches = matches + 1
            if client.connection_state == "lost" then
                fail("originating client connection was lost")
            end
            generation = client.connected_at
        end
    end
    if matches ~= 1 or not valid_generation(generation) then
        fail("originating client connection is missing or ambiguous")
    end
    return generation
end

-- Explicit at-load registration in the user-owned init.lua; nothing here writes
-- or reloads it. Use the SAME installed module/VM for launch and completion:
--   local popup = require("rex_whichkey")
--   popup.register_completion("whichkey_complete")
--   local menu = popup.new(defaults)
--   -- Optional second argument: THIS server's immutable endpoint override,
--   -- never ctx.server or a host label. Go resolves execution-host defaults.
--   -- The defaults owner supplies the baseline; setup cannot replace routing.
--   menu.setup(inline_options, optional_user_module_name)
--   menu.register_launch("whichkey") -- CLI, palette and key use the actual ctx
-- The completion action is NOT a script file or a popup-origin action.
function M.register_completion(name)
    if completion_action ~= nil then
        fail("completion action is already registered")
    end
    if launch_names[name] then
        fail("completion and launch actions must have distinct names")
    end
    dispatch.register_completion(name)
    completion_action = name
end

-- Defaults are supplied by the defaults owner, not invented named actions here.
-- Setup uses the delivered transactional loader/compiler and only then replaces
-- the installed invocation registry. Failure retains the last working setup.
function M.new(defaults, server_endpoint)
    if server_endpoint ~= nil and type(server_endpoint) ~= "string" then
        fail("the optional server-owned endpoint override must be a string")
    end
    -- Retain the optional immutable constructor binding outside mapping options
    -- and the actual action ctx. Only Go owns endpoint syntax/default policy;
    -- even an explicitly empty string is forwarded for fail-closed validation.
    local loader = config.new(defaults, function(candidate)
        return mapping.compile(candidate.spec)
    end)
    local registry, installed, launch_name
    local menu = {}

    function menu.setup(inline, module_name)
        local candidate, template = loader.setup(inline, module_name)
        local replacement = dispatch.new(template)
        registry, installed = replacement, candidate
    end

    -- At configuration load, install a named action in THIS original server VM.
    -- CLI origin data only constrains its actual ctx; keys/palette need no data.
    function menu.register_launch(name)
        if name == nil then
            name = "whichkey"
        end
        if type(name) ~= "string" or not name:match("^[a-z0-9_]+$") then
            fail("a valid launch configuration action name is required")
        end
        if launch_name ~= nil or launch_names[name] or name == completion_action then
            fail("launch action is already registered or conflicts with completion")
        end
        rex.action({
            name = name,
            title = "Which key",
            args = { origin = "string?" },
            run = function(ctx, args)
                if not plain(args) then
                    fail("launch arguments must be a table")
                end
                for key in pairs(args) do
                    if key ~= "origin" then
                        fail("unexpected launch argument")
                    end
                end
                -- Do not return any controller handle, reservation or secrets.
                menu.launch(ctx, args.origin)
            end,
        })
        launch_name, launch_names[name] = name, true
    end

    function menu.launch(ctx, origin_json)
        if completion_action == nil then
            fail("register the completion configuration action before launching")
        end
        if registry == nil then
            fail("setup must succeed before launching")
        end
        local server_kind = plain(ctx) and type(ctx.server)
        if
            not plain(ctx)
            or (server_kind ~= "string" and server_kind ~= "table" and server_kind ~= "userdata")
            or ctx.server == ""
            or not resource(ctx.session_id, "session")
            or not resource(ctx.block_id, "block")
            or not resource(ctx.client_id, "client")
        then
            fail("the original action context with full server/session/block/client identities is required")
        end
        local expected = origin_json ~= nil and launch_expectation(origin_json) or nil
        -- Capture actual opaque identity without stringifying it or making a
        -- replacement ctx. The caller's data cannot confer action authority.
        local actual_server, session_id, block_id, client_id = ctx.server, ctx.session_id, ctx.block_id, ctx.client_id
        local generation, options = registry, installed
        local started = os.time()
        if type(started) ~= "number" then
            fail("wall clock is unavailable")
        end
        local function check()
            local now = os.time()
            if
                registry ~= generation
                or installed ~= options
                or not plain(ctx)
                or not rawequal(ctx.server, actual_server)
                or ctx.session_id ~= session_id
                or ctx.block_id ~= block_id
                or ctx.client_id ~= client_id
                or type(now) ~= "number"
                or now < started
                or now >= started + 30
            then
                fail("origin or configuration changed during launch")
            end
            if
                expected
                and (
                    expected.session_id ~= session_id
                    or expected.block_id ~= block_id
                    or expected.client_id ~= client_id
                )
            then
                fail("launch origin does not match the actual action context")
            end
        end
        local connected_at = expected and expected.connected_at or nil
        local function check_connection()
            check()
            local current = connection(ctx)
            check() -- list_clients yields; compare the actual ctx again.
            if connected_at ~= nil and connected_at ~= current then
                fail("originating client connection changed during launch")
            end
            connected_at = current
        end
        local records, ticket
        local failure =
            "popup launch failed; check the exact origin IDs, connection generation and setup; invocation cancelled"
        local ok = pcall(function()
            -- Reject a supplied mismatched origin/generation BEFORE reserving or
            -- evaluating any user mapping. issue/bind retain placement checks.
            check_connection()
            records, ticket = generation.issue(ctx)
            check_connection()
            local resolved = configuration(options, records)
            -- rex.exec resolves the installed helper on THIS server's PATH.
            -- Only captured stdout contains the random capability. Never return
            -- it to the invoking action, print it, or propagate child errors.
            local prepare_argv = { "rex-whichkey", "prepare" }
            if server_endpoint ~= nil then
                prepare_argv[#prepare_argv + 1] = "--server"
                prepare_argv[#prepare_argv + 1] = server_endpoint
            end
            -- Only fixed guidance crosses the pcall boundary, including when
            -- exec/decode throws; never forward child output or error objects.
            failure =
                "server-side preparation failed; check the execution-host rex-whichkey helper, matching Rex CLI, running server/runtime environment and endpoint override or REX_SERVER; invocation cancelled"
            local prepared, prepare_err = rex.exec(prepare_argv, {
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
                or type(secret.server) ~= "string"
                or secret.server == ""
            then
                fail("invalid server-side invocation preparation")
            end
            failure = "popup launch failed; origin, connection or endpoint changed; invocation cancelled"
            check_connection()
            if expected and expected.server ~= secret.server then
                fail("launch endpoint differs from canonical server preparation")
            end
            generation.bind(ticket, ctx, secret.capability)
            check_connection() -- bind also yields; never start from a stale ctx.
            local handoff = encode({
                configuration = resolved,
                capability = secret.capability,
                server = secret.server,
                session_id = session_id,
                block_id = block_id,
                client_id = client_id,
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
            check()
            failure =
                "detached popup controller launch failed or origin changed; check the execution-host helper and context; invocation cancelled"
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
            check_connection()
        end)
        if not ok then
            if ticket ~= nil then
                generation.cancel(ticket)
            end
            -- Neither raw process diagnostics nor arbitrary user-code errors can
            -- expose a capability-bearing handoff. No fallback or retry.
            fail(failure)
        end
        -- Never return a ticket, capability, handle or callback to the caller.
    end

    return menu
end

return M
