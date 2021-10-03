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
		return token, err
	}

	return token, nil
}

func (tri *TokenRepositoryImpl) CheckToken(ctx context.Context, token string) (bool, error) {
	//	simple docs
	//* checking token valid or not.
	//****

	tokens := record.Token{}

	err := tri.client.Conn().Debug().WithContext(ctx).Where("tokens_id = ?", token).First(&tokens).Error
	if err != nil {
		return false, err
	}

	return true, nil
}



