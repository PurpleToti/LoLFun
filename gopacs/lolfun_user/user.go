package lolfun_user

import (
	lolfun_cookies "LoLFun/gopacs/cookie"
	"LoLFun/gopacs/data_utils"
	"errors"
	"strconv"
)

/*
User is a struct that represents a user and everyone using the server needs to be represented
by this struct.

You have to use it when interacting with the server.
*/

type User struct {
	Id string
}

func (user *User) Stringify() string {
	repr := "User{"
	repr += data_utils.GetFormattedKeyValue("Id", user.Id, "'")
	repr += "}"
	return repr
}

/*
UsersMap is a struct that lets you store a finite amount of users to keep acces to them
and manage them

You use this to store active users of a session.
*/

type UsersMap struct {
	Remaining int
	Map       map[string]*User
}

func (usersMap *UsersMap) Stringify() string {
	repr := "UserArray{"
	repr += data_utils.GetFormattedKeyValue("Remaining", strconv.Itoa(usersMap.Remaining), "'") + ","

	var users_repr []string
	for key, value := range usersMap.Map {
		users_repr = append(users_repr, data_utils.GetFormattedKeyValue(
			key, value.Stringify(), "'"),
		)
	}

	repr += data_utils.GetFormattedKeyValue("Map", data_utils.GetFormattedList(users_repr, "'"), "'")
	repr += "}"
	return repr
}

func CreateUsersMap() UsersMap {
	return UsersMap{
		Remaining: Number_of_users,
		Map:       make(map[string]*User),
	}
}

func (usersMap *UsersMap) AddUser(string_id string, user *User) error {
	if usersMap.Remaining < 1 {
		return errors.New("no space to add user")
	}

	usersMap.Map[string_id] = user
	usersMap.Remaining -= 1
	return nil
}

func (usersMap *UsersMap) AddTestUser() error {
	c := lolfun_cookies.GenerateCookie()
	c.Value = "TEST COOKIE VALUE"
	return usersMap.AddUser(
		c.Value,
		&User{
			Id: "TEST",
		},
	)
}
