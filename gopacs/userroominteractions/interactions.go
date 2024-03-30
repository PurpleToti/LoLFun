package userroominteractions

import (
	"LoLFun/gopacs/identification"
	"LoLFun/gopacs/rooms"
	"errors"
	"time"
)

func UserJoinRoom(user *identification.User, room *rooms.Room) error {
	for i := 0; i < rooms.Users_per_room; i++ {
		if room.Users_id[i] == "" {
			room.Users_id[i] = user.User_id
			room.Users_last_interaction[i] = time.Now()

			user.Room_id = room.Room_id

			return nil
		}
	}
	return errors.New("room is probably full")
}

func UserLeavesRoom(user *identification.User, room *rooms.Room) error {
	for i := 0; i < rooms.Users_per_room; i++ {
		if room.Users_id[i] == user.User_id {
			room.Users_id[i] = ""
			room.Users_last_interaction[i] = time.Now()

			user.Room_id = ""
			return nil
		}
	}
	return errors.New("user not in the room")
}

func UserIdJoinRoom(user_id string, room *rooms.Room) error {
	u, err := identification.GetUserFromMap(identification.Users_map, user_id)
	if err != nil {
		return err
	}
	return UserJoinRoom(u, room)
}

func UserIdJoinRoomId(user_id string, room_id string) error {
	r, err := rooms.GetRoomFromMap(rooms.Rooms_map, room_id)
	if err != nil {
		return err
	}
	return UserIdJoinRoom(user_id, r)
}

func UserJoinRoomId(user *identification.User, room_id string) error {
	r, err := rooms.GetRoomFromMap(rooms.Rooms_map, room_id)
	if err != nil {
		return err
	}
	return UserJoinRoom(user, r)
}
