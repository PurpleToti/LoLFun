package ciad

import "fmt"

func (room *Room) addUser(user *User) ExitCode {
	var room_pos int = -1
	for i := 0; i < Users_per_room; i++ {
		if room.Users[i] == nil {
			room_pos = i
		}

		if room.Users[i] == user {
			return EC_already_in_room
		}
	}

	if room_pos == -1 {
		return EC_room_full
	}

	room.Users[room_pos] = user

	return EC_ok
}

func (room *Room) getUserPosition(user *User) (int, ExitCode) {
	for i := 0; i < Users_per_room; i++ {
		if room.Users[i] == user {
			return i, EC_ok
		}
	}
	return -1, EC_user_not_in_room
}

func (room *Room) addMessage(message string, user *User) {
	room.Chat = append(room.Chat, user.Name+" : "+message+"\n\n")
}

func (room *Room) freeUser(user *User) ExitCode {
	for i := 0; i < Users_per_room; i++ {
		if user == room.Users[i] {
			room.Users[i] = nil
			ec := room.checkExistence()
			if ec != EC_ok {
				fmt.Println("Room deleted")
			}

			return EC_ok
		}
	}
	return EC_user_not_in_room
}

func (room *Room) checkExistence() ExitCode {
	for i := 0; i < Users_per_room; i++ {
		if room.Users[i] != nil {
			return EC_ok
		}
	}
	return EC_room_empty
}
