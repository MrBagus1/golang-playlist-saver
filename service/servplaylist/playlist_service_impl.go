package servplaylist

import (
	"context"
	"errors"
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
		listPlaylist, err := pl.PlaylistRepository.GetPlaylistByUserId(ctx,id)
		if err != nil {
			return dataPlaylist, err
		}
		if len(listPlaylist) > 10 {
			return dataPlaylist, errors.New("You're free user, max can't add 10 playlist")
		} else {
			CreatedPlaylist, err := pl.PlaylistRepository.CreatePlaylist(ctx, name, id)
			if err != nil{
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
	return dataPlaylist , errors.New("Something wrong with server")
}
