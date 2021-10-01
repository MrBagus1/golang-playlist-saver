package statusCtrl

import "github.com/labstack/echo/v4"

type StatusController interface {
	CronCheckToken(c echo.Context) error
}