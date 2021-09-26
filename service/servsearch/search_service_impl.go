package servsearch

import (
	"context"
	"errors"
	"log"
	"playlist-saver/repository/reposearch"
)

type SearchServiceImpl struct {
	SearchRepository reposearch.SearchRepository
}

func NewSearchService(SearchRepository reposearch.SearchRepository) SearchService {
	return &SearchServiceImpl{
		SearchRepository: SearchRepository,
	}
}

func (sr *SearchServiceImpl) GetByUrlId(ctx context.Context, Id string) YoutubeData {
	dataYoutube := YoutubeData{}
	if len(Id) == 0 {
		panic(errors.New("Id can't be blank"))
	}

	searchResult := sr.SearchRepository.GetByUrlId(ctx, Id)
	dataYoutube.FromRecordYoutubeNotArray(searchResult)

	return dataYoutube
}

func (sr *SearchServiceImpl) SearchYtByParam(ctx context.Context, Search string) []YoutubeData {
	dataYoutube := YoutubeData{}
	if len(Search) == 0 {
		panic(errors.New("What you want to searchCtrl, dont leave the params blank"))
	}

	dataFinal := make([]YoutubeData, 0)
	youtubeResult := sr.SearchRepository.SearchYtByParam(ctx, Search)
	log.Print("Print data dari Youtube Reesult Service", len(youtubeResult))
	dataYoutube.FromRecordYoutube(youtubeResult)

	dataFinal = append(dataFinal, dataYoutube)
	log.Print("Print daata Service", len(dataFinal))

	return dataFinal
}
