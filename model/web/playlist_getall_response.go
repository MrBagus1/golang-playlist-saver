package web

import (
	"gorm.io/gorm"
	"time"
)

type PlaylistGetResponse struct {
	Id             int                         `json:"id"`
	Name           string                      `json:"name"`
	UserId         int                         `json:"user_id"`
	CreatedAt      time.Time                   `json:"created_at"`
	UpdatedAt      time.Time                   `json:"updated_at"`
	DeletedAt      gorm.DeletedAt                   `json:"deleted_at"`
	PlaylistDetail []PlaylistDetailGetResponse `json:"playlist_detail"`
}

type PlaylistDetailGetResponse struct {
	Id            int         `json:"id"`
	PlaylistId    int         `json:"playlist_id"`
	YoutubeDataId string      `json:"youtube_data_id"`
	YoutubeData   YoutubeData `json:"youtube_data"`
}

type YoutubeData struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	ChannelId   string `json:"channel_id"`
	PublishedAt string `json:"published_at"`
	Description string `json:"description"`
}
