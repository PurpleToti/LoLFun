package main

import (
	"LoLFun/gopacs/command_handler"
	"LoLFun/gopacs/identification"
	"LoLFun/gopacs/rooms"
	"LoLFun/gopacs/userroominteractions"
	view_commandprompt "LoLFun/gopacs/views/commandprompt"
	view_profileinterface "LoLFun/gopacs/views/profileinterface"
	view_roominterface "LoLFun/gopacs/views/roominterface"
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

	e.GET("/user", userPage)
	e.GET("/user/latest", latestUserVersion)
	e.POST("/user/update", updateUser)

	e.GET("/room", roomPage)
	e.GET("/room/new", newRoom)
	e.POST("/room/join", joinRoom)
	e.GET("/room/latest", latestRoomVersion)

	e.GET("/commandprompt", commandPromptPage)
	e.GET("/handlecommand/:command", commandHandler)

	e.Logger.Fatal(e.Start("192.168.1.56:8080"))
}

func userPage(c echo.Context) error {
	user, err := identification.HandleIdentification(c)
	if err != nil {
		return err
	}

	return views_utils.UtilsRender(c, view_profileinterface.UserPage(user))
}

func latestUserVersion(c echo.Context) error {
	user, err := identification.HandleIdentification(c)
	if err != nil {
		return views_utils.UtilsRender(c, view_profileinterface.ProfilePostResponse(-1))
	}
	return views_utils.UtilsRender(c, view_profileinterface.ProfileDescDiv(user))
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

func roomPage(c echo.Context) error {
	user, err := identification.HandleIdentification(c)
	if err != nil {
		return err
	}

	room, err := rooms.GetRoomFromMap(rooms.Rooms_map, user.Room_id)
	if err != nil {
		room = nil
	}

	return views_utils.UtilsRender(c, view_roominterface.RoomPage(user, room))
}

func newRoom(c echo.Context) error {
	_, err := identification.HandleIdentification(c)
	if err != nil {
		return err
	}

	room := rooms.CreateRoom(rooms.Rooms_map)
	return views_utils.UtilsRender(c, view_roominterface.CreateRoomDivResponse(room))
}

func joinRoom(c echo.Context) error {
	user, err := identification.HandleIdentification(c)
	if err != nil {
		return views_utils.UtilsRender(c, view_roominterface.JoinRoomDivResponse(2))
	}

	room_id := c.FormValue("room_id")
	err = userroominteractions.UserJoinRoomId(user, room_id)
	if err != nil {
		return views_utils.UtilsRender(c, view_roominterface.JoinRoomDivResponse(1))
	}

	return views_utils.UtilsRender(c, view_roominterface.JoinRoomDivResponse(0))
}

func latestRoomVersion(c echo.Context) error {
	user, err := identification.HandleIdentification(c)
	if err != nil {
		return err
	}

	return views_utils.UtilsRender(c, view_roominterface.RoomDescDiv(user))
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
