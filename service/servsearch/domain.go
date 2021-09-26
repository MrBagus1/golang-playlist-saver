package servsearch

import (
	"log"
	"playlist-saver/model/record"
)

type YoutubeData struct {
	Id          string
	Title       string
	ChannelId   string
	PublishedAt string
	Description string
}

func (d *YoutubeData) FromRecordYoutube(data []record.YoutubeData) {
	youtubeDataRes := make([]*YoutubeData, 0)
	for _, item := range data {
		d.Id = item.Id
		d.Title = item.Title
		d.ChannelId = item.ChannelId
		d.PublishedAt = item.PublishedAt
		d.Description = item.Description
		youtubeDataRes = append(youtubeDataRes, d)
	}
	log.Print("LenYoutubeDataRess", len(youtubeDataRes))
}

func (d *YoutubeData) FromRecordYoutubeNotArray(data record.YoutubeData) {
	d.Id = data.Id
	d.Title = data.Title
	d.ChannelId = data.ChannelId
	d.PublishedAt = data.PublishedAt
	d.Description = data.Description
}
