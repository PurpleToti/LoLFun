package ciad

import (
	"errors"
	"time"
)

func UserJoinRoom(user *User, room *Room) error {
	for i := 0; i < Users_per_room; i++ {
		if room.Users_id[i] == "" {
			room.Users_id[i] = user.User_id
			room.Users_last_interaction[i] = time.Now()

			user.Room_id = room.Room_id

			return nil
		}
	}
	return errors.New("room is probably full")
}

func UserLeavesRoom(user *User, room *Room) error {
	for i := 0; i < Users_per_room; i++ {
		if room.Users_id[i] == user.User_id {
			room.Users_id[i] = ""
			room.Users_last_interaction[i] = time.Now()

			user.Room_id = ""
			return nil
		}
	}
	return errors.New("user not in the room")
}

func UserIdJoinRoom(user_id string, room *Room) error {
	u, err := GetUserFromMap(Users_map, user_id)
	if err != nil {
		return err
	}
	return UserJoinRoom(u, room)
}

func UserIdJoinRoomId(user_id string, room_id string) error {
	r, err := GetRoomFromMap(Rooms_map, room_id)
	if err != nil {
		return err
	}
	return UserIdJoinRoom(user_id, r)
}

func UserJoinRoomId(user *User, room_id string) error {
	r, err := GetRoomFromMap(Rooms_map, room_id)
	if err != nil {
		return err
	}
	return UserJoinRoom(user, r)
}
