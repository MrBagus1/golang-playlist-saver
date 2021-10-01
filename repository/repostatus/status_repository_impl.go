package repostatus

import (
	"context"
	"gopkg.in/guregu/null.v4"
	"playlist-saver/app/config/mysql"
	"playlist-saver/exceptions"
	"playlist-saver/model/record"
	"time"
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

func (st *StatusRepositoryImpl) GetPremiumStatus(ctx context.Context) ([]record.Status, error) {
	var status []record.Status

	err := st.client.Conn().Debug().WithContext(ctx).Where("name = ? AND expired_at < ?", "PREMIUM", time.Now()).Find(&status)
	if err != nil {
		return status, nil
	}
	return status, nil
}

func (st *StatusRepositoryImpl) UpdateStatus(ctx context.Context, userId int) error {
	status := record.Status{}
	err := st.client.Conn().Debug().WithContext(ctx).Model(&status).Where("user_id = ?", userId).Updates(map[string]interface{}{"name": "FREE", "token_id": null.Int{}, "expired_at": null.Time{}}).Error
	if err != nil {
		return err
	}

	return nil
}
