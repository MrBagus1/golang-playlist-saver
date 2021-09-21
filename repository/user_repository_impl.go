package repository

import (
	"context"
	"gorm.io/gorm"
	"playlist-saver/exceptions"
	"playlist-saver/model/domain"
)

type userRepositoryImpl struct {
	Conn *gorm.DB
}

func InitMysqlRepository(Conn *gorm.DB) UserRepository {
	return &userRepositoryImpl{Conn}
}

func (repository *userRepositoryImpl) Register(ctx context.Context, user domain.User) domain.User {
	err := repository.Conn.WithContext(ctx).Create(&user).Error
	exceptions.PanicIfError(err)
	return user
}

func (repository *userRepositoryImpl) Login(ctx context.Context, email string, password string) (domain.User, error) {
	repository.Conn.WithContext().Find()
}
