-- Internal Rex-runtime loader. mapping.go and docs/mapping-format.md own the
-- configuration rules; this module only selects and snapshots raw Lua options.
-- The composition owner supplies resolved defaults and the raw-spec compiler.
local M = {}

local fields = { position = true, title = true, icon = true, icons = true, spec = true }
local positions = { top = true, bottom = true, left = true, right = true, center = true }

local function fail(message)
    error("rex_whichkey.config: " .. message, 0)
end

local function plain_table(value, path)
    if type(value) ~= "table" or getmetatable(value) ~= nil then
        fail(path .. " must be a plain table")
    end
end

-- Lua 5.1 has no utf8 library. Match Go's valid-UTF-8/no-control text rule,
-- including C1 controls, overlong encodings, surrogates and out-of-range values.
local function text(value, path)
    if type(value) ~= "string" then
        fail(path .. " must be a string")
    end
    local i = 1
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
            fail(path .. " must be valid UTF-8")
        end
        for j = 1, count - 1 do
            local continuation = value:byte(i + j)
            if not continuation or continuation < 128 or continuation > 191 then
                fail(path .. " must be valid UTF-8")
            end
            code = code * 64 + continuation - 128
        end
        if code < minimum or code > 1114111 or (code >= 55296 and code <= 57343) then
            fail(path .. " must be valid UTF-8")
        end
        if code < 32 or (code >= 127 and code <= 159) then
            fail(path .. " must not contain terminal control characters")
        end
        i = i + count
    end
end

local function list(value, path)
    plain_table(value, path)
    local count = 0
    for key in pairs(value) do
        if type(key) ~= "number" or key < 1 or key % 1 ~= 0 then
            fail(path .. " must be a dense ordered list")
        end
        count = count + 1
    end
    for i = 1, count do
        if rawget(value, i) == nil then
            fail(path .. " must be a dense ordered list")
        end
    end
end

local function options(value, path, required)
    if value == nil and not required then
        return
    end
    plain_table(value, path)
    for key in pairs(value) do
        if not fields[key] then
            fail(path .. " contains an unknown option")
        end
    end
    for _, key in ipairs({ "position", "title", "icon", "icons", "spec" }) do
        local item = rawget(value, key)
        if item == nil then
            if required then
                fail(path .. "." .. key .. " is required in resolved defaults")
            end
        elseif key == "position" then
            if type(item) ~= "string" or not positions[item] then
                fail(path .. ".position must be top, bottom, left, right or center")
            end
        elseif key == "title" or key == "icon" then
            text(item, path .. "." .. key)
        elseif key == "icons" then
            plain_table(item, path .. ".icons")
            for marker, label in pairs(item) do
                if marker ~= "breadcrumb" and marker ~= "group" then
                    fail(path .. ".icons contains an unknown marker")
                end
                text(label, path .. ".icons." .. marker)
            end
        elseif key == "spec" then
            list(item, path .. ".spec")
        end
    end
end

-- Copy only the winning values; functions retain their Rex-runtime identity and
-- are never called here. The mapping compiler owns mapping/args validation.
local function copy(value, path, visiting, copied)
    if type(value) ~= "table" then
        return value
    end
    plain_table(value, path)
    visiting = visiting or {}
    copied = copied or {}
    if visiting[value] then
        fail(path .. " contains a cyclic table")
    end
    if copied[value] then
        return copied[value]
    end
    local result = {}
    visiting[value], copied[value] = true, result
    for key, item in pairs(value) do
        if type(key) ~= "string" and type(key) ~= "number" then
            fail(path .. " contains an invalid table key")
        end
        result[key] = copy(item, path .. " entry", visiting, copied)
    end
    visiting[value] = nil
    return result
end

-- Imported and inline options share the Go-owned shape. Presence, not truthiness,
-- selects overrides. No executable spec is evaluated by resolution.
function M.resolve(defaults, imported, inline)
    options(defaults, "defaults", true)
    options(imported, "imported", false)
    options(inline, "inline", false)
    local result = {}
    for _, key in ipairs({ "position", "title", "icon", "icons", "spec" }) do
        local value = rawget(defaults, key)
        if imported ~= nil and rawget(imported, key) ~= nil then
            value = rawget(imported, key)
        end
        if inline ~= nil and rawget(inline, key) ~= nil then
            value = rawget(inline, key)
        end
        result[key] = copy(value, key)
    end
    -- An explicitly supplied icons table replaces the whole object; missing
    -- markers become empty, not inherited default markers.
    if result.icons.breadcrumb == nil then
        result.icons.breadcrumb = ""
    end
    if result.icons.group == nil then
        result.icons.group = ""
    end
    return result
end

-- Use Rex's server-side require context (including siblings of its init.lua).
-- Never rewrite package.path, infer the client's home, or load an absolute file.
function M.load(module_name)
    if type(module_name) ~= "string" or module_name == "" then
        fail("import must be a nonempty Lua module name")
    end
    if module_name == "rex_whichkey" or module_name:sub(1, 13) == "rex_whichkey." then
        fail("user configuration must not use the rex_whichkey integration namespace")
    end
    if module_name:sub(1, 1) == "." or module_name:sub(-1) == "." or module_name:find("..", 1, true) then
        fail("import must be a dotted Lua module name, not a file path")
    end
    for part in module_name:gmatch("[^.]+") do
        if not part:match("^[A-Za-z_][A-Za-z0-9_]*$") then
            fail("import must be a dotted Lua module name, not a file path")
        end
    end
    local ok, value = pcall(require, module_name)
    if not ok then
        error("rex_whichkey.config: cannot import " .. module_name .. ": " .. tostring(value), 0)
    end
    plain_table(value, "imported module " .. module_name)
    options(value, "imported module " .. module_name, false)
    return value
end

-- Internal composition API, not a new user-facing setup/launch protocol.
-- defaults is the resolved baseline from the defaults owner, including its spec.
-- prepare(candidate) must validate/compile raw mappings without evaluating their
-- dynamic strings/conditions or invoking callbacks. It must throw on invalid
-- mappings and return a non-nil/non-false compiled template on success.
function M.new(defaults, prepare)
    if type(prepare) ~= "function" then
        fail("a raw-spec compiler is required")
    end
    local baseline = M.resolve(defaults)
    local installed, busy
    local loader = {}

    -- Repeated setup starts from baseline, never from the prior installed spec.
    -- Imported options precede inline options. Errors leave installed untouched.
    function loader.setup(inline, module_name)
        if busy then
            fail("setup is already in progress")
        end
        busy = true
        local ok, candidate, compiled = pcall(function()
            local imported
            if module_name ~= nil then
                imported = M.load(module_name)
            end
            local resolved = M.resolve(baseline, imported, inline)
            -- A compiler retaining or mutating its input cannot mutate the
            -- loader's installed raw options or defaults through that input.
            local prepared = prepare(copy(resolved, "candidate"))
            if prepared == nil or prepared == false then
                fail("raw-spec compilation returned no template")
            end
            return resolved, prepared
        end)
        busy = false
        if not ok then
            error(candidate, 0)
        end
        installed = { config = candidate, compiled = compiled }
        return copy(candidate, "installed configuration"), compiled
    end

    -- No successful setup yet is distinct from an explicitly empty menu.
    -- Compiled templates are compiler-owned and must be treated as immutable.
    function loader.current()
        if installed == nil then
            return nil
        end
        return copy(installed.config, "installed configuration"), installed.compiled
    end

    return loader
end

return M
