package utility

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type BaseReponses struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors, omitempty"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	response := BaseReponses{}
	response.Code = http.StatusOK
	response.Message = "Success"
	response.Data = data
	//log.Print(data)
	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, error int, err error) error {
	response := BaseReponses{}
	response.Code = error
	response.Message = "Failed"
	response.Errors = []string{err.Error()}

	return c.JSON(http.StatusBadRequest, response)
}
