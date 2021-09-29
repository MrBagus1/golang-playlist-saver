package reposearch

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"os"
	"playlist-saver/app/config/mysql"
	"playlist-saver/exceptions"
	"playlist-saver/model/record"
)

type SearchRepositoryImpl struct {
	client mysql.Client
}

func NewSearchRepository(client mysql.Client) SearchRepository {
	return &SearchRepositoryImpl{client}
}

func (searcher *SearchRepositoryImpl) GetByUrlId(ctx context.Context, Id string) record.YoutubeData {
	search := record.YoutubeData{}
	err := searcher.client.Conn().Debug().WithContext(ctx).Where("id = ? ", Id).First(&search).Error
	exceptions.PanicIfError(err)

	return search
}

func (searcher *SearchRepositoryImpl) SearchYtByParam(ctx context.Context, Search string) ([]record.YoutubeData, error) {
	API_KEY := os.Getenv("YT_SECRET")
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(API_KEY))
	if err != nil {
		return nil, err
	}
	youtubeData := make([]record.YoutubeData, 0)
	searching := youtubeService.Search.List([]string{"id", "snippet"}).Q(Search).MaxResults(10)
	response, err := searching.Do()
	if err !=nil {
		return nil,err
	}

	for _, item := range response.Items {
		youtubeRecord := record.YoutubeData{}
		youtubeRecord.YoutubeLink = item.Id.VideoId
		youtubeRecord.ChannelId = item.Id.ChannelId
		youtubeRecord.Title = item.Snippet.Title
		youtubeRecord.PublishedAt = item.Snippet.PublishedAt
		youtubeRecord.Description = item.Snippet.Description
		youtubeData = append(youtubeData, youtubeRecord)
	}

	return youtubeData,nil
}

func (searcher *SearchRepositoryImpl) SearchYtById(ctx context.Context, Id string) record.YoutubeData {
	API_KEY := os.Getenv("YT_SECRET")
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(API_KEY))
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	searching := youtubeService.Search.List([]string{"id", "snippet"}).Q(Id).MaxResults(1)
	response, err := searching.Do()
	exceptions.PanicIfError(err)
	youtubeRecord := record.YoutubeData{}
	for _, item := range response.Items {
		youtubeRecord.YoutubeLink = item.Id.VideoId
		youtubeRecord.ChannelId = item.Id.ChannelId
		youtubeRecord.Title = item.Snippet.Title
		youtubeRecord.PublishedAt = item.Snippet.PublishedAt
		youtubeRecord.Description = item.Snippet.Description
	}

	return youtubeRecord
}
