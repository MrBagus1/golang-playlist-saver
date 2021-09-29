package repouser

import (
	"context"
	"playlist-saver/app/config/mysql"
	"playlist-saver/exceptions"
	"playlist-saver/model/record"
)

type userRepositoryImpl struct {
	client mysql.Client
}

func NewUserRepository(client mysql.Client) UserRepository {
	return &userRepositoryImpl{client}
}

func (repository *userRepositoryImpl) Register(ctx context.Context, user record.User) (record.User, error) {
	err := repository.client.Conn().Debug().WithContext(ctx).Create(&user).Error
	if err != nil {
		return user,err
	}

	return user,nil
}

func (repository *userRepositoryImpl) Login(ctx context.Context, email string) (record.User, error) {
	user := record.User{}
	err := repository.client.Conn().Debug().WithContext(ctx).Preload("Status").Where("email = ? ", email).First(&user).Error
	if err != nil{
		return user, err
	}
	return user,nil
}

func (repository *userRepositoryImpl) UserFindById(ctx context.Context, id int) record.User{
	user := record.User{}
	err := repository.client.Conn().Debug().WithContext(ctx).Where("id = ?", id).First(&user).Error
	exceptions.PanicIfError(err)

	return user
}
