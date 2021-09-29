package web

import (
	"log"
	"playlist-saver/service/servsearch"
)

type YoutubeSearchResponse struct {
	YoutubeLink string `json:"id"`
	Title       string `json:"title"`
	ChannelId   string `json:"channel_id"`
	PublishedAt string `json:"published_at"`
	Description string `json:"description"`
}

func (response *YoutubeSearchResponse) FromDomainService(u []servsearch.YoutubeData) {
	responseData := make([]*YoutubeSearchResponse, 0)
	for _, value := range u {
		response.YoutubeLink = value.YoutubeLink
		response.Title = value.Title
		response.ChannelId = value.ChannelId
		response.PublishedAt = value.PublishedAt
		response.Description = value.Description
		responseData = append(responseData, response)
	}
	log.Print("Len dari Responsse Data sController :", len(responseData))

}
