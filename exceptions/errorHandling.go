package exceptions

import (
	"github.com/labstack/echo/v4"
	"log"
)

func ErrorHandler(err error, c echo.Context) {
	log.Println(err.Error())
}
