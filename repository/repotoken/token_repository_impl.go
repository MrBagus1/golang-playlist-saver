package repotoken

import (
	"context"
	"playlist-saver/app/config/mysql"
	"playlist-saver/model/record"
)

type TokenRepositoryImpl struct {
	client mysql.Client
}

func NewTokenRepository(client mysql.Client) TokenRepository {
	return &TokenRepositoryImpl{client}
}

func (tri *TokenRepositoryImpl) PostToken(ctx context.Context, token record.Token) error {
	//	simple docs
	//* admin post token
	//****
	err := tri.client.Conn().Debug().WithContext(ctx).Create(&token).Error
	if err != nil {
		return err
	}
	return nil
}

func (tri *TokenRepositoryImpl) GetToken(ctx context.Context) ([]record.Token, error) {
	//	simple docs
	//* admin post token
	//****
	var token []record.Token
	err := tri.client.Conn().Debug().WithContext(ctx).Find(&token).Error
	if err != nil {
		return token, nil
	}

	return token, nil
}
