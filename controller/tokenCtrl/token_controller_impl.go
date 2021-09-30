package tokenCtrl

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"playlist-saver/app/middleware"
	"playlist-saver/model/web"
	"playlist-saver/service/servtoken"
	"playlist-saver/utility"
)

type TokenControllerImpl struct {
	TokenService servtoken.TokenService
}

func NewTokenController(TokenService servtoken.TokenService) TokenController {
	return &TokenControllerImpl{TokenService: TokenService}
}

func (tcl *TokenControllerImpl) AddToken(c echo.Context) error {
	ctx := c.Request().Context()
	role := middleware.GetUserRole(c)

	if role != "ADMIN"{
		return utility.NewErrorResponse(c, http.StatusUnauthorized, errors.New("you're not admin, can't access this!"))
	}

	err := tcl.TokenService.AddToken(ctx)
	if err != nil {
		return err
	}

	return utility.NewSuccessResponse(c,errors.New("Success Generate premium Token!"))
}

func (tcl *TokenControllerImpl) GetToken(c echo.Context) error {
	ctx := c.Request().Context()
	role := middleware.GetUserRole(c)

	if role != "ADMIN"{
		return utility.NewErrorResponse(c, http.StatusUnauthorized, errors.New("you're not admin, can't access this!"))
	}

	dataFinal := make([]web.TokenResponse,0)
	token, err := tcl.TokenService.GetToken(ctx)
	if err != nil {
		return err
	}

	for _, values := range token {
		dataArr := web.TokenResponse{}
		dataArr.Id = values.Id
		dataArr.CreatedAt = values.CreatedAt
		dataArr.UpdatedAt = values.UpdatedAt
		dataFinal = append(dataFinal, dataArr)
	}

	return utility.NewSuccessResponse(c, dataFinal)
}
