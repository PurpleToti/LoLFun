package ciad

import (
	"time"

	"github.com/labstack/echo/v4"
)

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

	if lolfunctx.ContextUser.Room != nil {
		pos, ec := lolfunctx.ContextUser.Room.getUserPosition(lolfunctx.ContextUser)
		if ec == EC_ok {
			lolfunctx.ContextUser.Room.Users_last_interaction[pos] = time.Now()
		}
	}

	return lolfunctx
}
