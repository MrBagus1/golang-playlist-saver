package tokenCtrl

import "github.com/labstack/echo/v4"

type TokenController interface {
	AddToken(c echo.Context) error
	GetToken(c echo.Context)  error
}
