package lolfun_cookies

import (
	"time"
)

const User_identification_cookie string = "UserIdentificationCookie"
const Cookie_expire_time time.Duration = 30 * time.Minute
const Ascii_limit uint = 127
const Number_of_caracters uint = 10

var Ascii_repr_caracs [Number_of_caracters]rune = NewUserCookieValueArray()
