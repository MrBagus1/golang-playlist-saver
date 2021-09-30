package playlistCtrl

import "github.com/labstack/echo/v4"

type PlaylistController interface {
	CreatePlaylist(c echo.Context) error
	GetAllPlaylist(c echo.Context) error
	GetPlaylistByUserId(c echo.Context) error
	DeletePlaylist(c echo.Context) error
	EditPlaylist(c echo.Context) error
}
