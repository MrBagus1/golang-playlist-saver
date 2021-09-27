package servuser

import (
	"context"
	"errors"
	"log"
	"playlist-saver/app/middleware"
	"playlist-saver/exceptions"
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

func (service *UserServiceImpl) Register(ctx context.Context, dataUser User) User {
	err := dataUser.Validate()
	exceptions.PanicIfError(err)
	password, err := utility.HashPassword(dataUser.Password)
	exceptions.PanicIfError(err)
	dataUser.Password = password
	dataUser.Status.Name = "FREE"

	userRecord := dataUser.ToRecordUser()
	log.Print("Test Recored", userRecord)
	insertedUser := service.UserRepository.Register(ctx, userRecord)

	dataUser.FromRecordUser(insertedUser)

	return dataUser
}

func (service *UserServiceImpl) Login(ctx context.Context, email, password string) string {
	if len(email) == 0 || len(password) == 0 {
		panic(errors.New("email or password can't be blank"))
	}

	userRecord := service.UserRepository.Login(ctx, email)
	log.Print("Before hash", password)
	log.Print("After hash", userRecord.Password)
	match := utility.CheckPasswordHash(password, userRecord.Password)
	log.Print("Booleana", match)

	if !match {
		panic(errors.New("email or password not match"))
	}

	jwt := service.jwtAuth.GenerateToken(userRecord.Id, userRecord.Name, userRecord.Role)

	return jwt
}
