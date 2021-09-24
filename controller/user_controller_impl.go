package controller

import (
	"github.com/labstack/echo/v4"
	"playlist-saver/exceptions"
	"playlist-saver/model/web"
	"playlist-saver/service/servuser"
	"playlist-saver/utility"
)

type UserControllers struct {
	userService servuser.UserService
}

func NewUserController(userService servuser.UserService) UserController {
	return &UserControllers{userService: userService}
}

func (uc *UserControllers) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := web.UserRegisterRequest{}
	err := c.Bind(&req)
	exceptions.PanicIfError(err)

	srvDomain := req.ToDomainService()
	data := uc.userService.Register(ctx, srvDomain)
	res := web.UserRegisterResponse{}
	res.FromDomainService(data)
	return utility.NewSuccessResponse(c, res)
}

func (uc *UserControllers) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req := web.UserLoginRequest{}
	err := c.Bind(&req)
	exceptions.PanicIfError(err)

	data := uc.userService.Login(ctx, req.Email, req.Password)

	return utility.NewSuccessResponse(c, echo.Map{
		"Token": data,
	})
}
