package rooms

import "time"

const Users_per_room int = 5

var count uint = 0
var room_expire_time time.Duration = 5 * time.Minute

var Rooms_map map[string]*Room = make(map[string]*Room)
