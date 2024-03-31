package ciad

import "github.com/labstack/echo/v4"

type LoLFunContext struct {
	EchoContext echo.Context
	ContextUser *User
	UserRoom    *Room
}

func NewLoLFunContext(echoctx echo.Context) *LoLFunContext {
	lolfunctx := &LoLFunContext{
		EchoContext: echoctx,
		ContextUser: nil,
		UserRoom:    nil,
	}
	lolfunctx.ContextUser = handleIdentification(echoctx)
	lolfunctx.UserRoom = lolfunctx.ContextUser.Room
	return lolfunctx
}
