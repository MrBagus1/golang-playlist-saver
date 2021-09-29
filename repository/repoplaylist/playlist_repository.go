package repoplaylist

import (
	"context"
	"playlist-saver/model/record"
)

type PlaylistRepository interface {
	CreatePlaylist(ctx context.Context, name string, id int) (record.Playlist, error)
	GetPlaylistByUserId(ctx context.Context, id int) ([]record.Playlist, error)

}
