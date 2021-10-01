package userCtrl

import "github.com/labstack/echo/v4"

type UserController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	GetAllUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	AddTokenUser(c echo.Context) error
}
