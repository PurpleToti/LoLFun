package command_handler

import (
	lolfun_cookies "LoLFun/gopacs/cookie"
	"LoLFun/gopacs/lolfun_user"
	"LoLFun/gopacs/session"
)

func helloCommand() jsonCommandResponse {
	return jsonCommandResponse{
		to_display: "Hello World!",
		details:    "Just a classic hello world :)",
	}
}

func addTestUserCommand() jsonCommandResponse {
	session.LoLFunUsersMap.AddTestUser()
	return jsonCommandResponse{
		to_display: "New user succesfully added!",
		details:    "Adds a new Test user to the list of users of the current server session",
	}
}

func seeUsersCommand() jsonCommandResponse {
	return jsonCommandResponse{
		to_display: session.LoLFunUsersMap.Stringify(),
		details:    "Displays all current users in session",
	}
}

func addUserCommand(id_user string) jsonCommandResponse {
	new_user := lolfun_user.User{
		Id: id_user,
	}

	if err := session.LoLFunUsersMap.AddUser(lolfun_cookies.GenerateCookie(), &new_user); err != nil {
		return jsonCommandResponse{
			to_display: "Could not add the new user",
			details:    err.Error(),
		}
	}

	return jsonCommandResponse{
		to_display: "New user added : " + new_user.Stringify(),
		details:    "The new user was succesfully inserted in the current session list of users",
	}
}
