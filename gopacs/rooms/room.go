package rooms

import "time"

type Room struct {
	Room_id                string
	Last_interaction       time.Time
	Users_id               [users_per_room]string
	Users_last_interaction [users_per_room]time.Time
}

func emptyUsers_idArray() [users_per_room]string {
	var array [users_per_room]string
	for i := 0; i < users_per_room; i++ {
		array[i] = ""
	}
	return array
}

func emptyUsers_last_interactionArray() [users_per_room]time.Time {
	var array [users_per_room]time.Time
	for i := 0; i < users_per_room; i++ {
		array[i] = time.Now()
	}
	return array
}

func getNewRoomId() string {
	user_id := ""
	count_copy := count
	for {
		if count_copy <= 61 {
			break
		}
		user_id += string(65 + (count_copy % 61))
		count_copy /= 61
	}
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
	for i := 0; i < users_per_room; i++ {
		if room.Users_last_interaction[i].Before(time.Now().Add(-user_expire_time)) {
			room.Users_id[i] = ""
		}
	}
}
