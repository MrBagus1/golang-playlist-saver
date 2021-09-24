package service

import (
	"context"
	"playlist-saver/model/web"
)

type UserService interface {
	Register(ctx context.Context, request web.UserRegisterRequest) web.UserRegisterResponse
	Login(ctx context.Context, request web.UserLoginRequest) string
}
