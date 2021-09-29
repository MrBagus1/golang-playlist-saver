package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"playlist-saver/controller/playlistCtrl"
	"playlist-saver/controller/playlistDetailCtrl"
	"playlist-saver/controller/searchCtrl"
	"playlist-saver/controller/userCtrl"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	UserController     userCtrl.UserController
	SearchController   searchCtrl.SearchController
	PlaylistController playlistCtrl.PlaylistController
	DetailController   playlistDetailCtrl.PlaylistDetailController
}

func (c1 *ControllerList) Registration(e *echo.Echo) {

	apiV1 := e.Group("/api/v1")

	//	USER THINGS
	apiV1.POST("/users/register", c1.UserController.Register)
	apiV1.POST("/users/login", c1.UserController.Login)

	//	SEARCH
	apiV1.GET("/search", c1.SearchController.SearchYtByParam, middleware.JWTWithConfig(c1.JWTMiddleware))

	//CREATE PLAYLIST
	apiV1.POST("/playlist", c1.PlaylistController.CreatePlaylist, middleware.JWTWithConfig(c1.JWTMiddleware))

	//ADD YOUTUBE TO PLAYLIST
	apiV1.POST("/playlist/:playlistId/data", c1.DetailController.AddYoutubeToPlaylist, middleware.JWTWithConfig(c1.JWTMiddleware))
}
