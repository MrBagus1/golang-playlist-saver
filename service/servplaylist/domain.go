package servplaylist

import (
	"gopkg.in/guregu/null.v4"
	"playlist-saver/model/record"
	"time"
)

type Playlist struct {
	Id             int
	Name           string
	UserId         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      null.Time
	PlaylistDetail []PlaylistDetail
}

type PlaylistDetail struct {
	Id            int
	PlaylistId    int
	YoutubeDataId int
	YoutubeData   YoutubeData
}

type YoutubeData struct {
	Id          string
	Title       string
	ChannelId   string
	PublishedAt string
	Description string
}

//
//func (pl *Playlist) ToRecordPlaylist() record.User {
//	recordPlaylist := record.Playlist{}
//	recordPlaylist.Name = u.Name
//
//	return recordUser
//}

func (p *Playlist) FromRecordPlaylist(recordPlaylist record.Playlist) {
	p.Id = recordPlaylist.Id
	p.Name = recordPlaylist.Name
	p.UserId = recordPlaylist.UserId
	p.CreatedAt = recordPlaylist.CreatedAt
	p.UpdatedAt = recordPlaylist.UpdatedAt
	p.DeletedAt = recordPlaylist.DeletedAt
}
