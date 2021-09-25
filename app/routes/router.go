package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"playlist-saver/controller"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController controller.UserController
}

func (c1 *ControllerList) Registration(e *echo.Echo) {

	apiV1 := e.Group("/api/v1")

	//	USER THINGS
	apiV1.POST("/users/register", c1.UserController.Register)
	apiV1.POST("/users/login", c1.UserController.Login)
}
