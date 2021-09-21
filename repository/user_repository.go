package repository

import (
	"context"
	"playlist-saver/model/domain"
)

type UserRepository interface {
	Register(ctx context.Context, user domain.User) domain.User
	Login(ctx context.Context, email string, password string) (domain.User, error)
}
