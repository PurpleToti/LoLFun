package lolfun_cookies

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func GenerateCookie() *http.Cookie {
	return &http.Cookie{
		Name:  User_identification_cookie,
		Value: "NOVAL",
	}
}

func WriteUserCookie(c echo.Context, value string) error {
	cookie := new(http.Cookie)
	cookie.Name = User_identification_cookie
	cookie.Value = value
	cookie.Expires = time.Now().Add(Cookie_expire_time)
	c.SetCookie(cookie)
	return nil
}

func GenerateNewUserCookieValue() string {
	value := ""
	for i := 0; i < int(Number_of_caracters); i++ {
		value += string(Ascii_repr_caracs[i])
	}
	NextCookie()
	return value
}
