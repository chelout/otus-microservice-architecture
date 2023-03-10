package repository

import (
	"homework-4/internal/server"

	"github.com/labstack/echo/v4"
)

type User interface {
	CreateUser(ctx echo.Context) (server.User, error)
	DeleteUser(ctx echo.Context, userId int64) error
	FindUserById(ctx echo.Context, userId int64) (server.User, error)
	UpdateUser(ctx echo.Context, userId int64) (server.User, error)
}

//func (u server.User) CreateUser(ctx echo.Context) (server.User, error) {
//
//}