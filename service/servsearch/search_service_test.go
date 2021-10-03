package servsearch

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"playlist-saver/model/record"
	searchMock "playlist-saver/repository/reposearch/mocks"
	"testing"
)

var (
	searchRepository searchMock.SearchRepository
)

func setup() SearchService {
	searchService := NewSearchService(&searchRepository)
	return searchService
}

func TestGetByUrlId(t *testing.T){
	searchService := setup()
	t.Run("test case 1, success get url id ", func(t *testing.T) {
		expectedReturn := record.YoutubeData{
			Id:          1,
			YoutubeLink: "youtube.com/test",
			Title:       "Naruto Vs Sasuke",
			ChannelId:   "123123123",
			PublishedAt: "24 september 2000",
			Description: "Mantap Keren",
		}

		searchRepository.On("GetByUrlId", mock.Anything, mock.AnythingOfType("string")).Return(expectedReturn).Once()
		actualResult := searchService.GetByUrlId(context.Background(),"1")
		assert.Equal(t, actualResult.YoutubeLink, expectedReturn.YoutubeLink)
	})
}

func TestSearchYtByParam(t *testing.T){
	searchService := setup()
	t.Run("test case 1, success searching ", func(t *testing.T) {
		expectedReturn := []record.YoutubeData{
			{
				Id:          1,
				YoutubeLink: "youtube.com/test",
				Title:       "Naruto Vs Sasuke",
				ChannelId:   "123123123",
				PublishedAt: "24 september 2000",
				Description: "Mantap Keren",
			},
			{
				Id:          2,
				YoutubeLink: "youtube.com/test2",
				Title:       "Naruto Vs Sasuke2",
				ChannelId:   "1231231232222",
				PublishedAt: "24 september 2000222",
				Description: "Mantap Kere2222n",
			},
		}
		searchRepository.On("SearchYtByParam", mock.Anything, mock.AnythingOfType("string")).Return(expectedReturn, nil).Once()
		_, err := searchService.SearchYtByParam(context.Background(), "Naruto!")
		assert.Nil(t, err)
	})

	t.Run("test case 2, blank ", func(t *testing.T) {

		_, err := searchService.SearchYtByParam(context.Background(), "")
		assert.Error(t,err )
	})

	t.Run("test case 3, error db ", func(t *testing.T) {

		searchRepository.On("SearchYtByParam", mock.Anything, mock.AnythingOfType("string")).Return(nil, fmt.Errorf("something wrong with our server")).Once()
		_, err := searchService.SearchYtByParam(context.Background(), "test")
		assert.Equal(t, err, fmt.Errorf("something wrong with our server"))
	})
}