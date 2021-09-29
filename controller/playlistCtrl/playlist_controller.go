package playlistCtrl

import "github.com/labstack/echo/v4"

type PlaylistController interface {
	CreatePlaylist(c echo.Context) error
}
