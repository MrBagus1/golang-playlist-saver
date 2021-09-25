package repouser

import (
	"context"
	"playlist-saver/model/record"
)

type UserRepository interface {
	Register(ctx context.Context, user record.User) record.User
	Login(ctx context.Context, email string) record.User
}
