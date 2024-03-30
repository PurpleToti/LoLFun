package main

import (
	"LoLFun/gopacs/command_handler"
	"LoLFun/gopacs/identification"
	view_commandprompt "LoLFun/gopacs/views/commandprompt"
	view_profileinterface "LoLFun/gopacs/views/profileinterface"
	"LoLFun/gopacs/views_utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowCredentials: true,
		MaxAge:           3000,
	}))
	e.Static("/ressources", "ressources")

	e.GET("/", homePage)
	e.POST("/user/update", updateUser)

	e.GET("/commandprompt", commandPromptPage)
	e.GET("/handlecommand/:command", commandHandler)

	e.Logger.Fatal(e.Start("192.168.1.56:8080"))
}

func homePage(c echo.Context) error {
	user, err := identification.HandleIdentification(c)
	if err != nil {
		return err
	}

	return views_utils.UtilsRender(c, view_profileinterface.ProfileSubmitPage(user))
}

func updateUser(c echo.Context) error {
	user, err := identification.HandleIdentification(c)
	if err != nil {
		return views_utils.UtilsRender(c, view_profileinterface.ProfilePostResponse(-1))
	}

	new_username := c.FormValue("username")
	if len(new_username) < 1 {
		return views_utils.UtilsRender(c, view_profileinterface.ProfilePostResponse(1))
	}

	user.Name = new_username

	return views_utils.UtilsRender(c, view_profileinterface.ProfilePostResponse(0))
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
