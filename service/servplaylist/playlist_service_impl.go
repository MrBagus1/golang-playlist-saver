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

func (pl *PlaylistServiceImpl) CreatePlaylist(ctx context.Context, name string, id int) Playlist {
	dataPlaylist := Playlist{}
	if len(name) == 0 || len(strconv.Itoa(id)) == 0 {
		panic(errors.New("Name or Id can't be blank"))
	}
	CreatedPlaylist := pl.PlaylistRepository.CreatePlaylist(ctx, name, id)
	dataPlaylist.FromRecordPlaylist(CreatedPlaylist)
	return dataPlaylist
}
