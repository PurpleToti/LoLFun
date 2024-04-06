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
	ec := room.addUser(user)
	if ec != EC_ok {
		return ec
	}

	ec = user._leaveCurrentRoom()
	if ec != EC_ok {
		return ec
	}

	user.Room = room

	return EC_ok
}

func (user *User) SendMessageToRoom(message string) ExitCode {
	if user.Room == nil {
		return EC_user_not_in_room
	}
	user.Room.addMessage(message, user)
	return EC_ok
}

func (user *User) ChangeName(new_name string) ExitCode {
	if len(new_name) < 1 {
		return EC_bad_username
	}

	user.Name = new_name
	return EC_ok
}

func (user *User) _leaveCurrentRoom() ExitCode {
	if user.Room == nil {
		return EC_ok
	}
	ec := user.Room.freeUser(user)
	if ec != EC_ok {
		return ec
	}
	user.Room = nil
	return EC_ok
}
