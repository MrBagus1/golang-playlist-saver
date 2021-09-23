package repository

import (
	"context"
	"playlist-saver/app/config/mysql"
	"playlist-saver/exceptions"
	"playlist-saver/model/domain"
)

type userRepositoryImpl struct {
	client mysql.Client
}

func NewUserRepository(client mysql.Client) UserRepository {
	return &userRepositoryImpl{client}
}

func (repository *userRepositoryImpl) Register(ctx context.Context, user domain.User) domain.User {
	user.Status.Name = "testing"
	err := repository.client.Conn().Debug().WithContext(ctx).Create(&user).Error
	exceptions.PanicIfError(err)
	return user
}

func (repository *userRepositoryImpl) Login(ctx context.Context, email string) domain.User {
	user := domain.User{}
	err := repository.client.Conn().Debug().WithContext(ctx).Where("email = ? ", email).First(&user).Error
	exceptions.PanicIfError(err)

	return user
}
