package servuser

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"os"
	"playlist-saver/app/middleware"
	"playlist-saver/model/record"
	tokenMock "playlist-saver/repository/repotoken/mocks"
	userMock "playlist-saver/repository/repouser/mocks"
	"playlist-saver/utility"
	"testing"
	"time"
)

var (
	userRepository  userMock.UserRepository
	tokenRepository tokenMock.TokenRepository
	jwtAuth         *middleware.ConfigJWT
	service         UserService
)

func setup() UserService {

	jwtAuth := &middleware.ConfigJWT{
		SecretJWT: "testmock123",
		ExpiredIn: 2,
	}
	userService := NewUserService(&userRepository, jwtAuth, &tokenRepository)

	return userService
}

//Register(ctx context.Context, dataUser User) (User, error)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestRegister(t *testing.T) {
	userService := setup()
	t.Run("test case 1, valid test for register", func(t *testing.T) {
		domain := User{
			Name:     "Agus Janardana",
			Password: "agus123456",
			Email:    "bjanardana@google.com",
			Birthday: time.Now(),
			Gender:   "Laki-Laki",
		}
		expectedReturn := record.User{
			Id:       1,
			Name:     "Agus Janardana",
			Password: "agus123456",
			Email:    "bjanardana@google.com",
			Birthday: time.Now(),
			Gender:   "Laki-Laki",
		}
		userRepository.On("Register", mock.Anything, mock.Anything).Return(expectedReturn, nil).Once()

		actualResult, err := userService.Register(context.Background(), domain)
		assert.Nil(t, err)
		assert.Equal(t, expectedReturn.Name, actualResult.Name)
	})

	t.Run("test case 2, testing validation", func(t *testing.T) {
		domain := User{
			Name:     "Agus Janardana",
			Password: "agus123456",
			Email:    "bjanardana",
			Birthday: time.Now(),
			Gender:   "Laki-Laki",
		}
		expectedReturn := record.User{
			Id:       1,
			Name:     "Agus Janardana",
			Password: "agus123456",
			Email:    "bjanardana@google.com",
			Birthday: time.Now(),
			Gender:   "Laki-Laki",
		}
		userRepository.On("Register", mock.Anything, mock.Anything).Return(expectedReturn, nil).Once()
		actualResult, err := userService.Register(context.Background(), domain)
		assert.NotEqualValues(t, err, expectedReturn.Email, actualResult.Email)
		//assert.Equal(t,err, actualResult.Email)
	})
	t.Run("test case 3 email is bjanardana@google.com then make role admin", func(t *testing.T) {
		domain := User{
			Name:     "Agus Janardana",
			Password: "agus123456",
			Email:    "bjanardana@google.com",
			Birthday: time.Now(),
			Gender:   "Laki-Laki",
		}
		expectedReturn := record.User{
			Id:       1,
			Name:     "Agus Janardana",
			Password: "agus123456",
			Email:    "bjanardana@google.com",
			Birthday: time.Now(),
			Gender:   "Laki-Laki",
		}
		userRepository.On("Register", mock.Anything, mock.Anything).Return(expectedReturn, nil).Once()
		actualResult, err := userService.Register(context.Background(), domain)
		assert.Nil(t, err)
		assert.Equal(t, expectedReturn.Email, actualResult.Email)
		//assert.Equal(t,err, actualResult.Email)
	})
	//$2a$14$a8MQ/OJDgBUhXVgYiLOu5eY57CnzqvIC9zx.i6T9JtYQ6EXgSzc6a
}

func TestLogin(t *testing.T) {
	userService := setup()
	t.Run("test case 1, valid login", func(t *testing.T) {
		hashedPassword, _ := utility.HashPassword("agus123456")
		log.Print("testing password", hashedPassword)
		expectedReturn := record.User{
			Email:    "bjanardana@google.com",
			Password: hashedPassword,
		}
		//bjanardana@google.com
		userRepository.On("Login", mock.Anything, mock.AnythingOfType("string")).Return(expectedReturn, nil).Once()
		_, err := userService.Login(context.Background(), "bjanardana@google.com", "agus123456")
		assert.Nil(t, err)
	})

	t.Run("test case 2, invalid login password", func(t *testing.T) {
		hashedPassword, _ := utility.HashPassword("agus123456523")
		log.Print("testing password", hashedPassword)
		expectedReturn := record.User{
			Email:    "bjanardana@google.com",
			Password: hashedPassword,
		}
		//bjanardana@google.com
		userRepository.On("Login", mock.Anything, mock.AnythingOfType("string")).Return(expectedReturn, nil).Once()
		_, err := userService.Login(context.Background(), "bjanardana@google.com", "agus123456")
		errServiceLogin := errors.New("password doesn't match")
		assert.Equal(t, err, errServiceLogin)
	})

	t.Run("test case 3, blank password / email ", func(t *testing.T) {
		hashedPassword, _ := utility.HashPassword("agus123456523")
		log.Print("testing password", hashedPassword)
		expectedReturn := record.User{
			Email:    "bjanardana@google.com",
			Password: hashedPassword,
		}
		//bjanardana@google.com
		userRepository.On("Login", mock.Anything, mock.AnythingOfType("string")).Return(expectedReturn, nil).Once()
		_, err := userService.Login(context.Background(), "", "")
		errServiceLoginBlank := errors.New("dont't leave a blank form")
		assert.Equal(t, err, errServiceLoginBlank)
	})
}

func TestGetAllUser(t *testing.T) {
	userService := setup()
	t.Run("test case 1, valid test get all user", func(t *testing.T) {
		expectedReturn := []record.User{
			{
				Id:       1,
				Name:     "Testing1",
				Email:    "testingemail@ymail1.com",
				Birthday: time.Now(),
				Gender:   "Male",
				Role:     "ADMIN",
			},
			{
				Id:       2,
				Name:     "Testing2",
				Email:    "testingemail@ymail2.com",
				Birthday: time.Now(),
				Gender:   "Male",
				Role:     "ADMIN",
			},
		}
		//bjanardana@google.com
		userRepository.On("GetAllUser", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := userService.GetAllUser(context.Background(), "ADMIN")
		assert.Nil(t, err)
	})

	t.Run("test case 2, not admin ", func(t *testing.T) {
		expectedReturn := []record.User{
			{
				Id:       1,
				Name:     "Testing1",
				Email:    "testingemail@ymail1.com",
				Birthday: time.Now(),
				Gender:   "Male",
				Role:     "ADMIN",
			},
			{
				Id:       2,
				Name:     "Testing2",
				Email:    "testingemail@ymail2.com",
				Birthday: time.Now(),
				Gender:   "Male",
				Role:     "ADMIN",
			},
		}
		//bjanardana@google.com
		userRepository.On("GetAllUser", mock.Anything).Return(expectedReturn, nil).Once()
		_, err := userService.GetAllUser(context.Background(), "USER")
		errorRepositoryNotAdmin := errors.New("you're not admin, can't fetch all data")
		assert.Equal(t, err, errorRepositoryNotAdmin)
		//assert.Nil(t, err)
	})
}

func TestGetUserById(t *testing.T) {
	userService := setup()
	t.Run("test case 1, valid get user by id ", func(t *testing.T) {
		expectedReturn := record.User{
			Id:       1,
			Name:     "Testing1",
			Email:    "testingemail@ymail1.com",
			Birthday: time.Now(),
			Gender:   "Male",
			Role:     "ADMIN",
		}
		//bjanardana@google.com
		userRepository.On("UserFindById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturn, nil).Once()
		actualResult, err := userService.GetUserById(context.Background(), 1)
		assert.Equal(t, expectedReturn.Name, actualResult.Name)
		assert.Nil(t, err)
	})
}

func TestUpdateUser(t *testing.T) {
	userService := setup()
	t.Run("test case 1, valid update user ", func(t *testing.T) {
		domain := User {
			Id:       1,
			Name:     "Testing1",
			Email:    "testingemail@ymail1.com",
			Birthday: time.Now(),
			Gender:   "Male",
			Role:     "ADMIN",
		}
		expectedReturn := record.User{
			Id:       1,
			Name:     "Testing1",
			Email:    "testingemail@ymail1.com",
			Birthday: time.Now(),
			Gender:   "Male",
			Role:     "ADMIN",
		}
		//bjanardana@google.com
		userRepository.On("UserFindById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturn, nil).Once()
		userRepository.On("UpdateUser", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := userService.UpdateUser(context.Background(), domain, domain.Id)
		assert.Nil(t, err)
	})
}

func TestUserAddToken(t *testing.T) {
	userService := setup()
	t.Run("test case 1, user add token success ", func(t *testing.T) {
		//bjanardana@google.com
		tokenNumber := "testingToken1234123"
		tokenRepository.On("CheckToken",mock.Anything, tokenNumber).Return(true,nil)
		userRepository.On("UserAddToken", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(nil).Once()
		err := userService.UserAddToken(context.Background(), 1, 1, tokenNumber)
		assert.Nil(t, err)
	})


}