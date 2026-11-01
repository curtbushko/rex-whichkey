-- Internal consumer of mapping.go and docs/mapping-format.md. No wire envelope
-- or user setup API lives here. The launcher retains the opaque Lua ticket and
-- calls dispatch from the SAME originating Rex action after its menu completes.
-- A later CLI/RPC action cannot recreate that activation by copying ctx fields:
-- rex.invoke/client.queue inherit the ambient action, not an explicit ctx arg.
local evaluate = require("rex_whichkey.evaluate")
local M = {}
local current
local max_pending, lifetime = 64, 30

local function fail(message)
    error("rex_whichkey.dispatch: " .. message, 0)
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

local function origin(ctx)
    if not plain(ctx) then
        fail("a Rex action context is required")
    end
    local server_kind = type(ctx.server)
    if
        (server_kind ~= "string" and server_kind ~= "table" and server_kind ~= "userdata")
        or ctx.server == ""
        or not resource(ctx.session_id, "session")
        or not resource(ctx.block_id, "block")
        or not resource(ctx.client_id, "client")
    then
        fail("origin server, session, block and client identities are required")
    end
    -- server is an opaque Rex-provided identity, not a name to resolve against
    -- the local host. Keep it verbatim; all calls stay in this action's runtime.
    return {
        server = ctx.server,
        session_id = ctx.session_id,
        block_id = ctx.block_id,
        client_id = ctx.client_id,
        origin = ctx.origin,
    }
end

local function same_origin(ctx, expected)
    return plain(ctx)
        and rawequal(ctx.server, expected.server)
        and ctx.session_id == expected.session_id
        and ctx.block_id == expected.block_id
        and ctx.client_id == expected.client_id
        and ctx.origin == expected.origin
end

local function callback_context(expected)
    return {
        server = expected.server,
        session_id = expected.session_id,
        block_id = expected.block_id,
        client_id = expected.client_id,
        origin = expected.origin,
    }
end

-- One installed registry per required module. A replacement invalidates even
-- suspended dispatches before their next side effect. The composition owner must
-- call invalidate on configuration reload/teardown (including a reused module),
-- and create a replacement only after config compilation has succeeded.
function M.new(template)
    if not plain(template) or type(template.children) ~= "table" then
        fail("a compiled mapping template is required")
    end
    local registry = {}
    local entries, count, last_time = {}, 0, nil
    local closed = false

    local function release(ticket)
        local entry = entries[ticket]
        if entry then
            entries[ticket], count = nil, count - 1
            if entry.invocation then
                entry.invocation.release()
                entry.invocation = nil
            end
            entry.context = nil
        end
    end

    function registry.invalidate()
        closed = true
        for ticket in pairs(entries) do
            release(ticket)
        end
        template = nil
        if current == registry then
            current = nil
        end
    end

    -- Config Lua has no timer API. Expiry is lazy, bounded by max_pending even
    -- while idle, and checked again after every yielding API/user-code boundary.
    -- A backwards wall-clock jump invalidates the registry rather than extending
    -- any capability. No CPU-time clock is substituted for elapsed wall time.
    local function sweep()
        if closed then
            fail("runtime/configuration lifetime has ended")
        end
        local now = os.time()
        if type(now) ~= "number" or now ~= now or now == math.huge or now == -math.huge then
            registry.invalidate()
            fail("wall clock is unavailable")
        end
        if last_time and now < last_time then
            registry.invalidate()
            fail("wall clock moved backwards")
        end
        last_time = now
        for ticket, entry in pairs(entries) do
            if now >= entry.expires then
                release(ticket)
            end
        end
        return now
    end

    local function live(ticket, entry, ctx)
        sweep()
        if entries[ticket] ~= entry then
            fail("invalid, expired, cancelled or consumed invocation")
        end
        if not rawequal(ctx, entry.context) or not same_origin(ctx, entry.origin) then
            fail("dispatch must remain in its originating Rex action")
        end
    end

    local function inspect(ticket, entry, ctx)
        live(ticket, entry, ctx)
        local expected = entry.origin
        local view, err = rex.session.view({ session_id = expected.session_id })
        live(ticket, entry, ctx)
        if err ~= nil or type(view) ~= "table" or view.session_id ~= expected.session_id then
            fail("origin session is unavailable")
        end
        local placements = 0
        for _, window in ipairs(view.windows or {}) do
            for _, layer in ipairs(window.layers or {}) do
                for _, block in ipairs(layer.blocks or {}) do
                    if block.block_id == expected.block_id then
                        placements = placements + 1
                    end
                end
            end
        end
        if placements ~= 1 then
            fail("origin block is missing, detached or ambiguously placed")
        end
        -- Never use a client label, whichever app is now focused, or an arbitrary
        -- client attached to the server. Bind the connection as well as its ID.
        local result, client_err = rex.session.list_clients({ session_id = expected.session_id })
        live(ticket, entry, ctx)
        if client_err ~= nil or type(result) ~= "table" or type(result.clients) ~= "table" then
            fail("origin client is unavailable")
        end
        local matches, connected_at = 0, nil
        for _, client in ipairs(result.clients) do
            if client.client_id == expected.client_id then
                matches = matches + 1
                if client.connection_state == "lost" then
                    fail("origin client connection was lost")
                end
                connected_at = client.connected_at
            end
        end
        if matches ~= 1 or type(connected_at) ~= "string" or connected_at == "" then
            fail("origin client is no longer attached")
        end
        if entry.connected_at and entry.connected_at ~= connected_at then
            fail("origin client connection changed")
        end
        entry.connected_at = connected_at
    end

    -- The ticket is an unforgeable, runtime-local table identity held ONLY by
    -- the launcher, separate from callback_ref and all menu data. It deliberately
    -- has no fields, sequential/string bearer token, or transport representation.
    -- Return Go-shaped records separately; keep this ticket in the Rex runtime.
    function registry.issue(ctx)
        local now = sweep()
        if count >= max_pending then
            fail("too many pending invocations")
        end
        local ticket = {}
        local entry = { context = ctx, origin = origin(ctx), expires = now + lifetime, state = "evaluating" }
        entries[ticket], count = entry, count + 1
        local ok, records = pcall(function()
            inspect(ticket, entry, ctx)
            local data, invocation = evaluate.evaluate(template)
            -- Evaluation can yield or replace configuration. Always release its
            -- newly returned handle if the reserved invocation vanished meanwhile.
            entry.invocation = invocation
            live(ticket, entry, ctx)
            entry.state = "open"
            return data
        end)
        if not ok then
            if entry.invocation then
                entry.invocation.release()
                entry.invocation = nil
            end
            release(ticket)
            fail("invocation preparation failed")
        end
        return records, ticket
    end

    -- Validate only a data-only menu selection, not caller-supplied Mapping
    -- action/args/confirm metadata. This is an internal call, not a JSON decoder:
    -- its caller must reject duplicate/unknown JSON fields and null coercions.
    local function selection(value)
        if not plain(value) then
            fail("selection must be a plain data table")
        end
        for key in pairs(value) do
            if key ~= "keys" and key ~= "callback_ref" and key ~= "confirmed" then
                fail("unexpected selection field")
            end
        end
        if
            type(value.keys) ~= "string"
            or (value.callback_ref ~= nil and type(value.callback_ref) ~= "string")
            or (value.confirmed ~= nil and type(value.confirmed) ~= "boolean")
        then
            fail("invalid selection field type")
        end
        return value.keys, value.callback_ref, value.confirmed
    end

    function registry.dispatch(ticket, ctx, value)
        sweep()
        local entry = entries[ticket]
        if not entry or entry.state ~= "open" then
            fail("invalid, expired, consumed or reentrant invocation")
        end
        -- Claim before guards, Rex calls or callbacks can yield/reenter. Failed
        -- validation also consumes this invocation; nothing here retries actions.
        entry.state = "dispatching"
        local stage = "selection validation failed"
        local ok, result = pcall(function()
            live(ticket, entry, ctx)
            local keys, callback_ref, confirmed = selection(value)
            local retained = entry.invocation.resolve(keys, callback_ref, confirmed)
            stage = "origin validation failed"
            inspect(ticket, entry, ctx)
            -- Resolve placement afresh and focus the stable block, not the window
            -- captured at launch: a same-session move must not retarget dispatch.
            stage = "origin focus failed"
            local focused, err = rex.session.focus_block({
                session_id = entry.origin.session_id,
                block_id = entry.origin.block_id,
            })
            live(ticket, entry, ctx)
            if err ~= nil or type(focused) ~= "table" then
                fail("origin focus failed")
            end
            -- These APIs do not offer an atomic focus-and-execute transaction.
            -- Configuration/expiry/cancellation are checked after focus returns;
            -- concurrent external moves/disconnects can still race execution.
            if retained.callback then
                stage = "callback failed"
                return retained.callback(callback_context(entry.origin), retained.args)
            end
            if retained.action:find(".", 1, true) then
                -- Published Custom Actions/Lua API: dotted names belong to the
                -- client/block, while config actions have undotted names. Queue
                -- runs after the SAME originating action returns successfully.
                stage = "client action queue failed"
                local _, queue_err = rex.client.queue(retained.action, retained.args)
                if queue_err ~= nil then
                    fail("client action queue failed")
                end
                return nil
            end
            stage = "named action failed"
            local action_result, action_err = rex.invoke(retained.action, retained.args)
            if action_err ~= nil then
                fail("named action failed")
            end
            return action_result
        end)
        release(ticket)
        if not ok then
            -- User callbacks/conditions and Rex failures may throw arbitrary Lua
            -- values or terminal control text. Never surface those values here.
            fail(stage .. "; invocation consumed, not retried")
        end
        return result
    end

    -- Dismissal/error cleanup is local to the owning launcher. No selection,
    -- action, focus change or cascading Rex resource cleanup is performed.
    function registry.cancel(ticket)
        if not closed then
            sweep()
            release(ticket)
        end
    end

    if current then
        current.invalidate()
    end
    current = registry
    return registry
end

return M
