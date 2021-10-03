package servstatus

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"playlist-saver/model/record"
	statusMock "playlist-saver/repository/repostatus/mocks"
	"testing"
)

var (
	statusRepository statusMock.StatusRepository
)

func setup() StatusService {
	statusService := NewStatusService(&statusRepository)
	return statusService
}

func TestGetAllStatus(t *testing.T) {
	statusService := setup()
	t.Run("test case 1, get all status valid ", func(t *testing.T) {
		expectedReturn := []record.Status{
			{
				Id:      1,
				UserId:  1,
			},
			{
				Id:      2,
				UserId:  2,
			},
		}
		//bjanardana@google.com
		statusRepository.On("GetAllStatus", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := statusService.GetAllStatus(context.Background())
		assert.Nil(t, err)
	})

	t.Run("test case 2, get all status failed ", func(t *testing.T) {
		expectedReturn := []record.Status{}
		//bjanardana@google.com
		statusRepository.On("GetAllStatus", mock.Anything).Return(expectedReturn, fmt.Errorf("something wrong with the server")).Once()
		_, err := statusService.GetAllStatus(context.Background())
		assert.Equal(t, err, fmt.Errorf("something wrong with the server"))
	})

}

func TestCronPremiumChecker(t *testing.T) {
	statusService := setup()

	t.Run("test case 1, cron helper checker valid", func(t *testing.T) {
		expectedReturn := []record.Status{
			{
				Id:      1,
				UserId:  1,
				Name: "PREMIUM",
			},
			{
				Id:      2,
				UserId:  2,
				Name: "FREE",
			},
		}
		statusRepository.On("GetPremiumStatus", mock.Anything).Return(expectedReturn, nil).Once()
		statusRepository.On("UpdateStatus", mock.Anything, mock.AnythingOfType("int")).Return( nil).Twice()

		err := statusService.CronPremiumChecker(context.Background())
		assert.Nil(t, err)
	})

	t.Run("test case 2, get premium status error", func(t *testing.T) {
		expectedReturn := []record.Status{}
		statusRepository.On("GetPremiumStatus", mock.Anything).Return(expectedReturn, fmt.Errorf("something wrong with the server")).Once()

		err := statusService.CronPremiumChecker(context.Background())
		assert.Equal(t, err, fmt.Errorf("something wrong with the server"))
	})

	t.Run("test case 3, update status error", func(t *testing.T) {
		expectedReturn := []record.Status{
			{
				Id:      1,
				UserId:  1,
				Name: "PREMIUM",
			},
			{
				Id:      2,
				UserId:  2,
				Name: "FREE",
			},
		}
		statusRepository.On("GetPremiumStatus", mock.Anything).Return(expectedReturn, nil).Once()
		statusRepository.On("UpdateStatus", mock.Anything, mock.AnythingOfType("int")).Return(fmt.Errorf("something wrong with the server"))

		err := statusService.CronPremiumChecker(context.Background())
		assert.Equal(t, err,fmt.Errorf("something wrong with the server"))
	})

}