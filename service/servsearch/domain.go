package servsearch

import (
	"playlist-saver/model/record"
)

type YoutubeData struct {
	YoutubeLink string
	Title       string
	ChannelId   string
	PublishedAt string
	Description string
}

func (d *YoutubeData) FromRecordYoutubeNotArray(data record.YoutubeData) {
	d.YoutubeLink = data.YoutubeLink
	d.Title = data.Title
	d.ChannelId = data.ChannelId
	d.PublishedAt = data.PublishedAt
	d.Description = data.Description
}
