package command_handler

import (
	"LoLFun/gopacs/identification"
	"LoLFun/gopacs/rooms"
	"LoLFun/gopacs/userroominteractions"
)

func helloCommand() jsonCommandResponse {
	return jsonCommandResponse{
		to_display: "Hello World!",
		details:    "Just a classic hello world :)",
	}
}

func meCommand(user *identification.User) jsonCommandResponse {
	return jsonCommandResponse{
		to_display: user.Stringify(),
		details:    "To see how the server sees you...",
	}
}

func meNameCommand(user *identification.User, new_username string) jsonCommandResponse {
	user.Name = new_username
	return jsonCommandResponse{
		to_display: user.Stringify(),
		details:    "Change your username",
	}
}

func createRoomCommand(user *identification.User) jsonCommandResponse {
	r := rooms.CreateRoom(rooms.Rooms_map)
	err := userroominteractions.UserJoinRoom(user, r)
	if err != nil {
		return jsonCommandResponse{
			to_display: r.Stringify() + "\nerror joining newly created room",
			details:    "The room was created but the creator could not join the room succesfully",
		}
	}
	return jsonCommandResponse{
		to_display: r.Stringify(),
		details:    "The room was created and the creator joined the room succesfully",
	}
}

func meJoinRoomCommand(user *identification.User, room_id string) jsonCommandResponse {
	err := userroominteractions.UserJoinRoomId(user, room_id)
	if err != nil {
		return jsonCommandResponse{
			to_display: "Error joining room",
			details:    "The room could not be joined by user",
		}
	}
	return jsonCommandResponse{
		to_display: "Room joined succesfully",
		details:    "You joined an already existing room",
	}
}

func newUserCommand() jsonCommandResponse {
	u := identification.CreateUser(identification.Users_map)
	return jsonCommandResponse{
		to_display: u.Stringify(),
		details:    "A new user has been created",
	}
}
