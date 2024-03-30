package identification

import "time"

type UsersCleaner struct {
	Active    bool
	Users_map map[string]*User
}

func (users_cleaner *UsersCleaner) StartCleaning() {
	for users_cleaner.Active {
		cleanUsers(users_cleaner.Users_map)
		time.Sleep(1 * time.Minute)
	}
}

func cleanUsers(users_map map[string]*User) {
	for key, room := range users_map {
		if room.Last_interaction.Before(time.Now().Add(-user_expire_time)) {
			users_map[key] = nil
		}
	}
}
