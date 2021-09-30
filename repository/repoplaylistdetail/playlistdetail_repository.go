package repoplaylistdetail

import (
	"context"
	"playlist-saver/model/record"
)

type PlaylistDetailRepository interface {
	AddYoutubeToPlaylist(ctx context.Context, data record.PlaylistDetail) record.PlaylistDetail
	DeleteYoutubeDataFromPlaylist(ctx context.Context, id int) error
}