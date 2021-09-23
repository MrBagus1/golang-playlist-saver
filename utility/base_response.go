package utility

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type BaseReponses struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Errors []string    `json:"errors, omitempty"`
	Data   interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	response := BaseReponses{}
	response.Code = http.StatusOK
	response.Status = "200"
	response.Data = data
	log.Print(data)
	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, err error) error {
	response := BaseReponses{}
	response.Code = http.StatusBadRequest
	response.Status = "400"
	response.Errors = []string{err.Error()}

	return c.JSON(http.StatusOK, response)
}
