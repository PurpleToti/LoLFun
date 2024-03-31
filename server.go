package main

import (
	"LoLFun/gopacs/command_handler"
	view_commandprompt "LoLFun/gopacs/views/commandprompt"
	view_profileinterface "LoLFun/gopacs/views/profileinterface"
	view_roominterface "LoLFun/gopacs/views/roominterface"
	"LoLFun/gopacs/views_utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	ciad "LoLFun/gopacs/clients_interactions_and_data"
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

	e.GET("/user", userPage)
	e.GET("/user/latest", latestUserVersion)
	e.POST("/user/update", updateUser)

	e.GET("/room", roomPage)
	e.GET("/room/new", newRoom)
	e.POST("/room/join", joinRoom)
	e.GET("/room/latest", latestRoomVersion)
	e.GET("/room/chat", retrieveChatBox)
	e.POST("/room/send/", sendRoomMessage)

	e.GET("/commandprompt", commandPromptPage)
	e.GET("/handlecommand/:command", commandHandler)

	e.Logger.Fatal(e.Start("192.168.1.56:8080"))
}

func userPage(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	return views_utils.UtilsRender(c, view_profileinterface.UserPage(lolfunctx.ContextUser))
}

func latestUserVersion(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	return views_utils.UtilsRender(c, view_profileinterface.ProfileDescDiv(lolfunctx.ContextUser))
}

func updateUser(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)

	new_username := c.FormValue("username")
	if len(new_username) < 1 {
		return views_utils.UtilsRender(c, view_profileinterface.ProfilePostResponse(1))
	}

	lolfunctx.ContextUser.Name = new_username
	return views_utils.UtilsRender(c, view_profileinterface.ProfilePostResponse(0))
}

func roomPage(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	return views_utils.UtilsRender(c, view_roominterface.RoomPage(lolfunctx.ContextUser, lolfunctx.UserRoom))
}

func newRoom(c echo.Context) error {
	ciad.NewLoLFunContext(c)

	room := ciad.CreateNewRoom()
	return views_utils.UtilsRender(c, view_roominterface.CreateRoomDivResponse(room))
}

func joinRoom(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)

	room_id := c.FormValue("room_id")
	room, ec := ciad.GetRoomById(room_id)
	if ec != ciad.EC_ok {
		return views_utils.UtilsRender(c, view_roominterface.JoinRoomDivResponse(ec))
	}

	ec = lolfunctx.ContextUser.JoinRoom(room)
	return views_utils.UtilsRender(c, view_roominterface.JoinRoomDivResponse(ec))
}

func latestRoomVersion(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)

	return views_utils.UtilsRender(c, view_roominterface.RoomDescDiv(lolfunctx.ContextUser))
}

func retrieveChatBox(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	return views_utils.UtilsRender(c, view_roominterface.RoomChatDiv(lolfunctx.UserRoom))
}

func sendRoomMessage(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	message := c.FormValue("message")
	ec := lolfunctx.ContextUser.SendMessageToRoom(message)
	return views_utils.UtilsRender(c, view_roominterface.NewRoomMessageDivResponse(ec))
}

func commandPromptPage(c echo.Context) error {
	ciad.NewLoLFunContext(c)
	return views_utils.UtilsRender(c, view_commandprompt.AdminCommandPromptPage())
}

func commandHandler(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	var command string = c.Param("command")
	return command_handler.HandleCommand(c, command, lolfunctx.ContextUser)
}
