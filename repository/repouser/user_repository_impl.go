package repouser

import (
	"context"
	"errors"
	"log"
	"playlist-saver/app/config/mysql"
	"playlist-saver/model/record"
	"time"
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
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) Login(ctx context.Context, email string) (record.User, error) {
	user := record.User{}
	err := repository.client.Conn().Debug().WithContext(ctx).Preload("Status").Where("email = ? ", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repository *userRepositoryImpl) UserFindById(ctx context.Context, id int) (record.User, error) {
	user := record.User{}
	err := repository.client.Conn().Debug().WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err !=nil {
		return user, nil
	}

	return user, nil
}

func (repository *userRepositoryImpl) GetAllUser(ctx context.Context) ([]record.User, error) {
	var user []record.User
	err := repository.client.Conn().Debug().WithContext(ctx).Preload("Status").Find(&user).Error
	if err != nil {
		return user, errors.New("Something Wrong!")
	}
	return user, nil
}

func (repository *userRepositoryImpl) UpdateUser(ctx context.Context, user record.User, id int) error {
	err := repository.client.Conn().Debug().WithContext(ctx).Where("Id = ?", id).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *userRepositoryImpl) UserAddToken(ctx context.Context, id int, token int) error {
	user := record.Status{}
	err := user.TokenId.Scan(token)
	if err != nil {
		return err
	}
	log.Println("test id", id)
	// add expired day
	expired := time.Now().Add(5*time.Minute)
	err = user.ExpiredAt.Scan(expired)
	if err != nil {
		return err
	}
	user.Name = "PREMIUM"

	err = repository.client.Conn().WithContext(ctx).Where("user_id = ? ", id).Updates(&user).Error
	if err != nil {
		return err
	}

	return nil
}




