package servplaylistdetail

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"playlist-saver/model/record"
	playlistDetailMock "playlist-saver/repository/repoplaylistdetail/mocks"
	searchRepositoryMock "playlist-saver/repository/reposearch/mocks"
	"testing"
)

var (
	playlistDetailRepo playlistDetailMock.PlaylistDetailRepository
	searchRepository   searchRepositoryMock.SearchRepository
)

func setup() PlaylistDetailService {
	detailService := NewPlaylistDetail(&playlistDetailRepo, &searchRepository)
	return detailService
}

func TestAddYoutubeToPlaylist(t *testing.T) {
	detailService := setup()
	t.Run("test case 1, success add youtube to playlist ", func(t *testing.T) {
		domain := PlaylistDetail{
			Id:            1,
			PlaylistId:    1,
			YoutubeDataId: "231asdasdasd",
			YoutubeData: YoutubeData{
				YoutubeLink: "yotuube.com?",
				Title:       "Mahabrata",
				ChannelId:   "title",
				PublishedAt: "awokawkoaw",
				Description: "wamantap",
			},
		}

		expectedReturn := record.PlaylistDetail{
			Id:            1,
			PlaylistId:    1,
			YoutubeDataId: "231asdasdasd",
			YoutubeData: record.YoutubeData{
				YoutubeLink: "yotuube.com?",
				Title:       "Mahabrata",
				ChannelId:   "title",
				PublishedAt: "awokawkoaw",
				Description: "wamantap",
			},
		}
		searchRepository.On("SearchYtById", mock.Anything, mock.AnythingOfType("string")).Return(expectedReturn.YoutubeData).Once()
		playlistDetailRepo.On("AddYoutubeToPlaylist", mock.Anything, mock.Anything).Return(expectedReturn).Once()

		playlist, err := detailService.AddYoutubeToPlaylist(context.Background(),domain)

		assert.Equal(t, playlist.Id, expectedReturn.Id)
		assert.Nil(t, err)
	})

}

func TestDeleteYoutubeDataFromPlaylist(t *testing.T) {
	detailService := setup()
	t.Run("test case 1, success delete youtube data from playlist", func(t *testing.T) {
		playlistDetailRepo.On("DeleteYoutubeDataFromPlaylist", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()
		err := detailService.DeleteYoutubeDataFromPlaylist(context.Background(), 1)

		assert.Nil(t, err)
	})

	t.Run("test case 2, failed delete data from playlist detail", func(t *testing.T) {
		playlistDetailRepo.On("DeleteYoutubeDataFromPlaylist", mock.Anything, mock.AnythingOfType("int")).Return(fmt.Errorf("something wrong with our server")).Once()
		err := detailService.DeleteYoutubeDataFromPlaylist(context.Background(), 1)
		assert.Equal(t, err, fmt.Errorf("something wrong with our server"))
	})


}
