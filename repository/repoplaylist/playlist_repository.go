package repoplaylist

import (
	"context"
	"playlist-saver/model/record"
)

type PlaylistRepository interface {
	CreatePlaylist(ctx context.Context, name string, id int) (record.Playlist, error)
	GetPlaylistByUserId(ctx context.Context, id int) ([]record.Playlist, error)
	GetAllPlaylist(ctx context.Context)([]record.Playlist, error)
	DeletePlaylist(ctx context.Context,id int) error
	EditPlaylist(ctx context.Context, playlist record.Playlist, id int) (record.Playlist,error)
}
