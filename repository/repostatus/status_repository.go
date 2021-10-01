package repostatus

import (
	"context"
	"playlist-saver/model/record"
)

type StatusRepository interface {
	GetStatusByUserId(ctx context.Context, userid int) record.Status
	GetAllStatus(ctx context.Context) ([]record.Status, error)
	GetPremiumStatus(ctx context.Context) ([]record.Status, error)
	UpdateStatus(ctx context.Context,userId int) error
}