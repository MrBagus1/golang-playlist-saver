package playlistDetailCtrl

import "github.com/labstack/echo/v4"

type PlaylistDetailController interface {
	AddYoutubeToPlaylist(c echo.Context) error
	DeleteYoutubeDataFromPlaylist(c echo.Context) error
}