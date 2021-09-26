package searchCtrl

import "github.com/labstack/echo/v4"

type SearchController interface {
	SearchYtByParam(c echo.Context) error
}
