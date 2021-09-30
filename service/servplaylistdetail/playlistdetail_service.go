package servplaylistdetail

import "context"

type PlaylistDetailService interface {
	AddYoutubeToPlaylist(ctx context.Context, detail PlaylistDetail) (PlaylistDetail , error)
	DeleteYoutubeDataFromPlaylist(ctx context.Context, id int) error
}