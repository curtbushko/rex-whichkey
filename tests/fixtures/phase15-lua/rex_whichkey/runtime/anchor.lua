-- Normal installer-owned file module. Keep the logical loader path: namespace
-- symlink updates may select new assets next time, never within an invocation.
local source = debug.getinfo(1, "S").source
return { layout_revision = 1, source = source }
