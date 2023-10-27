package interfaces

import "github.com/labstack/echo/v4"

type IAuthHandlers interface {
	Login(ctx echo.Context) error
}

type IAuthService interface {
	CreateToken() error
	CreateUser() error
}
