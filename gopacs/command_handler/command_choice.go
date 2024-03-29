package command_handler

import (
	"github.com/labstack/echo/v4"
)

func HandleCommand(c echo.Context, raw_command string) error {
	main_command, _ := splitCommand(raw_command)
	switch main_command {
	case "hello":
		return echoStringResponse(c, helloCommand())
	}

	return echoStringResponse(c, jsonCommandResponse{
		to_display: "Command not found",
		details:    "No command corresponding to : '" + main_command + "'",
	})
}
