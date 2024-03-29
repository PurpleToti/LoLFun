package main

import (
	"LoLFun/gopacs/command_handler"
	view_commandprompt "LoLFun/gopacs/views/commandprompt"
	"LoLFun/gopacs/views_utils"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/ressources", "ressources")

	e.GET("/", rootPage)

	e.GET("/handlecommand/:command", commandHandler)

	e.Logger.Fatal(e.Start("localhost:8000"))
}

func rootPage(c echo.Context) error {
	return views_utils.UtilsRender(c, view_commandprompt.AdminCommandPromptPage())
}

func commandHandler(c echo.Context) error {
	var command string = c.Param("command")
	return command_handler.HandleCommand(c, command)
}
