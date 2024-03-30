package rooms

import "time"

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
