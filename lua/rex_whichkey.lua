-- Go owns formats, discovery and startup. This module keeps configuration,
-- callbacks and invocation authority in the original persistent server VM.
local function fail(message)
    error("rex_whichkey: " .. message, 0)
end

if ... ~= "rex_whichkey" then
    fail("helper/layout discovery failed; require the normal rex_whichkey namespace")
end

-- Only the opted-in normal file-loader namespace is supported. A source label
-- by itself does not authenticate arbitrary code or a hostile same-user loader.
local source = debug.getinfo(1, "S").source
local function file_source(label, suffix)
    if type(label) ~= "string" or #label > 4097 or label:sub(1, 1) ~= "@"
        or label:find("%c") or label:find("\\") or label:match("^@%a:") then
        fail("helper/layout discovery failed; install the normal integration assets")
    end
    local path = label:sub(2)
    for part in path:gmatch("[^/]+") do
        if part == ".." then
            fail("helper/layout discovery failed; install the normal integration assets")
        end
    end
    if path ~= suffix and path:sub(-#suffix - 1) ~= "/" .. suffix then
        fail("helper/layout discovery failed; install the normal integration assets")
    end
    return path, path:sub(1, #path - #suffix)
end
local _, namespace = file_source(source, "rex_whichkey.lua")
local file_loader = package.loaders and package.loaders[2]
local function native_loader(value)
    if type(value) ~= "function" then return false end
    local info = debug.getinfo(value, "S")
    -- Lua's native standard loader, not a user Lua replacement. Some Lua 5.1
    -- hosts label host-language functions G rather than C; neither is a chunk.
    return info and (info.what == "C" or info.what == "G")
end
if not native_loader(file_loader) or not native_loader(package.loaders[1])
    or debug.getinfo(2, "f").func ~= require or package.loaded["rex_whichkey.runtime.anchor"] ~= nil then
    fail("helper/layout discovery failed; use one normal file-backed integration namespace")
end
for name in pairs(package.preload) do
    if type(name) == "string" and (name == "rex_whichkey" or name:match("^rex_whichkey%.")) then
        fail("helper/layout discovery failed; preload integration modules are unsupported")
    end
end
-- Require must actually be able to use its standard file slot for these
-- modules. A custom later loader with an @ label for a missing file is not a
-- supported anchor. Compiled candidates are inspected, never executed here.
local function file_candidate(name, label)
    local ok, candidate = pcall(file_loader, name)
    if not ok or type(candidate) ~= "function" or debug.getinfo(candidate, "S").source ~= label then
        fail("helper/layout discovery failed; synthetic integration loaders are unsupported")
    end
end
file_candidate("rex_whichkey", source)
local anchor = require("rex_whichkey.runtime.anchor")
if type(anchor) ~= "table" or getmetatable(anchor) ~= nil then
    fail("helper/layout discovery failed; install the normal integration anchor")
end
file_candidate("rex_whichkey.runtime.anchor", anchor.source)
local anchor_path, anchor_namespace = file_source(anchor.source, "rex_whichkey/runtime/anchor.lua")
if anchor.layout_revision ~= 1 or anchor_namespace ~= namespace then
    fail("helper/layout discovery failed; use one normal integration asset tree")
end
local helper = anchor_path:sub(1, #anchor_path - #"anchor.lua") .. "rex-whichkey"
local config = require("rex_whichkey.config")
local mapping = require("rex_whichkey.mapping")
local dispatch = require("rex_whichkey.dispatch")
if not pcall(dispatch.validate_text, anchor_path) then
    fail("helper/layout discovery failed; invalid integration anchor path")
end
-- Reject shadowed/pinned modules from a different logical namespace; Go checks
-- canonical asset ownership/symlinks at each launch. Do not scan for replacements.
for name, module in pairs(package.loaded) do
    if type(name) == "string" and name:match("^rex_whichkey%.") and type(module) == "table" then
        local module_source
        for _, value in pairs(module) do
            if type(value) == "function" then
                local label = debug.getinfo(value, "S").source
                local _, root = file_source(label, name:gsub("%.", "/") .. ".lua")
                if root ~= namespace or (module_source and module_source ~= label) then
                    fail("helper/layout discovery failed; integration modules resolve from different trees")
                end
                module_source = label
            end
        end
        if module_source then file_candidate(name, module_source) end
    end
end

local M = {}
local completion_action
local launch_names = {}
local generation_serial = 0
local features = { "original-vm", "separate-caller-target", "single-use", "generation-bound", "retained-menu",
    "public-file-relay", "completion-receipt-v2", "selection-text-resource-rename" }
local small_limit, full_limit = 65536, 8388608
local compatibility = "runtime compatibility failed; update and load the integration through the normal configuration lifecycle; not retried"
local origin_failure = "origin or configuration changed; relaunch from the original Rex context; not retried"

local function plain(value)
    return type(value) == "table" and getmetatable(value) == nil
end
local function name_valid(name)
    return type(name) == "string" and #name <= 256 and name:match("^[a-z0-9_]+$") ~= nil
end
local function text_valid(value, limit)
    return type(value) == "string" and #value <= (limit or 1024) and pcall(dispatch.validate_text, value)
end
local function hash(value)
    return type(value) == "string" and #value == 64 and not value:find("[^0-9a-f]")
end
local function integer(value)
    return type(value) == "number" and value >= 0 and value <= 9007199254740991 and value == math.floor(value)
end
local function resource(value, kind)
    if type(value) ~= "string" then return false end
    local prefix, a, b, c, d, e = value:match("^([a-z]+):(%x+)%-(%x+)%-(%x+)%-(%x+)%-(%x+)$")
    return prefix == kind and #a == 8 and #b == 4 and #c == 4 and #d == 4 and #e == 12
end
local function encode(value, limit)
    local text, err = rex.json.encode(value)
    if err ~= nil or type(text) ~= "string" or #text > (limit or full_limit) then
        fail("cannot encode bounded invocation data")
    end
    return text
end
local function decode(text, schema, limit)
    local value = dispatch.decode_record(text, schema, limit)
    for key in pairs(schema) do
        if value[key] == nil then fail(compatibility) end
    end
    return value
end
-- Rex encodes empty tables as objects; spec must remain an ordered JSON array.
local function configuration(options, records)
    local values = {}
    for i, record in ipairs(records) do values[i] = encode(record) end
    local display = encode({ position = options.position, title = options.title, icon = options.icon, icons = options.icons })
    local result = display:sub(1, -2) .. ',"spec":[' .. table.concat(values, ",") .. "]}"
    if #result > 4194304 then fail("configuration exceeds the Go transport bound") end
    return result
end
local function freeze(ctx)
    if not plain(ctx) or not resource(ctx.session_id, "session") or not resource(ctx.block_id, "block")
        or not resource(ctx.client_id, "client") then fail(origin_failure) end
    local kind = type(ctx.server)
    if (kind ~= "nil" and kind ~= "string" and kind ~= "table" and kind ~= "userdata") or ctx.server == "" then
        fail(origin_failure)
    end
    return { server = ctx.server, session_id = ctx.session_id, block_id = ctx.block_id, client_id = ctx.client_id }
end
local function unchanged(ctx, actual)
    return plain(ctx) and rawequal(ctx.server, actual.server) and ctx.session_id == actual.session_id
        and ctx.block_id == actual.block_id and ctx.client_id == actual.client_id
end
local function metadata(check)
    local status, err = rex.call("server.status", {})
    check()
    if err ~= nil or not plain(status) or not plain(status.config) or not text_valid(status.instance_id)
        or not text_valid(status.config.hash) or not integer(status.config.generation) then fail(origin_failure) end
    return { server_instance = status.instance_id, config_generation = status.config.generation, config_hash = status.config.hash }
end
local function same_metadata(a, b)
    return a.server_instance == b.server_instance and a.config_generation == b.config_generation and a.config_hash == b.config_hash
end
local function inspect(ctx, actual, target, connected_at, check)
    local status = metadata(check)
    local view, err = rex.session.view({ session_id = actual.session_id })
    check()
    if err ~= nil or not plain(view) or view.session_id ~= actual.session_id or not plain(view.windows) then fail(origin_failure) end
    local count, window_id, layer_id = 0, nil, nil
    for _, window in ipairs(view.windows) do
        if not plain(window) or not plain(window.layers) then fail(origin_failure) end
        for _, layer in ipairs(window.layers) do
            if not plain(layer) or not plain(layer.blocks) then fail(origin_failure) end
            for _, block in ipairs(layer.blocks) do
                if not plain(block) then fail(origin_failure) end
                if block.block_id == actual.block_id then
                    count, window_id, layer_id = count + 1, window.window_id, layer.layer_id
                    if layer_id == nil then layer_id = "" end
                end
            end
        end
    end
    if count ~= 1 or not resource(window_id, "window") or (layer_id ~= "" and not resource(layer_id, "layer")) then
        fail(origin_failure)
    end
    local clients, client_err = rex.session.list_clients({ session_id = actual.session_id })
    check()
    if client_err ~= nil or not plain(clients) or not plain(clients.clients) then fail(origin_failure) end
    local matches, generation = 0, nil
    for _, client in ipairs(clients.clients) do
        if not plain(client) then fail(origin_failure) end
        if client.client_id == target then
            matches, generation = matches + 1, client.connected_at
            if (client.connection_state ~= nil and type(client.connection_state) ~= "string")
                or client.connection_state == "lost" or not plain(client.info) or client.info.kind ~= "app" then fail(origin_failure) end
        end
    end
    if matches ~= 1 or not text_valid(generation) or (connected_at ~= nil and generation ~= connected_at) then fail(origin_failure) end
    return { session_id = actual.session_id, block_id = actual.block_id, window_id = window_id, layer_id = layer_id,
        caller_id = actual.client_id, client_id = target, connected_at = generation }, status
end
local function same_origin(a, b)
    for key, value in pairs(a) do if b[key] ~= value then return false end end
    return true
end
local function consumer(name, entry, status)
    if not entry.registry or not completion_action then fail(compatibility) end
    local advertised = {}
    for i, feature in ipairs(features) do advertised[i] = feature end
    return { protocol_min = 1, protocol_max = 1, consumer_build = "rex-whichkey-lua-runtime-1-completion-2",
        features = advertised, launch_action = name, completion_action = completion_action,
        menu_generation = entry.generation, layout_revision = 1, server_instance = status.server_instance,
        config_generation = status.config_generation, config_hash = status.config_hash }
end
local function compatible(reply)
    if not integer(reply.protocol_min) or not integer(reply.protocol_max) or reply.protocol_min < 1
        or reply.protocol_min > 1 or reply.protocol_max < 1 or reply.layout_revision ~= 1
        or not text_valid(reply.producer_build) or not plain(reply.features) then fail(compatibility) end
    local seen = {}
    for _, feature in ipairs(reply.features) do
        local known = false
        for _, expected in ipairs(features) do if feature == expected then known = true end end
        if not known or seen[feature] then fail(compatibility) end
        seen[feature] = true
    end
    for _, required in ipairs(features) do if not seen[required] then fail(compatibility) end end
end
local function command(mode, input, schema, limit)
    local result, err = rex.exec({ helper, mode }, {
        stdin = encode(input, mode == "runtime-describe" and small_limit or full_limit), timeout = 3, stderr = false,
    })
    if err ~= nil or not plain(result) or result.ok ~= true or result.code ~= 0 or result.timed_out then fail(compatibility) end
    return decode(result.stdout, schema, limit)
end
local prepared_schema = { version = "integer", capability = "string", executable = "string", server = "string",
    receipt = "string", receipt_hash = "string", workspace = "string", expires = "integer" }
local function valid_preparation(value)
    if value.version ~= 1 or not hash(value.capability) or not hash(value.receipt_hash)
        or not text_valid(value.server) or not text_valid(value.executable, 4096) or value.executable:sub(1, 1) ~= "/"
        or not text_valid(value.workspace, 4096) or value.workspace:sub(1, 1) ~= "/"
        or value.executable ~= value.workspace .. "/runtime-image" or value.receipt == "" or #value.receipt > full_limit
        or not integer(value.expires) then fail(compatibility) end
end

function M.register_completion(name)
    if completion_action or not name_valid(name) or launch_names[name] or name == "whichkey_runtime" then
        fail("completion action is invalid, duplicate or reserved")
    end
    dispatch.register_completion(name)
    completion_action = name
end

function M.new(defaults, server_endpoint, execution)
    if execution == nil then execution = {} end
    if not plain(execution) then fail("execution options must be a table") end
    if execution.helper ~= nil or execution.rex_bin_dir ~= nil then
        fail("remove obsolete execution.helper/execution.rex_bin_dir wiring; install and load the normal integration assets")
    end
    for _ in pairs(execution) do fail("unsupported execution option") end
    if server_endpoint ~= nil and type(server_endpoint) ~= "string" then
        fail("the optional server-owned endpoint override must be a string")
    end
    local loader = config.new(defaults, function(candidate) return mapping.compile(candidate.spec) end)
    local entry = {}
    local menu = {}
    function menu.setup(inline, module_name)
        if generation_serial >= 9007199254740991 then fail("menu generation identity space exhausted") end
        local candidate, template = loader.setup(inline, module_name)
        local replacement
        replacement = dispatch.new(template, function() return entry.registry == replacement end)
        local previous = entry.registry
        generation_serial = generation_serial + 1
        entry.registry, entry.options, entry.generation = replacement, candidate, "menu:" .. string.format("%.0f", generation_serial)
        if previous then previous.invalidate() end
    end
    function menu.register_launch(name)
        if name == nil then name = "whichkey" end
        if not name_valid(name) or entry.name or launch_names[name] or name == completion_action or name == "whichkey_runtime" then
            fail("launch action is invalid, duplicate or reserved")
        end
        rex.action({ name = name, title = "Which key",
            args = { runtime = "string?", client_id = "string?", connected_at = "string?" },
            run = function(ctx, args)
                -- Public Go emits these exact three fields. The locator is an
                -- opaque string; never open its files or substitute a GUI here.
                if not plain(args) then fail(compatibility) end
                for key in pairs(args) do
                    if key ~= "runtime" and key ~= "client_id" and key ~= "connected_at" then fail(compatibility) end
                end
                if next(args) == nil then menu.launch(ctx) else menu.launch(ctx, args) end
            end,
        })
        entry.name, launch_names[name] = name, entry
    end
    function menu.launch(ctx, args)
        local stage = origin_failure
        local ticket, prepared, attempted_start
        local registry, options, menu_generation = entry.registry, entry.options, entry.generation
        local ok = pcall(function()
            if not registry or not entry.name or not completion_action then fail(compatibility) end
            local actual = freeze(ctx) -- Actual caller, never the nominated GUI.
            local started = os.time()
            if not integer(started) then fail(origin_failure) end
            local deadline = started + 30
            local function check()
                local now = os.time()
                if entry.registry ~= registry or entry.options ~= options or entry.generation ~= menu_generation
                    or not unchanged(ctx, actual) or not integer(now) or now < started or now >= deadline then fail(origin_failure) end
            end
            local public = args ~= nil
            local target, connected_at, expected_hash = actual.client_id, nil, nil
            if public then
                if not plain(args) then fail(compatibility) end
                local count = 0
                for key in pairs(args) do
                    if key ~= "runtime" and key ~= "client_id" and key ~= "connected_at" then fail(compatibility) end
                    count = count + 1
                end
                if count ~= 3 or type(args.runtime) ~= "string" or #args.runtime > small_limit
                    or not resource(args.client_id, "client") or not text_valid(args.connected_at) then fail(compatibility) end
                -- Validate locator syntax without reading nominated files; keep
                -- its original bytes for Go's authenticated local preflight.
                local locator = decode(args.runtime, { version = "integer", workspace = "string",
                    request_sha256 = "string", receipt_hash = "string" }, small_limit)
                if locator.version ~= 1 or not hash(locator.request_sha256) or not hash(locator.receipt_hash)
                    or not text_valid(locator.workspace, 4096) or locator.workspace:sub(1, 1) ~= "/" then fail(compatibility) end
                encode(args, small_limit)
                target, connected_at, expected_hash = args.client_id, args.connected_at, locator.receipt_hash
            end
            local origin, status = inspect(ctx, actual, target, connected_at, check)
            local advertisement = consumer(entry.name, entry, status)
            local function recheck()
                check()
                local current_origin, current_status = inspect(ctx, actual, target, origin.connected_at, check)
                if not same_origin(origin, current_origin) or not same_metadata(status, current_status) then fail(origin_failure) end
            end
            -- Preserve constructor intent byte-for-byte through describe/menu/
            -- bound. Only Go canonicalizes it against the retained endpoint.
            local input = { anchor = anchor_path, layout_revision = 1, consumer = advertisement, origin = origin,
                server = server_endpoint }
            if public then input.descriptor = args.runtime end
            local schema = { protocol_min = "integer", protocol_max = "integer", producer_build = "string",
                features = { array = "string" }, layout_revision = "integer", receipt_hash = "string" }
            if public then schema.version, schema.relay = "integer", "string" else schema.receipt = "string" end
            stage = "runtime preflight failed; check the execution-host assets, CLI, endpoint and loaded integration; not retried"
            local described = command("runtime-describe", input, schema, public and small_limit or full_limit)
            recheck()
            compatible(described)
            if not hash(described.receipt_hash) then fail(compatibility) end
            if public then
                if described.version ~= 1 or described.receipt_hash ~= expected_hash
                    or described.relay == "" or #described.relay > small_limit then fail(compatibility) end
            elseif described.receipt == "" or #described.receipt > full_limit then fail(compatibility) end
            -- Preflight was certain. Only now may user guards run or a ticket be
            -- reserved. The original start bounds even slow yielding evaluation.
            stage = "invocation evaluation failed or origin changed; invocation cancelled, not retried"
            local records
            records, ticket = registry.issue(ctx, { client_id = target, connected_at = origin.connected_at,
                window_id = origin.window_id, layer_id = origin.layer_id, metadata = status, expires = deadline })
            recheck()
            local resolved = configuration(options, records)
            stage = "runtime preparation failed; check the original deadline and integration; invocation cancelled, not retried"
            if public then
                input.operation, input.relay, input.configuration = "menu", described.relay, resolved
                prepared = command("runtime-relay", input, prepared_schema, full_limit)
            else
                prepared = command("prepare", { version = 1, configuration = resolved,
                    receipt = described.receipt, receipt_hash = described.receipt_hash }, prepared_schema, full_limit)
            end
            valid_preparation(prepared)
            deadline = math.min(deadline, prepared.expires) -- Never renew either owner's deadline.
            recheck()
            registry.bind(ticket, ctx, prepared)
            recheck()
            if os.time() >= deadline - 5 then fail(origin_failure) end
            if public then
                input.operation, input.configuration, input.receipt_hash = "bound", nil, prepared.receipt_hash
                local bound = command("runtime-relay", input, { version = "integer" }, small_limit)
                if bound.version ~= 1 then fail(compatibility) end
                recheck()
            else
                local handoff = encode({ configuration = resolved, capability = prepared.capability, server = prepared.server,
                    session_id = origin.session_id, block_id = origin.block_id, client_id = target,
                    connected_at = origin.connected_at, completion_action = completion_action, expires = prepared.expires,
                    receipt = prepared.receipt, receipt_hash = prepared.receipt_hash, workspace = prepared.workspace })
                recheck()
                if os.time() >= deadline - 5 then fail(origin_failure) end
                stage = "controller startup failed or uncertain; inspect the original invocation before relaunching; not retried"
                -- Transfer is irreversible even if the start/wait reply is lost.
                attempted_start = true
                local handle, err = rex.exec.start({ prepared.executable, "controller" }, {
                    stdin = handoff, detach = true, timeout = 35, stderr = false,
                })
                check()
                if err ~= nil or handle == nil then fail(stage) end
                local result, wait_err = handle:wait(35)
                check()
                if wait_err ~= nil or not plain(result) or result.ok ~= true or result.code ~= 0 or result.timed_out then fail(stage) end
                local accepted = decode(result.stdout, { version = "integer", status = "string" }, small_limit)
                if accepted.version ~= 1 or accepted.status ~= "accepted" then fail(stage) end
                recheck()
            end
        end)
        if not ok then
            -- Revoke only local, unclaimed authority. No retry or inference that
            -- a failed/uncertain launcher reply undid a possibly started child.
            if ticket and registry then pcall(registry.cancel, ticket) end
            if prepared and not args and not attempted_start then
                local released = pcall(function()
                    local result, err = rex.exec({ helper, "runtime-release" }, {
                        stdin = encode(prepared), timeout = 3, stdout = false, stderr = false,
                    })
                    if err ~= nil or not plain(result) or result.ok ~= true then fail("cleanup failed") end
                end)
                if not released then stage = "preparation cleanup failed; inspect owned invocation artifacts; not retried" end
            end
            fail(stage) -- Never raw child output, user exceptions or capabilities.
        end
        -- No handle, capability, receipt or arbitrary callback result escapes.
    end
    return menu
end

-- At-load registration advertises only already-retained code/menu state. It
-- performs read-only metadata checks, not evaluation, reservation or file I/O.
rex.action({ name = "whichkey_runtime", title = "Describe WhichKey runtime",
    args = { protocol_min = "integer", protocol_max = "integer", action = "string" },
    run = function(ctx, args)
        local ok, result = pcall(function()
            if not plain(args) then fail(compatibility) end
            local count = 0
            for key in pairs(args) do
                if key ~= "protocol_min" and key ~= "protocol_max" and key ~= "action" then fail(compatibility) end
                count = count + 1
            end
            if count ~= 3 or args.protocol_min ~= 1 or args.protocol_max ~= 1 or not name_valid(args.action) then fail(compatibility) end
            local entry = launch_names[args.action]
            if not entry then fail(compatibility) end
            local registry, generation = entry.registry, entry.generation
            local actual = freeze(ctx)
            local function check()
                if not unchanged(ctx, actual) or entry.registry ~= registry or entry.generation ~= generation then fail(origin_failure) end
            end
            encode(args, small_limit)
            local status = metadata(check)
            local value = consumer(args.action, entry, status)
            encode(value, small_limit)
            return value
        end)
        if not ok then fail(compatibility) end
        return result
    end,
})
return M
