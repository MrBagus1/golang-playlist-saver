package repouser

import (
	"context"
	"playlist-saver/model/record"
)

type UserRepository interface {
	Register(ctx context.Context, user record.User) (record.User, error)
	Login(ctx context.Context, email string) (record.User, error)
	UserFindById(ctx context.Context, id int) (record.User, error)
	GetAllUser(ctx context.Context) ([]record.User, error)
	UpdateUser(ctx context.Context, user record.User, id int) error
	UserAddToken(ctx context.Context, id int, token int) error
}
