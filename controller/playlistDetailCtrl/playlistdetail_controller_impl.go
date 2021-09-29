package playlistDetailCtrl

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"playlist-saver/model/web"
	"playlist-saver/service/servplaylistdetail"
	"playlist-saver/utility"
	"strconv"
)

type PlaylistDetailControllerImpl struct{
	PlaylistDetail servplaylistdetail.PlaylistDetailService
}

func NewPlaylistDetail(PlaylistDetail servplaylistdetail.PlaylistDetailService) PlaylistDetailController {
	return &PlaylistDetailControllerImpl{PlaylistDetail: PlaylistDetail}
}


func (pl *PlaylistDetailControllerImpl) AddYoutubeToPlaylist(c echo.Context) error {
	ctx := c.Request().Context()
	requestParams := c.Param("playlistId")
	req := web.DetailCreateRequest{}
	err := c.Bind(&req)
	if err != nil {
		return utility.NewErrorResponse(c,http.StatusBadRequest,errors.New("Something wrong with your request"))
	}
	//serv domain
	log.Println("REQUEST PARAMS", requestParams)
	atoi, err := strconv.Atoi(requestParams)
	log.Print("ATOI, " , atoi)
	if err != nil {
		return utility.NewErrorResponse(c, http.StatusBadRequest,errors.New("Error with your request params"))
	}

	servReq := servplaylistdetail.PlaylistDetail{
		PlaylistId: atoi,
		YoutubeDataId: req.YoutubeId ,
	}
	result, err := pl.PlaylistDetail.AddYoutubeToPlaylist(ctx,servReq)

	response := web.DetailCreateResponse{
		PlaylistId: result.PlaylistId,
		YoutubeId: result.YoutubeDataId,
	}

	return utility.NewSuccessResponse(c, response)
}

