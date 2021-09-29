package userCtrl

import (
	"github.com/labstack/echo/v4"
	"net/http"
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
	if err != nil{
		return utility.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	srvDomain := req.ToDomainService()
	data,err := uc.userService.Register(ctx, srvDomain)
	if err != nil {
		return utility.NewErrorResponse(c,http.StatusBadRequest,err)
	}

	res := web.UserRegisterResponse{}
	res.FromDomainService(data)
	return utility.NewSuccessResponse(c, res)
}

func (uc *UserControllers) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req := web.UserLoginRequest{}
	err := c.Bind(&req)
	if err != nil{
		return utility.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data,err := uc.userService.Login(ctx, req.Email, req.Password)
	if err !=nil {
		return utility.NewErrorResponse(c,http.StatusBadRequest,err)
	}

	return utility.NewSuccessResponse(c, echo.Map{
		"Token": data,
	})
}
