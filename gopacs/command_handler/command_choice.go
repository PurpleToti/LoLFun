package command_handler

import (
	"LoLFun/gopacs/identification"

	"github.com/labstack/echo/v4"
)

func HandleCommand(c echo.Context, raw_command string, user *identification.User) error {
	main_command, _ := splitCommand(raw_command)
	switch main_command {
	case "hello":
		return echoStringResponse(c, helloCommand())
	case "me":
		return echoStringResponse(c, meCommand(user))
	}

	return echoStringResponse(c, jsonCommandResponse{
		to_display: "Command not found",
		details:    "No command corresponding to : '" + main_command + "'",
	})
}
