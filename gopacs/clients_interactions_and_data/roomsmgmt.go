package ciad

import "time"

type RoomsCleaner struct {
	Active    bool
	Rooms_map map[string]*Room
	Interval  time.Duration
}

func (rooms_cleaner *RoomsCleaner) StartCleaning() {
	for rooms_cleaner.Active {
		cleanRooms(rooms_cleaner.Rooms_map)
		time.Sleep(rooms_cleaner.Interval)
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
