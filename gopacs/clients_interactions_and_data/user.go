package ciad

import (
	"LoLFun/gopacs/data_utils"
	"errors"
	"time"
)

/*
User struct
*/
type User struct {
	User_id          string
	Name             string
	Last_interaction time.Time
	Room             *Room
}

/*
User string format
*/
func (user *User) Stringify() string {
	repr := "User{"
	repr += data_utils.GetFormattedKeyValue("User_id", user.User_id, "'") + ","
	repr += data_utils.GetFormattedKeyValue("Name", user.Name, "'") + ","
	repr += data_utils.GetFormattedKeyValue("Last_interaction", user.Last_interaction.String(), "'") + ","
	if user.Room != nil {
		repr += data_utils.GetFormattedKeyValue("Room_id", user.Room.Room_id, "'")
	} else {
		repr += data_utils.GetFormattedKeyValue("Room_id", "", "'")
	}
	repr += "}"
	return repr
}

/*
User manipulation
*/

func _getNewUserId() string {
	user_id := ""
	count_copy := rune(count_user)
	for {
		user_id += string(65 + (count_copy % 61))
		count_copy /= 61
		if count_copy <= 61 {
			break
		}
	}
	count_user++
	return user_id
}

func _addNewUserToMap(umap map[string]*User) *User {
	new_user_id := _getNewUserId()
	new_user := &User{
		User_id:          new_user_id,
		Name:             "NaN",
		Last_interaction: time.Now(),
		Room:             nil,
	}
	umap[new_user.User_id] = new_user
	return new_user
}

func _getUserFromMap(umap map[string]*User, key string) (*User, error) {
	user, ok := umap[key]
	if !ok {
		return nil, errors.New("user id not a key of users map provided")
	}

	return user, nil
}

func CreateNewUser() *User {
	return _addNewUserToMap(Users_map)
}

func GetUserById(user_id string) (*User, error) {
	return _getUserFromMap(Users_map, user_id)
}
