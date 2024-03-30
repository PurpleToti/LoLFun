package rooms

import (
	"LoLFun/gopacs/data_utils"
	"errors"
	"time"
)

type Room struct {
	Room_id                string
	Last_interaction       time.Time
	Users_id               [Users_per_room]string
	Users_last_interaction [Users_per_room]time.Time
}

func (room *Room) Stringify() string {
	repr := "Room{"
	repr += data_utils.GetFormattedKeyValue("Room_id", room.Room_id, "'") + ","
	repr += data_utils.GetFormattedKeyValue("Last_interaction", room.Last_interaction.String(), "'") + ","
	repr += data_utils.GetFormattedList(room.Users_id[:], "'") + ","

	uli_str := make([]string, 0)
	for i := 0; i < Users_per_room; i++ {
		uli_str = append(uli_str, room.Users_last_interaction[i].String())
	}

	repr += data_utils.GetFormattedList(uli_str, "'")
	repr += "}"
	return repr
}

func emptyUsers_idArray() [Users_per_room]string {
	var array [Users_per_room]string
	for i := 0; i < Users_per_room; i++ {
		array[i] = ""
	}
	return array
}

func emptyUsers_last_interactionArray() [Users_per_room]time.Time {
	var array [Users_per_room]time.Time
	for i := 0; i < Users_per_room; i++ {
		array[i] = time.Now()
	}
	return array
}

func getNewRoomId() string {
	user_id := ""
	count_copy := rune(count)
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
	count++
	return user_id
}

func CreateRoom(users_map map[string]*Room) *Room {
	new_room_id := getNewRoomId()
	new_room := &Room{
		Room_id:                new_room_id,
		Last_interaction:       time.Now(),
		Users_id:               emptyUsers_idArray(),
		Users_last_interaction: emptyUsers_last_interactionArray(),
	}
	users_map[new_room.Room_id] = new_room
	return new_room
}

func cleanRoom(room *Room) {
	for i := 0; i < Users_per_room; i++ {
		if room.Users_last_interaction[i].Before(time.Now().Add(-room_expire_time)) {
			room.Users_id[i] = ""
		}
	}
}

func GetRoomFromMap(Rooms_map map[string]*Room, key string) (*Room, error) {
	room, ok := Rooms_map[key]
	if !ok {
		return nil, errors.New("user id not a key of users map provided")
	}

	return room, nil
}
