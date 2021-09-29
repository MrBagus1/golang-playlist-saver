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

func (sr *SearchServiceImpl) SearchYtByParam(ctx context.Context, Search string) ([]YoutubeData, error) {

	if len(Search) == 0 {
		return nil, errors.New("You want to search, but you leave the form blank.")
	}

	dataFinal := make([]YoutubeData, 0)
	youtubeResult, err := sr.SearchRepository.SearchYtByParam(ctx, Search)
	if err != nil{
		return nil, err
	}
	//log.Print("Print data dari Youtube Reesult Service", len(youtubeResult))
	//dataYoutube.FromRecordYoutube(youtubeResult)

	for _, values := range youtubeResult {
		dataYoutube := YoutubeData{}
		log.Println("ChannelService youtube", values.ChannelId)
		dataYoutube.YoutubeLink = values.YoutubeLink
		dataYoutube.Title = values.Title
		dataYoutube.ChannelId = values.ChannelId
		dataYoutube.PublishedAt = values.PublishedAt
		dataYoutube.Description = values.Description
		dataFinal = append(dataFinal, dataYoutube)
	}

	//log.Print("Print daata Serviceee", len(dataFinal))

	return dataFinal, nil
}
