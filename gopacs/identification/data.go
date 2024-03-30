package identification

import "time"

const user_cookie_name string = "user_id"

var user_expire_time time.Duration = 5 * time.Minute
var count rune = 0

var UsersMap map[string]*User = make(map[string]*User)
