package servplaylist

import (
	"gorm.io/gorm"
	"playlist-saver/model/record"
	"time"
)

type Playlist struct {
	Id             int
	Name           string
	UserId         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	PlaylistDetail []PlaylistDetail
}

type PlaylistDetail struct {
	Id            int
	PlaylistId    int
	YoutubeDataId string
	YoutubeData   YoutubeData
}

type YoutubeData struct {
	Id          int
	Title       string
	ChannelId   string
	PublishedAt string
	Description string
}

//func (p *Playlist) ArrayPlaylist(detail []record.PlaylistDetail) []PlaylistDetail {
//
//}

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

func (u *Playlist) ToRecordUser() record.Playlist {
	recordPlaylist := record.Playlist{}
	recordPlaylist.Id = u.Id
	recordPlaylist.Name = u.Name
	recordPlaylist.UserId = u.UserId
	recordPlaylist.CreatedAt = u.CreatedAt
	recordPlaylist.UpdatedAt = u.UpdatedAt
	recordPlaylist.DeletedAt = u.DeletedAt

	return recordPlaylist
}
