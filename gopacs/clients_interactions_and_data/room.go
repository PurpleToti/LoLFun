package ciad

import (
	"LoLFun/gopacs/data_utils"
	"time"
)

type Room struct {
	Room_id                string
	Last_interaction       time.Time
	Users                  [users_per_room]*User
	Users_last_interaction [users_per_room]time.Time
	Chat                   []string
}

func (room *Room) Stringify() string {
	repr := "Room{"
	repr += data_utils.GetFormattedKeyValue("Room_id", room.Room_id, "'") + ","
	repr += data_utils.GetFormattedKeyValue("Last_interaction", room.Last_interaction.String(), "'") + ","

	uids_str := make([]string, 0)
	for i := 0; i < users_per_room; i++ {
		if room.Users[i] != nil {
			uids_str = append(uids_str, room.Users[i].User_id)
		}
	}
	repr += data_utils.GetFormattedList(uids_str, "'") + ","

	uli_str := make([]string, 0)
	for i := 0; i < users_per_room; i++ {
		uli_str = append(uli_str, room.Users_last_interaction[i].String())
	}

	repr += data_utils.GetFormattedList(uli_str, "'")
	repr += "}"
	return repr
}

func _emptyUsersArray() [users_per_room]*User {
	var array [users_per_room]*User
	for i := 0; i < users_per_room; i++ {
		array[i] = nil
	}
	return array
}

func _emptyUsers_last_interactionArray() [users_per_room]time.Time {
	var array [users_per_room]time.Time
	for i := 0; i < users_per_room; i++ {
		array[i] = time.Now()
	}
	return array
}

func _getNewRoomId() string {
	user_id := ""
	count_copy := rune(count_room)
	for {
		ascii_carac := 65 + (count_copy % 61)
		if ascii_carac == 92 {
			user_id += "0"
		} else {
			user_id += string(ascii_carac)
		}
		count_copy /= 61
		if count_copy <= 61 {
			break
		}
	}
	count_room++
	return user_id
}

func _addNewRoomToMap(rmap map[string]*Room) *Room {
	new_room_id := _getNewRoomId()
	new_room := &Room{
		Room_id:                new_room_id,
		Last_interaction:       time.Now(),
		Users:                  _emptyUsersArray(),
		Users_last_interaction: _emptyUsers_last_interactionArray(),
		Chat:                   make([]string, 0),
	}
	rmap[new_room.Room_id] = new_room
	return new_room
}

func _getRoomFromMap(rmap map[string]*Room, key string) (*Room, ExitCode) {
	room, ok := rmap[key]
	if !ok {
		return nil, EC_room_not_in_map
	}

	return room, EC_ok
}

func CreateNewRoom() *Room {
	return _addNewRoomToMap(Rooms_map)
}

func GetRoomById(room_id string) (*Room, ExitCode) {
	return _getRoomFromMap(Rooms_map, room_id)
}
