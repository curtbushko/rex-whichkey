-- Consumer of mapping.go and docs/mapping-format.md, including the accepted
-- post-return bridge. Only the registered CONFIG action accepts completion;
-- callbacks, tickets, origin authority and snapshots never leave this Lua VM.
local evaluate = require("rex_whichkey.evaluate")
local mapping = require("rex_whichkey.mapping")
local M = {}
local current, complete_current
-- Go-owned InvocationLifetimeSeconds / MaxPendingInvocations, not setup options.
local max_pending, lifetime = 64, 30

local function fail(message)
    error("rex_whichkey.dispatch: " .. message, 0)
end

local function plain(value)
    return type(value) == "table" and getmetatable(value) == nil
end

local function capability(value)
    return type(value) == "string" and #value == 64 and not value:find("[^0-9a-f]")
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
    return {
        server = ctx.server,
        session_id = ctx.session_id,
        block_id = ctx.block_id,
        client_id = ctx.client_id,
        origin = ctx.origin,
    }
end

local function same_origin(ctx, expected)
    -- The activation's source (key/palette versus cli/api) is NOT identity.
    -- server remains Rex's opaque identity; never resolve it via a local fallback.
    return plain(ctx)
        and rawequal(ctx.server, expected.server)
        and ctx.session_id == expected.session_id
        and ctx.block_id == expected.block_id
        and ctx.client_id == expected.client_id
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

-- Lua 5.1 has no utf8 library. Reject malformed UTF-8 rather than accepting the
-- replacement characters used by permissive JSON decoders (see Completion).
local function codepoint(text, position)
    local first = text:byte(position)
    local size, code, minimum
    if not first then
        fail("invalid completion JSON")
    elseif first < 128 then
        size, code, minimum = 1, first, 0
    elseif first >= 194 and first <= 223 then
        size, code, minimum = 2, first - 192, 128
    elseif first >= 224 and first <= 239 then
        size, code, minimum = 3, first - 224, 2048
    elseif first >= 240 and first <= 244 then
        size, code, minimum = 4, first - 240, 65536
    else
        fail("invalid completion JSON")
    end
    for j = 1, size - 1 do
        local byte = text:byte(position + j)
        if not byte or byte < 128 or byte > 191 then
            fail("invalid completion JSON")
        end
        code = code * 64 + byte - 128
    end
    if code < minimum or code > 1114111 or (code >= 55296 and code <= 57343) then
        fail("invalid completion JSON")
    end
    return code, size
end

local function encode_codepoint(code)
    if code < 128 then
        return string.char(code)
    elseif code < 2048 then
        return string.char(192 + math.floor(code / 64), 128 + code % 64)
    elseif code < 65536 then
        return string.char(224 + math.floor(code / 4096), 128 + math.floor(code / 64) % 64, 128 + code % 64)
    end
    return string.char(
        240 + math.floor(code / 262144),
        128 + math.floor(code / 4096) % 64,
        128 + math.floor(code / 64) % 64,
        128 + code % 64
    )
end

local function reference(value)
    if type(value) ~= "string" or value == "" then
        fail("invalid completion callback reference")
    end
    local i, nonblank = 1, false
    while i <= #value do
        local code, size = codepoint(value, i)
        if code < 32 or (code >= 127 and code <= 159) then
            fail("invalid completion callback reference")
        end
        -- Go strings.TrimSpace's Unicode White_Space, after excluding controls.
        if
            code ~= 32
            and code ~= 160
            and code ~= 5760
            and not (code >= 8192 and code <= 8202)
            and code ~= 8232
            and code ~= 8233
            and code ~= 8239
            and code ~= 8287
            and code ~= 12288
        then
            nonblank = true
        end
        i = i + size
    end
    if not nonblank then
        fail("invalid completion callback reference")
    end
end

-- Strict raw-JSON consumer of Completion, not a Lua-owned wire format. Do not
-- rex.json.decode first: it erases duplicate keys and null versus missing fields.
-- Bound the transport message (64 KiB), including whitespace, before parsing it.
local function completion(text)
    if type(text) ~= "string" or #text > 65536 then
        fail("invalid completion JSON")
    end
    local position = 1
    local function invalid()
        fail("invalid completion JSON")
    end
    local function whitespace()
        while true do
            local byte = text:byte(position)
            if byte ~= 32 and byte ~= 9 and byte ~= 10 and byte ~= 13 then
                return
            end
            position = position + 1
        end
    end
    local function take(char)
        whitespace()
        if text:sub(position, position) ~= char then
            invalid()
        end
        position = position + 1
    end
    local function hex4()
        local hex = text:sub(position, position + 3)
        if #hex ~= 4 or hex:find("[^0-9a-fA-F]") then
            invalid()
        end
        position = position + 4
        return tonumber(hex, 16)
    end
    local escapes = { ['"'] = '"', ["\\"] = "\\", ["/"] = "/", b = "\b", f = "\f", n = "\n", r = "\r", t = "\t" }
    local function string_value()
        take('"')
        local parts = {}
        while position <= #text do
            local char = text:sub(position, position)
            if char == '"' then
                position = position + 1
                return table.concat(parts)
            elseif char == "\\" then
                position = position + 1
                local escape = text:sub(position, position)
                position = position + 1
                if escape == "u" then
                    local code = hex4()
                    if code >= 55296 and code <= 56319 then
                        if text:sub(position, position + 1) ~= "\\u" then
                            invalid()
                        end
                        position = position + 2
                        local low = hex4()
                        if low < 56320 or low > 57343 then
                            invalid()
                        end
                        code = 65536 + (code - 55296) * 1024 + low - 56320
                    elseif code >= 56320 and code <= 57343 then
                        invalid()
                    end
                    parts[#parts + 1] = encode_codepoint(code)
                elseif escapes[escape] then
                    parts[#parts + 1] = escapes[escape]
                else
                    invalid()
                end
            else
                local code, size = codepoint(text, position)
                if code < 32 then
                    invalid()
                end
                parts[#parts + 1] = text:sub(position, position + size - 1)
                position = position + size
            end
        end
        invalid()
    end
    local function boolean_value()
        whitespace()
        if text:sub(position, position + 3) == "true" then
            position = position + 4
            return true
        elseif text:sub(position, position + 4) == "false" then
            position = position + 5
            return false
        end
        invalid()
    end
    local function version_value()
        take("1") -- Go's integer version 1, not 1.0, 1e0, a string or boolean.
        return 1
    end
    local function object(readers)
        take("{")
        local fields = {}
        whitespace()
        if text:sub(position, position) == "}" then
            position = position + 1
            return fields
        end
        while true do
            local key = string_value()
            local reader = readers[key]
            if not reader or fields[key] ~= nil then
                invalid()
            end
            take(":")
            fields[key] = reader()
            whitespace()
            local char = text:sub(position, position)
            position = position + 1
            if char == "}" then
                return fields
            elseif char ~= "," then
                invalid()
            end
        end
    end
    local value = object({
        version = version_value,
        capability = string_value,
        kind = string_value,
        selection = function()
            return object({ keys = string_value, callback_ref = string_value, confirmed = boolean_value })
        end,
    })
    whitespace()
    if position <= #text or value.version ~= 1 or not capability(value.capability) then
        invalid()
    end
    if value.kind == "cancel" then
        if value.selection ~= nil then
            invalid()
        end
    elseif value.kind == "complete" then
        local selected = value.selection
        if not selected or type(selected.keys) ~= "string" or type(selected.confirmed) ~= "boolean" then
            invalid()
        end
        local _, canonical = mapping.parse_keys(selected.keys, "completion.keys")
        if canonical ~= selected.keys then
            invalid()
        end
        if selected.callback_ref ~= nil then
            reference(selected.callback_ref)
        end
    else
        invalid()
    end
    return value
end

-- Explicit opt-in at config LOAD; never rewrite or reload the user's init.lua.
-- Minimal composition contract (all functions run in this same server Lua VM):
--   local dispatch = require("rex_whichkey.dispatch")
--   dispatch.register_completion("whichkey_complete")
--   local registry = dispatch.new(compiled_template)
-- In the original action: records, ticket = registry.issue(ctx); obtain a fresh
-- capability from Go NewInvocationCapability via the trusted launcher, then
-- registry.bind(ticket, ctx, capability) BEFORE exposing the popup. On launch
-- failure registry.cancel(ticket). Never serialize ticket or put a capability in
-- argv/logs. 04-05 owns that trusted handoff and detached popup lifecycle.
--
-- Return completion through the installed action, not a rex do Lua script VM:
-- rex --server SERVER --client CLIENT do whichkey_complete --session SESSION
--     --block BLOCK --args -
-- stdin is {"message": RAW_COMPLETION_JSON_STRING}, encoded as JSON data. SERVER
-- must explicitly select the captured endpoint, and CLIENT/SESSION/BLOCK must be
-- captured full IDs, never defaults or popup IDs. Use argument arrays. This run
-- wrapper gets the ACTUAL fresh action ctx; copying ctx fields cannot rebind the
-- ambient rex.invoke/client.queue activation. No public dispatch(ctx, ...) exists.
function M.register_completion(name)
    if type(name) ~= "string" or not name:match("^[a-z0-9_]+$") then
        fail("a valid completion configuration action name is required")
    end
    rex.action({
        name = name,
        title = "Complete which-key invocation",
        args = { message = "string" },
        run = function(ctx, args)
            if not plain(args) or type(args.message) ~= "string" then
                fail("completion requires raw JSON message data")
            end
            for key in pairs(args) do
                if key ~= "message" then
                    fail("unexpected completion argument")
                end
            end
            local handler = complete_current
            if not handler then
                fail("no active invocation registry")
            end
            -- Never return arbitrary callback results to a capability-bearing
            -- transport, nor let the launcher swallow a failure and commit queues.
            handler(ctx, args.message)
        end,
    })
end

-- One installed configuration generation. Candidate compilation must succeed
-- before replacement; composition calls invalidate on reload/reuse/teardown.
function M.new(template)
    if not plain(template) or type(template.children) ~= "table" then
        fail("a compiled mapping template is required")
    end
    local registry = {}
    local entries, capabilities, count, last_time = {}, {}, 0, nil
    local closed = false

    local function release(ticket)
        local entry = entries[ticket]
        if entry then
            entries[ticket], count = nil, count - 1
            if entry.capability and capabilities[entry.capability] == ticket then
                capabilities[entry.capability] = nil
            end
            if entry.invocation then
                entry.invocation.release()
                entry.invocation = nil
            end
            entry.context, entry.snapshot, entry.capability = nil, nil, nil
        end
    end

    function registry.invalidate()
        closed = true
        for ticket in pairs(entries) do
            release(ticket)
        end
        template = nil
        if current == registry then
            current, complete_current = nil, nil
        end
    end

    -- Config Lua has no timer API. Lazy expiry bounds retained state even idle.
    -- A backwards wall-clock jump revokes the generation; CPU time is not elapsed
    -- wall time. Check after every yielding API/user-code boundary as well.
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
        if not same_origin(ctx, entry.origin) then
            fail("action is not targeted at the invocation origin")
        end
    end

    local function inspect(ticket, entry, ctx, require_focus)
        live(ticket, entry, ctx)
        local expected = entry.origin
        local view, err = rex.session.view({ session_id = expected.session_id })
        live(ticket, entry, ctx)
        if err ~= nil or type(view) ~= "table" or view.session_id ~= expected.session_id then
            fail("origin session is unavailable")
        end
        local placements, placed_window = 0, nil
        for _, window in ipairs(view.windows or {}) do
            for _, layer in ipairs(window.layers or {}) do
                for _, block in ipairs(layer.blocks or {}) do
                    if block.block_id == expected.block_id then
                        placements, placed_window = placements + 1, window
                    end
                end
            end
        end
        if placements ~= 1 or not resource(placed_window.window_id, "window") then
            fail("origin block is missing, detached or ambiguously placed")
        end
        if
            require_focus
            and (
                view.active_window_id ~= placed_window.window_id
                or placed_window.focused_block_id ~= expected.block_id
            )
        then
            fail("origin block no longer has focus")
        end
        local result, client_err = rex.session.list_clients({ session_id = expected.session_id })
        live(ticket, entry, ctx)
        if client_err ~= nil or type(result) ~= "table" or type(result.clients) ~= "table" then
            fail("origin client is unavailable")
        end
        local matches, connected_at = 0, nil
        for _, client in ipairs(result.clients) do
            if client.client_id == expected.client_id then
                matches = matches + 1
                -- API contract: missing/unknown state is connected; only lost is
                -- disconnected. connected_at is still mandatory generation data.
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

    -- Reserve BEFORE evaluation. The unforgeable table ticket stays private to
    -- the trusted launcher. It has no transport representation or bearer spelling.
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
            -- Invalidation may have released the reservation while user code
            -- yielded. Attach its new handle before checking, then release below.
            entry.invocation = invocation
            live(ticket, entry, ctx)
            local snapshot = {}
            for _, record in ipairs(data) do
                if record.action or record.callback_ref then
                    snapshot[record.keys] = { callback_ref = record.callback_ref, confirm = record.confirm == true }
                end
            end
            entry.snapshot, entry.state = snapshot, "prepared"
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

    -- Trusted launcher only, still inside the issuing action. Go generates the
    -- random capability; Lua checks its canonical representation, not entropy.
    -- A reservation is not completable until binding finishes successfully.
    function registry.bind(ticket, ctx, value)
        sweep()
        local entry = entries[ticket]
        if not entry or entry.state ~= "prepared" or not rawequal(ctx, entry.context) then
            fail("invalid invocation capability binding")
        end
        entry.state = "binding"
        local ok = pcall(function()
            live(ticket, entry, ctx)
            if not capability(value) or capabilities[value] ~= nil then
                fail("invalid or duplicate invocation capability")
            end
            inspect(ticket, entry, ctx)
            entry.capability, entry.state, entry.context = value, "open", nil
            capabilities[value] = ticket
        end)
        if not ok then
            release(ticket)
            fail("invocation capability binding failed")
        end
    end

    local function complete(ctx, raw)
        -- Decode before lookup/claim. Malformed/unknown capabilities and wrong
        -- origins cannot consume an unrelated reservation. No secrets in errors.
        local decoded, value = pcall(completion, raw)
        if not decoded then
            -- In particular, the shared key parser may include rejected input in
            -- its error. Never expose raw completion fragments or capabilities.
            fail("invalid completion JSON")
        end
        sweep()
        local ticket = capabilities[value.capability]
        local entry = ticket and entries[ticket]
        if not entry or entry.state ~= "open" or not same_origin(ctx, entry.origin) then
            fail("invalid completion authority or invocation state")
        end
        -- Claim BEFORE conditions or Rex calls yield/reenter. Do not remove from
        -- entries until done: in-progress work counts toward the bounded capacity
        -- and remains revocable by expiry/reload/local cancellation.
        entry.state = "claimed"
        -- Keep the spelling reserved until release, even though claimed state
        -- revokes its authority now. A concurrent bind cannot reuse a live token.
        local stage = "completion validation failed"
        local ok = pcall(function()
            live(ticket, entry, ctx)
            stage = "origin validation failed"
            -- Authenticate the retained connection before ANY user condition can
            -- run, not only before the final selected callback/action.
            inspect(ticket, entry, ctx)
            if value.kind == "cancel" then
                return -- no guards, selected action or focus change
            end
            stage = "completion validation failed"
            local selected = value.selection
            local record = entry.snapshot[selected.keys]
            if not record or record.callback_ref ~= selected.callback_ref or record.confirm ~= selected.confirmed then
                fail("selection does not match the retained invocation")
            end
            -- The evaluator owns original closures/args/guards. Its resolver is
            -- also single-use; bridge checks above strengthen its older internal
            -- acceptance of key aliases, absent flags and empty callback refs.
            local retained = entry.invocation.resolve(selected.keys, selected.callback_ref, selected.confirmed)
            live(ticket, entry, ctx)
            stage = "origin validation failed"
            inspect(ticket, entry, ctx)
            stage = "origin focus failed"
            local focused, err = rex.session.focus_block({
                session_id = entry.origin.session_id,
                block_id = entry.origin.block_id,
            })
            live(ticket, entry, ctx)
            if err ~= nil or type(focused) ~= "table" then
                fail("origin focus failed")
            end
            -- Revalidate connection and placement AFTER focus's yielding call.
            -- Rex has no atomic focus-and-execute transaction: an external move
            -- or disconnect can still race the final check. Never retry actions.
            stage = "post-focus origin validation failed"
            inspect(ticket, entry, ctx, true)
            live(ticket, entry, ctx)
            if retained.callback then
                stage = "callback failed"
                retained.callback(callback_context(entry.origin), retained.args)
            elseif retained.action:find(".", 1, true) then
                -- Custom Actions/Lua API: dotted names are client/block actions.
                -- Queue on this ACTUAL origin-targeted completion activation;
                -- the registered wrapper returns without swallowing any error.
                stage = "client action queue failed"
                local _, queue_err = rex.client.queue(retained.action, retained.args)
                if queue_err ~= nil then
                    fail("client action queue failed")
                end
            else
                stage = "named action failed"
                local _, action_err = rex.invoke(retained.action, retained.args)
                if action_err ~= nil then
                    fail("named action failed")
                end
            end
            live(ticket, entry, ctx)
        end)
        release(ticket)
        if not ok then
            -- User code and API errors can contain secrets/control text. Failed
            -- authenticated requests consume the entry, even if execution may
            -- already have happened. Nothing here reopens or retries it.
            fail(stage .. "; invocation consumed, not retried")
        end
    end

    -- Local launch-failure/dismissal cleanup, requiring the private table ticket.
    -- Remote dismissal uses the authenticated registered cancel message instead.
    function registry.cancel(ticket)
        if not closed then
            sweep()
            release(ticket)
        end
    end

    if current then
        current.invalidate()
    end
    current, complete_current = registry, complete
    return registry
end

return M
