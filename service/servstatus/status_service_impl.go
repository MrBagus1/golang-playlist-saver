package servstatus

import (
	"context"
	"playlist-saver/repository/repostatus"
)

type StatusServiceImpl struct {
	StatusRepository repostatus.StatusRepository
}

func NewStatusService(StatusRepository repostatus.StatusRepository) StatusService {
	return &StatusServiceImpl{
		StatusRepository: StatusRepository,
	}
}

func (ssi *StatusServiceImpl) GetStatusByUserId(ctx context.Context) Status {
	return Status{}
}

func (ssi *StatusServiceImpl) GetAllStatus(ctx context.Context) ([]Status, error) {
	statusRecord := make([]Status, 0)
	status, err := ssi.StatusRepository.GetAllStatus(ctx)
	if err != nil {
		return nil,err
	}
	for _, values := range status {
		statusResponse := Status{}
		statusResponse.Id = values.Id
		statusResponse.Name = values.Name
		statusResponse.TokenId = values.TokenId
		statusResponse.Token.TokensId = values.Token.TokensId
		statusRecord = append(statusRecord, statusResponse)
	}

	return statusRecord,nil
}
