package controller

import (
	"github.com/labstack/echo/v4"
	"playlist-saver/exceptions"
	"playlist-saver/model/web"
	"playlist-saver/service"
	"playlist-saver/utility"
)

type UserControllers struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllers{userService: userService}
}

func (uc *UserControllers) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := web.UserRegisterRequest{}
	err := c.Bind(&req)
	exceptions.PanicIfError(err)

	data := uc.userService.Register(ctx, req)

	return utility.NewSuccessResponse(c, data)
}

func (uc *UserControllers) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req := web.UserLoginRequest{}
	err := c.Bind(&req)
	exceptions.PanicIfError(err)

	data := uc.userService.Login(ctx, req)

	return utility.NewSuccessResponse(c, echo.Map{
		"Token": data,
	})
}
