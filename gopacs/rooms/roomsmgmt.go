package rooms

import "time"

var RoomsMap map[string]*Room = make(map[string]*Room)
var room_expire_time time.Duration = 5 * time.Minute

type RoomsCleaner struct {
	Active   bool
	RoomsMap map[string]*Room
}

func (roomsCleaner *RoomsCleaner) StartCleaning() {
	for roomsCleaner.Active {
		cleanRooms(roomsCleaner.RoomsMap)
		time.Sleep(1 * time.Minute)
	}
}

func cleanRooms(roomsMap map[string]*Room) {
	for key, room := range roomsMap {
		if room.Last_interaction.Before(time.Now().Add(-room_expire_time)) {
			roomsMap[key] = nil
		} else {
			cleanRoom(room)
		}
	}
}
