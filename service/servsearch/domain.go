package servsearch

import (
	"playlist-saver/model/record"
)

type YoutubeData struct {
	Id          string
	Title       string
	ChannelId   string
	PublishedAt string
	Description string
}

func (d *YoutubeData) FromRecordYoutubeNotArray(data record.YoutubeData) {
	d.Id = data.Id
	d.Title = data.Title
	d.ChannelId = data.ChannelId
	d.PublishedAt = data.PublishedAt
	d.Description = data.Description
}
