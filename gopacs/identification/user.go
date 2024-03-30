package identification

import (
	"LoLFun/gopacs/data_utils"
	"errors"
	"time"
)

type User struct {
	User_id          string
	Name             string
	Last_interaction time.Time
	Room_id          string
}

func (user *User) Stringify() string {
	repr := "User{"
	repr += data_utils.GetFormattedKeyValue("User_id", user.User_id, "'") + ","
	repr += data_utils.GetFormattedKeyValue("Name", user.Name, "'") + ","
	repr += data_utils.GetFormattedKeyValue("Last_interaction", user.Last_interaction.String(), "'") + ","
	repr += data_utils.GetFormattedKeyValue("Room_id", user.Room_id, "'")
	repr += "}"
	return repr
}

func getNewUserId() string {
	user_id := ""
	count_copy := rune(count)
	for {
		user_id += string(65 + (count_copy % 61))
		count_copy /= 61
		if count_copy <= 61 {
			break
		}
	}
	count++
	return user_id
}

func CreateUser(users_map map[string]*User) *User {
	new_user_id := getNewUserId()
	new_user := &User{
		User_id:          new_user_id,
		Name:             "NaN",
		Last_interaction: time.Now(),
		Room_id:          "",
	}
	users_map[new_user.User_id] = new_user
	return new_user
}

func GetUserFromMap(users_map map[string]*User, key string) (*User, error) {
	user, ok := users_map[key]
	if !ok {
		return nil, errors.New("user id not a key of users map provided")
	}

	return user, nil
}
