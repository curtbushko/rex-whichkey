-- Internal consumer of mapping.go and docs/mapping-format.md. Evaluate only in
-- the originating Rex Lua runtime, never in Go or while compiling configuration.
local mapping = require("rex_whichkey.mapping")
local M = {}
local next_invocation = 0

local function fail(path, message)
    error("rex_whichkey.evaluate: " .. path .. ": " .. message, 0)
end

-- Compiler-owned templates contain functions, but their tables are immutable.
-- Detach tables before evaluating user code and before returning menu metadata.
local function copy(value)
    if type(value) ~= "table" then
        return value
    end
    local result = {}
    for key, item in pairs(value) do
        result[key] = copy(item)
    end
    return result
end

local function call(fn, path)
    local ok, value = pcall(fn)
    if not ok then
        -- An arbitrary thrown Lua value is not safe terminal-facing error text.
        fail(path, "evaluation failed")
    end
    return value
end

local function condition(value, path)
    if type(value) == "function" then
        value = call(value, path)
        if value == nil then
            return false
        end
    end
    if type(value) ~= "boolean" then
        fail(path, "condition must resolve to a boolean or nil")
    end
    return value
end

-- Returns data-only Go-shaped records and a separate, nonserializable invocation
-- handle. The dispatch owner must retain that handle with the origin/runtime and
-- installed configuration identity, and release it on replacement, dismissal or
-- errors. This is not a wire envelope, registry or user-facing setup API.
function M.evaluate(template)
    if next_invocation >= 9007199254740991 then
        fail("invocation", "callback identity space exhausted")
    end
    next_invocation = next_invocation + 1
    local identity = string.format("%.0f", next_invocation)
    local snapshot = copy(template)
    local records, selections, contributions = {}, {}, {}
    local callback_count = 0

    local function walk(node, inherited, inherited_guards)
        local effective = mapping.inherit(inherited, {})
        local guards = copy(inherited_guards)
        local own = node.attributes
        -- Evaluate a scope before its strings or children. Inherited conditions
        -- and strings are already resolved, so redraws/children never call them.
        if own.cond ~= nil then
            local guard = { value = own.cond, path = node.path .. ".cond" }
            if not condition(guard.value, guard.path) then
                return
            end
            effective.cond = true
            guards[#guards + 1] = guard
        end
        for _, key in ipairs({ "desc", "group", "icon" }) do
            local value = own[key]
            if value ~= nil then
                if type(value) == "function" then
                    value = call(value, node.path .. "." .. key)
                end
                if value ~= nil and type(value) ~= "string" then
                    fail(node.path .. "." .. key, "must resolve to a string or nil")
                end
                -- A dynamic nil removes inheritance; it is not a fallback.
                effective[key] = value
            end
        end
        -- provider is the Go-owned data marker, never a target or callback.
        -- The compiler permits it only on a mapping, not an inherited scope.
        for _, key in ipairs({ "hidden", "confirm", "args", "provider" }) do
            if own[key] ~= nil then
                effective[key] = own[key]
            end
        end
        if node.children then
            for _, child in ipairs(node.children) do
                walk(child, effective, guards)
            end
            return
        end
        local record = mapping.inherit(effective, { keys = node.keys })
        if type(node.target) == "function" then
            callback_count = callback_count + 1
            record.callback_ref = "cb:" .. identity .. ":" .. string.format("%.0f", callback_count)
        elseif node.target ~= nil then
            record.action = node.target
        end
        records[#records + 1] = record
        contributions[node.keys] = contributions[node.keys] or {}
        for _, guard in ipairs(guards) do
            contributions[node.keys][#contributions[node.keys] + 1] = guard
        end
        if node.target ~= nil then
            -- coalesce below rejects duplicate targets before publishing a handle.
            selections[node.keys] = { target = node.target, args = copy(effective.args or {}) }
        end
    end

    for _, node in ipairs(snapshot.children) do
        walk(node, {}, {})
    end
    -- All active conflicts, resolved text safety and argument data are checked
    -- before anything escapes. Hidden records deliberately remain in this list.
    local evaluated = mapping.coalesce(records)
    for _, record in ipairs(evaluated) do
        local selection = selections[record.keys]
        if selection then
            -- Use final coalesced metadata for confirmation and callback identity,
            -- not caller-supplied fields or a partially evaluated source record.
            selection.record = copy(record)
            selection.guards = contributions[record.keys]
        end
    end

    local state = "open"
    local invocation = {}

    function invocation.release()
        state = "closed"
        selections, contributions, records, snapshot, evaluated = nil, nil, nil, nil, nil
    end

    -- Resolve once, without executing an action. The dispatch owner authenticates
    -- origin/config lifetime and restores focus before using the retained target.
    -- Only keys, the opaque callback reference (if any), and a boolean y-confirmed
    -- flag are accepted; replacement action names, code and args are never input.
    function invocation.resolve(keys, callback_ref, confirmed)
        if state ~= "open" then
            fail("selection", "invocation is closed or already resolving")
        end
        state = "resolving"
        local ok, result = pcall(function()
            local _, canonical = mapping.parse_keys(keys, "selection.keys")
            local selection = selections[canonical]
            if not selection then
                fail("selection", "no active action at these keys")
            end
            local record = selection.record
            if callback_ref == "" then
                callback_ref = nil
            end
            if callback_ref ~= record.callback_ref then
                fail("selection", "stale or mismatched callback reference")
            end
            if confirmed ~= nil and type(confirmed) ~= "boolean" then
                fail("selection", "confirmation must be a boolean")
            end
            if record.confirm == true and confirmed ~= true then
                fail("selection", "confirmation is required")
            end
            -- Recheck every contributing scope/declaration, including metadata
            -- coalesced onto this action. A new false/nil/error aborts dispatch.
            local checked = {}
            for _, guard in ipairs(selection.guards) do
                if not checked[guard.path] then
                    checked[guard.path] = true
                    if not condition(guard.value, guard.path) then
                        fail(guard.path, "selection is no longer enabled")
                    end
                end
            end
            if state ~= "resolving" then
                fail("selection", "invocation was released during condition evaluation")
            end
            local retained = { args = copy(selection.args), confirm = record.confirm == true }
            if record.callback_ref ~= nil then
                retained.callback_ref = record.callback_ref
                retained.callback = selection.target
            else
                retained.action = record.action
            end
            return retained
        end)
        -- Consume even failed selections; conditions cannot reenter dispatch or
        -- retry an uncertain operation. No callback/global registry outlives this.
        invocation.release()
        if not ok then
            error(result, 0)
        end
        return result
    end

    return copy(evaluated), invocation
end

return M
