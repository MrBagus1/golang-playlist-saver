package reposearch

import (
	"context"
	"playlist-saver/model/record"
)

type SearchRepository interface {
	GetByUrlId(ctx context.Context, Id string) record.YoutubeData
	SearchYtByParam(ctx context.Context, Search string) []record.YoutubeData
}
