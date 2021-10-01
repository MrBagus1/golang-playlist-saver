package repotoken

import (
	"context"
	"playlist-saver/model/record"
)

type TokenRepository interface {
	PostToken(ctx context.Context, token record.Token) error
	GetToken(ctx context.Context) ([]record.Token, error)
	CheckToken(ctx context.Context, token string) (bool,error)
}
