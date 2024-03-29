package identification

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const user_cookie_name string = "user_id"

func WriteCookie(c echo.Context, name string, value string, expires time.Time) error {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = expires
	c.SetCookie(cookie)
	return nil
}

func WriteUserCookie(c echo.Context, user_id string) error {
	err := WriteCookie(c, user_cookie_name, user_id, time.Now().Add(5*time.Minute))
	if err != nil {
		return errors.New("unexpected error while writing user cookie")
	}
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
