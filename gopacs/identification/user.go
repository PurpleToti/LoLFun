package identification

import (
	"LoLFun/gopacs/data_utils"
	"errors"
	"time"
)

var count rune = 0

type User struct {
	Name             string
	Last_interaction time.Time
}

func (user *User) Stringify() string {
	repr := "User{"
	repr += data_utils.GetFormattedKeyValue("Name", user.Name, "'") + ","
	repr += data_utils.GetFormattedKeyValue("LastInteraction", user.Last_interaction.String(), "'")
	repr += "}"
	return repr
}

func getNewUserId() string {
	count++
	return string(count)
}

func CreateUser(users_map map[string]*User) (*User, string) {
	new_user := &User{
		Name:             "NaN",
		Last_interaction: time.Now(),
	}
	new_user_id := getNewUserId()
	users_map[new_user_id] = new_user
	return new_user, new_user_id
}

func GetUserFromMap(users_map map[string]*User, key string) (*User, error) {
	user, ok := users_map[key]
	if !ok {
		return nil, errors.New("user id not a key of users map provided")
	}

	return user, nil
}
