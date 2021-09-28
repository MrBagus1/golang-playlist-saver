package record

import "gorm.io/gorm"

type YoutubeData struct {
	gorm.Model
	Id          string `gorm:"primary_key,not null"`
	Title       string
	ChannelId   string
	PublishedAt string
	Description string
}
