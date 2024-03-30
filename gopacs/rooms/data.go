package rooms

import "time"

const users_per_room int = 5

var count rune = 0
var user_expire_time time.Duration = 5 * time.Minute
var room_expire_time time.Duration = 5 * time.Minute

var RoomsMap map[string]*Room = make(map[string]*Room)
