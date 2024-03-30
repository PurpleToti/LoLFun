package identification

import (
	"time"

	"github.com/labstack/echo/v4"
)

var UsersMap map[string]*User = make(map[string]*User)

func createNewUser(c echo.Context) (*User, error) {
	// fmt.Println("Creating user...")

	new_user, new_user_id := CreateUser(UsersMap)
	WriteUserCookie(c, new_user_id)
	return new_user, nil
}

func refreshUser(c echo.Context, user *User) (*User, error) {
	// fmt.Println("Finding user...")

	user.Last_interaction = time.Now()
	cookie, err := GetUserCookie(c)
	if err != nil {
		return nil, err
	}
	cookie.Expires = time.Now().Add(5 * time.Minute)
	return user, nil
}

func HandleIdentification(c echo.Context) (*User, error) {
	if is_there, _ := UserCookieThere(c); is_there {
		user_id, err := GetUserCookieValue(c)
		if err != nil {
			return createNewUser(c)
		}
		user, err := GetUserFromMap(UsersMap, user_id)
		if err != nil {
			return createNewUser(c)
		}
		return refreshUser(c, user)
	} else {
		return createNewUser(c)
	}
}
