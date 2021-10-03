package servtoken

import (
	"context"
	"playlist-saver/model/record"
	"playlist-saver/repository/repotoken"
	"playlist-saver/utility"
)

type TokenServiceImpl struct{
	TokenRepository repotoken.TokenRepository
}


func NewTokenService(TokenRepository repotoken.TokenRepository) TokenService {
	return &TokenServiceImpl{
		TokenRepository: TokenRepository,
	}
}


func (tsi *TokenServiceImpl) AddToken(ctx context.Context) error {
	//* simple docs
	//* here admin where create 5 premium tokens at a time
	//******
	token := Token{}
	for i:=0; i<5; i++{
		tokens := utility.GenerateSecureToken(26)
		response := record.Token{
			TokensId: tokens,
			CreatedAt: token.CreatedAt,
			UpdatedAt: token.UpdatedAt,
		}
		err := tsi.TokenRepository.PostToken(ctx, response)
		if err != nil {
			return err
		}
	}

	return nil

}

func (tsi *TokenServiceImpl) GetToken(ctx context.Context) ([]Token, error) {
	//* simple docs
	//* here admin can get all more
	//******

	dataFinal := make([]Token, 0)

	TokenResult, err := tsi.TokenRepository.GetToken(ctx)
	if err != nil {
		return dataFinal, err
	}

	for _ , values := range TokenResult{
		response := Token{}
		response.Id = values.TokensId
		response.CreatedAt = values.CreatedAt
		response.UpdatedAt = values.UpdatedAt
		dataFinal = append(dataFinal, response)
	}

	return dataFinal, nil
}
