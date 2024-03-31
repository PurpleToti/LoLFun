package ciad

import "time"

type RoomsCleaner struct {
	Active   bool
	Rmap     map[string]*Room
	Interval time.Duration
}

func (rooms_cleaner *RoomsCleaner) StartCleaning() {
	for rooms_cleaner.Active {
		cleanRooms(rooms_cleaner.Rmap)
		time.Sleep(rooms_cleaner.Interval)
	}
}

func cleanRoom(room *Room) {
	for i := 0; i < Users_per_room; i++ {
		if room.Users_last_interaction[i].Before(time.Now().Add(-room_expire_time)) {
			room.Users[i] = nil
		}
	}
}

func cleanRooms(rooms_map map[string]*Room) {
	for key, room := range rooms_map {
		if room.Last_interaction.Before(time.Now().Add(-room_expire_time)) {
			rooms_map[key] = nil
		} else {
			cleanRoom(room)
		}
	}
}
