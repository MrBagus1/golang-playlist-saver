package web

import (
	"playlist-saver/service/servplaylist"
)

type PlaylistEditRequest struct {
	Name string `json:"name"`
}

func (request *PlaylistEditRequest) ToDomainService() servplaylist.Playlist {
	u := servplaylist.Playlist{}
	u.Name = request.Name
	return u
}