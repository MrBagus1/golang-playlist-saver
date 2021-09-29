package servplaylistdetail

import "context"

type PlaylistDetailService interface {
	AddYoutubeToPlaylist(ctx context.Context, detail PlaylistDetail) (PlaylistDetail , error)
}