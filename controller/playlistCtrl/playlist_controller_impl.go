package playlistCtrl

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"playlist-saver/app/middleware"
	"playlist-saver/exceptions"
	"playlist-saver/model/web"
	"playlist-saver/service/servplaylist"
	"playlist-saver/utility"
	"strconv"
)

type PlaylistControllerImpl struct {
	PlaylistService servplaylist.PlaylistService
}

func NewPlaylistController(PlaylistService servplaylist.PlaylistService) PlaylistController {
	return &PlaylistControllerImpl{PlaylistService: PlaylistService}
}

func (pl *PlaylistControllerImpl) CreatePlaylist(c echo.Context) error {
	ctx := c.Request().Context()
	ctxUserId := middleware.GetUserId(c)
	ctxUserStatus := middleware.GetUserStatus(c)
	req := web.PlaylistCreateRequest{}
	err := c.Bind(&req)
	exceptions.PanicIfError(err)

	data, err := pl.PlaylistService.CreatePlaylist(ctx, req.Name, ctxUserId, ctxUserStatus)
	if err != nil {
		return utility.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := web.PlaylistCreateResponse{
		Name:      data.Name,
		UserId:    data.UserId,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}

	return utility.NewSuccessResponse(c, response)
}

func (pl *PlaylistControllerImpl) GetAllPlaylist(c echo.Context) error {
	ctx := c.Request().Context()
	role := middleware.GetUserRole(c)

	if role != "ADMIN" {
		return utility.NewErrorResponse(c, http.StatusUnauthorized, errors.New("you're not admin, can't access this feature"))
	}

	finalRes := make([]web.PlaylistGetResponse, 0)

	playlist, err := pl.PlaylistService.GetAllPlaylist(ctx)
	if err != nil {
		return utility.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	for _, values := range playlist {
		dataSemi := web.PlaylistGetResponse{}
		dataSemi.Id = values.Id
		dataSemi.Name = values.Name
		dataSemi.UserId = values.UserId
		dataSemi.CreatedAt = values.CreatedAt
		dataSemi.UpdatedAt = values.UpdatedAt
		dataSemi.DeletedAt = values.DeletedAt
		for _, values2 := range values.PlaylistDetail {
			dataSemi2 := web.PlaylistDetailGetResponse{}
			dataSemi2.Id = values2.Id
			dataSemi2.PlaylistId = values2.PlaylistId
			dataSemi2.YoutubeDataId = values2.YoutubeDataId
			dataSemi2.YoutubeData.Id = values2.YoutubeData.Id
			dataSemi2.YoutubeData.Title = values2.YoutubeData.Title
			dataSemi2.YoutubeData.ChannelId = values2.YoutubeData.ChannelId
			dataSemi2.YoutubeData.PublishedAt = values2.YoutubeData.PublishedAt
			dataSemi2.YoutubeData.Description = values2.YoutubeData.Description
			dataSemi.PlaylistDetail = append(dataSemi.PlaylistDetail, dataSemi2)
		}
		finalRes = append(finalRes, dataSemi)
		log.Println("Finalres", finalRes)
	}
	return utility.NewSuccessResponse(c, finalRes)
}

func (pl *PlaylistControllerImpl) GetPlaylistByUserId(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	atoi, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	finalRes := make([]web.PlaylistGetResponse, 0)

	playlist, err := pl.PlaylistService.GetPlaylistByUserId(ctx, atoi)
	if err != nil {
		return utility.NewErrorResponse(c, http.StatusBadRequest, errors.New("can't find the id"))
	}

	for _, values := range playlist {
		dataSemi := web.PlaylistGetResponse{}
		dataSemi.Id = values.Id
		dataSemi.Name = values.Name
		dataSemi.UserId = values.UserId
		dataSemi.CreatedAt = values.CreatedAt
		dataSemi.UpdatedAt = values.UpdatedAt
		dataSemi.DeletedAt = values.DeletedAt
		for _, values2 := range values.PlaylistDetail {
			dataSemi2 := web.PlaylistDetailGetResponse{}
			dataSemi2.Id = values2.Id
			dataSemi2.PlaylistId = values2.PlaylistId
			dataSemi2.YoutubeDataId = values2.YoutubeDataId
			dataSemi2.YoutubeData.Id = values2.YoutubeData.Id
			dataSemi2.YoutubeData.Title = values2.YoutubeData.Title
			dataSemi2.YoutubeData.ChannelId = values2.YoutubeData.ChannelId
			dataSemi2.YoutubeData.PublishedAt = values2.YoutubeData.PublishedAt
			dataSemi2.YoutubeData.Description = values2.YoutubeData.Description
			dataSemi.PlaylistDetail = append(dataSemi.PlaylistDetail, dataSemi2)
		}
		finalRes = append(finalRes, dataSemi)
		log.Println("Finalres", finalRes)
	}
	return utility.NewSuccessResponse(c, finalRes)
}

func (pl *PlaylistControllerImpl) DeletePlaylist(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	atoi, err := strconv.Atoi(id)
	if err != nil {
		return utility.NewErrorResponse(c, http.StatusBadRequest, err)

	}
	err = pl.PlaylistService.DeletePlaylist(ctx, atoi)
	if err != nil {
		return utility.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return utility.NewSuccessResponse(c, "Success Delete Playlist !")
}

func (pl *PlaylistControllerImpl) EditPlaylist(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	atoi, err := strconv.Atoi(id)
	if err != nil {
		return utility.NewErrorResponse(c, http.StatusBadRequest, err)

	}

	request := web.PlaylistEditRequest{}
	if err := c.Bind(&request); err != nil {
		return utility.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	playlist, err := pl.PlaylistService.EditPlaylist(ctx, request.ToDomainService(), atoi)
	if err != nil {
		return utility.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := web.PlaylistEditResponse{
		Name: playlist.Name,
		UserId: playlist.UserId,
		CreatedAt: playlist.CreatedAt,
		UpdateAt: playlist.UpdatedAt,
	}

	return utility.NewSuccessResponse(c, response)
}
