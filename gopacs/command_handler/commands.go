package command_handler

import (
	"LoLFun/gopacs/identification"
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
