package command_handler

import (
	ciad "LoLFun/gopacs/clients_interactions_and_data"
	"fmt"
)

func helloCommand() jsonCommandResponse {
	return jsonCommandResponse{
		to_display: "Hello World!",
		details:    "Just a classic hello world :)",
	}
}

func meCommand(user *ciad.User) jsonCommandResponse {
	fmt.Println("request received from : " + user.Name)
	return jsonCommandResponse{
		to_display: user.Stringify(),
		details:    "To see how the server sees you...",
	}
}

func meNameCommand(user *ciad.User, new_username string) jsonCommandResponse {
	user.Name = new_username
	return jsonCommandResponse{
		to_display: user.Stringify(),
		details:    "Change your username",
	}
}

func createRoomCommand(user *ciad.User) jsonCommandResponse {
	room, ec := user.CreateAndJoinRoom()
	if ec != ciad.EC_ok {
		return jsonCommandResponse{
			to_display: room.Stringify() + "\nerror joining newly created room",
			details:    "The room was created but the creator could not join the room succesfully",
		}
	}
	return jsonCommandResponse{
		to_display: room.Stringify(),
		details:    "The room was created and the creator joined the room succesfully",
	}
}

func meJoinRoomCommand(user *ciad.User, room_id string) jsonCommandResponse {
	room, ec := ciad.GetRoomById(room_id)
	if ec != ciad.EC_ok {
		return jsonCommandResponse{
			to_display: "Error joining room",
			details:    "The room could not be joined by user",
		}
	} else {
		if ec := user.JoinRoom(room); ec == ciad.EC_ok {
			return jsonCommandResponse{
				to_display: "Room joined succesfully",
				details:    "You joined an already existing room",
			}
		}
	}

	return jsonCommandResponse{
		to_display: "Unexpected error while joining room",
		details:    "???",
	}
}

func newUserCommand() jsonCommandResponse {
	u := ciad.CreateNewUser()
	return jsonCommandResponse{
		to_display: u.Stringify(),
		details:    "A new user has been created",
	}
}
