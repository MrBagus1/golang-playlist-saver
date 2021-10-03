package servplaylistdetail

import (
	"context"
	"playlist-saver/repository/repoplaylistdetail"
	"playlist-saver/repository/reposearch"
)

type PlaylistDetailServiceImpl struct {
	PlaylistDetailRepository repoplaylistdetail.PlaylistDetailRepository
	SearchRepository         reposearch.SearchRepository
}

func NewPlaylistDetail(PlaylistDetailRepository repoplaylistdetail.PlaylistDetailRepository, SearchRepository reposearch.SearchRepository) PlaylistDetailService {
	return &PlaylistDetailServiceImpl{
		PlaylistDetailRepository: PlaylistDetailRepository,
		SearchRepository:         SearchRepository,
	}
}

func (pds *PlaylistDetailServiceImpl) AddYoutubeToPlaylist(ctx context.Context, detail PlaylistDetail) (PlaylistDetail, error) {
	search := pds.SearchRepository.SearchYtById(ctx, detail.YoutubeDataId)
	youtube := "youtube.com/watch?v="+search.YoutubeLink
	detailNew := PlaylistDetail{
		Id: detail.Id,
		PlaylistId: detail.PlaylistId,
		YoutubeDataId: detail.YoutubeDataId,
		YoutubeData: YoutubeData{
			YoutubeLink: youtube,
			Title: search.Title,
			ChannelId: search.ChannelId,
			PublishedAt: search.PublishedAt,
			Description: search.Description,
		},
	}

	addDetail := detailNew.ToRecordPlaylistDetail()
	resultPlaylist := pds.PlaylistDetailRepository.AddYoutubeToPlaylist(ctx, addDetail)
	detail.FromRecordPlaylistDetail(resultPlaylist)

	return detail, nil
}


func (pdr *PlaylistDetailServiceImpl) DeleteYoutubeDataFromPlaylist(ctx context.Context, id int) error {
	err := pdr.PlaylistDetailRepository.DeleteYoutubeDataFromPlaylist(ctx, id)
	if err != nil {
		return err
	}
	return nil
}