-- Lua adaptation of mapping.go and docs/mapping-format.md, not a wire schema.
-- compile(spec) snapshots raw notation without calling any user function. Its
-- internal tree retains local attributes so the evaluator can evaluate a scope
-- once, inherit its resolved values, and skip disabled scopes before children.
-- coalesce(records) accepts evaluated Go-shaped records from that evaluator.
-- Callback registration, evaluation, JSON encoding and dispatch belong elsewhere.
local M = {}

local attributes = {
    desc = true,
    group = true,
    icon = true,
    cond = true,
    hidden = true,
    confirm = true,
    args = true,
}

local function fail(path, message)
    error("rex_whichkey.mapping: " .. path .. ": " .. message, 0)
end

local function plain(value, path)
    if type(value) ~= "table" or getmetatable(value) ~= nil then
        fail(path, "must be a plain table")
    end
end

-- Decode explicitly for Rex runtimes without Lua's optional utf8 library.
local function characters(value, path, controls)
    if type(value) ~= "string" then
        fail(path, "must be a string")
    end
    local result, i = {}, 1
    while i <= #value do
        local first = value:byte(i)
        local count, code, minimum
        if first < 128 then
            count, code, minimum = 1, first, 0
        elseif first >= 194 and first <= 223 then
            count, code, minimum = 2, first - 192, 128
        elseif first >= 224 and first <= 239 then
            count, code, minimum = 3, first - 224, 2048
        elseif first >= 240 and first <= 244 then
            count, code, minimum = 4, first - 240, 65536
        else
            fail(path, "must be valid UTF-8")
        end
        for j = 1, count - 1 do
            local byte = value:byte(i + j)
            if not byte or byte < 128 or byte > 191 then
                fail(path, "must be valid UTF-8")
            end
            code = code * 64 + byte - 128
        end
        if code < minimum or code > 1114111 or (code >= 55296 and code <= 57343) then
            fail(path, "must be valid UTF-8")
        end
        if not controls and (code < 32 or (code >= 127 and code <= 159)) then
            fail(path, "must not contain terminal control characters")
        end
        result[#result + 1] = value:sub(i, i + count - 1)
        i = i + count
    end
    return result
end

local function text(value, path)
    characters(value, path, false)
    return value
end

local function nonblank(value)
    -- Go strings.TrimSpace includes Unicode White_Space, not just ASCII %s.
    local spaces = {
        [" "] = true,
        ["\194\160"] = true,
        ["\225\154\128"] = true,
        ["\226\128\128"] = true,
        ["\226\128\129"] = true,
        ["\226\128\130"] = true,
        ["\226\128\131"] = true,
        ["\226\128\132"] = true,
        ["\226\128\133"] = true,
        ["\226\128\134"] = true,
        ["\226\128\135"] = true,
        ["\226\128\136"] = true,
        ["\226\128\137"] = true,
        ["\226\128\138"] = true,
        ["\226\128\168"] = true,
        ["\226\128\169"] = true,
        ["\226\128\175"] = true,
        ["\226\129\159"] = true,
        ["\227\128\128"] = true,
    }
    for _, char in ipairs(characters(value, "text", false)) do
        if not spaces[char] then
            return true
        end
    end
    return false
end

local function literal(token)
    if token == " " then
        return "<Space>"
    elseif token == "<" then
        return "<lt>"
    end
    return token
end

-- Tokens are Unicode code points, not bytes, graphemes or display cells.
function M.parse_keys(sequence, path)
    path = path or "keys"
    characters(sequence, path, false)
    if sequence == "" then
        fail(path, "must be a nonempty full key sequence")
    end
    local tokens, canonical, i = {}, {}, 1
    while i <= #sequence do
        local token
        if sequence:sub(i, i) == "<" then
            local ending = sequence:find(">", i, true)
            if not ending then
                fail(path, "unclosed key notation; use <lt> for literal <")
            end
            local notation = sequence:sub(i, ending)
            local lower = notation:lower()
            if lower == "<space>" then
                token = " "
            elseif lower == "<lt>" then
                token = "<"
            else
                fail(path, "unsupported key notation " .. notation .. "; only <Space> and <lt> are supported")
            end
            i = ending + 1
        else
            local byte = sequence:byte(i)
            local size = byte < 128 and 1 or (byte < 224 and 2 or (byte < 240 and 3 or 4))
            token = sequence:sub(i, i + size - 1)
            i = i + size
        end
        tokens[#tokens + 1], canonical[#canonical + 1] = token, literal(token)
    end
    return tokens, table.concat(canonical)
end

local function dense(value, path, named)
    plain(value, path)
    local count = 0
    for key in pairs(value) do
        if type(key) == "number" then
            if key < 1 or key % 1 ~= 0 or key == math.huge then
                fail(path, "positional values must form a dense ordered list")
            end
            count = count + 1
        elseif not named or not attributes[key] then
            fail(path, "unsupported attribute or list key " .. tostring(key))
        end
    end
    for i = 1, count do
        if rawget(value, i) == nil then
            fail(path, "sparse positional list at index " .. i)
        end
    end
    return count
end

-- Snapshot data-only args. An empty table is an object; the encoder must retain
-- that distinction at args and nested object positions, rather than emit null.
local function data(value, path, visiting, object)
    local kind = type(value)
    if kind == "string" then
        characters(value, path, true)
        return value
    elseif kind == "boolean" then
        return value
    elseif kind == "number" then
        if value ~= value or value == math.huge or value == -math.huge then
            fail(path, "numbers must be finite")
        end
        return value
    elseif kind ~= "table" then
        fail(path, "args must contain only data, not " .. kind)
    end
    plain(value, path)
    if visiting[value] then
        fail(path, "cyclic argument table")
    end
    visiting[value] = true
    local result, count, strings, numbers = {}, 0, false, false
    for key, item in pairs(value) do
        if type(key) == "string" then
            characters(key, path .. " object key", true)
            strings = true
        elseif type(key) == "number" and key >= 1 and key % 1 == 0 and key ~= math.huge then
            numbers, count = true, count + 1
        else
            fail(path, "invalid argument object key or array index")
        end
        result[key] = data(item, path .. "[" .. tostring(key) .. "]", visiting, false)
    end
    if numbers and (strings or object) then
        fail(path, "must be a string-keyed object, not a mixed table or array")
    end
    if numbers then
        for i = 1, count do
            if rawget(value, i) == nil then
                fail(path, "argument arrays must be dense")
            end
        end
    end
    visiting[value] = nil
    return result
end

local function args(value, path)
    plain(value, path)
    return data(value, path, {}, true)
end

function M.inherit(parent, child)
    local result = {}
    for key, value in pairs(parent or {}) do
        result[key] = value
    end
    for key, value in pairs(child or {}) do
        result[key] = value
    end
    return result
end

local function raw_attributes(entry, path)
    local result = {}
    for key in pairs(attributes) do
        local value = rawget(entry, key)
        if value ~= nil then
            local kind = type(value)
            if key == "desc" or key == "group" or key == "icon" then
                if kind ~= "function" then
                    text(value, path .. "." .. key)
                end
            elseif key == "cond" then
                if kind ~= "boolean" and kind ~= "function" then
                    fail(path .. ".cond", "must be a boolean or function")
                end
            elseif key == "args" then
                value = args(value, path .. ".args")
            elseif kind ~= "boolean" then
                fail(path .. "." .. key, "must be a boolean")
            end
            result[key] = value
        end
    end
    return result
end

function M.compile(spec)
    local visiting = {}
    local function entry(value, path, inherited)
        plain(value, path)
        if visiting[value] then
            fail(path, "cyclic mapping scope")
        end
        visiting[value] = true
        local count = dense(value, path, true)
        local own = raw_attributes(value, path)
        local effective = M.inherit(inherited, own)
        local node = { path = path, attributes = own }
        local first = rawget(value, 1)
        if type(first) == "string" then
            if count > 2 then
                fail(
                    path,
                    "mapping accepts only keys and an optional action/callback; children require an attribute scope"
                )
            end
            node.tokens, node.keys = M.parse_keys(first, path .. "[1]")
            node.target = rawget(value, 2)
            if node.target ~= nil and type(node.target) ~= "function" then
                text(node.target, path .. "[2]")
                if not nonblank(node.target) then
                    fail(path .. "[2]", "named action must not be blank")
                end
            end
            -- Dynamic values and conditions are resolved by the evaluator before
            -- active-key conflict checks; do not execute them during setup.
            if node.target ~= nil then
                if type(effective.group) == "string" then
                    fail(
                        path,
                        "action "
                            .. node.keys
                            .. " cannot declare or inherit group; move group to a separate mapping"
                    )
                end
                if type(effective.desc) ~= "function" then
                    if effective.desc == nil or not nonblank(effective.desc) then
                        fail(path, "action " .. node.keys .. " requires a nonblank desc (possibly inherited)")
                    end
                end
            end
        else
            node.children = {}
            for i = 1, count do
                node.children[i] = entry(rawget(value, i), path .. "[" .. i .. "]", effective)
            end
        end
        visiting[value] = nil
        return node
    end
    local count = dense(spec, "spec", false)
    local template = { children = {} }
    visiting[spec] = true
    for i = 1, count do
        template.children[i] = entry(rawget(spec, i), "spec[" .. i .. "]", {})
    end
    -- Reject conflicts already knowable at setup, but never guess the result of
    -- a dynamic condition. The invocation's coalesce pass remains authoritative.
    local known = {}
    local function check(node, inherited, enabled)
        local effective = M.inherit(inherited, node.attributes)
        local certain = enabled and effective.cond ~= false and type(effective.cond) ~= "function"
        if node.children then
            -- A child override cannot establish that an unknown/disabled scope
            -- will be walked. Only unconditional active paths claim keys here.
            for _, child in ipairs(node.children) do
                check(child, effective, certain)
            end
        elseif certain then
            known[#known + 1] = {
                keys = node.keys,
                tokens = node.tokens,
                target = node.target,
                path = node.path,
                group = effective.group,
            }
        end
    end
    for _, node in ipairs(template.children) do
        check(node, {}, true)
    end
    local claims, prefixes, groups = {}, {}, {}
    for _, node in ipairs(known) do
        local prefix = ""
        for i, token in ipairs(node.tokens) do
            prefix = prefix .. literal(token)
            if i < #node.tokens then
                prefixes[prefix] = node.path
            end
        end
        if node.target ~= nil then
            if claims[node.keys] then
                fail(
                    node.path,
                    "duplicate action mapping " .. node.keys .. "; first declared at " .. claims[node.keys]
                )
            end
            claims[node.keys] = node.path
        end
        -- A dynamic group might resolve to absence; defer that case.
        if type(node.group) == "string" then
            groups[node.keys] = node.path
        end
    end
    for _, node in ipairs(known) do
        if claims[node.keys] and (prefixes[node.keys] or groups[node.keys]) then
            fail(
                node.path,
                "group/action conflict at "
                    .. node.keys
                    .. "; related declaration at "
                    .. (prefixes[node.keys] or groups[node.keys])
            )
        end
    end
    return template
end

-- The evaluator calls this only after applying scope inheritance, resolving all
-- strings/conditions and assigning invocation-owned opaque callback references.
-- Nothing executable is retained in the returned Go-shaped records.
function M.coalesce(records)
    local count = dense(records, "evaluated mappings", false)
    local result, nodes, origins, targets = {}, {}, {}, {}
    for i = 1, count do
        local path = "evaluated mappings[" .. i .. "]"
        local source = records[i]
        plain(source, path)
        local tokens, keys = M.parse_keys(source.keys, path .. ".keys")
        local record = { keys = keys }
        for key, value in pairs(source) do
            if key == "keys" then
                -- Canonicalized above.
            elseif
                key == "desc"
                or key == "group"
                or key == "icon"
                or key == "action"
                or key == "callback_ref"
            then
                record[key] = text(value, path .. "." .. key)
            elseif key == "hidden" or key == "confirm" or key == "cond" then
                if type(value) ~= "boolean" then
                    fail(path .. "." .. key, "must be an evaluated boolean")
                end
                record[key] = value
            elseif key == "args" then
                record.args = args(value, path .. ".args")
            else
                fail(path, "unknown evaluated mapping field " .. tostring(key))
            end
        end
        local actionable = (record.action ~= nil and record.action ~= "")
            or (record.callback_ref ~= nil and record.callback_ref ~= "")
        if record.action and record.action ~= "" and not nonblank(record.action) then
            fail(path, "action must not be blank")
        end
        if record.callback_ref and record.callback_ref ~= "" and not nonblank(record.callback_ref) then
            fail(path, "callback_ref must not be blank")
        end
        if record.action and record.action ~= "" and record.callback_ref and record.callback_ref ~= "" then
            fail(path, "action and callback_ref are mutually exclusive")
        end
        if actionable then
            if record.group ~= nil or not record.desc or not nonblank(record.desc) then
                fail(path, "action requires a nonblank desc and cannot declare group")
            end
        else
            -- Scope action-only attributes must not leak into a Go group record.
            record.args, record.confirm = nil, nil
        end
        if record.cond ~= false then
            local prefix = ""
            for _, token in ipairs(tokens) do
                prefix = prefix .. literal(token)
                if not nodes[prefix] then
                    nodes[prefix] = { keys = prefix }
                    result[#result + 1] = nodes[prefix]
                end
            end
            if actionable and targets[keys] then
                fail(path, "duplicate action mapping " .. keys .. "; first declared at " .. targets[keys])
            end
            local node = nodes[keys]
            if (actionable and node.group ~= nil) or (record.group ~= nil and targets[keys]) then
                fail(
                    path,
                    "group/action conflict at "
                        .. keys
                        .. "; earlier declaration at "
                        .. (origins[keys] or keys)
                )
            end
            if actionable then
                targets[keys] = path
            end
            -- Only supplied metadata overrides; false and empty strings survive.
            for key, value in pairs(record) do
                if key ~= "action" and key ~= "callback_ref" then
                    node[key] = value
                elseif value ~= "" then
                    node[key] = value
                end
            end
            origins[keys] = origins[keys] or path
        end
    end
    -- Prefix conflicts are order independent, including implicit groups.
    for _, record in ipairs(result) do
        if targets[record.keys] and (record.group ~= nil or not record.desc or not nonblank(record.desc)) then
            fail(
                origins[record.keys],
                "coalesced action " .. record.keys .. " requires a nonblank desc and no group"
            )
        end
        local tokens = M.parse_keys(record.keys)
        local prefix = ""
        for i = 1, #tokens - 1 do
            prefix = prefix .. literal(tokens[i])
            if targets[prefix] then
                fail(targets[prefix], "action " .. prefix .. " is also a submenu prefix of " .. record.keys)
            end
        end
    end
    return result
end

return M
