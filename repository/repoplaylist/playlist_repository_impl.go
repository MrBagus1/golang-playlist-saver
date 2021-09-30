package repoplaylist

import (
	"context"
	"playlist-saver/app/config/mysql"
	"playlist-saver/model/record"
)

type PlayListRepoImpl struct {
	client mysql.Client
}

func NewPlaylistRepository(client mysql.Client) PlaylistRepository {
	return &PlayListRepoImpl{client}
}

func (pl *PlayListRepoImpl) CreatePlaylist(ctx context.Context, name string, id int) (record.Playlist, error) {
	//panic("implement me")
	playlistRecord := record.Playlist{}
	playlistRecord.Name = name
	playlistRecord.UserId = id
	err := pl.client.Conn().Debug().WithContext(ctx).Create(&playlistRecord).Error
	if err != nil{
		return playlistRecord,err
	}

	return playlistRecord,nil
}

func (pl *PlayListRepoImpl) GetPlaylistByUserId(ctx context.Context, id int) ([]record.Playlist,error) {
	//panic("implement me")

	//playlistRecordArr := make([]record.Playlist, 0)
	var playlistRecord []record.Playlist
	//playlistRecord.Name = name
	//playlistRecord.UserId = id
	err := pl.client.Conn().Debug().WithContext(ctx).Where("user_id = ?", id).Preload("PlaylistDetail").Preload("PlaylistDetail.YoutubeData").Find(&playlistRecord).Error
	if err != nil{
		return nil, err
	}

	return playlistRecord, nil
}

func (pl *PlayListRepoImpl) GetAllPlaylist(ctx context.Context)([]record.Playlist, error){
	var playlistRecord []record.Playlist

	err := pl.client.Conn().Debug().WithContext(ctx).Preload("PlaylistDetail").Preload("PlaylistDetail.YoutubeData").Find(&playlistRecord).Error
	if err != nil {
		return nil, err
	}

	return playlistRecord,nil
}

func (pl *PlayListRepoImpl)DeletePlaylist(ctx context.Context,id int) error{
	playlistRecord := record.Playlist{}
	pl.client.Conn().Debug().WithContext(ctx).Where("Id", id).Delete(&playlistRecord)

	return nil
}
func (pl *PlayListRepoImpl) EditPlaylist(ctx context.Context, playlist record.Playlist, id int) (record.Playlist,error) {
	err := pl.client.Conn().Debug().WithContext(ctx).Where("Id = ?", id).Updates(&playlist).Error
	if err != nil {
		return playlist,err
	}
	return playlist,nil
}

