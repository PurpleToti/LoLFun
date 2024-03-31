package ciad

import (
	"time"

	"github.com/labstack/echo/v4"
)

func _createNewUser(c echo.Context) *User {
	// fmt.Println("Creating user...")

	new_user := CreateNewUser()
	WriteUserCookie(c, new_user.User_id)
	return new_user
}

func _refreshUser(c echo.Context, user *User) *User {
	// fmt.Println("Finding user...")

	user.Last_interaction = time.Now()
	WriteUserCookie(c, user.User_id)
	return user
}

func handleIdentification(c echo.Context) *User {
	if is_there, _ := UserCookieThere(c); is_there {
		user_id, err := GetUserCookieValue(c)
		if err != nil {
			return _createNewUser(c)
		}
		user, err := GetUserById(user_id)
		if err != nil {
			return _createNewUser(c)
		}
		return _refreshUser(c, user)
	} else {
		return _createNewUser(c)
	}
}

// DEPRECATED
func HandleIdentification(c echo.Context) *User {
	return handleIdentification(c)
}
