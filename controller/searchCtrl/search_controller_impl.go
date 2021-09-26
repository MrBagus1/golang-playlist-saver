package searchCtrl

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
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
	searchParams := c.Param("name")

	data := sc.SearchService.SearchYtByParam(ctx, searchParams)
	log.Print("Print datsa2 dari controoller", len(data))
	if len(data) == 0 {
		return utility.NewErrorResponse(c, errors.New("Not Found"))
	}
	res := web.YoutubeSearchResponse{}
	res.FromDomainService(data)
	return utility.NewSuccessResponse(c, res)
}
