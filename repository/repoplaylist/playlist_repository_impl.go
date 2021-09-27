package repoplaylist

import (
	"context"
	"playlist-saver/app/config/mysql"
	"playlist-saver/exceptions"
	"playlist-saver/model/record"
)

type PlayListRepoImpl struct {
	client mysql.Client
}

func NewPlaylistRepository(client mysql.Client) PlaylistRepository {
	return &PlayListRepoImpl{client}
}

func (pl *PlayListRepoImpl) CreatePlaylist(ctx context.Context, name string, id int) record.Playlist {
	//panic("implement me")
	playlistRecord := record.Playlist{}
	playlistRecord.Name = name
	playlistRecord.UserId = id
	err := pl.client.Conn().Debug().WithContext(ctx).Create(&playlistRecord).Error
	exceptions.PanicIfError(err)

	return playlistRecord
}
