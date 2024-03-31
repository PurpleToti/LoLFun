package ciad

import "time"

const users_per_room int = 5

var count_room uint = 0
var count_user uint = 0

var room_expire_time time.Duration = 5 * time.Minute
var user_expire_time time.Duration = 5 * time.Minute

var Rooms_map map[string]*Room = make(map[string]*Room)
var Users_map map[string]*User = make(map[string]*User)

const user_cookie_name string = "user_id"
