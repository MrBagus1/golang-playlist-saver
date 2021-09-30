package servtoken

import "context"

type TokenService interface {
	AddToken(ctx context.Context) error
	GetToken(ctx context.Context) ([]Token, error)
}
