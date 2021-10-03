package servtoken

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"playlist-saver/model/record"
	tokenMock "playlist-saver/repository/repotoken/mocks"
	"playlist-saver/utility"
	"testing"
	"time"
)

var (
	tokenRepository tokenMock.TokenRepository
)

func setup() TokenService {
	tokenService := NewTokenService(&tokenRepository)
	return tokenService
}

func TestAddToken(t *testing.T){
	TokenService := setup()
	t.Run("test case 1, admin add token valid ", func(t *testing.T) {
		tokenRepository.On("PostToken",mock.Anything, mock.Anything).Return(nil).Times(5)
		err := TokenService.AddToken(context.Background())
		assert.Nil(t, err)
	})
	t.Run("test case 2, admin  failed add token valid ", func(t *testing.T) {
		tokenRepository.On("PostToken",mock.Anything, mock.Anything).Return(fmt.Errorf("error adding token")).Once()
		err := TokenService.AddToken(context.Background())
		assert.Equal(t, err, fmt.Errorf("error adding token"))
	})
}

func TestGetToken(t *testing.T){
	TokenService := setup()
	t.Run("test case 1, admin  failed get all token ", func(t *testing.T){
		expectedReturn := make([]record.Token,0)
		domain := make([]Token, 0)
		tokenRepository.On("GetToken",mock.Anything).Return(expectedReturn,fmt.Errorf("Data not found")).Once()
		result, errorsServ := TokenService.GetToken(context.Background())
		assert.Equal(t, errorsServ , fmt.Errorf("Data not found"))
		assert.Equal(t, result, domain)
	})

	t.Run("test case 2, admin get all token ", func(t *testing.T){
		hashedToken := utility.GenerateSecureToken(26)
		expectedReturnValid := []record.Token{
			{
				TokensId: hashedToken,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				TokensId: hashedToken,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		tokenRepository.On("GetToken",mock.Anything).Return(expectedReturnValid,nil).Once()
		_, err := TokenService.GetToken(context.Background())
		assert.Nil(t, err)
	})


}