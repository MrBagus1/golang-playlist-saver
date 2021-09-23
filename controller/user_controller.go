package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}
