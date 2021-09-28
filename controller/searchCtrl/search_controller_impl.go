package searchCtrl

import (
	"errors"
	"github.com/labstack/echo/v4"
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

	data := sc.SearchService.SearchYtByParam(ctx, searchParams)
	//log.Print("Print datsa2 dari controoller", len(data))
	if len(data) == 0 {
		return utility.NewErrorResponse(c, errors.New("Not Found"))
	}

	responseData := make([]web.YoutubeSearchResponse, 0)
	for _, values := range data {
		res := web.YoutubeSearchResponse{}
		res.Id = values.Id
		res.Title = values.Title
		res.ChannelId = values.ChannelId
		res.PublishedAt = values.PublishedAt
		res.Description = values.Description
		responseData = append(responseData, res)
	}
	return utility.NewSuccessResponse(c, responseData)
}
