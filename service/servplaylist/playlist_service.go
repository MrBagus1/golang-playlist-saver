package servplaylist

import "context"

type PlaylistService interface {
	CreatePlaylist(ctx context.Context, name string, id int, role string) (Playlist , error)
	GetAllPlaylist(ctx context.Context)([]Playlist, error)
	GetPlaylistByUserId(ctx context.Context, id int)([]Playlist,error)
	DeletePlaylist(ctx context.Context, id int) error
	EditPlaylist(ctx context.Context, playlist Playlist, id int) (Playlist,error)
}
