package repostatus

import (
	"context"
	"playlist-saver/app/config/mysql"
	"playlist-saver/exceptions"
	"playlist-saver/model/record"
)

type StatusRepositoryImpl struct {
	client mysql.Client
}

func NewStatusRepository(client mysql.Client) StatusRepository {
	return &StatusRepositoryImpl{client}
}

func (st *StatusRepositoryImpl) GetStatusByUserId(ctx context.Context, userid int) record.Status {
	status := record.Status{}
	err := st.client.Conn().Debug().WithContext(ctx).Where("UserId = ? ", userid).First(&status).Error
	exceptions.PanicIfError(err)

	return status
}

func (st *StatusRepositoryImpl) GetAllStatus(ctx context.Context) ([]record.Status, error) {
	var status []record.Status

	err := st.client.Conn().Debug().WithContext(ctx).Find(&status).Error
	if err != nil {
		return status, err

	}

	return status, nil
}
