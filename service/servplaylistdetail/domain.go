package servplaylistdetail

import (
	"log"
	"playlist-saver/model/record"
)

type PlaylistDetail struct {
	Id            int
	PlaylistId    int
	YoutubeDataId string
	YoutubeData   YoutubeData
}

type YoutubeData struct {
	YoutubeLink string
	Title       string
	ChannelId   string
	PublishedAt string
	Description string
}

func (pd *PlaylistDetail) ToRecordPlaylistDetail() record.PlaylistDetail {
	recordUser := record.PlaylistDetail{}
	log.Print("pd youtube Id", pd.YoutubeDataId)
	recordUser.Id = pd.Id
	recordUser.PlaylistId = pd.PlaylistId
	recordUser.YoutubeDataId = pd.YoutubeDataId
	recordUser.YoutubeData.YoutubeLink = pd.YoutubeData.YoutubeLink
	recordUser.YoutubeData.Title = pd.YoutubeData.Title
	recordUser.YoutubeData.ChannelId = pd.YoutubeData.ChannelId
	recordUser.YoutubeData.PublishedAt = pd.YoutubeData.PublishedAt
	recordUser.YoutubeData.Description = pd.YoutubeData.Description

	return recordUser
}

func (pd *PlaylistDetail) FromRecordPlaylistDetail(record record.PlaylistDetail) {
	pd.Id = record.Id
	pd.PlaylistId = record.PlaylistId
	pd.YoutubeDataId = record.YoutubeDataId
	pd.YoutubeData.YoutubeLink = record.YoutubeData.YoutubeLink
	pd.YoutubeData.Title = record.YoutubeData.Title
	pd.YoutubeData.ChannelId = record.YoutubeData.ChannelId
	pd.YoutubeData.PublishedAt = record.YoutubeData.PublishedAt
	pd.YoutubeData.Description = record.YoutubeData.Description
}
