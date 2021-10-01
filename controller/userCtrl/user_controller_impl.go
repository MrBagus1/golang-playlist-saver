package userCtrl

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"playlist-saver/app/middleware"
	"playlist-saver/model/web"
	"playlist-saver/service/servstatus"
	"playlist-saver/service/servuser"
	"playlist-saver/utility"
)

type UserControllers struct {
	userService servuser.UserService
	statusService servstatus.StatusService
}

func NewUserController(userService servuser.UserService) UserController {
	return &UserControllers{userService: userService}
}

func (uc *UserControllers) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := web.UserRegisterRequest{}
	err := c.Bind(&req)
	if err != nil {
		return utility.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	srvDomain := req.ToDomainService()
	data, err := uc.userService.Register(ctx, srvDomain)
	if err != nil {
		return utility.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	res := web.UserRegisterResponse{}
	res.FromDomainService(data)
	return utility.NewSuccessResponse(c, res)
}

func (uc *UserControllers) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req := web.UserLoginRequest{}
	err := c.Bind(&req)
	if err != nil {
		return utility.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := uc.userService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return utility.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return utility.NewSuccessResponse(c, echo.Map{
		"Token": data,
	})
}

func (uc *UserControllers) GetAllUser(c echo.Context) error {
	ctx := c.Request().Context()
	admin := middleware.GetUserRole(c)

	user, err := uc.userService.GetAllUser(ctx, admin)
	if err != nil {
		return utility.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	if len(user) == 0 {
		return utility.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	finalUser := make([]web.UserGetAllResponse, 0)
	for _, values := range user {
		finalResponse := web.UserGetAllResponse{}
		finalResponse.Name = values.Name
		finalResponse.Email = values.Email
		finalResponse.Birthday = values.Birthday
		finalResponse.Role = values.Role
		finalResponse.CreatedAt = values.CreatedAt
		finalResponse.UpdatedAt = values.UpdatedAt
		finalResponse.Status.Id = values.Status.Id
		finalResponse.Status.Name = values.Status.Name
		finalResponse.Status.CreatedAt = values.Status.CreatedAt
		finalResponse.Status.UpdatedAt = values.Status.UpdatedAt
		finalUser = append(finalUser, finalResponse)
	}

	return utility.NewSuccessResponse(c, finalUser)
}

func (uc *UserControllers) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()
	id := middleware.GetUserId(c)
	request := web.UserUpdateRequest{}
	if err := c.Bind(&request); err != nil {
		return utility.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := uc.userService.UpdateUser(ctx,request.ToDomainService(),id)
	if err != nil {
		return utility.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	byId, err := uc.userService.GetUserById(ctx,id)
	if err != nil {
		return err
	}
	final := web.UserUpdateResponse{
		Name: byId.Name,
		Email: byId.Email,
		Birthday: byId.Birthday,
		Gender: byId.Gender,
	}
	return utility.NewSuccessResponse(c,final)
}

func (uc *UserControllers) AddTokenUser(c echo.Context) error {
	ctx := c.Request().Context()
	id := middleware.GetUserId(c)

	req := web.TokenAddRequest{}
	if err := c.Bind(&req); err != nil {
		return utility.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	//check status
	user, err := uc.userService.GetAllUser(ctx , "ADMIN")
	if err != nil {
		return err
	}
	for _ , values := range user {
		if values.Status.TokenId == req.Id {
			return utility.NewErrorResponse(c,http.StatusBadRequest,errors.New("token already in used!"))
		}
	}


	err = uc.userService.UserAddToken(ctx,id, req.Id, req.TokenId)
	if err != nil {
		return err

	}

	return utility.NewSuccessResponse(c,"you're premium now!")
}



