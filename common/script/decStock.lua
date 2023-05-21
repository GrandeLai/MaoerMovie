-- 1.参数列表
-- 1.1.放映argv
local showStockKey = ARGV[1]
-- 2.脚本业务
-- 2.1.判断该key是否存在
if(redis.call('exists', showStockKey) == 0) then
    -- 2.2.不存在该电影场次，则返回1
    return 1
end
-- 2.2.判断库存是否大于0
if(tonumber(redis.call('get', showStockKey)) <= 0) then
    -- 2.2.库存不足，则返回2
    return 2
end
-- 2.3.库存减1
redis.call('decrby', showStockKey, 1)
return 0
