package ciad

func (user *User) CreateAndJoinRoom() (*Room, ExitCode) {
	room := CreateNewRoom()
	ec := user.JoinRoom(room)
	if ec != EC_ok {
		return room, ec
	}

	return room, EC_ok
}

func (user *User) JoinRoom(room *Room) ExitCode {
	switch room.addUser(user) {
	case EC_ok:
		user.Room = room
		return EC_ok
	case EC_already_in_room:
		user.Room = room
		return EC_already_in_room
	case EC_room_full:
		return EC_room_full
	default:
		return EC_unexpected
	}
}
