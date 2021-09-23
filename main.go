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
	"playlist-saver/repository"
	"playlist-saver/service"
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
	userRepo := repository.NewUserRepository(mysqlClient)
	userService := service.NewUserService(userRepo, &ConfigJWT)
	userCtrl := controller.NewUserController(userService)

	routesInit := routes.ControllerList{
		JWTMiddleware:  ConfigJWT.Init(),
		UserController: userCtrl,
	}
	routesInit.Registration(e)

	// add middleware and routes
	// ...
	s := http.Server{
		Addr:    ":8080",
		Handler: e,
		//ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}
	log.Print("Server is running under", s.Addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
