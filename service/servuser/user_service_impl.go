package servuser

import (
	"context"
	"errors"
	"log"
	"playlist-saver/app/middleware"
	"playlist-saver/repository/repouser"
	"playlist-saver/utility"
)

type UserServiceImpl struct {
	UserRepository repouser.UserRepository
	jwtAuth        *middleware.ConfigJWT
}

func NewUserService(UserRepository repouser.UserRepository, jwtAuth *middleware.ConfigJWT) UserService {
	return &UserServiceImpl{
		UserRepository: UserRepository,
		jwtAuth:        jwtAuth,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, dataUser User) (User, error) {
	err := dataUser.Validate()
	if err != nil {
		return dataUser, err
	}
	password, err := utility.HashPassword(dataUser.Password)
	if err != nil{
		return dataUser, err
	}

	dataUser.Password = password
	dataUser.Status.Name = "FREE"

	userRecord := dataUser.ToRecordUser()
	log.Print("Test Recored", userRecord)
	insertedUser, err := service.UserRepository.Register(ctx, userRecord)
	if err != nil {
		return dataUser, err
	}
	dataUser.FromRecordUser(insertedUser)

	return dataUser,nil
}

func (service *UserServiceImpl) Login(ctx context.Context, email, password string) (string, error) {
	if len(email) == 0 || len(password) == 0 {
		return "", errors.New("Dont't leave a blank form.")
	}

	userRecord, err := service.UserRepository.Login(ctx, email)
	if err !=nil{
		return "", err
	}
	match := utility.CheckPasswordHash(password, userRecord.Password)

	if !match {
		return "", errors.New("Password doesn't match")
	}

	jwt := service.jwtAuth.GenerateToken(userRecord.Id, userRecord.Name, userRecord.Role, userRecord.Status.Name)

	return jwt,nil
}
