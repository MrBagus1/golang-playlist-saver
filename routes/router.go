package routes

import "github.com/labstack/echo/v4"

func NewRoute() *echo.Route {
	e := echo.New()
	ev1 := e.Group("api/v1")

	ev1.POST("user/register")
	//ev1.POST("user/login")

	return e
}
