package servplaylist

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"playlist-saver/model/record"
	playlistMock "playlist-saver/repository/repoplaylist/mocks"
	"testing"
	"time"
)

var (
	playlistRepo playlistMock.PlaylistRepository
)

func setup() PlaylistService {
	playlistService := NewPlaylistService(&playlistRepo)
	return playlistService
}

func TestAddYoutubeToPlaylist(t *testing.T) {
	playlistService := setup()
	t.Run("test case 1, success create playlist for free user! ", func(t *testing.T) {
		expectedReturnUid := []record.Playlist{
			{
				Id:        1,
				Name:      "wakwow",
				UserId:    3,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			{
				Id:        2,
				Name:      "wakwow",
				UserId:    4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
		}

		expectedReturnCreating := record.Playlist{
			Id:        1,
			Name:      "wakwow",
			UserId:    3,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		}
		playlistRepo.On("GetPlaylistByUserId", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnUid, nil).Once()
		playlistRepo.On("CreatePlaylist", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(expectedReturnCreating, nil).Once()

		playlist, err := playlistService.CreatePlaylist(context.Background(), "wakwow", 1, "FREE")

		assert.Equal(t, playlist.Name, expectedReturnCreating.Name)
		assert.Nil(t, err)
	})

	t.Run("test case 2, error len name | len id == 0 ! ", func(t *testing.T) {
		_, err := playlistService.CreatePlaylist(context.Background(), "", 1, "FREE")
		assert.Error(t, err)
	})

	t.Run("test case 3, status free, but get playlist error! ", func(t *testing.T) {
		expectedReturnUid := []record.Playlist{}

		playlistRepo.On("GetPlaylistByUserId", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnUid, fmt.Errorf("something wrong with the server")).Once()
		_, err := playlistService.CreatePlaylist(context.Background(), "wakwow", 1, "FREE")

		assert.Equal(t, err, fmt.Errorf("something wrong with the server"))
	})

	t.Run("test case 4, user free but added more than 10 playlist, error! ", func(t *testing.T) {
		expectedReturnUid := []record.Playlist{
			{
				Id:        1,
				Name:      "wakwow",
				UserId:    3,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			{
				Id:        2,
				Name:      "wakwow",
				UserId:    4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			{
				Id:        3,
				Name:      "wakwow",
				UserId:    4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			{
				Id:        4,
				Name:      "wakwow",
				UserId:    4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			{
				Id:        5,
				Name:      "wakwow",
				UserId:    4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			{
				Id:        6,
				Name:      "wakwow",
				UserId:    4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			{
				Id:        7,
				Name:      "wakwow",
				UserId:    4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			{
				Id:        8,
				Name:      "wakwow",
				UserId:    4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			{
				Id:        9,
				Name:      "wakwow",
				UserId:    4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			{
				Id:        10,
				Name:      "wakwow",
				UserId:    4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			{
				Id:        11,
				Name:      "wakwow",
				UserId:    4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			{
				Id:        12,
				Name:      "wakwow",
				UserId:    4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
		}

		playlistRepo.On("GetPlaylistByUserId", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnUid, nil).Once()
		_, err := playlistService.CreatePlaylist(context.Background(), "wakwow", 1, "FREE")
		assert.Error(t, err, "You're free user, max can't add 10 playlist")
	})

	t.Run("test case 5, user free but created playlist error! ", func(t *testing.T) {
		expectedReturnUid := []record.Playlist{
			{
				Id:        1,
				Name:      "wakwow",
				UserId:    3,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			{
				Id:        2,
				Name:      "wakwow",
				UserId:    4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			{
				Id:        3,
				Name:      "wakwow",
				UserId:    4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			{
				Id:        4,
				Name:      "wakwow",
				UserId:    4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
		}

		expectedReturnCreating := record.Playlist{}

		playlistRepo.On("GetPlaylistByUserId", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnUid, nil).Once()
		playlistRepo.On("CreatePlaylist", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(expectedReturnCreating, fmt.Errorf("something wrong with the server")).Once()

		_, err := playlistService.CreatePlaylist(context.Background(), "wakwow", 1, "FREE")
		assert.Equal(t, err, fmt.Errorf("something wrong with the server"))
	})

	t.Run("test case 6, success creating playlist premium user! ", func(t *testing.T) {
		expectedReturnCreating := record.Playlist{
			Id:        1,
			Name:      "wakwow",
			UserId:    3,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		}
		playlistRepo.On("CreatePlaylist", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(expectedReturnCreating, nil).Once()

		playlist, err := playlistService.CreatePlaylist(context.Background(), "wakwow", 1, "PREMIUM")

		assert.Equal(t, playlist.Name, expectedReturnCreating.Name)
		assert.Nil(t, err)
	})

	t.Run("test case 7, premium user but created playlist error server! ", func(t *testing.T) {
		expectedReturnCreating := record.Playlist{}
		playlistRepo.On("CreatePlaylist", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(expectedReturnCreating, fmt.Errorf("something wrong with our server")).Once()

		_, err := playlistService.CreatePlaylist(context.Background(), "wakwow", 1, "PREMIUM")
		assert.Equal(t, err, fmt.Errorf("something wrong with our server"))
	})

}

func TestGetAllPlaylist(t *testing.T) {
	playlistService := setup()

	t.Run("test case 1, success get all playlist! ", func(t *testing.T) {
		expectedReturn := []record.Playlist{
			{
				Id:        1,
				Name:      "testing1",
				UserId:    1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
				PlaylistDetail: []record.PlaylistDetail{
					{
						Id:            1,
						PlaylistId:    1,
						YoutubeDataId: "awokawkoaw",
						YoutubeData: record.YoutubeData{
							Id:          1,
							Title:       "Naruto",
							ChannelId:   "12",
							PublishedAt: "testing",
							Description: "asdasdad",
						},
					},
				},
			},
		}

		playlistRepo.On("GetAllPlaylist", mock.Anything).Return(expectedReturn, nil).Once()

		_, err := playlistService.GetAllPlaylist(context.Background())
		assert.Nil(t, err)
	})

	t.Run("test case 2, failed something wrong with db! ", func(t *testing.T) {
		//expectedReturn := []record.Playlist{
		//	{
		//		Id:        1,
		//		Name:      "testing1",
		//		UserId:    1,
		//		CreatedAt: time.Now(),
		//		UpdatedAt: time.Now(),
		//		DeletedAt: gorm.DeletedAt{},
		//		PlaylistDetail: []record.PlaylistDetail{
		//			{
		//				Id:            1,
		//				PlaylistId:    1,
		//				YoutubeDataId: "awokawkoaw",
		//				YoutubeData: record.YoutubeData{
		//					Id:          1,
		//					Title:       "Naruto",
		//					ChannelId:   "12",
		//					PublishedAt: "testing",
		//					Description: "asdasdad",
		//				},
		//			},
		//		},
		//	},
		//}

		playlistRepo.On("GetAllPlaylist", mock.Anything).Return(nil, fmt.Errorf("something wrong with our server")).Once()

		_, err := playlistService.GetAllPlaylist(context.Background())
		assert.Equal(t, err, fmt.Errorf("something wrong with our server"))
	})

}

func TestGetPlaylistByUserId(t *testing.T) {
	playlistService := setup()

	t.Run("test case 1, success get playlist owned by user id! ", func(t *testing.T) {
		expectedReturn := []record.Playlist{
			{
				Id:        1,
				Name:      "testing1",
				UserId:    1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
				PlaylistDetail: []record.PlaylistDetail{
					{
						Id:            1,
						PlaylistId:    1,
						YoutubeDataId: "awokawkoaw",
						YoutubeData: record.YoutubeData{
							Id:          1,
							Title:       "Naruto",
							ChannelId:   "12",
							PublishedAt: "testing",
							Description: "asdasdad",
						},
					},
				},
			},
		}
		playlistRepo.On("GetPlaylistByUserId", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturn, nil).Once()
		_, err := playlistService.GetPlaylistByUserId(context.Background(), 1)
		assert.Nil(t, err)
	})

	t.Run("test case 2, failed get playlist owned by user id ", func(t *testing.T) {
		expectedReturn := []record.Playlist{}
		playlistRepo.On("GetPlaylistByUserId", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturn, fmt.Errorf("something wrong with our server")).Once()
		_, err := playlistService.GetPlaylistByUserId(context.Background(), 1)
		assert.Equal(t, err, fmt.Errorf("something wrong with our server"))
	})
}

func TestDeletePlaylist(t *testing.T) {
	playlistService := setup()
	t.Run("test case 1, success delete playlist", func(t *testing.T) {
		playlistRepo.On("DeletePlaylist", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()
		err := playlistService.DeletePlaylist(context.Background(), 1)
		assert.Nil(t, err)
	})

	t.Run("test case 2, failed delete playlist", func(t *testing.T) {
		playlistRepo.On("DeletePlaylist", mock.Anything, mock.AnythingOfType("int")).Return(fmt.Errorf("something wrong with our server")).Once()
		err := playlistService.DeletePlaylist(context.Background(), 1)
		assert.Equal(t, err, fmt.Errorf("something wrong with our server"))
	})

}

func TestEditPlaylist(t *testing.T) {
	playlistService := setup()
	t.Run("test case 1, success update playlist", func(t *testing.T) {
		domain := Playlist{
			Id:        1,
			Name:      "testing1",
			UserId:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		}
		expectedReturn := record.Playlist{
			Id:        1,
			Name:      "testing1",
			UserId:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		}
		playlistRepo.On("EditPlaylist", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(expectedReturn, nil).Once()

		playlist, err := playlistService.EditPlaylist(context.Background(), domain, 1)

		assert.Nil(t, err)
		assert.Equal(t, playlist.Name, expectedReturn.Name)
	})

	t.Run("test case 2, failed update playlist", func(t *testing.T) {
		domain := Playlist{
			Id:        1,
			Name:      "testing1",
			UserId:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		}
		expectedReturn := record.Playlist{}
		playlistRepo.On("EditPlaylist", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(expectedReturn, fmt.Errorf("something wrong with the server")).Once()
		_ , err := playlistService.EditPlaylist(context.Background(), domain, 1)
		assert.Equal(t, err, fmt.Errorf("something wrong with the server"))

	})
}