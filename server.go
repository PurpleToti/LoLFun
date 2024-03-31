package main

import (
	view_room "LoLFun/gopacs/views/room"
	view_user "LoLFun/gopacs/views/user"
	"LoLFun/gopacs/views_utils"
	"fmt"
	"net/http"

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
	e.POST("/room/send/", sendRoomMessage)

	e.GET("/commandprompt", commandPromptPage)
	e.GET("/handlecommand/:command", commandHandler)

	e.Logger.Fatal(e.Start("192.168.1.56:8080"))
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
	fmt.Println(lolfunctx.ContextUser.Name)
	return c.String(http.StatusOK, "OK")

	//room_id := c.FormValue("room_id")
	//room, ec := ciad.GetRoomById(room_id)
	//if ec != ciad.EC_ok {
	//	return views_utils.UtilsRender(c, view_roominterface.JoinRoomDivResponse(ec))
	//}
	//
	//ec = lolfunctx.ContextUser.JoinRoom(room)
	//return views_utils.UtilsRender(c, view_roominterface.JoinRoomDivResponse(ec))
}

func latestRoomVersion(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	fmt.Println(lolfunctx.ContextUser.Name)
	return c.String(http.StatusOK, "OK")

	//return views_utils.UtilsRender(c, view_roominterface.RoomDescDiv(lolfunctx.ContextUser))
}

func retrieveChatBox(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	fmt.Println(lolfunctx.ContextUser.Name)
	return c.String(http.StatusOK, "OK")
	//return views_utils.UtilsRender(c, view_roominterface.RoomChatDiv(lolfunctx.UserRoom))
}

func sendRoomMessage(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	fmt.Println(lolfunctx.ContextUser.Name)
	return c.String(http.StatusOK, "OK")

	//message := c.FormValue("message")
	//ec := lolfunctx.ContextUser.SendMessageToRoom(message)
	//return views_utils.UtilsRender(c, view_roominterface.NewRoomMessageDivResponse(ec))
}

func commandPromptPage(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	fmt.Println(lolfunctx.ContextUser.Name)
	return c.String(http.StatusOK, "OK")
	//
	//return views_utils.UtilsRender(c, view_commandprompt.AdminCommandPromptPage())
}

func commandHandler(c echo.Context) error {
	lolfunctx := ciad.NewLoLFunContext(c)
	fmt.Println(lolfunctx.ContextUser.Name)
	return c.String(http.StatusOK, "OK")

	//var command string = c.Param("command")
	//return command_handler.HandleCommand(c, command, lolfunctx.ContextUser)
}
