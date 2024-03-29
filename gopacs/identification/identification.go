package identification

import (
	"github.com/labstack/echo/v4"
)

var UsersMap map[string]*User = make(map[string]*User)

func HandleIdentification(c echo.Context) (*User, error) {
	if is_there, _ := UserCookieThere(c); is_there {
		user_id, err := GetUserCookieValue(c)
		if err != nil {
			return nil, err
		}
		user, err := GetUserFromMap(UsersMap, user_id)
		if err != nil {
			return nil, err
		}
		return user, nil
	} else {
		user := CreateUser(UsersMap)
		return user, nil
	}
}
