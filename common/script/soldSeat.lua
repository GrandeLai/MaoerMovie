local soldSeats = KEYS[1]
local seat = ARGV[1]

local seatsStr = redis.call("GET", soldSeats)
local seats = {}
for s in string.gmatch(seatsStr, "%d+") do
    seats[#seats + 1] = s
end

seats[#seats + 1] = seat

table.sort(seats)

local newSeatsStr = table.concat(seats, ",")
redis.call("SET", soldSeats, newSeatsStr)