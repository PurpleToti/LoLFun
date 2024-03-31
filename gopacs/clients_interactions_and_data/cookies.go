package ciad

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func WriteCookie(c echo.Context, name string, value string, expires time.Time) {
	cookie := new(http.Cookie)
	cookie.Path = "/"
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = expires
	cookie.HttpOnly = true
	cookie.Secure = false
	cookie.SameSite = http.SameSiteLaxMode
	c.SetCookie(cookie)
}

func WriteUserCookie(c echo.Context, user_id string) error {
	WriteCookie(c, user_cookie_name, user_id, time.Now().Add(user_expire_time))
	return nil
}

func UserCookieThere(c echo.Context) (bool, error) {
	cookie, err := c.Cookie(user_cookie_name)
	if err != nil {
		return false, err
	}

	if cookie.Expires.Compare(time.Now()) == 1 {
		return false, nil
	}

	return true, nil
}

func GetUserCookie(c echo.Context) (*http.Cookie, error) {
	cookie, err := c.Cookie(user_cookie_name)
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func GetUserCookieValue(c echo.Context) (string, error) {
	cookie, err := c.Cookie(user_cookie_name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
