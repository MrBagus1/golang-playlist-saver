package playlistCtrl

import (
	"github.com/labstack/echo/v4"
	"playlist-saver/app/middleware"
	"playlist-saver/exceptions"
	"playlist-saver/model/web"
	"playlist-saver/service/servplaylist"
	"playlist-saver/utility"
)

type PlaylistControllerImpl struct {
	PlaylistService servplaylist.PlaylistService
}

func NewPlaylistController(PlaylistService servplaylist.PlaylistService) PlaylistController {
	return &PlaylistControllerImpl{PlaylistService: PlaylistService}
}

func (pl *PlaylistControllerImpl) CreatePlaylist(c echo.Context) error {
	ctx := c.Request().Context()
	ctxUserValue := middleware.GetUserId(c)
	req := web.PlaylistCreateRequest{}
	err := c.Bind(&req)
	exceptions.PanicIfError(err)

	data := pl.PlaylistService.CreatePlaylist(ctx, req.Name, ctxUserValue)

	response := web.PlaylistCreateResponse{
		Name:      data.Name,
		UserId:    data.UserId,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}

	return utility.NewSuccessResponse(c, response)
}
