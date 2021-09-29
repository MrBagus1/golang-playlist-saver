package repouser

import (
	"context"
	"playlist-saver/model/record"
)

type UserRepository interface {
	Register(ctx context.Context, user record.User) (record.User, error)
	Login(ctx context.Context, email string) (record.User, error)
	UserFindById(ctx context.Context, id int) record.User
}
