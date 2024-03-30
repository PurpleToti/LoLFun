package identification

import "time"

const user_cookie_name string = "user_id"

var user_expire_time time.Duration = 5 * time.Minute
var count int = 0

var Users_map map[string]*User = make(map[string]*User)
