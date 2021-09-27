package repoplaylist

import (
	"context"
	"playlist-saver/model/record"
)

type PlaylistRepository interface {
	CreatePlaylist(ctx context.Context, name string, id int) record.Playlist
}
