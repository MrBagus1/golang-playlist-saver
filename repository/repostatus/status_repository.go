package repostatus

import (
	"context"
	"playlist-saver/model/record"
)

type StatusRepository interface {
	GetStatusByUserId(ctx context.Context, userid int) record.Status
}