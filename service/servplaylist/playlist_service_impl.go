package servplaylist

import (
	"context"
	"errors"
	"log"
	"playlist-saver/repository/repoplaylist"
	"strconv"
)

type PlaylistServiceImpl struct {
	PlaylistRepository repoplaylist.PlaylistRepository
}

func NewPlaylistService(PlaylistRepository repoplaylist.PlaylistRepository) PlaylistService {
	return &PlaylistServiceImpl{
		PlaylistRepository: PlaylistRepository,
	}
}

func (pl *PlaylistServiceImpl) CreatePlaylist(ctx context.Context, name string, id int, status string) (Playlist, error) {
	dataPlaylist := Playlist{}

	if len(name) == 0 || len(strconv.Itoa(id)) == 0 {
		return dataPlaylist, errors.New("Don't leave name blank")
	}

	if status == "FREE" {
		listPlaylist, err := pl.PlaylistRepository.GetPlaylistByUserId(ctx, id)
		if err != nil {
			return dataPlaylist, err
		}
		if len(listPlaylist) > 10 {
			return dataPlaylist, errors.New("You're free user, max can't add 10 playlist")
		} else {
			CreatedPlaylist, err := pl.PlaylistRepository.CreatePlaylist(ctx, name, id)
			if err != nil {
				return dataPlaylist, err
			}
			dataPlaylist.FromRecordPlaylist(CreatedPlaylist)
			return dataPlaylist, nil
		}
	}

	if status == "PREMIUM" {
		CreatedPlaylist, err := pl.PlaylistRepository.CreatePlaylist(ctx, name, id)
		if err != nil {
			return dataPlaylist, err
		}
		dataPlaylist.FromRecordPlaylist(CreatedPlaylist)
		return dataPlaylist, nil
	}
	//simple docs, get data by user id, find the status free / premium
	// if free , find playlistbyId, if array <=10, then you can assign new playlkist
	// else return new errors, you are free user, can't add more thatn 10 playlist
	//if user status is PREMIUM, then they can add more thatn 10 playlist
	return dataPlaylist, errors.New("Something wrong with server")
}

func (pl *PlaylistServiceImpl) GetAllPlaylist(ctx context.Context) ([]Playlist, error) {

	dataFinalPlaylist := make([]Playlist, 0)
	dataResult, err := pl.PlaylistRepository.GetAllPlaylist(ctx)
	if err != nil {
		return nil, err
	}
	for _, values := range dataResult {
		dataSemi := Playlist{}
		dataSemi.Id = values.Id
		dataSemi.Name = values.Name
		dataSemi.UserId = values.UserId
		dataSemi.CreatedAt = values.CreatedAt
		dataSemi.UpdatedAt = values.UpdatedAt
		dataSemi.DeletedAt = values.DeletedAt
		//dataFinalPlaylistDetail := make([]PlaylistDetail,0)
		for _, values2 := range values.PlaylistDetail {
			dataSemi2 := PlaylistDetail{}
			dataSemi2.Id = values2.Id
			dataSemi2.PlaylistId = values2.PlaylistId
			dataSemi2.YoutubeDataId = values2.YoutubeDataId
			dataSemi2.YoutubeData.Id = values2.YoutubeData.Id
			dataSemi2.YoutubeData.Title = values2.YoutubeData.Title
			dataSemi2.YoutubeData.ChannelId = values2.YoutubeData.ChannelId
			dataSemi2.YoutubeData.PublishedAt = values2.YoutubeData.PublishedAt
			dataSemi2.YoutubeData.Description = values2.YoutubeData.Description
			log.Println("testest values2", values2.YoutubeData.Title)
			//dataFinalPlaylistDetail = append(dataFinalPlaylistDetail,dataSemi2)
			dataSemi.PlaylistDetail = append(dataSemi.PlaylistDetail, dataSemi2)
		}
		dataFinalPlaylist = append(dataFinalPlaylist, dataSemi)
		log.Println("Data semi", dataSemi)
	}
	log.Println("dataFinalService,", dataFinalPlaylist)

	return dataFinalPlaylist, nil
}

func (pl *PlaylistServiceImpl) GetPlaylistByUserId(ctx context.Context, id int) ([]Playlist, error) {
	dataFinalPlaylist := make([]Playlist, 0)
	dataResult, err := pl.PlaylistRepository.GetPlaylistByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	for _, values := range dataResult {
		dataSemi := Playlist{}
		dataSemi.Id = values.Id
		dataSemi.Name = values.Name
		dataSemi.UserId = values.UserId
		dataSemi.CreatedAt = values.CreatedAt
		dataSemi.UpdatedAt = values.UpdatedAt
		dataSemi.DeletedAt = values.DeletedAt
		//dataFinalPlaylistDetail := make([]PlaylistDetail,0)
		for _, values2 := range values.PlaylistDetail {
			dataSemi2 := PlaylistDetail{}
			dataSemi2.Id = values2.Id
			dataSemi2.PlaylistId = values2.PlaylistId
			dataSemi2.YoutubeDataId = values2.YoutubeDataId
			dataSemi2.YoutubeData.Id = values2.YoutubeData.Id
			dataSemi2.YoutubeData.Title = values2.YoutubeData.Title
			dataSemi2.YoutubeData.ChannelId = values2.YoutubeData.ChannelId
			dataSemi2.YoutubeData.PublishedAt = values2.YoutubeData.PublishedAt
			dataSemi2.YoutubeData.Description = values2.YoutubeData.Description
			log.Println("testest values2", values2.YoutubeData.Title)
			//dataFinalPlaylistDetail = append(dataFinalPlaylistDetail,dataSemi2)
			dataSemi.PlaylistDetail = append(dataSemi.PlaylistDetail, dataSemi2)
		}
		dataFinalPlaylist = append(dataFinalPlaylist, dataSemi)
		log.Println("Data semi", dataSemi)
	}
	log.Println("dataFinalService,", dataFinalPlaylist)

	return dataFinalPlaylist, nil
}

func (pl *PlaylistServiceImpl) DeletePlaylist(ctx context.Context, id int) error {
	err := pl.PlaylistRepository.DeletePlaylist(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
func (pl *PlaylistServiceImpl) EditPlaylist(ctx context.Context, playlist Playlist, id int) (Playlist, error) {
	editPlaylist, err := pl.PlaylistRepository.EditPlaylist(ctx, playlist.ToRecordUser(), id)
	if err != nil {
		return Playlist{}, err
	}
	playlist.FromRecordPlaylist(editPlaylist)

	return playlist,nil
}
