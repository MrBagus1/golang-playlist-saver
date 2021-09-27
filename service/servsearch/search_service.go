package servsearch

import (
	"context"
)

type SearchService interface {
	GetByUrlId(ctx context.Context, Id string) YoutubeData
	SearchYtByParam(ctx context.Context, Search string) []YoutubeData
}
