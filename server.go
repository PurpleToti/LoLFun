package main

import (
	"LoLFun/gopacs/command_handler"
	"LoLFun/gopacs/identification"
	view_commandprompt "LoLFun/gopacs/views/commandprompt"
	"LoLFun/gopacs/views_utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/ressources", "ressources")

	e.GET("/", func(c echo.Context) error {
		user, err := identification.HandleIdentification(c)
		if err != nil {
			return err
		}

		fmt.Println(user.Name)
		fmt.Println(user.Last_interaction)

		return c.String(http.StatusOK, "HOME PAGE.")
	})

	e.GET("/commandprompt", commandPromptPage)
	e.GET("/handlecommand/:command", commandHandler)

	e.Logger.Fatal(e.Start("localhost:8000"))
}

func commandPromptPage(c echo.Context) error {
	_, err := identification.HandleIdentification(c)
	if err != nil {
		return err
	}

	return views_utils.UtilsRender(c, view_commandprompt.AdminCommandPromptPage())
}

func commandHandler(c echo.Context) error {
	user, err := identification.HandleIdentification(c)
	if err != nil {
		return err
	}

	var command string = c.Param("command")
	return command_handler.HandleCommand(c, command, user)
}
