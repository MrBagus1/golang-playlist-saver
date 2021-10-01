package servuser

import (
	"context"
)

type UserService interface {
	Register(ctx context.Context, dataUser User) (User, error)
	Login(ctx context.Context, email, password string) (string, error)
	GetAllUser(ctx context.Context, admin string) ([]User, error)
	UpdateUser(ctx context.Context, user User, id int) error
	GetUserById(ctx context.Context, id int) (User, error)
	UserAddToken(ctx context.Context, id int, token int, tokenNumber string) error
}
