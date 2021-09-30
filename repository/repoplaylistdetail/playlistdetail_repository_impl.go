package repoplaylistdetail

import (
	"context"
	"gorm.io/gorm/clause"
	"playlist-saver/app/config/mysql"
	"playlist-saver/exceptions"
	"playlist-saver/model/record"
)

type PlaylistDetailRepositoryImpl struct{
	client mysql.Client
}


func NewPlaylistDetail(client mysql.Client) PlaylistDetailRepository {
	return &PlaylistDetailRepositoryImpl{client}
}
func (pdr *PlaylistDetailRepositoryImpl) AddYoutubeToPlaylist(ctx context.Context, data record.PlaylistDetail) record.PlaylistDetail {
	err := pdr.client.Conn().Debug().WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(&data).Error
	exceptions.PanicIfError(err)
	return data
}

func (pdr *PlaylistDetailRepositoryImpl) DeleteYoutubeDataFromPlaylist(ctx context.Context, id int) error {
	data := record.PlaylistDetail{}
	pdr.client.Conn().Debug().WithContext(ctx).Where("Id", id).Delete(&data)

	return nil
}


