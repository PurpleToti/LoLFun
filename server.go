package main

import (
	"LoLFun/gopacs/command_handler"
	view_commandprompt "LoLFun/gopacs/views/commandprompt"
	view_room "LoLFun/gopacs/views/room"
	view_user "LoLFun/gopacs/views/user"
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

	e.GET("/", testPage)

	e.GET("/user", userContent)
	e.GET("/user/latest", latestUserVersion)
	e.POST("/user/update", updateUser)

	e.GET("/room", roomContent)
	e.GET("/room/new", newRoom)
	e.POST("/room/join", joinRoom)
	e.GET("/room/latest", latestRoomVersion)
	e.GET("/room/chat", retrieveChatBox)
	e.POST("/room/send", sendRoomMessage)

	e.GET("/commandprompt", commandPromptPage)
	e.GET("/handlecommand/:command", commandHandler)

	e.Logger.Fatal(e.Start("localhost:8080"))
}

func testPage(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	return views_utils.UtilsRender(c, view_user.UserPage(lolfunctx))
}

func userContent(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	return views_utils.UtilsRender(c, view_user.UserContent(lolfunctx))
}

func latestUserVersion(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	return views_utils.UtilsRender(c, view_user.UserDescription(lolfunctx.ContextUser))
}

func updateUser(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)

	new_username := c.FormValue("username")
	ec := lolfunctx.ContextUser.ChangeName(new_username)
	return views_utils.UtilsRender(c, view_user.UserSettingsUpdateResponse(ec))
}

func roomContent(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	return views_utils.UtilsRender(c, view_room.RoomContent(lolfunctx))
}

func newRoom(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	_, ec := lolfunctx.ContextUser.CreateAndJoinRoom()
	return views_utils.UtilsRender(c, view_room.RoomManagementNewRoomResponse(ec))
}

func joinRoom(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	room_id := c.FormValue("room_id")
	room, ec := ciad.GetRoomById(room_id)
	if ec != ciad.EC_ok {
		return views_utils.UtilsRender(c, view_room.RoomManagementJoinRoomResponse(ec))
	}
	ec = lolfunctx.ContextUser.JoinRoom(room)
	return views_utils.UtilsRender(c, view_room.RoomManagementJoinRoomResponse(ec))
}

func latestRoomVersion(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	return views_utils.UtilsRender(c, view_room.UsersList(lolfunctx.UserRoom))
}

func retrieveChatBox(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	return views_utils.UtilsRender(c, view_room.RoomChatMessages(lolfunctx.UserRoom.Chat))
}

func sendRoomMessage(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	message := c.FormValue("message")
	ec := lolfunctx.ContextUser.SendMessageToRoom(message)
	return views_utils.UtilsRender(c, view_room.RoomChatNewMessageResponse(ec))
}

func commandPromptPage(c echo.Context) error {
	return views_utils.UtilsRender(c, view_commandprompt.AdminCommandPromptPage())
}

func commandHandler(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	var command string = c.Param("command")
	return command_handler.HandleCommand(c, command, lolfunctx.ContextUser)
}
