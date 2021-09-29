package servplaylist

import "context"

type PlaylistService interface {
	CreatePlaylist(ctx context.Context, name string, id int, role string) (Playlist , error)
}
