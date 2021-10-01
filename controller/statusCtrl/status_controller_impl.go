package statusCtrl

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"playlist-saver/service/servstatus"
	"playlist-saver/utility"
)

type StatusControllerImpl struct {
	StatusService servstatus.StatusService
}

func NewStatusController(StatusService servstatus.StatusService) StatusController {
	return &StatusControllerImpl{StatusService: StatusService}
}

func (scl *StatusControllerImpl) CronCheckToken(c echo.Context) error {
	ctx := c.Request().Context()

	err := scl.StatusService.CronPremiumChecker(ctx)
	if err != nil {
		return utility.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return utility.NewSuccessResponse(c, "Success check cron!")
}
