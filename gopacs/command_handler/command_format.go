package command_handler

import (
	"LoLFun/gopacs/data_utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type jsonCommandResponse struct {
	to_display string
	details    string
}

func jsonStringCommandResponse(response jsonCommandResponse) string {
	var string_response = "{"
	string_response += data_utils.GetFormattedKeyValue("to_display", response.to_display, "\"") + ","
	string_response += data_utils.GetFormattedKeyValue("details", response.details, "\"")
	string_response += "}"
	return string_response
}

func echoStringResponse(c echo.Context, response jsonCommandResponse) error {
	return c.String(http.StatusOK, jsonStringCommandResponse(response))
}

func splitCommand(raw_command string) (string, []string) {
	command_args := strings.Split(raw_command, " ")
	return command_args[0], command_args
}
