package searchCtrl

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"playlist-saver/model/web"
	"playlist-saver/service/servsearch"
	"playlist-saver/utility"
)

type SearchControllerImpl struct {
	SearchService servsearch.SearchService
}

func NewSearchController(SearchService servsearch.SearchService) SearchController {
	return &SearchControllerImpl{SearchService: SearchService}
}

func (sc *SearchControllerImpl) SearchYtByParam(c echo.Context) error {
	ctx := c.Request().Context()
	searchParams := c.QueryParam("q")

	data, err := sc.SearchService.SearchYtByParam(ctx, searchParams)
	if err != nil {
		return utility.NewErrorResponse(c,http.StatusBadRequest,err)
	}

	//log.Print("Print datsa2 dari controoller", len(data))
	if len(data) == 0 {
		return utility.NewErrorResponse(c, http.StatusBadRequest,errors.New("Not Found"))
	}

	responseData := make([]web.YoutubeSearchResponse, 0)
	for _, values := range data {
		log.Println("channel Id", values.ChannelId)
		res := web.YoutubeSearchResponse{}
		res.YoutubeLink = values.YoutubeLink
		res.Title = values.Title
		res.ChannelId = values.ChannelId
		res.PublishedAt = values.PublishedAt
		res.Description = values.Description
		responseData = append(responseData, res)
	}
	return utility.NewSuccessResponse(c, responseData)
}
