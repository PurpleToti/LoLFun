package command_handler

import (
	"github.com/labstack/echo/v4"
)

func HandleCommand(c echo.Context, raw_command string) error {
	main_command, command_args := splitCommand(raw_command)
	switch main_command {
	case "hello":
		return echoStringResponse(c, helloCommand())

	case "addTestUser":
		return echoStringResponse(c, addTestUserCommand())

	case "seeUsers":
		return echoStringResponse(c, seeUsersCommand())

	case "addUser":
		return echoStringResponse(c, addUserCommand(command_args[1]))
	}

	return echoStringResponse(c, jsonCommandResponse{
		to_display: "Command not found",
		details:    "No command corresponding to : '" + main_command + "'",
	})
}
