package main

import (
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"playlist-saver/app/config"
	"playlist-saver/app/config/mysql"
	_middleware "playlist-saver/app/middleware"
	"playlist-saver/app/routes"
	"playlist-saver/controller"
	"playlist-saver/exceptions"
	"playlist-saver/repository/repouser"
	"playlist-saver/service/servuser"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	exceptions.PanicIfError(err)
	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:       1 << 10, // 1 KB
		DisableStackAll: true,
	}))
	e.HTTPErrorHandler = exceptions.ErrorHandler

	cfg := config.New()
	mysqlClient := mysql.New(cfg)
	defer mysqlClient.Close()

	EXPIRED, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED"))

	ConfigJWT := _middleware.ConfigJWT{
		SecretJWT: os.Getenv("JWT_SECRET"),
		ExpiredIn: EXPIRED,
	}
	userRepo := repouser.NewUserRepository(mysqlClient)
	userService := servuser.NewUserService(userRepo, &ConfigJWT)
	userCtrl := controller.NewUserController(userService)

	routesInit := routes.ControllerList{
		JWTMiddleware:  ConfigJWT.Init(),
		UserController: userCtrl,
	}
	routesInit.Registration(e)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
