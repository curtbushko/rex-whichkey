-- Consumer of mapping.go and docs/mapping-format.md, including the accepted
-- post-return bridge. Only the registered CONFIG action accepts completion;
-- callbacks, tickets, origin authority and snapshots never leave this Lua VM.
local evaluate = require("rex_whichkey.evaluate")
local mapping = require("rex_whichkey.mapping")
local M = {}
local handlers = {}
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

local function inline_rename(record)
    return record.callback_ref == nil and not record.confirm
        and (record.keys == "tr" and record.action == "client.tab.rename"
            or record.keys == "wr" and record.action == "session.rename")
        and (record.args == nil or plain(record.args) and next(record.args) == nil)
end

local function activation(ctx)
    if not plain(ctx) then
        fail("a Rex action context is required")
    end
    local server_kind = type(ctx.server)
    if
        (server_kind ~= "nil" and server_kind ~= "string" and server_kind ~= "table" and server_kind ~= "userdata")
        or ctx.server == ""
        or not resource(ctx.session_id, "session")
        or not resource(ctx.block_id, "block")
        or not resource(ctx.client_id, "client")
    then
        fail("origin session, block and client identities are required")
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
    -- The capability authorizes a fresh caller in the same server/session/block,
    -- not impersonation of the retained GUI target. Each activation's own caller
    -- is frozen separately by live(). Older Rex omits server; compare absence
    -- exactly because the registry is already scoped to this server Lua VM.
    return plain(ctx)
        and rawequal(ctx.server, expected.server)
        and ctx.session_id == expected.session_id
        and ctx.block_id == expected.block_id
        and resource(ctx.client_id, "client")
end

local function callback_context(expected, ctx)
    return {
        server = expected.server,
        session_id = expected.session_id,
        block_id = expected.block_id,
        client_id = expected.client_id,
        request_client_id = ctx.client_id,
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

M.validate_text = reference

-- Strict raw-JSON consumer for Go's fixed records, not a Lua-owned wire format.
-- Do not rex.json.decode first: it erases duplicates and null versus missing.
-- Completion/small records are 64 KiB; full opaque preparation is 8 MiB.
local function decode_record(text, schema, limit)
    if type(text) ~= "string" or #text > (limit or 65536) then
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
        local parts, run_start = {}, position
        local function flush()
            if position > run_start then
                parts[#parts + 1] = text:sub(run_start, position - 1)
            end
        end
        while position <= #text do
            local char = text:sub(position, position)
            if char == '"' then
                flush()
                position = position + 1
                return table.concat(parts)
            elseif char == "\\" then
                flush()
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
                run_start = position
            else
                local code, size = codepoint(text, position)
                if code < 32 then
                    invalid()
                end
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
    local function integer_value()
        whitespace()
        local digits = text:sub(position):match("^%d+")
        if not digits or #digits > 16 or (#digits > 1 and digits:sub(1, 1) == "0") then
            invalid()
        end
        position = position + #digits
        local value = tonumber(digits)
        if not value or value > 9007199254740991 then
            invalid() -- Never round a Go identity through a Lua double.
        end
        return value
    end
    local read
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
            fields[key] = read(reader)
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
    read = function(kind)
        if kind == "string" then
            return string_value()
        elseif kind == "boolean" then
            return boolean_value()
        elseif kind == "integer" then
            return integer_value()
        elseif type(kind) == "table" and kind.array then
            take("[")
            local values = {}
            whitespace()
            if text:sub(position, position) == "]" then
                position = position + 1
                return values
            end
            while true do
                values[#values + 1] = read(kind.array)
                whitespace()
                local char = text:sub(position, position)
                position = position + 1
                if char == "]" then
                    return values
                elseif char ~= "," then
                    invalid()
                end
            end
        elseif type(kind) == "table" then
            return object(kind)
        end
        invalid()
    end
    local value = object(schema)
    whitespace()
    if position <= #text then
        invalid()
    end
    return value
end

-- Fixed Go-owned transport records only. Opaque receipt/relay JSON stays a
-- string; this parser never rebuilds their large numeric file identities.
M.decode_record = decode_record

local function completion(text)
    local value = decode_record(text, {
        version = "integer", capability = "string", receipt_hash = "string", kind = "string",
        selection = { keys = "string", callback_ref = "string", confirmed = "boolean", text = "string" },
    })
    local function invalid()
        fail("invalid completion JSON")
    end
    if value.version ~= 2 or not capability(value.capability) or not capability(value.receipt_hash) then
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
        if selected.text ~= nil then
            if #selected.text > 4096 or (selected.keys ~= "tr" and selected.keys ~= "wr")
                or selected.callback_ref ~= nil or selected.confirmed then
                invalid()
            end
            reference(selected.text) -- Same strict Unicode/nonblank/control rules as Go.
            local i, length = 1, 0
            while i <= #selected.text do
                local _, size = codepoint(selected.text, i)
                i, length = i + size, length + 1
            end
            if length > 128 then
                invalid()
            end
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
-- In the original action: registry.issue(ctx, {client_id=APP, connected_at=TIME}); obtain a fresh
-- capability from Go via the trusted launcher, then bind its exact preparation
-- BEFORE exposing the popup. Only unclaimed local failures may cancel a ticket.
-- Never serialize a ticket or put a capability in argv/logs; Go owns lifecycle.
--
-- Return completion through the installed action, not a rex do Lua script VM:
-- rex --server SERVER --client CLIENT do whichkey_complete --session SESSION
--     --block BLOCK --args -
-- stdin is {"message": RAW_COMPLETION_JSON_STRING}, encoded as JSON data. SERVER
-- must explicitly select the captured endpoint, and CLIENT/SESSION/BLOCK are
-- retained routing IDs, never defaults or popup IDs. The wrapper gets the ACTUAL
-- fresh CLI action ctx, whose client_id is not APP. Authenticate the capability
-- and exact server/session/block, then use the independently retained GUI binding
-- for explicit client.perform calls. No public dispatch(ctx, ...) exists.
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
            local ok, value = pcall(completion, args.message)
            if not ok then
                fail("invalid completion JSON")
            end
            -- Lookup without yielding or consuming another menu's entries.
            for _, handler in pairs(handlers) do
                if handler(ctx, value) then
                    return -- Never return a private/user callback result.
                end
            end
            fail("invalid completion authority or invocation state")
        end,
    })
end

-- One installed configuration generation. Candidate compilation must succeed
-- before replacement; composition calls invalidate on reload/reuse/teardown.
function M.new(template, is_current)
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
            entry.context, entry.activation, entry.origin, entry.snapshot, entry.capability = nil, nil, nil, nil, nil
            entry.receipt, entry.receipt_hash, entry.metadata = nil, nil, nil
        end
    end

    function registry.invalidate()
        closed = true
        for ticket in pairs(entries) do
            release(ticket)
        end
        template = nil
        handlers[registry] = nil
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
        if entries[ticket] ~= entry or (is_current and not is_current()) then
            fail("invalid, expired, cancelled or consumed invocation")
        end
        if not same_origin(ctx, entry.origin) or ctx.client_id ~= entry.activation.client_id then
            fail("action is not targeted at the invocation origin or its caller changed")
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
        local placements, placed_window, placed_layer = 0, nil, nil
        for _, window in ipairs(view.windows or {}) do
            for _, layer in ipairs(window.layers or {}) do
                for _, block in ipairs(layer.blocks or {}) do
                    if block.block_id == expected.block_id then
                        placements, placed_window, placed_layer = placements + 1, window, layer.layer_id
                        if placed_layer == nil then placed_layer = "" end
                    end
                end
            end
        end
        if placements ~= 1 or not resource(placed_window.window_id, "window")
            or placed_window.window_id ~= entry.window_id or placed_layer ~= entry.layer_id then
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
                if (client.connection_state ~= nil and type(client.connection_state) ~= "string")
                    or client.connection_state == "lost" or not plain(client.info) or client.info.kind ~= "app" then
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
        local status, status_err = rex.call("server.status", {})
        live(ticket, entry, ctx)
        local metadata = entry.metadata
        if status_err ~= nil or not plain(status) or not plain(status.config)
            or status.instance_id ~= metadata.server_instance or status.config.hash ~= metadata.config_hash
            or status.config.generation ~= metadata.config_generation then
            fail("origin server configuration changed")
        end
        return placed_window.window_id
    end

    -- Reserve BEFORE evaluation. The unforgeable table ticket stays private to
    -- the trusted launcher. It has no transport representation or bearer spelling.
    function registry.issue(ctx, target)
        local caller = activation(ctx)
        if not plain(target) or not resource(target.client_id, "client")
            or not resource(target.window_id, "window") or type(target.layer_id) ~= "string"
            or (target.layer_id ~= "" and not resource(target.layer_id, "layer"))
            or not plain(target.metadata) or type(target.expires) ~= "number" then
            fail("an exact GUI target is required")
        end
        local bound = {
            server = caller.server, session_id = caller.session_id, block_id = caller.block_id,
            client_id = target.client_id, origin = caller.origin,
        }
        local now = sweep()
        if count >= max_pending then
            fail("too many pending invocations")
        end
        local ticket = {}
        local entry = {
            context = ctx, activation = caller, origin = bound, connected_at = target.connected_at,
            expires = math.min(now + lifetime, target.expires), state = "evaluating",
            window_id = target.window_id, layer_id = target.layer_id, metadata = target.metadata,
        }
        entries[ticket], count = entry, count + 1
        local ok, records = pcall(function()
            entry.window_id = inspect(ticket, entry, ctx)
            local data, invocation = evaluate.evaluate(template)
            -- Invalidation may have released the reservation while user code
            -- yielded. Attach its new handle before checking, then release below.
            entry.invocation = invocation
            live(ticket, entry, ctx)
            local snapshot = {}
            for _, record in ipairs(data) do
                if record.action or record.callback_ref then
                    snapshot[record.keys] = {
                        callback_ref = record.callback_ref, confirm = record.confirm == true,
                        rename = inline_rename(record),
                    }
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
    function registry.bind(ticket, ctx, prepared)
        sweep()
        local entry = entries[ticket]
        if not entry or entry.state ~= "prepared" or not rawequal(ctx, entry.context) then
            fail("invalid invocation capability binding")
        end
        entry.state = "binding"
        local ok = pcall(function()
            live(ticket, entry, ctx)
            if not plain(prepared) or not capability(prepared.capability)
                or capabilities[prepared.capability] ~= nil or not capability(prepared.receipt_hash)
                or type(prepared.receipt) ~= "string" or prepared.receipt == "" or #prepared.receipt > 8388608
                or type(prepared.expires) ~= "number" then
                fail("invalid or duplicate invocation capability")
            end
            entry.expires = math.min(entry.expires, prepared.expires)
            inspect(ticket, entry, ctx)
            local value = prepared.capability
            entry.receipt, entry.receipt_hash = prepared.receipt, prepared.receipt_hash
            entry.capability, entry.state, entry.context = value, "open", nil
            capabilities[value] = ticket
        end)
        if not ok then
            release(ticket)
            fail("invocation capability binding failed")
        end
    end

    local function complete(ctx, value)
        local ticket = capabilities[value.capability]
        local entry = ticket and entries[ticket]
        if not entry then
            return false
        end
        -- Wrong origin/hash cannot claim an unrelated authentic invocation.
        if entry.state ~= "open" or not same_origin(ctx, entry.origin)
            or value.receipt_hash ~= entry.receipt_hash then
            fail("invalid completion authority or invocation state")
        end
        -- Claim BEFORE conditions or Rex calls yield/reenter. Do not remove from
        -- entries until done: in-progress work counts toward the bounded capacity
        -- and remains revocable by expiry/reload/local cancellation.
        local caller = activation(ctx)
        sweep()
        if entries[ticket] ~= entry or (is_current and not is_current()) then
            fail("invalid completion authority or invocation state")
        end
        entry.state = "claimed"
        -- A fresh CLI caller need not equal either the launch CLI or GUI target.
        entry.activation = caller
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
            if not record or record.callback_ref ~= selected.callback_ref or record.confirm ~= selected.confirmed
                or record.rename ~= (selected.text ~= nil) then
                fail("selection does not match the retained invocation")
            end
            -- The evaluator owns original closures/args/guards. Its resolver is
            -- also single-use; bridge checks above strengthen its older internal
            -- acceptance of key aliases, absent flags and empty callback refs.
            local retained = entry.invocation.resolve(selected.keys, selected.callback_ref, selected.confirmed)
            live(ticket, entry, ctx)
            stage = "origin validation failed"
            local before_focus = inspect(ticket, entry, ctx)
            if record.rename and before_focus ~= entry.window_id then
                fail("rename origin moved to another tab")
            end
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
            local current_window = inspect(ticket, entry, ctx, true)
            live(ticket, entry, ctx)
            if record.rename then
                -- Text supplies only a label. The original VM retains the
                -- action and exact session/window; never open an app dialog.
                stage = "rename origin changed"
                if current_window ~= entry.window_id then
                    fail("rename origin moved to another tab")
                end
                stage = "rename failed or uncertain"
                local result, rename_err
                if retained.action == "session.rename" and selected.keys == "wr" then
                    result, rename_err = rex.session.set_label({
                        session_id = entry.origin.session_id, label = selected.text,
                    })
                elseif retained.action == "client.tab.rename" and selected.keys == "tr" then
                    result, rename_err = rex.session.set_window_label({
                        session_id = entry.origin.session_id, window_id = entry.window_id, label = selected.text,
                    })
                else
                    fail("invalid retained rename action")
                end
                if rename_err ~= nil or not plain(result) then
                    fail("rename failed or uncertain")
                end
            elseif retained.callback then
                stage = "callback failed"
                retained.callback(callback_context(entry.origin, ctx), retained.args)
            elseif retained.action:find(".", 1, true) then
                -- Never queue on the transient CLI connection. The retained
                -- app binding is independent of both activation caller IDs.
                stage = "client action failed or uncertain"
                local result, perform_err = rex.call("client.perform", {
                    client_id = entry.origin.client_id,
                    session_id = entry.origin.session_id,
                    block_id = entry.origin.block_id,
                    action = retained.action,
                    args = retained.args,
                    timeout_ms = 3000,
                })
                if perform_err ~= nil or not plain(result) or result.performed_by ~= entry.origin.client_id then
                    fail("client action failed or uncertain")
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
        return true
    end

    -- Local launch-failure/dismissal cleanup, requiring the private table ticket.
    -- Remote dismissal uses the authenticated registered cancel message instead.
    function registry.cancel(ticket)
        if not closed then
            sweep()
            local entry = entries[ticket]
            if entry and entry.state ~= "claimed" then
                release(ticket)
            end
        end
    end

    handlers[registry] = complete
    return registry
end

return M
